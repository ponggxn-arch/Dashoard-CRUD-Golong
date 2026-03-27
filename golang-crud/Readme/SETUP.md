# Quick Setup Guide - Golang CRUD System

## Step-by-Step Installation

### 1. Install MySQL (if not already installed)
Download from: https://dev.mysql.com/downloads/mysql/

### 2. Create Database
Open MySQL command line or MySQL Workbench and run:
```sql
CREATE DATABASE golang_crud;
```

Or use the provided SQL file:
```bash
mysql -u root -p < database.sql
```

### 3. Install Go Dependencies
```bash
go mod tidy
```

### 4. Configure Database (if needed)
Edit `main.go` line 32 to match your MySQL credentials:
```go
dsn := "username:password@tcp(127.0.0.1:3306)/golang_crud?parseTime=true"
```

Common configurations:
- Local MySQL with password: `root:yourpassword@tcp(127.0.0.1:3306)/golang_crud?parseTime=true`
- Local MySQL without password: `root:@tcp(127.0.0.1:3306)/golang_crud?parseTime=true`
- Remote MySQL: `user:pass@tcp(192.168.1.100:3306)/golang_crud?parseTime=true`

### 5. Run the Application
```bash
go run main.go
```

You should see:
```
Database connected successfully!
Server starting on http://localhost:8080
Default login - Username: admin, Password: admin123
```

### 6. Access the Application
Open your browser and go to: http://localhost:8080

Login with:
- **Username**: admin
- **Password**: admin123

## Common Issues & Solutions

### Issue 1: "Error connecting to database"
**Solution**: 
- Check if MySQL is running
- Verify database credentials in `main.go`
- Ensure `golang_crud` database exists

### Issue 2: "Templates not found"
**Solution**: 
- Make sure you're running the command from the project root directory
- Verify that `templates/` folder exists with all HTML files

### Issue 3: Port 8080 already in use
**Solution**: 
- Change the port in `main.go` (last line):
  ```go
  log.Fatal(http.ListenAndServe(":8081", nil))
  ```

### Issue 4: Go packages not found
**Solution**: 
```bash
go mod download
go mod tidy
```

## Testing the System

1. **Login Test**: Use admin/admin123 credentials
2. **Add User Test**: Click "Add New User" and fill the form
3. **Edit User Test**: Click pencil icon on any user
4. **Delete User Test**: Click trash icon on any user
5. **SQL Injection Test**: Try entering `'; DROP TABLE users; --` in any field (it won't work! ✅)

## Building for Production

To create an executable:
```bash
go build -o crud-app.exe main.go
```

Then run:
```bash
./crud-app.exe
```

## Next Steps

- ✅ Change default admin password in database
- ✅ Add more fields to the user form
- ✅ Implement pagination for large datasets
- ✅ Add search/filter functionality
- ✅ Deploy to production server

---

Enjoy your Golang CRUD System! 🚀
