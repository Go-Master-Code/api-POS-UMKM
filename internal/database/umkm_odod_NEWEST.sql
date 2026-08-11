-- --------------------------------------------------------
-- Host:                         localhost
-- Server version:               5.7.43-log - MySQL Community Server (GPL)
-- Server OS:                    Win64
-- HeidiSQL Version:             12.8.0.6908
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for umkm_odod
CREATE DATABASE IF NOT EXISTS `umkm_odod` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
USE `umkm_odod`;

-- Dumping structure for table umkm_odod.activity_logs
CREATE TABLE IF NOT EXISTS `activity_logs` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `module` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `action` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `reference_id` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `reference_number` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_activity_logs_tenant` (`tenant_id`),
  KEY `idx_activity_logs_user` (`user_id`),
  KEY `idx_activity_logs_module` (`module`),
  KEY `idx_activity_logs_created_at` (`created_at`),
  CONSTRAINT `fk_activity_logs_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`),
  CONSTRAINT `fk_activity_logs_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.activity_logs: ~14 rows (approximately)
INSERT INTO `activity_logs` (`id`, `tenant_id`, `user_id`, `module`, `action`, `description`, `reference_id`, `reference_number`, `created_at`) VALUES
	('0091a01f-c0fb-4867-adae-cfc25c98ca50', '11111111-1111-1111-1111-111111111111', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'USER', 'UPDATE', 'Update Own Profile', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'heru.owner', '2026-06-15 00:44:14'),
	('060d7f8a-821e-4631-b4fb-1f477fe329d3', '11111111-1111-1111-1111-111111111111', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'PURCHASE RETURN', 'CREATE', 'Create Purchase Return P-RETUR-1781497398', '048ecf04-eee3-43ef-8e9b-9ed688dcd216', 'P-RETUR-1781497398', '2026-06-14 21:23:18'),
	('171d30a6-7ba6-4afa-aecf-c0289c7d8690', '11111111-1111-1111-1111-111111111111', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'USER', 'CHANGE PASSWORD', 'Change own password', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'AKBP Hendro Pandowo, M.Si.', '2026-06-15 00:26:22'),
	('272cbc6f-654d-4c27-845c-628e36f4729e', '11111111-1111-1111-1111-111111111111', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'USER', 'CHANGE PASSWORD', 'Change own password', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'heru.owner', '2026-07-06 03:13:45'),
	('414650f0-9279-48ac-9c84-a50c65fc8638', '11111111-1111-1111-1111-111111111111', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'USER', 'CHANGE PASSWORD', 'Change own password', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'heru.owner', '2026-07-06 03:57:08'),
	('6d676694-9bb4-488d-bee3-829636e7c331', '11111111-1111-1111-1111-111111111111', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'SALES', 'CREATE', 'Create Sales INV-1781495774', '155cf00c-3002-49cd-afc3-9723df7fb849', 'INV-1781495774', '2026-06-14 20:56:15'),
	('6e9ef33a-2676-4685-b4b7-be71c999d6e9', '11111111-1111-1111-1111-111111111111', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'SALES', 'CREATE', 'Create Sales INV-1786258125', '8e60bfab-58c2-4152-9c55-8d6fdf68d42e', 'INV-1786258125', '2026-08-09 06:48:46'),
	('7e8b3147-51ef-4dd5-b4e3-a3da4953316a', '11111111-1111-1111-1111-111111111111', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'SALES', 'CREATE', 'Create Sales INV-1786197188', '11c82d82-4587-4d24-87a2-e41e2911efec', 'INV-1786197188', '2026-08-08 13:53:09'),
	('98694d2a-ecc6-4dda-a76c-ec96147efa82', '11111111-1111-1111-1111-111111111111', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'USER', 'CHANGE PASSWORD', 'Change own password', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'AKBP Hendro Pandowo, M.Si.', '2026-06-15 00:43:35'),
	('9d5860fc-5eb2-43d9-89e7-08a93bb02bb3', '11111111-1111-1111-1111-111111111111', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'USER', 'UPDATE', 'Update Own Profile', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'heru.owner', '2026-06-14 23:48:38'),
	('b08a8442-250b-496f-ae52-eeae512ea952', '11111111-1111-1111-1111-111111111111', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'USER', 'CHANGE PASSWORD', 'Change own password', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'Heru Widodo', '2026-06-15 00:45:08'),
	('df5de65a-ad84-4731-9108-315b1bd5dd64', '11111111-1111-1111-1111-111111111111', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'USER', 'CHANGE PASSWORD', 'Change own password', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'Heru Widodo', '2026-06-15 00:44:57'),
	('df9792e7-b9e8-48de-8ad7-970b3782a210', '11111111-1111-1111-1111-111111111111', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'PURCHASE', 'CREATE', 'Create Purchase PO-1781495323', '5529269d-e8dc-42bb-95c0-adc201afa138', 'PO-1781495323', '2026-06-14 20:48:44'),
	('fe9edff4-0700-47c7-8996-e57e538f7e6a', '11111111-1111-1111-1111-111111111111', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'USER', 'CHANGE PASSWORD', 'Change own password', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 'heru.owner', '2026-07-06 03:56:29');

-- Dumping structure for table umkm_odod.catalog_categories
CREATE TABLE IF NOT EXISTS `catalog_categories` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_categories_name_per_tenant` (`tenant_id`,`name`),
  CONSTRAINT `fk_categories_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.catalog_categories: ~6 rows (approximately)
