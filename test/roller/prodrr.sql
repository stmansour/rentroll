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
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `AR`
--

LOCK TABLES `AR` WRITE;
/*!40000 ALTER TABLE `AR` DISABLE KEYS */;
INSERT INTO `AR` VALUES (2,1,'Application Fee Received',0,1,0,6,46,'Application fee taken, no assessment made','1900-01-01 00:00:00','9999-12-31 00:00:00',5,0.0000,'2017-12-05 17:24:15',0,'2017-11-10 23:24:23',0),(4,1,'Bad Debt Write-Off',0,2,0,71,9,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(5,1,'Bank Service Fee (Security Deposit Account)',0,2,0,72,4,'','1900-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-28 18:54:17',0,'2017-11-10 23:24:23',0),(6,1,'Bank Service Fee (Operating Account)',0,2,0,72,3,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(8,1,'Damage Fee',0,0,0,9,59,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(9,1,'Deposit to Security Deposit Account (FRB96953)',0,1,0,4,6,'','1900-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-28 18:55:11',0,'2017-11-10 23:24:23',0),(10,1,'Deposit to Operating Account (FRB54320)',0,1,0,3,6,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(11,1,'Electric Base Fee',0,0,0,9,36,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(12,1,'Electric Overage',0,0,0,9,37,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(13,1,'Eviction Fee Reimbursement',0,0,0,9,56,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(14,1,'Auto-Generated Floating Deposit Assessment',0,3,0,9,12,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(16,1,'Gas Base Fee',0,0,0,9,40,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(17,1,'Gas Base Overage',0,0,0,9,41,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(18,1,'Insufficient Funds Fee',0,0,0,9,48,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(19,1,'Late Fee',0,0,0,9,47,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(20,1,'Month to Month Fee',0,0,0,9,49,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(21,1,'No Show / Termination Fee',0,0,0,9,51,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(22,1,'Other Special Tenant Charges',0,0,0,9,61,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(23,1,'Pet Fee',0,0,0,9,52,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(24,1,'Pet Rent',0,0,0,9,53,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(25,1,'Receive a Payment',0,1,0,6,10,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(27,1,'Gross Scheduled Rent',0,0,0,9,17,'','1900-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-28 18:40:18',0,'2017-11-10 23:24:23',0),(28,1,'Security Deposit Assessment',0,0,0,9,11,'normal deposit','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(29,1,'Security Deposit Forfeiture',0,0,0,11,58,'Forfeit','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(30,1,'Security Deposit Refund from Operating Account',0,0,0,11,3,'Refund','1900-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-28 19:23:08',0,'2017-11-10 23:24:23',0),(31,1,'Special Cleaning Fee',0,0,0,9,55,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(32,1,'Tenant Expense Chargeback',0,0,0,9,54,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(33,1,'Vending Income from Credit Card',0,1,0,7,65,'','1900-01-01 00:00:00','9999-12-31 00:00:00',5,0.0000,'2017-11-28 19:26:13',0,'2017-11-10 23:24:23',0),(34,1,'Water and Sewer Base Fee',0,0,0,9,38,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(35,1,'Water and Sewer Overage',0,0,0,9,39,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(36,1,'Auto-Generated Application Fee Assessment',0,3,0,9,46,'','1900-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-28 18:57:23',0,'2017-11-10 23:24:23',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=66 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Assessments`
--

LOCK TABLES `Assessments` WRITE;
/*!40000 ALTER TABLE `Assessments` DISABLE KEYS */;
INSERT INTO `Assessments` VALUES (1,0,0,0,1,1,0,2,7000.0000,'2014-03-01 00:00:00','2014-03-01 00:00:00',0,0,0,'',28,2,'','2017-11-30 19:46:56',0,'2017-11-30 18:39:27',0),(2,0,0,0,1,3,0,4,8300.0000,'2016-07-01 00:00:00','2016-07-01 00:00:00',0,0,0,'',28,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:41:00',0),(3,0,0,0,1,1,0,2,3750.0000,'2017-01-01 00:00:00','2018-03-01 00:00:00',6,4,0,'',27,0,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(4,3,0,0,1,1,0,2,3750.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:31',0,'2017-11-30 18:43:20',0),(5,3,0,0,1,1,0,2,3750.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:31',0,'2017-11-30 18:43:20',0),(6,3,0,0,1,1,0,2,3750.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:31',0,'2017-11-30 18:43:20',0),(7,3,0,0,1,1,0,2,3750.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:31',0,'2017-11-30 18:43:20',0),(8,3,0,0,1,1,0,2,3750.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:31',0,'2017-11-30 18:43:20',0),(9,3,0,0,1,1,0,2,3750.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:32',0,'2017-11-30 18:43:20',0),(10,3,0,0,1,1,0,2,3750.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:32',0,'2017-11-30 18:43:20',0),(11,3,0,0,1,1,0,2,3750.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:32',0,'2017-11-30 18:43:20',0),(12,3,0,0,1,1,0,2,3750.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:32',0,'2017-11-30 18:43:20',0),(13,3,0,0,1,1,0,2,3750.0000,'2017-10-01 00:00:00','2017-10-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:32',0,'2017-11-30 18:43:20',0),(14,3,0,0,1,1,0,2,3750.0000,'2017-11-01 00:00:00','2017-11-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:32',0,'2017-11-30 18:43:20',0),(15,0,0,0,1,2,0,3,4000.0000,'2017-01-01 00:00:00','2018-01-01 00:00:00',6,4,0,'',27,0,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(16,15,0,0,1,2,0,3,4000.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 20:50:21',0,'2017-11-30 18:45:17',0),(17,15,0,0,1,2,0,3,4000.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 20:50:21',0,'2017-11-30 18:45:17',0),(18,15,0,0,1,2,0,3,4000.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 20:50:21',0,'2017-11-30 18:45:17',0),(19,15,0,0,1,2,0,3,4000.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 20:50:21',0,'2017-11-30 18:45:17',0),(20,15,0,0,1,2,0,3,4000.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 20:50:21',0,'2017-11-30 18:45:17',0),(21,15,0,0,1,2,0,3,4000.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 20:50:21',0,'2017-11-30 18:45:17',0),(22,15,0,0,1,2,0,3,4000.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 20:50:21',0,'2017-11-30 18:45:17',0),(23,15,0,0,1,2,0,3,4000.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 20:50:21',0,'2017-11-30 18:45:17',0),(24,15,0,0,1,2,0,3,4000.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 20:50:22',0,'2017-11-30 18:45:17',0),(25,15,0,0,1,2,0,3,4000.0000,'2017-10-01 00:00:00','2017-10-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 20:50:22',0,'2017-11-30 18:45:17',0),(26,15,0,0,1,2,0,3,4000.0000,'2017-11-01 00:00:00','2017-11-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 20:50:22',0,'2017-11-30 18:45:17',0),(27,0,0,0,1,3,0,4,4150.0000,'2017-01-01 00:00:00','2018-07-01 00:00:00',6,4,0,'',27,0,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(28,27,0,0,1,3,0,4,4150.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(29,27,0,0,1,3,0,4,4150.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(30,27,0,0,1,3,0,4,4150.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(31,27,0,0,1,3,0,4,4150.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(32,27,0,0,1,3,0,4,4150.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(33,27,0,0,1,3,0,4,4150.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(34,27,0,0,1,3,0,4,4150.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(35,27,0,0,1,3,0,4,4150.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(36,27,0,0,1,3,0,4,4150.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(37,27,0,0,1,3,0,4,4150.0000,'2017-10-01 00:00:00','2017-10-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(38,27,0,0,1,3,0,4,4150.0000,'2017-11-01 00:00:00','2017-11-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(39,3,0,0,1,1,0,2,3750.0000,'2017-12-01 00:00:00','2017-12-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:32',0,'2017-12-01 00:00:04',0),(40,15,0,0,1,2,0,3,4000.0000,'2017-12-01 00:00:00','2017-12-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 20:50:22',0,'2017-12-01 00:00:04',0),(41,27,0,0,1,3,0,4,4150.0000,'2017-12-01 00:00:00','2017-12-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-12-01 00:00:04',0),(42,0,0,0,1,2,0,3,628.4500,'2017-01-31 00:00:00','2017-01-31 00:00:00',0,0,0,'',12,2,'utilities reimbursement','2017-12-05 20:50:21',0,'2017-12-05 16:01:46',0),(43,0,0,0,1,2,0,3,175.0000,'2017-02-28 00:00:00','2017-02-28 00:00:00',0,0,0,'',12,2,'','2017-12-05 20:50:21',0,'2017-12-05 16:02:25',0),(44,0,0,0,1,2,0,3,175.0000,'2017-03-31 00:00:00','2017-03-31 00:00:00',0,0,0,'',12,2,'','2017-12-05 20:50:21',0,'2017-12-05 16:03:13',0),(45,0,0,0,1,2,0,3,81.7900,'2017-04-15 00:00:00','2017-04-15 00:00:00',0,0,0,'',12,2,'','2017-12-05 20:50:21',0,'2017-12-05 16:03:41',0),(46,0,0,0,1,2,0,3,409.2800,'2017-10-31 00:00:00','2017-10-31 00:00:00',0,0,0,'',12,2,'','2017-12-05 20:50:22',0,'2017-12-05 16:07:34',0),(47,0,0,0,1,2,0,3,4000.0000,'2016-11-01 00:00:00','2016-11-01 00:00:00',0,0,0,'',27,2,'','2017-12-05 20:50:21',0,'2017-12-05 16:15:10',0),(48,0,0,0,1,2,0,3,4000.0000,'2016-12-01 00:00:00','2016-12-01 00:00:00',0,0,0,'',27,2,'','2017-12-05 20:50:21',0,'2017-12-05 16:15:45',0),(49,0,0,0,1,2,0,3,350.0000,'2017-04-01 00:00:00','2018-12-31 00:00:00',6,4,0,'',11,0,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(50,49,0,0,1,2,0,3,350.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',11,2,'','2017-12-05 20:50:21',0,'2017-12-05 17:49:46',0),(51,49,0,0,1,2,0,3,350.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',11,2,'','2017-12-05 20:50:21',0,'2017-12-05 17:49:46',0),(52,49,0,0,1,2,0,3,350.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',11,2,'','2017-12-05 20:50:21',0,'2017-12-05 17:49:46',0),(53,49,0,0,1,2,0,3,350.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',11,2,'','2017-12-05 20:50:21',0,'2017-12-05 17:49:46',0),(54,49,0,0,1,2,0,3,350.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',11,2,'','2017-12-05 20:50:21',0,'2017-12-05 17:49:46',0),(55,49,0,0,1,2,0,3,350.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',11,2,'','2017-12-05 20:50:22',0,'2017-12-05 17:49:46',0),(56,49,0,0,1,2,0,3,350.0000,'2017-10-01 00:00:00','2017-10-02 00:00:00',6,4,0,'',11,2,'','2017-12-05 20:50:22',0,'2017-12-05 17:49:46',0),(57,49,0,0,1,2,0,3,350.0000,'2017-11-01 00:00:00','2017-11-02 00:00:00',6,4,0,'',11,2,'','2017-12-05 20:50:22',0,'2017-12-05 17:49:46',0),(58,49,0,0,1,2,0,3,350.0000,'2017-12-01 00:00:00','2017-12-02 00:00:00',6,4,0,'',11,2,'','2017-12-05 20:50:22',0,'2017-12-05 17:49:46',0),(59,0,0,0,1,2,0,3,4000.0000,'2016-10-01 00:00:00','2016-10-01 00:00:00',0,0,0,'',27,2,'','2017-12-05 20:50:21',0,'2017-12-05 18:03:15',0),(60,0,0,0,1,2,0,3,628.4500,'2017-12-05 00:00:00','2017-12-05 00:00:00',0,0,0,'',12,4,'Reversed by ASM00000061','2017-12-05 19:41:01',0,'2017-12-05 18:23:25',0),(61,0,60,0,1,2,0,3,-628.4500,'2017-12-05 00:00:00','2017-12-05 00:00:00',0,0,0,'',12,4,'Reversal of ASM00000060','2017-12-05 19:41:01',0,'2017-12-05 19:41:01',0),(62,3,0,0,1,1,0,2,3750.0000,'2018-01-01 00:00:00','2018-01-02 00:00:00',6,4,0,'',27,0,'','2018-01-01 00:00:00',0,'2018-01-01 00:00:00',0),(63,27,0,0,1,3,0,4,4150.0000,'2018-01-01 00:00:00','2018-01-02 00:00:00',6,4,0,'',27,0,'','2018-01-01 00:00:01',0,'2018-01-01 00:00:01',0),(64,49,0,0,1,2,0,3,0.0000,'2018-01-01 00:00:00','2018-01-01 00:00:00',6,4,0,'',11,0,'Prorated: 0 days out of 31','2018-01-01 00:00:01',0,'2018-01-01 00:00:01',0),(65,27,0,0,1,3,0,4,4150.0000,'2018-05-01 00:00:00','2018-05-02 00:00:00',6,4,0,'',27,0,'','2018-05-05 19:55:50',-1,'2018-05-05 19:55:50',-1);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
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
INSERT INTO `Business` VALUES (1,'REX','JGM First, LLC',6,4,4,0,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0,0);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `CARID` bigint(20) NOT NULL AUTO_INCREMENT,
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Deposit`
--

LOCK TABLES `Deposit` WRITE;
/*!40000 ALTER TABLE `Deposit` DISABLE KEYS */;
INSERT INTO `Deposit` VALUES (3,1,1,2,'2016-07-01',8300.0000,0.0000,0,'2017-11-30 19:17:32',0,'2017-11-30 19:17:32',0),(4,1,1,2,'2014-03-01',7000.0000,0.0000,0,'2017-11-30 19:25:23',0,'2017-11-30 19:25:23',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DepositMethod`
--

LOCK TABLES `DepositMethod` WRITE;
/*!40000 ALTER TABLE `DepositMethod` DISABLE KEYS */;
INSERT INTO `DepositMethod` VALUES (2,1,'Scanned/Electronic Batch','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(3,1,'ACH','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(4,1,'Wire','2017-11-30 18:54:20',0,'2017-11-10 23:24:23',0),(5,1,'EFT','2017-11-30 18:55:11',0,'2017-11-30 18:55:11',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DepositPart`
--

LOCK TABLES `DepositPart` WRITE;
/*!40000 ALTER TABLE `DepositPart` DISABLE KEYS */;
INSERT INTO `DepositPart` VALUES (5,3,1,3,'2017-11-30 19:17:32',0,'2017-11-30 19:17:32',0),(7,4,1,6,'2017-11-30 19:25:23',0,'2017-11-30 19:25:23',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Depository`
--

LOCK TABLES `Depository` WRITE;
/*!40000 ALTER TABLE `Depository` DISABLE KEYS */;
INSERT INTO `Depository` VALUES (1,1,3,'FRB 54320 Operating Account','80001054320','2017-12-05 17:07:03',0,'2017-11-10 23:24:23',0),(3,1,4,'FRB 96953 Security Deposits','80003196953','2017-12-05 17:07:33',0,'2017-12-04 20:06:16',0);
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Expense`
--

LOCK TABLES `Expense` WRITE;
/*!40000 ALTER TABLE `Expense` DISABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=75 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `GLAccount`
--

LOCK TABLES `GLAccount` WRITE;
/*!40000 ALTER TABLE `GLAccount` DISABLE KEYS */;
INSERT INTO `GLAccount` VALUES (1,0,1,0,0,'10000',2,'Cash','Cash',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(2,1,1,0,0,'10100',2,'Petty Cash','Cash',1,0,'','2017-11-28 18:32:14',0,'2017-11-10 23:24:22',0),(3,1,1,0,0,'10104',2,'FRB 54320 (operating account)','Bank Account',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(4,1,1,0,0,'10105',2,'FRB 96953 (security deposits)','Bank Account',1,0,'','2017-11-27 21:42:09',0,'2017-11-10 23:24:22',0),(6,1,1,0,0,'10999',2,'Undeposited Funds','Cash',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(7,0,1,0,0,'11000',2,'Credit Cards Funds in Transit','Cash',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(8,0,1,0,0,'12000',2,'Accounts Receivable','Accounts Receivable',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(9,8,1,0,0,'12001',2,'Rent Roll Receivables','Accounts Receivable',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(10,0,1,0,0,'12999',2,'Unapplied Funds','Asset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(11,0,1,0,0,'30000',2,'Security Deposit Liability','Liability Security Deposit',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(12,0,1,0,0,'30001',2,'Floating Security Deposits','Liability Security Deposit',1,0,'Sec Dep posted before rentable identified','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(13,0,1,0,0,'30100',2,'Collected Taxes','Liabilities',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(14,13,1,0,0,'30101',2,'Sales Taxes Collected','Liabilities',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(15,13,1,0,0,'30102',2,'Transient Occupancy Taxes Collected','Liabilities',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(16,13,1,0,0,'30199',2,'Other Collected Taxes','Liabilities',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(17,0,1,0,0,'41000',2,'Gross Scheduled Rent','Income',1,0,'','2017-11-28 18:38:33',0,'2017-11-10 23:24:22',0),(19,0,1,0,0,'41100',2,'Unit Income Offsets','Income Offset',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(20,19,1,0,0,'41101',2,'Vacancy','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(21,19,1,0,0,'41102',2,'Loss (Gain) to Lease','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(22,19,1,0,0,'41103',2,'Employee Concessions','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(23,19,1,0,0,'41104',2,'Resident Concessions','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(24,19,1,0,0,'41105',2,'Owner Concession','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(25,19,1,0,0,'41106',2,'Administrative Concession','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(26,19,1,0,0,'41107',2,'Off Line Renovations','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(27,19,1,0,0,'41108',2,'Off Line Maintenance','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(28,19,1,0,0,'41199',2,'Othe Income Offsets','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(29,0,1,0,0,'41200',2,'Service Fees','Income',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(30,29,1,0,0,'41201',2,'Broadcast and IT Services','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(31,29,1,0,0,'41202',2,'Food Services','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(32,29,1,0,0,'41203',2,'Linen Services','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(33,29,1,0,0,'41204',2,'Wash N Fold Services','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(34,29,1,0,0,'41299',2,'Other Service Fees','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(35,0,1,0,0,'41300',2,'Utility Fees','Income',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(36,35,1,0,0,'41301',2,'Electric Base Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(37,35,1,0,0,'41302',2,'Electric Overage','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(38,35,1,0,0,'41303',2,'Water and Sewer Base Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(39,35,1,0,0,'41304',2,'Water and Sewer Overage','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(40,35,1,0,0,'41305',2,'Gas Base Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(41,35,1,0,0,'41306',2,'Gas Overage','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(42,35,1,0,0,'41307',2,'Trash Collection Base Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(43,35,1,0,0,'41308',2,'Trash Collection Overage','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(44,35,1,0,0,'41399',2,'Other Utility Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(45,0,1,0,0,'41400',2,'Special Tenant Charges','Income',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(46,45,1,0,0,'41401',2,'Application Fees','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(47,45,1,0,0,'41402',2,'Late Fees','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(48,45,1,0,0,'41403',2,'Insufficient Funds Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(49,45,1,0,0,'41404',2,'Month to Month Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(50,45,1,0,0,'41405',2,'Rentable Specialties','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(51,45,1,0,0,'41406',2,'No Show or Termination Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(52,45,1,0,0,'41407',2,'Pet Fees','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(53,45,1,0,0,'41408',2,'Pet Rent','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(54,45,1,0,0,'41409',2,'Tenant Expense Chargeback','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(55,45,1,0,0,'41410',2,'Special Cleaning Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(56,45,1,0,0,'41411',2,'Eviction Fee Reimbursement','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(57,45,1,0,0,'41412',2,'Extra Person Charge','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(58,45,1,0,0,'41413',2,'Security Deposit Forfeiture','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(59,45,1,0,0,'41414',2,'Damage Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(60,45,1,0,0,'41415',2,'CAM Fees','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(61,45,1,0,0,'41499',2,'Other Special Tenant Charges','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(62,0,1,0,0,'42000',2,'Business Income','Income',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(63,62,1,0,0,'42100',2,'Convenience Store','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(64,62,1,0,0,'42200',2,'Fitness Center Revenue','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(65,62,1,0,0,'42300',2,'Vending Income','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(66,62,1,0,0,'42400',2,'Restaurant Sales','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(67,62,1,0,0,'42500',2,'Bar Sales','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(68,62,1,0,0,'42600',2,'Spa Sales','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(69,0,1,0,0,'50000',2,'Expenses','Expenses',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(70,69,1,0,0,'50001',2,'Cash Over/Short','Expenses',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(71,69,1,0,0,'50002',2,'Bad Debt','Expenses',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(72,69,1,0,0,'50003',2,'Bank Service Fee','Expenses',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(73,69,1,0,0,'50999',2,'Other Expenses','Expenses',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `InvoiceASMID` bigint(20) NOT NULL AUTO_INCREMENT,
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  `InvoicePayorID` bigint(20) NOT NULL AUTO_INCREMENT,
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
) ENGINE=InnoDB AUTO_INCREMENT=160 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Journal`
--

LOCK TABLES `Journal` WRITE;
/*!40000 ALTER TABLE `Journal` DISABLE KEYS */;
INSERT INTO `Journal` VALUES (1,1,'2014-03-01 00:00:00',7000.0000,1,1,'','2017-11-30 18:39:27',0,'2017-11-30 18:39:27',0),(2,1,'2016-07-01 00:00:00',8300.0000,1,2,'','2017-11-30 18:41:00',0,'2017-11-30 18:41:00',0),(3,1,'2017-01-01 00:00:00',3750.0000,1,4,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(4,1,'2017-02-01 00:00:00',3750.0000,1,5,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(5,1,'2017-03-01 00:00:00',3750.0000,1,6,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(6,1,'2017-04-01 00:00:00',3750.0000,1,7,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(7,1,'2017-05-01 00:00:00',3750.0000,1,8,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(8,1,'2017-06-01 00:00:00',3750.0000,1,9,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(9,1,'2017-07-01 00:00:00',3750.0000,1,10,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(10,1,'2017-08-01 00:00:00',3750.0000,1,11,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(11,1,'2017-09-01 00:00:00',3750.0000,1,12,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(12,1,'2017-10-01 00:00:00',3750.0000,1,13,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(13,1,'2017-11-01 00:00:00',3750.0000,1,14,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(14,1,'2017-01-01 00:00:00',4000.0000,1,16,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(15,1,'2017-02-01 00:00:00',4000.0000,1,17,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(16,1,'2017-03-01 00:00:00',4000.0000,1,18,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(17,1,'2017-04-01 00:00:00',4000.0000,1,19,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(18,1,'2017-05-01 00:00:00',4000.0000,1,20,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(19,1,'2017-06-01 00:00:00',4000.0000,1,21,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(20,1,'2017-07-01 00:00:00',4000.0000,1,22,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(21,1,'2017-08-01 00:00:00',4000.0000,1,23,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(22,1,'2017-09-01 00:00:00',4000.0000,1,24,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(23,1,'2017-10-01 00:00:00',4000.0000,1,25,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(24,1,'2017-11-01 00:00:00',4000.0000,1,26,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(25,1,'2017-01-01 00:00:00',4150.0000,1,28,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(26,1,'2017-02-01 00:00:00',4150.0000,1,29,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(27,1,'2017-03-01 00:00:00',4150.0000,1,30,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(28,1,'2017-04-01 00:00:00',4150.0000,1,31,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(29,1,'2017-05-01 00:00:00',4150.0000,1,32,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(30,1,'2017-06-01 00:00:00',4150.0000,1,33,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(31,1,'2017-07-01 00:00:00',4150.0000,1,34,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(32,1,'2017-08-01 00:00:00',4150.0000,1,35,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(33,1,'2017-09-01 00:00:00',4150.0000,1,36,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(34,1,'2017-10-01 00:00:00',4150.0000,1,37,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(35,1,'2017-11-01 00:00:00',4150.0000,1,38,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(36,1,'2014-03-01 00:00:00',7000.0000,2,1,'','2017-11-30 18:48:02',0,'2017-11-30 18:48:02',0),(37,1,'2016-07-01 00:00:00',8300.0000,2,2,'','2017-11-30 18:49:17',0,'2017-11-30 18:49:17',0),(38,1,'2016-07-01 00:00:00',8300.0000,2,3,'','2017-11-30 19:13:23',0,'2017-11-30 19:13:23',0),(39,1,'2016-07-01 00:00:00',-8300.0000,2,4,'','2017-11-30 19:13:59',0,'2017-11-30 19:13:59',0),(40,1,'2016-07-01 00:00:00',8300.0000,4,3,'auto-transfer for deposit DEP-1','2017-11-30 19:17:32',0,'2017-11-30 19:17:32',0),(41,1,'2014-03-01 00:00:00',-7000.0000,2,5,'','2017-11-30 19:23:47',0,'2017-11-30 19:23:47',0),(42,1,'2014-03-01 00:00:00',7000.0000,2,6,'','2017-11-30 19:24:32',0,'2017-11-30 19:24:32',0),(43,1,'2014-03-01 00:00:00',7000.0000,4,6,'auto-transfer for deposit DEP-1','2017-11-30 19:25:23',0,'2017-11-30 19:25:23',0),(44,1,'2017-01-01 00:00:00',3750.0000,2,7,'','2017-11-30 19:44:52',0,'2017-11-30 19:44:52',0),(45,1,'2014-03-01 00:00:00',7000.0000,2,6,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(46,1,'2017-01-01 00:00:00',3750.0000,2,7,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(47,1,'2017-12-01 00:00:00',3750.0000,1,39,'','2017-12-01 00:00:04',0,'2017-12-01 00:00:04',0),(48,1,'2017-12-01 00:00:00',4000.0000,1,40,'','2017-12-01 00:00:04',0,'2017-12-01 00:00:04',0),(49,1,'2017-12-01 00:00:00',4150.0000,1,41,'','2017-12-01 00:00:04',0,'2017-12-01 00:00:04',0),(50,1,'2017-01-31 00:00:00',628.4500,1,42,'','2017-12-05 16:01:46',0,'2017-12-05 16:01:46',0),(51,1,'2017-02-28 00:00:00',175.0000,1,43,'','2017-12-05 16:02:25',0,'2017-12-05 16:02:25',0),(52,1,'2017-03-31 00:00:00',175.0000,1,44,'','2017-12-05 16:03:13',0,'2017-12-05 16:03:13',0),(53,1,'2017-04-15 00:00:00',81.7900,1,45,'','2017-12-05 16:03:41',0,'2017-12-05 16:03:41',0),(54,1,'2017-10-31 00:00:00',409.2800,1,46,'','2017-12-05 16:07:34',0,'2017-12-05 16:07:34',0),(55,1,'2017-01-01 00:00:00',3750.0000,2,8,'','2017-12-05 16:09:37',0,'2017-12-05 16:09:37',0),(56,1,'2017-01-01 00:00:00',4150.0000,2,9,'','2017-12-05 16:10:06',0,'2017-12-05 16:10:06',0),(57,1,'2017-02-01 00:00:00',8350.0000,2,10,'','2017-12-05 16:12:02',0,'2017-12-05 16:12:02',0),(58,1,'2017-02-01 00:00:00',3750.0000,2,11,'','2017-12-05 16:12:32',0,'2017-12-05 16:12:32',0),(59,1,'2017-02-01 00:00:00',4150.0000,2,12,'','2017-12-05 16:13:44',0,'2017-12-05 16:13:44',0),(60,1,'2016-11-01 00:00:00',4000.0000,1,47,'','2017-12-05 16:15:10',0,'2017-12-05 16:15:10',0),(61,1,'2016-12-01 00:00:00',4000.0000,1,48,'','2017-12-05 16:15:45',0,'2017-12-05 16:15:45',0),(62,1,'2016-11-15 00:00:00',12000.0000,2,13,'','2017-12-05 16:16:37',0,'2017-12-05 16:16:37',0),(63,1,'2017-01-01 00:00:00',-3750.0000,2,14,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(64,1,'2017-12-05 16:19:04',-3750.0000,2,14,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(65,1,'2017-03-01 00:00:00',3750.0000,2,15,'','2017-12-05 16:20:28',0,'2017-12-05 16:20:28',0),(66,1,'2017-03-01 00:00:00',4150.0000,2,16,'','2017-12-05 16:22:08',0,'2017-12-05 16:22:08',0),(67,1,'2017-04-01 00:00:00',3750.0000,2,17,'','2017-12-05 16:22:50',0,'2017-12-05 16:22:50',0),(68,1,'2017-04-01 00:00:00',4150.0000,2,18,'','2017-12-05 16:23:12',0,'2017-12-05 16:23:12',0),(69,1,'2017-05-01 00:00:00',3750.0000,2,19,'','2017-12-05 16:24:00',0,'2017-12-05 16:24:00',0),(70,1,'2017-05-01 00:00:00',4150.0000,2,20,'','2017-12-05 16:24:18',0,'2017-12-05 16:24:18',0),(71,1,'2017-05-15 00:00:00',13131.7900,2,21,'','2017-12-05 16:26:21',0,'2017-12-05 16:26:21',0),(72,1,'2017-06-01 00:00:00',3750.0000,2,22,'','2017-12-05 16:27:03',0,'2017-12-05 16:27:03',0),(73,1,'2017-06-01 00:00:00',4150.0000,2,23,'','2017-12-05 16:27:16',0,'2017-12-05 16:27:16',0),(74,1,'2017-07-01 00:00:00',3750.0000,2,24,'','2017-12-05 16:27:58',0,'2017-12-05 16:27:58',0),(75,1,'2017-07-01 00:00:00',4150.0000,2,25,'','2017-12-05 16:28:12',0,'2017-12-05 16:28:12',0),(76,1,'2017-08-01 00:00:00',3750.0000,2,26,'','2017-12-05 16:29:16',0,'2017-12-05 16:29:16',0),(77,1,'2017-08-01 00:00:00',4150.0000,2,27,'','2017-12-05 16:29:33',0,'2017-12-05 16:29:33',0),(78,1,'2017-08-15 00:00:00',13050.0000,2,28,'','2017-12-05 16:29:59',0,'2017-12-05 16:29:59',0),(79,1,'2017-09-01 00:00:00',3750.0000,2,29,'','2017-12-05 16:30:33',0,'2017-12-05 16:30:33',0),(80,1,'2017-09-01 00:00:00',4150.0000,2,30,'','2017-12-05 16:30:51',0,'2017-12-05 16:30:51',0),(81,1,'2017-10-01 00:00:00',3750.0000,2,31,'','2017-12-05 16:31:42',0,'2017-12-05 16:31:42',0),(82,1,'2017-10-01 00:00:00',4150.0000,2,32,'','2017-12-05 16:31:56',0,'2017-12-05 16:31:56',0),(83,1,'2017-11-01 00:00:00',3750.0000,2,33,'','2017-12-05 16:32:48',0,'2017-12-05 16:32:48',0),(84,1,'2017-11-01 00:00:00',4150.0000,2,34,'','2017-12-05 16:33:11',0,'2017-12-05 16:33:11',0),(85,1,'2017-11-15 00:00:00',13459.2800,2,35,'','2017-12-05 16:40:59',0,'2017-12-05 16:40:59',0),(86,1,'2017-12-01 00:00:00',3750.0000,2,36,'','2017-12-05 16:42:24',0,'2017-12-05 16:42:24',0),(87,1,'2017-12-01 00:00:00',4150.0000,2,37,'','2017-12-05 16:42:35',0,'2017-12-05 16:42:35',0),(88,1,'2017-01-01 00:00:00',3750.0000,2,8,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(89,1,'2017-02-01 00:00:00',3750.0000,2,11,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(90,1,'2017-03-01 00:00:00',3750.0000,2,15,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(91,1,'2017-04-01 00:00:00',3750.0000,2,17,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(92,1,'2017-05-01 00:00:00',3750.0000,2,19,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(93,1,'2017-06-01 00:00:00',3750.0000,2,22,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(94,1,'2017-07-01 00:00:00',3750.0000,2,24,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(95,1,'2017-08-01 00:00:00',3750.0000,2,26,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(96,1,'2017-09-01 00:00:00',3750.0000,2,29,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(97,1,'2017-10-01 00:00:00',3750.0000,2,31,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(98,1,'2017-11-01 00:00:00',3750.0000,2,33,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(99,1,'2017-12-01 00:00:00',3750.0000,2,36,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(100,1,'2016-07-01 00:00:00',8300.0000,2,3,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(101,1,'2017-01-01 00:00:00',4150.0000,2,9,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(102,1,'2017-02-01 00:00:00',4150.0000,2,12,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(103,1,'2017-03-01 00:00:00',4150.0000,2,16,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(104,1,'2017-04-01 00:00:00',4150.0000,2,18,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(105,1,'2017-05-01 00:00:00',4150.0000,2,20,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(106,1,'2017-06-01 00:00:00',4150.0000,2,23,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(107,1,'2017-07-01 00:00:00',4150.0000,2,25,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(108,1,'2017-08-01 00:00:00',4150.0000,2,27,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(109,1,'2017-09-01 00:00:00',4150.0000,2,30,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(110,1,'2017-10-01 00:00:00',4150.0000,2,32,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(111,1,'2017-11-01 00:00:00',4150.0000,2,34,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(112,1,'2017-12-01 00:00:00',4150.0000,2,37,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(113,1,'2017-04-01 00:00:00',350.0000,1,50,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(114,1,'2017-05-01 00:00:00',350.0000,1,51,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(115,1,'2017-06-01 00:00:00',350.0000,1,52,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(116,1,'2017-07-01 00:00:00',350.0000,1,53,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(117,1,'2017-08-01 00:00:00',350.0000,1,54,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(118,1,'2017-09-01 00:00:00',350.0000,1,55,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(119,1,'2017-10-01 00:00:00',350.0000,1,56,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(120,1,'2017-11-01 00:00:00',350.0000,1,57,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(121,1,'2017-12-01 00:00:00',350.0000,1,58,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(122,1,'2016-10-01 00:00:00',4000.0000,1,59,'','2017-12-05 18:03:15',0,'2017-12-05 18:03:15',0),(123,1,'2017-12-05 00:00:00',628.4500,1,60,'','2017-12-05 18:23:25',0,'2017-12-05 18:23:25',0),(124,1,'2017-12-05 00:00:00',-628.4500,1,61,'','2017-12-05 19:41:01',0,'2017-12-05 19:41:01',0),(125,1,'2017-02-03 00:00:00',628.4500,2,38,'','2017-12-05 19:44:51',0,'2017-12-05 19:44:51',0),(126,1,'2016-10-03 00:00:00',4000.0000,2,39,'','2017-12-05 20:44:04',0,'2017-12-05 20:44:04',0),(127,1,'2016-10-03 00:00:00',4000.0000,2,39,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(128,1,'2016-11-11 00:00:00',4000.0000,2,13,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(129,1,'2016-11-11 00:00:00',4000.0000,2,13,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(130,1,'2016-11-11 00:00:00',4000.0000,2,13,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(131,1,'2017-02-03 00:00:00',628.4500,2,10,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(132,1,'2017-02-13 00:00:00',4000.0000,2,10,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(133,1,'2017-02-13 00:00:00',175.0000,2,10,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(134,1,'2017-02-13 00:00:00',3546.5500,2,10,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(135,1,'2017-02-13 00:00:00',453.4500,2,38,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(136,1,'2017-02-13 00:00:00',175.0000,2,38,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(137,1,'2017-05-12 00:00:00',4000.0000,2,21,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(138,1,'2017-05-12 00:00:00',350.0000,2,21,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(139,1,'2017-05-12 00:00:00',81.7900,2,21,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(140,1,'2017-05-12 00:00:00',4000.0000,2,21,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(141,1,'2017-05-12 00:00:00',350.0000,2,21,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(142,1,'2017-05-12 00:00:00',4000.0000,2,21,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(143,1,'2017-05-12 00:00:00',350.0000,2,21,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(144,1,'2017-08-15 00:00:00',4000.0000,2,28,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(145,1,'2017-08-15 00:00:00',350.0000,2,28,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(146,1,'2017-08-15 00:00:00',4000.0000,2,28,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(147,1,'2017-08-15 00:00:00',350.0000,2,28,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(148,1,'2017-08-15 00:00:00',4000.0000,2,28,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(149,1,'2017-08-15 00:00:00',350.0000,2,28,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(150,1,'2017-11-14 00:00:00',4000.0000,2,35,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(151,1,'2017-11-14 00:00:00',350.0000,2,35,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(152,1,'2017-11-14 00:00:00',409.2800,2,35,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(153,1,'2017-11-14 00:00:00',4000.0000,2,35,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(154,1,'2017-11-14 00:00:00',350.0000,2,35,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(155,1,'2017-11-14 00:00:00',4000.0000,2,35,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(156,1,'2017-11-14 00:00:00',350.0000,2,35,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(157,1,'2018-01-01 00:00:00',3750.0000,1,62,'','2018-01-01 00:00:01',0,'2018-01-01 00:00:01',0),(158,1,'2018-01-01 00:00:00',4150.0000,1,63,'','2018-01-01 00:00:01',0,'2018-01-01 00:00:01',0),(159,1,'2018-01-01 00:00:00',0.0000,1,64,'','2018-01-01 00:00:01',0,'2018-01-01 00:00:01',0);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`JAID`)
) ENGINE=InnoDB AUTO_INCREMENT=160 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `JournalAllocation`
--

LOCK TABLES `JournalAllocation` WRITE;
/*!40000 ALTER TABLE `JournalAllocation` DISABLE KEYS */;
INSERT INTO `JournalAllocation` VALUES (1,1,1,1,2,0,0,7000.0000,1,0,'d 12001 7000.00, c 30000 7000.00','2018-01-10 18:34:01',0,'2017-11-30 18:39:27',0),(2,1,2,3,4,0,0,8300.0000,2,0,'d 12001 8300.00, c 30000 8300.00','2018-01-10 18:34:01',0,'2017-11-30 18:41:00',0),(3,1,3,1,2,0,0,3750.0000,4,0,'d 12001 3750.00, c 41000 3750.00','2018-01-10 18:34:01',0,'2017-11-30 18:43:20',0),(4,1,4,1,2,0,0,3750.0000,5,0,'d 12001 3750.00, c 41000 3750.00','2018-01-10 18:34:01',0,'2017-11-30 18:43:20',0),(5,1,5,1,2,0,0,3750.0000,6,0,'d 12001 3750.00, c 41000 3750.00','2018-01-10 18:34:01',0,'2017-11-30 18:43:20',0),(6,1,6,1,2,0,0,3750.0000,7,0,'d 12001 3750.00, c 41000 3750.00','2018-01-10 18:34:01',0,'2017-11-30 18:43:20',0),(7,1,7,1,2,0,0,3750.0000,8,0,'d 12001 3750.00, c 41000 3750.00','2018-01-10 18:34:01',0,'2017-11-30 18:43:20',0),(8,1,8,1,2,0,0,3750.0000,9,0,'d 12001 3750.00, c 41000 3750.00','2018-01-10 18:34:01',0,'2017-11-30 18:43:20',0),(9,1,9,1,2,0,0,3750.0000,10,0,'d 12001 3750.00, c 41000 3750.00','2018-01-10 18:34:01',0,'2017-11-30 18:43:20',0),(10,1,10,1,2,0,0,3750.0000,11,0,'d 12001 3750.00, c 41000 3750.00','2018-01-10 18:34:01',0,'2017-11-30 18:43:20',0),(11,1,11,1,2,0,0,3750.0000,12,0,'d 12001 3750.00, c 41000 3750.00','2018-01-10 18:34:01',0,'2017-11-30 18:43:20',0),(12,1,12,1,2,0,0,3750.0000,13,0,'d 12001 3750.00, c 41000 3750.00','2018-01-10 18:34:01',0,'2017-11-30 18:43:20',0),(13,1,13,1,2,0,0,3750.0000,14,0,'d 12001 3750.00, c 41000 3750.00','2018-01-10 18:34:01',0,'2017-11-30 18:43:20',0),(14,1,14,2,3,0,0,4000.0000,16,0,'d 12001 4000.00, c 41000 4000.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:17',0),(15,1,15,2,3,0,0,4000.0000,17,0,'d 12001 4000.00, c 41000 4000.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:17',0),(16,1,16,2,3,0,0,4000.0000,18,0,'d 12001 4000.00, c 41000 4000.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:17',0),(17,1,17,2,3,0,0,4000.0000,19,0,'d 12001 4000.00, c 41000 4000.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:17',0),(18,1,18,2,3,0,0,4000.0000,20,0,'d 12001 4000.00, c 41000 4000.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:17',0),(19,1,19,2,3,0,0,4000.0000,21,0,'d 12001 4000.00, c 41000 4000.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:17',0),(20,1,20,2,3,0,0,4000.0000,22,0,'d 12001 4000.00, c 41000 4000.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:17',0),(21,1,21,2,3,0,0,4000.0000,23,0,'d 12001 4000.00, c 41000 4000.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:17',0),(22,1,22,2,3,0,0,4000.0000,24,0,'d 12001 4000.00, c 41000 4000.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:17',0),(23,1,23,2,3,0,0,4000.0000,25,0,'d 12001 4000.00, c 41000 4000.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:17',0),(24,1,24,2,3,0,0,4000.0000,26,0,'d 12001 4000.00, c 41000 4000.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:17',0),(25,1,25,3,4,0,0,4150.0000,28,0,'d 12001 4150.00, c 41000 4150.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:55',0),(26,1,26,3,4,0,0,4150.0000,29,0,'d 12001 4150.00, c 41000 4150.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:55',0),(27,1,27,3,4,0,0,4150.0000,30,0,'d 12001 4150.00, c 41000 4150.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:55',0),(28,1,28,3,4,0,0,4150.0000,31,0,'d 12001 4150.00, c 41000 4150.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:55',0),(29,1,29,3,4,0,0,4150.0000,32,0,'d 12001 4150.00, c 41000 4150.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:55',0),(30,1,30,3,4,0,0,4150.0000,33,0,'d 12001 4150.00, c 41000 4150.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:55',0),(31,1,31,3,4,0,0,4150.0000,34,0,'d 12001 4150.00, c 41000 4150.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:55',0),(32,1,32,3,4,0,0,4150.0000,35,0,'d 12001 4150.00, c 41000 4150.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:55',0),(33,1,33,3,4,0,0,4150.0000,36,0,'d 12001 4150.00, c 41000 4150.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:55',0),(34,1,34,3,4,0,0,4150.0000,37,0,'d 12001 4150.00, c 41000 4150.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:55',0),(35,1,35,3,4,0,0,4150.0000,38,0,'d 12001 4150.00, c 41000 4150.00','2018-01-10 18:34:01',0,'2017-11-30 18:45:55',0),(36,1,36,0,0,1,0,7000.0000,0,0,'d 10104 _, c 10999 _','2018-01-10 18:34:01',0,'2017-11-30 18:48:02',0),(37,1,37,0,0,4,0,8300.0000,0,0,'d 10104 _, c 10999 _','2018-01-10 18:34:01',0,'2017-11-30 18:49:17',0),(38,1,38,0,0,4,0,8300.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-11-30 19:13:23',0),(39,1,39,0,0,4,0,-8300.0000,0,0,'d 10104 _, c 10999 _','2018-01-10 18:34:01',0,'2017-11-30 19:13:59',0),(40,1,40,0,0,4,3,8300.0000,0,0,'d 10104 8300.0000, c 10999 8300.0000','2018-01-10 18:34:01',0,'2017-11-30 19:17:32',0),(41,1,41,0,0,1,0,-7000.0000,0,0,'d 10104 _, c 10999 _','2018-01-10 18:34:01',0,'2017-11-30 19:23:47',0),(42,1,42,0,0,1,0,7000.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-11-30 19:24:32',0),(43,1,43,0,0,1,6,7000.0000,0,0,'d 10104 7000.0000, c 10999 7000.0000','2018-01-10 18:34:01',0,'2017-11-30 19:25:23',0),(44,1,44,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-11-30 19:44:52',0),(45,1,45,1,2,1,6,7000.0000,1,0,'ASM(1) d 12999 7000.00,c 12001 7000.00','2018-01-10 18:34:01',0,'2017-11-30 19:46:56',0),(46,1,46,1,2,1,7,3750.0000,4,0,'ASM(4) d 12999 3750.00,c 12001 3750.00','2018-01-10 18:34:01',0,'2017-11-30 19:46:56',0),(47,1,47,1,2,0,0,3750.0000,39,0,'d 12001 3750.00, c 41000 3750.00','2018-01-10 18:34:01',0,'2017-12-01 00:00:04',0),(48,1,48,2,3,0,0,4000.0000,40,0,'d 12001 4000.00, c 41000 4000.00','2018-01-10 18:34:01',0,'2017-12-01 00:00:04',0),(49,1,49,3,4,0,0,4150.0000,41,0,'d 12001 4150.00, c 41000 4150.00','2018-01-10 18:34:01',0,'2017-12-01 00:00:04',0),(50,1,50,2,3,0,0,628.4500,42,0,'d 12001 628.45, c 41302 628.45','2018-01-10 18:34:01',0,'2017-12-05 16:01:46',0),(51,1,51,2,3,0,0,175.0000,43,0,'d 12001 175.00, c 41302 175.00','2018-01-10 18:34:01',0,'2017-12-05 16:02:25',0),(52,1,52,2,3,0,0,175.0000,44,0,'d 12001 175.00, c 41302 175.00','2018-01-10 18:34:01',0,'2017-12-05 16:03:13',0),(53,1,53,2,3,0,0,81.7900,45,0,'d 12001 81.79, c 41302 81.79','2018-01-10 18:34:01',0,'2017-12-05 16:03:41',0),(54,1,54,2,3,0,0,409.2800,46,0,'d 12001 409.28, c 41302 409.28','2018-01-10 18:34:01',0,'2017-12-05 16:07:34',0),(55,1,55,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:09:37',0),(56,1,56,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:10:06',0),(57,1,57,0,0,3,0,8350.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:12:02',0),(58,1,58,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:12:32',0),(59,1,59,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:13:44',0),(60,1,60,2,3,0,0,4000.0000,47,0,'d 12001 4000.00, c 41000 4000.00','2018-01-10 18:34:01',0,'2017-12-05 16:15:10',0),(61,1,61,2,3,0,0,4000.0000,48,0,'d 12001 4000.00, c 41000 4000.00','2018-01-10 18:34:01',0,'2017-12-05 16:15:45',0),(62,1,62,0,0,3,0,12000.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:16:37',0),(63,1,63,0,0,1,0,-3750.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:19:04',0),(64,1,64,1,2,1,14,-3750.0000,4,0,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2018-01-10 18:34:01',0,'2017-12-05 16:19:04',0),(65,1,65,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:20:28',0),(66,1,66,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:22:08',0),(67,1,67,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:22:50',0),(68,1,68,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:23:12',0),(69,1,69,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:24:00',0),(70,1,70,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:24:18',0),(71,1,71,0,0,3,0,13131.7900,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:26:21',0),(72,1,72,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:27:03',0),(73,1,73,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:27:16',0),(74,1,74,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:27:58',0),(75,1,75,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:28:12',0),(76,1,76,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:29:16',0),(77,1,77,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:29:33',0),(78,1,78,0,0,3,0,13050.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:29:59',0),(79,1,79,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:30:33',0),(80,1,80,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:30:51',0),(81,1,81,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:31:42',0),(82,1,82,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:31:56',0),(83,1,83,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:32:49',0),(84,1,84,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:33:11',0),(85,1,85,0,0,3,0,13459.2800,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:40:59',0),(86,1,86,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:42:24',0),(87,1,87,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 16:42:35',0),(88,1,88,1,2,1,8,3750.0000,4,0,'ASM(4) d 12999 3750.00,c 12001 3750.00','2018-01-10 18:34:01',0,'2017-12-05 16:59:31',0),(89,1,89,1,2,1,11,3750.0000,5,0,'ASM(5) d 12999 3750.00,c 12001 3750.00','2018-01-10 18:34:01',0,'2017-12-05 16:59:31',0),(90,1,90,1,2,1,15,3750.0000,6,0,'ASM(6) d 12999 3750.00,c 12001 3750.00','2018-01-10 18:34:01',0,'2017-12-05 16:59:31',0),(91,1,91,1,2,1,17,3750.0000,7,0,'ASM(7) d 12999 3750.00,c 12001 3750.00','2018-01-10 18:34:01',0,'2017-12-05 16:59:31',0),(92,1,92,1,2,1,19,3750.0000,8,0,'ASM(8) d 12999 3750.00,c 12001 3750.00','2018-01-10 18:34:01',0,'2017-12-05 16:59:32',0),(93,1,93,1,2,1,22,3750.0000,9,0,'ASM(9) d 12999 3750.00,c 12001 3750.00','2018-01-10 18:34:01',0,'2017-12-05 16:59:32',0),(94,1,94,1,2,1,24,3750.0000,10,0,'ASM(10) d 12999 3750.00,c 12001 3750.00','2018-01-10 18:34:01',0,'2017-12-05 16:59:32',0),(95,1,95,1,2,1,26,3750.0000,11,0,'ASM(11) d 12999 3750.00,c 12001 3750.00','2018-01-10 18:34:01',0,'2017-12-05 16:59:32',0),(96,1,96,1,2,1,29,3750.0000,12,0,'ASM(12) d 12999 3750.00,c 12001 3750.00','2018-01-10 18:34:01',0,'2017-12-05 16:59:32',0),(97,1,97,1,2,1,31,3750.0000,13,0,'ASM(13) d 12999 3750.00,c 12001 3750.00','2018-01-10 18:34:01',0,'2017-12-05 16:59:32',0),(98,1,98,1,2,1,33,3750.0000,14,0,'ASM(14) d 12999 3750.00,c 12001 3750.00','2018-01-10 18:34:01',0,'2017-12-05 16:59:32',0),(99,1,99,1,2,1,36,3750.0000,39,0,'ASM(39) d 12999 3750.00,c 12001 3750.00','2018-01-10 18:34:01',0,'2017-12-05 16:59:32',0),(100,1,100,3,4,4,3,8300.0000,2,0,'ASM(2) d 12999 8300.00,c 12001 8300.00','2018-01-10 18:34:01',0,'2017-12-05 17:06:25',0),(101,1,101,3,4,4,9,4150.0000,28,0,'ASM(28) d 12999 4150.00,c 12001 4150.00','2018-01-10 18:34:01',0,'2017-12-05 17:06:25',0),(102,1,102,3,4,4,12,4150.0000,29,0,'ASM(29) d 12999 4150.00,c 12001 4150.00','2018-01-10 18:34:01',0,'2017-12-05 17:06:25',0),(103,1,103,3,4,4,16,4150.0000,30,0,'ASM(30) d 12999 4150.00,c 12001 4150.00','2018-01-10 18:34:01',0,'2017-12-05 17:06:25',0),(104,1,104,3,4,4,18,4150.0000,31,0,'ASM(31) d 12999 4150.00,c 12001 4150.00','2018-01-10 18:34:01',0,'2017-12-05 17:06:25',0),(105,1,105,3,4,4,20,4150.0000,32,0,'ASM(32) d 12999 4150.00,c 12001 4150.00','2018-01-10 18:34:01',0,'2017-12-05 17:06:25',0),(106,1,106,3,4,4,23,4150.0000,33,0,'ASM(33) d 12999 4150.00,c 12001 4150.00','2018-01-10 18:34:01',0,'2017-12-05 17:06:25',0),(107,1,107,3,4,4,25,4150.0000,34,0,'ASM(34) d 12999 4150.00,c 12001 4150.00','2018-01-10 18:34:01',0,'2017-12-05 17:06:25',0),(108,1,108,3,4,4,27,4150.0000,35,0,'ASM(35) d 12999 4150.00,c 12001 4150.00','2018-01-10 18:34:01',0,'2017-12-05 17:06:25',0),(109,1,109,3,4,4,30,4150.0000,36,0,'ASM(36) d 12999 4150.00,c 12001 4150.00','2018-01-10 18:34:01',0,'2017-12-05 17:06:25',0),(110,1,110,3,4,4,32,4150.0000,37,0,'ASM(37) d 12999 4150.00,c 12001 4150.00','2018-01-10 18:34:01',0,'2017-12-05 17:06:25',0),(111,1,111,3,4,4,34,4150.0000,38,0,'ASM(38) d 12999 4150.00,c 12001 4150.00','2018-01-10 18:34:01',0,'2017-12-05 17:06:25',0),(112,1,112,3,4,4,37,4150.0000,41,0,'ASM(41) d 12999 4150.00,c 12001 4150.00','2018-01-10 18:34:01',0,'2017-12-05 17:06:25',0),(113,1,113,2,3,0,0,350.0000,50,0,'d 12001 350.00, c 41301 350.00','2018-01-10 18:34:01',0,'2017-12-05 17:49:46',0),(114,1,114,2,3,0,0,350.0000,51,0,'d 12001 350.00, c 41301 350.00','2018-01-10 18:34:01',0,'2017-12-05 17:49:46',0),(115,1,115,2,3,0,0,350.0000,52,0,'d 12001 350.00, c 41301 350.00','2018-01-10 18:34:01',0,'2017-12-05 17:49:46',0),(116,1,116,2,3,0,0,350.0000,53,0,'d 12001 350.00, c 41301 350.00','2018-01-10 18:34:01',0,'2017-12-05 17:49:46',0),(117,1,117,2,3,0,0,350.0000,54,0,'d 12001 350.00, c 41301 350.00','2018-01-10 18:34:01',0,'2017-12-05 17:49:46',0),(118,1,118,2,3,0,0,350.0000,55,0,'d 12001 350.00, c 41301 350.00','2018-01-10 18:34:01',0,'2017-12-05 17:49:46',0),(119,1,119,2,3,0,0,350.0000,56,0,'d 12001 350.00, c 41301 350.00','2018-01-10 18:34:01',0,'2017-12-05 17:49:46',0),(120,1,120,2,3,0,0,350.0000,57,0,'d 12001 350.00, c 41301 350.00','2018-01-10 18:34:01',0,'2017-12-05 17:49:46',0),(121,1,121,2,3,0,0,350.0000,58,0,'d 12001 350.00, c 41301 350.00','2018-01-10 18:34:01',0,'2017-12-05 17:49:46',0),(122,1,122,2,3,0,0,4000.0000,59,0,'d 12001 4000.00, c 41000 4000.00','2018-01-10 18:34:01',0,'2017-12-05 18:03:15',0),(123,1,123,2,3,0,0,628.4500,60,0,'d 12001 628.45, c 41302 628.45','2018-01-10 18:34:01',0,'2017-12-05 18:23:25',0),(124,1,124,2,3,0,0,-628.4500,61,0,'d 12001 -628.45, c 41302 -628.45','2018-01-10 18:34:01',0,'2017-12-05 19:41:01',0),(125,1,125,0,0,3,0,628.4500,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 19:44:51',0),(126,1,126,0,0,3,0,4000.0000,0,0,'d 10999 _, c 12999 _','2018-01-10 18:34:01',0,'2017-12-05 20:44:04',0),(127,1,127,2,3,3,39,4000.0000,59,0,'ASM(59) d 12999 4000.00,c 12001 4000.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(128,1,128,2,3,3,13,4000.0000,47,0,'ASM(47) d 12999 4000.00,c 12001 4000.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(129,1,129,2,3,3,13,4000.0000,48,0,'ASM(48) d 12999 4000.00,c 12001 4000.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(130,1,130,2,3,3,13,4000.0000,16,0,'ASM(16) d 12999 4000.00,c 12001 4000.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(131,1,131,2,3,3,10,628.4500,42,0,'ASM(42) d 12999 628.45,c 12001 628.45','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(132,1,132,2,3,3,10,4000.0000,17,0,'ASM(17) d 12999 4000.00,c 12001 4000.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(133,1,133,2,3,3,10,175.0000,43,0,'ASM(43) d 12999 175.00,c 12001 175.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(134,1,134,2,3,3,10,3546.5500,18,0,'ASM(18) d 12999 3546.55,c 12001 3546.55','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(135,1,135,2,3,3,38,453.4500,18,0,'ASM(18) d 12999 453.45,c 12001 453.45','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(136,1,136,2,3,3,38,175.0000,44,0,'ASM(44) d 12999 175.00,c 12001 175.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(137,1,137,2,3,3,21,4000.0000,19,0,'ASM(19) d 12999 4000.00,c 12001 4000.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(138,1,138,2,3,3,21,350.0000,50,0,'ASM(50) d 12999 350.00,c 12001 350.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(139,1,139,2,3,3,21,81.7900,45,0,'ASM(45) d 12999 81.79,c 12001 81.79','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(140,1,140,2,3,3,21,4000.0000,20,0,'ASM(20) d 12999 4000.00,c 12001 4000.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(141,1,141,2,3,3,21,350.0000,51,0,'ASM(51) d 12999 350.00,c 12001 350.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(142,1,142,2,3,3,21,4000.0000,21,0,'ASM(21) d 12999 4000.00,c 12001 4000.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(143,1,143,2,3,3,21,350.0000,52,0,'ASM(52) d 12999 350.00,c 12001 350.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(144,1,144,2,3,3,28,4000.0000,22,0,'ASM(22) d 12999 4000.00,c 12001 4000.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(145,1,145,2,3,3,28,350.0000,53,0,'ASM(53) d 12999 350.00,c 12001 350.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(146,1,146,2,3,3,28,4000.0000,23,0,'ASM(23) d 12999 4000.00,c 12001 4000.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:21',0),(147,1,147,2,3,3,28,350.0000,54,0,'ASM(54) d 12999 350.00,c 12001 350.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:22',0),(148,1,148,2,3,3,28,4000.0000,24,0,'ASM(24) d 12999 4000.00,c 12001 4000.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:22',0),(149,1,149,2,3,3,28,350.0000,55,0,'ASM(55) d 12999 350.00,c 12001 350.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:22',0),(150,1,150,2,3,3,35,4000.0000,25,0,'ASM(25) d 12999 4000.00,c 12001 4000.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:22',0),(151,1,151,2,3,3,35,350.0000,56,0,'ASM(56) d 12999 350.00,c 12001 350.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:22',0),(152,1,152,2,3,3,35,409.2800,46,0,'ASM(46) d 12999 409.28,c 12001 409.28','2018-01-10 18:34:01',0,'2017-12-05 20:50:22',0),(153,1,153,2,3,3,35,4000.0000,26,0,'ASM(26) d 12999 4000.00,c 12001 4000.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:22',0),(154,1,154,2,3,3,35,350.0000,57,0,'ASM(57) d 12999 350.00,c 12001 350.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:22',0),(155,1,155,2,3,3,35,4000.0000,40,0,'ASM(40) d 12999 4000.00,c 12001 4000.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:22',0),(156,1,156,2,3,3,35,350.0000,58,0,'ASM(58) d 12999 350.00,c 12001 350.00','2018-01-10 18:34:01',0,'2017-12-05 20:50:22',0),(157,1,157,1,2,0,0,3750.0000,62,0,'d 12001 3750.00, c 41000 3750.00','2018-01-10 18:34:01',0,'2018-01-01 00:00:01',0),(158,1,158,3,4,0,0,4150.0000,63,0,'d 12001 4150.00, c 41000 4150.00','2018-01-10 18:34:01',0,'2018-01-01 00:00:01',0),(159,1,159,2,3,0,0,0.0000,64,0,'d 12001 0.00, c 41301 0.00','2018-01-10 18:34:01',0,'2018-01-01 00:00:01',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=307 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerEntry`
--

LOCK TABLES `LedgerEntry` WRITE;
/*!40000 ALTER TABLE `LedgerEntry` DISABLE KEYS */;
INSERT INTO `LedgerEntry` VALUES (1,1,1,1,9,2,1,0,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 18:39:27',0,'2017-11-30 18:39:27',0),(2,1,1,1,11,2,1,0,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 18:39:27',0,'2017-11-30 18:39:27',0),(3,1,2,2,9,4,3,0,'2016-07-01 00:00:00',8300.0000,'','2017-11-30 18:41:00',0,'2017-11-30 18:41:00',0),(4,1,2,2,11,4,3,0,'2016-07-01 00:00:00',-8300.0000,'','2017-11-30 18:41:00',0,'2017-11-30 18:41:00',0),(5,1,3,3,9,2,1,0,'2017-01-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(6,1,3,3,17,2,1,0,'2017-01-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(7,1,4,4,9,2,1,0,'2017-02-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(8,1,4,4,17,2,1,0,'2017-02-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(9,1,5,5,9,2,1,0,'2017-03-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(10,1,5,5,17,2,1,0,'2017-03-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(11,1,6,6,9,2,1,0,'2017-04-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(12,1,6,6,17,2,1,0,'2017-04-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(13,1,7,7,9,2,1,0,'2017-05-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(14,1,7,7,17,2,1,0,'2017-05-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(15,1,8,8,9,2,1,0,'2017-06-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(16,1,8,8,17,2,1,0,'2017-06-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(17,1,9,9,9,2,1,0,'2017-07-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(18,1,9,9,17,2,1,0,'2017-07-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(19,1,10,10,9,2,1,0,'2017-08-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(20,1,10,10,17,2,1,0,'2017-08-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(21,1,11,11,9,2,1,0,'2017-09-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(22,1,11,11,17,2,1,0,'2017-09-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(23,1,12,12,9,2,1,0,'2017-10-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(24,1,12,12,17,2,1,0,'2017-10-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(25,1,13,13,9,2,1,0,'2017-11-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(26,1,13,13,17,2,1,0,'2017-11-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(27,1,14,14,9,3,2,0,'2017-01-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(28,1,14,14,17,3,2,0,'2017-01-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(29,1,15,15,9,3,2,0,'2017-02-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(30,1,15,15,17,3,2,0,'2017-02-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(31,1,16,16,9,3,2,0,'2017-03-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(32,1,16,16,17,3,2,0,'2017-03-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(33,1,17,17,9,3,2,0,'2017-04-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(34,1,17,17,17,3,2,0,'2017-04-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(35,1,18,18,9,3,2,0,'2017-05-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(36,1,18,18,17,3,2,0,'2017-05-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(37,1,19,19,9,3,2,0,'2017-06-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(38,1,19,19,17,3,2,0,'2017-06-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(39,1,20,20,9,3,2,0,'2017-07-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(40,1,20,20,17,3,2,0,'2017-07-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(41,1,21,21,9,3,2,0,'2017-08-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(42,1,21,21,17,3,2,0,'2017-08-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(43,1,22,22,9,3,2,0,'2017-09-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(44,1,22,22,17,3,2,0,'2017-09-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(45,1,23,23,9,3,2,0,'2017-10-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(46,1,23,23,17,3,2,0,'2017-10-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(47,1,24,24,9,3,2,0,'2017-11-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(48,1,24,24,17,3,2,0,'2017-11-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(49,1,25,25,9,4,3,0,'2017-01-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(50,1,25,25,17,4,3,0,'2017-01-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(51,1,26,26,9,4,3,0,'2017-02-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(52,1,26,26,17,4,3,0,'2017-02-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(53,1,27,27,9,4,3,0,'2017-03-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(54,1,27,27,17,4,3,0,'2017-03-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(55,1,28,28,9,4,3,0,'2017-04-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(56,1,28,28,17,4,3,0,'2017-04-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(57,1,29,29,9,4,3,0,'2017-05-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(58,1,29,29,17,4,3,0,'2017-05-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(59,1,30,30,9,4,3,0,'2017-06-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(60,1,30,30,17,4,3,0,'2017-06-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(61,1,31,31,9,4,3,0,'2017-07-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(62,1,31,31,17,4,3,0,'2017-07-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(63,1,32,32,9,4,3,0,'2017-08-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(64,1,32,32,17,4,3,0,'2017-08-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(65,1,33,33,9,4,3,0,'2017-09-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(66,1,33,33,17,4,3,0,'2017-09-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(67,1,34,34,9,4,3,0,'2017-10-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(68,1,34,34,17,4,3,0,'2017-10-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(69,1,35,35,9,4,3,0,'2017-11-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(70,1,35,35,17,4,3,0,'2017-11-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(71,1,36,36,3,0,0,1,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 18:48:02',0,'2017-11-30 18:48:02',0),(72,1,36,36,6,0,0,1,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 18:48:02',0,'2017-11-30 18:48:02',0),(73,1,37,37,3,0,0,4,'2016-07-01 00:00:00',8300.0000,'','2017-11-30 18:49:17',0,'2017-11-30 18:49:17',0),(74,1,37,37,6,0,0,4,'2016-07-01 00:00:00',-8300.0000,'','2017-11-30 18:49:17',0,'2017-11-30 18:49:17',0),(75,1,38,38,6,0,0,4,'2016-07-01 00:00:00',8300.0000,'','2017-11-30 19:13:23',0,'2017-11-30 19:13:23',0),(76,1,38,38,10,0,0,4,'2016-07-01 00:00:00',-8300.0000,'','2017-11-30 19:13:23',0,'2017-11-30 19:13:23',0),(77,1,39,39,3,0,0,4,'2016-07-01 00:00:00',-8300.0000,'','2017-11-30 19:13:59',0,'2017-11-30 19:13:59',0),(78,1,39,39,6,0,0,4,'2016-07-01 00:00:00',8300.0000,'','2017-11-30 19:13:59',0,'2017-11-30 19:13:59',0),(79,1,40,40,3,0,0,4,'2016-07-01 00:00:00',8300.0000,'','2017-11-30 19:17:32',0,'2017-11-30 19:17:32',0),(80,1,40,40,6,0,0,4,'2016-07-01 00:00:00',-8300.0000,'','2017-11-30 19:17:32',0,'2017-11-30 19:17:32',0),(81,1,41,41,3,0,0,1,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 19:23:47',0,'2017-11-30 19:23:47',0),(82,1,41,41,6,0,0,1,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 19:23:47',0,'2017-11-30 19:23:47',0),(83,1,42,42,6,0,0,1,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 19:24:32',0,'2017-11-30 19:24:32',0),(84,1,42,42,10,0,0,1,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 19:24:32',0,'2017-11-30 19:24:32',0),(85,1,43,43,3,0,0,1,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 19:25:23',0,'2017-11-30 19:25:23',0),(86,1,43,43,6,0,0,1,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 19:25:23',0,'2017-11-30 19:25:23',0),(87,1,44,44,6,0,0,1,'2017-01-01 00:00:00',3750.0000,'','2017-11-30 19:44:52',0,'2017-11-30 19:44:52',0),(88,1,44,44,10,0,0,1,'2017-01-01 00:00:00',-3750.0000,'','2017-11-30 19:44:52',0,'2017-11-30 19:44:52',0),(89,1,45,45,10,2,1,1,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(90,1,45,45,9,2,1,1,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(91,1,46,46,10,2,1,1,'2017-01-01 00:00:00',3750.0000,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(92,1,46,46,9,2,1,1,'2017-01-01 00:00:00',-3750.0000,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(93,1,50,50,9,3,2,0,'2017-01-31 00:00:00',628.4500,'','2017-12-05 16:01:46',0,'2017-12-05 16:01:46',0),(94,1,50,50,37,3,2,0,'2017-01-31 00:00:00',-628.4500,'','2017-12-05 16:01:46',0,'2017-12-05 16:01:46',0),(95,1,51,51,9,3,2,0,'2017-02-28 00:00:00',175.0000,'','2017-12-05 16:02:25',0,'2017-12-05 16:02:25',0),(96,1,51,51,37,3,2,0,'2017-02-28 00:00:00',-175.0000,'','2017-12-05 16:02:25',0,'2017-12-05 16:02:25',0),(97,1,52,52,9,3,2,0,'2017-03-31 00:00:00',175.0000,'','2017-12-05 16:03:13',0,'2017-12-05 16:03:13',0),(98,1,52,52,37,3,2,0,'2017-03-31 00:00:00',-175.0000,'','2017-12-05 16:03:13',0,'2017-12-05 16:03:13',0),(99,1,53,53,9,3,2,0,'2017-04-15 00:00:00',81.7900,'','2017-12-05 16:03:41',0,'2017-12-05 16:03:41',0),(100,1,53,53,37,3,2,0,'2017-04-15 00:00:00',-81.7900,'','2017-12-05 16:03:41',0,'2017-12-05 16:03:41',0),(101,1,54,54,9,3,2,0,'2017-10-31 00:00:00',409.2800,'','2017-12-05 16:07:34',0,'2017-12-05 16:07:34',0),(102,1,54,54,37,3,2,0,'2017-10-31 00:00:00',-409.2800,'','2017-12-05 16:07:34',0,'2017-12-05 16:07:34',0),(103,1,55,55,6,0,0,1,'2017-01-01 00:00:00',3750.0000,'','2017-12-05 16:09:37',0,'2017-12-05 16:09:37',0),(104,1,55,55,10,0,0,1,'2017-01-01 00:00:00',-3750.0000,'','2017-12-05 16:09:37',0,'2017-12-05 16:09:37',0),(105,1,56,56,6,0,0,4,'2017-01-01 00:00:00',4150.0000,'','2017-12-05 16:10:06',0,'2017-12-05 16:10:06',0),(106,1,56,56,10,0,0,4,'2017-01-01 00:00:00',-4150.0000,'','2017-12-05 16:10:06',0,'2017-12-05 16:10:06',0),(107,1,57,57,6,0,0,3,'2017-02-01 00:00:00',8350.0000,'','2017-12-05 16:12:02',0,'2017-12-05 16:12:02',0),(108,1,57,57,10,0,0,3,'2017-02-01 00:00:00',-8350.0000,'','2017-12-05 16:12:02',0,'2017-12-05 16:12:02',0),(109,1,58,58,6,0,0,1,'2017-02-01 00:00:00',3750.0000,'','2017-12-05 16:12:32',0,'2017-12-05 16:12:32',0),(110,1,58,58,10,0,0,1,'2017-02-01 00:00:00',-3750.0000,'','2017-12-05 16:12:32',0,'2017-12-05 16:12:32',0),(111,1,59,59,6,0,0,4,'2017-02-01 00:00:00',4150.0000,'','2017-12-05 16:13:44',0,'2017-12-05 16:13:44',0),(112,1,59,59,10,0,0,4,'2017-02-01 00:00:00',-4150.0000,'','2017-12-05 16:13:44',0,'2017-12-05 16:13:44',0),(113,1,60,60,9,3,2,0,'2016-11-01 00:00:00',4000.0000,'','2017-12-05 16:15:10',0,'2017-12-05 16:15:10',0),(114,1,60,60,17,3,2,0,'2016-11-01 00:00:00',-4000.0000,'','2017-12-05 16:15:10',0,'2017-12-05 16:15:10',0),(115,1,61,61,9,3,2,0,'2016-12-01 00:00:00',4000.0000,'','2017-12-05 16:15:45',0,'2017-12-05 16:15:45',0),(116,1,61,61,17,3,2,0,'2016-12-01 00:00:00',-4000.0000,'','2017-12-05 16:15:45',0,'2017-12-05 16:15:45',0),(117,1,62,62,6,0,0,3,'2016-11-15 00:00:00',12000.0000,'','2017-12-05 16:16:37',0,'2017-12-05 16:16:37',0),(118,1,62,62,10,0,0,3,'2016-11-15 00:00:00',-12000.0000,'','2017-12-05 16:16:37',0,'2017-12-05 16:16:37',0),(119,1,63,63,6,0,0,1,'2017-01-01 00:00:00',-3750.0000,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(120,1,63,63,10,0,0,1,'2017-01-01 00:00:00',3750.0000,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(121,1,64,64,10,2,1,1,'2017-01-01 00:00:00',-3750.0000,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(122,1,64,64,9,2,1,1,'2017-01-01 00:00:00',3750.0000,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(123,1,65,65,6,0,0,1,'2017-03-01 00:00:00',3750.0000,'','2017-12-05 16:20:28',0,'2017-12-05 16:20:28',0),(124,1,65,65,10,0,0,1,'2017-03-01 00:00:00',-3750.0000,'','2017-12-05 16:20:28',0,'2017-12-05 16:20:28',0),(125,1,66,66,6,0,0,4,'2017-03-01 00:00:00',4150.0000,'','2017-12-05 16:22:08',0,'2017-12-05 16:22:08',0),(126,1,66,66,10,0,0,4,'2017-03-01 00:00:00',-4150.0000,'','2017-12-05 16:22:08',0,'2017-12-05 16:22:08',0),(127,1,67,67,6,0,0,1,'2017-04-01 00:00:00',3750.0000,'','2017-12-05 16:22:50',0,'2017-12-05 16:22:50',0),(128,1,67,67,10,0,0,1,'2017-04-01 00:00:00',-3750.0000,'','2017-12-05 16:22:50',0,'2017-12-05 16:22:50',0),(129,1,68,68,6,0,0,4,'2017-04-01 00:00:00',4150.0000,'','2017-12-05 16:23:12',0,'2017-12-05 16:23:12',0),(130,1,68,68,10,0,0,4,'2017-04-01 00:00:00',-4150.0000,'','2017-12-05 16:23:12',0,'2017-12-05 16:23:12',0),(131,1,69,69,6,0,0,1,'2017-05-01 00:00:00',3750.0000,'','2017-12-05 16:24:00',0,'2017-12-05 16:24:00',0),(132,1,69,69,10,0,0,1,'2017-05-01 00:00:00',-3750.0000,'','2017-12-05 16:24:01',0,'2017-12-05 16:24:01',0),(133,1,70,70,6,0,0,4,'2017-05-01 00:00:00',4150.0000,'','2017-12-05 16:24:18',0,'2017-12-05 16:24:18',0),(134,1,70,70,10,0,0,4,'2017-05-01 00:00:00',-4150.0000,'','2017-12-05 16:24:18',0,'2017-12-05 16:24:18',0),(135,1,71,71,6,0,0,3,'2017-05-15 00:00:00',13131.7900,'','2017-12-05 16:26:21',0,'2017-12-05 16:26:21',0),(136,1,71,71,10,0,0,3,'2017-05-15 00:00:00',-13131.7900,'','2017-12-05 16:26:21',0,'2017-12-05 16:26:21',0),(137,1,72,72,6,0,0,1,'2017-06-01 00:00:00',3750.0000,'','2017-12-05 16:27:03',0,'2017-12-05 16:27:03',0),(138,1,72,72,10,0,0,1,'2017-06-01 00:00:00',-3750.0000,'','2017-12-05 16:27:03',0,'2017-12-05 16:27:03',0),(139,1,73,73,6,0,0,4,'2017-06-01 00:00:00',4150.0000,'','2017-12-05 16:27:16',0,'2017-12-05 16:27:16',0),(140,1,73,73,10,0,0,4,'2017-06-01 00:00:00',-4150.0000,'','2017-12-05 16:27:16',0,'2017-12-05 16:27:16',0),(141,1,74,74,6,0,0,1,'2017-07-01 00:00:00',3750.0000,'','2017-12-05 16:27:58',0,'2017-12-05 16:27:58',0),(142,1,74,74,10,0,0,1,'2017-07-01 00:00:00',-3750.0000,'','2017-12-05 16:27:58',0,'2017-12-05 16:27:58',0),(143,1,75,75,6,0,0,4,'2017-07-01 00:00:00',4150.0000,'','2017-12-05 16:28:13',0,'2017-12-05 16:28:13',0),(144,1,75,75,10,0,0,4,'2017-07-01 00:00:00',-4150.0000,'','2017-12-05 16:28:13',0,'2017-12-05 16:28:13',0),(145,1,76,76,6,0,0,1,'2017-08-01 00:00:00',3750.0000,'','2017-12-05 16:29:16',0,'2017-12-05 16:29:16',0),(146,1,76,76,10,0,0,1,'2017-08-01 00:00:00',-3750.0000,'','2017-12-05 16:29:16',0,'2017-12-05 16:29:16',0),(147,1,77,77,6,0,0,4,'2017-08-01 00:00:00',4150.0000,'','2017-12-05 16:29:33',0,'2017-12-05 16:29:33',0),(148,1,77,77,10,0,0,4,'2017-08-01 00:00:00',-4150.0000,'','2017-12-05 16:29:33',0,'2017-12-05 16:29:33',0),(149,1,78,78,6,0,0,3,'2017-08-15 00:00:00',13050.0000,'','2017-12-05 16:29:59',0,'2017-12-05 16:29:59',0),(150,1,78,78,10,0,0,3,'2017-08-15 00:00:00',-13050.0000,'','2017-12-05 16:29:59',0,'2017-12-05 16:29:59',0),(151,1,79,79,6,0,0,1,'2017-09-01 00:00:00',3750.0000,'','2017-12-05 16:30:33',0,'2017-12-05 16:30:33',0),(152,1,79,79,10,0,0,1,'2017-09-01 00:00:00',-3750.0000,'','2017-12-05 16:30:33',0,'2017-12-05 16:30:33',0),(153,1,80,80,6,0,0,4,'2017-09-01 00:00:00',4150.0000,'','2017-12-05 16:30:51',0,'2017-12-05 16:30:51',0),(154,1,80,80,10,0,0,4,'2017-09-01 00:00:00',-4150.0000,'','2017-12-05 16:30:51',0,'2017-12-05 16:30:51',0),(155,1,81,81,6,0,0,1,'2017-10-01 00:00:00',3750.0000,'','2017-12-05 16:31:42',0,'2017-12-05 16:31:42',0),(156,1,81,81,10,0,0,1,'2017-10-01 00:00:00',-3750.0000,'','2017-12-05 16:31:42',0,'2017-12-05 16:31:42',0),(157,1,82,82,6,0,0,4,'2017-10-01 00:00:00',4150.0000,'','2017-12-05 16:31:56',0,'2017-12-05 16:31:56',0),(158,1,82,82,10,0,0,4,'2017-10-01 00:00:00',-4150.0000,'','2017-12-05 16:31:56',0,'2017-12-05 16:31:56',0),(159,1,83,83,6,0,0,1,'2017-11-01 00:00:00',3750.0000,'','2017-12-05 16:32:49',0,'2017-12-05 16:32:49',0),(160,1,83,83,10,0,0,1,'2017-11-01 00:00:00',-3750.0000,'','2017-12-05 16:32:49',0,'2017-12-05 16:32:49',0),(161,1,84,84,6,0,0,4,'2017-11-01 00:00:00',4150.0000,'','2017-12-05 16:33:11',0,'2017-12-05 16:33:11',0),(162,1,84,84,10,0,0,4,'2017-11-01 00:00:00',-4150.0000,'','2017-12-05 16:33:11',0,'2017-12-05 16:33:11',0),(163,1,85,85,6,0,0,3,'2017-11-15 00:00:00',13459.2800,'','2017-12-05 16:40:59',0,'2017-12-05 16:40:59',0),(164,1,85,85,10,0,0,3,'2017-11-15 00:00:00',-13459.2800,'','2017-12-05 16:40:59',0,'2017-12-05 16:40:59',0),(165,1,86,86,6,0,0,1,'2017-12-01 00:00:00',3750.0000,'','2017-12-05 16:42:24',0,'2017-12-05 16:42:24',0),(166,1,86,86,10,0,0,1,'2017-12-01 00:00:00',-3750.0000,'','2017-12-05 16:42:24',0,'2017-12-05 16:42:24',0),(167,1,87,87,6,0,0,4,'2017-12-01 00:00:00',4150.0000,'','2017-12-05 16:42:35',0,'2017-12-05 16:42:35',0),(168,1,87,87,10,0,0,4,'2017-12-01 00:00:00',-4150.0000,'','2017-12-05 16:42:35',0,'2017-12-05 16:42:35',0),(169,1,88,88,10,2,1,1,'2017-01-01 00:00:00',3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(170,1,88,88,9,2,1,1,'2017-01-01 00:00:00',-3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(171,1,89,89,10,2,1,1,'2017-02-01 00:00:00',3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(172,1,89,89,9,2,1,1,'2017-02-01 00:00:00',-3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(173,1,90,90,10,2,1,1,'2017-03-01 00:00:00',3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(174,1,90,90,9,2,1,1,'2017-03-01 00:00:00',-3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(175,1,91,91,10,2,1,1,'2017-04-01 00:00:00',3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(176,1,91,91,9,2,1,1,'2017-04-01 00:00:00',-3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(177,1,92,92,10,2,1,1,'2017-05-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(178,1,92,92,9,2,1,1,'2017-05-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(179,1,93,93,10,2,1,1,'2017-06-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(180,1,93,93,9,2,1,1,'2017-06-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(181,1,94,94,10,2,1,1,'2017-07-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(182,1,94,94,9,2,1,1,'2017-07-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(183,1,95,95,10,2,1,1,'2017-08-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(184,1,95,95,9,2,1,1,'2017-08-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(185,1,96,96,10,2,1,1,'2017-09-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(186,1,96,96,9,2,1,1,'2017-09-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(187,1,97,97,10,2,1,1,'2017-10-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(188,1,97,97,9,2,1,1,'2017-10-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(189,1,98,98,10,2,1,1,'2017-11-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(190,1,98,98,9,2,1,1,'2017-11-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(191,1,99,99,10,2,1,1,'2017-12-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(192,1,99,99,9,2,1,1,'2017-12-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(193,1,100,100,10,4,3,4,'2016-07-01 00:00:00',8300.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(194,1,100,100,9,4,3,4,'2016-07-01 00:00:00',-8300.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(195,1,101,101,10,4,3,4,'2017-01-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(196,1,101,101,9,4,3,4,'2017-01-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(197,1,102,102,10,4,3,4,'2017-02-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(198,1,102,102,9,4,3,4,'2017-02-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(199,1,103,103,10,4,3,4,'2017-03-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(200,1,103,103,9,4,3,4,'2017-03-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(201,1,104,104,10,4,3,4,'2017-04-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(202,1,104,104,9,4,3,4,'2017-04-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(203,1,105,105,10,4,3,4,'2017-05-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(204,1,105,105,9,4,3,4,'2017-05-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(205,1,106,106,10,4,3,4,'2017-06-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(206,1,106,106,9,4,3,4,'2017-06-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(207,1,107,107,10,4,3,4,'2017-07-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(208,1,107,107,9,4,3,4,'2017-07-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(209,1,108,108,10,4,3,4,'2017-08-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(210,1,108,108,9,4,3,4,'2017-08-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(211,1,109,109,10,4,3,4,'2017-09-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(212,1,109,109,9,4,3,4,'2017-09-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(213,1,110,110,10,4,3,4,'2017-10-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(214,1,110,110,9,4,3,4,'2017-10-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(215,1,111,111,10,4,3,4,'2017-11-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(216,1,111,111,9,4,3,4,'2017-11-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(217,1,112,112,10,4,3,4,'2017-12-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(218,1,112,112,9,4,3,4,'2017-12-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(219,1,113,113,9,3,2,0,'2017-04-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(220,1,113,113,36,3,2,0,'2017-04-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(221,1,114,114,9,3,2,0,'2017-05-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(222,1,114,114,36,3,2,0,'2017-05-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(223,1,115,115,9,3,2,0,'2017-06-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(224,1,115,115,36,3,2,0,'2017-06-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(225,1,116,116,9,3,2,0,'2017-07-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(226,1,116,116,36,3,2,0,'2017-07-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(227,1,117,117,9,3,2,0,'2017-08-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(228,1,117,117,36,3,2,0,'2017-08-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(229,1,118,118,9,3,2,0,'2017-09-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(230,1,118,118,36,3,2,0,'2017-09-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(231,1,119,119,9,3,2,0,'2017-10-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(232,1,119,119,36,3,2,0,'2017-10-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(233,1,120,120,9,3,2,0,'2017-11-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(234,1,120,120,36,3,2,0,'2017-11-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(235,1,121,121,9,3,2,0,'2017-12-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(236,1,121,121,36,3,2,0,'2017-12-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(237,1,122,122,9,3,2,0,'2016-10-01 00:00:00',4000.0000,'','2017-12-05 18:03:15',0,'2017-12-05 18:03:15',0),(238,1,122,122,17,3,2,0,'2016-10-01 00:00:00',-4000.0000,'','2017-12-05 18:03:15',0,'2017-12-05 18:03:15',0),(239,1,123,123,9,3,2,0,'2017-12-05 00:00:00',628.4500,'','2017-12-05 18:23:25',0,'2017-12-05 18:23:25',0),(240,1,123,123,37,3,2,0,'2017-12-05 00:00:00',-628.4500,'','2017-12-05 18:23:25',0,'2017-12-05 18:23:25',0),(241,1,124,124,9,3,2,0,'2017-12-05 00:00:00',-628.4500,'','2017-12-05 19:41:01',0,'2017-12-05 19:41:01',0),(242,1,124,124,37,3,2,0,'2017-12-05 00:00:00',628.4500,'','2017-12-05 19:41:01',0,'2017-12-05 19:41:01',0),(243,1,125,125,6,0,0,3,'2017-02-03 00:00:00',628.4500,'','2017-12-05 19:44:51',0,'2017-12-05 19:44:51',0),(244,1,125,125,10,0,0,3,'2017-02-03 00:00:00',-628.4500,'','2017-12-05 19:44:51',0,'2017-12-05 19:44:51',0),(245,1,126,126,6,0,0,3,'2016-10-03 00:00:00',4000.0000,'','2017-12-05 20:44:04',0,'2017-12-05 20:44:04',0),(246,1,126,126,10,0,0,3,'2016-10-03 00:00:00',-4000.0000,'','2017-12-05 20:44:04',0,'2017-12-05 20:44:04',0),(247,1,127,127,10,3,2,3,'2016-10-03 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(248,1,127,127,9,3,2,3,'2016-10-03 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(249,1,128,128,10,3,2,3,'2016-11-11 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(250,1,128,128,9,3,2,3,'2016-11-11 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(251,1,129,129,10,3,2,3,'2016-11-11 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(252,1,129,129,9,3,2,3,'2016-11-11 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(253,1,130,130,10,3,2,3,'2016-11-11 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(254,1,130,130,9,3,2,3,'2016-11-11 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(255,1,131,131,10,3,2,3,'2017-02-03 00:00:00',628.4500,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(256,1,131,131,9,3,2,3,'2017-02-03 00:00:00',-628.4500,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(257,1,132,132,10,3,2,3,'2017-02-13 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(258,1,132,132,9,3,2,3,'2017-02-13 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(259,1,133,133,10,3,2,3,'2017-02-13 00:00:00',175.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(260,1,133,133,9,3,2,3,'2017-02-13 00:00:00',-175.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(261,1,134,134,10,3,2,3,'2017-02-13 00:00:00',3546.5500,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(262,1,134,134,9,3,2,3,'2017-02-13 00:00:00',-3546.5500,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(263,1,135,135,10,3,2,3,'2017-02-13 00:00:00',453.4500,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(264,1,135,135,9,3,2,3,'2017-02-13 00:00:00',-453.4500,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(265,1,136,136,10,3,2,3,'2017-02-13 00:00:00',175.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(266,1,136,136,9,3,2,3,'2017-02-13 00:00:00',-175.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(267,1,137,137,10,3,2,3,'2017-05-12 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(268,1,137,137,9,3,2,3,'2017-05-12 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(269,1,138,138,10,3,2,3,'2017-05-12 00:00:00',350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(270,1,138,138,9,3,2,3,'2017-05-12 00:00:00',-350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(271,1,139,139,10,3,2,3,'2017-05-12 00:00:00',81.7900,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(272,1,139,139,9,3,2,3,'2017-05-12 00:00:00',-81.7900,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(273,1,140,140,10,3,2,3,'2017-05-12 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(274,1,140,140,9,3,2,3,'2017-05-12 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(275,1,141,141,10,3,2,3,'2017-05-12 00:00:00',350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(276,1,141,141,9,3,2,3,'2017-05-12 00:00:00',-350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(277,1,142,142,10,3,2,3,'2017-05-12 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(278,1,142,142,9,3,2,3,'2017-05-12 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(279,1,143,143,10,3,2,3,'2017-05-12 00:00:00',350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(280,1,143,143,9,3,2,3,'2017-05-12 00:00:00',-350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(281,1,144,144,10,3,2,3,'2017-08-15 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(282,1,144,144,9,3,2,3,'2017-08-15 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(283,1,145,145,10,3,2,3,'2017-08-15 00:00:00',350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(284,1,145,145,9,3,2,3,'2017-08-15 00:00:00',-350.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(285,1,146,146,10,3,2,3,'2017-08-15 00:00:00',4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(286,1,146,146,9,3,2,3,'2017-08-15 00:00:00',-4000.0000,'','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(287,1,147,147,10,3,2,3,'2017-08-15 00:00:00',350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(288,1,147,147,9,3,2,3,'2017-08-15 00:00:00',-350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(289,1,148,148,10,3,2,3,'2017-08-15 00:00:00',4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(290,1,148,148,9,3,2,3,'2017-08-15 00:00:00',-4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(291,1,149,149,10,3,2,3,'2017-08-15 00:00:00',350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(292,1,149,149,9,3,2,3,'2017-08-15 00:00:00',-350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(293,1,150,150,10,3,2,3,'2017-11-14 00:00:00',4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(294,1,150,150,9,3,2,3,'2017-11-14 00:00:00',-4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(295,1,151,151,10,3,2,3,'2017-11-14 00:00:00',350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(296,1,151,151,9,3,2,3,'2017-11-14 00:00:00',-350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(297,1,152,152,10,3,2,3,'2017-11-14 00:00:00',409.2800,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(298,1,152,152,9,3,2,3,'2017-11-14 00:00:00',-409.2800,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(299,1,153,153,10,3,2,3,'2017-11-14 00:00:00',4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(300,1,153,153,9,3,2,3,'2017-11-14 00:00:00',-4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(301,1,154,154,10,3,2,3,'2017-11-14 00:00:00',350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(302,1,154,154,9,3,2,3,'2017-11-14 00:00:00',-350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(303,1,155,155,10,3,2,3,'2017-11-14 00:00:00',4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(304,1,155,155,9,3,2,3,'2017-11-14 00:00:00',-4000.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(305,1,156,156,10,3,2,3,'2017-11-14 00:00:00',350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(306,1,156,156,9,3,2,3,'2017-11-14 00:00:00',-350.0000,'','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=91 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerMarker`
--

LOCK TABLES `LedgerMarker` WRITE;
/*!40000 ALTER TABLE `LedgerMarker` DISABLE KEYS */;
INSERT INTO `LedgerMarker` VALUES (1,1,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(2,2,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(3,3,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(4,4,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(6,6,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(7,7,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(8,8,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(9,9,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(10,10,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(11,11,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(12,12,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(13,13,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(14,14,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(15,15,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(16,16,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(17,17,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(19,19,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(20,20,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(21,21,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(22,22,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(23,23,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(24,24,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(25,25,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(26,26,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(27,27,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(28,28,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(29,29,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(30,30,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(31,31,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(32,32,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(33,33,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(34,34,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(35,35,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(36,36,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(37,37,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(38,38,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(39,39,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(40,40,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(41,41,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(42,42,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(43,43,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(44,44,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(45,45,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(46,46,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(47,47,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(48,48,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(49,49,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(50,50,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(51,51,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(52,52,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(53,53,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(54,54,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(55,55,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(56,56,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(57,57,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(58,58,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(59,59,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(60,60,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(61,61,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(62,62,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(63,63,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(64,64,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(65,65,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(66,66,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(67,67,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(68,68,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(69,69,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(70,70,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(71,71,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(72,72,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(73,73,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(75,0,0,1,0,0,'2017-11-28 00:00:00',0.0000,3,'2017-11-28 18:14:19',0,'2017-11-28 18:14:19',0),(76,0,1,0,0,1,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:15:20',0,'2017-11-30 18:15:20',0),(77,0,1,0,0,2,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:15:28',0,'2017-11-30 18:15:28',0),(78,0,1,0,0,3,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:16:10',0,'2017-11-30 18:16:10',0),(79,0,1,0,0,4,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:16:17',0,'2017-11-30 18:16:17',0),(80,0,1,0,0,5,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:16:28',0,'2017-11-30 18:16:28',0),(81,0,0,2,0,0,'2014-03-01 00:00:00',0.0000,3,'2017-11-30 18:29:02',0,'2017-11-30 18:17:55',0),(82,0,1,2,1,0,'2014-03-01 00:00:00',0.0000,3,'2017-11-30 18:20:15',0,'2017-11-30 18:20:15',0),(83,0,1,0,0,6,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:24:52',0,'2017-11-30 18:24:52',0),(84,0,0,3,0,0,'2016-10-01 00:00:00',0.0000,3,'2017-11-30 18:33:53',0,'2017-11-30 18:29:25',0),(85,0,1,3,2,0,'2016-10-01 00:00:00',0.0000,3,'2017-11-30 18:32:13',0,'2017-11-30 18:32:13',0),(86,0,0,4,0,0,'2016-07-01 00:00:00',0.0000,3,'2017-11-30 18:37:13',0,'2017-11-30 18:33:59',0),(87,0,1,4,3,0,'2016-07-01 00:00:00',0.0000,3,'2017-11-30 18:34:33',0,'2017-11-30 18:34:33',0),(88,0,0,5,0,0,'2017-11-30 00:00:00',0.0000,3,'2017-11-30 18:37:24',0,'2017-11-30 18:37:24',0),(89,0,0,6,0,0,'2018-01-03 00:00:00',0.0000,3,'2018-01-03 07:07:58',0,'2018-01-03 07:07:58',0),(90,0,0,7,0,0,'2018-01-03 00:00:00',0.0000,3,'2018-01-03 07:08:22',0,'2018-01-03 07:08:22',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PaymentType`
--

LOCK TABLES `PaymentType` WRITE;
/*!40000 ALTER TABLE `PaymentType` DISABLE KEYS */;
INSERT INTO `PaymentType` VALUES (1,1,'Cash','Cash','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(2,1,'Check','Personal check from payor','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(3,1,'VISA','Credit card charge','2017-11-10 23:24:23',0,'2017-12-03 08:28:05',0),(4,1,'AMEX','American Express credit card','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(5,1,'ACH','Bank transfer','2017-11-10 23:24:23',0,'2017-11-28 17:56:08',0),(6,1,'Wire','Wire transfer','2017-11-30 19:42:29',0,'2017-11-30 19:42:29',0);
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
  PRIMARY KEY (`TCID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Payor`
--

LOCK TABLES `Payor` WRITE;
/*!40000 ALTER TABLE `Payor` DISABLE KEYS */;
INSERT INTO `Payor` VALUES (1,1,'',0.0000,0,1,0,'','',0.0000,'2017-11-30 18:15:20',0,'2017-11-30 18:15:20',0),(2,1,'',0.0000,0,1,0,'','',0.0000,'2017-11-30 18:15:28',0,'2017-11-30 18:15:28',0),(3,1,'',0.0000,0,1,0,'','',0.0000,'2017-11-30 18:16:10',0,'2017-11-30 18:16:10',0),(4,1,'',0.0000,0,1,0,'','',0.0000,'2017-11-30 18:16:17',0,'2017-11-30 18:16:17',0),(5,1,'',0.0000,0,1,0,'','',0.0000,'2017-11-30 18:16:28',0,'2017-11-30 18:16:28',0),(6,1,'',0.0000,0,1,0,'','',0.0000,'2017-11-30 18:24:52',0,'2017-11-30 18:24:52',0);
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
  PRIMARY KEY (`TCID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Prospect`
--

LOCK TABLES `Prospect` WRITE;
/*!40000 ALTER TABLE `Prospect` DISABLE KEYS */;
INSERT INTO `Prospect` VALUES (1,1,'','','','','','','','1900-01-01',0,0,'','','',0,0,'','1900-01-01',0,0,'','','',0,'','','','',0,'','','2017-11-30 18:15:20',0,'2017-11-30 18:15:20',0),(2,1,'','','','','','','','1900-01-01',0,0,'','','',0,0,'','1900-01-01',0,0,'','','',0,'','','','',0,'','','2017-11-30 18:15:28',0,'2017-11-30 18:15:28',0),(3,1,'','','','','','','','1900-01-01',0,0,'','','',0,0,'','1900-01-01',0,0,'','','',0,'','','','',0,'','','2017-11-30 18:16:10',0,'2017-11-30 18:16:10',0),(4,1,'','','','','','','','1900-01-01',0,0,'','','',0,0,'','1900-01-01',0,0,'','','',0,'','','','',0,'','','2017-11-30 18:16:17',0,'2017-11-30 18:16:17',0),(5,1,'','','','','','','','1900-01-01',0,0,'','','',0,0,'','1900-01-01',0,0,'','','',0,'','','','',0,'','','2017-11-30 18:16:28',0,'2017-11-30 18:16:28',0),(6,1,'','','','','','','','1900-01-01',0,0,'','','',0,0,'','1900-01-01',0,0,'','','',0,'','','','',0,'','','2017-11-30 18:24:52',0,'2017-11-30 18:24:52',0);
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
INSERT INTO `Receipt` VALUES (1,0,1,1,2,0,0,0,'2014-03-01 00:00:00','1234',7000.0000,'',10,'',4,'Reversed by receipt RCPT00000005','','2017-11-30 19:24:46',0,'2017-11-30 18:48:02',0),(2,0,1,4,2,0,0,0,'2016-07-01 00:00:00','2345',8300.0000,'',10,'',4,'Reversed by receipt RCPT00000004','','2017-11-30 19:16:53',0,'2017-11-30 18:49:17',0),(3,0,1,4,2,0,3,0,'2016-07-01 00:00:00','2456',8300.0000,'',25,'ASM(2) d 12999 8300.00,c 12001 8300.00',2,'','','2017-12-05 17:06:25',0,'2017-11-30 19:13:23',0),(4,2,1,4,2,0,0,0,'2016-07-01 00:00:00','2345',-8300.0000,'',10,'',4,'Reversal of receipt RCPT00000002','','2017-11-30 19:16:53',0,'2017-11-30 19:13:59',0),(5,1,1,1,2,0,0,0,'2014-03-01 00:00:00','1234',-7000.0000,'',10,'',4,'Reversal of receipt RCPT00000001','','2017-11-30 19:24:46',0,'2017-11-30 19:23:47',0),(6,0,1,1,2,0,4,0,'2014-03-01 00:00:00','3457',7000.0000,'',25,'ASM(1) d 12999 7000.00,c 12001 7000.00',2,'','','2017-11-30 19:46:56',0,'2017-11-30 19:24:32',0),(7,0,1,1,6,0,0,0,'2017-01-01 00:00:00','2354',3750.0000,'',25,'ASM(4) d 12999 3750.00,c 12001 3750.00',4,'Reversed by receipt RCPT00000014','','2017-12-05 16:19:04',0,'2017-11-30 19:44:52',0),(8,0,1,1,6,0,0,0,'2017-01-01 00:00:00','',3750.0000,'',25,'ASM(4) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:31',0,'2017-12-05 16:09:37',0),(9,0,1,4,2,0,0,0,'2017-01-01 00:00:00','',4150.0000,'',25,'ASM(28) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:10:06',0),(10,0,1,3,6,0,0,0,'2017-02-01 00:00:00','',8350.0000,'',25,'ASM(42) d 12999 628.45,c 12001 628.45,ASM(17) d 12999 4000.00,c 12001 4000.00,ASM(43) d 12999 175.00,c 12001 175.00,ASM(18) d 12999 3546.55,c 12001 3546.55',2,'','','2017-12-05 20:50:21',0,'2017-12-05 16:12:02',0),(11,0,1,1,6,0,0,0,'2017-02-01 00:00:00','',3750.0000,'',25,'ASM(5) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:31',0,'2017-12-05 16:12:32',0),(12,0,1,4,2,0,0,0,'2017-02-01 00:00:00','',4150.0000,'',25,'ASM(29) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:13:44',0),(13,0,1,3,6,0,0,0,'2016-11-15 00:00:00','',12000.0000,'',25,'ASM(47) d 12999 4000.00,c 12001 4000.00,ASM(48) d 12999 4000.00,c 12001 4000.00,ASM(16) d 12999 4000.00,c 12001 4000.00',2,'3 month rent in advance','','2017-12-05 20:50:21',0,'2017-12-05 16:16:37',0),(14,7,1,1,6,0,0,0,'2017-01-01 00:00:00','2354',-3750.0000,'',25,'ASM(4) d 12999 3750.00,c 12001 3750.00',4,'Reversal of receipt RCPT00000007','','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(15,0,1,1,6,0,0,0,'2017-03-01 00:00:00','',3750.0000,'',25,'ASM(6) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:31',0,'2017-12-05 16:20:28',0),(16,0,1,4,2,0,0,0,'2017-03-01 00:00:00','',4150.0000,'',25,'ASM(30) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:22:08',0),(17,0,1,1,6,0,0,0,'2017-04-01 00:00:00','',3750.0000,'',25,'ASM(7) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:31',0,'2017-12-05 16:22:50',0),(18,0,1,4,2,0,0,0,'2017-04-01 00:00:00','',4150.0000,'',25,'ASM(31) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:23:12',0),(19,0,1,1,6,0,0,0,'2017-05-01 00:00:00','',3750.0000,'',25,'ASM(8) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:32',0,'2017-12-05 16:24:00',0),(20,0,1,4,2,0,0,0,'2017-05-01 00:00:00','',4150.0000,'',25,'ASM(32) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:24:18',0),(21,0,1,3,6,0,0,0,'2017-05-15 00:00:00','',13131.7900,'',25,'ASM(19) d 12999 4000.00,c 12001 4000.00,ASM(50) d 12999 350.00,c 12001 350.00,ASM(45) d 12999 81.79,c 12001 81.79,ASM(20) d 12999 4000.00,c 12001 4000.00,ASM(51) d 12999 350.00,c 12001 350.00,ASM(21) d 12999 4000.00,c 12001 4000.00,ASM(52) d 12999 350.00,c 12001 350.00',2,'3 month rent in advance and utilities overage','','2017-12-05 20:50:21',0,'2017-12-05 16:26:21',0),(22,0,1,1,6,0,0,0,'2017-06-01 00:00:00','',3750.0000,'',25,'ASM(9) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:32',0,'2017-12-05 16:27:03',0),(23,0,1,4,2,0,0,0,'2017-06-01 00:00:00','',4150.0000,'',25,'ASM(33) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:27:16',0),(24,0,1,1,6,0,0,0,'2017-07-01 00:00:00','',3750.0000,'',25,'ASM(10) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:32',0,'2017-12-05 16:27:58',0),(25,0,1,4,2,0,0,0,'2017-07-01 00:00:00','',4150.0000,'',25,'ASM(34) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:28:12',0),(26,0,1,1,6,0,0,0,'2017-08-01 00:00:00','',3750.0000,'',25,'ASM(11) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:32',0,'2017-12-05 16:29:16',0),(27,0,1,4,2,0,0,0,'2017-08-01 00:00:00','',4150.0000,'',25,'ASM(35) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:29:33',0),(28,0,1,3,6,0,0,0,'2017-08-15 00:00:00','',13050.0000,'',25,'ASM(22) d 12999 4000.00,c 12001 4000.00,ASM(53) d 12999 350.00,c 12001 350.00,ASM(23) d 12999 4000.00,c 12001 4000.00,ASM(54) d 12999 350.00,c 12001 350.00,ASM(24) d 12999 4000.00,c 12001 4000.00,ASM(55) d 12999 350.00,c 12001 350.00',2,'3 month rent in advance','','2017-12-05 20:50:22',0,'2017-12-05 16:29:59',0),(29,0,1,1,6,0,0,0,'2017-09-01 00:00:00','',3750.0000,'',25,'ASM(12) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:32',0,'2017-12-05 16:30:33',0),(30,0,1,4,2,0,0,0,'2017-09-01 00:00:00','',4150.0000,'',25,'ASM(36) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:30:51',0),(31,0,1,1,6,0,0,0,'2017-10-01 00:00:00','',3750.0000,'',25,'ASM(13) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:32',0,'2017-12-05 16:31:42',0),(32,0,1,4,2,0,0,0,'2017-10-01 00:00:00','',4150.0000,'',25,'ASM(37) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:31:56',0),(33,0,1,1,6,0,0,0,'2017-11-01 00:00:00','',3750.0000,'',25,'ASM(14) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:32',0,'2017-12-05 16:32:48',0),(34,0,1,4,2,0,0,0,'2017-11-01 00:00:00','',4150.0000,'',25,'ASM(38) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:33:11',0),(35,0,1,3,6,0,0,0,'2017-11-15 00:00:00','',13459.2800,'',25,'ASM(25) d 12999 4000.00,c 12001 4000.00,ASM(56) d 12999 350.00,c 12001 350.00,ASM(46) d 12999 409.28,c 12001 409.28,ASM(26) d 12999 4000.00,c 12001 4000.00,ASM(57) d 12999 350.00,c 12001 350.00,ASM(40) d 12999 4000.00,c 12001 4000.00,ASM(58) d 12999 350.00,c 12001 350.00',2,'3 month rent in advance and utilities overage','','2017-12-05 20:50:22',0,'2017-12-05 16:40:59',0),(36,0,1,1,6,0,0,0,'2017-12-01 00:00:00','',3750.0000,'',25,'',2,'','Kirsten Read','2017-12-06 11:51:39',0,'2017-12-05 16:42:24',0),(37,0,1,4,2,0,0,0,'2017-12-01 00:00:00','',4150.0000,'',25,'',2,'','','2017-12-06 11:51:45',0,'2017-12-05 16:42:35',0),(38,0,1,3,6,0,0,0,'2017-02-03 00:00:00','',628.4500,'',25,'ASM(18) d 12999 453.45,c 12001 453.45,ASM(44) d 12999 175.00,c 12001 175.00',2,'','','2017-12-05 20:50:21',0,'2017-12-05 19:44:51',0),(39,0,1,3,2,0,0,0,'2016-10-03 00:00:00','',4000.0000,'',25,'ASM(59) d 12999 4000.00,c 12001 4000.00',2,'','','2017-12-05 20:50:21',0,'2017-12-05 20:44:04',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ReceiptAllocation`
--

LOCK TABLES `ReceiptAllocation` WRITE;
/*!40000 ALTER TABLE `ReceiptAllocation` DISABLE KEYS */;
INSERT INTO `ReceiptAllocation` VALUES (1,1,1,0,'2014-03-01 00:00:00',7000.0000,0,4,'d 10104 _, c 10999 _','2017-11-30 19:23:47',0,'2017-11-30 18:48:02',0),(2,2,1,0,'2016-07-01 00:00:00',8300.0000,0,4,'d 10104 _, c 10999 _','2017-11-30 19:13:59',0,'2017-11-30 18:49:17',0),(3,3,1,0,'2016-07-01 00:00:00',8300.0000,0,0,'d 10999 _, c 12999 _','2017-11-30 19:13:23',0,'2017-11-30 19:13:23',0),(4,4,1,0,'2016-07-01 00:00:00',-8300.0000,0,4,'d 10104 _, c 10999 _','2017-11-30 19:13:59',0,'2017-11-30 19:13:59',0),(5,3,1,0,'2016-07-01 00:00:00',8300.0000,0,0,'d 10104 _, c 10999 _','2017-11-30 19:17:32',0,'2017-11-30 19:17:32',0),(6,5,1,0,'2014-03-01 00:00:00',-7000.0000,0,4,'d 10104 _, c 10999 _','2017-11-30 19:23:47',0,'2017-11-30 19:23:47',0),(7,6,1,0,'2014-03-01 00:00:00',7000.0000,0,0,'d 10999 _, c 12999 _','2017-11-30 19:24:32',0,'2017-11-30 19:24:32',0),(8,6,1,0,'2014-03-01 00:00:00',7000.0000,0,0,'d 10104 _, c 10999 _','2017-11-30 19:25:24',0,'2017-11-30 19:25:24',0),(9,7,1,0,'2017-01-01 00:00:00',3750.0000,0,4,'d 10999 _, c 12999 _','2017-12-05 16:19:04',0,'2017-11-30 19:44:52',0),(10,6,1,2,'2014-03-01 00:00:00',7000.0000,1,0,'ASM(1) d 12999 7000.00,c 12001 7000.00','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(11,7,1,2,'2017-01-01 00:00:00',3750.0000,4,4,'ASM(4) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:19:04',0,'2017-11-30 19:46:56',0),(12,8,1,0,'2017-01-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:09:37',0,'2017-12-05 16:09:37',0),(13,9,1,0,'2017-01-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:10:06',0,'2017-12-05 16:10:06',0),(14,10,1,0,'2017-02-01 00:00:00',8350.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:12:02',0,'2017-12-05 16:12:02',0),(15,11,1,0,'2017-02-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:12:32',0,'2017-12-05 16:12:32',0),(16,12,1,0,'2017-02-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:13:44',0,'2017-12-05 16:13:44',0),(17,13,1,0,'2016-11-15 00:00:00',12000.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:16:37',0,'2017-12-05 16:16:37',0),(18,14,1,0,'2017-01-01 00:00:00',-3750.0000,0,4,'d 10999 _, c 12999 _','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(19,14,1,2,'2017-12-05 16:19:04',-3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(20,15,1,0,'2017-03-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:20:28',0,'2017-12-05 16:20:28',0),(21,16,1,0,'2017-03-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:22:08',0,'2017-12-05 16:22:08',0),(22,17,1,0,'2017-04-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:22:50',0,'2017-12-05 16:22:50',0),(23,18,1,0,'2017-04-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:23:12',0,'2017-12-05 16:23:12',0),(24,19,1,0,'2017-05-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:24:00',0,'2017-12-05 16:24:00',0),(25,20,1,0,'2017-05-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:24:18',0,'2017-12-05 16:24:18',0),(26,21,1,0,'2017-05-15 00:00:00',13131.7900,0,0,'d 10999 _, c 12999 _','2017-12-05 16:26:21',0,'2017-12-05 16:26:21',0),(27,22,1,0,'2017-06-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:27:03',0,'2017-12-05 16:27:03',0),(28,23,1,0,'2017-06-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:27:16',0,'2017-12-05 16:27:16',0),(29,24,1,0,'2017-07-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:27:58',0,'2017-12-05 16:27:58',0),(30,25,1,0,'2017-07-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:28:12',0,'2017-12-05 16:28:12',0),(31,26,1,0,'2017-08-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:29:16',0,'2017-12-05 16:29:16',0),(32,27,1,0,'2017-08-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:29:33',0,'2017-12-05 16:29:33',0),(33,28,1,0,'2017-08-15 00:00:00',13050.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:29:59',0,'2017-12-05 16:29:59',0),(34,29,1,0,'2017-09-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:30:33',0,'2017-12-05 16:30:33',0),(35,30,1,0,'2017-09-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:30:51',0,'2017-12-05 16:30:51',0),(36,31,1,0,'2017-10-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:31:42',0,'2017-12-05 16:31:42',0),(37,32,1,0,'2017-10-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:31:56',0,'2017-12-05 16:31:56',0),(38,33,1,0,'2017-11-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:32:48',0,'2017-12-05 16:32:48',0),(39,34,1,0,'2017-11-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:33:11',0,'2017-12-05 16:33:11',0),(40,35,1,0,'2017-11-15 00:00:00',13459.2800,0,0,'d 10999 _, c 12999 _','2017-12-05 16:40:59',0,'2017-12-05 16:40:59',0),(41,36,1,0,'2017-12-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:42:24',0,'2017-12-05 16:42:24',0),(42,37,1,0,'2017-12-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:42:35',0,'2017-12-05 16:42:35',0),(43,8,1,2,'2017-01-01 00:00:00',3750.0000,4,0,'ASM(4) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(44,11,1,2,'2017-02-01 00:00:00',3750.0000,5,0,'ASM(5) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(45,15,1,2,'2017-03-01 00:00:00',3750.0000,6,0,'ASM(6) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(46,17,1,2,'2017-04-01 00:00:00',3750.0000,7,0,'ASM(7) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(47,19,1,2,'2017-05-01 00:00:00',3750.0000,8,0,'ASM(8) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(48,22,1,2,'2017-06-01 00:00:00',3750.0000,9,0,'ASM(9) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(49,24,1,2,'2017-07-01 00:00:00',3750.0000,10,0,'ASM(10) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(50,26,1,2,'2017-08-01 00:00:00',3750.0000,11,0,'ASM(11) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(51,29,1,2,'2017-09-01 00:00:00',3750.0000,12,0,'ASM(12) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(52,31,1,2,'2017-10-01 00:00:00',3750.0000,13,0,'ASM(13) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(53,33,1,2,'2017-11-01 00:00:00',3750.0000,14,0,'ASM(14) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(54,36,1,2,'2017-12-01 00:00:00',3750.0000,39,0,'ASM(39) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(55,3,1,4,'2016-07-01 00:00:00',8300.0000,2,0,'ASM(2) d 12999 8300.00,c 12001 8300.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(56,9,1,4,'2017-01-01 00:00:00',4150.0000,28,0,'ASM(28) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(57,12,1,4,'2017-02-01 00:00:00',4150.0000,29,0,'ASM(29) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(58,16,1,4,'2017-03-01 00:00:00',4150.0000,30,0,'ASM(30) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(59,18,1,4,'2017-04-01 00:00:00',4150.0000,31,0,'ASM(31) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(60,20,1,4,'2017-05-01 00:00:00',4150.0000,32,0,'ASM(32) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(61,23,1,4,'2017-06-01 00:00:00',4150.0000,33,0,'ASM(33) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(62,25,1,4,'2017-07-01 00:00:00',4150.0000,34,0,'ASM(34) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(63,27,1,4,'2017-08-01 00:00:00',4150.0000,35,0,'ASM(35) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(64,30,1,4,'2017-09-01 00:00:00',4150.0000,36,0,'ASM(36) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(65,32,1,4,'2017-10-01 00:00:00',4150.0000,37,0,'ASM(37) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(66,34,1,4,'2017-11-01 00:00:00',4150.0000,38,0,'ASM(38) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(67,37,1,4,'2017-12-01 00:00:00',4150.0000,41,0,'ASM(41) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(68,38,1,0,'2017-02-03 00:00:00',628.4500,0,0,'d 10999 _, c 12999 _','2017-12-05 19:44:51',0,'2017-12-05 19:44:51',0),(69,39,1,0,'2016-10-03 00:00:00',4000.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 20:44:04',0,'2017-12-05 20:44:04',0),(70,39,1,3,'2016-10-03 00:00:00',4000.0000,59,0,'ASM(59) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(71,13,1,3,'2016-11-11 00:00:00',4000.0000,47,0,'ASM(47) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(72,13,1,3,'2016-12-01 00:00:00',4000.0000,48,0,'ASM(48) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(73,13,1,3,'2017-01-01 00:00:00',4000.0000,16,0,'ASM(16) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(74,10,1,3,'2017-02-03 00:00:00',628.4500,42,0,'ASM(42) d 12999 628.45,c 12001 628.45','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(75,10,1,3,'2017-02-13 00:00:00',4000.0000,17,0,'ASM(17) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(76,10,1,3,'2017-02-28 00:00:00',175.0000,43,0,'ASM(43) d 12999 175.00,c 12001 175.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(77,10,1,3,'2017-03-01 00:00:00',3546.5500,18,0,'ASM(18) d 12999 3546.55,c 12001 3546.55','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(78,38,1,3,'2017-03-01 00:00:00',453.4500,18,0,'ASM(18) d 12999 453.45,c 12001 453.45','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(79,38,1,3,'2017-03-31 00:00:00',175.0000,44,0,'ASM(44) d 12999 175.00,c 12001 175.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(80,21,1,3,'2017-05-12 00:00:00',4000.0000,19,0,'ASM(19) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(81,21,1,3,'2017-05-12 00:00:00',350.0000,50,0,'ASM(50) d 12999 350.00,c 12001 350.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(82,21,1,3,'2017-05-12 00:00:00',81.7900,45,0,'ASM(45) d 12999 81.79,c 12001 81.79','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(83,21,1,3,'2017-05-12 00:00:00',4000.0000,20,0,'ASM(20) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(84,21,1,3,'2017-05-12 00:00:00',350.0000,51,0,'ASM(51) d 12999 350.00,c 12001 350.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(85,21,1,3,'2017-06-01 00:00:00',4000.0000,21,0,'ASM(21) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(86,21,1,3,'2017-06-01 00:00:00',350.0000,52,0,'ASM(52) d 12999 350.00,c 12001 350.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(87,28,1,3,'2017-08-15 00:00:00',4000.0000,22,0,'ASM(22) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(88,28,1,3,'2017-08-15 00:00:00',350.0000,53,0,'ASM(53) d 12999 350.00,c 12001 350.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(89,28,1,3,'2017-08-15 00:00:00',4000.0000,23,0,'ASM(23) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(90,28,1,3,'2017-08-15 00:00:00',350.0000,54,0,'ASM(54) d 12999 350.00,c 12001 350.00','2017-12-05 20:50:21',0,'2017-12-05 20:50:21',0),(91,28,1,3,'2017-09-01 00:00:00',4000.0000,24,0,'ASM(24) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(92,28,1,3,'2017-09-01 00:00:00',350.0000,55,0,'ASM(55) d 12999 350.00,c 12001 350.00','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(93,35,1,3,'2017-11-14 00:00:00',4000.0000,25,0,'ASM(25) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(94,35,1,3,'2017-11-14 00:00:00',350.0000,56,0,'ASM(56) d 12999 350.00,c 12001 350.00','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(95,35,1,3,'2017-11-14 00:00:00',409.2800,46,0,'ASM(46) d 12999 409.28,c 12001 409.28','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(96,35,1,3,'2017-11-14 00:00:00',4000.0000,26,0,'ASM(26) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(97,35,1,3,'2017-11-14 00:00:00',350.0000,57,0,'ASM(57) d 12999 350.00,c 12001 350.00','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(98,35,1,3,'2017-12-01 00:00:00',4000.0000,40,0,'ASM(40) d 12999 4000.00,c 12001 4000.00','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0),(99,35,1,3,'2017-12-01 00:00:00',350.0000,58,0,'ASM(58) d 12999 350.00,c 12001 350.00','2017-12-05 20:50:22',0,'2017-12-05 20:50:22',0);
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
INSERT INTO `Rentable` VALUES (1,1,'309 Rexford',1,0,'0000-00-00 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0,''),(2,1,'309 1/2 Rexford',1,0,'0000-00-00 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0,''),(3,1,'311 Rexford',1,0,'0000-00-00 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0,''),(4,1,'311 1/2 Rexford',1,0,'0000-00-00 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0,'');
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
INSERT INTO `RentableMarketRate` VALUES (1,1,1,3750.0000,'2014-01-01 00:00:00','9999-03-01 00:00:00','2018-01-10 18:32:37',0,'2017-11-28 03:44:18',0),(2,2,1,4000.0000,'2014-01-01 00:00:00','9999-05-01 00:00:00','2018-01-10 18:32:37',0,'2017-11-28 03:44:18',0),(3,3,1,4150.0000,'2014-01-01 00:00:00','9999-04-01 00:00:00','2018-01-10 18:32:37',0,'2017-11-28 03:44:18',0),(4,4,1,2500.0000,'2014-01-01 00:00:00','9999-01-01 00:00:00','2018-01-10 18:32:37',0,'2017-11-28 03:44:18',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableStatus`
--

LOCK TABLES `RentableStatus` WRITE;
/*!40000 ALTER TABLE `RentableStatus` DISABLE KEYS */;
INSERT INTO `RentableStatus` VALUES (1,1,1,1,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0),(2,2,1,1,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0),(3,3,1,1,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0),(4,4,1,4,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableTypeRef`
--

LOCK TABLES `RentableTypeRef` WRITE;
/*!40000 ALTER TABLE `RentableTypeRef` DISABLE KEYS */;
INSERT INTO `RentableTypeRef` VALUES (1,1,1,1,0,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0),(2,2,1,2,0,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0),(3,3,1,3,0,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0),(4,4,1,4,0,0,'2014-01-01 00:00:00','9999-01-01 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0);
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
INSERT INTO `RentableTypes` VALUES (1,1,'Rex1','309 Rexford',6,4,4,0,0,'2017-11-28 03:44:18',0,'2017-11-28 03:44:18',0),(2,1,'Rex2','309 1/2 Rexford',6,4,4,0,0,'2017-11-28 03:44:18',0,'2017-11-28 03:44:18',0),(3,1,'Rex3','311 Rexford',6,4,4,0,0,'2017-11-28 03:44:18',0,'2017-11-28 03:44:18',0),(4,1,'Rex4','311 1/2 Rexford',6,4,4,0,0,'2017-11-28 03:44:18',0,'2017-11-28 03:44:18',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableUsers`
--

LOCK TABLES `RentableUsers` WRITE;
/*!40000 ALTER TABLE `RentableUsers` DISABLE KEYS */;
INSERT INTO `RentableUsers` VALUES (1,1,1,1,'2014-03-01','2018-03-01','2018-01-10 18:32:36',0,'2017-11-30 18:22:29',0),(2,1,1,2,'2014-03-01','2018-03-01','2018-01-10 18:32:36',0,'2017-11-30 18:23:03',0),(4,1,1,6,'2014-03-01','2018-03-01','2018-01-10 18:32:36',0,'2017-11-30 18:30:18',0),(5,2,1,3,'2016-10-01','2018-01-01','2018-01-10 18:32:36',0,'2017-11-30 18:33:01',0),(6,3,1,4,'2016-07-01','2018-07-01','2018-01-10 18:32:36',0,'2017-11-30 18:36:10',0),(7,3,1,5,'2016-07-01','2018-07-01','2018-01-10 18:32:36',0,'2017-11-30 18:36:30',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreement`
--

LOCK TABLES `RentalAgreement` WRITE;
/*!40000 ALTER TABLE `RentalAgreement` DISABLE KEYS */;
INSERT INTO `RentalAgreement` VALUES (2,0,1,0,'2014-03-01','2018-03-01','2014-03-01','2018-03-01','2014-03-01','2018-03-01','2014-03-01',0,0,2,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2017-11-30 18:30:41',0,'2017-11-30 18:17:55',0),(3,0,1,0,'2016-10-01','2018-01-01','2016-10-01','2018-01-01','2016-10-01','2018-01-01','2016-10-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2017-11-30 18:33:53',0,'2017-11-30 18:29:25',0),(4,0,1,0,'2016-07-01','2018-07-01','2016-07-01','2018-07-01','2016-07-01','2018-07-01','2016-07-01',0,0,2,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2017-11-30 18:37:13',0,'2017-11-30 18:33:59',0),(6,0,1,0,'2018-01-03','2019-01-03','2018-01-03','2019-01-03','2018-01-03','2019-01-03','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-03 07:07:58',0,'2018-01-03 07:07:58',0),(7,0,1,0,'2018-01-03','2019-01-03','2018-01-03','2019-01-03','2018-01-03','2019-01-03','2018-01-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','',0,'2018-01-03 07:08:22',0,'2018-01-03 07:08:22',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementPayors`
--

LOCK TABLES `RentalAgreementPayors` WRITE;
/*!40000 ALTER TABLE `RentalAgreementPayors` DISABLE KEYS */;
INSERT INTO `RentalAgreementPayors` VALUES (1,2,1,1,'2014-03-01','2018-03-01',0,'2018-01-10 18:32:35',0,'2017-11-30 18:21:00',0),(2,2,1,2,'2017-11-30','2018-03-01',0,'2018-01-10 18:32:35',0,'2017-11-30 18:21:57',0),(3,2,1,2,'2018-03-01','2018-03-01',0,'2018-01-10 18:32:35',0,'2017-11-30 18:28:09',0),(4,3,1,3,'2016-10-01','2018-01-01',0,'2018-01-10 18:32:35',0,'2017-11-30 18:32:33',0),(5,4,1,4,'2016-07-01','2018-02-01',0,'2018-01-10 18:32:35',0,'2017-11-30 18:35:19',0),(6,4,1,5,'2016-07-01','2018-07-01',0,'2018-01-10 18:32:35',0,'2017-11-30 18:35:41',0);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RARID`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementRentables`
--

LOCK TABLES `RentalAgreementRentables` WRITE;
/*!40000 ALTER TABLE `RentalAgreementRentables` DISABLE KEYS */;
INSERT INTO `RentalAgreementRentables` VALUES (1,2,1,1,0,0,3750.0000,'2014-03-01','2018-03-01','2018-01-10 18:32:35',0,'2017-11-30 18:20:15',0),(2,3,1,2,0,0,4000.0000,'2016-10-01','2018-01-01','2018-01-10 18:32:35',0,'2017-11-30 18:32:13',0),(3,4,1,3,0,0,4150.0000,'2016-07-01','2018-07-01','2018-01-10 18:32:35',0,'2017-11-30 18:34:33',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TWS`
--

LOCK TABLES `TWS` WRITE;
/*!40000 ALTER TABLE `TWS` DISABLE KEYS */;
INSERT INTO `TWS` VALUES (2,'CreateAssessmentInstances','','CreateAssessmentInstances','2018-01-17 00:00:00','ip-172-31-51-141.ec2.internal',4,'2018-01-16 00:00:02','2018-01-16 00:00:02','2017-11-27 21:24:27','2018-01-16 00:00:02'),(3,'CleanRARBalanceCache','','CleanRARBalanceCache','2018-01-16 18:34:07','ip-172-31-51-141.ec2.internal',4,'2018-01-16 18:29:07','2018-01-16 18:29:07','2017-11-30 17:39:57','2018-01-16 18:29:06'),(4,'CleanSecDepBalanceCache','','CleanSecDepBalanceCache','2018-01-16 18:33:47','ip-172-31-51-141.ec2.internal',4,'2018-01-16 18:28:47','2018-01-16 18:28:47','2017-11-30 17:39:57','2018-01-16 18:28:46'),(5,'CleanAcctSliceCache','','CleanAcctSliceCache','2018-01-16 18:33:27','ip-172-31-51-141.ec2.internal',4,'2018-01-16 18:28:27','2018-01-16 18:28:27','2017-11-30 17:39:57','2018-01-16 18:28:26'),(6,'CleanARSliceCache','','CleanARSliceCache','2018-01-16 18:37:56','ip-172-31-51-141.ec2.internal',4,'2018-01-16 18:32:47','2018-01-16 18:32:56','2017-11-30 17:39:57','2018-01-16 18:32:55'),(7,'CreateAssessmentInstances','','CreateAssessmentInstances','2018-01-11 00:00:00','ip-172-31-56-225.ec2.internal',4,'2018-01-10 22:45:36','2018-01-10 22:45:36','2018-01-10 22:45:25','2018-01-10 22:45:35'),(8,'CleanRARBalanceCache','','CleanRARBalanceCache','2018-01-10 23:56:52','ip-172-31-56-225.ec2.internal',4,'2018-01-10 23:51:52','2018-01-10 23:51:52','2018-01-10 22:45:25','2018-01-10 23:51:52'),(9,'CleanSecDepBalanceCache','','CleanSecDepBalanceCache','2018-01-10 23:56:52','ip-172-31-56-225.ec2.internal',4,'2018-01-10 23:51:52','2018-01-10 23:51:52','2018-01-10 22:45:25','2018-01-10 23:51:52'),(10,'CleanAcctSliceCache','','CleanAcctSliceCache','2018-01-10 23:56:52','ip-172-31-56-225.ec2.internal',4,'2018-01-10 23:51:52','2018-01-10 23:51:52','2018-01-10 22:45:25','2018-01-10 23:51:52'),(11,'CleanARSliceCache','','CleanARSliceCache','2018-01-10 23:56:52','ip-172-31-56-225.ec2.internal',4,'2018-01-10 23:51:52','2018-01-10 23:51:52','2018-01-10 22:45:25','2018-01-10 23:51:52'),(12,'CreateAssessmentInstances','','CreateAssessmentInstances','2018-01-12 00:00:00','ip-172-31-52-71.ec2.internal',4,'2018-01-11 00:00:05','2018-01-11 00:00:05','2018-01-10 22:47:08','2018-01-11 00:00:05'),(13,'CleanRARBalanceCache','','CleanRARBalanceCache','2018-01-11 00:19:05','ip-172-31-52-71.ec2.internal',4,'2018-01-11 00:14:05','2018-01-11 00:14:05','2018-01-10 22:47:08','2018-01-11 00:14:05'),(14,'CleanSecDepBalanceCache','','CleanSecDepBalanceCache','2018-01-11 00:19:05','ip-172-31-52-71.ec2.internal',4,'2018-01-11 00:14:05','2018-01-11 00:14:05','2018-01-10 22:47:08','2018-01-11 00:14:05'),(15,'CleanAcctSliceCache','','CleanAcctSliceCache','2018-01-11 00:19:05','ip-172-31-52-71.ec2.internal',4,'2018-01-11 00:14:05','2018-01-11 00:14:05','2018-01-10 22:47:08','2018-01-11 00:14:05'),(16,'CleanARSliceCache','','CleanARSliceCache','2018-01-11 00:19:05','ip-172-31-52-71.ec2.internal',4,'2018-01-11 00:14:05','2018-01-11 00:14:05','2018-01-10 22:47:08','2018-01-11 00:14:05');
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
INSERT INTO `Transactant` VALUES (1,1,0,'Aaron','','Read','','',0,'','','','','','','','','','','',0,'','2017-11-30 18:15:20',0,'2017-11-30 18:15:20',0),(2,1,0,'Kirsten','','Read','','',0,'','','','','','','','','','','',0,'','2017-11-30 18:15:28',0,'2017-11-30 18:15:28',0),(3,1,0,'Alex','','Vahabzadeh','','Beaumont Partners LP',1,'','','','','','','','','','','',0,'','2017-11-30 18:17:13',0,'2017-11-30 18:16:10',0),(4,1,0,'Kevin','','Mills','','',0,'','','','','','','','','','','',0,'','2017-11-30 18:16:17',0,'2017-11-30 18:16:17',0),(5,1,0,'Lauren','','Beck','','',0,'','','','','','','','','','','',0,'','2017-11-30 18:16:28',0,'2017-11-30 18:16:28',0),(6,1,0,'Child','','Read','','',0,'','','','','','','','','','','',0,'','2017-11-30 18:24:52',0,'2017-11-30 18:24:52',0);
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
  PRIMARY KEY (`TCID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` VALUES (1,1,0,'1900-01-01','','','','','',1,0,'',0,'2017-11-30 18:15:20',0,'2017-11-30 18:15:20',0),(2,1,0,'1900-01-01','','','','','',1,0,'',0,'2017-11-30 18:15:28',0,'2017-11-30 18:15:28',0),(3,1,0,'1900-01-01','','','','','',1,0,'',0,'2017-11-30 18:16:10',0,'2017-11-30 18:16:10',0),(4,1,0,'1900-01-01','','','','','',1,0,'',0,'2017-11-30 18:16:17',0,'2017-11-30 18:16:17',0),(5,1,0,'1900-01-01','','','','','',1,0,'',0,'2017-11-30 18:16:28',0,'2017-11-30 18:16:28',0),(6,1,0,'1900-01-01','','','','','',1,0,'',0,'2017-11-30 18:24:52',0,'2017-11-30 18:24:52',0);
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

-- Dump completed on 2018-06-20 14:16:08
