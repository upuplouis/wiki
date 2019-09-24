CREATE DATABASE IF NOT EXISTS `grpc`;
USE `grpc`
DROP TABLE IF EXISTS `grpc_gateway`;
CREATE TABLE `grpc_gateway` (
    `id` int unsigned NOT NULL  AUTO_INCREMENT,
    `name` varchar(45) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;