INSERT INTO `catalog_categories` (`id`, `tenant_id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('08e850d4-6590-4cef-91bd-df349c271019', '11111111-1111-1111-1111-111111111111', 'Kue Basah', '2026-05-20 07:54:18', '2026-05-20 07:54:18', NULL),
	('3b6c015f-9304-4673-9c96-8d8287c01cf6', '11111111-1111-1111-1111-111111111111', 'Makanan kering', '2026-07-30 04:33:45', '2026-07-30 04:44:26', NULL),
	('44444444-4444-4444-4444-444444444441', '11111111-1111-1111-1111-111111111111', 'Snack', '2026-05-18 09:21:24', '2026-05-20 07:22:22', NULL),
	('44444444-4444-4444-4444-444444444442', '11111111-1111-1111-1111-111111111111', 'Minuman', '2026-05-18 09:21:24', '2026-05-20 06:50:33', NULL),
	('8a36bc96-32af-4f1e-b908-7c114d92786f', '11111111-1111-1111-1111-111111111111', 'Frozen Food', '2026-07-30 04:30:46', '2026-07-30 04:43:59', NULL),
	('db37583c-5c4e-4b27-8d1e-739883f2c242', '11111111-1111-1111-1111-111111111111', 'Makanan kucing', '2026-07-29 14:02:08', '2026-07-30 04:29:24', '2026-07-30 04:29:24');

-- Dumping structure for table umkm_odod.catalog_items
CREATE TABLE IF NOT EXISTS `catalog_items` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `category_id` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_items_name_per_tenant` (`tenant_id`,`name`),
  KEY `fk_items_tenant` (`tenant_id`),
  KEY `fk_items_category` (`category_id`),
  CONSTRAINT `fk_items_category` FOREIGN KEY (`category_id`) REFERENCES `catalog_categories` (`id`),
  CONSTRAINT `fk_items_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.catalog_items: ~7 rows (approximately)
INSERT INTO `catalog_items` (`id`, `tenant_id`, `category_id`, `name`, `description`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('2b463c6f-fb30-4531-867a-ac33b7162fbb', '11111111-1111-1111-1111-111111111111', '08e850d4-6590-4cef-91bd-df349c271019', 'Dodol Garut', 'Hanya tahan 3 bulan', 0, '2026-07-30 12:43:41', '2026-07-30 12:47:49', '2026-07-30 12:47:50'),
	('55555555-5555-5555-5555-555555555551', '11111111-1111-1111-1111-111111111111', '44444444-4444-4444-4444-444444444441', 'Keripik Singkong', 'Keripik singkong premium', 1, '2026-05-18 09:21:44', '2026-05-20 07:34:53', NULL),
	('55555555-5555-5555-5555-555555555552', '11111111-1111-1111-1111-111111111111', '44444444-4444-4444-4444-444444444441', 'Keripik Pisang', 'Keripik pisang renyah', 1, '2026-05-18 09:21:44', '2026-05-20 07:34:23', NULL),
	('592fcb7b-4451-46e9-ba1d-3039a70ee61d', 'f27e441f-5385-4b8d-b2e2-88b8615a4634', '44444444-4444-4444-4444-444444444442', 'Fruit tea', 'Rasa Black Currant', 0, '2026-05-20 07:37:28', '2026-05-20 07:49:08', NULL),
	('77b4a60f-813f-46e3-a8a8-6609a6609da7', '11111111-1111-1111-1111-111111111111', '44444444-4444-4444-4444-444444444441', 'Kerupuk Kulit', 'Dorokdok sapi', 1, '2026-07-30 12:52:07', '2026-07-30 12:52:07', NULL),
	('c1124d3b-2894-420a-8878-556f96af06db', '11111111-1111-1111-1111-111111111111', '08e850d4-6590-4cef-91bd-df349c271019', 'Nagasari', 'Nagasari pisang ambon', 1, '2026-05-20 07:55:51', '2026-05-20 07:55:51', NULL),
	('ed32d30e-37b1-4ff7-b3e1-d80bf55b86c4', '11111111-1111-1111-1111-111111111111', '44444444-4444-4444-4444-444444444441', 'Kue sagu', 'Kue sagu putih', 1, '2026-07-30 12:52:23', '2026-08-07 01:36:02', NULL);

-- Dumping structure for table umkm_odod.item_attributes
CREATE TABLE IF NOT EXISTS `item_attributes` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_item_attributes_name_per_tenant` (`tenant_id`,`name`),
  CONSTRAINT `fk_item_attributes_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.item_attributes: ~2 rows (approximately)
INSERT INTO `item_attributes` (`id`, `tenant_id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('66666666-6666-6666-6666-666666666661', '11111111-1111-1111-1111-111111111111', 'Rasa', '2026-05-18 09:21:58', '2026-05-18 09:21:58', NULL),
	('66666666-6666-6666-6666-666666666662', '11111111-1111-1111-1111-111111111111', 'Berat', '2026-05-18 09:21:58', '2026-05-18 09:21:58', NULL);

-- Dumping structure for table umkm_odod.item_attribute_values
CREATE TABLE IF NOT EXISTS `item_attribute_values` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `attribute_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_attribute_values_tenant` (`tenant_id`),
  KEY `fk_attribute_values_attribute` (`attribute_id`),
  CONSTRAINT `fk_attribute_values_attribute` FOREIGN KEY (`attribute_id`) REFERENCES `item_attributes` (`id`),
  CONSTRAINT `fk_attribute_values_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.item_attribute_values: ~4 rows (approximately)
INSERT INTO `item_attribute_values` (`id`, `tenant_id`, `attribute_id`, `value`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('77777777-7777-7777-7777-777777777771', '11111111-1111-1111-1111-111111111111', '66666666-6666-6666-6666-666666666661', 'Pedas Daun Jeruk', '2026-05-18 09:22:08', '2026-05-18 09:22:08', NULL),
	('77777777-7777-7777-7777-777777777772', '11111111-1111-1111-1111-111111111111', '66666666-6666-6666-6666-666666666661', 'Asin Bawang', '2026-05-18 09:22:08', '2026-05-18 09:22:08', NULL),
	('77777777-7777-7777-7777-777777777773', '11111111-1111-1111-1111-111111111111', '66666666-6666-6666-6666-666666666662', '250 Gram', '2026-05-18 09:22:08', '2026-05-18 09:22:08', NULL),
	('77777777-7777-7777-7777-777777777774', '11111111-1111-1111-1111-111111111111', '66666666-6666-6666-6666-666666666662', '500 Gram', '2026-05-18 09:22:08', '2026-05-18 09:22:08', NULL);

-- Dumping structure for table umkm_odod.item_variants
CREATE TABLE IF NOT EXISTS `item_variants` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sku` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `barcode` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `variant_name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
  `minimum_stock` decimal(18,2) NOT NULL DEFAULT '0.00',
  `cost_price` decimal(18,2) NOT NULL DEFAULT '0.00',
  `selling_price` decimal(18,2) NOT NULL DEFAULT '0.00',
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_variants_sku_per_tenant` (`tenant_id`,`sku`) USING BTREE,
  UNIQUE KEY `uq_variant_name_per_item` (`tenant_id`,`item_id`,`variant_name`),
  KEY `fk_variants_item` (`item_id`),
  CONSTRAINT `fk_variants_item` FOREIGN KEY (`item_id`) REFERENCES `catalog_items` (`id`),
  CONSTRAINT `fk_variants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.item_variants: ~5 rows (approximately)
INSERT INTO `item_variants` (`id`, `tenant_id`, `item_id`, `sku`, `barcode`, `variant_name`, `minimum_stock`, `cost_price`, `selling_price`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('29454d2b-8cff-4e9a-9c47-6865dc7aa36f', '11111111-1111-1111-1111-111111111111', '55555555-5555-5555-5555-555555555551', 'KS-AW-01', '26172871', 'Ayam Woku', 10.00, 11800.00, 19800.00, 1, '2026-08-07 01:39:17', '2026-08-08 11:03:17', NULL),
	('76a3c5f8-4340-42a4-8aa0-db591a791c58', '11111111-1111-1111-1111-111111111111', '55555555-5555-5555-5555-555555555551', 'KS-PB-500', '899100000003', 'Pedas Balado', 34.00, 13000.00, 15000.00, 1, '2026-05-20 10:45:44', '2026-08-08 08:05:54', NULL),
	('88888888-8888-8888-8888-888888888881', '11111111-1111-1111-1111-111111111111', '55555555-5555-5555-5555-555555555551', 'KS-PDJ-250', '899100000001', 'Pedas Daun Jeruk 250gr', 5.00, 13000.00, 15000.00, 1, '2026-05-18 09:22:19', '2026-06-02 08:23:38', NULL),
	('88888888-8888-8888-8888-888888888882', '11111111-1111-1111-1111-111111111111', '55555555-5555-5555-5555-555555555551', 'KS-AB-500', '899100000002', 'Asin Bawang 500gr', 10.00, 12000.00, 14000.00, 1, '2026-05-18 09:22:19', '2026-08-06 15:44:45', NULL),
	('f3afc64d-5c56-43a8-927a-2deec7a5ca69', '11111111-1111-1111-1111-111111111111', '55555555-5555-5555-5555-555555555551', 'KS-SP-11', '899100000004', 'Sapi Panggang', 12.00, 13000.00, 17500.00, 1, '2026-08-08 07:44:46', '2026-08-08 11:02:41', NULL);

-- Dumping structure for table umkm_odod.item_variant_attribute_values
CREATE TABLE IF NOT EXISTS `item_variant_attribute_values` (
  `variant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `attribute_value_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`variant_id`,`attribute_value_id`),
  KEY `fk_variant_attribute_value` (`attribute_value_id`),
  KEY `FK_item_variant_attribute_values_tenants` (`tenant_id`),
  CONSTRAINT `FK_item_variant_attribute_values_tenants` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`),
  CONSTRAINT `fk_variant_attribute_value` FOREIGN KEY (`attribute_value_id`) REFERENCES `item_attribute_values` (`id`),
  CONSTRAINT `fk_variant_attribute_variant` FOREIGN KEY (`variant_id`) REFERENCES `item_variants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.item_variant_attribute_values: ~4 rows (approximately)
INSERT INTO `item_variant_attribute_values` (`variant_id`, `attribute_value_id`, `tenant_id`) VALUES
	('88888888-8888-8888-8888-888888888881', '77777777-7777-7777-7777-777777777771', '11111111-1111-1111-1111-111111111111'),
	('88888888-8888-8888-8888-888888888881', '77777777-7777-7777-7777-777777777773', '11111111-1111-1111-1111-111111111111'),
	('88888888-8888-8888-8888-888888888882', '77777777-7777-7777-7777-777777777772', '11111111-1111-1111-1111-111111111111'),
	('88888888-8888-8888-8888-888888888882', '77777777-7777-7777-7777-777777777774', '11111111-1111-1111-1111-111111111111');

-- Dumping structure for table umkm_odod.price_histories
CREATE TABLE IF NOT EXISTS `price_histories` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_variant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `price_type` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `price` decimal(18,2) NOT NULL,
  `effective_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_by` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_prices_tenant` (`tenant_id`),
  KEY `fk_prices_variant` (`item_variant_id`),
  KEY `fk_prices_created_by` (`created_by`),
  CONSTRAINT `fk_prices_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_prices_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`),
  CONSTRAINT `fk_prices_variant` FOREIGN KEY (`item_variant_id`) REFERENCES `item_variants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.price_histories: ~2 rows (approximately)
INSERT INTO `price_histories` (`id`, `tenant_id`, `item_variant_id`, `price_type`, `price`, `effective_date`, `created_by`, `created_at`) VALUES
	('99999999-9999-9999-9999-999999999991', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SELLING', 18000.00, '2026-05-20 06:36:19', '2416854f-55e6-423c-a2ec-8154c9431cd6', '2026-05-18 09:23:03'),
	('99999999-9999-9999-9999-999999999992', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888882', 'SELLING', 25000.00, '2026-05-20 06:36:20', '2416854f-55e6-423c-a2ec-8154c9431cd6', '2026-05-18 09:23:03');

-- Dumping structure for table umkm_odod.purchases
CREATE TABLE IF NOT EXISTS `purchases` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `purchase_number` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `supplier_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `invoice_number` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `subtotal` decimal(18,2) NOT NULL DEFAULT '0.00',
  `discount_amount` decimal(18,2) NOT NULL DEFAULT '0.00',
  `tax_amount` decimal(18,2) NOT NULL DEFAULT '0.00',
  `grand_total` decimal(18,2) NOT NULL DEFAULT '0.00',
  `notes` text COLLATE utf8mb4_unicode_ci,
  `created_by` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_purchase_number` (`purchase_number`) USING BTREE,
  KEY `idx_purchase_tenant` (`tenant_id`) USING BTREE,
  KEY `idx_purchase_supplier` (`supplier_id`) USING BTREE,
  KEY `idx_purchase_created_by` (`created_by`) USING BTREE,
  KEY `idx_purchase_created_at` (`created_at`) USING BTREE,
  CONSTRAINT `fk_purchase_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_purchase_supplier` FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`),
  CONSTRAINT `fk_purchase_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.purchases: ~2 rows (approximately)
INSERT INTO `purchases` (`id`, `tenant_id`, `purchase_number`, `supplier_id`, `invoice_number`, `subtotal`, `discount_amount`, `tax_amount`, `grand_total`, `notes`, `created_by`, `created_at`, `updated_at`) VALUES
	('4f3f3907-2536-4587-9d32-a2ae4751a9aa', '11111111-1111-1111-1111-111111111111', 'PO-1780935589', '4efba307-0fbd-4501-bb8d-3f9212c14fb9', 'SUP-INV-001', 169000.00, 5000.00, 16900.00, 180900.00, 'Pembelian stok mingguan', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-07-28 16:19:49', '2026-07-30 04:26:06'),
	('d1b0e60f-b894-4320-b2e0-fae446d06e36', '11111111-1111-1111-1111-111111111111', 'PO-1780934118', '4efba307-0fbd-4501-bb8d-3f9212c14fb9', 'SUP-INV-002', 169000.00, 0.00, 0.00, 169000.00, 'Pembelian stok mingguan', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-07-29 15:55:18', '2026-07-30 04:26:09');

-- Dumping structure for table umkm_odod.purchase_items
CREATE TABLE IF NOT EXISTS `purchase_items` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `purchase_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_variant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_name_snapshot` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `variant_name_snapshot` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sku_snapshot` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `qty` decimal(18,2) NOT NULL,
  `cost_price` decimal(18,2) NOT NULL,
  `discount_amount` decimal(18,2) NOT NULL DEFAULT '0.00',
  `subtotal` decimal(18,2) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_purchase_item_tenant` (`tenant_id`) USING BTREE,
  KEY `idx_purchase_item_purchase` (`purchase_id`) USING BTREE,
  KEY `idx_purchase_item_variant` (`item_variant_id`) USING BTREE,
  CONSTRAINT `fk_purchase_item_purchase` FOREIGN KEY (`purchase_id`) REFERENCES `purchases` (`id`),
  CONSTRAINT `fk_purchase_item_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`),
  CONSTRAINT `fk_purchase_item_variant` FOREIGN KEY (`item_variant_id`) REFERENCES `item_variants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.purchase_items: ~4 rows (approximately)
