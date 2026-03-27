# ระบบจัดการข้อมูล Golang CRUD

## 📋 ภาพรวมของระบบ

ระบบจัดการข้อมูลที่พัฒนาด้วยภาษา Golang พร้อมคุณสมบัติครบถ้วนตามที่ร้องขอ

## ✨ คุณสมบัติทั้งหมด (ตามที่ร้องขอ)

### ✅ มีหน้า Login
- หน้า Login ที่สวยงามพร้อม Gradient สีม่วง
- ระบบ Session สำหรับจดจำการเข้าสู่ระบบ
- ป้องกันการเข้าถึงโดยไม่ได้รับอนุญาต
- **Username เริ่มต้น:** admin
- **Password เริ่มต้น:** admin123

### ✅ แสดงข้อมูลได้
- ตารางแสดงข้อมูลผู้ใช้ทั้งหมด
- รองรับการแสดงผลบนมือถือ (Responsive)
- แสดงข้อมูล: ID, ชื่อ, อีเมล, เบอร์โทร, ที่อยู่, วันที่สร้าง
- มีปุ่มแก้ไขและลบในแต่ละแถว

### ✅ เพิ่มข้อมูล (มีป้องกัน SQL Injection)
- ฟอร์มเพิ่มข้อมูลผู้ใช้ใหม่
- ตรวจสอบข้อมูลก่อนบันทึก
- **ใช้ Prepared Statements ป้องกัน SQL Injection**
- แสดงการแจ้งเตือนเมื่อบันทึกสำเร็จด้วย SweetAlert

```go
// ตัวอย่างการป้องกัน SQL Injection
stmt, _ := db.Prepare("INSERT INTO users (name, email, phone, address) VALUES (?, ?, ?, ?)")
stmt.Exec(name, email, phone, address)
```

### ✅ แก้ไขข้อมูล (มีป้องกัน SQL Injection)
- ฟอร์มแก้ไขพร้อมข้อมูลเดิมที่กรอกไว้แล้ว
- **ใช้ Prepared Statements ป้องกัน SQL Injection**
- แสดงการแจ้งเตือนเมื่ออัพเดทสำเร็จ

```go
// ตัวอย่างการป้องกัน SQL Injection
stmt, _ := db.Prepare("UPDATE users SET name = ?, email = ?, phone = ?, address = ? WHERE id = ?")
stmt.Exec(name, email, phone, address, id)
```

### ✅ ลบข้อมูล (มีป้องกัน SQL Injection)
- ถามยืนยันก่อนลบด้วย SweetAlert
- ลบด้วย AJAX ไม่ต้อง Reload หน้า
- **ใช้ Prepared Statements ป้องกัน SQL Injection**
- แสดงการแจ้งเตือนเมื่อลบสำเร็จ

```go
// ตัวอย่างการป้องกัน SQL Injection
stmt, _ := db.Prepare("DELETE FROM users WHERE id = ?")
stmt.Exec(id)
```

## 🛠️ เทคโนโลยีที่ใช้ (Tech Stack)

### Backend
- **Golang** - ภาษาหลักในการพัฒนา
- **MySQL** - ฐานข้อมูล
- **go-sql-driver/mysql** - Database driver
- **gorilla/sessions** - Session management

### Frontend
- **Bootstrap 5** - CSS Framework
- **SweetAlert2** - การแจ้งเตือนที่สวยงาม
- **Bootstrap Icons** - ไอคอน

## 📦 โครงสร้างโปรเจค

```
golang-crud/
├── main.go                 # ไฟล์หลักของโปรแกรม (310 บรรทัด)
├── go.mod                  # การจัดการ Dependencies
├── database.sql            # Script สำหรับสร้างฐานข้อมูล
├── README.md              # คู่มือภาษาอังกฤษ
├── README_TH.md           # คู่มือภาษาไทย (ไฟล์นี้)
├── SETUP.md               # คู่มือติดตั้ง
├── USAGE_GUIDE.md         # คู่มือการใช้งาน
├── FEATURES.md            # รายละเอียดคุณสมบัติ
├── start.ps1              # Script เริ่มต้นอัตโนมัติ (Windows)
└── templates/
    ├── layout.html        # Template หลัก
    ├── login.html         # หน้า Login
    ├── list.html          # หน้าแสดงรายการ
    └── form.html          # ฟอร์มเพิ่ม/แก้ไข
```

## 🚀 วิธีติดตั้งและใช้งาน

