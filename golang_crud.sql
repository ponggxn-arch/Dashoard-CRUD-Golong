-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Oct 15, 2025 at 05:07 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.0.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang_crud`
--

-- --------------------------------------------------------

--
-- Table structure for table `admins`
--

CREATE TABLE `admins` (
  `id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `admins`
--

INSERT INTO `admins` (`id`, `username`, `password`) VALUES
(1, 'admin', 'admin1234');

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `name`, `description`, `created_at`) VALUES
(1, 'เครื่องดื่ม', '', '2025-10-14 06:42:15'),
(2, 'อาหารเเห้ง', '', '2025-10-14 06:44:49'),
(3, 'ของใช้ในบ้าน', '', '2025-10-14 06:48:18'),
(4, 'ของใช้ส่วนตัว', '', '2025-10-14 06:56:21'),
(5, 'เครื่องเขียน', '', '2025-10-15 14:36:14');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `code` varchar(50) NOT NULL,
  `name` varchar(200) NOT NULL,
  `category_id` int(11) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `unit` varchar(50) DEFAULT 'ชิ้น',
  `price` decimal(10,2) DEFAULT 0.00,
  `cost` decimal(10,2) DEFAULT 0.00,
  `stock_quantity` int(11) DEFAULT 0,
  `min_stock` int(11) DEFAULT 10,
  `max_stock` int(11) DEFAULT 1000,
  `location` varchar(100) DEFAULT NULL,
  `barcode` varchar(100) DEFAULT NULL,
  `status` enum('active','inactive','discontinued') DEFAULT 'active',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `code`, `name`, `category_id`, `description`, `unit`, `price`, `cost`, `stock_quantity`, `min_stock`, `max_stock`, `location`, `barcode`, `status`, `created_at`, `updated_at`) VALUES
(1, '5818', 'Coca-Cola 325 มล.', 1, 'น้ำอัดลมรสโคล่า ขนาด 325 มล.', 'ชิ้น', 30.00, 20.00, 20, 50, 1000, 'A-01, ชั้น 2', '8850121234567', 'active', '2025-10-14 06:43:08', '2025-10-15 14:48:42'),
(2, '5819', 'น้ำดื่มคริสตัล 600 มล.', 1, 'น้ำดื่มบรรจุขวด ขนาด 600 มิลลิลิตร', 'ชิ้น', 10.00, 6.00, 10, 20, 1000, 'B-02, ชั้น 1', '8851959134562', 'active', '2025-10-14 06:44:19', '2025-10-15 14:58:08'),
(3, '5820', 'มาม่ารสต้มยำกุ้ง 55 ก.', 2, 'มาม่า บะหมี่กึ่งสำเร็จรูป รสต้มยำกุ้ง 55 ก. แพ็ค 10', 'แพ็ค', 62.00, 55.00, 5, 50, 2000, 'C-01, ชั้น 2', '8850987123456', 'active', '2025-10-14 06:47:16', '2025-10-15 14:47:55'),
(5, '5821', 'น้ำยาล้างจานซันไลต์ 500 มล.', 3, 'น้ำยาล้างจาน สูตรมะนาว เข้มข้น 500 มิลลิลิตร', 'ชิ้น', 35.00, 25.00, 150, 10, 500, 'D-03, ชั้น 1', '8851932298765', 'active', '2025-10-14 06:49:30', '2025-10-14 06:49:30'),
(7, '5823', 'สบู่ลักส์ สีชมพู 110 กรัม', 4, 'สบู่ก้อนกลิ่นหอมกุหลาบ 110 กรัม', 'ชิ้น', 20.00, 12.00, 0, 15, 1000, 'F-02, ชั้น 2', '8850912133333', 'inactive', '2025-10-14 06:57:27', '2025-10-15 14:39:15'),
(8, '5896', 'กาแฟ 3-in-1 สูตรเข้มข้น (10 ซอง/กล่อง)', 1, 'กาแฟปรุงสำเร็จชนิดผง รสชาติเข้มข้น หวานน้อย', 'กล่อง', 95.00, 70.00, 170, 40, 1000, 'DRINK-01-A', '	8859998887776', 'active', '2025-10-15 14:34:55', '2025-10-15 14:38:32'),
(9, '5426', 'ยาสีฟันสมุนไพร สูตรดั้งเดิม 100 กรัม', 4, 'ลดกลิ่นปาก บำรุงเหงือกและฟัน', 'ชิ้น', 89.00, 55.00, 400, 80, 1200, 'TOIL-G-01', '8851122334455', 'active', '2025-10-15 14:41:18', '2025-10-15 14:41:29');

