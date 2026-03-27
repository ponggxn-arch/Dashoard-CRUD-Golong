# Golang CRUD System - Usage Guide

## 🚀 Getting Started

### Step 1: Start the Application
```bash
go run main.go
```

You should see:
```
Database connected successfully!
Server starting on http://localhost:8080
Default login - Username: admin, Password: admin123
```

### Step 2: Open Your Browser
Navigate to: **http://localhost:8080**

---

## 📖 Page-by-Page Guide

### 1️⃣ Login Page (`/login`)

**Features:**
- Beautiful gradient purple background
- Shield lock icon
- Username and password fields
- Default credentials displayed
- Error notification for wrong credentials

**What to do:**
1. Enter username: `admin`
2. Enter password: `admin123`
3. Click "Login" button
4. You'll be redirected to the main page

**Security Features:**
- ✅ SQL injection protected
- ✅ Session-based authentication
- ✅ Auto-redirect if already logged in

---

### 2️⃣ User List Page (`/`)

**Features:**
- Navigation bar with Home, Add User, and Logout links
- Card-based design
- Responsive table showing all users
- Action buttons for each user (Edit/Delete)
- Empty state message if no users exist

**What you'll see:**
- User ID
- Name
- Email
- Phone
- Address
- Created At date
- Edit (yellow pencil icon) and Delete (red trash icon) buttons

**Actions available:**
1. **Add New User** - Click button in header
2. **Edit User** - Click pencil icon
3. **Delete User** - Click trash icon (with confirmation)
4. **Logout** - Click logout in navbar

**Success Messages:**
- Green notification after creating a user
- Green notification after updating a user

---

### 3️⃣ Add User Page (`/add`)

**Features:**
- Clean form with Bootstrap styling
- Icons for each field
- Required fields marked with red asterisk (*)
- Validation before submission
- Cancel button to go back

**Form Fields:**
1. **Name*** - Required, text input
2. **Email*** - Required, email validation
3. **Phone** - Optional, tel input
4. **Address** - Optional, textarea

**What to do:**
1. Fill in Name (required)
2. Fill in Email (required)
3. Optionally add Phone
4. Optionally add Address
5. Click "Create" button
6. You'll see a success notification
7. Redirected to user list

**Validation:**
- ✅ Name and Email are required
- ✅ Email format validated
- ✅ SweetAlert shows errors if validation fails

**Security:**
- ✅ SQL injection protected with prepared statements

---

### 4️⃣ Edit User Page (`/edit?id=X`)

**Features:**
- Same form as Add User
- Pre-filled with existing data
- Shows "Edit User" in header
- Update button instead of Create
- Cancel button to abort

**What to do:**
1. Modify any field you want to change
2. Click "Update" button
3. You'll see a success notification
4. Redirected to user list with updated data

**Security:**
- ✅ SQL injection protected
- ✅ Authentication required
- ✅ Invalid ID redirects safely

---

### 5️⃣ Delete User (AJAX Action)

**Features:**
- Beautiful confirmation dialog with SweetAlert
- Warning icon
- "Are you sure?" message
- Confirm and Cancel buttons
- No page reload (uses AJAX)

**What happens:**
1. Click trash icon on any user
2. SweetAlert confirmation appears
3. Click "Yes, delete it!"
4. AJAX request sent to server
5. Success notification appears
6. Page auto-refreshes to show updated list

**Security:**
- ✅ SQL injection protected
- ✅ Confirmation required (prevents accidental deletion)
- ✅ Server-side validation

---

## 🎯 Common Use Cases

### Use Case 1: Add a New Customer
1. Login to the system
2. Click "Add New User" button
3. Enter customer details:
   - Name: "John Doe"
   - Email: "john@example.com"
   - Phone: "0812345678"
   - Address: "123 Main St, Bangkok"
4. Click "Create"
5. ✅ Customer added!

### Use Case 2: Update Customer Information
1. Find the customer in the list
2. Click the yellow pencil icon
3. Modify the information (e.g., update phone number)
4. Click "Update"
5. ✅ Information updated!

### Use Case 3: Remove Old Records
1. Find the user to delete
2. Click the red trash icon
3. Read the confirmation dialog
4. Click "Yes, delete it!"
5. ✅ User deleted!

### Use Case 4: Test SQL Injection Protection
1. Go to Add User form
2. Try entering malicious SQL in Name field:
   ```sql
   Robert'); DROP TABLE users; --
   ```
