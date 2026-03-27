# 📦 Inventory Management System (ระบบจัดการสินค้าคงคลัง)

ระบบจัดการสินค้าคงคลังแบบครบวงจร พัฒนาด้วย **Golang** และ **MySQL**

## ✨ ฟีเจอร์หลัก

### 🔐 ระบบ Authentication
- เข้าสู่ระบบด้วย Username/Password
- แยกสิทธิ์การใช้งาน (Admin / Staff)
- Session Management

### 📦 จัดการสินค้า (Product Management)
- ✅ เพิ่ม / ลบ / แก้ไข ข้อมูลสินค้า
- ✅ รหัสสินค้า, ชื่อ, ราคา, หมวดหมู่, หน่วยนับ
- ✅ กำหนดสต็อกขั้นต่ำ และสต็อกสูงสุด
- ✅ ระบุตำแหน่งจัดเก็บในคลัง
- ✅ รองรับบาร์โค้ด

### ➕ รับสินค้าเข้า (Stock In)
- บันทึกจำนวนที่นำเข้า
- ระบุแหล่งที่มา/ซัพพลายเออร์
- เลขที่เอกสารอ้างอิง (PO)
- อัปเดตสต็อกอัตโนมัติ

### ➖ เบิกสินค้าออก (Stock Out)
- บันทึกการเบิกใช้หรือจำหน่าย
- ระบุปลายทาง/ผู้เบิก
- ตรวจสอบสต็อกไม่เพียงพอ
- อัปเดตสต็อกอัตโนมัติ

### 🔎 ค้นหาสินค้า
- ค้นหาด้วยชื่อสินค้า
- ค้นหาด้วยรหัสสินค้า
- กรองตามหมวดหมู่
- กรองตามสถานะ

### 📉 แจ้งเตือนสินค้าคงเหลือต่ำ
- แสดงรายการสินค้าที่ต่ำกว่าค่าขั้นต่ำ
- แจ้งเตือนสินค้าหมดสต็อก
- สีสัญลักษณ์ที่ชัดเจน (แดง/เหลือง/เขียว)

### 📊 Dashboard และสถิติ
- ภาพรวมสินค้าทั้งหมด
- สินค้าสต็อกต่ำ
- สินค้าหมดสต็อก
- มูลค่าสต็อกรวม
- ประวัติการเคลื่อนไหวล่าสุด
- เมนูด่วนสำหรับการทำงาน

### 📅 รายงานและประวัติ
- ประวัติการรับเข้า-เบิกออก
- แสดงตามช่วงเวลา
- กรองตามประเภทการเคลื่อนไหว
- ข้อมูลผู้บันทึกและวันเวลา

### 👥 จัดการผู้ใช้งาน
- เพิ่ม / ลบ / แก้ไขผู้ใช้
- ระบบเดิมจาก User Management

## 🚀 การติดตั้งและใช้งาน

### ขั้นตอนที่ 1: เตรียม Database
```bash
# เปิด MySQL และสร้างฐานข้อมูล
mysql -u root -p
CREATE DATABASE golang_crud;
USE golang_crud;
SOURCE inventory_database.sql;
```

### ขั้นตอนที่ 2: ติดตั้ง Dependencies
```bash
go mod init golang-crud
go get github.com/go-sql-driver/mysql
go get github.com/gorilla/sessions
```

### ขั้นตอนที่ 3: แก้ไขการเชื่อมต่อ Database
แก้ไขไฟล์ `main.go` บรรทัด 40:
```go
dsn := "root:your_password@tcp(127.0.0.1:3306)/golang_crud?parseTime=true"
```

### ขั้นตอนที่ 4: รันโปรแกรม
```bash
go run main.go
```

หรือใช้ PowerShell script:
```powershell
.\start.ps1
```

### ขั้นตอนที่ 5: เข้าใช้งานระบบ
เปิดเว็บเบราว์เซอร์ที่: **http://localhost:8080**

**ข้อมูล Login:**
- Username: `admin`
- Password: `admin123`

