-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jul 03, 2024 at 05:50 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.1.25

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `agen`
--

-- --------------------------------------------------------

--
-- Table structure for table `admin`
--

CREATE TABLE `admin` (
  `id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `admin`
--

INSERT INTO `admin` (`id`, `username`, `password`) VALUES
(1, 'Admin', '$2a$10$6rBQpJm7tAp2/FyoICnCWeshO3f4oYGBeGgMNhJiBrVK5TG/YU2jq');

-- --------------------------------------------------------

--
-- Table structure for table `cart`
--

CREATE TABLE `cart` (
  `id` int(11) NOT NULL,
  `product_name` varchar(255) NOT NULL,
  `kuantitas` int(11) NOT NULL DEFAULT 1,
  `timestamp` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `pengguna`
--

CREATE TABLE `pengguna` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `Time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `pengguna`
--

INSERT INTO `pengguna` (`id`, `username`, `password`, `Time`) VALUES
(1, '31720566533', '$2a$10$VW6MfZakzwpagsb/4tGtrO50Fry6ipGezkaptpmncBezBzvCblWxi', '2024-06-13 18:21:08'),
(2, 'afni', '$2a$10$6lz1LT6HlAoDBdbdf1CRVunkVV1b3TEozBWHxGLUhLWi/4MkdeehG', '2024-06-13 18:21:08'),
(3, 'salwa', '$2a$10$M2.sFDAsaAmc0krJH53ktOaJ7bgHfGT6pdiXW/PbcJH240xqmyvw.', '2024-06-13 18:21:08'),
(4, 'papa', '$2a$10$WKjddvb64ymlWiYh3/HE7eD07JINfQ.0af4TQux38KLv0TheIx9Ia', '2024-06-13 18:21:08'),
(5, 'kekey', '$2a$10$X1bFKHUF.ny6ww/0OEQ2eeIWFC2Mnefi8rKgu/mZqmEdkYSeExMje', '2024-06-13 18:21:08'),
(19, 'mama', '$2a$10$MlS4O..AFX9vCsPQiptHNuqvkRSyJCQL.LwsDpPSmnf1bo.vsurpa', '2024-06-13 18:21:08'),
(25, 'sxlwxz', '$2a$10$Kc0KoWs4CTKnaM4U8e81C./hbowtfIBOsO1oLI/SVmBRdXZyKXMXK', '2024-06-16 04:40:10'),
(26, 'gumai', '$2a$10$zAULXEn1phNmYlcQ1lDmsuV.fDWJWF97F1jyQoFr4Zf4oX0zsgEwC', '2024-06-18 13:15:10'),
(29, 'koba', '$2a$10$fgZ2urOSl6Jf7g8Yj9HPl.dc74z60CoQtkTlKcjc/mpMcdZ0tkc6.', '2024-06-20 18:33:41'),
(30, 'afnipz', '$2a$10$5kLoWFXGkhaOMv3q6StxhOkDxs4b4duE/SwUfqgaZ6jLxMfFYQ2wi', '2024-06-21 06:25:48'),
(33, 'mir', '$2a$10$bNaWzol11ZktrHXjq6M0yuOnh06ay7i7ohIROAwqfan9Dxtr4M41O', '2024-06-28 20:04:59'),
(34, 'amir', '$2a$10$sC/Nh9ogf5CdH8.dt/NKNurOtVy71KHEPAtauSbzL94eT2IIfCeeS', '2024-06-30 04:35:11'),
(35, 'Gita', '$2a$10$d262rwv5QRFQfLtIYxnIauUrDcEIy.z8fFKFsVMSGTuvfZnHBKIqG', '2024-07-01 14:53:08'),
(38, 'Dilla', '$2a$10$5r2phCqdKcYtdAxUP5ktrOnf9tsSS3e/YpNBO.KOLhFCow4/vSKVq', '2024-07-03 15:40:12');

-- --------------------------------------------------------

--
-- Table structure for table `product`
--

CREATE TABLE `product` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `harga` float DEFAULT NULL,
  `deskripsi` tinytext DEFAULT NULL,
  `tanggal` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `product`
--

INSERT INTO `product` (`id`, `nama`, `harga`, `deskripsi`, `tanggal`) VALUES
(2, 'Djarum', 16000, 'Rokok Mahal', '2024-07-03 15:19:07'),
(5, 'Kopi ABC', 2000, 'Kopi', '2024-06-30 05:29:12'),
(9, 'Indomie', 3000, 'Indomie Indonesia', '2024-06-21 06:27:32'),
(18, 'Spagethie', 200000, ' Spagethi enakkk', '2024-06-21 06:28:27'),
(24, 'Cap Kaki', 13000, '     Drink', '2024-06-28 16:26:41'),
(28, 'Gula Pasir', 9800, '          Untuk ngeteh', '2024-06-29 14:39:49'),
(32, 'Garam', 1000, 'Gudang garam jaya', '2024-07-03 08:21:09'),
(33, 'Cheetos', 2500, 'Cikian', '2024-07-03 08:27:27'),
(34, 'Milkita', 1500, 'Manis bikin ketagihan', '2024-07-03 15:21:15');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` bigint(20) NOT NULL,
  `username` int(11) NOT NULL,
  `product_name` varchar(255) NOT NULL,
  `quantity` int(11) NOT NULL,
  `unit_price` int(11) NOT NULL,
  `total_price` int(11) NOT NULL,
  `total_amount` int(11) NOT NULL,
  `timestamp` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `username`, `product_name`, `quantity`, `unit_price`, `total_price`, `total_amount`, `timestamp`) VALUES
(16, 0, 'Kopi ABC', 5, 2000, 10000, 13000, '2024-06-29 13:25:03'),
(20, 0, 'Kopi', 1, 300, 300, 23600, '2024-06-29 13:25:03'),
(21, 0, 'Cap Kaki', 1, 13000, 13000, 23600, '2024-06-29 13:25:03'),
(22, 0, 'Spagethie', 2, 200000, 400000, 417300, '2024-06-29 13:25:03'),
(23, 0, 'Gula', 1, 15000, 15000, 417300, '2024-06-29 13:25:03'),
(24, 0, 'Kopi', 1, 300, 300, 417300, '2024-06-29 13:25:03'),
(25, 0, 'Kopi ABC', 1, 2000, 2000, 417300, '2024-06-29 13:25:03'),
(26, 0, 'Mill', 1, 15000, 15000, 15000, '2024-06-29 13:25:03'),
(27, 0, 'Kopi ABC', 1, 2000, 2000, 2000, '2024-06-29 13:25:03'),
(36, 0, 'Zaky', 1, 1200, 1200, 40200, '2024-06-29 13:25:03'),
(1719666517954526758, 0, 'Kopi ABC', 1, 2000, 2000, 87000, '2024-06-29 23:21:11'),
(1719666517954526759, 0, 'Djarum', 5, 16000, 80000, 87000, '2024-06-29 23:21:11');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- Indexes for table `cart`
--
ALTER TABLE `cart`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_name` (`product_name`);

--
-- Indexes for table `pengguna`
--
ALTER TABLE `pengguna`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `nama` (`nama`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admin`
--
ALTER TABLE `admin`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `cart`
--
ALTER TABLE `cart`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=153;

--
-- AUTO_INCREMENT for table `pengguna`
--
ALTER TABLE `pengguna`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=39;

--
-- AUTO_INCREMENT for table `product`
--
ALTER TABLE `product`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=36;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1719666517954526762;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `cart`
--
ALTER TABLE `cart`
  ADD CONSTRAINT `cart_ibfk_1` FOREIGN KEY (`product_name`) REFERENCES `product` (`nama`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
