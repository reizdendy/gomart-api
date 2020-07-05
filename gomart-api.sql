CREATE DATABASE  IF NOT EXISTS `gomart` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `gomart`;
-- MySQL dump 10.13  Distrib 8.0.20, for Win64 (x86_64)
--
-- Host: localhost    Database: gomart
-- ------------------------------------------------------
-- Server version	8.0.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `m_category`
--

DROP TABLE IF EXISTS `m_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_category` (
  `category_id` varchar(36) NOT NULL,
  `category_name` varchar(45) DEFAULT NULL,
  `status` char(1) DEFAULT 'A',
  PRIMARY KEY (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_category`
--

LOCK TABLES `m_category` WRITE;
/*!40000 ALTER TABLE `m_category` DISABLE KEYS */;
INSERT  IGNORE INTO `m_category` VALUES ('00389671-fd78-4435-86cd-3a5c6c6f2901','sepatu','A'),('3a5ecbe2-9e03-4ce1-9ba6-3491c9cb4b3c','sandal','A'),('400c8744-bc3a-11ea-b81f-dc4a3e5ea168','Kaos','A'),('ab51c1be-3d08-4960-8268-352d8020c5cb','sandal','I'),('c5eee472-bc39-11ea-b81f-dc4a3e5ea168','Jaket','I'),('d3ca6d6d-4769-45b1-8b5b-da363560bd1d','snack','I');
/*!40000 ALTER TABLE `m_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_product`
--

DROP TABLE IF EXISTS `m_product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_product` (
  `product_id` varchar(36) NOT NULL,
  `product_code` varchar(5) DEFAULT NULL,
  `product_name` varchar(45) DEFAULT NULL,
  `product_category` varchar(36) NOT NULL,
  `product_price` int DEFAULT NULL,
  `status` char(1) DEFAULT 'A',
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_product`
--

LOCK TABLES `m_product` WRITE;
/*!40000 ALTER TABLE `m_product` DISABLE KEYS */;
INSERT  IGNORE INTO `m_product` VALUES ('958f029a-27e9-444b-83bc-fbcc2d4d391b','snik','sepatu nike','00389671-fd78-4435-86cd-3a5c6c6f2901',399999,'I'),('a138b4a2-1379-498f-b3e6-ee7b51ff86de','sadi','sepatu adidas','00389671-fd78-4435-86cd-3a5c6c6f2901',500000,'A'),('ccea537e-b79a-480a-8fef-30713d3d503e','pr1','kaos enigma black','400c8744-bc3a-11ea-b81f-dc4a3e5ea168',59999,'A'),('d7e10b1a-8e8a-47f7-9c6c-27f15c73d575','pr2','kaos enigma white','400c8744-bc3a-11ea-b81f-dc4a3e5ea168',79999,'A');
/*!40000 ALTER TABLE `m_product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_transaction`
--

DROP TABLE IF EXISTS `m_transaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_transaction` (
  `id` varchar(36) NOT NULL,
  `transaction_date` datetime DEFAULT NULL,
  `product_id` varchar(36) DEFAULT NULL,
  `qty` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_product_trans_idx` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_transaction`
--

LOCK TABLES `m_transaction` WRITE;
/*!40000 ALTER TABLE `m_transaction` DISABLE KEYS */;
INSERT  IGNORE INTO `m_transaction` VALUES ('19c0e151-bc5d-11ea-b81f-dc4a3e5ea168','2020-07-02 19:20:22','ccea537e-b79a-480a-8fef-30713d3d503e',5),('9965594e-08b3-41e2-98e4-57c6b68cc484','2020-02-07 22:08:22','a138b4a2-1379-498f-b3e6-ee7b51ff86de',3),('aee50c12-bc48-11ea-b81f-dc4a3e5ea168','2020-07-02 00:00:00','d7e10b1a-8e8a-47f7-9c6c-27f15c73d575',2),('cd12aa1a-bc48-11ea-b81f-dc4a3e5ea168','2020-07-02 00:00:00','ccea537e-b79a-480a-8fef-30713d3d503e',3),('e0b4abb3-cd4f-4825-91b4-8cdff0fb89b4','2020-02-07 22:09:02','a138b4a2-1379-498f-b3e6-ee7b51ff86de',1);
/*!40000 ALTER TABLE `m_transaction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `v_report`
--

DROP TABLE IF EXISTS `v_report`;
/*!50001 DROP VIEW IF EXISTS `v_report`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `v_report` AS SELECT 
 1 AS `transaction_date`,
 1 AS `product_name`,
 1 AS `product_price`,
 1 AS `qty`,
 1 AS `total`*/;
SET character_set_client = @saved_cs_client;

--
-- Final view structure for view `v_report`
--

/*!50001 DROP VIEW IF EXISTS `v_report`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `v_report` AS select `t`.`transaction_date` AS `transaction_date`,`p`.`product_name` AS `product_name`,`p`.`product_price` AS `product_price`,`t`.`qty` AS `qty`,sum((`t`.`qty` * `p`.`product_price`)) AS `total` from (`m_transaction` `t` join `m_product` `p` on((`t`.`product_id` = `p`.`product_id`))) group by `p`.`product_id`,`t`.`transaction_date` order by `t`.`transaction_date` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-07-02 22:17:40
