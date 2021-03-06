-- MySQL dump 10.13  Distrib 5.7.10, for Win32 (AMD64)
--
-- Host: localhost    Database: hundred_dev
-- ------------------------------------------------------
-- Server version	5.7.10-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `audit`
--

DROP TABLE IF EXISTS `audit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `audit` (
  `id` int(18) NOT NULL AUTO_INCREMENT,
  `special` tinyint(2) DEFAULT NULL,
  `sso` int(11) DEFAULT NULL,
  `count` tinyint(4) DEFAULT NULL,
  `status` tinyint(2) DEFAULT NULL,
  `monad_id` int(11) DEFAULT NULL,
  `relational_id` int(11) DEFAULT NULL,
  `proposer_sso` int(11) DEFAULT NULL,
  `proposer_count` tinyint(4) DEFAULT NULL,
  `proposer_monad_id` int(11) DEFAULT NULL,
  `proposer_relational_id` int(11) DEFAULT NULL,
  `create` datetime DEFAULT NULL,
  `operation` datetime DEFAULT NULL,
  `isnewmonad` tinyint(2) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `audit`
--

LOCK TABLES `audit` WRITE;
/*!40000 ALTER TABLE `audit` DISABLE KEYS */;
INSERT INTO `audit` VALUES (6,1,3,0,0,0,0,4,1,2,2,'2016-05-08 10:13:15','2016-05-08 10:13:15',0),(7,1,3,0,0,0,0,6,1,3,3,'2016-05-08 23:55:13','2016-05-08 23:55:13',0),(9,1,3,0,0,0,0,7,1,4,4,'2016-05-08 23:55:32','2016-05-08 23:55:32',0),(10,0,8,1,0,5,5,7,0,9,4,'2016-05-08 23:55:35','2016-05-08 23:55:35',1),(11,0,8,1,0,5,5,4,0,10,2,'2016-05-09 20:55:43','2016-05-09 20:55:43',1),(12,0,9,1,0,6,6,13,0,11,9,'2016-05-09 21:36:16','2016-05-09 21:36:16',1);
/*!40000 ALTER TABLE `audit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message`
--

DROP TABLE IF EXISTS `message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `message` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `m_id` int(11) DEFAULT NULL,
  `r_id` int(11) DEFAULT NULL,
  `m_class` int(11) DEFAULT NULL,
  `status` tinyint(2) DEFAULT NULL,
  `typee` tinyint(2) DEFAULT NULL,
  `content` varchar(1000) DEFAULT NULL,
  `create` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message`
--

LOCK TABLES `message` WRITE;
/*!40000 ALTER TABLE `message` DISABLE KEYS */;
/*!40000 ALTER TABLE `message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `monad`
--

DROP TABLE IF EXISTS `monad`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `monad` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `is_main` tinyint(2) DEFAULT NULL,
  `main_monad` int(11) DEFAULT NULL,
  `state` tinyint(2) DEFAULT NULL,
  `count` smallint(4) DEFAULT NULL,
  `class` tinyint(2) DEFAULT NULL,
  `pertain` varchar(255) DEFAULT NULL,
  `parent_monad` int(11) DEFAULT NULL,
  `create` datetime DEFAULT NULL,
  `freeze` datetime DEFAULT NULL,
  `un_freeze` datetime DEFAULT NULL,
  `upgrade` datetime DEFAULT NULL,
  `audit` datetime DEFAULT NULL,
  `audit_number` varchar(255) DEFAULT NULL,
  `unfreeze_period_count` tinyint(2) DEFAULT '0',
  `task` tinyint(2) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `monad`
--