INSERT INTO `purchase_items` (`id`, `tenant_id`, `purchase_id`, `item_variant_id`, `item_name_snapshot`, `variant_name_snapshot`, `sku_snapshot`, `qty`, `cost_price`, `discount_amount`, `subtotal`, `created_at`) VALUES
	('06377b76-b7de-485d-ab3e-8e20de6cbd63', '11111111-1111-1111-1111-111111111111', 'd1b0e60f-b894-4320-b2e0-fae446d06e36', '88888888-8888-8888-8888-888888888882', 'Keripik Singkong', 'Asin Bawang 500gr', 'KS-AB-500', 5.00, 10000.00, 0.00, 50000.00, '2026-06-08 15:55:18'),
	('802ad168-06c6-4fc8-8d95-c44aec8f1ac5', '11111111-1111-1111-1111-111111111111', '4f3f3907-2536-4587-9d32-a2ae4751a9aa', '88888888-8888-8888-8888-888888888881', 'Keripik Singkong', 'Pedas Daun Jeruk 250gr', 'KS-PDJ-250', 5.00, 10000.00, 0.00, 50000.00, '2026-06-08 16:19:49'),
	('94624e7c-97ef-4e65-a1ed-356a6247bf01', '11111111-1111-1111-1111-111111111111', 'd1b0e60f-b894-4320-b2e0-fae446d06e36', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'Keripik Singkong', 'Pedas Balado', 'KS-PB-500', 10.00, 12000.00, 1000.00, 119000.00, '2026-06-08 15:55:18'),
	('da05d590-c091-4e18-ad93-bed0fb341f2e', '11111111-1111-1111-1111-111111111111', '4f3f3907-2536-4587-9d32-a2ae4751a9aa', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'Keripik Singkong', 'Pedas Balado', 'KS-PB-500', 10.00, 12000.00, 1000.00, 119000.00, '2026-06-08 16:19:49');

