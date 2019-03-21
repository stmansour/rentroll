-- MySQL dump 10.13  Distrib 5.7.22, for osx10.12 (x86_64)
--
-- Host: localhost    Database: rentroll.tbf
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
  `DefaultRentCycle` smallint(6) NOT NULL DEFAULT '0',
  `DefaultProrationCycle` smallint(6) NOT NULL DEFAULT '0',
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
INSERT INTO `AR` VALUES (2,1,'Application Fee Received',0,1,0,6,46,'Application fee taken, no assessment made','1900-01-01 00:00:00','9999-12-31 00:00:00',5,0.0000,0,0,'2017-12-05 17:24:15',0,'2017-11-10 23:24:23',0),(4,1,'Bad Debt Write-Off',0,2,0,71,9,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(5,1,'Bank Service Fee (Security Deposit Account)',0,2,0,72,4,'','1900-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-28 18:54:17',0,'2017-11-10 23:24:23',0),(6,1,'Bank Service Fee (Operating Account)',0,2,0,72,3,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(8,1,'Damage Fee',0,0,0,9,59,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(9,1,'Deposit to Security Deposit Account (FRB96953)',0,1,0,4,6,'','1900-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-28 18:55:11',0,'2017-11-10 23:24:23',0),(10,1,'Deposit to Operating Account (FRB54320)',0,1,0,3,6,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(11,1,'Electric Base Fee',0,0,0,9,36,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(12,1,'Electric Overage',0,0,0,9,37,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(13,1,'Eviction Fee Reimbursement',0,0,0,9,56,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(14,1,'Auto-Generated Floating Deposit Assessment',0,3,0,9,12,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(16,1,'Gas Base Fee',0,0,0,9,40,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(17,1,'Gas Base Overage',0,0,0,9,41,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(18,1,'Insufficient Funds Fee',0,0,0,9,48,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(19,1,'Late Fee',0,0,0,9,47,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(20,1,'Month to Month Fee',0,0,0,9,49,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(21,1,'No Show / Termination Fee',0,0,0,9,51,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(22,1,'Other Special Tenant Charges',0,0,0,9,61,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(23,1,'Pet Fee',0,0,0,9,52,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(24,1,'Pet Rent',0,0,0,9,53,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(25,1,'Receive a Payment',0,1,0,6,10,'','1900-01-01 00:00:00','9999-12-28 00:00:00',0,0.0000,0,0,'2018-11-07 23:11:39',198,'2017-11-10 23:24:23',0),(27,1,'Gross Scheduled Rent',0,0,0,9,17,'','1900-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-28 18:40:18',0,'2017-11-10 23:24:23',0),(28,1,'Security Deposit Assessment',0,0,0,9,11,'normal deposit','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(29,1,'Security Deposit Forfeiture',0,0,0,11,58,'Forfeit','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(30,1,'Security Deposit Refund from Operating Account',0,0,0,11,3,'Refund','1900-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-28 19:23:08',0,'2017-11-10 23:24:23',0),(31,1,'Special Cleaning Fee',0,0,0,9,55,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(32,1,'Tenant Expense Chargeback',0,0,0,9,54,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(33,1,'Vending Income from Credit Card',0,1,0,7,65,'','1900-01-01 00:00:00','9999-12-31 00:00:00',5,0.0000,0,0,'2017-11-28 19:26:13',0,'2017-11-10 23:24:23',0),(34,1,'Water and Sewer Base Fee',0,0,0,9,38,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(35,1,'Water and Sewer Overage',0,0,0,9,39,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-27 21:50:36',0,'2017-11-10 23:24:23',0),(36,1,'Auto-Generated Application Fee Assessment',0,3,0,9,46,'','1900-01-01 00:00:00','9999-12-31 00:00:00',0,0.0000,0,0,'2017-11-28 18:57:23',0,'2017-11-10 23:24:23',0);
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
  `AssocElemType` bigint(20) NOT NULL DEFAULT '0',
  `AssocElemID` bigint(20) NOT NULL DEFAULT '0',
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
) ENGINE=InnoDB AUTO_INCREMENT=111 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Assessments`
--

LOCK TABLES `Assessments` WRITE;
/*!40000 ALTER TABLE `Assessments` DISABLE KEYS */;
INSERT INTO `Assessments` VALUES (1,0,0,0,1,1,0,0,2,7000.0000,'2014-03-01 00:00:00','2014-03-01 00:00:00',0,0,0,'',28,2,'','2017-11-30 19:46:56',0,'2017-11-30 18:39:27',0),(2,0,0,0,1,3,0,0,4,8300.0000,'2016-07-01 00:00:00','2016-07-01 00:00:00',0,0,0,'',28,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:41:00',0),(3,0,0,0,1,1,0,0,2,3750.0000,'2017-01-01 00:00:00','2018-02-01 00:00:00',6,4,0,'',27,0,'','2018-02-27 19:38:38',200,'2017-11-30 18:43:20',0),(4,3,0,0,1,1,0,0,2,3750.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:31',0,'2017-11-30 18:43:20',0),(5,3,0,0,1,1,0,0,2,3750.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:31',0,'2017-11-30 18:43:20',0),(6,3,0,0,1,1,0,0,2,3750.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:31',0,'2017-11-30 18:43:20',0),(7,3,0,0,1,1,0,0,2,3750.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:31',0,'2017-11-30 18:43:20',0),(8,3,0,0,1,1,0,0,2,3750.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:31',0,'2017-11-30 18:43:20',0),(9,3,0,0,1,1,0,0,2,3750.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:32',0,'2017-11-30 18:43:20',0),(10,3,0,0,1,1,0,0,2,3750.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:32',0,'2017-11-30 18:43:20',0),(11,3,0,0,1,1,0,0,2,3750.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:32',0,'2017-11-30 18:43:20',0),(12,3,0,0,1,1,0,0,2,3750.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:32',0,'2017-11-30 18:43:20',0),(13,3,0,0,1,1,0,0,2,3750.0000,'2017-10-01 00:00:00','2017-10-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:32',0,'2017-11-30 18:43:20',0),(14,3,0,0,1,1,0,0,2,3750.0000,'2017-11-01 00:00:00','2017-11-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:32',0,'2017-11-30 18:43:20',0),(15,0,0,0,1,2,0,0,3,4000.0000,'2017-01-01 00:00:00','2018-01-01 00:00:00',6,4,0,'',27,0,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(16,15,0,0,1,2,0,0,3,4000.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',27,2,'','2018-02-27 20:16:52',200,'2017-11-30 18:45:17',0),(17,15,0,0,1,2,0,0,3,4000.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',27,2,'','2018-02-27 20:16:52',200,'2017-11-30 18:45:17',0),(18,15,0,0,1,2,0,0,3,4000.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',27,2,'','2018-02-27 20:16:52',200,'2017-11-30 18:45:17',0),(19,15,0,0,1,2,0,0,3,4000.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',27,2,'','2018-02-27 20:16:52',200,'2017-11-30 18:45:17',0),(20,15,0,0,1,2,0,0,3,4000.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',27,2,'','2018-02-27 20:16:52',200,'2017-11-30 18:45:17',0),(21,15,0,0,1,2,0,0,3,4000.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',27,2,'','2018-02-27 20:16:52',200,'2017-11-30 18:45:17',0),(22,15,0,0,1,2,0,0,3,4000.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',27,2,'','2018-02-27 20:16:52',200,'2017-11-30 18:45:17',0),(23,15,0,0,1,2,0,0,3,4000.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',27,2,'','2018-02-27 20:16:53',200,'2017-11-30 18:45:17',0),(24,15,0,0,1,2,0,0,3,4000.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',27,2,'','2018-02-27 20:16:53',200,'2017-11-30 18:45:17',0),(25,15,0,0,1,2,0,0,3,4000.0000,'2017-10-01 00:00:00','2017-10-02 00:00:00',6,4,0,'',27,2,'','2018-02-27 20:16:53',200,'2017-11-30 18:45:17',0),(26,15,0,0,1,2,0,0,3,4000.0000,'2017-11-01 00:00:00','2017-11-02 00:00:00',6,4,0,'',27,2,'','2018-02-27 20:16:53',200,'2017-11-30 18:45:17',0),(27,0,0,0,1,3,0,0,4,4150.0000,'2017-01-01 00:00:00','2018-02-27 00:00:00',6,4,0,'',27,0,'','2018-05-30 19:17:48',200,'2017-11-30 18:45:55',0),(28,27,0,0,1,3,0,0,4,4150.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(29,27,0,0,1,3,0,0,4,4150.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(30,27,0,0,1,3,0,0,4,4150.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(31,27,0,0,1,3,0,0,4,4150.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(32,27,0,0,1,3,0,0,4,4150.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(33,27,0,0,1,3,0,0,4,4150.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(34,27,0,0,1,3,0,0,4,4150.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(35,27,0,0,1,3,0,0,4,4150.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(36,27,0,0,1,3,0,0,4,4150.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(37,27,0,0,1,3,0,0,4,4150.0000,'2017-10-01 00:00:00','2017-10-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(38,27,0,0,1,3,0,0,4,4150.0000,'2017-11-01 00:00:00','2017-11-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-11-30 18:45:55',0),(39,3,0,0,1,1,0,0,2,3750.0000,'2017-12-01 00:00:00','2017-12-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 16:59:32',0,'2017-12-01 00:00:04',0),(40,15,0,0,1,2,0,0,3,4000.0000,'2017-12-01 00:00:00','2017-12-02 00:00:00',6,4,0,'',27,2,'','2018-02-28 17:19:56',211,'2017-12-01 00:00:04',0),(41,27,0,0,1,3,0,0,4,4150.0000,'2017-12-01 00:00:00','2017-12-02 00:00:00',6,4,0,'',27,2,'','2017-12-05 17:06:25',0,'2017-12-01 00:00:04',0),(42,0,0,0,1,2,0,0,3,628.4500,'2017-01-31 00:00:00','2017-01-31 00:00:00',0,0,0,'',12,2,'utilities reimbursement','2018-02-27 20:16:52',200,'2017-12-05 16:01:46',0),(43,0,0,0,1,2,0,0,3,175.0000,'2017-02-28 00:00:00','2017-02-28 00:00:00',0,0,0,'',12,2,'','2018-02-27 20:16:52',200,'2017-12-05 16:02:25',0),(44,0,0,0,1,2,0,0,3,175.0000,'2017-03-31 00:00:00','2017-03-31 00:00:00',0,0,0,'',12,2,'','2018-02-27 20:16:52',200,'2017-12-05 16:03:13',0),(45,0,0,0,1,2,0,0,3,81.7900,'2017-04-15 00:00:00','2017-04-15 00:00:00',0,0,0,'',12,2,'','2018-02-27 20:16:52',200,'2017-12-05 16:03:41',0),(46,0,0,0,1,2,0,0,3,409.2800,'2017-10-31 00:00:00','2017-10-31 00:00:00',0,0,0,'',12,2,'','2018-02-27 20:16:53',200,'2017-12-05 16:07:34',0),(47,0,0,0,1,2,0,0,3,4000.0000,'2016-11-01 00:00:00','2016-11-01 00:00:00',0,0,0,'',27,2,'','2018-02-27 20:16:52',200,'2017-12-05 16:15:10',0),(48,0,0,0,1,2,0,0,3,4000.0000,'2016-12-01 00:00:00','2016-12-01 00:00:00',0,0,0,'',27,2,'','2018-02-27 20:16:52',200,'2017-12-05 16:15:45',0),(49,0,0,0,1,2,0,0,3,350.0000,'2017-04-01 00:00:00','2018-12-31 00:00:00',6,4,0,'',11,0,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(50,49,0,0,1,2,0,0,3,350.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',11,2,'','2018-02-27 20:16:52',200,'2017-12-05 17:49:46',0),(51,49,0,0,1,2,0,0,3,350.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',11,2,'','2018-02-27 20:16:52',200,'2017-12-05 17:49:46',0),(52,49,0,0,1,2,0,0,3,350.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',11,2,'','2018-02-27 20:16:52',200,'2017-12-05 17:49:46',0),(53,49,0,0,1,2,0,0,3,350.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',11,2,'','2018-02-27 20:16:52',200,'2017-12-05 17:49:46',0),(54,49,0,0,1,2,0,0,3,350.0000,'2017-08-01 00:00:00','2017-08-02 00:00:00',6,4,0,'',11,2,'','2018-02-27 20:16:53',200,'2017-12-05 17:49:46',0),(55,49,0,0,1,2,0,0,3,350.0000,'2017-09-01 00:00:00','2017-09-02 00:00:00',6,4,0,'',11,2,'','2018-02-27 20:16:53',200,'2017-12-05 17:49:46',0),(56,49,0,0,1,2,0,0,3,350.0000,'2017-10-01 00:00:00','2017-10-02 00:00:00',6,4,0,'',11,2,'','2018-02-27 20:16:53',200,'2017-12-05 17:49:46',0),(57,49,0,0,1,2,0,0,3,350.0000,'2017-11-01 00:00:00','2017-11-02 00:00:00',6,4,0,'',11,2,'','2018-02-27 20:16:53',200,'2017-12-05 17:49:46',0),(58,49,0,0,1,2,0,0,3,350.0000,'2017-12-01 00:00:00','2017-12-02 00:00:00',6,4,0,'',11,2,'','2018-02-28 17:19:56',211,'2017-12-05 17:49:46',0),(59,0,0,0,1,2,0,0,3,4000.0000,'2016-10-01 00:00:00','2016-10-01 00:00:00',0,0,0,'',27,2,'','2018-02-27 20:16:52',200,'2017-12-05 18:03:15',0),(60,0,0,0,1,2,0,0,3,628.4500,'2017-12-05 00:00:00','2017-12-05 00:00:00',0,0,0,'',12,4,'Reversed by ASM00000061','2017-12-05 19:41:01',0,'2017-12-05 18:23:25',0),(61,0,60,0,1,2,0,0,3,-628.4500,'2017-12-05 00:00:00','2017-12-05 00:00:00',0,0,0,'',12,4,'Reversal of ASM00000060','2017-12-05 19:41:01',0,'2017-12-05 19:41:01',0),(62,3,0,0,1,1,0,0,2,3750.0000,'2018-01-01 00:00:00','2018-01-02 00:00:00',6,4,0,'',27,2,'','2018-02-28 21:17:34',200,'2018-02-20 19:54:46',-99),(63,27,0,0,1,3,0,0,4,4150.0000,'2018-02-01 00:00:00','2018-02-02 00:00:00',6,4,0,'',27,2,'','2018-02-28 21:17:57',200,'2018-02-21 00:00:09',-99),(65,0,0,0,1,3,0,0,4,4150.0000,'2018-01-01 00:00:00','2018-01-01 00:00:00',0,0,0,'',27,2,'fixes manual deletion','2018-02-28 21:17:57',200,'2018-02-28 17:39:47',211),(66,0,0,0,1,2,0,0,3,4000.0000,'2018-01-01 00:00:00','2018-01-01 00:00:00',0,0,0,'',27,2,'','2018-02-28 21:17:49',200,'2018-02-28 17:41:00',211),(67,0,0,0,1,2,0,0,3,4000.0000,'2018-02-01 00:00:00','2018-02-01 00:00:00',0,0,0,'',27,2,'','2018-02-28 21:17:49',200,'2018-02-28 17:41:40',211),(68,0,0,0,1,2,0,0,3,350.0000,'2018-01-01 00:00:00','2018-01-01 00:00:00',0,0,0,'',11,2,'','2018-02-28 21:17:49',200,'2018-02-28 21:12:57',200),(69,0,0,0,1,2,0,0,3,350.0000,'2018-02-01 00:00:00','2018-02-01 00:00:00',0,0,0,'',11,2,'','2018-02-28 21:17:49',200,'2018-02-28 21:13:21',200),(70,27,0,0,1,3,0,0,4,4150.0000,'2018-03-01 00:00:00','2018-03-02 00:00:00',6,4,0,'',27,4,'Reversed by ASM00000080','2018-05-30 19:21:11',200,'2018-04-03 17:19:38',0),(71,49,0,0,1,2,0,0,3,350.0000,'2018-03-01 00:00:00','2018-03-02 00:00:00',6,4,0,'',11,4,'Reversed by ASM00000077','2018-05-30 19:20:48',200,'2018-04-03 17:19:38',0),(72,27,0,0,1,3,0,0,4,4150.0000,'2018-04-01 00:00:00','2018-04-02 00:00:00',6,4,0,'',27,4,'Reversed by ASM00000081','2018-05-30 19:21:11',200,'2018-04-03 17:19:48',0),(73,49,0,0,1,2,0,0,3,350.0000,'2018-04-01 00:00:00','2018-04-02 00:00:00',6,4,0,'',11,4,'Reversed by ASM00000078','2018-05-30 19:20:48',200,'2018-04-03 17:19:48',0),(74,27,0,0,1,3,0,0,4,4150.0000,'2018-05-01 00:00:00','2018-05-02 00:00:00',6,4,0,'',27,4,'Reversed by ASM00000076','2018-05-30 19:20:22',200,'2018-05-02 20:49:41',-1),(75,49,0,0,1,2,0,0,3,350.0000,'2018-05-01 00:00:00','2018-05-02 00:00:00',6,4,0,'',11,4,'Reversed by ASM00000079','2018-05-30 19:20:48',200,'2018-05-02 20:49:41',-1),(76,27,74,0,1,3,0,0,4,-4150.0000,'2018-05-01 00:00:00','2018-05-02 00:00:00',6,4,0,'',27,4,'Reversal of ASM00000074','2018-05-30 19:20:22',200,'2018-05-30 19:20:22',200),(77,49,71,0,1,2,0,0,3,-350.0000,'2018-03-01 00:00:00','2018-03-02 00:00:00',6,4,0,'',11,4,'Reversal of ASM00000071','2018-05-30 19:20:48',200,'2018-05-30 19:20:48',200),(78,49,73,0,1,2,0,0,3,-350.0000,'2018-04-01 00:00:00','2018-04-02 00:00:00',6,4,0,'',11,4,'Reversal of ASM00000073','2018-05-30 19:20:48',200,'2018-05-30 19:20:48',200),(79,49,75,0,1,2,0,0,3,-350.0000,'2018-05-01 00:00:00','2018-05-02 00:00:00',6,4,0,'',11,4,'Reversal of ASM00000075','2018-05-30 19:20:48',200,'2018-05-30 19:20:48',200),(80,27,70,0,1,3,0,0,4,-4150.0000,'2018-03-01 00:00:00','2018-03-02 00:00:00',6,4,0,'',27,4,'Reversal of ASM00000070','2018-05-30 19:21:11',200,'2018-05-30 19:21:11',200),(81,27,72,0,1,3,0,0,4,-4150.0000,'2018-04-01 00:00:00','2018-04-02 00:00:00',6,4,0,'',27,4,'Reversal of ASM00000072','2018-05-30 19:21:11',200,'2018-05-30 19:21:11',200),(82,0,0,0,1,2,0,0,3,350.0000,'2018-03-01 00:00:00','2018-12-31 00:00:00',6,4,0,'',11,0,'','2018-05-30 19:45:28',200,'2018-05-30 19:45:28',200),(83,82,0,0,1,2,0,0,3,350.0000,'2018-03-01 00:00:00','2018-03-02 00:00:00',6,4,0,'',11,2,'','2018-05-30 20:04:19',200,'2018-05-30 19:45:28',200),(84,82,0,0,1,2,0,0,3,350.0000,'2018-04-01 00:00:00','2018-04-02 00:00:00',6,4,0,'',11,2,'','2018-05-30 20:09:39',200,'2018-05-30 19:45:29',200),(85,82,0,0,1,2,0,0,3,350.0000,'2018-05-01 00:00:00','2018-05-02 00:00:00',6,4,0,'',11,2,'','2018-05-30 20:09:39',200,'2018-05-30 19:45:29',200),(86,0,0,0,1,2,0,0,3,4000.0000,'2018-03-01 00:00:00','2018-12-31 00:00:00',6,4,0,'',27,0,'','2018-05-30 19:50:19',200,'2018-05-30 19:50:19',200),(87,86,0,0,1,2,0,0,3,4000.0000,'2018-03-01 00:00:00','2018-03-02 00:00:00',6,4,0,'',27,2,'','2018-05-30 20:04:19',200,'2018-05-30 19:50:19',200),(88,86,0,0,1,2,0,0,3,4000.0000,'2018-04-01 00:00:00','2018-04-02 00:00:00',6,4,0,'',27,2,'','2018-05-30 20:09:39',200,'2018-05-30 19:50:19',200),(89,86,0,0,1,2,0,0,3,4000.0000,'2018-05-01 00:00:00','2018-05-02 00:00:00',6,4,0,'',27,2,'','2018-05-30 20:09:39',200,'2018-05-30 19:50:19',200),(90,49,0,0,1,2,0,0,3,350.0000,'2018-06-01 00:00:00','2018-06-02 00:00:00',6,4,0,'',11,0,'','2018-06-01 00:03:21',-1,'2018-06-01 00:03:21',-1),(91,82,0,0,1,2,0,0,3,350.0000,'2018-06-01 00:00:00','2018-06-02 00:00:00',6,4,0,'',11,0,'','2018-06-01 00:03:21',-1,'2018-06-01 00:03:21',-1),(92,86,0,0,1,2,0,0,3,4000.0000,'2018-06-01 00:00:00','2018-06-02 00:00:00',6,4,0,'',27,0,'','2018-06-01 00:03:21',-1,'2018-06-01 00:03:21',-1),(93,49,0,0,1,2,0,0,3,350.0000,'2018-07-01 00:00:00','2018-07-02 00:00:00',6,4,0,'',11,0,'','2018-07-01 00:07:32',-1,'2018-07-01 00:07:32',-1),(94,82,0,0,1,2,0,0,3,350.0000,'2018-07-01 00:00:00','2018-07-02 00:00:00',6,4,0,'',11,0,'','2018-07-01 00:07:32',-1,'2018-07-01 00:07:32',-1),(95,86,0,0,1,2,0,0,3,4000.0000,'2018-07-01 00:00:00','2018-07-02 00:00:00',6,4,0,'',27,0,'','2018-07-01 00:07:32',-1,'2018-07-01 00:07:32',-1),(96,49,0,0,1,2,0,0,3,350.0000,'2018-08-01 00:00:00','2018-08-02 00:00:00',6,4,0,'',11,0,'','2018-08-01 21:23:21',-1,'2018-08-01 21:23:21',-1),(97,82,0,0,1,2,0,0,3,350.0000,'2018-08-01 00:00:00','2018-08-02 00:00:00',6,4,0,'',11,0,'','2018-08-01 21:23:21',-1,'2018-08-01 21:23:21',-1),(98,86,0,0,1,2,0,0,3,4000.0000,'2018-08-01 00:00:00','2018-08-02 00:00:00',6,4,0,'',27,0,'','2018-08-01 21:23:21',-1,'2018-08-01 21:23:21',-1),(99,49,0,0,1,2,0,0,3,350.0000,'2018-09-01 00:00:00','2018-09-02 00:00:00',6,4,0,'',11,0,'','2018-09-25 18:38:59',-1,'2018-09-25 18:38:59',-1),(100,82,0,0,1,2,0,0,3,350.0000,'2018-09-01 00:00:00','2018-09-02 00:00:00',6,4,0,'',11,0,'','2018-09-25 18:38:59',-1,'2018-09-25 18:38:59',-1),(101,86,0,0,1,2,0,0,3,4000.0000,'2018-09-01 00:00:00','2018-09-02 00:00:00',6,4,0,'',27,0,'','2018-09-25 18:38:59',-1,'2018-09-25 18:38:59',-1),(102,49,0,0,1,2,0,0,3,350.0000,'2018-10-01 00:00:00','2018-10-01 00:00:00',6,4,0,'',11,0,'','2018-10-01 18:39:17',-1,'2018-10-01 18:39:17',-1),(103,82,0,0,1,2,0,0,3,350.0000,'2018-10-01 00:00:00','2018-10-01 00:00:00',6,4,0,'',11,0,'','2018-10-01 18:39:17',-1,'2018-10-01 18:39:17',-1),(104,86,0,0,1,2,0,0,3,4000.0000,'2018-10-01 00:00:00','2018-10-01 00:00:00',6,4,0,'',27,0,'','2018-10-01 18:39:17',-1,'2018-10-01 18:39:17',-1),(105,49,0,0,1,2,0,0,3,350.0000,'2018-11-01 00:00:00','2018-11-01 00:00:00',6,4,0,'',11,0,'','2018-11-01 18:40:56',-1,'2018-11-01 18:40:56',-1),(106,82,0,0,1,2,0,0,3,350.0000,'2018-11-01 00:00:00','2018-11-01 00:00:00',6,4,0,'',11,0,'','2018-11-01 18:40:56',-1,'2018-11-01 18:40:56',-1),(107,86,0,0,1,2,0,0,3,4000.0000,'2018-11-01 00:00:00','2018-11-01 00:00:00',6,4,0,'',27,0,'','2018-11-01 18:40:56',-1,'2018-11-01 18:40:56',-1),(108,49,0,0,1,2,0,0,3,327.4200,'2018-12-01 00:00:00','2018-12-01 00:00:00',6,4,0,'',11,0,'prorated for 29 of 31 days','2018-12-01 18:42:24',-1,'2018-12-01 18:42:24',-1),(109,82,0,0,1,2,0,0,3,327.4200,'2018-12-01 00:00:00','2018-12-01 00:00:00',6,4,0,'',11,0,'prorated for 29 of 31 days','2018-12-01 18:42:24',-1,'2018-12-01 18:42:24',-1),(110,86,0,0,1,2,0,0,3,3741.9400,'2018-12-01 00:00:00','2018-12-01 00:00:00',6,4,0,'',27,0,'prorated for 29 of 31 days','2018-12-01 18:42:24',-1,'2018-12-01 18:42:24',-1);
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
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
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
INSERT INTO `Business` VALUES (1,'REX','JGM First, LLC',6,4,4,0,1,'2018-02-28 08:57:03',0,'2017-11-10 23:24:22',0);
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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
  `CARID` bigint(20) NOT NULL AUTO_INCREMENT,
  `ElementType` bigint(20) NOT NULL,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `ID` bigint(20) NOT NULL,
  `CID` bigint(20) NOT NULL,
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`CARID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Expense`
--

LOCK TABLES `Expense` WRITE;
/*!40000 ALTER TABLE `Expense` DISABLE KEYS */;
INSERT INTO `Expense` VALUES (1,0,1,2,3,15.0000,'2016-10-01 00:00:00','',6,0,'','2018-02-27 20:16:16',200,'2018-02-27 20:16:16',200);
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Flow`
--

LOCK TABLES `Flow` WRITE;
/*!40000 ALTER TABLE `Flow` DISABLE KEYS */;
INSERT INTO `Flow` VALUES (1,1,'R56WK33TGS9E97MUJ449','RA',3,'{\"tie\": {\"people\": [{\"BID\": 1, \"PRID\": 2, \"TMPTCID\": 1}]}, \"meta\": {\"RAID\": 3, \"RAFLAGS\": 0, \"HavePets\": false, \"ActiveUID\": 0, \"Approver1\": 0, \"Approver2\": 0, \"MoveInUID\": 0, \"ActiveDate\": \"1970-01-01 00:00:00 UTC\", \"ActiveName\": \"UID-0\", \"LastTMPVID\": 0, \"MoveInDate\": \"1970-01-01 00:00:00 UTC\", \"MoveInName\": \"UID-0\", \"LastTMPTCID\": 1, \"DocumentDate\": \"1970-01-01 00:00:00 UTC\", \"HaveVehicles\": false, \"LastTMPASMID\": 18, \"LastTMPPETID\": 0, \"Approver1Name\": \"UID-0\", \"Approver2Name\": \"UID-0\", \"DecisionDate1\": \"1970-01-01 00:00:00 UTC\", \"DecisionDate2\": \"1970-01-01 00:00:00 UTC\", \"TerminatorUID\": 0, \"DeclineReason1\": 0, \"DeclineReason2\": 0, \"TerminatorName\": \"UID-0\", \"NoticeToMoveUID\": 0, \"TerminationDate\": \"1970-01-01 00:00:00 UTC\", \"NoticeToMoveDate\": \"1970-01-01 00:00:00 UTC\", \"NoticeToMoveName\": \"UID-0\", \"ApplicationReadyUID\": 0, \"ApplicationReadyDate\": \"1970-01-01 00:00:00 UTC\", \"ApplicationReadyName\": \"UID-0\", \"NoticeToMoveReported\": \"1970-01-01 00:00:00 UTC\", \"LeaseTerminationReason\": 0}, \"pets\": [], \"dates\": {\"BID\": 1, \"CSAgent\": 0, \"RentStop\": \"12/31/2018\", \"RentStart\": \"10/1/2016\", \"AgreementStop\": \"12/31/2018\", \"AgreementStart\": \"10/1/2016\", \"PossessionStop\": \"12/31/2018\", \"PossessionStart\": \"10/1/2016\"}, \"people\": [{\"BID\": 1, \"City\": \"\", \"TCID\": 3, \"State\": \"\", \"Points\": 0, \"Address\": \"\", \"Comment\": \"\", \"Country\": \"\", \"Evicted\": false, \"TMPTCID\": 1, \"Website\": \"\", \"Address2\": \"\", \"Industry\": 0, \"IsRenter\": true, \"LastName\": \"Vahabzadeh\", \"CellPhone\": \"\", \"Convicted\": false, \"FirstName\": \"Alex\", \"IsCompany\": true, \"WorkPhone\": \"\", \"Bankruptcy\": false, \"EvictedDes\": \"\", \"IsOccupant\": true, \"MiddleName\": \"\", \"Occupation\": \"\", \"PostalCode\": \"\", \"TaxpayorID\": \"\", \"CompanyCity\": \"\", \"CompanyName\": \"Beaumont Partners LP\", \"CreditLimit\": 0, \"DateofBirth\": \"1/1/1900\", \"GrossIncome\": 0, \"IsGuarantor\": false, \"SourceSLSID\": 0, \"CompanyEmail\": \"\", \"CompanyPhone\": \"\", \"CompanyState\": \"\", \"ConvictedDes\": \"\", \"PrimaryEmail\": \"\", \"PriorAddress\": \"\", \"SpecialNeeds\": \"\", \"BankruptcyDes\": \"\", \"PreferredName\": \"\", \"CompanyAddress\": \"\", \"CurrentAddress\": \"\", \"DriversLicense\": \"\", \"SecondaryEmail\": \"\", \"OtherPreferences\": \"\", \"ThirdPartySource\": \"0\", \"CompanyPostalCode\": \"\", \"PriorLandLordName\": \"\", \"EligibleFutureUser\": true, \"CurrentLandLordName\": \"\", \"EligibleFuturePayor\": true, \"EmergencyContactName\": \"\", \"PriorLandLordPhoneNo\": \"\", \"PriorReasonForMoving\": 0, \"AlternateEmailAddress\": \"\", \"EmergencyContactEmail\": \"\", \"CurrentLandLordPhoneNo\": \"\", \"CurrentReasonForMoving\": 0, \"PriorLengthOfResidency\": \"\", \"EmergencyContactAddress\": \"\", \"CurrentLengthOfResidency\": \"\", \"EmergencyContactTelephone\": \"\"}], \"vehicles\": [], \"rentables\": [{\"BID\": 1, \"RID\": 2, \"Fees\": [{\"ARID\": 27, \"Stop\": \"1/1/2018\", \"ASMID\": 15, \"Start\": \"1/1/2017\", \"ARName\": \"Gross Scheduled Rent\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 1, \"RentCycle\": 6, \"TransOccTax\": 0, \"ContractAmount\": 4000, \"AtSigningPreTax\": 0}, {\"ARID\": 12, \"Stop\": \"1/31/2017\", \"ASMID\": 42, \"Start\": \"1/31/2017\", \"ARName\": \"Electric Overage\", \"Comment\": \"utilities reimbursement\", \"SalesTax\": 0, \"TMPASMID\": 2, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": 628.45, \"AtSigningPreTax\": 0}, {\"ARID\": 12, \"Stop\": \"2/28/2017\", \"ASMID\": 43, \"Start\": \"2/28/2017\", \"ARName\": \"Electric Overage\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 3, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": 175, \"AtSigningPreTax\": 0}, {\"ARID\": 12, \"Stop\": \"3/31/2017\", \"ASMID\": 44, \"Start\": \"3/31/2017\", \"ARName\": \"Electric Overage\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 4, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": 175, \"AtSigningPreTax\": 0}, {\"ARID\": 12, \"Stop\": \"4/15/2017\", \"ASMID\": 45, \"Start\": \"4/15/2017\", \"ARName\": \"Electric Overage\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 5, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": 81.79, \"AtSigningPreTax\": 0}, {\"ARID\": 12, \"Stop\": \"10/31/2017\", \"ASMID\": 46, \"Start\": \"10/31/2017\", \"ARName\": \"Electric Overage\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 6, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": 409.28, \"AtSigningPreTax\": 0}, {\"ARID\": 27, \"Stop\": \"11/1/2016\", \"ASMID\": 47, \"Start\": \"11/1/2016\", \"ARName\": \"Gross Scheduled Rent\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 7, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": 4000, \"AtSigningPreTax\": 0}, {\"ARID\": 27, \"Stop\": \"12/1/2016\", \"ASMID\": 48, \"Start\": \"12/1/2016\", \"ARName\": \"Gross Scheduled Rent\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 8, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": 4000, \"AtSigningPreTax\": 0}, {\"ARID\": 11, \"Stop\": \"12/31/2018\", \"ASMID\": 49, \"Start\": \"4/1/2017\", \"ARName\": \"Electric Base Fee\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 9, \"RentCycle\": 6, \"TransOccTax\": 0, \"ContractAmount\": 350, \"AtSigningPreTax\": 0}, {\"ARID\": 27, \"Stop\": \"10/1/2016\", \"ASMID\": 59, \"Start\": \"10/1/2016\", \"ARName\": \"Gross Scheduled Rent\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 10, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": 4000, \"AtSigningPreTax\": 0}, {\"ARID\": 12, \"Stop\": \"12/5/2017\", \"ASMID\": 60, \"Start\": \"12/5/2017\", \"ARName\": \"Electric Overage\", \"Comment\": \"Reversed by ASM00000061\", \"SalesTax\": 0, \"TMPASMID\": 11, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": 628.45, \"AtSigningPreTax\": 0}, {\"ARID\": 12, \"Stop\": \"12/5/2017\", \"ASMID\": 61, \"Start\": \"12/5/2017\", \"ARName\": \"Electric Overage\", \"Comment\": \"Reversal of ASM00000060\", \"SalesTax\": 0, \"TMPASMID\": 12, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": -628.45, \"AtSigningPreTax\": 0}, {\"ARID\": 27, \"Stop\": \"1/1/2018\", \"ASMID\": 66, \"Start\": \"1/1/2018\", \"ARName\": \"Gross Scheduled Rent\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 13, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": 4000, \"AtSigningPreTax\": 0}, {\"ARID\": 27, \"Stop\": \"2/1/2018\", \"ASMID\": 67, \"Start\": \"2/1/2018\", \"ARName\": \"Gross Scheduled Rent\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 14, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": 4000, \"AtSigningPreTax\": 0}, {\"ARID\": 11, \"Stop\": \"1/1/2018\", \"ASMID\": 68, \"Start\": \"1/1/2018\", \"ARName\": \"Electric Base Fee\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 15, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": 350, \"AtSigningPreTax\": 0}, {\"ARID\": 11, \"Stop\": \"2/1/2018\", \"ASMID\": 69, \"Start\": \"2/1/2018\", \"ARName\": \"Electric Base Fee\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 16, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": 350, \"AtSigningPreTax\": 0}, {\"ARID\": 11, \"Stop\": \"12/31/2018\", \"ASMID\": 82, \"Start\": \"3/1/2018\", \"ARName\": \"Electric Base Fee\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 17, \"RentCycle\": 6, \"TransOccTax\": 0, \"ContractAmount\": 350, \"AtSigningPreTax\": 0}, {\"ARID\": 27, \"Stop\": \"12/31/2018\", \"ASMID\": 86, \"Start\": \"3/1/2018\", \"ARName\": \"Gross Scheduled Rent\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 18, \"RentCycle\": 6, \"TransOccTax\": 0, \"ContractAmount\": 4000, \"AtSigningPreTax\": 0}], \"RTID\": 2, \"RTFLAGS\": 0, \"SalesTax\": 0, \"RentCycle\": 6, \"TransOccTax\": 0, \"RentableName\": \"309 1/2 Rexford\", \"AtSigningPreTax\": 0}], \"parentchild\": []}','2018-08-02 01:15:13',211,'2018-08-02 01:15:13',211);
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
INSERT INTO `GLAccount` VALUES (1,0,1,0,0,'10000','Cash','Cash',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(2,1,1,0,0,'10100','Petty Cash','Cash',1,0,'','2017-11-28 18:32:14',0,'2017-11-10 23:24:22',0),(3,1,1,0,0,'10104','FRB 54320 (operating account)','Bank Account',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(4,1,1,0,0,'10105','FRB 96953 (security deposits)','Bank Account',1,0,'','2017-11-27 21:42:09',0,'2017-11-10 23:24:22',0),(6,1,1,0,0,'10999','Undeposited Funds','Cash',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(7,0,1,0,0,'11000','Credit Cards Funds in Transit','Cash',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(8,0,1,0,0,'12000','Accounts Receivable','Accounts Receivable',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(9,8,1,0,0,'12001','Rent Roll Receivables','Accounts Receivable',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(10,0,1,0,0,'12999','Unapplied Funds','Asset',1,0,'','2018-11-29 22:24:05',198,'2017-11-10 23:24:22',0),(11,0,1,0,0,'30000','Security Deposit Liability','Liability Security Deposit',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(12,0,1,0,0,'30001','Floating Security Deposits','Liability Security Deposit',1,0,'Sec Dep posted before rentable identified','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(13,0,1,0,0,'30100','Collected Taxes','Liabilities',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(14,13,1,0,0,'30101','Sales Taxes Collected','Liabilities',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(15,13,1,0,0,'30102','Transient Occupancy Taxes Collected','Liabilities',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(16,13,1,0,0,'30199','Other Collected Taxes','Liabilities',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(17,0,1,0,0,'41000','Gross Scheduled Rent','Income',1,0,'','2017-11-28 18:38:33',0,'2017-11-10 23:24:22',0),(19,0,1,0,0,'41100','Unit Income Offsets','Income Offset',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(20,19,1,0,0,'41101','Vacancy','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(21,19,1,0,0,'41102','Loss (Gain) to Lease','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(22,19,1,0,0,'41103','Employee Concessions','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(23,19,1,0,0,'41104','Resident Concessions','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(24,19,1,0,0,'41105','Owner Concession','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(25,19,1,0,0,'41106','Administrative Concession','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(26,19,1,0,0,'41107','Off Line Renovations','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(27,19,1,0,0,'41108','Off Line Maintenance','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(28,19,1,0,0,'41199','Othe Income Offsets','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(29,0,1,0,0,'41200','Service Fees','Income',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(30,29,1,0,0,'41201','Broadcast and IT Services','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(31,29,1,0,0,'41202','Food Services','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(32,29,1,0,0,'41203','Linen Services','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(33,29,1,0,0,'41204','Housekeeping Revenue','Income',1,0,'','2018-11-29 22:27:18',198,'2017-11-10 23:24:22',0),(34,29,1,0,0,'41299','Other Service Fees','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(35,0,1,0,0,'41300','Utility Fees','Income',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(36,35,1,0,0,'41301','Electric Base Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(37,35,1,0,0,'41302','Electric Overage','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(38,35,1,0,0,'41303','Water and Sewer Base Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(39,35,1,0,0,'41304','Water and Sewer Overage','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(40,35,1,0,0,'41305','Gas Base Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(41,35,1,0,0,'41306','Gas Overage','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(42,35,1,0,0,'41307','Trash Collection Base Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(43,35,1,0,0,'41308','Trash Collection Overage','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(44,35,1,0,0,'41399','Other Utility Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(45,0,1,0,0,'41400','Special Tenant Charges','Income',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(46,45,1,0,0,'41401','Application Fees','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(47,45,1,0,0,'41402','Late Fees','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(48,45,1,0,0,'41403','Insufficient Funds Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(49,45,1,0,0,'41404','Month to Month Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(50,45,1,0,0,'41405','Rentable Specialties','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(51,45,1,0,0,'41406','No Show or Termination Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(52,45,1,0,0,'41407','Pet Fees','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(53,45,1,0,0,'41408','Pet Rent','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(54,45,1,0,0,'41409','Tenant Expense Chargeback','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(55,45,1,0,0,'41410','Special Cleaning Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(56,45,1,0,0,'41411','Eviction Fee Reimbursement','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(57,45,1,0,0,'41412','Extra Person Charge','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(58,45,1,0,0,'41413','Security Deposit Forfeiture','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(59,45,1,0,0,'41414','Damage Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(60,45,1,0,0,'41415','CAM Fees','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(61,45,1,0,0,'41499','Other Special Tenant Charges','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(62,0,1,0,0,'42000','Business Income','Income',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(63,62,1,0,0,'42100','Convenience Store','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(64,62,1,0,0,'42200','Fitness Center Revenue','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(65,62,1,0,0,'42300','Vending Income','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(66,62,1,0,0,'42400','Restaurant Sales','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(67,62,1,0,0,'42500','Bar Sales','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(68,62,1,0,0,'42600','Spa Sales','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(69,0,1,0,0,'50000','Expenses','Expenses',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(70,69,1,0,0,'50001','Cash Over/Short','Expenses',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(71,69,1,0,0,'50002','Bad Debt','Expenses',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(72,69,1,0,0,'50003','Bank Service Fee','Expenses',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(73,69,1,0,0,'50999','Other Expenses','Expenses',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0);
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
  `InvoiceASMID` bigint(20) NOT NULL AUTO_INCREMENT,
  `InvoiceNo` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `ASMID` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`InvoiceASMID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
  `InvoicePayorID` bigint(20) NOT NULL AUTO_INCREMENT,
  `InvoiceNo` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `PID` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`InvoicePayorID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
) ENGINE=InnoDB AUTO_INCREMENT=221 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Journal`
--

LOCK TABLES `Journal` WRITE;
/*!40000 ALTER TABLE `Journal` DISABLE KEYS */;
INSERT INTO `Journal` VALUES (1,1,'2014-03-01 00:00:00',7000.0000,1,1,'','2017-11-30 18:39:27',0,'2017-11-30 18:39:27',0),(2,1,'2016-07-01 00:00:00',8300.0000,1,2,'','2017-11-30 18:41:00',0,'2017-11-30 18:41:00',0),(3,1,'2017-01-01 00:00:00',3750.0000,1,4,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(4,1,'2017-02-01 00:00:00',3750.0000,1,5,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(5,1,'2017-03-01 00:00:00',3750.0000,1,6,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(6,1,'2017-04-01 00:00:00',3750.0000,1,7,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(7,1,'2017-05-01 00:00:00',3750.0000,1,8,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(8,1,'2017-06-01 00:00:00',3750.0000,1,9,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(9,1,'2017-07-01 00:00:00',3750.0000,1,10,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(10,1,'2017-08-01 00:00:00',3750.0000,1,11,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(11,1,'2017-09-01 00:00:00',3750.0000,1,12,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(12,1,'2017-10-01 00:00:00',3750.0000,1,13,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(13,1,'2017-11-01 00:00:00',3750.0000,1,14,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(14,1,'2017-01-01 00:00:00',4000.0000,1,16,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(15,1,'2017-02-01 00:00:00',4000.0000,1,17,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(16,1,'2017-03-01 00:00:00',4000.0000,1,18,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(17,1,'2017-04-01 00:00:00',4000.0000,1,19,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(18,1,'2017-05-01 00:00:00',4000.0000,1,20,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(19,1,'2017-06-01 00:00:00',4000.0000,1,21,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(20,1,'2017-07-01 00:00:00',4000.0000,1,22,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(21,1,'2017-08-01 00:00:00',4000.0000,1,23,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(22,1,'2017-09-01 00:00:00',4000.0000,1,24,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(23,1,'2017-10-01 00:00:00',4000.0000,1,25,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(24,1,'2017-11-01 00:00:00',4000.0000,1,26,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(25,1,'2017-01-01 00:00:00',4150.0000,1,28,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(26,1,'2017-02-01 00:00:00',4150.0000,1,29,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(27,1,'2017-03-01 00:00:00',4150.0000,1,30,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(28,1,'2017-04-01 00:00:00',4150.0000,1,31,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(29,1,'2017-05-01 00:00:00',4150.0000,1,32,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(30,1,'2017-06-01 00:00:00',4150.0000,1,33,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(31,1,'2017-07-01 00:00:00',4150.0000,1,34,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(32,1,'2017-08-01 00:00:00',4150.0000,1,35,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(33,1,'2017-09-01 00:00:00',4150.0000,1,36,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(34,1,'2017-10-01 00:00:00',4150.0000,1,37,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(35,1,'2017-11-01 00:00:00',4150.0000,1,38,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(36,1,'2014-03-01 00:00:00',7000.0000,2,1,'','2017-11-30 18:48:02',0,'2017-11-30 18:48:02',0),(37,1,'2016-07-01 00:00:00',8300.0000,2,2,'','2017-11-30 18:49:17',0,'2017-11-30 18:49:17',0),(38,1,'2016-07-01 00:00:00',8300.0000,2,3,'','2017-11-30 19:13:23',0,'2017-11-30 19:13:23',0),(39,1,'2016-07-01 00:00:00',-8300.0000,2,4,'','2017-11-30 19:13:59',0,'2017-11-30 19:13:59',0),(40,1,'2016-07-01 00:00:00',8300.0000,4,3,'auto-transfer for deposit DEP-1','2017-11-30 19:17:32',0,'2017-11-30 19:17:32',0),(41,1,'2014-03-01 00:00:00',-7000.0000,2,5,'','2017-11-30 19:23:47',0,'2017-11-30 19:23:47',0),(42,1,'2014-03-01 00:00:00',7000.0000,2,6,'','2017-11-30 19:24:32',0,'2017-11-30 19:24:32',0),(43,1,'2014-03-01 00:00:00',7000.0000,4,6,'auto-transfer for deposit DEP-1','2017-11-30 19:25:23',0,'2017-11-30 19:25:23',0),(44,1,'2017-01-01 00:00:00',3750.0000,2,7,'','2017-11-30 19:44:52',0,'2017-11-30 19:44:52',0),(45,1,'2014-03-01 00:00:00',7000.0000,2,6,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(46,1,'2017-01-01 00:00:00',3750.0000,2,7,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(47,1,'2017-12-01 00:00:00',3750.0000,1,39,'','2017-12-01 00:00:04',0,'2017-12-01 00:00:04',0),(48,1,'2017-12-01 00:00:00',4000.0000,1,40,'','2017-12-01 00:00:04',0,'2017-12-01 00:00:04',0),(49,1,'2017-12-01 00:00:00',4150.0000,1,41,'','2017-12-01 00:00:04',0,'2017-12-01 00:00:04',0),(50,1,'2017-01-31 00:00:00',628.4500,1,42,'','2017-12-05 16:01:46',0,'2017-12-05 16:01:46',0),(51,1,'2017-02-28 00:00:00',175.0000,1,43,'','2017-12-05 16:02:25',0,'2017-12-05 16:02:25',0),(52,1,'2017-03-31 00:00:00',175.0000,1,44,'','2017-12-05 16:03:13',0,'2017-12-05 16:03:13',0),(53,1,'2017-04-15 00:00:00',81.7900,1,45,'','2017-12-05 16:03:41',0,'2017-12-05 16:03:41',0),(54,1,'2017-10-31 00:00:00',409.2800,1,46,'','2017-12-05 16:07:34',0,'2017-12-05 16:07:34',0),(55,1,'2017-01-01 00:00:00',3750.0000,2,8,'','2017-12-05 16:09:37',0,'2017-12-05 16:09:37',0),(56,1,'2017-01-01 00:00:00',4150.0000,2,9,'','2017-12-05 16:10:06',0,'2017-12-05 16:10:06',0),(57,1,'2017-02-01 00:00:00',8350.0000,2,10,'','2017-12-05 16:12:02',0,'2017-12-05 16:12:02',0),(58,1,'2017-02-01 00:00:00',3750.0000,2,11,'','2017-12-05 16:12:32',0,'2017-12-05 16:12:32',0),(59,1,'2017-02-01 00:00:00',4150.0000,2,12,'','2017-12-05 16:13:44',0,'2017-12-05 16:13:44',0),(60,1,'2016-11-01 00:00:00',4000.0000,1,47,'','2017-12-05 16:15:10',0,'2017-12-05 16:15:10',0),(61,1,'2016-12-01 00:00:00',4000.0000,1,48,'','2017-12-05 16:15:45',0,'2017-12-05 16:15:45',0),(62,1,'2016-11-15 00:00:00',12000.0000,2,13,'','2017-12-05 16:16:37',0,'2017-12-05 16:16:37',0),(63,1,'2017-01-01 00:00:00',-3750.0000,2,14,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(64,1,'2017-12-05 16:19:04',-3750.0000,2,14,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(65,1,'2017-03-01 00:00:00',3750.0000,2,15,'','2017-12-05 16:20:28',0,'2017-12-05 16:20:28',0),(66,1,'2017-03-01 00:00:00',4150.0000,2,16,'','2017-12-05 16:22:08',0,'2017-12-05 16:22:08',0),(67,1,'2017-04-01 00:00:00',3750.0000,2,17,'','2017-12-05 16:22:50',0,'2017-12-05 16:22:50',0),(68,1,'2017-04-01 00:00:00',4150.0000,2,18,'','2017-12-05 16:23:12',0,'2017-12-05 16:23:12',0),(69,1,'2017-05-01 00:00:00',3750.0000,2,19,'','2017-12-05 16:24:00',0,'2017-12-05 16:24:00',0),(70,1,'2017-05-01 00:00:00',4150.0000,2,20,'','2017-12-05 16:24:18',0,'2017-12-05 16:24:18',0),(71,1,'2017-05-15 00:00:00',13131.7900,2,21,'','2017-12-05 16:26:21',0,'2017-12-05 16:26:21',0),(72,1,'2017-06-01 00:00:00',3750.0000,2,22,'','2017-12-05 16:27:03',0,'2017-12-05 16:27:03',0),(73,1,'2017-06-01 00:00:00',4150.0000,2,23,'','2017-12-05 16:27:16',0,'2017-12-05 16:27:16',0),(74,1,'2017-07-01 00:00:00',3750.0000,2,24,'','2017-12-05 16:27:58',0,'2017-12-05 16:27:58',0),(75,1,'2017-07-01 00:00:00',4150.0000,2,25,'','2017-12-05 16:28:12',0,'2017-12-05 16:28:12',0),(76,1,'2017-08-01 00:00:00',3750.0000,2,26,'','2017-12-05 16:29:16',0,'2017-12-05 16:29:16',0),(77,1,'2017-08-01 00:00:00',4150.0000,2,27,'','2017-12-05 16:29:33',0,'2017-12-05 16:29:33',0),(78,1,'2017-08-15 00:00:00',13050.0000,2,28,'','2017-12-05 16:29:59',0,'2017-12-05 16:29:59',0),(79,1,'2017-09-01 00:00:00',3750.0000,2,29,'','2017-12-05 16:30:33',0,'2017-12-05 16:30:33',0),(80,1,'2017-09-01 00:00:00',4150.0000,2,30,'','2017-12-05 16:30:51',0,'2017-12-05 16:30:51',0),(81,1,'2017-10-01 00:00:00',3750.0000,2,31,'','2017-12-05 16:31:42',0,'2017-12-05 16:31:42',0),(82,1,'2017-10-01 00:00:00',4150.0000,2,32,'','2017-12-05 16:31:56',0,'2017-12-05 16:31:56',0),(83,1,'2017-11-01 00:00:00',3750.0000,2,33,'','2017-12-05 16:32:48',0,'2017-12-05 16:32:48',0),(84,1,'2017-11-01 00:00:00',4150.0000,2,34,'','2017-12-05 16:33:11',0,'2017-12-05 16:33:11',0),(85,1,'2017-11-15 00:00:00',13459.2800,2,35,'','2017-12-05 16:40:59',0,'2017-12-05 16:40:59',0),(86,1,'2017-12-01 00:00:00',3750.0000,2,36,'','2017-12-05 16:42:24',0,'2017-12-05 16:42:24',0),(87,1,'2017-12-01 00:00:00',4150.0000,2,37,'','2017-12-05 16:42:35',0,'2017-12-05 16:42:35',0),(88,1,'2017-01-01 00:00:00',3750.0000,2,8,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(89,1,'2017-02-01 00:00:00',3750.0000,2,11,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(90,1,'2017-03-01 00:00:00',3750.0000,2,15,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(91,1,'2017-04-01 00:00:00',3750.0000,2,17,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(92,1,'2017-05-01 00:00:00',3750.0000,2,19,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(93,1,'2017-06-01 00:00:00',3750.0000,2,22,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(94,1,'2017-07-01 00:00:00',3750.0000,2,24,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(95,1,'2017-08-01 00:00:00',3750.0000,2,26,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(96,1,'2017-09-01 00:00:00',3750.0000,2,29,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(97,1,'2017-10-01 00:00:00',3750.0000,2,31,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(98,1,'2017-11-01 00:00:00',3750.0000,2,33,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(99,1,'2017-12-01 00:00:00',3750.0000,2,36,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(100,1,'2016-07-01 00:00:00',8300.0000,2,3,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(101,1,'2017-01-01 00:00:00',4150.0000,2,9,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(102,1,'2017-02-01 00:00:00',4150.0000,2,12,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(103,1,'2017-03-01 00:00:00',4150.0000,2,16,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(104,1,'2017-04-01 00:00:00',4150.0000,2,18,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(105,1,'2017-05-01 00:00:00',4150.0000,2,20,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(106,1,'2017-06-01 00:00:00',4150.0000,2,23,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(107,1,'2017-07-01 00:00:00',4150.0000,2,25,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(108,1,'2017-08-01 00:00:00',4150.0000,2,27,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(109,1,'2017-09-01 00:00:00',4150.0000,2,30,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(110,1,'2017-10-01 00:00:00',4150.0000,2,32,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(111,1,'2017-11-01 00:00:00',4150.0000,2,34,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(112,1,'2017-12-01 00:00:00',4150.0000,2,37,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(113,1,'2017-04-01 00:00:00',350.0000,1,50,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(114,1,'2017-05-01 00:00:00',350.0000,1,51,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(115,1,'2017-06-01 00:00:00',350.0000,1,52,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(116,1,'2017-07-01 00:00:00',350.0000,1,53,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(117,1,'2017-08-01 00:00:00',350.0000,1,54,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(118,1,'2017-09-01 00:00:00',350.0000,1,55,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(119,1,'2017-10-01 00:00:00',350.0000,1,56,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(120,1,'2017-11-01 00:00:00',350.0000,1,57,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(121,1,'2017-12-01 00:00:00',350.0000,1,58,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(122,1,'2016-10-01 00:00:00',4000.0000,1,59,'','2017-12-05 18:03:15',0,'2017-12-05 18:03:15',0),(123,1,'2017-12-05 00:00:00',628.4500,1,60,'','2017-12-05 18:23:25',0,'2017-12-05 18:23:25',0),(124,1,'2017-12-05 00:00:00',-628.4500,1,61,'','2017-12-05 19:41:01',0,'2017-12-05 19:41:01',0),(125,1,'2017-02-03 00:00:00',628.4500,2,38,'','2017-12-05 19:44:51',0,'2017-12-05 19:44:51',0),(126,1,'2018-02-01 00:00:00',3750.0000,1,62,'','2018-02-20 19:54:46',-99,'2018-02-20 19:54:46',-99),(127,1,'2018-02-01 00:00:00',4150.0000,1,63,'','2018-02-21 00:00:09',-99,'2018-02-21 00:00:09',-99),(128,1,'2018-02-01 00:00:00',-387.5000,1,64,'','2018-02-22 00:00:08',-99,'2018-02-22 00:00:08',-99),(129,1,'2016-10-01 00:00:00',15.0000,3,1,'','2018-02-27 20:16:16',200,'2018-02-27 20:16:16',200),(130,1,'2018-02-27 00:00:00',4000.0000,2,13,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(131,1,'2018-02-27 00:00:00',4000.0000,2,13,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(132,1,'2018-02-27 00:00:00',4000.0000,2,13,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(133,1,'2018-02-27 00:00:00',4000.0000,2,10,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(134,1,'2018-02-27 00:00:00',628.4500,2,10,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(135,1,'2018-02-27 00:00:00',3721.5500,2,10,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(136,1,'2018-02-27 00:00:00',278.4500,2,38,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(137,1,'2018-02-27 00:00:00',175.0000,2,38,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(138,1,'2018-02-27 00:00:00',175.0000,2,38,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(139,1,'2018-02-27 00:00:00',3825.0000,2,21,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(140,1,'2018-02-27 00:00:00',175.0000,2,21,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(141,1,'2018-02-27 00:00:00',4000.0000,2,21,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(142,1,'2018-02-27 00:00:00',350.0000,2,21,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(143,1,'2018-02-27 00:00:00',81.7900,2,21,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(144,1,'2018-02-27 00:00:00',4000.0000,2,21,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(145,1,'2018-02-27 00:00:00',350.0000,2,21,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(146,1,'2018-02-27 00:00:00',350.0000,2,21,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(147,1,'2018-02-27 00:00:00',3650.0000,2,28,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(148,1,'2018-02-27 00:00:00',350.0000,2,28,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(149,1,'2018-02-27 00:00:00',4000.0000,2,28,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(150,1,'2018-02-27 00:00:00',350.0000,2,28,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(151,1,'2018-02-27 00:00:00',4000.0000,2,28,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(152,1,'2018-02-27 00:00:00',350.0000,2,28,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(153,1,'2018-02-27 00:00:00',350.0000,2,28,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(154,1,'2018-02-27 00:00:00',3650.0000,2,35,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(155,1,'2018-02-27 00:00:00',350.0000,2,35,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(156,1,'2018-02-27 00:00:00',4000.0000,2,35,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(157,1,'2018-02-27 00:00:00',350.0000,2,35,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(158,1,'2018-02-27 00:00:00',409.2800,2,35,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(159,1,'2018-02-27 00:00:00',4000.0000,2,35,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(160,1,'2018-02-27 00:00:00',350.0000,2,35,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(161,1,'2018-02-27 00:00:00',350.0000,2,35,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(162,1,'2016-10-03 00:00:00',4000.0000,2,39,'','2018-02-28 17:18:33',211,'2018-02-28 17:18:33',211),(163,1,'2017-12-01 00:00:00',3650.0000,2,39,'','2018-02-28 17:19:56',211,'2018-02-28 17:19:56',211),(164,1,'2017-12-01 00:00:00',350.0000,2,39,'','2018-02-28 17:19:56',211,'2018-02-28 17:19:56',211),(165,1,'2018-01-01 00:00:00',4150.0000,1,65,'','2018-02-28 17:39:47',211,'2018-02-28 17:39:47',211),(166,1,'2018-01-01 00:00:00',4000.0000,1,66,'','2018-02-28 17:41:00',211,'2018-02-28 17:41:00',211),(167,1,'2018-02-01 00:00:00',4000.0000,1,67,'','2018-02-28 17:41:40',211,'2018-02-28 17:41:40',211),(168,1,'2018-01-01 00:00:00',350.0000,1,68,'','2018-02-28 21:12:57',200,'2018-02-28 21:12:57',200),(169,1,'2018-02-01 00:00:00',350.0000,1,69,'','2018-02-28 21:13:21',200,'2018-02-28 21:13:21',200),(170,1,'2018-01-01 00:00:00',3750.0000,2,40,'','2018-02-28 21:16:01',200,'2018-02-28 21:16:01',200),(171,1,'2018-01-01 00:00:00',4150.0000,2,41,'','2018-02-28 21:16:23',200,'2018-02-28 21:16:23',200),(172,1,'2018-02-01 00:00:00',4150.0000,2,42,'','2018-02-28 21:16:42',200,'2018-02-28 21:16:42',200),(173,1,'2018-02-23 00:00:00',13050.0000,2,43,'','2018-02-28 21:17:07',200,'2018-02-28 21:17:07',200),(174,1,'2018-02-28 00:00:00',3750.0000,2,40,'','2018-02-28 21:17:34',200,'2018-02-28 21:17:34',200),(175,1,'2018-02-28 00:00:00',4000.0000,2,43,'','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(176,1,'2018-02-28 00:00:00',350.0000,2,43,'','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(177,1,'2018-02-28 00:00:00',4000.0000,2,43,'','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(178,1,'2018-02-28 00:00:00',350.0000,2,43,'','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(179,1,'2018-02-28 00:00:00',4150.0000,2,41,'','2018-02-28 21:17:57',200,'2018-02-28 21:17:57',200),(180,1,'2018-02-28 00:00:00',4150.0000,2,42,'','2018-02-28 21:17:57',200,'2018-02-28 21:17:57',200),(181,1,'2018-03-01 00:00:00',4150.0000,1,70,'','2018-04-03 17:19:38',0,'2018-04-03 17:19:38',0),(182,1,'2018-03-01 00:00:00',350.0000,1,71,'','2018-04-03 17:19:38',0,'2018-04-03 17:19:38',0),(183,1,'2018-04-01 00:00:00',4150.0000,1,72,'','2018-04-03 17:19:48',0,'2018-04-03 17:19:48',0),(184,1,'2018-04-01 00:00:00',350.0000,1,73,'','2018-04-03 17:19:48',0,'2018-04-03 17:19:48',0),(185,1,'2018-05-01 00:00:00',4150.0000,1,74,'','2018-05-02 20:49:41',-1,'2018-05-02 20:49:41',-1),(186,1,'2018-05-01 00:00:00',350.0000,1,75,'','2018-05-02 20:49:41',-1,'2018-05-02 20:49:41',-1),(187,1,'2018-03-01 00:00:00',350.0000,1,83,'','2018-05-30 19:45:28',200,'2018-05-30 19:45:28',200),(188,1,'2018-04-01 00:00:00',350.0000,1,84,'','2018-05-30 19:45:29',200,'2018-05-30 19:45:29',200),(189,1,'2018-05-01 00:00:00',350.0000,1,85,'','2018-05-30 19:45:29',200,'2018-05-30 19:45:29',200),(190,1,'2018-03-01 00:00:00',4000.0000,1,87,'','2018-05-30 19:50:19',200,'2018-05-30 19:50:19',200),(191,1,'2018-04-01 00:00:00',4000.0000,1,88,'','2018-05-30 19:50:19',200,'2018-05-30 19:50:19',200),(192,1,'2018-05-01 00:00:00',4000.0000,1,89,'','2018-05-30 19:50:19',200,'2018-05-30 19:50:19',200),(193,1,'2018-03-01 00:00:00',350.0000,2,43,'','2018-05-30 20:04:19',200,'2018-05-30 20:04:19',200),(194,1,'2018-03-01 00:00:00',4000.0000,2,43,'','2018-05-30 20:04:19',200,'2018-05-30 20:04:19',200),(195,1,'2018-05-20 00:00:00',13050.0000,2,44,'','2018-05-30 20:08:28',200,'2018-05-30 20:08:28',200),(196,1,'2018-05-20 00:00:00',350.0000,2,44,'','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(197,1,'2018-05-20 00:00:00',4000.0000,2,44,'','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(198,1,'2018-05-20 00:00:00',350.0000,2,44,'','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(199,1,'2018-05-20 00:00:00',4000.0000,2,44,'','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(200,1,'2018-06-01 00:00:00',350.0000,1,90,'','2018-06-01 00:03:21',-1,'2018-06-01 00:03:21',-1),(201,1,'2018-06-01 00:00:00',350.0000,1,91,'','2018-06-01 00:03:21',-1,'2018-06-01 00:03:21',-1),(202,1,'2018-06-01 00:00:00',4000.0000,1,92,'','2018-06-01 00:03:21',-1,'2018-06-01 00:03:21',-1),(203,1,'2018-07-01 00:00:00',350.0000,1,93,'','2018-07-01 00:07:32',-1,'2018-07-01 00:07:32',-1),(204,1,'2018-07-01 00:00:00',350.0000,1,94,'','2018-07-01 00:07:32',-1,'2018-07-01 00:07:32',-1),(205,1,'2018-07-01 00:00:00',4000.0000,1,95,'','2018-07-01 00:07:32',-1,'2018-07-01 00:07:32',-1),(206,1,'2018-08-01 00:00:00',350.0000,1,96,'','2018-08-01 21:23:21',-1,'2018-08-01 21:23:21',-1),(207,1,'2018-08-01 00:00:00',350.0000,1,97,'','2018-08-01 21:23:21',-1,'2018-08-01 21:23:21',-1),(208,1,'2018-08-01 00:00:00',4000.0000,1,98,'','2018-08-01 21:23:21',-1,'2018-08-01 21:23:21',-1),(209,1,'2018-09-01 00:00:00',350.0000,1,99,'','2018-09-25 18:38:59',-1,'2018-09-25 18:38:59',-1),(210,1,'2018-09-01 00:00:00',350.0000,1,100,'','2018-09-25 18:38:59',-1,'2018-09-25 18:38:59',-1),(211,1,'2018-09-01 00:00:00',4000.0000,1,101,'','2018-09-25 18:38:59',-1,'2018-09-25 18:38:59',-1),(212,1,'2018-10-01 00:00:00',350.0000,1,102,'','2018-10-01 18:39:17',-1,'2018-10-01 18:39:17',-1),(213,1,'2018-10-01 00:00:00',350.0000,1,103,'','2018-10-01 18:39:17',-1,'2018-10-01 18:39:17',-1),(214,1,'2018-10-01 00:00:00',4000.0000,1,104,'','2018-10-01 18:39:17',-1,'2018-10-01 18:39:17',-1),(215,1,'2018-11-01 00:00:00',350.0000,1,105,'','2018-11-01 18:40:56',-1,'2018-11-01 18:40:56',-1),(216,1,'2018-11-01 00:00:00',350.0000,1,106,'','2018-11-01 18:40:56',-1,'2018-11-01 18:40:56',-1),(217,1,'2018-11-01 00:00:00',4000.0000,1,107,'','2018-11-01 18:40:56',-1,'2018-11-01 18:40:56',-1),(218,1,'2018-12-01 00:00:00',327.4200,1,108,'','2018-12-01 18:42:24',-1,'2018-12-01 18:42:24',-1),(219,1,'2018-12-01 00:00:00',327.4200,1,109,'','2018-12-01 18:42:24',-1,'2018-12-01 18:42:24',-1),(220,1,'2018-12-01 00:00:00',3741.9400,1,110,'','2018-12-01 18:42:24',-1,'2018-12-01 18:42:24',-1);
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
) ENGINE=InnoDB AUTO_INCREMENT=218 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `JournalAllocation`
--

LOCK TABLES `JournalAllocation` WRITE;
/*!40000 ALTER TABLE `JournalAllocation` DISABLE KEYS */;
INSERT INTO `JournalAllocation` VALUES (1,1,1,1,2,0,0,7000.0000,1,0,'d 12001 7000.00, c 30000 7000.00','2017-11-30 18:39:27',0,'2018-02-23 08:47:27',0),(2,1,2,3,4,0,0,8300.0000,2,0,'d 12001 8300.00, c 30000 8300.00','2017-11-30 18:41:00',0,'2018-02-23 08:47:27',0),(3,1,3,1,2,0,0,3750.0000,4,0,'d 12001 3750.00, c 41000 3750.00','2017-11-30 18:43:20',0,'2018-02-23 08:47:27',0),(4,1,4,1,2,0,0,3750.0000,5,0,'d 12001 3750.00, c 41000 3750.00','2017-11-30 18:43:20',0,'2018-02-23 08:47:27',0),(5,1,5,1,2,0,0,3750.0000,6,0,'d 12001 3750.00, c 41000 3750.00','2017-11-30 18:43:20',0,'2018-02-23 08:47:27',0),(6,1,6,1,2,0,0,3750.0000,7,0,'d 12001 3750.00, c 41000 3750.00','2017-11-30 18:43:20',0,'2018-02-23 08:47:27',0),(7,1,7,1,2,0,0,3750.0000,8,0,'d 12001 3750.00, c 41000 3750.00','2017-11-30 18:43:20',0,'2018-02-23 08:47:27',0),(8,1,8,1,2,0,0,3750.0000,9,0,'d 12001 3750.00, c 41000 3750.00','2017-11-30 18:43:20',0,'2018-02-23 08:47:27',0),(9,1,9,1,2,0,0,3750.0000,10,0,'d 12001 3750.00, c 41000 3750.00','2017-11-30 18:43:20',0,'2018-02-23 08:47:27',0),(10,1,10,1,2,0,0,3750.0000,11,0,'d 12001 3750.00, c 41000 3750.00','2017-11-30 18:43:20',0,'2018-02-23 08:47:27',0),(11,1,11,1,2,0,0,3750.0000,12,0,'d 12001 3750.00, c 41000 3750.00','2017-11-30 18:43:20',0,'2018-02-23 08:47:27',0),(12,1,12,1,2,0,0,3750.0000,13,0,'d 12001 3750.00, c 41000 3750.00','2017-11-30 18:43:20',0,'2018-02-23 08:47:27',0),(13,1,13,1,2,0,0,3750.0000,14,0,'d 12001 3750.00, c 41000 3750.00','2017-11-30 18:43:20',0,'2018-02-23 08:47:27',0),(14,1,14,2,3,0,0,4000.0000,16,0,'d 12001 4000.00, c 41000 4000.00','2017-11-30 18:45:17',0,'2018-02-23 08:47:27',0),(15,1,15,2,3,0,0,4000.0000,17,0,'d 12001 4000.00, c 41000 4000.00','2017-11-30 18:45:17',0,'2018-02-23 08:47:27',0),(16,1,16,2,3,0,0,4000.0000,18,0,'d 12001 4000.00, c 41000 4000.00','2017-11-30 18:45:17',0,'2018-02-23 08:47:27',0),(17,1,17,2,3,0,0,4000.0000,19,0,'d 12001 4000.00, c 41000 4000.00','2017-11-30 18:45:17',0,'2018-02-23 08:47:27',0),(18,1,18,2,3,0,0,4000.0000,20,0,'d 12001 4000.00, c 41000 4000.00','2017-11-30 18:45:17',0,'2018-02-23 08:47:27',0),(19,1,19,2,3,0,0,4000.0000,21,0,'d 12001 4000.00, c 41000 4000.00','2017-11-30 18:45:17',0,'2018-02-23 08:47:27',0),(20,1,20,2,3,0,0,4000.0000,22,0,'d 12001 4000.00, c 41000 4000.00','2017-11-30 18:45:17',0,'2018-02-23 08:47:27',0),(21,1,21,2,3,0,0,4000.0000,23,0,'d 12001 4000.00, c 41000 4000.00','2017-11-30 18:45:17',0,'2018-02-23 08:47:27',0),(22,1,22,2,3,0,0,4000.0000,24,0,'d 12001 4000.00, c 41000 4000.00','2017-11-30 18:45:17',0,'2018-02-23 08:47:27',0),(23,1,23,2,3,0,0,4000.0000,25,0,'d 12001 4000.00, c 41000 4000.00','2017-11-30 18:45:17',0,'2018-02-23 08:47:27',0),(24,1,24,2,3,0,0,4000.0000,26,0,'d 12001 4000.00, c 41000 4000.00','2017-11-30 18:45:17',0,'2018-02-23 08:47:27',0),(25,1,25,3,4,0,0,4150.0000,28,0,'d 12001 4150.00, c 41000 4150.00','2017-11-30 18:45:55',0,'2018-02-23 08:47:27',0),(26,1,26,3,4,0,0,4150.0000,29,0,'d 12001 4150.00, c 41000 4150.00','2017-11-30 18:45:55',0,'2018-02-23 08:47:27',0),(27,1,27,3,4,0,0,4150.0000,30,0,'d 12001 4150.00, c 41000 4150.00','2017-11-30 18:45:55',0,'2018-02-23 08:47:27',0),(28,1,28,3,4,0,0,4150.0000,31,0,'d 12001 4150.00, c 41000 4150.00','2017-11-30 18:45:55',0,'2018-02-23 08:47:27',0),(29,1,29,3,4,0,0,4150.0000,32,0,'d 12001 4150.00, c 41000 4150.00','2017-11-30 18:45:55',0,'2018-02-23 08:47:27',0),(30,1,30,3,4,0,0,4150.0000,33,0,'d 12001 4150.00, c 41000 4150.00','2017-11-30 18:45:55',0,'2018-02-23 08:47:27',0),(31,1,31,3,4,0,0,4150.0000,34,0,'d 12001 4150.00, c 41000 4150.00','2017-11-30 18:45:55',0,'2018-02-23 08:47:27',0),(32,1,32,3,4,0,0,4150.0000,35,0,'d 12001 4150.00, c 41000 4150.00','2017-11-30 18:45:55',0,'2018-02-23 08:47:27',0),(33,1,33,3,4,0,0,4150.0000,36,0,'d 12001 4150.00, c 41000 4150.00','2017-11-30 18:45:55',0,'2018-02-23 08:47:27',0),(34,1,34,3,4,0,0,4150.0000,37,0,'d 12001 4150.00, c 41000 4150.00','2017-11-30 18:45:55',0,'2018-02-23 08:47:27',0),(35,1,35,3,4,0,0,4150.0000,38,0,'d 12001 4150.00, c 41000 4150.00','2017-11-30 18:45:55',0,'2018-02-23 08:47:27',0),(36,1,36,0,0,1,0,7000.0000,0,0,'d 10104 _, c 10999 _','2017-11-30 18:48:02',0,'2018-02-23 08:47:27',0),(37,1,37,0,0,4,0,8300.0000,0,0,'d 10104 _, c 10999 _','2017-11-30 18:49:17',0,'2018-02-23 08:47:27',0),(38,1,38,0,0,4,0,8300.0000,0,0,'d 10999 _, c 12999 _','2017-11-30 19:13:23',0,'2018-02-23 08:47:27',0),(39,1,39,0,0,4,0,-8300.0000,0,0,'d 10104 _, c 10999 _','2017-11-30 19:13:59',0,'2018-02-23 08:47:27',0),(40,1,40,0,0,4,3,8300.0000,0,0,'d 10104 8300.0000, c 10999 8300.0000','2017-11-30 19:17:32',0,'2018-02-23 08:47:27',0),(41,1,41,0,0,1,0,-7000.0000,0,0,'d 10104 _, c 10999 _','2017-11-30 19:23:47',0,'2018-02-23 08:47:27',0),(42,1,42,0,0,1,0,7000.0000,0,0,'d 10999 _, c 12999 _','2017-11-30 19:24:32',0,'2018-02-23 08:47:27',0),(43,1,43,0,0,1,6,7000.0000,0,0,'d 10104 7000.0000, c 10999 7000.0000','2017-11-30 19:25:23',0,'2018-02-23 08:47:27',0),(44,1,44,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2017-11-30 19:44:52',0,'2018-02-23 08:47:27',0),(45,1,45,1,2,1,6,7000.0000,1,0,'ASM(1) d 12999 7000.00,c 12001 7000.00','2017-11-30 19:46:56',0,'2018-02-23 08:47:27',0),(46,1,46,1,2,1,7,3750.0000,4,0,'ASM(4) d 12999 3750.00,c 12001 3750.00','2017-11-30 19:46:56',0,'2018-02-23 08:47:27',0),(47,1,47,1,2,0,0,3750.0000,39,0,'d 12001 3750.00, c 41000 3750.00','2017-12-01 00:00:04',0,'2018-02-23 08:47:27',0),(48,1,48,2,3,0,0,4000.0000,40,0,'d 12001 4000.00, c 41000 4000.00','2017-12-01 00:00:04',0,'2018-02-23 08:47:27',0),(49,1,49,3,4,0,0,4150.0000,41,0,'d 12001 4150.00, c 41000 4150.00','2017-12-01 00:00:04',0,'2018-02-23 08:47:27',0),(50,1,50,2,3,0,0,628.4500,42,0,'d 12001 628.45, c 41302 628.45','2017-12-05 16:01:46',0,'2018-02-23 08:47:27',0),(51,1,51,2,3,0,0,175.0000,43,0,'d 12001 175.00, c 41302 175.00','2017-12-05 16:02:25',0,'2018-02-23 08:47:27',0),(52,1,52,2,3,0,0,175.0000,44,0,'d 12001 175.00, c 41302 175.00','2017-12-05 16:03:13',0,'2018-02-23 08:47:27',0),(53,1,53,2,3,0,0,81.7900,45,0,'d 12001 81.79, c 41302 81.79','2017-12-05 16:03:41',0,'2018-02-23 08:47:27',0),(54,1,54,2,3,0,0,409.2800,46,0,'d 12001 409.28, c 41302 409.28','2017-12-05 16:07:34',0,'2018-02-23 08:47:27',0),(55,1,55,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:09:37',0,'2018-02-23 08:47:27',0),(56,1,56,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:10:06',0,'2018-02-23 08:47:27',0),(57,1,57,0,0,3,0,8350.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:12:02',0,'2018-02-23 08:47:27',0),(58,1,58,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:12:32',0,'2018-02-23 08:47:27',0),(59,1,59,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:13:44',0,'2018-02-23 08:47:27',0),(60,1,60,2,3,0,0,4000.0000,47,0,'d 12001 4000.00, c 41000 4000.00','2017-12-05 16:15:10',0,'2018-02-23 08:47:27',0),(61,1,61,2,3,0,0,4000.0000,48,0,'d 12001 4000.00, c 41000 4000.00','2017-12-05 16:15:45',0,'2018-02-23 08:47:27',0),(62,1,62,0,0,3,0,12000.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:16:37',0,'2018-02-23 08:47:27',0),(63,1,63,0,0,1,0,-3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:19:04',0,'2018-02-23 08:47:27',0),(64,1,64,1,2,1,14,-3750.0000,4,0,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2017-12-05 16:19:04',0,'2018-02-23 08:47:27',0),(65,1,65,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:20:28',0,'2018-02-23 08:47:27',0),(66,1,66,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:22:08',0,'2018-02-23 08:47:27',0),(67,1,67,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:22:50',0,'2018-02-23 08:47:27',0),(68,1,68,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:23:12',0,'2018-02-23 08:47:27',0),(69,1,69,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:24:00',0,'2018-02-23 08:47:27',0),(70,1,70,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:24:18',0,'2018-02-23 08:47:27',0),(71,1,71,0,0,3,0,13131.7900,0,0,'d 10999 _, c 12999 _','2017-12-05 16:26:21',0,'2018-02-23 08:47:27',0),(72,1,72,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:27:03',0,'2018-02-23 08:47:27',0),(73,1,73,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:27:16',0,'2018-02-23 08:47:27',0),(74,1,74,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:27:58',0,'2018-02-23 08:47:27',0),(75,1,75,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:28:12',0,'2018-02-23 08:47:27',0),(76,1,76,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:29:16',0,'2018-02-23 08:47:27',0),(77,1,77,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:29:33',0,'2018-02-23 08:47:27',0),(78,1,78,0,0,3,0,13050.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:29:59',0,'2018-02-23 08:47:27',0),(79,1,79,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:30:33',0,'2018-02-23 08:47:27',0),(80,1,80,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:30:51',0,'2018-02-23 08:47:27',0),(81,1,81,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:31:42',0,'2018-02-23 08:47:27',0),(82,1,82,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:31:56',0,'2018-02-23 08:47:27',0),(83,1,83,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:32:49',0,'2018-02-23 08:47:27',0),(84,1,84,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:33:11',0,'2018-02-23 08:47:27',0),(85,1,85,0,0,3,0,13459.2800,0,0,'d 10999 _, c 12999 _','2017-12-05 16:40:59',0,'2018-02-23 08:47:27',0),(86,1,86,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:42:24',0,'2018-02-23 08:47:27',0),(87,1,87,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:42:35',0,'2018-02-23 08:47:27',0),(88,1,88,1,2,1,8,3750.0000,4,0,'ASM(4) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:31',0,'2018-02-23 08:47:27',0),(89,1,89,1,2,1,11,3750.0000,5,0,'ASM(5) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:31',0,'2018-02-23 08:47:27',0),(90,1,90,1,2,1,15,3750.0000,6,0,'ASM(6) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:31',0,'2018-02-23 08:47:27',0),(91,1,91,1,2,1,17,3750.0000,7,0,'ASM(7) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:31',0,'2018-02-23 08:47:27',0),(92,1,92,1,2,1,19,3750.0000,8,0,'ASM(8) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2018-02-23 08:47:27',0),(93,1,93,1,2,1,22,3750.0000,9,0,'ASM(9) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2018-02-23 08:47:27',0),(94,1,94,1,2,1,24,3750.0000,10,0,'ASM(10) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2018-02-23 08:47:27',0),(95,1,95,1,2,1,26,3750.0000,11,0,'ASM(11) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2018-02-23 08:47:27',0),(96,1,96,1,2,1,29,3750.0000,12,0,'ASM(12) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2018-02-23 08:47:27',0),(97,1,97,1,2,1,31,3750.0000,13,0,'ASM(13) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2018-02-23 08:47:27',0),(98,1,98,1,2,1,33,3750.0000,14,0,'ASM(14) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2018-02-23 08:47:27',0),(99,1,99,1,2,1,36,3750.0000,39,0,'ASM(39) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2018-02-23 08:47:27',0),(100,1,100,3,4,4,3,8300.0000,2,0,'ASM(2) d 12999 8300.00,c 12001 8300.00','2017-12-05 17:06:25',0,'2018-02-23 08:47:27',0),(101,1,101,3,4,4,9,4150.0000,28,0,'ASM(28) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2018-02-23 08:47:27',0),(102,1,102,3,4,4,12,4150.0000,29,0,'ASM(29) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2018-02-23 08:47:27',0),(103,1,103,3,4,4,16,4150.0000,30,0,'ASM(30) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2018-02-23 08:47:27',0),(104,1,104,3,4,4,18,4150.0000,31,0,'ASM(31) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2018-02-23 08:47:27',0),(105,1,105,3,4,4,20,4150.0000,32,0,'ASM(32) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2018-02-23 08:47:27',0),(106,1,106,3,4,4,23,4150.0000,33,0,'ASM(33) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2018-02-23 08:47:27',0),(107,1,107,3,4,4,25,4150.0000,34,0,'ASM(34) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2018-02-23 08:47:27',0),(108,1,108,3,4,4,27,4150.0000,35,0,'ASM(35) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2018-02-23 08:47:27',0),(109,1,109,3,4,4,30,4150.0000,36,0,'ASM(36) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2018-02-23 08:47:27',0),(110,1,110,3,4,4,32,4150.0000,37,0,'ASM(37) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2018-02-23 08:47:27',0),(111,1,111,3,4,4,34,4150.0000,38,0,'ASM(38) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2018-02-23 08:47:27',0),(112,1,112,3,4,4,37,4150.0000,41,0,'ASM(41) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2018-02-23 08:47:27',0),(113,1,113,2,3,0,0,350.0000,50,0,'d 12001 350.00, c 41301 350.00','2017-12-05 17:49:46',0,'2018-02-23 08:47:27',0),(114,1,114,2,3,0,0,350.0000,51,0,'d 12001 350.00, c 41301 350.00','2017-12-05 17:49:46',0,'2018-02-23 08:47:27',0),(115,1,115,2,3,0,0,350.0000,52,0,'d 12001 350.00, c 41301 350.00','2017-12-05 17:49:46',0,'2018-02-23 08:47:27',0),(116,1,116,2,3,0,0,350.0000,53,0,'d 12001 350.00, c 41301 350.00','2017-12-05 17:49:46',0,'2018-02-23 08:47:27',0),(117,1,117,2,3,0,0,350.0000,54,0,'d 12001 350.00, c 41301 350.00','2017-12-05 17:49:46',0,'2018-02-23 08:47:27',0),(118,1,118,2,3,0,0,350.0000,55,0,'d 12001 350.00, c 41301 350.00','2017-12-05 17:49:46',0,'2018-02-23 08:47:27',0),(119,1,119,2,3,0,0,350.0000,56,0,'d 12001 350.00, c 41301 350.00','2017-12-05 17:49:46',0,'2018-02-23 08:47:27',0),(120,1,120,2,3,0,0,350.0000,57,0,'d 12001 350.00, c 41301 350.00','2017-12-05 17:49:46',0,'2018-02-23 08:47:27',0),(121,1,121,2,3,0,0,350.0000,58,0,'d 12001 350.00, c 41301 350.00','2017-12-05 17:49:46',0,'2018-02-23 08:47:27',0),(122,1,122,2,3,0,0,4000.0000,59,0,'d 12001 4000.00, c 41000 4000.00','2017-12-05 18:03:15',0,'2018-02-23 08:47:27',0),(123,1,123,2,3,0,0,628.4500,60,0,'d 12001 628.45, c 41302 628.45','2017-12-05 18:23:25',0,'2018-02-23 08:47:27',0),(124,1,124,2,3,0,0,-628.4500,61,0,'d 12001 -628.45, c 41302 -628.45','2017-12-05 19:41:01',0,'2018-02-23 08:47:27',0),(125,1,125,0,0,3,0,628.4500,0,0,'d 10999 _, c 12999 _','2017-12-05 19:44:51',0,'2018-02-23 08:47:27',0),(126,1,129,2,3,0,0,15.0000,0,1,'d 50003 15.00, c 10104 15.00','2018-02-27 20:16:16',200,'2018-02-27 20:16:16',200),(127,1,130,2,3,3,13,4000.0000,59,0,'ASM(59) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(128,1,131,2,3,3,13,4000.0000,47,0,'ASM(47) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(129,1,132,2,3,3,13,4000.0000,48,0,'ASM(48) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(130,1,133,2,3,3,10,4000.0000,16,0,'ASM(16) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(131,1,134,2,3,3,10,628.4500,42,0,'ASM(42) d 12999 628.45,c 12001 628.45','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(132,1,135,2,3,3,10,3721.5500,17,0,'ASM(17) d 12999 3721.55,c 12001 3721.55','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(133,1,136,2,3,3,38,278.4500,17,0,'ASM(17) d 12999 278.45,c 12001 278.45','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(134,1,137,2,3,3,38,175.0000,43,0,'ASM(43) d 12999 175.00,c 12001 175.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(135,1,138,2,3,3,38,175.0000,18,0,'ASM(18) d 12999 175.00,c 12001 175.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(136,1,139,2,3,3,21,3825.0000,18,0,'ASM(18) d 12999 3825.00,c 12001 3825.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(137,1,140,2,3,3,21,175.0000,44,0,'ASM(44) d 12999 175.00,c 12001 175.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(138,1,141,2,3,3,21,4000.0000,19,0,'ASM(19) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(139,1,142,2,3,3,21,350.0000,50,0,'ASM(50) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(140,1,143,2,3,3,21,81.7900,45,0,'ASM(45) d 12999 81.79,c 12001 81.79','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(141,1,144,2,3,3,21,4000.0000,20,0,'ASM(20) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(142,1,145,2,3,3,21,350.0000,51,0,'ASM(51) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(143,1,146,2,3,3,21,350.0000,21,0,'ASM(21) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(144,1,147,2,3,3,28,3650.0000,21,0,'ASM(21) d 12999 3650.00,c 12001 3650.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(145,1,148,2,3,3,28,350.0000,52,0,'ASM(52) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(146,1,149,2,3,3,28,4000.0000,22,0,'ASM(22) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(147,1,150,2,3,3,28,350.0000,53,0,'ASM(53) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(148,1,151,2,3,3,28,4000.0000,23,0,'ASM(23) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(149,1,152,2,3,3,28,350.0000,54,0,'ASM(54) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(150,1,153,2,3,3,28,350.0000,24,0,'ASM(24) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(151,1,154,2,3,3,35,3650.0000,24,0,'ASM(24) d 12999 3650.00,c 12001 3650.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(152,1,155,2,3,3,35,350.0000,55,0,'ASM(55) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(153,1,156,2,3,3,35,4000.0000,25,0,'ASM(25) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(154,1,157,2,3,3,35,350.0000,56,0,'ASM(56) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(155,1,158,2,3,3,35,409.2800,46,0,'ASM(46) d 12999 409.28,c 12001 409.28','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(156,1,159,2,3,3,35,4000.0000,26,0,'ASM(26) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(157,1,160,2,3,3,35,350.0000,57,0,'ASM(57) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(158,1,161,2,3,3,35,350.0000,40,0,'ASM(40) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(159,1,162,0,0,3,0,4000.0000,0,0,'d 10999 _, c 12999 _','2018-02-28 17:18:33',211,'2018-02-28 17:18:33',211),(160,1,163,2,3,3,39,3650.0000,40,0,'ASM(40) d 12999 3650.00,c 12001 3650.00','2018-02-28 17:19:56',211,'2018-02-28 17:19:56',211),(161,1,164,2,3,3,39,350.0000,58,0,'ASM(58) d 12999 350.00,c 12001 350.00','2018-02-28 17:19:56',211,'2018-02-28 17:19:56',211),(162,1,165,3,4,0,0,4150.0000,65,0,'d 12001 4150.00, c 41000 4150.00','2018-02-28 17:39:47',211,'2018-02-28 17:39:47',211),(163,1,166,2,3,0,0,4000.0000,66,0,'d 12001 4000.00, c 41000 4000.00','2018-02-28 17:41:00',211,'2018-02-28 17:41:00',211),(164,1,167,2,3,0,0,4000.0000,67,0,'d 12001 4000.00, c 41000 4000.00','2018-02-28 17:41:40',211,'2018-02-28 17:41:40',211),(165,1,168,2,3,0,0,350.0000,68,0,'d 12001 350.00, c 41301 350.00','2018-02-28 21:12:57',200,'2018-02-28 21:12:57',200),(166,1,169,2,3,0,0,350.0000,69,0,'d 12001 350.00, c 41301 350.00','2018-02-28 21:13:21',200,'2018-02-28 21:13:21',200),(167,1,170,0,0,1,0,3750.0000,0,0,'d 10999 _, c 12999 _','2018-02-28 21:16:01',200,'2018-02-28 21:16:01',200),(168,1,171,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2018-02-28 21:16:23',200,'2018-02-28 21:16:23',200),(169,1,172,0,0,4,0,4150.0000,0,0,'d 10999 _, c 12999 _','2018-02-28 21:16:42',200,'2018-02-28 21:16:42',200),(170,1,173,0,0,3,0,13050.0000,0,0,'d 10999 _, c 12999 _','2018-02-28 21:17:07',200,'2018-02-28 21:17:07',200),(171,1,174,1,2,1,40,3750.0000,62,0,'ASM(62) d 12999 3750.00,c 12001 3750.00','2018-02-28 21:17:35',200,'2018-02-28 21:17:35',200),(172,1,175,2,3,3,43,4000.0000,66,0,'ASM(66) d 12999 4000.00,c 12001 4000.00','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(173,1,176,2,3,3,43,350.0000,68,0,'ASM(68) d 12999 350.00,c 12001 350.00','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(174,1,177,2,3,3,43,4000.0000,67,0,'ASM(67) d 12999 4000.00,c 12001 4000.00','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(175,1,178,2,3,3,43,350.0000,69,0,'ASM(69) d 12999 350.00,c 12001 350.00','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(176,1,179,3,4,4,41,4150.0000,65,0,'ASM(65) d 12999 4150.00,c 12001 4150.00','2018-02-28 21:17:57',200,'2018-02-28 21:17:57',200),(177,1,180,3,4,4,42,4150.0000,63,0,'ASM(63) d 12999 4150.00,c 12001 4150.00','2018-02-28 21:17:57',200,'2018-02-28 21:17:57',200),(178,1,181,3,4,0,0,4150.0000,70,0,'d 12001 4150.00, c 41000 4150.00','2018-04-03 17:19:38',0,'2018-04-03 17:19:38',0),(179,1,182,2,3,0,0,350.0000,71,0,'d 12001 350.00, c 41301 350.00','2018-04-03 17:19:38',0,'2018-04-03 17:19:38',0),(180,1,183,3,4,0,0,4150.0000,72,0,'d 12001 4150.00, c 41000 4150.00','2018-04-03 17:19:48',0,'2018-04-03 17:19:48',0),(181,1,184,2,3,0,0,350.0000,73,0,'d 12001 350.00, c 41301 350.00','2018-04-03 17:19:48',0,'2018-04-03 17:19:48',0),(182,1,185,3,4,0,0,4150.0000,74,0,'d 12001 4150.00, c 41000 4150.00','2018-05-02 20:49:41',-1,'2018-05-02 20:49:41',-1),(183,1,186,2,3,0,0,350.0000,75,0,'d 12001 350.00, c 41301 350.00','2018-05-02 20:49:41',-1,'2018-05-02 20:49:41',-1),(184,1,187,2,3,0,0,350.0000,83,0,'d 12001 350.00, c 41301 350.00','2018-05-30 19:45:29',200,'2018-05-30 19:45:29',200),(185,1,188,2,3,0,0,350.0000,84,0,'d 12001 350.00, c 41301 350.00','2018-05-30 19:45:29',200,'2018-05-30 19:45:29',200),(186,1,189,2,3,0,0,350.0000,85,0,'d 12001 350.00, c 41301 350.00','2018-05-30 19:45:29',200,'2018-05-30 19:45:29',200),(187,1,190,2,3,0,0,4000.0000,87,0,'d 12001 4000.00, c 41000 4000.00','2018-05-30 19:50:19',200,'2018-05-30 19:50:19',200),(188,1,191,2,3,0,0,4000.0000,88,0,'d 12001 4000.00, c 41000 4000.00','2018-05-30 19:50:19',200,'2018-05-30 19:50:19',200),(189,1,192,2,3,0,0,4000.0000,89,0,'d 12001 4000.00, c 41000 4000.00','2018-05-30 19:50:19',200,'2018-05-30 19:50:19',200),(190,1,193,2,3,3,43,350.0000,83,0,'ASM(83) d 12999 350.00,c 12001 350.00','2018-05-30 20:04:19',200,'2018-05-30 20:04:19',200),(191,1,194,2,3,3,43,4000.0000,87,0,'ASM(87) d 12999 4000.00,c 12001 4000.00','2018-05-30 20:04:19',200,'2018-05-30 20:04:19',200),(192,1,195,0,0,3,0,13050.0000,0,0,'d 10999 _, c 12999 _','2018-05-30 20:08:28',200,'2018-05-30 20:08:28',200),(193,1,196,2,3,3,44,350.0000,84,0,'ASM(84) d 12999 350.00,c 12001 350.00','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(194,1,197,2,3,3,44,4000.0000,88,0,'ASM(88) d 12999 4000.00,c 12001 4000.00','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(195,1,198,2,3,3,44,350.0000,85,0,'ASM(85) d 12999 350.00,c 12001 350.00','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(196,1,199,2,3,3,44,4000.0000,89,0,'ASM(89) d 12999 4000.00,c 12001 4000.00','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(197,1,200,2,3,0,0,350.0000,90,0,'d 12001 350.00, c 41301 350.00','2018-06-01 00:03:21',-1,'2018-06-01 00:03:21',-1),(198,1,201,2,3,0,0,350.0000,91,0,'d 12001 350.00, c 41301 350.00','2018-06-01 00:03:21',-1,'2018-06-01 00:03:21',-1),(199,1,202,2,3,0,0,4000.0000,92,0,'d 12001 4000.00, c 41000 4000.00','2018-06-01 00:03:21',-1,'2018-06-01 00:03:21',-1),(200,1,203,2,3,0,0,350.0000,93,0,'d 12001 350.00, c 41301 350.00','2018-07-01 00:07:32',-1,'2018-07-01 00:07:32',-1),(201,1,204,2,3,0,0,350.0000,94,0,'d 12001 350.00, c 41301 350.00','2018-07-01 00:07:32',-1,'2018-07-01 00:07:32',-1),(202,1,205,2,3,0,0,4000.0000,95,0,'d 12001 4000.00, c 41000 4000.00','2018-07-01 00:07:32',-1,'2018-07-01 00:07:32',-1),(203,1,206,2,3,0,0,350.0000,96,0,'d 12001 350.00, c 41301 350.00','2018-08-01 21:23:21',-1,'2018-08-01 21:23:21',-1),(204,1,207,2,3,0,0,350.0000,97,0,'d 12001 350.00, c 41301 350.00','2018-08-01 21:23:21',-1,'2018-08-01 21:23:21',-1),(205,1,208,2,3,0,0,4000.0000,98,0,'d 12001 4000.00, c 41000 4000.00','2018-08-01 21:23:21',-1,'2018-08-01 21:23:21',-1),(206,1,209,2,3,0,0,350.0000,99,0,'d 12001 350.00, c 41301 350.00','2018-09-25 18:38:59',-1,'2018-09-25 18:38:59',-1),(207,1,210,2,3,0,0,350.0000,100,0,'d 12001 350.00, c 41301 350.00','2018-09-25 18:38:59',-1,'2018-09-25 18:38:59',-1),(208,1,211,2,3,0,0,4000.0000,101,0,'d 12001 4000.00, c 41000 4000.00','2018-09-25 18:38:59',-1,'2018-09-25 18:38:59',-1),(209,1,212,2,3,0,0,350.0000,102,0,'d 12001 350.00, c 41301 350.00','2018-10-01 18:39:17',-1,'2018-10-01 18:39:17',-1),(210,1,213,2,3,0,0,350.0000,103,0,'d 12001 350.00, c 41301 350.00','2018-10-01 18:39:17',-1,'2018-10-01 18:39:17',-1),(211,1,214,2,3,0,0,4000.0000,104,0,'d 12001 4000.00, c 41000 4000.00','2018-10-01 18:39:17',-1,'2018-10-01 18:39:17',-1),(212,1,215,2,3,0,0,350.0000,105,0,'d 12001 350.00, c 41301 350.00','2018-11-01 18:40:56',-1,'2018-11-01 18:40:56',-1),(213,1,216,2,3,0,0,350.0000,106,0,'d 12001 350.00, c 41301 350.00','2018-11-01 18:40:56',-1,'2018-11-01 18:40:56',-1),(214,1,217,2,3,0,0,4000.0000,107,0,'d 12001 4000.00, c 41000 4000.00','2018-11-01 18:40:56',-1,'2018-11-01 18:40:56',-1),(215,1,218,2,3,0,0,327.4200,108,0,'d 12001 327.42, c 41301 327.42','2018-12-01 18:42:24',-1,'2018-12-01 18:42:24',-1),(216,1,219,2,3,0,0,327.4200,109,0,'d 12001 327.42, c 41301 327.42','2018-12-01 18:42:24',-1,'2018-12-01 18:42:24',-1),(217,1,220,2,3,0,0,3741.9400,110,0,'d 12001 3741.94, c 41000 3741.94','2018-12-01 18:42:24',-1,'2018-12-01 18:42:24',-1);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
) ENGINE=InnoDB AUTO_INCREMENT=375 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerEntry`
--

LOCK TABLES `LedgerEntry` WRITE;
/*!40000 ALTER TABLE `LedgerEntry` DISABLE KEYS */;
INSERT INTO `LedgerEntry` VALUES (1,1,1,1,9,2,1,0,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 18:39:27',0,'2017-11-30 18:39:27',0),(2,1,1,1,11,2,1,0,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 18:39:27',0,'2017-11-30 18:39:27',0),(3,1,2,2,9,4,3,0,'2016-07-01 00:00:00',8300.0000,'','2017-11-30 18:41:00',0,'2017-11-30 18:41:00',0),(4,1,2,2,11,4,3,0,'2016-07-01 00:00:00',-8300.0000,'','2017-11-30 18:41:00',0,'2017-11-30 18:41:00',0),(5,1,3,3,9,2,1,0,'2017-01-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(6,1,3,3,17,2,1,0,'2017-01-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(7,1,4,4,9,2,1,0,'2017-02-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(8,1,4,4,17,2,1,0,'2017-02-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(9,1,5,5,9,2,1,0,'2017-03-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(10,1,5,5,17,2,1,0,'2017-03-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(11,1,6,6,9,2,1,0,'2017-04-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(12,1,6,6,17,2,1,0,'2017-04-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(13,1,7,7,9,2,1,0,'2017-05-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(14,1,7,7,17,2,1,0,'2017-05-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(15,1,8,8,9,2,1,0,'2017-06-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(16,1,8,8,17,2,1,0,'2017-06-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(17,1,9,9,9,2,1,0,'2017-07-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(18,1,9,9,17,2,1,0,'2017-07-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(19,1,10,10,9,2,1,0,'2017-08-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(20,1,10,10,17,2,1,0,'2017-08-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(21,1,11,11,9,2,1,0,'2017-09-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(22,1,11,11,17,2,1,0,'2017-09-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(23,1,12,12,9,2,1,0,'2017-10-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(24,1,12,12,17,2,1,0,'2017-10-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(25,1,13,13,9,2,1,0,'2017-11-01 00:00:00',3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(26,1,13,13,17,2,1,0,'2017-11-01 00:00:00',-3750.0000,'','2017-11-30 18:43:20',0,'2017-11-30 18:43:20',0),(27,1,14,14,9,3,2,0,'2017-01-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(28,1,14,14,17,3,2,0,'2017-01-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(29,1,15,15,9,3,2,0,'2017-02-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(30,1,15,15,17,3,2,0,'2017-02-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(31,1,16,16,9,3,2,0,'2017-03-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(32,1,16,16,17,3,2,0,'2017-03-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(33,1,17,17,9,3,2,0,'2017-04-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(34,1,17,17,17,3,2,0,'2017-04-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(35,1,18,18,9,3,2,0,'2017-05-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(36,1,18,18,17,3,2,0,'2017-05-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(37,1,19,19,9,3,2,0,'2017-06-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(38,1,19,19,17,3,2,0,'2017-06-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(39,1,20,20,9,3,2,0,'2017-07-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(40,1,20,20,17,3,2,0,'2017-07-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(41,1,21,21,9,3,2,0,'2017-08-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(42,1,21,21,17,3,2,0,'2017-08-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(43,1,22,22,9,3,2,0,'2017-09-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(44,1,22,22,17,3,2,0,'2017-09-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(45,1,23,23,9,3,2,0,'2017-10-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(46,1,23,23,17,3,2,0,'2017-10-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(47,1,24,24,9,3,2,0,'2017-11-01 00:00:00',4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(48,1,24,24,17,3,2,0,'2017-11-01 00:00:00',-4000.0000,'','2017-11-30 18:45:17',0,'2017-11-30 18:45:17',0),(49,1,25,25,9,4,3,0,'2017-01-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(50,1,25,25,17,4,3,0,'2017-01-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(51,1,26,26,9,4,3,0,'2017-02-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(52,1,26,26,17,4,3,0,'2017-02-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(53,1,27,27,9,4,3,0,'2017-03-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(54,1,27,27,17,4,3,0,'2017-03-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(55,1,28,28,9,4,3,0,'2017-04-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(56,1,28,28,17,4,3,0,'2017-04-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(57,1,29,29,9,4,3,0,'2017-05-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(58,1,29,29,17,4,3,0,'2017-05-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(59,1,30,30,9,4,3,0,'2017-06-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(60,1,30,30,17,4,3,0,'2017-06-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(61,1,31,31,9,4,3,0,'2017-07-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(62,1,31,31,17,4,3,0,'2017-07-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(63,1,32,32,9,4,3,0,'2017-08-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(64,1,32,32,17,4,3,0,'2017-08-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(65,1,33,33,9,4,3,0,'2017-09-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(66,1,33,33,17,4,3,0,'2017-09-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(67,1,34,34,9,4,3,0,'2017-10-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(68,1,34,34,17,4,3,0,'2017-10-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(69,1,35,35,9,4,3,0,'2017-11-01 00:00:00',4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(70,1,35,35,17,4,3,0,'2017-11-01 00:00:00',-4150.0000,'','2017-11-30 18:45:55',0,'2017-11-30 18:45:55',0),(71,1,36,36,3,0,0,1,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 18:48:02',0,'2017-11-30 18:48:02',0),(72,1,36,36,6,0,0,1,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 18:48:02',0,'2017-11-30 18:48:02',0),(73,1,37,37,3,0,0,4,'2016-07-01 00:00:00',8300.0000,'','2017-11-30 18:49:17',0,'2017-11-30 18:49:17',0),(74,1,37,37,6,0,0,4,'2016-07-01 00:00:00',-8300.0000,'','2017-11-30 18:49:17',0,'2017-11-30 18:49:17',0),(75,1,38,38,6,0,0,4,'2016-07-01 00:00:00',8300.0000,'','2017-11-30 19:13:23',0,'2017-11-30 19:13:23',0),(76,1,38,38,10,0,0,4,'2016-07-01 00:00:00',-8300.0000,'','2017-11-30 19:13:23',0,'2017-11-30 19:13:23',0),(77,1,39,39,3,0,0,4,'2016-07-01 00:00:00',-8300.0000,'','2017-11-30 19:13:59',0,'2017-11-30 19:13:59',0),(78,1,39,39,6,0,0,4,'2016-07-01 00:00:00',8300.0000,'','2017-11-30 19:13:59',0,'2017-11-30 19:13:59',0),(79,1,40,40,3,0,0,4,'2016-07-01 00:00:00',8300.0000,'','2017-11-30 19:17:32',0,'2017-11-30 19:17:32',0),(80,1,40,40,6,0,0,4,'2016-07-01 00:00:00',-8300.0000,'','2017-11-30 19:17:32',0,'2017-11-30 19:17:32',0),(81,1,41,41,3,0,0,1,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 19:23:47',0,'2017-11-30 19:23:47',0),(82,1,41,41,6,0,0,1,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 19:23:47',0,'2017-11-30 19:23:47',0),(83,1,42,42,6,0,0,1,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 19:24:32',0,'2017-11-30 19:24:32',0),(84,1,42,42,10,0,0,1,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 19:24:32',0,'2017-11-30 19:24:32',0),(85,1,43,43,3,0,0,1,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 19:25:23',0,'2017-11-30 19:25:23',0),(86,1,43,43,6,0,0,1,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 19:25:23',0,'2017-11-30 19:25:23',0),(87,1,44,44,6,0,0,1,'2017-01-01 00:00:00',3750.0000,'','2017-11-30 19:44:52',0,'2017-11-30 19:44:52',0),(88,1,44,44,10,0,0,1,'2017-01-01 00:00:00',-3750.0000,'','2017-11-30 19:44:52',0,'2017-11-30 19:44:52',0),(89,1,45,45,10,2,1,1,'2014-03-01 00:00:00',7000.0000,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(90,1,45,45,9,2,1,1,'2014-03-01 00:00:00',-7000.0000,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(91,1,46,46,10,2,1,1,'2017-01-01 00:00:00',3750.0000,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(92,1,46,46,9,2,1,1,'2017-01-01 00:00:00',-3750.0000,'','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(93,1,50,50,9,3,2,0,'2017-01-31 00:00:00',628.4500,'','2017-12-05 16:01:46',0,'2017-12-05 16:01:46',0),(94,1,50,50,37,3,2,0,'2017-01-31 00:00:00',-628.4500,'','2017-12-05 16:01:46',0,'2017-12-05 16:01:46',0),(95,1,51,51,9,3,2,0,'2017-02-28 00:00:00',175.0000,'','2017-12-05 16:02:25',0,'2017-12-05 16:02:25',0),(96,1,51,51,37,3,2,0,'2017-02-28 00:00:00',-175.0000,'','2017-12-05 16:02:25',0,'2017-12-05 16:02:25',0),(97,1,52,52,9,3,2,0,'2017-03-31 00:00:00',175.0000,'','2017-12-05 16:03:13',0,'2017-12-05 16:03:13',0),(98,1,52,52,37,3,2,0,'2017-03-31 00:00:00',-175.0000,'','2017-12-05 16:03:13',0,'2017-12-05 16:03:13',0),(99,1,53,53,9,3,2,0,'2017-04-15 00:00:00',81.7900,'','2017-12-05 16:03:41',0,'2017-12-05 16:03:41',0),(100,1,53,53,37,3,2,0,'2017-04-15 00:00:00',-81.7900,'','2017-12-05 16:03:41',0,'2017-12-05 16:03:41',0),(101,1,54,54,9,3,2,0,'2017-10-31 00:00:00',409.2800,'','2017-12-05 16:07:34',0,'2017-12-05 16:07:34',0),(102,1,54,54,37,3,2,0,'2017-10-31 00:00:00',-409.2800,'','2017-12-05 16:07:34',0,'2017-12-05 16:07:34',0),(103,1,55,55,6,0,0,1,'2017-01-01 00:00:00',3750.0000,'','2017-12-05 16:09:37',0,'2017-12-05 16:09:37',0),(104,1,55,55,10,0,0,1,'2017-01-01 00:00:00',-3750.0000,'','2017-12-05 16:09:37',0,'2017-12-05 16:09:37',0),(105,1,56,56,6,0,0,4,'2017-01-01 00:00:00',4150.0000,'','2017-12-05 16:10:06',0,'2017-12-05 16:10:06',0),(106,1,56,56,10,0,0,4,'2017-01-01 00:00:00',-4150.0000,'','2017-12-05 16:10:06',0,'2017-12-05 16:10:06',0),(107,1,57,57,6,0,0,3,'2017-02-01 00:00:00',8350.0000,'','2017-12-05 16:12:02',0,'2017-12-05 16:12:02',0),(108,1,57,57,10,0,0,3,'2017-02-01 00:00:00',-8350.0000,'','2017-12-05 16:12:02',0,'2017-12-05 16:12:02',0),(109,1,58,58,6,0,0,1,'2017-02-01 00:00:00',3750.0000,'','2017-12-05 16:12:32',0,'2017-12-05 16:12:32',0),(110,1,58,58,10,0,0,1,'2017-02-01 00:00:00',-3750.0000,'','2017-12-05 16:12:32',0,'2017-12-05 16:12:32',0),(111,1,59,59,6,0,0,4,'2017-02-01 00:00:00',4150.0000,'','2017-12-05 16:13:44',0,'2017-12-05 16:13:44',0),(112,1,59,59,10,0,0,4,'2017-02-01 00:00:00',-4150.0000,'','2017-12-05 16:13:44',0,'2017-12-05 16:13:44',0),(113,1,60,60,9,3,2,0,'2016-11-01 00:00:00',4000.0000,'','2017-12-05 16:15:10',0,'2017-12-05 16:15:10',0),(114,1,60,60,17,3,2,0,'2016-11-01 00:00:00',-4000.0000,'','2017-12-05 16:15:10',0,'2017-12-05 16:15:10',0),(115,1,61,61,9,3,2,0,'2016-12-01 00:00:00',4000.0000,'','2017-12-05 16:15:45',0,'2017-12-05 16:15:45',0),(116,1,61,61,17,3,2,0,'2016-12-01 00:00:00',-4000.0000,'','2017-12-05 16:15:45',0,'2017-12-05 16:15:45',0),(117,1,62,62,6,0,0,3,'2016-11-15 00:00:00',12000.0000,'','2017-12-05 16:16:37',0,'2017-12-05 16:16:37',0),(118,1,62,62,10,0,0,3,'2016-11-15 00:00:00',-12000.0000,'','2017-12-05 16:16:37',0,'2017-12-05 16:16:37',0),(119,1,63,63,6,0,0,1,'2017-01-01 00:00:00',-3750.0000,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(120,1,63,63,10,0,0,1,'2017-01-01 00:00:00',3750.0000,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(121,1,64,64,10,2,1,1,'2017-01-01 00:00:00',-3750.0000,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(122,1,64,64,9,2,1,1,'2017-01-01 00:00:00',3750.0000,'','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(123,1,65,65,6,0,0,1,'2017-03-01 00:00:00',3750.0000,'','2017-12-05 16:20:28',0,'2017-12-05 16:20:28',0),(124,1,65,65,10,0,0,1,'2017-03-01 00:00:00',-3750.0000,'','2017-12-05 16:20:28',0,'2017-12-05 16:20:28',0),(125,1,66,66,6,0,0,4,'2017-03-01 00:00:00',4150.0000,'','2017-12-05 16:22:08',0,'2017-12-05 16:22:08',0),(126,1,66,66,10,0,0,4,'2017-03-01 00:00:00',-4150.0000,'','2017-12-05 16:22:08',0,'2017-12-05 16:22:08',0),(127,1,67,67,6,0,0,1,'2017-04-01 00:00:00',3750.0000,'','2017-12-05 16:22:50',0,'2017-12-05 16:22:50',0),(128,1,67,67,10,0,0,1,'2017-04-01 00:00:00',-3750.0000,'','2017-12-05 16:22:50',0,'2017-12-05 16:22:50',0),(129,1,68,68,6,0,0,4,'2017-04-01 00:00:00',4150.0000,'','2017-12-05 16:23:12',0,'2017-12-05 16:23:12',0),(130,1,68,68,10,0,0,4,'2017-04-01 00:00:00',-4150.0000,'','2017-12-05 16:23:12',0,'2017-12-05 16:23:12',0),(131,1,69,69,6,0,0,1,'2017-05-01 00:00:00',3750.0000,'','2017-12-05 16:24:00',0,'2017-12-05 16:24:00',0),(132,1,69,69,10,0,0,1,'2017-05-01 00:00:00',-3750.0000,'','2017-12-05 16:24:01',0,'2017-12-05 16:24:01',0),(133,1,70,70,6,0,0,4,'2017-05-01 00:00:00',4150.0000,'','2017-12-05 16:24:18',0,'2017-12-05 16:24:18',0),(134,1,70,70,10,0,0,4,'2017-05-01 00:00:00',-4150.0000,'','2017-12-05 16:24:18',0,'2017-12-05 16:24:18',0),(135,1,71,71,6,0,0,3,'2017-05-15 00:00:00',13131.7900,'','2017-12-05 16:26:21',0,'2017-12-05 16:26:21',0),(136,1,71,71,10,0,0,3,'2017-05-15 00:00:00',-13131.7900,'','2017-12-05 16:26:21',0,'2017-12-05 16:26:21',0),(137,1,72,72,6,0,0,1,'2017-06-01 00:00:00',3750.0000,'','2017-12-05 16:27:03',0,'2017-12-05 16:27:03',0),(138,1,72,72,10,0,0,1,'2017-06-01 00:00:00',-3750.0000,'','2017-12-05 16:27:03',0,'2017-12-05 16:27:03',0),(139,1,73,73,6,0,0,4,'2017-06-01 00:00:00',4150.0000,'','2017-12-05 16:27:16',0,'2017-12-05 16:27:16',0),(140,1,73,73,10,0,0,4,'2017-06-01 00:00:00',-4150.0000,'','2017-12-05 16:27:16',0,'2017-12-05 16:27:16',0),(141,1,74,74,6,0,0,1,'2017-07-01 00:00:00',3750.0000,'','2017-12-05 16:27:58',0,'2017-12-05 16:27:58',0),(142,1,74,74,10,0,0,1,'2017-07-01 00:00:00',-3750.0000,'','2017-12-05 16:27:58',0,'2017-12-05 16:27:58',0),(143,1,75,75,6,0,0,4,'2017-07-01 00:00:00',4150.0000,'','2017-12-05 16:28:13',0,'2017-12-05 16:28:13',0),(144,1,75,75,10,0,0,4,'2017-07-01 00:00:00',-4150.0000,'','2017-12-05 16:28:13',0,'2017-12-05 16:28:13',0),(145,1,76,76,6,0,0,1,'2017-08-01 00:00:00',3750.0000,'','2017-12-05 16:29:16',0,'2017-12-05 16:29:16',0),(146,1,76,76,10,0,0,1,'2017-08-01 00:00:00',-3750.0000,'','2017-12-05 16:29:16',0,'2017-12-05 16:29:16',0),(147,1,77,77,6,0,0,4,'2017-08-01 00:00:00',4150.0000,'','2017-12-05 16:29:33',0,'2017-12-05 16:29:33',0),(148,1,77,77,10,0,0,4,'2017-08-01 00:00:00',-4150.0000,'','2017-12-05 16:29:33',0,'2017-12-05 16:29:33',0),(149,1,78,78,6,0,0,3,'2017-08-15 00:00:00',13050.0000,'','2017-12-05 16:29:59',0,'2017-12-05 16:29:59',0),(150,1,78,78,10,0,0,3,'2017-08-15 00:00:00',-13050.0000,'','2017-12-05 16:29:59',0,'2017-12-05 16:29:59',0),(151,1,79,79,6,0,0,1,'2017-09-01 00:00:00',3750.0000,'','2017-12-05 16:30:33',0,'2017-12-05 16:30:33',0),(152,1,79,79,10,0,0,1,'2017-09-01 00:00:00',-3750.0000,'','2017-12-05 16:30:33',0,'2017-12-05 16:30:33',0),(153,1,80,80,6,0,0,4,'2017-09-01 00:00:00',4150.0000,'','2017-12-05 16:30:51',0,'2017-12-05 16:30:51',0),(154,1,80,80,10,0,0,4,'2017-09-01 00:00:00',-4150.0000,'','2017-12-05 16:30:51',0,'2017-12-05 16:30:51',0),(155,1,81,81,6,0,0,1,'2017-10-01 00:00:00',3750.0000,'','2017-12-05 16:31:42',0,'2017-12-05 16:31:42',0),(156,1,81,81,10,0,0,1,'2017-10-01 00:00:00',-3750.0000,'','2017-12-05 16:31:42',0,'2017-12-05 16:31:42',0),(157,1,82,82,6,0,0,4,'2017-10-01 00:00:00',4150.0000,'','2017-12-05 16:31:56',0,'2017-12-05 16:31:56',0),(158,1,82,82,10,0,0,4,'2017-10-01 00:00:00',-4150.0000,'','2017-12-05 16:31:56',0,'2017-12-05 16:31:56',0),(159,1,83,83,6,0,0,1,'2017-11-01 00:00:00',3750.0000,'','2017-12-05 16:32:49',0,'2017-12-05 16:32:49',0),(160,1,83,83,10,0,0,1,'2017-11-01 00:00:00',-3750.0000,'','2017-12-05 16:32:49',0,'2017-12-05 16:32:49',0),(161,1,84,84,6,0,0,4,'2017-11-01 00:00:00',4150.0000,'','2017-12-05 16:33:11',0,'2017-12-05 16:33:11',0),(162,1,84,84,10,0,0,4,'2017-11-01 00:00:00',-4150.0000,'','2017-12-05 16:33:11',0,'2017-12-05 16:33:11',0),(163,1,85,85,6,0,0,3,'2017-11-15 00:00:00',13459.2800,'','2017-12-05 16:40:59',0,'2017-12-05 16:40:59',0),(164,1,85,85,10,0,0,3,'2017-11-15 00:00:00',-13459.2800,'','2017-12-05 16:40:59',0,'2017-12-05 16:40:59',0),(165,1,86,86,6,0,0,1,'2017-12-01 00:00:00',3750.0000,'','2017-12-05 16:42:24',0,'2017-12-05 16:42:24',0),(166,1,86,86,10,0,0,1,'2017-12-01 00:00:00',-3750.0000,'','2017-12-05 16:42:24',0,'2017-12-05 16:42:24',0),(167,1,87,87,6,0,0,4,'2017-12-01 00:00:00',4150.0000,'','2017-12-05 16:42:35',0,'2017-12-05 16:42:35',0),(168,1,87,87,10,0,0,4,'2017-12-01 00:00:00',-4150.0000,'','2017-12-05 16:42:35',0,'2017-12-05 16:42:35',0),(169,1,88,88,10,2,1,1,'2017-01-01 00:00:00',3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(170,1,88,88,9,2,1,1,'2017-01-01 00:00:00',-3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(171,1,89,89,10,2,1,1,'2017-02-01 00:00:00',3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(172,1,89,89,9,2,1,1,'2017-02-01 00:00:00',-3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(173,1,90,90,10,2,1,1,'2017-03-01 00:00:00',3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(174,1,90,90,9,2,1,1,'2017-03-01 00:00:00',-3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(175,1,91,91,10,2,1,1,'2017-04-01 00:00:00',3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(176,1,91,91,9,2,1,1,'2017-04-01 00:00:00',-3750.0000,'','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(177,1,92,92,10,2,1,1,'2017-05-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(178,1,92,92,9,2,1,1,'2017-05-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(179,1,93,93,10,2,1,1,'2017-06-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(180,1,93,93,9,2,1,1,'2017-06-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(181,1,94,94,10,2,1,1,'2017-07-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(182,1,94,94,9,2,1,1,'2017-07-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(183,1,95,95,10,2,1,1,'2017-08-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(184,1,95,95,9,2,1,1,'2017-08-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(185,1,96,96,10,2,1,1,'2017-09-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(186,1,96,96,9,2,1,1,'2017-09-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(187,1,97,97,10,2,1,1,'2017-10-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(188,1,97,97,9,2,1,1,'2017-10-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(189,1,98,98,10,2,1,1,'2017-11-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(190,1,98,98,9,2,1,1,'2017-11-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(191,1,99,99,10,2,1,1,'2017-12-01 00:00:00',3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(192,1,99,99,9,2,1,1,'2017-12-01 00:00:00',-3750.0000,'','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(193,1,100,100,10,4,3,4,'2016-07-01 00:00:00',8300.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(194,1,100,100,9,4,3,4,'2016-07-01 00:00:00',-8300.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(195,1,101,101,10,4,3,4,'2017-01-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(196,1,101,101,9,4,3,4,'2017-01-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(197,1,102,102,10,4,3,4,'2017-02-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(198,1,102,102,9,4,3,4,'2017-02-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(199,1,103,103,10,4,3,4,'2017-03-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(200,1,103,103,9,4,3,4,'2017-03-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(201,1,104,104,10,4,3,4,'2017-04-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(202,1,104,104,9,4,3,4,'2017-04-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(203,1,105,105,10,4,3,4,'2017-05-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(204,1,105,105,9,4,3,4,'2017-05-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(205,1,106,106,10,4,3,4,'2017-06-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(206,1,106,106,9,4,3,4,'2017-06-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(207,1,107,107,10,4,3,4,'2017-07-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(208,1,107,107,9,4,3,4,'2017-07-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(209,1,108,108,10,4,3,4,'2017-08-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(210,1,108,108,9,4,3,4,'2017-08-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(211,1,109,109,10,4,3,4,'2017-09-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(212,1,109,109,9,4,3,4,'2017-09-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(213,1,110,110,10,4,3,4,'2017-10-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(214,1,110,110,9,4,3,4,'2017-10-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(215,1,111,111,10,4,3,4,'2017-11-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(216,1,111,111,9,4,3,4,'2017-11-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(217,1,112,112,10,4,3,4,'2017-12-01 00:00:00',4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(218,1,112,112,9,4,3,4,'2017-12-01 00:00:00',-4150.0000,'','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(219,1,113,113,9,3,2,0,'2017-04-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(220,1,113,113,36,3,2,0,'2017-04-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(221,1,114,114,9,3,2,0,'2017-05-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(222,1,114,114,36,3,2,0,'2017-05-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(223,1,115,115,9,3,2,0,'2017-06-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(224,1,115,115,36,3,2,0,'2017-06-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(225,1,116,116,9,3,2,0,'2017-07-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(226,1,116,116,36,3,2,0,'2017-07-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(227,1,117,117,9,3,2,0,'2017-08-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(228,1,117,117,36,3,2,0,'2017-08-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(229,1,118,118,9,3,2,0,'2017-09-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(230,1,118,118,36,3,2,0,'2017-09-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(231,1,119,119,9,3,2,0,'2017-10-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(232,1,119,119,36,3,2,0,'2017-10-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(233,1,120,120,9,3,2,0,'2017-11-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(234,1,120,120,36,3,2,0,'2017-11-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(235,1,121,121,9,3,2,0,'2017-12-01 00:00:00',350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(236,1,121,121,36,3,2,0,'2017-12-01 00:00:00',-350.0000,'','2017-12-05 17:49:46',0,'2017-12-05 17:49:46',0),(237,1,122,122,9,3,2,0,'2016-10-01 00:00:00',4000.0000,'','2017-12-05 18:03:15',0,'2017-12-05 18:03:15',0),(238,1,122,122,17,3,2,0,'2016-10-01 00:00:00',-4000.0000,'','2017-12-05 18:03:15',0,'2017-12-05 18:03:15',0),(239,1,123,123,9,3,2,0,'2017-12-05 00:00:00',628.4500,'','2017-12-05 18:23:25',0,'2017-12-05 18:23:25',0),(240,1,123,123,37,3,2,0,'2017-12-05 00:00:00',-628.4500,'','2017-12-05 18:23:25',0,'2017-12-05 18:23:25',0),(241,1,124,124,9,3,2,0,'2017-12-05 00:00:00',-628.4500,'','2017-12-05 19:41:01',0,'2017-12-05 19:41:01',0),(242,1,124,124,37,3,2,0,'2017-12-05 00:00:00',628.4500,'','2017-12-05 19:41:01',0,'2017-12-05 19:41:01',0),(243,1,125,125,6,0,0,3,'2017-02-03 00:00:00',628.4500,'','2017-12-05 19:44:51',0,'2017-12-05 19:44:51',0),(244,1,125,125,10,0,0,3,'2017-02-03 00:00:00',-628.4500,'','2017-12-05 19:44:51',0,'2017-12-05 19:44:51',0),(245,1,129,126,72,3,2,0,'2016-10-01 00:00:00',15.0000,'','2018-02-27 20:16:16',200,'2018-02-27 20:16:16',200),(246,1,129,126,3,3,2,0,'2016-10-01 00:00:00',-15.0000,'','2018-02-27 20:16:16',200,'2018-02-27 20:16:16',200),(247,1,130,127,10,3,2,3,'2018-02-27 00:00:00',4000.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(248,1,130,127,9,3,2,3,'2018-02-27 00:00:00',-4000.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(249,1,131,128,10,3,2,3,'2018-02-27 00:00:00',4000.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(250,1,131,128,9,3,2,3,'2018-02-27 00:00:00',-4000.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(251,1,132,129,10,3,2,3,'2018-02-27 00:00:00',4000.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(252,1,132,129,9,3,2,3,'2018-02-27 00:00:00',-4000.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(253,1,133,130,10,3,2,3,'2018-02-27 00:00:00',4000.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(254,1,133,130,9,3,2,3,'2018-02-27 00:00:00',-4000.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(255,1,134,131,10,3,2,3,'2018-02-27 00:00:00',628.4500,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(256,1,134,131,9,3,2,3,'2018-02-27 00:00:00',-628.4500,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(257,1,135,132,10,3,2,3,'2018-02-27 00:00:00',3721.5500,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(258,1,135,132,9,3,2,3,'2018-02-27 00:00:00',-3721.5500,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(259,1,136,133,10,3,2,3,'2018-02-27 00:00:00',278.4500,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(260,1,136,133,9,3,2,3,'2018-02-27 00:00:00',-278.4500,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(261,1,137,134,10,3,2,3,'2018-02-27 00:00:00',175.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(262,1,137,134,9,3,2,3,'2018-02-27 00:00:00',-175.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(263,1,138,135,10,3,2,3,'2018-02-27 00:00:00',175.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(264,1,138,135,9,3,2,3,'2018-02-27 00:00:00',-175.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(265,1,139,136,10,3,2,3,'2018-02-27 00:00:00',3825.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(266,1,139,136,9,3,2,3,'2018-02-27 00:00:00',-3825.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(267,1,140,137,10,3,2,3,'2018-02-27 00:00:00',175.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(268,1,140,137,9,3,2,3,'2018-02-27 00:00:00',-175.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(269,1,141,138,10,3,2,3,'2018-02-27 00:00:00',4000.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(270,1,141,138,9,3,2,3,'2018-02-27 00:00:00',-4000.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(271,1,142,139,10,3,2,3,'2018-02-27 00:00:00',350.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(272,1,142,139,9,3,2,3,'2018-02-27 00:00:00',-350.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(273,1,143,140,10,3,2,3,'2018-02-27 00:00:00',81.7900,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(274,1,143,140,9,3,2,3,'2018-02-27 00:00:00',-81.7900,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(275,1,144,141,10,3,2,3,'2018-02-27 00:00:00',4000.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(276,1,144,141,9,3,2,3,'2018-02-27 00:00:00',-4000.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(277,1,145,142,10,3,2,3,'2018-02-27 00:00:00',350.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(278,1,145,142,9,3,2,3,'2018-02-27 00:00:00',-350.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(279,1,146,143,10,3,2,3,'2018-02-27 00:00:00',350.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(280,1,146,143,9,3,2,3,'2018-02-27 00:00:00',-350.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(281,1,147,144,10,3,2,3,'2018-02-27 00:00:00',3650.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(282,1,147,144,9,3,2,3,'2018-02-27 00:00:00',-3650.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(283,1,148,145,10,3,2,3,'2018-02-27 00:00:00',350.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(284,1,148,145,9,3,2,3,'2018-02-27 00:00:00',-350.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(285,1,149,146,10,3,2,3,'2018-02-27 00:00:00',4000.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(286,1,149,146,9,3,2,3,'2018-02-27 00:00:00',-4000.0000,'','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(287,1,150,147,10,3,2,3,'2018-02-27 00:00:00',350.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(288,1,150,147,9,3,2,3,'2018-02-27 00:00:00',-350.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(289,1,151,148,10,3,2,3,'2018-02-27 00:00:00',4000.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(290,1,151,148,9,3,2,3,'2018-02-27 00:00:00',-4000.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(291,1,152,149,10,3,2,3,'2018-02-27 00:00:00',350.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(292,1,152,149,9,3,2,3,'2018-02-27 00:00:00',-350.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(293,1,153,150,10,3,2,3,'2018-02-27 00:00:00',350.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(294,1,153,150,9,3,2,3,'2018-02-27 00:00:00',-350.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(295,1,154,151,10,3,2,3,'2018-02-27 00:00:00',3650.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(296,1,154,151,9,3,2,3,'2018-02-27 00:00:00',-3650.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(297,1,155,152,10,3,2,3,'2018-02-27 00:00:00',350.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(298,1,155,152,9,3,2,3,'2018-02-27 00:00:00',-350.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(299,1,156,153,10,3,2,3,'2018-02-27 00:00:00',4000.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(300,1,156,153,9,3,2,3,'2018-02-27 00:00:00',-4000.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(301,1,157,154,10,3,2,3,'2018-02-27 00:00:00',350.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(302,1,157,154,9,3,2,3,'2018-02-27 00:00:00',-350.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(303,1,158,155,10,3,2,3,'2018-02-27 00:00:00',409.2800,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(304,1,158,155,9,3,2,3,'2018-02-27 00:00:00',-409.2800,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(305,1,159,156,10,3,2,3,'2018-02-27 00:00:00',4000.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(306,1,159,156,9,3,2,3,'2018-02-27 00:00:00',-4000.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(307,1,160,157,10,3,2,3,'2018-02-27 00:00:00',350.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(308,1,160,157,9,3,2,3,'2018-02-27 00:00:00',-350.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(309,1,161,158,10,3,2,3,'2018-02-27 00:00:00',350.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(310,1,161,158,9,3,2,3,'2018-02-27 00:00:00',-350.0000,'','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(311,1,162,159,6,0,0,3,'2016-10-03 00:00:00',4000.0000,'','2018-02-28 17:18:33',211,'2018-02-28 17:18:33',211),(312,1,162,159,10,0,0,3,'2016-10-03 00:00:00',-4000.0000,'','2018-02-28 17:18:33',211,'2018-02-28 17:18:33',211),(313,1,163,160,10,3,2,3,'2017-12-01 00:00:00',3650.0000,'','2018-02-28 17:19:56',211,'2018-02-28 17:19:56',211),(314,1,163,160,9,3,2,3,'2017-12-01 00:00:00',-3650.0000,'','2018-02-28 17:19:56',211,'2018-02-28 17:19:56',211),(315,1,164,161,10,3,2,3,'2017-12-01 00:00:00',350.0000,'','2018-02-28 17:19:56',211,'2018-02-28 17:19:56',211),(316,1,164,161,9,3,2,3,'2017-12-01 00:00:00',-350.0000,'','2018-02-28 17:19:56',211,'2018-02-28 17:19:56',211),(317,1,165,162,9,4,3,0,'2018-01-01 00:00:00',4150.0000,'','2018-02-28 17:39:47',211,'2018-02-28 17:39:47',211),(318,1,165,162,17,4,3,0,'2018-01-01 00:00:00',-4150.0000,'','2018-02-28 17:39:47',211,'2018-02-28 17:39:47',211),(319,1,166,163,9,3,2,0,'2018-01-01 00:00:00',4000.0000,'','2018-02-28 17:41:00',211,'2018-02-28 17:41:00',211),(320,1,166,163,17,3,2,0,'2018-01-01 00:00:00',-4000.0000,'','2018-02-28 17:41:00',211,'2018-02-28 17:41:00',211),(321,1,167,164,9,3,2,0,'2018-02-01 00:00:00',4000.0000,'','2018-02-28 17:41:40',211,'2018-02-28 17:41:40',211),(322,1,167,164,17,3,2,0,'2018-02-01 00:00:00',-4000.0000,'','2018-02-28 17:41:40',211,'2018-02-28 17:41:40',211),(323,1,168,165,9,3,2,0,'2018-01-01 00:00:00',350.0000,'','2018-02-28 21:12:57',200,'2018-02-28 21:12:57',200),(324,1,168,165,36,3,2,0,'2018-01-01 00:00:00',-350.0000,'','2018-02-28 21:12:57',200,'2018-02-28 21:12:57',200),(325,1,169,166,9,3,2,0,'2018-02-01 00:00:00',350.0000,'','2018-02-28 21:13:21',200,'2018-02-28 21:13:21',200),(326,1,169,166,36,3,2,0,'2018-02-01 00:00:00',-350.0000,'','2018-02-28 21:13:21',200,'2018-02-28 21:13:21',200),(327,1,170,167,6,0,0,1,'2018-01-01 00:00:00',3750.0000,'','2018-02-28 21:16:01',200,'2018-02-28 21:16:01',200),(328,1,170,167,10,0,0,1,'2018-01-01 00:00:00',-3750.0000,'','2018-02-28 21:16:01',200,'2018-02-28 21:16:01',200),(329,1,171,168,6,0,0,4,'2018-01-01 00:00:00',4150.0000,'','2018-02-28 21:16:23',200,'2018-02-28 21:16:23',200),(330,1,171,168,10,0,0,4,'2018-01-01 00:00:00',-4150.0000,'','2018-02-28 21:16:23',200,'2018-02-28 21:16:23',200),(331,1,172,169,6,0,0,4,'2018-02-01 00:00:00',4150.0000,'','2018-02-28 21:16:42',200,'2018-02-28 21:16:42',200),(332,1,172,169,10,0,0,4,'2018-02-01 00:00:00',-4150.0000,'','2018-02-28 21:16:42',200,'2018-02-28 21:16:42',200),(333,1,173,170,6,0,0,3,'2018-02-23 00:00:00',13050.0000,'','2018-02-28 21:17:07',200,'2018-02-28 21:17:07',200),(334,1,173,170,10,0,0,3,'2018-02-23 00:00:00',-13050.0000,'','2018-02-28 21:17:07',200,'2018-02-28 21:17:07',200),(335,1,174,171,10,2,1,1,'2018-02-28 00:00:00',3750.0000,'','2018-02-28 21:17:35',200,'2018-02-28 21:17:35',200),(336,1,174,171,9,2,1,1,'2018-02-28 00:00:00',-3750.0000,'','2018-02-28 21:17:35',200,'2018-02-28 21:17:35',200),(337,1,175,172,10,3,2,3,'2018-02-28 00:00:00',4000.0000,'','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(338,1,175,172,9,3,2,3,'2018-02-28 00:00:00',-4000.0000,'','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(339,1,176,173,10,3,2,3,'2018-02-28 00:00:00',350.0000,'','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(340,1,176,173,9,3,2,3,'2018-02-28 00:00:00',-350.0000,'','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(341,1,177,174,10,3,2,3,'2018-02-28 00:00:00',4000.0000,'','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(342,1,177,174,9,3,2,3,'2018-02-28 00:00:00',-4000.0000,'','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(343,1,178,175,10,3,2,3,'2018-02-28 00:00:00',350.0000,'','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(344,1,178,175,9,3,2,3,'2018-02-28 00:00:00',-350.0000,'','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(345,1,179,176,10,4,3,4,'2018-02-28 00:00:00',4150.0000,'','2018-02-28 21:17:57',200,'2018-02-28 21:17:57',200),(346,1,179,176,9,4,3,4,'2018-02-28 00:00:00',-4150.0000,'','2018-02-28 21:17:57',200,'2018-02-28 21:17:57',200),(347,1,180,177,10,4,3,4,'2018-02-28 00:00:00',4150.0000,'','2018-02-28 21:17:57',200,'2018-02-28 21:17:57',200),(348,1,180,177,9,4,3,4,'2018-02-28 00:00:00',-4150.0000,'','2018-02-28 21:17:57',200,'2018-02-28 21:17:57',200),(349,1,187,184,9,3,2,0,'2018-03-01 00:00:00',350.0000,'','2018-05-30 19:45:29',200,'2018-05-30 19:45:29',200),(350,1,187,184,36,3,2,0,'2018-03-01 00:00:00',-350.0000,'','2018-05-30 19:45:29',200,'2018-05-30 19:45:29',200),(351,1,188,185,9,3,2,0,'2018-04-01 00:00:00',350.0000,'','2018-05-30 19:45:29',200,'2018-05-30 19:45:29',200),(352,1,188,185,36,3,2,0,'2018-04-01 00:00:00',-350.0000,'','2018-05-30 19:45:29',200,'2018-05-30 19:45:29',200),(353,1,189,186,9,3,2,0,'2018-05-01 00:00:00',350.0000,'','2018-05-30 19:45:29',200,'2018-05-30 19:45:29',200),(354,1,189,186,36,3,2,0,'2018-05-01 00:00:00',-350.0000,'','2018-05-30 19:45:29',200,'2018-05-30 19:45:29',200),(355,1,190,187,9,3,2,0,'2018-03-01 00:00:00',4000.0000,'','2018-05-30 19:50:19',200,'2018-05-30 19:50:19',200),(356,1,190,187,17,3,2,0,'2018-03-01 00:00:00',-4000.0000,'','2018-05-30 19:50:19',200,'2018-05-30 19:50:19',200),(357,1,191,188,9,3,2,0,'2018-04-01 00:00:00',4000.0000,'','2018-05-30 19:50:19',200,'2018-05-30 19:50:19',200),(358,1,191,188,17,3,2,0,'2018-04-01 00:00:00',-4000.0000,'','2018-05-30 19:50:19',200,'2018-05-30 19:50:19',200),(359,1,192,189,9,3,2,0,'2018-05-01 00:00:00',4000.0000,'','2018-05-30 19:50:19',200,'2018-05-30 19:50:19',200),(360,1,192,189,17,3,2,0,'2018-05-01 00:00:00',-4000.0000,'','2018-05-30 19:50:19',200,'2018-05-30 19:50:19',200),(361,1,193,190,10,3,2,3,'2018-03-01 00:00:00',350.0000,'','2018-05-30 20:04:19',200,'2018-05-30 20:04:19',200),(362,1,193,190,9,3,2,3,'2018-03-01 00:00:00',-350.0000,'','2018-05-30 20:04:19',200,'2018-05-30 20:04:19',200),(363,1,194,191,10,3,2,3,'2018-03-01 00:00:00',4000.0000,'','2018-05-30 20:04:19',200,'2018-05-30 20:04:19',200),(364,1,194,191,9,3,2,3,'2018-03-01 00:00:00',-4000.0000,'','2018-05-30 20:04:19',200,'2018-05-30 20:04:19',200),(365,1,195,192,6,0,0,3,'2018-05-20 00:00:00',13050.0000,'','2018-05-30 20:08:28',200,'2018-05-30 20:08:28',200),(366,1,195,192,10,0,0,3,'2018-05-20 00:00:00',-13050.0000,'','2018-05-30 20:08:28',200,'2018-05-30 20:08:28',200),(367,1,196,193,10,3,2,3,'2018-05-20 00:00:00',350.0000,'','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(368,1,196,193,9,3,2,3,'2018-05-20 00:00:00',-350.0000,'','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(369,1,197,194,10,3,2,3,'2018-05-20 00:00:00',4000.0000,'','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(370,1,197,194,9,3,2,3,'2018-05-20 00:00:00',-4000.0000,'','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(371,1,198,195,10,3,2,3,'2018-05-20 00:00:00',350.0000,'','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(372,1,198,195,9,3,2,3,'2018-05-20 00:00:00',-350.0000,'','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(373,1,199,196,10,3,2,3,'2018-05-20 00:00:00',4000.0000,'','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(374,1,199,196,9,3,2,3,'2018-05-20 00:00:00',-4000.0000,'','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200);
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
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerMarker`
--

LOCK TABLES `LedgerMarker` WRITE;
/*!40000 ALTER TABLE `LedgerMarker` DISABLE KEYS */;
INSERT INTO `LedgerMarker` VALUES (1,1,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(2,2,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(3,3,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(4,4,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(6,6,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(7,7,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(8,8,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(9,9,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(10,10,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(11,11,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(12,12,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(13,13,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(14,14,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(15,15,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(16,16,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(17,17,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(19,19,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(20,20,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(21,21,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(22,22,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(23,23,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(24,24,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(25,25,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(26,26,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(27,27,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(28,28,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(29,29,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(30,30,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(31,31,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(32,32,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(33,33,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(34,34,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(35,35,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(36,36,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(37,37,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(38,38,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(39,39,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(40,40,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(41,41,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(42,42,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(43,43,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(44,44,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(45,45,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(46,46,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(47,47,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(48,48,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(49,49,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(50,50,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(51,51,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(52,52,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(53,53,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(54,54,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(55,55,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(56,56,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(57,57,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(58,58,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(59,59,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(60,60,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(61,61,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(62,62,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(63,63,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(64,64,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(65,65,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(66,66,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(67,67,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(68,68,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(69,69,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(70,70,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(71,71,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(72,72,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(73,73,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(75,0,0,1,0,0,'2017-11-28 00:00:00',0.0000,3,'2017-11-28 18:14:19',0,'2017-11-28 18:14:19',0),(76,0,1,0,0,1,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:15:20',0,'2017-11-30 18:15:20',0),(77,0,1,0,0,2,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:15:28',0,'2017-11-30 18:15:28',0),(78,0,1,0,0,3,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:16:10',0,'2017-11-30 18:16:10',0),(79,0,1,0,0,4,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:16:17',0,'2017-11-30 18:16:17',0),(80,0,1,0,0,5,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:16:28',0,'2017-11-30 18:16:28',0),(81,0,0,2,0,0,'2014-03-01 00:00:00',0.0000,3,'2017-11-30 18:29:02',0,'2017-11-30 18:17:55',0),(82,0,1,2,1,0,'2014-03-01 00:00:00',0.0000,3,'2017-11-30 18:20:15',0,'2017-11-30 18:20:15',0),(83,0,1,0,0,6,'1970-01-01 00:00:00',0.0000,3,'2017-11-30 18:24:52',0,'2017-11-30 18:24:52',0),(84,0,0,3,0,0,'2016-10-01 00:00:00',0.0000,3,'2017-11-30 18:33:53',0,'2017-11-30 18:29:25',0),(85,0,1,3,2,0,'2016-10-01 00:00:00',0.0000,3,'2017-11-30 18:32:13',0,'2017-11-30 18:32:13',0),(86,0,0,4,0,0,'2016-07-01 00:00:00',0.0000,3,'2017-11-30 18:37:13',0,'2017-11-30 18:33:59',0),(87,0,1,4,3,0,'2016-07-01 00:00:00',0.0000,3,'2017-11-30 18:34:33',0,'2017-11-30 18:34:33',0),(88,0,0,5,0,0,'2017-11-30 00:00:00',0.0000,3,'2017-11-30 18:37:24',0,'2017-11-30 18:37:24',0);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
  `TaxpayorID` char(128) NOT NULL DEFAULT '',
  `CreditLimit` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `EligibleFuturePayor` tinyint(1) NOT NULL DEFAULT '1',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
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
INSERT INTO `Payor` VALUES (1,1,'',0.0000,1,0,'',0.0000,'2017-11-30 18:15:20',0,'2017-11-30 18:15:20',0),(2,1,'',0.0000,1,0,'',0.0000,'2017-11-30 18:15:28',0,'2017-11-30 18:15:28',0),(3,1,'',0.0000,1,0,'',0.0000,'2017-11-30 18:16:10',0,'2017-11-30 18:16:10',0),(4,1,'',0.0000,1,0,'',0.0000,'2017-11-30 18:16:17',0,'2017-11-30 18:16:17',0),(5,1,'',0.0000,1,0,'',0.0000,'2017-11-30 18:16:28',0,'2017-11-30 18:16:28',0),(6,1,'',0.0000,1,0,'',0.0000,'2017-11-30 18:24:52',0,'2017-11-30 18:24:52',0);
/*!40000 ALTER TABLE `Payor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Pets`
--

DROP TABLE IF EXISTS `Pets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Pets` (
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
-- Dumping data for table `Pets`
--

LOCK TABLES `Pets` WRITE;
/*!40000 ALTER TABLE `Pets` DISABLE KEYS */;
/*!40000 ALTER TABLE `Pets` ENABLE KEYS */;
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
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `EvictedDes` varchar(2048) NOT NULL DEFAULT '',
  `ConvictedDes` varchar(2048) NOT NULL DEFAULT '',
  `BankruptcyDes` varchar(2048) NOT NULL DEFAULT '',
  `OtherPreferences` varchar(1024) NOT NULL DEFAULT '',
  `SpecialNeeds` varchar(1024) NOT NULL DEFAULT '',
  `FollowUpDate` date NOT NULL DEFAULT '1970-01-01',
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
  `ThirdPartySource` varchar(100) NOT NULL DEFAULT '',
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
INSERT INTO `Prospect` VALUES (1,1,'','','','','','','',0,'','','','','','1900-01-01','','','',0,'','','','',0,'','','0','2017-11-30 18:15:20',0,'2017-11-30 18:15:20',0),(2,1,'','','','','','','',0,'','','','','','1900-01-01','','','',0,'','','','',0,'','','0','2017-11-30 18:15:28',0,'2017-11-30 18:15:28',0),(3,1,'','','','','','','',0,'','','','','','1900-01-01','','','',0,'','','','',0,'','','0','2017-11-30 18:16:10',0,'2017-11-30 18:16:10',0),(4,1,'','','','','','','',0,'','','','','','1900-01-01','','','',0,'','','','',0,'','','0','2017-11-30 18:16:17',0,'2017-11-30 18:16:17',0),(5,1,'','','','','','','',0,'','','','','','1900-01-01','','','',0,'','','','',0,'','','0','2017-11-30 18:16:28',0,'2017-11-30 18:16:28',0),(6,1,'','','','','','','',0,'','','','','','1900-01-01','','','',0,'','','','',0,'','','0','2017-11-30 18:24:52',0,'2017-11-30 18:24:52',0);
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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
  `RPRRTRateID` bigint(20) NOT NULL AUTO_INCREMENT,
  `RPRID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RTID` bigint(20) NOT NULL DEFAULT '0',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `Val` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RPRRTRateID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
  `RPRSPRateID` bigint(20) NOT NULL AUTO_INCREMENT,
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
  PRIMARY KEY (`RPRSPRateID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Receipt`
--

LOCK TABLES `Receipt` WRITE;
/*!40000 ALTER TABLE `Receipt` DISABLE KEYS */;
INSERT INTO `Receipt` VALUES (1,0,1,1,2,0,0,0,'2014-03-01 00:00:00','1234',7000.0000,'',10,'',4,'Reversed by receipt RCPT00000005','','2017-11-30 19:24:46',0,'2017-11-30 18:48:02',0),(2,0,1,4,2,0,0,0,'2016-07-01 00:00:00','2345',8300.0000,'',10,'',4,'Reversed by receipt RCPT00000004','','2017-11-30 19:16:53',0,'2017-11-30 18:49:17',0),(3,0,1,4,2,0,3,0,'2016-07-01 00:00:00','2456',8300.0000,'',25,'ASM(2) d 12999 8300.00,c 12001 8300.00',2,'','','2017-12-05 17:06:25',0,'2017-11-30 19:13:23',0),(4,2,1,4,2,0,0,0,'2016-07-01 00:00:00','2345',-8300.0000,'',10,'',4,'Reversal of receipt RCPT00000002','','2017-11-30 19:16:53',0,'2017-11-30 19:13:59',0),(5,1,1,1,2,0,0,0,'2014-03-01 00:00:00','1234',-7000.0000,'',10,'',4,'Reversal of receipt RCPT00000001','','2017-11-30 19:24:46',0,'2017-11-30 19:23:47',0),(6,0,1,1,2,0,4,0,'2014-03-01 00:00:00','3457',7000.0000,'',25,'ASM(1) d 12999 7000.00,c 12001 7000.00',2,'','','2017-11-30 19:46:56',0,'2017-11-30 19:24:32',0),(7,0,1,1,6,0,0,0,'2017-01-01 00:00:00','2354',3750.0000,'',25,'ASM(4) d 12999 3750.00,c 12001 3750.00',4,'Reversed by receipt RCPT00000014','','2017-12-05 16:19:04',0,'2017-11-30 19:44:52',0),(8,0,1,1,6,0,0,0,'2017-01-01 00:00:00','',3750.0000,'',25,'ASM(4) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:31',0,'2017-12-05 16:09:37',0),(9,0,1,4,2,0,0,0,'2017-01-01 00:00:00','',4150.0000,'',25,'ASM(28) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:10:06',0),(10,0,1,3,6,0,0,3,'2017-02-01 00:00:00','',8350.0000,'',25,'ASM(16) d 12999 4000.00,c 12001 4000.00,ASM(42) d 12999 628.45,c 12001 628.45,ASM(17) d 12999 3721.55,c 12001 3721.55',2,'','','2018-02-27 21:48:57',200,'2017-12-05 16:12:02',0),(11,0,1,1,6,0,0,0,'2017-02-01 00:00:00','',3750.0000,'',25,'ASM(5) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:31',0,'2017-12-05 16:12:32',0),(12,0,1,4,2,0,0,0,'2017-02-01 00:00:00','',4150.0000,'',25,'ASM(29) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:13:44',0),(13,0,1,3,6,0,0,3,'2016-11-15 00:00:00','',12000.0000,'',25,'ASM(59) d 12999 4000.00,c 12001 4000.00,ASM(47) d 12999 4000.00,c 12001 4000.00,ASM(48) d 12999 4000.00,c 12001 4000.00',2,'3 month rent in advance','','2018-02-27 21:48:57',200,'2017-12-05 16:16:37',0),(14,7,1,1,6,0,0,0,'2017-01-01 00:00:00','2354',-3750.0000,'',25,'ASM(4) d 12999 3750.00,c 12001 3750.00',4,'Reversal of receipt RCPT00000007','','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(15,0,1,1,6,0,0,0,'2017-03-01 00:00:00','',3750.0000,'',25,'ASM(6) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:31',0,'2017-12-05 16:20:28',0),(16,0,1,4,2,0,0,0,'2017-03-01 00:00:00','',4150.0000,'',25,'ASM(30) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:22:08',0),(17,0,1,1,6,0,0,0,'2017-04-01 00:00:00','',3750.0000,'',25,'ASM(7) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:31',0,'2017-12-05 16:22:50',0),(18,0,1,4,2,0,0,0,'2017-04-01 00:00:00','',4150.0000,'',25,'ASM(31) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:23:12',0),(19,0,1,1,6,0,0,0,'2017-05-01 00:00:00','',3750.0000,'',25,'ASM(8) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:32',0,'2017-12-05 16:24:00',0),(20,0,1,4,2,0,0,0,'2017-05-01 00:00:00','',4150.0000,'',25,'ASM(32) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:24:18',0),(21,0,1,3,6,0,0,3,'2017-05-15 00:00:00','',13131.7900,'',25,'ASM(18) d 12999 3825.00,c 12001 3825.00,ASM(44) d 12999 175.00,c 12001 175.00,ASM(19) d 12999 4000.00,c 12001 4000.00,ASM(50) d 12999 350.00,c 12001 350.00,ASM(45) d 12999 81.79,c 12001 81.79,ASM(20) d 12999 4000.00,c 12001 4000.00,ASM(51) d 12999 350.00,c 12001 350.00,ASM(21) d 12999 350.00,c 12001 350.00',2,'3 month rent in advance and utilities overage','','2018-02-27 21:48:57',200,'2017-12-05 16:26:21',0),(22,0,1,1,6,0,0,0,'2017-06-01 00:00:00','',3750.0000,'',25,'ASM(9) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:32',0,'2017-12-05 16:27:03',0),(23,0,1,4,2,0,0,0,'2017-06-01 00:00:00','',4150.0000,'',25,'ASM(33) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:27:16',0),(24,0,1,1,6,0,0,0,'2017-07-01 00:00:00','',3750.0000,'',25,'ASM(10) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:32',0,'2017-12-05 16:27:58',0),(25,0,1,4,2,0,0,0,'2017-07-01 00:00:00','',4150.0000,'',25,'ASM(34) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:28:12',0),(26,0,1,1,6,0,0,0,'2017-08-01 00:00:00','',3750.0000,'',25,'ASM(11) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:32',0,'2017-12-05 16:29:16',0),(27,0,1,4,2,0,0,0,'2017-08-01 00:00:00','',4150.0000,'',25,'ASM(35) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:29:33',0),(28,0,1,3,6,0,0,3,'2017-08-15 00:00:00','',13050.0000,'',25,'ASM(21) d 12999 3650.00,c 12001 3650.00,ASM(52) d 12999 350.00,c 12001 350.00,ASM(22) d 12999 4000.00,c 12001 4000.00,ASM(53) d 12999 350.00,c 12001 350.00,ASM(23) d 12999 4000.00,c 12001 4000.00,ASM(54) d 12999 350.00,c 12001 350.00,ASM(24) d 12999 350.00,c 12001 350.00',2,'3 month rent in advance','','2018-02-27 21:48:57',200,'2017-12-05 16:29:59',0),(29,0,1,1,6,0,0,0,'2017-09-01 00:00:00','',3750.0000,'',25,'ASM(12) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:32',0,'2017-12-05 16:30:33',0),(30,0,1,4,2,0,0,0,'2017-09-01 00:00:00','',4150.0000,'',25,'ASM(36) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:30:51',0),(31,0,1,1,6,0,0,0,'2017-10-01 00:00:00','',3750.0000,'',25,'ASM(13) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:32',0,'2017-12-05 16:31:42',0),(32,0,1,4,2,0,0,0,'2017-10-01 00:00:00','',4150.0000,'',25,'ASM(37) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:31:56',0),(33,0,1,1,6,0,0,0,'2017-11-01 00:00:00','',3750.0000,'',25,'ASM(14) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:32',0,'2017-12-05 16:32:48',0),(34,0,1,4,2,0,0,0,'2017-11-01 00:00:00','',4150.0000,'',25,'ASM(38) d 12999 4150.00,c 12001 4150.00',2,'','Lauren Beck','2017-12-05 17:06:25',0,'2017-12-05 16:33:11',0),(35,0,1,3,6,0,0,3,'2017-11-15 00:00:00','',13459.2800,'',25,'ASM(24) d 12999 3650.00,c 12001 3650.00,ASM(55) d 12999 350.00,c 12001 350.00,ASM(25) d 12999 4000.00,c 12001 4000.00,ASM(56) d 12999 350.00,c 12001 350.00,ASM(46) d 12999 409.28,c 12001 409.28,ASM(26) d 12999 4000.00,c 12001 4000.00,ASM(57) d 12999 350.00,c 12001 350.00,ASM(40) d 12999 350.00,c 12001 350.00',2,'3 month rent in advance and utilities overage','','2018-02-27 21:48:57',200,'2017-12-05 16:40:59',0),(36,0,1,1,6,0,0,0,'2017-12-01 00:00:00','',3750.0000,'',25,'ASM(39) d 12999 3750.00,c 12001 3750.00',2,'','Kirsten Read','2017-12-05 16:59:32',0,'2017-12-05 16:42:24',0),(37,0,1,4,2,0,0,0,'2017-12-01 00:00:00','',4150.0000,'',25,'ASM(41) d 12999 4150.00,c 12001 4150.00',2,'','','2017-12-05 17:06:25',0,'2017-12-05 16:42:35',0),(38,0,1,3,6,0,0,3,'2017-02-03 00:00:00','',628.4500,'',25,'ASM(17) d 12999 278.45,c 12001 278.45,ASM(43) d 12999 175.00,c 12001 175.00,ASM(18) d 12999 175.00,c 12001 175.00',2,'','','2018-02-27 21:48:57',200,'2017-12-05 19:44:51',0),(39,0,1,3,2,0,0,0,'2016-10-03 00:00:00','1234',4000.0000,'',25,'ASM(40) d 12999 3650.00,c 12001 3650.00,ASM(58) d 12999 350.00,c 12001 350.00',2,'','','2018-02-28 17:19:56',211,'2018-02-28 17:18:33',211),(40,0,1,1,6,0,0,0,'2018-01-01 00:00:00','999',3750.0000,'',25,'ASM(62) d 12999 3750.00,c 12001 3750.00',2,'','','2018-02-28 21:17:34',200,'2018-02-28 21:16:01',200),(41,0,1,4,2,0,0,0,'2018-01-01 00:00:00','8888',4150.0000,'',25,'ASM(65) d 12999 4150.00,c 12001 4150.00',2,'','','2018-02-28 21:17:57',200,'2018-02-28 21:16:23',200),(42,0,1,4,2,0,0,0,'2018-02-01 00:00:00','889',4150.0000,'',25,'ASM(63) d 12999 4150.00,c 12001 4150.00',2,'','','2018-02-28 21:17:57',200,'2018-02-28 21:16:42',200),(43,0,1,3,6,0,0,0,'2018-02-23 00:00:00','9898',13050.0000,'',25,'ASM(66) d 12999 4000.00,c 12001 4000.00,ASM(68) d 12999 350.00,c 12001 350.00,ASM(67) d 12999 4000.00,c 12001 4000.00,ASM(69) d 12999 350.00,c 12001 350.00,ASM(83) d 12999 350.00,c 12001 350.00,ASM(87) d 12999 4000.00,c 12001 4000.00',2,'','','2018-05-30 20:04:19',200,'2018-02-28 21:17:07',200),(44,0,1,3,6,0,0,0,'2018-05-20 00:00:00','999X',13050.0000,'',25,'ASM(84) d 12999 350.00,c 12001 350.00,ASM(88) d 12999 4000.00,c 12001 4000.00,ASM(85) d 12999 350.00,c 12001 350.00,ASM(89) d 12999 4000.00,c 12001 4000.00',1,'','','2018-05-30 20:09:39',200,'2018-05-30 20:08:28',200);
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
) ENGINE=InnoDB AUTO_INCREMENT=122 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ReceiptAllocation`
--

LOCK TABLES `ReceiptAllocation` WRITE;
/*!40000 ALTER TABLE `ReceiptAllocation` DISABLE KEYS */;
INSERT INTO `ReceiptAllocation` VALUES (1,1,1,0,'2014-03-01 00:00:00',7000.0000,0,4,'d 10104 _, c 10999 _','2017-11-30 19:23:47',0,'2017-11-30 18:48:02',0),(2,2,1,0,'2016-07-01 00:00:00',8300.0000,0,4,'d 10104 _, c 10999 _','2017-11-30 19:13:59',0,'2017-11-30 18:49:17',0),(3,3,1,0,'2016-07-01 00:00:00',8300.0000,0,0,'d 10999 _, c 12999 _','2017-11-30 19:13:23',0,'2017-11-30 19:13:23',0),(4,4,1,0,'2016-07-01 00:00:00',-8300.0000,0,4,'d 10104 _, c 10999 _','2017-11-30 19:13:59',0,'2017-11-30 19:13:59',0),(5,3,1,0,'2016-07-01 00:00:00',8300.0000,0,0,'d 10104 _, c 10999 _','2017-11-30 19:17:32',0,'2017-11-30 19:17:32',0),(6,5,1,0,'2014-03-01 00:00:00',-7000.0000,0,4,'d 10104 _, c 10999 _','2017-11-30 19:23:47',0,'2017-11-30 19:23:47',0),(7,6,1,0,'2014-03-01 00:00:00',7000.0000,0,0,'d 10999 _, c 12999 _','2017-11-30 19:24:32',0,'2017-11-30 19:24:32',0),(8,6,1,0,'2014-03-01 00:00:00',7000.0000,0,0,'d 10104 _, c 10999 _','2017-11-30 19:25:24',0,'2017-11-30 19:25:24',0),(9,7,1,0,'2017-01-01 00:00:00',3750.0000,0,4,'d 10999 _, c 12999 _','2017-12-05 16:19:04',0,'2017-11-30 19:44:52',0),(10,6,1,2,'2014-03-01 00:00:00',7000.0000,1,0,'ASM(1) d 12999 7000.00,c 12001 7000.00','2017-11-30 19:46:56',0,'2017-11-30 19:46:56',0),(11,7,1,2,'2017-01-01 00:00:00',3750.0000,4,4,'ASM(4) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:19:04',0,'2017-11-30 19:46:56',0),(12,8,1,0,'2017-01-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:09:37',0,'2017-12-05 16:09:37',0),(13,9,1,0,'2017-01-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:10:06',0,'2017-12-05 16:10:06',0),(14,10,1,0,'2017-02-01 00:00:00',8350.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:12:02',0,'2017-12-05 16:12:02',0),(15,11,1,0,'2017-02-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:12:32',0,'2017-12-05 16:12:32',0),(16,12,1,0,'2017-02-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:13:44',0,'2017-12-05 16:13:44',0),(17,13,1,0,'2016-11-15 00:00:00',12000.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:16:37',0,'2017-12-05 16:16:37',0),(18,14,1,0,'2017-01-01 00:00:00',-3750.0000,0,4,'d 10999 _, c 12999 _','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(19,14,1,2,'2017-12-05 16:19:04',-3750.0000,4,4,'ASM(4) d 12999 -3750.00,ASM(4) c 12001 -3750.00','2017-12-05 16:19:04',0,'2017-12-05 16:19:04',0),(20,15,1,0,'2017-03-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:20:28',0,'2017-12-05 16:20:28',0),(21,16,1,0,'2017-03-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:22:08',0,'2017-12-05 16:22:08',0),(22,17,1,0,'2017-04-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:22:50',0,'2017-12-05 16:22:50',0),(23,18,1,0,'2017-04-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:23:12',0,'2017-12-05 16:23:12',0),(24,19,1,0,'2017-05-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:24:00',0,'2017-12-05 16:24:00',0),(25,20,1,0,'2017-05-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:24:18',0,'2017-12-05 16:24:18',0),(26,21,1,0,'2017-05-15 00:00:00',13131.7900,0,0,'d 10999 _, c 12999 _','2017-12-05 16:26:21',0,'2017-12-05 16:26:21',0),(27,22,1,0,'2017-06-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:27:03',0,'2017-12-05 16:27:03',0),(28,23,1,0,'2017-06-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:27:16',0,'2017-12-05 16:27:16',0),(29,24,1,0,'2017-07-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:27:58',0,'2017-12-05 16:27:58',0),(30,25,1,0,'2017-07-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:28:12',0,'2017-12-05 16:28:12',0),(31,26,1,0,'2017-08-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:29:16',0,'2017-12-05 16:29:16',0),(32,27,1,0,'2017-08-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:29:33',0,'2017-12-05 16:29:33',0),(33,28,1,0,'2017-08-15 00:00:00',13050.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:29:59',0,'2017-12-05 16:29:59',0),(34,29,1,0,'2017-09-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:30:33',0,'2017-12-05 16:30:33',0),(35,30,1,0,'2017-09-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:30:51',0,'2017-12-05 16:30:51',0),(36,31,1,0,'2017-10-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:31:42',0,'2017-12-05 16:31:42',0),(37,32,1,0,'2017-10-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:31:56',0,'2017-12-05 16:31:56',0),(38,33,1,0,'2017-11-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:32:48',0,'2017-12-05 16:32:48',0),(39,34,1,0,'2017-11-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:33:11',0,'2017-12-05 16:33:11',0),(40,35,1,0,'2017-11-15 00:00:00',13459.2800,0,0,'d 10999 _, c 12999 _','2017-12-05 16:40:59',0,'2017-12-05 16:40:59',0),(41,36,1,0,'2017-12-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:42:24',0,'2017-12-05 16:42:24',0),(42,37,1,0,'2017-12-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2017-12-05 16:42:35',0,'2017-12-05 16:42:35',0),(43,8,1,2,'2017-01-01 00:00:00',3750.0000,4,0,'ASM(4) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(44,11,1,2,'2017-02-01 00:00:00',3750.0000,5,0,'ASM(5) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(45,15,1,2,'2017-03-01 00:00:00',3750.0000,6,0,'ASM(6) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(46,17,1,2,'2017-04-01 00:00:00',3750.0000,7,0,'ASM(7) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(47,19,1,2,'2017-05-01 00:00:00',3750.0000,8,0,'ASM(8) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:31',0,'2017-12-05 16:59:31',0),(48,22,1,2,'2017-06-01 00:00:00',3750.0000,9,0,'ASM(9) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(49,24,1,2,'2017-07-01 00:00:00',3750.0000,10,0,'ASM(10) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(50,26,1,2,'2017-08-01 00:00:00',3750.0000,11,0,'ASM(11) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(51,29,1,2,'2017-09-01 00:00:00',3750.0000,12,0,'ASM(12) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(52,31,1,2,'2017-10-01 00:00:00',3750.0000,13,0,'ASM(13) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(53,33,1,2,'2017-11-01 00:00:00',3750.0000,14,0,'ASM(14) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(54,36,1,2,'2017-12-01 00:00:00',3750.0000,39,0,'ASM(39) d 12999 3750.00,c 12001 3750.00','2017-12-05 16:59:32',0,'2017-12-05 16:59:32',0),(55,3,1,4,'2016-07-01 00:00:00',8300.0000,2,0,'ASM(2) d 12999 8300.00,c 12001 8300.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(56,9,1,4,'2017-01-01 00:00:00',4150.0000,28,0,'ASM(28) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(57,12,1,4,'2017-02-01 00:00:00',4150.0000,29,0,'ASM(29) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(58,16,1,4,'2017-03-01 00:00:00',4150.0000,30,0,'ASM(30) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(59,18,1,4,'2017-04-01 00:00:00',4150.0000,31,0,'ASM(31) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(60,20,1,4,'2017-05-01 00:00:00',4150.0000,32,0,'ASM(32) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(61,23,1,4,'2017-06-01 00:00:00',4150.0000,33,0,'ASM(33) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(62,25,1,4,'2017-07-01 00:00:00',4150.0000,34,0,'ASM(34) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(63,27,1,4,'2017-08-01 00:00:00',4150.0000,35,0,'ASM(35) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(64,30,1,4,'2017-09-01 00:00:00',4150.0000,36,0,'ASM(36) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(65,32,1,4,'2017-10-01 00:00:00',4150.0000,37,0,'ASM(37) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(66,34,1,4,'2017-11-01 00:00:00',4150.0000,38,0,'ASM(38) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(67,37,1,4,'2017-12-01 00:00:00',4150.0000,41,0,'ASM(41) d 12999 4150.00,c 12001 4150.00','2017-12-05 17:06:25',0,'2017-12-05 17:06:25',0),(68,38,1,0,'2017-02-03 00:00:00',628.4500,0,0,'d 10999 _, c 12999 _','2017-12-05 19:44:51',0,'2017-12-05 19:44:51',0),(69,13,1,3,'2018-02-27 00:00:00',4000.0000,59,0,'ASM(59) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(70,13,1,3,'2018-02-27 00:00:00',4000.0000,47,0,'ASM(47) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(71,13,1,3,'2018-02-27 00:00:00',4000.0000,48,0,'ASM(48) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(72,10,1,3,'2018-02-27 00:00:00',4000.0000,16,0,'ASM(16) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(73,10,1,3,'2018-02-27 00:00:00',628.4500,42,0,'ASM(42) d 12999 628.45,c 12001 628.45','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(74,10,1,3,'2018-02-27 00:00:00',3721.5500,17,0,'ASM(17) d 12999 3721.55,c 12001 3721.55','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(75,38,1,3,'2018-02-27 00:00:00',278.4500,17,0,'ASM(17) d 12999 278.45,c 12001 278.45','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(76,38,1,3,'2018-02-27 00:00:00',175.0000,43,0,'ASM(43) d 12999 175.00,c 12001 175.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(77,38,1,3,'2018-02-27 00:00:00',175.0000,18,0,'ASM(18) d 12999 175.00,c 12001 175.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(78,21,1,3,'2018-02-27 00:00:00',3825.0000,18,0,'ASM(18) d 12999 3825.00,c 12001 3825.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(79,21,1,3,'2018-02-27 00:00:00',175.0000,44,0,'ASM(44) d 12999 175.00,c 12001 175.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(80,21,1,3,'2018-02-27 00:00:00',4000.0000,19,0,'ASM(19) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(81,21,1,3,'2018-02-27 00:00:00',350.0000,50,0,'ASM(50) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(82,21,1,3,'2018-02-27 00:00:00',81.7900,45,0,'ASM(45) d 12999 81.79,c 12001 81.79','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(83,21,1,3,'2018-02-27 00:00:00',4000.0000,20,0,'ASM(20) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(84,21,1,3,'2018-02-27 00:00:00',350.0000,51,0,'ASM(51) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(85,21,1,3,'2018-02-27 00:00:00',350.0000,21,0,'ASM(21) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(86,28,1,3,'2018-02-27 00:00:00',3650.0000,21,0,'ASM(21) d 12999 3650.00,c 12001 3650.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(87,28,1,3,'2018-02-27 00:00:00',350.0000,52,0,'ASM(52) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(88,28,1,3,'2018-02-27 00:00:00',4000.0000,22,0,'ASM(22) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(89,28,1,3,'2018-02-27 00:00:00',350.0000,53,0,'ASM(53) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:52',200,'2018-02-27 20:16:52',200),(90,28,1,3,'2018-02-27 00:00:00',4000.0000,23,0,'ASM(23) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(91,28,1,3,'2018-02-27 00:00:00',350.0000,54,0,'ASM(54) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(92,28,1,3,'2018-02-27 00:00:00',350.0000,24,0,'ASM(24) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(93,35,1,3,'2018-02-27 00:00:00',3650.0000,24,0,'ASM(24) d 12999 3650.00,c 12001 3650.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(94,35,1,3,'2018-02-27 00:00:00',350.0000,55,0,'ASM(55) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(95,35,1,3,'2018-02-27 00:00:00',4000.0000,25,0,'ASM(25) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(96,35,1,3,'2018-02-27 00:00:00',350.0000,56,0,'ASM(56) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(97,35,1,3,'2018-02-27 00:00:00',409.2800,46,0,'ASM(46) d 12999 409.28,c 12001 409.28','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(98,35,1,3,'2018-02-27 00:00:00',4000.0000,26,0,'ASM(26) d 12999 4000.00,c 12001 4000.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(99,35,1,3,'2018-02-27 00:00:00',350.0000,57,0,'ASM(57) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(100,35,1,3,'2018-02-27 00:00:00',350.0000,40,0,'ASM(40) d 12999 350.00,c 12001 350.00','2018-02-27 20:16:53',200,'2018-02-27 20:16:53',200),(101,39,1,0,'2016-10-03 00:00:00',4000.0000,0,0,'d 10999 _, c 12999 _','2018-02-28 17:18:33',211,'2018-02-28 17:18:33',211),(102,39,1,3,'2017-12-01 00:00:00',3650.0000,40,0,'ASM(40) d 12999 3650.00,c 12001 3650.00','2018-02-28 17:19:56',211,'2018-02-28 17:19:56',211),(103,39,1,3,'2017-12-01 00:00:00',350.0000,58,0,'ASM(58) d 12999 350.00,c 12001 350.00','2018-02-28 17:19:56',211,'2018-02-28 17:19:56',211),(104,40,1,0,'2018-01-01 00:00:00',3750.0000,0,0,'d 10999 _, c 12999 _','2018-02-28 21:16:01',200,'2018-02-28 21:16:01',200),(105,41,1,0,'2018-01-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2018-02-28 21:16:23',200,'2018-02-28 21:16:23',200),(106,42,1,0,'2018-02-01 00:00:00',4150.0000,0,0,'d 10999 _, c 12999 _','2018-02-28 21:16:42',200,'2018-02-28 21:16:42',200),(107,43,1,0,'2018-02-23 00:00:00',13050.0000,0,0,'d 10999 _, c 12999 _','2018-02-28 21:17:07',200,'2018-02-28 21:17:07',200),(108,40,1,2,'2018-02-28 00:00:00',3750.0000,62,0,'ASM(62) d 12999 3750.00,c 12001 3750.00','2018-02-28 21:17:34',200,'2018-02-28 21:17:34',200),(109,43,1,3,'2018-02-28 00:00:00',4000.0000,66,0,'ASM(66) d 12999 4000.00,c 12001 4000.00','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(110,43,1,3,'2018-02-28 00:00:00',350.0000,68,0,'ASM(68) d 12999 350.00,c 12001 350.00','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(111,43,1,3,'2018-02-28 00:00:00',4000.0000,67,0,'ASM(67) d 12999 4000.00,c 12001 4000.00','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(112,43,1,3,'2018-02-28 00:00:00',350.0000,69,0,'ASM(69) d 12999 350.00,c 12001 350.00','2018-02-28 21:17:49',200,'2018-02-28 21:17:49',200),(113,41,1,4,'2018-02-28 00:00:00',4150.0000,65,0,'ASM(65) d 12999 4150.00,c 12001 4150.00','2018-02-28 21:17:57',200,'2018-02-28 21:17:57',200),(114,42,1,4,'2018-02-28 00:00:00',4150.0000,63,0,'ASM(63) d 12999 4150.00,c 12001 4150.00','2018-02-28 21:17:57',200,'2018-02-28 21:17:57',200),(115,43,1,3,'2018-03-01 00:00:00',350.0000,83,0,'ASM(83) d 12999 350.00,c 12001 350.00','2018-05-30 20:04:19',200,'2018-05-30 20:04:19',200),(116,43,1,3,'2018-03-01 00:00:00',4000.0000,87,0,'ASM(87) d 12999 4000.00,c 12001 4000.00','2018-05-30 20:04:19',200,'2018-05-30 20:04:19',200),(117,44,1,0,'2018-05-20 00:00:00',13050.0000,0,0,'d 10999 _, c 12999 _','2018-05-30 20:08:28',200,'2018-05-30 20:08:28',200),(118,44,1,3,'2018-05-20 00:00:00',350.0000,84,0,'ASM(84) d 12999 350.00,c 12001 350.00','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(119,44,1,3,'2018-05-20 00:00:00',4000.0000,88,0,'ASM(88) d 12999 4000.00,c 12001 4000.00','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(120,44,1,3,'2018-05-20 00:00:00',350.0000,85,0,'ASM(85) d 12999 350.00,c 12001 350.00','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200),(121,44,1,3,'2018-05-20 00:00:00',4000.0000,89,0,'ASM(89) d 12999 4000.00,c 12001 4000.00','2018-05-30 20:09:39',200,'2018-05-30 20:09:39',200);
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
  `PRID` bigint(20) NOT NULL DEFAULT '0',
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
INSERT INTO `Rentable` VALUES (1,1,0,'309 Rexford',1,0,'0000-00-00 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0,''),(2,1,0,'309 1/2 Rexford',1,0,'0000-00-00 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0,''),(3,1,0,'311 Rexford',1,0,'0000-00-00 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0,''),(4,1,0,'311 1/2 Rexford',1,0,'0000-00-00 00:00:00','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0,'');
/*!40000 ALTER TABLE `Rentable` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentableLeaseStatus`
--

DROP TABLE IF EXISTS `RentableLeaseStatus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentableLeaseStatus` (
  `RLID` bigint(20) NOT NULL AUTO_INCREMENT,
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `LeaseStatus` smallint(6) NOT NULL DEFAULT '0',
  `DtStart` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtStop` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `Comment` varchar(2048) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RLID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableLeaseStatus`
--

LOCK TABLES `RentableLeaseStatus` WRITE;
/*!40000 ALTER TABLE `RentableLeaseStatus` DISABLE KEYS */;
/*!40000 ALTER TABLE `RentableLeaseStatus` ENABLE KEYS */;
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
INSERT INTO `RentableMarketRate` VALUES (1,1,1,3750.0000,'2014-01-01 00:00:00','9999-03-01 00:00:00','2018-02-23 09:03:24',0,'2017-11-28 03:44:18',0),(2,2,1,4000.0000,'2014-01-01 00:00:00','9999-05-01 00:00:00','2018-02-23 09:03:24',0,'2017-11-28 03:44:18',0),(3,3,1,4150.0000,'2014-01-01 00:00:00','9999-04-01 00:00:00','2018-02-23 09:03:24',0,'2017-11-28 03:44:18',0),(4,4,1,2500.0000,'2014-01-01 00:00:00','9999-01-01 00:00:00','2018-02-23 09:03:24',0,'2017-11-28 03:44:18',0);
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
  `RSPRefID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `RSPID` bigint(20) NOT NULL DEFAULT '0',
  `DtStart` date NOT NULL DEFAULT '1970-01-01',
  `DtStop` date NOT NULL DEFAULT '1970-01-01',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RSPRefID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableSpecialtyRef`
--

LOCK TABLES `RentableSpecialtyRef` WRITE;
/*!40000 ALTER TABLE `RentableSpecialtyRef` DISABLE KEYS */;
/*!40000 ALTER TABLE `RentableSpecialtyRef` ENABLE KEYS */;
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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
-- Table structure for table `RentableUseStatus`
--

DROP TABLE IF EXISTS `RentableUseStatus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentableUseStatus` (
  `RSID` bigint(20) NOT NULL AUTO_INCREMENT,
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `UseStatus` smallint(6) NOT NULL DEFAULT '0',
  `DtStart` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtStop` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `Comment` varchar(2048) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RSID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableUseStatus`
--

LOCK TABLES `RentableUseStatus` WRITE;
/*!40000 ALTER TABLE `RentableUseStatus` DISABLE KEYS */;
INSERT INTO `RentableUseStatus` VALUES (1,1,1,1,'2014-01-01 00:00:00','9999-01-01 00:00:00','','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0),(2,2,1,1,'2014-01-01 00:00:00','9999-01-01 00:00:00','','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0),(3,3,1,1,'2014-01-01 00:00:00','9999-01-01 00:00:00','','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0),(4,4,1,4,'2014-01-01 00:00:00','9999-01-01 00:00:00','','2017-11-28 03:52:45',0,'2017-11-28 03:52:45',0);
/*!40000 ALTER TABLE `RentableUseStatus` ENABLE KEYS */;
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
INSERT INTO `RentableUsers` VALUES (1,1,1,1,'2014-03-01','2018-02-01','2018-02-27 19:39:18',200,'2017-11-30 18:22:29',0),(2,1,1,2,'2014-03-01','2018-02-01','2018-02-27 19:39:27',200,'2017-11-30 18:23:03',0),(4,1,1,6,'2014-03-01','2018-02-01','2018-02-27 19:39:34',200,'2017-11-30 18:30:18',0),(5,2,1,3,'2016-10-01','2018-01-01','2018-02-23 09:00:52',0,'2017-11-30 18:33:01',0),(6,3,1,4,'2016-07-01','2018-07-01','2018-02-23 09:00:52',0,'2017-11-30 18:36:10',0),(7,3,1,5,'2016-07-01','2018-07-01','2018-02-23 09:00:52',0,'2017-11-30 18:36:30',0);
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
  `PRAID` bigint(20) NOT NULL DEFAULT '0',
  `ORIGIN` bigint(20) NOT NULL DEFAULT '0',
  `RATID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `NLID` bigint(20) NOT NULL DEFAULT '0',
  `DocumentDate` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
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
  `DesiredUsageStartDate` date NOT NULL DEFAULT '1970-01-01',
  `RentableTypePreference` bigint(20) NOT NULL DEFAULT '0',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `ApplicationReadyUID` bigint(20) NOT NULL DEFAULT '0',
  `ApplicationReadyDate` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `Approver1` bigint(20) NOT NULL DEFAULT '0',
  `DecisionDate1` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DeclineReason1` bigint(20) NOT NULL DEFAULT '0',
  `Approver2` bigint(20) NOT NULL DEFAULT '0',
  `DecisionDate2` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DeclineReason2` bigint(20) NOT NULL DEFAULT '0',
  `MoveInUID` bigint(20) NOT NULL DEFAULT '0',
  `MoveInDate` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `ActiveUID` bigint(20) NOT NULL DEFAULT '0',
  `ActiveDate` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `Outcome` bigint(20) NOT NULL DEFAULT '0',
  `NoticeToMoveUID` bigint(20) NOT NULL DEFAULT '0',
  `NoticeToMoveDate` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `NoticeToMoveReported` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `TerminatorUID` bigint(20) NOT NULL DEFAULT '0',
  `TerminationDate` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `LeaseTerminationReason` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RAID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreement`
--

LOCK TABLES `RentalAgreement` WRITE;
/*!40000 ALTER TABLE `RentalAgreement` DISABLE KEYS */;
INSERT INTO `RentalAgreement` VALUES (2,0,0,0,1,0,'1970-01-01 00:00:00','2014-03-01','2018-02-01','2014-03-01','2018-02-01','2014-03-01','2018-02-01','2014-03-01',0,0,2,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','','1970-01-01',0,0,0,'1970-01-01 00:00:00',0,'1970-01-01 00:00:00',0,0,'1970-01-01 00:00:00',0,0,'1970-01-01 00:00:00',0,'1970-01-01 00:00:00',0,0,'1970-01-01 00:00:00','1970-01-01 00:00:00',0,'1970-01-01 00:00:00',0,'2018-02-27 19:38:38',200,'2017-11-30 18:17:55',0),(3,0,0,0,1,0,'1970-01-01 00:00:00','2016-10-01','2018-12-31','2016-10-01','2018-12-31','2016-10-01','2018-12-31','2016-10-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','','1970-01-01',0,0,0,'1970-01-01 00:00:00',0,'1970-01-01 00:00:00',0,0,'1970-01-01 00:00:00',0,0,'1970-01-01 00:00:00',0,'1970-01-01 00:00:00',0,0,'1970-01-01 00:00:00','1970-01-01 00:00:00',0,'1970-01-01 00:00:00',0,'2018-02-27 19:42:16',200,'2017-11-30 18:29:25',0),(4,0,0,0,1,0,'1970-01-01 00:00:00','2016-07-01','2018-06-29','2016-07-01','2018-02-27','2016-07-01','2018-02-27','2016-07-01',0,0,2,'permitted to break lease 2/28/2018 without penalty',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,0,'1900-01-01','','','','1900-01-01','','1900-01-01','','1970-01-01',0,0,0,'1970-01-01 00:00:00',0,'1970-01-01 00:00:00',0,0,'1970-01-01 00:00:00',0,0,'1970-01-01 00:00:00',0,'1970-01-01 00:00:00',0,0,'1970-01-01 00:00:00','1970-01-01 00:00:00',0,'1970-01-01 00:00:00',0,'2018-05-30 19:17:48',200,'2017-11-30 18:33:59',0);
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
INSERT INTO `RentalAgreementPayors` VALUES (1,2,1,1,'2014-03-01','2018-02-01',0,'2018-02-27 19:39:12',200,'2017-11-30 18:21:00',0),(2,2,1,2,'2017-11-30','2018-03-01',0,'2018-02-23 09:02:11',0,'2017-11-30 18:21:57',0),(3,2,1,2,'2018-03-01','2018-03-01',0,'2018-02-23 09:02:11',0,'2017-11-30 18:28:09',0),(4,3,1,3,'2016-10-01','2018-12-31',0,'2018-02-27 19:42:13',200,'2017-11-30 18:32:33',0),(5,4,1,4,'2016-07-01','2018-02-01',0,'2018-02-23 09:02:11',0,'2017-11-30 18:35:19',0),(6,4,1,5,'2016-07-01','2018-07-01',0,'2018-02-23 09:02:11',0,'2017-11-30 18:35:41',0);
/*!40000 ALTER TABLE `RentalAgreementPayors` ENABLE KEYS */;
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
INSERT INTO `RentalAgreementRentables` VALUES (1,2,1,1,0,0,3750.0000,'2014-03-01','2018-02-01','2018-02-27 19:38:58',200,'2017-11-30 18:20:15',0),(2,3,1,2,0,0,4000.0000,'2016-10-01','2018-12-31','2018-02-27 21:16:18',200,'2017-11-30 18:32:13',0),(3,4,1,3,0,0,4150.0000,'2016-07-01','2018-07-01','2018-02-23 08:58:08',0,'2017-11-30 18:34:33',0);
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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
) ENGINE=InnoDB AUTO_INCREMENT=388 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SLString`
--

LOCK TABLES `SLString` WRITE;
/*!40000 ALTER TABLE `SLString` DISABLE KEYS */;
INSERT INTO `SLString` VALUES (1,1,1,'4Walls','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(2,1,1,'Apartment Finder Blue Book','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(3,1,1,'Apartment Guide','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(4,1,1,'Apartment Locator','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(5,1,1,'Apartment Map','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(6,1,1,'ApartmentFinder.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(7,1,1,'ApartmentGuide.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(8,1,1,'ApartmentGuyze.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(9,1,1,'ApartmentHomeLiving.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(10,1,1,'ApartmentLints.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(11,1,1,'ApartmentMag.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(12,1,1,'ApartmentMarketer.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(13,1,1,'ApartmentMatching.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(14,1,1,'ApartmentRatings.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(15,1,1,'ApartmentSearch.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(16,1,1,'ApartmentShowcase.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(17,1,1,'Apartments.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(18,1,1,'Apartments24-7.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(19,1,1,'ApartmentsNationwide.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(20,1,1,'ApartmentsPlus.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(21,1,1,'Brochure/Flyer','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(22,1,1,'CitySearch.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(23,1,1,'CollegeRentals.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(24,1,1,'CraigsList.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(25,1,1,'Current resident','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(26,1,1,'Direct Mail - Conventional','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(27,1,1,'Direct Mail - FullService','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(28,1,1,'Drive by','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(29,1,1,'EasyRent.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(30,1,1,'El Nacional','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(31,1,1,'EliteRenting.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(32,1,1,'For Rent Magazine','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(33,1,1,'ForRent.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(34,1,1,'Google Internet Program','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(35,1,1,'Google.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(36,1,1,'HotPads.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(37,1,1,'LivingChoices.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(38,1,1,'Local Line Rolloer','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(39,1,1,'Locator Service','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(40,1,1,'Move.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(41,1,1,'MoveForFree.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(42,1,1,'MyNewPlace.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(43,1,1,'Oklahoma Gazette','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(44,1,1,'Oodle.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(45,1,1,'Other','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(46,1,1,'Other','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(47,1,1,'Other OneSite property','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(48,1,1,'Other property','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(49,1,1,'Other publication','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(50,1,1,'Other site','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(51,1,1,'PMC-owned Website','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(52,1,1,'PeopleWithPets.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(53,1,1,'Preferred employer program','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(54,1,1,'Prior resident','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(55,1,1,'Property website','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(56,1,1,'Radio Advertising','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(57,1,1,'Referral companies/merchants','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(58,1,1,'Rent.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(59,1,1,'RentAndMove.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(60,1,1,'RentClicks.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(61,1,1,'RentJungle.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(62,1,1,'RentNet.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(63,1,1,'RentWiki.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(64,1,1,'Rentals.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(65,1,1,'Rentping.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(66,1,1,'Roomster.net','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(67,1,1,'Senior Living Magazine','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(68,1,1,'Site-owned website','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(69,1,1,'TV Advertising','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(70,1,1,'Tinker Take Off','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(71,1,1,'UMoveFree.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(72,1,1,'Unknown/Would not give','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(73,1,1,'Yahoo.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(74,1,1,'Yellow pages','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(75,1,2,'Criminal background','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(76,1,2,'No credit history','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(77,1,2,'No employment history','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(78,1,2,'No poor credit history','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(79,1,2,'No poor employment history','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(80,1,2,'No poor rental history','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(81,1,2,'No rental history','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(82,1,2,'Other','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(83,1,3,'Abandoned Apartment','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(84,1,3,'Acquired a pet','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(85,1,3,'Added a roommate','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(86,1,3,'Amenities lacking','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(87,1,3,'Bought condominium','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(88,1,3,'Bought home','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(89,1,3,'Bought townhome','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(90,1,3,'Changed jobs','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(91,1,3,'Closer to airport','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(92,1,3,'Closer to town/city','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(93,1,3,'Closer to work','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(94,1,3,'Corporate or short term lease only','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(95,1,3,'Death or illness','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(96,1,3,'Dissatisfied for another reason','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(97,1,3,'Divorce','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(98,1,3,'Employment transfer','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(99,1,3,'Evicted for another reason','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(100,1,3,'Evicted for criminal reasons','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(101,1,3,'Evicted for non-compliance with community policies','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(102,1,3,'Evicted for non-payment of rent','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(103,1,3,'Generally unhappy with property','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(104,1,3,'Getting married','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(105,1,3,'High utility costs','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(106,1,3,'Leaving/graduating school','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(107,1,3,'Lifestyle change for another reason','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(108,1,3,'Loss of employment from the PMC','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(109,1,3,'Lost a job','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(110,1,3,'Lost a roommate','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(111,1,3,'Marital status change','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(112,1,3,'Military transfer','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(113,1,3,'Money problems','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(114,1,3,'Moving closer to home','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(115,1,3,'Moving home','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(116,1,3,'No reason given','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(117,1,3,'Noise problem','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(118,1,3,'Non-renewal of lease','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(119,1,3,'Other','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(120,1,3,'Parking problems','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(121,1,3,'Personal reasons/concerns','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(122,1,3,'Property disaster','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(123,1,3,'Rental increase','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(124,1,3,'Rentin home','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(125,1,3,'Returning/going to school','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(126,1,3,'Road construction','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(127,1,3,'Selling/old house','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(128,1,3,'Skipped during eviction process','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(129,1,3,'Skipped without notice','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(130,1,4,'ADA accessible','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(131,1,4,'Amenities lacking','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(132,1,4,'Color palette','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(133,1,4,'Drive up appeal','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(134,1,4,'Furniture','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(135,1,4,'Lease term','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(136,1,4,'Location to employment','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(137,1,4,'Location to family','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(138,1,4,'Location to shopping and entertainment','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(139,1,4,'Meets square footage needs','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(140,1,4,'Personnel','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(141,1,4,'Pet allowances','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(142,1,4,'Point of lease e-commerce offers','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(143,1,4,'Priing','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(144,1,4,'Public transportation','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(145,1,4,'School district','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(146,1,4,'Special','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(147,1,5,'Amenities ^ Amenities lacking','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(148,1,5,'Amenities ^ Bedroom size','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(149,1,5,'Amenities ^ Color scheme','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(150,1,5,'Amenities ^ Competition has better amenities','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(151,1,5,'Amenities ^ Objection to floor plan','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(152,1,5,'Cost ^ Competition is less expensive','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(153,1,5,'Cost ^ No specials/concessions','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(154,1,5,'Cost ^ Too expensive','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(155,1,5,'Inactive ^ Inactive','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(156,1,5,'Location ^ Location','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(157,1,5,'Location ^ Road construction','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(158,1,5,'Location ^ Too close to highway','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(159,1,5,'Not available ^ Unit/floor plan not available','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(160,1,5,'Not interested ^ Bought/rented house instead','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(161,1,5,'Not interested ^ Changed their mind','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(162,1,5,'Not interested ^ Not interested','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(163,1,5,'Not qualified ^ Credit rating below standard','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(164,1,5,'Not qualified ^ Criminal history not allowed','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(165,1,5,'Not qualified ^ Does not meet property criteria','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(166,1,5,'Not qualified ^ Oversized/unallowable pet','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(167,1,5,'Not qualified ^ Rental history not allowed','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(168,1,5,'Not qualified ^ Roommate/spouse unqualified','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(169,1,5,'Not qualified ^ Too many occupants','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(170,1,2,'Application declined','2018-06-30 00:53:45',0,'2018-06-30 00:53:45',0),(171,1,6,'Application was declined','2018-07-23 16:14:19',0,'2018-07-23 16:14:19',-99998),(172,1,6,'Rental Agreement was updated','2018-07-23 16:14:19',0,'2018-07-23 16:14:19',-99998),(175,1,7,'Accountants','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(176,1,7,'Advertising/Public Relations','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(177,1,7,'Aerospace','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(178,1,7,'Agribusiness','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(179,1,7,'Agricultural Services & Products','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(180,1,7,'Agriculture','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(181,1,7,'Air Transport','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(182,1,7,'Air Transport Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(183,1,7,'Airlines','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(184,1,7,'Alcoholic Beverages','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(185,1,7,'Alternative Energy Production & Services','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(186,1,7,'Architectural Services','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(187,1,7,'Attorneys/Law Firms','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(188,1,7,'Auto Dealers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(189,1,7,'Auto Dealers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(190,1,7,'Auto Manufacturers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(191,1,7,'Automotive','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(192,1,7,'Banking','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(193,1,7,'Banks','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(194,1,7,'Banks','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(195,1,7,'Bars & Restaurants','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(196,1,7,'Beer','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(197,1,7,'Books','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(198,1,7,'Broadcasters','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(199,1,7,'Builders/General Contractors','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(200,1,7,'Builders/Residential','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(201,1,7,'Building Materials & Equipment','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(202,1,7,'Building Trade Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(203,1,7,'Business Associations','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(204,1,7,'Business Services','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(205,1,7,'Cable & Satellite TV Production & Distribution','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(206,1,7,'Candidate Committees','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(207,1,7,'Candidate Committees','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(208,1,7,'Candidate Committees','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(209,1,7,'Car Dealers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(210,1,7,'Car Dealers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(211,1,7,'Car Manufacturers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(212,1,7,'Casinos / Gambling','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(213,1,7,'Cattle Ranchers/Livestock','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(214,1,7,'Chemical & Related Manufacturing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(215,1,7,'Chiropractors','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(216,1,7,'Civil Servants/Public Officials','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(217,1,7,'Clergy & Religious Organizations','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(218,1,7,'Clothing Manufacturing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(219,1,7,'Coal Mining','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(220,1,7,'Colleges','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(221,1,7,'Commercial Banks','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(222,1,7,'Commercial TV & Radio Stations','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(223,1,7,'Communications/Electronics','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(224,1,7,'Computer Software','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(225,1,7,'Conservative/Republican','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(226,1,7,'Construction','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(227,1,7,'Construction Services','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(228,1,7,'Construction Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(229,1,7,'Credit Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(230,1,7,'Crop Production & Basic Processing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(231,1,7,'Cruise Lines','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(232,1,7,'Cruise Ships & Lines','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(233,1,7,'Dairy','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(234,1,7,'Defense','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(235,1,7,'Defense Aerospace','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(236,1,7,'Defense Electronics','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(237,1,7,'Defense/Foreign Policy Advocates','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(238,1,7,'Democratic Candidate Committees','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(239,1,7,'Democratic Leadership PACs','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(240,1,7,'Democratic/Liberal','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(241,1,7,'Dentists','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(242,1,7,'Doctors & Other Health Professionals','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(243,1,7,'Drug Manufacturers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(244,1,7,'Education','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(245,1,7,'Electric Utilities','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(246,1,7,'Electronics Manufacturing & Equipment','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(247,1,7,'Electronics','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(248,1,7,'Energy & Natural Resources','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(249,1,7,'Entertainment Industry','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(250,1,7,'Environment','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(251,1,7,'Farm Bureaus','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(252,1,7,'Farming','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(253,1,7,'Finance / Credit Companies','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(254,1,7,'Finance','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(255,1,7,'Food & Beverage','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(256,1,7,'Food Processing & Sales','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(257,1,7,'Food Products Manufacturing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(258,1,7,'Food Stores','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(259,1,7,'For-profit Education','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(260,1,7,'For-profit Prisons','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(261,1,7,'Foreign & Defense Policy','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(262,1,7,'Forestry & Forest Products','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(263,1,7,'Foundations','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(264,1,7,'Funeral Services','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(265,1,7,'Gambling & Casinos','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(266,1,7,'Gambling','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(267,1,7,'Garbage Collection/Waste Management','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(268,1,7,'Gas & Oil','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(269,1,7,'Gay & Lesbian Rights & Issues','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(270,1,7,'General Contractors','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(271,1,7,'Government Employee Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(272,1,7,'Government Employees','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(273,1,7,'Gun Control','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(274,1,7,'Gun Rights','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(275,1,7,'Health','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(276,1,7,'Health Professionals','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(277,1,7,'Health Services/HMOs','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(278,1,7,'Hedge Funds','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(279,1,7,'HMOs & Health Care Services','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(280,1,7,'Home Builders','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(281,1,7,'Hospitals & Nursing Homes','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(282,1,7,'Hotels','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(283,1,7,'Human Rights','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(284,1,7,'Ideological/Single-Issue','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(285,1,7,'Indian Gaming','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(286,1,7,'Industrial Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(287,1,7,'Insurance','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(288,1,7,'Internet','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(289,1,7,'Israel Policy','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(290,1,7,'Labor','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(291,1,7,'Lawyers & Lobbyists','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(292,1,7,'Lawyers / Law Firms','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(293,1,7,'Leadership PACs','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(294,1,7,'Liberal/Democratic','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(295,1,7,'Liquor','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(296,1,7,'Livestock','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(297,1,7,'Lobbyists','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(298,1,7,'Lodging / Tourism','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(299,1,7,'Logging','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(300,1,7,'Manufacturing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(301,1,7,'Marine Transport','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(302,1,7,'Meat processing & products','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(303,1,7,'Medical Supplies','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(304,1,7,'Mining','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(305,1,7,'Misc Business','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(306,1,7,'Misc Finance','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(307,1,7,'Misc Manufacturing & Distributing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(308,1,7,'Misc Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(309,1,7,'Miscellaneous Defense','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(310,1,7,'Miscellaneous Services','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(311,1,7,'Mortgage Bankers & Brokers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(312,1,7,'Motion Picture Production & Distribution','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(313,1,7,'Music Production','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(314,1,7,'Natural Gas Pipelines','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(315,1,7,'Newspaper','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(316,1,7,'Non-profits','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(317,1,7,'Nurses','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(318,1,7,'Nursing Homes/Hospitals','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(319,1,7,'Nutritional & Dietary Supplements','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(320,1,7,'Oil & Gas','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(321,1,7,'Payday Lenders','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(322,1,7,'Pharmaceutical Manufacturing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(323,1,7,'Pharmaceuticals / Health Products','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(324,1,7,'Phone Companies','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(325,1,7,'Physicians & Other Health Professionals','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(326,1,7,'Postal Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(327,1,7,'Poultry & Eggs','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(328,1,7,'Power Utilities','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(329,1,7,'Printing & Publishing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(330,1,7,'Private Equity & Investment Firms','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(331,1,7,'Pro-Israel','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(332,1,7,'Professional Sports','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(333,1,7,'Progressive/Democratic','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(334,1,7,'Public Employees','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(335,1,7,'Public Sector Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(336,1,7,'Publishing & Printing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(337,1,7,'Radio/TV Stations','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(338,1,7,'Railroads','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(339,1,7,'Real Estate','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(340,1,7,'Record Companies/Singers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(341,1,7,'Recorded Music & Music Production','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(342,1,7,'Recreation / Live Entertainment','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(343,1,7,'Religious Organizations/Clergy','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(344,1,7,'Republican Candidate Committees','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(345,1,7,'Republican Leadership PACs','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(346,1,7,'Republican/Conservative','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(347,1,7,'Residential Construction','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(348,1,7,'Restaurants & Drinking Establishments','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(349,1,7,'Retail Sales','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(350,1,7,'Retired','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(351,1,7,'Savings & Loans','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(352,1,7,'Schools/Education','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(353,1,7,'Sea Transport','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(354,1,7,'Securities & Investment','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(355,1,7,'Special Trade Contractors','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(356,1,7,'Sports','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(357,1,7,'Steel Production','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(358,1,7,'Stock Brokers/Investment Industry','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(359,1,7,'Student Loan Companies','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(360,1,7,'Sugar Cane & Sugar Beets','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(361,1,7,'Teachers Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(362,1,7,'Teachers/Education','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(363,1,7,'Telecom Services & Equipment','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(364,1,7,'Telephone Utilities','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(365,1,7,'Textiles','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(366,1,7,'Timber','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(367,1,7,'Tobacco','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(368,1,7,'Transportation','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(369,1,7,'Transportation Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(370,1,7,'Trash Collection/Waste Management','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(371,1,7,'Trucking','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(372,1,7,'TV / Movies / Music','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(373,1,7,'TV Production','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(374,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(375,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(376,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(377,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(378,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(379,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(380,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(381,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(382,1,7,'Universities','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(383,1,7,'Vegetables & Fruits','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(384,1,7,'Venture Capital','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(385,1,7,'Waste Management','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(386,1,7,'Wine','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(387,1,7,'Women\'s Issues','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `StringList`
--

LOCK TABLES `StringList` WRITE;
/*!40000 ALTER TABLE `StringList` DISABLE KEYS */;
INSERT INTO `StringList` VALUES (1,1,'HowFound','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(2,1,'ApplDeny','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(3,1,'WhyLeaving','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(4,1,'WhyChoose','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(5,1,'ProspectLost','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(6,1,'RollerMsgs','2018-07-23 16:14:19',-99998,'2018-07-23 16:14:19',-99998),(7,1,'Industries','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0);
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
-- Table structure for table `TBind`
--

DROP TABLE IF EXISTS `TBind`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `TBind` (
  `TBID` bigint(20) NOT NULL AUTO_INCREMENT,
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `SourceElemType` bigint(20) NOT NULL DEFAULT '0',
  `SourceElemID` bigint(20) NOT NULL DEFAULT '0',
  `AssocElemType` bigint(20) NOT NULL DEFAULT '0',
  `AssocElemID` bigint(20) NOT NULL DEFAULT '0',
  `DtStart` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtStop` datetime NOT NULL DEFAULT '2066-01-01 00:00:00',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`TBID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TBind`
--

LOCK TABLES `TBind` WRITE;
/*!40000 ALTER TABLE `TBind` DISABLE KEYS */;
/*!40000 ALTER TABLE `TBind` ENABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=120 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TWS`
--

LOCK TABLES `TWS` WRITE;
/*!40000 ALTER TABLE `TWS` DISABLE KEYS */;
INSERT INTO `TWS` VALUES (2,'CreateAssessmentInstances','','CreateAssessmentInstances','2018-07-18 00:09:48','ip-172-31-51-141.ec2.internal',4,'2018-07-17 00:09:48','2018-07-17 00:09:48','2017-11-27 21:24:27','2018-07-17 00:09:47'),(3,'CleanRARBalanceCache','','CleanRARBalanceCache','2018-05-02 22:03:40','ip-172-31-51-141.ec2.internal',4,'2018-05-02 21:58:40','2018-05-02 21:58:40','2017-11-30 17:39:57','2018-05-02 21:58:39'),(4,'CleanSecDepBalanceCache','','CleanSecDepBalanceCache','2018-05-02 22:03:50','ip-172-31-51-141.ec2.internal',4,'2018-05-02 21:58:50','2018-05-02 21:58:50','2017-11-30 17:39:57','2018-05-02 21:58:49'),(5,'CleanAcctSliceCache','','CleanAcctSliceCache','2018-05-02 22:03:50','ip-172-31-51-141.ec2.internal',4,'2018-05-02 21:58:50','2018-05-02 21:58:50','2017-11-30 17:39:57','2018-05-02 21:58:49'),(6,'CleanARSliceCache','','CleanARSliceCache','2018-05-02 22:03:50','ip-172-31-51-141.ec2.internal',4,'2018-05-02 21:58:50','2018-05-02 21:58:50','2017-11-30 17:39:57','2018-05-02 21:58:49'),(29,'ManualTaskBot','','ManualTaskBot','2018-07-18 00:09:38','ip-172-31-51-141.ec2.internal',4,'2018-07-17 00:09:38','2018-07-17 00:09:38','2018-05-02 22:01:22','2018-07-17 00:09:37'),(30,'RARBcacheBot','','RARBcacheBot','2018-07-17 00:17:28','ip-172-31-51-141.ec2.internal',4,'2018-07-17 00:12:28','2018-07-17 00:12:28','2018-05-02 22:01:22','2018-07-17 00:12:27'),(31,'SecDepCacheBot','','SecDepCacheBot','2018-07-17 00:17:28','ip-172-31-51-141.ec2.internal',4,'2018-07-17 00:12:28','2018-07-17 00:12:28','2018-05-02 22:01:22','2018-07-17 00:12:27'),(32,'AcctSliceCacheBot','','AcctSliceCacheBot','2018-07-17 00:17:28','ip-172-31-51-141.ec2.internal',4,'2018-07-17 00:12:28','2018-07-17 00:12:28','2018-05-02 22:01:22','2018-07-17 00:12:27'),(33,'ARSliceCacheBot','','ARSliceCacheBot','2018-07-17 00:17:18','ip-172-31-51-141.ec2.internal',4,'2018-07-17 00:12:18','2018-07-17 00:12:18','2018-05-02 22:01:22','2018-07-17 00:12:17'),(34,'ARSliceCacheBot','','ARSliceCacheBot','2018-08-06 20:16:02','ip-172-31-60-42.ec2.internal',4,'2018-08-06 20:11:02','2018-08-06 20:11:02','2018-08-01 21:23:11','2018-08-06 20:11:01'),(35,'CreateAssessmentInstances','','CreateAssessmentInstances','2018-08-06 21:23:54','ip-172-31-60-42.ec2.internal',4,'2018-08-05 21:23:54','2018-08-05 21:23:54','2018-08-01 21:23:11','2018-08-05 21:23:54'),(36,'ManualTaskBot','','ManualTaskBot','2018-08-06 21:23:54','ip-172-31-60-42.ec2.internal',4,'2018-08-05 21:23:54','2018-08-05 21:23:54','2018-08-01 21:23:11','2018-08-05 21:23:54'),(37,'RARBcacheBot','','RARBcacheBot','2018-08-06 20:15:52','ip-172-31-60-42.ec2.internal',4,'2018-08-06 20:10:52','2018-08-06 20:10:52','2018-08-01 21:23:11','2018-08-06 20:10:51'),(38,'SecDepCacheBot','','SecDepCacheBot','2018-08-06 20:15:12','ip-172-31-60-42.ec2.internal',4,'2018-08-06 20:10:12','2018-08-06 20:10:12','2018-08-01 21:23:11','2018-08-06 20:10:11'),(39,'AcctSliceCacheBot','','AcctSliceCacheBot','2018-08-06 20:13:22','ip-172-31-60-42.ec2.internal',4,'2018-08-06 20:08:22','2018-08-06 20:08:22','2018-08-01 21:23:11','2018-08-06 20:08:21'),(40,'TLInstanceBot','','TLInstanceBot','2018-08-03 00:24:39','ip-172-31-57-175.ec2.internal',4,'2018-08-02 00:24:39','2018-08-02 00:24:39','2018-08-02 00:24:13','2018-08-02 00:24:39'),(41,'ManualTaskBot','','ManualTaskBot','2018-08-03 00:24:39','ip-172-31-57-175.ec2.internal',4,'2018-08-02 00:24:39','2018-08-02 00:24:39','2018-08-02 00:24:13','2018-08-02 00:24:39'),(42,'RARBcacheBot','','RARBcacheBot','2018-08-02 00:29:39','ip-172-31-57-175.ec2.internal',4,'2018-08-02 00:24:39','2018-08-02 00:24:39','2018-08-02 00:24:13','2018-08-02 00:24:39'),(43,'SecDepCacheBot','','SecDepCacheBot','2018-08-02 00:29:39','ip-172-31-57-175.ec2.internal',4,'2018-08-02 00:24:39','2018-08-02 00:24:39','2018-08-02 00:24:13','2018-08-02 00:24:39'),(44,'AcctSliceCacheBot','','AcctSliceCacheBot','2018-08-02 00:29:39','ip-172-31-57-175.ec2.internal',4,'2018-08-02 00:24:39','2018-08-02 00:24:39','2018-08-02 00:24:13','2018-08-02 00:24:39'),(45,'AssessmentBot','','AssessmentBot','2018-08-03 00:24:39','ip-172-31-57-175.ec2.internal',4,'2018-08-02 00:24:39','2018-08-02 00:24:39','2018-08-02 00:24:13','2018-08-02 00:24:39'),(46,'ARSliceCacheBot','','ARSliceCacheBot','2018-08-02 00:29:39','ip-172-31-57-175.ec2.internal',4,'2018-08-02 00:24:39','2018-08-02 00:24:39','2018-08-02 00:24:13','2018-08-02 00:24:39'),(47,'TLReportBot','','TLReportBot','2018-08-02 00:26:39','ip-172-31-57-175.ec2.internal',4,'2018-08-02 00:24:39','2018-08-02 00:24:39','2018-08-02 00:24:13','2018-08-02 00:24:39'),(48,'ARSliceCacheBot','','ARSliceCacheBot','2018-08-02 02:04:48','ip-172-31-50-31.ec2.internal',4,'2018-08-02 01:59:48','2018-08-02 01:59:48','2018-08-02 01:08:59','2018-08-02 01:59:48'),(49,'TLReportBot','','TLReportBot','2018-08-02 02:05:31','ip-172-31-50-31.ec2.internal',4,'2018-08-02 02:03:31','2018-08-02 02:03:31','2018-08-02 01:08:59','2018-08-02 02:03:31'),(50,'TLInstanceBot','','TLInstanceBot','2018-08-03 01:09:09','ip-172-31-50-31.ec2.internal',4,'2018-08-02 01:09:09','2018-08-02 01:09:09','2018-08-02 01:08:59','2018-08-02 01:09:09'),(51,'ManualTaskBot','','ManualTaskBot','2018-08-03 01:09:09','ip-172-31-50-31.ec2.internal',4,'2018-08-02 01:09:09','2018-08-02 01:09:09','2018-08-02 01:08:59','2018-08-02 01:09:09'),(52,'RARBcacheBot','','RARBcacheBot','2018-08-02 02:04:48','ip-172-31-50-31.ec2.internal',4,'2018-08-02 01:59:48','2018-08-02 01:59:48','2018-08-02 01:08:59','2018-08-02 01:59:48'),(53,'AcctSliceCacheBot','','AcctSliceCacheBot','2018-08-02 02:04:48','ip-172-31-50-31.ec2.internal',4,'2018-08-02 01:59:48','2018-08-02 01:59:48','2018-08-02 01:08:59','2018-08-02 01:59:48'),(54,'AssessmentBot','','AssessmentBot','2018-08-03 01:09:09','ip-172-31-50-31.ec2.internal',4,'2018-08-02 01:09:09','2018-08-02 01:09:10','2018-08-02 01:08:59','2018-08-02 01:09:09'),(55,'SecDepCacheBot','','SecDepCacheBot','2018-08-02 02:04:48','ip-172-31-50-31.ec2.internal',4,'2018-08-02 01:59:48','2018-08-02 01:59:48','2018-08-02 01:08:59','2018-08-02 01:59:48'),(56,'SecDepCacheBot','','SecDepCacheBot','2018-08-02 18:24:50','ip-172-31-56-225.ec2.internal',4,'2018-08-02 18:19:50','2018-08-02 18:19:50','2018-08-02 18:14:30','2018-08-02 18:19:49'),(57,'ARSliceCacheBot','','ARSliceCacheBot','2018-08-02 18:24:50','ip-172-31-56-225.ec2.internal',4,'2018-08-02 18:19:50','2018-08-02 18:19:50','2018-08-02 18:14:30','2018-08-02 18:19:49'),(58,'TLReportBot','','TLReportBot','2018-08-02 18:21:50','ip-172-31-56-225.ec2.internal',4,'2018-08-02 18:19:50','2018-08-02 18:19:50','2018-08-02 18:14:30','2018-08-02 18:19:49'),(59,'TLInstanceBot','','TLInstanceBot','2018-08-03 18:14:40','ip-172-31-56-225.ec2.internal',4,'2018-08-02 18:14:40','2018-08-02 18:14:40','2018-08-02 18:14:30','2018-08-02 18:14:40'),(60,'ManualTaskBot','','ManualTaskBot','2018-08-03 18:14:40','ip-172-31-56-225.ec2.internal',4,'2018-08-02 18:14:40','2018-08-02 18:14:40','2018-08-02 18:14:30','2018-08-02 18:14:40'),(61,'AssessmentBot','','AssessmentBot','2018-08-03 18:14:40','ip-172-31-56-225.ec2.internal',4,'2018-08-02 18:14:40','2018-08-02 18:14:40','2018-08-02 18:14:30','2018-08-02 18:14:40'),(62,'RARBcacheBot','','RARBcacheBot','2018-08-02 18:24:50','ip-172-31-56-225.ec2.internal',4,'2018-08-02 18:19:50','2018-08-02 18:19:50','2018-08-02 18:14:30','2018-08-02 18:19:49'),(63,'AcctSliceCacheBot','','AcctSliceCacheBot','2018-08-02 18:24:50','ip-172-31-56-225.ec2.internal',4,'2018-08-02 18:19:50','2018-08-02 18:19:50','2018-08-02 18:14:30','2018-08-02 18:19:49'),(64,'AcctSliceCacheBot','','AcctSliceCacheBot','2018-08-06 17:48:09','ip-172-31-53-195.ec2.internal',4,'2018-08-06 17:43:09','2018-08-06 17:43:09','2018-08-03 03:56:23','2018-08-06 17:43:09'),(65,'TLReportBot','','TLReportBot','2018-08-06 17:47:29','ip-172-31-53-195.ec2.internal',4,'2018-08-06 17:45:29','2018-08-06 17:45:29','2018-08-03 03:56:23','2018-08-06 17:45:29'),(66,'AssessmentBot','','AssessmentBot','2018-08-07 03:56:37','ip-172-31-53-195.ec2.internal',4,'2018-08-06 03:56:37','2018-08-06 03:56:37','2018-08-03 03:56:23','2018-08-06 03:56:37'),(67,'SecDepCacheBot','','SecDepCacheBot','2018-08-06 17:47:59','ip-172-31-53-195.ec2.internal',4,'2018-08-06 17:42:59','2018-08-06 17:42:59','2018-08-03 03:56:23','2018-08-06 17:42:59'),(68,'TLInstanceBot','','TLInstanceBot','2018-08-07 03:56:37','ip-172-31-53-195.ec2.internal',4,'2018-08-06 03:56:37','2018-08-06 03:56:37','2018-08-03 03:56:23','2018-08-06 03:56:37'),(69,'ManualTaskBot','','ManualTaskBot','2018-08-07 03:56:37','ip-172-31-53-195.ec2.internal',4,'2018-08-06 03:56:37','2018-08-06 03:56:37','2018-08-03 03:56:23','2018-08-06 03:56:37'),(70,'RARBcacheBot','','RARBcacheBot','2018-08-06 17:47:39','ip-172-31-53-195.ec2.internal',4,'2018-08-06 17:42:39','2018-08-06 17:42:39','2018-08-03 03:56:23','2018-08-06 17:42:39'),(71,'ARSliceCacheBot','','ARSliceCacheBot','2018-08-06 17:51:39','ip-172-31-53-195.ec2.internal',4,'2018-08-06 17:46:39','2018-08-06 17:46:39','2018-08-03 03:56:23','2018-08-06 17:46:39'),(72,'SecDepCacheBot','','SecDepCacheBot','2018-08-06 20:42:24','ip-172-31-3-229.ec2.internal',4,'2018-08-06 20:37:24','2018-08-06 20:37:24','2018-08-06 20:37:13','2018-08-06 20:37:23'),(73,'ARSliceCacheBot','','ARSliceCacheBot','2018-08-06 20:42:24','ip-172-31-3-229.ec2.internal',4,'2018-08-06 20:37:24','2018-08-06 20:37:24','2018-08-06 20:37:13','2018-08-06 20:37:23'),(74,'TLReportBot','','TLReportBot','2018-08-06 20:41:34','ip-172-31-3-229.ec2.internal',4,'2018-08-06 20:39:34','2018-08-06 20:39:34','2018-08-06 20:37:13','2018-08-06 20:39:33'),(75,'TLInstanceBot','','TLInstanceBot','2018-08-07 20:37:24','ip-172-31-3-229.ec2.internal',4,'2018-08-06 20:37:24','2018-08-06 20:37:24','2018-08-06 20:37:13','2018-08-06 20:37:23'),(76,'ManualTaskBot','','ManualTaskBot','2018-08-07 20:37:24','ip-172-31-3-229.ec2.internal',4,'2018-08-06 20:37:24','2018-08-06 20:37:24','2018-08-06 20:37:13','2018-08-06 20:37:23'),(77,'AssessmentBot','','AssessmentBot','2018-08-07 20:37:24','ip-172-31-3-229.ec2.internal',4,'2018-08-06 20:37:24','2018-08-06 20:37:24','2018-08-06 20:37:13','2018-08-06 20:37:23'),(78,'AcctSliceCacheBot','','AcctSliceCacheBot','2018-08-06 20:42:24','ip-172-31-3-229.ec2.internal',4,'2018-08-06 20:37:24','2018-08-06 20:37:24','2018-08-06 20:37:13','2018-08-06 20:37:23'),(79,'RARBcacheBot','','RARBcacheBot','2018-08-06 20:42:24','ip-172-31-3-229.ec2.internal',4,'2018-08-06 20:37:24','2018-08-06 20:37:24','2018-08-06 20:37:13','2018-08-06 20:37:23'),(80,'RARBcacheBot','','RARBcacheBot','2018-08-06 23:59:43','ip-172-31-59-125.ec2.internal',4,'2018-08-06 23:54:43','2018-08-06 23:54:43','2018-08-06 23:39:02','2018-08-06 23:54:42'),(81,'SecDepCacheBot','','SecDepCacheBot','2018-08-06 23:59:43','ip-172-31-59-125.ec2.internal',4,'2018-08-06 23:54:43','2018-08-06 23:54:43','2018-08-06 23:39:02','2018-08-06 23:54:42'),(82,'ARSliceCacheBot','','ARSliceCacheBot','2018-08-06 23:59:43','ip-172-31-59-125.ec2.internal',4,'2018-08-06 23:54:43','2018-08-06 23:54:43','2018-08-06 23:39:02','2018-08-06 23:54:42'),(83,'TLInstanceBot','','TLInstanceBot','2018-08-07 23:39:13','ip-172-31-59-125.ec2.internal',4,'2018-08-06 23:39:13','2018-08-06 23:39:13','2018-08-06 23:39:02','2018-08-06 23:39:12'),(84,'ManualTaskBot','','ManualTaskBot','2018-08-07 23:39:13','ip-172-31-59-125.ec2.internal',4,'2018-08-06 23:39:13','2018-08-06 23:39:13','2018-08-06 23:39:02','2018-08-06 23:39:12'),(85,'AssessmentBot','','AssessmentBot','2018-08-07 23:39:13','ip-172-31-59-125.ec2.internal',4,'2018-08-06 23:39:13','2018-08-06 23:39:13','2018-08-06 23:39:02','2018-08-06 23:39:12'),(86,'AcctSliceCacheBot','','AcctSliceCacheBot','2018-08-06 23:59:43','ip-172-31-59-125.ec2.internal',4,'2018-08-06 23:54:43','2018-08-06 23:54:43','2018-08-06 23:39:02','2018-08-06 23:54:42'),(87,'TLReportBot','','TLReportBot','2018-08-06 23:58:33','ip-172-31-59-125.ec2.internal',4,'2018-08-06 23:56:33','2018-08-06 23:56:33','2018-08-06 23:39:02','2018-08-06 23:56:32'),(88,'ARSliceCacheBot','','ARSliceCacheBot','2018-08-07 00:19:47','ip-172-31-62-139.ec2.internal',4,'2018-08-07 00:14:47','2018-08-07 00:14:47','2018-08-07 00:04:37','2018-08-07 00:14:47'),(89,'TLReportBot','','TLReportBot','2018-08-07 00:16:47','ip-172-31-62-139.ec2.internal',4,'2018-08-07 00:14:47','2018-08-07 00:14:47','2018-08-07 00:04:37','2018-08-07 00:14:47'),(90,'ManualTaskBot','','ManualTaskBot','2018-08-08 00:04:47','ip-172-31-62-139.ec2.internal',4,'2018-08-07 00:04:47','2018-08-07 00:04:47','2018-08-07 00:04:37','2018-08-07 00:04:47'),(91,'AssessmentBot','','AssessmentBot','2018-08-08 00:04:47','ip-172-31-62-139.ec2.internal',4,'2018-08-07 00:04:47','2018-08-07 00:04:47','2018-08-07 00:04:37','2018-08-07 00:04:47'),(92,'RARBcacheBot','','RARBcacheBot','2018-08-07 00:19:47','ip-172-31-62-139.ec2.internal',4,'2018-08-07 00:14:47','2018-08-07 00:14:47','2018-08-07 00:04:37','2018-08-07 00:14:47'),(93,'SecDepCacheBot','','SecDepCacheBot','2018-08-07 00:19:47','ip-172-31-62-139.ec2.internal',4,'2018-08-07 00:14:47','2018-08-07 00:14:47','2018-08-07 00:04:37','2018-08-07 00:14:47'),(94,'AcctSliceCacheBot','','AcctSliceCacheBot','2018-08-07 00:19:47','ip-172-31-62-139.ec2.internal',4,'2018-08-07 00:14:47','2018-08-07 00:14:47','2018-08-07 00:04:37','2018-08-07 00:14:47'),(95,'TLInstanceBot','','TLInstanceBot','2018-08-08 00:04:47','ip-172-31-62-139.ec2.internal',4,'2018-08-07 00:04:47','2018-08-07 00:04:47','2018-08-07 00:04:37','2018-08-07 00:04:47'),(96,'AssessmentBot','','AssessmentBot','2018-08-08 00:22:30','ip-172-31-53-126.ec2.internal',4,'2018-08-07 00:22:30','2018-08-07 00:22:30','2018-08-07 00:22:20','2018-08-07 00:22:30'),(97,'ARSliceCacheBot','','ARSliceCacheBot','2018-08-07 06:03:05','ip-172-31-53-126.ec2.internal',4,'2018-08-07 05:58:05','2018-08-07 05:58:05','2018-08-07 00:22:20','2018-08-07 05:58:04'),(98,'ManualTaskBot','','ManualTaskBot','2018-08-08 00:22:30','ip-172-31-53-126.ec2.internal',4,'2018-08-07 00:22:30','2018-08-07 00:22:30','2018-08-07 00:22:20','2018-08-07 00:22:30'),(99,'TLReportBot','','TLReportBot','2018-08-07 06:01:15','ip-172-31-53-126.ec2.internal',4,'2018-08-07 05:59:15','2018-08-07 05:59:15','2018-08-07 00:22:20','2018-08-07 05:59:14'),(100,'TLInstanceBot','','TLInstanceBot','2018-08-08 00:22:30','ip-172-31-53-126.ec2.internal',4,'2018-08-07 00:22:30','2018-08-07 00:22:30','2018-08-07 00:22:20','2018-08-07 00:22:30'),(101,'RARBcacheBot','','RARBcacheBot','2018-08-07 06:02:45','ip-172-31-53-126.ec2.internal',4,'2018-08-07 05:57:45','2018-08-07 05:57:45','2018-08-07 00:22:20','2018-08-07 05:57:44'),(102,'SecDepCacheBot','','SecDepCacheBot','2018-08-07 06:02:35','ip-172-31-53-126.ec2.internal',4,'2018-08-07 05:57:35','2018-08-07 05:57:35','2018-08-07 00:22:20','2018-08-07 05:57:34'),(103,'AcctSliceCacheBot','','AcctSliceCacheBot','2018-08-07 06:02:25','ip-172-31-53-126.ec2.internal',4,'2018-08-07 05:57:25','2018-08-07 05:57:25','2018-08-07 00:22:20','2018-08-07 05:57:24'),(104,'RARBcacheBot','','RARBcacheBot','2018-08-07 08:41:35','ip-172-31-62-72.ec2.internal',4,'2018-08-07 08:36:35','2018-08-07 08:36:35','2018-08-07 06:34:13','2018-08-07 08:36:35'),(105,'AcctSliceCacheBot','','AcctSliceCacheBot','2018-08-07 08:41:35','ip-172-31-62-72.ec2.internal',4,'2018-08-07 08:36:35','2018-08-07 08:36:35','2018-08-07 06:34:13','2018-08-07 08:36:35'),(106,'AssessmentBot','','AssessmentBot','2018-08-08 06:34:24','ip-172-31-62-72.ec2.internal',4,'2018-08-07 06:34:24','2018-08-07 06:34:24','2018-08-07 06:34:13','2018-08-07 06:34:23'),(107,'SecDepCacheBot','','SecDepCacheBot','2018-08-07 08:41:35','ip-172-31-62-72.ec2.internal',4,'2018-08-07 08:36:35','2018-08-07 08:36:35','2018-08-07 06:34:13','2018-08-07 08:36:35'),(108,'ARSliceCacheBot','','ARSliceCacheBot','2018-08-07 08:41:35','ip-172-31-62-72.ec2.internal',4,'2018-08-07 08:36:35','2018-08-07 08:36:35','2018-08-07 06:34:13','2018-08-07 08:36:35'),(109,'TLReportBot','','TLReportBot','2018-08-07 08:41:35','ip-172-31-62-72.ec2.internal',4,'2018-08-07 08:39:35','2018-08-07 08:39:35','2018-08-07 06:34:13','2018-08-07 08:39:35'),(110,'TLInstanceBot','','TLInstanceBot','2018-08-08 06:34:24','ip-172-31-62-72.ec2.internal',4,'2018-08-07 06:34:24','2018-08-07 06:34:24','2018-08-07 06:34:13','2018-08-07 06:34:23'),(111,'ManualTaskBot','','ManualTaskBot','2018-08-08 06:34:24','ip-172-31-62-72.ec2.internal',4,'2018-08-07 06:34:24','2018-08-07 06:34:24','2018-08-07 06:34:13','2018-08-07 06:34:23'),(112,'RARBcacheBot','','RARBcacheBot','2019-01-31 07:15:37','ip-172-31-61-67.ec2.internal',4,'2019-01-31 07:10:37','2019-01-31 07:10:37','2018-09-25 18:38:48','2019-01-31 07:10:36'),(113,'ARSliceCacheBot','','ARSliceCacheBot','2019-01-31 07:15:47','ip-172-31-61-67.ec2.internal',4,'2019-01-31 07:10:47','2019-01-31 07:10:47','2018-09-25 18:38:48','2019-01-31 07:10:46'),(114,'TLReportBot','','TLReportBot','2019-01-31 07:14:37','ip-172-31-61-67.ec2.internal',4,'2019-01-31 07:12:37','2019-01-31 07:12:37','2018-09-25 18:38:48','2019-01-31 07:12:36'),(115,'TLInstanceBot','','TLInstanceBot','2019-01-31 18:45:34','ip-172-31-61-67.ec2.internal',4,'2019-01-30 18:45:34','2019-01-30 18:45:34','2018-09-25 18:38:48','2019-01-30 18:45:33'),(116,'AssessmentBot','','AssessmentBot','2019-01-31 18:45:34','ip-172-31-61-67.ec2.internal',4,'2019-01-30 18:45:34','2019-01-30 18:45:34','2018-09-25 18:38:48','2019-01-30 18:45:33'),(117,'AcctSliceCacheBot','','AcctSliceCacheBot','2019-01-31 07:15:07','ip-172-31-61-67.ec2.internal',4,'2019-01-31 07:10:07','2019-01-31 07:10:07','2018-09-25 18:38:48','2019-01-31 07:10:06'),(118,'ManualTaskBot','','ManualTaskBot','2019-01-31 18:45:34','ip-172-31-61-67.ec2.internal',4,'2019-01-30 18:45:34','2019-01-30 18:45:34','2018-09-25 18:38:48','2019-01-30 18:45:33'),(119,'SecDepCacheBot','','SecDepCacheBot','2019-01-31 07:14:47','ip-172-31-61-67.ec2.internal',4,'2019-01-31 07:09:47','2019-01-31 07:09:47','2018-09-25 18:38:48','2019-01-31 07:09:46');
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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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
  `AlternateEmailAddress` varchar(100) NOT NULL DEFAULT '',
  `EligibleFutureUser` tinyint(1) NOT NULL DEFAULT '1',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `Industry` bigint(20) NOT NULL DEFAULT '0',
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
INSERT INTO `User` VALUES (1,1,0,'1900-01-01','','','','','',1,0,0,0,'2017-11-30 18:15:20',0,'2017-11-30 18:15:20',0),(2,1,0,'1900-01-01','','','','','',1,0,0,0,'2017-11-30 18:15:28',0,'2017-11-30 18:15:28',0),(3,1,0,'1900-01-01','','','','','',1,0,0,0,'2017-11-30 18:16:10',0,'2017-11-30 18:16:10',0),(4,1,0,'1900-01-01','','','','','',1,0,0,0,'2017-11-30 18:16:17',0,'2017-11-30 18:16:17',0),(5,1,0,'1900-01-01','','','','','',1,0,0,0,'2017-11-30 18:16:28',0,'2017-11-30 18:16:28',0),(6,1,0,'1900-01-01','','','','','',1,0,0,0,'2017-11-30 18:24:52',0,'2017-11-30 18:24:52',0);
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

-- Dump completed on 2019-01-30 23:55:01
