-- MySQL dump 10.13  Distrib 5.7.22, for Linux (x86_64)
--
-- Host: localhost    Database: rentroll
-- ------------------------------------------------------
-- Server version	5.7.22-0ubuntu0.16.04.1-log

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
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `AR`
--

LOCK TABLES `AR` WRITE;
/*!40000 ALTER TABLE `AR` DISABLE KEYS */;
INSERT INTO `AR` VALUES (1,1,'Rent Taxable',0,3,8,16,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:23:21',0,'2017-08-16 23:19:30',0),(2,1,'Rent Non-Taxable',0,3,8,17,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:23:11',0,'2017-08-16 23:20:05',0),(3,1,'Electric Overage',0,3,8,36,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:26:13',0,'2017-08-16 23:22:28',0),(4,1,'Electric Base Fee',0,3,8,35,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:25:44',0,'2017-08-16 23:25:44',0),(7,1,'Water and Sewer Base Fee',0,3,8,37,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:29:50',0,'2017-08-16 23:29:50',0),(8,1,'Water and Sewer Overage',0,3,8,38,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:30:22',0,'2017-08-16 23:30:22',0),(9,1,'Gas Base Fee',0,3,8,39,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:31:02',0,'2017-08-16 23:31:02',0),(10,1,'Gas Base Overage',0,3,8,40,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:31:27',0,'2017-08-16 23:31:27',0),(11,1,'Application Fee',0,3,8,45,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:36:16',0,'2017-08-16 23:32:46',0),(12,1,'Late Fee',0,3,8,46,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:33:10',0,'2017-08-16 23:33:10',0),(13,1,'Month to Month Fee',0,3,8,48,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:33:34',0,'2017-08-16 23:33:34',0),(14,1,'Insufficient Funds Fee',0,3,8,47,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:34:00',0,'2017-08-16 23:34:00',0),(15,1,'No Show / Termination Fee',0,3,8,50,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:35:03',0,'2017-08-16 23:35:03',0),(16,1,'Pet Fee',0,3,8,51,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:35:23',0,'2017-08-16 23:35:23',0),(17,1,'Pet Rent',0,3,8,52,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:41:40',0,'2017-08-16 23:35:54',0),(18,1,'Tenant Expense Chargeback',0,3,8,53,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:36:59',0,'2017-08-16 23:36:59',0),(19,1,'Special Cleaning Fee',0,3,8,54,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:37:27',0,'2017-08-16 23:37:27',0),(20,1,'Eviction Fee Reimbursement',0,3,8,55,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:37:59',0,'2017-08-16 23:37:59',0),(21,1,'Security Deposit Forfeiture',0,3,11,57,'Forfeit','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-09-21 02:21:35',0,'2017-08-16 23:38:38',0),(22,1,'Damage Fee',0,3,8,58,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:39:18',0,'2017-08-16 23:39:18',0),(23,1,'Other Special Tenant Charges',0,3,8,60,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-16 23:39:57',0,'2017-08-16 23:39:57',0),(24,1,'Security Deposit Assessment',0,3,8,11,'normal deposit','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-09-25 21:21:05',0,'2017-08-16 23:40:26',0),(25,1,'Bad Debt Write-Off',2,3,70,8,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-22 13:38:33',0,'2017-08-22 13:38:33',0),(26,1,'Bank Service Fee (Operating Account)',2,3,71,3,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-23 14:10:11',0,'2017-08-22 13:40:32',0),(27,1,'Receive a Payment',1,3,5,10,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-22 13:41:19',0,'2017-08-22 13:41:19',0),(28,1,'Deposit to Operating Account (FRB54320)',1,3,3,5,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-22 13:42:29',0,'2017-08-22 13:42:29',0),(29,1,'Deposit to Deposit Account (FRB96953)',1,3,4,5,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-22 13:43:44',0,'2017-08-22 13:43:22',0),(30,1,'Apply Payment',1,3,10,8,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-22 14:06:06',0,'2017-08-22 14:05:19',0),(31,1,'Bank Service Fee (Deposit Account)',2,3,71,4,'','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-08-23 14:11:03',0,'2017-08-23 14:11:03',0),(32,1,'Broken Window charge',0,3,8,58,'','2017-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-09-15 06:42:13',0,'2017-09-15 06:42:13',0),(33,1,'Vending Income',1,3,6,64,'','2010-01-01 00:00:00','9999-12-31 00:00:00',5,0.0000,'2017-09-29 21:56:01',0,'2017-09-18 19:37:22',0),(34,1,'Security Deposit Refund',0,3,11,73,'Refund','2014-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-09-21 02:22:30',0,'2017-09-21 02:22:30',0),(35,1,'Floating Security Deposit',1,3,5,75,'','2014-01-01 00:00:00','9999-12-31 00:00:00',5,0.0000,'2017-09-29 21:57:50',0,'2017-09-22 21:20:49',0),(37,1,'Application Fee (no assessment)',1,3,6,45,'Application Fee taken, no assessment made','2017-01-01 00:00:00','9999-12-31 00:00:00',5,0.0000,'2017-10-01 21:01:43',0,'2017-10-01 21:01:29',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Assessments`
--

LOCK TABLES `Assessments` WRITE;
/*!40000 ALTER TABLE `Assessments` DISABLE KEYS */;
INSERT INTO `Assessments` VALUES (1,0,0,0,1,2,0,3,4000.0000,'2016-10-01 00:00:00','2018-01-01 00:00:00',6,4,0,'',2,0,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(2,1,0,0,1,2,0,3,4000.0000,'2016-10-01 00:00:00','2016-10-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:40:14',0,'2017-08-22 18:24:26',0),(3,1,0,0,1,2,0,3,4000.0000,'2016-11-01 00:00:00','2016-11-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:40:14',0,'2017-08-22 18:24:26',0),(4,1,0,0,1,2,0,3,4000.0000,'2016-12-01 00:00:00','2016-12-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:40:14',0,'2017-08-22 18:24:26',0),(5,1,0,0,1,2,0,3,4000.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:40:14',0,'2017-08-22 18:24:26',0),(6,1,0,0,1,2,0,3,4000.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:53:58',0,'2017-08-22 18:24:26',0),(7,1,0,0,1,2,0,3,4000.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:53:58',0,'2017-08-22 18:24:26',0),(8,1,0,0,1,2,0,3,4000.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:53:58',0,'2017-08-22 18:24:26',0),(9,1,0,0,1,2,0,3,4000.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:53:58',0,'2017-08-22 18:24:26',0),(10,1,0,0,1,2,0,3,4000.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:53:58',0,'2017-08-22 18:24:26',0),(11,1,0,0,1,2,0,3,4000.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:53:58',0,'2017-08-22 18:24:26',0),(12,1,0,0,1,2,0,3,4000.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',2,2,'','2017-08-22 18:53:58',0,'2017-08-22 18:24:26',0),(13,0,0,0,1,2,0,3,628.4500,'2017-01-01 00:00:00','2017-01-01 00:00:00',0,0,0,'',4,2,'utilities allowance','2017-08-22 18:53:58',0,'2017-08-22 18:27:44',0),(14,0,0,0,1,2,0,3,175.0000,'2017-02-01 00:00:00','2017-02-01 00:00:00',0,0,0,'',4,2,'utilities','2017-08-22 18:53:58',0,'2017-08-22 18:28:39',0),(15,0,0,0,1,2,0,3,175.0000,'2017-02-01 00:00:00','2017-02-01 00:00:00',0,0,0,'',4,4,'Reversed by ASM00000024','2017-08-22 18:52:07',0,'2017-08-22 18:47:56',0),(16,0,0,0,1,2,0,3,175.0000,'2017-03-01 00:00:00','2017-03-01 00:00:00',0,0,0,'',4,2,'utilities','2017-08-22 18:53:58',0,'2017-08-22 18:48:32',0),(17,0,0,0,1,2,0,3,81.7900,'2017-04-01 00:00:00','2017-04-01 00:00:00',0,0,0,'',4,2,'retro utilities','2017-08-22 18:53:58',0,'2017-08-22 18:49:22',0),(18,0,0,0,1,2,0,3,350.0000,'2017-04-01 00:00:00','2017-12-31 00:00:00',6,4,0,'',4,0,'monthly utilities allowance','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(19,18,0,0,1,2,0,3,350.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',4,2,'monthly utilities allowance','2017-08-22 18:53:58',0,'2017-08-22 18:50:08',0),(20,18,0,0,1,2,0,3,350.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',4,2,'monthly utilities allowance','2017-08-22 18:53:58',0,'2017-08-22 18:50:08',0),(21,18,0,0,1,2,0,3,350.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',4,2,'monthly utilities allowance','2017-08-22 18:53:58',0,'2017-08-22 18:50:08',0),(22,18,0,0,1,2,0,3,350.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',4,2,'monthly utilities allowance','2017-08-22 18:53:58',0,'2017-08-22 18:50:08',0),(23,18,0,0,1,2,0,3,350.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',4,2,'monthly utilities allowance','2017-08-22 18:53:58',0,'2017-08-22 18:50:08',0),(24,0,15,0,1,2,0,3,-175.0000,'2017-02-01 00:00:00','2017-02-01 00:00:00',0,0,0,'',4,4,'Reversal of ASM00000015','2017-08-22 18:52:07',0,'2017-08-22 18:52:07',0),(25,1,0,0,1,2,0,3,4000.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',2,2,'','2017-10-17 19:07:35',0,'2017-09-01 00:00:08',0),(26,18,0,0,1,2,0,3,350.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',4,2,'monthly utilities allowance','2017-10-17 19:07:35',0,'2017-09-01 00:00:08',0),(27,0,0,0,1,1,0,1,4000.0000,'2017-03-01 00:00:00','2018-03-01 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(28,27,0,0,1,1,0,1,4000.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(29,27,0,0,1,1,0,1,4000.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(30,27,0,0,1,1,0,1,4000.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(31,27,0,0,1,1,0,1,4000.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(32,27,0,0,1,1,0,1,4000.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(33,27,0,0,1,1,0,1,4000.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(34,27,0,0,1,1,0,1,4000.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(35,0,0,0,1,3,0,2,4150.0000,'2017-01-01 00:00:00','2018-01-01 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(36,35,0,0,1,3,0,2,4150.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(37,35,0,0,1,3,0,2,4150.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(38,35,0,0,1,3,0,2,4150.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(39,35,0,0,1,3,0,2,4150.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(40,35,0,0,1,3,0,2,4150.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(41,35,0,0,1,3,0,2,4150.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(42,35,0,0,1,3,0,2,4150.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(43,35,0,0,1,3,0,2,4150.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(44,35,0,0,1,3,0,2,4150.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',2,0,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(45,0,0,0,1,1,0,1,200.0000,'2017-09-14 00:00:00','2017-09-14 00:00:00',0,0,0,'',3,2,'','2017-09-15 05:44:13',0,'2017-09-15 05:42:55',0),(46,0,0,0,1,1,0,1,3500.0000,'2014-03-01 00:00:00','2014-03-01 00:00:00',0,0,0,'',24,0,'','2017-09-15 23:30:54',0,'2017-09-15 23:30:54',0),(47,0,0,0,1,1,0,1,3500.0000,'2014-03-02 00:00:00','2014-03-02 00:00:00',0,0,0,'',24,0,'','2017-09-15 23:41:23',0,'2017-09-15 23:41:23',0),(48,0,0,0,1,1,0,1,3000.0000,'2017-01-01 00:00:00','2017-01-01 00:00:00',0,0,0,'',24,0,'','2017-09-16 00:27:23',0,'2017-09-16 00:27:23',0),(49,0,0,0,1,0,0,8,50.0000,'2017-09-18 00:00:00','2017-09-18 00:00:00',0,0,0,'',11,2,'','2017-09-26 05:29:10',0,'2017-09-18 18:07:30',0),(50,0,0,0,1,0,0,9,50.0000,'2017-09-18 00:00:00','2017-09-18 00:00:00',0,0,0,'',11,2,'','2017-09-26 05:33:06',0,'2017-09-18 18:08:54',0),(51,0,0,0,1,0,0,10,50.0000,'2017-09-18 00:00:00','2017-09-18 00:00:00',0,0,0,'',11,2,'','2017-09-26 05:33:10',0,'2017-09-18 18:25:13',0),(52,0,0,0,1,0,0,11,50.0000,'2017-09-18 00:00:00','2017-09-18 00:00:00',0,0,0,'',11,2,'','2017-09-26 05:33:14',0,'2017-09-18 18:25:58',0),(53,0,0,0,1,0,0,12,50.0000,'2017-09-18 00:00:00','2017-09-18 00:00:00',0,0,0,'',11,4,'Reversed by ASM00000055','2017-09-25 20:51:37',0,'2017-09-18 18:26:24',0),(54,0,0,0,1,0,0,12,75.0000,'2017-09-03 00:00:00','2017-09-03 00:00:00',0,0,0,'',23,2,'Parking Ticket','2017-09-26 05:33:18',0,'2017-09-25 20:51:01',0),(55,0,53,0,1,0,0,12,-50.0000,'2017-09-18 00:00:00','2017-09-18 00:00:00',0,0,0,'',11,4,'Reversal of ASM00000053','2017-09-25 20:51:37',0,'2017-09-25 20:51:37',0),(56,0,0,0,1,0,0,14,100.0000,'2017-09-25 00:00:00','2017-09-25 00:00:00',0,0,0,'',36,4,'Reversed by ASM00000062','2017-10-03 00:47:03',0,'2017-09-25 21:23:20',0),(57,0,0,0,1,0,0,14,5000.0000,'2017-08-02 00:00:00','2017-08-02 00:00:00',0,0,0,'',36,0,'Wants the next unit type X that becomes available','2017-09-25 21:25:30',0,'2017-09-25 21:25:30',0),(58,1,0,0,1,2,0,3,4000.0000,'2017-10-01 00:00:00','2017-10-02 00:00:00',6,4,0,'',2,0,'','2017-10-01 20:55:22',0,'2017-10-01 20:55:22',0),(59,18,0,0,1,2,0,3,350.0000,'2017-10-01 00:00:00','2017-10-02 00:00:00',6,4,0,'',4,0,'monthly utilities allowance','2017-10-01 20:55:22',0,'2017-10-01 20:55:22',0),(60,27,0,0,1,1,0,1,4000.0000,'2017-10-01 00:00:00','2017-10-02 00:00:00',6,4,0,'',2,0,'','2017-10-01 20:55:22',0,'2017-10-01 20:55:22',0),(61,35,0,0,1,3,0,2,4150.0000,'2017-10-01 00:00:00','2017-10-02 00:00:00',6,4,0,'',2,0,'','2017-10-01 20:55:22',0,'2017-10-01 20:55:22',0),(62,0,56,0,1,0,0,14,-100.0000,'2017-09-25 00:00:00','2017-09-25 00:00:00',0,0,0,'',36,4,'Reversal of ASM00000056','2017-10-03 00:47:03',0,'2017-10-03 00:47:03',0),(63,0,0,0,1,5,0,18,10000.0000,'2017-10-01 00:00:00','2017-10-01 00:00:00',0,0,0,'',24,0,'','2017-10-04 03:46:11',0,'2017-10-04 03:36:50',0),(64,0,0,0,1,5,0,4,2258.0600,'2017-10-01 00:00:00','2017-10-01 00:00:00',0,0,0,'',2,0,'Prorate: 14 of 31 days','2017-10-05 03:43:20',0,'2017-10-05 03:43:20',0);
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
INSERT INTO `Business` VALUES (1,'REX','JGM First, LLC',6,4,4,'2017-08-16 05:49:58',0,'2017-08-16 05:49:58',0,0);
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
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `CARID` bigint(20) NOT NULL AUTO_INCREMENT,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`CARID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `CustomAttrRef`
--

LOCK TABLES `CustomAttrRef` WRITE;
/*!40000 ALTER TABLE `CustomAttrRef` DISABLE KEYS */;
INSERT INTO `CustomAttrRef` VALUES (5,1,1,1,'2017-09-15 07:57:55',0,1,'2018-01-01 10:09:57',0),(5,1,2,2,'2017-09-15 07:57:55',0,2,'2018-01-01 10:09:57',0),(5,1,3,3,'2017-09-15 07:57:55',0,3,'2018-01-01 10:09:57',0),(5,1,4,4,'2017-09-15 07:57:55',0,4,'2018-01-01 10:09:57',0),(5,1,5,5,'2017-09-15 07:57:55',0,5,'2018-01-01 10:09:57',0),(5,1,6,6,'2017-09-15 07:57:55',0,6,'2018-01-01 10:09:57',0);
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
-- Table structure for table `FlowPart`
--

DROP TABLE IF EXISTS `FlowPart`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `FlowPart` (
  `FlowPartID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `Flow` varchar(50) NOT NULL DEFAULT '',
  `FlowID` varchar(50) NOT NULL DEFAULT '',
  `PartType` smallint(6) NOT NULL DEFAULT '0',
  `Data` json DEFAULT NULL,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`FlowPartID`),
  UNIQUE KEY `FlowPartUnique` (`FlowPartID`,`BID`,`FlowID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `FlowPart`
--

LOCK TABLES `FlowPart` WRITE;
/*!40000 ALTER TABLE `FlowPart` DISABLE KEYS */;
/*!40000 ALTER TABLE `FlowPart` ENABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=109 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Journal`
--

LOCK TABLES `Journal` WRITE;
/*!40000 ALTER TABLE `Journal` DISABLE KEYS */;
INSERT INTO `Journal` VALUES (1,1,'2016-10-01 00:00:00',4000.0000,1,2,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(2,1,'2016-11-01 00:00:00',4000.0000,1,3,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(3,1,'2016-12-01 00:00:00',4000.0000,1,4,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(4,1,'2017-01-01 00:00:00',4000.0000,1,5,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(5,1,'2017-02-01 00:00:00',4000.0000,1,6,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(6,1,'2017-03-01 00:00:00',4000.0000,1,7,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(7,1,'2017-04-01 00:00:00',4000.0000,1,8,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(8,1,'2017-05-01 00:00:00',4000.0000,1,9,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(9,1,'2017-06-01 00:00:00',4000.0000,1,10,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(10,1,'2017-07-01 00:00:00',4000.0000,1,11,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(11,1,'2017-08-01 00:00:00',4000.0000,1,12,'','2017-08-22 18:24:26',0,'2017-08-22 18:24:26',0),(12,1,'2017-01-01 00:00:00',628.4500,1,13,'','2017-08-22 18:27:44',0,'2017-08-22 18:27:44',0),(13,1,'2017-02-01 00:00:00',175.0000,1,14,'','2017-08-22 18:28:39',0,'2017-08-22 18:28:39',0),(14,1,'2016-10-03 00:00:00',4000.0000,2,1,'','2017-08-22 18:29:52',0,'2017-08-22 18:29:52',0),(15,1,'2016-11-11 00:00:00',12000.0000,2,2,'','2017-08-22 18:31:17',0,'2017-08-22 18:31:17',0),(16,1,'2016-11-11 00:00:00',15.0000,3,1,'','2017-08-22 18:32:26',0,'2017-08-22 18:32:26',0),(17,1,'2016-10-03 00:00:00',4000.0000,2,1,'','2017-08-22 18:40:14',0,'2017-08-22 18:40:14',0),(18,1,'2016-11-11 00:00:00',4000.0000,2,2,'','2017-08-22 18:40:14',0,'2017-08-22 18:40:14',0),(19,1,'2016-12-03 00:00:00',4000.0000,2,2,'','2017-09-21 21:30:14',0,'2017-08-22 18:40:14',0),(20,1,'2017-01-03 00:00:00',4000.0000,2,2,'','2017-09-21 21:30:33',0,'2017-08-22 18:40:14',0),(21,1,'2017-02-03 00:00:00',628.4500,2,3,'','2017-08-22 18:43:32',0,'2017-08-22 18:43:32',0),(22,1,'2017-02-13 00:00:00',8350.0000,2,4,'','2017-08-22 18:44:27',0,'2017-08-22 18:44:27',0),(23,1,'2017-05-12 00:00:00',13131.7900,2,5,'','2017-08-22 18:45:12',0,'2017-08-22 18:45:12',0),(24,1,'2017-08-15 00:00:00',13050.0000,2,6,'','2017-08-22 18:45:41',0,'2017-08-22 18:45:41',0),(25,1,'2017-02-01 00:00:00',175.0000,1,15,'','2017-08-22 18:47:56',0,'2017-08-22 18:47:56',0),(26,1,'2017-03-01 00:00:00',175.0000,1,16,'','2017-08-22 18:48:32',0,'2017-08-22 18:48:32',0),(27,1,'2017-04-01 00:00:00',81.7900,1,17,'','2017-08-22 18:49:22',0,'2017-08-22 18:49:22',0),(28,1,'2017-04-01 00:00:00',350.0000,1,19,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(29,1,'2017-05-01 00:00:00',350.0000,1,20,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(30,1,'2017-06-01 00:00:00',350.0000,1,21,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(31,1,'2017-07-01 00:00:00',350.0000,1,22,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(32,1,'2017-08-01 00:00:00',350.0000,1,23,'','2017-08-22 18:50:08',0,'2017-08-22 18:50:08',0),(33,1,'2017-02-01 00:00:00',-175.0000,1,24,'','2017-08-22 18:52:07',0,'2017-08-22 18:52:07',0),(34,1,'2017-02-03 00:00:00',628.4500,2,3,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(35,1,'2017-02-13 00:00:00',4000.0000,2,4,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(36,1,'2017-02-13 00:00:00',175.0000,2,4,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(37,1,'2017-02-13 00:00:00',4000.0000,2,4,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(38,1,'2017-02-13 00:00:00',175.0000,2,4,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(39,1,'2017-05-12 00:00:00',4000.0000,2,5,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(40,1,'2017-05-12 00:00:00',81.7900,2,5,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(41,1,'2017-05-12 00:00:00',350.0000,2,5,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(42,1,'2017-05-12 00:00:00',4000.0000,2,5,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(43,1,'2017-05-12 00:00:00',350.0000,2,5,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(44,1,'2017-05-12 00:00:00',4000.0000,2,5,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(45,1,'2017-05-12 00:00:00',350.0000,2,5,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(46,1,'2017-08-15 00:00:00',4000.0000,2,6,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(47,1,'2017-08-15 00:00:00',350.0000,2,6,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(48,1,'2017-08-15 00:00:00',4000.0000,2,6,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(49,1,'2017-08-15 00:00:00',350.0000,2,6,'','2017-08-22 18:53:58',0,'2017-08-22 18:53:58',0),(50,1,'2017-09-01 00:00:00',4000.0000,1,25,'','2017-09-01 00:00:08',0,'2017-09-01 00:00:08',0),(51,1,'2017-09-01 00:00:00',350.0000,1,26,'','2017-09-01 00:00:08',0,'2017-09-01 00:00:08',0),(52,1,'2017-03-01 00:00:00',4000.0000,1,28,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(53,1,'2017-04-01 00:00:00',4000.0000,1,29,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(54,1,'2017-05-01 00:00:00',4000.0000,1,30,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(55,1,'2017-06-01 00:00:00',4000.0000,1,31,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(56,1,'2017-07-01 00:00:00',4000.0000,1,32,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(57,1,'2017-08-01 00:00:00',4000.0000,1,33,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(58,1,'2017-09-01 00:00:00',4000.0000,1,34,'','2017-09-15 04:14:32',0,'2017-09-15 04:14:32',0),(59,1,'2017-01-01 00:00:00',4150.0000,1,36,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(60,1,'2017-02-01 00:00:00',4150.0000,1,37,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(61,1,'2017-03-01 00:00:00',4150.0000,1,38,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(62,1,'2017-04-01 00:00:00',4150.0000,1,39,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(63,1,'2017-05-01 00:00:00',4150.0000,1,40,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(64,1,'2017-06-01 00:00:00',4150.0000,1,41,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(65,1,'2017-07-01 00:00:00',4150.0000,1,42,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(66,1,'2017-08-01 00:00:00',4150.0000,1,43,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(67,1,'2017-09-01 00:00:00',4150.0000,1,44,'','2017-09-15 04:16:34',0,'2017-09-15 04:16:34',0),(68,1,'2017-09-14 00:00:00',200.0000,1,45,'','2017-09-15 05:42:55',0,'2017-09-15 05:42:55',0),(69,1,'2017-09-14 00:00:00',125.0000,2,7,'','2017-09-15 05:43:22',0,'2017-09-15 05:43:22',0),(70,1,'2017-09-14 00:00:00',75.0000,2,8,'','2017-09-15 05:43:38',0,'2017-09-15 05:43:38',0),(71,1,'2017-09-14 00:00:00',125.0000,2,7,'','2017-09-15 05:44:13',0,'2017-09-15 05:44:13',0),(72,1,'2017-09-14 00:00:00',75.0000,2,8,'','2017-09-15 05:44:13',0,'2017-09-15 05:44:13',0),(73,1,'2014-03-01 00:00:00',3500.0000,1,46,'','2017-09-15 23:30:54',0,'2017-09-15 23:30:54',0),(74,1,'2014-03-02 00:00:00',3500.0000,1,47,'','2017-09-15 23:41:23',0,'2017-09-15 23:41:23',0),(75,1,'2017-01-01 00:00:00',3000.0000,1,48,'','2017-09-16 00:27:23',0,'2017-09-16 00:27:23',0),(76,1,'2017-09-18 00:00:00',50.0000,1,49,'','2017-09-18 18:07:30',0,'2017-09-18 18:07:30',0),(77,1,'2017-09-18 00:00:00',50.0000,1,50,'','2017-09-18 18:08:54',0,'2017-09-18 18:08:54',0),(78,1,'2017-09-18 00:00:00',50.0000,1,51,'','2017-09-18 18:25:13',0,'2017-09-18 18:25:13',0),(79,1,'2017-09-18 00:00:00',50.0000,1,52,'','2017-09-18 18:25:58',0,'2017-09-18 18:25:58',0),(80,1,'2017-09-18 00:00:00',50.0000,1,53,'','2017-09-18 18:26:24',0,'2017-09-18 18:26:24',0),(81,1,'2017-09-18 00:00:00',3525.0000,2,9,'','2017-09-18 19:49:58',0,'2017-09-18 19:49:58',0),(82,1,'2017-09-03 00:00:00',75.0000,1,54,'','2017-09-25 20:51:01',0,'2017-09-25 20:51:01',0),(83,1,'2017-09-18 00:00:00',-50.0000,1,55,'','2017-09-25 20:51:37',0,'2017-09-25 20:51:37',0),(84,1,'2017-09-25 00:00:00',100.0000,1,56,'','2017-09-25 21:23:20',0,'2017-09-25 21:23:20',0),(85,1,'2017-08-02 00:00:00',5000.0000,1,57,'','2017-09-25 21:25:30',0,'2017-09-25 21:25:30',0),(86,1,'2017-09-25 00:00:00',50.0000,2,10,'','2017-09-26 05:28:20',0,'2017-09-26 05:28:20',0),(87,1,'2017-09-25 00:00:00',50.0000,2,10,'','2017-09-26 05:29:10',0,'2017-09-26 05:29:10',0),(88,1,'2017-09-25 00:00:00',50.0000,2,11,'','2017-09-26 05:30:24',0,'2017-09-26 05:30:24',0),(89,1,'2017-09-25 00:00:00',50.0000,2,12,'','2017-09-26 05:31:32',0,'2017-09-26 05:31:32',0),(90,1,'2017-09-25 00:00:00',50.0000,2,13,'','2017-09-26 05:32:00',0,'2017-09-26 05:32:00',0),(91,1,'2017-09-25 00:00:00',75.0000,2,14,'','2017-09-26 05:32:57',0,'2017-09-26 05:32:57',0),(92,1,'2017-09-25 00:00:00',50.0000,2,11,'','2017-09-26 05:33:06',0,'2017-09-26 05:33:06',0),(93,1,'2017-09-25 00:00:00',50.0000,2,12,'','2017-09-26 05:33:10',0,'2017-09-26 05:33:10',0),(94,1,'2017-09-25 00:00:00',50.0000,2,13,'','2017-09-26 05:33:14',0,'2017-09-26 05:33:14',0),(95,1,'2017-09-25 00:00:00',75.0000,2,14,'','2017-09-26 05:33:18',0,'2017-09-26 05:33:18',0),(96,1,'2017-09-27 00:00:00',100.0000,2,15,'','2017-09-27 18:26:00',0,'2017-09-27 18:26:00',0),(97,1,'2017-09-18 00:00:00',-3525.0000,2,16,'','2017-09-29 18:12:25',0,'2017-09-29 18:12:25',0),(98,1,'2017-09-18 00:00:00',3525.0000,2,17,'','2017-09-29 18:12:25',0,'2017-09-29 18:12:25',0),(99,1,'2017-10-01 00:00:00',4000.0000,1,58,'','2017-10-01 20:55:22',0,'2017-10-01 20:55:22',0),(100,1,'2017-10-01 00:00:00',350.0000,1,59,'','2017-10-01 20:55:22',0,'2017-10-01 20:55:22',0),(101,1,'2017-10-01 00:00:00',4000.0000,1,60,'','2017-10-01 20:55:22',0,'2017-10-01 20:55:22',0),(102,1,'2017-10-01 00:00:00',4150.0000,1,61,'','2017-10-01 20:55:22',0,'2017-10-01 20:55:22',0),(103,1,'2017-09-01 00:00:00',50.0000,2,18,'','2017-10-01 21:04:24',0,'2017-10-01 21:04:24',0),(104,1,'2017-09-25 00:00:00',0.0000,1,62,'','2017-10-03 00:47:03',0,'2017-10-03 00:47:03',0),(105,1,'2017-10-15 00:00:00',10000.0000,1,63,'','2017-10-04 03:36:50',0,'2017-10-04 03:36:50',0),(106,1,'2017-10-01 00:00:00',2258.0600,1,64,'','2017-10-05 03:43:20',0,'2017-10-05 03:43:20',0),(107,1,'2017-10-17 00:00:00',4000.0000,2,6,'','2017-10-17 19:07:35',0,'2017-10-17 19:07:35',0),(108,1,'2017-10-17 00:00:00',350.0000,2,6,'','2017-10-17 19:07:35',0,'2017-10-17 19:07:35',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=109 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `JournalAllocation`
--

LOCK TABLES `JournalAllocation` WRITE;
/*!40000 ALTER TABLE `JournalAllocation` DISABLE KEYS */;
INSERT INTO `JournalAllocation` VALUES (1,1,1,2,3,0,0,4000.0000,2,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0,'2018-01-01 10:09:57',0),(2,1,2,2,3,0,0,4000.0000,3,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0,'2018-01-01 10:09:57',0),(3,1,3,2,3,0,0,4000.0000,4,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0,'2018-01-01 10:09:57',0),(4,1,4,2,3,0,0,4000.0000,5,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0,'2018-01-01 10:09:57',0),(5,1,5,2,3,0,0,4000.0000,6,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0,'2018-01-01 10:09:57',0),(6,1,6,2,3,0,0,4000.0000,7,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0,'2018-01-01 10:09:57',0),(7,1,7,2,3,0,0,4000.0000,8,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0,'2018-01-01 10:09:57',0),(8,1,8,2,3,0,0,4000.0000,9,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0,'2018-01-01 10:09:57',0),(9,1,9,2,3,0,0,4000.0000,10,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0,'2018-01-01 10:09:57',0),(10,1,10,2,3,0,0,4000.0000,11,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0,'2018-01-01 10:09:57',0),(11,1,11,2,3,0,0,4000.0000,12,0,'d 12001 4000.00, c 41001 4000.00','2017-08-22 18:24:26',0,'2018-01-01 10:09:57',0),(12,1,12,2,3,0,0,628.4500,13,0,'d 12001 628.45, c 41301 628.45','2017-08-22 18:27:44',0,'2018-01-01 10:09:57',0),(13,1,13,2,3,0,0,175.0000,14,0,'d 12001 175.00, c 41301 175.00','2017-08-22 18:28:39',0,'2018-01-01 10:09:57',0),(14,1,14,0,0,6,0,4000.0000,0,0,'d 10999 _, c 12999 _','2017-08-22 18:29:52',0,'2018-01-01 10:09:57',0),(15,1,15,0,0,6,0,12000.0000,0,0,'d 10999 _, c 12999 _','2017-08-22 18:31:17',0,'2018-01-01 10:09:57',0),(16,1,16,2,3,0,0,15.0000,0,1,'d 50003 15.00, c 12001 15.00','2017-08-22 18:32:26',0,'2018-01-01 10:09:57',0),(17,1,17,2,3,6,1,4000.0000,2,0,'ASM(2) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:40:14',0,'2018-01-01 10:09:57',0),(18,1,18,2,3,6,2,4000.0000,3,0,'ASM(3) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:40:14',0,'2018-01-01 10:09:57',0),(19,1,19,2,3,6,2,4000.0000,4,0,'ASM(4) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:40:14',0,'2018-01-01 10:09:57',0),(20,1,20,2,3,6,2,4000.0000,5,0,'ASM(5) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:40:14',0,'2018-01-01 10:09:57',0),(21,1,21,0,0,6,0,628.4500,0,0,'d 10999 _, c 12999 _','2017-08-22 18:43:32',0,'2018-01-01 10:09:57',0),(22,1,22,0,0,6,0,8350.0000,0,0,'d 10999 _, c 12999 _','2017-08-22 18:44:27',0,'2018-01-01 10:09:57',0),(23,1,23,0,0,6,0,13131.7900,0,0,'d 10999 _, c 12999 _','2017-08-22 18:45:12',0,'2018-01-01 10:09:57',0),(24,1,24,0,0,6,0,13050.0000,0,0,'d 10999 _, c 12999 _','2017-08-22 18:45:41',0,'2018-01-01 10:09:57',0),(25,1,25,2,3,0,0,175.0000,15,0,'d 12001 175.00, c 41301 175.00','2017-08-22 18:47:56',0,'2018-01-01 10:09:57',0),(26,1,26,2,3,0,0,175.0000,16,0,'d 12001 175.00, c 41301 175.00','2017-08-22 18:48:32',0,'2018-01-01 10:09:57',0),(27,1,27,2,3,0,0,81.7900,17,0,'d 12001 81.79, c 41301 81.79','2017-08-22 18:49:22',0,'2018-01-01 10:09:57',0),(28,1,28,2,3,0,0,350.0000,19,0,'d 12001 350.00, c 41301 350.00','2017-08-22 18:50:08',0,'2018-01-01 10:09:57',0),(29,1,29,2,3,0,0,350.0000,20,0,'d 12001 350.00, c 41301 350.00','2017-08-22 18:50:08',0,'2018-01-01 10:09:57',0),(30,1,30,2,3,0,0,350.0000,21,0,'d 12001 350.00, c 41301 350.00','2017-08-22 18:50:08',0,'2018-01-01 10:09:57',0),(31,1,31,2,3,0,0,350.0000,22,0,'d 12001 350.00, c 41301 350.00','2017-08-22 18:50:08',0,'2018-01-01 10:09:57',0),(32,1,32,2,3,0,0,350.0000,23,0,'d 12001 350.00, c 41301 350.00','2017-08-22 18:50:08',0,'2018-01-01 10:09:57',0),(33,1,33,2,3,0,0,-175.0000,24,0,'d 12001 -175.00, c 41301 -175.00','2017-08-22 18:52:07',0,'2018-01-01 10:09:57',0),(34,1,34,2,3,6,3,628.4500,13,0,'ASM(13) d 12999 628.45,c 12001 628.45','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(35,1,35,2,3,6,4,4000.0000,6,0,'ASM(6) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(36,1,36,2,3,6,4,175.0000,14,0,'ASM(14) d 12999 175.00,c 12001 175.00','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(37,1,37,2,3,6,4,4000.0000,7,0,'ASM(7) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(38,1,38,2,3,6,4,175.0000,16,0,'ASM(16) d 12999 175.00,c 12001 175.00','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(39,1,39,2,3,6,5,4000.0000,8,0,'ASM(8) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(40,1,40,2,3,6,5,81.7900,17,0,'ASM(17) d 12999 81.79,c 12001 81.79','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(41,1,41,2,3,6,5,350.0000,19,0,'ASM(19) d 12999 350.00,c 12001 350.00','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(42,1,42,2,3,6,5,4000.0000,9,0,'ASM(9) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(43,1,43,2,3,6,5,350.0000,20,0,'ASM(20) d 12999 350.00,c 12001 350.00','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(44,1,44,2,3,6,5,4000.0000,10,0,'ASM(10) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(45,1,45,2,3,6,5,350.0000,21,0,'ASM(21) d 12999 350.00,c 12001 350.00','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(46,1,46,2,3,6,6,4000.0000,11,0,'ASM(11) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(47,1,47,2,3,6,6,350.0000,22,0,'ASM(22) d 12999 350.00,c 12001 350.00','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(48,1,48,2,3,6,6,4000.0000,12,0,'ASM(12) d 12999 4000.00,c 12001 4000.00','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(49,1,49,2,3,6,6,350.0000,23,0,'ASM(23) d 12999 350.00,c 12001 350.00','2017-08-22 18:53:58',0,'2018-01-01 10:09:57',0),(50,1,50,2,3,0,0,4000.0000,25,0,'d 12001 4000.00, c 41001 4000.00','2017-09-01 00:00:08',0,'2018-01-01 10:09:57',0),(51,1,51,2,3,0,0,350.0000,26,0,'d 12001 350.00, c 41301 350.00','2017-09-01 00:00:08',0,'2018-01-01 10:09:57',0),(52,1,52,1,1,0,0,4000.0000,28,0,'d 12001 4000.00, c 41001 4000.00','2017-09-15 04:14:32',0,'2018-01-01 10:09:57',0),(53,1,53,1,1,0,0,4000.0000,29,0,'d 12001 4000.00, c 41001 4000.00','2017-09-15 04:14:32',0,'2018-01-01 10:09:57',0),(54,1,54,1,1,0,0,4000.0000,30,0,'d 12001 4000.00, c 41001 4000.00','2017-09-15 04:14:32',0,'2018-01-01 10:09:57',0),(55,1,55,1,1,0,0,4000.0000,31,0,'d 12001 4000.00, c 41001 4000.00','2017-09-15 04:14:32',0,'2018-01-01 10:09:57',0),(56,1,56,1,1,0,0,4000.0000,32,0,'d 12001 4000.00, c 41001 4000.00','2017-09-15 04:14:32',0,'2018-01-01 10:09:57',0),(57,1,57,1,1,0,0,4000.0000,33,0,'d 12001 4000.00, c 41001 4000.00','2017-09-15 04:14:32',0,'2018-01-01 10:09:57',0),(58,1,58,1,1,0,0,4000.0000,34,0,'d 12001 4000.00, c 41001 4000.00','2017-09-15 04:14:32',0,'2018-01-01 10:09:57',0),(59,1,59,3,2,0,0,4150.0000,36,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0,'2018-01-01 10:09:57',0),(60,1,60,3,2,0,0,4150.0000,37,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0,'2018-01-01 10:09:57',0),(61,1,61,3,2,0,0,4150.0000,38,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0,'2018-01-01 10:09:57',0),(62,1,62,3,2,0,0,4150.0000,39,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0,'2018-01-01 10:09:57',0),(63,1,63,3,2,0,0,4150.0000,40,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0,'2018-01-01 10:09:57',0),(64,1,64,3,2,0,0,4150.0000,41,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0,'2018-01-01 10:09:57',0),(65,1,65,3,2,0,0,4150.0000,42,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0,'2018-01-01 10:09:57',0),(66,1,66,3,2,0,0,4150.0000,43,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0,'2018-01-01 10:09:57',0),(67,1,67,3,2,0,0,4150.0000,44,0,'d 12001 4150.00, c 41001 4150.00','2017-09-15 04:16:34',0,'2018-01-01 10:09:57',0),(68,1,68,1,1,0,0,200.0000,45,0,'d 12001 200.00, c 41302 200.00','2017-09-15 05:42:55',0,'2018-01-01 10:09:57',0),(69,1,69,0,0,1,0,125.0000,0,0,'d 10999 _, c 12999 _','2017-09-15 05:43:22',0,'2018-01-01 10:09:57',0),(70,1,70,0,0,1,0,75.0000,0,0,'d 10999 _, c 12999 _','2017-09-15 05:43:38',0,'2018-01-01 10:09:57',0),(71,1,71,1,1,1,7,125.0000,45,0,'ASM(45) d 12999 125.00,c 12001 125.00','2017-09-15 05:44:13',0,'2018-01-01 10:09:57',0),(72,1,72,1,1,1,8,75.0000,45,0,'ASM(45) d 12999 75.00,c 12001 75.00','2017-09-15 05:44:13',0,'2018-01-01 10:09:57',0),(73,1,73,1,1,0,0,3500.0000,46,0,'d 12001 3500.00, c 30000 3500.00','2017-09-15 23:30:54',0,'2018-01-01 10:09:57',0),(74,1,74,1,1,0,0,3500.0000,47,0,'d 12001 3500.00, c 30000 3500.00','2017-09-15 23:41:23',0,'2018-01-01 10:09:57',0),(75,1,75,1,1,0,0,3000.0000,48,0,'d 12001 3000.00, c 30000 3000.00','2017-09-16 00:27:23',0,'2018-01-01 10:09:57',0),(76,1,76,0,8,0,0,50.0000,49,0,'d 12001 50.00, c 41401 50.00','2017-09-18 18:07:30',0,'2018-01-01 10:09:57',0),(77,1,77,0,9,0,0,50.0000,50,0,'d 12001 50.00, c 41401 50.00','2017-09-18 18:08:54',0,'2018-01-01 10:09:57',0),(78,1,78,0,10,0,0,50.0000,51,0,'d 12001 50.00, c 41401 50.00','2017-09-18 18:25:13',0,'2018-01-01 10:09:57',0),(79,1,79,0,11,0,0,50.0000,52,0,'d 12001 50.00, c 41401 50.00','2017-09-18 18:25:58',0,'2018-01-01 10:09:57',0),(80,1,80,0,12,0,0,50.0000,53,0,'d 12001 50.00, c 41401 50.00','2017-09-18 18:26:24',0,'2018-01-01 10:09:57',0),(81,1,81,0,0,14,0,3525.0000,0,0,'d 11000 _, c 42300 _','2017-09-18 19:49:58',0,'2018-01-01 10:09:57',0),(82,1,82,0,12,0,0,75.0000,54,0,'d 12001 75.00, c 41499 75.00','2017-09-25 20:51:01',0,'2018-01-01 10:09:57',0),(83,1,83,0,12,0,0,-50.0000,55,0,'d 12001 -50.00, c 41401 -50.00','2017-09-25 20:51:37',0,'2018-01-01 10:09:57',0),(84,1,84,0,14,0,0,100.0000,56,0,'d 12001 100.00, c 30000 100.00','2017-09-25 21:23:20',0,'2018-01-01 10:09:57',0),(85,1,85,0,14,0,0,5000.0000,57,0,'d 12001 5000.00, c 30000 5000.00','2017-09-25 21:25:30',0,'2018-01-01 10:09:57',0),(86,1,86,0,0,9,0,50.0000,0,0,'d 10999 _, c 12999 _','2017-09-26 05:28:20',0,'2018-01-01 10:09:57',0),(87,1,87,0,8,9,10,50.0000,49,0,'ASM(49) d 12999 50.00,c 12001 50.00','2017-09-26 05:29:10',0,'2018-01-01 10:09:57',0),(88,1,88,0,0,10,0,50.0000,0,0,'d 10999 _, c 12999 _','2017-09-26 05:30:24',0,'2018-01-01 10:09:57',0),(89,1,89,0,0,11,0,50.0000,0,0,'d 10999 _, c 12999 _','2017-09-26 05:31:32',0,'2018-01-01 10:09:57',0),(90,1,90,0,0,12,0,50.0000,0,0,'d 10999 _, c 12999 _','2017-09-26 05:32:00',0,'2018-01-01 10:09:57',0),(91,1,91,0,0,13,0,75.0000,0,0,'d 10999 _, c 12999 _','2017-09-26 05:32:57',0,'2018-01-01 10:09:57',0),(92,1,92,0,9,10,11,50.0000,50,0,'ASM(50) d 12999 50.00,c 12001 50.00','2017-09-26 05:33:06',0,'2018-01-01 10:09:57',0),(93,1,93,0,10,11,12,50.0000,51,0,'ASM(51) d 12999 50.00,c 12001 50.00','2017-09-26 05:33:10',0,'2018-01-01 10:09:57',0),(94,1,94,0,11,12,13,50.0000,52,0,'ASM(52) d 12999 50.00,c 12001 50.00','2017-09-26 05:33:14',0,'2018-01-01 10:09:57',0),(95,1,95,0,12,13,14,75.0000,54,0,'ASM(54) d 12999 75.00,c 12001 75.00','2017-09-26 05:33:18',0,'2018-01-01 10:09:57',0),(96,1,96,0,0,15,0,100.0000,0,0,'d 10999 _, c 30001 _','2017-09-27 18:26:00',0,'2018-01-01 10:09:57',0),(97,1,97,0,0,14,0,-3525.0000,0,0,'d 11000 _, c 42300 _','2017-09-29 18:12:25',0,'2018-01-01 10:09:57',0),(98,1,98,0,0,14,0,3525.0000,0,0,'d 11000 _, c 42300 _','2017-09-29 18:12:25',0,'2018-01-01 10:09:57',0),(99,1,99,2,3,0,0,4000.0000,58,0,'d 12001 4000.00, c 41001 4000.00','2017-10-01 20:55:22',0,'2018-01-01 10:09:57',0),(100,1,100,2,3,0,0,350.0000,59,0,'d 12001 350.00, c 41301 350.00','2017-10-01 20:55:22',0,'2018-01-01 10:09:57',0),(101,1,101,1,1,0,0,4000.0000,60,0,'d 12001 4000.00, c 41001 4000.00','2017-10-01 20:55:22',0,'2018-01-01 10:09:57',0),(102,1,102,3,2,0,0,4150.0000,61,0,'d 12001 4150.00, c 41001 4150.00','2017-10-01 20:55:22',0,'2018-01-01 10:09:57',0),(103,1,103,0,0,16,0,50.0000,0,0,'d 11000 _, c 41401 _','2017-10-01 21:04:24',0,'2018-01-01 10:09:57',0),(104,1,104,0,14,0,0,0.0000,62,0,'','2017-10-03 00:47:03',0,'2018-01-01 10:09:57',0),(105,1,105,0,18,0,0,10000.0000,63,0,'d 12001 10000.00, c 30000 10000.00','2017-10-04 03:36:50',0,'2018-01-01 10:09:57',0),(106,1,106,5,4,0,0,2258.0600,64,0,'d 12001 2258.06, c 41001 2258.06','2017-10-05 03:43:20',0,'2018-01-01 10:09:57',0),(107,1,107,2,3,6,6,4000.0000,25,0,'ASM(25) d 12999 4000.00,c 12001 4000.00','2017-10-17 19:07:35',0,'2018-01-01 10:09:57',0),(108,1,108,2,3,6,6,350.0000,26,0,'ASM(26) d 12999 350.00,c 12001 350.00','2017-10-17 19:07:35',0,'2018-01-01 10:09:57',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=379 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerEntry`
--

LOCK TABLES `LedgerEntry` WRITE;
/*!40000 ALTER TABLE `LedgerEntry` DISABLE KEYS */;
INSERT INTO `LedgerEntry` VALUES (1,1,1,1,9,2,1,0,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 18:39:27',0,'2017-11-30 18:39:27',0),(2,1,1,1,11,2,1,0,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 18:39:27',0,'2017-11-30 18:39:27',0),(3,1,2,2,9,4,3,0,'2016-07-01 00:00:00',8300.0000,'','2017-11-30 18:41:00',0,'2017-11-30 18:41:00',0),(4,1,2,2,11,4,3,0,'2016-07-01 00:00:00',-8300.0000,'','2017-11-30 18:41:00',0,'2017-11-30 18:41:00',0),(5,1,3,3,9,2,1,0,'2017-01-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(6,1,3,3,17,2,1,0,'2017-01-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(7,1,4,4,9,2,1,0,'2017-02-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(8,1,4,4,17,2,1,0,'2017-02-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(9,1,5,5,9,2,1,0,'2017-03-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(10,1,5,5,17,2,1,0,'2017-03-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(11,1,6,6,9,2,1,0,'2017-04-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(12,1,6,6,17,2,1,0,'2017-04-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(13,1,7,7,9,2,1,0,'2017-05-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(14,1,7,7,17,2,1,0,'2017-05-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(15,1,8,8,9,2,1,0,'2017-06-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(16,1,8,8,17,2,1,0,'2017-06-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(17,1,9,9,9,2,1,0,'2017-07-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(18,1,9,9,17,2,1,0,'2017-07-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(19,1,10,10,9,2,1,0,'2017-08-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(20,1,10,10,17,2,1,0,'2017-08-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(21,1,11,11,9,2,1,0,'2017-09-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(22,1,11,11,17,2,1,0,'2017-09-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(23,1,12,12,9,2,1,0,'2017-10-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(24,1,12,12,17,2,1,0,'2017-10-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(25,1,13,13,9,2,1,0,'2017-11-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(26,1,13,13,17,2,1,0,'2017-11-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(27,1,14,14,9,3,2,0,'2017-01-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(28,1,14,14,17,3,2,0,'2017-01-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(29,1,15,15,9,3,2,0,'2017-02-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(30,1,15,15,17,3,2,0,'2017-02-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(31,1,16,16,9,3,2,0,'2017-03-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(32,1,16,16,17,3,2,0,'2017-03-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(33,1,17,17,9,3,2,0,'2017-04-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(34,1,17,17,17,3,2,0,'2017-04-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(35,1,18,18,9,3,2,0,'2017-05-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(36,1,18,18,17,3,2,0,'2017-05-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(37,1,19,19,9,3,2,0,'2017-06-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(38,1,19,19,17,3,2,0,'2017-06-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(39,1,20,20,9,3,2,0,'2017-07-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(40,1,20,20,17,3,2,0,'2017-07-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(41,1,21,21,9,3,2,0,'2017-08-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(42,1,21,21,17,3,2,0,'2017-08-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(43,1,22,22,9,3,2,0,'2017-09-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(44,1,22,22,17,3,2,0,'2017-09-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(45,1,23,23,9,3,2,0,'2017-10-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(46,1,23,23,17,3,2,0,'2017-10-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(47,1,24,24,9,3,2,0,'2017-11-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(48,1,24,24,17,3,2,0,'2017-11-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(49,1,25,25,9,4,3,0,'2017-01-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(50,1,25,25,17,4,3,0,'2017-01-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(51,1,26,26,9,4,3,0,'2017-02-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(52,1,26,26,17,4,3,0,'2017-02-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(53,1,27,27,9,4,3,0,'2017-03-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(54,1,27,27,17,4,3,0,'2017-03-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(55,1,28,28,9,4,3,0,'2017-04-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(56,1,28,28,17,4,3,0,'2017-04-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(57,1,29,29,9,4,3,0,'2017-05-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(58,1,29,29,17,4,3,0,'2017-05-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(59,1,30,30,9,4,3,0,'2017-06-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(60,1,30,30,17,4,3,0,'2017-06-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(61,1,31,31,9,4,3,0,'2017-07-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(62,1,31,31,17,4,3,0,'2017-07-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(63,1,32,32,9,4,3,0,'2017-08-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(64,1,32,32,17,4,3,0,'2017-08-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(65,1,33,33,9,4,3,0,'2017-09-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(66,1,33,33,17,4,3,0,'2017-09-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(67,1,34,34,9,4,3,0,'2017-10-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(68,1,34,34,17,4,3,0,'2017-10-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(69,1,35,35,9,4,3,0,'2017-11-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(70,1,35,35,17,4,3,0,'2017-11-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(71,1,36,36,3,0,0,1,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 18:48:02',0,'2017-11-30 18:48:02',0),(72,1,36,36,6,0,0,1,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 18:48:02',0,'2017-11-30 18:48:02',0),(73,1,37,37,3,0,0,4,'2016-07-01 00:00:00',8300.0000,'','2017-11-30 18:49:17',0,'2017-11-30 18:49:17',0),(74,1,37,37,6,0,0,4,'2016-07-01 00:00:00',-8300.0000,'','2017-11-30 18:49:17',0,'2017-11-30 18:49:17',0),(75,1,38,38,6,0,0,4,'2016-07-01 00:00:00',8300.0000,'','2017-11-30 19:13:23',0,'2017-11-30 19:13:23',0),(76,1,38,38,10,0,0,4,'2016-07-01 00:00:00',-8300.0000,'','2017-11-30 19:13:23',0,'2017-11-30 19:13:23',0),(77,1,39,39,3,0,0,4,'2016-07-01 00:00:00',-8300.0000,'','2017-11-30 19:13:59',0,'2017-11-30 19:13:59',0),(78,1,39,39,6,0,0,4,'2016-07-01 00:00:00',8300.0000,'','2017-11-30 19:13:59',0,'2017-11-30 19:13:59',0),(79,1,40,40,3,0,0,4,'2016-07-01 00:00:00',8300.0000,'','2017-11-30 19:17:32',0,'2017-11-30 19:17:32',0),(80,1,40,40,6,0,0,4,'2016-07-01 00:00:00',-8300.0000,'','2017-11-30 19:17:32',0,'2017-11-30 19:17:32',0),(81,1,41,41,3,0,0,1,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 19:23:47',0,'2017-11-30 19:23:47',0),(82,1,41,41,6,0,0,1,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 19:23:47',0,'2017-11-30 19:23:47',0),(83,1,42,42,6,0,0,1,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 19:24:32',0,'2017-11-30 19:24:32',0),(84,1,42,42,10,0,0,1,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 19:24:32',0,'2017-11-30 19:24:32',0),(85,1,43,43,3,0,0,1,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 19:25:23',0,'2017-11-30 19:25:23',0),(86,1,43,43,6,0,0,1,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 19:25:23',0,'2017-11-30 19:25:23',0),(87,1,44,44,6,0,0,1,'2017-01-01 00:00:00',3750.0000,'','2017-11-30 19:44:52',0,'2017-11-30 19:44:52',0),(88,1,44,44,10,0,0,1,'2017-01-01 00:00:00',-3750.0000,'','2017-11-30 19:44:52',0,'2017-11-30 19:44:52',0),(89,1,45,45,10,2,1,1,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(90,1,45,45,9,2,1,1,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(91,1,46,46,10,2,1,1,'2017-01-01 00:00:00',3750.0000,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(92,1,46,46,9,2,1,1,'2017-01-01 00:00:00',-3750.0000,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(93,1,50,50,9,3,2,0,'2017-01-31 00:00:00',628.4500,'','2017-12-05 16:01:46',0,'2017-12-05 16:01:46',0),(94,1,50,50,37,3,2,0,'2017-01-31 00:00:00',-628.4500,'','2017-12-05 16:01:46',0,'2017-12-05 16:01:46',0),(95,1,51,51,9,3,2,0,'2017-02-28 00:00:00',175.0000,'','2017-12-05 16:02:25',0,'2017-12-05 16:02:25',0),(96,1,51,51,37,3,2,0,'2017-02-28 00:00:00',-175.0000,'','2017-12-05 16:02:25',0,'2017-12-05 16:02:25',0),(97,1,52,52,9,3,2,0,'2017-03-31 00:00:00',175.0000,'','2017-12-05 16:03:13',0,'2017-12-05 16:03:13',0),(98,1,52,52,37,3,2,0,'2017-03-31 00:00:00',-175.0000,'','2017-12-05 16:03:13',0,'2017-12-05 16:03:13',0),(99,1,53,53,9,3,2,0,'2017-04-15 00:00:00',81.7900,'','2017-12-05 16:03:41',0,'2017-12-05 16:03:41',0),(100,1,53,53,37,3,2,0,'2017-04-15 00:00:00',-81.7900,'','2017-12-05 16:03:41',0,'2017-12-05 16:03:41',0),(101,1,54,54,9,3,2,0,'2017-10-31 00:00:00',409.2800,'','2017-12-05 16:07:34',0,'2017-12-05 16:07:34',0),(102,1,54,54,37,3,2,0,'2017-10-31 00:00:00',-409.2800,'','2017-12-05 16:07:34',0,'2017-12-05 16:07:34',0),(103,1,55,55,6,0,0,1,'2017-01-01 00:00:00',3750.0000,'','2017-12-05 16:09:37',0,'2017-12-05 16:09:37',0),(104,1,55,55,10,0,0,1,'2017-01-01 00:00:00',-3750.0000,'','2017-12-05 16:09:37',0,'2017-12-05 16:09:37',0),(105,1,56,56,6,0,0,4,'2017-01-01 00:00:00',4150.0000,'','2017-12-05 16:10:06',0,'2017-12-05 16:10:06',0),(106,1,56,56,10,0,0,4,'2017-01-01 00:00:00',-4150.0000,'','2017-12-05 16:10:06',0,'2017-12-05 16:10:06',0),(107,1,57,57,6,0,0,3,'2017-02-01 00:00:00',8350.0000,'','2017-12-05 16:12:02',0,'2017-12-05 16:12:02',0),(108,1,57,57,10,0,0,3,'2017-02-01 00:00:00',-8350.0000,'','2017-12-05 16:12:02',0,'2017-12-05 16:12:02',0),(109,1,58,58,6,0,0,1,'2017-02-01 00:00:00',3750.0000,'','2017-12-05 16:12:32',0,'2017-12-05 16:12:32',0),(110,1,58,58,10,0,0,1,'2017-02-01 00:00:00',-3750.0000,'','2017-12-05 16:12:32',0,'2017-12-05 16:12:32',0),(111,1,59,59,6,0,0,4,'2017-02-01 00:00:00',4150.0000,'','2017-12-05 16:13:44',0,'2017-12-05 16:13:44',0),(112,1,59,59,10,0,0,4,'2017-02-01 00:00:00',-4150.0000,'','2017-12-05 16:13:44',0,'2017-12-05 16:13:44',0),(113,1,60,60,9,3,2,0,'2016-11-01 00:00:00',4000.0000,'','2017-12-05 16:15:10',0,'2017-12-05 16:15:10',0),(114,1,60,60,17,3,2,0,'2016-11-01 00:00:00',-4000.0000,'','2017-12-05 16:15:10',0,'2017-12-05 16:15:10',0),(115,1,61,61,9,3,2,0,'2016-12-01 00:00:00',4000.0000,'','2017-12-05 16:15:45',0,'2017-12-05 16:15:45',0),(116,1,61,61,17,3,2,0,'2016-12-01 00:00:00',-4000.0000,'','2017-12-05 16:15:45',0,'2017-12-05 16:15:45',0),(117,1,62,62,6,0,0,3,'2016-11-15 00:00:00',12000.0000,'','2017-12-05 16:16:37',0,'2017-12-05 16:16:37',0),(118,1,62,62,10,0,0,3,'2016-11-15 00:00:00',-12000.0000,'','2017-12-05 16:16:37',0,'2017-12-05 16:16:37',0),(119,1,63,63,6,0,0,1,'2017-01-01 00:00:00',-3750.0000,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(120,1,63,63,10,0,0,1,'2017-01-01 00:00:00',3750.0000,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(121,1,64,64,10,2,1,1,'2017-01-01 00:00:00',-3750.0000,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(122,1,64,64,9,2,1,1,'2017-01-01 00:00:00',3750.0000,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(123,1,65,65,6,0,0,1,'2017-03-01 00:00:00',3750.0000,'','2017-12-05 16:20:28',0,'2017-12-05 16:20:28',0),(124,1,65,65,10,0,0,1,'2017-03-01 00:00:00',-3750.0000,'','2017-12-05 16:20:28',0,'2017-12-05 16:20:28',0),(125,1,66,66,6,0,0,4,'2017-03-01 00:00:00',4150.0000,'','2017-12-05 16:22:08',0,'2017-12-05 16:22:08',0),(126,1,66,66,10,0,0,4,'2017-03-01 00:00:00',-4150.0000,'','2017-12-05 16:22:08',0,'2017-12-05 16:22:08',0),(127,1,67,67,6,0,0,1,'2017-04-01 00:00:00',3750.0000,'','2017-12-05 16:22:50',0,'2017-12-05 16:22:50',0),(128,1,67,67,10,0,0,1,'2017-04-01 00:00:00',-3750.0000,'','2017-12-05 16:22:50',0,'2017-12-05 16:22:50',0),(129,1,68,68,6,0,0,4,'2017-04-01 00:00:00',4150.0000,'','2017-12-05 16:23:12',0,'2017-12-05 16:23:12',0),(130,1,68,68,10,0,0,4,'2017-04-01 00:00:00',-4150.0000,'','2017-12-05 16:23:12',0,'2017-12-05 16:23:12',0),(131,1,69,69,6,0,0,1,'2017-05-01 00:00:00',3750.0000,'','2017-12-05 16:24:00',0,'2017-12-05 16:24:00',0),(132,1,69,69,10,0,0,1,'2017-05-01 00:00:00',-3750.0000,'','2017-12-05 16:24:01',0,'2017-12-05 16:24:01',0),(133,1,70,70,6,0,0,4,'2017-05-01 00:00:00',4150.0000,'','2017-12-05 16:24:18',0,'2017-12-05 16:24:18',0),(134,1,70,70,10,0,0,4,'2017-05-01 00:00:00',-4150.0000,'','2017-12-05 16:24:18',0,'2017-12-05 16:24:18',0),(135,1,71,71,6,0,0,3,'2017-05-15 00:00:00',13131.7900,'','2017-12-05 16:26:21',0,'2017-12-05 16:26:21',0),(136,1,71,71,10,0,0,3,'2017-05-15 00:00:00',-13131.7900,'','2017-12-05 16:26:21',0,'2017-12-05 16:26:21',0),(137,1,72,72,6,0,0,1,'2017-06-01 00:00:00',3750.0000,'','2017-12-05 16:27:03',0,'2017-12-05 16:27:03',0),(138,1,72,72,10,0,0,1,'2017-06-01 00:00:00',-3750.0000,'','2017-12-05 16:27:03',0,'2017-12-05 16:27:03',0),(139,1,73,73,6,0,0,4,'2017-06-01 00:00:00',4150.0000,'','2017-12-05 16:27:16',0,'2017-12-05 16:27:16',0),(140,1,73,73,10,0,0,4,'2017-06-01 00:00:00',-4150.0000,'','2017-12-05 16:27:16',0,'2017-12-05 16:27:16',0),(141,1,74,74,6,0,0,1,'2017-07-01 00:00:00',3750.0000,'','2017-12-05 16:27:58',0,'2017-12-05 16:27:58',0),(142,1,74,74,10,0,0,1,'2017-07-01 00:00:00',-3750.0000,'','2017-12-05 16:27:58',0,'2017-12-05 16:27:58',0),(143,1,75,75,6,0,0,4,'2017-07-01 00:00:00',4150.0000,'','2017-12-05 16:28:13',0,'2017-12-05 16:28:13',0),(144,1,75,75,10,0,0,4,'2017-07-01 00:00:00',-4150.0000,'','2017-12-05 16:28:13',0,'2017-12-05 16:28:13',0),(145,1,76,76,6,0,0,1,'2017-08-01 00:00:00',3750.0000,'','2017-12-05 16:29:16',0,'2017-12-05 16:29:16',0),(146,1,76,76,10,0,0,1,'2017-08-01 00:00:00',-3750.0000,'','2017-12-05 16:29:16',0,'2017-12-05 16:29:16',0),(147,1,77,77,6,0,0,4,'2017-08-01 00:00:00',4150.0000,'','2017-12-05 16:29:33',0,'2017-12-05 16:29:33',0),(148,1,77,77,10,0,0,4,'2017-08-01 00:00:00',-4150.0000,'','2017-12-05 16:29:33',0,'2017-12-05 16:29:33',0),(149,1,78,78,6,0,0,3,'2017-08-15 00:00:00',13050.0000,'','2017-12-05 16:29:59',0,'2017-12-05 16:29:59',0),(150,1,78,78,10,0,0,3,'2017-08-15 00:00:00',-13050.0000,'','2017-12-05 16:29:59',0,'2017-12-05 16:29:59',0),(151,1,79,79,6,0,0,1,'2017-09-01 00:00:00',3750.0000,'','2017-12-05 16:30:33',0,'2017-12-05 16:30:33',0),(152,1,79,79,10,0,0,1,'2017-09-01 00:00:00',-3750.0000,'','2017-12-05 16:30:33',0,'2017-12-05 16:30:33',0),(153,1,80,80,6,0,0,4,'2017-09-01 00:00:00',4150.0000,'','2017-12-05 16:30:51',0,'2017-12-05 16:30:51',0),(154,1,80,80,10,0,0,4,'2017-09-01 00:00:00',-4150.0000,'','2017-12-05 16:30:51',0,'2017-12-05 16:30:51',0),(155,1,81,81,6,0,0,1,'2017-10-01 00:00:00',3750.0000,'','2017-12-05 16:31:42',0,'2017-12-05 16:31:42',0),(156,1,81,81,10,0,0,1,'2017-10-01 00:00:00',-3750.0000,'','2017-12-05 16:31:42',0,'2017-12-05 16:31:42',0),(157,1,82,82,6,0,0,4,'2017-10-01 00:00:00',4150.0000,'','2017-12-05 16:31:56',0,'2017-12-05 16:31:56',0),(158,1,82,82,10,0,0,4,'2017-10-01 00:00:00',-4150.0000,'','2017-12-05 16:31:56',0,'2017-12-05 16:31:56',0),(159,1,83,83,6,0,0,1,'2017-11-01 00:00:00',3750.0000,'','2017-12-05 16:32:49',0,'2017-12-05 16:32:49',0),(160,1,83,83,10,0,0,1,'2017-11-01 00:00:00',-3750.0000,'','2017-12-05 16:32:49',0,'2017-12-05 16:32:49',0),(161,1,84,84,6,0,0,4,'2017-11-01 00:00:00',4150.0000,'','2017-12-05 16:33:11',0,'2017-12-05 16:33:11',0),(162,1,84,84,10,0,0,4,'2017-11-01 00:00:00',-4150.0000,'','2017-12-05 16:33:11',0,'2017-12-05 16:33:11',0),(163,1,85,85,6,0,0,3,'2017-11-15 00:00:00',13459.2800,'','2017-12-05 16:40:59',0,'2017-12-05 16:40:59',0),(164,1,85,85,10,0,0,3,'2017-11-15 00:00:00',-13459.2800,'','2017-12-05 16:40:59',0,'2017-12-05 16:40:59',0),(165,1,86,86,6,0,0,1,'2017-12-01 00:00:00',3750.0000,'','2017-12-05 16:42:24',0,'2017-12-05 16:42:24',0),(166,1,86,86,10,0,0,1,'2017-12-01 00:00:00',-3750.0000,'','2017-12-05 16:42:24',0,'2017-12-05 16:42:24',0),(167,1,87,87,6,0,0,4,'2017-12-01 00:00:00',4150.0000,'','2017-12-05 16:42:35',0,'2017-12-05 16:42:35',0),(168,1,87,87,10,0,0,4,'2017-12-01 00:00:00',-4150.0000,'','2017-12-05 16:42:35',0,'2017-12-05 16:42:35',0),(169,1,88,88,10,2,1,1,'2017-01-01 00:00:00',3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(170,1,88,88,9,2,1,1,'2017-01-01 00:00:00',-3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(171,1,89,89,10,2,1,1,'2017-02-01 00:00:00',3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(172,1,89,89,9,2,1,1,'2017-02-01 00:00:00',-3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(173,1,90,90,10,2,1,1,'2017-03-01 00:00:00',3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(174,1,90,90,9,2,1,1,'2017-03-01 00:00:00',-3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(175,1,91,91,10,2,1,1,'2017-04-01 00:00:00',3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(176,1,91,91,9,2,1,1,'2017-04-01 00:00:00',-3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(177,1,92,92,10,2,1,1,'2017-05-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(178,1,92,92,9,2,1,1,'2017-05-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(179,1,93,93,10,2,1,1,'2017-06-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(180,1,93,93,9,2,1,1,'2017-06-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(181,1,94,94,10,2,1,1,'2017-07-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(182,1,94,94,9,2,1,1,'2017-07-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(183,1,95,95,10,2,1,1,'2017-08-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(184,1,95,95,9,2,1,1,'2017-08-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(185,1,96,96,10,2,1,1,'2017-09-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(186,1,96,96,9,2,1,1,'2017-09-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(187,1,97,97,10,2,1,1,'2017-10-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(188,1,97,97,9,2,1,1,'2017-10-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(189,1,98,98,10,2,1,1,'2017-11-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(190,1,98,98,9,2,1,1,'2017-11-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(191,1,99,99,10,2,1,1,'2017-12-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(192,1,99,99,9,2,1,1,'2017-12-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(193,1,100,100,10,4,3,4,'2016-07-01 00:00:00',8300.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(194,1,100,100,9,4,3,4,'2016-07-01 00:00:00',-8300.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(195,1,101,101,10,4,3,4,'2017-01-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(196,1,101,101,9,4,3,4,'2017-01-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(197,1,102,102,10,4,3,4,'2017-02-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(198,1,102,102,9,4,3,4,'2017-02-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(199,1,103,103,10,4,3,4,'2017-03-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(200,1,103,103,9,4,3,4,'2017-03-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(201,1,104,104,10,4,3,4,'2017-04-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(202,1,104,104,9,4,3,4,'2017-04-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(203,1,105,105,10,4,3,4,'2017-05-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(204,1,105,105,9,4,3,4,'2017-05-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(205,1,106,106,10,4,3,4,'2017-06-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(206,1,106,106,9,4,3,4,'2017-06-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(207,1,107,107,10,4,3,4,'2017-07-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(208,1,107,107,9,4,3,4,'2017-07-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(209,1,108,108,10,4,3,4,'2017-08-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(210,1,108,108,9,4,3,4,'2017-08-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(211,1,109,109,10,4,3,4,'2017-09-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(212,1,109,109,9,4,3,4,'2017-09-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(213,1,110,110,10,4,3,4,'2017-10-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(214,1,110,110,9,4,3,4,'2017-10-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(215,1,111,111,10,4,3,4,'2017-11-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(216,1,111,111,9,4,3,4,'2017-11-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(217,1,112,112,10,4,3,4,'2017-12-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(218,1,112,112,9,4,3,4,'2017-12-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(219,1,113,113,9,3,2,0,'2017-04-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(220,1,113,113,36,3,2,0,'2017-04-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(221,1,114,114,9,3,2,0,'2017-05-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(222,1,114,114,36,3,2,0,'2017-05-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(223,1,115,115,9,3,2,0,'2017-06-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(224,1,115,115,36,3,2,0,'2017-06-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(225,1,116,116,9,3,2,0,'2017-07-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(226,1,116,116,36,3,2,0,'2017-07-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(227,1,117,117,9,3,2,0,'2017-08-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(228,1,117,117,36,3,2,0,'2017-08-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(229,1,118,118,9,3,2,0,'2017-09-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(230,1,118,118,36,3,2,0,'2017-09-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(231,1,119,119,9,3,2,0,'2017-10-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(232,1,119,119,36,3,2,0,'2017-10-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(233,1,120,120,9,3,2,0,'2017-11-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(234,1,120,120,36,3,2,0,'2017-11-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(235,1,121,121,9,3,2,0,'2017-12-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(236,1,121,121,36,3,2,0,'2017-12-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(237,1,122,122,9,3,2,0,'2016-10-01 00:00:00',4000.0000,'','2017-12-05 18:03:15',0,'2017-12-05 18:03:15',0),(238,1,122,122,17,3,2,0,'2016-10-01 00:00:00',-4000.0000,'','2017-12-05 18:03:15',0,'2017-12-05 18:03:15',0),(239,1,123,123,9,3,2,0,'2017-12-05 00:00:00',628.4500,'','2017-12-05 18:23:25',0,'2017-12-05 18:23:25',0),(240,1,123,123,37,3,2,0,'2017-12-05 00:00:00',-628.4500,'','2017-12-05 18:23:25',0,'2017-12-05 18:23:25',0),(241,1,124,124,9,3,2,0,'2017-12-05 00:00:00',-628.4500,'','2017-12-05 19:41:01',0,'2017-12-05 19:41:01',0),(242,1,124,124,37,3,2,0,'2017-12-05 00:00:00',628.4500,'','2017-12-05 19:41:01',0,'2017-12-05 19:41:01',0),(243,1,125,125,6,0,0,3,'2017-02-03 00:00:00',628.4500,'','2017-12-05 19:44:51',0,'2017-12-05 19:44:51',0),(244,1,125,125,10,0,0,3,'2017-02-03 00:00:00',-628.4500,'','2017-12-05 19:44:51',0,'2017-12-05 19:44:51',0),(245,1,126,126,6,0,0,3,'2016-10-03 00:00:00',4000.0000,'','2017-12-05 20:44:04',0,'2017-12-05 20:44:04',0),(246,1,126,126,10,0,0,3,'2016-10-03 00:00:00',-4000.0000,'','2017-12-05 20:44:04',0,'2017-12-05 20:44:04',0),(247,1,127,127,10,3,2,3,'2016-10-03 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(248,1,127,127,9,3,2,3,'2016-10-03 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(249,1,128,128,10,3,2,3,'2016-11-11 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(250,1,128,128,9,3,2,3,'2016-11-11 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(251,1,129,129,10,3,2,3,'2016-11-11 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(252,1,129,129,9,3,2,3,'2016-11-11 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(253,1,130,130,10,3,2,3,'2016-11-11 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(254,1,130,130,9,3,2,3,'2016-11-11 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(255,1,131,131,10,3,2,3,'2017-02-03 00:00:00',628.4500,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(256,1,131,131,9,3,2,3,'2017-02-03 00:00:00',-628.4500,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(257,1,132,132,10,3,2,3,'2017-02-13 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(258,1,132,132,9,3,2,3,'2017-02-13 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(259,1,133,133,10,3,2,3,'2017-02-13 00:00:00',175.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(260,1,133,133,9,3,2,3,'2017-02-13 00:00:00',-175.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(261,1,134,134,10,3,2,3,'2017-02-13 00:00:00',3546.5500,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(262,1,134,134,9,3,2,3,'2017-02-13 00:00:00',-3546.5500,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(263,1,135,135,10,3,2,3,'2017-02-13 00:00:00',453.4500,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(264,1,135,135,9,3,2,3,'2017-02-13 00:00:00',-453.4500,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(265,1,136,136,10,3,2,3,'2017-02-13 00:00:00',175.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(266,1,136,136,9,3,2,3,'2017-02-13 00:00:00',-175.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(267,1,137,137,10,3,2,3,'2017-05-12 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(268,1,137,137,9,3,2,3,'2017-05-12 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(269,1,138,138,10,3,2,3,'2017-05-12 00:00:00',350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(270,1,138,138,9,3,2,3,'2017-05-12 00:00:00',-350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(271,1,139,139,10,3,2,3,'2017-05-12 00:00:00',81.7900,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(272,1,139,139,9,3,2,3,'2017-05-12 00:00:00',-81.7900,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(273,1,140,140,10,3,2,3,'2017-05-12 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(274,1,140,140,9,3,2,3,'2017-05-12 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(275,1,141,141,10,3,2,3,'2017-05-12 00:00:00',350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(276,1,141,141,9,3,2,3,'2017-05-12 00:00:00',-350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(277,1,142,142,10,3,2,3,'2017-05-12 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(278,1,142,142,9,3,2,3,'2017-05-12 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(279,1,143,143,10,3,2,3,'2017-05-12 00:00:00',350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(280,1,143,143,9,3,2,3,'2017-05-12 00:00:00',-350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(281,1,144,144,10,3,2,3,'2017-08-15 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(282,1,144,144,9,3,2,3,'2017-08-15 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(283,1,145,145,10,3,2,3,'2017-08-15 00:00:00',350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(284,1,145,145,9,3,2,3,'2017-08-15 00:00:00',-350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(285,1,146,146,10,3,2,3,'2017-08-15 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(286,1,146,146,9,3,2,3,'2017-08-15 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(287,1,147,147,10,3,2,3,'2017-08-15 00:00:00',350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(288,1,147,147,9,3,2,3,'2017-08-15 00:00:00',-350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(289,1,148,148,10,3,2,3,'2017-08-15 00:00:00',4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(290,1,148,148,9,3,2,3,'2017-08-15 00:00:00',-4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(291,1,149,149,10,3,2,3,'2017-08-15 00:00:00',350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(292,1,149,149,9,3,2,3,'2017-08-15 00:00:00',-350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(293,1,150,150,10,3,2,3,'2017-11-14 00:00:00',4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(294,1,150,150,9,3,2,3,'2017-11-14 00:00:00',-4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(295,1,151,151,10,3,2,3,'2017-11-14 00:00:00',350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(296,1,151,151,9,3,2,3,'2017-11-14 00:00:00',-350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(297,1,152,152,10,3,2,3,'2017-11-14 00:00:00',409.2800,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(298,1,152,152,9,3,2,3,'2017-11-14 00:00:00',-409.2800,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(299,1,153,153,10,3,2,3,'2017-11-14 00:00:00',4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(300,1,153,153,9,3,2,3,'2017-11-14 00:00:00',-4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(301,1,154,154,10,3,2,3,'2017-11-14 00:00:00',350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(302,1,154,154,9,3,2,3,'2017-11-14 00:00:00',-350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(303,1,155,155,10,3,2,3,'2017-11-14 00:00:00',4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(304,1,155,155,9,3,2,3,'2017-11-14 00:00:00',-4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(305,1,156,156,10,3,2,3,'2017-11-14 00:00:00',350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(306,1,156,156,9,3,2,3,'2017-11-14 00:00:00',-350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(307,1,163,163,10,3,2,3,'2017-05-12 00:00:00',-350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(308,1,163,163,9,3,2,3,'2017-05-12 00:00:00',350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(309,1,164,164,10,3,2,3,'2017-05-12 00:00:00',-350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(310,1,164,164,9,3,2,3,'2017-05-12 00:00:00',350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(311,1,165,165,10,3,2,3,'2017-05-12 00:00:00',-350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(312,1,165,165,9,3,2,3,'2017-05-12 00:00:00',350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(313,1,166,166,10,3,2,3,'2017-08-15 00:00:00',-350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(314,1,166,166,9,3,2,3,'2017-08-15 00:00:00',350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(315,1,167,167,10,3,2,3,'2017-08-15 00:00:00',-350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(316,1,167,167,9,3,2,3,'2017-08-15 00:00:00',350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(317,1,168,168,10,3,2,3,'2017-08-15 00:00:00',-350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(318,1,168,168,9,3,2,3,'2017-08-15 00:00:00',350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(319,1,169,169,10,3,2,3,'2017-11-14 00:00:00',-350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(320,1,169,169,9,3,2,3,'2017-11-14 00:00:00',350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(321,1,170,170,10,3,2,3,'2017-11-14 00:00:00',-350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(322,1,170,170,9,3,2,3,'2017-11-14 00:00:00',350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(323,1,171,171,10,3,2,3,'2017-11-14 00:00:00',-350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(324,1,171,171,9,3,2,3,'2017-11-14 00:00:00',350.0000,'','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(325,1,172,172,10,2,1,1,'2017-01-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(326,1,172,172,9,2,1,1,'2017-01-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(327,1,173,173,10,2,1,1,'2017-01-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(328,1,173,173,9,2,1,1,'2017-01-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(329,1,174,174,10,2,1,1,'2017-01-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(330,1,174,174,9,2,1,1,'2017-01-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(331,1,175,175,10,2,1,1,'2017-02-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(332,1,175,175,9,2,1,1,'2017-02-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(333,1,176,176,10,2,1,1,'2017-03-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(334,1,176,176,9,2,1,1,'2017-03-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(335,1,177,177,10,2,1,1,'2017-04-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(336,1,177,177,9,2,1,1,'2017-04-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(337,1,178,178,10,2,1,1,'2017-05-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(338,1,178,178,9,2,1,1,'2017-05-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(339,1,179,179,10,2,1,1,'2017-06-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(340,1,179,179,9,2,1,1,'2017-06-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(341,1,180,180,10,2,1,1,'2017-07-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(342,1,180,180,9,2,1,1,'2017-07-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(343,1,181,181,10,2,1,1,'2017-08-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(344,1,181,181,9,2,1,1,'2017-08-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(345,1,182,182,10,2,1,1,'2017-09-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(346,1,182,182,9,2,1,1,'2017-09-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(347,1,183,183,10,2,1,1,'2017-10-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(348,1,183,183,9,2,1,1,'2017-10-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(349,1,184,184,10,2,1,1,'2017-11-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(350,1,184,184,9,2,1,1,'2017-11-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(351,1,185,185,10,2,1,1,'2017-12-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(352,1,185,185,9,2,1,1,'2017-12-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(353,1,186,186,9,2,1,0,'2017-01-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(354,1,186,186,17,2,1,0,'2017-01-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(355,1,187,187,9,2,1,0,'2017-02-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(356,1,187,187,17,2,1,0,'2017-02-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(357,1,188,188,9,2,1,0,'2017-03-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(358,1,188,188,17,2,1,0,'2017-03-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(359,1,189,189,9,2,1,0,'2017-04-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(360,1,189,189,17,2,1,0,'2017-04-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(361,1,190,190,9,2,1,0,'2017-05-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(362,1,190,190,17,2,1,0,'2017-05-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(363,1,191,191,9,2,1,0,'2017-06-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(364,1,191,191,17,2,1,0,'2017-06-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(365,1,192,192,9,2,1,0,'2017-07-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(366,1,192,192,17,2,1,0,'2017-07-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(367,1,193,193,9,2,1,0,'2017-08-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(368,1,193,193,17,2,1,0,'2017-08-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(369,1,194,194,9,2,1,0,'2017-09-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(370,1,194,194,17,2,1,0,'2017-09-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(371,1,195,195,9,2,1,0,'2017-10-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(372,1,195,195,17,2,1,0,'2017-10-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(373,1,196,196,9,2,1,0,'2017-11-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(374,1,196,196,17,2,1,0,'2017-11-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(375,1,197,197,9,2,1,0,'2017-12-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(376,1,197,197,17,2,1,0,'2017-12-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(377,1,198,198,9,2,1,0,'2018-01-01 00:00:00',3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(378,1,198,198,17,2,1,0,'2018-01-01 00:00:00',-3750.0000,'','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211);
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
) ENGINE=InnoDB AUTO_INCREMENT=496 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerMarker`
--

LOCK TABLES `LedgerMarker` WRITE;
/*!40000 ALTER TABLE `LedgerMarker` DISABLE KEYS */;
INSERT INTO `LedgerMarker` VALUES (1,1,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(2,2,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(3,3,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(4,4,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(6,6,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(7,7,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(8,8,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(9,9,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(10,10,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(11,11,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(12,12,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(13,13,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(14,14,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(15,15,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(16,16,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(17,17,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(19,19,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(20,20,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(21,21,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(22,22,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(23,23,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(24,24,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(25,25,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(26,26,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(27,27,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(28,28,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(29,29,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(30,30,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(31,31,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(32,32,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(33,33,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(34,34,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(35,35,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(36,36,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(37,37,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(38,38,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(39,39,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(40,40,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(41,41,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(42,42,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(43,43,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(44,44,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(45,45,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(46,46,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(47,47,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(48,48,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(49,49,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(50,50,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(51,51,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(52,52,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(53,53,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(54,54,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(55,55,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(56,56,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(57,57,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(58,58,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(59,59,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(60,60,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(61,61,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(62,62,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(63,63,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(64,64,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(65,65,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(66,66,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(67,67,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(68,68,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(69,69,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(70,70,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(71,71,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(72,72,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(73,73,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(75,0,0,1,0,0,'2017-11-28 00:00:00',0.0000,3,'2017-11-28 18:14:19',0,'2017-11-28 18:14:19',0),(76,0,1,0,0,1,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:15:20',0,'2017-11-30 18:15:20',0),(77,0,1,0,0,2,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:15:28',0,'2017-11-30 18:15:28',0),(78,0,1,0,0,3,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:16:10',0,'2017-11-30 18:16:10',0),(79,0,1,0,0,4,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:16:17',0,'2017-11-30 18:16:17',0),(80,0,1,0,0,5,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:16:28',0,'2017-11-30 18:16:28',0),(81,0,0,2,0,0,'2014-03-01 00:00:00',0.0000,3,'2017-11-30 18:29:02',0,'2017-11-30 18:17:55',0),(82,0,1,2,1,0,'2014-03-01 00:00:00',0.0000,3,'2017-11-30 18:20:15',0,'2017-11-30 18:20:15',0),(83,0,1,0,0,6,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:24:52',0,'2017-11-30 18:24:52',0),(84,0,0,3,0,0,'2016-10-01 00:00:00',0.0000,3,'2017-11-30 18:33:53',0,'2017-11-30 18:29:25',0),(85,0,1,3,2,0,'2016-10-01 00:00:00',0.0000,3,'2017-11-30 18:32:13',0,'2017-11-30 18:32:13',0),(86,0,0,4,0,0,'2016-07-01 00:00:00',0.0000,3,'2017-11-30 18:37:13',0,'2017-11-30 18:33:59',0),(87,0,1,4,3,0,'2016-07-01 00:00:00',0.0000,3,'2017-11-30 18:34:33',0,'2017-11-30 18:34:33',0),(88,0,0,5,0,0,'2017-11-30 00:00:00',0.0000,3,'2017-11-30 18:37:24',0,'2017-11-30 18:37:24',0),(89,0,0,6,0,0,'2018-01-03 00:00:00',0.0000,3,'2018-01-03 07:07:58',0,'2018-01-03 07:07:58',0),(90,0,0,7,0,0,'2018-01-03 00:00:00',0.0000,3,'2018-01-03 07:08:22',0,'2018-01-03 07:08:22',0),(91,75,2,0,0,0,'1970-01-01 00:00:00',174719.5200,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(92,76,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(93,77,2,0,0,0,'1970-01-01 00:00:00',15300.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(94,78,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(95,79,2,0,0,0,'1970-01-01 00:00:00',159419.5200,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(96,80,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(97,81,2,0,0,0,'1970-01-01 00:00:00',-11900.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(98,82,2,0,0,0,'1970-01-01 00:00:00',-11900.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(99,83,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(100,84,2,0,0,0,'1970-01-01 00:00:00',-15300.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(101,85,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(102,86,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(103,87,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(104,88,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(105,89,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(106,90,2,0,0,0,'1970-01-01 00:00:00',-142900.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(107,91,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(108,92,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(109,93,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(110,94,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(111,95,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(112,96,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(113,97,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(114,98,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(115,99,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(116,100,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(117,101,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(118,102,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(119,103,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(120,104,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(121,105,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(122,106,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(123,107,2,0,0,0,'1970-01-01 00:00:00',-4619.5200,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(124,108,2,0,0,0,'1970-01-01 00:00:00',-3150.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(125,109,2,0,0,0,'1970-01-01 00:00:00',-1469.5200,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(126,110,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(127,111,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(128,112,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(129,113,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(130,114,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(131,115,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(132,116,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(133,117,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(134,118,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(135,119,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(136,120,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(137,121,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(138,122,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(139,123,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(140,124,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(141,125,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(142,126,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(143,127,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(144,128,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(145,129,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(146,130,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(147,131,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(148,132,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(149,133,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(150,134,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(151,135,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(152,136,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(153,137,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(154,138,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(155,139,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(156,140,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(157,141,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(158,142,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(159,143,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(160,144,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(161,145,2,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(162,146,3,0,0,0,'1970-01-01 00:00:00',174719.5200,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(164,148,3,0,0,0,'1970-01-01 00:00:00',15300.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(165,149,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(166,150,3,0,0,0,'1970-01-01 00:00:00',159419.5200,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(167,151,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(168,152,3,0,0,0,'1970-01-01 00:00:00',-11900.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(169,153,3,0,0,0,'1970-01-01 00:00:00',-11900.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(170,154,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(171,155,3,0,0,0,'1970-01-01 00:00:00',-15300.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(172,156,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(173,157,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(174,158,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(175,159,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(176,160,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(177,161,3,0,0,0,'1970-01-01 00:00:00',-142900.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(178,162,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(179,163,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(180,164,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(181,165,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(182,166,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(183,167,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(184,168,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(185,169,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(186,170,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(187,171,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(188,172,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(189,173,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(190,174,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(191,175,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(192,176,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(193,177,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(194,178,3,0,0,0,'1970-01-01 00:00:00',-4619.5200,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(195,179,3,0,0,0,'1970-01-01 00:00:00',-3150.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(196,180,3,0,0,0,'1970-01-01 00:00:00',-1469.5200,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(197,181,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(198,182,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(199,183,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(200,184,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(201,185,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(202,186,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(203,187,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(204,188,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(205,189,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(206,190,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(207,191,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(208,192,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(209,193,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(210,194,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(211,195,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(212,196,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(213,197,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(214,198,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(215,199,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(216,200,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(217,201,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(218,202,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(219,203,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(220,204,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(221,205,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(222,206,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(223,207,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(224,208,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(225,209,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(226,210,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(227,211,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(228,212,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(229,213,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(230,214,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(231,215,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(232,216,3,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(233,217,4,0,0,0,'1970-01-01 00:00:00',174719.5200,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(234,218,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(235,219,4,0,0,0,'1970-01-01 00:00:00',15300.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(236,220,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(237,221,4,0,0,0,'1970-01-01 00:00:00',159419.5200,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(238,222,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(239,223,4,0,0,0,'1970-01-01 00:00:00',-11900.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(240,224,4,0,0,0,'1970-01-01 00:00:00',-11900.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(241,225,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(242,226,4,0,0,0,'1970-01-01 00:00:00',-15300.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(243,227,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(244,228,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(245,229,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(246,230,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(247,231,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(248,232,4,0,0,0,'1970-01-01 00:00:00',-142900.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(249,233,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(250,234,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(251,235,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(252,236,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(253,237,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(254,238,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(255,239,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(256,240,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(257,241,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(258,242,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(259,243,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(260,244,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(261,245,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(262,246,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(263,247,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(264,248,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(265,249,4,0,0,0,'1970-01-01 00:00:00',-4619.5200,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(266,250,4,0,0,0,'1970-01-01 00:00:00',-3150.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(267,251,4,0,0,0,'1970-01-01 00:00:00',-1469.5200,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(268,252,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(269,253,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(270,254,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(271,255,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(272,256,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(273,257,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(274,258,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(275,259,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(276,260,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(277,261,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(278,262,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(279,263,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(280,264,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(281,265,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(282,266,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(283,267,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(284,268,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(285,269,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(286,270,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(287,271,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(288,272,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(289,273,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(290,274,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(291,275,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(292,276,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(293,277,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(294,278,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(295,279,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(296,280,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(297,281,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(298,282,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(299,283,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(300,284,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(301,285,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(302,286,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(303,287,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:44',0,'2018-01-17 18:49:44',0),(304,288,5,0,0,0,'1970-01-01 00:00:00',174719.5200,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(305,289,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(306,290,5,0,0,0,'1970-01-01 00:00:00',15300.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(307,291,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(308,292,5,0,0,0,'1970-01-01 00:00:00',159419.5200,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(309,293,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(310,294,5,0,0,0,'1970-01-01 00:00:00',-11900.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(311,295,5,0,0,0,'1970-01-01 00:00:00',-11900.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(312,296,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(313,297,5,0,0,0,'1970-01-01 00:00:00',-15300.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(314,298,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(315,299,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(316,300,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(317,301,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(318,302,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(319,303,5,0,0,0,'1970-01-01 00:00:00',-142900.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(320,304,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(321,305,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(322,306,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(323,307,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(324,308,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(325,309,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(326,310,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(327,311,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(328,312,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(329,313,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(330,314,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(331,315,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(332,316,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(333,317,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(334,318,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(335,319,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(336,320,5,0,0,0,'1970-01-01 00:00:00',-4619.5200,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(337,321,5,0,0,0,'1970-01-01 00:00:00',-3150.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(338,322,5,0,0,0,'1970-01-01 00:00:00',-1469.5200,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(339,323,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(340,324,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(341,325,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(342,326,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(343,327,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(344,328,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(345,329,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(346,330,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(347,331,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(348,332,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(349,333,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(350,334,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(351,335,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(352,336,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(353,337,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(354,338,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(355,339,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(356,340,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(357,341,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(358,342,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(359,343,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(360,344,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(361,345,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(362,346,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(363,347,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(364,348,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(365,349,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(366,350,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(367,351,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(368,352,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(369,353,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(370,354,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(371,355,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(372,356,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(373,357,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(374,358,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(375,359,6,0,0,0,'1970-01-01 00:00:00',174719.5200,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(376,360,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(377,361,6,0,0,0,'1970-01-01 00:00:00',15300.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(378,362,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(379,363,6,0,0,0,'1970-01-01 00:00:00',159419.5200,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(380,364,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(381,365,6,0,0,0,'1970-01-01 00:00:00',-11900.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(382,366,6,0,0,0,'1970-01-01 00:00:00',-11900.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(383,367,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(384,368,6,0,0,0,'1970-01-01 00:00:00',-15300.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(385,369,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(386,370,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(387,371,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(388,372,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(389,373,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(390,374,6,0,0,0,'1970-01-01 00:00:00',-142900.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(391,375,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(392,376,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(393,377,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(394,378,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(395,379,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(396,380,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(397,381,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(398,382,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(399,383,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(400,384,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(401,385,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(402,386,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(403,387,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(404,388,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(405,389,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(406,390,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(407,391,6,0,0,0,'1970-01-01 00:00:00',-4619.5200,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(408,392,6,0,0,0,'1970-01-01 00:00:00',-3150.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(409,393,6,0,0,0,'1970-01-01 00:00:00',-1469.5200,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(410,394,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(411,395,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(412,396,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(413,397,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(414,398,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(415,399,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(416,400,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(417,401,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(418,402,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(419,403,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(420,404,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(421,405,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(422,406,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(423,407,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(424,408,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(425,409,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(426,410,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(427,411,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(428,412,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(429,413,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(430,414,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(431,415,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(432,416,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(433,417,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(434,418,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(435,419,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(436,420,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(437,421,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(438,422,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(439,423,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(440,424,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(441,425,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(442,426,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(443,427,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(444,428,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(445,429,6,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 18:49:45',0,'2018-01-17 18:49:45',0),(446,0,2,0,0,7,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 22:22:13',200,'2018-01-17 22:22:13',200),(447,0,2,0,0,8,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 22:23:47',200,'2018-01-17 22:23:47',200),(448,0,2,0,0,9,'1970-01-01 00:00:00',0.0000,3,'2018-01-17 22:25:56',200,'2018-01-17 22:25:56',200),(449,0,0,8,0,0,'2015-04-29 00:00:00',0.0000,3,'2018-01-17 22:30:29',200,'2018-01-17 22:26:30',200),(450,0,2,8,5,0,'2018-01-01 00:00:00',0.0000,3,'2018-01-17 23:13:37',200,'2018-01-17 23:13:37',200),(451,0,2,8,6,0,'2018-01-01 00:00:00',0.0000,3,'2018-01-24 20:51:54',200,'2018-01-24 20:51:54',200),(452,0,0,9,0,0,'2018-01-24 00:00:00',0.0000,3,'2018-01-24 20:58:06',200,'2018-01-24 20:58:06',200),(453,0,0,10,0,0,'2018-01-24 00:00:00',0.0000,3,'2018-01-24 21:12:53',200,'2018-01-24 21:12:53',200),(454,430,4,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-24 21:26:43',200,'2018-01-24 21:26:43',200),(455,0,0,11,0,0,'2015-04-01 00:00:00',0.0000,3,'2018-02-14 21:50:08',200,'2018-01-25 20:30:36',200),(456,0,4,11,7,0,'2015-04-01 00:00:00',0.0000,3,'2018-01-25 20:32:46',200,'2018-01-25 20:32:46',200),(457,0,0,12,0,0,'2015-04-01 00:00:00',0.0000,3,'2018-01-25 20:35:55',200,'2018-01-25 20:34:31',200),(458,0,4,12,7,0,'2015-04-01 00:00:00',0.0000,3,'2018-01-25 20:35:16',200,'2018-01-25 20:35:16',200),(459,0,4,0,0,10,'1970-01-01 00:00:00',0.0000,3,'2018-01-25 20:39:56',200,'2018-01-25 20:39:56',200),(460,0,0,13,0,0,'2018-01-25 00:00:00',0.0000,3,'2018-01-25 20:40:06',200,'2018-01-25 20:40:06',200),(461,0,0,14,0,0,'2018-01-25 00:00:00',0.0000,3,'2018-01-25 21:31:40',211,'2018-01-25 21:31:40',211),(462,0,0,15,0,0,'2018-01-25 00:00:00',0.0000,3,'2018-01-25 21:32:28',211,'2018-01-25 21:32:28',211),(463,0,0,16,0,0,'2018-01-25 00:00:00',0.0000,3,'2018-01-25 21:32:30',200,'2018-01-25 21:32:30',200),(464,0,0,17,0,0,'2018-01-25 00:00:00',0.0000,3,'2018-01-25 21:33:21',200,'2018-01-25 21:33:21',200),(465,0,6,14,15,0,'2018-01-25 00:00:00',0.0000,3,'2018-01-25 21:34:06',211,'2018-01-25 21:34:06',211),(466,0,6,0,0,11,'1970-01-01 00:00:00',0.0000,3,'2018-01-25 21:36:56',211,'2018-01-25 21:36:56',211),(467,431,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-25 21:57:29',200,'2018-01-25 21:57:29',200),(468,0,5,0,0,12,'1970-01-01 00:00:00',0.0000,3,'2018-01-25 22:24:15',200,'2018-01-25 22:24:15',200),(469,0,5,0,0,13,'1970-01-01 00:00:00',0.0000,3,'2018-01-25 22:25:10',200,'2018-01-25 22:25:10',200),(470,0,5,0,0,14,'1970-01-01 00:00:00',0.0000,3,'2018-01-25 22:25:45',200,'2018-01-25 22:25:45',200),(471,0,0,18,0,0,'2007-12-24 00:00:00',0.0000,3,'2018-01-25 22:29:26',200,'2018-01-25 22:25:58',200),(472,0,5,18,16,0,'2007-12-24 00:00:00',0.0000,3,'2018-01-25 22:27:39',200,'2018-01-25 22:27:39',200),(473,0,0,19,0,0,'2015-06-01 00:00:00',0.0000,3,'2018-01-25 22:31:50',200,'2018-01-25 22:29:34',200),(474,0,5,19,17,0,'2015-06-01 00:00:00',0.0000,3,'2018-01-25 22:30:32',200,'2018-01-25 22:30:32',200),(475,0,0,20,0,0,'2009-06-05 00:00:00',0.0000,3,'2018-01-25 22:37:11',200,'2018-01-25 22:31:58',200),(476,432,5,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-01-25 22:34:38',200,'2018-01-25 22:34:38',200),(477,0,5,20,20,0,'2009-06-05 00:00:00',0.0000,3,'2018-01-25 22:35:49',200,'2018-01-25 22:35:49',200),(478,0,3,0,0,15,'1970-01-01 00:00:00',0.0000,3,'2018-01-31 21:48:40',200,'2018-01-31 21:48:40',200),(479,0,3,0,0,16,'1970-01-01 00:00:00',0.0000,3,'2018-01-31 21:50:24',200,'2018-01-31 21:50:24',200),(480,0,0,21,0,0,'2010-01-04 00:00:00',0.0000,3,'2018-01-31 21:54:52',200,'2018-01-31 21:51:12',200),(481,0,3,21,13,0,'2010-01-01 00:00:00',0.0000,3,'2018-01-31 21:53:59',200,'2018-01-31 21:53:59',200),(482,0,0,22,0,0,'2010-01-01 00:00:00',0.0000,3,'2018-01-31 21:56:28',200,'2018-01-31 21:54:58',200),(483,0,3,22,9,0,'2018-01-31 00:00:00',0.0000,3,'2018-01-31 21:55:58',200,'2018-01-31 21:55:58',200),(484,0,0,23,0,0,'2010-01-01 00:00:00',0.0000,3,'2018-01-31 21:58:23',200,'2018-01-31 21:56:32',200),(485,0,3,23,10,0,'2018-01-31 00:00:00',0.0000,3,'2018-01-31 21:57:33',200,'2018-01-31 21:57:33',200),(486,0,0,24,0,0,'2010-01-01 00:00:00',0.0000,3,'2018-01-31 22:00:58',200,'2018-01-31 21:59:51',200),(487,0,3,24,11,0,'2018-01-31 00:00:00',0.0000,3,'2018-01-31 22:00:53',200,'2018-01-31 22:00:53',200),(488,0,0,25,0,0,'2010-01-01 00:00:00',0.0000,3,'2018-01-31 22:03:50',200,'2018-01-31 22:01:03',200),(489,0,3,25,12,0,'2018-01-01 00:00:00',0.0000,3,'2018-01-31 22:02:52',200,'2018-01-31 22:02:52',200),(490,0,0,26,0,0,'2010-01-01 00:00:00',0.0000,3,'2018-01-31 22:16:48',200,'2018-01-31 22:09:50',200),(491,0,0,27,0,0,'2018-01-31 00:00:00',0.0000,3,'2018-01-31 22:10:59',200,'2018-01-31 22:10:59',200),(492,0,3,26,14,0,'2018-01-01 00:00:00',0.0000,3,'2018-01-31 22:15:01',200,'2018-01-31 22:15:01',200),(493,0,3,26,14,0,'2018-01-01 00:00:00',0.0000,3,'2018-01-31 22:16:18',200,'2018-01-31 22:16:18',200),(494,0,3,0,0,17,'1970-01-01 00:00:00',0.0000,3,'2018-01-31 22:17:45',200,'2018-01-31 22:17:45',200),(495,0,1,2,3,0,'2014-03-01 00:00:00',0.0000,3,'2018-02-20 21:27:16',211,'2018-02-20 21:27:16',211);
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
  `Active` smallint(6) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
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
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PaymentType`
--

LOCK TABLES `PaymentType` WRITE;
/*!40000 ALTER TABLE `PaymentType` DISABLE KEYS */;
INSERT INTO `PaymentType` VALUES (1,1,'Cash','Cash','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(2,1,'Check','Personal check from payor','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(3,1,'VISA','Credit card charge','2017-11-10 23:24:23',0,'2017-12-03 08:28:05',0),(4,1,'AMEX','American Express credit card','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(5,1,'ACH','Bank transfer','2017-11-10 23:24:23',0,'2017-11-28 17:56:08',0),(6,1,'Wire','Wire transfer','2017-11-30 19:42:29',0,'2017-11-30 19:42:29',0),(7,2,'EFT','Electronic Funds Transfer','2018-01-24 20:54:27',200,'2018-01-24 20:54:27',200),(8,2,'Check','Paper Check','2018-01-24 20:54:48',200,'2018-01-24 20:54:48',200),(9,2,'ACH','Overnight Funds Transfer','2018-01-24 20:55:11',200,'2018-01-24 20:55:11',200),(10,2,'Wire','Wire Transfer','2018-01-24 20:55:40',200,'2018-01-24 20:55:40',200),(11,4,'EFT','Electronic Funds Transfer','2018-01-24 21:28:36',200,'2018-01-24 21:28:36',200),(12,4,'ACH','Overnight Funds Transfer','2018-01-24 21:28:51',200,'2018-01-24 21:28:51',200),(13,4,'Wire','Wired Funds','2018-01-24 21:29:03',200,'2018-01-24 21:29:03',200),(14,4,'Check','Paper Check','2018-01-24 21:29:25',200,'2018-01-24 21:29:42',200),(15,6,'EFT','Electronic Funds Transfer','2018-01-25 20:51:20',200,'2018-01-25 20:51:20',200),(16,6,'ACH','Overnight Funds Transfer','2018-01-25 20:51:32',200,'2018-01-25 20:51:32',200),(17,6,'Wire','Wired Funds ','2018-01-25 20:51:53',200,'2018-01-25 20:51:53',200),(18,6,'Check','Paper Check','2018-01-25 20:52:21',200,'2018-01-25 20:52:45',200),(19,3,'ACH','Overnight Transfers','2018-01-25 21:08:13',200,'2018-01-25 21:08:13',200),(20,3,'EFT','Electronic Funds Transfer','2018-01-25 21:08:22',200,'2018-01-25 21:08:22',200),(21,3,'Wire','Wired Funds Transfer','2018-01-25 21:08:32',200,'2018-01-25 21:08:32',200),(22,3,'Check','Paper Check','2018-01-25 21:08:46',200,'2018-01-25 21:08:46',200),(23,5,'ACH','Overnight Funds Transfer','2018-01-25 22:04:23',200,'2018-01-25 22:04:23',200),(24,5,'EFT','Electronic Funds Transfer','2018-01-25 22:04:33',200,'2018-01-25 22:04:33',200),(25,5,'Wire','Wired Funds Transfer','2018-01-25 22:04:43',200,'2018-01-25 22:04:43',200),(26,5,'Check','Paper Check','2018-01-25 22:11:52',200,'2018-01-25 22:11:52',200);
/*!40000 ALTER TABLE `PaymentType` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Payor`
--

DROP TABLE IF EXISTS `Payor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Payor` (
  `TCID` bigint(20) NOT NULL,
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
INSERT INTO `Payor` VALUES (1,1,'',0.0000,0,1,'2017-11-30 18:15:20',0,'2017-11-30 18:15:20',0),(2,1,'',0.0000,0,1,'2017-11-30 18:15:28',0,'2017-11-30 18:15:28',0),(3,1,'',0.0000,0,1,'2017-11-30 18:16:10',0,'2017-11-30 18:16:10',0),(4,1,'',0.0000,0,1,'2017-11-30 18:16:17',0,'2017-11-30 18:16:17',0),(5,1,'',0.0000,0,1,'2017-11-30 18:16:28',0,'2017-11-30 18:16:28',0),(6,1,'',0.0000,0,1,'2017-11-30 18:24:52',0,'2017-11-30 18:24:52',0),(7,2,'',0.0000,0,1,'2018-01-17 22:22:13',200,'2018-01-17 22:22:13',200),(8,2,'',0.0000,0,1,'2018-01-17 22:23:47',200,'2018-01-17 22:23:47',200),(9,2,'',0.0000,0,1,'2018-01-17 22:25:56',200,'2018-01-17 22:25:56',200),(10,4,'',0.0000,0,1,'2018-01-25 20:39:56',200,'2018-01-25 20:39:56',200),(11,6,'',0.0000,0,1,'2018-01-25 21:36:56',211,'2018-01-25 21:36:56',211),(12,5,'',0.0000,0,1,'2018-01-25 22:24:15',200,'2018-01-25 22:24:15',200),(13,5,'',0.0000,0,1,'2018-01-25 22:25:10',200,'2018-01-25 22:25:10',200),(14,5,'',0.0000,0,1,'2018-01-25 22:25:45',200,'2018-01-25 22:25:45',200),(15,3,'',0.0000,0,1,'2018-01-31 21:48:40',200,'2018-01-31 21:48:40',200),(16,3,'',0.0000,0,1,'2018-01-31 21:50:24',200,'2018-01-31 21:50:24',200),(17,3,'',0.0000,0,1,'2018-01-31 22:17:45',200,'2018-01-31 22:17:45',200);
/*!40000 ALTER TABLE `Payor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Prospect`
--

DROP TABLE IF EXISTS `Prospect`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Prospect` (
  `TCID` bigint(20) NOT NULL,
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
INSERT INTO `Prospect` VALUES (1,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-11-30 18:15:20',0,'2017-11-30 18:15:20',0),(2,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-11-30 18:15:28',0,'2017-11-30 18:15:28',0),(3,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-11-30 18:16:10',0,'2017-11-30 18:16:10',0),(4,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-11-30 18:16:17',0,'2017-11-30 18:16:17',0),(5,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-11-30 18:16:28',0,'2017-11-30 18:16:28',0),(6,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-11-30 18:24:52',0,'2017-11-30 18:24:52',0),(7,2,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2018-01-17 22:22:13',200,'2018-01-17 22:22:13',200),(8,2,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2018-01-17 22:23:47',200,'2018-01-17 22:23:47',200),(9,2,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2018-01-17 22:25:56',200,'2018-01-17 22:25:56',200),(10,4,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2018-01-25 20:39:56',200,'2018-01-25 20:39:56',200),(11,6,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2018-01-25 21:36:56',211,'2018-01-25 21:36:56',211),(12,5,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2018-01-25 22:24:15',200,'2018-01-25 22:24:15',200),(13,5,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2018-01-25 22:25:10',200,'2018-01-25 22:25:10',200),(14,5,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2018-01-25 22:25:45',200,'2018-01-25 22:25:45',200),(15,3,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2018-01-31 21:48:40',200,'2018-01-31 21:48:40',200),(16,3,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2018-01-31 21:50:24',200,'2018-01-31 21:50:24',200),(17,3,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2018-01-31 22:17:45',200,'2018-01-31 22:17:45',200);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `RPRRTRateID` bigint(20) NOT NULL AUTO_INCREMENT,
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `RPRSPRateID` bigint(20) NOT NULL AUTO_INCREMENT,
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
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Receipt`
--

LOCK TABLES `Receipt` WRITE;
/*!40000 ALTER TABLE `Receipt` DISABLE KEYS */;
INSERT INTO `Receipt` VALUES (1,0,1,1,2,0,0,0,'2014-03-01 00:00:00','1234',7000.0000,'',10,'',4,'Reversed by receipt RCPT00000005','','2017-11-30 19:24:46',0,'2017-11-30 18:48:02',0),(2,0,1,4,2,0,0,0,'2016-07-01 00:00:00','2345',8300.0000,'',10,'',4,'Reversed by receipt RCPT00000004','','2017-11-30 19:16:53',0,'2017-11-30 18:49:17',0),(3,0,1,4,2,0,3,0,'2016-07-01 00:00:00','2456',8300.0000,'',25,'ASM(2) d 12999 8300.00,c 12001 8300.00',2,'','','2017-12-05 17:06:25',0,'2017-11-30 19:13:23',0),(4,2,1,4,2,0,0,0,'2016-07-01 00:00:00','2345',-8300.0000,'',10,'',4,'Reversal of receipt RCPT00000002','','2017-11-30 19:16:53',0,'2017-11-30 19:13:59',0),(5,1,1,1,2,0,0,0,'2014-03-01 00:00:00','1234',-7000.0000,'',10,'',4,'Reversal of receipt RCPT00000001','','2017-11-30 19:24:46',0,'2017-11-30 19:23:47',0),(6,0,1,1,2,0,4,0,'2014-03-01 00:00:00','3457',7000.0000,'',25,'ASM(1) d 12999 7000.00,c 12001 7000.00',2,'','','2017-11-30 19:46:56',0,'2017-11-30 19:24:32',0),(7,0,1,1,6,0,0,0,'2017-01-01 00:00:00','2354',3750.0000,'',25,'ASM(4) d 12999 3750.00,c 12001 3750.00,ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00',4,'Reversed by receipt RCPT00000014','','2018-02-16 22:14:38',211,'2017-11-30 19:44:52',0),(8,0,1,1,6,0,0,0,'2017-01-01 00:00:00','',3750.0000,'',25,'ASM(4) d 12999 3750.00,c 12001 3750.00,ASM(4) d 12999 3750.00,ASM(4) c 12001 3750.00,ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00,ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00,ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00,ASM(4) d 12999 3750.00,ASM(4) c 12001 3750.00,ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00,ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00',1,'','Kirsten Read','2018-02-16 22:14:38',211,'2017-12-05 16:09:37',0),(9,0,1,4,2,0,0,0,'2017-01-01 00:00:00','',4150.0000,'',25,'ASM(28) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:10:06',0),(10,0,1,3,6,0,0,0,'2017-02-01 00:00:00','',8350.0000,'',25,'ASM(42) d 12999 628.45,c 12001 628.45,ASM(17) d 12999 4000.00,c 12001 4000.00,ASM(43) d 12999 175.00,c 12001 175.00,ASM(18) d 12999 3546.55,c 12001 3546.55',2,'','','2017-12-05 20:50:21',0,'2017-12-05 16:12:02',0),(11,0,1,1,6,0,0,0,'2017-02-01 00:00:00','',3750.0000,'',25,'ASM(5) d 12999 3750.00,c 12001 3750.00,ASM(5) d 12999 -3750.00,ASM(5) c 12001 -3750.00',0,'','Kirsten Read','2018-02-16 22:14:38',211,'2017-12-05 16:12:32',0),(12,0,1,4,2,0,0,0,'2017-02-01 00:00:00','',4150.0000,'',25,'ASM(29) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:13:44',0),(13,0,1,3,6,0,0,0,'2016-11-15 00:00:00','',12000.0000,'',25,'ASM(47) d 12999 4000.00,c 12001 4000.00,ASM(48) d 12999 4000.00,c 12001 4000.00,ASM(16) d 12999 4000.00,c 12001 4000.00',2,'3 month rent in advance','','2017-12-05 20:50:21',0,'2017-12-05 16:16:37',0),(14,7,1,1,6,0,0,0,'2017-01-01 00:00:00','2354',-3750.0000,'',25,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00,ASM(4) d 12999 3750.00,ASM(4) c 12001 3750.00,ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00,ASM(4) d 12999 3750.00,ASM(4) c 12001 3750.00',4,'Reversal of receipt RCPT00000007','','2018-02-16 22:14:38',211,'2017-12-05 16:19:04',0),(15,0,1,1,6,0,0,0,'2017-03-01 00:00:00','',3750.0000,'',25,'ASM(6) d 12999 3750.00,c 12001 3750.00,ASM(6) d 12999 -3750.00,ASM(6) c 12001 -3750.00',0,'','Kirsten Read','2018-02-16 22:14:38',211,'2017-12-05 16:20:28',0),(16,0,1,4,2,0,0,0,'2017-03-01 00:00:00','',4150.0000,'',25,'ASM(30) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:22:08',0),(17,0,1,1,6,0,0,0,'2017-04-01 00:00:00','',3750.0000,'',25,'ASM(7) d 12999 3750.00,c 12001 3750.00,ASM(7) d 12999 -3750.00,ASM(7) c 12001 -3750.00',0,'','Kirsten Read','2018-02-16 22:14:38',211,'2017-12-05 16:22:50',0),(18,0,1,4,2,0,0,0,'2017-04-01 00:00:00','',4150.0000,'',25,'ASM(31) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:23:12',0),(19,0,1,1,6,0,0,0,'2017-05-01 00:00:00','',3750.0000,'',25,'ASM(8) d 12999 3750.00,c 12001 3750.00,ASM(8) d 12999 -3750.00,ASM(8) c 12001 -3750.00',0,'','Kirsten Read','2018-02-16 22:14:38',211,'2017-12-05 16:24:00',0),(20,0,1,4,2,0,0,0,'2017-05-01 00:00:00','',4150.0000,'',25,'ASM(32) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:24:18',0),(21,0,1,3,6,0,0,0,'2017-05-15 00:00:00','',13131.7900,'',25,'ASM(19) d 12999 4000.00,c 12001 4000.00,ASM(20) d 12999 4000.00,c 12001 4000.00,ASM(21) d 12999 4000.00,c 12001 4000.00,ASM(50) d 12999 350.00,c 12001 350.00,ASM(51) d 12999 350.00,c 12001 350.00,ASM(52) d 12999 350.00,c 12001 350.00,ASM(45) d 12999 81.79,c 12001 81.79,ASM(50) d 12999 -350.00,ASM(50) c 12001 -350.00,ASM(51) d 12999 -350.00,ASM(51) c 12001 -350.00,ASM(52) d 12999 -350.00,ASM(52) c 12001 -350.00',1,'3 month rent in advance and utilities overage','','2018-02-15 21:49:51',211,'2017-12-05 16:26:21',0),(22,0,1,1,6,0,0,0,'2017-06-01 00:00:00','',3750.0000,'',25,'ASM(9) d 12999 3750.00,c 12001 3750.00,ASM(9) d 12999 -3750.00,ASM(9) c 12001 -3750.00',0,'','Kirsten Read','2018-02-16 22:14:38',211,'2017-12-05 16:27:03',0),(23,0,1,4,2,0,0,0,'2017-06-01 00:00:00','',4150.0000,'',25,'ASM(33) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:27:16',0),(24,0,1,1,6,0,0,0,'2017-07-01 00:00:00','',3750.0000,'',25,'ASM(10) d 12999 3750.00,c 12001 3750.00,ASM(10) d 12999 -3750.00,ASM(10) c 12001 -3750.00',0,'','Kirsten Read','2018-02-16 22:14:38',211,'2017-12-05 16:27:58',0),(25,0,1,4,2,0,0,0,'2017-07-01 00:00:00','',4150.0000,'',25,'ASM(34) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:28:12',0),(26,0,1,1,6,0,0,0,'2017-08-01 00:00:00','',3750.0000,'',25,'ASM(11) d 12999 3750.00,c 12001 3750.00,ASM(11) d 12999 -3750.00,ASM(11) c 12001 -3750.00',0,'','Kirsten Read','2018-02-16 22:14:38',211,'2017-12-05 16:29:16',0),(27,0,1,4,2,0,0,0,'2017-08-01 00:00:00','',4150.0000,'',25,'ASM(35) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:29:33',0),(28,0,1,3,6,0,0,0,'2017-08-15 00:00:00','',13050.0000,'',25,'ASM(22) d 12999 4000.00,c 12001 4000.00,ASM(23) d 12999 4000.00,c 12001 4000.00,ASM(24) d 12999 4000.00,c 12001 4000.00,ASM(53) d 12999 350.00,c 12001 350.00,ASM(54) d 12999 350.00,c 12001 350.00,ASM(55) d 12999 350.00,c 12001 350.00,ASM(53) d 12999 -350.00,ASM(53) c 12001 -350.00,ASM(54) d 12999 -350.00,ASM(54) c 12001 -350.00,ASM(55) d 12999 -350.00,ASM(55) c 12001 -350.00',1,'3 month rent in advance','','2018-02-15 21:49:51',211,'2017-12-05 16:29:59',0),(29,0,1,1,6,0,0,0,'2017-09-01 00:00:00','',3750.0000,'',25,'ASM(12) d 12999 3750.00,c 12001 3750.00,ASM(12) d 12999 -3750.00,ASM(12) c 12001 -3750.00',0,'','Kirsten Read','2018-02-16 22:14:38',211,'2017-12-05 16:30:33',0),(30,0,1,4,2,0,0,0,'2017-09-01 00:00:00','',4150.0000,'',25,'ASM(36) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:30:51',0),(31,0,1,1,6,0,0,0,'2017-10-01 00:00:00','',3750.0000,'',25,'ASM(13) d 12999 3750.00,c 12001 3750.00,ASM(13) d 12999 -3750.00,ASM(13) c 12001 -3750.00',0,'','Kirsten Read','2018-02-16 22:14:38',211,'2017-12-05 16:31:42',0),(32,0,1,4,2,0,0,0,'2017-10-01 00:00:00','',4150.0000,'',25,'ASM(37) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:31:56',0),(33,0,1,1,6,0,0,0,'2017-11-01 00:00:00','',3750.0000,'',25,'ASM(14) d 12999 3750.00,c 12001 3750.00,ASM(14) d 12999 -3750.00,ASM(14) c 12001 -3750.00',0,'','Kirsten Read','2018-02-16 22:14:38',211,'2017-12-05 16:32:48',0),(34,0,1,4,2,0,0,0,'2017-11-01 00:00:00','',4150.0000,'',25,'ASM(38) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:33:11',0),(35,0,1,3,6,0,0,0,'2017-11-15 00:00:00','',13459.2800,'',25,'ASM(25) d 12999 4000.00,c 12001 4000.00,ASM(26) d 12999 4000.00,c 12001 4000.00,ASM(40) d 12999 4000.00,c 12001 4000.00,ASM(46) d 12999 409.28,c 12001 409.28,ASM(56) d 12999 350.00,c 12001 350.00,ASM(57) d 12999 350.00,c 12001 350.00,ASM(58) d 12999 350.00,c 12001 350.00,ASM(56) d 12999 -350.00,ASM(56) c 12001 -350.00,ASM(57) d 12999 -350.00,ASM(57) c 12001 -350.00,ASM(58) d 12999 -350.00,ASM(58) c 12001 -350.00',1,'3 month rent in advance and utilities overage','','2018-02-15 21:49:51',211,'2017-12-05 16:40:59',0),(36,0,1,1,6,0,0,0,'2017-12-01 00:00:00','',3750.0000,'',25,'ASM(39) d 12999 3750.00,c 12001 3750.00,ASM(39) d 12999 -3750.00,ASM(39) c 12001 -3750.00',0,'','Kirsten Read','2018-02-16 22:14:38',211,'2017-12-05 16:42:24',0),(37,0,1,4,2,0,0,0,'2017-12-01 00:00:00','',4150.0000,'',25,'',2,'','','2017-12-06 11:51:45',0,'2017-12-05 16:42:35',0),(38,0,1,3,6,0,0,0,'2017-02-03 00:00:00','',628.4500,'',25,'ASM(18) d 12999 453.45,c 12001 453.45,ASM(44) d 12999 175.00,c 12001 175.00',2,'','','2017-12-05 20:50:21',0,'2017-12-05 19:44:51',0),(39,0,1,3,2,0,0,0,'2016-10-03 00:00:00','',4000.0000,'',25,'ASM(59) d 12999 4000.00,c 12001 4000.00',2,'','','2017-12-05 20:50:21',0,'2017-12-05 20:44:04',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=141 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ReceiptAllocation`
--

LOCK TABLES `ReceiptAllocation` WRITE;
/*!40000 ALTER TABLE `ReceiptAllocation` DISABLE KEYS */;
INSERT INTO `ReceiptAllocation` VALUES (1,1,1,0,'2014-03-01 00:00:00',7000.0000,0,4,'d 10104 _, c 10999 _','2017-11-30 19:23:47',0,'2017-11-30 18:48:02',0),(2,2,1,0,'2016-07-01 00:00:00',8300.0000,0,4,'d 10104 _, c 10999 _','2017-11-30 19:13:59',0,'2017-11-30 18:49:17',0),(3,3,1,0,'2016-07-01 00:00:00',8300.0000,0,0,'d 10999 _, c 12999 _','2017-11-30 19:13:23',0,'2017-11-30 19:13:23',0),(4,4,1,0,'2016-07-01 00:00:00',-8300.0000,0,4,'d 10104 _, c 10999 _','2017-11-30 19:13:59',0,'2017-11-30 19:13:59',0),(5,3,1,0,'2016-07-01 00:00:00',8300.0000,0,0,'d 10104 _, c 10999 _','2017-11-30 19:17:32',0,'2017-11-30 19:17:32',0),(6,5,1,0,'2014-03-01 00:00:00',-7000.0000,0,4,'d 10104 _, c 10999 _','2017-11-30 19:23:47',0,'2017-11-30 19:23:47',0),(7,6,1,0,'2014-03-01 00:00:00',7000.0000,0,0,'d 10999 _, c 12999 _','2017-11-30 19:24:32',0,'2017-11-30 19:24:32',0),(8,6,1,0,'2014-03-01 00:00:00',7000.0000,0,0,'d 10104 _, c 10999 _','2017-11-30 19:25:24',0,'2017-11-30 19:25:24',0),(9,7,1,0,'2017-01-01 00:00:00',3750.0000,0,4,'d 10999 _, c 12999 _','2017-12-05 16:19:04',0,'2017-11-30 19:44:52',0),(10,6,1,2,'2014-03-01 00:00:00',7000.0000,1,0,'ASM(1) d 12999 7000.00,c 12001 7000.00','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(11,7,1,2,'2017-01-01 00:00:00',3750.0000,4,4,'ASM(4) d 12999 3750.00,c 12001 3750.00','2018-02-16 22:14:38',211,'2017-11-30 19:46:56',0),(12,8,1,0,'2017-01-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:09:37',0,'2017-12-05 16:09:37',0),(13,9,1,0,'2017-01-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:10:06',0,'2017-12-05 16:10:06',0),(14,10,1,0,'2017-02-01 00:00:00',8350.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:12:02',0,'2017-12-05 16:12:02',0),(15,11,1,0,'2017-02-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:12:32',0,'2017-12-05 16:12:32',0),(16,12,1,0,'2017-02-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:13:44',0,'2017-12-05 16:13:44',0),(17,13,1,0,'2016-11-15 00:00:00',12000.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:16:37',0,'2017-12-05 16:16:37',0),(18,14,1,0,'2017-01-01 00:00:00',-3750.0000,0,4,'d 10999 _, c 12999 _','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(19,14,1,2,'2017-12-05 16:19:04',-3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2017-12-05 16:19:04',0),(20,15,1,0,'2017-03-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:20:28',0,'2017-12-05 16:20:28',0),(21,16,1,0,'2017-03-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:22:08',0,'2017-12-05 16:22:08',0),(22,17,1,0,'2017-04-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:22:50',0,'2017-12-05 16:22:50',0),(23,18,1,0,'2017-04-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:23:12',0,'2017-12-05 16:23:12',0),(24,19,1,0,'2017-05-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:24:00',0,'2017-12-05 16:24:00',0),(25,20,1,0,'2017-05-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:24:18',0,'2017-12-05 16:24:18',0),(26,21,1,0,'2017-05-15 00:00:00',13131.7900,0,0,'d 10999 _, c 12999 _','2017-12-05 16:26:21',0,'2017-12-05 16:26:21',0),(27,22,1,0,'2017-06-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:27:03',0,'2017-12-05 16:27:03',0),(28,23,1,0,'2017-06-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:27:16',0,'2017-12-05 16:27:16',0),(29,24,1,0,'2017-07-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:27:58',0,'2017-12-05 16:27:58',0),(30,25,1,0,'2017-07-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:28:12',0,'2017-12-05 16:28:12',0),(31,26,1,0,'2017-08-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:29:16',0,'2017-12-05 16:29:16',0),(32,27,1,0,'2017-08-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:29:33',0,'2017-12-05 16:29:33',0),(33,28,1,0,'2017-08-15 00:00:00',13050.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:29:59',0,'2017-12-05 16:29:59',0),(34,29,1,0,'2017-09-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:30:33',0,'2017-12-05 16:30:33',0),(35,30,1,0,'2017-09-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:30:51',0,'2017-12-05 16:30:51',0),(36,31,1,0,'2017-10-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:31:42',0,'2017-12-05 16:31:42',0),(37,32,1,0,'2017-10-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:31:56',0,'2017-12-05 16:31:56',0),(38,33,1,0,'2017-11-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:32:48',0,'2017-12-05 16:32:48',0),(39,34,1,0,'2017-11-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:33:11',0,'2017-12-05 16:33:11',0),(40,35,1,0,'2017-11-15 00:00:00',13459.2800,0,0,'d 10999 _, c 12999 _','2017-12-05 16:40:59',0,'2017-12-05 16:40:59',0),(41,36,1,0,'2017-12-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:42:24',0,'2017-12-05 16:42:24',0),(42,37,1,0,'2017-12-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:42:35',0,'2017-12-05 16:42:35',0),(43,8,1,2,'2017-01-01 00:00:00',3750.0000,4,4,'ASM(4) d 12999 3750.00,c 12001 3750.00','2018-02-16 22:14:38',211,'2017-12-05 16:59:31',0),(44,11,1,2,'2017-02-01 00:00:00',3750.0000,5,4,'ASM(5) d 12999 3750.00,c 12001 3750.00','2018-02-16 22:14:38',211,'2017-12-05 16:59:31',0),(45,15,1,2,'2017-03-01 00:00:00',3750.0000,6,4,'ASM(6) d 12999 3750.00,c 12001 3750.00','2018-02-16 22:14:38',211,'2017-12-05 16:59:31',0),(46,17,1,2,'2017-04-01 00:00:00',3750.0000,7,4,'ASM(7) d 12999 3750.00,c 12001 3750.00','2018-02-16 22:14:38',211,'2017-12-05 16:59:31',0),(47,19,1,2,'2017-05-01 00:00:00',3750.0000,8,4,'ASM(8) d 12999 3750.00,c 12001 3750.00','2018-02-16 22:14:38',211,'2017-12-05 16:59:31',0),(48,22,1,2,'2017-06-01 00:00:00',3750.0000,9,4,'ASM(9) d 12999 3750.00,c 12001 3750.00','2018-02-16 22:14:38',211,'2017-12-05 16:59:32',0),(49,24,1,2,'2017-07-01 00:00:00',3750.0000,10,4,'ASM(10) d 12999 3750.00,c 12001 3750.00','2018-02-16 22:14:38',211,'2017-12-05 16:59:32',0),(50,26,1,2,'2017-08-01 00:00:00',3750.0000,11,4,'ASM(11) d 12999 3750.00,c 12001 3750.00','2018-02-16 22:14:38',211,'2017-12-05 16:59:32',0),(51,29,1,2,'2017-09-01 00:00:00',3750.0000,12,4,'ASM(12) d 12999 3750.00,c 12001 3750.00','2018-02-16 22:14:38',211,'2017-12-05 16:59:32',0),(52,31,1,2,'2017-10-01 00:00:00',3750.0000,13,4,'ASM(13) d 12999 3750.00,c 12001 3750.00','2018-02-16 22:14:38',211,'2017-12-05 16:59:32',0),(53,33,1,2,'2017-11-01 00:00:00',3750.0000,14,4,'ASM(14) d 12999 3750.00,c 12001 3750.00','2018-02-16 22:14:38',211,'2017-12-05 16:59:32',0),(54,36,1,2,'2017-12-01 00:00:00',3750.0000,39,4,'ASM(39) d 12999 3750.00,c 12001 3750.00','2018-02-16 22:14:38',211,'2017-12-05 16:59:32',0),(55,3,1,4,'2016-07-01 00:00:00',8300.0000,2,0,'ASM(2) d 12999 8300.00,c 12001 8300.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(56,9,1,4,'2017-01-01 00:00:00',4150.0000,28,0,'ASM(28) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(57,12,1,4,'2017-02-01 00:00:00',4150.0000,29,0,'ASM(29) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(58,16,1,4,'2017-03-01 00:00:00',4150.0000,30,0,'ASM(30) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(59,18,1,4,'2017-04-01 00:00:00',4150.0000,31,0,'ASM(31) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(60,20,1,4,'2017-05-01 00:00:00',4150.0000,32,0,'ASM(32) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(61,23,1,4,'2017-06-01 00:00:00',4150.0000,33,0,'ASM(33) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(62,25,1,4,'2017-07-01 00:00:00',4150.0000,34,0,'ASM(34) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(63,27,1,4,'2017-08-01 00:00:00',4150.0000,35,0,'ASM(35) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(64,30,1,4,'2017-09-01 00:00:00',4150.0000,36,0,'ASM(36) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(65,32,1,4,'2017-10-01 00:00:00',4150.0000,37,0,'ASM(37) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(66,34,1,4,'2017-11-01 00:00:00',4150.0000,38,0,'ASM(38) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(67,37,1,4,'2017-12-01 00:00:00',4150.0000,41,0,'ASM(41) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(68,38,1,0,'2017-02-03 00:00:00',628.4500,0,0,'d 10999 _, c 12999 _','2017-12-05 19:44:51',0,'2017-12-05 19:44:51',0),(69,39,1,0,'2016-10-03 00:00:00',4000.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 20:44:04',0,'2017-12-05 20:44:04',0),(70,39,1,3,'2016-10-03 00:00:00',4000.0000,59,0,'ASM(59) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(71,13,1,3,'2016-11-11 00:00:00',4000.0000,47,0,'ASM(47) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(72,13,1,3,'2016-12-01 00:00:00',4000.0000,48,0,'ASM(48) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(73,13,1,3,'2017-01-01 00:00:00',4000.0000,16,0,'ASM(16) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(74,10,1,3,'2017-02-03 00:00:00',628.4500,42,0,'ASM(42) d 12999 628.45,c 12001 628.45','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(75,10,1,3,'2017-02-13 00:00:00',4000.0000,17,0,'ASM(17) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(76,10,1,3,'2017-02-28 00:00:00',175.0000,43,0,'ASM(43) d 12999 175.00,c 12001 175.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(77,10,1,3,'2017-03-01 00:00:00',3546.5500,18,0,'ASM(18) d 12999 3546.55,c 12001 3546.55','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(78,38,1,3,'2017-03-01 00:00:00',453.4500,18,0,'ASM(18) d 12999 453.45,c 12001 453.45','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(79,38,1,3,'2017-03-31 00:00:00',175.0000,44,0,'ASM(44) d 12999 175.00,c 12001 175.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(80,21,1,3,'2017-05-12 00:00:00',4000.0000,19,0,'ASM(19) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(81,21,1,3,'2017-05-12 00:00:00',350.0000,50,4,'ASM(50) d 12999 350.00,c 12001 350.00','2018-02-15 21:49:51',211,'2017-12-05 20:50:21',0),(82,21,1,3,'2017-05-12 00:00:00',81.7900,45,0,'ASM(45) d 12999 81.79,c 12001 81.79','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(83,21,1,3,'2017-05-12 00:00:00',4000.0000,20,0,'ASM(20) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(84,21,1,3,'2017-05-12 00:00:00',350.0000,51,4,'ASM(51) d 12999 350.00,c 12001 350.00','2018-02-15 21:49:51',211,'2017-12-05 20:50:21',0),(85,21,1,3,'2017-06-01 00:00:00',4000.0000,21,0,'ASM(21) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(86,21,1,3,'2017-06-01 00:00:00',350.0000,52,4,'ASM(52) d 12999 350.00,c 12001 350.00','2018-02-15 21:49:51',211,'2017-12-05 20:50:21',0),(87,28,1,3,'2017-08-15 00:00:00',4000.0000,22,0,'ASM(22) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(88,28,1,3,'2017-08-15 00:00:00',350.0000,53,4,'ASM(53) d 12999 350.00,c 12001 350.00','2018-02-15 21:49:51',211,'2017-12-05 20:50:21',0),(89,28,1,3,'2017-08-15 00:00:00',4000.0000,23,0,'ASM(23) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(90,28,1,3,'2017-08-15 00:00:00',350.0000,54,4,'ASM(54) d 12999 350.00,c 12001 350.00','2018-02-15 21:49:51',211,'2017-12-05 20:50:21',0),(91,28,1,3,'2017-09-01 00:00:00',4000.0000,24,0,'ASM(24) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(92,28,1,3,'2017-09-01 00:00:00',350.0000,55,4,'ASM(55) d 12999 350.00,c 12001 350.00','2018-02-15 21:49:51',211,'2017-12-05 20:50:22',0),(93,35,1,3,'2017-11-14 00:00:00',4000.0000,25,0,'ASM(25) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(94,35,1,3,'2017-11-14 00:00:00',350.0000,56,4,'ASM(56) d 12999 350.00,c 12001 350.00','2018-02-15 21:49:51',211,'2017-12-05 20:50:22',0),(95,35,1,3,'2017-11-14 00:00:00',409.2800,46,0,'ASM(46) d 12999 409.28,c 12001 409.28','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(96,35,1,3,'2017-11-14 00:00:00',4000.0000,26,0,'ASM(26) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(97,35,1,3,'2017-11-14 00:00:00',350.0000,57,4,'ASM(57) d 12999 350.00,c 12001 350.00','2018-02-15 21:49:51',211,'2017-12-05 20:50:22',0),(98,35,1,3,'2017-12-01 00:00:00',4000.0000,40,0,'ASM(40) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(99,35,1,3,'2017-12-01 00:00:00',350.0000,58,4,'ASM(58) d 12999 350.00,c 12001 350.00','2018-02-15 21:49:51',211,'2017-12-05 20:50:22',0),(100,21,1,3,'2018-02-15 21:49:52',-350.0000,50,4,'ASM(50) d 12999 -350.00,ASM(50) c 12001 -350.00','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(101,21,1,3,'2018-02-15 21:49:52',-350.0000,51,4,'ASM(51) d 12999 -350.00,ASM(51) c 12001 -350.00','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(102,21,1,3,'2018-02-15 21:49:52',-350.0000,52,4,'ASM(52) d 12999 -350.00,ASM(52) c 12001 -350.00','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(103,28,1,3,'2018-02-15 21:49:52',-350.0000,53,4,'ASM(53) d 12999 -350.00,ASM(53) c 12001 -350.00','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(104,28,1,3,'2018-02-15 21:49:52',-350.0000,54,4,'ASM(54) d 12999 -350.00,ASM(54) c 12001 -350.00','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(105,28,1,3,'2018-02-15 21:49:52',-350.0000,55,4,'ASM(55) d 12999 -350.00,ASM(55) c 12001 -350.00','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(106,35,1,3,'2018-02-15 21:49:52',-350.0000,56,4,'ASM(56) d 12999 -350.00,ASM(56) c 12001 -350.00','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(107,35,1,3,'2018-02-15 21:49:52',-350.0000,57,4,'ASM(57) d 12999 -350.00,ASM(57) c 12001 -350.00','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(108,35,1,3,'2018-02-15 21:49:52',-350.0000,58,4,'ASM(58) d 12999 -350.00,ASM(58) c 12001 -350.00','2018-02-15 21:49:51',211,'2018-02-15 21:49:51',211),(109,7,1,2,'2018-02-16 22:14:38',-3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(110,14,1,2,'2018-02-16 22:14:38',3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(111,8,1,2,'2018-02-16 22:14:38',-3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(112,7,1,2,'2018-02-16 22:14:38',-3750.0000,4,4,'ASM(4) d 12999 3750.00,ASM(4) c 12001 3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(113,14,1,2,'2018-02-16 22:14:38',3750.0000,4,4,'ASM(4) d 12999 3750.00,ASM(4) c 12001 3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(114,8,1,2,'2018-02-16 22:14:38',-3750.0000,4,4,'ASM(4) d 12999 3750.00,ASM(4) c 12001 3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(115,7,1,2,'2018-02-16 22:14:38',3750.0000,4,4,'ASM(4) d 12999 3750.00,ASM(4) c 12001 3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(116,14,1,2,'2018-02-16 22:14:38',-3750.0000,4,4,'ASM(4) d 12999 3750.00,ASM(4) c 12001 3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(117,8,1,2,'2018-02-16 22:14:38',3750.0000,4,4,'ASM(4) d 12999 3750.00,ASM(4) c 12001 3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(118,7,1,2,'2018-02-16 22:14:38',-3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(119,14,1,2,'2018-02-16 22:14:38',3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(120,8,1,2,'2018-02-16 22:14:38',-3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(121,7,1,2,'2018-02-16 22:14:38',3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(122,14,1,2,'2018-02-16 22:14:38',-3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(123,8,1,2,'2018-02-16 22:14:38',3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(124,7,1,2,'2018-02-16 22:14:38',3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(125,14,1,2,'2018-02-16 22:14:38',-3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(126,8,1,2,'2018-02-16 22:14:38',3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(127,7,1,2,'2018-02-16 22:14:38',-3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(128,14,1,2,'2018-02-16 22:14:38',3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(129,8,1,2,'2018-02-16 22:14:38',-3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(130,11,1,2,'2018-02-16 22:14:38',-3750.0000,5,4,'ASM(5) d 12999 -3750.00,ASM(5) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(131,15,1,2,'2018-02-16 22:14:38',-3750.0000,6,4,'ASM(6) d 12999 -3750.00,ASM(6) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(132,17,1,2,'2018-02-16 22:14:38',-3750.0000,7,4,'ASM(7) d 12999 -3750.00,ASM(7) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(133,19,1,2,'2018-02-16 22:14:38',-3750.0000,8,4,'ASM(8) d 12999 -3750.00,ASM(8) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(134,22,1,2,'2018-02-16 22:14:38',-3750.0000,9,4,'ASM(9) d 12999 -3750.00,ASM(9) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(135,24,1,2,'2018-02-16 22:14:38',-3750.0000,10,4,'ASM(10) d 12999 -3750.00,ASM(10) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(136,26,1,2,'2018-02-16 22:14:38',-3750.0000,11,4,'ASM(11) d 12999 -3750.00,ASM(11) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(137,29,1,2,'2018-02-16 22:14:38',-3750.0000,12,4,'ASM(12) d 12999 -3750.00,ASM(12) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(138,31,1,2,'2018-02-16 22:14:38',-3750.0000,13,4,'ASM(13) d 12999 -3750.00,ASM(13) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(139,33,1,2,'2018-02-16 22:14:38',-3750.0000,14,4,'ASM(14) d 12999 -3750.00,ASM(14) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211),(140,36,1,2,'2018-02-16 22:14:38',-3750.0000,39,4,'ASM(39) d 12999 -3750.00,ASM(39) c 12001 -3750.00','2018-02-16 22:14:38',211,'2018-02-16 22:14:38',211);
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
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Rentable`
--

LOCK TABLES `Rentable` WRITE;
/*!40000 ALTER TABLE `Rentable` DISABLE KEYS */;
INSERT INTO `Rentable` VALUES (1,1,'309 Rexford',1,0,'0000-00-00 00:00:00','2018-02-16 22:12:59',211,'2017-11-28 03:52:45',0,''),(2,1,'309 1/2 Rexford',1,0,'0000-00-00 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0,''),(3,1,'311 Rexford',1,0,'0000-00-00 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0,''),(4,1,'311 1/2 Rexford',1,0,'0000-00-00 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0,''),(5,2,'Nickelodeon Office Building',0,0,'0000-00-00 00:00:00','2018-01-18 08:19:29',267,'2018-01-17 23:09:38',200,''),(6,2,'Nickelodeon Office Garage',0,0,'0000-00-00 00:00:00','2018-01-17 23:23:58',200,'2018-01-17 23:23:58',200,''),(7,4,'Nickelodeon Animation Building',0,0,'0000-00-00 00:00:00','2018-01-25 20:30:05',200,'2018-01-25 20:30:05',200,''),(8,4,'Nicelodeon Animation Garage',0,0,'0000-00-00 00:00:00','2018-01-25 20:30:24',200,'2018-01-25 20:30:24',200,''),(9,3,'PAC Unit A',0,0,'0000-00-00 00:00:00','2018-01-25 21:16:41',200,'2018-01-25 21:16:41',200,''),(10,3,'PAC Unit B',0,0,'0000-00-00 00:00:00','2018-01-25 21:17:24',200,'2018-01-25 21:17:24',200,''),(11,3,'PAC Unit C',0,0,'0000-00-00 00:00:00','2018-01-25 21:17:55',200,'2018-01-25 21:17:55',200,''),(12,3,'PAC Unit D',0,0,'0000-00-00 00:00:00','2018-01-25 21:18:23',200,'2018-01-25 21:18:23',200,''),(13,3,'NHV Unit E',0,0,'0000-00-00 00:00:00','2018-01-25 21:19:11',200,'2018-01-25 21:19:11',200,''),(14,3,'BVW Nursery',0,0,'0000-00-00 00:00:00','2018-01-25 21:46:00',211,'2018-01-25 21:20:08',200,''),(15,6,'Summitridge Rehab',1,0,'0000-00-00 00:00:00','2018-01-25 21:31:13',200,'2018-01-25 21:30:32',211,''),(16,5,'Unit A',1,0,'0000-00-00 00:00:00','2018-01-25 22:17:26',200,'2018-01-25 22:17:26',200,''),(17,5,'Unit B',1,0,'0000-00-00 00:00:00','2018-01-25 22:17:34',200,'2018-01-25 22:17:34',200,''),(18,5,'Unit C',1,0,'0000-00-00 00:00:00','2018-01-25 22:17:41',200,'2018-01-25 22:17:41',200,''),(19,5,'Unit D',1,0,'0000-00-00 00:00:00','2018-01-25 22:17:55',200,'2018-01-25 22:17:55',200,''),(20,5,'Units C-D',1,0,'0000-00-00 00:00:00','2018-01-25 22:20:39',200,'2018-01-25 22:20:39',200,''),(21,5,'Unit A Property Tax Reimbursement',0,0,'0000-00-00 00:00:00','2018-02-14 21:59:46',200,'2018-02-14 21:59:46',200,''),(22,5,'Unit A Property Tax Reimbursement',0,0,'0000-00-00 00:00:00','2018-02-14 21:59:51',200,'2018-02-14 21:59:51',200,''),(23,5,'Unit A Property Tax Reimbursement',2,0,'0000-00-00 00:00:00','2018-02-14 22:00:34',200,'2018-02-14 21:59:52',200,'');
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RMRID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableMarketRate`
--

LOCK TABLES `RentableMarketRate` WRITE;
/*!40000 ALTER TABLE `RentableMarketRate` DISABLE KEYS */;
INSERT INTO `RentableMarketRate` VALUES (1,1,1,3750.0000,'2014-01-01 00:00:00','9999-03-01 00:00:00','2018-01-10 18:32:37',0,'2017-11-28 03:44:18',0),(2,2,1,4000.0000,'2014-01-01 00:00:00','9999-05-01 00:00:00','2018-01-18 09:09:22',267,'2017-11-28 03:44:18',0),(3,3,1,4150.0000,'2014-01-01 00:00:00','9999-04-01 00:00:00','2018-01-10 18:32:37',0,'2017-11-28 03:44:18',0),(4,4,1,2500.0000,'2014-01-01 00:00:00','9999-01-01 00:00:00','2018-01-10 18:32:37',0,'2017-11-28 03:44:18',0);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
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
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableStatus`
--

LOCK TABLES `RentableStatus` WRITE;
/*!40000 ALTER TABLE `RentableStatus` DISABLE KEYS */;
INSERT INTO `RentableStatus` VALUES (1,1,1,1,0,'2014-01-01 00:00:00','9000-01-31 00:00:00','1900-01-01','2018-02-16 22:13:14',211,'2017-11-28 03:52:45',0),(2,2,1,1,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0),(3,3,1,1,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0),(4,4,1,4,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0),(7,5,2,1,0,'2018-01-17 00:00:00','2036-09-01 00:00:00','0000-00-00','2018-01-18 08:19:29',267,'2018-01-18 08:19:29',267),(8,6,2,1,0,'2018-01-01 00:00:00','2036-09-01 00:00:00','0000-00-00','2018-01-24 20:50:44',200,'2018-01-24 20:50:44',200),(9,7,4,1,0,'2018-01-25 20:30:05','2036-09-01 00:00:00','0000-00-00','2018-01-25 20:30:05',200,'2018-01-25 20:30:05',200),(11,8,4,1,0,'2015-04-01 00:00:00','2036-09-01 00:00:00','1900-01-01','2018-02-14 21:50:55',200,'2018-01-25 20:34:09',200),(12,9,3,1,0,'2018-01-25 21:16:41','9999-01-01 00:00:00','0000-00-00','2018-01-25 21:16:41',200,'2018-01-25 21:16:41',200),(13,10,3,1,0,'2018-01-25 21:17:24','9999-01-01 00:00:00','0000-00-00','2018-01-25 21:17:24',200,'2018-01-25 21:17:24',200),(14,11,3,1,0,'2018-01-25 21:17:55','9999-01-01 00:00:00','0000-00-00','2018-01-25 21:17:55',200,'2018-01-25 21:17:55',200),(15,12,3,1,0,'2018-01-25 21:18:23','9999-01-01 00:00:00','0000-00-00','2018-01-25 21:18:23',200,'2018-01-25 21:18:23',200),(16,13,3,1,0,'2018-01-25 21:19:12','9999-01-01 00:00:00','0000-00-00','2018-01-25 21:19:11',200,'2018-01-25 21:19:11',200),(19,14,3,1,0,'2010-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2018-01-25 21:23:04',200,'2018-01-25 21:23:04',200),(21,15,6,1,0,'2018-01-25 00:00:00','9999-01-01 00:00:00','0000-00-00','2018-01-25 21:31:04',200,'2018-01-25 21:31:04',200),(22,16,5,1,0,'2018-01-25 22:17:26','9999-01-01 00:00:00','0000-00-00','2018-01-25 22:17:26',200,'2018-01-25 22:17:26',200),(23,17,5,1,0,'2018-01-25 22:17:34','9999-01-01 00:00:00','0000-00-00','2018-01-25 22:17:34',200,'2018-01-25 22:17:34',200),(26,20,5,1,0,'2018-01-25 22:20:39','9999-01-01 00:00:00','0000-00-00','2018-01-25 22:20:39',200,'2018-01-25 22:20:39',200),(27,18,5,5,0,'2018-01-25 00:00:00','2018-01-25 22:17:41','0000-00-00','2018-01-25 22:20:57',200,'2018-01-25 22:20:57',200),(28,18,5,1,0,'2018-01-25 22:17:41','9999-01-01 00:00:00','0000-00-00','2018-01-25 22:20:57',200,'2018-01-25 22:20:57',200),(33,19,5,5,0,'2010-01-01 00:00:00','2010-01-25 00:00:00','0000-00-00','2018-01-25 22:22:10',200,'2018-01-25 22:22:10',200),(34,19,5,1,0,'2010-01-25 00:00:00','9999-01-01 00:00:00','0000-00-00','2018-01-25 22:22:10',200,'2018-01-25 22:22:10',200);
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
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableTypeRef`
--

LOCK TABLES `RentableTypeRef` WRITE;
/*!40000 ALTER TABLE `RentableTypeRef` DISABLE KEYS */;
INSERT INTO `RentableTypeRef` VALUES (1,1,1,1,0,0,'2014-01-01 00:00:00','9000-01-31 00:00:00','2018-02-16 22:12:59',211,'2017-11-28 03:52:45',0),(2,2,1,2,0,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0),(3,3,1,3,0,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0),(4,4,1,4,0,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0),(7,5,2,5,0,0,'2018-01-17 00:00:00','2036-09-01 00:00:00','2018-01-18 08:19:29',267,'2018-01-18 08:19:29',267),(8,6,2,6,0,0,'2018-01-01 00:00:00','2036-09-01 00:00:00','2018-01-24 20:50:44',200,'2018-01-24 20:50:44',200),(9,7,4,7,0,0,'2018-01-25 20:30:05','2036-09-01 00:00:00','2018-01-25 20:30:05',200,'2018-01-25 20:30:05',200),(11,8,4,8,0,0,'2015-04-01 00:00:00','2036-09-01 00:00:00','2018-01-25 20:34:09',200,'2018-01-25 20:34:09',200),(12,9,3,10,0,0,'2018-01-25 21:16:41','9999-01-01 00:00:00','2018-01-25 21:16:41',200,'2018-01-25 21:16:41',200),(13,10,3,10,0,0,'2018-01-25 21:17:24','9999-01-01 00:00:00','2018-01-25 21:17:24',200,'2018-01-25 21:17:24',200),(14,11,3,10,0,0,'2018-01-25 21:17:55','9999-01-01 00:00:00','2018-01-25 21:17:55',200,'2018-01-25 21:17:55',200),(15,12,3,10,0,0,'2018-01-25 21:18:23','9999-01-01 00:00:00','2018-01-25 21:18:23',200,'2018-01-25 21:18:23',200),(16,13,3,10,0,0,'2018-01-25 21:19:12','9999-01-01 00:00:00','2018-01-25 21:19:11',200,'2018-01-25 21:19:11',200),(27,15,6,9,0,0,'2018-01-25 00:00:00','9999-01-01 00:00:00','2018-01-25 21:31:04',200,'2018-01-25 21:31:04',200),(28,14,3,11,0,0,'2018-01-25 00:00:00','2018-01-25 21:20:09','2018-01-25 21:46:00',211,'2018-01-25 21:46:00',211),(30,16,5,12,0,0,'2018-01-25 22:17:26','9999-01-01 00:00:00','2018-01-25 22:17:26',200,'2018-01-25 22:17:26',200),(31,17,5,12,0,0,'2018-01-25 22:17:34','9999-01-01 00:00:00','2018-01-25 22:17:34',200,'2018-01-25 22:17:34',200),(34,20,5,12,0,0,'2018-01-25 22:20:39','9999-01-01 00:00:00','2018-01-25 22:20:39',200,'2018-01-25 22:20:39',200),(35,18,5,12,0,0,'2018-01-25 00:00:00','9999-01-01 00:00:00','2018-01-25 22:20:57',200,'2018-01-25 22:20:57',200),(37,19,5,12,0,0,'2010-01-01 00:00:00','9999-01-01 00:00:00','2018-01-25 22:21:46',200,'2018-01-25 22:21:46',200);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
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
  `ARID` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RTID`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableTypes`
--

LOCK TABLES `RentableTypes` WRITE;
/*!40000 ALTER TABLE `RentableTypes` DISABLE KEYS */;
INSERT INTO `RentableTypes` VALUES (1,1,'Rex1','309 Rexford',6,4,4,1,0,0,'2017-11-28 03:44:18',0,'2017-11-28 03:44:18',0),(2,1,'Rex2','309 1/2 Rexford',6,4,4,1,0,0,'2018-01-18 09:09:22',267,'2017-11-28 03:44:18',0),(3,1,'Rex3','311 Rexford',6,4,4,1,0,0,'2017-11-28 03:44:18',0,'2017-11-28 03:44:18',0),(4,1,'Rex4','311 1/2 Rexford',6,4,4,1,0,0,'2017-11-28 03:44:18',0,'2017-11-28 03:44:18',0),(5,2,'Office Building','Nickelodeon Office',6,4,6,0,0,0,'2018-01-24 20:49:23',200,'2018-01-17 22:37:07',200),(6,2,'Garage','Nickelodeon Garage',6,4,6,0,0,0,'2018-01-24 20:49:44',200,'2018-01-17 22:39:39',200),(7,4,'Office Building','Animation Office',6,4,6,0,0,0,'2018-01-24 21:33:56',200,'2018-01-24 21:33:56',200),(8,4,'Garage','Animation Garage',6,4,6,0,0,0,'2018-01-24 21:34:20',200,'2018-01-24 21:34:20',200),(9,6,'Commercial Rehabilitation Facility','Rehabilitation Facility',6,4,6,0,0,0,'2018-01-25 20:58:33',200,'2018-01-25 20:58:33',200),(10,3,'Condominium Warehouse Unit','Condominium Warehouse Unit',6,4,6,0,0,0,'2018-01-25 21:15:36',200,'2018-01-25 21:15:36',200),(11,3,'Tenants in Common','Tenants in Common',6,4,6,0,0,0,'2018-01-25 21:43:29',200,'2018-01-25 21:20:50',200),(12,5,'Condominium Warehouse','Condominium Warehouse',6,4,6,0,0,0,'2018-01-25 22:16:33',200,'2018-01-25 22:16:33',200);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RUID`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableUsers`
--

LOCK TABLES `RentableUsers` WRITE;
/*!40000 ALTER TABLE `RentableUsers` DISABLE KEYS */;
INSERT INTO `RentableUsers` VALUES (1,1,1,1,'2014-03-01','2018-02-01','2018-02-14 20:50:32',200,'2017-11-30 18:22:29',0),(2,1,1,2,'2014-03-01','2018-02-01','2018-02-14 20:50:40',200,'2017-11-30 18:23:03',0),(4,1,1,6,'2014-03-01','2018-02-01','2018-02-14 20:50:47',200,'2017-11-30 18:30:18',0),(5,2,1,3,'2016-10-01','2018-01-01','2018-01-10 18:32:36',0,'2017-11-30 18:33:01',0),(6,3,1,4,'2016-07-01','2018-07-01','2018-01-10 18:32:36',0,'2017-11-30 18:36:10',0),(7,3,1,5,'2016-07-01','2018-07-01','2018-01-10 18:32:36',0,'2017-11-30 18:36:30',0),(8,5,2,7,'2016-09-01','2036-09-01','2018-01-17 23:16:51',200,'2018-01-17 23:16:51',200);
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
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreement`
--

LOCK TABLES `RentalAgreement` WRITE;
/*!40000 ALTER TABLE `RentalAgreement` DISABLE KEYS */;
INSERT INTO `RentalAgreement` VALUES (2,0,1,0,'2014-03-01','2018-02-01','2014-03-01','2018-02-01','2014-03-01','2018-02-01','2014-03-01',0,0,2,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-02-14 20:50:47',200,'2017-11-30 18:17:55',0),(3,0,1,0,'2016-10-01','2018-01-01','2016-10-01','2018-01-01','2016-10-01','2018-01-01','2016-10-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2017-11-30 18:33:53',0,'2017-11-30 18:29:25',0),(4,0,1,0,'2016-07-01','2018-07-01','2016-07-01','2018-07-01','2016-07-01','2018-07-01','2016-07-01',0,0,2,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2017-11-30 18:37:13',0,'2017-11-30 18:33:59',0),(6,0,1,0,'2018-01-03','2019-01-03','2018-01-03','2019-01-03','2018-01-03','2019-01-03','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-03 07:07:58',0,'2018-01-03 07:07:58',0),(7,0,1,0,'2018-01-03','2019-01-03','2018-01-03','2019-01-03','2018-01-03','2019-01-03','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-03 07:08:22',0,'2018-01-03 07:08:22',0),(8,0,2,0,'2015-04-29','2036-09-01','2016-09-01','2036-09-01','2016-09-01','2036-09-01','2018-01-01',0,0,2,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-17 23:17:32',200,'2018-01-17 22:26:30',200),(9,0,2,0,'2018-01-24','2019-01-24','2018-01-24','2019-01-24','2018-01-24','2019-01-24','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-24 20:58:06',200,'2018-01-24 20:58:06',200),(10,0,2,0,'2018-01-24','2019-01-24','2018-01-24','2019-01-24','2018-01-24','2019-01-24','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-24 21:12:53',200,'2018-01-24 21:12:53',200),(11,0,4,0,'2015-04-01','2036-09-01','2015-04-01','2036-09-01','2018-01-01','2018-02-01','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-02-14 21:50:08',200,'2018-01-25 20:30:36',200),(12,0,4,0,'2015-04-01','2036-09-01','2015-04-01','2036-09-01','2015-04-01','2036-09-01','2018-01-01',0,0,2,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-25 20:35:54',200,'2018-01-25 20:34:31',200),(13,0,4,0,'2018-01-25','2019-01-25','2018-01-25','2019-01-25','2018-01-25','2019-01-25','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-25 20:40:06',200,'2018-01-25 20:40:06',200),(14,0,6,0,'2018-01-25','2019-01-25','2018-01-25','2019-01-25','2018-01-25','2019-01-25','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-25 21:31:40',211,'2018-01-25 21:31:40',211),(18,0,5,0,'2007-12-24','2022-12-23','2007-12-24','2022-12-23','2007-12-24','2022-12-23','2018-01-01',0,0,2,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-25 22:29:26',200,'2018-01-25 22:25:58',200),(19,0,5,0,'2015-06-01','2025-05-31','2015-06-01','2025-05-31','2015-06-01','2025-05-31','2018-01-01',0,0,2,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-25 22:31:50',200,'2018-01-25 22:29:34',200),(20,0,5,0,'2009-06-05','2019-06-04','2009-06-05','2019-06-04','2009-06-05','2019-06-04','2018-01-01',0,0,2,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-25 22:37:11',200,'2018-01-25 22:31:58',200),(21,0,3,0,'2010-01-04','9999-01-01','2010-01-04','9999-01-01','2010-01-04','9999-01-01','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-31 21:54:52',200,'2018-01-31 21:51:12',200),(22,0,3,0,'2010-01-01','9999-01-01','2010-01-01','9999-01-01','2010-01-01','9999-01-01','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-31 21:56:28',200,'2018-01-31 21:54:58',200),(23,0,3,0,'2010-01-01','9999-01-01','2010-01-01','9999-01-01','2010-01-01','9999-01-01','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-31 21:58:23',200,'2018-01-31 21:56:32',200),(24,0,3,0,'2010-01-01','9999-01-01','2010-01-01','9999-01-01','2010-01-01','9999-01-01','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-31 22:00:58',200,'2018-01-31 21:59:51',200),(25,0,3,0,'2010-01-01','9999-01-01','2010-01-01','9999-01-01','2010-01-01','9999-01-01','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-31 22:03:50',200,'2018-01-31 22:01:03',200),(26,0,3,0,'2010-01-01','9999-01-01','2010-01-01','9999-01-01','2010-01-01','9999-01-01','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-31 22:16:48',200,'2018-01-31 22:09:50',200),(27,0,3,0,'2018-01-31','2019-01-31','2018-01-31','2019-01-31','2018-01-31','2019-01-31','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-31 22:10:59',200,'2018-01-31 22:10:59',200);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RAPID`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementPayors`
--

LOCK TABLES `RentalAgreementPayors` WRITE;
/*!40000 ALTER TABLE `RentalAgreementPayors` DISABLE KEYS */;
INSERT INTO `RentalAgreementPayors` VALUES (1,2,1,1,'2014-03-01','2018-02-01',0,'2018-02-14 20:50:24',200,'2017-11-30 18:21:00',0),(2,2,1,2,'2017-11-30','2018-03-01',0,'2018-01-10 18:32:35',0,'2017-11-30 18:21:57',0),(3,2,1,2,'2018-03-01','2018-03-01',0,'2018-01-10 18:32:35',0,'2017-11-30 18:28:09',0),(4,3,1,3,'2016-10-01','2018-01-01',0,'2018-01-10 18:32:35',0,'2017-11-30 18:32:33',0),(5,4,1,4,'2016-07-01','2018-02-01',0,'2018-01-10 18:32:35',0,'2017-11-30 18:35:19',0),(6,4,1,5,'2016-07-01','2018-07-01',0,'2018-01-10 18:32:35',0,'2017-11-30 18:35:41',0),(7,8,2,8,'2016-09-01','2016-09-01',0,'2018-01-17 23:15:55',200,'2018-01-17 23:15:55',200),(8,8,2,8,'2018-01-01','2036-09-01',0,'2018-01-24 20:52:30',200,'2018-01-24 20:52:30',200),(9,14,6,11,'2018-01-25','2019-01-25',0,'2018-01-25 21:37:28',211,'2018-01-25 21:37:28',211),(10,18,5,12,'2017-12-24','2022-12-23',0,'2018-01-25 22:28:34',200,'2018-01-25 22:28:34',200),(11,19,5,13,'2015-06-01','2025-05-31',0,'2018-01-25 22:30:54',200,'2018-01-25 22:30:54',200),(12,20,5,14,'2009-06-05','2019-06-04',0,'2018-01-25 22:36:21',200,'2018-01-25 22:36:21',200),(13,21,3,16,'2010-01-01','9999-01-01',0,'2018-01-31 21:54:38',200,'2018-01-31 21:54:38',200),(14,22,3,15,'2018-01-31','2019-01-31',0,'2018-01-31 21:56:16',200,'2018-01-31 21:56:16',200),(15,23,3,15,'2018-01-31','9999-01-01',0,'2018-01-31 21:58:13',200,'2018-01-31 21:58:13',200),(16,25,3,15,'2018-01-31','9999-01-01',0,'2018-01-31 22:03:20',200,'2018-01-31 22:03:20',200),(17,24,3,15,'2018-01-01','9999-01-01',0,'2018-01-31 22:04:15',200,'2018-01-31 22:04:15',200),(18,26,3,17,'2010-01-01','9999-01-01',0,'2018-01-31 22:18:24',200,'2018-01-31 22:18:24',200),(19,11,4,10,'2018-01-01','2036-09-01',0,'2018-02-14 21:48:52',200,'2018-02-14 21:48:52',200);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RARID`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementRentables`
--

LOCK TABLES `RentalAgreementRentables` WRITE;
/*!40000 ALTER TABLE `RentalAgreementRentables` DISABLE KEYS */;
INSERT INTO `RentalAgreementRentables` VALUES (1,2,1,1,0,3750.0000,'2014-03-01','2018-02-01','2018-02-14 20:50:14',200,'2017-11-30 18:20:15',0),(2,3,1,2,0,4000.0000,'2016-10-01','2018-01-01','2018-01-10 18:32:35',0,'2017-11-30 18:32:13',0),(3,4,1,3,0,4150.0000,'2016-07-01','2018-07-01','2018-01-10 18:32:35',0,'2017-11-30 18:34:33',0),(4,8,2,5,0,192924.9200,'2018-01-01','2036-09-01','2018-01-17 23:13:37',200,'2018-01-17 23:13:37',200),(5,8,2,6,0,34501.5000,'2018-01-01','2036-09-01','2018-01-24 20:51:54',200,'2018-01-24 20:51:54',200),(6,11,4,7,0,144670.0000,'2015-04-01','2036-09-01','2018-01-25 20:32:46',200,'2018-01-25 20:32:46',200),(7,12,4,7,0,144670.0000,'2015-04-01','2036-09-01','2018-01-25 20:35:16',200,'2018-01-25 20:35:16',200),(8,14,6,15,0,30692.0000,'2018-01-25','2019-01-25','2018-01-25 21:34:06',211,'2018-01-25 21:34:06',211),(9,18,5,16,0,135146.4500,'2007-12-24','2022-12-23','2018-01-25 22:27:39',200,'2018-01-25 22:27:39',200),(10,19,5,17,0,47289.0000,'2015-06-01','2025-05-31','2018-01-25 22:30:31',200,'2018-01-25 22:30:31',200),(11,20,5,20,0,58422.6000,'2009-06-05','2019-06-04','2018-01-25 22:35:49',200,'2018-01-25 22:35:49',200),(12,21,3,13,0,6857.0100,'2010-01-01','9999-01-01','2018-01-31 21:53:59',200,'2018-01-31 21:53:59',200),(13,22,3,9,0,444.0500,'2018-01-31','2019-01-31','2018-01-31 21:55:58',200,'2018-01-31 21:55:58',200),(14,23,3,10,0,3932.9900,'2018-01-31','2018-12-31','2018-01-31 21:57:33',200,'2018-01-31 21:57:33',200),(15,24,3,11,0,1533.5700,'2018-01-31','2018-12-31','2018-01-31 22:00:53',200,'2018-01-31 22:00:53',200),(16,25,3,12,0,2818.6200,'2018-01-01','2018-12-31','2018-01-31 22:02:52',200,'2018-01-31 22:02:52',200),(18,26,3,14,0,601.2300,'2018-01-01','2018-12-31','2018-01-31 22:16:18',200,'2018-01-31 22:16:18',200),(19,2,1,3,0,3750.0000,'2014-03-01','2108-02-01','2018-02-20 21:27:16',211,'2018-02-20 21:27:16',211);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementTemplate`
--

LOCK TABLES `RentalAgreementTemplate` WRITE;
/*!40000 ALTER TABLE `RentalAgreementTemplate` DISABLE KEYS */;
INSERT INTO `RentalAgreementTemplate` VALUES (1,1,'Agreement3722.1.doc','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(2,1,'Agreement4421.2.doc','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(3,1,'Agreement4980.3.doc','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(4,1,'Agreement5342.7.doc','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SubAR`
--

LOCK TABLES `SubAR` WRITE;
/*!40000 ALTER TABLE `SubAR` DISABLE KEYS */;
INSERT INTO `SubAR` VALUES (1,15,14,1,'2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(2,37,36,1,'2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TWS`
--

LOCK TABLES `TWS` WRITE;
/*!40000 ALTER TABLE `TWS` DISABLE KEYS */;
INSERT INTO `TWS` VALUES (2,'CreateAssessmentInstances','','CreateAssessmentInstances','2018-02-16 00:00:00','ip-172-31-51-141.ec2.internal',4,'2018-02-15 00:00:04','2018-02-15 00:00:04','2017-11-27 21:24:27','2018-02-15 00:00:03'),(3,'CleanRARBalanceCache','','CleanRARBalanceCache','2018-02-15 18:36:18','ip-172-31-51-141.ec2.internal',4,'2018-02-15 18:31:18','2018-02-15 18:31:18','2017-11-30 17:39:57','2018-02-15 18:31:18'),(4,'CleanSecDepBalanceCache','','CleanSecDepBalanceCache','2018-02-15 18:39:29','ip-172-31-51-141.ec2.internal',4,'2018-02-15 18:34:29','2018-02-15 18:34:29','2017-11-30 17:39:57','2018-02-15 18:34:28'),(5,'CleanAcctSliceCache','','CleanAcctSliceCache','2018-02-15 18:39:19','ip-172-31-51-141.ec2.internal',4,'2018-02-15 18:34:19','2018-02-15 18:34:19','2017-11-30 17:39:57','2018-02-15 18:34:18'),(6,'CleanARSliceCache','','CleanARSliceCache','2018-02-15 18:40:19','ip-172-31-51-141.ec2.internal',4,'2018-02-15 18:35:19','2018-02-15 18:35:19','2017-11-30 17:39:57','2018-02-15 18:35:18'),(17,'CreateAssessmentInstances','','CreateAssessmentInstances','2018-02-21 00:00:00','Steves-MacBook-Pro-2.local',4,'2018-02-20 20:36:10','2018-02-20 20:36:10','2018-02-15 14:01:24','2018-02-20 12:36:09'),(18,'CleanRARBalanceCache','','CleanRARBalanceCache','2018-02-20 21:30:43','Steves-MacBook-Pro-2.local',4,'2018-02-20 21:25:43','2018-02-20 21:25:43','2018-02-15 14:01:24','2018-02-20 13:25:43'),(19,'CleanSecDepBalanceCache','','CleanSecDepBalanceCache','2018-02-20 21:30:43','Steves-MacBook-Pro-2.local',4,'2018-02-20 21:25:43','2018-02-20 21:25:43','2018-02-15 14:01:24','2018-02-20 13:25:43'),(20,'CleanAcctSliceCache','','CleanAcctSliceCache','2018-02-20 21:30:43','Steves-MacBook-Pro-2.local',4,'2018-02-20 21:25:43','2018-02-20 21:25:43','2018-02-15 14:01:24','2018-02-20 13:25:43'),(21,'CleanARSliceCache','','CleanARSliceCache','2018-02-20 21:30:43','Steves-MacBook-Pro-2.local',4,'2018-02-20 21:25:43','2018-02-20 21:25:43','2018-02-15 14:01:24','2018-02-20 13:25:43');
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
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Transactant`
--

LOCK TABLES `Transactant` WRITE;
/*!40000 ALTER TABLE `Transactant` DISABLE KEYS */;
INSERT INTO `Transactant` VALUES (1,1,0,'Aaron','','Read','','',0,'','','','','','','','','','','','2017-11-30 18:15:20',0,'2017-11-30 18:15:20',0),(2,1,0,'Kirsten','','Read','','',0,'','','','','','','','','','','','2017-11-30 18:15:28',0,'2017-11-30 18:15:28',0),(3,1,0,'Alex','','Vahabzadeh','','Beaumont Partners LP',1,'','','','','','','','','','','','2017-11-30 18:17:13',0,'2017-11-30 18:16:10',0),(4,1,0,'Kevin','','Mills','','',0,'','','','','','','','','','','','2017-11-30 18:16:17',0,'2017-11-30 18:16:17',0),(5,1,0,'Lauren','','Beck','','',0,'','','','','','','','','','','','2017-11-30 18:16:28',0,'2017-11-30 18:16:28',0),(6,1,0,'Child','','Read','','',0,'','','','','','','','','','','','2017-11-30 18:24:52',0,'2017-11-30 18:24:52',0),(7,2,0,'','','','','Viacom, Inc.',1,'','','','','1515 Broadway','','New York','NY','10036','US','','2018-01-17 22:22:13',200,'2018-01-17 22:22:13',200),(8,2,0,'','','','','Viacom Media Networks',1,'','','','','1515 Broadway','','New YOrk','NY','10036','','','2018-01-17 22:23:47',200,'2018-01-17 22:23:47',200),(9,2,0,'','','','','Nickelodeon',1,'','','','','203 W Olive Avenue','','Burbank','CA','91502','','','2018-01-17 22:25:56',200,'2018-01-17 22:25:56',200),(10,4,0,'','','','','Viacom, Inc.',1,'','','','','','','','','','','','2018-01-25 20:39:56',200,'2018-01-25 20:39:56',200),(11,6,0,'','','','','Southern California Recovery',1,'','','','','','','','','','','','2018-01-25 21:36:56',211,'2018-01-25 21:36:56',211),(12,5,0,'','','','','HD Supply',1,'','','','','','','','','','','','2018-01-25 22:24:15',200,'2018-01-25 22:24:15',200),(13,5,0,'','','','','Technocel Sub to Purchasing 411 Inc.',1,'','','','','','','','','','','','2018-01-25 22:25:10',200,'2018-01-25 22:25:10',200),(14,5,0,'','','','','Illumination Dynamics Inc.',1,'','','','','','','','','','','','2018-01-25 22:25:45',200,'2018-01-25 22:25:45',200),(15,3,0,'','','','','Accord/PAC Members, LLC',1,'','','','','','','','','','','','2018-01-31 21:48:40',200,'2018-01-31 21:48:40',200),(16,3,0,'','','','','Dreadnaught Properties Los Angeles, LLC',1,'','','','','','','','','','','','2018-01-31 21:50:24',200,'2018-01-31 21:50:24',200),(17,3,0,'','','','','BrightView Companies, LLC',1,'','','','','','','','','','','','2018-01-31 22:17:45',200,'2018-01-31 22:17:45',200);
/*!40000 ALTER TABLE `Transactant` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `User` (
  `TCID` bigint(20) NOT NULL,
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
INSERT INTO `User` VALUES (1,1,0,'1900-01-01','','','','','',1,'',0,'2017-11-30 18:15:20',0,'2017-11-30 18:15:20',0),(2,1,0,'1900-01-01','','','','','',1,'',0,'2017-11-30 18:15:28',0,'2017-11-30 18:15:28',0),(3,1,0,'1900-01-01','','','','','',1,'',0,'2017-11-30 18:16:10',0,'2017-11-30 18:16:10',0),(4,1,0,'1900-01-01','','','','','',1,'',0,'2017-11-30 18:16:17',0,'2017-11-30 18:16:17',0),(5,1,0,'1900-01-01','','','','','',1,'',0,'2017-11-30 18:16:28',0,'2017-11-30 18:16:28',0),(6,1,0,'1900-01-01','','','','','',1,'',0,'2017-11-30 18:24:52',0,'2017-11-30 18:24:52',0),(7,2,0,'1900-01-01','','','','','',1,'',0,'2018-01-17 22:22:13',200,'2018-01-17 22:22:13',200),(8,2,0,'1900-01-01','','','','','',1,'',0,'2018-01-17 22:23:47',200,'2018-01-17 22:23:47',200),(9,2,0,'1900-01-01','','','','','',1,'',0,'2018-01-17 22:25:56',200,'2018-01-17 22:25:56',200),(10,4,0,'1900-01-01','','','','','',1,'',0,'2018-01-25 20:39:56',200,'2018-01-25 20:39:56',200),(11,6,0,'1900-01-01','','','','','',1,'',0,'2018-01-25 21:36:56',211,'2018-01-25 21:36:56',211),(12,5,0,'1900-01-01','','','','','',1,'',0,'2018-01-25 22:24:15',200,'2018-01-25 22:24:15',200),(13,5,0,'1900-01-01','','','','','',1,'',0,'2018-01-25 22:25:10',200,'2018-01-25 22:25:10',200),(14,5,0,'1900-01-01','','','','','',1,'',0,'2018-01-25 22:25:45',200,'2018-01-25 22:25:45',200),(15,3,0,'1900-01-01','','','','','',1,'',0,'2018-01-31 21:48:40',200,'2018-01-31 21:48:40',200),(16,3,0,'1900-01-01','','','','','',1,'',0,'2018-01-31 21:50:24',200,'2018-01-31 21:50:24',200),(17,3,0,'1900-01-01','','','','','',1,'',0,'2018-01-31 22:17:45',200,'2018-01-31 22:17:45',200);
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

-- Dump completed on 2018-05-16 14:44:38
