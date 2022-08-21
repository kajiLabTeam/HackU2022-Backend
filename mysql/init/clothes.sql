DROP DATABASE IF EXISTS clothesdb;
CREATE DATABASE clothesdb character set utf8 collate utf8_general_ci;

use clothesdb;
DROP TABLE IF EXISTS user;
CREATE TABLE users (
    `id` VARCHAR(20),
    `name` VARCHAR(50),
    `gender` INT(1),
    `age` VARCHAR(20),
    `height` INT(4) UNSIGNED,
    `uuid` VARCHAR(40),
    `mail` VARCHAR(50),
    `icon` VARCHAR(150),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

DROP TABLE IF EXISTS likes;
CREATE TABLE likes (
    `id` VARCHAR(20),
    `coordinate_id` VARCHAR(20),
    `liked_user_id` VARCHAR(20),
    `user_id` VARCHAR(20),
    `lat` DOUBLE(9,6) DEFAULT NULL,
    `lng` DOUBLE(9,6) DEFAULT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

DROP TABLE IF EXISTS coordinate;
CREATE TABLE coordinates (
    `id` INT NOT NULL AUTO_INCREMENT,
    `coordinate_id` VARCHAR(20),
    `user_id` VARCHAR(20),
    `put_flag` INT(1),
    `public` INT(1),
    `image` VARCHAR(50),
    `category` VARCHAR(50),
    `brand` VARCHAR(50),
    `price` VARCHAR(50),
    `ble` VARCHAR(40),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
