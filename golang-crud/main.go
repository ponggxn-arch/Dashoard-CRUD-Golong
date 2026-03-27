package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

// Database connection
var db *sql.DB
var store = sessions.NewCookieStore([]byte("your-secret-key-change-this-in-production"))

// User struct
type User struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	Address   string
	CreatedAt time.Time
}

// Category struct
type Category struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
}

// Product struct
type Product struct {
	ID            int
	Code          string
	Name          string
	CategoryID    int
	CategoryName  string
	Description   string
	Unit          string
	Price         float64
	Cost          float64
	StockQuantity int
	MinStock      int
	MaxStock      int
	Location      string
	Barcode       string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// StockMovement struct
type StockMovement struct {
	ID                int
	ProductID         int
	ProductCode       string
	ProductName       string
	MovementType      string // 'in', 'out', 'adjust'
	Quantity          int
	PreviousStock     int
	NewStock          int
	ReferenceNo       string
	SourceDestination string
	Reason            string
	CreatedBy         string
	CreatedAt         time.Time
}

// Dashboard struct
type Dashboard struct {
	TotalProducts      int
	LowStockProducts   int
	OutOfStockProducts int
	TotalValue         float64
	RecentMovements    []StockMovement
	LowStockItems      []Product
}

// Initialize database
func initDB() {
	var err error
	// Change these credentials to match your MySQL setup
	dsn := "root:@tcp(127.0.0.1:3306)/golang_crud?parseTime=true"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	fmt.Println("Database connected successfully!")

	// Create table if not exists
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL,
		phone VARCHAR(20),
		address TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	// Create admin table
	createAdminTableQuery := `
	CREATE TABLE IF NOT EXISTS admins (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL
	)`
	_, err = db.Exec(createAdminTableQuery)
	if err != nil {
		log.Fatal("Error creating admin table:", err)
	}

	// Create categories table
	createCategories := `
	CREATE TABLE IF NOT EXISTS categories (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL UNIQUE,
		description TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err = db.Exec(createCategories)
	if err != nil {
		log.Fatal("Error creating categories table:", err)
	}

	// Create products table
	createProducts := `
	CREATE TABLE IF NOT EXISTS products (
		id INT AUTO_INCREMENT PRIMARY KEY,
		code VARCHAR(50) NOT NULL UNIQUE,
		name VARCHAR(200) NOT NULL,
		category_id INT,
		description TEXT,
		unit VARCHAR(50) DEFAULT 'ชิ้น',
		price DECIMAL(10,2) DEFAULT 0.00,
		cost DECIMAL(10,2) DEFAULT 0.00,
		stock_quantity INT DEFAULT 0,
		min_stock INT DEFAULT 10,
		max_stock INT DEFAULT 1000,
		location VARCHAR(100),
		barcode VARCHAR(100),
		status ENUM('active', 'inactive', 'discontinued') DEFAULT 'active',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
	)`
	_, err = db.Exec(createProducts)
	if err != nil {
		log.Fatal("Error creating products table:", err)
	}

	// Create stock_movements table
	createMovements := `
	CREATE TABLE IF NOT EXISTS stock_movements (
		id INT AUTO_INCREMENT PRIMARY KEY,
		product_id INT NOT NULL,
		movement_type ENUM('in', 'out', 'adjust') NOT NULL,
		quantity INT NOT NULL,
		previous_stock INT NOT NULL,
		new_stock INT NOT NULL,
		reference_no VARCHAR(100),
		source_destination VARCHAR(200),
		reason TEXT,
		created_by VARCHAR(50),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
	)`
	_, err = db.Exec(createMovements)
	if err != nil {
		log.Fatal("Error creating stock_movements table:", err)
	}

	// Insert default admin (username: admin, password: admin123)
	db.Exec("INSERT IGNORE INTO admins (username, password) VALUES (?, ?)", "admin", "admin123")
}

// Middleware to check if user is logged in
func authRequired(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

// Login page handler
func loginPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil)
}

// Login handler
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Use prepared statement to prevent SQL injection
	var dbPassword string
	err := db.QueryRow("SELECT password FROM admins WHERE username = ?", username).Scan(&dbPassword)

	if err != nil || dbPassword != password {
		http.Redirect(w, r, "/login?error=1", http.StatusSeeOther)
		return
	}

	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = true
	session.Values["username"] = username
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Logout handler
func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// List all users
func listUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, email, phone, address, created_at FROM users ORDER BY id DESC")
	if err != nil {
		log.Println("Error querying users:", err)
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Address, &user.CreatedAt)
		if err != nil {
			log.Println("Error scanning user:", err)
			continue
		}
		users = append(users, user)
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/list.html"))
	tmpl.ExecuteTemplate(w, "layout", users)
}

// Show add form
func addUserForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/form.html"))
	data := map[string]interface{}{
		"Action": "add",
		"User":   User{},
	}
	tmpl.ExecuteTemplate(w, "layout", data)
}

// Create user
func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	address := r.FormValue("address")

	// Use prepared statement to prevent SQL injection
	stmt, err := db.Prepare("INSERT INTO users (name, email, phone, address) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println("Error preparing statement:", err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, email, phone, address)
	if err != nil {
		log.Println("Error executing statement:", err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/?success=created", http.StatusSeeOther)
}

// Show edit form
func editUserForm(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var user User
	// Use prepared statement to prevent SQL injection
	err := db.QueryRow("SELECT id, name, email, phone, address FROM users WHERE id = ?", id).Scan(
		&user.ID, &user.Name, &user.Email, &user.Phone, &user.Address,
	)
	if err != nil {
		log.Println("Error fetching user:", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/form.html"))
	data := map[string]interface{}{
		"Action": "edit",
		"User":   user,
	}
	tmpl.ExecuteTemplate(w, "layout", data)
}

// Update user
func updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	id := r.FormValue("id")
	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	address := r.FormValue("address")

	// Use prepared statement to prevent SQL injection
	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ?, phone = ?, address = ? WHERE id = ?")
	if err != nil {
		log.Println("Error preparing statement:", err)
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, email, phone, address, id)
	if err != nil {
		log.Println("Error executing statement:", err)
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/?success=updated", http.StatusSeeOther)
}

// Delete user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	id := r.FormValue("id")

	// Use prepared statement to prevent SQL injection
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		log.Println("Error preparing statement:", err)
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error executing statement:", err)
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))
}

// ==================== PRODUCT MANAGEMENT ====================

// List all products
func listProducts(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	categoryID := r.URL.Query().Get("category")
	status := r.URL.Query().Get("status")

	query := `
		SELECT p.id, p.code, p.name, p.category_id, COALESCE(c.name, '') as category_name,
		       p.unit, p.price, p.cost, p.stock_quantity, p.min_stock, p.location, p.status
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE 1=1
	`
	args := []interface{}{}

	if search != "" {
		query += " AND (p.name LIKE ? OR p.code LIKE ?)"
		searchParam := "%" + search + "%"
		args = append(args, searchParam, searchParam)
	}

	if categoryID != "" {
		query += " AND p.category_id = ?"
		args = append(args, categoryID)
	}

	if status != "" {
		query += " AND p.status = ?"
		args = append(args, status)
	} else {
		query += " AND p.status = 'active'"
	}

	query += " ORDER BY p.id DESC"

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("Error querying products:", err)
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Code, &p.Name, &p.CategoryID, &p.CategoryName,
			&p.Unit, &p.Price, &p.Cost, &p.StockQuantity, &p.MinStock, &p.Location, &p.Status)
		if err != nil {
			log.Println("Error scanning product:", err)
			continue
		}
		products = append(products, p)
	}

	// Get categories for filter
	categories, _ := getCategories()

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/product_list.html"))
	data := map[string]interface{}{
		"Products":   products,
		"Categories": categories,
		"Search":     search,
		"CategoryID": categoryID,
	}
	tmpl.ExecuteTemplate(w, "layout", data)
}

// Get all categories
func getCategories() ([]Category, error) {
	rows, err := db.Query("SELECT id, name, description FROM categories ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		err := rows.Scan(&c.ID, &c.Name, &c.Description)
		if err != nil {
			continue
		}
		categories = append(categories, c)
	}
	return categories, nil
}

// Show add product form
func addProductForm(w http.ResponseWriter, r *http.Request) {
	categories, _ := getCategories()

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/product_form.html"))
	data := map[string]interface{}{
		"Action":     "add",
		"Product":    Product{},
		"Categories": categories,
	}
	tmpl.ExecuteTemplate(w, "layout", data)
}

// Create product
func createProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/products", http.StatusSeeOther)
		return
	}

	r.ParseForm()

	stmt, err := db.Prepare(`
		INSERT INTO products (code, name, category_id, description, unit, price, cost, 
							  stock_quantity, min_stock, max_stock, location, barcode, status)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Println("Error preparing statement:", err)
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		r.FormValue("code"),
		r.FormValue("name"),
		r.FormValue("category_id"),
		r.FormValue("description"),
		r.FormValue("unit"),
		r.FormValue("price"),
		r.FormValue("cost"),
		r.FormValue("stock_quantity"),
		r.FormValue("min_stock"),
		r.FormValue("max_stock"),
		r.FormValue("location"),
		r.FormValue("barcode"),
		"active",
	)

	if err != nil {
		log.Println("Error executing statement:", err)
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/products?success=created", http.StatusSeeOther)
}

// Show edit product form
func editProductForm(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Redirect(w, r, "/products", http.StatusSeeOther)
		return
	}

	var p Product
	err := db.QueryRow(`
		SELECT p.id, p.code, p.name, p.category_id, p.description, p.unit, p.price, p.cost,
		       p.stock_quantity, p.min_stock, p.max_stock, p.location, p.barcode, p.status
		FROM products p WHERE p.id = ?
	`, id).Scan(
		&p.ID, &p.Code, &p.Name, &p.CategoryID, &p.Description, &p.Unit, &p.Price, &p.Cost,
		&p.StockQuantity, &p.MinStock, &p.MaxStock, &p.Location, &p.Barcode, &p.Status,
	)

	if err != nil {
		log.Println("Error fetching product:", err)
		http.Redirect(w, r, "/products", http.StatusSeeOther)
		return
	}

	categories, _ := getCategories()

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/product_form.html"))
	data := map[string]interface{}{
		"Action":     "edit",
		"Product":    p,
		"Categories": categories,
	}
	tmpl.ExecuteTemplate(w, "layout", data)
}

