-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 20, 2021 at 02:12 AM
-- Server version: 10.4.21-MariaDB
-- PHP Version: 8.0.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_quiz3`
--

-- --------------------------------------------------------

--
-- Table structure for table `book`
--

CREATE TABLE `book` (
  `id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` longtext NOT NULL,
  `image_url` varchar(255) NOT NULL,
  `release_year` int(11) NOT NULL,
  `price` varchar(255) NOT NULL,
  `total_page` int(11) NOT NULL,
  `thickness` varchar(255) DEFAULT NULL,
  `category_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `book`
--

INSERT INTO `book` (`id`, `title`, `description`, `image_url`, `release_year`, `price`, `total_page`, `thickness`, `category_id`, `created_at`, `updated_at`) VALUES
(2, 'Laskar Pelangi', 'ini adalah sebuah buku untuk dibaca', 'www.google.com', 2019, '50000', 800, '', 3, '2021-12-20 07:04:38', '2021-12-20 07:04:38'),
(3, 'Laskarrrrrr!!!!!!', 'ini adalah sebuah buku untuk dibaca', 'www.google.com', 2019, '50000', 400, 'tebal', 1, '2021-12-20 07:38:44', '2021-12-20 07:38:44'),
(4, 'Laskarrrrrr!!!!!!', 'ini adalah sebuah buku untuk dibaca', 'http://www.google.com', 2019, '50000', 400, 'tebal', 1, '2021-12-20 07:47:45', '2021-12-20 07:47:45'),
(5, 'Laskarrrrrr!!!!!!', 'ini adalah sebuah buku untuk dibaca', 'http://www.google.com', 2019, '50000', 100, 'tipis', 1, '2021-12-20 07:48:06', '2021-12-20 07:48:06'),
(6, 'Laskarrrrrr!!!!!!', 'ini adalah sebuah buku untuk dibaca', 'http://www.google.com', 2019, '50000', 101, 'sedang', 1, '2021-12-20 07:48:12', '2021-12-20 07:48:12');

-- --------------------------------------------------------

--
-- Table structure for table `category`
--

CREATE TABLE `category` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `category`
--

INSERT INTO `category` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Novel', '2021-12-20 02:07:06', '2021-12-20 02:07:06'),
(3, 'Pelajaran', '2021-12-20 02:34:05', '2021-12-20 02:34:05');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `book`
--
ALTER TABLE `book`
  ADD PRIMARY KEY (`id`),
  ADD KEY `category_id` (`category_id`);

--
-- Indexes for table `category`
--
ALTER TABLE `category`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `book`
--
ALTER TABLE `book`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `category`
--
ALTER TABLE `category`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `book`
--
ALTER TABLE `book`
  ADD CONSTRAINT `book_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