## 📁 โครงสร้างโปรเจค
```
golang-crud/
├── main.go                      # ไฟล์หลักของโปรแกรม
├── inventory_database.sql       # SQL Schema สำหรับสร้างฐานข้อมูล
├── start.ps1                    # Script สำหรับรันโปรแกรม (Windows)
├── templates/
│   ├── layout.html              # Layout หลักพร้อมเมนู
│   ├── login.html               # หน้า Login
│   ├── dashboard.html           # หน้า Dashboard
│   ├── product_list.html        # รายการสินค้า
│   ├── product_form.html        # ฟอร์มเพิ่ม/แก้ไขสินค้า
│   ├── stock_in.html            # ฟอร์มรับสินค้าเข้า
│   ├── stock_out.html           # ฟอร์มเบิกสินค้าออก
│   ├── stock_movements.html     # ประวัติการเคลื่อนไหว
│   ├── list.html                # รายการ Users
│   └── form.html                # ฟอร์ม User
└── README_INVENTORY.md          # เอกสารนี้
```

## 💾 โครงสร้างฐานข้อมูล

### ตาราง `categories` (หมวดหมู่สินค้า)
- id, name, description, created_at

### ตาราง `products` (ข้อมูลสินค้า)
- id, code, name, category_id
- description, unit, price, cost
- stock_quantity, min_stock, max_stock
- location, barcode, status
- created_at, updated_at

### ตาราง `stock_movements` (การเคลื่อนไหวสต็อก)
- id, product_id, movement_type
- quantity, previous_stock, new_stock
- reference_no, source_destination
- reason, created_by, created_at

### ตาราง `users` (ผู้ใช้งานทั่วไป)
- id, name, email, phone, address, created_at

### ตาราง `admins` (ผู้ดูแลระบบ)
- id, username, password

## 🎯 การใช้งานหลัก

### 1️⃣ เพิ่มสินค้าใหม่
1. คลิกเมนู "สินค้า" → "เพิ่มสินค้า"
2. กรอกข้อมูล: รหัส, ชื่อ, หมวดหมู่, ราคา
3. ระบุสต็อกขั้นต่ำและตำแหน่งจัดเก็บ
4. คลิก "บันทึก"

### 2️⃣ รับสินค้าเข้า
1. คลิกเมนู "จัดการสต็อก" → "รับสินค้าเข้า"
2. เลือกสินค้าที่ต้องการ
3. ระบุจำนวนและแหล่งที่มา
4. ระบบจะอัปเดตสต็อกอัตโนมัติ

### 3️⃣ เบิกสินค้าออก
1. คลิกเมนู "จัดการสต็อก" → "เบิกสินค้าออก"
2. เลือกสินค้าและระบุจำนวน
3. ระบุปลายทางและเหตุผล
4. ระบบตรวจสอบสต็อกและอัปเดต

### 4️⃣ ตรวจสอบสินค้าคงเหลือต่ำ
- ดูได้จาก Dashboard ในส่วน "สินค้าคงเหลือต่ำ"
- สินค้าที่มีสต็อก ≤ ค่าขั้นต่ำ จะแสดงเป็นสีเหลือง/แดง

### 5️⃣ ดูประวัติการเคลื่อนไหว
- คลิก "จัดการสต็อก" → "ประวัติการเคลื่อนไหว"
- สามารถกรองตามประเภท (รับเข้า/เบิกออก)

## 🛠 เทคโนโลยีที่ใช้

- **Backend:** Go (Golang)
- **Database:** MySQL
- **Frontend:** Bootstrap 5, Bootstrap Icons
- **JavaScript:** SweetAlert2
- **Session:** Gorilla Sessions

## 📝 Features Checklist

- ✅ Authentication System
- ✅ Product Management (CRUD)
- ✅ Category Management
- ✅ Stock In (รับสินค้าเข้า)
- ✅ Stock Out (เบิกสินค้าออก)
- ✅ Stock Movement History
- ✅ Low Stock Alert
- ✅ Search & Filter Products
- ✅ Dashboard with Statistics
- ✅ Real-time Stock Tracking
- ✅ User Management
- ✅ Responsive Design
- ✅ Transaction Management

## 🎨 สีสัญลักษณ์

| สี | ความหมาย |
|---|---|
| 🟢 เขียว | สต็อกปกติ |
| 🟡 เหลือง | สต็อกต่ำ (≤ ค่าขั้นต่ำ) |
| 🔴 แดง | หมดสต็อก (= 0) |

## 👨‍💻 ผู้พัฒนา

Developed with ❤️ using Go and MySQL

## 📄 License

MIT License - Feel free to use this project for learning purposes.

---

**หมายเหตุ:** ระบบนี้พัฒนาเพื่อการศึกษาและใช้งานจริงในองค์กรขนาดเล็ก-กลาง