// Update product
func updateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/products", http.StatusSeeOther)
		return
	}

	r.ParseForm()

	stmt, err := db.Prepare(`
		UPDATE products SET code=?, name=?, category_id=?, description=?, unit=?, price=?, cost=?,
							min_stock=?, max_stock=?, location=?, barcode=?, status=?
		WHERE id=?
	`)
	if err != nil {
		log.Println("Error preparing statement:", err)
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		r.FormValue("code"),
		r.FormValue("name"),
		r.FormValue("category_id"),
		r.FormValue("description"),
		r.FormValue("unit"),
		r.FormValue("price"),
		r.FormValue("cost"),
		r.FormValue("min_stock"),
		r.FormValue("max_stock"),
		r.FormValue("location"),
		r.FormValue("barcode"),
		r.FormValue("status"),
		r.FormValue("id"),
	)

	if err != nil {
		log.Println("Error executing statement:", err)
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/products?success=updated", http.StatusSeeOther)
}

// Delete product
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/products", http.StatusSeeOther)
		return
	}

	id := r.FormValue("id")

	stmt, err := db.Prepare("UPDATE products SET status='inactive' WHERE id=?")
	if err != nil {
		log.Println("Error preparing statement:", err)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success": false}`))
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error executing statement:", err)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success": false}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))
}

