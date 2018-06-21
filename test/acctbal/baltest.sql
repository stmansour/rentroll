-- MySQL dump 10.13  Distrib 5.7.22, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: rentroll
-- ------------------------------------------------------
-- Server version	5.7.22

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
  `SubARID` bigint(20) NOT NULL DEFAULT '0',
  `ARType` smallint(6) NOT NULL DEFAULT '0',
  `RARequired` smallint(6) NOT NULL DEFAULT '0',
  `DebitLID` bigint(20) NOT NULL DEFAULT '0',
  `CreditLID` bigint(20) NOT NULL DEFAULT '0',
  `Description` varchar(1024) NOT NULL DEFAULT '',
  `DtStart` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtStop` datetime NOT NULL DEFAULT '9999-12-31 00:00:00',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `DefaultAmount` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`ARID`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `AR`
--

LOCK TABLES `AR` WRITE;
/*!40000 ALTER TABLE `AR` DISABLE KEYS */;
INSERT INTO `AR` VALUES (1,1,'Taxable Rent Assessment',0,0,3,5,14,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 01:41:40',0,'2017-07-05 01:30:25',0),(2,1,'Non-Taxable Rent Assessment',0,0,3,5,15,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 01:41:22',0,'2017-07-05 01:31:15',0),(3,1,'Broadcast and IT Assessment',0,0,3,5,27,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 01:41:06',0,'2017-07-05 01:41:06',0),(4,1,'Electric Base Fee Assessment',0,0,3,5,33,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 01:42:38',0,'2017-07-05 01:42:38',0),(5,1,'Water and Sewer Base Fee',0,0,3,5,35,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 01:54:05',0,'2017-07-05 01:54:05',0),(6,1,'Water and Sewer Overage Assessment',0,0,3,5,36,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 01:54:50',0,'2017-07-05 01:54:50',0),(7,1,'Application Fee',0,0,3,5,43,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 01:56:23',0,'2017-07-05 01:56:23',0),(8,1,'Late Fees',0,0,3,5,44,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 01:57:12',0,'2017-07-05 01:57:12',0),(9,1,'Insufficient Funds Fee',0,0,3,5,45,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 01:58:45',0,'2017-07-05 01:58:45',0),(10,1,'Month to Month Fee',0,0,3,5,46,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 02:00:06',0,'2017-07-05 02:00:06',0),(11,1,'Termination No-Show Fee',0,0,3,5,48,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 02:02:12',0,'2017-07-05 02:02:12',0),(12,1,'Pet Fee',0,0,3,5,49,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 02:03:16',0,'2017-07-05 02:03:16',0),(13,1,'Pet Rent',0,0,3,5,50,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 02:03:48',0,'2017-07-05 02:03:48',0),(14,1,'Expense Reimbursement',0,0,3,5,51,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 02:05:20',0,'2017-07-05 02:05:20',0),(15,1,'Maintenance Fee',0,0,3,5,52,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 02:07:00',0,'2017-07-05 02:06:47',0),(16,1,'Eviction Reimbursement Fee',0,0,3,5,53,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 02:07:58',0,'2017-07-05 02:07:58',0),(17,1,'Bad Debt Write-Off',0,0,3,71,5,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 02:09:46',0,'2017-07-05 02:09:46',0),(18,1,'Forfeited Security Deposit',0,0,3,10,55,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 02:10:55',0,'2017-07-05 02:10:55',0),(19,1,'Deposit to Operating Account',0,1,3,3,72,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 13:49:57',0,'2017-07-05 02:19:47',0),(20,1,'Deposit to Security Deposit Account',0,1,3,4,72,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 13:50:06',0,'2017-07-05 02:20:18',0),(21,1,'Security Deposit Assessment',0,0,3,5,10,'','2000-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-05 02:21:17',0,'2017-07-05 02:21:17',0),(23,1,'Commission Offset to Collection',0,0,3,74,5,'','2017-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-12 16:43:39',0,'2017-07-12 16:43:39',0),(24,1,'Tenant Withholding',0,0,3,75,5,'','2017-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-12 16:44:12',0,'2017-07-12 16:44:12',0),(25,1,'Cash Over/Short',0,0,3,70,5,'enter Short as a positive number','2017-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-12 16:45:13',0,'2017-07-12 16:45:13',0),(26,1,'Bank Services',0,0,3,76,5,'','2017-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-07-12 16:45:50',0,'2017-07-12 16:45:50',0),(27,1,'Bank Services Fee',0,2,3,73,5,'','2010-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-11 02:42:55',0,'2017-08-09 04:11:35',0);
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
  `AGRCPTID` bigint(20) NOT NULL DEFAULT '0',
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
) ENGINE=InnoDB AUTO_INCREMENT=67 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Assessments`
--

