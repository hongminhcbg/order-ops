DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `order_number` VARCHAR(128) NOT NULL UNIQUE,
  `customer_name` VARCHAR(255) NOT NULL,
  `quantiny` INT(10) NOT NULL,
  `phone` VARCHAR(32) NOT NULL,
  `address1` VARCHAR(100) DEFAULT '',
  `address2` VARCHAR(100)  DEFAULT '',
  `city` VARCHAR(32)  DEFAULT '',
  `state` VARCHAR(100)  DEFAULT '',
  `postal_code` VARCHAR(100) DEFAULT '',
  `country` VARCHAR(100) DEFAULT 'VN',
  `tracking_number` VARCHAR(255) DEFAULT '',
  `url` VARCHAR(255) DEFAULT '',
  `partner_tracking_number` VARCHAR(255) DEFAULT '',
  `status` TINYINT(3) DEFAULT 0,
  `note` VARCHAR(255) DEFAULT '',
  `begin_shipping` TIMESTAMP DEFAULT NOW(),
  `time_completed` TIMESTAMP DEFAULT NOW(),
  `created_at`                 DATETIME    DEFAULT NOW(),
  `updated_at`                 DATETIME    DEFAULT NOW() ON UPDATE NOW()
  );
  ALTER TABLE
    orders
    CHARACTER SET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;