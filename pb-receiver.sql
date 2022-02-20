-- MySQL dump 10.13  Distrib 8.0.28, for Linux (x86_64)
--
-- Host: mysql.db.gotechnology.io    Database: postback-receiver
-- ------------------------------------------------------
-- Server version	8.0.17

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `incoming_postback`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `incoming_postback` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `trackerid` int(11) DEFAULT NULL,
  `cnv_status` varchar(100) DEFAULT '',
  `payout` float DEFAULT '0',
  `currency` varchar(20) DEFAULT '',
  `url_query` json NOT NULL,
  `request_ip` varchar(50) DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `clickid` varchar(100) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `incoming_postback_tracker_id_fk` (`trackerid`),
  CONSTRAINT `incoming_postback_tracker_id_fk` FOREIGN KEY (`trackerid`) REFERENCES `tracker` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `send_postback`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `send_postback` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `incoming_postback_id` int(11) NOT NULL,
  `tracker_id` int(11) NOT NULL,
  `request_url` varchar(1500) DEFAULT NULL,
  `response_body` text,
  `response_code` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `postback_idx` (`incoming_postback_id`),
  KEY `tracker_idx` (`tracker_id`) /*!80000 INVISIBLE */,
  CONSTRAINT `postback` FOREIGN KEY (`incoming_postback_id`) REFERENCES `incoming_postback` (`id`),
  CONSTRAINT `tracker` FOREIGN KEY (`tracker_id`) REFERENCES `tracker` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `send_postback_failed`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `send_postback_failed` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `incoming_postback_id` int(11) NOT NULL,
  `tracker_id` int(11) NOT NULL,
  `request_url` varchar(1500) DEFAULT NULL,
  `response_body` text,
  `response_code` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_request_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `number_of_attempts` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `tracker_idx` (`incoming_postback_id`) /*!80000 INVISIBLE */,
  KEY `tracker_idx1` (`tracker_id`),
  CONSTRAINT `postback-failed` FOREIGN KEY (`incoming_postback_id`) REFERENCES `incoming_postback` (`id`),
  CONSTRAINT `tracker-failed` FOREIGN KEY (`tracker_id`) REFERENCES `tracker` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tracker`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tracker` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `is_active` tinyint(4) NOT NULL,
  `tracker_name` varchar(50) DEFAULT NULL,
  `postback_template` varchar(1500) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-02-15 16:46:46