### ขั้นตอนที่ 1: ติดตั้ง Dependencies
```bash
go mod download
go mod tidy
```

### ขั้นตอนที่ 2: สร้างฐานข้อมูล MySQL
```sql
CREATE DATABASE golang_crud;
```

หรือใช้ไฟล์ SQL ที่เตรียมไว้:
```bash
mysql -u root -p < database.sql
```

### ขั้นตอนที่ 3: ตั้งค่าการเชื่อมต่อฐานข้อมูล
แก้ไขไฟล์ `main.go` บรรทัดที่ 32:
```go
dsn := "root:รหัสผ่าน@tcp(127.0.0.1:3306)/golang_crud?parseTime=true"
```

### ขั้นตอนที่ 4: รันโปรแกรม
```bash
go run main.go
```

### ขั้นตอนที่ 5: เปิดเว็บเบราว์เซอร์
ไปที่: **http://localhost:8080**

เข้าสู่ระบบด้วย:
- **Username:** admin
- **Password:** admin123

## 🔐 การป้องกัน SQL Injection

### ทำไมต้องป้องกัน SQL Injection?

SQL Injection คือการโจมตีที่ผู้ไม่หวังดีแทรกคำสั่ง SQL ผ่านช่องทางรับข้อมูล เช่น:

```sql
-- ผู้โจมตีอาจพยายามใส่ค่านี้ในช่อง Username
admin' OR '1'='1

-- หรือในช่องข้อมูลอื่นๆ
Robert'); DROP TABLE users; --
```

### วิธีป้องกัน: ใช้ Prepared Statements

โปรเจคนี้ใช้ **Prepared Statements** ในทุกคำสั่ง SQL:

```go
// ❌ วิธีที่ไม่ปลอดภัย (อันตราย!)
query := "SELECT * FROM users WHERE id = " + userInput
db.Query(query)

// ✅ วิธีที่ปลอดภัย (แนะนำ)
stmt, err := db.Prepare("SELECT * FROM users WHERE id = ?")
stmt.QueryRow(userInput)
```

### การป้องกันในโปรเจคนี้:

1. **Login** - ใช้ Prepared Statement ✅
2. **เพิ่มข้อมูล** - ใช้ Prepared Statement ✅
3. **แก้ไขข้อมูล** - ใช้ Prepared Statement ✅
4. **ลบข้อมูล** - ใช้ Prepared Statement ✅
5. **ดึงข้อมูล** - ใช้ Prepared Statement ✅

## 📱 หน้าจอต่างๆ ในระบบ

### 1. หน้า Login (`/login`)
- ดีไซน์สวยงามพร้อม Gradient
- กรอก Username และ Password
- มีการแจ้งเตือนหากกรอกผิด
- **ป้องกัน SQL Injection ✅**

### 2. หน้ารายการผู้ใช้ (`/`)
- แสดงตารางข้อมูลผู้ใช้ทั้งหมด
- มีปุ่ม "เพิ่มผู้ใช้ใหม่"
- มีปุ่มแก้ไข (ไอคอนดินสอ) และลบ (ไอคอนถังขยะ)
- **ป้องกัน SQL Injection ✅**

### 3. หน้าเพิ่มข้อมูล (`/add`)
- ฟอร์มสำหรับกรอกข้อมูลผู้ใช้ใหม่
- ช่องที่จำเป็นมีเครื่องหมาย * สีแดง
- มีการ Validate ข้อมูลก่อนส่ง
- **ป้องกัน SQL Injection ✅**

### 4. หน้าแก้ไขข้อมูล (`/edit?id=X`)
- ฟอร์มพร้อมข้อมูลเดิมกรอกไว้แล้ว
- แก้ไขได้ทันที
- มีปุ่มยกเลิก
- **ป้องกัน SQL Injection ✅**

### 5. การลบข้อมูล (AJAX)
- มี Dialog ถามยืนยันก่อนลบ
- ใช้ AJAX ไม่ต้องโหลดหน้าใหม่
- แจ้งเตือนเมื่อลบสำเร็จ
- **ป้องกัน SQL Injection ✅**

## 🎨 UI/UX Features

### Bootstrap 5
- ✅ Responsive (ใช้ได้ทั้งมือถือและคอมพิวเตอร์)
- ✅ Card Design ที่สวยงาม
- ✅ Navigation Bar แบบ Professional
- ✅ ตารางข้อมูลที่อ่านง่าย
- ✅ ปุ่มต่างๆ มี Hover Effect