LOCK TABLES `Assessments` WRITE;
/*!40000 ALTER TABLE `Assessments` DISABLE KEYS */;
INSERT INTO `Assessments` VALUES (1,0,0,0,1,1,0,1,3750.0000,'2016-03-01 00:00:00','2018-03-01 00:00:00',6,4,0,'',2,0,'lease renewal','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(2,1,0,0,1,1,0,1,3750.0000,'2016-03-01 00:00:00','2016-03-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(3,1,0,0,1,1,0,1,3750.0000,'2016-04-01 00:00:00','2016-04-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(4,1,0,0,1,1,0,1,3750.0000,'2016-05-01 00:00:00','2016-05-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(5,1,0,0,1,1,0,1,3750.0000,'2016-06-01 00:00:00','2016-06-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(6,1,0,0,1,1,0,1,3750.0000,'2016-07-01 00:00:00','2016-07-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(7,1,0,0,1,1,0,1,3750.0000,'2016-08-01 00:00:00','2016-08-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(8,1,0,0,1,1,0,1,3750.0000,'2016-09-01 00:00:00','2016-09-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(9,1,0,0,1,1,0,1,3750.0000,'2016-10-01 00:00:00','2016-10-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(10,1,0,0,1,1,0,1,3750.0000,'2016-11-01 00:00:00','2016-11-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(11,1,0,0,1,1,0,1,3750.0000,'2016-12-01 00:00:00','2016-12-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(12,1,0,0,1,1,0,1,3750.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(13,1,0,0,1,1,0,1,3750.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(14,1,0,0,1,1,0,1,3750.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(15,1,0,0,1,1,0,1,3750.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(16,1,0,0,1,1,0,1,3750.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(17,1,0,0,1,1,0,1,3750.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(18,1,0,0,1,1,0,1,3750.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',2,2,'lease renewal','2017-07-18 21:47:42',0,'2017-07-18 19:10:52',0),(19,0,0,0,1,3,0,4,4150.0000,'2016-07-01 00:00:00','2018-07-01 00:00:00',6,4,0,'',2,0,'initial lease','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(20,19,0,0,1,3,0,4,4150.0000,'2016-07-01 00:00:00','2016-07-02 00:00:00',6,4,0,'',2,2,'initial lease','2017-07-18 21:49:52',0,'2017-07-18 19:12:53',0),(21,19,0,0,1,3,0,4,4150.0000,'2016-08-01 00:00:00','2016-08-02 00:00:00',6,4,0,'',2,2,'initial lease','2017-07-18 21:49:52',0,'2017-07-18 19:12:53',0),(22,19,0,0,1,3,0,4,4150.0000,'2016-09-01 00:00:00','2016-09-02 00:00:00',6,4,0,'',2,2,'initial lease','2017-07-18 21:49:52',0,'2017-07-18 19:12:53',0),(23,19,0,0,1,3,0,4,4150.0000,'2016-10-01 00:00:00','2016-10-02 00:00:00',6,4,0,'',2,2,'initial lease','2017-07-18 21:49:53',0,'2017-07-18 19:12:53',0),(24,19,0,0,1,3,0,4,4150.0000,'2016-11-01 00:00:00','2016-11-02 00:00:00',6,4,0,'',2,2,'initial lease','2017-07-18 21:49:53',0,'2017-07-18 19:12:53',0),(25,19,0,0,1,3,0,4,4150.0000,'2016-12-01 00:00:00','2016-12-02 00:00:00',6,4,0,'',2,2,'initial lease','2017-07-18 21:49:53',0,'2017-07-18 19:12:53',0),(26,19,0,0,1,3,0,4,4150.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',2,2,'initial lease','2017-07-18 21:49:53',0,'2017-07-18 19:12:53',0),(27,19,0,0,1,3,0,4,4150.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',2,2,'initial lease','2017-07-18 21:49:53',0,'2017-07-18 19:12:53',0),(28,19,0,0,1,3,0,4,4150.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',2,2,'initial lease','2017-07-18 21:49:53',0,'2017-07-18 19:12:53',0),(29,19,0,0,1,3,0,4,4150.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',2,2,'initial lease','2017-07-18 21:49:53',0,'2017-07-18 19:12:53',0),(30,19,0,0,1,3,0,4,4150.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',2,2,'initial lease','2017-07-18 21:49:53',0,'2017-07-18 19:12:53',0),(31,19,0,0,1,3,0,4,4150.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',2,2,'initial lease','2017-07-18 21:49:53',0,'2017-07-18 19:12:53',0),(32,19,0,0,1,3,0,4,4150.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',2,0,'initial lease','2017-07-24 08:06:52',0,'2017-07-18 19:12:53',0),(33,0,0,0,1,2,0,5,4000.0000,'2016-10-01 00:00:00','2017-12-31 00:00:00',6,4,0,'',2,0,'initial, month-to-month','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(34,33,0,0,1,2,0,5,4000.0000,'2016-10-01 00:00:00','2016-10-02 00:00:00',6,4,0,'',2,2,'initial, month-to-month','2017-07-18 21:55:07',0,'2017-07-18 19:13:43',0),(35,33,0,0,1,2,0,5,4000.0000,'2016-11-01 00:00:00','2016-11-02 00:00:00',6,4,0,'',2,2,'initial, month-to-month','2017-07-18 21:55:07',0,'2017-07-18 19:13:43',0),(36,33,0,0,1,2,0,5,4000.0000,'2016-12-01 00:00:00','2016-12-02 00:00:00',6,4,0,'',2,2,'initial, month-to-month','2017-07-18 21:55:07',0,'2017-07-18 19:13:43',0),(37,33,0,0,1,2,0,5,4000.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',2,2,'initial, month-to-month','2017-07-18 21:55:07',0,'2017-07-18 19:13:43',0),(38,33,0,0,1,2,0,5,4000.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',2,2,'initial, month-to-month','2017-07-18 21:55:07',0,'2017-07-18 19:13:43',0),(39,33,0,0,1,2,0,5,4000.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',2,2,'initial, month-to-month','2017-07-18 21:55:07',0,'2017-07-18 19:13:43',0),(40,33,0,0,1,2,0,5,4000.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',2,2,'initial, month-to-month','2017-07-18 21:55:07',0,'2017-07-18 19:13:43',0),(41,33,0,0,1,2,0,5,4000.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',2,2,'initial, month-to-month','2017-07-18 21:55:07',0,'2017-07-18 19:13:43',0),(42,33,0,0,1,2,0,5,4000.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',2,2,'initial, month-to-month','2017-07-18 21:55:07',0,'2017-07-18 19:13:43',0),(43,33,0,0,1,2,0,5,4000.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',2,0,'initial, month-to-month','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(44,0,0,0,1,2,0,5,628.4500,'2017-01-24 00:00:00','2017-01-24 00:00:00',0,0,0,'',14,2,'utilities reimbursement','2017-07-18 21:55:07',0,'2017-07-18 19:21:53',0),(45,0,0,0,1,2,0,5,175.0000,'2017-02-01 00:00:00','2017-02-12 00:00:00',0,0,0,'',14,2,'utilities charge','2017-07-18 21:55:07',0,'2017-07-18 19:23:02',0),(46,0,0,0,1,2,0,5,160.0000,'2017-03-01 00:00:00','2017-03-01 00:00:00',0,0,0,'',14,4,'Reversed by ASM00000057','2017-07-20 21:54:47',0,'2017-07-18 19:24:22',0),(47,0,0,0,1,2,0,5,350.0000,'2017-04-01 00:00:00','2017-04-01 00:00:00',0,0,0,'',14,2,'utilities charge','2017-07-18 21:55:07',0,'2017-07-18 19:25:28',0),(48,0,0,0,1,2,0,5,81.7900,'2017-04-01 00:00:00','2017-04-01 00:00:00',0,0,0,'',14,2,'retro utilities reimbursement','2017-07-18 21:55:07',0,'2017-07-18 19:26:10',0),(49,0,0,0,1,2,0,5,350.0000,'2017-05-01 00:00:00','2017-05-01 00:00:00',0,0,0,'',14,2,'utilities charge','2017-07-18 21:55:07',0,'2017-07-18 19:27:02',0),(50,0,0,0,1,2,0,5,350.0000,'2017-06-01 00:00:00','2017-06-01 00:00:00',0,0,0,'',14,1,'utilities charge','2017-07-18 21:55:07',0,'2017-07-18 19:27:48',0),(51,0,0,0,1,1,0,1,7000.0000,'2014-03-01 00:00:00','2014-03-01 00:00:00',0,0,0,'',21,2,'initial security deposit','2017-07-18 21:47:42',0,'2017-07-18 19:31:09',0),(52,0,0,0,1,3,0,4,8300.0000,'2016-07-01 00:00:00','2016-07-01 00:00:00',0,0,0,'',21,2,'initial security deposit','2017-07-18 21:49:52',0,'2017-07-18 19:32:01',0),(53,0,0,0,1,2,0,5,350.0000,'2017-07-01 00:00:00','2017-07-01 00:00:00',0,0,0,'',14,0,'utilities charge','2017-07-18 19:32:32',0,'2017-07-18 19:32:32',0),(54,0,0,0,1,2,0,5,15.0000,'2016-11-11 00:00:00','2016-11-11 00:00:00',0,0,0,'',26,4,'Reversed by ASM00000062','2017-08-09 21:22:14',0,'2017-07-18 19:53:22',0),(55,0,0,0,1,2,0,5,15.0000,'2017-02-13 00:00:00','2017-02-13 00:00:00',0,0,0,'',26,2,'$15 wire fee withheld by bank','2017-07-18 21:55:07',0,'2017-07-18 19:57:33',0),(56,0,0,0,1,2,0,5,15.0000,'2017-05-12 00:00:00','2017-05-12 00:00:00',0,0,0,'',26,2,'$15 wire fee withheld by bank','2017-07-18 21:55:07',0,'2017-07-18 20:00:26',0),(57,0,46,0,1,2,0,5,-160.0000,'2017-03-01 00:00:00','2017-03-01 00:00:00',0,0,0,'',14,6,'Reversal of ASM00000046','2017-07-20 21:54:47',0,'2017-07-20 21:54:47',0),(58,0,0,0,1,2,0,5,175.0000,'2017-03-01 00:00:00','2017-03-01 00:00:00',0,0,0,'',14,0,'utilities reimbursement','2017-07-20 21:55:39',0,'2017-07-20 21:55:39',0),(59,1,0,0,1,1,0,1,3750.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',2,0,'lease renewal','2017-08-02 16:53:39',0,'2017-08-02 16:53:39',0),(60,19,0,0,1,3,0,4,4150.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',2,0,'initial lease','2017-08-02 16:53:39',0,'2017-08-02 16:53:39',0),(61,33,0,0,1,2,0,5,4000.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',2,0,'initial, month-to-month','2017-08-02 16:53:39',0,'2017-08-02 16:53:39',0),(62,0,54,0,1,2,0,5,-15.0000,'2016-11-11 00:00:00','2016-11-11 00:00:00',0,0,0,'',26,6,'Reversal of ASM00000054','2017-08-09 21:22:14',0,'2017-08-09 21:22:14',0),(63,0,0,0,1,2,0,5,15.0000,'2016-11-11 00:00:00','2016-11-11 00:00:00',0,0,0,'',27,3,'$15 wire fee withheld by bank','2017-08-09 21:22:14',0,'2017-08-09 21:22:14',0),(64,1,0,0,1,1,0,1,3750.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',2,0,'lease renewal','2017-09-19 18:59:07',0,'2017-09-19 18:59:07',0),(65,19,0,0,1,3,0,4,4150.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',2,0,'initial lease','2017-09-19 18:59:07',0,'2017-09-19 18:59:07',0),(66,33,0,0,1,2,0,5,4000.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',2,0,'initial, month-to-month','2017-09-19 18:59:07',0,'2017-09-19 18:59:07',0);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
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
  `ClosePeriodTLID` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`BID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Business`
--

LOCK TABLES `Business` WRITE;
/*!40000 ALTER TABLE `Business` DISABLE KEYS */;
INSERT INTO `Business` VALUES (1,'REX','JGM First, LLC',6,4,4,0,'2017-06-13 05:39:46',0,'2017-06-14 18:26:46',0,0);
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
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0'
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
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0'
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
-- Table structure for table `BusinessProperties`
--

DROP TABLE IF EXISTS `BusinessProperties`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `BusinessProperties` (
  `BPID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Name` varchar(100) NOT NULL DEFAULT '',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `Data` json DEFAULT NULL,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`BPID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `BusinessProperties`
--

LOCK TABLES `BusinessProperties` WRITE;
/*!40000 ALTER TABLE `BusinessProperties` DISABLE KEYS */;
/*!40000 ALTER TABLE `BusinessProperties` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ClosePeriod`
--

DROP TABLE IF EXISTS `ClosePeriod`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ClosePeriod` (
  `CPID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `TLID` bigint(20) NOT NULL DEFAULT '0',
  `Dt` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`CPID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ClosePeriod`
--

LOCK TABLES `ClosePeriod` WRITE;
/*!40000 ALTER TABLE `ClosePeriod` DISABLE KEYS */;
/*!40000 ALTER TABLE `ClosePeriod` ENABLE KEYS */;
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `CustomAttr`
--

LOCK TABLES `CustomAttr` WRITE;
/*!40000 ALTER TABLE `CustomAttr` DISABLE KEYS */;
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
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `CARID` bigint(20) NOT NULL AUTO_INCREMENT,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`CARID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `CustomAttrRef`
--

LOCK TABLES `CustomAttrRef` WRITE;
/*!40000 ALTER TABLE `CustomAttrRef` DISABLE KEYS */;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DepositMethod`
--

LOCK TABLES `DepositMethod` WRITE;
/*!40000 ALTER TABLE `DepositMethod` DISABLE KEYS */;
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
INSERT INTO `Depository` VALUES (1,1,3,'FRB Operating  Account Deposit','4320','2017-07-05 02:13:56',0,'2017-07-05 02:13:37',0),(2,1,4,'FRB Security Deposit Account Deposit','6953','2017-07-05 02:14:34',0,'2017-07-05 02:14:34',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Expense`
--

LOCK TABLES `Expense` WRITE;
/*!40000 ALTER TABLE `Expense` DISABLE KEYS */;
INSERT INTO `Expense` VALUES (1,0,1,2,5,15.0000,'2016-11-11 00:00:00','',27,0,'','2017-08-10 19:00:35',0,'2017-08-10 19:00:35',0),(2,0,1,2,5,15.0000,'2017-05-12 00:00:00','',27,0,'','2017-08-11 01:55:11',0,'2017-08-11 01:55:11',0);
/*!40000 ALTER TABLE `Expense` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Flow`
--

DROP TABLE IF EXISTS `Flow`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Flow` (
  `FlowID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `UserRefNo` varchar(50) NOT NULL DEFAULT '',
  `FlowType` varchar(50) NOT NULL DEFAULT '',
  `ID` bigint(20) NOT NULL DEFAULT '0',
  `Data` json DEFAULT NULL,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`FlowID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Flow`
--

LOCK TABLES `Flow` WRITE;
/*!40000 ALTER TABLE `Flow` DISABLE KEYS */;
/*!40000 ALTER TABLE `Flow` ENABLE KEYS */;
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
  `AllowPost` tinyint(1) NOT NULL DEFAULT '0',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `Description` varchar(1024) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`LID`)
) ENGINE=InnoDB AUTO_INCREMENT=74 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `GLAccount`
--

LOCK TABLES `GLAccount` WRITE;
/*!40000 ALTER TABLE `GLAccount` DISABLE KEYS */;
INSERT INTO `GLAccount` VALUES (1,0,1,0,0,'10000',2,'Cash','Cash',0,0,'','2017-07-05 00:41:12',0,'2017-07-05 00:41:12',0),(2,1,1,0,0,'10001',2,'Petty Cash','Cash',1,0,'','2017-07-05 00:45:49',0,'2017-07-05 00:42:01',0),(3,9,1,0,0,'10104',2,'FRB Operating Account 4320','Cash',1,0,'','2017-07-05 00:46:29',0,'2017-07-05 00:42:28',0),(4,9,1,0,0,'10105',2,'FRB Security Deposit Account 6953','Cash',1,0,'','2017-07-05 00:46:39',0,'2017-07-05 00:42:58',0),(5,0,1,0,0,'11000',2,'Accounts Receivable','Accounts Receivable',0,0,'','2017-07-05 00:43:32',0,'2017-07-05 00:43:32',0),(8,1,1,0,0,'10200',2,'Credit Card Clearing','Cash',0,0,'','2017-07-05 00:50:00',0,'2017-07-05 00:45:39',0),(9,1,1,0,0,'10100',2,'Bank Accounts','Cash',0,0,'','2017-07-05 00:49:52',0,'2017-07-05 00:46:17',0),(10,0,1,0,0,'31000',2,'Security Deposits','Liabilities',1,0,'','2017-08-09 21:05:19',0,'2017-07-05 00:47:30',0),(11,0,1,0,0,'32000',2,'Collected Taxes','Liabilities',0,0,'','2017-08-09 21:05:38',0,'2017-07-05 00:48:23',0),(12,11,1,0,0,'32001',2,'Sales Tax Collected','Liabilities',1,0,'','2017-08-09 21:06:45',0,'2017-07-05 00:48:50',0),(13,11,1,0,0,'32002',2,'TOT Taxes Collected','Liabilities',1,0,'','2017-08-09 21:06:45',0,'2017-07-05 00:49:15',0),(14,0,1,0,0,'40000',2,'Gross Scheduled Rent Taxable','Income',1,0,'','2017-07-05 00:51:44',0,'2017-07-05 00:50:55',0),(15,0,1,0,0,'40001',2,'Gross Scheduled Rent  Non-Taxable','Income',1,0,'','2017-07-05 00:51:53',0,'2017-07-05 00:51:30',0),(16,0,1,0,0,'41000',2,'Income Offsets','Income Offsets',0,0,'','2017-07-07 19:13:04',0,'2017-07-05 00:52:25',0),(17,16,1,0,0,'41001',2,'Vacancy','Income Offsets',1,0,'','2017-07-07 19:13:04',0,'2017-07-05 00:52:55',0),(18,16,1,0,0,'41002',2,'Loss to Lease','Income Offsets',1,0,'','2017-07-07 19:13:04',0,'2017-07-05 00:53:22',0),(19,16,1,0,0,'41003',2,'Employee Concessions','Income Offsets',1,0,'','2017-07-07 19:13:04',0,'2017-07-05 00:53:53',0),(20,16,1,0,0,'41004',2,'Resident Concessions','Income Offsets',1,0,'','2017-07-07 19:13:04',0,'2017-07-05 00:54:34',0),(21,16,1,0,0,'41005',2,'Owner Concession','Income Offsets',1,0,'','2017-07-07 19:13:04',0,'2017-07-05 00:55:44',0),(22,16,1,0,0,'41006',2,'Administrative Concession','Income Offsets',1,0,'','2017-07-07 19:13:04',0,'2017-07-05 00:56:17',0),(23,16,1,0,0,'41007',2,'Off Line Renovations','Income Offsets',1,0,'','2017-07-07 19:13:04',0,'2017-07-05 00:57:03',0),(24,16,1,0,0,'41008',2,'Off Line Maintenance','Income Offsets',1,0,'','2017-07-07 19:13:04',0,'2017-07-05 00:57:33',0),(25,16,1,0,0,'41999',2,'Other Income Offsets','Income Offsets',1,0,'','2017-07-07 19:13:04',0,'2017-07-05 00:58:09',0),(26,0,1,0,0,'42000',2,'Service Fees','Other Income',0,0,'','2017-07-07 19:12:08',0,'2017-07-05 00:58:49',0),(27,26,1,0,0,'42001',2,'Broadcast and IT Services','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 00:59:50',0),(28,26,1,0,0,'42002',2,'Food Service','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:00:17',0),(29,26,1,0,0,'42003',2,'Linen Service','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:00:45',0),(30,26,1,0,0,'42004',2,'Wash N Fold Laundry','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:01:07',0),(31,26,1,0,0,'42999',2,'Other Service Fees','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:02:00',0),(32,0,1,0,0,'43000',2,'Utility Fees','Other Income',0,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:02:31',0),(33,32,1,0,0,'43001',2,'Electric Base Fee','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:02:58',0),(34,32,1,0,0,'43002',2,'Electric Overage','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:03:23',0),(35,32,1,0,0,'43003',2,'Water and Sewer Base Fee','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:04:03',0),(36,0,1,0,0,'43004',2,'Water and Sewer Overage','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:04:30',0),(37,32,1,0,0,'43005',2,'Gas Base Fee','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:04:58',0),(38,32,1,0,0,'43006',2,'Gas Overage','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:05:16',0),(39,32,1,0,0,'43007',2,'Trash Collection Base','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:05:38',0),(40,32,1,0,0,'43008',2,'Trash Collection Overage','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:06:01',0),(41,32,1,0,0,'43999',2,'Other Utility Fees','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:06:45',0),(42,0,1,0,0,'44000',2,'Miscellaneous Income','Other Income',0,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:08:18',0),(43,0,1,0,0,'44001',2,'Application Fees','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:08:41',0),(44,42,1,0,0,'44002',2,'Late Fees','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:09:04',0),(45,42,1,0,0,'44003',2,'Insufficient Funds Fee','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:09:56',0),(46,42,1,0,0,'44004',2,'Month to Month Fee','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:10:38',0),(47,42,1,0,0,'44005',2,'Unit Specialties','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:11:22',0),(48,42,1,0,0,'44006',2,'Termination and No-Show Fee','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:11:48',0),(49,42,1,0,0,'44007',2,'Pet Fee','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:12:20',0),(50,42,1,0,0,'44008',2,'Pet Rent','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:13:04',0),(51,42,1,0,0,'44009',2,'Expense Reimbursement','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:13:34',0),(52,42,1,0,0,'44010',2,'Maintenance Fee','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:14:14',0),(53,42,1,0,0,'44011',2,'Eviction Reimbursement','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:14:39',0),(54,42,1,0,0,'44012',2,'Extra Person Charge','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:15:02',0),(55,42,1,0,0,'44013',2,'Forfeited Security Deposit','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:15:29',0),(56,42,1,0,0,'44014',2,'CAM Fee','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:15:57',0),(57,42,1,0,0,'44999',2,'Other Miscellaneous Income','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:17:06',0),(58,0,1,0,0,'45000',2,'Tax Collections','Other Income',0,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:18:25',0),(59,58,1,0,0,'45001',2,'Sales Tax Collected','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:18:50',0),(60,58,1,0,0,'45002',2,'TOT Tax Collection','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:19:16',0),(61,0,1,0,0,'46000',2,'Business Income','Other Income',0,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:19:52',0),(62,61,1,0,0,'46001',2,'Convenience Store Income','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:20:20',0),(63,61,1,0,0,'46002',2,'Fitness Center Revenue','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:20:47',0),(64,61,1,0,0,'46003',2,'Vending Income','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:21:16',0),(65,61,1,0,0,'46004',2,'Restaurant Sales','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:22:00',0),(66,61,1,0,0,'46005',2,'Bar Sales','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:22:37',0),(67,61,1,0,0,'46006',2,'Spa Sales','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:23:00',0),(68,61,1,0,0,'46999',2,'Other Business Income','Other Income',1,0,'','2017-07-07 19:12:08',0,'2017-07-05 01:23:34',0),(69,0,1,0,0,'50000',2,'Expenses','Expense Account',0,1,'','2017-08-09 04:08:32',0,'2017-07-05 01:23:58',0),(70,69,1,0,0,'50001',2,'Cash Over/Short','Expense Account',1,1,'','2017-08-09 04:08:51',0,'2017-07-05 01:24:25',0),(71,69,1,0,0,'50002',2,'Bad Debt','Expense Account',1,1,'','2017-08-09 04:09:09',0,'2017-07-05 01:25:00',0),(72,1,1,0,0,'10999',2,'Unapplied Funds','Cash',1,0,'','2017-07-05 13:43:58',0,'2017-07-05 13:43:58',0),(73,69,1,0,0,'50003',2,'Bank Services Fee','Expense Account',1,1,'','2017-08-09 04:10:19',0,'2017-08-09 04:10:19',0);
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
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `InvoiceASMID` bigint(20) NOT NULL AUTO_INCREMENT,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`InvoiceASMID`)
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
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `InvoicePayorID` bigint(20) NOT NULL AUTO_INCREMENT,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`InvoicePayorID`)
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
) ENGINE=InnoDB AUTO_INCREMENT=166 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Journal`
--

LOCK TABLES `Journal` WRITE;
/*!40000 ALTER TABLE `Journal` DISABLE KEYS */;
INSERT INTO `Journal` VALUES (1,1,'2016-03-01 00:00:00',3750.0000,1,2,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(2,1,'2016-04-01 00:00:00',3750.0000,1,3,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(3,1,'2016-05-01 00:00:00',3750.0000,1,4,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(4,1,'2016-06-01 00:00:00',3750.0000,1,5,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(5,1,'2016-07-01 00:00:00',3750.0000,1,6,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(6,1,'2016-08-01 00:00:00',3750.0000,1,7,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(7,1,'2016-09-01 00:00:00',3750.0000,1,8,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(8,1,'2016-10-01 00:00:00',3750.0000,1,9,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(9,1,'2016-11-01 00:00:00',3750.0000,1,10,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(10,1,'2016-12-01 00:00:00',3750.0000,1,11,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(11,1,'2017-01-01 00:00:00',3750.0000,1,12,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(12,1,'2017-02-01 00:00:00',3750.0000,1,13,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(13,1,'2017-03-01 00:00:00',3750.0000,1,14,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(14,1,'2017-04-01 00:00:00',3750.0000,1,15,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(15,1,'2017-05-01 00:00:00',3750.0000,1,16,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(16,1,'2017-06-01 00:00:00',3750.0000,1,17,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(17,1,'2017-07-01 00:00:00',3750.0000,1,18,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(18,1,'2016-07-01 00:00:00',4150.0000,1,20,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(19,1,'2016-08-01 00:00:00',4150.0000,1,21,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(20,1,'2016-09-01 00:00:00',4150.0000,1,22,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(21,1,'2016-10-01 00:00:00',4150.0000,1,23,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(22,1,'2016-11-01 00:00:00',4150.0000,1,24,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(23,1,'2016-12-01 00:00:00',4150.0000,1,25,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(24,1,'2017-01-01 00:00:00',4150.0000,1,26,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(25,1,'2017-02-01 00:00:00',4150.0000,1,27,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(26,1,'2017-03-01 00:00:00',4150.0000,1,28,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(27,1,'2017-04-01 00:00:00',4150.0000,1,29,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(28,1,'2017-05-01 00:00:00',4150.0000,1,30,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(29,1,'2017-06-01 00:00:00',4150.0000,1,31,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(30,1,'2017-07-01 00:00:00',4150.0000,1,32,'','2017-07-18 19:12:54',0,'2017-07-18 19:12:54',0),(31,1,'2016-10-01 00:00:00',4000.0000,1,34,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(32,1,'2016-11-01 00:00:00',4000.0000,1,35,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(33,1,'2016-12-01 00:00:00',4000.0000,1,36,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(34,1,'2017-01-01 00:00:00',4000.0000,1,37,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(35,1,'2017-02-01 00:00:00',4000.0000,1,38,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(36,1,'2017-03-01 00:00:00',4000.0000,1,39,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(37,1,'2017-04-01 00:00:00',4000.0000,1,40,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(38,1,'2017-05-01 00:00:00',4000.0000,1,41,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(39,1,'2017-06-01 00:00:00',4000.0000,1,42,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(40,1,'2017-07-01 00:00:00',4000.0000,1,43,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(41,1,'2017-01-24 00:00:00',628.4500,1,44,'','2017-07-18 19:21:53',0,'2017-07-18 19:21:53',0),(42,1,'2017-02-01 00:00:00',175.0000,1,45,'','2017-07-18 19:23:02',0,'2017-07-18 19:23:02',0),(43,1,'2017-03-01 00:00:00',160.0000,1,46,'','2017-07-18 19:24:22',0,'2017-07-18 19:24:22',0),(44,1,'2017-04-01 00:00:00',350.0000,1,47,'','2017-07-18 19:25:28',0,'2017-07-18 19:25:28',0),(45,1,'2017-04-01 00:00:00',81.7900,1,48,'','2017-07-18 19:26:10',0,'2017-07-18 19:26:10',0),(46,1,'2017-05-01 00:00:00',350.0000,1,49,'','2017-07-18 19:27:02',0,'2017-07-18 19:27:02',0),(47,1,'2017-06-01 00:00:00',350.0000,1,50,'','2017-07-18 19:27:48',0,'2017-07-18 19:27:48',0),(48,1,'2014-03-01 00:00:00',7000.0000,1,51,'','2017-07-18 19:31:09',0,'2017-07-18 19:31:09',0),(49,1,'2016-07-01 00:00:00',8300.0000,1,52,'','2017-07-18 19:32:01',0,'2017-07-18 19:32:01',0),(50,1,'2017-07-01 00:00:00',350.0000,1,53,'','2017-07-18 19:32:32',0,'2017-07-18 19:32:32',0),(51,1,'2014-03-01 00:00:00',7000.0000,2,1,'','2017-07-18 19:35:44',0,'2017-07-18 19:35:44',0),(52,1,'2016-03-01 00:00:00',3750.0000,2,2,'','2017-07-18 19:36:47',0,'2017-07-18 19:36:47',0),(53,1,'2016-04-01 00:00:00',3750.0000,2,3,'','2017-07-18 19:37:04',0,'2017-07-18 19:37:04',0),(54,1,'2017-05-01 00:00:00',3750.0000,2,4,'','2017-07-18 19:37:26',0,'2017-07-18 19:37:26',0),(55,1,'2017-05-01 00:00:00',-3750.0000,2,5,'','2017-07-18 19:38:20',0,'2017-07-18 19:38:20',0),(56,1,'2016-05-01 00:00:00',3750.0000,2,6,'','2017-07-18 19:38:20',0,'2017-07-18 19:38:20',0),(57,1,'2016-05-01 00:00:00',3750.0000,2,7,'','2017-07-18 19:38:55',0,'2017-07-18 19:38:55',0),(58,1,'2016-06-01 00:00:00',3750.0000,2,8,'','2017-07-18 19:39:17',0,'2017-07-18 19:39:17',0),(59,1,'2016-07-01 00:00:00',3750.0000,2,9,'','2017-07-18 19:39:41',0,'2017-07-18 19:39:41',0),(60,1,'2016-08-01 00:00:00',3750.0000,2,10,'','2017-07-18 19:40:00',0,'2017-07-18 19:40:00',0),(61,1,'2016-09-01 00:00:00',3750.0000,2,11,'','2017-07-18 19:40:17',0,'2017-07-18 19:40:17',0),(62,1,'2016-10-01 00:00:00',3750.0000,2,12,'','2017-07-18 19:40:33',0,'2017-07-18 19:40:33',0),(63,1,'2016-11-01 00:00:00',3750.0000,2,13,'','2017-07-18 19:40:50',0,'2017-07-18 19:40:50',0),(64,1,'2016-12-01 00:00:00',3750.0000,2,14,'','2017-07-18 19:41:03',0,'2017-07-18 19:41:03',0),(65,1,'2017-01-01 00:00:00',3750.0000,2,15,'','2017-07-18 19:41:23',0,'2017-07-18 19:41:23',0),(66,1,'2017-02-01 00:00:00',3750.0000,2,16,'','2017-07-18 19:41:38',0,'2017-07-18 19:41:38',0),(67,1,'2017-03-01 00:00:00',3750.0000,2,17,'','2017-07-18 19:42:11',0,'2017-07-18 19:42:11',0),(68,1,'2017-04-01 00:00:00',3750.0000,2,18,'','2017-07-18 19:42:27',0,'2017-07-18 19:42:27',0),(69,1,'2017-05-01 00:00:00',3750.0000,2,19,'','2017-07-18 19:42:43',0,'2017-07-18 19:42:43',0),(70,1,'2017-06-01 00:00:00',3750.0000,2,20,'','2017-07-18 19:43:06',0,'2017-07-18 19:43:06',0),(71,1,'2017-07-01 00:00:00',3750.0000,2,21,'','2017-07-18 19:43:23',0,'2017-07-18 19:43:23',0),(72,1,'2016-05-01 00:00:00',-3750.0000,2,22,'','2017-07-18 19:44:38',0,'2017-07-18 19:44:38',0),(73,1,'2016-07-01 00:00:00',8300.0000,2,23,'','2017-07-18 19:45:29',0,'2017-07-18 19:45:29',0),(74,1,'2016-07-01 00:00:00',4150.0000,2,24,'','2017-07-18 19:45:55',0,'2017-07-18 19:45:55',0),(75,1,'2016-08-01 00:00:00',4150.0000,2,25,'','2017-07-18 19:46:13',0,'2017-07-18 19:46:13',0),(76,1,'2016-09-01 00:00:00',4150.0000,2,26,'','2017-07-18 19:46:31',0,'2017-07-18 19:46:31',0),(77,1,'2016-10-01 00:00:00',4150.0000,2,27,'','2017-07-18 19:46:50',0,'2017-07-18 19:46:50',0),(78,1,'2016-11-01 00:00:00',4150.0000,2,28,'','2017-07-18 19:47:09',0,'2017-07-18 19:47:09',0),(79,1,'2016-12-01 00:00:00',4150.0000,2,29,'','2017-07-18 19:47:28',0,'2017-07-18 19:47:28',0),(80,1,'2017-01-01 00:00:00',4150.0000,2,30,'','2017-07-18 19:47:59',0,'2017-07-18 19:47:59',0),(81,1,'2017-02-01 00:00:00',4150.0000,2,31,'','2017-07-18 19:48:18',0,'2017-07-18 19:48:18',0),(82,1,'2017-03-01 00:00:00',4150.0000,2,32,'','2017-07-18 19:48:38',0,'2017-07-18 19:48:38',0),(83,1,'2017-04-01 00:00:00',4150.0000,2,33,'','2017-07-18 19:48:56',0,'2017-07-18 19:48:56',0),(84,1,'2017-05-01 00:00:00',4150.0000,2,34,'','2017-07-18 19:49:12',0,'2017-07-18 19:49:12',0),(85,1,'2017-06-01 00:00:00',4150.0000,2,35,'','2017-07-18 19:49:28',0,'2017-07-18 19:49:28',0),(86,1,'2017-07-01 00:00:00',4150.0000,2,36,'','2017-07-18 19:49:51',0,'2017-07-18 19:49:51',0),(87,1,'2016-10-03 00:00:00',4000.0000,2,37,'','2017-07-18 19:51:26',0,'2017-07-18 19:51:26',0),(88,1,'2016-11-11 00:00:00',11985.0000,2,38,'','2017-07-18 19:52:33',0,'2017-07-18 19:52:33',0),(89,1,'2016-11-11 00:00:00',0.0000,1,54,'','2017-07-18 19:53:22',0,'2017-07-18 19:53:22',0),(90,1,'2017-02-03 00:00:00',628.4500,2,39,'','2017-07-18 19:54:30',0,'2017-07-18 19:54:30',0),(91,1,'2017-02-13 00:00:00',8335.0000,2,40,'','2017-07-18 19:56:46',0,'2017-07-18 19:56:46',0),(92,1,'2017-02-13 00:00:00',0.0000,1,55,'','2017-07-18 19:57:33',0,'2017-07-18 19:57:33',0),(93,1,'2017-05-12 00:00:00',13116.7900,2,41,'','2017-07-18 19:59:41',0,'2017-07-18 19:59:41',0),(94,1,'2017-05-12 00:00:00',0.0000,1,56,'','2017-07-18 20:00:26',0,'2017-07-18 20:00:26',0),(95,1,'2014-03-01 00:00:00',7000.0000,2,1,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(96,1,'2016-03-01 00:00:00',3750.0000,2,2,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(97,1,'2016-04-01 00:00:00',3750.0000,2,3,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(98,1,'2016-05-01 00:00:00',3750.0000,2,7,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(99,1,'2016-06-01 00:00:00',3750.0000,2,8,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(100,1,'2016-07-01 00:00:00',3750.0000,2,9,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(101,1,'2016-08-01 00:00:00',3750.0000,2,10,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(102,1,'2016-09-01 00:00:00',3750.0000,2,11,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(103,1,'2016-10-01 00:00:00',3750.0000,2,12,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(104,1,'2016-11-01 00:00:00',3750.0000,2,13,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(105,1,'2016-12-01 00:00:00',3750.0000,2,14,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(106,1,'2017-01-01 00:00:00',3750.0000,2,15,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(107,1,'2017-02-01 00:00:00',3750.0000,2,16,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(108,1,'2017-03-01 00:00:00',3750.0000,2,17,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(109,1,'2017-04-01 00:00:00',3750.0000,2,18,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(110,1,'2017-05-01 00:00:00',3750.0000,2,19,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(111,1,'2017-06-01 00:00:00',3750.0000,2,20,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(112,1,'2017-07-01 00:00:00',3750.0000,2,21,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(113,1,'2016-07-01 00:00:00',4150.0000,2,23,'','2017-07-18 21:49:52',0,'2017-07-18 21:49:52',0),(114,1,'2016-07-01 00:00:00',4150.0000,2,23,'','2017-07-18 21:49:52',0,'2017-07-18 21:49:52',0),(115,1,'2016-07-01 00:00:00',4150.0000,2,24,'','2017-07-18 21:49:52',0,'2017-07-18 21:49:52',0),(116,1,'2016-08-01 00:00:00',4150.0000,2,25,'','2017-07-18 21:49:52',0,'2017-07-18 21:49:52',0),(117,1,'2016-09-01 00:00:00',4150.0000,2,26,'','2017-07-18 21:49:52',0,'2017-07-18 21:49:52',0),(118,1,'2016-10-01 00:00:00',4150.0000,2,27,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(119,1,'2016-11-01 00:00:00',4150.0000,2,28,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(120,1,'2016-12-01 00:00:00',4150.0000,2,29,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(121,1,'2017-01-01 00:00:00',4150.0000,2,30,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(122,1,'2017-02-01 00:00:00',4150.0000,2,31,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(123,1,'2017-03-01 00:00:00',4150.0000,2,32,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(124,1,'2017-04-01 00:00:00',4150.0000,2,33,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(125,1,'2017-05-01 00:00:00',4150.0000,2,34,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(126,1,'2017-06-01 00:00:00',4150.0000,2,35,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(127,1,'2017-07-01 00:00:00',4150.0000,2,36,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(128,1,'2016-10-03 00:00:00',4000.0000,2,37,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(129,1,'2016-11-11 00:00:00',4000.0000,2,38,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(130,1,'2016-11-11 00:00:00',15.0000,2,38,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(131,1,'2016-11-11 00:00:00',4000.0000,2,38,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(132,1,'2016-11-11 00:00:00',3970.0000,2,38,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(133,1,'2017-02-03 00:00:00',30.0000,2,39,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(134,1,'2017-02-03 00:00:00',598.4500,2,39,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(135,1,'2017-02-13 00:00:00',30.0000,2,40,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(136,1,'2017-02-13 00:00:00',4000.0000,2,40,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(137,1,'2017-02-13 00:00:00',175.0000,2,40,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(138,1,'2017-02-13 00:00:00',15.0000,2,40,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(139,1,'2017-02-13 00:00:00',4000.0000,2,40,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(140,1,'2017-02-13 00:00:00',115.0000,2,40,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(141,1,'2017-05-12 00:00:00',45.0000,2,41,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(142,1,'2017-05-12 00:00:00',4000.0000,2,41,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(143,1,'2017-05-12 00:00:00',350.0000,2,41,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(144,1,'2017-05-12 00:00:00',81.7900,2,41,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(145,1,'2017-05-12 00:00:00',4000.0000,2,41,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(146,1,'2017-05-12 00:00:00',350.0000,2,41,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(147,1,'2017-05-12 00:00:00',15.0000,2,41,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(148,1,'2017-05-12 00:00:00',4000.0000,2,41,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(149,1,'2017-05-12 00:00:00',275.0000,2,41,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(150,1,'2017-03-01 00:00:00',-160.0000,1,57,'','2017-07-20 21:54:47',0,'2017-07-20 21:54:47',0),(151,1,'2017-07-20 21:54:47',-115.0000,1,57,'','2017-07-20 21:54:47',0,'2017-07-20 21:54:47',0),(152,1,'2017-07-20 21:54:47',-45.0000,1,57,'','2017-07-20 21:54:47',0,'2017-07-20 21:54:47',0),(153,1,'2017-03-01 00:00:00',175.0000,1,58,'','2017-07-20 21:55:39',0,'2017-07-20 21:55:39',0),(154,1,'2017-07-01 00:00:00',-4150.0000,2,42,'','2017-07-24 08:06:52',0,'2017-07-24 08:06:52',0),(155,1,'2017-07-24 08:06:53',-4150.0000,2,42,'','2017-07-24 08:06:52',0,'2017-07-24 08:06:52',0),(156,1,'2017-07-01 00:00:00',4150.0000,2,43,'','2017-07-24 08:06:52',0,'2017-07-24 08:06:52',0),(157,1,'2017-08-01 00:00:00',3750.0000,1,59,'','2017-08-02 16:53:39',0,'2017-08-02 16:53:39',0),(158,1,'2017-08-01 00:00:00',4150.0000,1,60,'','2017-08-02 16:53:39',0,'2017-08-02 16:53:39',0),(159,1,'2017-08-01 00:00:00',4000.0000,1,61,'','2017-08-02 16:53:39',0,'2017-08-02 16:53:39',0),(160,1,'2016-11-11 00:00:00',0.0000,1,62,'','2017-08-09 21:22:14',0,'2017-08-09 21:22:14',0),(161,1,'2017-08-09 21:22:15',-15.0000,1,62,'','2017-08-09 21:22:14',0,'2017-08-09 21:22:14',0),(162,1,'2016-11-11 00:00:00',15.0000,1,63,'','2017-08-09 21:22:14',0,'2017-08-09 21:22:14',0),(163,1,'2017-09-01 00:00:00',3750.0000,1,64,'','2017-09-19 18:59:07',0,'2017-09-19 18:59:07',0),(164,1,'2017-09-01 00:00:00',4150.0000,1,65,'','2017-09-19 18:59:07',0,'2017-09-19 18:59:07',0),(165,1,'2017-09-01 00:00:00',4000.0000,1,66,'','2017-09-19 18:59:07',0,'2017-09-19 18:59:07',0);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`JAID`)
) ENGINE=InnoDB AUTO_INCREMENT=166 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `JournalAllocation`
--

LOCK TABLES `JournalAllocation` WRITE;
/*!40000 ALTER TABLE `JournalAllocation` DISABLE KEYS */;
INSERT INTO `JournalAllocation` VALUES (1,1,1,1,1,0,0,3750.0000,2,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(2,1,2,1,1,0,0,3750.0000,3,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(3,1,3,1,1,0,0,3750.0000,4,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(4,1,4,1,1,0,0,3750.0000,5,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(5,1,5,1,1,0,0,3750.0000,6,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(6,1,6,1,1,0,0,3750.0000,7,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(7,1,7,1,1,0,0,3750.0000,8,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(8,1,8,1,1,0,0,3750.0000,9,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(9,1,9,1,1,0,0,3750.0000,10,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(10,1,10,1,1,0,0,3750.0000,11,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(11,1,11,1,1,0,0,3750.0000,12,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(12,1,12,1,1,0,0,3750.0000,13,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(13,1,13,1,1,0,0,3750.0000,14,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(14,1,14,1,1,0,0,3750.0000,15,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(15,1,15,1,1,0,0,3750.0000,16,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(16,1,16,1,1,0,0,3750.0000,17,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(17,1,17,1,1,0,0,3750.0000,18,0,'d 11000 3750.00, c 40001 3750.00','2017-07-18 19:10:52',0,'2018-01-01 09:59:48',0),(18,1,18,3,4,0,0,4150.0000,20,0,'d 11000 4150.00, c 40001 4150.00','2017-07-18 19:12:53',0,'2018-01-01 09:59:48',0),(19,1,19,3,4,0,0,4150.0000,21,0,'d 11000 4150.00, c 40001 4150.00','2017-07-18 19:12:53',0,'2018-01-01 09:59:48',0),(20,1,20,3,4,0,0,4150.0000,22,0,'d 11000 4150.00, c 40001 4150.00','2017-07-18 19:12:53',0,'2018-01-01 09:59:48',0),(21,1,21,3,4,0,0,4150.0000,23,0,'d 11000 4150.00, c 40001 4150.00','2017-07-18 19:12:53',0,'2018-01-01 09:59:48',0),(22,1,22,3,4,0,0,4150.0000,24,0,'d 11000 4150.00, c 40001 4150.00','2017-07-18 19:12:53',0,'2018-01-01 09:59:48',0),(23,1,23,3,4,0,0,4150.0000,25,0,'d 11000 4150.00, c 40001 4150.00','2017-07-18 19:12:53',0,'2018-01-01 09:59:48',0),(24,1,24,3,4,0,0,4150.0000,26,0,'d 11000 4150.00, c 40001 4150.00','2017-07-18 19:12:53',0,'2018-01-01 09:59:48',0),(25,1,25,3,4,0,0,4150.0000,27,0,'d 11000 4150.00, c 40001 4150.00','2017-07-18 19:12:53',0,'2018-01-01 09:59:48',0),(26,1,26,3,4,0,0,4150.0000,28,0,'d 11000 4150.00, c 40001 4150.00','2017-07-18 19:12:53',0,'2018-01-01 09:59:48',0),(27,1,27,3,4,0,0,4150.0000,29,0,'d 11000 4150.00, c 40001 4150.00','2017-07-18 19:12:53',0,'2018-01-01 09:59:48',0),(28,1,28,3,4,0,0,4150.0000,30,0,'d 11000 4150.00, c 40001 4150.00','2017-07-18 19:12:53',0,'2018-01-01 09:59:48',0),(29,1,29,3,4,0,0,4150.0000,31,0,'d 11000 4150.00, c 40001 4150.00','2017-07-18 19:12:53',0,'2018-01-01 09:59:48',0),(30,1,30,3,4,0,0,4150.0000,32,0,'d 11000 4150.00, c 40001 4150.00','2017-07-18 19:12:54',0,'2018-01-01 09:59:48',0),(31,1,31,2,5,0,0,4000.0000,34,0,'d 11000 4000.00, c 40001 4000.00','2017-07-18 19:13:43',0,'2018-01-01 09:59:48',0),(32,1,32,2,5,0,0,4000.0000,35,0,'d 11000 4000.00, c 40001 4000.00','2017-07-18 19:13:43',0,'2018-01-01 09:59:48',0),(33,1,33,2,5,0,0,4000.0000,36,0,'d 11000 4000.00, c 40001 4000.00','2017-07-18 19:13:43',0,'2018-01-01 09:59:48',0),(34,1,34,2,5,0,0,4000.0000,37,0,'d 11000 4000.00, c 40001 4000.00','2017-07-18 19:13:43',0,'2018-01-01 09:59:48',0),(35,1,35,2,5,0,0,4000.0000,38,0,'d 11000 4000.00, c 40001 4000.00','2017-07-18 19:13:43',0,'2018-01-01 09:59:48',0),(36,1,36,2,5,0,0,4000.0000,39,0,'d 11000 4000.00, c 40001 4000.00','2017-07-18 19:13:43',0,'2018-01-01 09:59:48',0),(37,1,37,2,5,0,0,4000.0000,40,0,'d 11000 4000.00, c 40001 4000.00','2017-07-18 19:13:43',0,'2018-01-01 09:59:48',0),(38,1,38,2,5,0,0,4000.0000,41,0,'d 11000 4000.00, c 40001 4000.00','2017-07-18 19:13:43',0,'2018-01-01 09:59:48',0),(39,1,39,2,5,0,0,4000.0000,42,0,'d 11000 4000.00, c 40001 4000.00','2017-07-18 19:13:43',0,'2018-01-01 09:59:48',0),(40,1,40,2,5,0,0,4000.0000,43,0,'d 11000 4000.00, c 40001 4000.00','2017-07-18 19:13:43',0,'2018-01-01 09:59:48',0),(41,1,41,2,5,0,0,628.4500,44,0,'d 11000 628.45, c 44009 628.45','2017-07-18 19:21:53',0,'2018-01-01 09:59:48',0),(42,1,42,2,5,0,0,175.0000,45,0,'d 11000 175.00, c 44009 175.00','2017-07-18 19:23:02',0,'2018-01-01 09:59:48',0),(43,1,43,2,5,0,0,160.0000,46,0,'d 11000 160.00, c 44009 160.00','2017-07-18 19:24:22',0,'2018-01-01 09:59:48',0),(44,1,44,2,5,0,0,350.0000,47,0,'d 11000 350.00, c 44009 350.00','2017-07-18 19:25:28',0,'2018-01-01 09:59:48',0),(45,1,45,2,5,0,0,81.7900,48,0,'d 11000 81.79, c 44009 81.79','2017-07-18 19:26:10',0,'2018-01-01 09:59:48',0),(46,1,46,2,5,0,0,350.0000,49,0,'d 11000 350.00, c 44009 350.00','2017-07-18 19:27:02',0,'2018-01-01 09:59:48',0),(47,1,47,2,5,0,0,350.0000,50,0,'d 11000 350.00, c 44009 350.00','2017-07-18 19:27:48',0,'2018-01-01 09:59:48',0),(48,1,48,1,1,0,0,7000.0000,51,0,'d 11000 7000.00, c 31000 7000.00','2017-07-18 19:31:09',0,'2018-01-01 09:59:48',0),(49,1,49,3,4,0,0,8300.0000,52,0,'d 11000 8300.00, c 31000 8300.00','2017-07-18 19:32:01',0,'2018-01-01 09:59:48',0),(50,1,50,2,5,0,0,350.0000,53,0,'d 11000 350.00, c 44009 350.00','2017-07-18 19:32:32',0,'2018-01-01 09:59:48',0),(51,1,51,0,0,1,0,7000.0000,0,0,'d 10105 _, c 10999 _','2017-07-18 19:35:44',0,'2018-01-01 09:59:48',0),(52,1,52,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:36:47',0,'2018-01-01 09:59:48',0),(53,1,53,0,0,0,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:37:04',0,'2018-01-01 09:59:48',0),(54,1,54,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:37:26',0,'2018-01-01 09:59:48',0),(55,1,55,0,0,1,0,-3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:38:20',0,'2018-01-01 09:59:48',0),(56,1,56,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:38:20',0,'2018-01-01 09:59:48',0),(57,1,57,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:38:55',0,'2018-01-01 09:59:48',0),(58,1,58,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:39:17',0,'2018-01-01 09:59:48',0),(59,1,59,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:39:41',0,'2018-01-01 09:59:48',0),(60,1,60,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:40:00',0,'2018-01-01 09:59:48',0),(61,1,61,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:40:17',0,'2018-01-01 09:59:48',0),(62,1,62,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:40:33',0,'2018-01-01 09:59:48',0),(63,1,63,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:40:50',0,'2018-01-01 09:59:48',0),(64,1,64,0,0,0,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:41:03',0,'2018-01-01 09:59:48',0),(65,1,65,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:41:23',0,'2018-01-01 09:59:48',0),(66,1,66,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:41:38',0,'2018-01-01 09:59:48',0),(67,1,67,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:42:11',0,'2018-01-01 09:59:48',0),(68,1,68,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:42:27',0,'2018-01-01 09:59:48',0),(69,1,69,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:42:43',0,'2018-01-01 09:59:48',0),(70,1,70,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:43:06',0,'2018-01-01 09:59:48',0),(71,1,71,0,0,1,0,3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:43:23',0,'2018-01-01 09:59:48',0),(72,1,72,0,0,1,0,-3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:44:38',0,'2018-01-01 09:59:48',0),(73,1,73,0,0,3,0,8300.0000,0,0,'d 10105 _, c 10999 _','2017-07-18 19:45:29',0,'2018-01-01 09:59:48',0),(74,1,74,0,0,3,0,4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:45:55',0,'2018-01-01 09:59:48',0),(75,1,75,0,0,3,0,4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:46:13',0,'2018-01-01 09:59:48',0),(76,1,76,0,0,3,0,4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:46:31',0,'2018-01-01 09:59:48',0),(77,1,77,0,0,3,0,4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:46:50',0,'2018-01-01 09:59:48',0),(78,1,78,0,0,3,0,4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:47:09',0,'2018-01-01 09:59:48',0),(79,1,79,0,0,3,0,4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:47:28',0,'2018-01-01 09:59:48',0),(80,1,80,0,0,3,0,4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:47:59',0,'2018-01-01 09:59:48',0),(81,1,81,0,0,3,0,4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:48:18',0,'2018-01-01 09:59:48',0),(82,1,82,0,0,3,0,4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:48:38',0,'2018-01-01 09:59:48',0),(83,1,83,0,0,3,0,4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:48:56',0,'2018-01-01 09:59:48',0),(84,1,84,0,0,3,0,4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:49:12',0,'2018-01-01 09:59:48',0),(85,1,85,0,0,3,0,4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:49:28',0,'2018-01-01 09:59:48',0),(86,1,86,0,0,3,0,4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:49:51',0,'2018-01-01 09:59:48',0),(87,1,87,0,0,6,0,4000.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:51:26',0,'2018-01-01 09:59:48',0),(88,1,88,0,0,6,0,11985.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:52:33',0,'2018-01-01 09:59:48',0),(89,1,89,2,5,0,0,0.0000,54,0,'d _ 0.00, c 11000 15.00','2017-07-18 19:53:22',0,'2018-01-01 09:59:48',0),(90,1,90,0,0,6,0,628.4500,0,0,'d 10104 _, c 10999 _','2017-07-18 19:54:30',0,'2018-01-01 09:59:48',0),(91,1,91,0,0,6,0,8335.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:56:46',0,'2018-01-01 09:59:48',0),(92,1,92,2,5,0,0,0.0000,55,0,'d _ 0.00, c 11000 15.00','2017-07-18 19:57:33',0,'2018-01-01 09:59:48',0),(93,1,93,0,0,6,0,13116.7900,0,0,'d 10104 _, c 10999 _','2017-07-18 19:59:41',0,'2018-01-01 09:59:48',0),(94,1,94,2,5,0,0,0.0000,56,0,'d _ 0.00, c 11000 15.00','2017-07-18 20:00:26',0,'2018-01-01 09:59:48',0),(95,1,95,1,1,1,1,7000.0000,51,0,'ASM(51) d 10999 7000.00,c 11000 7000.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(96,1,96,1,1,1,2,3750.0000,2,0,'ASM(2) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(97,1,97,1,1,1,3,3750.0000,3,0,'ASM(3) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(98,1,98,1,1,1,7,3750.0000,4,0,'ASM(4) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(99,1,99,1,1,1,8,3750.0000,5,0,'ASM(5) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(100,1,100,1,1,1,9,3750.0000,6,0,'ASM(6) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(101,1,101,1,1,1,10,3750.0000,7,0,'ASM(7) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(102,1,102,1,1,1,11,3750.0000,8,0,'ASM(8) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(103,1,103,1,1,1,12,3750.0000,9,0,'ASM(9) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(104,1,104,1,1,1,13,3750.0000,10,0,'ASM(10) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(105,1,105,1,1,1,14,3750.0000,11,0,'ASM(11) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(106,1,106,1,1,1,15,3750.0000,12,0,'ASM(12) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(107,1,107,1,1,1,16,3750.0000,13,0,'ASM(13) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(108,1,108,1,1,1,17,3750.0000,14,0,'ASM(14) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(109,1,109,1,1,1,18,3750.0000,15,0,'ASM(15) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(110,1,110,1,1,1,19,3750.0000,16,0,'ASM(16) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(111,1,111,1,1,1,20,3750.0000,17,0,'ASM(17) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(112,1,112,1,1,1,21,3750.0000,18,0,'ASM(18) d 10999 3750.00,c 11000 3750.00','2017-07-18 21:47:42',0,'2018-01-01 09:59:48',0),(113,1,113,3,4,3,23,4150.0000,20,0,'ASM(20) d 10999 4150.00,c 11000 4150.00','2017-07-18 21:49:52',0,'2018-01-01 09:59:48',0),(114,1,114,3,4,3,23,4150.0000,52,0,'ASM(52) d 10999 4150.00,c 11000 4150.00','2017-07-18 21:49:52',0,'2018-01-01 09:59:48',0),(115,1,115,3,4,3,24,4150.0000,52,0,'ASM(52) d 10999 4150.00,c 11000 4150.00','2017-07-18 21:49:52',0,'2018-01-01 09:59:48',0),(116,1,116,3,4,3,25,4150.0000,21,0,'ASM(21) d 10999 4150.00,c 11000 4150.00','2017-07-18 21:49:52',0,'2018-01-01 09:59:48',0),(117,1,117,3,4,3,26,4150.0000,22,0,'ASM(22) d 10999 4150.00,c 11000 4150.00','2017-07-18 21:49:52',0,'2018-01-01 09:59:48',0),(118,1,118,3,4,3,27,4150.0000,23,0,'ASM(23) d 10999 4150.00,c 11000 4150.00','2017-07-18 21:49:53',0,'2018-01-01 09:59:48',0),(119,1,119,3,4,3,28,4150.0000,24,0,'ASM(24) d 10999 4150.00,c 11000 4150.00','2017-07-18 21:49:53',0,'2018-01-01 09:59:48',0),(120,1,120,3,4,3,29,4150.0000,25,0,'ASM(25) d 10999 4150.00,c 11000 4150.00','2017-07-18 21:49:53',0,'2018-01-01 09:59:48',0),(121,1,121,3,4,3,30,4150.0000,26,0,'ASM(26) d 10999 4150.00,c 11000 4150.00','2017-07-18 21:49:53',0,'2018-01-01 09:59:48',0),(122,1,122,3,4,3,31,4150.0000,27,0,'ASM(27) d 10999 4150.00,c 11000 4150.00','2017-07-18 21:49:53',0,'2018-01-01 09:59:48',0),(123,1,123,3,4,3,32,4150.0000,28,0,'ASM(28) d 10999 4150.00,c 11000 4150.00','2017-07-18 21:49:53',0,'2018-01-01 09:59:48',0),(124,1,124,3,4,3,33,4150.0000,29,0,'ASM(29) d 10999 4150.00,c 11000 4150.00','2017-07-18 21:49:53',0,'2018-01-01 09:59:48',0),(125,1,125,3,4,3,34,4150.0000,30,0,'ASM(30) d 10999 4150.00,c 11000 4150.00','2017-07-18 21:49:53',0,'2018-01-01 09:59:48',0),(126,1,126,3,4,3,35,4150.0000,31,0,'ASM(31) d 10999 4150.00,c 11000 4150.00','2017-07-18 21:49:53',0,'2018-01-01 09:59:48',0),(127,1,127,3,4,3,36,4150.0000,32,0,'ASM(32) d 10999 4150.00,c 11000 4150.00','2017-07-18 21:49:53',0,'2018-01-01 09:59:48',0),(128,1,128,2,5,6,37,4000.0000,34,0,'ASM(34) d 10999 4000.00,c 11000 4000.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(129,1,129,2,5,6,38,4000.0000,35,0,'ASM(35) d 10999 4000.00,c 11000 4000.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(130,1,130,2,5,6,38,15.0000,54,0,'ASM(54) d 10999 15.00,c  15.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(131,1,131,2,5,6,38,4000.0000,36,0,'ASM(36) d 10999 4000.00,c 11000 4000.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(132,1,132,2,5,6,38,3970.0000,37,0,'ASM(37) d 10999 3970.00,c 11000 3970.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(133,1,133,2,5,6,39,30.0000,37,0,'ASM(37) d 10999 30.00,c 11000 30.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(134,1,134,2,5,6,39,598.4500,44,0,'ASM(44) d 10999 598.45,c 11000 598.45','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(135,1,135,2,5,6,40,30.0000,44,0,'ASM(44) d 10999 30.00,c 11000 30.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(136,1,136,2,5,6,40,4000.0000,38,0,'ASM(38) d 10999 4000.00,c 11000 4000.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(137,1,137,2,5,6,40,175.0000,45,0,'ASM(45) d 10999 175.00,c 11000 175.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(138,1,138,2,5,6,40,15.0000,55,0,'ASM(55) d 10999 15.00,c  15.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(139,1,139,2,5,6,40,4000.0000,39,0,'ASM(39) d 10999 4000.00,c 11000 4000.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(140,1,140,2,5,6,40,115.0000,46,0,'ASM(46) d 10999 115.00,c 11000 115.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(141,1,141,2,5,6,41,45.0000,46,0,'ASM(46) d 10999 45.00,c 11000 45.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(142,1,142,2,5,6,41,4000.0000,40,0,'ASM(40) d 10999 4000.00,c 11000 4000.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(143,1,143,2,5,6,41,350.0000,47,0,'ASM(47) d 10999 350.00,c 11000 350.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(144,1,144,2,5,6,41,81.7900,48,0,'ASM(48) d 10999 81.79,c 11000 81.79','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(145,1,145,2,5,6,41,4000.0000,41,0,'ASM(41) d 10999 4000.00,c 11000 4000.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(146,1,146,2,5,6,41,350.0000,49,0,'ASM(49) d 10999 350.00,c 11000 350.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(147,1,147,2,5,6,41,15.0000,56,0,'ASM(56) d 10999 15.00,c  15.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(148,1,148,2,5,6,41,4000.0000,42,0,'ASM(42) d 10999 4000.00,c 11000 4000.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(149,1,149,2,5,6,41,275.0000,50,0,'ASM(50) d 10999 275.00,c 11000 275.00','2017-07-18 21:55:07',0,'2018-01-01 09:59:48',0),(150,1,150,2,5,0,0,-160.0000,57,0,'d 11000 -160.00, c 44009 -160.00','2017-07-20 21:54:47',0,'2018-01-01 09:59:48',0),(151,1,151,2,5,6,40,-115.0000,46,0,'ASM(46) d 10999 -115.00,ASM(46) c 11000 -115.00','2017-07-20 21:54:47',0,'2018-01-01 09:59:48',0),(152,1,152,2,5,6,41,-45.0000,46,0,'ASM(46) d 10999 -45.00,ASM(46) c 11000 -45.00','2017-07-20 21:54:47',0,'2018-01-01 09:59:48',0),(153,1,153,2,5,0,0,175.0000,58,0,'d 11000 175.00, c 44009 175.00','2017-07-20 21:55:39',0,'2018-01-01 09:59:48',0),(154,1,154,0,0,3,0,-4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-24 08:06:52',0,'2018-01-01 09:59:48',0),(155,1,155,3,4,3,42,-4150.0000,32,0,'ASM(32) d 10999 -4150.00,ASM(32) c 11000 -4150.00','2017-07-24 08:06:52',0,'2018-01-01 09:59:48',0),(156,1,156,0,0,3,0,4150.0000,0,0,'d 10105 _, c 10999 _','2017-07-24 08:06:52',0,'2018-01-01 09:59:48',0),(157,1,157,1,1,0,0,3750.0000,59,0,'d 11000 3750.00, c 40001 3750.00','2017-08-02 16:53:39',0,'2018-01-01 09:59:48',0),(158,1,158,3,4,0,0,4150.0000,60,0,'d 11000 4150.00, c 40001 4150.00','2017-08-02 16:53:39',0,'2018-01-01 09:59:48',0),(159,1,159,2,5,0,0,4000.0000,61,0,'d 11000 4000.00, c 40001 4000.00','2017-08-02 16:53:39',0,'2018-01-01 09:59:48',0),(160,1,160,2,5,0,0,0.0000,62,0,'c 11000 -15.00','2017-08-09 21:22:14',0,'2018-01-01 09:59:48',0),(161,1,161,2,5,6,38,-15.0000,54,0,'ASM(54) d 10999 -15.00','2017-08-09 21:22:14',0,'2018-01-01 09:59:48',0),(162,1,162,2,5,0,0,15.0000,63,0,'d 50003 15.00, c 11000 15.00','2017-08-09 21:22:14',0,'2018-01-01 09:59:48',0),(163,1,163,1,1,0,0,3750.0000,64,0,'d 11000 3750.00, c 40001 3750.00','2017-09-19 18:59:07',0,'2018-01-01 09:59:48',0),(164,1,164,3,4,0,0,4150.0000,65,0,'d 11000 4150.00, c 40001 4150.00','2017-09-19 18:59:07',0,'2018-01-01 09:59:48',0),(165,1,165,2,5,0,0,4000.0000,66,0,'d 11000 4000.00, c 40001 4000.00','2017-09-19 18:59:07',0,'2018-01-01 09:59:48',0);
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
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0'
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
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0'
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
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0'
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
) ENGINE=InnoDB AUTO_INCREMENT=315 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerEntry`
--

LOCK TABLES `LedgerEntry` WRITE;
/*!40000 ALTER TABLE `LedgerEntry` DISABLE KEYS */;
INSERT INTO `LedgerEntry` VALUES (1,1,1,1,5,1,1,0,'2016-03-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(2,1,1,1,15,1,1,0,'2016-03-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(3,1,2,2,5,1,1,0,'2016-04-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(4,1,2,2,15,1,1,0,'2016-04-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(5,1,3,3,5,1,1,0,'2016-05-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(6,1,3,3,15,1,1,0,'2016-05-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(7,1,4,4,5,1,1,0,'2016-06-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(8,1,4,4,15,1,1,0,'2016-06-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(9,1,5,5,5,1,1,0,'2016-07-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(10,1,5,5,15,1,1,0,'2016-07-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(11,1,6,6,5,1,1,0,'2016-08-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(12,1,6,6,15,1,1,0,'2016-08-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(13,1,7,7,5,1,1,0,'2016-09-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(14,1,7,7,15,1,1,0,'2016-09-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(15,1,8,8,5,1,1,0,'2016-10-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(16,1,8,8,15,1,1,0,'2016-10-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(17,1,9,9,5,1,1,0,'2016-11-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(18,1,9,9,15,1,1,0,'2016-11-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(19,1,10,10,5,1,1,0,'2016-12-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(20,1,10,10,15,1,1,0,'2016-12-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(21,1,11,11,5,1,1,0,'2017-01-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(22,1,11,11,15,1,1,0,'2017-01-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(23,1,12,12,5,1,1,0,'2017-02-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(24,1,12,12,15,1,1,0,'2017-02-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(25,1,13,13,5,1,1,0,'2017-03-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(26,1,13,13,15,1,1,0,'2017-03-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(27,1,14,14,5,1,1,0,'2017-04-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(28,1,14,14,15,1,1,0,'2017-04-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(29,1,15,15,5,1,1,0,'2017-05-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(30,1,15,15,15,1,1,0,'2017-05-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(31,1,16,16,5,1,1,0,'2017-06-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(32,1,16,16,15,1,1,0,'2017-06-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(33,1,17,17,5,1,1,0,'2017-07-01 00:00:00',3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(34,1,17,17,15,1,1,0,'2017-07-01 00:00:00',-3750.0000,'','2017-07-18 19:10:52',0,'2017-07-18 19:10:52',0),(35,1,18,18,5,4,3,0,'2016-07-01 00:00:00',4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(36,1,18,18,15,4,3,0,'2016-07-01 00:00:00',-4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(37,1,19,19,5,4,3,0,'2016-08-01 00:00:00',4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(38,1,19,19,15,4,3,0,'2016-08-01 00:00:00',-4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(39,1,20,20,5,4,3,0,'2016-09-01 00:00:00',4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(40,1,20,20,15,4,3,0,'2016-09-01 00:00:00',-4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(41,1,21,21,5,4,3,0,'2016-10-01 00:00:00',4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(42,1,21,21,15,4,3,0,'2016-10-01 00:00:00',-4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(43,1,22,22,5,4,3,0,'2016-11-01 00:00:00',4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(44,1,22,22,15,4,3,0,'2016-11-01 00:00:00',-4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(45,1,23,23,5,4,3,0,'2016-12-01 00:00:00',4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(46,1,23,23,15,4,3,0,'2016-12-01 00:00:00',-4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(47,1,24,24,5,4,3,0,'2017-01-01 00:00:00',4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(48,1,24,24,15,4,3,0,'2017-01-01 00:00:00',-4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(49,1,25,25,5,4,3,0,'2017-02-01 00:00:00',4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(50,1,25,25,15,4,3,0,'2017-02-01 00:00:00',-4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(51,1,26,26,5,4,3,0,'2017-03-01 00:00:00',4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(52,1,26,26,15,4,3,0,'2017-03-01 00:00:00',-4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(53,1,27,27,5,4,3,0,'2017-04-01 00:00:00',4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(54,1,27,27,15,4,3,0,'2017-04-01 00:00:00',-4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(55,1,28,28,5,4,3,0,'2017-05-01 00:00:00',4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(56,1,28,28,15,4,3,0,'2017-05-01 00:00:00',-4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(57,1,29,29,5,4,3,0,'2017-06-01 00:00:00',4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(58,1,29,29,15,4,3,0,'2017-06-01 00:00:00',-4150.0000,'','2017-07-18 19:12:53',0,'2017-07-18 19:12:53',0),(59,1,30,30,5,4,3,0,'2017-07-01 00:00:00',4150.0000,'','2017-07-18 19:12:54',0,'2017-07-18 19:12:54',0),(60,1,30,30,15,4,3,0,'2017-07-01 00:00:00',-4150.0000,'','2017-07-18 19:12:54',0,'2017-07-18 19:12:54',0),(61,1,31,31,5,5,2,0,'2016-10-01 00:00:00',4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(62,1,31,31,15,5,2,0,'2016-10-01 00:00:00',-4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(63,1,32,32,5,5,2,0,'2016-11-01 00:00:00',4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(64,1,32,32,15,5,2,0,'2016-11-01 00:00:00',-4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(65,1,33,33,5,5,2,0,'2016-12-01 00:00:00',4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(66,1,33,33,15,5,2,0,'2016-12-01 00:00:00',-4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(67,1,34,34,5,5,2,0,'2017-01-01 00:00:00',4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(68,1,34,34,15,5,2,0,'2017-01-01 00:00:00',-4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(69,1,35,35,5,5,2,0,'2017-02-01 00:00:00',4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(70,1,35,35,15,5,2,0,'2017-02-01 00:00:00',-4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(71,1,36,36,5,5,2,0,'2017-03-01 00:00:00',4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(72,1,36,36,15,5,2,0,'2017-03-01 00:00:00',-4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(73,1,37,37,5,5,2,0,'2017-04-01 00:00:00',4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(74,1,37,37,15,5,2,0,'2017-04-01 00:00:00',-4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(75,1,38,38,5,5,2,0,'2017-05-01 00:00:00',4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(76,1,38,38,15,5,2,0,'2017-05-01 00:00:00',-4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(77,1,39,39,5,5,2,0,'2017-06-01 00:00:00',4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(78,1,39,39,15,5,2,0,'2017-06-01 00:00:00',-4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(79,1,40,40,5,5,2,0,'2017-07-01 00:00:00',4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(80,1,40,40,15,5,2,0,'2017-07-01 00:00:00',-4000.0000,'','2017-07-18 19:13:43',0,'2017-07-18 19:13:43',0),(81,1,41,41,5,5,2,0,'2017-01-24 00:00:00',628.4500,'','2017-07-18 19:21:53',0,'2017-07-18 19:21:53',0),(82,1,41,41,51,5,2,0,'2017-01-24 00:00:00',-628.4500,'','2017-07-18 19:21:53',0,'2017-07-18 19:21:53',0),(83,1,42,42,5,5,2,0,'2017-02-01 00:00:00',175.0000,'','2017-07-18 19:23:02',0,'2017-07-18 19:23:02',0),(84,1,42,42,51,5,2,0,'2017-02-01 00:00:00',-175.0000,'','2017-07-18 19:23:02',0,'2017-07-18 19:23:02',0),(85,1,43,43,5,5,2,0,'2017-03-01 00:00:00',160.0000,'','2017-07-18 19:24:22',0,'2017-07-18 19:24:22',0),(86,1,43,43,51,5,2,0,'2017-03-01 00:00:00',-160.0000,'','2017-07-18 19:24:22',0,'2017-07-18 19:24:22',0),(87,1,44,44,5,5,2,0,'2017-04-01 00:00:00',350.0000,'','2017-07-18 19:25:28',0,'2017-07-18 19:25:28',0),(88,1,44,44,51,5,2,0,'2017-04-01 00:00:00',-350.0000,'','2017-07-18 19:25:28',0,'2017-07-18 19:25:28',0),(89,1,45,45,5,5,2,0,'2017-04-01 00:00:00',81.7900,'','2017-07-18 19:26:10',0,'2017-07-18 19:26:10',0),(90,1,45,45,51,5,2,0,'2017-04-01 00:00:00',-81.7900,'','2017-07-18 19:26:10',0,'2017-07-18 19:26:10',0),(91,1,46,46,5,5,2,0,'2017-05-01 00:00:00',350.0000,'','2017-07-18 19:27:02',0,'2017-07-18 19:27:02',0),(92,1,46,46,51,5,2,0,'2017-05-01 00:00:00',-350.0000,'','2017-07-18 19:27:02',0,'2017-07-18 19:27:02',0),(93,1,47,47,5,5,2,0,'2017-06-01 00:00:00',350.0000,'','2017-07-18 19:27:48',0,'2017-07-18 19:27:48',0),(94,1,47,47,51,5,2,0,'2017-06-01 00:00:00',-350.0000,'','2017-07-18 19:27:48',0,'2017-07-18 19:27:48',0),(95,1,48,48,5,1,1,0,'2014-03-01 00:00:00',7000.0000,'','2017-07-18 19:31:09',0,'2017-07-18 19:31:09',0),(96,1,48,48,10,1,1,0,'2014-03-01 00:00:00',-7000.0000,'','2017-07-18 19:31:09',0,'2017-07-18 19:31:09',0),(97,1,49,49,5,4,3,0,'2016-07-01 00:00:00',8300.0000,'','2017-07-18 19:32:01',0,'2017-07-18 19:32:01',0),(98,1,49,49,10,4,3,0,'2016-07-01 00:00:00',-8300.0000,'','2017-07-18 19:32:01',0,'2017-07-18 19:32:01',0),(99,1,50,50,5,5,2,0,'2017-07-01 00:00:00',350.0000,'','2017-07-18 19:32:32',0,'2017-07-18 19:32:32',0),(100,1,50,50,51,5,2,0,'2017-07-01 00:00:00',-350.0000,'','2017-07-18 19:32:32',0,'2017-07-18 19:32:32',0),(101,1,51,51,4,0,0,1,'2014-03-01 00:00:00',7000.0000,'','2017-07-18 19:35:44',0,'2017-07-18 19:35:44',0),(102,1,51,51,72,0,0,1,'2014-03-01 00:00:00',-7000.0000,'','2017-07-18 19:35:44',0,'2017-07-18 19:35:44',0),(103,1,52,52,3,0,0,1,'2016-03-01 00:00:00',3750.0000,'','2017-07-18 19:36:47',0,'2017-07-18 19:36:47',0),(104,1,52,52,72,0,0,1,'2016-03-01 00:00:00',-3750.0000,'','2017-07-18 19:36:47',0,'2017-07-18 19:36:47',0),(105,1,53,53,3,0,0,0,'2016-04-01 00:00:00',3750.0000,'','2017-07-18 19:37:04',0,'2017-07-18 19:37:04',0),(106,1,53,53,72,0,0,0,'2016-04-01 00:00:00',-3750.0000,'','2017-07-18 19:37:04',0,'2017-07-18 19:37:04',0),(107,1,54,54,3,0,0,1,'2017-05-01 00:00:00',3750.0000,'','2017-07-18 19:37:26',0,'2017-07-18 19:37:26',0),(108,1,54,54,72,0,0,1,'2017-05-01 00:00:00',-3750.0000,'','2017-07-18 19:37:26',0,'2017-07-18 19:37:26',0),(109,1,55,55,3,0,0,1,'2017-05-01 00:00:00',-3750.0000,'','2017-07-18 19:38:20',0,'2017-07-18 19:38:20',0),(110,1,55,55,72,0,0,1,'2017-05-01 00:00:00',3750.0000,'','2017-07-18 19:38:20',0,'2017-07-18 19:38:20',0),(111,1,56,56,3,0,0,1,'2016-05-01 00:00:00',3750.0000,'','2017-07-18 19:38:20',0,'2017-07-18 19:38:20',0),(112,1,56,56,72,0,0,1,'2016-05-01 00:00:00',-3750.0000,'','2017-07-18 19:38:20',0,'2017-07-18 19:38:20',0),(113,1,57,57,3,0,0,1,'2016-05-01 00:00:00',3750.0000,'','2017-07-18 19:38:55',0,'2017-07-18 19:38:55',0),(114,1,57,57,72,0,0,1,'2016-05-01 00:00:00',-3750.0000,'','2017-07-18 19:38:55',0,'2017-07-18 19:38:55',0),(115,1,58,58,3,0,0,1,'2016-06-01 00:00:00',3750.0000,'','2017-07-18 19:39:17',0,'2017-07-18 19:39:17',0),(116,1,58,58,72,0,0,1,'2016-06-01 00:00:00',-3750.0000,'','2017-07-18 19:39:17',0,'2017-07-18 19:39:17',0),(117,1,59,59,3,0,0,1,'2016-07-01 00:00:00',3750.0000,'','2017-07-18 19:39:41',0,'2017-07-18 19:39:41',0),(118,1,59,59,72,0,0,1,'2016-07-01 00:00:00',-3750.0000,'','2017-07-18 19:39:41',0,'2017-07-18 19:39:41',0),(119,1,60,60,3,0,0,1,'2016-08-01 00:00:00',3750.0000,'','2017-07-18 19:40:00',0,'2017-07-18 19:40:00',0),(120,1,60,60,72,0,0,1,'2016-08-01 00:00:00',-3750.0000,'','2017-07-18 19:40:00',0,'2017-07-18 19:40:00',0),(121,1,61,61,3,0,0,1,'2016-09-01 00:00:00',3750.0000,'','2017-07-18 19:40:17',0,'2017-07-18 19:40:17',0),(122,1,61,61,72,0,0,1,'2016-09-01 00:00:00',-3750.0000,'','2017-07-18 19:40:17',0,'2017-07-18 19:40:17',0),(123,1,62,62,3,0,0,1,'2016-10-01 00:00:00',3750.0000,'','2017-07-18 19:40:33',0,'2017-07-18 19:40:33',0),(124,1,62,62,72,0,0,1,'2016-10-01 00:00:00',-3750.0000,'','2017-07-18 19:40:33',0,'2017-07-18 19:40:33',0),(125,1,63,63,3,0,0,1,'2016-11-01 00:00:00',3750.0000,'','2017-07-18 19:40:50',0,'2017-07-18 19:40:50',0),(126,1,63,63,72,0,0,1,'2016-11-01 00:00:00',-3750.0000,'','2017-07-18 19:40:50',0,'2017-07-18 19:40:50',0),(127,1,64,64,3,0,0,0,'2016-12-01 00:00:00',3750.0000,'','2017-07-18 19:41:03',0,'2017-07-18 19:41:03',0),(128,1,64,64,72,0,0,0,'2016-12-01 00:00:00',-3750.0000,'','2017-07-18 19:41:03',0,'2017-07-18 19:41:03',0),(129,1,65,65,3,0,0,1,'2017-01-01 00:00:00',3750.0000,'','2017-07-18 19:41:23',0,'2017-07-18 19:41:23',0),(130,1,65,65,72,0,0,1,'2017-01-01 00:00:00',-3750.0000,'','2017-07-18 19:41:23',0,'2017-07-18 19:41:23',0),(131,1,66,66,3,0,0,1,'2017-02-01 00:00:00',3750.0000,'','2017-07-18 19:41:38',0,'2017-07-18 19:41:38',0),(132,1,66,66,72,0,0,1,'2017-02-01 00:00:00',-3750.0000,'','2017-07-18 19:41:38',0,'2017-07-18 19:41:38',0),(133,1,67,67,3,0,0,1,'2017-03-01 00:00:00',3750.0000,'','2017-07-18 19:42:11',0,'2017-07-18 19:42:11',0),(134,1,67,67,72,0,0,1,'2017-03-01 00:00:00',-3750.0000,'','2017-07-18 19:42:11',0,'2017-07-18 19:42:11',0),(135,1,68,68,3,0,0,1,'2017-04-01 00:00:00',3750.0000,'','2017-07-18 19:42:27',0,'2017-07-18 19:42:27',0),(136,1,68,68,72,0,0,1,'2017-04-01 00:00:00',-3750.0000,'','2017-07-18 19:42:27',0,'2017-07-18 19:42:27',0),(137,1,69,69,3,0,0,1,'2017-05-01 00:00:00',3750.0000,'','2017-07-18 19:42:43',0,'2017-07-18 19:42:43',0),(138,1,69,69,72,0,0,1,'2017-05-01 00:00:00',-3750.0000,'','2017-07-18 19:42:43',0,'2017-07-18 19:42:43',0),(139,1,70,70,3,0,0,1,'2017-06-01 00:00:00',3750.0000,'','2017-07-18 19:43:06',0,'2017-07-18 19:43:06',0),(140,1,70,70,72,0,0,1,'2017-06-01 00:00:00',-3750.0000,'','2017-07-18 19:43:06',0,'2017-07-18 19:43:06',0),(141,1,71,71,3,0,0,1,'2017-07-01 00:00:00',3750.0000,'','2017-07-18 19:43:23',0,'2017-07-18 19:43:23',0),(142,1,71,71,72,0,0,1,'2017-07-01 00:00:00',-3750.0000,'','2017-07-18 19:43:23',0,'2017-07-18 19:43:23',0),(143,1,72,72,3,0,0,1,'2016-05-01 00:00:00',-3750.0000,'','2017-07-18 19:44:38',0,'2017-07-18 19:44:38',0),(144,1,72,72,72,0,0,1,'2016-05-01 00:00:00',3750.0000,'','2017-07-18 19:44:38',0,'2017-07-18 19:44:38',0),(145,1,73,73,4,0,0,3,'2016-07-01 00:00:00',8300.0000,'','2017-07-18 19:45:29',0,'2017-07-18 19:45:29',0),(146,1,73,73,72,0,0,3,'2016-07-01 00:00:00',-8300.0000,'','2017-07-18 19:45:29',0,'2017-07-18 19:45:29',0),(147,1,74,74,3,0,0,3,'2016-07-01 00:00:00',4150.0000,'','2017-07-18 19:45:55',0,'2017-07-18 19:45:55',0),(148,1,74,74,72,0,0,3,'2016-07-01 00:00:00',-4150.0000,'','2017-07-18 19:45:55',0,'2017-07-18 19:45:55',0),(149,1,75,75,3,0,0,3,'2016-08-01 00:00:00',4150.0000,'','2017-07-18 19:46:13',0,'2017-07-18 19:46:13',0),(150,1,75,75,72,0,0,3,'2016-08-01 00:00:00',-4150.0000,'','2017-07-18 19:46:13',0,'2017-07-18 19:46:13',0),(151,1,76,76,3,0,0,3,'2016-09-01 00:00:00',4150.0000,'','2017-07-18 19:46:31',0,'2017-07-18 19:46:31',0),(152,1,76,76,72,0,0,3,'2016-09-01 00:00:00',-4150.0000,'','2017-07-18 19:46:31',0,'2017-07-18 19:46:31',0),(153,1,77,77,3,0,0,3,'2016-10-01 00:00:00',4150.0000,'','2017-07-18 19:46:50',0,'2017-07-18 19:46:50',0),(154,1,77,77,72,0,0,3,'2016-10-01 00:00:00',-4150.0000,'','2017-07-18 19:46:50',0,'2017-07-18 19:46:50',0),(155,1,78,78,3,0,0,3,'2016-11-01 00:00:00',4150.0000,'','2017-07-18 19:47:09',0,'2017-07-18 19:47:09',0),(156,1,78,78,72,0,0,3,'2016-11-01 00:00:00',-4150.0000,'','2017-07-18 19:47:09',0,'2017-07-18 19:47:09',0),(157,1,79,79,3,0,0,3,'2016-12-01 00:00:00',4150.0000,'','2017-07-18 19:47:28',0,'2017-07-18 19:47:28',0),(158,1,79,79,72,0,0,3,'2016-12-01 00:00:00',-4150.0000,'','2017-07-18 19:47:28',0,'2017-07-18 19:47:28',0),(159,1,80,80,3,0,0,3,'2017-01-01 00:00:00',4150.0000,'','2017-07-18 19:47:59',0,'2017-07-18 19:47:59',0),(160,1,80,80,72,0,0,3,'2017-01-01 00:00:00',-4150.0000,'','2017-07-18 19:47:59',0,'2017-07-18 19:47:59',0),(161,1,81,81,3,0,0,3,'2017-02-01 00:00:00',4150.0000,'','2017-07-18 19:48:18',0,'2017-07-18 19:48:18',0),(162,1,81,81,72,0,0,3,'2017-02-01 00:00:00',-4150.0000,'','2017-07-18 19:48:18',0,'2017-07-18 19:48:18',0),(163,1,82,82,3,0,0,3,'2017-03-01 00:00:00',4150.0000,'','2017-07-18 19:48:38',0,'2017-07-18 19:48:38',0),(164,1,82,82,72,0,0,3,'2017-03-01 00:00:00',-4150.0000,'','2017-07-18 19:48:38',0,'2017-07-18 19:48:38',0),(165,1,83,83,3,0,0,3,'2017-04-01 00:00:00',4150.0000,'','2017-07-18 19:48:56',0,'2017-07-18 19:48:56',0),(166,1,83,83,72,0,0,3,'2017-04-01 00:00:00',-4150.0000,'','2017-07-18 19:48:56',0,'2017-07-18 19:48:56',0),(167,1,84,84,3,0,0,3,'2017-05-01 00:00:00',4150.0000,'','2017-07-18 19:49:12',0,'2017-07-18 19:49:12',0),(168,1,84,84,72,0,0,3,'2017-05-01 00:00:00',-4150.0000,'','2017-07-18 19:49:12',0,'2017-07-18 19:49:12',0),(169,1,85,85,3,0,0,3,'2017-06-01 00:00:00',4150.0000,'','2017-07-18 19:49:28',0,'2017-07-18 19:49:28',0),(170,1,85,85,72,0,0,3,'2017-06-01 00:00:00',-4150.0000,'','2017-07-18 19:49:28',0,'2017-07-18 19:49:28',0),(171,1,86,86,3,0,0,3,'2017-07-01 00:00:00',4150.0000,'','2017-07-18 19:49:51',0,'2017-07-18 19:49:51',0),(172,1,86,86,72,0,0,3,'2017-07-01 00:00:00',-4150.0000,'','2017-07-18 19:49:51',0,'2017-07-18 19:49:51',0),(173,1,87,87,3,0,0,6,'2016-10-03 00:00:00',4000.0000,'','2017-07-18 19:51:26',0,'2017-07-18 19:51:26',0),(174,1,87,87,72,0,0,6,'2016-10-03 00:00:00',-4000.0000,'','2017-07-18 19:51:26',0,'2017-07-18 19:51:26',0),(175,1,88,88,3,0,0,6,'2016-11-11 00:00:00',11985.0000,'','2017-07-18 19:52:33',0,'2017-07-18 19:52:33',0),(176,1,88,88,72,0,0,6,'2016-11-11 00:00:00',-11985.0000,'','2017-07-18 19:52:33',0,'2017-07-18 19:52:33',0),(177,1,89,89,5,5,2,0,'2016-11-11 00:00:00',-15.0000,'','2017-07-18 19:53:22',0,'2017-07-18 19:53:22',0),(178,1,90,90,3,0,0,6,'2017-02-03 00:00:00',628.4500,'','2017-07-18 19:54:30',0,'2017-07-18 19:54:30',0),(179,1,90,90,72,0,0,6,'2017-02-03 00:00:00',-628.4500,'','2017-07-18 19:54:30',0,'2017-07-18 19:54:30',0),(180,1,91,91,3,0,0,6,'2017-02-13 00:00:00',8335.0000,'','2017-07-18 19:56:46',0,'2017-07-18 19:56:46',0),(181,1,91,91,72,0,0,6,'2017-02-13 00:00:00',-8335.0000,'','2017-07-18 19:56:46',0,'2017-07-18 19:56:46',0),(182,1,92,92,5,5,2,0,'2017-02-13 00:00:00',-15.0000,'','2017-07-18 19:57:33',0,'2017-07-18 19:57:33',0),(183,1,93,93,3,0,0,6,'2017-05-12 00:00:00',13116.7900,'','2017-07-18 19:59:41',0,'2017-07-18 19:59:41',0),(184,1,93,93,72,0,0,6,'2017-05-12 00:00:00',-13116.7900,'','2017-07-18 19:59:41',0,'2017-07-18 19:59:41',0),(185,1,94,94,5,5,2,0,'2017-05-12 00:00:00',-15.0000,'','2017-07-18 20:00:26',0,'2017-07-18 20:00:26',0),(186,1,95,95,72,1,1,1,'2014-03-01 00:00:00',7000.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(187,1,95,95,5,1,1,1,'2014-03-01 00:00:00',-7000.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(188,1,96,96,72,1,1,1,'2016-03-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(189,1,96,96,5,1,1,1,'2016-03-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(190,1,97,97,72,1,1,1,'2016-04-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(191,1,97,97,5,1,1,1,'2016-04-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(192,1,98,98,72,1,1,1,'2016-05-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(193,1,98,98,5,1,1,1,'2016-05-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(194,1,99,99,72,1,1,1,'2016-06-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(195,1,99,99,5,1,1,1,'2016-06-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(196,1,100,100,72,1,1,1,'2016-07-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(197,1,100,100,5,1,1,1,'2016-07-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(198,1,101,101,72,1,1,1,'2016-08-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(199,1,101,101,5,1,1,1,'2016-08-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(200,1,102,102,72,1,1,1,'2016-09-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(201,1,102,102,5,1,1,1,'2016-09-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(202,1,103,103,72,1,1,1,'2016-10-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(203,1,103,103,5,1,1,1,'2016-10-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(204,1,104,104,72,1,1,1,'2016-11-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(205,1,104,104,5,1,1,1,'2016-11-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(206,1,105,105,72,1,1,1,'2016-12-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(207,1,105,105,5,1,1,1,'2016-12-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(208,1,106,106,72,1,1,1,'2017-01-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(209,1,106,106,5,1,1,1,'2017-01-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(210,1,107,107,72,1,1,1,'2017-02-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(211,1,107,107,5,1,1,1,'2017-02-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(212,1,108,108,72,1,1,1,'2017-03-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(213,1,108,108,5,1,1,1,'2017-03-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(214,1,109,109,72,1,1,1,'2017-04-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(215,1,109,109,5,1,1,1,'2017-04-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(216,1,110,110,72,1,1,1,'2017-05-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(217,1,110,110,5,1,1,1,'2017-05-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(218,1,111,111,72,1,1,1,'2017-06-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(219,1,111,111,5,1,1,1,'2017-06-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(220,1,112,112,72,1,1,1,'2017-07-01 00:00:00',3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(221,1,112,112,5,1,1,1,'2017-07-01 00:00:00',-3750.0000,'','2017-07-18 21:47:42',0,'2017-07-18 21:47:42',0),(222,1,113,113,72,4,3,3,'2016-07-01 00:00:00',4150.0000,'','2017-07-18 21:49:52',0,'2017-07-18 21:49:52',0),(223,1,113,113,5,4,3,3,'2016-07-01 00:00:00',-4150.0000,'','2017-07-18 21:49:52',0,'2017-07-18 21:49:52',0),(224,1,114,114,72,4,3,3,'2016-07-01 00:00:00',4150.0000,'','2017-07-18 21:49:52',0,'2017-07-18 21:49:52',0),(225,1,114,114,5,4,3,3,'2016-07-01 00:00:00',-4150.0000,'','2017-07-18 21:49:52',0,'2017-07-18 21:49:52',0),(226,1,115,115,72,4,3,3,'2016-07-01 00:00:00',4150.0000,'','2017-07-18 21:49:52',0,'2017-07-18 21:49:52',0),(227,1,115,115,5,4,3,3,'2016-07-01 00:00:00',-4150.0000,'','2017-07-18 21:49:52',0,'2017-07-18 21:49:52',0),(228,1,116,116,72,4,3,3,'2016-08-01 00:00:00',4150.0000,'','2017-07-18 21:49:52',0,'2017-07-18 21:49:52',0),(229,1,116,116,5,4,3,3,'2016-08-01 00:00:00',-4150.0000,'','2017-07-18 21:49:52',0,'2017-07-18 21:49:52',0),(230,1,117,117,72,4,3,3,'2016-09-01 00:00:00',4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(231,1,117,117,5,4,3,3,'2016-09-01 00:00:00',-4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(232,1,118,118,72,4,3,3,'2016-10-01 00:00:00',4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(233,1,118,118,5,4,3,3,'2016-10-01 00:00:00',-4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(234,1,119,119,72,4,3,3,'2016-11-01 00:00:00',4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(235,1,119,119,5,4,3,3,'2016-11-01 00:00:00',-4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(236,1,120,120,72,4,3,3,'2016-12-01 00:00:00',4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(237,1,120,120,5,4,3,3,'2016-12-01 00:00:00',-4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(238,1,121,121,72,4,3,3,'2017-01-01 00:00:00',4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(239,1,121,121,5,4,3,3,'2017-01-01 00:00:00',-4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(240,1,122,122,72,4,3,3,'2017-02-01 00:00:00',4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(241,1,122,122,5,4,3,3,'2017-02-01 00:00:00',-4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(242,1,123,123,72,4,3,3,'2017-03-01 00:00:00',4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(243,1,123,123,5,4,3,3,'2017-03-01 00:00:00',-4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(244,1,124,124,72,4,3,3,'2017-04-01 00:00:00',4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(245,1,124,124,5,4,3,3,'2017-04-01 00:00:00',-4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(246,1,125,125,72,4,3,3,'2017-05-01 00:00:00',4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(247,1,125,125,5,4,3,3,'2017-05-01 00:00:00',-4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(248,1,126,126,72,4,3,3,'2017-06-01 00:00:00',4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(249,1,126,126,5,4,3,3,'2017-06-01 00:00:00',-4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(250,1,127,127,72,4,3,3,'2017-07-01 00:00:00',4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(251,1,127,127,5,4,3,3,'2017-07-01 00:00:00',-4150.0000,'','2017-07-18 21:49:53',0,'2017-07-18 21:49:53',0),(252,1,128,128,72,5,2,6,'2016-10-03 00:00:00',4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(253,1,128,128,5,5,2,6,'2016-10-03 00:00:00',-4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(254,1,129,129,72,5,2,6,'2016-11-11 00:00:00',4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(255,1,129,129,5,5,2,6,'2016-11-11 00:00:00',-4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(256,1,130,130,72,5,2,6,'2016-11-11 00:00:00',15.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(257,1,130,130,0,5,2,6,'2016-11-11 00:00:00',-15.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(258,1,131,131,72,5,2,6,'2016-11-11 00:00:00',4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(259,1,131,131,5,5,2,6,'2016-11-11 00:00:00',-4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(260,1,132,132,72,5,2,6,'2016-11-11 00:00:00',3970.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(261,1,132,132,5,5,2,6,'2016-11-11 00:00:00',-3970.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(262,1,133,133,72,5,2,6,'2017-02-03 00:00:00',30.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(263,1,133,133,5,5,2,6,'2017-02-03 00:00:00',-30.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(264,1,134,134,72,5,2,6,'2017-02-03 00:00:00',598.4500,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(265,1,134,134,5,5,2,6,'2017-02-03 00:00:00',-598.4500,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(266,1,135,135,72,5,2,6,'2017-02-13 00:00:00',30.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(267,1,135,135,5,5,2,6,'2017-02-13 00:00:00',-30.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(268,1,136,136,72,5,2,6,'2017-02-13 00:00:00',4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(269,1,136,136,5,5,2,6,'2017-02-13 00:00:00',-4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(270,1,137,137,72,5,2,6,'2017-02-13 00:00:00',175.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(271,1,137,137,5,5,2,6,'2017-02-13 00:00:00',-175.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(272,1,138,138,72,5,2,6,'2017-02-13 00:00:00',15.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(273,1,138,138,0,5,2,6,'2017-02-13 00:00:00',-15.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(274,1,139,139,72,5,2,6,'2017-02-13 00:00:00',4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(275,1,139,139,5,5,2,6,'2017-02-13 00:00:00',-4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(276,1,140,140,72,5,2,6,'2017-02-13 00:00:00',115.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(277,1,140,140,5,5,2,6,'2017-02-13 00:00:00',-115.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(278,1,141,141,72,5,2,6,'2017-05-12 00:00:00',45.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(279,1,141,141,5,5,2,6,'2017-05-12 00:00:00',-45.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(280,1,142,142,72,5,2,6,'2017-05-12 00:00:00',4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(281,1,142,142,5,5,2,6,'2017-05-12 00:00:00',-4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(282,1,143,143,72,5,2,6,'2017-05-12 00:00:00',350.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(283,1,143,143,5,5,2,6,'2017-05-12 00:00:00',-350.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(284,1,144,144,72,5,2,6,'2017-05-12 00:00:00',81.7900,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(285,1,144,144,5,5,2,6,'2017-05-12 00:00:00',-81.7900,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(286,1,145,145,72,5,2,6,'2017-05-12 00:00:00',4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(287,1,145,145,5,5,2,6,'2017-05-12 00:00:00',-4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(288,1,146,146,72,5,2,6,'2017-05-12 00:00:00',350.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(289,1,146,146,5,5,2,6,'2017-05-12 00:00:00',-350.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(290,1,147,147,72,5,2,6,'2017-05-12 00:00:00',15.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(291,1,147,147,0,5,2,6,'2017-05-12 00:00:00',-15.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(292,1,148,148,72,5,2,6,'2017-05-12 00:00:00',4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(293,1,148,148,5,5,2,6,'2017-05-12 00:00:00',-4000.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(294,1,149,149,72,5,2,6,'2017-05-12 00:00:00',275.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(295,1,149,149,5,5,2,6,'2017-05-12 00:00:00',-275.0000,'','2017-07-18 21:55:07',0,'2017-07-18 21:55:07',0),(296,1,150,150,5,5,2,0,'2017-03-01 00:00:00',-160.0000,'','2017-07-20 21:54:47',0,'2017-07-20 21:54:47',0),(297,1,150,150,51,5,2,0,'2017-03-01 00:00:00',160.0000,'','2017-07-20 21:54:47',0,'2017-07-20 21:54:47',0),(298,1,151,151,72,5,2,6,'2017-02-13 00:00:00',-115.0000,'','2017-07-20 21:54:47',0,'2017-07-20 21:54:47',0),(299,1,151,151,5,5,2,6,'2017-02-13 00:00:00',115.0000,'','2017-07-20 21:54:47',0,'2017-07-20 21:54:47',0),(300,1,152,152,72,5,2,6,'2017-05-12 00:00:00',-45.0000,'','2017-07-20 21:54:47',0,'2017-07-20 21:54:47',0),(301,1,152,152,5,5,2,6,'2017-05-12 00:00:00',45.0000,'','2017-07-20 21:54:47',0,'2017-07-20 21:54:47',0),(302,1,153,153,5,5,2,0,'2017-03-01 00:00:00',175.0000,'','2017-07-20 21:55:39',0,'2017-07-20 21:55:39',0),(303,1,153,153,51,5,2,0,'2017-03-01 00:00:00',-175.0000,'','2017-07-20 21:55:39',0,'2017-07-20 21:55:39',0),(304,1,154,154,3,0,0,3,'2017-07-01 00:00:00',-4150.0000,'','2017-07-24 08:06:52',0,'2017-07-24 08:06:52',0),(305,1,154,154,72,0,0,3,'2017-07-01 00:00:00',4150.0000,'','2017-07-24 08:06:52',0,'2017-07-24 08:06:52',0),(306,1,155,155,72,4,3,3,'2017-07-01 00:00:00',-4150.0000,'','2017-07-24 08:06:52',0,'2017-07-24 08:06:52',0),(307,1,155,155,5,4,3,3,'2017-07-01 00:00:00',4150.0000,'','2017-07-24 08:06:52',0,'2017-07-24 08:06:52',0),(308,1,156,156,4,0,0,3,'2017-07-01 00:00:00',4150.0000,'','2017-07-24 08:06:52',0,'2017-07-24 08:06:52',0),(309,1,156,156,72,0,0,3,'2017-07-01 00:00:00',-4150.0000,'','2017-07-24 08:06:52',0,'2017-07-24 08:06:52',0),(310,1,160,160,5,5,2,0,'2016-11-11 00:00:00',15.0000,'','2017-08-09 21:22:14',0,'2017-08-09 21:22:14',0),(311,1,161,161,72,5,2,6,'2016-11-11 00:00:00',-15.0000,'','2017-08-09 21:22:14',0,'2017-08-09 21:22:14',0),(312,1,161,161,0,5,2,6,'2016-11-11 00:00:00',15.0000,'','2017-08-09 21:22:14',0,'2017-08-09 21:22:14',0),(313,1,162,162,73,5,2,0,'2016-11-11 00:00:00',15.0000,'','2017-08-09 21:22:14',0,'2017-08-09 21:22:14',0),(314,1,162,162,5,5,2,0,'2016-11-11 00:00:00',-15.0000,'','2017-08-09 21:22:14',0,'2017-08-09 21:22:14',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=678 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerMarker`
--

LOCK TABLES `LedgerMarker` WRITE;
/*!40000 ALTER TABLE `LedgerMarker` DISABLE KEYS */;
INSERT INTO `LedgerMarker` VALUES (1,1,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-02 18:06:00',0,'2017-06-14 18:26:49',0),(2,2,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-02 18:06:00',0,'2017-06-14 18:26:49',0),(3,3,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-02 18:06:00',0,'2017-06-14 18:26:49',0),(4,4,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-02 18:06:00',0,'2017-06-14 18:26:49',0),(5,5,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-02 18:06:00',0,'2017-06-14 18:26:49',0),(6,6,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-02 18:06:00',0,'2017-06-14 18:26:49',0),(7,7,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-02 18:06:00',0,'2017-06-14 18:26:49',0),(8,8,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-02 18:06:00',0,'2017-06-14 18:26:49',0),(9,9,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(10,10,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(11,11,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(12,12,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(13,13,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(14,14,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(15,15,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(16,16,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(17,17,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(18,18,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(19,19,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(20,20,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(21,21,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(22,22,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(23,23,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(24,24,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(25,25,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(26,26,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(27,27,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(28,28,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(29,29,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(30,30,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(31,31,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(32,32,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(33,33,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(34,34,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(633,35,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:43:45',0,'2017-07-08 01:43:45',0),(634,36,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(635,37,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(636,38,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(637,39,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(638,40,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(639,41,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(640,42,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(641,43,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(642,44,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(643,45,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(644,46,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(645,47,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(646,48,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(647,49,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(648,50,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(649,51,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(650,52,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(651,53,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(652,54,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(653,55,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(654,56,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(655,57,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(656,58,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(657,59,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(658,60,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(659,61,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(660,62,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(661,63,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(662,64,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(663,65,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:27',0,'2017-07-08 01:46:27',0),(664,66,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:28',0,'2017-07-08 01:46:28',0),(665,67,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:28',0,'2017-07-08 01:46:28',0),(666,68,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:28',0,'2017-07-08 01:46:28',0),(667,69,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:28',0,'2017-07-08 01:46:28',0),(668,70,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:28',0,'2017-07-08 01:46:28',0),(669,71,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:28',0,'2017-07-08 01:46:28',0),(670,72,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-07-08 01:46:28',0,'2017-07-08 01:46:28',0),(671,0,0,1,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-02 18:06:00',0,'2017-07-27 19:55:19',0),(672,0,0,4,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-02 18:06:00',0,'2017-07-27 19:55:19',0),(673,0,0,5,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-02 18:06:00',0,'2017-07-27 19:55:19',0),(674,73,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-08-09 04:10:19',0,'2017-08-09 04:10:19',0),(675,0,1,1,1,0,'2014-03-01 00:00:00',0.0000,3,'2017-09-19 18:59:03',0,'2017-09-19 18:59:03',0),(676,0,1,5,2,0,'2016-10-01 00:00:00',0.0000,3,'2017-09-19 18:59:06',0,'2017-09-19 18:59:06',0),(677,0,1,4,3,0,'2016-07-01 00:00:00',0.0000,3,'2017-09-19 18:59:08',0,'2017-09-19 18:59:08',0);
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
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0'
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
-- Table structure for table `MRHistory`
--

DROP TABLE IF EXISTS `MRHistory`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `MRHistory` (
  `MRHID` bigint(20) NOT NULL AUTO_INCREMENT,
  `MRStatus` smallint(6) NOT NULL DEFAULT '0',
  `DtMRStart` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `DtMRStop` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`MRHID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `MRHistory`
--

LOCK TABLES `MRHistory` WRITE;
/*!40000 ALTER TABLE `MRHistory` DISABLE KEYS */;
/*!40000 ALTER TABLE `MRHistory` ENABLE KEYS */;
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
  `Active` tinyint(1) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PaymentType`
--

LOCK TABLES `PaymentType` WRITE;
/*!40000 ALTER TABLE `PaymentType` DISABLE KEYS */;
INSERT INTO `PaymentType` VALUES (1,1,'Check','Check','2017-07-05 02:15:10',0,'2017-07-05 02:15:24',0),(2,1,'ACH','ACH Transfer','2017-07-05 02:15:41',0,'2017-07-05 02:15:41',0),(3,1,'Wire','Wire Transfer','2017-07-05 02:15:52',0,'2017-07-05 02:15:52',0),(4,1,'Money Order','Money Order','2017-07-05 02:16:05',0,'2017-07-05 02:16:05',0);
/*!40000 ALTER TABLE `PaymentType` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Payor`
--

DROP TABLE IF EXISTS `Payor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Payor` (
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `TaxpayorID` varchar(25) NOT NULL DEFAULT '',
  `CreditLimit` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `ThirdPartySource` bigint(20) NOT NULL DEFAULT '0',
  `EligibleFuturePayor` tinyint(1) NOT NULL DEFAULT '1',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `SSN` char(128) NOT NULL DEFAULT '',
  `DriversLicense` char(128) NOT NULL DEFAULT '',
  `GrossIncome` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `TCID` bigint(20) NOT NULL,
  PRIMARY KEY (`TCID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Payor`
--

LOCK TABLES `Payor` WRITE;
/*!40000 ALTER TABLE `Payor` DISABLE KEYS */;
INSERT INTO `Payor` VALUES (1,'',0.0000,0,1,0,'','',0.0000,'2017-06-13 19:39:18',0,'2017-06-14 18:26:50',0,1),(1,'',0.0000,0,1,0,'','',0.0000,'2017-06-13 19:40:59',0,'2017-06-14 18:26:50',0,2),(1,'',0.0000,0,1,0,'','',0.0000,'2017-06-15 16:35:44',0,'2017-06-15 16:35:44',0,3),(1,'',0.0000,0,1,0,'','',0.0000,'2017-06-15 16:36:27',0,'2017-06-15 16:36:27',0,4),(1,'',0.0000,0,1,0,'','',0.0000,'2017-06-15 16:38:32',0,'2017-06-15 16:38:32',0,5),(1,'',0.0000,0,1,0,'','',0.0000,'2017-06-15 16:50:13',0,'2017-06-15 16:50:13',0,6);
/*!40000 ALTER TABLE `Payor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Prospect`
--

DROP TABLE IF EXISTS `Prospect`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Prospect` (
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `CompanyAddress` varchar(100) NOT NULL DEFAULT '',
  `CompanyCity` varchar(100) NOT NULL DEFAULT '',
  `CompanyState` varchar(100) NOT NULL DEFAULT '',
  `CompanyPostalCode` varchar(100) NOT NULL DEFAULT '',
  `CompanyEmail` varchar(100) NOT NULL DEFAULT '',
  `CompanyPhone` varchar(100) NOT NULL DEFAULT '',
  `Occupation` varchar(100) NOT NULL DEFAULT '',
  `DesiredUsageStartDate` date NOT NULL DEFAULT '1970-01-01',
  `RentableTypePreference` bigint(20) NOT NULL DEFAULT '0',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `EvictedDes` varchar(2048) NOT NULL DEFAULT '',
  `ConvictedDes` varchar(2048) NOT NULL DEFAULT '',
  `BankruptcyDes` varchar(2048) NOT NULL DEFAULT '',
  `Approver` bigint(20) NOT NULL DEFAULT '0',
  `DeclineReasonSLSID` bigint(20) NOT NULL DEFAULT '0',
  `OtherPreferences` varchar(1024) NOT NULL DEFAULT '',
  `FollowUpDate` date NOT NULL DEFAULT '1970-01-01',
  `CSAgent` bigint(20) NOT NULL DEFAULT '0',
  `OutcomeSLSID` bigint(20) NOT NULL DEFAULT '0',
  `CurrentAddress` varchar(200) NOT NULL DEFAULT '',
  `CurrentLandLordName` varchar(100) NOT NULL DEFAULT '',
  `CurrentLandLordPhoneNo` varchar(20) NOT NULL DEFAULT '',
  `CurrentReasonForMoving` bigint(20) NOT NULL DEFAULT '0',
  `CurrentLengthOfResidency` varchar(100) NOT NULL DEFAULT '',
  `PriorAddress` varchar(200) NOT NULL DEFAULT '',
  `PriorLandLordName` varchar(100) NOT NULL DEFAULT '',
  `PriorLandLordPhoneNo` varchar(20) NOT NULL DEFAULT '',
  `PriorReasonForMoving` bigint(20) NOT NULL DEFAULT '0',
  `PriorLengthOfResidency` varchar(100) NOT NULL DEFAULT '',
  `CommissionableThirdParty` text NOT NULL,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `TCID` bigint(20) NOT NULL,
  PRIMARY KEY (`TCID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Prospect`
--

LOCK TABLES `Prospect` WRITE;
/*!40000 ALTER TABLE `Prospect` DISABLE KEYS */;
INSERT INTO `Prospect` VALUES (1,'','','','','','','','1900-01-01',0,0,'','','',0,0,'','1900-01-01',0,0,'','','',0,'','','','',0,'','','2017-06-13 19:39:18',0,'2017-06-14 18:26:50',0,1),(1,'','','','','','','','1900-01-01',0,0,'','','',0,0,'','1900-01-01',0,0,'','','',0,'','','','',0,'','','2017-06-13 19:40:59',0,'2017-06-14 18:26:50',0,2),(1,'','','','','','','','1900-01-01',0,0,'','','',0,0,'','1900-01-01',0,0,'','','',0,'','','','',0,'','','2017-06-15 16:35:44',0,'2017-06-15 16:35:44',0,3),(1,'','','','','','','','1900-01-01',0,0,'','','',0,0,'','1900-01-01',0,0,'','','',0,'','','','',0,'','','2017-06-15 16:36:27',0,'2017-06-15 16:36:27',0,4),(1,'','','','','','','','1900-01-01',0,0,'','','',0,0,'','1900-01-01',0,0,'','','',0,'','','','',0,'','','2017-06-15 16:38:32',0,'2017-06-15 16:38:32',0,5),(1,'','','','','','','','1900-01-01',0,0,'','','',0,0,'','1900-01-01',0,0,'','','',0,'','','','',0,'','','2017-06-15 16:50:13',0,'2017-06-15 16:50:13',0,6);
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
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0'
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
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `RPRRTRateID` bigint(20) NOT NULL AUTO_INCREMENT,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RPRRTRateID`)
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
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `RPRSPRateID` bigint(20) NOT NULL AUTO_INCREMENT,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RPRSPRateID`)
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
  `RAID` bigint(20) NOT NULL DEFAULT '0',
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
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Receipt`
--

LOCK TABLES `Receipt` WRITE;
/*!40000 ALTER TABLE `Receipt` DISABLE KEYS */;
INSERT INTO `Receipt` VALUES (1,0,1,1,1,0,0,0,'2014-03-01 00:00:00','',7000.0000,'',20,'ASM(51) d 10999 7000.00,c 11000 7000.00',2,'initial security deposit','','2017-07-18 21:47:42',0,'2017-07-18 19:35:44',0),(2,0,1,1,1,0,0,0,'2016-03-01 00:00:00','',3750.0000,'',19,'ASM(2) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:36:47',0),(3,0,1,1,3,0,0,0,'2016-04-01 00:00:00','',3750.0000,'',19,'ASM(3) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:37:04',0),(4,0,1,1,3,0,0,0,'2017-05-01 00:00:00','',3750.0000,'',19,'',4,'Reversed by receipt RCPT00000005','','2017-07-18 19:38:20',0,'2017-07-18 19:37:26',0),(5,4,1,1,3,0,0,0,'2017-05-01 00:00:00','',-3750.0000,'',19,'',4,'Reversal of receipt RCPT00000004','','2017-07-18 19:38:20',0,'2017-07-18 19:38:20',0),(6,0,1,1,3,0,0,0,'2016-05-01 00:00:00','',3750.0000,'',19,'',4,'Reversed by receipt RCPT00000022','','2017-07-18 19:44:38',0,'2017-07-18 19:38:20',0),(7,0,1,1,3,0,0,0,'2016-05-01 00:00:00','',3750.0000,'',19,'ASM(4) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:38:55',0),(8,0,1,1,3,0,0,0,'2016-06-01 00:00:00','',3750.0000,'',19,'ASM(5) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:39:17',0),(9,0,1,1,3,0,0,0,'2016-07-01 00:00:00','',3750.0000,'',19,'ASM(6) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:39:41',0),(10,0,1,1,3,0,0,0,'2016-08-01 00:00:00','',3750.0000,'',19,'ASM(7) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:40:00',0),(11,0,1,1,3,0,0,0,'2016-09-01 00:00:00','',3750.0000,'',19,'ASM(8) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:40:17',0),(12,0,1,1,3,0,0,0,'2016-10-01 00:00:00','',3750.0000,'',19,'ASM(9) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:40:33',0),(13,0,1,1,3,0,0,0,'2016-11-01 00:00:00','',3750.0000,'',19,'ASM(10) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:40:50',0),(14,0,1,1,3,0,0,0,'2016-12-01 00:00:00','',3750.0000,'',19,'ASM(11) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:41:03',0),(15,0,1,1,3,0,0,0,'2017-01-01 00:00:00','',3750.0000,'',19,'ASM(12) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:41:23',0),(16,0,1,1,3,0,0,0,'2017-02-01 00:00:00','',3750.0000,'',19,'ASM(13) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:41:38',0),(17,0,1,1,3,0,0,0,'2017-03-01 00:00:00','',3750.0000,'',19,'ASM(14) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:42:11',0),(18,0,1,1,3,0,0,0,'2017-04-01 00:00:00','',3750.0000,'',19,'ASM(15) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:42:27',0),(19,0,1,1,3,0,0,0,'2017-05-01 00:00:00','',3750.0000,'',19,'ASM(16) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:42:43',0),(20,0,1,1,3,0,0,0,'2017-06-01 00:00:00','',3750.0000,'',19,'ASM(17) d 10999 3750.00,c 11000 3750.00',2,'','','2017-07-18 21:47:42',0,'2017-07-18 19:43:06',0),(21,0,1,1,3,0,0,0,'2017-07-01 00:00:00','',3750.0000,'',19,'',2,'','','2017-07-24 07:36:17',0,'2017-07-18 19:43:23',0),(22,6,1,1,3,0,0,0,'2016-05-01 00:00:00','',-3750.0000,'',19,'',4,'Reversal of receipt RCPT00000006','','2017-07-18 19:44:38',0,'2017-07-18 19:44:38',0),(23,0,1,3,1,0,0,0,'2016-07-01 00:00:00','',8300.0000,'',20,'ASM(20) d 10999 4150.00,c 11000 4150.00,ASM(52) d 10999 4150.00,c 11000 4150.00',2,'initial security deposit','','2017-07-18 21:49:52',0,'2017-07-18 19:45:28',0),(24,0,1,3,1,0,0,0,'2016-07-01 00:00:00','',4150.0000,'',19,'ASM(52) d 10999 4150.00,c 11000 4150.00',2,'','','2017-07-18 21:49:52',0,'2017-07-18 19:45:55',0),(25,0,1,3,1,0,0,0,'2016-08-01 00:00:00','',4150.0000,'',19,'ASM(21) d 10999 4150.00,c 11000 4150.00',2,'','','2017-07-18 21:49:52',0,'2017-07-18 19:46:13',0),(26,0,1,3,1,0,0,0,'2016-09-01 00:00:00','',4150.0000,'',19,'ASM(22) d 10999 4150.00,c 11000 4150.00',2,'','','2017-07-18 21:49:52',0,'2017-07-18 19:46:31',0),(27,0,1,3,1,0,0,0,'2016-10-01 00:00:00','',4150.0000,'',19,'ASM(23) d 10999 4150.00,c 11000 4150.00',2,'','','2017-07-18 21:49:53',0,'2017-07-18 19:46:50',0),(28,0,1,3,1,0,0,0,'2016-11-01 00:00:00','',4150.0000,'',19,'ASM(24) d 10999 4150.00,c 11000 4150.00',2,'','','2017-07-18 21:49:53',0,'2017-07-18 19:47:09',0),(29,0,1,3,1,0,0,0,'2016-12-01 00:00:00','',4150.0000,'',19,'ASM(25) d 10999 4150.00,c 11000 4150.00',2,'','','2017-07-18 21:49:53',0,'2017-07-18 19:47:28',0),(30,0,1,3,1,0,0,0,'2017-01-01 00:00:00','',4150.0000,'',19,'ASM(26) d 10999 4150.00,c 11000 4150.00',2,'','','2017-07-18 21:49:53',0,'2017-07-18 19:47:59',0),(31,0,1,3,1,0,0,0,'2017-02-01 00:00:00','',4150.0000,'',19,'ASM(27) d 10999 4150.00,c 11000 4150.00',2,'','','2017-07-18 21:49:53',0,'2017-07-18 19:48:18',0),(32,0,1,3,1,0,0,0,'2017-03-01 00:00:00','',4150.0000,'',19,'ASM(28) d 10999 4150.00,c 11000 4150.00',2,'','','2017-07-18 21:49:53',0,'2017-07-18 19:48:38',0),(33,0,1,3,1,0,0,0,'2017-04-01 00:00:00','',4150.0000,'',19,'ASM(29) d 10999 4150.00,c 11000 4150.00',2,'','','2017-07-18 21:49:53',0,'2017-07-18 19:48:56',0),(34,0,1,3,1,0,0,0,'2017-05-01 00:00:00','',4150.0000,'',19,'ASM(30) d 10999 4150.00,c 11000 4150.00',2,'','','2017-07-18 21:49:53',0,'2017-07-18 19:49:12',0),(35,0,1,3,1,0,0,0,'2017-06-01 00:00:00','',4150.0000,'',19,'ASM(31) d 10999 4150.00,c 11000 4150.00',2,'','','2017-07-18 21:49:53',0,'2017-07-18 19:49:28',0),(36,0,1,3,1,0,0,0,'2017-07-01 00:00:00','',4150.0000,'',19,'ASM(32) d 10999 4150.00,c 11000 4150.00',4,'Reversed by receipt RCPT00000042','','2017-07-24 08:06:52',0,'2017-07-18 19:49:51',0),(37,0,1,6,1,0,0,0,'2016-10-03 00:00:00','',4000.0000,'',19,'ASM(34) d 10999 4000.00,c 11000 4000.00',2,'','','2017-07-18 21:55:07',0,'2017-07-18 19:51:26',0),(38,0,1,6,3,0,0,0,'2016-11-11 00:00:00','',11985.0000,'',19,'ASM(35) d 10999 4000.00,c 11000 4000.00,ASM(36) d 10999 4000.00,c 11000 4000.00,ASM(37) d 10999 3970.00,c 11000 3970.00,ASM(54) d 10999 15.00,c  15.00,ASM(54) d 10999 -15.00',1,'pre-pay 3 months; less $15 wire fee','','2017-08-09 21:22:14',0,'2017-07-18 19:52:33',0),(39,0,1,6,1,0,0,0,'2017-02-03 00:00:00','',628.4500,'',19,'ASM(37) d 10999 30.00,c 11000 30.00,ASM(44) d 10999 598.45,c 11000 598.45',2,'utilities reimbursement','','2017-07-18 21:55:07',0,'2017-07-18 19:54:30',0),(40,0,1,6,3,0,0,0,'2017-02-13 00:00:00','',8335.0000,'',19,'ASM(38) d 10999 4000.00,c 11000 4000.00,ASM(39) d 10999 4000.00,c 11000 4000.00,ASM(45) d 10999 175.00,c 11000 175.00,ASM(46) d 10999 115.00,c 11000 115.00,ASM(44) d 10999 30.00,c 11000 30.00,ASM(55) d 10999 15.00,c  15.00,ASM(46) d 10999 -115.00,ASM(46) c 11000 -115.00',1,'2 month rent/utilities less $15 wire fee','','2017-07-20 21:54:47',0,'2017-07-18 19:56:46',0),(41,0,1,6,3,0,0,0,'2017-05-12 00:00:00','',13116.7900,'',19,'ASM(40) d 10999 4000.00,c 11000 4000.00,ASM(41) d 10999 4000.00,c 11000 4000.00,ASM(42) d 10999 4000.00,c 11000 4000.00,ASM(47) d 10999 350.00,c 11000 350.00,ASM(49) d 10999 350.00,c 11000 350.00,ASM(50) d 10999 275.00,c 11000 275.00,ASM(48) d 10999 81.79,c 11000 81.79,ASM(46) d 10999 45.00,c 11000 45.00,ASM(46) d 10999 -45.00,ASM(46) c 11000 -45.00,ASM(56) d 10999 15.00,c  15.00,ASM(46) d 10999 -115.00,ASM(46) c 11000 -115.00,ASM(46) d 10999 -45.00,ASM(46) c 11000 -45.00',1,'APR/MAY/JUN rent and utilities less $15 wire fee','','2017-07-20 21:54:47',0,'2017-07-18 19:59:41',0),(42,36,1,3,1,0,0,0,'2017-07-01 00:00:00','',-4150.0000,'',19,'ASM(32) d 10999 4150.00,c 11000 4150.00',6,'Reversal of receipt RCPT00000036','','2017-07-24 08:06:52',0,'2017-07-24 08:06:52',0),(43,0,1,3,1,0,0,0,'2017-07-01 00:00:00','',4150.0000,'',20,'',2,'','','2017-07-24 08:06:52',0,'2017-07-24 08:06:52',0);
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
  `AcctRule` varchar(150) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RCPAID`)
) ENGINE=InnoDB AUTO_INCREMENT=107 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ReceiptAllocation`
--

LOCK TABLES `ReceiptAllocation` WRITE;
/*!40000 ALTER TABLE `ReceiptAllocation` DISABLE KEYS */;
INSERT INTO `ReceiptAllocation` VALUES (1,1,1,1,'2014-03-01 00:00:00',7000.0000,0,0,'d 10105 _, c 10999 _','2017-07-27 23:25:30',0,'2017-07-18 19:35:44',0),(2,2,1,1,'2016-03-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:36:47',0),(3,3,1,1,'2016-04-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:37:04',0),(4,4,1,1,'2017-05-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:37:26',0),(5,5,1,1,'2017-05-01 00:00:00',-3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:23:16',0,'2017-07-18 19:38:20',0),(6,6,1,1,'2016-05-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:38:20',0),(7,7,1,1,'2016-05-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:38:55',0),(8,8,1,1,'2016-06-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:39:17',0),(9,9,1,1,'2016-07-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:39:41',0),(10,10,1,1,'2016-08-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:40:00',0),(11,11,1,1,'2016-09-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:40:17',0),(12,12,1,1,'2016-10-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:40:33',0),(13,13,1,1,'2016-11-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:40:50',0),(14,14,1,1,'2016-12-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:41:03',0),(15,15,1,1,'2017-01-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:41:23',0),(16,16,1,1,'2017-02-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:41:38',0),(17,17,1,1,'2017-03-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:42:11',0),(18,18,1,1,'2017-04-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:42:27',0),(19,19,1,1,'2017-05-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:42:43',0),(20,20,1,1,'2017-06-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:43:06',0),(21,21,1,1,'2017-07-01 00:00:00',3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:21:39',0,'2017-07-18 19:43:23',0),(22,22,1,1,'2016-05-01 00:00:00',-3750.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:23:16',0,'2017-07-18 19:44:38',0),(23,23,1,4,'2016-07-01 00:00:00',8300.0000,0,0,'d 10105 _, c 10999 _','2017-07-27 23:26:18',0,'2017-07-18 19:45:29',0),(24,24,1,4,'2016-07-01 00:00:00',4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:22:37',0,'2017-07-18 19:45:55',0),(25,25,1,4,'2016-08-01 00:00:00',4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:22:37',0,'2017-07-18 19:46:13',0),(26,26,1,4,'2016-09-01 00:00:00',4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:22:37',0,'2017-07-18 19:46:31',0),(27,27,1,4,'2016-10-01 00:00:00',4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:22:37',0,'2017-07-18 19:46:50',0),(28,28,1,4,'2016-11-01 00:00:00',4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:22:37',0,'2017-07-18 19:47:09',0),(29,29,1,4,'2016-12-01 00:00:00',4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:22:37',0,'2017-07-18 19:47:28',0),(30,30,1,4,'2017-01-01 00:00:00',4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:22:37',0,'2017-07-18 19:47:59',0),(31,31,1,4,'2017-02-01 00:00:00',4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:22:37',0,'2017-07-18 19:48:18',0),(32,32,1,4,'2017-03-01 00:00:00',4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:22:37',0,'2017-07-18 19:48:38',0),(33,33,1,4,'2017-04-01 00:00:00',4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:22:37',0,'2017-07-18 19:48:56',0),(34,34,1,4,'2017-05-01 00:00:00',4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:22:37',0,'2017-07-18 19:49:12',0),(35,35,1,4,'2017-06-01 00:00:00',4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:22:37',0,'2017-07-18 19:49:28',0),(36,36,1,4,'2017-07-01 00:00:00',4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:22:37',0,'2017-07-18 19:49:51',0),(37,37,1,5,'2016-10-03 00:00:00',4000.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:23:07',0,'2017-07-18 19:51:26',0),(38,38,1,0,'2016-11-11 00:00:00',11985.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:52:33',0,'2017-07-18 19:52:33',0),(39,39,1,0,'2017-02-03 00:00:00',628.4500,0,0,'d 10104 _, c 10999 _','2017-07-18 19:54:30',0,'2017-07-18 19:54:30',0),(40,40,1,0,'2017-02-13 00:00:00',8335.0000,0,0,'d 10104 _, c 10999 _','2017-07-18 19:56:46',0,'2017-07-18 19:56:46',0),(41,41,1,0,'2017-05-12 00:00:00',13116.7900,0,0,'d 10104 _, c 10999 _','2017-07-18 19:59:41',0,'2017-07-18 19:59:41',0),(42,1,1,1,'2014-03-01 00:00:00',7000.0000,51,0,'ASM(51) d 10999 7000.00,c 11000 7000.00','2017-07-27 22:41:23',0,'2017-07-18 21:47:42',0),(43,2,1,1,'2016-03-01 00:00:00',3750.0000,2,0,'ASM(2) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(44,3,1,1,'2016-04-01 00:00:00',3750.0000,3,0,'ASM(3) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(45,7,1,1,'2016-05-01 00:00:00',3750.0000,4,0,'ASM(4) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(46,8,1,1,'2016-06-01 00:00:00',3750.0000,5,0,'ASM(5) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(47,9,1,1,'2016-07-01 00:00:00',3750.0000,6,0,'ASM(6) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(48,10,1,1,'2016-08-01 00:00:00',3750.0000,7,0,'ASM(7) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(49,11,1,1,'2016-09-01 00:00:00',3750.0000,8,0,'ASM(8) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(50,12,1,1,'2016-10-01 00:00:00',3750.0000,9,0,'ASM(9) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(51,13,1,1,'2016-11-01 00:00:00',3750.0000,10,0,'ASM(10) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(52,14,1,1,'2016-12-01 00:00:00',3750.0000,11,0,'ASM(11) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(53,15,1,1,'2017-01-01 00:00:00',3750.0000,12,0,'ASM(12) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(54,16,1,1,'2017-02-01 00:00:00',3750.0000,13,0,'ASM(13) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(55,17,1,1,'2017-03-01 00:00:00',3750.0000,14,0,'ASM(14) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(56,18,1,1,'2017-04-01 00:00:00',3750.0000,15,0,'ASM(15) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(57,19,1,1,'2017-05-01 00:00:00',3750.0000,16,0,'ASM(16) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(58,20,1,1,'2017-06-01 00:00:00',3750.0000,17,0,'ASM(17) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(59,21,1,1,'2017-07-01 00:00:00',3750.0000,18,0,'ASM(18) d 10999 3750.00,c 11000 3750.00','2017-07-27 22:36:19',0,'2017-07-18 21:47:42',0),(60,23,1,4,'2016-07-01 00:00:00',4150.0000,20,0,'ASM(20) d 10999 4150.00,c 11000 4150.00','2017-07-27 22:36:53',0,'2017-07-18 21:49:52',0),(61,23,1,4,'2016-07-01 00:00:00',4150.0000,52,0,'ASM(52) d 10999 4150.00,c 11000 4150.00','2017-07-27 22:36:53',0,'2017-07-18 21:49:52',0),(62,24,1,4,'2016-07-01 00:00:00',4150.0000,52,0,'ASM(52) d 10999 4150.00,c 11000 4150.00','2017-07-27 22:36:53',0,'2017-07-18 21:49:52',0),(63,25,1,4,'2016-08-01 00:00:00',4150.0000,21,0,'ASM(21) d 10999 4150.00,c 11000 4150.00','2017-07-27 22:36:53',0,'2017-07-18 21:49:52',0),(64,26,1,4,'2016-09-01 00:00:00',4150.0000,22,0,'ASM(22) d 10999 4150.00,c 11000 4150.00','2017-07-27 22:36:53',0,'2017-07-18 21:49:52',0),(65,27,1,4,'2016-10-01 00:00:00',4150.0000,23,0,'ASM(23) d 10999 4150.00,c 11000 4150.00','2017-07-27 22:36:53',0,'2017-07-18 21:49:53',0),(66,28,1,4,'2016-11-01 00:00:00',4150.0000,24,0,'ASM(24) d 10999 4150.00,c 11000 4150.00','2017-07-27 22:36:53',0,'2017-07-18 21:49:53',0),(67,29,1,4,'2016-12-01 00:00:00',4150.0000,25,0,'ASM(25) d 10999 4150.00,c 11000 4150.00','2017-07-27 22:36:53',0,'2017-07-18 21:49:53',0),(68,30,1,4,'2017-01-01 00:00:00',4150.0000,26,0,'ASM(26) d 10999 4150.00,c 11000 4150.00','2017-07-27 22:36:53',0,'2017-07-18 21:49:53',0),(69,31,1,4,'2017-02-01 00:00:00',4150.0000,27,0,'ASM(27) d 10999 4150.00,c 11000 4150.00','2017-07-27 22:36:53',0,'2017-07-18 21:49:53',0),(70,32,1,4,'2017-03-01 00:00:00',4150.0000,28,0,'ASM(28) d 10999 4150.00,c 11000 4150.00','2017-07-27 22:36:53',0,'2017-07-18 21:49:53',0),(71,33,1,4,'2017-04-01 00:00:00',4150.0000,29,0,'ASM(29) d 10999 4150.00,c 11000 4150.00','2017-07-27 22:36:53',0,'2017-07-18 21:49:53',0),(72,34,1,4,'2017-05-01 00:00:00',4150.0000,30,0,'ASM(30) d 10999 4150.00,c 11000 4150.00','2017-07-27 22:36:53',0,'2017-07-18 21:49:53',0),(73,35,1,4,'2017-06-01 00:00:00',4150.0000,31,0,'ASM(31) d 10999 4150.00,c 11000 4150.00','2017-07-27 22:36:53',0,'2017-07-18 21:49:53',0),(74,36,1,4,'2017-07-01 00:00:00',4150.0000,32,0,'ASM(32) d 10999 4150.00,c 11000 4150.00','2017-07-27 22:36:53',0,'2017-07-18 21:49:53',0),(75,37,1,5,'2016-10-03 00:00:00',4000.0000,34,0,'ASM(34) d 10999 4000.00,c 11000 4000.00','2017-07-27 22:37:20',0,'2017-07-18 21:55:07',0),(76,38,1,5,'2016-11-11 00:00:00',4000.0000,35,0,'ASM(35) d 10999 4000.00,c 11000 4000.00','2017-07-27 22:37:20',0,'2017-07-18 21:55:07',0),(77,38,1,5,'2016-11-11 00:00:00',15.0000,54,4,'ASM(54) d 10999 15.00,c  15.00','2017-08-09 21:22:14',0,'2017-07-18 21:55:07',0),(78,38,1,5,'2016-11-11 00:00:00',4000.0000,36,0,'ASM(36) d 10999 4000.00,c 11000 4000.00','2017-07-27 22:37:20',0,'2017-07-18 21:55:07',0),(79,38,1,5,'2016-11-11 00:00:00',3970.0000,37,0,'ASM(37) d 10999 3970.00,c 11000 3970.00','2017-07-27 23:02:22',0,'2017-07-18 21:55:07',0),(80,39,1,5,'2017-02-03 00:00:00',30.0000,37,0,'ASM(37) d 10999 30.00,c 11000 30.00','2017-07-27 23:02:45',0,'2017-07-18 21:55:07',0),(81,39,1,5,'2017-02-03 00:00:00',598.4500,44,0,'ASM(44) d 10999 598.45,c 11000 598.45','2017-07-27 23:04:28',0,'2017-07-18 21:55:07',0),(82,40,1,5,'2017-02-13 00:00:00',30.0000,44,0,'ASM(44) d 10999 30.00,c 11000 30.00','2017-07-27 23:04:28',0,'2017-07-18 21:55:07',0),(83,40,1,5,'2017-02-13 00:00:00',4000.0000,38,0,'ASM(38) d 10999 4000.00,c 11000 4000.00','2017-07-27 22:37:20',0,'2017-07-18 21:55:07',0),(84,40,1,5,'2017-02-13 00:00:00',175.0000,45,0,'ASM(45) d 10999 175.00,c 11000 175.00','2017-07-27 23:04:28',0,'2017-07-18 21:55:07',0),(85,40,1,5,'2017-02-13 00:00:00',15.0000,55,0,'ASM(55) d 10999 15.00,c  15.00','2017-07-27 23:04:28',0,'2017-07-18 21:55:07',0),(86,40,1,5,'2017-02-13 00:00:00',4000.0000,39,0,'ASM(39) d 10999 4000.00,c 11000 4000.00','2017-07-27 22:37:20',0,'2017-07-18 21:55:07',0),(87,40,1,5,'2017-02-13 00:00:00',115.0000,46,4,'ASM(46) d 10999 115.00,c 11000 115.00','2017-07-27 23:08:26',0,'2017-07-18 21:55:07',0),(88,41,1,5,'2017-05-12 00:00:00',45.0000,46,4,'ASM(46) d 10999 45.00,c 11000 45.00','2017-07-27 23:09:26',0,'2017-07-18 21:55:07',0),(89,41,1,5,'2017-05-12 00:00:00',4000.0000,40,0,'ASM(40) d 10999 4000.00,c 11000 4000.00','2017-07-27 22:37:20',0,'2017-07-18 21:55:07',0),(90,41,1,5,'2017-05-12 00:00:00',350.0000,47,0,'ASM(47) d 10999 350.00,c 11000 350.00','2017-07-27 23:09:30',0,'2017-07-18 21:55:07',0),(91,41,1,5,'2017-05-12 00:00:00',81.7900,48,0,'ASM(48) d 10999 81.79,c 11000 81.79','2017-07-27 23:09:37',0,'2017-07-18 21:55:07',0),(92,41,1,5,'2017-05-12 00:00:00',4000.0000,41,0,'ASM(41) d 10999 4000.00,c 11000 4000.00','2017-07-27 22:37:20',0,'2017-07-18 21:55:07',0),(93,41,1,5,'2017-05-12 00:00:00',350.0000,49,0,'ASM(49) d 10999 350.00,c 11000 350.00','2017-07-27 23:09:43',0,'2017-07-18 21:55:07',0),(94,41,1,5,'2017-05-12 00:00:00',15.0000,56,0,'ASM(56) d 10999 15.00,c  15.00','2017-07-27 23:09:49',0,'2017-07-18 21:55:07',0),(95,41,1,5,'2017-05-12 00:00:00',4000.0000,42,0,'ASM(42) d 10999 4000.00,c 11000 4000.00','2017-07-27 22:37:20',0,'2017-07-18 21:55:07',0),(96,41,1,5,'2017-05-12 00:00:00',275.0000,50,0,'ASM(50) d 10999 275.00,c 11000 275.00','2017-07-27 23:09:53',0,'2017-07-18 21:55:07',0),(97,40,1,5,'2017-07-20 21:54:47',-115.0000,46,4,'ASM(46) d 10999 -115.00,ASM(46) c 11000 -115.00','2017-07-27 23:10:06',0,'2017-07-20 21:54:47',0),(98,41,1,5,'2017-07-20 21:54:47',-45.0000,46,4,'ASM(46) d 10999 -115.00,ASM(46) c 11000 -115.00','2017-07-27 23:10:12',0,'2017-07-20 21:54:47',0),(99,40,1,5,'2017-07-20 21:54:47',-115.0000,46,4,'ASM(46) d 10999 -45.00,ASM(46) c 11000 -45.00','2017-07-27 23:10:16',0,'2017-07-20 21:54:47',0),(100,41,1,5,'2017-07-20 21:54:47',-45.0000,46,4,'ASM(46) d 10999 -45.00,ASM(46) c 11000 -45.00','2017-07-27 23:10:24',0,'2017-07-20 21:54:47',0),(101,40,1,5,'2017-07-20 21:54:47',115.0000,46,4,'ASM(46) d 10999 -45.00,ASM(46) c 11000 -45.00','2017-07-27 23:10:29',0,'2017-07-20 21:54:47',0),(102,41,1,5,'2017-07-20 21:54:47',45.0000,46,4,'ASM(46) d 10999 -45.00,ASM(46) c 11000 -45.00','2017-07-27 23:10:32',0,'2017-07-20 21:54:47',0),(103,42,1,4,'2017-07-01 00:00:00',-4150.0000,0,0,'d 10104 _, c 10999 _','2017-07-27 23:22:44',0,'2017-07-24 08:06:52',0),(104,42,1,4,'2017-07-24 08:06:53',-4150.0000,32,0,'ASM(32) d 10999 -4150.00,ASM(32) c 11000 -4150.00','2017-07-27 22:39:56',0,'2017-07-24 08:06:52',0),(105,43,1,4,'2017-07-01 00:00:00',4150.0000,0,0,'d 10105 _, c 10999 _','2017-07-27 23:22:37',0,'2017-07-24 08:06:52',0),(106,38,1,5,'2017-08-09 21:22:15',-15.0000,54,4,'ASM(54) d 10999 -15.00','2017-08-09 21:22:14',0,'2017-08-09 21:22:14',0);
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
  `MRStatus` smallint(6) NOT NULL DEFAULT '0',
  `DtMRStart` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `Comment` varchar(2048) NOT NULL DEFAULT '',
  PRIMARY KEY (`RID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Rentable`
--

LOCK TABLES `Rentable` WRITE;
/*!40000 ALTER TABLE `Rentable` DISABLE KEYS */;
INSERT INTO `Rentable` VALUES (1,1,'309 S Rexford',1,0,'2017-10-10 04:29:04','2017-06-13 19:34:58',0,'2017-06-14 18:26:51',0,''),(2,1,'309 1/2 S Rexford',1,0,'2017-10-10 04:29:04','2017-06-13 20:02:10',0,'2017-06-14 18:26:51',0,''),(3,1,'311 S Rexford',1,0,'2017-10-10 04:29:04','2017-06-13 20:02:33',0,'2017-06-14 18:26:51',0,''),(4,1,'311 1/2 S Rexford',1,0,'2017-10-10 04:29:04','2017-06-13 20:03:01',0,'2017-06-14 18:26:51',0,'');
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RMRID`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableMarketRate`
--

LOCK TABLES `RentableMarketRate` WRITE;
/*!40000 ALTER TABLE `RentableMarketRate` DISABLE KEYS */;
INSERT INTO `RentableMarketRate` VALUES (1,1,1,3500.0000,'2014-01-01 00:00:00','3000-01-01 00:00:00','2017-06-14 18:26:52',0,'2018-01-01 09:59:49',0),(2,2,1,3550.0000,'2014-01-01 00:00:00','3000-01-01 00:00:00','2017-06-14 18:26:52',0,'2018-01-01 09:59:49',0),(3,3,1,4400.0000,'2014-01-01 00:00:00','3000-01-01 00:00:00','2017-06-14 18:26:52',0,'2018-01-01 09:59:49',0),(4,4,1,2500.0000,'2014-01-01 00:00:00','2017-06-15 05:09:09','2017-06-14 18:26:52',0,'2018-01-01 09:59:49',0),(51,4,1,0.0000,'2017-06-15 05:09:09','2017-06-15 05:17:56','2017-06-15 05:09:09',0,'2018-01-01 09:59:49',0),(52,4,1,2400.0000,'2017-06-15 05:17:56','9998-12-31 00:00:00','2017-06-15 05:17:55',0,'2018-01-01 09:59:49',0);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
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
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `RSPRefID` bigint(20) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`RSPRefID`)
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
  `UseStatus` smallint(6) NOT NULL DEFAULT '0',
  `LeaseStatus` smallint(6) NOT NULL DEFAULT '0',
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
INSERT INTO `RentableStatus` VALUES (5,1,1,1,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-06-20 16:09:45',0,'2017-06-20 16:09:45',0),(6,2,1,1,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-06-20 16:20:18',0,'2017-06-20 16:20:18',0),(7,3,1,1,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-06-20 16:20:30',0,'2017-06-20 16:20:30',0),(8,4,1,4,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-06-20 16:20:41',0,'2017-06-20 16:20:41',0);
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
INSERT INTO `RentableTypeRef` VALUES (5,1,1,1,0,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','2017-06-20 16:09:45',0,'2017-06-20 16:09:45',0),(6,2,1,2,0,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','2017-06-20 16:20:18',0,'2017-06-20 16:20:18',0),(7,3,1,3,0,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','2017-06-20 16:20:30',0,'2017-06-20 16:20:30',0),(8,4,1,4,0,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','2017-06-20 16:20:41',0,'2017-06-20 16:20:41',0);
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
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0'
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
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `ARID` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RTID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableTypes`
--

LOCK TABLES `RentableTypes` WRITE;
/*!40000 ALTER TABLE `RentableTypes` DISABLE KEYS */;
INSERT INTO `RentableTypes` VALUES (1,1,'Rex1','309 Rexford',6,4,4,0,0,'2017-06-13 05:39:46',0,'2017-06-14 18:26:53',0),(2,1,'Rex2','309 1/2 Rexford',6,4,4,0,0,'2017-06-13 05:39:46',0,'2017-06-14 18:26:53',0),(3,1,'Rex3','311 Rexford',6,4,4,0,0,'2017-06-13 05:39:46',0,'2017-06-14 18:26:53',0),(4,1,'Rex4','311 1/2 Rexford',6,4,4,0,0,'2017-06-15 05:43:52',0,'2017-06-14 18:26:53',0);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RUID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableUsers`
--

LOCK TABLES `RentableUsers` WRITE;
/*!40000 ALTER TABLE `RentableUsers` DISABLE KEYS */;
INSERT INTO `RentableUsers` VALUES (1,1,1,2,'2014-03-01','2018-03-01','2017-07-07 22:41:49',0,'2018-01-01 09:59:49',0),(2,2,1,5,'2016-10-01','2017-12-31','2017-07-18 18:58:12',0,'2018-01-01 09:59:49',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreement`
--

LOCK TABLES `RentalAgreement` WRITE;
/*!40000 ALTER TABLE `RentalAgreement` DISABLE KEYS */;
INSERT INTO `RentalAgreement` VALUES (1,0,1,0,'2014-03-01','2018-03-01','2014-03-01','2018-03-01','2014-03-01','2018-03-01','2016-03-01',2,1,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2017-07-18 18:54:37',0,'2017-07-07 22:25:57',0),(4,0,1,0,'2016-07-01','2018-07-01','2016-07-01','2018-07-01','2016-07-01','2018-07-01','2017-07-01',2,1,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2017-07-18 18:56:40',0,'2017-07-18 18:54:44',0),(5,0,1,0,'2016-10-01','2017-12-31','2016-10-01','2017-12-31','2016-10-01','2017-12-31','2016-10-01',1,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2017-07-18 18:58:15',0,'2017-07-18 18:56:42',0);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RAPID`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementPayors`
--

LOCK TABLES `RentalAgreementPayors` WRITE;
/*!40000 ALTER TABLE `RentalAgreementPayors` DISABLE KEYS */;
INSERT INTO `RentalAgreementPayors` VALUES (1,1,1,1,'2014-03-01','2018-03-01',0,'2017-07-07 22:40:32',0,'2018-01-01 09:59:50',0),(2,1,1,2,'2014-03-01','2018-03-01',0,'2017-07-18 18:54:16',0,'2018-01-01 09:59:50',0),(3,4,1,3,'2016-07-01','2018-07-01',0,'2017-07-18 18:56:11',0,'2018-01-01 09:59:50',0),(4,4,1,4,'2016-07-01','2018-07-01',0,'2017-07-18 18:56:28',0,'2018-01-01 09:59:50',0),(5,5,1,6,'2016-10-01','2017-12-31',0,'2017-07-18 18:57:44',0,'2018-01-01 09:59:50',0);
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
  `TCID` bigint(20) NOT NULL DEFAULT '0',
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
  `PRID` bigint(20) NOT NULL DEFAULT '0',
  `CLID` bigint(20) NOT NULL DEFAULT '0',
  `ContractRent` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `RARDtStart` date NOT NULL DEFAULT '1970-01-01',
  `RARDtStop` date NOT NULL DEFAULT '1970-01-01',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RARID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementRentables`
--

LOCK TABLES `RentalAgreementRentables` WRITE;
/*!40000 ALTER TABLE `RentalAgreementRentables` DISABLE KEYS */;
INSERT INTO `RentalAgreementRentables` VALUES (2,1,1,1,0,0,3750.0000,'2016-03-01','2018-03-01','2017-07-18 18:53:59',0,'2018-01-01 09:59:50',0),(3,4,1,3,0,0,4150.0000,'2016-07-01','2018-07-01','2017-07-18 18:55:39',0,'2018-01-01 09:59:50',0),(4,5,1,2,0,0,4000.0000,'2016-10-01','2018-12-31','2017-07-18 18:57:24',0,'2018-01-01 09:59:50',0);
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
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0'
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
-- Table structure for table `SubAR`
--

DROP TABLE IF EXISTS `SubAR`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `SubAR` (
  `SARID` bigint(20) NOT NULL AUTO_INCREMENT,
  `ARID` bigint(20) NOT NULL DEFAULT '0',
  `SubARID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`SARID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SubAR`
--

LOCK TABLES `SubAR` WRITE;
/*!40000 ALTER TABLE `SubAR` DISABLE KEYS */;
/*!40000 ALTER TABLE `SubAR` ENABLE KEYS */;
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
INSERT INTO `TWS` VALUES (1,'CreateAssessmentInstances','','CreateAssessmentInstances','2017-09-29 00:00:00','Steves-MacBook-Pro-2.local',4,'2017-09-28 03:46:54','2017-09-28 03:46:54','2017-07-18 09:03:33','2017-09-27 20:46:54'),(2,'CreateAssessmentInstances','','CreateAssessmentInstances','2017-07-26 00:00:00','ip-172-31-51-141.ec2.internal',4,'2017-07-25 00:00:06','2017-07-25 00:00:06','2017-07-19 04:14:03','2017-07-25 00:00:06');
/*!40000 ALTER TABLE `TWS` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Task`
--

DROP TABLE IF EXISTS `Task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Task` (
  `TID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `TLID` bigint(20) NOT NULL DEFAULT '0',
  `Name` varchar(256) NOT NULL DEFAULT '',
  `Worker` varchar(80) NOT NULL DEFAULT '',
  `DtDue` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtPreDue` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtDone` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtPreDone` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `DoneUID` bigint(20) NOT NULL DEFAULT '0',
  `PreDoneUID` bigint(20) NOT NULL DEFAULT '0',
  `Comment` varchar(2048) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`TID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Task`
--

LOCK TABLES `Task` WRITE;
/*!40000 ALTER TABLE `Task` DISABLE KEYS */;
/*!40000 ALTER TABLE `Task` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `TaskDescriptor`
--

DROP TABLE IF EXISTS `TaskDescriptor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `TaskDescriptor` (
  `TDID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `TLDID` bigint(20) NOT NULL DEFAULT '0',
  `Name` varchar(256) NOT NULL DEFAULT '',
  `Worker` varchar(80) NOT NULL DEFAULT '',
  `EpochDue` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `EpochPreDue` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `Comment` varchar(2048) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`TDID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TaskDescriptor`
--

LOCK TABLES `TaskDescriptor` WRITE;
/*!40000 ALTER TABLE `TaskDescriptor` DISABLE KEYS */;
/*!40000 ALTER TABLE `TaskDescriptor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `TaskList`
--

DROP TABLE IF EXISTS `TaskList`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `TaskList` (
  `TLID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `PTLID` bigint(20) NOT NULL DEFAULT '0',
  `TLDID` bigint(20) NOT NULL DEFAULT '0',
  `Name` varchar(256) NOT NULL DEFAULT '',
  `Cycle` bigint(20) NOT NULL DEFAULT '0',
  `DtDue` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtPreDue` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtDone` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtPreDone` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `DoneUID` bigint(20) NOT NULL DEFAULT '0',
  `PreDoneUID` bigint(20) NOT NULL DEFAULT '0',
  `EmailList` varchar(2048) NOT NULL DEFAULT '',
  `DtLastNotify` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DurWait` bigint(20) NOT NULL DEFAULT '86400000000000',
  `Comment` varchar(2048) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`TLID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TaskList`
--

LOCK TABLES `TaskList` WRITE;
/*!40000 ALTER TABLE `TaskList` DISABLE KEYS */;
/*!40000 ALTER TABLE `TaskList` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `TaskListDefinition`
--

DROP TABLE IF EXISTS `TaskListDefinition`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `TaskListDefinition` (
  `TLDID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Name` varchar(256) NOT NULL DEFAULT '',
  `Cycle` bigint(20) NOT NULL DEFAULT '0',
  `Epoch` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `EpochDue` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `EpochPreDue` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `EmailList` varchar(2048) NOT NULL DEFAULT '',
  `DurWait` bigint(20) NOT NULL DEFAULT '86400000000000',
  `Comment` varchar(2048) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`TLDID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TaskListDefinition`
--

LOCK TABLES `TaskListDefinition` WRITE;
/*!40000 ALTER TABLE `TaskListDefinition` DISABLE KEYS */;
/*!40000 ALTER TABLE `TaskListDefinition` ENABLE KEYS */;
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
  `IsCompany` tinyint(1) NOT NULL DEFAULT '0',
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
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `Comment` varchar(2048) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`TCID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Transactant`
--

LOCK TABLES `Transactant` WRITE;
/*!40000 ALTER TABLE `Transactant` DISABLE KEYS */;
INSERT INTO `Transactant` VALUES (1,1,0,'Aaron','','Read','','',0,'read.aaron@gmail.com','','','1-469-307-7095','','','','','','','',0,'','2017-06-15 16:33:59',0,'2017-06-14 18:26:55',0),(2,1,0,'Kirsten','','Read','','',0,'klmrda@gmail.com','','','1-469-693-9933','','','','','','','',0,'','2017-06-15 16:34:39',0,'2017-06-14 18:26:55',0),(3,1,0,'Kevin','','Mills','','',0,'kevinmillsesq@aol.com','','','1-424-234-3535','','','','','','','',0,'','2017-06-15 16:35:44',0,'2017-06-15 16:35:44',0),(4,1,0,'Lauren','','Beck','','',0,'laurensbeck@aol.com','','','1-310-948-6442','','','','','','','',0,'','2017-06-15 16:36:27',0,'2017-06-15 16:36:27',0),(5,1,0,'Alex','','Vahabzadeh','','Beaumont Partners LLC',0,'av@beaumont-partners.ch','','44-79-203-354-77','1-202-550-2477','118 Rue du Rhone','','1204 Geneva','','','Switzerland','',0,'','2017-06-15 16:50:29',0,'2017-06-15 16:38:32',0),(6,1,0,'','','','','Beaumont Partners LLC',1,'av@beaumont-partners.ch','scigler@bvgroup.com','44-79-203-354-77','1-202-550-2477','118 Rue du Rhone','','1204 Geneva','','','Switzerland','',0,'','2017-06-15 16:50:13',0,'2017-06-15 16:50:13',0);
/*!40000 ALTER TABLE `Transactant` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `User` (
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Points` bigint(20) NOT NULL DEFAULT '0',
  `DateofBirth` date NOT NULL DEFAULT '1970-01-01',
  `EmergencyContactName` varchar(100) NOT NULL DEFAULT '',
  `EmergencyContactAddress` varchar(100) NOT NULL DEFAULT '',
  `EmergencyContactTelephone` varchar(100) NOT NULL DEFAULT '',
  `EmergencyContactEmail` varchar(100) NOT NULL DEFAULT '',
  `AlternateAddress` varchar(100) NOT NULL DEFAULT '',
  `EligibleFutureUser` tinyint(1) NOT NULL DEFAULT '1',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `Industry` varchar(100) NOT NULL DEFAULT '',
  `SourceSLSID` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `TCID` bigint(20) NOT NULL,
  PRIMARY KEY (`TCID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` VALUES (1,0,'1900-01-01','','','','','',1,0,'',0,'2017-06-13 19:39:18',0,'2017-06-14 18:26:55',0,1),(1,0,'1900-01-01','','','','','',1,0,'',0,'2017-06-13 19:40:59',0,'2017-06-14 18:26:55',0,2),(1,0,'1900-01-01','','','','','',1,0,'',0,'2017-06-15 16:35:44',0,'2017-06-15 16:35:44',0,3),(1,0,'1900-01-01','','','','','',1,0,'',0,'2017-06-15 16:36:27',0,'2017-06-15 16:36:27',0,4),(1,0,'1900-01-01','','','','','',1,0,'',0,'2017-06-15 16:38:32',0,'2017-06-15 16:38:32',0,5),(1,0,'1900-01-01','','','','','',1,0,'',0,'2017-06-15 16:50:13',0,'2017-06-15 16:50:13',0,6);
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
  `VIN` varchar(20) NOT NULL DEFAULT '',
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

-- Dump completed on 2018-06-20 14:16:00