-- --------------------------------------------------------

--
-- Table structure for table `stock_movements`
--

CREATE TABLE `stock_movements` (
  `id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `movement_type` enum('in','out','adjust') NOT NULL,
  `quantity` int(11) NOT NULL,
  `previous_stock` int(11) NOT NULL,
  `new_stock` int(11) NOT NULL,
  `reference_no` varchar(100) DEFAULT NULL,
  `source_destination` varchar(200) DEFAULT NULL,
  `reason` text DEFAULT NULL,
  `created_by` varchar(50) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `stock_movements`
--

INSERT INTO `stock_movements` (`id`, `product_id`, `movement_type`, `quantity`, `previous_stock`, `new_stock`, `reference_no`, `source_destination`, `reason`, `created_by`, `created_at`) VALUES
(1, 3, 'in', 150, 500, 650, 'PO-2025-002', ' บริษัท ไทยเพรซิเดนท์ฟูดส์ จำกัด (มหาชน)', '', 'admin', '2025-10-14 06:50:25'),
(2, 1, 'in', 200, 100, 300, 'PO-2025-001', 'บริษัท โคคา-โคล่า (ประเทศไทย) จำกัด', 'รับสินค้าเข้าคลังรอบต้นเดือนตุลาคม 2025', 'admin', '2025-10-14 06:51:26'),
(3, 1, 'out', 50, 300, 250, 'REQ-2025-015', 'แผนกจัดเลี้ยง', 'เบิกไปใช้สำหรับงานสัมมนาภายในบริษัท', 'admin', '2025-10-14 06:52:16'),
(4, 3, 'out', 630, 650, 20, 'REQ-2025-080', 'ร้านขายของชำ', '', 'admin', '2025-10-14 06:54:55'),
(5, 7, 'out', 250, 250, 0, 'PO-2025-788', 'ร้านขายของชำ', 'รอเบิกใหม่สันที่ 14/10/2568', 'admin', '2025-10-14 06:58:49'),
(6, 8, 'out', 80, 250, 170, 'PO-2025-789', 'ร้านขายของชำ', '', 'admin', '2025-10-15 14:38:32'),
(7, 3, 'out', 15, 20, 5, 'REQ-2025-755', 'ร้านขายของชำ', '', 'admin', '2025-10-15 14:47:55'),
(8, 1, 'out', 230, 250, 20, 'REQ-2025-771', 'ร้านขายของชำ', '', 'admin', '2025-10-15 14:48:22'),
(9, 2, 'out', 190, 200, 10, 'REQ-2025-788', 'ร้านขายของชำ', '', 'admin', '2025-10-15 14:58:08');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `address` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admins`
--
ALTER TABLE `admins`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`),
  ADD KEY `category_id` (`category_id`);

--
-- Indexes for table `stock_movements`
--
ALTER TABLE `stock_movements`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_id` (`product_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admins`
--
ALTER TABLE `admins`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `stock_movements`
--
ALTER TABLE `stock_movements`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE SET NULL;

--
-- Constraints for table `stock_movements`
--
ALTER TABLE `stock_movements`
  ADD CONSTRAINT `stock_movements_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