### SweetAlert2
- ✅ การแจ้งเตือนสำเร็จ (สีเขียว)
- ✅ การแจ้งเตือนข้อผิดพลาด (สีแดง)
- ✅ Dialog ยืนยันการลบ (สีเหลือง)
- ✅ Auto-dismiss หลังจากผ่านไป 2 วินาที
- ✅ Animation ที่สวยงาม

## 📊 ฐานข้อมูล

### ตาราง users
```sql
- id (รหัสอัตโนมัติ)
- name (ชื่อผู้ใช้)
- email (อีเมล)
- phone (เบอร์โทรศัพท์)
- address (ที่อยู่)
- created_at (วันที่สร้าง)
```

### ตาราง admins
```sql
- id (รหัสอัตโนมัติ)
- username (ชื่อผู้ใช้)
- password (รหัสผ่าน)
```

## 🧪 การทดสอบระบบ

### ทดสอบ SQL Injection - Login
ลองใส่ใน Username:
```
admin' OR '1'='1
```
**ผลลัพธ์:** Login ไม่สำเร็จ ✅ (ปลอดภัย)

### ทดสอบ SQL Injection - เพิ่มข้อมูล
ลองใส่ในช่องชื่อ:
```
Robert'); DROP TABLE users; --
```
**ผลลัพธ์:** ข้อมูลถูกเก็บเป็นข้อความธรรมดา, ไม่มี SQL ถูก Execute ✅

### ทดสอบการเข้าถึงโดยไม่ได้ Login
1. Logout จากระบบ
2. ลองเข้า: `http://localhost:8080/add`
**ผลลัพธ์:** ถูก Redirect ไปหน้า Login ✅

## 🔧 การปรับแต่ง

### เปลี่ยน Port
แก้ไขไฟล์ `main.go` บรรทัดสุดท้าย:
```go
log.Fatal(http.ListenAndServe(":8081", nil))
```

### เปลี่ยนรหัสผ่าน Admin
```sql
UPDATE admins SET password = 'รหัสใหม่' WHERE username = 'admin';
```

### เพิ่ม Admin อีกคน
```sql
INSERT INTO admins (username, password) VALUES ('admin2', 'password123');
```

## 📚 เอกสารเพิ่มเติม

- **README.md** - คู่มือภาษาอังกฤษฉบับสมบูรณ์
- **SETUP.md** - คู่มือติดตั้งละเอียด
- **USAGE_GUIDE.md** - คู่มือการใช้งานทีละขั้นตอน
- **FEATURES.md** - รายละเอียดคุณสมบัติทั้งหมด

## ❓ แก้ปัญหา

### ปัญหา: เชื่อมต่อฐานข้อมูลไม่ได้
**วิธีแก้:**
- ตรวจสอบว่า MySQL เปิดอยู่หรือไม่
- ตรวจสอบ Username/Password ในไฟล์ `main.go`
- ตรวจสอบว่าฐานข้อมูล `golang_crud` ถูกสร้างแล้ว

### ปัญหา: Login ไม่ได้
**วิธีแก้:**
- ตรวจสอบว่าพิมพ์ Username: `admin` ถูกต้อง
- ตรวจสอบว่าพิมพ์ Password: `admin123` ถูกต้อง

### ปัญหา: หา Template ไม่เจอ
**วิธีแก้:**
- ตรวจสอบว่าโฟลเดอร์ `templates/` มีอยู่
- ตรวจสอบว่ารันคำสั่งในโฟลเดอร์ที่ถูกต้อง

## ✅ สรุป

ระบบนี้มีคุณสมบัติครบถ้วนตามที่ร้องขอ:

- ✅ หน้า Login ที่สวยงาม
- ✅ แสดงข้อมูลได้อย่างครบถ้วน
- ✅ เพิ่มข้อมูลพร้อมป้องกัน SQL Injection
- ✅ แก้ไขข้อมูลพร้อมป้องกัน SQL Injection
- ✅ ลบข้อมูลพร้อมป้องกัน SQL Injection
- ✅ ใช้ Bootstrap สำหรับ UI
- ✅ ใช้ SweetAlert สำหรับการแจ้งเตือน
- ✅ ใช้ MySQL เป็นฐานข้อมูล

**พร้อมใช้งานทันที!** 🚀

## 📞 ติดต่อ

หากมีคำถามหรือพบปัญหา สามารถเปิด Issue ใน GitHub ได้เลยครับ

---

**พัฒนาด้วย ❤️ โดยใช้ Golang**