LOCK TABLES `monad` WRITE;
/*!40000 ALTER TABLE `monad` DISABLE KEYS */;
INSERT INTO `monad` VALUES (1,1,1,1,0,1,'1',0,'2016-05-08 10:06:55','0001-01-01 00:00:00','0001-01-01 00:00:00','2016-05-08 10:06:55','0001-01-01 00:00:00','',0,0),(2,1,1,1,2,1,'2',0,'2016-05-08 10:06:56','0001-01-01 00:00:00','0001-01-01 00:00:00','2016-05-08 10:06:56','0001-01-01 00:00:00','',0,1),(3,1,3,1,2,1,'3',2,'2016-05-08 10:11:54','0001-01-01 00:00:00','0001-01-01 00:00:00','2016-05-08 10:11:54','0001-01-01 00:00:00','',0,1),(4,1,4,1,2,1,'4',2,'2016-05-08 10:12:00','0001-01-01 00:00:00','0001-01-01 00:00:00','2016-05-08 10:12:00','0001-01-01 00:00:00','',0,1),(5,1,5,1,0,1,'5',3,'2016-05-08 10:12:06','0001-01-01 00:00:00','0001-01-01 00:00:00','2016-05-08 10:12:06','0001-01-01 00:00:00','',0,0),(6,1,6,1,0,1,'6',3,'2016-05-08 10:12:12','0001-01-01 00:00:00','0001-01-01 00:00:00','2016-05-08 10:12:12','0001-01-01 00:00:00','',0,0),(7,1,7,1,0,1,'7',4,'2016-05-08 10:12:24','0001-01-01 00:00:00','0001-01-01 00:00:00','2016-05-08 10:12:24','0001-01-01 00:00:00','',0,0),(8,0,3,1,0,1,'3',4,'2016-05-08 23:55:16','0001-01-01 00:00:00','0001-01-01 00:00:00','2016-05-08 23:55:16','0001-01-01 00:00:00','',0,0),(9,0,4,0,0,0,'4',5,'2016-05-08 23:55:35','0001-01-01 00:00:00','0001-01-01 00:00:00','2016-05-08 23:55:35','0001-01-01 00:00:00','',0,0),(10,0,2,0,0,0,'2',5,'2016-05-09 20:55:43','0001-01-01 00:00:00','0001-01-01 00:00:00','2016-05-09 20:55:43','0001-01-01 00:00:00','',0,0),(11,1,11,NULL,0,0,'9',6,'2016-05-09 21:36:15','0001-01-01 00:00:00','0001-01-01 00:00:00','2016-05-09 21:36:15','0001-01-01 00:00:00','',0,0);
/*!40000 ALTER TABLE `monad` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `relaadmin`
--

DROP TABLE IF EXISTS `relaadmin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `relaadmin` (
  `ssoid` int(11) NOT NULL,
  `relaid` int(11) NOT NULL DEFAULT '0',
  `income` int(11) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `relaadmin`
--

LOCK TABLES `relaadmin` WRITE;
/*!40000 ALTER TABLE `relaadmin` DISABLE KEYS */;
INSERT INTO `relaadmin` VALUES (3,0,0),(2,1,0),(2,2,0);
/*!40000 ALTER TABLE `relaadmin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `relational`
--

DROP TABLE IF EXISTS `relational`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `relational` (
  `id` int(18) NOT NULL AUTO_INCREMENT,
  `current_monad` int(11) DEFAULT NULL,
  `recommand_total` int(11) DEFAULT NULL,
  `role` tinyint(4) DEFAULT NULL,
  `mobile` varchar(11) DEFAULT NULL,
  `referrer` varchar(255) DEFAULT NULL,
  `sso_id` int(11) DEFAULT NULL,
  `prev_info` varchar(255) DEFAULT NULL,
  `history_track` varchar(255) DEFAULT NULL,
  `history_monads` varchar(255) DEFAULT NULL,
  `freetime` varchar(255) DEFAULT NULL,
  `create` datetime DEFAULT NULL,
  `prev` datetime DEFAULT NULL,
  `income` int(6) DEFAULT NULL,
  `spending` int(6) DEFAULT NULL,
  `loss` int(6) DEFAULT NULL,
  `status` tinyint(2) DEFAULT '1',
  `prev_new_monad` datetime DEFAULT NULL,
  `one_count` smallint(5) DEFAULT NULL,
  `two_count` smallint(5) DEFAULT NULL,
  `three_count` smallint(5) DEFAULT NULL,
  `monad_count` smallint(5) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `relational`
--

LOCK TABLES `relational` WRITE;
/*!40000 ALTER TABLE `relational` DISABLE KEYS */;
INSERT INTO `relational` VALUES (1,1,0,0,'18650710067','top',5,'','','','','2016-05-08 10:06:55','2016-05-08 10:06:55',0,0,0,1,'2016-05-08 10:06:55',0,0,0,NULL),(2,2,0,0,'18059244379','top',4,'','','','','2016-05-08 10:06:56','2016-05-08 10:06:56',200,0,0,1,'2016-05-08 12:06:56',0,0,0,NULL),(3,3,0,0,'18011110001','2',6,'','','','','2016-05-08 10:10:26','2016-05-08 23:55:19',200,100,0,1,'2016-05-08 12:11:54',0,0,0,NULL),(4,4,0,0,'18011110002','2',7,'','','','','2016-05-08 10:10:45','2016-05-08 23:55:35',200,0,0,1,'2016-05-08 12:12:00',0,0,0,NULL),(5,5,0,0,'18011110003','2',8,'','','','','2016-05-08 10:10:57','2016-05-08 10:10:57',0,0,0,1,'2016-05-08 10:12:06',0,0,0,NULL),(6,6,0,0,'18011110004','2',9,'','','','','2016-05-08 10:11:10','2016-05-08 10:11:10',0,0,0,1,'2016-05-08 10:12:12',0,0,0,NULL),(7,7,0,0,'18011110005','2',10,'','','','','2016-05-08 10:11:33','2016-05-08 10:11:33',0,0,0,1,'2016-05-08 10:12:24',0,0,0,NULL),(8,0,0,0,'18011110006','2',11,'','','','','2016-05-08 10:13:08','2016-05-08 10:13:08',0,0,0,0,'0001-01-01 00:00:00',0,0,0,NULL),(9,11,0,0,'18022220000','2',13,'','','','','2016-05-09 21:36:05','2016-05-09 21:36:05',0,0,0,0,'2016-05-09 21:36:16',0,0,0,NULL);
/*!40000 ALTER TABLE `relational` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `state`
--

DROP TABLE IF EXISTS `state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `state` (
  `token` varchar(255) DEFAULT NULL,
  `userjson` varchar(255) DEFAULT NULL,
  `overdue` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `state`
--

LOCK TABLES `state` WRITE;
/*!40000 ALTER TABLE `state` DISABLE KEYS */;
INSERT INTO `state` VALUES ('b0a8e9f8','23','2016-02-05 11:51:19'),('9347fa28','3','2016-02-06 08:21:42'),('8f9435ce','3','2016-02-06 08:21:44'),('c23f48bd','3','2016-02-06 10:45:18'),('879a25b0','3','2016-02-06 11:10:53'),('d35eea68','3','2016-02-06 14:20:31'),('b9e2ae77','3','2016-02-06 14:20:39'),('f4fe7242','3','2016-02-06 14:20:39'),('aede48f8','3','2016-02-06 16:07:20'),('95f96702','3','2016-02-06 17:00:05'),('959fa6b8','3','2016-02-06 17:00:06'),('f71dc9b0','3','2016-02-06 17:00:10'),('650ce0c3','3','2016-02-06 17:18:28'),('49b8743c','3','2016-02-06 17:18:35'),('ecbaa181','3','2016-02-06 17:18:38'),('ee0841bc','3','2016-02-06 17:18:45'),('40a62a28','3','2016-02-06 17:18:59'),('c9f3c301','41','2016-02-06 18:13:21'),('d713febd','41','2016-02-06 18:14:23'),('ed32be1c','41','2016-02-06 18:17:12'),('aec7a00e','41','2016-02-06 18:18:47'),('d785012a','41','2016-02-06 18:18:52'),('ff8e5064','41','2016-02-06 18:18:58'),('cb4f706e','45','2016-02-06 18:31:57'),('c503bc2e','16','2016-02-07 12:21:11'),('46544ce2','16','2016-02-07 12:32:01'),('4da883d3','16','2016-02-07 12:43:18'),('f7c33456','16','2016-02-07 19:58:47'),('00d86353','24','2016-02-12 12:15:41'),('0600f936','14','2016-02-14 21:38:59'),('d2ed3d19','32','2016-02-15 00:47:49'),('67526ca5','32','2016-02-15 06:17:46'),('742a4e52','30','2016-02-15 07:25:58'),('6e9c9b58','50','2016-02-15 12:49:59'),('e71d9111','50','2016-02-15 14:23:18'),('c82246ef','76','2016-02-16 00:41:38'),('3ae5eedd','76','2016-02-16 00:41:56'),('774c9130','76','2016-02-16 00:41:58'),('e0a35b71','76','2016-02-16 00:41:59'),('42b2f08a','61','2016-02-16 01:05:47'),('c7e37c54','61','2016-02-16 01:07:33'),('026fb10b','61','2016-02-16 01:07:37'),('f8b38015','61','2016-02-16 01:07:45'),('0d7c8c30','80','2016-02-16 02:36:46'),('be7f5e95','61','2016-02-16 03:05:50'),('5e091294','80','2016-02-16 03:05:57'),('1609ba34','80','2016-02-16 03:05:58'),('39168792','80','2016-02-16 03:05:58'),('a24a8f87','80','2016-02-16 03:06:27'),('4c2d520a','61','2016-02-16 03:41:24'),('3c02feda','61','2016-02-16 03:51:11'),('bd4a2698','85','2016-02-16 04:04:04'),('1d89cd80','85','2016-02-16 04:15:13'),('02b1e3e0','85','2016-02-16 04:16:39'),('a116c76d','85','2016-02-16 04:17:06'),('10ccfb19','85','2016-02-16 04:18:06'),('dd76599f','85','2016-02-16 04:18:19'),('7f08dfad','78','2016-02-16 05:29:22'),('cca55336','59','2016-02-16 05:35:12'),('c16b6e0d','75','2016-02-16 05:58:52'),('4b3c9a1b','75','2016-02-16 05:58:52'),('7e416bf0','75','2016-02-16 05:58:52'),('63a9c59d','75','2016-02-16 05:58:53'),('fe1790fa','75','2016-02-16 05:58:53'),('ba4d0213','75','2016-02-16 05:58:53'),('0d5c1b7d','75','2016-02-16 05:58:54'),('287fa7ac','75','2016-02-16 05:58:54'),('3d7838ad','75','2016-02-16 05:58:54'),('2d63af51','75','2016-02-16 05:58:54'),('359ad4d5','75','2016-02-16 05:58:54'),('99d8e04d','75','2016-02-16 05:58:54'),('cae02247','75','2016-02-16 05:58:54'),('ebf8066d','75','2016-02-16 05:58:54'),('14909351','75','2016-02-16 06:03:55'),('c28b0385','92','2016-02-16 06:24:30'),('442b370d','75','2016-02-16 07:01:49'),('dbeedf6d','94','2016-02-16 07:31:28'),('1a86d5b5','61','2016-02-16 07:59:28'),('68c550a7','94','2016-02-16 08:08:51'),('bdf2a29f','89','2016-02-16 08:31:59'),('0b3fe458','78','2016-02-16 08:55:33'),('16b5e804','89','2016-02-16 08:57:20'),('9065f233','98','2016-02-16 09:30:26'),('5c3438af','98','2016-02-16 09:30:26'),('d7b08bd5','98','2016-02-16 09:30:51'),('0506105f','98','2016-02-16 09:30:51'),('5c7ade01','100','2016-02-16 10:46:37'),('76b4822e','100','2016-02-16 11:00:45'),('4c68de94','56','2016-02-17 03:01:38'),('96a0ff17','105','2016-02-17 05:08:43'),('5b6947d5','105','2016-02-17 05:09:26'),('15c8351b','107','2016-02-17 06:20:51'),('32287343','107','2016-02-17 06:21:19'),('99623bf3','107','2016-02-17 06:24:39'),('de8cd81b','107','2016-02-17 06:25:00'),('ee7f2989','107','2016-02-17 06:25:01'),('5f4f8a0d','107','2016-02-17 06:25:03'),('6e2d93a7','107','2016-02-17 06:25:29'),('c65c8a58','105','2016-02-17 07:17:07'),('1a797fa4','51','2016-02-17 07:54:26'),('830f463f','101','2016-02-17 10:46:16'),('f594dc61','63','2016-02-17 11:21:17'),('f6bfab71','63','2016-02-17 11:26:06'),('9f1dfd97','63','2016-02-17 11:26:15'),('e1a83242','63','2016-02-17 11:31:05'),('c9fbd1cc','63','2016-02-17 11:31:20'),('e4ac8dd4','42','2016-02-17 11:46:44'),('45ed3bd7','58','2016-02-17 21:01:27'),('743e196d','58','2016-02-17 21:04:18'),('47e72993','58','2016-02-17 21:04:20'),('4d002a34','58','2016-02-17 21:04:21'),('2696bf8e','58','2016-02-17 21:07:55'),('eeca7a22','66','2016-02-17 21:27:09'),('8ee85f6b','120','2016-02-17 21:46:26'),('1d5156bb','58','2016-02-17 22:04:47'),('da20d8fb','124','2016-02-17 22:45:34'),('ca9b2f3d','118','2016-02-17 22:49:06'),('07680fd5','69','2016-02-18 00:08:31'),('335888ce','73','2016-02-18 00:56:16'),('c5d287b0','81','2016-02-18 01:14:36'),('7ea67ca6','81','2016-02-18 01:19:41'),('84fd272d','73','2016-02-18 01:20:09'),('15f94d6f','81','2016-02-18 01:22:48'),('5fd51e44','81','2016-02-18 01:33:21'),('a936d4a7','81','2016-02-18 01:33:22'),('299631f4','81','2016-02-18 01:33:31'),('eead3418','73','2016-02-18 01:36:51'),('6e181591','73','2016-02-18 01:36:52'),('98c73520','131','2016-02-18 01:48:48'),('429c38b3','131','2016-02-18 01:54:45'),('037505dc','73','2016-02-18 02:06:11'),('38719ea7','131','2016-02-18 02:16:12'),('31148f96','135','2016-02-18 03:18:02'),('bfccded2','120','2016-02-18 03:41:41'),('3abe62e3','118','2016-02-18 04:12:00'),('105e626e','118','2016-02-18 04:18:57'),('95a9f064','124','2016-02-18 04:32:32'),('d44e5cdb','124','2016-02-18 04:33:06'),('7207a9b8','81','2016-02-18 05:19:09'),('86b2d649','81','2016-02-18 05:19:11'),('539e15f2','97','2016-02-18 05:28:39'),('72c32340','81','2016-02-18 06:00:29'),('5572085e','73','2016-02-18 06:08:24'),('48e58a22','73','2016-02-18 06:16:20'),('17c258ed','73','2016-02-18 06:23:18'),('aefe9c11','73','2016-02-18 06:23:57'),('e71596ca','73','2016-02-18 06:29:05'),('810bb47e','73','2016-02-18 06:29:07'),('ce40fb66','73','2016-02-18 06:29:08'),('14aa06eb','73','2016-02-18 06:29:09'),('b2597be9','73','2016-02-18 06:29:10'),('e458cc07','73','2016-02-18 06:29:10'),('5e51e8bb','81','2016-02-18 07:59:27'),('7d05c0fd','148','2016-02-18 10:27:21'),('82fae51c','150','2016-02-18 11:21:20'),('dfdbc0fe','150','2016-02-18 11:21:20'),('c1a4ea54','136','2016-02-18 11:36:41'),('a6b5395d','144','2016-02-18 11:47:16'),('7900f479','151','2016-02-18 11:51:48'),('59cd8a90','144','2016-02-18 11:55:07'),('e87954c4','144','2016-02-18 11:55:25'),('b1a6fce0','144','2016-02-18 11:55:27'),('eab2a4af','127','2016-02-18 21:44:49'),('c9f98523','157','2016-02-18 22:53:19'),('3528e65d','156','2016-02-18 22:57:19'),('b3450ce2','157','2016-02-18 22:58:46'),('aa0a1267','157','2016-02-18 23:07:22'),('b2e0634c','157','2016-02-18 23:22:26'),('60c6282b','157','2016-02-18 23:36:14'),('df964ae6','162','2016-02-19 00:09:28'),('aa18d14a','159','2016-02-19 00:14:38'),('547b3a81','48','2016-02-19 00:39:24'),('8fffef4d','103','2016-02-19 02:08:52'),('543b4d5e','163','2016-02-19 02:27:05'),('3ae7f974','113','2016-02-19 02:31:41'),('1a4046a2','88','2016-02-19 02:56:12'),('3eea48d4','132','2016-02-19 04:51:08'),('e48f6a8d','48','2016-02-19 05:00:38'),('1252603d','169','2016-02-19 05:00:58'),('1a9392af','169','2016-02-19 05:01:04'),('6645fe38','48','2016-02-19 05:15:04'),('0c059f25','171','2016-02-19 05:36:19'),('6c5bcda1','171','2016-02-19 05:39:47'),('bfed90aa','171','2016-02-19 05:39:54'),('943b8393','171','2016-02-19 05:39:55'),('59add575','171','2016-02-19 05:39:56'),('e8c43408','171','2016-02-19 05:41:51'),('16510cad','171','2016-02-19 05:41:53'),('ec56cdd6','171','2016-02-19 05:41:57'),('f9ade1e3','48','2016-02-19 06:34:02'),('8e375450','174','2016-02-19 07:37:32'),('be9a0b12','174','2016-02-19 07:41:55'),('162773aa','178','2016-02-19 08:26:46'),('cdeb5f5e','178','2016-02-19 08:27:06'),('d06f193c','178','2016-02-19 08:27:08'),('6f1fa1df','178','2016-02-19 08:27:09'),('057e0367','178','2016-02-19 08:27:09'),('e421d293','178','2016-02-19 08:27:10'),('9a1eff69','178','2016-02-19 08:28:17'),('dec8ba62','178','2016-02-19 08:28:18'),('d75701f8','178','2016-02-19 08:28:20'),('a036673d','178','2016-02-19 08:28:21'),('3c7b500e','178','2016-02-19 08:28:22'),('d28524ac','178','2016-02-19 08:28:24'),('31bfcf39','178','2016-02-19 08:28:33'),('a1acf11d','178','2016-02-19 08:28:35'),('b9f268bb','178','2016-02-19 08:28:35'),('70954e6a','178','2016-02-19 08:28:36'),('7f8799be','178','2016-02-19 08:28:46'),('dfa9513f','178','2016-02-19 08:29:09'),('932d2de0','178','2016-02-19 08:36:31'),('cde9471f','114','2016-02-19 11:00:25'),('770589db','113','2016-02-19 11:09:00'),('3ed65b6f','181','2016-02-19 11:41:39'),('512f433c','181','2016-02-19 11:57:46'),('d01cb5e0','70','2016-02-19 19:59:54'),('74837807','133','2016-02-19 23:37:26'),('5ecaaebe','185','2016-02-20 01:10:44'),('14b18db5','72','2016-02-20 07:24:36'),('9c2a7b7f','0','2016-02-20 20:47:11'),('84411414','0','2016-02-20 21:00:30'),('c7b166e6','0','2016-02-20 21:00:31'),('6bdfec69','191','2016-02-21 08:52:04'),('2cf4a5a7','200','2016-02-21 09:39:39'),('dd2e3bb7','165','2016-02-21 13:21:12'),('3cd216b5','191','2016-02-21 20:45:13'),('c736a04d','203','2016-02-22 05:19:02'),('639d54e2','203','2016-02-22 05:19:08'),('c3d4b4c2','203','2016-02-22 05:19:17'),('19c8e28b','203','2016-02-22 05:19:32'),('72e2f3b7','203','2016-02-22 05:19:33'),('4941ee25','203','2016-02-22 05:19:46'),('c2602d87','203','2016-02-22 05:19:49'),('926d184f','203','2016-02-22 05:20:55'),('1e416aed','203','2016-02-22 05:20:59'),('65574f8c','115','2016-02-22 09:57:02'),('ae12815a','194','2016-02-22 21:49:00'),('612b0075','194','2016-02-22 21:49:10'),('1bba1b9f','194','2016-02-22 21:49:12'),('9b157e96','194','2016-02-22 21:49:14'),('2aa79130','194','2016-02-22 21:55:41'),('119131ad','194','2016-02-22 22:47:48'),('d1dc6929','204','2016-02-23 03:32:25'),('09ee6bc6','204','2016-02-23 03:43:02'),('4c8c4c9d','204','2016-02-23 03:46:58'),('4e25ab78','204','2016-02-23 04:04:16'),('3750503b','199','2016-02-23 08:37:53'),('53c232f1','199','2016-02-23 08:44:32'),('d4b8ea7d','188','2016-02-24 21:25:38'),('f1eee3b4','190','2016-02-24 21:32:12'),('a0271275','125','2016-02-25 06:38:51'),('9f855dd2','167','2016-02-26 05:57:52'),('20f44275','167','2016-02-26 08:10:07'),('3dba1764','167','2016-02-26 08:11:23'),('d603eae7','167','2016-02-26 08:12:42'),('475332ab','167','2016-02-26 08:14:28'),('f2348b7b','208','2016-02-27 01:44:44'),('e3b8ee06','208','2016-02-27 01:45:24'),('5345bd1b','208','2016-02-27 01:45:26'),('31eba06b','208','2016-02-27 01:45:30'),('98bff900','208','2016-02-27 01:45:50'),('46cffd2e','208','2016-02-27 01:47:03'),('3493b7b6','208','2016-02-27 01:47:20'),('8b958405','208','2016-02-27 01:47:25'),('2ac2846c','208','2016-02-27 01:47:39'),('641317bc','208','2016-02-27 01:47:52'),('2da51ff6','208','2016-02-27 01:47:54'),('7fe9fbd4','208','2016-02-27 01:47:56'),('4977abf6','193','2016-02-28 04:37:03'),('40d1e07a','214','2016-02-28 04:47:24'),('a21d144f','214','2016-02-28 04:47:42'),('fe251bc8','214','2016-02-28 04:47:54'),('0743ca55','192','2016-02-29 02:13:36'),('e02b09fb','192','2016-02-29 02:13:38'),('fe6f5baf','164','2016-03-01 08:27:09'),('69df53ea','221','2016-03-01 08:42:52'),('bc1140db','222','2016-03-01 08:50:21'),('2d7601db','202','2016-03-01 10:35:07'),('c2cf0e37','219','2016-03-01 11:33:25'),('50245221','219','2016-03-01 11:33:27'),('d5f5ae1a','219','2016-03-01 12:08:45'),('0d80ef1e','219','2016-03-01 12:08:47'),('4e48b22a','219','2016-03-01 18:32:12'),('fca1ca42','212','2016-03-01 19:52:33'),('65606b8b','170','2016-03-01 23:29:11'),('6874e1c3','170','2016-03-01 23:30:23'),('a77c6ca3','166','2016-03-02 05:20:23'),('7e226297','186','2016-03-02 06:18:47'),('9590d24d','186','2016-03-02 06:20:56'),('22ab7477','186','2016-03-02 06:20:59'),('215adda6','186','2016-03-02 06:21:11'),('f497a9a7','186','2016-03-02 06:22:05'),('5f80a9ba','186','2016-03-02 06:31:21'),('674adc4d','186','2016-03-02 06:31:47'),('3f1e9430','186','2016-03-02 06:31:47'),('5df044aa','186','2016-03-02 06:31:47'),('7cd67905','186','2016-03-02 06:32:17'),('d11ddfda','186','2016-03-02 06:36:35'),('55acb16c','186','2016-03-02 06:39:01'),('b57ce79f','186','2016-03-02 06:44:41'),('87d7b763','186','2016-03-02 06:44:58'),('78c2f11a','186','2016-03-02 06:46:27'),('89ea58fe','186','2016-03-02 06:49:12'),('64e757a3','186','2016-03-02 06:50:57'),('e97531a9','186','2016-03-02 07:00:00'),('678abf03','186','2016-03-02 07:00:17'),('2e2dad50','230','2016-03-02 07:05:39'),('27378444','230','2016-03-02 07:14:20'),('be301b8d','230','2016-03-02 07:14:41'),('d682d59a','34','2016-03-02 07:44:36'),('a11b1932','186','2016-03-02 07:48:05'),('002c946d','186','2016-03-02 07:52:32'),('3cae7930','33','2016-03-02 06:50:25'),('409a9afa','227','2016-03-02 17:17:27'),('3a66d951','227','2016-03-02 17:17:50'),('decbb100','227','2016-03-02 17:19:00'),('480a9b0e','227','2016-03-02 17:29:47'),('f8af4fc3','227','2016-03-02 17:31:15'),('8c188ac7','227','2016-03-02 17:31:55'),('99be07de','248','2016-03-02 18:27:08'),('507bb9a7','247','2016-03-02 18:31:38'),('973431a7','138','2016-03-02 18:39:01'),('6bf1091f','138','2016-03-02 18:39:14'),('26c57d15','234','2016-03-02 18:43:45'),('256f5fa6','234','2016-03-02 18:44:00'),('4afbf38c','234','2016-03-02 18:44:02'),('83634545','195','2016-03-02 22:25:28'),('3c62c92b','195','2016-03-02 22:29:11'),('44750381','207','2016-03-02 22:32:01'),('0d144cb1','207','2016-03-02 22:33:28'),('192a547e','207','2016-03-02 22:34:21'),('02ecbc8b','206','2016-03-02 23:39:03'),('831cb145','206','2016-03-02 23:39:39'),('8a3ef2ba','206','2016-03-02 23:39:40'),('3b6c84d6','38','2016-03-02 23:40:59'),('7fd2eded','184','2016-03-03 01:06:31'),('04c94a7d','184','2016-03-03 01:06:36'),('fb8017c5','184','2016-03-03 01:09:12'),('0f24fd2d','184','2016-03-03 01:10:22'),('d57370bb','13','2016-03-03 01:34:29'),('234308ac','26','2016-03-03 02:35:20'),('e69c2f32','26','2016-03-03 02:35:20'),('dc9fe738','26','2016-03-03 02:45:49'),('034afb71','254','2016-03-03 02:51:52'),('f3f1baa0','26','2016-03-03 03:03:20'),('6c891ba1','176','2016-03-03 19:02:14'),('1798e4f8','259','2016-03-03 19:13:44'),('8e64f0f8','259','2016-03-03 19:15:18'),('71f69f64','259','2016-03-03 19:16:25'),('60c4112f','259','2016-03-03 19:18:08'),('ca199599','259','2016-03-03 19:18:20'),('e9f89464','259','2016-03-03 19:21:56'),('aad062f8','54','2016-03-03 19:53:48'),('2b8357c9','54','2016-03-03 19:55:04'),('732a8c47','54','2016-03-03 19:55:32'),('0b283ac7','54','2016-03-03 19:55:37'),('e196d963','54','2016-03-03 19:55:50'),('f9b17747','196','2016-03-03 21:10:27'),('75b492f5','196','2016-03-03 21:17:36'),('987a0e48','196','2016-03-03 21:19:39'),('e21240b3','196','2016-03-03 21:30:54'),('6c7dcf64','196','2016-03-03 21:35:59'),('ea4de7b6','196','2016-03-03 21:37:13'),('dadc94bf','196','2016-03-03 21:38:20'),('e82b9fc9','196','2016-03-03 21:40:11'),('f50a935d','47','2016-03-03 22:14:08'),('4b89668b','231','2016-03-03 22:24:20'),('353f6c52','231','2016-03-03 22:25:13'),('c2204f41','231','2016-03-03 22:26:43'),('0a629f00','231','2016-03-03 22:28:30'),('3f332543','198','2016-03-03 23:03:19'),('d3c58c4c','198','2016-03-03 23:03:45'),('d84f0548','270','2016-03-03 23:09:03'),('cd16365f','270','2016-03-03 23:09:13'),('7e958015','196','2016-03-04 01:02:28'),('c7d37f86','196','2016-03-04 01:02:30'),('200864c1','224','2016-03-04 01:57:55'),('01d32d60','224','2016-03-04 01:58:17'),('e9e4d95b','228','2016-03-04 02:31:40'),('f8c7a5dc','228','2016-03-04 02:31:43'),('ce3cbbbb','29','2016-03-04 04:45:19'),('89f97aba','29','2016-03-04 04:45:32'),('3a8224c6','29','2016-03-04 04:45:38'),('a3ed7b46','29','2016-03-04 04:45:40'),('f2a8cc2f','31','2016-03-04 04:46:14'),('1728fa99','31','2016-03-04 04:46:15'),('7256ec51','31','2016-03-04 04:46:16'),('a8013a24','31','2016-03-04 04:46:17'),('41227ce5','31','2016-03-04 04:46:18'),('1380f8ee','15','2016-03-04 04:51:39'),('2f362430','15','2016-03-04 04:51:42'),('c484b38a','15','2016-03-04 04:51:42'),('5d4c7193','29','2016-03-04 04:53:38'),('eefd5044','29','2016-03-04 04:53:41'),('a785cdb1','39','2016-03-04 04:54:04'),('7aa5b882','29','2016-03-04 04:55:40'),('c66376fc','29','2016-03-04 04:55:42'),('cb23d4af','29','2016-03-04 04:55:43'),('f7f2088b','29','2016-03-04 04:55:43'),('a4e3d84f','29','2016-03-04 04:55:44'),('d10b57f0','29','2016-03-04 04:55:44'),('8b77f49a','29','2016-03-04 04:55:44'),('4a549884','29','2016-03-04 04:55:44'),('9b3c0df8','29','2016-03-04 04:55:44'),('143ded67','29','2016-03-04 04:55:45'),('85c94813','29','2016-03-04 04:55:46'),('157d617d','29','2016-03-04 04:55:47'),('14e0f9c5','29','2016-03-04 04:55:47'),('9d6b9dc6','29','2016-03-04 04:55:47'),('f8aa2980','29','2016-03-04 04:55:47'),('52ed2579','29','2016-03-04 04:55:48'),('8ea478d9','29','2016-03-04 04:55:48'),('da1c5ba6','29','2016-03-04 04:55:48'),('0a9a6844','29','2016-03-04 04:55:48'),('4345d247','29','2016-03-04 04:55:49'),('f2449e58','29','2016-03-04 04:55:49'),('f544a9e6','29','2016-03-04 04:55:49'),('aa02dd2c','29','2016-03-04 04:55:49'),('32412724','29','2016-03-04 04:55:49'),('94061688','44','2016-03-04 05:44:36'),('73bad337','29','2016-03-04 06:23:13'),('6ed7423c','31','2016-03-04 06:23:51'),('9761f017','249','2016-03-04 14:45:53'),('7adfda73','5','2016-03-04 14:48:02'),('1b58a761','5','2016-03-04 14:48:07'),('29e2a7bb','117','2016-03-04 16:59:23'),('db70b698','276','2016-03-04 17:11:45'),('951ab4c0','276','2016-03-04 17:17:49'),('48937322','276','2016-03-04 17:18:28'),('d2dedf15','276','2016-03-04 17:20:07'),('2c128a84','276','2016-03-04 17:20:27'),('089286bc','276','2016-03-04 17:20:30'),('1a1d9552','276','2016-03-04 17:20:38'),('09cbeb2b','276','2016-03-04 17:21:01'),('b3824995','276','2016-03-04 17:21:04'),('d3e074ce','276','2016-03-04 17:21:13'),('953354ee','276','2016-03-04 17:43:38'),('0ca430ca','19','2016-03-04 17:48:35'),('97fb96be','233','2016-03-04 17:50:38'),('2905f33e','233','2016-03-04 17:51:38'),('1a97da84','233','2016-03-04 17:51:54'),('9d0c9310','233','2016-03-04 17:52:05'),('4446bdcc','233','2016-03-04 17:52:17'),('661bbc48','233','2016-03-04 17:52:38'),('6d13520b','277','2016-03-04 17:57:38'),('3a3a9a9a','238','2016-03-04 18:03:49'),('9a9f23de','238','2016-03-04 18:04:03'),('2c9b8af3','238','2016-03-04 18:04:09'),('4e743736','238','2016-03-04 18:04:16'),('33038e65','238','2016-03-04 18:04:22'),('74d05d84','238','2016-03-04 18:04:30'),('758e9f8e','211','2016-03-04 18:08:13'),('23619300','211','2016-03-04 18:08:20'),('7ff14d2b','5','2016-03-04 18:12:08'),('4262a5dd','277','2016-03-04 18:17:54'),('8b5eef9c','277','2016-03-04 18:22:15'),('6b0730eb','277','2016-03-04 18:22:59'),('d9c631c4','277','2016-03-04 18:23:03'),('8a402d06','277','2016-03-04 18:23:10'),('649617fe','277','2016-03-04 18:23:12'),('4b619d75','277','2016-03-04 18:24:12'),('3124be08','277','2016-03-04 18:25:31'),('02dd9253','236','2016-03-04 18:26:14'),('1d615b54','236','2016-03-04 18:28:53'),('ac2df362','236','2016-03-04 18:29:03'),('d9b4d387','236','2016-03-04 18:29:15'),('de970c01','236','2016-03-04 18:29:23'),('762db4d6','236','2016-03-04 18:29:54'),('e70c4d6b','236','2016-03-04 18:29:56'),('3d5e2c14','12','2016-03-04 19:11:57'),('1fde975e','12','2016-03-04 19:12:03'),('0d038796','12','2016-03-04 19:12:06'),('7929609a','12','2016-03-04 19:14:09'),('bdf474b4','12','2016-03-04 19:14:38'),('2aef99b3','19','2016-03-04 23:28:57'),('6b530047','19','2016-03-04 23:30:10'),('f8e4bf5a','211','2016-03-04 23:37:06'),('e52f3aae','211','2016-03-04 23:37:12'),('d100afd4','211','2016-03-04 23:37:44'),('f84bb493','211','2016-03-04 23:37:51'),('d398c3a9','211','2016-03-04 23:46:29'),('960db80d','218','2016-03-05 16:00:36'),('340f9ca0','9','2016-03-06 02:40:11'),('90e9c7e5','9','2016-03-06 02:40:21'),('d61442ec','9','2016-03-06 02:48:04'),('0b6da56b','9','2016-03-06 03:00:28'),('23e97a88','9','2016-03-06 03:20:37'),('6cefd242','9','2016-03-06 03:27:09'),('1dc58905','52','2016-03-06 06:01:21'),('0fc4c3e6','20','2016-03-06 23:05:40'),('76cea42f','244','2016-03-06 23:37:17'),('ce10b3fe','27','2016-03-07 20:58:02'),('0387c43f','168','2016-03-07 21:04:05'),('4452f591','168','2016-03-07 21:04:14'),('d08bda0f','168','2016-03-07 21:04:15'),('e5955856','168','2016-03-07 21:04:15'),('cfb31a2b','27','2016-03-07 21:04:25'),('2d542b0a','168','2016-03-07 21:04:54'),('22c9ca25','168','2016-03-07 21:06:26'),('4f8dea2f','168','2016-03-07 21:06:26'),('1f3a9329','168','2016-03-07 21:06:36'),('ca069859','168','2016-03-07 21:06:37'),('b9a5d377','168','2016-03-07 21:06:37'),('931e7e0a','168','2016-03-07 21:07:30'),('546a0a17','168','2016-03-07 21:10:02'),('f3188b69','168','2016-03-07 21:10:09'),('17ff6e13','168','2016-03-07 21:10:10'),('735bb871','168','2016-03-07 21:12:38'),('b12697fa','168','2016-03-07 21:18:12'),('8fdcd843','168','2016-03-07 21:18:13'),('eef77c50','168','2016-03-07 21:18:32'),('b49ead15','168','2016-03-07 21:18:32'),('5ce944e8','168','2016-03-07 21:18:50'),('8c968f9f','168','2016-03-07 21:19:15'),('64c8b34a','168','2016-03-07 21:20:19'),('f51a1007','168','2016-03-07 21:20:38'),('98c8d6a7','168','2016-03-07 21:20:40'),('061b9fbe','168','2016-03-07 21:20:42'),('a807e7a9','168','2016-03-07 21:20:54'),('040f469c','168','2016-03-07 21:21:24'),('97e3cad7','168','2016-03-07 21:21:26'),('90171d55','168','2016-03-07 21:21:31'),('1259a5a6','168','2016-03-07 21:21:37'),('78117e07','168','2016-03-07 21:21:47'),('53a360aa','168','2016-03-07 21:21:49'),('1ae38f66','251','2016-03-09 04:01:32'),('176cfd96','258','2016-03-11 18:04:56'),('3781f26d','235','2016-03-13 04:21:26'),('b9786462','112','2016-03-19 05:09:40'),('0716ee1b','241','2016-03-19 13:06:25'),('9d0d723c','241','2016-03-19 13:06:41'),('98462f13','7','2016-05-08 10:26:35'),('1cd4918c','4','2016-05-10 11:02:14');
/*!40000 ALTER TABLE `state` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-05-09 23:08:06
