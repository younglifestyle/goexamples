create database if not exists articles_demo_dev
  DEFAULT CHARACTER SET utf8
  DEFAULT COLLATE utf8_general_ci;
USE articles_demo_dev;
SET NAMES utf8;
SET character_set_server = utf8;

DROP TABLE if exists articles;
CREATE TABLE `articles` (
  `id` int NOT NULL,
  `title` VARCHAR(30) unsigned default 0,
  `body` VARCHAR(800) unsigned default 0,
  `created_on` int unsigned NOT NULL default 1,
  `updated_on` int unsigned NOT NULL default 1,
  PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