3. Submit the form
4. ✅ The input is safely stored as plain text
5. No SQL injection occurs!

---

## 🔐 Security Testing

### Test 1: SQL Injection - Login Form
Try entering in username:
```
admin' OR '1'='1
```
**Expected Result:** Login fails ✅ (Protected by prepared statements)

### Test 2: SQL Injection - Add User
Try entering in Name field:
```
'; DELETE FROM users WHERE '1'='1
```
**Expected Result:** Name is stored as-is, no SQL executed ✅

### Test 3: Unauthorized Access
1. Logout from the system
2. Try to access: `http://localhost:8080/add`
**Expected Result:** Redirected to login page ✅

### Test 4: Invalid Edit
Try accessing: `http://localhost:8080/edit?id=99999`
**Expected Result:** Redirected to home page ✅

---

## 💡 Tips & Tricks

### Tip 1: Quick Data Entry
Use Tab key to navigate between form fields quickly.

### Tip 2: View Database
To view data directly in MySQL:
```sql
USE golang_crud;
SELECT * FROM users;
```

### Tip 3: Reset Admin Password
```sql
UPDATE admins SET password = 'newpassword' WHERE username = 'admin';
```

### Tip 4: Clear All Users
```sql
TRUNCATE TABLE users;
```

### Tip 5: Add More Admins
```sql
INSERT INTO admins (username, password) VALUES ('newadmin', 'password123');
```

---

## ⚙️ Keyboard Shortcuts

| Action | Shortcut |
|--------|----------|
| Focus Name field | Tab (from Email) |
| Submit form | Enter (when in form) |
| Cancel | Esc (in SweetAlert) |
| Confirm delete | Enter (in SweetAlert) |

---

## 🐛 Troubleshooting

### Problem: Can't login
**Solution:** 
- Check if username is exactly `admin`
- Check if password is exactly `admin123`
- Check MySQL connection

### Problem: User not appearing after creation
**Solution:**
- Check browser console for errors
- Verify MySQL connection
- Refresh the page (F5)

### Problem: Delete button not working
**Solution:**
- Check browser console for JavaScript errors
- Ensure JavaScript is enabled
- Try clearing browser cache

### Problem: Validation errors
**Solution:**
- Ensure Name and Email fields are filled
- Use valid email format (name@domain.com)

---

## 📊 Understanding the UI

### Color Codes
- 🔵 **Blue (Primary)** - Main actions, headers, navigation
- 🟡 **Yellow (Warning)** - Edit actions
- 🔴 **Red (Danger)** - Delete actions
- 🟢 **Green (Success)** - Success notifications
- ⚪ **Gray (Secondary)** - Cancel actions

### Icons
- 🔐 **Shield Lock** - Security/Login
- 📊 **Database** - System logo
- ➕ **Plus Circle** - Add new
- ✏️ **Pencil** - Edit
- 🗑️ **Trash** - Delete
- 🏠 **House** - Home
- 🚪 **Box Arrow** - Logout
- 👤 **Person** - User/Name
- 📧 **Envelope** - Email
- 📞 **Telephone** - Phone
- 📍 **Geo** - Address

---

## 🎓 Best Practices

### Do's ✅
- ✅ Always logout when finished
- ✅ Fill required fields before submitting
- ✅ Confirm before deleting
- ✅ Use valid email formats
- ✅ Keep your admin password secure

### Don'ts ❌
- ❌ Don't share admin credentials
- ❌ Don't skip confirmation dialogs
- ❌ Don't enter special characters in phone numbers
- ❌ Don't close browser while submitting forms

---

## 📱 Mobile Usage

The system is fully responsive and works on:
- 📱 Mobile phones (iOS/Android)
- 📱 Tablets
- 💻 Laptops
- 🖥️ Desktop computers

**Mobile Tips:**
- Navigation menu collapses to hamburger icon
- Tables scroll horizontally
- Forms stack vertically for easier input

---

## 🎉 Success Criteria

You're using the system correctly if:
- ✅ You can login successfully
- ✅ You can see the user list
- ✅ You can add new users
- ✅ You can edit existing users
- ✅ You can delete users (with confirmation)
- ✅ You see success notifications
- ✅ SQL injection attempts don't work

---

**Enjoy your Golang CRUD System!** 🚀

For technical details, see: `FEATURES.md`  
For setup help, see: `SETUP.md`  
For general info, see: `README.md`