// ==================== STOCK MANAGEMENT ====================

// Show stock in form
func stockInForm(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("product_id")

	var product Product
	if productID != "" {
		db.QueryRow(`
			SELECT id, code, name, unit, stock_quantity 
			FROM products WHERE id=? AND status='active'
		`, productID).Scan(&product.ID, &product.Code, &product.Name, &product.Unit, &product.StockQuantity)
	}

	// Get all active products
	rows, _ := db.Query(`
		SELECT id, code, name, unit, stock_quantity 
		FROM products WHERE status='active' ORDER BY name
	`)
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		rows.Scan(&p.ID, &p.Code, &p.Name, &p.Unit, &p.StockQuantity)
		products = append(products, p)
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/stock_in.html"))
	data := map[string]interface{}{
		"Product":  product,
		"Products": products,
	}
	tmpl.ExecuteTemplate(w, "layout", data)
}

// Process stock in
func processStockIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/stock/in", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	session, _ := store.Get(r, "session")
	username := session.Values["username"].(string)

	productID := r.FormValue("product_id")
	quantity := r.FormValue("quantity")
	referenceNo := r.FormValue("reference_no")
	source := r.FormValue("source")
	reason := r.FormValue("reason")

	// Get current stock
	var currentStock int
	err := db.QueryRow("SELECT stock_quantity FROM products WHERE id=?", productID).Scan(&currentStock)
	if err != nil {
		log.Println("Error fetching product:", err)
		http.Redirect(w, r, "/stock/in?error=1", http.StatusSeeOther)
		return
	}

	// Calculate new stock
	var qtyInt int
	fmt.Sscanf(quantity, "%d", &qtyInt)
	newStock := currentStock + qtyInt

	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		http.Redirect(w, r, "/stock/in?error=1", http.StatusSeeOther)
		return
	}

	// Update product stock
	_, err = tx.Exec("UPDATE products SET stock_quantity=? WHERE id=?", newStock, productID)
	if err != nil {
		tx.Rollback()
		log.Println("Error updating stock:", err)
		http.Redirect(w, r, "/stock/in?error=1", http.StatusSeeOther)
		return
	}

	// Insert stock movement
	_, err = tx.Exec(`
		INSERT INTO stock_movements (product_id, movement_type, quantity, previous_stock, 
									 new_stock, reference_no, source_destination, reason, created_by)
		VALUES (?, 'in', ?, ?, ?, ?, ?, ?, ?)
	`, productID, qtyInt, currentStock, newStock, referenceNo, source, reason, username)

	if err != nil {
		tx.Rollback()
		log.Println("Error inserting movement:", err)
		http.Redirect(w, r, "/stock/in?error=1", http.StatusSeeOther)
		return
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		log.Println("Error committing transaction:", err)
		http.Redirect(w, r, "/stock/in?error=1", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/stock/movements?success=stock_in", http.StatusSeeOther)
}

// Show stock out form
func stockOutForm(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("product_id")

	var product Product
	if productID != "" {
		db.QueryRow(`
			SELECT id, code, name, unit, stock_quantity 
			FROM products WHERE id=? AND status='active'
		`, productID).Scan(&product.ID, &product.Code, &product.Name, &product.Unit, &product.StockQuantity)
	}

	// Get all active products
	rows, _ := db.Query(`
		SELECT id, code, name, unit, stock_quantity 
		FROM products WHERE status='active' ORDER BY name
	`)
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		rows.Scan(&p.ID, &p.Code, &p.Name, &p.Unit, &p.StockQuantity)
		products = append(products, p)
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/stock_out.html"))
	data := map[string]interface{}{
		"Product":  product,
		"Products": products,
	}
	tmpl.ExecuteTemplate(w, "layout", data)
}

// Process stock out
func processStockOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/stock/out", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	session, _ := store.Get(r, "session")
	username := session.Values["username"].(string)

	productID := r.FormValue("product_id")
	quantity := r.FormValue("quantity")
	referenceNo := r.FormValue("reference_no")
	destination := r.FormValue("destination")
	reason := r.FormValue("reason")

	// Get current stock
	var currentStock int
	err := db.QueryRow("SELECT stock_quantity FROM products WHERE id=?", productID).Scan(&currentStock)
	if err != nil {
		log.Println("Error fetching product:", err)
		http.Redirect(w, r, "/stock/out?error=1", http.StatusSeeOther)
		return
	}

	// Calculate new stock
	var qtyInt int
	fmt.Sscanf(quantity, "%d", &qtyInt)

	if qtyInt > currentStock {
		http.Redirect(w, r, "/stock/out?error=insufficient", http.StatusSeeOther)
		return
	}

	newStock := currentStock - qtyInt

	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		http.Redirect(w, r, "/stock/out?error=1", http.StatusSeeOther)
		return
	}

	// Update product stock
	_, err = tx.Exec("UPDATE products SET stock_quantity=? WHERE id=?", newStock, productID)
	if err != nil {
		tx.Rollback()
		log.Println("Error updating stock:", err)
		http.Redirect(w, r, "/stock/out?error=1", http.StatusSeeOther)
		return
	}

	// Insert stock movement
	_, err = tx.Exec(`
		INSERT INTO stock_movements (product_id, movement_type, quantity, previous_stock, 
									 new_stock, reference_no, source_destination, reason, created_by)
		VALUES (?, 'out', ?, ?, ?, ?, ?, ?, ?)
	`, productID, qtyInt, currentStock, newStock, referenceNo, destination, reason, username)

	if err != nil {
		tx.Rollback()
		log.Println("Error inserting movement:", err)
		http.Redirect(w, r, "/stock/out?error=1", http.StatusSeeOther)
		return
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		log.Println("Error committing transaction:", err)
		http.Redirect(w, r, "/stock/out?error=1", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/stock/movements?success=stock_out", http.StatusSeeOther)
}

// List stock movements
func listStockMovements(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("product_id")
	movementType := r.URL.Query().Get("type")

	query := `
		SELECT sm.id, sm.product_id, p.code, p.name, sm.movement_type, sm.quantity,
		       sm.previous_stock, sm.new_stock, sm.reference_no, sm.source_destination,
		       sm.reason, sm.created_by, sm.created_at
		FROM stock_movements sm
		JOIN products p ON sm.product_id = p.id
		WHERE 1=1
	`
	args := []interface{}{}

	if productID != "" {
		query += " AND sm.product_id = ?"
		args = append(args, productID)
	}

	if movementType != "" {
		query += " AND sm.movement_type = ?"
		args = append(args, movementType)
	}

	query += " ORDER BY sm.created_at DESC LIMIT 100"

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("Error querying movements:", err)
		http.Error(w, "Error fetching movements", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var movements []StockMovement
	for rows.Next() {
		var m StockMovement
		err := rows.Scan(&m.ID, &m.ProductID, &m.ProductCode, &m.ProductName,
			&m.MovementType, &m.Quantity, &m.PreviousStock, &m.NewStock,
			&m.ReferenceNo, &m.SourceDestination, &m.Reason, &m.CreatedBy, &m.CreatedAt)
		if err != nil {
			log.Println("Error scanning movement:", err)
			continue
		}
		movements = append(movements, m)
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/stock_movements.html"))
	tmpl.ExecuteTemplate(w, "layout", movements)
}

//========== CATEGORY MANAGEMENT ==========

// List all categories
func listCategories(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, description, created_at FROM categories ORDER BY name")
	if err != nil {
		log.Println("Error querying categories:", err)
		http.Error(w, "Error fetching categories", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt)
		if err != nil {
			log.Println("Error scanning category:", err)
			continue
		}
		categories = append(categories, c)
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/category_list.html"))
	tmpl.ExecuteTemplate(w, "layout", categories)
}

// Show add category form
func addCategoryForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/category_form.html"))
	data := map[string]interface{}{
		"Action":   "add",
		"Category": Category{},
	}
	tmpl.ExecuteTemplate(w, "layout", data)
}

// Create category
func createCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
		return
	}

	r.ParseForm()

	stmt, err := db.Prepare("INSERT INTO categories (name, description) VALUES (?, ?)")
	if err != nil {
		log.Println("Error preparing statement:", err)
		http.Error(w, "Error creating category", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.FormValue("name"), r.FormValue("description"))
	if err != nil {
		log.Println("Error executing statement:", err)
		http.Error(w, "Error creating category", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/categories?success=created", http.StatusSeeOther)
}

// Show edit category form
func editCategoryForm(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
		return
	}

	var c Category
	err := db.QueryRow("SELECT id, name, description FROM categories WHERE id = ?", id).Scan(
		&c.ID, &c.Name, &c.Description,
	)

	if err != nil {
		log.Println("Error fetching category:", err)
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/category_form.html"))
	data := map[string]interface{}{
		"Action":   "edit",
		"Category": c,
	}
	tmpl.ExecuteTemplate(w, "layout", data)
}

// Update category
func updateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
		return
	}

	r.ParseForm()

	stmt, err := db.Prepare("UPDATE categories SET name=?, description=? WHERE id=?")
	if err != nil {
		log.Println("Error preparing statement:", err)
		http.Error(w, "Error updating category", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.FormValue("name"), r.FormValue("description"), r.FormValue("id"))
	if err != nil {
		log.Println("Error executing statement:", err)
		http.Error(w, "Error updating category", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/categories?success=updated", http.StatusSeeOther)
}

// Delete category
func deleteCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
		return
	}

	id := r.FormValue("id")

	// Check if category is being used by products
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM products WHERE category_id = ?", id).Scan(&count)
	if err != nil {
		log.Println("Error checking category usage:", err)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success": false, "message": "Error checking category"}`))
		return
	}

	if count > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success": false, "message": "Cannot delete category. It is being used by products."}`))
		return
	}

	stmt, err := db.Prepare("DELETE FROM categories WHERE id=?")
	if err != nil {
		log.Println("Error preparing statement:", err)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success": false}`))
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error executing statement:", err)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success": false}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))
}