-- Dumping structure for table umkm_odod.purchase_returns
CREATE TABLE IF NOT EXISTS `purchase_returns` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `purchase_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `return_number` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `reason` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `notes` text COLLATE utf8mb4_unicode_ci,
  `created_by` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_purchase_returns_tenant` (`tenant_id`),
  KEY `idx_purchase_returns_purchase` (`purchase_id`),
  CONSTRAINT `fk_purchase_returns_purchase` FOREIGN KEY (`purchase_id`) REFERENCES `purchases` (`id`),
  CONSTRAINT `fk_purchase_returns_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.purchase_returns: ~3 rows (approximately)
INSERT INTO `purchase_returns` (`id`, `tenant_id`, `purchase_id`, `return_number`, `reason`, `notes`, `created_by`, `created_at`) VALUES
	('048ecf04-eee3-43ef-8e9b-9ed688dcd216', '11111111-1111-1111-1111-111111111111', 'd1b0e60f-b894-4320-b2e0-fae446d06e36', 'P-RETUR-1781497398', 'plastik sobek', 'bahan terlalu tipis', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-06-14 21:23:18'),
	('49d03c48-8d30-4610-b8bd-d3fb4a082cc8', '11111111-1111-1111-1111-111111111111', 'd1b0e60f-b894-4320-b2e0-fae446d06e36', 'P-RETUR-1781140525', 'plastik sobek', 'bahan terlalu tipis', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-06-10 18:15:26'),
	('9fd956eb-0f22-46a9-890a-e4dffc07136e', '11111111-1111-1111-1111-111111111111', '4f3f3907-2536-4587-9d32-a2ae4751a9aa', 'P-RETUR-1781065285', 'melempem', 'seal tidak rapat', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-06-09 21:21:26');

-- Dumping structure for table umkm_odod.purchase_return_items
CREATE TABLE IF NOT EXISTS `purchase_return_items` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `purchase_return_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_variant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `qty` decimal(18,2) NOT NULL,
  `notes` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_purchase_return_items_tenant` (`tenant_id`),
  KEY `fk_purchase_return_items_return` (`purchase_return_id`),
  CONSTRAINT `fk_purchase_return_items_return` FOREIGN KEY (`purchase_return_id`) REFERENCES `purchase_returns` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.purchase_return_items: ~4 rows (approximately)
INSERT INTO `purchase_return_items` (`id`, `tenant_id`, `purchase_return_id`, `item_variant_id`, `qty`, `notes`, `created_at`) VALUES
	('4a802abc-7b9d-466e-8739-88eb94a8c10f', '11111111-1111-1111-1111-111111111111', '9fd956eb-0f22-46a9-890a-e4dffc07136e', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 3.00, '', '2026-06-09 21:21:26'),
	('8d5aaab3-d168-4d45-8fd0-592d84088f63', '11111111-1111-1111-1111-111111111111', '9fd956eb-0f22-46a9-890a-e4dffc07136e', '88888888-8888-8888-8888-888888888881', 2.00, '', '2026-06-09 21:21:26'),
	('c97b88cc-f277-4ec5-b221-d39ced0f8b1a', '11111111-1111-1111-1111-111111111111', '048ecf04-eee3-43ef-8e9b-9ed688dcd216', '88888888-8888-8888-8888-888888888882', 5.00, '', '2026-06-14 21:23:18'),
	('d947faac-4420-44c5-b19b-4116bcd469d5', '11111111-1111-1111-1111-111111111111', '49d03c48-8d30-4610-b8bd-d3fb4a082cc8', '88888888-8888-8888-8888-888888888882', 5.00, '', '2026-06-10 18:15:26');

-- Dumping structure for table umkm_odod.roles
CREATE TABLE IF NOT EXISTS `roles` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_roles_name_per_tenant` (`tenant_id`,`name`),
  CONSTRAINT `fk_roles_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.roles: ~5 rows (approximately)
INSERT INTO `roles` (`id`, `tenant_id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('22222222-2222-2222-2222-222222222221', '11111111-1111-1111-1111-111111111112', 'OWNER', '2026-05-18 09:20:36', '2026-05-18 16:35:19', NULL),
	('22222222-2222-2222-2222-222222222222', '11111111-1111-1111-1111-111111111111', 'CASHIER', '2026-05-18 09:20:36', '2026-05-18 09:20:36', NULL),
	('22222222-2222-2222-2222-222222222223', '11111111-1111-1111-1111-111111111111', 'ADMIN', '2026-06-02 10:01:02', '2026-06-02 10:01:02', NULL),
	('ab88f666-b2d6-482c-bdcd-be29915e5395', 'f27e441f-5385-4b8d-b2e2-88b8615a4634', 'SUPER_ADMIN', '2026-05-19 02:02:01', '2026-05-28 10:32:27', NULL),
	('dc064501-5798-4c91-8308-0198561ceae3', 'f27e441f-5385-4b8d-b2e2-88b8615a4634', 'ADMIN', '2026-05-19 01:56:21', '2026-05-19 02:17:01', NULL);

