-- MySQL dump 10.13  Distrib 5.7.16, for osx10.12 (x86_64)
--
-- Host: localhost    Database: rentroll
-- ------------------------------------------------------
-- Server version	5.7.16

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
-- Table structure for table `AR`
--

DROP TABLE IF EXISTS `AR`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `AR` (
  `ARID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Name` varchar(100) NOT NULL DEFAULT '',
  `ARType` smallint(6) NOT NULL DEFAULT '0',
  `RARequired` smallint(6) NOT NULL DEFAULT '0',
  `DebitLID` bigint(20) NOT NULL DEFAULT '0',
  `CreditLID` bigint(20) NOT NULL DEFAULT '0',
  `Description` varchar(1024) NOT NULL DEFAULT '',
  `DtStart` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtStop` datetime NOT NULL DEFAULT '2066-01-01 00:00:00',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `DefaultAmount` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`ARID`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `AR`
--

LOCK TABLES `AR` WRITE;
/*!40000 ALTER TABLE `AR` DISABLE KEYS */;
INSERT INTO `AR` VALUES (1,1,'Rent Taxable',0,3,8,16,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:23:21',0,'2017-08-16 23:19:30',0),(2,1,'Rent Non-Taxable',0,3,8,17,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:23:11',0,'2017-08-16 23:20:05',0),(3,1,'Electric Overage',0,3,8,36,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:26:13',0,'2017-08-16 23:22:28',0),(4,1,'Electric Base Fee',0,3,8,35,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:25:44',0,'2017-08-16 23:25:44',0),(7,1,'Water and Sewer Base Fee',0,3,8,37,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:29:50',0,'2017-08-16 23:29:50',0),(8,1,'Water and Sewer Overage',0,3,8,38,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:30:22',0,'2017-08-16 23:30:22',0),(9,1,'Gas Base Fee',0,3,8,39,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:31:02',0,'2017-08-16 23:31:02',0),(10,1,'Gas Base Overage',0,3,8,40,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:31:27',0,'2017-08-16 23:31:27',0),(11,1,'Application Fee',0,3,8,45,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:36:16',0,'2017-08-16 23:32:46',0),(12,1,'Late Fee',0,3,8,46,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:33:10',0,'2017-08-16 23:33:10',0),(13,1,'Month to Month Fee',0,3,8,48,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:33:34',0,'2017-08-16 23:33:34',0),(14,1,'Insufficient Funds Fee',0,3,8,47,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:34:00',0,'2017-08-16 23:34:00',0),(15,1,'No Show / Termination Fee',0,3,8,50,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:35:03',0,'2017-08-16 23:35:03',0),(16,1,'Pet Fee',0,3,8,51,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:35:23',0,'2017-08-16 23:35:23',0),(17,1,'Pet Rent',0,3,8,52,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:41:40',0,'2017-08-16 23:35:54',0),(18,1,'Tenant Expense Chargeback',0,3,8,53,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:36:59',0,'2017-08-16 23:36:59',0),(19,1,'Special Cleaning Fee',0,3,8,54,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:37:27',0,'2017-08-16 23:37:27',0),(20,1,'Eviction Fee Reimbursement',0,3,8,55,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:37:59',0,'2017-08-16 23:37:59',0),(21,1,'Security Deposit Forfeiture',0,3,11,57,'Forfeit','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-09-21 02:21:35',0,'2017-08-16 23:38:38',0),(22,1,'Damage Fee',0,3,8,58,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:39:18',0,'2017-08-16 23:39:18',0),(23,1,'Other Special Tenant Charges',0,3,8,60,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:39:57',0,'2017-08-16 23:39:57',0),(24,1,'Security Deposit Assessment',0,3,8,11,'normal deposit','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-09-25 21:21:05',0,'2017-08-16 23:40:26',0),(25,1,'Bad Debt Write-Off',2,3,70,8,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-22 13:38:33',0,'2017-08-22 13:38:33',0),(26,1,'Bank Service Fee (Operating Account)',2,3,71,3,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-23 14:10:11',0,'2017-08-22 13:40:32',0),(27,1,'Receive a Payment',1,3,5,10,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-22 13:41:19',0,'2017-08-22 13:41:19',0),(28,1,'Deposit to Operating Account (FRB54320)',1,3,3,5,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-22 13:42:29',0,'2017-08-22 13:42:29',0),(29,1,'Deposit to Deposit Account (FRB96953)',1,3,4,5,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-22 13:43:44',0,'2017-08-22 13:43:22',0),(30,1,'Apply Payment',1,3,10,8,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-22 14:06:06',0,'2017-08-22 14:05:19',0),(31,1,'Bank Service Fee (Deposit Account)',2,3,71,4,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-23 14:11:03',0,'2017-08-23 14:11:03',0),(32,1,'Broken Window charge',0,3,8,58,'','2017-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-09-15 06:42:13',0,'2017-09-15 06:42:13',0),(33,1,'Vending Income',1,3,6,64,'','2010-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-09-18 19:47:42',0,'2017-09-18 19:37:22',0),(34,1,'Security Deposit Refund',0,3,11,73,'Refund','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-09-21 02:22:30',0,'2017-09-21 02:22:30',0),(35,1,'Floating Security Deposit',1,3,5,75,'','2014-01-01 00:00:00','9999-12-31 00:00:00',1,0.0000,'2017-09-27 05:16:31',0,'2017-09-22 21:20:49',0);
/*!40000 ALTER TABLE `AR` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `AssessmentTax`
--

DROP TABLE IF EXISTS `AssessmentTax`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `AssessmentTax` (
  `ASMID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `TAXID` bigint(20) NOT NULL DEFAULT '0',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `OverrideTaxApprover` mediumint(9) NOT NULL DEFAULT '0',
  `OverrideAmount` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `AssessmentTax`
--

LOCK TABLES `AssessmentTax` WRITE;
/*!40000 ALTER TABLE `AssessmentTax` DISABLE KEYS */;
/*!40000 ALTER TABLE `AssessmentTax` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Assessments`
--

DROP TABLE IF EXISTS `Assessments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Assessments` (
  `ASMID` bigint(20) NOT NULL AUTO_INCREMENT,
  `PASMID` bigint(20) NOT NULL DEFAULT '0',
  `RPASMID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `ATypeLID` bigint(20) NOT NULL DEFAULT '0',
  `RAID` bigint(20) NOT NULL DEFAULT '0',
  `Amount` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `Start` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `Stop` datetime NOT NULL DEFAULT '2066-01-01 00:00:00',
  `RentCycle` smallint(6) NOT NULL DEFAULT '0',
  `ProrationCycle` smallint(6) NOT NULL DEFAULT '0',
  `InvoiceNo` bigint(20) NOT NULL DEFAULT '0',
  `AcctRule` varchar(200) NOT NULL DEFAULT '',
  `ARID` bigint(20) NOT NULL DEFAULT '0',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `Comment` varchar(256) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`ASMID`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Assessments`
--

LOCK TABLES `Assessments` WRITE;
/*!40000 ALTER TABLE `Assessments` DISABLE KEYS */;
INSERT INTO `Assessments` VALUES (1,0,0,1,2,0,3,4000.0000,'2016-10-01 00:00:00','2018-01-01 00:00:00',6,4,0,'',2,0,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(2,1,0,1,2,0,3,4000.0000,'2016-10-01 00:00:00','2016-10-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:40:14',0,'2017-08-22 18:24:26',0),(3,1,0,1,2,0,3,4000.0000,'2016-11-01 00:00:00','2016-11-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:40:14',0,'2017-08-22 18:24:26',0),(4,1,0,1,2,0,3,4000.0000,'2016-12-01 00:00:00','2016-12-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:40:14',0,'2017-08-22 18:24:26',0),(5,1,0,1,2,0,3,4000.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:40:14',0,'2017-08-22 18:24:26',0),(6,1,0,1,2,0,3,4000.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:53:58',0,'2017-08-22 18:24:26',0),(7,1,0,1,2,0,3,4000.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:53:58',0,'2017-08-22 18:24:26',0),(8,1,0,1,2,0,3,4000.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:53:58',0,'2017-08-22 18:24:26',0),(9,1,0,1,2,0,3,4000.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:53:58',0,'2017-08-22 18:24:26',0),(10,1,0,1,2,0,3,4000.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:53:58',0,'2017-08-22 18:24:26',0),(11,1,0,1,2,0,3,4000.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:53:58',0,'2017-08-22 18:24:26',0),(12,1,0,1,2,0,3,4000.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:53:58',0,'2017-08-22 18:24:26',0),(13,0,0,1,2,0,3,628.4500,'2017-01-01 00:00:00','2017-01-01 00:00:00',0,0,0,'',4,2,'utilities allowance','2017-08-22 18:53:58',0,'2017-08-22 18:27:44',0),(14,0,0,1,2,0,3,175.0000,'2017-02-01 00:00:00','2017-02-01 00:00:00',0,0,0,'',4,2,'utilities','2017-08-22 18:53:58',0,'2017-08-22 18:28:39',0),(15,0,0,1,2,0,3,175.0000,'2017-02-01 00:00:00','2017-02-01 00:00:00',0,0,0,'',4,4,'Reversed by ASM00000024','2017-08-22 18:52:07',0,'2017-08-22 18:47:56',0),(16,0,0,1,2,0,3,175.0000,'2017-03-01 00:00:00','2017-03-01 00:00:00',0,0,0,'',4,2,'utilities','2017-08-22 18:53:58',0,'2017-08-22 18:48:32',0),(17,0,0,1,2,0,3,81.7900,'2017-04-01 00:00:00','2017-04-01 00:00:00',0,0,0,'',4,2,'retro utilities','2017-08-22 18:53:58',0,'2017-08-22 18:49:22',0),(18,0,0,1,2,0,3,350.0000,'2017-04-01 00:00:00','2017-12-31 00:00:00',6,4,0,'',4,0,'monthly utilities allowance','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(19,18,0,1,2,0,3,350.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',4,2,'monthly utilities allowance','2017-08-22 18:53:58',0,'2017-08-22 18:50:08',0),(20,18,0,1,2,0,3,350.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',4,2,'monthly utilities allowance','2017-08-22 18:53:58',0,'2017-08-22 18:50:08',0),(21,18,0,1,2,0,3,350.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',4,2,'monthly utilities allowance','2017-08-22 18:53:58',0,'2017-08-22 18:50:08',0),(22,18,0,1,2,0,3,350.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',4,2,'monthly utilities allowance','2017-08-22 18:53:58',0,'2017-08-22 18:50:08',0),(23,18,0,1,2,0,3,350.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',4,2,'monthly utilities allowance','2017-08-22 18:53:58',0,'2017-08-22 18:50:08',0),(24,0,15,1,2,0,3,-175.0000,'2017-02-01 00:00:00','2017-02-01 00:00:00',0,0,0,'',4,4,'Reversal of ASM00000015','2017-08-22 18:52:07',0,'2017-08-22 18:52:07',0),(25,1,0,1,2,0,3,4000.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',2,0,'','2017-09-01 00:00:08',0,'2017-09-01 00:00:08',0),(26,18,0,1,2,0,3,350.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',4,0,'monthly utilities allowance','2017-09-01 00:00:08',0,'2017-09-01 00:00:08',0),(27,0,0,1,1,0,1,4000.0000,'2017-03-01 00:00:00','2018-03-01 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(28,27,0,1,1,0,1,4000.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(29,27,0,1,1,0,1,4000.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(30,27,0,1,1,0,1,4000.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(31,27,0,1,1,0,1,4000.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(32,27,0,1,1,0,1,4000.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(33,27,0,1,1,0,1,4000.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(34,27,0,1,1,0,1,4000.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(35,0,0,1,3,0,2,4150.0000,'2017-01-01 00:00:00','2018-01-01 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(36,35,0,1,3,0,2,4150.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(37,35,0,1,3,0,2,4150.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(38,35,0,1,3,0,2,4150.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(39,35,0,1,3,0,2,4150.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(40,35,0,1,3,0,2,4150.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(41,35,0,1,3,0,2,4150.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(42,35,0,1,3,0,2,4150.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(43,35,0,1,3,0,2,4150.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(44,35,0,1,3,0,2,4150.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(45,0,0,1,1,0,1,200.0000,'2017-09-14 00:00:00','2017-09-14 00:00:00',0,0,0,'',3,2,'','2017-09-15 05:44:13',0,'2017-09-15 05:42:55',0),(46,0,0,1,1,0,1,3500.0000,'2014-03-01 00:00:00','2014-03-01 00:00:00',0,0,0,'',24,0,'','2017-09-15 23:30:54',0,'2017-09-15 23:30:54',0),(47,0,0,1,1,0,1,3500.0000,'2014-03-02 00:00:00','2014-03-02 00:00:00',0,0,0,'',24,0,'','2017-09-15 23:41:23',0,'2017-09-15 23:41:23',0),(48,0,0,1,1,0,1,3000.0000,'2017-01-01 00:00:00','2017-01-01 00:00:00',0,0,0,'',24,0,'','2017-09-16 00:27:23',0,'2017-09-16 00:27:23',0),(49,0,0,1,0,0,8,50.0000,'2017-09-18 00:00:00','2017-09-18 00:00:00',0,0,0,'',11,2,'','2017-09-26 05:29:10',0,'2017-09-18 18:07:30',0),(50,0,0,1,0,0,9,50.0000,'2017-09-18 00:00:00','2017-09-18 00:00:00',0,0,0,'',11,2,'','2017-09-26 05:33:06',0,'2017-09-18 18:08:54',0),(51,0,0,1,0,0,10,50.0000,'2017-09-18 00:00:00','2017-09-18 00:00:00',0,0,0,'',11,2,'','2017-09-26 05:33:10',0,'2017-09-18 18:25:13',0),(52,0,0,1,0,0,11,50.0000,'2017-09-18 00:00:00','2017-09-18 00:00:00',0,0,0,'',11,2,'','2017-09-26 05:33:14',0,'2017-09-18 18:25:58',0),(53,0,0,1,0,0,12,50.0000,'2017-09-18 00:00:00','2017-09-18 00:00:00',0,0,0,'',11,4,'Reversed by ASM00000055','2017-09-25 20:51:37',0,'2017-09-18 18:26:24',0),(54,0,0,1,0,0,12,75.0000,'2017-09-03 00:00:00','2017-09-03 00:00:00',0,0,0,'',23,2,'Parking Ticket','2017-09-26 05:33:18',0,'2017-09-25 20:51:01',0),(55,0,53,1,0,0,12,-50.0000,'2017-09-18 00:00:00','2017-09-18 00:00:00',0,0,0,'',11,4,'Reversal of ASM00000053','2017-09-25 20:51:37',0,'2017-09-25 20:51:37',0),(56,0,0,1,0,0,14,100.0000,'2017-09-25 00:00:00','2017-09-25 00:00:00',0,0,0,'',36,0,'','2017-09-25 21:23:20',0,'2017-09-25 21:23:20',0),(57,0,0,1,0,0,14,5000.0000,'2017-08-02 00:00:00','2017-08-02 00:00:00',0,0,0,'',36,0,'Wants the next unit type X that becomes available','2017-09-25 21:25:30',0,'2017-09-25 21:25:30',0);
/*!40000 ALTER TABLE `Assessments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `AvailabilityTypes`
--

DROP TABLE IF EXISTS `AvailabilityTypes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `AvailabilityTypes` (
  `AVAILID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL,
  `Name` varchar(100) NOT NULL DEFAULT '',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`AVAILID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `AvailabilityTypes`
--

LOCK TABLES `AvailabilityTypes` WRITE;
/*!40000 ALTER TABLE `AvailabilityTypes` DISABLE KEYS */;
/*!40000 ALTER TABLE `AvailabilityTypes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Building`
--

DROP TABLE IF EXISTS `Building`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Building` (
  `BLDGID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Address` varchar(100) NOT NULL DEFAULT '',
  `Address2` varchar(100) NOT NULL DEFAULT '',
  `City` varchar(100) NOT NULL DEFAULT '',
  `State` char(25) NOT NULL DEFAULT '',
  `PostalCode` varchar(100) NOT NULL DEFAULT '',
  `Country` varchar(100) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`BLDGID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Building`
--

LOCK TABLES `Building` WRITE;
/*!40000 ALTER TABLE `Building` DISABLE KEYS */;
/*!40000 ALTER TABLE `Building` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Business`
--

DROP TABLE IF EXISTS `Business`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Business` (
  `BID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BUD` varchar(100) NOT NULL DEFAULT '',
  `Name` varchar(100) NOT NULL DEFAULT '',
  `DefaultRentCycle` smallint(6) NOT NULL DEFAULT '0',
  `DefaultProrationCycle` smallint(6) NOT NULL DEFAULT '0',
  `DefaultGSRPC` smallint(6) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`BID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Business`
--

LOCK TABLES `Business` WRITE;
/*!40000 ALTER TABLE `Business` DISABLE KEYS */;
INSERT INTO `Business` VALUES (1,'REX','JGM First, LLC',6,4,4,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0);
/*!40000 ALTER TABLE `Business` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `BusinessAssessments`
--

DROP TABLE IF EXISTS `BusinessAssessments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `BusinessAssessments` (
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `ATypeLID` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `BusinessAssessments`
--

LOCK TABLES `BusinessAssessments` WRITE;
/*!40000 ALTER TABLE `BusinessAssessments` DISABLE KEYS */;
/*!40000 ALTER TABLE `BusinessAssessments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `BusinessPaymentTypes`
--

DROP TABLE IF EXISTS `BusinessPaymentTypes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `BusinessPaymentTypes` (
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `PMTID` mediumint(9) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `BusinessPaymentTypes`
--

LOCK TABLES `BusinessPaymentTypes` WRITE;
/*!40000 ALTER TABLE `BusinessPaymentTypes` DISABLE KEYS */;
/*!40000 ALTER TABLE `BusinessPaymentTypes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `CommissionLedger`
--

DROP TABLE IF EXISTS `CommissionLedger`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `CommissionLedger` (
  `CLID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RAID` bigint(20) NOT NULL DEFAULT '0',
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `Salesperson` varchar(100) NOT NULL DEFAULT '',
  `Percent` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `Amount` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `PaymentDueDate` date NOT NULL DEFAULT '1970-01-01',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`CLID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `CommissionLedger`
--

LOCK TABLES `CommissionLedger` WRITE;
/*!40000 ALTER TABLE `CommissionLedger` DISABLE KEYS */;
/*!40000 ALTER TABLE `CommissionLedger` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `CustomAttr`
--

DROP TABLE IF EXISTS `CustomAttr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `CustomAttr` (
  `CID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Type` smallint(6) NOT NULL DEFAULT '0',
  `Name` varchar(100) NOT NULL DEFAULT '',
  `Value` varchar(256) NOT NULL DEFAULT '',
  `Units` varchar(256) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`CID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `CustomAttr`
--

LOCK TABLES `CustomAttr` WRITE;
/*!40000 ALTER TABLE `CustomAttr` DISABLE KEYS */;
INSERT INTO `CustomAttr` VALUES (1,1,1,'Square Feet','1215','sqft','2017-09-15 07:57:52',0,'2017-09-15 07:57:52',0),(2,1,1,'Square Feet','1239','sqft','2017-09-15 07:57:52',0,'2017-09-15 07:57:52',0),(3,1,1,'Square Feet','1433','sqft','2017-09-15 07:57:52',0,'2017-09-15 07:57:52',0),(4,1,1,'Square Feet','827','sqft','2017-09-15 07:57:52',0,'2017-09-15 07:57:52',0),(5,1,1,'Square Feet','1285','sqft','2017-09-15 07:57:52',0,'2017-09-15 07:57:52',0),(6,1,1,'Square Feet','1175','sqft','2017-09-15 07:57:52',0,'2017-09-15 07:57:52',0);
/*!40000 ALTER TABLE `CustomAttr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `CustomAttrRef`
--

DROP TABLE IF EXISTS `CustomAttrRef`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `CustomAttrRef` (
  `ElementType` bigint(20) NOT NULL,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `ID` bigint(20) NOT NULL,
  `CID` bigint(20) NOT NULL,
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `CustomAttrRef`
--

LOCK TABLES `CustomAttrRef` WRITE;
/*!40000 ALTER TABLE `CustomAttrRef` DISABLE KEYS */;
INSERT INTO `CustomAttrRef` VALUES (5,1,1,1,'2017-09-15 07:57:55',0),(5,1,2,2,'2017-09-15 07:57:55',0),(5,1,3,3,'2017-09-15 07:57:55',0),(5,1,4,4,'2017-09-15 07:57:55',0),(5,1,5,5,'2017-09-15 07:57:55',0),(5,1,6,6,'2017-09-15 07:57:55',0);
/*!40000 ALTER TABLE `CustomAttrRef` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DemandSource`
--

DROP TABLE IF EXISTS `DemandSource`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `DemandSource` (
  `SourceSLSID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Name` varchar(100) DEFAULT NULL,
  `Industry` varchar(100) DEFAULT NULL,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`SourceSLSID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DemandSource`
--

LOCK TABLES `DemandSource` WRITE;
/*!40000 ALTER TABLE `DemandSource` DISABLE KEYS */;
/*!40000 ALTER TABLE `DemandSource` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Deposit`
--

DROP TABLE IF EXISTS `Deposit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Deposit` (
  `DID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `DEPID` bigint(20) NOT NULL DEFAULT '0',
  `DPMID` bigint(20) NOT NULL DEFAULT '0',
  `Dt` date NOT NULL DEFAULT '1970-01-01',
  `Amount` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `ClearedAmount` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`DID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Deposit`
--

LOCK TABLES `Deposit` WRITE;
/*!40000 ALTER TABLE `Deposit` DISABLE KEYS */;
/*!40000 ALTER TABLE `Deposit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DepositMethod`
--

DROP TABLE IF EXISTS `DepositMethod`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `DepositMethod` (
  `DPMID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Method` varchar(50) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`DPMID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DepositMethod`
--

LOCK TABLES `DepositMethod` WRITE;
/*!40000 ALTER TABLE `DepositMethod` DISABLE KEYS */;
INSERT INTO `DepositMethod` VALUES (1,1,'Hand Delivered','2017-08-21 20:24:11',0,'2017-08-21 20:24:11',0);
/*!40000 ALTER TABLE `DepositMethod` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DepositPart`
--

DROP TABLE IF EXISTS `DepositPart`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `DepositPart` (
  `DPID` bigint(20) NOT NULL AUTO_INCREMENT,
  `DID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RCPTID` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`DPID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DepositPart`
--

LOCK TABLES `DepositPart` WRITE;
/*!40000 ALTER TABLE `DepositPart` DISABLE KEYS */;
/*!40000 ALTER TABLE `DepositPart` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Depository`
--

DROP TABLE IF EXISTS `Depository`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Depository` (
  `DEPID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `LID` bigint(20) NOT NULL DEFAULT '0',
  `Name` varchar(256) DEFAULT NULL,
  `AccountNo` varchar(256) DEFAULT NULL,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`DEPID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Depository`
--

LOCK TABLES `Depository` WRITE;
/*!40000 ALTER TABLE `Depository` DISABLE KEYS */;
INSERT INTO `Depository` VALUES (1,1,3,'FRB Operating Account','80001054320','2017-08-16 22:07:41',0,'2017-08-16 22:07:41',0),(2,1,4,'FRB Tenant Deposit Account','80003196953','2017-08-16 22:09:06',0,'2017-08-16 22:09:06',0);
/*!40000 ALTER TABLE `Depository` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Expense`
--

DROP TABLE IF EXISTS `Expense`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Expense` (
  `EXPID` bigint(20) NOT NULL AUTO_INCREMENT,
  `RPEXPID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `RAID` bigint(20) NOT NULL DEFAULT '0',
  `Amount` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `Dt` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `AcctRule` varchar(200) NOT NULL DEFAULT '',
  `ARID` bigint(20) NOT NULL DEFAULT '0',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `Comment` varchar(256) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`EXPID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Expense`
--

LOCK TABLES `Expense` WRITE;
/*!40000 ALTER TABLE `Expense` DISABLE KEYS */;
INSERT INTO `Expense` VALUES (1,0,1,2,3,15.0000,'2016-11-11 00:00:00','',26,0,'wire fee charged by bank','2017-08-22 18:32:26',0,'2017-08-22 18:32:26',0);
/*!40000 ALTER TABLE `Expense` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `GLAccount`
--

DROP TABLE IF EXISTS `GLAccount`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `GLAccount` (
  `LID` bigint(20) NOT NULL AUTO_INCREMENT,
  `PLID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RAID` bigint(20) NOT NULL DEFAULT '0',
  `TCID` bigint(20) NOT NULL DEFAULT '0',
  `GLNumber` varchar(100) NOT NULL DEFAULT '',
  `Status` smallint(6) NOT NULL DEFAULT '0',
  `Name` varchar(100) NOT NULL DEFAULT '',
  `AcctType` varchar(100) NOT NULL DEFAULT '',
  `AllowPost` smallint(6) NOT NULL DEFAULT '0',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `Description` varchar(1024) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`LID`)
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `GLAccount`
--

LOCK TABLES `GLAccount` WRITE;
/*!40000 ALTER TABLE `GLAccount` DISABLE KEYS */;
INSERT INTO `GLAccount` VALUES (1,0,1,0,0,'10000',2,'Cash','Cash',0,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(2,0,1,0,0,'10100',2,'Petty Cash','Cash',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(3,1,1,0,0,'10104',2,'FRB 54320 (operating account)','Bank Account',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(4,1,1,0,0,'10105',2,'FRB 96953 (deposit account)','Bank Account',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(5,1,1,0,0,'10999',2,'Undeposited Funds','Cash',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(6,0,1,0,0,'11000',2,'Credit Cards Funds in Transit','Cash',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(7,0,1,0,0,'12000',2,'Accounts Receivable','Accounts Receivable',0,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(8,7,1,0,0,'12001',2,'Rent Roll Receivables','Accounts Receivable',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(10,0,1,0,0,'12999',2,'Unapplied Funds','Asset',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(11,0,1,0,0,'30000',2,'Security Deposit Liability','Liability Security Deposit',1,0,'','2017-09-21 02:18:08',0,'2017-08-16 05:49:58',0),(12,0,1,0,0,'30100',2,'Collected Taxes','Liabilities',0,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(13,12,1,0,0,'30101',2,'Sales Taxes Collected','Liabilities',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(14,12,1,0,0,'30102',2,'Transient Occupancy Taxes Collected','Liabilities',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(15,12,1,0,0,'30199',2,'Other Collected Taxes','Liabilities',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(16,0,1,0,0,'41000',2,'Gross Scheduled Rent-Taxable','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(17,0,1,0,0,'41001',2,'Gross Scheduled Rent-Not Taxable','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(18,0,1,0,0,'41100',2,'Unit Income Offsets','Income Offset',0,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(19,18,1,0,0,'41101',2,'Vacancy','Income Offset',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(20,18,1,0,0,'41102',2,'Loss (Gain) to Lease','Income Offset',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(21,18,1,0,0,'41103',2,'Employee Concessions','Income Offset',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(22,18,1,0,0,'41104',2,'Resident Concessions','Income Offset',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(23,18,1,0,0,'41105',2,'Owner Concession','Income Offset',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(24,18,1,0,0,'41106',2,'Administrative Concession','Income Offset',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(25,18,1,0,0,'41107',2,'Off Line Renovations','Income Offset',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(26,18,1,0,0,'41108',2,'Off Line Maintenance','Income Offset',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(27,18,1,0,0,'41199',2,'Othe Income Offsets','Income Offset',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(28,0,1,0,0,'41200',2,'Service Fees','Income',0,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(29,28,1,0,0,'41201',2,'Broadcast and IT Services','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(30,28,1,0,0,'41202',2,'Food Services','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(31,28,1,0,0,'41203',2,'Linen Services','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(32,28,1,0,0,'41204',2,'Wash N Fold Services','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(33,28,1,0,0,'41299',2,'Other Service Fees','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(34,0,1,0,0,'41300',2,'Utility Fees','Income',0,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(35,34,1,0,0,'41301',2,'Electric Base Fee','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(36,34,1,0,0,'41302',2,'Electric Overage','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(37,34,1,0,0,'41303',2,'Water and Sewer Base Fee','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(38,34,1,0,0,'41304',2,'Water and Sewer Overage','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(39,34,1,0,0,'41305',2,'Gas Base Fee','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(40,34,1,0,0,'41306',2,'Gas Overage','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(41,34,1,0,0,'41307',2,'Trash Collection Base Fee','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(42,34,1,0,0,'41308',2,'Trash Collection Overage','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(43,34,1,0,0,'41399',2,'Other Utility Fee','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(44,0,1,0,0,'41400',2,'Special Tenant Charges','Income',0,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(45,44,1,0,0,'41401',2,'Application Fees','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(46,44,1,0,0,'41402',2,'Late Fees','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(47,44,1,0,0,'41403',2,'Insufficient Funds Fee','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(48,44,1,0,0,'41404',2,'Month to Month Fee','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(49,44,1,0,0,'41405',2,'Rentable Specialties','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(50,44,1,0,0,'41406',2,'No Show or Termination Fee','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(51,44,1,0,0,'41407',2,'Pet Fees','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(52,44,1,0,0,'41408',2,'Pet Rent','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(53,44,1,0,0,'41409',2,'Tenant Expense Chargeback','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(54,44,1,0,0,'41410',2,'Special Cleaning Fee','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(55,44,1,0,0,'41411',2,'Eviction Fee Reimbursement','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(56,44,1,0,0,'41412',2,'Extra Person Charge','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(57,44,1,0,0,'41413',2,'Security Deposit Forfeiture','Income',1,0,'','2017-09-22 20:27:13',0,'2017-08-16 05:49:58',0),(58,44,1,0,0,'41414',2,'Damage Fee','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(59,44,1,0,0,'41415',2,'CAM Fees','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(60,44,1,0,0,'41499',2,'Other Special Tenant Charges','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(61,0,1,0,0,'42000',2,'Business Income','Income',0,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(62,61,1,0,0,'42100',2,'Convenience Store','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(63,61,1,0,0,'42200',2,'Fitness Center Revenue','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(64,61,1,0,0,'42300',2,'Vending Income','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(65,61,1,0,0,'42400',2,'Restaurant Sales','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(66,61,1,0,0,'42500',2,'Bar Sales','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(67,61,1,0,0,'42600',2,'Spa Sales','Income',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(68,0,1,0,0,'50000',2,'Expenses','Expenses',0,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(69,68,1,0,0,'50001',2,'Cash Over/Short','Expenses',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(70,68,1,0,0,'50002',2,'Bad Debt','Expenses',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(71,68,1,0,0,'50003',2,'Bank Service Fee','Expenses',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(72,68,1,0,0,'50999',2,'Other Expenses','Expenses',1,0,'','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(73,1,1,0,0,'10199',2,'Security Deposit Refund','Cash',1,0,'','2017-09-22 20:28:07',0,'2017-09-21 02:19:32',0),(74,0,1,0,0,'999911',2,'test 1','Cash',1,0,'laskdjf','2017-09-22 20:59:16',0,'2017-09-22 20:59:16',0),(75,0,1,0,0,'30001',2,'Floating Security Deposits','Liability Security Deposit',1,0,'Sec Dep posted before rentable identified','2017-09-22 21:19:03',0,'2017-09-22 21:19:03',0);
/*!40000 ALTER TABLE `GLAccount` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Invoice`
--

DROP TABLE IF EXISTS `Invoice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Invoice` (
  `InvoiceNo` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Dt` date NOT NULL DEFAULT '1970-01-01',
  `DtDue` date NOT NULL DEFAULT '1970-01-01',
  `Amount` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `DeliveredBy` varchar(256) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`InvoiceNo`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Invoice`
--

LOCK TABLES `Invoice` WRITE;
/*!40000 ALTER TABLE `Invoice` DISABLE KEYS */;
/*!40000 ALTER TABLE `Invoice` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `InvoiceAssessment`
--

DROP TABLE IF EXISTS `InvoiceAssessment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `InvoiceAssessment` (
  `InvoiceNo` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `ASMID` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `InvoiceAssessment`
--

LOCK TABLES `InvoiceAssessment` WRITE;
/*!40000 ALTER TABLE `InvoiceAssessment` DISABLE KEYS */;
/*!40000 ALTER TABLE `InvoiceAssessment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `InvoicePayor`
--

DROP TABLE IF EXISTS `InvoicePayor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `InvoicePayor` (
  `InvoiceNo` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `PID` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `InvoicePayor`
--

LOCK TABLES `InvoicePayor` WRITE;
/*!40000 ALTER TABLE `InvoicePayor` DISABLE KEYS */;
/*!40000 ALTER TABLE `InvoicePayor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Journal`
--

DROP TABLE IF EXISTS `Journal`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Journal` (
  `JID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Dt` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `Amount` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `Type` smallint(6) NOT NULL DEFAULT '0',
  `ID` bigint(20) NOT NULL DEFAULT '0',
  `Comment` varchar(256) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`JID`)
) ENGINE=InnoDB AUTO_INCREMENT=96 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Journal`
--

LOCK TABLES `Journal` WRITE;
/*!40000 ALTER TABLE `Journal` DISABLE KEYS */;
INSERT INTO `Journal` VALUES (1,1,'2016-10-01 00:00:00',4000.0000,1,2,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(2,1,'2016-11-01 00:00:00',4000.0000,1,3,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(3,1,'2016-12-01 00:00:00',4000.0000,1,4,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(4,1,'2017-01-01 00:00:00',4000.0000,1,5,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(5,1,'2017-02-01 00:00:00',4000.0000,1,6,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(6,1,'2017-03-01 00:00:00',4000.0000,1,7,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(7,1,'2017-04-01 00:00:00',4000.0000,1,8,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(8,1,'2017-05-01 00:00:00',4000.0000,1,9,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(9,1,'2017-06-01 00:00:00',4000.0000,1,10,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(10,1,'2017-07-01 00:00:00',4000.0000,1,11,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(11,1,'2017-08-01 00:00:00',4000.0000,1,12,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(12,1,'2017-01-01 00:00:00',628.4500,1,13,'','2017-08-22 18:27:44',0,'2017-08-22 18:27:44',0),(13,1,'2017-02-01 00:00:00',175.0000,1,14,'','2017-08-22 18:28:39',0,'2017-08-22 18:28:39',0),(14,1,'2016-10-03 00:00:00',4000.0000,2,1,'','2017-08-22 18:29:52',0,'2017-08-22 18:29:52',0),(15,1,'2016-11-11 00:00:00',12000.0000,2,2,'','2017-08-22 18:31:17',0,'2017-08-22 18:31:17',0),(16,1,'2016-11-11 00:00:00',15.0000,3,1,'','2017-08-22 18:32:26',0,'2017-08-22 18:32:26',0),(17,1,'2016-10-03 00:00:00',4000.0000,2,1,'','2017-08-22 18:40:14',0,'2017-08-22 18:40:14',0),(18,1,'2016-11-11 00:00:00',4000.0000,2,2,'','2017-08-22 18:40:14',0,'2017-08-22 18:40:14',0),(19,1,'2016-12-03 00:00:00',4000.0000,2,2,'','2017-09-21 21:30:14',0,'2017-08-22 18:40:14',0),(20,1,'2017-01-03 00:00:00',4000.0000,2,2,'','2017-09-21 21:30:33',0,'2017-08-22 18:40:14',0),(21,1,'2017-02-03 00:00:00',628.4500,2,3,'','2017-08-22 18:43:32',0,'2017-08-22 18:43:32',0),(22,1,'2017-02-13 00:00:00',8350.0000,2,4,'','2017-08-22 18:44:27',0,'2017-08-22 18:44:27',0),(23,1,'2017-05-12 00:00:00',13131.7900,2,5,'','2017-08-22 18:45:12',0,'2017-08-22 18:45:12',0),(24,1,'2017-08-15 00:00:00',13050.0000,2,6,'','2017-08-22 18:45:41',0,'2017-08-22 18:45:41',0),(25,1,'2017-02-01 00:00:00',175.0000,1,15,'','2017-08-22 18:47:56',0,'2017-08-22 18:47:56',0),(26,1,'2017-03-01 00:00:00',175.0000,1,16,'','2017-08-22 18:48:32',0,'2017-08-22 18:48:32',0),(27,1,'2017-04-01 00:00:00',81.7900,1,17,'','2017-08-22 18:49:22',0,'2017-08-22 18:49:22',0),(28,1,'2017-04-01 00:00:00',350.0000,1,19,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(29,1,'2017-05-01 00:00:00',350.0000,1,20,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(30,1,'2017-06-01 00:00:00',350.0000,1,21,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(31,1,'2017-07-01 00:00:00',350.0000,1,22,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(32,1,'2017-08-01 00:00:00',350.0000,1,23,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(33,1,'2017-02-01 00:00:00',-175.0000,1,24,'','2017-08-22 18:52:07',0,'2017-08-22 18:52:07',0),(34,1,'2017-02-03 00:00:00',628.4500,2,3,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(35,1,'2017-02-13 00:00:00',4000.0000,2,4,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(36,1,'2017-02-13 00:00:00',175.0000,2,4,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(37,1,'2017-02-13 00:00:00',4000.0000,2,4,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(38,1,'2017-02-13 00:00:00',175.0000,2,4,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(39,1,'2017-05-12 00:00:00',4000.0000,2,5,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(40,1,'2017-05-12 00:00:00',81.7900,2,5,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(41,1,'2017-05-12 00:00:00',350.0000,2,5,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(42,1,'2017-05-12 00:00:00',4000.0000,2,5,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(43,1,'2017-05-12 00:00:00',350.0000,2,5,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(44,1,'2017-05-12 00:00:00',4000.0000,2,5,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(45,1,'2017-05-12 00:00:00',350.0000,2,5,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(46,1,'2017-08-15 00:00:00',4000.0000,2,6,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(47,1,'2017-08-15 00:00:00',350.0000,2,6,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(48,1,'2017-08-15 00:00:00',4000.0000,2,6,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(49,1,'2017-08-15 00:00:00',350.0000,2,6,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(50,1,'2017-09-01 00:00:00',4000.0000,1,25,'','2017-09-01 00:00:08',0,'2017-09-01 00:00:08',0),(51,1,'2017-09-01 00:00:00',350.0000,1,26,'','2017-09-01 00:00:08',0,'2017-09-01 00:00:08',0),(52,1,'2017-03-01 00:00:00',4000.0000,1,28,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(53,1,'2017-04-01 00:00:00',4000.0000,1,29,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(54,1,'2017-05-01 00:00:00',4000.0000,1,30,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(55,1,'2017-06-01 00:00:00',4000.0000,1,31,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(56,1,'2017-07-01 00:00:00',4000.0000,1,32,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(57,1,'2017-08-01 00:00:00',4000.0000,1,33,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(58,1,'2017-09-01 00:00:00',4000.0000,1,34,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(59,1,'2017-01-01 00:00:00',4150.0000,1,36,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(60,1,'2017-02-01 00:00:00',4150.0000,1,37,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(61,1,'2017-03-01 00:00:00',4150.0000,1,38,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(62,1,'2017-04-01 00:00:00',4150.0000,1,39,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(63,1,'2017-05-01 00:00:00',4150.0000,1,40,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(64,1,'2017-06-01 00:00:00',4150.0000,1,41,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(65,1,'2017-07-01 00:00:00',4150.0000,1,42,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(66,1,'2017-08-01 00:00:00',4150.0000,1,43,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(67,1,'2017-09-01 00:00:00',4150.0000,1,44,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(68,1,'2017-09-14 00:00:00',200.0000,1,45,'','2017-09-15 05:42:55',0,'2017-09-15 05:42:55',0),(69,1,'2017-09-14 00:00:00',125.0000,2,7,'','2017-09-15 05:43:22',0,'2017-09-15 05:43:22',0),(70,1,'2017-09-14 00:00:00',75.0000,2,8,'','2017-09-15 05:43:38',0,'2017-09-15 05:43:38',0),(71,1,'2017-09-14 00:00:00',125.0000,2,7,'','2017-09-15 05:44:13',0,'2017-09-15 05:44:13',0),(72,1,'2017-09-14 00:00:00',75.0000,2,8,'','2017-09-15 05:44:13',0,'2017-09-15 05:44:13',0),(73,1,'2014-03-01 00:00:00',3500.0000,1,46,'','2017-09-15 23:30:54',0,'2017-09-15 23:30:54',0),(74,1,'2014-03-02 00:00:00',3500.0000,1,47,'','2017-09-15 23:41:23',0,'2017-09-15 23:41:23',0),(75,1,'2017-01-01 00:00:00',3000.0000,1,48,'','2017-09-16 00:27:23',0,'2017-09-16 00:27:23',0),(76,1,'2017-09-18 00:00:00',50.0000,1,49,'','2017-09-18 18:07:30',0,'2017-09-18 18:07:30',0),(77,1,'2017-09-18 00:00:00',50.0000,1,50,'','2017-09-18 18:08:54',0,'2017-09-18 18:08:54',0),(78,1,'2017-09-18 00:00:00',50.0000,1,51,'','2017-09-18 18:25:13',0,'2017-09-18 18:25:13',0),(79,1,'2017-09-18 00:00:00',50.0000,1,52,'','2017-09-18 18:25:58',0,'2017-09-18 18:25:58',0),(80,1,'2017-09-18 00:00:00',50.0000,1,53,'','2017-09-18 18:26:24',0,'2017-09-18 18:26:24',0),(81,1,'2017-09-18 00:00:00',3525.0000,2,9,'','2017-09-18 19:49:58',0,'2017-09-18 19:49:58',0),(82,1,'2017-09-03 00:00:00',75.0000,1,54,'','2017-09-25 20:51:01',0,'2017-09-25 20:51:01',0),(83,1,'2017-09-18 00:00:00',-50.0000,1,55,'','2017-09-25 20:51:37',0,'2017-09-25 20:51:37',0),(84,1,'2017-09-25 00:00:00',100.0000,1,56,'','2017-09-25 21:23:20',0,'2017-09-25 21:23:20',0),(85,1,'2017-08-02 00:00:00',5000.0000,1,57,'','2017-09-25 21:25:30',0,'2017-09-25 21:25:30',0),(86,1,'2017-09-25 00:00:00',50.0000,2,10,'','2017-09-26 05:28:20',0,'2017-09-26 05:28:20',0),(87,1,'2017-09-25 00:00:00',50.0000,2,10,'','2017-09-26 05:29:10',0,'2017-09-26 05:29:10',0),(88,1,'2017-09-25 00:00:00',50.0000,2,11,'','2017-09-26 05:30:24',0,'2017-09-26 05:30:24',0),(89,1,'2017-09-25 00:00:00',50.0000,2,12,'','2017-09-26 05:31:32',0,'2017-09-26 05:31:32',0),(90,1,'2017-09-25 00:00:00',50.0000,2,13,'','2017-09-26 05:32:00',0,'2017-09-26 05:32:00',0),(91,1,'2017-09-25 00:00:00',75.0000,2,14,'','2017-09-26 05:32:57',0,'2017-09-26 05:32:57',0),(92,1,'2017-09-25 00:00:00',50.0000,2,11,'','2017-09-26 05:33:06',0,'2017-09-26 05:33:06',0),(93,1,'2017-09-25 00:00:00',50.0000,2,12,'','2017-09-26 05:33:10',0,'2017-09-26 05:33:10',0),(94,1,'2017-09-25 00:00:00',50.0000,2,13,'','2017-09-26 05:33:14',0,'2017-09-26 05:33:14',0),(95,1,'2017-09-25 00:00:00',75.0000,2,14,'','2017-09-26 05:33:18',0,'2017-09-26 05:33:18',0);
/*!40000 ALTER TABLE `Journal` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `JournalAllocation`
--

DROP TABLE IF EXISTS `JournalAllocation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `JournalAllocation` (
  `JAID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `JID` bigint(20) NOT NULL DEFAULT '0',
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `RAID` bigint(20) NOT NULL DEFAULT '0',
  `TCID` bigint(20) NOT NULL DEFAULT '0',
  `RCPTID` bigint(20) NOT NULL DEFAULT '0',
  `Amount` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `ASMID` bigint(20) NOT NULL DEFAULT '0',
  `EXPID` bigint(20) NOT NULL DEFAULT '0',
  `AcctRule` varchar(200) NOT NULL DEFAULT '',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`JAID`)
) ENGINE=InnoDB AUTO_INCREMENT=96 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `JournalAllocation`
--

LOCK TABLES `JournalAllocation` WRITE;
/*!40000 ALTER TABLE `JournalAllocation` DISABLE KEYS */;
INSERT INTO `JournalAllocation` VALUES (1,1,1,2,3,0,0,4000.0000,2,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0),(2,1,2,2,3,0,0,4000.0000,3,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0),(3,1,3,2,3,0,0,4000.0000,4,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0),(4,1,4,2,3,0,0,4000.0000,5,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0),(5,1,5,2,3,0,0,4000.0000,6,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0),(6,1,6,2,3,0,0,4000.0000,7,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0),(7,1,7,2,3,0,0,4000.0000,8,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0),(8,1,8,2,3,0,0,4000.0000,9,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0),(9,1,9,2,3,0,0,4000.0000,10,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0),(10,1,10,2,3,0,0,4000.0000,11,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0),(11,1,11,2,3,0,0,4000.0000,12,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0),(12,1,12,2,3,0,0,628.4500,13,0,'d 12001 628.45, c 41301 628.45','2017-08-22 18:27:44',0),(13,1,13,2,3,0,0,175.0000,14,0,'d 12001 175.00, c 41301 175.00','2017-08-22 18:28:39',0),(14,1,14,0,0,6,0,4000.0000,0,0,'d 10999 _, c 12999 _','2017-08-22 18:29:52',0),(15,1,15,0,0,6,0,12000.0000,0,0,'d 10999 _, c 12999 _','2017-08-22 18:31:17',0),(16,1,16,2,3,0,0,15.0000,0,1,'d 50003 15.00, c 12001 15.00','2017-08-22 18:32:26',0),(17,1,17,2,3,6,1,4000.0000,2,0,'ASM(2) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:40:14',0),(18,1,18,2,3,6,2,4000.0000,3,0,'ASM(3) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:40:14',0),(19,1,19,2,3,6,2,4000.0000,4,0,'ASM(4) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:40:14',0),(20,1,20,2,3,6,2,4000.0000,5,0,'ASM(5) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:40:14',0),(21,1,21,0,0,6,0,628.4500,0,0,'d 10999 _, c 12999 _','2017-08-22 18:43:32',0),(22,1,22,0,0,6,0,8350.0000,0,0,'d 10999 _, c 12999 _','2017-08-22 18:44:27',0),(23,1,23,0,0,6,0,13131.7900,0,0,'d 10999 _, c 12999 _','2017-08-22 18:45:12',0),(24,1,24,0,0,6,0,13050.0000,0,0,'d 10999 _, c 12999 _','2017-08-22 18:45:41',0),(25,1,25,2,3,0,0,175.0000,15,0,'d 12001 175.00, c 41301 175.00','2017-08-22 18:47:56',0),(26,1,26,2,3,0,0,175.0000,16,0,'d 12001 175.00, c 41301 175.00','2017-08-22 18:48:32',0),(27,1,27,2,3,0,0,81.7900,17,0,'d 12001 81.79, c 41301 81.79','2017-08-22 18:49:22',0),(28,1,28,2,3,0,0,350.0000,19,0,'d 12001 350.00, c 41301 350.00','2017-08-22 18:50:08',0),(29,1,29,2,3,0,0,350.0000,20,0,'d 12001 350.00, c 41301 350.00','2017-08-22 18:50:08',0),(30,1,30,2,3,0,0,350.0000,21,0,'d 12001 350.00, c 41301 350.00','2017-08-22 18:50:08',0),(31,1,31,2,3,0,0,350.0000,22,0,'d 12001 350.00, c 41301 350.00','2017-08-22 18:50:08',0),(32,1,32,2,3,0,0,350.0000,23,0,'d 12001 350.00, c 41301 350.00','2017-08-22 18:50:08',0),(33,1,33,2,3,0,0,-175.0000,24,0,'d 12001 -175.00, c 41301 -175.00','2017-08-22 18:52:07',0),(34,1,34,2,3,6,3,628.4500,13,0,'ASM(13) d 12999 628.45,c 12001 628.45','2017-08-22 18:53:58',0),(35,1,35,2,3,6,4,4000.0000,6,0,'ASM(6) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0),(36,1,36,2,3,6,4,175.0000,14,0,'ASM(14) d 12999 175.00,c 12001 175.00','2017-08-22 18:53:58',0),(37,1,37,2,3,6,4,4000.0000,7,0,'ASM(7) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0),(38,1,38,2,3,6,4,175.0000,16,0,'ASM(16) d 12999 175.00,c 12001 175.00','2017-08-22 18:53:58',0),(39,1,39,2,3,6,5,4000.0000,8,0,'ASM(8) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0),(40,1,40,2,3,6,5,81.7900,17,0,'ASM(17) d 12999 81.79,c 12001 81.79','2017-08-22 18:53:58',0),(41,1,41,2,3,6,5,350.0000,19,0,'ASM(19) d 12999 350.00,c 12001 350.00','2017-08-22 18:53:58',0),(42,1,42,2,3,6,5,4000.0000,9,0,'ASM(9) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0),(43,1,43,2,3,6,5,350.0000,20,0,'ASM(20) d 12999 350.00,c 12001 350.00','2017-08-22 18:53:58',0),(44,1,44,2,3,6,5,4000.0000,10,0,'ASM(10) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0),(45,1,45,2,3,6,5,350.0000,21,0,'ASM(21) d 12999 350.00,c 12001 350.00','2017-08-22 18:53:58',0),(46,1,46,2,3,6,6,4000.0000,11,0,'ASM(11) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0),(47,1,47,2,3,6,6,350.0000,22,0,'ASM(22) d 12999 350.00,c 12001 350.00','2017-08-22 18:53:58',0),(48,1,48,2,3,6,6,4000.0000,12,0,'ASM(12) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0),(49,1,49,2,3,6,6,350.0000,23,0,'ASM(23) d 12999 350.00,c 12001 350.00','2017-08-22 18:53:58',0),(50,1,50,2,3,0,0,4000.0000,25,0,'d 12001 4000.00, c 41001 4000.00','2017-09-01 00:00:08',0),(51,1,51,2,3,0,0,350.0000,26,0,'d 12001 350.00, c 41301 350.00','2017-09-01 00:00:08',0),(52,1,52,1,1,0,0,4000.0000,28,0,'d 12001 4000.00, c 41001 4000.00','2017-09-15 04:14:32',0),(53,1,53,1,1,0,0,4000.0000,29,0,'d 12001 4000.00, c 41001 4000.00','2017-09-15 04:14:32',0),(54,1,54,1,1,0,0,4000.0000,30,0,'d 12001 4000.00, c 41001 4000.00','2017-09-15 04:14:32',0),(55,1,55,1,1,0,0,4000.0000,31,0,'d 12001 4000.00, c 41001 4000.00','2017-09-15 04:14:32',0),(56,1,56,1,1,0,0,4000.0000,32,0,'d 12001 4000.00, c 41001 4000.00','2017-09-15 04:14:32',0),(57,1,57,1,1,0,0,4000.0000,33,0,'d 12001 4000.00, c 41001 4000.00','2017-09-15 04:14:32',0),(58,1,58,1,1,0,0,4000.0000,34,0,'d 12001 4000.00, c 41001 4000.00','2017-09-15 04:14:32',0),(59,1,59,3,2,0,0,4150.0000,36,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0),(60,1,60,3,2,0,0,4150.0000,37,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0),(61,1,61,3,2,0,0,4150.0000,38,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0),(62,1,62,3,2,0,0,4150.0000,39,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0),(63,1,63,3,2,0,0,4150.0000,40,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0),(64,1,64,3,2,0,0,4150.0000,41,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0),(65,1,65,3,2,0,0,4150.0000,42,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0),(66,1,66,3,2,0,0,4150.0000,43,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0),(67,1,67,3,2,0,0,4150.0000,44,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0),(68,1,68,1,1,0,0,200.0000,45,0,'d 12001 200.00, c 41302 200.00','2017-09-15 05:42:55',0),(69,1,69,0,0,1,0,125.0000,0,0,'d 10999 _, c 12999 _','2017-09-15 05:43:22',0),(70,1,70,0,0,1,0,75.0000,0,0,'d 10999 _, c 12999 _','2017-09-15 05:43:38',0),(71,1,71,1,1,1,7,125.0000,45,0,'ASM(45) d 12999 125.00,c 12001 125.00','2017-09-15 05:44:13',0),(72,1,72,1,1,1,8,75.0000,45,0,'ASM(45) d 12999 75.00,c 12001 75.00','2017-09-15 05:44:13',0),(73,1,73,1,1,0,0,3500.0000,46,0,'d 12001 3500.00, c 30000 3500.00','2017-09-15 23:30:54',0),(74,1,74,1,1,0,0,3500.0000,47,0,'d 12001 3500.00, c 30000 3500.00','2017-09-15 23:41:23',0),(75,1,75,1,1,0,0,3000.0000,48,0,'d 12001 3000.00, c 30000 3000.00','2017-09-16 00:27:23',0),(76,1,76,0,8,0,0,50.0000,49,0,'d 12001 50.00, c 41401 50.00','2017-09-18 18:07:30',0),(77,1,77,0,9,0,0,50.0000,50,0,'d 12001 50.00, c 41401 50.00','2017-09-18 18:08:54',0),(78,1,78,0,10,0,0,50.0000,51,0,'d 12001 50.00, c 41401 50.00','2017-09-18 18:25:13',0),(79,1,79,0,11,0,0,50.0000,52,0,'d 12001 50.00, c 41401 50.00','2017-09-18 18:25:58',0),(80,1,80,0,12,0,0,50.0000,53,0,'d 12001 50.00, c 41401 50.00','2017-09-18 18:26:24',0),(81,1,81,0,0,14,0,3525.0000,0,0,'d 11000 _, c 42300 _','2017-09-18 19:49:58',0),(82,1,82,0,12,0,0,75.0000,54,0,'d 12001 75.00, c 41499 75.00','2017-09-25 20:51:01',0),(83,1,83,0,12,0,0,-50.0000,55,0,'d 12001 -50.00, c 41401 -50.00','2017-09-25 20:51:37',0),(84,1,84,0,14,0,0,100.0000,56,0,'d 12001 100.00, c 30000 100.00','2017-09-25 21:23:20',0),(85,1,85,0,14,0,0,5000.0000,57,0,'d 12001 5000.00, c 30000 5000.00','2017-09-25 21:25:30',0),(86,1,86,0,0,9,0,50.0000,0,0,'d 10999 _, c 12999 _','2017-09-26 05:28:20',0),(87,1,87,0,8,9,10,50.0000,49,0,'ASM(49) d 12999 50.00,c 12001 50.00','2017-09-26 05:29:10',0),(88,1,88,0,0,10,0,50.0000,0,0,'d 10999 _, c 12999 _','2017-09-26 05:30:24',0),(89,1,89,0,0,11,0,50.0000,0,0,'d 10999 _, c 12999 _','2017-09-26 05:31:32',0),(90,1,90,0,0,12,0,50.0000,0,0,'d 10999 _, c 12999 _','2017-09-26 05:32:00',0),(91,1,91,0,0,13,0,75.0000,0,0,'d 10999 _, c 12999 _','2017-09-26 05:32:57',0),(92,1,92,0,9,10,11,50.0000,50,0,'ASM(50) d 12999 50.00,c 12001 50.00','2017-09-26 05:33:06',0),(93,1,93,0,10,11,12,50.0000,51,0,'ASM(51) d 12999 50.00,c 12001 50.00','2017-09-26 05:33:10',0),(94,1,94,0,11,12,13,50.0000,52,0,'ASM(52) d 12999 50.00,c 12001 50.00','2017-09-26 05:33:14',0),(95,1,95,0,12,13,14,75.0000,54,0,'ASM(54) d 12999 75.00,c 12001 75.00','2017-09-26 05:33:18',0);
/*!40000 ALTER TABLE `JournalAllocation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `JournalAudit`
--

DROP TABLE IF EXISTS `JournalAudit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `JournalAudit` (
  `JID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `UID` mediumint(9) NOT NULL DEFAULT '0',
  `ModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `JournalAudit`
--

LOCK TABLES `JournalAudit` WRITE;
/*!40000 ALTER TABLE `JournalAudit` DISABLE KEYS */;
/*!40000 ALTER TABLE `JournalAudit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `JournalMarker`
--

DROP TABLE IF EXISTS `JournalMarker`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `JournalMarker` (
  `JMID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `State` smallint(6) NOT NULL DEFAULT '0',
  `DtStart` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtStop` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`JMID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `JournalMarker`
--

LOCK TABLES `JournalMarker` WRITE;
/*!40000 ALTER TABLE `JournalMarker` DISABLE KEYS */;
/*!40000 ALTER TABLE `JournalMarker` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `JournalMarkerAudit`
--

DROP TABLE IF EXISTS `JournalMarkerAudit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `JournalMarkerAudit` (
  `JMID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `UID` mediumint(9) NOT NULL DEFAULT '0',
  `ModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `JournalMarkerAudit`
--

LOCK TABLES `JournalMarkerAudit` WRITE;
/*!40000 ALTER TABLE `JournalMarkerAudit` DISABLE KEYS */;
/*!40000 ALTER TABLE `JournalMarkerAudit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `LeadSource`
--

DROP TABLE IF EXISTS `LeadSource`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `LeadSource` (
  `LSID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Name` varchar(100) DEFAULT NULL,
  `IndustrySLID` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`LSID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LeadSource`
--

LOCK TABLES `LeadSource` WRITE;
/*!40000 ALTER TABLE `LeadSource` DISABLE KEYS */;
/*!40000 ALTER TABLE `LeadSource` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `LedgerAudit`
--

DROP TABLE IF EXISTS `LedgerAudit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `LedgerAudit` (
  `LEID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `UID` mediumint(9) NOT NULL DEFAULT '0',
  `ModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerAudit`
--

LOCK TABLES `LedgerAudit` WRITE;
/*!40000 ALTER TABLE `LedgerAudit` DISABLE KEYS */;
/*!40000 ALTER TABLE `LedgerAudit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `LedgerEntry`
--

DROP TABLE IF EXISTS `LedgerEntry`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `LedgerEntry` (
  `LEID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `JID` bigint(20) NOT NULL DEFAULT '0',
  `JAID` bigint(20) NOT NULL DEFAULT '0',
  `LID` bigint(20) NOT NULL DEFAULT '0',
  `RAID` bigint(20) NOT NULL DEFAULT '0',
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `TCID` bigint(20) NOT NULL DEFAULT '0',
  `Dt` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `Amount` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `Comment` varchar(256) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`LEID`)
) ENGINE=InnoDB AUTO_INCREMENT=187 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerEntry`
--

LOCK TABLES `LedgerEntry` WRITE;
/*!40000 ALTER TABLE `LedgerEntry` DISABLE KEYS */;
INSERT INTO `LedgerEntry` VALUES (1,1,1,1,8,3,2,0,'2016-10-01 00:00:00',4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(2,1,1,1,17,3,2,0,'2016-10-01 00:00:00',-4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(3,1,2,2,8,3,2,0,'2016-11-01 00:00:00',4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(4,1,2,2,17,3,2,0,'2016-11-01 00:00:00',-4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(5,1,3,3,8,3,2,0,'2016-12-01 00:00:00',4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(6,1,3,3,17,3,2,0,'2016-12-01 00:00:00',-4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(7,1,4,4,8,3,2,0,'2017-01-01 00:00:00',4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(8,1,4,4,17,3,2,0,'2017-01-01 00:00:00',-4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(9,1,5,5,8,3,2,0,'2017-02-01 00:00:00',4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(10,1,5,5,17,3,2,0,'2017-02-01 00:00:00',-4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(11,1,6,6,8,3,2,0,'2017-03-01 00:00:00',4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(12,1,6,6,17,3,2,0,'2017-03-01 00:00:00',-4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(13,1,7,7,8,3,2,0,'2017-04-01 00:00:00',4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(14,1,7,7,17,3,2,0,'2017-04-01 00:00:00',-4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(15,1,8,8,8,3,2,0,'2017-05-01 00:00:00',4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(16,1,8,8,17,3,2,0,'2017-05-01 00:00:00',-4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(17,1,9,9,8,3,2,0,'2017-06-01 00:00:00',4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(18,1,9,9,17,3,2,0,'2017-06-01 00:00:00',-4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(19,1,10,10,8,3,2,0,'2017-07-01 00:00:00',4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(20,1,10,10,17,3,2,0,'2017-07-01 00:00:00',-4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(21,1,11,11,8,3,2,0,'2017-08-01 00:00:00',4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(22,1,11,11,17,3,2,0,'2017-08-01 00:00:00',-4000.0000,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(23,1,12,12,8,3,2,0,'2017-01-01 00:00:00',628.4500,'','2017-08-22 18:27:44',0,'2017-08-22 18:27:44',0),(24,1,12,12,35,3,2,0,'2017-01-01 00:00:00',-628.4500,'','2017-08-22 18:27:44',0,'2017-08-22 18:27:44',0),(25,1,13,13,8,3,2,0,'2017-02-01 00:00:00',175.0000,'','2017-08-22 18:28:39',0,'2017-08-22 18:28:39',0),(26,1,13,13,35,3,2,0,'2017-02-01 00:00:00',-175.0000,'','2017-08-22 18:28:39',0,'2017-08-22 18:28:39',0),(27,1,14,14,5,0,0,6,'2016-10-03 00:00:00',4000.0000,'','2017-08-22 18:29:52',0,'2017-08-22 18:29:52',0),(28,1,14,14,10,0,0,6,'2016-10-03 00:00:00',-4000.0000,'','2017-08-22 18:29:52',0,'2017-08-22 18:29:52',0),(29,1,15,15,5,0,0,6,'2016-11-11 00:00:00',12000.0000,'','2017-08-22 18:31:17',0,'2017-08-22 18:31:17',0),(30,1,15,15,10,0,0,6,'2016-11-11 00:00:00',-12000.0000,'','2017-08-22 18:31:17',0,'2017-08-22 18:31:17',0),(31,1,16,16,71,3,2,0,'2016-11-11 00:00:00',15.0000,'','2017-08-22 18:32:26',0,'2017-08-22 18:32:26',0),(32,1,16,16,8,3,2,0,'2016-11-11 00:00:00',-15.0000,'','2017-08-22 18:32:26',0,'2017-08-22 18:32:26',0),(33,1,17,17,10,3,2,6,'2016-10-03 00:00:00',4000.0000,'','2017-08-22 18:40:14',0,'2017-08-22 18:40:14',0),(34,1,17,17,8,3,2,6,'2016-10-03 00:00:00',-4000.0000,'','2017-08-22 18:40:14',0,'2017-08-22 18:40:14',0),(35,1,18,18,10,3,2,6,'2016-11-11 00:00:00',4000.0000,'','2017-08-22 18:40:14',0,'2017-08-22 18:40:14',0),(36,1,18,18,8,3,2,6,'2016-11-11 00:00:00',-4000.0000,'','2017-08-22 18:40:14',0,'2017-08-22 18:40:14',0),(37,1,19,19,10,3,2,6,'2016-11-11 00:00:00',4000.0000,'','2017-08-22 18:40:14',0,'2017-08-22 18:40:14',0),(38,1,19,19,8,3,2,6,'2016-11-11 00:00:00',-4000.0000,'','2017-08-22 18:40:14',0,'2017-08-22 18:40:14',0),(39,1,20,20,10,3,2,6,'2016-11-11 00:00:00',4000.0000,'','2017-08-22 18:40:14',0,'2017-08-22 18:40:14',0),(40,1,20,20,8,3,2,6,'2016-11-11 00:00:00',-4000.0000,'','2017-08-22 18:40:14',0,'2017-08-22 18:40:14',0),(41,1,21,21,5,0,0,6,'2017-02-03 00:00:00',628.4500,'','2017-08-22 18:43:32',0,'2017-08-22 18:43:32',0),(42,1,21,21,10,0,0,6,'2017-02-03 00:00:00',-628.4500,'','2017-08-22 18:43:32',0,'2017-08-22 18:43:32',0),(43,1,22,22,5,0,0,6,'2017-02-13 00:00:00',8350.0000,'','2017-08-22 18:44:27',0,'2017-08-22 18:44:27',0),(44,1,22,22,10,0,0,6,'2017-02-13 00:00:00',-8350.0000,'','2017-08-22 18:44:27',0,'2017-08-22 18:44:27',0),(45,1,23,23,5,0,0,6,'2017-05-12 00:00:00',13131.7900,'','2017-08-22 18:45:12',0,'2017-08-22 18:45:12',0),(46,1,23,23,10,0,0,6,'2017-05-12 00:00:00',-13131.7900,'','2017-08-22 18:45:12',0,'2017-08-22 18:45:12',0),(47,1,24,24,5,0,0,6,'2017-08-15 00:00:00',13050.0000,'','2017-08-22 18:45:41',0,'2017-08-22 18:45:41',0),(48,1,24,24,10,0,0,6,'2017-08-15 00:00:00',-13050.0000,'','2017-08-22 18:45:41',0,'2017-08-22 18:45:41',0),(49,1,25,25,8,3,2,0,'2017-02-01 00:00:00',175.0000,'','2017-08-22 18:47:56',0,'2017-08-22 18:47:56',0),(50,1,25,25,35,3,2,0,'2017-02-01 00:00:00',-175.0000,'','2017-08-22 18:47:56',0,'2017-08-22 18:47:56',0),(51,1,26,26,8,3,2,0,'2017-03-01 00:00:00',175.0000,'','2017-08-22 18:48:32',0,'2017-08-22 18:48:32',0),(52,1,26,26,35,3,2,0,'2017-03-01 00:00:00',-175.0000,'','2017-08-22 18:48:32',0,'2017-08-22 18:48:32',0),(53,1,27,27,8,3,2,0,'2017-04-01 00:00:00',81.7900,'','2017-08-22 18:49:22',0,'2017-08-22 18:49:22',0),(54,1,27,27,35,3,2,0,'2017-04-01 00:00:00',-81.7900,'','2017-08-22 18:49:22',0,'2017-08-22 18:49:22',0),(55,1,28,28,8,3,2,0,'2017-04-01 00:00:00',350.0000,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(56,1,28,28,35,3,2,0,'2017-04-01 00:00:00',-350.0000,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(57,1,29,29,8,3,2,0,'2017-05-01 00:00:00',350.0000,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(58,1,29,29,35,3,2,0,'2017-05-01 00:00:00',-350.0000,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(59,1,30,30,8,3,2,0,'2017-06-01 00:00:00',350.0000,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(60,1,30,30,35,3,2,0,'2017-06-01 00:00:00',-350.0000,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(61,1,31,31,8,3,2,0,'2017-07-01 00:00:00',350.0000,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(62,1,31,31,35,3,2,0,'2017-07-01 00:00:00',-350.0000,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(63,1,32,32,8,3,2,0,'2017-08-01 00:00:00',350.0000,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(64,1,32,32,35,3,2,0,'2017-08-01 00:00:00',-350.0000,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(65,1,33,33,8,3,2,0,'2017-02-01 00:00:00',-175.0000,'','2017-08-22 18:52:07',0,'2017-08-22 18:52:07',0),(66,1,33,33,35,3,2,0,'2017-02-01 00:00:00',175.0000,'','2017-08-22 18:52:07',0,'2017-08-22 18:52:07',0),(67,1,34,34,10,3,2,6,'2017-02-03 00:00:00',628.4500,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(68,1,34,34,8,3,2,6,'2017-02-03 00:00:00',-628.4500,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(69,1,35,35,10,3,2,6,'2017-02-13 00:00:00',4000.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(70,1,35,35,8,3,2,6,'2017-02-13 00:00:00',-4000.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(71,1,36,36,10,3,2,6,'2017-02-13 00:00:00',175.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(72,1,36,36,8,3,2,6,'2017-02-13 00:00:00',-175.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(73,1,37,37,10,3,2,6,'2017-02-13 00:00:00',4000.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(74,1,37,37,8,3,2,6,'2017-02-13 00:00:00',-4000.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(75,1,38,38,10,3,2,6,'2017-02-13 00:00:00',175.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(76,1,38,38,8,3,2,6,'2017-02-13 00:00:00',-175.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(77,1,39,39,10,3,2,6,'2017-05-12 00:00:00',4000.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(78,1,39,39,8,3,2,6,'2017-05-12 00:00:00',-4000.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(79,1,40,40,10,3,2,6,'2017-05-12 00:00:00',81.7900,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(80,1,40,40,8,3,2,6,'2017-05-12 00:00:00',-81.7900,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(81,1,41,41,10,3,2,6,'2017-05-12 00:00:00',350.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(82,1,41,41,8,3,2,6,'2017-05-12 00:00:00',-350.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(83,1,42,42,10,3,2,6,'2017-05-12 00:00:00',4000.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(84,1,42,42,8,3,2,6,'2017-05-12 00:00:00',-4000.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(85,1,43,43,10,3,2,6,'2017-05-12 00:00:00',350.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(86,1,43,43,8,3,2,6,'2017-05-12 00:00:00',-350.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(87,1,44,44,10,3,2,6,'2017-05-12 00:00:00',4000.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(88,1,44,44,8,3,2,6,'2017-05-12 00:00:00',-4000.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(89,1,45,45,10,3,2,6,'2017-05-12 00:00:00',350.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(90,1,45,45,8,3,2,6,'2017-05-12 00:00:00',-350.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(91,1,46,46,10,3,2,6,'2017-08-15 00:00:00',4000.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(92,1,46,46,8,3,2,6,'2017-08-15 00:00:00',-4000.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(93,1,47,47,10,3,2,6,'2017-08-15 00:00:00',350.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(94,1,47,47,8,3,2,6,'2017-08-15 00:00:00',-350.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(95,1,48,48,10,3,2,6,'2017-08-15 00:00:00',4000.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(96,1,48,48,8,3,2,6,'2017-08-15 00:00:00',-4000.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(97,1,49,49,10,3,2,6,'2017-08-15 00:00:00',350.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(98,1,49,49,8,3,2,6,'2017-08-15 00:00:00',-350.0000,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(99,1,52,52,8,1,1,0,'2017-03-01 00:00:00',4000.0000,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(100,1,52,52,17,1,1,0,'2017-03-01 00:00:00',-4000.0000,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(101,1,53,53,8,1,1,0,'2017-04-01 00:00:00',4000.0000,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(102,1,53,53,17,1,1,0,'2017-04-01 00:00:00',-4000.0000,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(103,1,54,54,8,1,1,0,'2017-05-01 00:00:00',4000.0000,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(104,1,54,54,17,1,1,0,'2017-05-01 00:00:00',-4000.0000,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(105,1,55,55,8,1,1,0,'2017-06-01 00:00:00',4000.0000,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(106,1,55,55,17,1,1,0,'2017-06-01 00:00:00',-4000.0000,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(107,1,56,56,8,1,1,0,'2017-07-01 00:00:00',4000.0000,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(108,1,56,56,17,1,1,0,'2017-07-01 00:00:00',-4000.0000,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(109,1,57,57,8,1,1,0,'2017-08-01 00:00:00',4000.0000,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(110,1,57,57,17,1,1,0,'2017-08-01 00:00:00',-4000.0000,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(111,1,58,58,8,1,1,0,'2017-09-01 00:00:00',4000.0000,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(112,1,58,58,17,1,1,0,'2017-09-01 00:00:00',-4000.0000,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(113,1,59,59,8,2,3,0,'2017-01-01 00:00:00',4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(114,1,59,59,17,2,3,0,'2017-01-01 00:00:00',-4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(115,1,60,60,8,2,3,0,'2017-02-01 00:00:00',4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(116,1,60,60,17,2,3,0,'2017-02-01 00:00:00',-4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(117,1,61,61,8,2,3,0,'2017-03-01 00:00:00',4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(118,1,61,61,17,2,3,0,'2017-03-01 00:00:00',-4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(119,1,62,62,8,2,3,0,'2017-04-01 00:00:00',4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(120,1,62,62,17,2,3,0,'2017-04-01 00:00:00',-4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(121,1,63,63,8,2,3,0,'2017-05-01 00:00:00',4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(122,1,63,63,17,2,3,0,'2017-05-01 00:00:00',-4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(123,1,64,64,8,2,3,0,'2017-06-01 00:00:00',4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(124,1,64,64,17,2,3,0,'2017-06-01 00:00:00',-4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(125,1,65,65,8,2,3,0,'2017-07-01 00:00:00',4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(126,1,65,65,17,2,3,0,'2017-07-01 00:00:00',-4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(127,1,66,66,8,2,3,0,'2017-08-01 00:00:00',4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(128,1,66,66,17,2,3,0,'2017-08-01 00:00:00',-4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(129,1,67,67,8,2,3,0,'2017-09-01 00:00:00',4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(130,1,67,67,17,2,3,0,'2017-09-01 00:00:00',-4150.0000,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(131,1,68,68,8,1,1,0,'2017-09-14 00:00:00',200.0000,'','2017-09-15 05:42:55',0,'2017-09-15 05:42:55',0),(132,1,68,68,36,1,1,0,'2017-09-14 00:00:00',-200.0000,'','2017-09-15 05:42:55',0,'2017-09-15 05:42:55',0),(133,1,69,69,5,0,0,1,'2017-09-14 00:00:00',125.0000,'','2017-09-15 05:43:22',0,'2017-09-15 05:43:22',0),(134,1,69,69,10,0,0,1,'2017-09-14 00:00:00',-125.0000,'','2017-09-15 05:43:22',0,'2017-09-15 05:43:22',0),(135,1,70,70,5,0,0,1,'2017-09-14 00:00:00',75.0000,'','2017-09-15 05:43:38',0,'2017-09-15 05:43:38',0),(136,1,70,70,10,0,0,1,'2017-09-14 00:00:00',-75.0000,'','2017-09-15 05:43:38',0,'2017-09-15 05:43:38',0),(137,1,71,71,10,1,1,1,'2017-09-14 00:00:00',125.0000,'','2017-09-15 05:44:13',0,'2017-09-15 05:44:13',0),(138,1,71,71,8,1,1,1,'2017-09-14 00:00:00',-125.0000,'','2017-09-15 05:44:13',0,'2017-09-15 05:44:13',0),(139,1,72,72,10,1,1,1,'2017-09-14 00:00:00',75.0000,'','2017-09-15 05:44:13',0,'2017-09-15 05:44:13',0),(140,1,72,72,8,1,1,1,'2017-09-14 00:00:00',-75.0000,'','2017-09-15 05:44:13',0,'2017-09-15 05:44:13',0),(141,1,73,73,8,1,1,0,'2014-03-01 00:00:00',3500.0000,'','2017-09-15 23:30:54',0,'2017-09-15 23:30:54',0),(142,1,73,73,11,1,1,0,'2014-03-01 00:00:00',-3500.0000,'','2017-09-15 23:30:54',0,'2017-09-15 23:30:54',0),(143,1,74,74,8,1,1,0,'2014-03-02 00:00:00',3500.0000,'','2017-09-15 23:41:23',0,'2017-09-15 23:41:23',0),(144,1,74,74,11,1,1,0,'2014-03-02 00:00:00',-3500.0000,'','2017-09-15 23:41:23',0,'2017-09-15 23:41:23',0),(145,1,75,75,8,1,1,0,'2017-01-01 00:00:00',3000.0000,'','2017-09-16 00:27:23',0,'2017-09-16 00:27:23',0),(146,1,75,75,11,1,1,0,'2017-01-01 00:00:00',-3000.0000,'','2017-09-16 00:27:23',0,'2017-09-16 00:27:23',0),(147,1,76,76,8,8,0,0,'2017-09-18 00:00:00',50.0000,'','2017-09-18 18:07:30',0,'2017-09-18 18:07:30',0),(148,1,76,76,45,8,0,0,'2017-09-18 00:00:00',-50.0000,'','2017-09-18 18:07:30',0,'2017-09-18 18:07:30',0),(149,1,77,77,8,9,0,0,'2017-09-18 00:00:00',50.0000,'','2017-09-18 18:08:54',0,'2017-09-18 18:08:54',0),(150,1,77,77,45,9,0,0,'2017-09-18 00:00:00',-50.0000,'','2017-09-18 18:08:54',0,'2017-09-18 18:08:54',0),(151,1,78,78,8,10,0,0,'2017-09-18 00:00:00',50.0000,'','2017-09-18 18:25:13',0,'2017-09-18 18:25:13',0),(152,1,78,78,45,10,0,0,'2017-09-18 00:00:00',-50.0000,'','2017-09-18 18:25:13',0,'2017-09-18 18:25:13',0),(153,1,79,79,8,11,0,0,'2017-09-18 00:00:00',50.0000,'','2017-09-18 18:25:58',0,'2017-09-18 18:25:58',0),(154,1,79,79,45,11,0,0,'2017-09-18 00:00:00',-50.0000,'','2017-09-18 18:25:58',0,'2017-09-18 18:25:58',0),(155,1,80,80,8,12,0,0,'2017-09-18 00:00:00',50.0000,'','2017-09-18 18:26:24',0,'2017-09-18 18:26:24',0),(156,1,80,80,45,12,0,0,'2017-09-18 00:00:00',-50.0000,'','2017-09-18 18:26:24',0,'2017-09-18 18:26:24',0),(157,1,81,81,6,0,0,14,'2017-09-18 00:00:00',3525.0000,'','2017-09-18 19:49:58',0,'2017-09-18 19:49:58',0),(158,1,81,81,64,0,0,14,'2017-09-18 00:00:00',-3525.0000,'','2017-09-18 19:49:58',0,'2017-09-18 19:49:58',0),(159,1,82,82,8,12,0,0,'2017-09-03 00:00:00',75.0000,'','2017-09-25 20:51:01',0,'2017-09-25 20:51:01',0),(160,1,82,82,60,12,0,0,'2017-09-03 00:00:00',-75.0000,'','2017-09-25 20:51:01',0,'2017-09-25 20:51:01',0),(161,1,83,83,8,12,0,0,'2017-09-18 00:00:00',-50.0000,'','2017-09-25 20:51:37',0,'2017-09-25 20:51:37',0),(162,1,83,83,45,12,0,0,'2017-09-18 00:00:00',50.0000,'','2017-09-25 20:51:37',0,'2017-09-25 20:51:37',0),(163,1,84,84,8,14,0,0,'2017-09-25 00:00:00',100.0000,'','2017-09-25 21:23:20',0,'2017-09-25 21:23:20',0),(164,1,84,84,11,14,0,0,'2017-09-25 00:00:00',-100.0000,'','2017-09-25 21:23:20',0,'2017-09-25 21:23:20',0),(165,1,85,85,8,14,0,0,'2017-08-02 00:00:00',5000.0000,'','2017-09-25 21:25:30',0,'2017-09-25 21:25:30',0),(166,1,85,85,11,14,0,0,'2017-08-02 00:00:00',-5000.0000,'','2017-09-25 21:25:30',0,'2017-09-25 21:25:30',0),(167,1,86,86,5,0,0,9,'2017-09-25 00:00:00',50.0000,'','2017-09-26 05:28:20',0,'2017-09-26 05:28:20',0),(168,1,86,86,10,0,0,9,'2017-09-25 00:00:00',-50.0000,'','2017-09-26 05:28:20',0,'2017-09-26 05:28:20',0),(169,1,87,87,10,8,0,9,'2017-09-25 00:00:00',50.0000,'','2017-09-26 05:29:10',0,'2017-09-26 05:29:10',0),(170,1,87,87,8,8,0,9,'2017-09-25 00:00:00',-50.0000,'','2017-09-26 05:29:10',0,'2017-09-26 05:29:10',0),(171,1,88,88,5,0,0,10,'2017-09-25 00:00:00',50.0000,'','2017-09-26 05:30:24',0,'2017-09-26 05:30:24',0),(172,1,88,88,10,0,0,10,'2017-09-25 00:00:00',-50.0000,'','2017-09-26 05:30:24',0,'2017-09-26 05:30:24',0),(173,1,89,89,5,0,0,11,'2017-09-25 00:00:00',50.0000,'','2017-09-26 05:31:32',0,'2017-09-26 05:31:32',0),(174,1,89,89,10,0,0,11,'2017-09-25 00:00:00',-50.0000,'','2017-09-26 05:31:32',0,'2017-09-26 05:31:32',0),(175,1,90,90,5,0,0,12,'2017-09-25 00:00:00',50.0000,'','2017-09-26 05:32:00',0,'2017-09-26 05:32:00',0),(176,1,90,90,10,0,0,12,'2017-09-25 00:00:00',-50.0000,'','2017-09-26 05:32:00',0,'2017-09-26 05:32:00',0),(177,1,91,91,5,0,0,13,'2017-09-25 00:00:00',75.0000,'','2017-09-26 05:32:57',0,'2017-09-26 05:32:57',0),(178,1,91,91,10,0,0,13,'2017-09-25 00:00:00',-75.0000,'','2017-09-26 05:32:57',0,'2017-09-26 05:32:57',0),(179,1,92,92,10,9,0,10,'2017-09-25 00:00:00',50.0000,'','2017-09-26 05:33:06',0,'2017-09-26 05:33:06',0),(180,1,92,92,8,9,0,10,'2017-09-25 00:00:00',-50.0000,'','2017-09-26 05:33:06',0,'2017-09-26 05:33:06',0),(181,1,93,93,10,10,0,11,'2017-09-25 00:00:00',50.0000,'','2017-09-26 05:33:10',0,'2017-09-26 05:33:10',0),(182,1,93,93,8,10,0,11,'2017-09-25 00:00:00',-50.0000,'','2017-09-26 05:33:10',0,'2017-09-26 05:33:10',0),(183,1,94,94,10,11,0,12,'2017-09-25 00:00:00',50.0000,'','2017-09-26 05:33:14',0,'2017-09-26 05:33:14',0),(184,1,94,94,8,11,0,12,'2017-09-25 00:00:00',-50.0000,'','2017-09-26 05:33:14',0,'2017-09-26 05:33:14',0),(185,1,95,95,10,12,0,13,'2017-09-25 00:00:00',75.0000,'','2017-09-26 05:33:18',0,'2017-09-26 05:33:18',0),(186,1,95,95,8,12,0,13,'2017-09-25 00:00:00',-75.0000,'','2017-09-26 05:33:18',0,'2017-09-26 05:33:18',0);
/*!40000 ALTER TABLE `LedgerEntry` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `LedgerMarker`
--

DROP TABLE IF EXISTS `LedgerMarker`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `LedgerMarker` (
  `LMID` bigint(20) NOT NULL AUTO_INCREMENT,
  `LID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RAID` bigint(20) NOT NULL DEFAULT '0',
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `TCID` bigint(20) NOT NULL DEFAULT '0',
  `Dt` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `Balance` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `State` smallint(6) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`LMID`)
) ENGINE=InnoDB AUTO_INCREMENT=104 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerMarker`
--

LOCK TABLES `LedgerMarker` WRITE;
/*!40000 ALTER TABLE `LedgerMarker` DISABLE KEYS */;
INSERT INTO `LedgerMarker` VALUES (1,1,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(2,2,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(3,3,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(4,4,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(5,5,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(6,6,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(7,7,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(8,8,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(10,10,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(11,11,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(12,12,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(13,13,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(14,14,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(15,15,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(16,16,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(17,17,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(18,18,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(19,19,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(20,20,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(21,21,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(22,22,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(23,23,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(24,24,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(25,25,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(26,26,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(27,27,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(28,28,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(29,29,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(30,30,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(31,31,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(32,32,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(33,33,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(34,34,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(35,35,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(36,36,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(37,37,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(38,38,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(39,39,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(40,40,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(41,41,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(42,42,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(43,43,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(44,44,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(45,45,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(46,46,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(47,47,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(48,48,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(49,49,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(50,50,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(51,51,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(52,52,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(53,53,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(54,54,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(55,55,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(56,56,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(57,57,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(58,58,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(59,59,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(60,60,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(61,61,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(62,62,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(63,63,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(64,64,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(65,65,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(66,66,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(67,67,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(68,68,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(69,69,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(70,70,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(71,71,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(72,72,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(73,0,0,3,0,0,'2016-10-01 00:00:00',0.0000,3,'2017-09-14 21:22:55',0,'2017-08-22 18:38:08',0),(74,0,0,1,0,0,'2014-03-01 00:00:00',0.0000,3,'2017-09-14 21:22:55',0,'2017-08-22 20:14:08',0),(75,0,0,2,0,0,'2016-07-01 00:00:00',0.0000,3,'2017-09-14 21:22:55',0,'2017-08-22 20:14:44',0),(76,0,1,0,0,7,'1970-01-01 00:00:00',0.0000,3,'2017-09-15 03:54:33',0,'2017-09-15 03:54:33',0),(77,0,1,0,0,8,'1970-01-01 00:00:00',0.0000,3,'2017-09-15 03:55:54',0,'2017-09-15 03:55:54',0),(78,0,0,4,0,0,'2017-09-01 00:00:00',0.0000,3,'2017-09-15 04:51:58',0,'2017-09-15 03:57:13',0),(79,0,0,5,0,0,'2017-09-14 00:00:00',0.0000,3,'2017-09-15 03:59:26',0,'2017-09-15 03:59:26',0),(80,0,1,0,0,9,'1970-01-01 00:00:00',0.0000,3,'2017-09-18 02:53:22',0,'2017-09-18 02:53:22',0),(81,0,1,0,0,10,'1970-01-01 00:00:00',0.0000,3,'2017-09-18 02:53:56',0,'2017-09-18 02:53:56',0),(82,0,1,0,0,11,'1970-01-01 00:00:00',0.0000,3,'2017-09-18 02:55:19',0,'2017-09-18 02:55:19',0),(83,0,1,0,0,12,'1970-01-01 00:00:00',0.0000,3,'2017-09-18 02:56:13',0,'2017-09-18 02:56:13',0),(84,0,1,0,0,13,'1970-01-01 00:00:00',0.0000,3,'2017-09-18 02:56:42',0,'2017-09-18 02:56:42',0),(85,0,0,6,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-09-18 02:58:00',0,'2017-09-18 02:57:26',0),(86,0,0,7,0,0,'2017-09-17 00:00:00',0.0000,3,'2017-09-18 02:59:35',0,'2017-09-18 02:59:35',0),(87,0,0,8,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-09-18 03:30:08',0,'2017-09-18 03:28:34',0),(88,0,0,9,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-09-18 03:37:06',0,'2017-09-18 03:31:00',0),(89,0,0,10,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-09-18 03:38:01',0,'2017-09-18 03:37:25',0),(90,0,0,11,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-09-18 03:40:17',0,'2017-09-18 03:39:12',0),(91,0,0,12,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-09-18 03:40:59',0,'2017-09-18 03:40:34',0),(92,0,0,13,0,0,'2017-09-17 00:00:00',0.0000,3,'2017-09-18 03:41:56',0,'2017-09-18 03:41:56',0),(93,0,1,0,0,14,'1970-01-01 00:00:00',0.0000,3,'2017-09-18 19:33:38',0,'2017-09-18 19:33:38',0),(94,0,1,1,1,0,'2014-03-01 00:00:00',0.0000,3,'2017-09-20 17:30:10',0,'2017-09-20 17:30:10',0),(95,0,1,3,2,0,'2016-10-01 00:00:00',0.0000,3,'2017-09-20 17:31:05',0,'2017-09-20 17:31:05',0),(96,0,1,2,3,0,'2016-07-01 00:00:00',0.0000,3,'2017-09-20 17:31:36',0,'2017-09-20 17:31:36',0),(97,0,1,4,5,0,'2017-09-14 00:00:00',0.0000,3,'2017-09-20 17:32:17',0,'2017-09-20 17:32:17',0),(98,0,1,5,6,0,'2017-10-01 00:00:00',0.0000,3,'2017-09-20 17:32:48',0,'2017-09-20 17:32:48',0),(99,73,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-09-21 02:19:32',0,'2017-09-21 02:19:32',0),(100,74,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-09-22 20:59:16',0,'2017-09-22 20:59:16',0),(101,75,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-09-22 21:19:03',0,'2017-09-22 21:19:03',0),(102,0,1,0,0,15,'1970-01-01 00:00:00',0.0000,3,'2017-09-25 21:06:26',0,'2017-09-25 21:06:26',0),(103,0,0,14,0,0,'2017-08-01 00:00:00',0.0000,3,'2017-09-25 21:23:51',0,'2017-09-25 21:07:18',0);
/*!40000 ALTER TABLE `LedgerMarker` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `LedgerMarkerAudit`
--

DROP TABLE IF EXISTS `LedgerMarkerAudit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `LedgerMarkerAudit` (
  `LMID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `UID` mediumint(9) NOT NULL DEFAULT '0',
  `ModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerMarkerAudit`
--

LOCK TABLES `LedgerMarkerAudit` WRITE;
/*!40000 ALTER TABLE `LedgerMarkerAudit` DISABLE KEYS */;
/*!40000 ALTER TABLE `LedgerMarkerAudit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `NoteList`
--

DROP TABLE IF EXISTS `NoteList`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `NoteList` (
  `NLID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`NLID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `NoteList`
--

LOCK TABLES `NoteList` WRITE;
/*!40000 ALTER TABLE `NoteList` DISABLE KEYS */;
/*!40000 ALTER TABLE `NoteList` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `NoteType`
--

DROP TABLE IF EXISTS `NoteType`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `NoteType` (
  `NTID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Name` varchar(128) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`NTID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `NoteType`
--

LOCK TABLES `NoteType` WRITE;
/*!40000 ALTER TABLE `NoteType` DISABLE KEYS */;
/*!40000 ALTER TABLE `NoteType` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Notes`
--

DROP TABLE IF EXISTS `Notes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Notes` (
  `NID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `NLID` bigint(20) NOT NULL DEFAULT '0',
  `PNID` bigint(20) NOT NULL DEFAULT '0',
  `NTID` bigint(20) NOT NULL DEFAULT '0',
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `RAID` bigint(20) NOT NULL DEFAULT '0',
  `TCID` bigint(20) NOT NULL DEFAULT '0',
  `Comment` varchar(1024) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`NID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Notes`
--

LOCK TABLES `Notes` WRITE;
/*!40000 ALTER TABLE `Notes` DISABLE KEYS */;
/*!40000 ALTER TABLE `Notes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `OtherDeliverables`
--

DROP TABLE IF EXISTS `OtherDeliverables`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `OtherDeliverables` (
  `ODID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Name` varchar(256) DEFAULT NULL,
  `Active` smallint(6) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`ODID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `OtherDeliverables`
--

LOCK TABLES `OtherDeliverables` WRITE;
/*!40000 ALTER TABLE `OtherDeliverables` DISABLE KEYS */;
/*!40000 ALTER TABLE `OtherDeliverables` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `PaymentType`
--

DROP TABLE IF EXISTS `PaymentType`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `PaymentType` (
  `PMTID` mediumint(9) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL,
  `Name` varchar(100) NOT NULL DEFAULT '',
  `Description` varchar(256) NOT NULL DEFAULT '',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`PMTID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PaymentType`
--

LOCK TABLES `PaymentType` WRITE;
/*!40000 ALTER TABLE `PaymentType` DISABLE KEYS */;
INSERT INTO `PaymentType` VALUES (1,1,'Check','','2017-08-16 22:05:41',0,'2017-08-16 22:05:41',0),(2,1,'ACH','','2017-08-16 22:05:46',0,'2017-08-16 22:05:46',0),(3,1,'Money Order','','2017-08-16 22:05:53',0,'2017-08-16 22:05:53',0),(4,1,'Wire','','2017-08-16 22:05:57',0,'2017-08-16 22:05:57',0),(5,1,'Credit Card','','2017-08-16 22:06:09',0,'2017-08-16 22:06:09',0),(6,1,'Cash','limited to certain vending income sources','2017-08-16 23:03:40',0,'2017-08-16 23:03:40',0);
/*!40000 ALTER TABLE `PaymentType` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Payor`
--

DROP TABLE IF EXISTS `Payor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Payor` (
  `TCID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `TaxpayorID` varchar(25) NOT NULL DEFAULT '',
  `CreditLimit` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `AccountRep` bigint(20) NOT NULL DEFAULT '0',
  `EligibleFuturePayor` smallint(6) NOT NULL DEFAULT '1',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`TCID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Payor`
--

LOCK TABLES `Payor` WRITE;
/*!40000 ALTER TABLE `Payor` DISABLE KEYS */;
INSERT INTO `Payor` VALUES (1,1,'',0.0000,0,1,'2017-08-16 14:38:42',0,'2017-08-16 14:38:42',0),(2,1,'',0.0000,0,1,'2017-08-16 14:40:13',0,'2017-08-16 14:40:13',0),(3,1,'',0.0000,0,1,'2017-08-16 14:41:42',0,'2017-08-16 14:41:42',0),(4,1,'',0.0000,0,1,'2017-08-16 14:43:10',0,'2017-08-16 14:43:10',0),(5,1,'',0.0000,0,1,'2017-08-16 14:47:19',0,'2017-08-16 14:47:19',0),(6,1,'',0.0000,0,1,'2017-08-16 14:49:36',0,'2017-08-16 14:49:36',0),(7,1,'',0.0000,0,1,'2017-09-15 03:54:33',0,'2017-09-15 03:54:33',0),(8,1,'',0.0000,0,1,'2017-09-15 03:55:54',0,'2017-09-15 03:55:54',0),(9,1,'',0.0000,0,1,'2017-09-18 02:53:22',0,'2017-09-18 02:53:22',0),(10,1,'',0.0000,0,1,'2017-09-18 02:53:56',0,'2017-09-18 02:53:56',0),(11,1,'',0.0000,0,1,'2017-09-18 02:55:19',0,'2017-09-18 02:55:19',0),(12,1,'',0.0000,0,1,'2017-09-18 02:56:13',0,'2017-09-18 02:56:13',0),(13,1,'',0.0000,0,1,'2017-09-18 02:56:42',0,'2017-09-18 02:56:42',0),(14,1,'',0.0000,0,1,'2017-09-18 19:33:38',0,'2017-09-18 19:33:38',0),(15,1,'',0.0000,0,1,'2017-09-25 21:06:26',0,'2017-09-25 21:06:26',0);
/*!40000 ALTER TABLE `Payor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Prospect`
--

DROP TABLE IF EXISTS `Prospect`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Prospect` (
  `TCID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `EmployerName` varchar(100) NOT NULL DEFAULT '',
  `EmployerStreetAddress` varchar(100) NOT NULL DEFAULT '',
  `EmployerCity` varchar(100) NOT NULL DEFAULT '',
  `EmployerState` varchar(100) NOT NULL DEFAULT '',
  `EmployerPostalCode` varchar(100) NOT NULL DEFAULT '',
  `EmployerEmail` varchar(100) NOT NULL DEFAULT '',
  `EmployerPhone` varchar(100) NOT NULL DEFAULT '',
  `Occupation` varchar(100) NOT NULL DEFAULT '',
  `ApplicationFee` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `DesiredUsageStartDate` date NOT NULL DEFAULT '1970-01-01',
  `RentableTypePreference` bigint(20) NOT NULL DEFAULT '0',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `Approver` bigint(20) NOT NULL DEFAULT '0',
  `DeclineReasonSLSID` bigint(20) NOT NULL DEFAULT '0',
  `OtherPreferences` varchar(1024) NOT NULL DEFAULT '',
  `FollowUpDate` date NOT NULL DEFAULT '1970-01-01',
  `CSAgent` bigint(20) NOT NULL DEFAULT '0',
  `OutcomeSLSID` bigint(20) NOT NULL DEFAULT '0',
  `FloatingDeposit` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `RAID` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`TCID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Prospect`
--

LOCK TABLES `Prospect` WRITE;
/*!40000 ALTER TABLE `Prospect` DISABLE KEYS */;
INSERT INTO `Prospect` VALUES (1,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-08-16 14:38:42',0,'2017-08-16 14:38:42',0),(2,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-08-16 14:40:13',0,'2017-08-16 14:40:13',0),(3,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-08-16 14:41:42',0,'2017-08-16 14:41:42',0),(4,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-08-16 14:43:10',0,'2017-08-16 14:43:10',0),(5,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-08-16 14:47:19',0,'2017-08-16 14:47:19',0),(6,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-08-16 14:49:36',0,'2017-08-16 14:49:36',0),(7,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-09-15 03:54:33',0,'2017-09-15 03:54:33',0),(8,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-09-15 03:55:54',0,'2017-09-15 03:55:54',0),(9,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-09-18 02:53:22',0,'2017-09-18 02:53:22',0),(10,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-09-18 02:53:56',0,'2017-09-18 02:53:56',0),(11,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-09-18 02:55:19',0,'2017-09-18 02:55:19',0),(12,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-09-18 02:56:13',0,'2017-09-18 02:56:13',0),(13,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-09-18 02:56:42',0,'2017-09-18 02:56:42',0),(14,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-09-18 19:33:38',0,'2017-09-18 19:33:38',0),(15,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-09-25 21:06:26',0,'2017-09-25 21:06:26',0);
/*!40000 ALTER TABLE `Prospect` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RatePlan`
--

DROP TABLE IF EXISTS `RatePlan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RatePlan` (
  `RPID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Name` varchar(100) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RPID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RatePlan`
--

LOCK TABLES `RatePlan` WRITE;
/*!40000 ALTER TABLE `RatePlan` DISABLE KEYS */;
/*!40000 ALTER TABLE `RatePlan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RatePlanOD`
--

DROP TABLE IF EXISTS `RatePlanOD`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RatePlanOD` (
  `RPRID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `ODID` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RatePlanOD`
--

LOCK TABLES `RatePlanOD` WRITE;
/*!40000 ALTER TABLE `RatePlanOD` DISABLE KEYS */;
/*!40000 ALTER TABLE `RatePlanOD` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RatePlanRef`
--

DROP TABLE IF EXISTS `RatePlanRef`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RatePlanRef` (
  `RPRID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RPID` bigint(20) NOT NULL DEFAULT '0',
  `DtStart` date DEFAULT '1970-01-01',
  `DtStop` date DEFAULT '1970-01-01',
  `FeeAppliesAge` smallint(6) NOT NULL DEFAULT '0',
  `MaxNoFeeUsers` smallint(6) NOT NULL DEFAULT '0',
  `AdditionalUserFee` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `PromoCode` varchar(100) DEFAULT NULL,
  `CancellationFee` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RPRID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RatePlanRef`
--

LOCK TABLES `RatePlanRef` WRITE;
/*!40000 ALTER TABLE `RatePlanRef` DISABLE KEYS */;
/*!40000 ALTER TABLE `RatePlanRef` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RatePlanRefRTRate`
--

DROP TABLE IF EXISTS `RatePlanRefRTRate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RatePlanRefRTRate` (
  `RPRID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RTID` bigint(20) NOT NULL DEFAULT '0',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `Val` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RatePlanRefRTRate`
--

LOCK TABLES `RatePlanRefRTRate` WRITE;
/*!40000 ALTER TABLE `RatePlanRefRTRate` DISABLE KEYS */;
/*!40000 ALTER TABLE `RatePlanRefRTRate` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RatePlanRefSPRate`
--

DROP TABLE IF EXISTS `RatePlanRefSPRate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RatePlanRefSPRate` (
  `RPRID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RTID` bigint(20) NOT NULL DEFAULT '0',
  `RSPID` bigint(20) NOT NULL DEFAULT '0',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `Val` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RatePlanRefSPRate`
--

LOCK TABLES `RatePlanRefSPRate` WRITE;
/*!40000 ALTER TABLE `RatePlanRefSPRate` DISABLE KEYS */;
/*!40000 ALTER TABLE `RatePlanRefSPRate` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Receipt`
--

DROP TABLE IF EXISTS `Receipt`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Receipt` (
  `RCPTID` bigint(20) NOT NULL AUTO_INCREMENT,
  `PRCPTID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `TCID` bigint(20) NOT NULL DEFAULT '0',
  `PMTID` bigint(20) NOT NULL DEFAULT '0',
  `DEPID` bigint(20) NOT NULL DEFAULT '0',
  `DID` bigint(20) NOT NULL DEFAULT '0',
  `Dt` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DocNo` varchar(50) NOT NULL DEFAULT '',
  `Amount` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `AcctRuleReceive` varchar(215) NOT NULL DEFAULT '',
  `ARID` bigint(20) NOT NULL DEFAULT '0',
  `AcctRuleApply` varchar(4096) NOT NULL DEFAULT '',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `Comment` varchar(256) NOT NULL DEFAULT '',
  `OtherPayorName` varchar(128) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RCPTID`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Receipt`
--

LOCK TABLES `Receipt` WRITE;
/*!40000 ALTER TABLE `Receipt` DISABLE KEYS */;
INSERT INTO `Receipt` VALUES (1,0,1,6,1,0,0,'2016-10-03 00:00:00','',4000.0000,'',27,'ASM(2) d 12999 4000.00,c 12001 4000.00',2,'','','2017-08-22 18:40:14',0,'2017-08-22 18:29:52',0),(2,0,1,6,4,0,0,'2016-11-11 00:00:00','',12000.0000,'',27,'ASM(3) d 12999 4000.00,c 12001 4000.00,ASM(4) d 12999 4000.00,c 12001 4000.00,ASM(5) d 12999 4000.00,c 12001 4000.00',2,'','','2017-08-22 18:40:14',0,'2017-08-22 18:31:17',0),(3,0,1,6,4,0,0,'2017-02-03 00:00:00','',628.4500,'',27,'ASM(13) d 12999 628.45,c 12001 628.45',2,'utilities','','2017-08-22 18:53:58',0,'2017-08-22 18:43:32',0),(4,0,1,6,4,0,0,'2017-02-13 00:00:00','',8350.0000,'',27,'ASM(6) d 12999 4000.00,c 12001 4000.00,ASM(14) d 12999 175.00,c 12001 175.00,ASM(7) d 12999 4000.00,c 12001 4000.00,ASM(16) d 12999 175.00,c 12001 175.00',2,'2 month rent and utilities','','2017-08-22 18:53:58',0,'2017-08-22 18:44:27',0),(5,0,1,6,4,0,0,'2017-05-12 00:00:00','',13131.7900,'',27,'ASM(8) d 12999 4000.00,c 12001 4000.00,ASM(17) d 12999 81.79,c 12001 81.79,ASM(19) d 12999 350.00,c 12001 350.00,ASM(9) d 12999 4000.00,c 12001 4000.00,ASM(20) d 12999 350.00,c 12001 350.00,ASM(10) d 12999 4000.00,c 12001 4000.00,ASM(21) d 12999 350.00,c 12001 350.00',2,'3 month rent and utilities','','2017-08-22 18:53:58',0,'2017-08-22 18:45:12',0),(6,0,1,6,4,0,0,'2017-08-15 00:00:00','',13050.0000,'',27,'ASM(11) d 12999 4000.00,c 12001 4000.00,ASM(22) d 12999 350.00,c 12001 350.00,ASM(12) d 12999 4000.00,c 12001 4000.00,ASM(23) d 12999 350.00,c 12001 350.00',1,'3 month rent and utilities','','2017-08-22 18:53:58',0,'2017-08-22 18:45:41',0),(7,0,1,1,1,0,0,'2017-09-14 00:00:00','1234',125.0000,'',27,'ASM(45) d 12999 125.00,c 12001 125.00',2,'','','2017-09-15 05:44:13',0,'2017-09-15 05:43:22',0),(8,0,1,1,1,0,0,'2017-09-14 00:00:00','3456',75.0000,'',27,'ASM(45) d 12999 75.00,c 12001 75.00',2,'','','2017-09-15 05:44:13',0,'2017-09-15 05:43:38',0),(9,0,1,14,5,0,0,'2017-09-18 00:00:00','3452345345',3525.0000,'',33,'',0,'','','2017-09-18 19:49:58',0,'2017-09-18 19:49:58',0),(10,0,1,9,1,0,0,'2017-09-25 00:00:00','34346',50.0000,'',27,'ASM(49) d 12999 50.00,c 12001 50.00',2,'','','2017-09-26 05:29:10',0,'2017-09-26 05:28:20',0),(11,0,1,10,1,0,0,'2017-09-25 00:00:00','3432',50.0000,'',27,'ASM(50) d 12999 50.00,c 12001 50.00',2,'','','2017-09-26 05:33:06',0,'2017-09-26 05:30:24',0),(12,0,1,11,1,0,0,'2017-09-25 00:00:00','6585',50.0000,'',27,'ASM(51) d 12999 50.00,c 12001 50.00',2,'','','2017-09-26 05:33:10',0,'2017-09-26 05:31:31',0),(13,0,1,12,1,0,0,'2017-09-25 00:00:00','4443',50.0000,'',27,'ASM(52) d 12999 50.00,c 12001 50.00',2,'','','2017-09-26 05:33:14',0,'2017-09-26 05:32:00',0),(14,0,1,13,1,0,0,'2017-09-25 00:00:00','7245',75.0000,'',27,'ASM(54) d 12999 75.00,c 12001 75.00',2,'','','2017-09-26 05:33:18',0,'2017-09-26 05:32:57',0);
/*!40000 ALTER TABLE `Receipt` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ReceiptAllocation`
--

DROP TABLE IF EXISTS `ReceiptAllocation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ReceiptAllocation` (
  `RCPAID` bigint(20) NOT NULL AUTO_INCREMENT,
  `RCPTID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RAID` bigint(20) NOT NULL DEFAULT '0',
  `Dt` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `Amount` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `ASMID` bigint(20) NOT NULL DEFAULT '0',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `AcctRule` varchar(150) DEFAULT NULL,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RCPAID`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ReceiptAllocation`
--

LOCK TABLES `ReceiptAllocation` WRITE;
/*!40000 ALTER TABLE `ReceiptAllocation` DISABLE KEYS */;
INSERT INTO `ReceiptAllocation` VALUES (1,1,1,0,'2016-10-03 00:00:00',4000.0000,0,0,'d 10999 _, c 12999 _','2017-08-22 18:29:52',0,'2017-08-22 18:29:52',0),(2,2,1,0,'2016-11-11 00:00:00',12000.0000,0,0,'d 10999 _, c 12999 _','2017-08-22 18:31:17',0,'2017-08-22 18:31:17',0),(3,1,1,3,'2016-10-03 00:00:00',4000.0000,2,0,'ASM(2) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:40:14',0,'2017-08-22 18:40:14',0),(4,2,1,3,'2016-11-11 00:00:00',4000.0000,3,0,'ASM(3) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:40:14',0,'2017-08-22 18:40:14',0),(5,2,1,3,'2016-12-03 00:00:00',4000.0000,4,0,'ASM(4) d 12999 4000.00,c 12001 4000.00','2017-09-21 21:33:03',0,'2017-08-22 18:40:14',0),(6,2,1,3,'2017-01-03 00:00:00',4000.0000,5,0,'ASM(5) d 12999 4000.00,c 12001 4000.00','2017-09-21 21:33:33',0,'2017-08-22 18:40:14',0),(7,3,1,0,'2017-02-03 00:00:00',628.4500,0,0,'d 10999 _, c 12999 _','2017-08-22 18:43:32',0,'2017-08-22 18:43:32',0),(8,4,1,0,'2017-02-13 00:00:00',8350.0000,0,0,'d 10999 _, c 12999 _','2017-08-22 18:44:27',0,'2017-08-22 18:44:27',0),(9,5,1,0,'2017-05-12 00:00:00',13131.7900,0,0,'d 10999 _, c 12999 _','2017-08-22 18:45:12',0,'2017-08-22 18:45:12',0),(10,6,1,0,'2017-08-15 00:00:00',13050.0000,0,0,'d 10999 _, c 12999 _','2017-08-22 18:45:41',0,'2017-08-22 18:45:41',0),(11,3,1,3,'2017-02-03 00:00:00',628.4500,13,0,'ASM(13) d 12999 628.45,c 12001 628.45','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(12,4,1,3,'2017-02-13 00:00:00',4000.0000,6,0,'ASM(6) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(13,4,1,3,'2017-02-13 00:00:00',175.0000,14,0,'ASM(14) d 12999 175.00,c 12001 175.00','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(14,4,1,3,'2017-02-13 00:00:00',4000.0000,7,0,'ASM(7) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(15,4,1,3,'2017-02-13 00:00:00',175.0000,16,0,'ASM(16) d 12999 175.00,c 12001 175.00','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(16,5,1,3,'2017-05-12 00:00:00',4000.0000,8,0,'ASM(8) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(17,5,1,3,'2017-05-12 00:00:00',81.7900,17,0,'ASM(17) d 12999 81.79,c 12001 81.79','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(18,5,1,3,'2017-05-12 00:00:00',350.0000,19,0,'ASM(19) d 12999 350.00,c 12001 350.00','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(19,5,1,3,'2017-05-12 00:00:00',4000.0000,9,0,'ASM(9) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(20,5,1,3,'2017-05-12 00:00:00',350.0000,20,0,'ASM(20) d 12999 350.00,c 12001 350.00','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(21,5,1,3,'2017-05-12 00:00:00',4000.0000,10,0,'ASM(10) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(22,5,1,3,'2017-05-12 00:00:00',350.0000,21,0,'ASM(21) d 12999 350.00,c 12001 350.00','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(23,6,1,3,'2017-08-15 00:00:00',4000.0000,11,0,'ASM(11) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(24,6,1,3,'2017-08-15 00:00:00',350.0000,22,0,'ASM(22) d 12999 350.00,c 12001 350.00','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(25,6,1,3,'2017-08-15 00:00:00',4000.0000,12,0,'ASM(12) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(26,6,1,3,'2017-08-15 00:00:00',350.0000,23,0,'ASM(23) d 12999 350.00,c 12001 350.00','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(27,7,1,0,'2017-09-14 00:00:00',125.0000,0,0,'d 10999 _, c 12999 _','2017-09-15 05:43:22',0,'2017-09-15 05:43:22',0),(28,8,1,0,'2017-09-14 00:00:00',75.0000,0,0,'d 10999 _, c 12999 _','2017-09-15 05:43:38',0,'2017-09-15 05:43:38',0),(29,7,1,1,'2017-09-14 00:00:00',125.0000,45,0,'ASM(45) d 12999 125.00,c 12001 125.00','2017-09-15 05:44:13',0,'2017-09-15 05:44:13',0),(30,8,1,1,'2017-09-14 00:00:00',75.0000,45,0,'ASM(45) d 12999 75.00,c 12001 75.00','2017-09-15 05:44:13',0,'2017-09-15 05:44:13',0),(31,9,1,0,'2017-09-18 00:00:00',3525.0000,0,0,'d 11000 _, c 42300 _','2017-09-18 19:49:58',0,'2017-09-18 19:49:58',0),(32,10,1,0,'2017-09-25 00:00:00',50.0000,0,0,'d 10999 _, c 12999 _','2017-09-26 05:28:20',0,'2017-09-26 05:28:20',0),(33,10,1,8,'2017-09-25 00:00:00',50.0000,49,0,'ASM(49) d 12999 50.00,c 12001 50.00','2017-09-26 05:29:09',0,'2017-09-26 05:29:09',0),(34,11,1,0,'2017-09-25 00:00:00',50.0000,0,0,'d 10999 _, c 12999 _','2017-09-26 05:30:24',0,'2017-09-26 05:30:24',0),(35,12,1,0,'2017-09-25 00:00:00',50.0000,0,0,'d 10999 _, c 12999 _','2017-09-26 05:31:32',0,'2017-09-26 05:31:32',0),(36,13,1,0,'2017-09-25 00:00:00',50.0000,0,0,'d 10999 _, c 12999 _','2017-09-26 05:32:00',0,'2017-09-26 05:32:00',0),(37,14,1,0,'2017-09-25 00:00:00',75.0000,0,0,'d 10999 _, c 12999 _','2017-09-26 05:32:57',0,'2017-09-26 05:32:57',0),(38,11,1,9,'2017-09-25 00:00:00',50.0000,50,0,'ASM(50) d 12999 50.00,c 12001 50.00','2017-09-26 05:33:06',0,'2017-09-26 05:33:06',0),(39,12,1,10,'2017-09-25 00:00:00',50.0000,51,0,'ASM(51) d 12999 50.00,c 12001 50.00','2017-09-26 05:33:10',0,'2017-09-26 05:33:10',0),(40,13,1,11,'2017-09-25 00:00:00',50.0000,52,0,'ASM(52) d 12999 50.00,c 12001 50.00','2017-09-26 05:33:14',0,'2017-09-26 05:33:14',0),(41,14,1,12,'2017-09-25 00:00:00',75.0000,54,0,'ASM(54) d 12999 75.00,c 12001 75.00','2017-09-26 05:33:18',0,'2017-09-26 05:33:18',0);
/*!40000 ALTER TABLE `ReceiptAllocation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Rentable`
--

DROP TABLE IF EXISTS `Rentable`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Rentable` (
  `RID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RentableName` varchar(100) NOT NULL DEFAULT '',
  `AssignmentTime` smallint(6) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RID`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Rentable`
--

LOCK TABLES `Rentable` WRITE;
/*!40000 ALTER TABLE `Rentable` DISABLE KEYS */;
INSERT INTO `Rentable` VALUES (1,1,'309 S. Rexford Drive',1,'2017-08-16 23:13:05',0,'2017-08-16 05:49:58',0),(2,1,'309-1/2 S. Rexford Drive',1,'2017-08-16 23:13:36',0,'2017-08-16 05:49:58',0),(3,1,'311 S. Rexford Drive',1,'2017-08-16 23:13:50',0,'2017-08-16 05:49:58',0),(4,1,'311-1/2 S. Rexford Drive',1,'2017-08-16 23:15:05',0,'2017-08-16 05:49:58',0),(5,1,'315 S. Rexford Drive',2,'2017-09-15 03:51:40',0,'2017-09-15 03:51:40',0),(6,1,'317 S. Rexford Drive',2,'2017-09-15 03:52:13',0,'2017-09-15 03:52:13',0),(7,1,'CP 309',1,'2017-09-15 05:27:04',0,'2017-09-15 05:27:04',0),(8,1,'Washing Machines',2,'2017-09-18 19:45:49',0,'2017-09-18 19:45:49',0);
/*!40000 ALTER TABLE `Rentable` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentableMarketRate`
--

DROP TABLE IF EXISTS `RentableMarketRate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentableMarketRate` (
  `RMRID` bigint(20) NOT NULL AUTO_INCREMENT,
  `RTID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `MarketRate` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `DtStart` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtStop` datetime NOT NULL DEFAULT '9999-12-31 23:59:59',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RMRID`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableMarketRate`
--

LOCK TABLES `RentableMarketRate` WRITE;
/*!40000 ALTER TABLE `RentableMarketRate` DISABLE KEYS */;
INSERT INTO `RentableMarketRate` VALUES (1,1,1,3750.0000,'2016-03-01 00:00:00','9999-01-01 00:00:00','2017-08-16 05:49:58',0),(2,2,1,4000.0000,'2016-10-01 00:00:00','9999-01-01 00:00:00','2017-08-16 05:49:58',0),(3,3,1,4150.0000,'2016-07-01 00:00:00','9999-01-01 00:00:00','2017-08-16 05:49:58',0),(4,4,1,2500.0000,'2016-01-01 00:00:00','9999-01-01 00:00:00','2017-08-16 05:49:58',0),(5,5,1,5000.0000,'2017-01-01 00:00:00','9999-12-31 00:00:00','2017-09-14 21:10:59',0),(6,6,1,4500.0000,'2017-01-01 00:00:00','9999-12-31 00:00:00','2017-09-14 21:13:45',0),(7,0,1,35.0000,'2017-01-01 00:00:00','9999-12-31 00:00:00','2017-09-15 05:19:58',0),(8,8,1,20.0000,'2017-01-01 00:00:00','9999-12-31 00:00:00','2017-09-15 05:20:54',0),(9,9,1,35.0000,'2017-01-01 00:00:00','9999-12-31 00:00:00','2017-09-15 05:25:13',0);
/*!40000 ALTER TABLE `RentableMarketRate` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentableSpecialty`
--

DROP TABLE IF EXISTS `RentableSpecialty`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentableSpecialty` (
  `RSPID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL,
  `Name` varchar(100) NOT NULL DEFAULT '',
  `Fee` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `Description` varchar(256) NOT NULL DEFAULT '',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RSPID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableSpecialty`
--

LOCK TABLES `RentableSpecialty` WRITE;
/*!40000 ALTER TABLE `RentableSpecialty` DISABLE KEYS */;
/*!40000 ALTER TABLE `RentableSpecialty` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentableSpecialtyRef`
--

DROP TABLE IF EXISTS `RentableSpecialtyRef`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentableSpecialtyRef` (
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `RSPID` bigint(20) NOT NULL DEFAULT '0',
  `DtStart` date NOT NULL DEFAULT '1970-01-01',
  `DtStop` date NOT NULL DEFAULT '1970-01-01',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableSpecialtyRef`
--

LOCK TABLES `RentableSpecialtyRef` WRITE;
/*!40000 ALTER TABLE `RentableSpecialtyRef` DISABLE KEYS */;
/*!40000 ALTER TABLE `RentableSpecialtyRef` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentableStatus`
--

DROP TABLE IF EXISTS `RentableStatus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentableStatus` (
  `RSID` bigint(20) NOT NULL AUTO_INCREMENT,
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Status` smallint(6) NOT NULL DEFAULT '0',
  `DtStart` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtStop` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtNoticeToVacate` date NOT NULL DEFAULT '1970-01-01',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RSID`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableStatus`
--

LOCK TABLES `RentableStatus` WRITE;
/*!40000 ALTER TABLE `RentableStatus` DISABLE KEYS */;
INSERT INTO `RentableStatus` VALUES (1,1,1,1,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(2,2,1,1,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(3,3,1,1,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(4,4,1,4,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(5,5,1,1,'2017-09-15 03:51:41','9999-01-01 00:00:00','0000-00-00','2017-09-15 03:51:40',0,'2017-09-15 03:51:40',0),(6,6,1,1,'2017-09-15 03:52:14','9999-01-01 00:00:00','0000-00-00','2017-09-15 03:52:13',0,'2017-09-15 03:52:13',0),(7,7,1,1,'2017-09-15 05:27:04','9999-01-01 00:00:00','0000-00-00','2017-09-15 05:27:04',0,'2017-09-15 05:27:04',0),(8,8,1,1,'2017-09-18 19:45:49','9999-01-01 00:00:00','0000-00-00','2017-09-18 19:45:49',0,'2017-09-18 19:45:49',0);
/*!40000 ALTER TABLE `RentableStatus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentableTypeRef`
--

DROP TABLE IF EXISTS `RentableTypeRef`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentableTypeRef` (
  `RTRID` bigint(20) NOT NULL AUTO_INCREMENT,
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RTID` bigint(20) NOT NULL DEFAULT '0',
  `OverrideRentCycle` bigint(20) NOT NULL DEFAULT '0',
  `OverrideProrationCycle` bigint(20) NOT NULL DEFAULT '0',
  `DtStart` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtStop` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RTRID`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableTypeRef`
--

LOCK TABLES `RentableTypeRef` WRITE;
/*!40000 ALTER TABLE `RentableTypeRef` DISABLE KEYS */;
INSERT INTO `RentableTypeRef` VALUES (1,1,1,1,0,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(2,2,1,2,0,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(3,3,1,3,0,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(4,4,1,4,0,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0),(5,5,1,5,0,0,'2017-09-15 03:51:41','9999-01-01 00:00:00','2017-09-15 03:51:40',0,'2017-09-15 03:51:40',0),(6,6,1,6,0,0,'2017-09-15 03:52:14','9999-01-01 00:00:00','2017-09-15 03:52:13',0,'2017-09-15 03:52:13',0),(7,7,1,1,0,0,'2017-09-15 05:27:04','2018-03-01 00:00:00','2017-09-15 05:27:04',0,'2017-09-15 05:27:04',0),(8,8,1,10,0,0,'2017-09-18 19:45:49','9999-01-01 00:00:00','2017-09-18 19:45:49',0,'2017-09-18 19:45:49',0);
/*!40000 ALTER TABLE `RentableTypeRef` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentableTypeTax`
--

DROP TABLE IF EXISTS `RentableTypeTax`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentableTypeTax` (
  `RTID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `TAXID` bigint(20) NOT NULL DEFAULT '0',
  `DtStart` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtStop` datetime NOT NULL DEFAULT '9999-12-31 23:59:59',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableTypeTax`
--

LOCK TABLES `RentableTypeTax` WRITE;
/*!40000 ALTER TABLE `RentableTypeTax` DISABLE KEYS */;
/*!40000 ALTER TABLE `RentableTypeTax` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentableTypes`
--

DROP TABLE IF EXISTS `RentableTypes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentableTypes` (
  `RTID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Style` char(255) NOT NULL DEFAULT '',
  `Name` varchar(256) NOT NULL DEFAULT '',
  `RentCycle` bigint(20) NOT NULL DEFAULT '0',
  `Proration` bigint(20) NOT NULL DEFAULT '0',
  `GSRPC` bigint(20) NOT NULL DEFAULT '0',
  `ManageToBudget` smallint(6) NOT NULL DEFAULT '0',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RTID`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableTypes`
--

LOCK TABLES `RentableTypes` WRITE;
/*!40000 ALTER TABLE `RentableTypes` DISABLE KEYS */;
INSERT INTO `RentableTypes` VALUES (1,1,'309','309 S. Rexford Drive',6,4,4,1,0,'2017-08-16 23:09:23',0,'2017-08-16 05:49:58',0),(2,1,'309-1/2','309-1/2 S. Rexford Drive',6,4,4,1,0,'2017-08-16 23:10:20',0,'2017-08-16 05:49:58',0),(3,1,'311','311  S. Rexford Drive',6,4,4,1,0,'2017-08-16 23:10:42',0,'2017-08-16 05:49:58',0),(4,1,'311-1/2','311-1/2  S. Rexford Drive',6,4,4,1,0,'2017-08-16 23:11:02',0,'2017-08-16 05:49:58',0),(5,1,'315','315 S. Rexford',6,4,4,1,0,'2017-09-14 21:10:59',0,'2017-09-14 21:10:59',0),(6,1,'317','317 S. Rexford',6,4,4,1,0,'2017-09-14 21:13:45',0,'2017-09-14 21:13:45',0),(7,1,'Carport Covered','Covered Car port',6,4,4,1,1,'2017-09-15 05:23:35',0,'2017-09-15 05:19:58',0),(8,1,'Carport NoCover','Carport',6,4,4,1,0,'2017-09-15 05:20:54',0,'2017-09-15 05:20:54',0),(9,1,'Covered Carport','Open car port',6,4,4,1,0,'2017-09-15 05:25:13',0,'2017-09-15 05:25:13',0),(10,1,'Vending','Washing Machines',0,0,0,0,0,'2017-09-18 19:44:01',0,'2017-09-18 19:44:01',0);
/*!40000 ALTER TABLE `RentableTypes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentableUsers`
--

DROP TABLE IF EXISTS `RentableUsers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentableUsers` (
  `RUID` bigint(20) NOT NULL AUTO_INCREMENT,
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `TCID` bigint(20) NOT NULL DEFAULT '0',
  `DtStart` date NOT NULL DEFAULT '1970-01-01',
  `DtStop` date NOT NULL DEFAULT '1970-01-01',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RUID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableUsers`
--

LOCK TABLES `RentableUsers` WRITE;
/*!40000 ALTER TABLE `RentableUsers` DISABLE KEYS */;
INSERT INTO `RentableUsers` VALUES (1,2,1,5,'2016-10-01','2017-12-31','2017-08-22 18:18:55',0),(2,1,1,1,'2014-03-01','2018-03-01','2017-09-14 21:20:15',0),(3,1,1,2,'2014-03-01','2018-03-01','2017-09-14 21:20:44',0),(4,3,1,3,'2016-07-01','2018-07-01','2017-09-14 21:25:59',0),(5,3,1,4,'2016-07-01','2018-07-01','2017-09-14 21:26:54',0),(6,5,1,7,'2017-09-14','2018-09-14','2017-09-15 03:58:33',0);
/*!40000 ALTER TABLE `RentableUsers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentalAgreement`
--

DROP TABLE IF EXISTS `RentalAgreement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentalAgreement` (
  `RAID` bigint(20) NOT NULL AUTO_INCREMENT,
  `RATID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `NLID` bigint(20) NOT NULL DEFAULT '0',
  `AgreementStart` date NOT NULL DEFAULT '1970-01-01',
  `AgreementStop` date NOT NULL DEFAULT '1970-01-01',
  `PossessionStart` date NOT NULL DEFAULT '1970-01-01',
  `PossessionStop` date NOT NULL DEFAULT '1970-01-01',
  `RentStart` date NOT NULL DEFAULT '1970-01-01',
  `RentStop` date NOT NULL DEFAULT '1970-01-01',
  `RentCycleEpoch` date NOT NULL DEFAULT '1970-01-01',
  `UnspecifiedAdults` smallint(6) NOT NULL DEFAULT '0',
  `UnspecifiedChildren` smallint(6) NOT NULL DEFAULT '0',
  `Renewal` smallint(6) NOT NULL DEFAULT '0',
  `SpecialProvisions` varchar(1024) NOT NULL DEFAULT '',
  `LeaseType` bigint(20) NOT NULL DEFAULT '0',
  `ExpenseAdjustmentType` bigint(20) NOT NULL DEFAULT '0',
  `ExpensesStop` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `ExpenseStopCalculation` varchar(128) NOT NULL DEFAULT '',
  `BaseYearEnd` date NOT NULL DEFAULT '1970-01-01',
  `ExpenseAdjustment` date NOT NULL DEFAULT '1970-01-01',
  `EstimatedCharges` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `RateChange` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `CSAgent` bigint(20) NOT NULL DEFAULT '0',
  `NextRateChange` date NOT NULL DEFAULT '1970-01-01',
  `PermittedUses` varchar(128) NOT NULL DEFAULT '',
  `ExclusiveUses` varchar(128) NOT NULL DEFAULT '',
  `ExtensionOption` varchar(128) NOT NULL DEFAULT '',
  `ExtensionOptionNotice` date NOT NULL DEFAULT '1970-01-01',
  `ExpansionOption` varchar(128) NOT NULL DEFAULT '',
  `ExpansionOptionNotice` date NOT NULL DEFAULT '1970-01-01',
  `RightOfFirstRefusal` varchar(128) NOT NULL DEFAULT '',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RAID`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreement`
--

LOCK TABLES `RentalAgreement` WRITE;
/*!40000 ALTER TABLE `RentalAgreement` DISABLE KEYS */;
INSERT INTO `RentalAgreement` VALUES (1,0,1,0,'2014-03-01','2018-03-01','2014-03-01','2018-03-01','2014-03-01','2018-03-01','2014-03-01',0,1,2,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2017-08-22 18:13:03',0,'2017-08-22 18:06:59',0),(2,0,1,0,'2016-07-01','2018-07-01','2016-07-01','2018-07-01','2016-07-01','2018-07-01','2016-07-01',0,1,0,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2017-08-22 18:16:04',0,'2017-08-22 18:13:14',0),(3,0,1,0,'2016-10-01','2017-12-31','2016-10-01','2017-12-31','2016-10-01','2017-12-31','2016-10-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2017-08-22 18:19:04',0,'2017-08-22 18:16:13',0),(4,0,1,0,'2017-09-14','2018-09-14','2017-09-14','2018-09-14','2017-09-14','2018-09-14','2017-09-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2017-09-15 03:57:13',0,'2017-09-15 03:57:13',0),(5,0,1,0,'2017-10-01','2018-10-01','2017-10-01','2018-10-01','2017-10-01','2018-10-01','2017-10-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2017-09-15 04:02:09',0,'2017-09-15 03:59:26',0),(8,0,1,0,'1970-01-01','1970-01-01','1970-01-01','1970-01-01','1970-01-01','1970-01-01','1970-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',1,'2017-09-25 20:42:42',0,'2017-09-18 03:28:34',0),(9,0,1,0,'1970-01-01','1970-01-01','1970-01-01','1970-01-01','1970-01-01','1970-01-01','1970-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',1,'2017-09-25 20:46:36',0,'2017-09-18 03:31:00',0),(10,0,1,0,'1970-01-01','1970-01-01','1970-01-01','1970-01-01','1970-01-01','1970-01-01','1970-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2017-09-18 03:38:01',0,'2017-09-18 03:37:25',0),(11,0,1,0,'1970-01-01','1970-01-01','1970-01-01','1970-01-01','1970-01-01','1970-01-01','1970-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',1,'2017-09-25 20:47:21',0,'2017-09-18 03:39:12',0),(12,0,1,0,'1970-01-01','1970-01-01','1970-01-01','1970-01-01','1970-01-01','1970-01-01','1970-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',1,'2017-09-25 20:47:35',0,'2017-09-18 03:40:34',0),(14,0,1,0,'2017-08-01','2018-09-25','2017-08-01','2018-09-25','2017-08-01','2018-09-25','2017-08-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2017-09-25 21:23:51',0,'2017-09-25 21:07:18',0);
/*!40000 ALTER TABLE `RentalAgreement` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentalAgreementPayors`
--

DROP TABLE IF EXISTS `RentalAgreementPayors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentalAgreementPayors` (
  `RAPID` bigint(20) NOT NULL AUTO_INCREMENT,
  `RAID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `TCID` bigint(20) NOT NULL DEFAULT '0',
  `DtStart` date NOT NULL DEFAULT '1970-01-01',
  `DtStop` date NOT NULL DEFAULT '1970-01-01',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RAPID`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementPayors`
--

LOCK TABLES `RentalAgreementPayors` WRITE;
/*!40000 ALTER TABLE `RentalAgreementPayors` DISABLE KEYS */;
INSERT INTO `RentalAgreementPayors` VALUES (1,1,1,1,'2014-03-01','2018-03-01',0,'2017-08-22 18:12:25',0),(2,1,1,2,'2014-03-01','2018-03-01',0,'2017-08-22 18:12:47',0),(3,2,1,3,'2016-07-01','2018-07-01',0,'2017-08-22 18:15:36',0),(4,2,1,4,'2016-07-01','2018-07-01',0,'2017-08-22 18:15:55',0),(5,3,1,6,'2016-10-01','2017-12-31',0,'2017-08-22 18:18:21',0),(6,4,1,7,'2017-09-14','2018-09-14',0,'2017-09-15 03:58:14',0),(7,5,1,8,'2017-10-01','2018-10-01',0,'2017-09-15 04:01:39',0),(8,8,1,9,'2017-09-17','2018-09-17',0,'2017-09-18 03:29:25',0),(9,9,1,10,'2017-09-17','2018-09-17',0,'2017-09-18 03:36:26',0),(10,10,1,11,'2017-09-17','2018-09-17',0,'2017-09-18 03:37:54',0),(11,11,1,12,'2017-09-17','2018-09-17',0,'2017-09-18 03:39:52',0),(12,12,1,13,'2017-09-17','2018-09-17',0,'2017-09-18 03:40:55',0),(13,14,1,15,'2017-09-01','2018-09-25',0,'2017-09-25 21:07:41',0);
/*!40000 ALTER TABLE `RentalAgreementPayors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentalAgreementPets`
--

DROP TABLE IF EXISTS `RentalAgreementPets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentalAgreementPets` (
  `PETID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RAID` bigint(20) NOT NULL DEFAULT '0',
  `Type` varchar(100) NOT NULL DEFAULT '',
  `Breed` varchar(100) NOT NULL DEFAULT '',
  `Color` varchar(100) NOT NULL DEFAULT '',
  `Weight` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `Name` varchar(100) NOT NULL DEFAULT '',
  `DtStart` date NOT NULL DEFAULT '1970-01-01',
  `DtStop` date NOT NULL DEFAULT '1970-01-01',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`PETID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementPets`
--

LOCK TABLES `RentalAgreementPets` WRITE;
/*!40000 ALTER TABLE `RentalAgreementPets` DISABLE KEYS */;
/*!40000 ALTER TABLE `RentalAgreementPets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentalAgreementRentables`
--

DROP TABLE IF EXISTS `RentalAgreementRentables`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentalAgreementRentables` (
  `RARID` bigint(20) NOT NULL AUTO_INCREMENT,
  `RAID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `CLID` bigint(20) NOT NULL DEFAULT '0',
  `ContractRent` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `RARDtStart` date NOT NULL DEFAULT '1970-01-01',
  `RARDtStop` date NOT NULL DEFAULT '1970-01-01',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RARID`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementRentables`
--

LOCK TABLES `RentalAgreementRentables` WRITE;
/*!40000 ALTER TABLE `RentalAgreementRentables` DISABLE KEYS */;
INSERT INTO `RentalAgreementRentables` VALUES (1,1,1,1,0,3750.0000,'2014-03-01','2018-03-01','2017-08-22 18:12:00',0),(2,2,1,3,0,4150.0000,'2016-07-01','2018-07-01','2017-08-22 18:14:23',0),(3,3,1,2,0,4000.0000,'2016-10-01','2017-12-31','2017-08-22 18:18:02',0),(4,4,1,5,0,5000.0000,'2017-09-14','2018-09-14','2017-09-15 03:57:58',0),(5,5,1,6,0,4500.0000,'2017-10-01','2018-10-01','2017-09-15 04:01:22',0);
/*!40000 ALTER TABLE `RentalAgreementRentables` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentalAgreementTax`
--

DROP TABLE IF EXISTS `RentalAgreementTax`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentalAgreementTax` (
  `RAID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `DtStart` date NOT NULL DEFAULT '1970-01-01',
  `DtStop` date NOT NULL DEFAULT '1970-01-01',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementTax`
--

LOCK TABLES `RentalAgreementTax` WRITE;
/*!40000 ALTER TABLE `RentalAgreementTax` DISABLE KEYS */;
/*!40000 ALTER TABLE `RentalAgreementTax` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentalAgreementTemplate`
--

DROP TABLE IF EXISTS `RentalAgreementTemplate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentalAgreementTemplate` (
  `RATID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RATemplateName` varchar(100) DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RATID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementTemplate`
--

LOCK TABLES `RentalAgreementTemplate` WRITE;
/*!40000 ALTER TABLE `RentalAgreementTemplate` DISABLE KEYS */;
/*!40000 ALTER TABLE `RentalAgreementTemplate` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `SLString`
--

DROP TABLE IF EXISTS `SLString`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `SLString` (
  `SLSID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `SLID` bigint(20) NOT NULL DEFAULT '0',
  `Value` varchar(256) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`SLSID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SLString`
--

LOCK TABLES `SLString` WRITE;
/*!40000 ALTER TABLE `SLString` DISABLE KEYS */;
/*!40000 ALTER TABLE `SLString` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `StringList`
--

DROP TABLE IF EXISTS `StringList`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `StringList` (
  `SLID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Name` varchar(50) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`SLID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `StringList`
--

LOCK TABLES `StringList` WRITE;
/*!40000 ALTER TABLE `StringList` DISABLE KEYS */;
/*!40000 ALTER TABLE `StringList` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `TWS`
--

DROP TABLE IF EXISTS `TWS`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `TWS` (
  `TWSID` bigint(20) NOT NULL AUTO_INCREMENT,
  `Owner` varchar(256) NOT NULL DEFAULT '',
  `OwnerData` varchar(256) NOT NULL DEFAULT '',
  `WorkerName` varchar(256) NOT NULL DEFAULT '',
  `ActivateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `Node` varchar(256) NOT NULL DEFAULT '',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `DtActivated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `DtCompleted` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `DtCreate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `DtLastUpdate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`TWSID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TWS`
--

LOCK TABLES `TWS` WRITE;
/*!40000 ALTER TABLE `TWS` DISABLE KEYS */;
INSERT INTO `TWS` VALUES (1,'CreateAssessmentInstances','','CreateAssessmentInstances','2017-09-09 00:00:00','ip-172-31-51-141.ec2.internal',4,'2017-09-08 00:00:06','2017-09-08 00:00:06','2017-08-16 06:07:23','2017-09-08 00:00:05'),(2,'CreateAssessmentInstances','','CreateAssessmentInstances','2017-09-28 00:00:00','Steves-MacBook-Pro-2.local',4,'2017-09-27 04:17:37','2017-09-27 04:17:37','2017-09-14 14:12:38','2017-09-26 21:17:37');
/*!40000 ALTER TABLE `TWS` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Tax`
--

DROP TABLE IF EXISTS `Tax`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Tax` (
  `TAXID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Name` varchar(50) DEFAULT NULL,
  `TaxingAuthority` varchar(100) DEFAULT NULL,
  `TaxingAuthorityAddress` varchar(256) DEFAULT NULL,
  `FilingDate` date NOT NULL DEFAULT '1970-01-01',
  `FilingCycle` bigint(20) NOT NULL DEFAULT '0',
  `Instructions` varchar(1024) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`TAXID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Tax`
--

LOCK TABLES `Tax` WRITE;
/*!40000 ALTER TABLE `Tax` DISABLE KEYS */;
/*!40000 ALTER TABLE `Tax` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `TaxRate`
--

DROP TABLE IF EXISTS `TaxRate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `TaxRate` (
  `TAXID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `DtStart` date NOT NULL DEFAULT '1970-01-01',
  `DtStop` date NOT NULL DEFAULT '1970-01-01',
  `Rate` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `Fee` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `Formula` varchar(256) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TaxRate`
--

LOCK TABLES `TaxRate` WRITE;
/*!40000 ALTER TABLE `TaxRate` DISABLE KEYS */;
/*!40000 ALTER TABLE `TaxRate` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Transactant`
--

DROP TABLE IF EXISTS `Transactant`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Transactant` (
  `TCID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `NLID` bigint(20) NOT NULL DEFAULT '0',
  `FirstName` varchar(100) NOT NULL DEFAULT '',
  `MiddleName` varchar(100) NOT NULL DEFAULT '',
  `LastName` varchar(100) NOT NULL DEFAULT '',
  `PreferredName` varchar(100) NOT NULL DEFAULT '',
  `CompanyName` varchar(100) NOT NULL DEFAULT '',
  `IsCompany` smallint(6) NOT NULL DEFAULT '0',
  `PrimaryEmail` varchar(100) NOT NULL DEFAULT '',
  `SecondaryEmail` varchar(100) NOT NULL DEFAULT '',
  `WorkPhone` varchar(100) NOT NULL DEFAULT '',
  `CellPhone` varchar(100) NOT NULL DEFAULT '',
  `Address` varchar(100) NOT NULL DEFAULT '',
  `Address2` varchar(100) NOT NULL DEFAULT '',
  `City` varchar(100) NOT NULL DEFAULT '',
  `State` char(25) NOT NULL DEFAULT '',
  `PostalCode` varchar(100) NOT NULL DEFAULT '',
  `Country` varchar(100) NOT NULL DEFAULT '',
  `Website` varchar(100) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`TCID`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Transactant`
--

LOCK TABLES `Transactant` WRITE;
/*!40000 ALTER TABLE `Transactant` DISABLE KEYS */;
INSERT INTO `Transactant` VALUES (1,1,0,'Aaron','','Read','','',0,'read.aaron@gmail.com','','','469-307-7095','309 S Rexford Drive','','Beverly Hills','CA','90212','','','2017-08-16 14:38:42',0,'2017-08-16 14:38:42',0),(2,1,0,'Kirsten','','Read','','',0,'klmrda@gmail.com','','','469-693-9933','309 S Rexford Drive','','Beverly Hills','CA','90212','','','2017-08-16 14:40:13',0,'2017-08-16 14:40:13',0),(3,1,0,'Kevin','','Mills','','',0,'kevinmillsesq@aol.com','','','424-234-3535','311 S Rexford Drive','','Beverly Hills','CA','90212','','','2017-08-16 14:41:42',0,'2017-08-16 14:41:42',0),(4,1,0,'Lauren','','Beck','','',0,'laurensbeck@aol.com','','','310-948-6442','311 S Rexford Drive','','Beverly Hills','CA','90212','','','2017-08-16 14:43:10',0,'2017-08-16 14:43:10',0),(5,1,0,'Alex','','Vahabzadeh','','Beaumont Partners, Ltd.',0,'av@beaumont-partners.ch','','','202-550-2477','309 1/2 S Rexford Drive','','Beverly Hills','CA','90212','','','2017-08-16 14:47:19',0,'2017-08-16 14:47:19',0),(6,1,0,'','','','','Beaumont Partners, Ltd.',1,'','','44 79 203 354 77','','118 Rue du Rhone','','1204 Geneva','','','Switzerland','www.beaumont-partners.ch','2017-08-16 14:50:01',0,'2017-08-16 14:49:36',0),(7,1,0,'Edna','','Jones','Edna','',0,'edna@rexford.com','','','','315 S. Rexford Drive','','Beverly Hills','CA','90210','USA','','2017-09-15 03:54:33',0,'2017-09-15 03:54:33',0),(8,1,0,'Vera','','Jones','Vera','',0,'vera@rexford.com','','','','317 S. Rexford Drive','','Beverly Hills','CA','90210','USA','','2017-09-15 03:55:54',0,'2017-09-15 03:55:54',0),(9,1,0,'Ginger','','Smith','','',0,'','','','','ginger@gilligansisle.com','','','','','','','2017-09-18 02:53:22',0,'2017-09-18 02:53:22',0),(10,1,0,'Herman','','Mann','','',0,'hmann@gmail.com','','','','','','','','','','','2017-09-18 02:53:56',0,'2017-09-18 02:53:56',0),(11,1,0,'Contessa','','Livingston','Tessa','',0,'tessa@gmail.com','','','','','','','','','','','2017-09-18 02:55:19',0,'2017-09-18 02:55:19',0),(12,1,0,'Jillian','','Timmons','Jill','',0,'jill@xyz.com','','','','','','','','','','','2017-09-18 02:56:13',0,'2017-09-18 02:56:13',0),(13,1,0,'Ariel','','Simmons','','',0,'ariel@abc.com','','','','','','','','','','','2017-09-18 02:56:42',0,'2017-09-18 02:56:42',0),(14,1,0,'','','','','Vending Machines',1,'hostcompany@hostcompany.com','','','','','','','','','','','2017-09-18 19:33:38',0,'2017-09-18 19:33:38',0),(15,1,0,'William','','Smith','Bill','',0,'bill@smith.com','','','','','','','CA','90210','USA','','2017-09-25 21:06:26',0,'2017-09-25 21:06:26',0);
/*!40000 ALTER TABLE `Transactant` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `User` (
  `TCID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Points` bigint(20) NOT NULL DEFAULT '0',
  `DateofBirth` date NOT NULL DEFAULT '1970-01-01',
  `EmergencyContactName` varchar(100) NOT NULL DEFAULT '',
  `EmergencyContactAddress` varchar(100) NOT NULL DEFAULT '',
  `EmergencyContactTelephone` varchar(100) NOT NULL DEFAULT '',
  `EmergencyEmail` varchar(100) NOT NULL DEFAULT '',
  `AlternateAddress` varchar(100) NOT NULL DEFAULT '',
  `EligibleFutureUser` smallint(6) NOT NULL DEFAULT '1',
  `Industry` varchar(100) NOT NULL DEFAULT '',
  `SourceSLSID` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`TCID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` VALUES (1,1,0,'1900-01-01','','','','','',1,'',0,'2017-08-16 14:38:42',0,'2017-08-16 14:38:42',0),(2,1,0,'1900-01-01','','','','','',1,'',0,'2017-08-16 14:40:13',0,'2017-08-16 14:40:13',0),(3,1,0,'1900-01-01','','','','','',1,'',0,'2017-08-16 14:41:42',0,'2017-08-16 14:41:42',0),(4,1,0,'1900-01-01','','','','','',1,'',0,'2017-08-16 14:43:10',0,'2017-08-16 14:43:10',0),(5,1,0,'1900-01-01','','','','','',1,'',0,'2017-08-16 14:47:19',0,'2017-08-16 14:47:19',0),(6,1,0,'1900-01-01','','','','','',1,'',0,'2017-08-16 14:49:36',0,'2017-08-16 14:49:36',0),(7,1,0,'1900-01-01','','','','','',1,'',0,'2017-09-15 03:54:33',0,'2017-09-15 03:54:33',0),(8,1,0,'1900-01-01','','','','','',1,'',0,'2017-09-15 03:55:54',0,'2017-09-15 03:55:54',0),(9,1,0,'1900-01-01','','','','','',1,'',0,'2017-09-18 02:53:22',0,'2017-09-18 02:53:22',0),(10,1,0,'1900-01-01','','','','','',1,'',0,'2017-09-18 02:53:56',0,'2017-09-18 02:53:56',0),(11,1,0,'1900-01-01','','','','','',1,'',0,'2017-09-18 02:55:19',0,'2017-09-18 02:55:19',0),(12,1,0,'1900-01-01','','','','','',1,'',0,'2017-09-18 02:56:13',0,'2017-09-18 02:56:13',0),(13,1,0,'1900-01-01','','','','','',1,'',0,'2017-09-18 02:56:42',0,'2017-09-18 02:56:42',0),(14,1,0,'1900-01-01','','','','','',1,'',0,'2017-09-18 19:33:38',0,'2017-09-18 19:33:38',0),(15,1,0,'1900-01-01','','','','','',1,'',0,'2017-09-25 21:06:26',0,'2017-09-25 21:06:26',0);
/*!40000 ALTER TABLE `User` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Vehicle`
--

DROP TABLE IF EXISTS `Vehicle`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Vehicle` (
  `VID` bigint(20) NOT NULL AUTO_INCREMENT,
  `TCID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `VehicleType` varchar(80) NOT NULL DEFAULT '',
  `VehicleMake` varchar(80) NOT NULL DEFAULT '',
  `VehicleModel` varchar(80) NOT NULL DEFAULT '',
  `VehicleColor` varchar(80) NOT NULL DEFAULT '',
  `VehicleYear` bigint(20) NOT NULL DEFAULT '0',
  `LicensePlateState` varchar(80) NOT NULL DEFAULT '',
  `LicensePlateNumber` varchar(80) NOT NULL DEFAULT '',
  `ParkingPermitNumber` varchar(80) NOT NULL DEFAULT '',
  `DtStart` date NOT NULL DEFAULT '1970-01-01',
  `DtStop` date NOT NULL DEFAULT '1970-01-01',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`VID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Vehicle`
--

LOCK TABLES `Vehicle` WRITE;
/*!40000 ALTER TABLE `Vehicle` DISABLE KEYS */;
/*!40000 ALTER TABLE `Vehicle` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-09-26 22:18:54