//========== DASHBOARD & REPORTS ==========

func dashboard(w http.ResponseWriter, r *http.Request) {
	var dash Dashboard

	// Total products
	db.QueryRow("SELECT COUNT(*) FROM products WHERE status='active'").Scan(&dash.TotalProducts)

	// Low stock products
	db.QueryRow("SELECT COUNT(*) FROM products WHERE status='active' AND stock_quantity <= min_stock AND stock_quantity > 0").Scan(&dash.LowStockProducts)

	// Out of stock
	db.QueryRow("SELECT COUNT(*) FROM products WHERE status='active' AND stock_quantity = 0").Scan(&dash.OutOfStockProducts)

	// Total inventory value
	db.QueryRow("SELECT COALESCE(SUM(stock_quantity * price), 0) FROM products WHERE status='active'").Scan(&dash.TotalValue)

	// Recent movements - only essential fields to save space, limit to 5
	rows, _ := db.Query(`
		SELECT sm.id, p.name, sm.movement_type, sm.quantity, sm.created_at
		FROM stock_movements sm
		JOIN products p ON sm.product_id = p.id
		ORDER BY sm.created_at DESC LIMIT 5
	`)
	defer rows.Close()

	for rows.Next() {
		var m StockMovement
		// Scan only the fields we selected
		rows.Scan(&m.ID, &m.ProductName, &m.MovementType, &m.Quantity, &m.CreatedAt)
		dash.RecentMovements = append(dash.RecentMovements, m)
	}

	// Low stock items
	rows2, _ := db.Query(`
		SELECT p.id, p.code, p.name, COALESCE(c.name, '') as category,
		       p.unit, p.stock_quantity, p.min_stock, p.location
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.status='active' AND p.stock_quantity <= p.min_stock
		ORDER BY p.stock_quantity ASC LIMIT 10
	`)
	defer rows2.Close()

	for rows2.Next() {
		var p Product
		rows2.Scan(&p.ID, &p.Code, &p.Name, &p.CategoryName,
			&p.Unit, &p.StockQuantity, &p.MinStock, &p.Location)
		dash.LowStockItems = append(dash.LowStockItems, p)
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/dashboard.html"))
	tmpl.ExecuteTemplate(w, "layout", dash)
}

func main() {
	// Initialize database
	initDB()
	defer db.Close()

	// Routes - Authentication
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/auth/login", login)
	http.HandleFunc("/logout", logout)

	// Routes - Dashboard
	http.HandleFunc("/", authRequired(dashboard))
	http.HandleFunc("/dashboard", authRequired(dashboard))

	// Routes - User Management
	http.HandleFunc("/users", authRequired(listUsers))
	http.HandleFunc("/users/add", authRequired(addUserForm))
	http.HandleFunc("/users/create", authRequired(createUser))
	http.HandleFunc("/users/edit", authRequired(editUserForm))
	http.HandleFunc("/users/update", authRequired(updateUser))
	http.HandleFunc("/users/delete", authRequired(deleteUser))

	// Routes - Product Management
	http.HandleFunc("/products", authRequired(listProducts))
	http.HandleFunc("/products/add", authRequired(addProductForm))
	http.HandleFunc("/products/create", authRequired(createProduct))
	http.HandleFunc("/products/edit", authRequired(editProductForm))
	http.HandleFunc("/products/update", authRequired(updateProduct))
	http.HandleFunc("/products/delete", authRequired(deleteProduct))

	// Routes - Category Management
	http.HandleFunc("/categories", authRequired(listCategories))
	http.HandleFunc("/categories/add", authRequired(addCategoryForm))
	http.HandleFunc("/categories/create", authRequired(createCategory))
	http.HandleFunc("/categories/edit", authRequired(editCategoryForm))
	http.HandleFunc("/categories/update", authRequired(updateCategory))
	http.HandleFunc("/categories/delete", authRequired(deleteCategory))

	// Routes - Stock Management
	http.HandleFunc("/stock/in", authRequired(stockInForm))
	http.HandleFunc("/stock/in/process", authRequired(processStockIn))
	http.HandleFunc("/stock/out", authRequired(stockOutForm))
	http.HandleFunc("/stock/out/process", authRequired(processStockOut))
	http.HandleFunc("/stock/movements", authRequired(listStockMovements))

	fmt.Println("========================================")
	fmt.Println("📦 Inventory Management System")
	fmt.Println("========================================")
	fmt.Println("🌐 Server: http://localhost:8080")
	// fmt.Println("👤 Username: admin")
	// fmt.Println("🔑 Password: admin1234")
	fmt.Println("========================================")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