-- Dumping structure for table umkm_odod.sales
CREATE TABLE IF NOT EXISTS `sales` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `invoice_number` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `customer_name` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `cashier_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `subtotal` decimal(18,2) NOT NULL DEFAULT '0.00',
  `discount_amount` decimal(18,2) NOT NULL DEFAULT '0.00',
  `tax_amount` decimal(18,2) NOT NULL DEFAULT '0.00',
  `grand_total` decimal(18,2) NOT NULL DEFAULT '0.00',
  `payment_method` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `payment_status` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `notes` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_sales_invoice_per_tenant` (`tenant_id`,`invoice_number`),
  KEY `fk_sales_cashier` (`cashier_id`),
  CONSTRAINT `fk_sales_cashier` FOREIGN KEY (`cashier_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_sales_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.sales: ~6 rows (approximately)
INSERT INTO `sales` (`id`, `tenant_id`, `invoice_number`, `customer_name`, `cashier_id`, `subtotal`, `discount_amount`, `tax_amount`, `grand_total`, `payment_method`, `payment_status`, `notes`, `created_at`, `updated_at`) VALUES
	('11c82d82-4587-4d24-87a2-e41e2911efec', '11111111-1111-1111-1111-111111111111', 'INV-1786197188', 'Budi Santoso', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 104000.00, 5000.00, 10400.00, 109400.00, 'QRIS', 'PAID', 'Pembelian sore hari', '2026-08-08 13:53:08', '2026-08-08 13:53:09'),
	('4cf1c87a-0014-47e6-9b46-d3f4ec980377', '11111111-1111-1111-1111-111111111111', 'INV-1779948538', 'Budi', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 27000.00, 500.00, 2700.00, 29200.00, 'CASH', 'PAID', 'Pembelian 1 item', '2026-07-28 06:08:58', '2026-07-30 04:26:46'),
	('8e60bfab-58c2-4152-9c55-8d6fdf68d42e', '11111111-1111-1111-1111-111111111111', 'INV-1786258125', 'Budi', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 55500.00, 500.00, 5500.00, 60500.00, 'CASH', 'PAID', 'Pembelian 1 item', '2026-08-09 06:48:46', '2026-08-09 06:48:46'),
	('bcec37b4-a8bb-44a2-a039-cbe8cb06a530', '11111111-1111-1111-1111-111111111111', 'INV-1779941000', 'Budi Santoso', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 104000.00, 5000.00, 10400.00, 109400.00, 'QRIS', 'PAID', 'Pembelian sore hari', '2026-07-29 04:03:21', '2026-07-30 04:26:49'),
	('ecf9ac1a-6c21-4881-80d5-b741e62ad205', '11111111-1111-1111-1111-111111111111', 'INV-1779948742', 'Andi', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', 437000.00, 10000.00, 43700.00, 470700.00, 'TRANSFER', 'UNPAID', 'Pembelian grosir', '2026-07-30 06:12:22', '2026-07-30 04:26:53');

-- Dumping structure for table umkm_odod.sale_items
CREATE TABLE IF NOT EXISTS `sale_items` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sale_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_variant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_name_snapshot` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `variant_name_snapshot` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sku_snapshot` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `qty` decimal(18,2) NOT NULL,
  `unit_price` decimal(18,2) NOT NULL,
  `discount_amount` decimal(18,2) NOT NULL DEFAULT '0.00',
  `subtotal` decimal(18,2) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_sale_items_tenant` (`tenant_id`),
  KEY `fk_sale_items_sale` (`sale_id`),
  KEY `fk_sale_items_variant` (`item_variant_id`),
  CONSTRAINT `fk_sale_items_sale` FOREIGN KEY (`sale_id`) REFERENCES `sales` (`id`),
  CONSTRAINT `fk_sale_items_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`),
  CONSTRAINT `fk_sale_items_variant` FOREIGN KEY (`item_variant_id`) REFERENCES `item_variants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.sale_items: ~10 rows (approximately)
INSERT INTO `sale_items` (`id`, `tenant_id`, `sale_id`, `item_variant_id`, `item_name_snapshot`, `variant_name_snapshot`, `sku_snapshot`, `qty`, `unit_price`, `discount_amount`, `subtotal`, `created_at`) VALUES
	('13574740-8adf-4885-842c-7009b8567392', '11111111-1111-1111-1111-111111111111', 'ecf9ac1a-6c21-4881-80d5-b741e62ad205', '88888888-8888-8888-8888-888888888882', 'Keripik Singkong', 'Asin Bawang 500gr', 'KS-AB-500', 10.00, 14000.00, 2000.00, 138000.00, '2026-05-28 06:12:22'),
	('1d10f8cb-e35d-4cab-9035-762eeee4cfdb', '11111111-1111-1111-1111-111111111111', 'bcec37b4-a8bb-44a2-a039-cbe8cb06a530', '88888888-8888-8888-8888-888888888881', 'Keripik Singkong', 'Pedas Daun Jeruk 250gr', 'KS-PDJ-250', 5.00, 15000.00, 0.00, 75000.00, '2026-05-28 04:03:21'),
	('1f17845a-f5b6-48b2-a93c-b45daae9ccfa', '11111111-1111-1111-1111-111111111111', 'ecf9ac1a-6c21-4881-80d5-b741e62ad205', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'Keripik Singkong', 'Pedas Balado', 'KS-PB-500', 7.00, 15000.00, 1000.00, 104000.00, '2026-05-28 06:12:22'),
	('3d09c56a-e883-46bb-a25e-36d0b34735c5', '11111111-1111-1111-1111-111111111111', '8e60bfab-58c2-4152-9c55-8d6fdf68d42e', '88888888-8888-8888-8888-888888888882', 'Keripik Singkong', 'Asin Bawang 500gr', 'KS-AB-500', 4.00, 14000.00, 500.00, 55500.00, '2026-08-09 06:48:46'),
	('714c55df-935f-4a23-980a-716457985a5d', '11111111-1111-1111-1111-111111111111', '11c82d82-4587-4d24-87a2-e41e2911efec', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'Keripik Singkong', 'Pedas Balado', 'KS-PB-500', 2.00, 15000.00, 1000.00, 29000.00, '2026-08-08 13:53:09'),
	('8bca2936-1417-4d3a-b2d4-4cd119637451', '11111111-1111-1111-1111-111111111111', 'bcec37b4-a8bb-44a2-a039-cbe8cb06a530', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'Keripik Singkong', 'Pedas Balado', 'KS-PB-500', 2.00, 15000.00, 1000.00, 29000.00, '2026-05-28 04:03:21'),
	('adcb2348-f629-4dda-a8c0-535039e03942', '11111111-1111-1111-1111-111111111111', '4cf1c87a-0014-47e6-9b46-d3f4ec980377', '88888888-8888-8888-8888-888888888882', 'Keripik Singkong', 'Asin Bawang 500gr', 'KS-AB-500', 2.00, 14000.00, 1000.00, 27000.00, '2026-05-28 06:08:58'),
	('beddda53-3faa-46b7-a4dc-6fb00befb751', '11111111-1111-1111-1111-111111111111', 'ecf9ac1a-6c21-4881-80d5-b741e62ad205', '88888888-8888-8888-8888-888888888881', 'Keripik Singkong', 'Pedas Daun Jeruk 250gr', 'KS-PDJ-250', 13.00, 15000.00, 0.00, 195000.00, '2026-05-28 06:12:22'),
	('e76754af-f91d-46d9-acc8-c0aca1423d02', '11111111-1111-1111-1111-111111111111', '11c82d82-4587-4d24-87a2-e41e2911efec', '88888888-8888-8888-8888-888888888881', 'Keripik Singkong', 'Pedas Daun Jeruk 250gr', 'KS-PDJ-250', 5.00, 15000.00, 0.00, 75000.00, '2026-08-08 13:53:09');

-- Dumping structure for table umkm_odod.stock_movements
CREATE TABLE IF NOT EXISTS `stock_movements` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_variant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `movement_type` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `qty` decimal(18,2) NOT NULL,
  `reference_type` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `reference_id` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `notes` text COLLATE utf8mb4_unicode_ci,
  `created_by` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_stock_tenant` (`tenant_id`),
  KEY `fk_stock_variant` (`item_variant_id`),
  KEY `fk_stock_created_by` (`created_by`),
  CONSTRAINT `fk_stock_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_stock_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`),
  CONSTRAINT `fk_stock_variant` FOREIGN KEY (`item_variant_id`) REFERENCES `item_variants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.stock_movements: ~63 rows (approximately)
INSERT INTO `stock_movements` (`id`, `tenant_id`, `item_variant_id`, `movement_type`, `qty`, `reference_type`, `reference_id`, `notes`, `created_by`, `created_at`) VALUES
	('1012475d-2a7f-4fd5-88dc-8e399e3ee3c6', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -10.00, '', '', 'Bapak Udin', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-21 09:23:40'),
	('119de780-c9bb-48fc-bbad-3fd8ab8b2118', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -7.00, 'SALE', 'ecf9ac1a-6c21-4881-80d5-b741e62ad205', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 06:12:22'),
	('1776a9be-4f3a-4a9e-9c89-ee71eaee709f', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', '0f79d701-b78f-452e-a0cc-62c7d2ac35e8', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:20:06'),
	('1b85f780-2583-4589-bb9e-4401431b2866', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888882', 'PURCHASE', 5.00, 'PURCHASE', 'd1b0e60f-b894-4320-b2e0-fae446d06e36', 'purchase transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-06-08 15:55:18'),
	('1e38c273-cb4b-47db-be33-9a1f53739337', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', 'bcec37b4-a8bb-44a2-a039-cbe8cb06a530', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 04:03:21'),
	('1f880299-0e64-4c4b-a3e8-15be0ab2052a', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'REDUCE', -3.00, 'ADJUSTMENT', '', 'Reason: kencing tikus | Notes: plastik berbau', '64ee38a5-23c6-4a82-b1ad-d488ccc0d8e6', '2026-05-29 02:06:48'),
	('204020ba-96b8-4394-8033-26d0cd6cd286', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '2d61cc53-394b-4ab8-9354-dbd8f7212150', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:23:39'),
	('219bc491-cc1e-4dcd-be2b-f57016e12bb5', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -1.00, 'SALE', '9b34dc3c-db00-400b-a09a-d0eb9ef21d52', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:06:49'),
	('251daba8-f80e-4f18-8bfc-1fca2da08c50', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '0565b794-0d7f-4ce4-92d6-13aa5cd6007f', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:19:06'),
	('26bc763c-c689-47c8-829a-ac91c987d33a', '11111111-1111-1111-1111-111111111111', 'f3afc64d-5c56-43a8-927a-2deec7a5ca69', 'INITIAL STOCK', 9.00, 'ITEM_VARIANT', 'f3afc64d-5c56-43a8-927a-2deec7a5ca69', 'Initial Stock', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-08-08 07:44:46'),
	('26ffdab0-e6b7-4103-8110-48e44b38eebe', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -3.00, 'SALE', '609c139e-4a73-4453-8462-c6e3d39150ba', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:35:57'),
	('27514e4a-a357-45ba-adc3-9275289b22e8', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', 'b5c9ed63-63f0-46c1-95e3-c77fb096049d', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:26:25'),
	('27c2aa35-82aa-4410-b218-43213e8e9aa0', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -5.00, 'SALE', '58171b7c-b6f8-49ee-b4a3-41d8fa84f5f9', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:05:04'),
	('28bc8dbc-c94a-403e-bbaf-d8318a6617f7', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '1500c864-83a8-49e0-a704-f962ee94fe57', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:25:09'),
	('28c420a2-846c-45e0-9ed3-4f74c6799d2d', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'PURCHASE', 150.00, '', '', 'Stock Awal', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 06:11:41'),
	('2d6f8aac-8b17-4368-b744-a1b14b917823', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', '1500c864-83a8-49e0-a704-f962ee94fe57', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:25:09'),
	('33c0109a-37fc-4b85-a02c-501e193c5e35', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -10.00, 'SALE', '55f23f83-a2d0-45a2-a913-078c617d04c8', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:32:07'),
	('3d73a734-e367-4b76-a120-499f73fb1b4e', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', '974dd35c-53fe-4a40-ae94-2f8ca4229c48', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:21:27'),
	('480f2f60-dd50-41d7-8078-243cebad4196', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '11c82d82-4587-4d24-87a2-e41e2911efec', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-08-08 13:53:09'),
	('4de4596b-b553-43ee-9414-af7f78a1625c', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -1.00, 'SALE', 'db3c1825-a8b3-449a-a999-f7e5dc8c1e6e', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:05:13'),
	('515406e0-5306-4fe0-88cb-d405c508b27d', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888882', 'SALE', -4.00, 'SALE', '8e60bfab-58c2-4152-9c55-8d6fdf68d42e', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-08-09 06:48:46'),
	('54e130e8-f38a-4612-86d5-c4a1df869fbe', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'PURCHASE', 15.00, 'PURCHASE', '6a208d10-09a5-4e0b-8abe-c293602f0po9', 'Repack kedua', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-21 09:15:48'),
	('58481476-b28d-469b-88bb-88328d67568e', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', 'c4585fcf-5364-479d-b6dd-5ab2d414fd10', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:12:06'),
	('5c044814-78e5-4ce0-8589-72377b417899', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '49e45d2b-7a13-4906-a292-b7d870bb1a8e', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:15:35'),
	('5c2fb1ed-7339-473c-ba3f-4608a94d6ad1', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -13.00, 'SALE', 'b50b0a06-422e-446c-865c-643a43beab27', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:22:20'),
	('5fa4dd52-b554-4eec-8ddd-1e08156b6425', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -13.00, 'SALE', 'ecf9ac1a-6c21-4881-80d5-b741e62ad205', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 06:12:22'),
	('64cddab1-8093-41f9-9df3-790b76eb8a74', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -1.00, 'SALE', 'c4585fcf-5364-479d-b6dd-5ab2d414fd10', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:12:06'),
	('6eace372-d348-422a-8335-a6c7bb44690d', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', 'db3c1825-a8b3-449a-a999-f7e5dc8c1e6e', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:05:13'),
	('704bf63c-9bfb-4cd2-af8b-726897576591', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -4.00, 'SALE', 'c8a4629d-c6fc-46c8-bc9d-608bec8c554e', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:16:02'),
	('705e9eed-9def-4588-98b6-66cbd32e83d4', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '9b34dc3c-db00-400b-a09a-d0eb9ef21d52', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:06:49'),
	('745b00da-60ad-48dd-8016-737c9af4a588', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', '2d61cc53-394b-4ab8-9354-dbd8f7212150', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:23:39'),
	('7bb76694-03f8-4b12-b00c-ebf66213c8ff', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888882', 'SALE', -10.00, 'SALE', 'ecf9ac1a-6c21-4881-80d5-b741e62ad205', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 06:12:22'),
	('807eb585-0eb7-4b03-883a-fb27b17de013', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', '11c82d82-4587-4d24-87a2-e41e2911efec', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-08-08 13:53:09'),
	('81b5b7b6-a4b7-40b5-8762-b70c338e1d19', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '974dd35c-53fe-4a40-ae94-2f8ca4229c48', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:21:27'),
	('82a27008-3799-4e33-88e8-3e2eccecb123', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'PURCHASE', 5.00, 'PURCHASE', '4f3f3907-2536-4587-9d32-a2ae4751a9aa', 'purchase transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-06-08 16:19:49'),
	('884ede0d-7385-4224-81a5-9e9ffc23fe17', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', 'b5c9ed63-63f0-46c1-95e3-c77fb096049d', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:26:25'),
	('8df813f3-3e51-49c2-9396-c244c53ccc4c', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '7564b59d-35cd-4b52-afe4-deddd48456c9', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:24:41'),
	('9ed20b5d-9f8c-42e8-b36a-d853c75bfb68', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -11.00, 'SALE', '98dc6b96-260b-4b08-9bc3-6ef84025706d', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:59:47'),
	('a3ab7610-6093-4ea4-9a45-6c3330d33ad0', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -1.00, 'SALE', 'c0988bf7-cd16-49ce-a33d-9ecb54d506db', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:12:46'),
	('a47a1f1a-246b-408c-8402-b206118ce58d', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -3.00, '', '', 'Haji Sapari', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-21 09:25:59'),
	('a702ef07-ac1a-4765-9ccf-1dd300616435', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '6a208d10-09a5-4e0b-8abe-c293602f0abe', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 11:58:26'),
	('a718cc1a-59de-438d-b5db-a14a5b2a8d6f', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', '0f79d701-b78f-452e-a0cc-62c7d2ac35e8', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:20:06'),
	('ac65475f-ad0f-4294-8781-b0928654459d', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'PURCHASE', 13.00, 'PURCHASE', '6a208d10-09a5-4e0b-8abe-c293602f0fae', 'Hasil repack', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-21 09:14:27'),
	('b8d6362f-4cb1-47ff-b704-29609046b15f', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888882', 'SALE', -2.00, 'SALE', '4cf1c87a-0014-47e6-9b46-d3f4ec980377', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 06:08:58'),
	('bc8b4ad8-dc92-43b4-b09f-11e0094bbf39', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -13.00, 'SALE', 'abbeae5a-20a5-44d5-883c-065773b4e524', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 12:53:06'),
	('c82aec49-2d32-43a8-9f6a-fe4397fc9394', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', '0565b794-0d7f-4ce4-92d6-13aa5cd6007f', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:19:06'),
	('cccccccc-cccc-cccc-cccc-ccccccccccc1', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'OPENING', 100.00, 'SYSTEM', NULL, 'Stok awal', '2416854f-55e6-423c-a2ec-8154c9431cd6', '2026-05-18 09:24:30'),
	('cccccccc-cccc-cccc-cccc-ccccccccccc2', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -1.00, 'SALE', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Penjualan invoice INV-20260518-0001', '2416854f-55e6-423c-a2ec-8154c9431cd6', '2026-05-18 09:24:30'),
	('d0f869a8-ed0c-4c53-8798-f7a2098d3baa', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'REDUCE', -2.00, 'ADJUSTMENT', '', 'plastik terbuka', '64ee38a5-23c6-4a82-b1ad-d488ccc0d8e6', '2026-05-28 16:24:38'),
	('d426f189-813d-4501-804d-456710fa3543', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -20.00, '', '', 'Ibu Linda', '64ee38a5-23c6-4a82-b1ad-d488ccc0d8e6', '2026-06-02 08:25:23'),
	('d908d5e0-5ff5-486b-874b-47df2622b965', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888882', 'PURCHASE', 500.00, '', '', 'Stock Awal', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 06:08:47'),
	('da10c4ec-dd3b-44c9-88e7-85eeaa343dc1', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'PURCHASE', 10.00, 'PURCHASE', '4f3f3907-2536-4587-9d32-a2ae4751a9aa', 'purchase transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-06-08 16:19:49'),
	('db4093f9-c16e-4a85-afb4-3ba21828af5a', '11111111-1111-1111-1111-111111111111', '29454d2b-8cff-4e9a-9c47-6865dc7aa36f', 'INITIAL STOCK', 9.00, 'ITEM_VARIANT', '29454d2b-8cff-4e9a-9c47-6865dc7aa36f', 'Initial Stock', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-08-07 01:39:17'),
	('e0b28d2a-51f7-4d05-a243-11cea30dd8c2', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'PURCHASE', 99.00, 'PURCHASE', '6a208d10-09a5-4e0b-8abe-c293602f018h', 'Stock awal', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-21 09:02:13'),
	('e331dfb6-6a53-4eea-8869-3447d116b043', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'ADD', 10.00, 'ADJUSTMENT', '', 'barang masuk angin', '64ee38a5-23c6-4a82-b1ad-d488ccc0d8e6', '2026-05-28 16:23:07'),
	('e7bc843e-7178-4148-8ea0-417db75819f2', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'REDUCE', -1.00, 'ADJUSTMENT', '', 'plastik terbuka', '64ee38a5-23c6-4a82-b1ad-d488ccc0d8e6', '2026-05-29 01:56:02'),
	('eae61992-c31c-4b65-99f4-fbe568378f1f', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', 'b4a64b8a-c2ac-4c57-8a67-14e65058a7e4', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:08:45'),
	('f3c363d4-0425-44a0-a1ea-80475000f3b2', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -1.00, 'SALE', 'b4a64b8a-c2ac-4c57-8a67-14e65058a7e4', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:08:45'),
	('f7802881-51b8-4cd5-8e80-86aa73110fd8', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', '49e45d2b-7a13-4906-a292-b7d870bb1a8e', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:15:35'),
	('f8dea795-4049-415d-8e19-f0d0e8d7e45b', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'SALE', -2.00, 'SALE', 'bcec37b4-a8bb-44a2-a039-cbe8cb06a530', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-28 04:03:21'),
	('f90c5b7d-497f-4cd1-9385-4583a19b7aab', '11111111-1111-1111-1111-111111111111', '88888888-8888-8888-8888-888888888881', 'SALE', -5.00, 'SALE', '7564b59d-35cd-4b52-afe4-deddd48456c9', 'sale transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-05-26 13:24:41'),
	('fbdcb0c4-200e-45e0-aa79-0278b6caf7ac', '11111111-1111-1111-1111-111111111111', '76a3c5f8-4340-42a4-8aa0-db591a791c58', 'PURCHASE', 10.00, 'PURCHASE', 'd1b0e60f-b894-4320-b2e0-fae446d06e36', 'purchase transaction', '9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '2026-06-08 15:55:18');

-- Dumping structure for table umkm_odod.suppliers
CREATE TABLE IF NOT EXISTS `suppliers` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_supplier_name_per_tenant` (`tenant_id`,`name`),
  KEY `idx_suppliers_tenant` (`tenant_id`),
  KEY `idx_suppliers_name` (`name`),
  KEY `idx_suppliers_active` (`is_active`),
  CONSTRAINT `fk_suppliers_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.suppliers: ~3 rows (approximately)
INSERT INTO `suppliers` (`id`, `tenant_id`, `name`, `phone`, `address`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('4e0c76e8-e258-440d-9ed3-a8831f3c3296', '11111111-1111-1111-1111-111111111111', 'Lusi Oktaviani', '089981451237', 'Jl. Kalidam No. 45', 1, '2026-06-02 08:32:51', '2026-06-02 08:32:51', NULL),
	('4efba307-0fbd-4501-bb8d-3f9212c14fb9', '11111111-1111-1111-1111-111111111111', 'Keripik Ma Icih', '081248150369', 'Jl. Pasirkaliki No. 33', 1, '2026-06-02 10:04:20', '2026-06-02 10:04:20', NULL),
	('98f6686a-48b9-478e-b709-e8177a483bb8', '11111111-1111-1111-1111-111111111111', 'Amir Machmud', '081547841302', 'Jl. Cibaduyut No. 121', 1, '2026-06-02 11:10:00', '2026-06-02 11:25:59', '2026-06-02 11:25:59');

-- Dumping structure for table umkm_odod.tenants
CREATE TABLE IF NOT EXISTS `tenants` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
  `owner_name` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `address` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `logo` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `currency` char(3) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'IDR',
  `time_zone` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'Asia/Jakarta',
  `receipt_footer` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.tenants: ~4 rows (approximately)
INSERT INTO `tenants` (`id`, `name`, `owner_name`, `phone`, `email`, `address`, `logo`, `currency`, `time_zone`, `receipt_footer`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('11111111-1111-1111-1111-111111111111', 'Keripik Mang Odod', 'Ade Londok', '083182718171', 'sukmana@odod.co.id', 'Jl. Kademangan 96, Cimahi, Jawa Barat', '96bb4253-18a6-4ca9-b32a-fe924bb58f12.png', 'IDR', 'Asia/Jakarta', 'Hatur nuhun!', '2026-05-18 09:20:02', '2026-07-29 12:08:27', NULL),
	('11111111-1111-1111-1111-111111111112', 'Warung Bakso Petruk', 'Andri Kosasih', '081234567890', 'owner.baksopetruk@gmail.com', 'Jl. Permata Kopo No. 121, Kab. Bandung', NULL, 'IDR', 'Asia/Jakarta', 'Terima kasih telah berbelanja.\nBarang yang sudah dibeli tidak dapat ditukar atau dikembalikan.', '2026-05-18 09:57:40', '2026-07-29 01:57:41', NULL),
	('b83ed04b-f5b0-4109-a54d-de2916da7e0b', 'Aay Guevara', 'Hendro Kartiko', '081172638373', 'hendro@gmail.com', 'Jl. Kamarung no 12', NULL, 'IDR', 'Asia/Jakarta', 'Thank you for shopping.\r\n\r\nItems sold are not refundable.\r\n\r\nSee you again!', '2026-05-18 15:08:31', '2026-07-28 07:39:32', NULL),
	('f27e441f-5385-4b8d-b2e2-88b8615a4634', 'Bakso Solo', 'Evory Gunawan', '123456788', 'evory@gmail.com', 'Jl. Singgasana Pradana No.111', NULL, 'IDR', 'Asia/Jakarta', 'Terima kasih sudah berbelanja.\r\n\r\nItems sold are not refundable.\r\n\r\nSee you again!', '2026-05-18 10:29:53', '2026-07-28 07:39:34', NULL);

-- Dumping structure for table umkm_odod.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tenant_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `role_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `full_name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
  `username` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` char(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `phone` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `last_login_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_users_username_per_tenant` (`tenant_id`,`username`),
  KEY `fk_users_role` (`role_id`),
  CONSTRAINT `fk_users_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`),
  CONSTRAINT `fk_users_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table umkm_odod.users: ~21 rows (approximately)
INSERT INTO `users` (`id`, `tenant_id`, `role_id`, `full_name`, `username`, `password`, `phone`, `is_active`, `last_login_at`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('219f83ac-169c-40ed-8e2c-7a7795827102', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'andi', 'andi', '$2a$10$ezR6S1hCV87xgoMru8RqLOABeSHs710Lqo7DeLzvvePUL51GU2MWa', '089918272817', 0, NULL, '2026-07-03 03:35:10', '2026-07-03 03:35:10', NULL),
	('2416854f-55e6-423c-a2ec-8154c9431cd6', 'f27e441f-5385-4b8d-b2e2-88b8615a4634', 'dc064501-5798-4c91-8308-0198561ceae3', 'Budi Karmana', 'budi.admin', '$2a$10$NgbIGU78TzIZ.n2An5RIJe557VhXcHLzlQy1vOl1lVIUx4h4CVKnm', '087718273837', 1, NULL, '2026-05-20 06:35:42', '2026-05-20 06:35:42', NULL),
	('28dcfdbd-0532-45ad-a4ee-7204af4f0c51', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222223', 'dicky', 'dicky', '$2a$10$fOTwHIciA4kBTWnKi2H4mOXC2brXADhXhT/bDEDbyRrRsNXGPZdIC', '182738172', 1, NULL, '2026-07-03 06:42:37', '2026-07-03 11:40:43', '2026-07-03 11:40:44'),
	('375b98d3-9d07-483c-b9ea-7b2ce088ae48', 'f27e441f-5385-4b8d-b2e2-88b8615a4634', 'ab88f666-b2d6-482c-bdcd-be29915e5395', 'Kevin SA', 'kevin.sa', '$2a$10$ERZtRhd/MuxKn6pvMkFuS.8TfxCdyC4B8Zn4C4PYsOvspoPbCAzxW', '089945712017', 1, '2026-07-02 07:23:05', '2026-05-28 10:26:35', '2026-07-02 07:23:05', NULL),
	('64ee38a5-23c6-4a82-b1ad-d488ccc0d8e6', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222223', 'Kiki Sutisna', 'kiki.admin', '$2a$10$kN/7usrkKbIjHPPNZnLYCenuD7fLF3BKBbRTiLiIwLfP2yyVQfOVu', '081321114109', 1, '2026-07-18 12:49:05', '2026-05-28 09:39:41', '2026-07-18 12:49:05', NULL),
	('6a43e7fc-273f-4f77-a11b-d4318b40ac90', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'jamaludin', 'hantu', '$2a$10$2Xgh2W3kSoZygOE1vHuLK.q/QhRHnvv1DRGvkLAcqf.VyGlBWmDlK', '081192819181', 0, NULL, '2026-07-03 10:05:36', '2026-07-03 11:09:56', '2026-07-03 11:09:57'),
	('6b13d065-1766-422d-a30f-1762527e90ed', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'user12', 'user12', '$2a$10$.xAY0jlmqyUp8fuxNeZDcu/1uMor5hJjhHKBINKDrSxDBPB8uEovC', '123456789', 1, NULL, '2026-07-02 11:30:41', '2026-07-03 11:13:37', '2026-07-03 11:13:37'),
	('6eca3f5b-ddaa-419a-beab-e3a86e688071', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'heri', 'heri', '$2a$10$WB59lUmHI5Kkj3kUXod/a.G58bxdy5O.p0mwPZ03zXYDLH/9Oe/q.', '18273817', 1, NULL, '2026-07-03 06:43:53', '2026-07-03 11:13:51', '2026-07-03 11:13:52'),
	('750da97e-727e-4e54-a7b7-033ebd2b5f3c', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'tantri', 'tantri', '$2a$10$ohz0xlYI/1r/7QOdeEBlwe9ZWd.ilecNc3EySc7DD5AvktipWe6kC', '087718271817', 1, NULL, '2026-07-03 03:22:24', '2026-07-03 03:22:24', NULL),
	('76b97a14-fb57-4488-b6a6-755ed9e52706', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'user4', 'user4', '$2a$10$gchHxjBxOa.UL3yib0s5i.nOJLg.JvKAuIXjqcLnSSy9K0bKzc9W6', '123456789', 1, NULL, '2026-07-02 11:28:59', '2026-07-03 11:13:39', '2026-07-03 11:13:40'),
	('8ca6bd08-70ed-4b66-aea6-206f533824a5', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222223', 'Indra Hermawan', 'ihey', '$2a$10$cphwCrdJjf.jGUoASzIiV.MuU6XLgy6KtRP9Ct6MxLXKCtAdtZ2h6', '91829182928', 1, NULL, '2026-07-29 07:38:57', '2026-07-29 07:40:27', '2026-07-29 07:40:28'),
	('8deb396e-542e-41b8-9225-33a82f567642', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222223', 'Teguh Wijaya', 'akuntansi', '$2a$10$vpjNnq9OcVD4tzkXzrXXH./lFH23KGL7/.mqNc.XZMS5TznA0/tkS', '19282918', 1, NULL, '2026-07-29 07:23:31', '2026-07-29 07:23:31', NULL),
	('9b2ffb0e-c9f8-4fdf-94d6-71fffbc14014', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'Ibuk Bayik', 'ibuk', '$2a$10$zzxoebf5PHd5wFVTrYcyq.XWdoxVE2bOndrmDVw.ELwHZEDV0Mrju', '081192819182', 1, NULL, '2026-07-03 03:09:54', '2026-07-03 11:40:34', '2026-07-03 11:40:34'),
	('9d66b41e-e2d0-4074-bf40-881b156fb14d', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'bambang', 'bambang', '$2a$10$sjMJVaMEJK6L3Gvu61sgTukFjCqJi0GLYRx0hiEeExapB44lNkYh2', '18263821267', 0, NULL, '2026-07-03 06:40:03', '2026-07-03 06:40:03', NULL),
	('9dc150d4-ce87-4e5f-91c1-d0b5b7330ba7', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222221', 'Heru Wibowo', 'heru.owner', '$2a$10$TTp0I1S7i.fCgyRdJqrAEuDol/qsFEpCW9E7mrPooLq07ZgkNvc9e', '081345712017', 1, '2026-08-09 06:01:11', '2026-05-20 06:47:35', '2026-08-09 06:01:11', NULL),
	('a1a9b296-7286-4922-9f6c-c2531d61e8aa', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'heryan', 'heryan', '$2a$10$amomxq0K4RE5KDjgg5jEOeM7jnyXXj6AVVMIYkhjQPVJ0PYLLoF8S', '128123821', 1, '2026-07-18 12:49:23', '2026-07-03 07:27:30', '2026-07-18 12:49:23', NULL),
	('a4956ed7-8c4e-4728-9a31-42182819eb57', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222223', 'maman', 'maman', '$2a$10$EgcNIsJpa/jbWHjveB3QrexDOcjvk29AVELZia7ZU/EavGwwFUmR.', '+62112134555', 1, NULL, '2026-07-03 08:39:15', '2026-07-03 11:10:10', '2026-07-03 11:10:11'),
	('ab36bc0e-ba87-4f14-973e-63d0686e090e', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'kurweni ukar', 'karyana', '$2a$10$T24WBCm2D95U3UJZn8CZ7e0dT2jO1ogZUJF25qgif7DH1qCAk0lyu', '081192829181', 1, NULL, '2026-07-03 08:28:36', '2026-07-08 06:52:39', NULL),
	('d757bbc4-f4b5-4ae1-b040-47676f5e6038', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'mandra', 'mandra', '$2a$10$vtVhg4Nv5GdsP14beXs3aOku4Pl9nRUbcYdxeX0DoFXH7dQx6JXMm', '217282718', 1, NULL, '2026-07-03 07:23:57', '2026-07-08 06:53:07', '2026-07-08 06:53:07'),
	('dd750b80-03fe-4a67-acc5-fb62565bf75e', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222223', 'henri', 'henri', '$2a$10$Gd4/aPIqWOQiPZnInJ1xMe.T.Uc/oLSjesLj3LoC1CqqyCP14RLg2', '1728372918', 1, NULL, '2026-07-03 03:15:43', '2026-07-03 03:15:43', NULL),
	('e51a7024-6f91-486e-ac40-37e22130f115', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'indro', 'indro', '$2a$10$J4NRz5iIrYTqUQWYyo3MO.S6Iac6GtsMoQCe62JXyBgXvT3RhNqP2', '123456789', 1, NULL, '2026-07-03 03:30:08', '2026-07-03 03:30:08', NULL),
	('e5536d07-ebf1-42ce-aa90-df80cc6bd388', '11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'maman abdurahman', 'maman_a', '$2a$10$wlXiDR0OCDl.b23MfR9ftuVpMwFtiAdGyGYab3FoFeIFfw8zkqr4C', '0899182738171', 0, NULL, '2026-07-03 03:11:11', '2026-07-03 03:12:52', NULL);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
