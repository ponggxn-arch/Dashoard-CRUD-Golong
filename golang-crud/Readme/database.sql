-- Create database
CREATE DATABASE IF NOT EXISTS golang_crud;

-- Use the database
USE golang_crud;

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(20),
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create admins table
CREATE TABLE IF NOT EXISTS admins (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

-- Insert default admin credentials
-- Username: admin, Password: admin123
INSERT INTO admins (username, password) VALUES ('admin', 'admin123')
ON DUPLICATE KEY UPDATE username='admin';

-- Insert sample data (optional)
INSERT INTO users (name, email, phone, address) VALUES
('John Doe', 'john@example.com', '0812345678', '123 Main St, Bangkok'),
('Jane Smith', 'jane@example.com', '0823456789', '456 Oak Ave, Chiang Mai'),
('Bob Wilson', 'bob@example.com', '0834567890', '789 Pine Rd, Phuket');
