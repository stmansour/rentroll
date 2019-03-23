-- MySQL dump 10.13  Distrib 5.7.22, for osx10.12 (x86_64)
--
-- Host: localhost    Database: rentroll
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
  `DefaultRentCycle` smallint(6) NOT NULL DEFAULT '0',
  `DefaultProrationCycle` smallint(6) NOT NULL DEFAULT '0',
  PRIMARY KEY (`ARID`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `AR`
--

LOCK TABLES `AR` WRITE;
/*!40000 ALTER TABLE `AR` DISABLE KEYS */;
INSERT INTO `AR` VALUES (1,1,'Application Fee',0,0,0,9,46,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,15.0000,'2018-08-15 02:53:15',211,'2017-11-10 23:24:23',0,0,0),(2,1,'Application Fee (no assessment)',0,1,0,7,46,'Application fee taken, no assessment made','0000-00-00 00:00:00','9999-12-31 00:00:00',5,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(3,1,'Apply Payment',0,1,0,10,9,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 02:59:12',211,'2017-11-10 23:24:23',0,0,0),(4,1,'Bad Debt Write-Off',0,2,0,71,9,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(5,1,'Bank Service Fee (Deposit Account)',0,2,0,72,4,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(6,1,'Bank Service Fee (Operating Account)',0,2,0,72,3,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(7,1,'Broken Window charge',0,0,0,9,59,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:00:40',211,'2017-11-10 23:24:23',0,0,0),(8,1,'Damage Fee',0,0,0,9,59,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:01:10',211,'2017-11-10 23:24:23',0,0,0),(9,1,'Deposit to Deposit Account (FRB96953)',0,1,0,4,6,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(10,1,'Deposit to Operating Account (FRB54320)',0,1,0,3,6,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(11,1,'Electric Base Fee',0,0,0,9,36,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(12,1,'Electric Overage',0,0,0,9,37,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:01:56',211,'2017-11-10 23:24:23',0,0,0),(13,1,'Eviction Fee Reimbursement',0,0,0,9,56,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(14,1,'Auto-Generated Floating Deposit Assessment',0,3,0,9,12,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(15,1,'Receive Floating Security Deposit',0,1,0,6,9,'','0000-00-00 00:00:00','9999-12-31 00:00:00',13,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(16,1,'Gas Base Fee',0,0,0,9,40,'','1900-01-01 00:00:00','9999-12-29 00:00:00',2,50.0000,'2018-07-23 16:16:36',211,'2017-11-10 23:24:23',0,6,4),(17,1,'Gas Base Overage',0,0,0,9,41,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:02:36',211,'2017-11-10 23:24:23',0,0,0),(18,1,'Insufficient Funds Fee',0,0,0,9,48,'','1900-01-01 00:00:00','9999-12-29 00:00:00',64,25.0000,'2018-08-15 03:03:11',211,'2017-11-10 23:24:23',0,0,0),(19,1,'Late Fee',0,0,0,9,47,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,5.0000,'2018-08-15 03:03:22',211,'2017-11-10 23:24:23',0,0,0),(20,1,'Month to Month Fee',0,0,0,9,49,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(21,1,'No Show / Termination Fee',0,0,0,9,51,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:03:52',211,'2017-11-10 23:24:23',0,0,0),(22,1,'Other Special Tenant Charges',0,0,0,9,61,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(23,1,'Pet Fee',0,0,0,9,52,'','1900-01-01 00:00:00','9999-12-31 00:00:00',192,50.0000,'2018-07-04 04:13:35',211,'2017-11-10 23:24:23',0,0,0),(24,1,'Pet Rent',0,0,0,9,53,'','1900-01-01 00:00:00','9999-12-30 00:00:00',144,10.0000,'2018-07-23 16:16:52',211,'2017-11-10 23:24:23',0,6,4),(25,1,'Receive a Payment',0,1,0,6,10,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(26,1,'Rent Non-Taxable',0,0,0,9,18,'','1900-01-01 00:00:00','9999-12-30 00:00:00',16,0.0000,'2018-07-23 16:14:54',211,'2017-11-10 23:24:23',0,6,4),(27,1,'Rent Taxable',0,0,0,9,17,'','1900-01-01 00:00:00','9999-12-30 00:00:00',16,0.0000,'2018-07-23 16:15:28',211,'2017-11-10 23:24:23',0,4,0),(28,1,'Security Deposit Assessment',0,0,0,9,11,'normal deposit','1900-01-01 00:00:00','9999-12-30 00:00:00',96,0.0000,'2018-08-15 03:04:36',211,'2017-11-10 23:24:23',0,0,0),(29,1,'Security Deposit Forfeiture',0,0,0,11,58,'Forfeit','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(30,1,'Security Deposit Refund',0,0,0,11,5,'Refund','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(31,1,'Special Cleaning Fee',0,0,0,9,55,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(32,1,'Tenant Expense Chargeback',0,0,0,9,54,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(33,1,'Vending Income',0,1,0,7,65,'','0000-00-00 00:00:00','9999-12-31 00:00:00',5,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(34,1,'Water and Sewer Base Fee',0,0,0,9,38,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(35,1,'Water and Sewer Overage',0,0,0,9,39,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:05:54',211,'2017-11-10 23:24:23',0,0,0),(36,1,'Auto-gen Application Fee Asmt',0,3,0,9,46,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(37,1,'Receive Application Fee (auto-gen asmt)',0,1,0,6,9,'Application fee taken, autogen asmt','0000-00-00 00:00:00','9999-12-31 00:00:00',13,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(38,1,'XFER  Operating to SecDep',0,2,0,4,3,'Move money from Operating acct to Sec Dep','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(39,1,'Vehicle Registration Fee',0,0,3,9,75,'','2018-01-01 00:00:00','9999-12-31 00:00:00',320,10.0000,'2018-07-04 04:14:56',211,'2018-07-03 02:47:47',211,0,0),(40,1,'Rent ST000',0,0,0,9,18,'Default rent assessment for rentable type RType000','1970-01-01 00:00:00','9999-12-30 00:00:00',16,1000.0000,'2018-08-15 03:07:17',211,'2018-07-27 06:49:30',0,6,4),(41,1,'Rent ST001',0,0,0,9,18,'Default rent assessment for rentable type RType001','1970-01-01 00:00:00','9999-12-30 00:00:00',16,1500.0000,'2018-08-15 03:07:28',211,'2018-07-27 06:49:30',0,6,4),(42,1,'Rent ST002',0,0,0,9,18,'Default rent assessment for rentable type RType002','1970-01-01 00:00:00','9999-12-30 00:00:00',16,1750.0000,'2018-08-15 03:07:40',211,'2018-07-27 06:49:30',0,6,4),(43,1,'Rent ST003',0,0,0,9,18,'Default rent assessment for rentable type RType003','1970-01-01 00:00:00','9999-12-30 00:00:00',16,2500.0000,'2018-08-15 03:16:30',211,'2018-07-27 06:49:30',0,6,4),(44,1,'Rent CP000',0,0,0,9,18,'Default rent assessment for rentable type Car Port 000','1970-01-01 00:00:00','9999-12-30 00:00:00',16,35.0000,'2018-08-15 03:16:45',211,'2018-07-27 06:49:30',0,6,4),(45,1,'Rent ST000',0,0,0,9,18,'Default rent assessment for rentable type RType000','1970-01-01 00:00:00','9999-12-31 00:00:00',20,1000.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,6,4),(46,1,'Rent ST001',0,0,0,9,18,'Default rent assessment for rentable type RType001','1970-01-01 00:00:00','9999-12-31 00:00:00',20,1500.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,6,4),(47,1,'Rent ST002',0,0,0,9,18,'Default rent assessment for rentable type RType002','1970-01-01 00:00:00','9999-12-31 00:00:00',20,1750.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,6,4),(48,1,'Rent ST003',0,0,0,9,18,'Default rent assessment for rentable type RType003','1970-01-01 00:00:00','9999-12-31 00:00:00',20,2500.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,6,4),(49,1,'Rent CP000',0,0,0,9,18,'Default rent assessment for rentable type Car Port 000','1970-01-01 00:00:00','9999-12-31 00:00:00',20,35.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,0,0);
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
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Assessments`
--

LOCK TABLES `Assessments` WRITE;
/*!40000 ALTER TABLE `Assessments` DISABLE KEYS */;
INSERT INTO `Assessments` VALUES (1,0,0,0,1,1,0,0,1,1000.0000,'2018-03-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',26,0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,0,0,0,1,1,0,0,1,2000.0000,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',28,2,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:31',0),(3,0,0,0,1,1,14,1,1,10.0000,'2018-03-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',24,8,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,0,0,0,1,1,0,0,1,571.4300,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',26,2,'prorated for 16 of 28 days','2018-10-16 19:35:32',0,'2018-10-16 19:35:31',0),(5,0,0,0,1,1,14,1,1,50.0000,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',23,2,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:31',0),(6,0,0,0,1,1,14,1,1,5.7100,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',24,2,'prorated for 16 of 28 days','2018-10-16 19:35:32',0,'2018-10-16 19:35:31',0),(7,0,0,0,1,1,15,1,1,10.0000,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',39,2,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:31',0),(8,0,0,0,1,1,15,2,1,10.0000,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',39,2,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:31',0),(9,0,0,0,1,2,0,0,2,1000.0000,'2018-03-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',26,0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(10,0,0,0,1,2,0,0,2,2000.0000,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',28,2,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:31',0),(11,0,0,0,1,2,0,0,2,571.4300,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',26,2,'prorated for 16 of 28 days','2018-10-16 19:35:32',0,'2018-10-16 19:35:31',0),(12,0,0,0,1,2,15,3,2,10.0000,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',39,2,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:31',0),(13,0,0,0,1,3,0,0,3,1000.0000,'2018-03-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',26,0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(14,0,0,0,1,3,0,0,3,2000.0000,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',28,2,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:31',0),(15,0,0,0,1,3,14,2,3,10.0000,'2018-03-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',24,8,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(16,0,0,0,1,3,0,0,3,571.4300,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',26,2,'prorated for 16 of 28 days','2018-10-16 19:35:32',0,'2018-10-16 19:35:31',0),(17,0,0,0,1,3,14,2,3,50.0000,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',23,2,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:31',0),(18,0,0,0,1,3,14,2,3,5.7100,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',24,2,'prorated for 16 of 28 days','2018-10-16 19:35:32',0,'2018-10-16 19:35:31',0),(19,0,0,0,1,3,15,4,3,10.0000,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',39,2,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:31',0),(20,0,0,0,1,3,15,5,3,10.0000,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',39,2,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(21,0,0,0,1,4,0,0,4,1500.0000,'2018-03-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',26,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(22,0,0,0,1,4,0,0,4,3000.0000,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',28,2,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(23,0,0,0,1,4,14,3,4,10.0000,'2018-03-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',24,8,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(24,0,0,0,1,4,0,0,4,857.1400,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',26,2,'prorated for 16 of 28 days','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(25,0,0,0,1,4,14,3,4,50.0000,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',23,2,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(26,0,0,0,1,4,14,3,4,5.7100,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',24,2,'prorated for 16 of 28 days','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(27,0,0,0,1,4,15,6,4,10.0000,'2018-02-13 00:00:00','2018-02-13 00:00:00',0,0,0,'',39,2,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(28,1,0,0,1,1,0,0,1,1000.0000,'2018-10-01 00:00:00','2018-10-01 00:00:00',6,4,0,'',26,0,'','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(29,3,0,0,1,1,14,1,1,10.0000,'2018-10-01 00:00:00','2018-10-01 00:00:00',6,4,0,'',24,0,'','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(30,9,0,0,1,2,0,0,2,1000.0000,'2018-10-01 00:00:00','2018-10-01 00:00:00',6,4,0,'',26,0,'','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(31,13,0,0,1,3,0,0,3,1000.0000,'2018-10-01 00:00:00','2018-10-01 00:00:00',6,4,0,'',26,0,'','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(32,15,0,0,1,3,14,2,3,10.0000,'2018-10-01 00:00:00','2018-10-01 00:00:00',6,4,0,'',24,0,'','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(33,21,0,0,1,4,0,0,4,1500.0000,'2018-10-01 00:00:00','2018-10-01 00:00:00',6,4,0,'',26,0,'','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(34,23,0,0,1,4,14,3,4,10.0000,'2018-10-01 00:00:00','2018-10-01 00:00:00',6,4,0,'',24,0,'','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(35,0,0,0,1,5,12,5,5,19.3500,'2018-10-20 00:00:00','2018-10-21 00:00:00',0,4,0,'',16,0,'prorated for 12 of 31 days','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(36,0,0,0,1,5,12,5,5,50.0000,'2018-11-01 00:00:00','2019-11-01 00:00:00',6,4,0,'',16,0,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(37,0,0,0,1,5,12,5,5,580.6500,'2018-10-20 00:00:00','2018-10-21 00:00:00',0,4,0,'',46,0,'prorated for 12 of 31 days','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(38,0,0,0,1,5,12,5,5,1500.0000,'2018-11-01 00:00:00','2019-11-01 00:00:00',6,4,0,'',46,0,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(39,0,0,0,1,5,12,5,5,3000.0000,'2018-10-16 00:00:00','2018-10-17 00:00:00',0,0,0,'',28,0,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(40,0,0,0,1,5,14,31,5,50.0000,'2018-10-20 00:00:00','2018-10-21 00:00:00',0,0,0,'',23,8,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(41,0,0,0,1,5,14,31,5,3.8700,'2018-10-20 00:00:00','2018-10-21 00:00:00',0,4,0,'',24,8,'prorated for 12 of 31 days','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(42,0,0,0,1,5,14,31,5,10.0000,'2018-11-01 00:00:00','2019-11-01 00:00:00',6,4,0,'',24,8,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(43,0,0,0,1,5,15,0,5,10.0000,'2018-10-20 00:00:00','2018-10-21 00:00:00',0,0,0,'',39,16,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(44,1,0,0,1,1,0,0,1,1000.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',26,0,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(45,3,0,0,1,1,14,1,1,10.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',24,0,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(46,9,0,0,1,2,0,0,2,1000.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',26,0,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(47,13,0,0,1,3,0,0,3,1000.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',26,0,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(48,15,0,0,1,3,14,2,3,10.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',24,0,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(49,21,0,0,1,4,0,0,4,1500.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',26,0,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(50,23,0,0,1,4,14,3,4,10.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',24,0,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(51,36,0,0,1,5,12,5,5,50.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',16,0,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(52,38,0,0,1,5,12,5,5,1500.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',46,0,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(53,42,0,0,1,5,14,31,5,10.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',24,0,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(54,21,0,0,1,4,0,0,4,1500.0000,'2019-03-01 00:00:00','2019-03-01 00:00:00',6,4,0,'',26,0,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(55,1,0,0,1,1,0,0,1,1000.0000,'2019-03-01 00:00:00','2019-03-01 00:00:00',6,4,0,'',26,0,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(56,9,0,0,1,2,0,0,2,1000.0000,'2019-03-01 00:00:00','2019-03-01 00:00:00',6,4,0,'',26,0,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(57,13,0,0,1,3,0,0,3,1000.0000,'2019-03-01 00:00:00','2019-03-01 00:00:00',6,4,0,'',26,0,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(58,3,0,0,1,1,14,1,1,10.0000,'2019-03-01 00:00:00','2019-03-01 00:00:00',6,4,0,'',24,0,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(59,15,0,0,1,3,14,2,3,10.0000,'2019-03-01 00:00:00','2019-03-01 00:00:00',6,4,0,'',24,0,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(60,23,0,0,1,4,14,3,4,10.0000,'2019-03-01 00:00:00','2019-03-01 00:00:00',6,4,0,'',24,0,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(61,38,0,0,1,5,12,5,5,1500.0000,'2019-03-01 00:00:00','2019-03-01 00:00:00',6,4,0,'',46,0,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(62,36,0,0,1,5,12,5,5,50.0000,'2019-03-01 00:00:00','2019-03-01 00:00:00',6,4,0,'',16,0,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(63,42,0,0,1,5,14,31,5,10.0000,'2019-03-01 00:00:00','2019-03-01 00:00:00',6,4,0,'',24,0,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1);
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
INSERT INTO `Business` VALUES (1,'REX','JGM First, LLC',6,4,4,1,'2018-06-05 23:06:51',0,'2017-11-10 23:24:22',0,1);
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `BusinessProperties`
--

LOCK TABLES `BusinessProperties` WRITE;
/*!40000 ALTER TABLE `BusinessProperties` DISABLE KEYS */;
INSERT INTO `BusinessProperties` VALUES (1,1,'general',0,'{\"Epochs\": {\"Daily\": \"2017-01-01T00:00:00Z\", \"Weekly\": \"2017-01-01T00:00:00Z\", \"Yearly\": \"2017-01-01T00:00:00Z\", \"Monthly\": \"2017-01-01T00:00:00Z\", \"Quarterly\": \"2017-01-01T00:00:00Z\"}, \"PetFees\": [\"Pet Fee\", \"Pet Rent\"], \"VehicleFees\": [\"Vehicle Registration Fee\"]}','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Deposit`
--

LOCK TABLES `Deposit` WRITE;
/*!40000 ALTER TABLE `Deposit` DISABLE KEYS */;
INSERT INTO `Deposit` VALUES (1,1,2,1,'2018-02-13',2798.5600,0.0000,0,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(2,1,1,1,'2018-02-13',9000.0000,0.0000,0,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0);
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
INSERT INTO `DepositMethod` VALUES (1,1,'Hand Delivery','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(2,1,'Scanned/Electronic Batch','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(3,1,'ACH','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(4,1,'US Mail','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DepositPart`
--

LOCK TABLES `DepositPart` WRITE;
/*!40000 ALTER TABLE `DepositPart` DISABLE KEYS */;
INSERT INTO `DepositPart` VALUES (1,1,1,2,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(2,1,1,3,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(3,1,1,4,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(4,1,1,5,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(5,1,1,6,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(6,1,1,8,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(7,1,1,9,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(8,1,1,11,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(9,1,1,12,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(10,1,1,13,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(11,1,1,14,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(12,1,1,15,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(13,1,1,17,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(14,1,1,18,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(15,1,1,19,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(16,1,1,20,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(17,2,1,1,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(18,2,1,7,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(19,2,1,10,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(20,2,1,16,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0);
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
INSERT INTO `Depository` VALUES (1,1,3,'Wells Fargo','987654321','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(2,1,4,'Bank Of America','12345678','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Flow`
--

LOCK TABLES `Flow` WRITE;
/*!40000 ALTER TABLE `Flow` DISABLE KEYS */;
INSERT INTO `Flow` VALUES (3,1,'FV9M2N2695EFPG795KQ4','RA',1,'{\"tie\": {\"people\": [{\"PRID\": 1, \"TMPTCID\": 1}]}, \"meta\": {\"BID\": 1, \"RAID\": 1, \"RAFLAGS\": 51, \"HavePets\": false, \"ActiveUID\": 0, \"Approver1\": 211, \"Approver2\": 211, \"MoveInUID\": 211, \"ActiveDate\": \"1900-01-01 00:00:00 UTC\", \"ActiveName\": \"\", \"LastTMPVID\": 2, \"MoveInDate\": \"2018-10-19 03:02:00 UTC\", \"MoveInName\": \"Steve Mansour\", \"LastTMPTCID\": 1, \"DocumentDate\": \"2018-10-18 00:00:00 UTC\", \"HaveVehicles\": false, \"LastTMPASMID\": 6, \"LastTMPPETID\": 1, \"Approver1Name\": \"Steve Mansour\", \"Approver2Name\": \"Steve Mansour\", \"DecisionDate1\": \"2018-10-19 03:02:00 UTC\", \"DecisionDate2\": \"2018-10-19 03:02:00 UTC\", \"TerminatorUID\": 0, \"DeclineReason1\": 0, \"DeclineReason2\": 0, \"TerminatorName\": \"UID-0\", \"NoticeToMoveUID\": 0, \"TerminationDate\": \"1900-01-01 00:00:00 UTC\", \"NoticeToMoveDate\": \"1900-01-01 00:00:00 UTC\", \"NoticeToMoveName\": \"UID-0\", \"ApplicationReadyUID\": 211, \"ApplicationReadyDate\": \"2018-10-19 03:02:00 UTC\", \"ApplicationReadyName\": \"Steve Mansour\", \"NoticeToMoveReported\": \"1900-01-01 00:00:00 UTC\", \"LeaseTerminationReason\": 0}, \"pets\": [{\"Fees\": [{\"ARID\": 24, \"Stop\": \"10/19/2018\", \"ASMID\": 3, \"Start\": \"10/19/2018\", \"ARName\": \"Pet Rent\", \"Comment\": \"prorated for 13 of 31 days\", \"SalesTax\": 0, \"TMPASMID\": 3, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": 4.19, \"ProrationCycle\": 0, \"AtSigningPreTax\": 0}, {\"ARID\": 24, \"Stop\": \"2/29/2020\", \"ASMID\": 3, \"Start\": \"11/1/2018\", \"ARName\": \"Pet Rent\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 4, \"RentCycle\": 6, \"TransOccTax\": 0, \"ContractAmount\": 10, \"ProrationCycle\": 4, \"AtSigningPreTax\": 0}], \"Name\": \"Oreo\", \"Type\": \"dog\", \"Breed\": \"English Cocker Spaniel\", \"Color\": \"Grey\", \"PETID\": 1, \"Weight\": 40, \"TMPTCID\": 1, \"TMPPETID\": 1}], \"dates\": {\"CSAgent\": 283, \"RentStop\": \"2/29/2020\", \"RentStart\": \"10/19/2018\", \"AgreementStop\": \"2/29/2020\", \"AgreementStart\": \"10/19/2018\", \"PossessionStop\": \"2/29/2020\", \"PossessionStart\": \"10/19/2018\"}, \"people\": [{\"City\": \"Chattanooga\", \"TCID\": 1, \"State\": \"VT\", \"Points\": 0, \"Address\": \"59256 Apache\", \"Comment\": \"\", \"Country\": \"USA\", \"Evicted\": false, \"TMPTCID\": 1, \"Website\": \"\", \"Address2\": \"\", \"Industry\": 304, \"IsRenter\": true, \"LastName\": \"Pearson\", \"CellPhone\": \"(483) 719-8541\", \"Convicted\": false, \"FirstName\": \"Pablo\", \"IsCompany\": false, \"WorkPhone\": \"(597) 731-5597\", \"Bankruptcy\": false, \"EvictedDes\": \"\", \"IsOccupant\": true, \"MiddleName\": \"Chantell\", \"Occupation\": \"sound effects technician\", \"PostalCode\": \"98081\", \"TaxpayorID\": \"01129584\", \"CompanyCity\": \"FortWayne\", \"CompanyName\": \"Capital One Financial Corp.\", \"CreditLimit\": 5563, \"DateofBirth\": \"7/17/1974\", \"GrossIncome\": 73697, \"IsGuarantor\": false, \"SourceSLSID\": 26, \"CompanyEmail\": \"CapitalOneFinancialCorpF3900@gmail.com\", \"CompanyPhone\": \"(895) 638-0659\", \"CompanyState\": \"WV\", \"ConvictedDes\": \"\", \"PrimaryEmail\": \"PabloPearson178@hotmail.com\", \"PriorAddress\": \"36865 Oak, Bethlehem, CO 41318\", \"SpecialNeeds\": \"\", \"BankruptcyDes\": \"\", \"PreferredName\": \"Shenita\", \"CompanyAddress\": \"82595 Park\", \"CurrentAddress\": \"8144 Aloha, New Orleans, IL 02081\", \"DriversLicense\": \"D1408541\", \"SecondaryEmail\": \"PabloPearson203@aol.com\", \"OtherPreferences\": \"\", \"ThirdPartySource\": \"Laci Guthrie\", \"CompanyPostalCode\": \"84059\", \"PriorLandLordName\": \"Carolina Wyatt\", \"EligibleFutureUser\": false, \"CurrentLandLordName\": \"Otis Oliver\", \"EligibleFuturePayor\": true, \"EmergencyContactName\": \"Sigrid Sanford\", \"PriorLandLordPhoneNo\": \"(511) 768-4962\", \"PriorReasonForMoving\": 128, \"AlternateEmailAddress\": \"35824 Sycamore,Cathedral City,TX 31847\", \"EmergencyContactEmail\": \"SSanford3254@comcast.net\", \"CurrentLandLordPhoneNo\": \"(562) 975-6102\", \"CurrentReasonForMoving\": 84, \"PriorLengthOfResidency\": \"7 years 2 months\", \"EmergencyContactAddress\": \"30059 Aspen,Indianapolis,VA 27887\", \"CurrentLengthOfResidency\": \"7 years 8 months\", \"EmergencyContactTelephone\": \"(956) 631-6546\"}], \"vehicles\": [{\"VID\": 1, \"VIN\": \"WAU2F7X8N9DLXDEM\", \"Fees\": [], \"TMPVID\": 1, \"TMPTCID\": 1, \"VehicleMake\": \"Mazda\", \"VehicleType\": \"car\", \"VehicleYear\": 1998, \"VehicleColor\": \"Turquoise\", \"VehicleModel\": \"B-Series\", \"LicensePlateState\": \"VA\", \"LicensePlateNumber\": \"CUH0864\", \"ParkingPermitNumber\": \"2860465\"}, {\"VID\": 2, \"VIN\": \"X1MO1JIW0ZPPN3PK\", \"Fees\": [], \"TMPVID\": 2, \"TMPTCID\": 1, \"VehicleMake\": \"Ford\", \"VehicleType\": \"car\", \"VehicleYear\": 2011, \"VehicleColor\": \"Navy\", \"VehicleModel\": \"F-Series Super Duty\", \"LicensePlateState\": \"MA\", \"LicensePlateNumber\": \"C3Q722U\", \"ParkingPermitNumber\": \"6386881\"}], \"rentables\": [{\"RID\": 1, \"Fees\": [{\"ARID\": 26, \"Stop\": \"10/19/2018\", \"ASMID\": 1, \"Start\": \"10/19/2018\", \"ARName\": \"Rent Non-Taxable\", \"Comment\": \"prorated for 13 of 31 days\", \"SalesTax\": 0, \"TMPASMID\": 5, \"RentCycle\": 0, \"TransOccTax\": 0, \"ContractAmount\": 419.35, \"ProrationCycle\": 0, \"AtSigningPreTax\": 0}, {\"ARID\": 26, \"Stop\": \"2/29/2020\", \"ASMID\": 1, \"Start\": \"11/1/2018\", \"ARName\": \"Rent Non-Taxable\", \"Comment\": \"\", \"SalesTax\": 0, \"TMPASMID\": 6, \"RentCycle\": 6, \"TransOccTax\": 0, \"ContractAmount\": 1000, \"ProrationCycle\": 4, \"AtSigningPreTax\": 0}], \"RTID\": 1, \"RTFLAGS\": 4, \"SalesTax\": 0, \"RentCycle\": 6, \"TransOccTax\": 0, \"RentableName\": \"Rentable001\", \"AtSigningPreTax\": 0}], \"parentchild\": []}','2018-10-19 03:02:45',211,'2018-10-19 03:02:24',211);
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
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `GLAccount`
--

LOCK TABLES `GLAccount` WRITE;
/*!40000 ALTER TABLE `GLAccount` DISABLE KEYS */;
INSERT INTO `GLAccount` VALUES (1,0,1,0,0,'10000','Cash','Cash',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(2,0,1,0,0,'10100','Petty Cash','Cash',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(3,1,1,0,0,'10104','FRB 54320 (operating account)','Bank Account',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(4,1,1,0,0,'10105','FRB 96953 (deposit account)','Bank Account',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(5,1,1,0,0,'10199','Security Deposit Refund','Cash',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(6,1,1,0,0,'10999','Undeposited Funds','Cash',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(7,0,1,0,0,'11000','Credit Cards Funds in Transit','Cash',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(8,0,1,0,0,'12000','Accounts Receivable','Accounts Receivable',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(9,8,1,0,0,'12001','Rent Roll Receivables','Accounts Receivable',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(10,0,1,0,0,'12999','Unapplied Funds','Asset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(11,0,1,0,0,'30000','Security Deposit Liability','Liability Security Deposit',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(12,0,1,0,0,'30001','Floating Security Deposits','Liability Security Deposit',1,0,'Sec Dep posted before rentable identified','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(13,0,1,0,0,'30100','Collected Taxes','Liabilities',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(14,13,1,0,0,'30101','Sales Taxes Collected','Liabilities',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(15,13,1,0,0,'30102','Transient Occupancy Taxes Collected','Liabilities',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(16,13,1,0,0,'30199','Other Collected Taxes','Liabilities',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(17,0,1,0,0,'41000','Gross Scheduled Rent-Taxable','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(18,0,1,0,0,'41001','Gross Scheduled Rent-Not Taxable','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(19,0,1,0,0,'41100','Unit Income Offsets','Income Offset',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(20,19,1,0,0,'41101','Vacancy','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(21,19,1,0,0,'41102','Loss (Gain) to Lease','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(22,19,1,0,0,'41103','Employee Concessions','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(23,19,1,0,0,'41104','Resident Concessions','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(24,19,1,0,0,'41105','Owner Concession','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(25,19,1,0,0,'41106','Administrative Concession','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(26,19,1,0,0,'41107','Off Line Renovations','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(27,19,1,0,0,'41108','Off Line Maintenance','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(28,19,1,0,0,'41199','Othe Income Offsets','Income Offset',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(29,0,1,0,0,'41200','Service Fees','Income',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(30,29,1,0,0,'41201','Broadcast and IT Services','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(31,29,1,0,0,'41202','Food Services','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(32,29,1,0,0,'41203','Linen Services','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(33,29,1,0,0,'41204','Wash N Fold Services','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(34,29,1,0,0,'41299','Other Service Fees','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(35,0,1,0,0,'41300','Utility Fees','Income',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(36,35,1,0,0,'41301','Electric Base Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(37,35,1,0,0,'41302','Electric Overage','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(38,35,1,0,0,'41303','Water and Sewer Base Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(39,35,1,0,0,'41304','Water and Sewer Overage','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(40,35,1,0,0,'41305','Gas Base Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(41,35,1,0,0,'41306','Gas Overage','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(42,35,1,0,0,'41307','Trash Collection Base Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(43,35,1,0,0,'41308','Trash Collection Overage','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(44,35,1,0,0,'41399','Other Utility Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(45,0,1,0,0,'41400','Special Tenant Charges','Income',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(46,45,1,0,0,'41401','Application Fees','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(47,45,1,0,0,'41402','Late Fees','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(48,45,1,0,0,'41403','Insufficient Funds Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(49,45,1,0,0,'41404','Month to Month Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(50,45,1,0,0,'41405','Rentable Specialties','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(51,45,1,0,0,'41406','No Show or Termination Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(52,45,1,0,0,'41407','Pet Fees','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(53,45,1,0,0,'41408','Pet Rent','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(54,45,1,0,0,'41409','Tenant Expense Chargeback','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(55,45,1,0,0,'41410','Special Cleaning Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(56,45,1,0,0,'41411','Eviction Fee Reimbursement','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(57,45,1,0,0,'41412','Extra Person Charge','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(58,45,1,0,0,'41413','Security Deposit Forfeiture','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(59,45,1,0,0,'41414','Damage Fee','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(60,45,1,0,0,'41415','CAM Fees','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(61,45,1,0,0,'41499','Other Special Tenant Charges','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(62,0,1,0,0,'42000','Business Income','Income',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(63,62,1,0,0,'42100','Convenience Store','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(64,62,1,0,0,'42200','Fitness Center Revenue','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(65,62,1,0,0,'42300','Vending Income','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(66,62,1,0,0,'42400','Restaurant Sales','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(67,62,1,0,0,'42500','Bar Sales','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(68,62,1,0,0,'42600','Spa Sales','Income',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(69,0,1,0,0,'50000','Expenses','Expenses',0,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(70,69,1,0,0,'50001','Cash Over/Short','Expenses',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(71,69,1,0,0,'50002','Bad Debt','Expenses',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(72,69,1,0,0,'50003','Bank Service Fee','Expenses',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(73,69,1,0,0,'50999','Other Expenses','Expenses',1,0,'','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(74,0,1,0,0,'999911','test 1','Cash',1,0,'laskdjf','2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(75,45,1,0,0,'41416','Vehicle Fees','Income',1,0,'Vehicle Registration fees','2018-07-03 02:45:37',211,'2018-07-03 02:45:37',211);
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
) ENGINE=InnoDB AUTO_INCREMENT=114 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Journal`
--

LOCK TABLES `Journal` WRITE;
/*!40000 ALTER TABLE `Journal` DISABLE KEYS */;
INSERT INTO `Journal` VALUES (1,1,'2018-02-13 00:00:00',2000.0000,1,2,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,1,'2018-02-13 00:00:00',571.4300,1,4,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,1,'2018-02-13 00:00:00',50.0000,1,5,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,1,'2018-02-13 00:00:00',5.7100,1,6,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(5,1,'2018-02-13 00:00:00',10.0000,1,7,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(6,1,'2018-02-13 00:00:00',10.0000,1,8,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(7,1,'2018-02-13 00:00:00',2000.0000,1,10,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(8,1,'2018-02-13 00:00:00',571.4300,1,11,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(9,1,'2018-02-13 00:00:00',10.0000,1,12,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(10,1,'2018-02-13 00:00:00',2000.0000,1,14,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(11,1,'2018-02-13 00:00:00',571.4300,1,16,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(12,1,'2018-02-13 00:00:00',50.0000,1,17,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(13,1,'2018-02-13 00:00:00',5.7100,1,18,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(14,1,'2018-02-13 00:00:00',10.0000,1,19,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(15,1,'2018-02-13 00:00:00',10.0000,1,20,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(16,1,'2018-02-13 00:00:00',3000.0000,1,22,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(17,1,'2018-02-13 00:00:00',857.1400,1,24,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(18,1,'2018-02-13 00:00:00',50.0000,1,25,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(19,1,'2018-02-13 00:00:00',5.7100,1,26,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(20,1,'2018-02-13 00:00:00',10.0000,1,27,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(21,1,'2018-02-13 00:00:00',2000.0000,2,1,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(22,1,'2018-02-13 00:00:00',571.4300,2,2,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(23,1,'2018-02-13 00:00:00',50.0000,2,3,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(24,1,'2018-02-13 00:00:00',5.7100,2,4,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(25,1,'2018-02-13 00:00:00',10.0000,2,5,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(26,1,'2018-02-13 00:00:00',10.0000,2,6,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(27,1,'2018-02-13 00:00:00',2000.0000,2,7,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(28,1,'2018-02-13 00:00:00',571.4300,2,8,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(29,1,'2018-02-13 00:00:00',10.0000,2,9,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(30,1,'2018-02-13 00:00:00',2000.0000,2,10,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(31,1,'2018-02-13 00:00:00',571.4300,2,11,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(32,1,'2018-02-13 00:00:00',50.0000,2,12,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(33,1,'2018-02-13 00:00:00',5.7100,2,13,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(34,1,'2018-02-13 00:00:00',10.0000,2,14,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(35,1,'2018-02-13 00:00:00',10.0000,2,15,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(36,1,'2018-02-13 00:00:00',3000.0000,2,16,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(37,1,'2018-02-13 00:00:00',857.1400,2,17,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(38,1,'2018-02-13 00:00:00',50.0000,2,18,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(39,1,'2018-02-13 00:00:00',5.7100,2,19,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(40,1,'2018-02-13 00:00:00',10.0000,2,20,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(41,1,'2018-02-13 00:00:00',2000.0000,2,1,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(42,1,'2018-02-13 00:00:00',571.4300,2,2,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(43,1,'2018-02-13 00:00:00',50.0000,2,3,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(44,1,'2018-02-13 00:00:00',5.7100,2,4,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(45,1,'2018-02-13 00:00:00',10.0000,2,5,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(46,1,'2018-02-13 00:00:00',10.0000,2,6,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(47,1,'2018-02-13 00:00:00',2000.0000,2,7,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(48,1,'2018-02-13 00:00:00',571.4300,2,8,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(49,1,'2018-02-13 00:00:00',10.0000,2,9,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(50,1,'2018-02-13 00:00:00',2000.0000,2,10,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(51,1,'2018-02-13 00:00:00',571.4300,2,11,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(52,1,'2018-02-13 00:00:00',50.0000,2,12,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(53,1,'2018-02-13 00:00:00',5.7100,2,13,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(54,1,'2018-02-13 00:00:00',10.0000,2,14,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(55,1,'2018-02-13 00:00:00',10.0000,2,15,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(56,1,'2018-02-13 00:00:00',3000.0000,2,16,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(57,1,'2018-02-13 00:00:00',857.1400,2,17,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(58,1,'2018-02-13 00:00:00',50.0000,2,18,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(59,1,'2018-02-13 00:00:00',5.7100,2,19,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(60,1,'2018-02-13 00:00:00',10.0000,2,20,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(61,1,'2018-02-13 00:00:00',571.4300,4,2,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(62,1,'2018-02-13 00:00:00',50.0000,4,3,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(63,1,'2018-02-13 00:00:00',5.7100,4,4,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(64,1,'2018-02-13 00:00:00',10.0000,4,5,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(65,1,'2018-02-13 00:00:00',10.0000,4,6,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(66,1,'2018-02-13 00:00:00',571.4300,4,8,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(67,1,'2018-02-13 00:00:00',10.0000,4,9,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(68,1,'2018-02-13 00:00:00',571.4300,4,11,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(69,1,'2018-02-13 00:00:00',50.0000,4,12,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(70,1,'2018-02-13 00:00:00',5.7100,4,13,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(71,1,'2018-02-13 00:00:00',10.0000,4,14,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(72,1,'2018-02-13 00:00:00',10.0000,4,15,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(73,1,'2018-02-13 00:00:00',857.1400,4,17,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(74,1,'2018-02-13 00:00:00',50.0000,4,18,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(75,1,'2018-02-13 00:00:00',5.7100,4,19,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(76,1,'2018-02-13 00:00:00',10.0000,4,20,'auto-transfer for deposit DEP-2','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(77,1,'2018-02-13 00:00:00',2000.0000,4,1,'auto-transfer for deposit DEP-1','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(78,1,'2018-02-13 00:00:00',2000.0000,4,7,'auto-transfer for deposit DEP-1','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(79,1,'2018-02-13 00:00:00',2000.0000,4,10,'auto-transfer for deposit DEP-1','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(80,1,'2018-02-13 00:00:00',3000.0000,4,16,'auto-transfer for deposit DEP-1','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(81,1,'2018-10-01 00:00:00',1000.0000,1,28,'','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(82,1,'2018-10-01 00:00:00',10.0000,1,29,'','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(83,1,'2018-10-01 00:00:00',1000.0000,1,30,'','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(84,1,'2018-10-01 00:00:00',1000.0000,1,31,'','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(85,1,'2018-10-01 00:00:00',10.0000,1,32,'','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(86,1,'2018-10-01 00:00:00',1500.0000,1,33,'','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(87,1,'2018-10-01 00:00:00',10.0000,1,34,'','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(88,1,'2018-10-20 00:00:00',19.3500,1,35,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(89,1,'2018-10-20 00:00:00',580.6500,1,37,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(90,1,'2018-10-16 00:00:00',3000.0000,1,39,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(91,1,'2018-10-20 00:00:00',50.0000,1,40,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(92,1,'2018-10-20 00:00:00',3.8700,1,41,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(93,1,'2018-10-20 00:00:00',10.0000,1,43,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(94,1,'2019-02-01 00:00:00',1000.0000,1,44,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(95,1,'2019-02-01 00:00:00',10.0000,1,45,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(96,1,'2019-02-01 00:00:00',1000.0000,1,46,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(97,1,'2019-02-01 00:00:00',1000.0000,1,47,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(98,1,'2019-02-01 00:00:00',10.0000,1,48,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(99,1,'2019-02-01 00:00:00',1500.0000,1,49,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(100,1,'2019-02-01 00:00:00',10.0000,1,50,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(101,1,'2019-02-01 00:00:00',50.0000,1,51,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(102,1,'2019-02-01 00:00:00',1500.0000,1,52,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(103,1,'2019-02-01 00:00:00',10.0000,1,53,'','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(104,1,'2019-03-01 00:00:00',1500.0000,1,54,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(105,1,'2019-03-01 00:00:00',1000.0000,1,55,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(106,1,'2019-03-01 00:00:00',1000.0000,1,56,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(107,1,'2019-03-01 00:00:00',1000.0000,1,57,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(108,1,'2019-03-01 00:00:00',10.0000,1,58,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(109,1,'2019-03-01 00:00:00',10.0000,1,59,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(110,1,'2019-03-01 00:00:00',10.0000,1,60,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(111,1,'2019-03-01 00:00:00',1500.0000,1,61,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(112,1,'2019-03-01 00:00:00',50.0000,1,62,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(113,1,'2019-03-01 00:00:00',10.0000,1,63,'','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1);
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
) ENGINE=InnoDB AUTO_INCREMENT=114 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `JournalAllocation`
--

LOCK TABLES `JournalAllocation` WRITE;
/*!40000 ALTER TABLE `JournalAllocation` DISABLE KEYS */;
INSERT INTO `JournalAllocation` VALUES (1,1,1,1,1,0,0,2000.0000,2,0,'d 12001 2000.00, c 30000 2000.00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,1,2,1,1,0,0,571.4300,4,0,'d 12001 571.43, c 41001 571.43','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,1,3,1,1,0,0,50.0000,5,0,'d 12001 50.00, c 41407 50.00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,1,4,1,1,0,0,5.7100,6,0,'d 12001 5.71, c 41408 5.71','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(5,1,5,1,1,0,0,10.0000,7,0,'d 12001 10.00, c 41416 10.00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(6,1,6,1,1,0,0,10.0000,8,0,'d 12001 10.00, c 41416 10.00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(7,1,7,2,2,0,0,2000.0000,10,0,'d 12001 2000.00, c 30000 2000.00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(8,1,8,2,2,0,0,571.4300,11,0,'d 12001 571.43, c 41001 571.43','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(9,1,9,2,2,0,0,10.0000,12,0,'d 12001 10.00, c 41416 10.00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(10,1,10,3,3,0,0,2000.0000,14,0,'d 12001 2000.00, c 30000 2000.00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(11,1,11,3,3,0,0,571.4300,16,0,'d 12001 571.43, c 41001 571.43','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(12,1,12,3,3,0,0,50.0000,17,0,'d 12001 50.00, c 41407 50.00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(13,1,13,3,3,0,0,5.7100,18,0,'d 12001 5.71, c 41408 5.71','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(14,1,14,3,3,0,0,10.0000,19,0,'d 12001 10.00, c 41416 10.00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(15,1,15,3,3,0,0,10.0000,20,0,'d 12001 10.00, c 41416 10.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(16,1,16,4,4,0,0,3000.0000,22,0,'d 12001 3000.00, c 30000 3000.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(17,1,17,4,4,0,0,857.1400,24,0,'d 12001 857.14, c 41001 857.14','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(18,1,18,4,4,0,0,50.0000,25,0,'d 12001 50.00, c 41407 50.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(19,1,19,4,4,0,0,5.7100,26,0,'d 12001 5.71, c 41408 5.71','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(20,1,20,4,4,0,0,10.0000,27,0,'d 12001 10.00, c 41416 10.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(21,1,21,0,0,1,0,2000.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(22,1,22,0,0,1,0,571.4300,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(23,1,23,0,0,1,0,50.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(24,1,24,0,0,1,0,5.7100,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(25,1,25,0,0,1,0,10.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(26,1,26,0,0,1,0,10.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(27,1,27,0,0,2,0,2000.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(28,1,28,0,0,2,0,571.4300,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(29,1,29,0,0,2,0,10.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(30,1,30,0,0,3,0,2000.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(31,1,31,0,0,3,0,571.4300,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(32,1,32,0,0,3,0,50.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(33,1,33,0,0,3,0,5.7100,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(34,1,34,0,0,3,0,10.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(35,1,35,0,0,3,0,10.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(36,1,36,0,0,4,0,3000.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(37,1,37,0,0,4,0,857.1400,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(38,1,38,0,0,4,0,50.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(39,1,39,0,0,4,0,5.7100,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(40,1,40,0,0,4,0,10.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(41,1,41,1,1,1,1,2000.0000,2,0,'ASM(2) d 12999 2000.00,c 12001 2000.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(42,1,42,1,1,1,2,571.4300,4,0,'ASM(4) d 12999 571.43,c 12001 571.43','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(43,1,43,1,1,1,3,50.0000,5,0,'ASM(5) d 12999 50.00,c 12001 50.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(44,1,44,1,1,1,4,5.7100,6,0,'ASM(6) d 12999 5.71,c 12001 5.71','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(45,1,45,1,1,1,5,10.0000,7,0,'ASM(7) d 12999 10.00,c 12001 10.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(46,1,46,1,1,1,6,10.0000,8,0,'ASM(8) d 12999 10.00,c 12001 10.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(47,1,47,2,2,2,7,2000.0000,10,0,'ASM(10) d 12999 2000.00,c 12001 2000.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(48,1,48,2,2,2,8,571.4300,11,0,'ASM(11) d 12999 571.43,c 12001 571.43','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(49,1,49,2,2,2,9,10.0000,12,0,'ASM(12) d 12999 10.00,c 12001 10.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(50,1,50,3,3,3,10,2000.0000,14,0,'ASM(14) d 12999 2000.00,c 12001 2000.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(51,1,51,3,3,3,11,571.4300,16,0,'ASM(16) d 12999 571.43,c 12001 571.43','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(52,1,52,3,3,3,12,50.0000,17,0,'ASM(17) d 12999 50.00,c 12001 50.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(53,1,53,3,3,3,13,5.7100,18,0,'ASM(18) d 12999 5.71,c 12001 5.71','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(54,1,54,3,3,3,14,10.0000,19,0,'ASM(19) d 12999 10.00,c 12001 10.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(55,1,55,3,3,3,15,10.0000,20,0,'ASM(20) d 12999 10.00,c 12001 10.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(56,1,56,4,4,4,16,3000.0000,22,0,'ASM(22) d 12999 3000.00,c 12001 3000.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(57,1,57,4,4,4,17,857.1400,24,0,'ASM(24) d 12999 857.14,c 12001 857.14','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(58,1,58,4,4,4,18,50.0000,25,0,'ASM(25) d 12999 50.00,c 12001 50.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(59,1,59,4,4,4,19,5.7100,26,0,'ASM(26) d 12999 5.71,c 12001 5.71','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(60,1,60,4,4,4,20,10.0000,27,0,'ASM(27) d 12999 10.00,c 12001 10.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(61,1,61,0,1,1,2,571.4300,4,0,'d 10105 571.4300, c 10999 571.4300','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(62,1,62,0,1,1,3,50.0000,5,0,'d 10105 50.0000, c 10999 50.0000','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(63,1,63,0,1,1,4,5.7100,6,0,'d 10105 5.7100, c 10999 5.7100','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(64,1,64,0,1,1,5,10.0000,7,0,'d 10105 10.0000, c 10999 10.0000','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(65,1,65,0,1,1,6,10.0000,8,0,'d 10105 10.0000, c 10999 10.0000','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(66,1,66,0,2,2,8,571.4300,11,0,'d 10105 571.4300, c 10999 571.4300','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(67,1,67,0,2,2,9,10.0000,12,0,'d 10105 10.0000, c 10999 10.0000','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(68,1,68,0,3,3,11,571.4300,16,0,'d 10105 571.4300, c 10999 571.4300','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(69,1,69,0,3,3,12,50.0000,17,0,'d 10105 50.0000, c 10999 50.0000','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(70,1,70,0,3,3,13,5.7100,18,0,'d 10105 5.7100, c 10999 5.7100','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(71,1,71,0,3,3,14,10.0000,19,0,'d 10105 10.0000, c 10999 10.0000','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(72,1,72,0,3,3,15,10.0000,20,0,'d 10105 10.0000, c 10999 10.0000','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(73,1,73,0,4,4,17,857.1400,24,0,'d 10105 857.1400, c 10999 857.1400','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(74,1,74,0,4,4,18,50.0000,25,0,'d 10105 50.0000, c 10999 50.0000','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(75,1,75,0,4,4,19,5.7100,26,0,'d 10105 5.7100, c 10999 5.7100','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(76,1,76,0,4,4,20,10.0000,27,0,'d 10105 10.0000, c 10999 10.0000','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(77,1,77,0,1,1,1,2000.0000,2,0,'d 10104 2000.0000, c 10999 2000.0000','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(78,1,78,0,2,2,7,2000.0000,10,0,'d 10104 2000.0000, c 10999 2000.0000','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(79,1,79,0,3,3,10,2000.0000,14,0,'d 10104 2000.0000, c 10999 2000.0000','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(80,1,80,0,4,4,16,3000.0000,22,0,'d 10104 3000.0000, c 10999 3000.0000','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(81,1,81,1,1,0,0,1000.0000,28,0,'d 12001 1000.00, c 41001 1000.00','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(82,1,82,1,1,0,0,10.0000,29,0,'d 12001 10.00, c 41408 10.00','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(83,1,83,2,2,0,0,1000.0000,30,0,'d 12001 1000.00, c 41001 1000.00','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(84,1,84,3,3,0,0,1000.0000,31,0,'d 12001 1000.00, c 41001 1000.00','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(85,1,85,3,3,0,0,10.0000,32,0,'d 12001 10.00, c 41408 10.00','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(86,1,86,4,4,0,0,1500.0000,33,0,'d 12001 1500.00, c 41001 1500.00','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(87,1,87,4,4,0,0,10.0000,34,0,'d 12001 10.00, c 41408 10.00','2018-10-16 19:35:38',-1,'2018-10-16 19:35:38',-1),(88,1,88,5,5,0,0,19.3500,35,0,'d 12001 19.35, c 41305 19.35','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(89,1,89,5,5,0,0,580.6500,37,0,'d 12001 580.65, c 41001 580.65','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(90,1,90,5,5,0,0,3000.0000,39,0,'d 12001 3000.00, c 30000 3000.00','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(91,1,91,5,5,0,0,50.0000,40,0,'d 12001 50.00, c 41407 50.00','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(92,1,92,5,5,0,0,3.8700,41,0,'d 12001 3.87, c 41408 3.87','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(93,1,93,5,5,0,0,10.0000,43,0,'d 12001 10.00, c 41416 10.00','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(94,1,94,1,1,0,0,1000.0000,44,0,'d 12001 1000.00, c 41001 1000.00','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(95,1,95,1,1,0,0,10.0000,45,0,'d 12001 10.00, c 41408 10.00','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(96,1,96,2,2,0,0,1000.0000,46,0,'d 12001 1000.00, c 41001 1000.00','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(97,1,97,3,3,0,0,1000.0000,47,0,'d 12001 1000.00, c 41001 1000.00','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(98,1,98,3,3,0,0,10.0000,48,0,'d 12001 10.00, c 41408 10.00','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(99,1,99,4,4,0,0,1500.0000,49,0,'d 12001 1500.00, c 41001 1500.00','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(100,1,100,4,4,0,0,10.0000,50,0,'d 12001 10.00, c 41408 10.00','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(101,1,101,5,5,0,0,50.0000,51,0,'d 12001 50.00, c 41305 50.00','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(102,1,102,5,5,0,0,1500.0000,52,0,'d 12001 1500.00, c 41001 1500.00','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(103,1,103,5,5,0,0,10.0000,53,0,'d 12001 10.00, c 41408 10.00','2019-02-12 17:11:24',-1,'2019-02-12 17:11:24',-1),(104,1,104,4,4,0,0,1500.0000,54,0,'d 12001 1500.00, c 41001 1500.00','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(105,1,105,1,1,0,0,1000.0000,55,0,'d 12001 1000.00, c 41001 1000.00','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(106,1,106,2,2,0,0,1000.0000,56,0,'d 12001 1000.00, c 41001 1000.00','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(107,1,107,3,3,0,0,1000.0000,57,0,'d 12001 1000.00, c 41001 1000.00','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(108,1,108,1,1,0,0,10.0000,58,0,'d 12001 10.00, c 41408 10.00','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(109,1,109,3,3,0,0,10.0000,59,0,'d 12001 10.00, c 41408 10.00','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(110,1,110,4,4,0,0,10.0000,60,0,'d 12001 10.00, c 41408 10.00','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(111,1,111,5,5,0,0,1500.0000,61,0,'d 12001 1500.00, c 41001 1500.00','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(112,1,112,5,5,0,0,50.0000,62,0,'d 12001 50.00, c 41305 50.00','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1),(113,1,113,5,5,0,0,10.0000,63,0,'d 12001 10.00, c 41408 10.00','2019-03-01 18:41:00',-1,'2019-03-01 18:41:00',-1);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
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
) ENGINE=InnoDB AUTO_INCREMENT=173 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerEntry`
--

LOCK TABLES `LedgerEntry` WRITE;
/*!40000 ALTER TABLE `LedgerEntry` DISABLE KEYS */;
INSERT INTO `LedgerEntry` VALUES (1,1,1,1,9,1,1,0,'2018-02-13 00:00:00',2000.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,1,1,1,11,1,1,0,'2018-02-13 00:00:00',-2000.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,1,2,2,9,1,1,0,'2018-02-13 00:00:00',571.4300,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,1,2,2,18,1,1,0,'2018-02-13 00:00:00',-571.4300,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(5,1,3,3,9,1,1,0,'2018-02-13 00:00:00',50.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(6,1,3,3,52,1,1,0,'2018-02-13 00:00:00',-50.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(7,1,4,4,9,1,1,0,'2018-02-13 00:00:00',5.7100,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(8,1,4,4,53,1,1,0,'2018-02-13 00:00:00',-5.7100,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(9,1,5,5,9,1,1,0,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(10,1,5,5,75,1,1,0,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(11,1,6,6,9,1,1,0,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(12,1,6,6,75,1,1,0,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(13,1,7,7,9,2,2,0,'2018-02-13 00:00:00',2000.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(14,1,7,7,11,2,2,0,'2018-02-13 00:00:00',-2000.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(15,1,8,8,9,2,2,0,'2018-02-13 00:00:00',571.4300,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(16,1,8,8,18,2,2,0,'2018-02-13 00:00:00',-571.4300,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(17,1,9,9,9,2,2,0,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(18,1,9,9,75,2,2,0,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(19,1,10,10,9,3,3,0,'2018-02-13 00:00:00',2000.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(20,1,10,10,11,3,3,0,'2018-02-13 00:00:00',-2000.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(21,1,11,11,9,3,3,0,'2018-02-13 00:00:00',571.4300,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(22,1,11,11,18,3,3,0,'2018-02-13 00:00:00',-571.4300,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(23,1,12,12,9,3,3,0,'2018-02-13 00:00:00',50.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(24,1,12,12,52,3,3,0,'2018-02-13 00:00:00',-50.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(25,1,13,13,9,3,3,0,'2018-02-13 00:00:00',5.7100,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(26,1,13,13,53,3,3,0,'2018-02-13 00:00:00',-5.7100,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(27,1,14,14,9,3,3,0,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(28,1,14,14,75,3,3,0,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(29,1,15,15,9,3,3,0,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(30,1,15,15,75,3,3,0,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(31,1,16,16,9,4,4,0,'2018-02-13 00:00:00',3000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(32,1,16,16,11,4,4,0,'2018-02-13 00:00:00',-3000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(33,1,17,17,9,4,4,0,'2018-02-13 00:00:00',857.1400,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(34,1,17,17,18,4,4,0,'2018-02-13 00:00:00',-857.1400,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(35,1,18,18,9,4,4,0,'2018-02-13 00:00:00',50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(36,1,18,18,52,4,4,0,'2018-02-13 00:00:00',-50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(37,1,19,19,9,4,4,0,'2018-02-13 00:00:00',5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(38,1,19,19,53,4,4,0,'2018-02-13 00:00:00',-5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(39,1,20,20,9,4,4,0,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(40,1,20,20,75,4,4,0,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(41,1,21,21,6,0,0,1,'2018-02-13 00:00:00',2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(42,1,21,21,10,0,0,1,'2018-02-13 00:00:00',-2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(43,1,22,22,6,0,0,1,'2018-02-13 00:00:00',571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(44,1,22,22,10,0,0,1,'2018-02-13 00:00:00',-571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(45,1,23,23,6,0,0,1,'2018-02-13 00:00:00',50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(46,1,23,23,10,0,0,1,'2018-02-13 00:00:00',-50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(47,1,24,24,6,0,0,1,'2018-02-13 00:00:00',5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(48,1,24,24,10,0,0,1,'2018-02-13 00:00:00',-5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(49,1,25,25,6,0,0,1,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(50,1,25,25,10,0,0,1,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(51,1,26,26,6,0,0,1,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(52,1,26,26,10,0,0,1,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(53,1,27,27,6,0,0,2,'2018-02-13 00:00:00',2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(54,1,27,27,10,0,0,2,'2018-02-13 00:00:00',-2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(55,1,28,28,6,0,0,2,'2018-02-13 00:00:00',571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(56,1,28,28,10,0,0,2,'2018-02-13 00:00:00',-571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(57,1,29,29,6,0,0,2,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(58,1,29,29,10,0,0,2,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(59,1,30,30,6,0,0,3,'2018-02-13 00:00:00',2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(60,1,30,30,10,0,0,3,'2018-02-13 00:00:00',-2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(61,1,31,31,6,0,0,3,'2018-02-13 00:00:00',571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(62,1,31,31,10,0,0,3,'2018-02-13 00:00:00',-571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(63,1,32,32,6,0,0,3,'2018-02-13 00:00:00',50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(64,1,32,32,10,0,0,3,'2018-02-13 00:00:00',-50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(65,1,33,33,6,0,0,3,'2018-02-13 00:00:00',5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(66,1,33,33,10,0,0,3,'2018-02-13 00:00:00',-5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(67,1,34,34,6,0,0,3,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(68,1,34,34,10,0,0,3,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(69,1,35,35,6,0,0,3,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(70,1,35,35,10,0,0,3,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(71,1,36,36,6,0,0,4,'2018-02-13 00:00:00',3000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(72,1,36,36,10,0,0,4,'2018-02-13 00:00:00',-3000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(73,1,37,37,6,0,0,4,'2018-02-13 00:00:00',857.1400,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(74,1,37,37,10,0,0,4,'2018-02-13 00:00:00',-857.1400,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(75,1,38,38,6,0,0,4,'2018-02-13 00:00:00',50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(76,1,38,38,10,0,0,4,'2018-02-13 00:00:00',-50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(77,1,39,39,6,0,0,4,'2018-02-13 00:00:00',5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(78,1,39,39,10,0,0,4,'2018-02-13 00:00:00',-5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(79,1,40,40,6,0,0,4,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(80,1,40,40,10,0,0,4,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(81,1,41,41,10,1,1,1,'2018-02-13 00:00:00',2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(82,1,41,41,9,1,1,1,'2018-02-13 00:00:00',-2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(83,1,42,42,10,1,1,1,'2018-02-13 00:00:00',571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(84,1,42,42,9,1,1,1,'2018-02-13 00:00:00',-571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(85,1,43,43,10,1,1,1,'2018-02-13 00:00:00',50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(86,1,43,43,9,1,1,1,'2018-02-13 00:00:00',-50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(87,1,44,44,10,1,1,1,'2018-02-13 00:00:00',5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(88,1,44,44,9,1,1,1,'2018-02-13 00:00:00',-5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(89,1,45,45,10,1,1,1,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(90,1,45,45,9,1,1,1,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(91,1,46,46,10,1,1,1,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(92,1,46,46,9,1,1,1,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(93,1,47,47,10,2,2,2,'2018-02-13 00:00:00',2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(94,1,47,47,9,2,2,2,'2018-02-13 00:00:00',-2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(95,1,48,48,10,2,2,2,'2018-02-13 00:00:00',571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(96,1,48,48,9,2,2,2,'2018-02-13 00:00:00',-571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(97,1,49,49,10,2,2,2,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(98,1,49,49,9,2,2,2,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(99,1,50,50,10,3,3,3,'2018-02-13 00:00:00',2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(100,1,50,50,9,3,3,3,'2018-02-13 00:00:00',-2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(101,1,51,51,10,3,3,3,'2018-02-13 00:00:00',571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(102,1,51,51,9,3,3,3,'2018-02-13 00:00:00',-571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(103,1,52,52,10,3,3,3,'2018-02-13 00:00:00',50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(104,1,52,52,9,3,3,3,'2018-02-13 00:00:00',-50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(105,1,53,53,10,3,3,3,'2018-02-13 00:00:00',5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(106,1,53,53,9,3,3,3,'2018-02-13 00:00:00',-5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(107,1,54,54,10,3,3,3,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(108,1,54,54,9,3,3,3,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(109,1,55,55,10,3,3,3,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(110,1,55,55,9,3,3,3,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(111,1,56,56,10,4,4,4,'2018-02-13 00:00:00',3000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(112,1,56,56,9,4,4,4,'2018-02-13 00:00:00',-3000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(113,1,57,57,10,4,4,4,'2018-02-13 00:00:00',857.1400,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(114,1,57,57,9,4,4,4,'2018-02-13 00:00:00',-857.1400,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(115,1,58,58,10,4,4,4,'2018-02-13 00:00:00',50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(116,1,58,58,9,4,4,4,'2018-02-13 00:00:00',-50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(117,1,59,59,10,4,4,4,'2018-02-13 00:00:00',5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(118,1,59,59,9,4,4,4,'2018-02-13 00:00:00',-5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(119,1,60,60,10,4,4,4,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(120,1,60,60,9,4,4,4,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(121,1,61,61,4,1,0,1,'2018-02-13 00:00:00',571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(122,1,61,61,6,1,0,1,'2018-02-13 00:00:00',-571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(123,1,62,62,4,1,0,1,'2018-02-13 00:00:00',50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(124,1,62,62,6,1,0,1,'2018-02-13 00:00:00',-50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(125,1,63,63,4,1,0,1,'2018-02-13 00:00:00',5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(126,1,63,63,6,1,0,1,'2018-02-13 00:00:00',-5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(127,1,64,64,4,1,0,1,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(128,1,64,64,6,1,0,1,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(129,1,65,65,4,1,0,1,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(130,1,65,65,6,1,0,1,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(131,1,66,66,4,2,0,2,'2018-02-13 00:00:00',571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(132,1,66,66,6,2,0,2,'2018-02-13 00:00:00',-571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(133,1,67,67,4,2,0,2,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(134,1,67,67,6,2,0,2,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(135,1,68,68,4,3,0,3,'2018-02-13 00:00:00',571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(136,1,68,68,6,3,0,3,'2018-02-13 00:00:00',-571.4300,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(137,1,69,69,4,3,0,3,'2018-02-13 00:00:00',50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(138,1,69,69,6,3,0,3,'2018-02-13 00:00:00',-50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(139,1,70,70,4,3,0,3,'2018-02-13 00:00:00',5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(140,1,70,70,6,3,0,3,'2018-02-13 00:00:00',-5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(141,1,71,71,4,3,0,3,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(142,1,71,71,6,3,0,3,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(143,1,72,72,4,3,0,3,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(144,1,72,72,6,3,0,3,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(145,1,73,73,4,4,0,4,'2018-02-13 00:00:00',857.1400,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(146,1,73,73,6,4,0,4,'2018-02-13 00:00:00',-857.1400,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(147,1,74,74,4,4,0,4,'2018-02-13 00:00:00',50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(148,1,74,74,6,4,0,4,'2018-02-13 00:00:00',-50.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(149,1,75,75,4,4,0,4,'2018-02-13 00:00:00',5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(150,1,75,75,6,4,0,4,'2018-02-13 00:00:00',-5.7100,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(151,1,76,76,4,4,0,4,'2018-02-13 00:00:00',10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(152,1,76,76,6,4,0,4,'2018-02-13 00:00:00',-10.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(153,1,77,77,3,1,0,1,'2018-02-13 00:00:00',2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(154,1,77,77,6,1,0,1,'2018-02-13 00:00:00',-2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(155,1,78,78,3,2,0,2,'2018-02-13 00:00:00',2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(156,1,78,78,6,2,0,2,'2018-02-13 00:00:00',-2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(157,1,79,79,3,3,0,3,'2018-02-13 00:00:00',2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(158,1,79,79,6,3,0,3,'2018-02-13 00:00:00',-2000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(159,1,80,80,3,4,0,4,'2018-02-13 00:00:00',3000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(160,1,80,80,6,4,0,4,'2018-02-13 00:00:00',-3000.0000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(161,1,88,88,9,5,5,0,'2018-10-20 00:00:00',19.3500,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(162,1,88,88,40,5,5,0,'2018-10-20 00:00:00',-19.3500,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(163,1,89,89,9,5,5,0,'2018-10-20 00:00:00',580.6500,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(164,1,89,89,18,5,5,0,'2018-10-20 00:00:00',-580.6500,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(165,1,90,90,9,5,5,0,'2018-10-16 00:00:00',3000.0000,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(166,1,90,90,11,5,5,0,'2018-10-16 00:00:00',-3000.0000,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(167,1,91,91,9,5,5,0,'2018-10-20 00:00:00',50.0000,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(168,1,91,91,52,5,5,0,'2018-10-20 00:00:00',-50.0000,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(169,1,92,92,9,5,5,0,'2018-10-20 00:00:00',3.8700,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(170,1,92,92,53,5,5,0,'2018-10-20 00:00:00',-3.8700,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(171,1,93,93,9,5,5,0,'2018-10-20 00:00:00',10.0000,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(172,1,93,93,75,5,5,0,'2018-10-20 00:00:00',-10.0000,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211);
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
) ENGINE=InnoDB AUTO_INCREMENT=86 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerMarker`
--

LOCK TABLES `LedgerMarker` WRITE;
/*!40000 ALTER TABLE `LedgerMarker` DISABLE KEYS */;
INSERT INTO `LedgerMarker` VALUES (1,1,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(2,2,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(3,3,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(4,4,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(5,5,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(6,6,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(7,7,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(8,8,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(9,9,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(10,10,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(11,11,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(12,12,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(13,13,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(14,14,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(15,15,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(16,16,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(17,17,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(18,18,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(19,19,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(20,20,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(21,21,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(22,22,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(23,23,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(24,24,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(25,25,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(26,26,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(27,27,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(28,28,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(29,29,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(30,30,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(31,31,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(32,32,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(33,33,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(34,34,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(35,35,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(36,36,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(37,37,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(38,38,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(39,39,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(40,40,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(41,41,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(42,42,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(43,43,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(44,44,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(45,45,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(46,46,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(47,47,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(48,48,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(49,49,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(50,50,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(51,51,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(52,52,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(53,53,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(54,54,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(55,55,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(56,56,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(57,57,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(58,58,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(59,59,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(60,60,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(61,61,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(62,62,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(63,63,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(64,64,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(65,65,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(66,66,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(67,67,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(68,68,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(69,69,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(70,70,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(71,71,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(72,72,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(73,73,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(74,74,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(75,75,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-07-03 02:45:37',211,'2018-07-03 02:45:37',211),(76,0,1,1,0,0,'2018-01-30 00:00:00',0.0000,3,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(77,0,1,1,1,0,'2018-01-30 00:00:00',0.0000,3,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(78,0,1,2,0,0,'2018-01-30 00:00:00',0.0000,3,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(79,0,1,2,2,0,'2018-01-30 00:00:00',0.0000,3,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(80,0,1,3,0,0,'2018-01-30 00:00:00',0.0000,3,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(81,0,1,3,3,0,'2018-01-30 00:00:00',0.0000,3,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(82,0,1,4,0,0,'2018-01-30 00:00:00',0.0000,3,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(83,0,1,4,4,0,'2018-01-30 00:00:00',0.0000,3,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(84,0,1,5,0,0,'2018-10-20 00:00:00',0.0000,3,'2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(85,0,1,0,0,41,'1970-01-01 00:00:00',0.0000,3,'2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211);
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PaymentType`
--

LOCK TABLES `PaymentType` WRITE;
/*!40000 ALTER TABLE `PaymentType` DISABLE KEYS */;
INSERT INTO `PaymentType` VALUES (1,1,'Cash','Cash','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(2,1,'Check','Personal check from payor','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(3,1,'VISA','Credit card charge','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(4,1,'AMEX','American Express credit card','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0),(5,1,'ACH','','2017-11-10 23:24:23',0,'2017-11-10 23:24:23',0);
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
INSERT INTO `Payor` VALUES (1,1,'007663c575d2ef50a2e5936fe9f88fc761b4c6ebd3293cad4051f5d63d56c418145db97f',5563.0000,1,0,'84f1af725145509fce028f022a28dfdfbd81bb48352280e3f4701e81ee42bb9b1d618862',73697.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,1,'8db365f9a9565868587dfafb783d5fe593eac6908cc0bbcf334921512f00712ec5a39e78',14729.0000,1,0,'05ec07db893b1d9c85758185b7267738e466a09f629d8808da44663d060a234644d9fc6f',99563.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,1,'be0264bd86c93fd1f2320b186190cd6df8fb1d1be088b799888210356b4f15d0bade67a5',8355.0000,1,0,'05226fe2b7789eae3ae0811b3da782e56c517af42410a72e3f574c6a97036b36be155b8a',12815.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,1,'d9a9d75523983c24b9569afeca1a8376eb18277e6ffeed6d7367f0095fc2c3f14b28b102',12034.0000,1,0,'f6bdabde02cd5e8011ca352662790447d24260cf5595516010295bd99aeadbcd8eabf8bb',116086.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(5,1,'00886cd68b418081ab0eb5eab4d8052c9ce54250a6e3b09520aa7f53b0ab4d0d6f5cd3f5',28733.0000,1,0,'ca5cb000c87256b7f40820aac0fc3491a255005c0454fc1264833d5e8f2b8c3821b607f2',73876.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(6,1,'48dbcb5c9bc1938d96e8d05ef7d6a9c33381a66af37a6ca34963de74a64ae8cfe2d01313',11518.0000,1,0,'8925d7c3669341162f7e205c4243ad0152960e3c5f8a907fd3166a5b0981e7eaee30c04f',56963.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(7,1,'27c481b805b803464efb063932dd255eb4efe54b08e4c555ca16c6470655cff72a63f3ae',14541.0000,1,0,'49ef2ac7590784d1b42e931363bb883317a9877fa165cd25d95e642c9ad5ddb9b9b12c6f',26757.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(8,1,'99e970ec98bbf41d7518919c7899dbc78b949c93e5a68e168020da636ebc2a7bbad5db09',9856.0000,1,0,'63a17756420cbcd95ac250c2da310b0d56036f8463d8b5b0662dc2be212acc46c2cb03cf',83398.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(9,1,'95d2d6820b3dff91111cb992af565680567ce12410133393520becfba3dc160129435538',24661.0000,1,0,'f010d48ca26aa377141f10ba663312d40c01ab4c361f9db7623425989ab100fe5e826963',12206.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(10,1,'ed64e83ae095f6d7553c8f6eb465a912cd4e788f3162727090450eddb40a03387c2a701f',1075.0000,1,0,'0fd78ba2870dd37ef89b9695d00d97a256ba8c753a36288ac814ceb4e50db89d1abcb91f',63453.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(11,1,'28c46c589ecb595f418b5d0c09ded8819f7bcfd41498d08a3f67a6c69cc7a68ac5b93077',6261.0000,1,0,'75fa7c8e9728b527388f8ec8aad44071852fb580eee2a8348d145e13c2d5e1a53a675d32',111482.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(12,1,'245710c289339e6eaa166ca8b1f3f00fe09e771ce50104bf7305cedfc4979bf2f108563e',5710.0000,1,0,'cc9bc1e69c50371eb883c6821d9fcd4d643d57e2f6b7d3cd1f28f4b886aebc83b6447950',34703.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(13,1,'7e17c75e8ae9aca316b1e07912ff1871a697fa6a9185ed14c3c5ddf02e8fd2fb68777f74',10079.0000,1,0,'66be319d498cc58ec3c6c8794c1685b4bc45c54a9605e4b9bf409762fe7a92656a01d258',100244.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(14,1,'b0af9d8b237233afde32778b3bb2735c423861e6757c9305ecb8cdd1a84520b58c7f74a3',24180.0000,1,0,'2089789e24b330184050c4ed9ef9b9179f22f797558ab65bf864388a8fd6b65d29bf9774',144526.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(15,1,'1d91e6b2afc41bc923e28db48d84210d81bc9a10ef633d38c733163d415e22fb811370cf',25659.0000,1,0,'2e7210eb66d724e024fc8837ce27c373ed28f57282223023c56394ec2e89aad7fca69932',86669.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(16,1,'7a8ee594d3b74db33e54ad5e5a1253d83f69672aa1e439743be577554c5bfaf7becb4fa8',18565.0000,1,0,'8e6a98ce76927d645ab96cd76b02ea04d4940b2dd4509a030e4a551ec008c04e39cfd002',131384.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(17,1,'a6365cde0aeaf5b40bdf537f2b0167e0bcab1044d5512eefb55bac17e767cebd6a98a286',21394.0000,1,0,'5dda3d18618573878a76ff4326607534e8a562a6143e956c26e2fe430ef7d7af178cf423',104020.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(18,1,'558475371c254a5a39574b01bf0b4e112b68fc19827bab3526ab1803cf08b0e9697f23ef',24934.0000,1,0,'573983243983d409ddf1ca8eff47a20832d6644c39678019dc65ef03b469357b2c6f17a5',40401.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(19,1,'aee5185e5115ab2d4a0d6736495af2edb0a89bed29bec3c3740adbce799ce1e348101505',28496.0000,1,0,'d2dc2a706d3b5222782bcf30e447af7c60cd9a37f98a5471c6132c400fe5dc3c43a2976b',79907.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(20,1,'754ebc7de9fca59937954c4e02976c46ff993c0b9b70ffbac29ebc7aedf67faaed49582d',22850.0000,1,0,'71dfb7b213e95c367222342ccd6278dd92bd1577df4a91d15db670f0b16167fc5013cc84',73838.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(21,1,'da1368a9205a368d6291eef70c4d1c39d251534569a6603bd5606d55cda9f633d8bf2b1d',17311.0000,1,0,'95cc862963ec7ab6ec29dc0a6bc7ffa499fcd6a37a522fc5b4e6bbb445e5b2e5fb6404ce',19739.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(22,1,'d0647fed2b22ea3b7006f24bc258cd0e85adb3ae005d40125f9a0bdeee8274a7a0cec869',15013.0000,1,0,'bdfb05193cda39acfb1aecdf35e0ebc03714a7266180d02c33510d3b9856b2ad120ce3db',50777.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(23,1,'eb00f13a6526834efb3303ebdb263b3896e893db2f6fc3a26faa1b66533b6616a33d4358',29717.0000,1,0,'6175797be9be2ddc4b728b80069ef2deb670c2da5d0d0043441960a77449dff11c04d0c8',90134.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(24,1,'a0b04c23d51f771ece20f9f6510d42ebad1480931176af02e17d076472d6ca7254c464ae',28228.0000,1,0,'33ae2228d89d612f4171fc1a465227341c839b682ab254ef86f2f28f6100866eb15bc8c5',39646.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(25,1,'358b1f82587ce5d5186e52f9ad708b9d7f27733331bfe2abcbfdf86a3bf7d70b9bf85844',18037.0000,1,0,'4311cebf136b3f6b1435215cef695b957d14d7ae228dbef41b0544199edddc41dd13c088',30265.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(26,1,'cfb43e98aeb8adaa6e0c1cf5cd9d0baf3bbf6d072b57a5c141c87f624093d7dfa7c7f558',17940.0000,1,0,'05b737e80947c89510dd56305bdec2819a70c4ed91046c77c0202b699ffdb65384a73e02',64579.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(27,1,'f7fd8a4b4b2f6df53a8ff36c06d515e0434c726899fb15aeefc9a80a45079a6e8ce19b9d',4031.0000,1,0,'acb82c29dc53b16fc77f9bb96531f87ad1e7520ee2456e7fbc947f872515e1e6cf16e3ff',26604.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(28,1,'a0e954f94e63b175a2f2c1b2d43a5b5d13e7f08f142ef90b7d356c79c9e031ffeeca4c4c',16188.0000,1,0,'9c67f9045fc83454dd2518052fd2830f92568bff90e68802a5780e7c2cc5eba9166214e7',97740.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(29,1,'84bed19a397183cf0ec7d6e879a34acd87c89437b4240edb2aeb254dbddb6887b9324ff8',21459.0000,1,0,'da4685f508a368b9305b1837b1a0793924a61cf9b7e539a4afcf10233e4cd4d3a0d9f4b1',51311.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(30,1,'344ace43541bd2bfe6a1085e406e8eb7437f8c2fbd5bc12b11c9a1bf529a5e62088c942c',29372.0000,1,0,'9e9329eb53cc77d0a25e77a88a8b53dcc2dc3612c2bcd0012600f30d5e7174507871478a',35813.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(31,1,'ff89dd195d1d78ae89972cb23b73e5f7755d3df2e200dd2eaab8ed8b47d69d98acc87407',8660.0000,1,0,'433ce50b539bd6fca0ea6017ad5768bfb456b9a09b493b121da50d7b9e39244ed85fd094',103768.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(32,1,'17768350fd0019b2e08982b226894c7f926f30a97e7d6c50cfcf88e0d5df5d2425e6cb77',26871.0000,1,0,'389f66ad715010ba774263d95ce65beb0fc218ae034e85a2e301ddeac8fedaab86dba2a4',125332.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(33,1,'42b70716e24e01ef1ee388eff506be19ca5498b8327e259471a0bec50e25a3dd9adbc2d0',17693.0000,1,0,'409e01652596e7de38dc2e1e4290772e2f0d8339608fe0d13af8e487dce903fb09e4f210',118050.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(34,1,'1c67d93f26c8839c66de413adc50dbc839e81a86cb87680ca209e9bdfda1765c76654e94',7181.0000,1,0,'d6d1c3aaecde6b2994e281c5da6a35c468807d29973def36ae202fb1fa83de981abe11f5',122094.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(35,1,'5fe19360382a63337fe2701cfdf578ae485369aea26c3e038fd2e20a865aa1c484f93bb4',25547.0000,1,0,'c8e1f95e172475f0b41c355b4e7a1296741f976502b7f6edcd6394995b8183366f1e373d',143858.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(36,1,'21933ee15a668acf7a57b277a01ad690926a0e05ab6bf9ba33515cd03add78f6feb15c18',8971.0000,1,0,'3eda9938d2fce0de6e492163cb114b601dbbb10c01553c0a1655e134ecc5e2c20d89d343',99369.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(37,1,'2e39d9c3bbaaf1bf6fc81177ed3f81caca63a7d2184f729a174771c01e3cf2bdea6d27d3',9674.0000,1,0,'bf759572482af0c375e48b152cce994c061f875e88d175317b2ecfa281382265beeec65d',10443.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(38,1,'7f3eec8a6e67c4af1e1f921a8d52bb656f33d06fea93f400f2459aa580245e70d38a190f',19527.0000,1,0,'ff218d040571bce70b838bc224289ae3a8c5bdd78c718500471246777acfd840c3114ce9',73111.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(39,1,'9b307f5e139abe98debac7d02558f65f502564d8b161b83e2622d2ea8a288e77d4de037f',25619.0000,1,0,'4d7eab2395b1f5f43a399b5c737e91142dde90b439659ceaa4aed61250e6081b3efdf070',81565.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(40,1,'b0efe1870bf34e9d70f66d852ec33e3422ff81796bf8c2ff1658dde5029933a49f4c5d5a',16130.0000,1,0,'df7bdb8f8743dc2b62cf3490354cf6bd763b2cf0c442602895a4a4d0b13b6c9aad62c093',132348.0000,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(41,1,'0f5b287f48159603e60044081f81707ff1c2d8f069e4e07ed9ee61aebc70719d4b0bd3836ac04a',0.0000,1,0,'3e081969f0eb9b1fd10b91e35ededc06702e909e39800c3f30fc68b5e3a799ed9c8bcae8b4',68000.0000,'2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211);
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
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Pets`
--

LOCK TABLES `Pets` WRITE;
/*!40000 ALTER TABLE `Pets` DISABLE KEYS */;
INSERT INTO `Pets` VALUES (1,1,1,1,'dog','English Cocker Spaniel','Grey',40.0000,'Oreo','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,1,3,3,'dog','Bullmastiff','Harlequin',20.0000,'Gus','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,1,4,4,'cat','Siberian cat','tortoiseshell',23.0000,'Boots ','2018-02-13','2020-03-01','2018-10-16 19:35:32',0,'2018-10-16 19:35:31',0),(4,1,0,5,'dog','Welsh Springer Spaniel','Grey',40.0000,'Bruno','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(5,1,0,6,'cat','Donskoy cat','seal point',21.0000,'Smokey ','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(6,1,0,7,'dog','Shiba Inu','Red',35.0000,'Brady','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(7,1,0,7,'dog','Doberman Pinscher','Harlequin',34.0000,'Champ','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(8,1,0,9,'cat','Mexican Hairless Cat','lynx point',23.0000,'Lucky ','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(9,1,0,13,'dog','Komondor','Tuxedo',50.0000,'Yoda','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(10,1,0,13,'dog','Cairn Terrier','Spotted',32.0000,'Joey','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(11,1,0,16,'cat','American Ringtail','Dilute',10.0000,'bailey ','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(12,1,0,16,'cat','Thai Lilac','chocolate point',11.0000,'SUGAR ','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(13,1,0,17,'dog','Greater Swiss Mountain Dog','Spotted',50.0000,'Tucker','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(14,1,0,18,'dog','Miniature Schnauzer puppy','Brown',34.0000,'Scout','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(15,1,0,19,'dog','Lhasa Apso','Harlequin',52.0000,'Frankie','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(16,1,0,20,'cat','Abyssinian cat','Harlequin',7.0000,'Charlie ','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(17,1,0,23,'cat','Sokoke','lynx point',12.0000,'Misty ','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(18,1,0,23,'cat','Birman','Tuxedo',15.0000,'Fluffy ','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(19,1,0,25,'dog','Belgian Sheepdog','Gold',28.0000,'Brutus','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(20,1,0,26,'dog','Border Terrier','Fawn',36.0000,'Bentley','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(21,1,0,27,'cat','Foldex cat','white',8.0000,'Garfield ','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(22,1,0,28,'dog','Spinone Italiano','Tricolor',30.0000,'Diesel','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(23,1,0,29,'dog','Affenpinscher','Tricolor',21.0000,'Ziggy','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(24,1,0,31,'cat','Persian cat','Dilute',6.0000,'mittens ','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(25,1,0,32,'dog','Scottish Deerhound','Harlequin',22.0000,'Shadow','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(26,1,0,35,'dog','Tibetan Mastiff','Fawn',26.0000,'Duke','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(27,1,0,36,'dog','Cesky Terrier','Fawn',16.0000,'Cash','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(28,1,0,38,'cat','Isle of Man Longhair','Smoke',22.0000,'Angel ','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(29,1,0,39,'dog','American Eskimo Dog (Toy)','Blue',39.0000,'Brody','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(30,1,0,40,'dog','Azawakh','Cream',50.0000,'Joey','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(31,1,5,41,'dog','beagle','tri-color',30.0000,'Barney','2018-10-20','9999-12-31','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211);
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
INSERT INTO `Prospect` VALUES (1,1,'82595 Park','FortWayne','WV','84059','CapitalOneFinancialCorpF3900@gmail.com','(895) 638-0659','sound effects technician',0,'','','','','','0000-00-00','8144 Aloha, New Orleans, IL 02081','Otis Oliver','(562) 975-6102',84,'7 years 8 months','36865 Oak, Bethlehem, CO 41318','Carolina Wyatt','(511) 768-4962',128,'7 years 2 months','','Laci Guthrie','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,1,'19805 Fourth','NewBedford','MN','03300','ONewBedford3691@yahoo.com','(761) 861-1856','tailor, dressmaker (tailor and dressmaker)',0,'','','','','','0000-00-00','95559 Dogwood, Sioux Falls, NE 10694','Mickie Arnold','(151) 573-1434',116,'9 years 2 months','46011 11th, Houma, NM 78511','Shavon Trujillo','(959) 628-6397',89,'2 years 2 months','','Meda Jarvis','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,1,'2750 Ridge','Chicago','WA','33274','ActernaCorpC7870@abiz.com','(300) 109-6273','safety and communication electrician',0,'','','','','','0000-00-00','97394 Main, Tallahassee, NM 11211','Julene Oneal','(892) 198-0495',92,'3 years 9 months','46350 South Carolina, Boulder, ME 31445','Ashanti Steele','(660) 625-2127',104,'9 months','','Carmella Garner','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,1,'67414 Lee','Plano','NJ','65466','UPlano1698@abiz.com','(187) 948-6345','textiles dyer',0,'','','','','','0000-00-00','95888 Center, Paterson, MI 11528','Sunshine Howard','(826) 635-9941',93,'6 years 9 months','84663 Pine, Columbia, MT 86258','Elisabeth Lawrence','(233) 725-2068',110,'7 years 11 months','','Loria Gamble','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(5,1,'75371 Sycamore','Bremerton','IN','32888','McLeodUSAIncorporatedBremerton312@abiz.com','(795) 299-3904','knitter',0,'','','','','','0000-00-00','80657 8th, Tyler, MI 92790','Amy Chambers','(102) 749-1700',127,'7 years 6 months','81504 Pinon, Lorain, WI 93015','Lucilla Graham','(499) 819-1610',107,'7 years 7 months','','Isabella Hays','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(6,1,'60895 Aspen','SantaAna','SD','66831','RSantaAna5903@abiz.com','(587) 354-4343','systems engineer (manufacturing)',0,'','','','','','0000-00-00','79095 9th, Fremont, NY 15429','Sulema Charles','(223) 617-7767',91,'7 months','10048 Eleventh, Santa Clarita, ND 25356','Heide Myers','(805) 955-6564',117,'10 years 4 months','','Vinita Walls','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(7,1,'6376 Pecan','Phoenix','MI','15026','StarwoodHotels&ResortsWorldwideIncP5656@bdiddy.com','(787) 742-0806','referee (umpire)',0,'','','','','','0000-00-00','10730 County Line, Lorain, MN 86413','Branden Noble','(809) 824-2231',110,'8 years 11 months','96775 Church, Rancho Cucamonga, AL 03090','Louann Morrison','(404) 336-0971',123,'4 years 9 months','','Dia Collier','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(8,1,'74084 7th','BonitaSprings','OR','24147','EcolabIncBonitaSprings69@comcast.net','(561) 464-8369','photographic reporter (press photographer)',0,'','','','','','0000-00-00','33021 Pleasant, Albany, CA 74078','Millie Brock','(933) 331-9918',97,'4 months','43571 Dogwood, Santa Ana, NY 64324','Mendy Decker','(301) 268-9926',99,'5 years 11 months','','Eduardo Reed','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(9,1,'65197 Main','Fargo','VA','43721','SFargo2158@gmail.com','(180) 523-4423','rolling-mill operator (rolling-mill worker)',0,'','','','','','0000-00-00','16664 Laurel, El Paso, ME 27189','Berniece Crane','(839) 692-2273',106,'4 months','60350 Laurel, Eugene, PA 82199','Alva Leonard','(413) 185-8687',91,'1 years 7 months','','Moshe Mckenzie','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(10,1,'46702 Hemlock','SiouxCity','WA','04538','RockwellAutomationIncSiouxCity203@gmail.com','(411) 172-8185','stuntman (stuntwoman)',0,'','','','','','0000-00-00','47735 Apache, St. Paul, FL 49703','Erin Miller','(884) 580-4994',111,'9 years 8 months','73125 Center, Kenosha, PA 89355','Veronica Clemons','(344) 433-2751',105,'5 years 3 months','','Annie Brady','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(11,1,'40162 Maple','Winston','DC','60156','SWinston3802@yahoo.com','(857) 675-6177','boilermaker/fitter',0,'','','','','','0000-00-00','95248 Hemlock, Tucson, MI 28266','Marivel Dixon','(616) 547-6394',127,'8 years 7 months','30509 South Dakota, Olathe, MT 89828','Janine Valencia','(225) 610-9767',121,'5 years 2 months','','Ching Keith','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(12,1,'65672 Narragansett','MissionViejo','SC','35746','URSCorporationM4987@abiz.com','(140) 434-7866','fortune teller',0,'','','','','','0000-00-00','15057 2nd, Corpus Christi, WV 71563','Tyree Mcneil','(348) 826-7479',87,'5 years 2 months','59634 A, Lancaster, IN 14376','Margene Hartman','(396) 400-7456',111,'8 months','','Deidra Maddox','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(13,1,'67141 Lake','LongBeach','VA','65094','LLongBeach6988@yahoo.com','(548) 625-8367','primary school teacher',0,'','','','','','0000-00-00','63889 Hickory, Sarasota, ND 61577','Jinny Mcguire','(255) 434-9498',115,'6 years 4 months','52273 Mesquite, Sioux Falls, IN 07463','Rudolf Sherman','(868) 936-3557',113,'7 years 9 months','','Janette Short','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(14,1,'40275 Airport','Mesquite','KS','60953','SMesquite6108@yahoo.com','(639) 783-6407','power station/systems operator (energy s equipment operator)',0,'','','','','','0000-00-00','60384 Church, Fairfield, SC 71137','Tawna Vargas','(711) 416-9034',86,'4 years 11 months','22819 Bay, Columbus, TX 43133','Caroline Woods','(154) 256-7281',95,'7 years 3 months','','Jeanelle Curtis','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(15,1,'94335 Willow','Normal','NH','38643','GreenPointFinancialCorpN5307@aol.com','(870) 780-1002','rail transport worker',0,'','','','','','0000-00-00','2216 Lake, Thornton, AL 53891','Isadora Craig','(390) 856-5930',128,'10 years 9 months','80082 Aspen, Independence, CT 02002','Charlie Perez','(233) 941-7892',99,'2 years 10 months','','Benny Branch','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(16,1,'89277 Broadway','Kenosha','IL','09107','TheMayDepartmentStoresCompanyKenosha164@comcast.net','(299) 450-4199','stable hand, groom',0,'','','','','','0000-00-00','81438 Hickory, Tacoma, AZ 07940','Nisha Kaufman','(934) 179-5821',126,'4 years 7 months','56128 4th, Harlingen, DC 06503','Deeann Hebert','(651) 744-4238',90,'7 years 3 months','','Juliann Flores','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(17,1,'51022 12th','Richmond','MD','61598','AudiovoxCorporationRichmond261@comcast.net','(460) 907-0621','flight attendant (steward/ess, cabin staff)',0,'','','','','','0000-00-00','99356 10th, Lorain, RI 67425','Brett Crosby','(822) 888-8925',121,'9 months','67516 3rd, Sebastian, CA 81351','Sidney Mcconnell','(435) 479-9667',119,'4 years 8 months','','Antonetta Johns','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(18,1,'49671 New Mexico','KansasCity','ND','58010','GenuityIncK8627@hotmail.com','(640) 763-2534','heating and ventilating fitter (air-conditioning fitter)',0,'','','','','','0000-00-00','70492 Maple, Coral Springs, KS 03410','Jana Moran','(222) 816-5749',114,'11 months','30179 Sunset, Boston, NM 95285','Kathline Mills','(934) 434-9801',119,'11 months','','Stephany Underwood','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(19,1,'42278 Airport','SimiValley','MO','48553','BSimiValley8717@hotmail.com','(532) 533-7954','fish warden (water keeper/bailiff)',0,'','','','','','0000-00-00','53001 Hill, Little Rock, ID 18591','Glendora Dickson','(333) 615-3543',83,'1 years 5 months','76984 Quail, Escondido, ND 88582','Analisa Logan','(891) 652-0138',99,'2 years 7 months','','Frank Downs','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(20,1,'24661 3rd','ElkGrove','MS','86137','GElkGrove5060@aol.com','(871) 866-2478','dust control technician',0,'','','','','','0000-00-00','82471 Walnut, Santa Ana, CT 69271','Sherie Boyd','(189) 273-2437',98,'4 years 4 months','26801 West Virginia, Nashville, OK 15894','Carmel Williams','(577) 715-9270',103,'7 years 5 months','','Milagro Burns','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(21,1,'5793 Zoo','Denton','DE','12079','ADenton5889@hotmail.com','(575) 685-6702','keeper of records (archivist)',0,'','','','','','0000-00-00','78870 New Mexico, Colorado Springs, MA 12066','Gretchen Levine','(996) 350-7598',126,'2 years 5 months','31710 10th, Pasadena, MS 21270','Krystyna Rogers','(299) 213-4819',96,'3 years 3 months','','Mikki Cline','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(22,1,'54255 7th','Brownsville','IL','98981','FBrownsville9309@aol.com','(632) 867-2101','electroceramic production operative',0,'','','','','','0000-00-00','98088 North, Saint Petersburg, CT 06052','Reed Mueller','(886) 719-9677',116,'8 years 10 months','39031 Hampton, Simi Valley, MI 17175','Chin Bauer','(356) 799-4829',115,'2 years 3 months','','Roseanne Greer','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(23,1,'20630 Pioneer','StPetersburg','VT','73749','Albertson\'sIncStPetersburg207@abiz.com','(796) 386-4098','worker in the paper industry',0,'','','','','','0000-00-00','97582 Park, Apple Valley, DE 01528','Genna Baldwin','(663) 219-2243',113,'7 years 5 months','60468 Meadow, Pembroke Pines, MN 92818','Quinn Aguirre','(788) 101-2819',107,'8 years 6 months','','Marvella Stafford','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(24,1,'81325 County Line','SouthBend','KY','64547','JackInTheBoxIncSouthBend752@aol.com','(588) 621-9816','woodcutting manager',0,'','','','','','0000-00-00','81433 New Hampshire, Boston, NJ 93612','Daysi Ortiz','(981) 587-8284',127,'2 years 8 months','30198 West Virginia, High Point, CT 21532','Easter Mccarty','(562) 873-9386',94,'2 years 6 months','','Deandra Hays','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(25,1,'82219 Thirteenth','Houston','AZ','25786','BigLotsIncH2132@comcast.net','(139) 914-1161','metal refiner',0,'','','','','','0000-00-00','728 Evergreen, New London, WY 47051','Elizabeth Myers','(832) 937-5344',88,'7 years 2 months','67484 Cedar, Anaheim, AK 58076','Selene Santiago','(767) 685-0657',97,'5 years 8 months','','Marianela Castro','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(26,1,'99534 1st','Fresno','SC','90364','EnergyEastCorporationFresno395@bdiddy.com','(554) 427-1093','postmaster/postmistress',0,'','','','','','0000-00-00','35482 North, Saint Paul, ID 92305','Eva Brown','(320) 108-2767',92,'4 years 10 months','92851 Narragansett, Waterloo, CA 99183','Mariah Sandoval','(861) 778-7857',119,'9 years 8 months','','Elvina Hopkins','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(27,1,'2280 Ridge','Dayton','MD','92258','HDayton1960@abiz.com','(983) 289-5901','varnisher (painter-varnisher)',0,'','','','','','0000-00-00','52350 Shore, Cleveland, NE 63767','Mika Emerson','(344) 574-0372',108,'10 years 10 months','34636 S 400, Lafayette, WA 43231','Mark Patterson','(994) 642-1757',120,'8 years 5 months','','Hsiu Sears','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(28,1,'47620 Navajo','Kenosha','OR','81223','ProgressiveCorporationK4473@comcast.net','(836) 354-3319','speech therapist',0,'','','','','','0000-00-00','63665 Kukui, Sioux City, MS 17342','Sheri Howell','(164) 383-3521',126,'3 years 11 months','85887 Bay, Fresno, WV 04208','Carlotta Stanton','(536) 667-1779',95,'7 years 10 months','','Nanci Jarvis','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(29,1,'1635 Dogwood','Clarksville','NJ','53710','SaraLeeCorpC8846@abiz.com','(316) 818-9229','municipal services worker (communal service worker)',0,'','','','','','0000-00-00','2340 Ninth, Greensboro, DC 94535','Hoyt Cannon','(576) 497-7358',91,'7 years 10 months','96397 Broadway, McHenry, CO 90440','Vivian Perez','(594) 362-9062',110,'5 years 5 months','','Kory Lott','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(30,1,'81930 Cherry','Pasadena','KS','34415','EPasadena4493@aol.com','(280) 255-4604','fine artist',0,'','','','','','0000-00-00','22258 Lincoln, Fullerton, CA 89371','Renae Mayer','(538) 892-6604',104,'9 years 11 months','87765 Johnson, Carrollton, DC 83039','Phil Macias','(578) 349-6598',99,'8 years 11 months','','Glynda Rivas','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(31,1,'92521 Pine','Sunnyvale','AK','61359','Anheuser-BuschCompaniesIncSunnyvale66@gmail.com','(771) 678-6479','career diplomat/ diplomat',0,'','','','','','0000-00-00','57518 A, Hemet, AK 96720','Flo Fischer','(743) 681-4063',124,'7 years 11 months','10802 Sycamore, Bellevue, AZ 40783','Xiao Hayes','(152) 818-2320',127,'2 years 9 months','','Kathie Leonard','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(32,1,'45893 Lehua','CedarRapids','LA','58010','ECedarRapids4722@bdiddy.com','(832) 795-6683','actuary',0,'','','','','','0000-00-00','49595 South Dakota, Tempe, LA 40565','Audry Romero','(362) 108-1818',126,'1 years 3 months','7120 Church, Provo, KS 94162','Earnestine Whitaker','(738) 416-8291',118,'10 years 5 months','','Norah Wilcox','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(33,1,'20461 Wood','Lafayette','DE','56756','LucentTechnologiesIncLafayette777@aol.com','(364) 230-6713','well digger',0,'','','','','','0000-00-00','4770 Pleasant, Mesa, PA 95695','Kristen Mcgowan','(161) 193-2810',86,'3 years 9 months','27774 Jackson, Lake Charles, CA 78666','Miguel Rivera','(412) 376-7897',98,'6 years 2 months','','Stanton Park','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(34,1,'25220 Twelfth','Pomona','NY','21092','Pennzoil-QuakerStateCompanyP9606@bdiddy.com','(841) 323-4364','glass jewellery maker',0,'','','','','','0000-00-00','54408 Pioneer, South Lyon, MD 78831','Alan Gomez','(950) 395-6815',125,'7 years 2 months','8849 Eleventh, Bremerton, DC 97577','Michelle Patterson','(475) 255-3542',107,'5 years 6 months','','Ofelia Bates','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(35,1,'82339 Highland','StLouis','PA','71162','PeoplesEnergyCorpS9167@abiz.com','(783) 895-1465','airline clerk (airline ticket agent)',0,'','','','','','0000-00-00','35382 Kukui, Vallejo, MS 43447','Krysten Griffith','(461) 859-3908',126,'9 months','55996 Lakeview, Hayward, AR 90292','Georgia Gaines','(234) 382-8761',88,'1 years 3 months','','Clelia Stout','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(36,1,'49909 A','Tyler','ND','69888','ATyler3543@hotmail.com','(908) 113-4137','travel courier (tourist guide)',0,'','','','','','0000-00-00','70109 Pleasant, Fort Wayne, VA 38318','Glenn Lang','(791) 632-4149',108,'2 years 9 months','90882 Johnson, Escondido, MS 23756','Larae Terrell','(100) 836-3038',108,'7 years 2 months','','Lavinia Thompson','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(37,1,'32169 Wilson','Hesperia','VT','58652','SearsRoebuck&CoHesperia728@aol.com','(932) 215-8494','primary school teacher',0,'','','','','','0000-00-00','40059 Zoo, Jersey City, MD 58675','Ellen Cortez','(701) 713-4801',86,'1 years 7 months','19847 Elm, Youngstown, IN 22181','Tanner Velazquez','(331) 610-5922',87,'9 years 11 months','','Valentina Lang','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(38,1,'29831 Willow','Raleigh','UT','01393','UniversalCorporationR5779@hotmail.com','(293) 303-5770','telecommunications technician',0,'','','','','None','0000-00-00','71664 Lincoln, Orange, OR 78470','Buffy Dillon','(753) 776-4855',110,'1 years 10 months','97031 Canyon, Eugene, MS 78996','Asley Benson','(636) 987-0692',98,'10 years 4 months','','Maryalice Hunter','2018-10-17 06:23:18',211,'2018-10-16 19:35:31',0),(39,1,'85552 Maple','Fayetteville','OK','60260','CFayetteville5451@aol.com','(469) 784-8921','storekeeper (warehouse keeper)',0,'','','','','','0000-00-00','49328 Aloha, Anaheim, AK 73922','Trisha Henson','(117) 837-4535',104,'10 years 6 months','99503 11th, Lacey, NV 47029','Hilario Knowles','(579) 447-9018',89,'1 years 6 months','','Adrianne Maddox','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(40,1,'56838 Oak','Spartanburg','WV','31079','USpartanburg9994@abiz.com','(932) 967-0908','boiler operator (boiler attendant)',0,'','','','','','0000-00-00','82002 Hillside, Davidson County, DE 92954','Dotty Weeks','(203) 931-4462',106,'8 years 8 months','5281 Apache, Mission Viejo, OH 31464','Audrey Vaughn','(556) 921-6131',107,'8 years 7 months','','Linh Santiago','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(41,1,'','','','','','','Pilot',0,'','','','','None','0000-00-00','123 Elm Street','Emmit Bailer','987-654-3210',123,'2 years','','','',0,'','','','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211);
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
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Receipt`
--

LOCK TABLES `Receipt` WRITE;
/*!40000 ALTER TABLE `Receipt` DISABLE KEYS */;
INSERT INTO `Receipt` VALUES (1,0,1,1,2,2,2,1,'2018-02-13 00:00:00','157480',2000.0000,'',25,'ASM(2) d 12999 2000.00,c 12001 2000.00',2,'payment for ASM-2','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(2,0,1,1,2,1,1,1,'2018-02-13 00:00:00','196514',571.4300,'',25,'ASM(4) d 12999 571.43,c 12001 571.43',2,'payment for ASM-4','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(3,0,1,1,2,1,1,1,'2018-02-13 00:00:00','82595',50.0000,'',25,'ASM(5) d 12999 50.00,c 12001 50.00',2,'payment for ASM-5','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(4,0,1,1,2,1,1,1,'2018-02-13 00:00:00','197605',5.7100,'',25,'ASM(6) d 12999 5.71,c 12001 5.71',2,'payment for ASM-6','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(5,0,1,1,2,1,1,1,'2018-02-13 00:00:00','929262',10.0000,'',25,'ASM(7) d 12999 10.00,c 12001 10.00',2,'payment for ASM-7','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(6,0,1,1,2,1,1,1,'2018-02-13 00:00:00','809866',10.0000,'',25,'ASM(8) d 12999 10.00,c 12001 10.00',2,'payment for ASM-8','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(7,0,1,2,2,2,2,2,'2018-02-13 00:00:00','177687',2000.0000,'',25,'ASM(10) d 12999 2000.00,c 12001 2000.00',2,'payment for ASM-10','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(8,0,1,2,2,1,1,2,'2018-02-13 00:00:00','632933',571.4300,'',25,'ASM(11) d 12999 571.43,c 12001 571.43',2,'payment for ASM-11','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(9,0,1,2,2,1,1,2,'2018-02-13 00:00:00','842890',10.0000,'',25,'ASM(12) d 12999 10.00,c 12001 10.00',2,'payment for ASM-12','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(10,0,1,3,2,2,2,3,'2018-02-13 00:00:00','551237',2000.0000,'',25,'ASM(14) d 12999 2000.00,c 12001 2000.00',2,'payment for ASM-14','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(11,0,1,3,2,1,1,3,'2018-02-13 00:00:00','68528',571.4300,'',25,'ASM(16) d 12999 571.43,c 12001 571.43',2,'payment for ASM-16','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(12,0,1,3,2,1,1,3,'2018-02-13 00:00:00','539462',50.0000,'',25,'ASM(17) d 12999 50.00,c 12001 50.00',2,'payment for ASM-17','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(13,0,1,3,2,1,1,3,'2018-02-13 00:00:00','897323',5.7100,'',25,'ASM(18) d 12999 5.71,c 12001 5.71',2,'payment for ASM-18','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(14,0,1,3,2,1,1,3,'2018-02-13 00:00:00','486864',10.0000,'',25,'ASM(19) d 12999 10.00,c 12001 10.00',2,'payment for ASM-19','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(15,0,1,3,2,1,1,3,'2018-02-13 00:00:00','481923',10.0000,'',25,'ASM(20) d 12999 10.00,c 12001 10.00',2,'payment for ASM-20','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(16,0,1,4,2,2,2,4,'2018-02-13 00:00:00','443447',3000.0000,'',25,'ASM(22) d 12999 3000.00,c 12001 3000.00',2,'payment for ASM-22','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(17,0,1,4,2,1,1,4,'2018-02-13 00:00:00','641777',857.1400,'',25,'ASM(24) d 12999 857.14,c 12001 857.14',2,'payment for ASM-24','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(18,0,1,4,2,1,1,4,'2018-02-13 00:00:00','79311',50.0000,'',25,'ASM(25) d 12999 50.00,c 12001 50.00',2,'payment for ASM-25','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(19,0,1,4,2,1,1,4,'2018-02-13 00:00:00','614879',5.7100,'',25,'ASM(26) d 12999 5.71,c 12001 5.71',2,'payment for ASM-26','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(20,0,1,4,2,1,1,4,'2018-02-13 00:00:00','485412',10.0000,'',25,'ASM(27) d 12999 10.00,c 12001 10.00',2,'payment for ASM-27','','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ReceiptAllocation`
--

LOCK TABLES `ReceiptAllocation` WRITE;
/*!40000 ALTER TABLE `ReceiptAllocation` DISABLE KEYS */;
INSERT INTO `ReceiptAllocation` VALUES (1,1,1,1,'2018-02-13 00:00:00',2000.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(2,2,1,1,'2018-02-13 00:00:00',571.4300,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(3,3,1,1,'2018-02-13 00:00:00',50.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(4,4,1,1,'2018-02-13 00:00:00',5.7100,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(5,5,1,1,'2018-02-13 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(6,6,1,1,'2018-02-13 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(7,7,1,2,'2018-02-13 00:00:00',2000.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(8,8,1,2,'2018-02-13 00:00:00',571.4300,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(9,9,1,2,'2018-02-13 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(10,10,1,3,'2018-02-13 00:00:00',2000.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(11,11,1,3,'2018-02-13 00:00:00',571.4300,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(12,12,1,3,'2018-02-13 00:00:00',50.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(13,13,1,3,'2018-02-13 00:00:00',5.7100,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(14,14,1,3,'2018-02-13 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(15,15,1,3,'2018-02-13 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(16,16,1,4,'2018-02-13 00:00:00',3000.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(17,17,1,4,'2018-02-13 00:00:00',857.1400,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(18,18,1,4,'2018-02-13 00:00:00',50.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(19,19,1,4,'2018-02-13 00:00:00',5.7100,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(20,20,1,4,'2018-02-13 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(21,1,1,1,'2018-02-13 00:00:00',2000.0000,2,0,'ASM(2) d 12999 2000.00,c 12001 2000.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(22,2,1,1,'2018-02-13 00:00:00',571.4300,4,0,'ASM(4) d 12999 571.43,c 12001 571.43','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(23,3,1,1,'2018-02-13 00:00:00',50.0000,5,0,'ASM(5) d 12999 50.00,c 12001 50.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(24,4,1,1,'2018-02-13 00:00:00',5.7100,6,0,'ASM(6) d 12999 5.71,c 12001 5.71','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(25,5,1,1,'2018-02-13 00:00:00',10.0000,7,0,'ASM(7) d 12999 10.00,c 12001 10.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(26,6,1,1,'2018-02-13 00:00:00',10.0000,8,0,'ASM(8) d 12999 10.00,c 12001 10.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(27,7,1,2,'2018-02-13 00:00:00',2000.0000,10,0,'ASM(10) d 12999 2000.00,c 12001 2000.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(28,8,1,2,'2018-02-13 00:00:00',571.4300,11,0,'ASM(11) d 12999 571.43,c 12001 571.43','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(29,9,1,2,'2018-02-13 00:00:00',10.0000,12,0,'ASM(12) d 12999 10.00,c 12001 10.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(30,10,1,3,'2018-02-13 00:00:00',2000.0000,14,0,'ASM(14) d 12999 2000.00,c 12001 2000.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(31,11,1,3,'2018-02-13 00:00:00',571.4300,16,0,'ASM(16) d 12999 571.43,c 12001 571.43','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(32,12,1,3,'2018-02-13 00:00:00',50.0000,17,0,'ASM(17) d 12999 50.00,c 12001 50.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(33,13,1,3,'2018-02-13 00:00:00',5.7100,18,0,'ASM(18) d 12999 5.71,c 12001 5.71','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(34,14,1,3,'2018-02-13 00:00:00',10.0000,19,0,'ASM(19) d 12999 10.00,c 12001 10.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(35,15,1,3,'2018-02-13 00:00:00',10.0000,20,0,'ASM(20) d 12999 10.00,c 12001 10.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(36,16,1,4,'2018-02-13 00:00:00',3000.0000,22,0,'ASM(22) d 12999 3000.00,c 12001 3000.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(37,17,1,4,'2018-02-13 00:00:00',857.1400,24,0,'ASM(24) d 12999 857.14,c 12001 857.14','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(38,18,1,4,'2018-02-13 00:00:00',50.0000,25,0,'ASM(25) d 12999 50.00,c 12001 50.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(39,19,1,4,'2018-02-13 00:00:00',5.7100,26,0,'ASM(26) d 12999 5.71,c 12001 5.71','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(40,20,1,4,'2018-02-13 00:00:00',10.0000,27,0,'ASM(27) d 12999 10.00,c 12001 10.00','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(41,2,1,1,'2018-02-13 00:00:00',571.4300,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(42,3,1,1,'2018-02-13 00:00:00',50.0000,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(43,4,1,1,'2018-02-13 00:00:00',5.7100,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(44,5,1,1,'2018-02-13 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(45,6,1,1,'2018-02-13 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(46,8,1,2,'2018-02-13 00:00:00',571.4300,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(47,9,1,2,'2018-02-13 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(48,11,1,3,'2018-02-13 00:00:00',571.4300,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(49,12,1,3,'2018-02-13 00:00:00',50.0000,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(50,13,1,3,'2018-02-13 00:00:00',5.7100,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(51,14,1,3,'2018-02-13 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(52,15,1,3,'2018-02-13 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(53,17,1,4,'2018-02-13 00:00:00',857.1400,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(54,18,1,4,'2018-02-13 00:00:00',50.0000,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(55,19,1,4,'2018-02-13 00:00:00',5.7100,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(56,20,1,4,'2018-02-13 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(57,1,1,1,'2018-02-13 00:00:00',2000.0000,0,0,'d 10104 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(58,7,1,2,'2018-02-13 00:00:00',2000.0000,0,0,'d 10104 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(59,10,1,3,'2018-02-13 00:00:00',2000.0000,0,0,'d 10104 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(60,16,1,4,'2018-02-13 00:00:00',3000.0000,0,0,'d 10104 _, c 10999 _','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Rentable`
--

LOCK TABLES `Rentable` WRITE;
/*!40000 ALTER TABLE `Rentable` DISABLE KEYS */;
INSERT INTO `Rentable` VALUES (1,1,0,'Rentable001',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(2,1,0,'Rentable002',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(3,1,0,'Rentable003',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(4,1,0,'Rentable004',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(5,1,0,'Rentable005',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(6,1,0,'Rentable006',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(7,1,0,'Rentable007',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(8,1,0,'Rentable008',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(9,1,0,'Rentable009',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(10,1,0,'Rentable010',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(11,1,0,'Rentable011',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(12,1,0,'Rentable012',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(13,1,0,'CP001',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(14,1,0,'CP002',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(15,1,0,'CP003',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(16,1,0,'CP004',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(17,1,0,'CP005',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(18,1,0,'CP006',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(19,1,0,'CP007',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(20,1,0,'CP008',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(21,1,0,'CP009',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(22,1,0,'CP010',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(23,1,0,'CP011',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,''),(24,1,0,'CP012',0,0,'0000-00-00 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0,'');
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
  `FirstName` varchar(50) NOT NULL DEFAULT '',
  `LastName` varchar(50) NOT NULL DEFAULT '',
  `Email` varchar(100) NOT NULL DEFAULT '',
  `Phone` varchar(100) NOT NULL DEFAULT '',
  `Address` varchar(100) NOT NULL DEFAULT '',
  `Address2` varchar(100) NOT NULL DEFAULT '',
  `City` varchar(100) NOT NULL DEFAULT '',
  `State` char(25) NOT NULL DEFAULT '',
  `PostalCode` varchar(100) NOT NULL DEFAULT '',
  `Country` varchar(100) NOT NULL DEFAULT '',
  `CCName` varchar(100) NOT NULL DEFAULT '',
  `CCType` varchar(100) NOT NULL DEFAULT '',
  `CCNumber` varchar(100) NOT NULL DEFAULT '',
  `CCExpMonth` varchar(100) NOT NULL DEFAULT '',
  `CCExpYear` varchar(100) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RLID`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableLeaseStatus`
--

LOCK TABLES `RentableLeaseStatus` WRITE;
/*!40000 ALTER TABLE `RentableLeaseStatus` DISABLE KEYS */;
INSERT INTO `RentableLeaseStatus` VALUES (1,1,1,0,'2017-01-01 00:00:00','2018-02-13 00:00:00','','','','','','','','','','','','','','','','','2019-02-12 17:13:53',0,'2019-02-09 06:26:31',0),(2,1,1,1,'2018-02-13 00:00:00','2020-03-01 00:00:00','','','','','','','','','','','','','','','','','2019-03-01 18:41:26',0,'2019-02-09 06:26:31',0),(3,1,1,2,'2020-03-01 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','2019-03-01 18:41:26',0,'2019-02-09 06:26:31',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableMarketRate`
--

LOCK TABLES `RentableMarketRate` WRITE;
/*!40000 ALTER TABLE `RentableMarketRate` DISABLE KEYS */;
INSERT INTO `RentableMarketRate` VALUES (1,1,0,1000.0000,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,2,0,1500.0000,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,3,0,1750.0000,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,4,0,2500.0000,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(5,5,0,35.0000,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableTypeRef`
--

LOCK TABLES `RentableTypeRef` WRITE;
/*!40000 ALTER TABLE `RentableTypeRef` DISABLE KEYS */;
INSERT INTO `RentableTypeRef` VALUES (1,1,1,1,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,2,1,1,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,3,1,1,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,4,1,2,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(5,5,1,2,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(6,6,1,2,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(7,7,1,3,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(8,8,1,3,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(9,9,1,3,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(10,10,1,4,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(11,11,1,4,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(12,12,1,4,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(13,13,1,5,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(14,14,1,5,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(15,15,1,5,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(16,16,1,5,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(17,17,1,5,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(18,18,1,5,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(19,19,1,5,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(20,20,1,5,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(21,21,1,5,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(22,22,1,5,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(23,23,1,5,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(24,24,1,5,0,0,'2018-01-01 00:00:00','3001-01-01 00:00:00','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableTypes`
--

LOCK TABLES `RentableTypes` WRITE;
/*!40000 ALTER TABLE `RentableTypes` DISABLE KEYS */;
INSERT INTO `RentableTypes` VALUES (1,1,'1BR','1 Bed 1 Bath',6,4,4,4,45,'2018-10-16 19:40:53',0,'2018-10-16 19:35:31',0),(2,1,'2BR','2 Bed 2 Bath',6,4,4,4,46,'2018-10-16 19:40:53',0,'2018-10-16 19:35:31',0),(3,1,'3BR','3 Bed 2.5 Bath',6,4,4,4,47,'2018-10-16 19:40:53',0,'2018-10-16 19:35:31',0),(4,1,'4BR','4 BR 3 Bath',6,4,4,4,48,'2018-10-16 19:40:53',0,'2018-10-16 19:35:31',0),(5,1,'CP','Car Port',6,4,4,6,49,'2018-10-16 19:40:53',0,'2018-10-16 19:35:31',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableUseStatus`
--

LOCK TABLES `RentableUseStatus` WRITE;
/*!40000 ALTER TABLE `RentableUseStatus` DISABLE KEYS */;
INSERT INTO `RentableUseStatus` VALUES (1,1,1,1,'2018-01-01 00:00:00','3001-01-01 00:00:00','','2018-09-19 18:14:38',0,'2018-09-19 18:14:38',0),(2,2,1,1,'2018-01-01 00:00:00','3001-01-01 00:00:00','Test','2018-09-19 18:14:38',0,'2018-09-19 18:14:38',0),(3,3,1,1,'2018-01-01 00:00:00','3001-01-01 00:00:00','Test','2018-09-19 18:14:38',0,'2018-09-19 18:14:38',0);
/*!40000 ALTER TABLE `RentableUseStatus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RentableUseType`
--

DROP TABLE IF EXISTS `RentableUseType`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RentableUseType` (
  `UTID` bigint(20) NOT NULL AUTO_INCREMENT,
  `RID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `UseType` smallint(6) NOT NULL DEFAULT '0',
  `DtStart` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtStop` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `Comment` varchar(2048) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`UTID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableUseType`
--

LOCK TABLES `RentableUseType` WRITE;
/*!40000 ALTER TABLE `RentableUseType` DISABLE KEYS */;
/*!40000 ALTER TABLE `RentableUseType` ENABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableUsers`
--

LOCK TABLES `RentableUsers` WRITE;
/*!40000 ALTER TABLE `RentableUsers` DISABLE KEYS */;
INSERT INTO `RentableUsers` VALUES (1,1,1,1,'2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,2,1,2,'2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,3,1,3,'2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,4,1,4,'2018-02-13','2020-03-01','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(5,5,1,41,'2018-10-20','2019-11-01','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(6,5,1,38,'2018-10-20','2019-11-01','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211);
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreement`
--

LOCK TABLES `RentalAgreement` WRITE;
/*!40000 ALTER TABLE `RentalAgreement` DISABLE KEYS */;
INSERT INTO `RentalAgreement` VALUES (1,0,0,1,1,0,'2018-02-13 00:00:00','2018-02-13','2020-03-01','2018-02-13','2020-03-01','2018-02-13','2020-03-01','2018-03-01',2,1,2,'',0,0,0.0000,'','0000-00-00','0000-00-00',0.0000,0.0000,283,'0000-00-00','','','','0000-00-00','','0000-00-00','','0000-00-00',0,52,250,'2018-02-13 00:00:00',213,'2018-02-13 00:00:00',0,114,'2018-02-13 00:00:00',0,250,'2018-02-13 00:00:00',56,'2018-02-13 00:00:00',0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,0,0,1,1,0,'2018-02-13 00:00:00','2018-02-13','2020-03-01','2018-02-13','2020-03-01','2018-02-13','2020-03-01','2018-03-01',1,2,2,'',0,0,0.0000,'','0000-00-00','0000-00-00',0.0000,0.0000,242,'0000-00-00','','','','0000-00-00','','0000-00-00','','0000-00-00',0,52,70,'2018-02-13 00:00:00',225,'2018-02-13 00:00:00',0,252,'2018-02-13 00:00:00',0,61,'2018-02-13 00:00:00',142,'2018-02-13 00:00:00',0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,0,0,1,1,0,'2018-02-13 00:00:00','2018-02-13','2020-03-01','2018-02-13','2020-03-01','2018-02-13','2020-03-01','2018-03-01',2,2,2,'',0,0,0.0000,'','0000-00-00','0000-00-00',0.0000,0.0000,101,'0000-00-00','','','','0000-00-00','','0000-00-00','','0000-00-00',0,52,169,'2018-02-13 00:00:00',220,'2018-02-13 00:00:00',0,32,'2018-02-13 00:00:00',0,39,'2018-02-13 00:00:00',19,'2018-02-13 00:00:00',0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,0,0,1,1,0,'2018-02-13 00:00:00','2018-02-13','2020-03-01','2018-02-13','2020-03-01','2018-02-13','2020-03-01','2018-03-01',1,2,2,'',0,0,0.0000,'','0000-00-00','0000-00-00',0.0000,0.0000,141,'0000-00-00','','','','0000-00-00','','0000-00-00','','0000-00-00',0,52,183,'2018-02-13 00:00:00',103,'2018-02-13 00:00:00',0,37,'2018-02-13 00:00:00',0,275,'2018-02-13 00:00:00',116,'2018-02-13 00:00:00',0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(5,0,0,0,1,0,'2018-10-20 00:00:00','2018-10-20','2019-11-01','2018-10-20','2019-11-01','2018-10-20','2019-11-01','0000-00-00',0,0,0,'',0,0,0.0000,'','0000-00-00','0000-00-00',0.0000,0.0000,211,'0000-00-00','','','','0000-00-00','','0000-00-00','','0000-00-00',0,52,211,'2018-10-17 06:21:00',211,'2018-10-17 06:22:00',0,211,'2018-10-17 06:22:00',0,211,'2018-10-17 06:22:00',211,'2018-10-17 06:23:00',0,0,'1900-01-01 00:00:00','1900-01-01 00:00:00',0,'1900-01-01 00:00:00',0,'2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211);
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
INSERT INTO `RentalAgreementPayors` VALUES (1,1,1,1,'2018-02-13','2020-03-01',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,2,1,2,'2018-02-13','2020-03-01',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,3,1,3,'2018-02-13','2020-03-01',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,4,1,4,'2018-02-13','2020-03-01',0,'2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(5,5,1,41,'2018-10-20','2019-11-01',0,'2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211),(6,5,1,38,'2018-10-20','2019-11-01',0,'2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211);
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementRentables`
--

LOCK TABLES `RentalAgreementRentables` WRITE;
/*!40000 ALTER TABLE `RentalAgreementRentables` DISABLE KEYS */;
INSERT INTO `RentalAgreementRentables` VALUES (1,1,1,1,0,0,1000.0000,'2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,2,1,2,0,0,1000.0000,'2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,3,1,3,0,0,1000.0000,'2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,4,1,4,0,0,1500.0000,'2018-02-13','2020-03-01','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(5,5,1,5,0,0,0.0000,'2018-10-20','2019-10-31','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211);
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
) ENGINE=InnoDB AUTO_INCREMENT=73 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TBind`
--

LOCK TABLES `TBind` WRITE;
/*!40000 ALTER TABLE `TBind` DISABLE KEYS */;
INSERT INTO `TBind` VALUES (1,1,1,1,15,1,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,1,1,1,15,2,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,1,1,1,14,1,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,1,1,2,15,3,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(5,1,1,3,15,4,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(6,1,1,3,15,5,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(7,1,1,3,14,2,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(8,1,1,4,15,6,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(9,1,1,4,14,3,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(10,1,1,5,15,7,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(11,1,1,5,15,8,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(12,1,1,5,14,4,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(13,1,1,6,15,9,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(14,1,1,6,14,5,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(15,1,1,7,15,10,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(16,1,1,7,14,6,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(17,1,1,7,14,7,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(18,1,1,8,15,11,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(19,1,1,9,15,12,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(20,1,1,9,14,8,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(21,1,1,10,15,13,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(22,1,1,11,15,14,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(23,1,1,12,15,15,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(24,1,1,13,15,16,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(25,1,1,13,14,9,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(26,1,1,13,14,10,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(27,1,1,14,15,17,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(28,1,1,15,15,18,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(29,1,1,16,15,19,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(30,1,1,16,15,20,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(31,1,1,16,14,11,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(32,1,1,16,14,12,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(33,1,1,17,15,21,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(34,1,1,17,14,13,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(35,1,1,18,15,22,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(36,1,1,18,14,14,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(37,1,1,19,15,23,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(38,1,1,19,14,15,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(39,1,1,20,14,16,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(40,1,1,21,15,24,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(41,1,1,23,15,25,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(42,1,1,23,14,17,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(43,1,1,23,14,18,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(44,1,1,24,15,26,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(45,1,1,25,15,27,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(46,1,1,25,14,19,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(47,1,1,26,15,28,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(48,1,1,26,14,20,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(49,1,1,27,15,29,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(50,1,1,27,14,21,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(51,1,1,28,15,30,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(52,1,1,28,14,22,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(53,1,1,29,14,23,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(54,1,1,30,15,31,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(55,1,1,31,15,32,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(56,1,1,31,14,24,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(57,1,1,32,15,33,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(58,1,1,32,14,25,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(59,1,1,33,15,34,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(60,1,1,34,15,35,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(61,1,1,35,15,36,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(62,1,1,35,14,26,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(63,1,1,36,15,37,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(64,1,1,36,14,27,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(65,1,1,37,15,38,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(66,1,1,38,15,39,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(67,1,1,38,14,28,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(68,1,1,39,15,40,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(69,1,1,39,14,29,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(70,1,1,40,15,41,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(71,1,1,40,14,30,'2018-02-13 00:00:00','9999-12-31 00:00:00',0,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(72,1,1,41,14,31,'2018-10-20 00:00:00','9999-12-31 00:00:00',0,'2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211);
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
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TWS`
--

LOCK TABLES `TWS` WRITE;
/*!40000 ALTER TABLE `TWS` DISABLE KEYS */;
INSERT INTO `TWS` VALUES (1,'CreateAssessmentInstances','','CreateAssessmentInstances','2018-02-25 00:00:00','Steves-MacBook-Pro-2.local',4,'2018-02-24 01:19:54','2018-02-24 01:19:54','2017-11-10 15:24:21','2018-02-23 17:19:53'),(2,'CleanRARBalanceCache','','CleanRARBalanceCache','2018-02-24 05:09:45','Steves-MacBook-Pro-2.local',4,'2018-02-24 05:04:45','2018-02-24 05:04:45','2018-02-23 17:19:43','2018-02-23 21:04:45'),(3,'CleanSecDepBalanceCache','','CleanSecDepBalanceCache','2018-02-24 05:09:45','Steves-MacBook-Pro-2.local',4,'2018-02-24 05:04:45','2018-02-24 05:04:45','2018-02-23 17:19:43','2018-02-23 21:04:45'),(4,'CleanAcctSliceCache','','CleanAcctSliceCache','2018-02-24 05:09:45','Steves-MacBook-Pro-2.local',4,'2018-02-24 05:04:45','2018-02-24 05:04:45','2018-02-23 17:19:43','2018-02-23 21:04:45'),(5,'CleanARSliceCache','','CleanARSliceCache','2018-02-24 05:09:45','Steves-MacBook-Pro-2.local',4,'2018-02-24 05:04:45','2018-02-24 05:04:45','2018-02-23 17:19:43','2018-02-23 21:04:45'),(6,'RARBcacheBot','','RARBcacheBot','2019-03-01 18:46:00','Steves-MacBook-Pro-2.local',4,'2019-03-01 18:41:00','2019-03-01 18:41:00','2018-06-02 13:09:58','2019-03-01 10:41:00'),(7,'ARSliceCacheBot','','ARSliceCacheBot','2019-03-01 18:46:00','Steves-MacBook-Pro-2.local',4,'2019-03-01 18:41:00','2019-03-01 18:41:00','2018-06-02 13:09:58','2019-03-01 10:41:00'),(8,'TLReportBot','','TLReportBot','2019-03-01 18:43:00','Steves-MacBook-Pro-2.local',4,'2019-03-01 18:41:00','2019-03-01 18:41:00','2018-06-02 13:09:58','2019-03-01 10:41:00'),(9,'ManualTaskBot','','ManualTaskBot','2019-03-02 18:41:00','Steves-MacBook-Pro-2.local',4,'2019-03-01 18:41:00','2019-03-01 18:41:00','2018-06-02 13:09:58','2019-03-01 10:41:00'),(10,'AssessmentBot','','AssessmentBot','2019-03-02 18:41:00','Steves-MacBook-Pro-2.local',4,'2019-03-01 18:41:00','2019-03-01 18:41:00','2018-06-02 13:09:58','2019-03-01 10:41:00'),(11,'SecDepCacheBot','','SecDepCacheBot','2019-03-01 18:46:00','Steves-MacBook-Pro-2.local',4,'2019-03-01 18:41:00','2019-03-01 18:41:00','2018-06-02 13:09:58','2019-03-01 10:41:00'),(12,'AcctSliceCacheBot','','AcctSliceCacheBot','2019-03-01 18:46:00','Steves-MacBook-Pro-2.local',4,'2019-03-01 18:41:00','2019-03-01 18:41:00','2018-06-02 13:09:58','2019-03-01 10:41:00'),(13,'TLInstanceBot','','TLInstanceBot','2019-03-02 18:41:00','Steves-MacBook-Pro-2.local',4,'2019-03-01 18:41:00','2019-03-01 18:41:00','2018-06-02 13:09:58','2019-03-01 10:41:00');
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
) ENGINE=InnoDB AUTO_INCREMENT=111 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Task`
--

LOCK TABLES `Task` WRITE;
/*!40000 ALTER TABLE `Task` DISABLE KEYS */;
INSERT INTO `Task` VALUES (1,1,1,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','2018-02-28 20:00:00','2018-02-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(2,1,1,'Review all receivables for accuracy','ManualTaskBot','2018-02-28 20:00:00','2018-02-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(3,1,1,'Compare total cash deposits to bank statement','ManualTaskBot','2018-02-28 20:00:00','2018-02-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(4,1,1,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','2018-02-28 07:00:00','2018-02-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(5,1,1,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','2018-02-28 07:00:00','2018-02-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(6,1,1,'Make certain that all suspense accounts have been closed out','ManualTaskBot','2018-02-28 07:00:00','2018-02-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(7,1,1,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','2018-02-28 07:00:00','2018-02-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(8,1,1,'Print Rent Roll Activity Report','ManualTaskBot','2018-02-28 07:00:00','2018-02-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(9,1,1,'Print Rent Roll Report','ManualTaskBot','2018-02-28 07:00:00','2018-02-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(10,1,1,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','2018-02-28 07:00:00','2018-02-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(11,1,2,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','2018-03-31 20:00:00','2018-03-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(12,1,2,'Review all receivables for accuracy','ManualTaskBot','2018-03-31 20:00:00','2018-03-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(13,1,2,'Compare total cash deposits to bank statement','ManualTaskBot','2018-03-31 20:00:00','2018-03-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(14,1,2,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','2018-03-31 07:00:00','2018-03-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(15,1,2,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','2018-03-31 07:00:00','2018-03-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(16,1,2,'Make certain that all suspense accounts have been closed out','ManualTaskBot','2018-03-31 07:00:00','2018-03-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(17,1,2,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','2018-03-31 07:00:00','2018-03-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(18,1,2,'Print Rent Roll Activity Report','ManualTaskBot','2018-03-31 07:00:00','2018-03-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(19,1,2,'Print Rent Roll Report','ManualTaskBot','2018-03-31 07:00:00','2018-03-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(20,1,2,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','2018-03-31 07:00:00','2018-03-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(21,1,3,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','2018-04-30 20:00:00','2018-04-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(22,1,3,'Review all receivables for accuracy','ManualTaskBot','2018-04-30 20:00:00','2018-04-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(23,1,3,'Compare total cash deposits to bank statement','ManualTaskBot','2018-04-30 20:00:00','2018-04-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(24,1,3,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','2018-04-30 07:00:00','2018-04-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(25,1,3,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','2018-04-30 07:00:00','2018-04-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(26,1,3,'Make certain that all suspense accounts have been closed out','ManualTaskBot','2018-04-30 07:00:00','2018-04-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(27,1,3,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','2018-04-30 07:00:00','2018-04-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(28,1,3,'Print Rent Roll Activity Report','ManualTaskBot','2018-04-30 07:00:00','2018-04-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(29,1,3,'Print Rent Roll Report','ManualTaskBot','2018-04-30 07:00:00','2018-04-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(30,1,3,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','2018-04-30 07:00:00','2018-04-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(31,1,4,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','2018-05-31 20:00:00','2018-05-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(32,1,4,'Review all receivables for accuracy','ManualTaskBot','2018-05-31 20:00:00','2018-05-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(33,1,4,'Compare total cash deposits to bank statement','ManualTaskBot','2018-05-31 20:00:00','2018-05-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(34,1,4,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','2018-05-31 07:00:00','2018-05-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(35,1,4,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','2018-05-31 07:00:00','2018-05-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(36,1,4,'Make certain that all suspense accounts have been closed out','ManualTaskBot','2018-05-31 07:00:00','2018-05-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(37,1,4,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','2018-05-31 07:00:00','2018-05-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(38,1,4,'Print Rent Roll Activity Report','ManualTaskBot','2018-05-31 07:00:00','2018-05-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(39,1,4,'Print Rent Roll Report','ManualTaskBot','2018-05-31 07:00:00','2018-05-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(40,1,4,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','2018-05-31 07:00:00','2018-05-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(41,1,5,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','2018-06-30 20:00:00','2018-06-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(42,1,5,'Review all receivables for accuracy','ManualTaskBot','2018-06-30 20:00:00','2018-06-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(43,1,5,'Compare total cash deposits to bank statement','ManualTaskBot','2018-06-30 20:00:00','2018-06-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(44,1,5,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','2018-06-30 07:00:00','2018-06-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(45,1,5,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','2018-06-30 07:00:00','2018-06-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(46,1,5,'Make certain that all suspense accounts have been closed out','ManualTaskBot','2018-06-30 07:00:00','2018-06-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(47,1,5,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','2018-06-30 07:00:00','2018-06-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(48,1,5,'Print Rent Roll Activity Report','ManualTaskBot','2018-06-30 07:00:00','2018-06-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(49,1,5,'Print Rent Roll Report','ManualTaskBot','2018-06-30 07:00:00','2018-06-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(50,1,5,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','2018-06-30 07:00:00','2018-06-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(51,1,6,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','2018-07-31 20:00:00','2018-07-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(52,1,6,'Review all receivables for accuracy','ManualTaskBot','2018-07-31 20:00:00','2018-07-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(53,1,6,'Compare total cash deposits to bank statement','ManualTaskBot','2018-07-31 20:00:00','2018-07-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(54,1,6,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','2018-07-31 07:00:00','2018-07-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(55,1,6,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','2018-07-31 07:00:00','2018-07-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(56,1,6,'Make certain that all suspense accounts have been closed out','ManualTaskBot','2018-07-31 07:00:00','2018-07-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(57,1,6,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','2018-07-31 07:00:00','2018-07-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(58,1,6,'Print Rent Roll Activity Report','ManualTaskBot','2018-07-31 07:00:00','2018-07-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(59,1,6,'Print Rent Roll Report','ManualTaskBot','2018-07-31 07:00:00','2018-07-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(60,1,6,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','2018-07-31 07:00:00','2018-07-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(61,1,7,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','2018-08-31 20:00:00','2018-08-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(62,1,7,'Review all receivables for accuracy','ManualTaskBot','2018-08-31 20:00:00','2018-08-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(63,1,7,'Compare total cash deposits to bank statement','ManualTaskBot','2018-08-31 20:00:00','2018-08-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(64,1,7,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','2018-08-31 07:00:00','2018-08-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(65,1,7,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','2018-08-31 07:00:00','2018-08-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(66,1,7,'Make certain that all suspense accounts have been closed out','ManualTaskBot','2018-08-31 07:00:00','2018-08-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(67,1,7,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','2018-08-31 07:00:00','2018-08-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(68,1,7,'Print Rent Roll Activity Report','ManualTaskBot','2018-08-31 07:00:00','2018-08-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(69,1,7,'Print Rent Roll Report','ManualTaskBot','2018-08-31 07:00:00','2018-08-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(70,1,7,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','2018-08-31 07:00:00','2018-08-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(71,1,8,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','2018-09-30 20:00:00','2018-09-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(72,1,8,'Review all receivables for accuracy','ManualTaskBot','2018-09-30 20:00:00','2018-09-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(73,1,8,'Compare total cash deposits to bank statement','ManualTaskBot','2018-09-30 20:00:00','2018-09-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(74,1,8,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','2018-09-30 07:00:00','2018-09-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(75,1,8,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','2018-09-30 07:00:00','2018-09-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(76,1,8,'Make certain that all suspense accounts have been closed out','ManualTaskBot','2018-09-30 07:00:00','2018-09-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(77,1,8,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','2018-09-30 07:00:00','2018-09-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(78,1,8,'Print Rent Roll Activity Report','ManualTaskBot','2018-09-30 07:00:00','2018-09-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(79,1,8,'Print Rent Roll Report','ManualTaskBot','2018-09-30 07:00:00','2018-09-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(80,1,8,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','2018-09-30 07:00:00','2018-09-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(81,1,9,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','2018-10-31 20:00:00','2018-10-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(82,1,9,'Review all receivables for accuracy','ManualTaskBot','2018-10-31 20:00:00','2018-10-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(83,1,9,'Compare total cash deposits to bank statement','ManualTaskBot','2018-10-31 20:00:00','2018-10-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(84,1,9,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','2018-10-31 07:00:00','2018-10-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(85,1,9,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','2018-10-31 07:00:00','2018-10-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(86,1,9,'Make certain that all suspense accounts have been closed out','ManualTaskBot','2018-10-31 07:00:00','2018-10-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(87,1,9,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','2018-10-31 07:00:00','2018-10-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(88,1,9,'Print Rent Roll Activity Report','ManualTaskBot','2018-10-31 07:00:00','2018-10-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(89,1,9,'Print Rent Roll Report','ManualTaskBot','2018-10-31 07:00:00','2018-10-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(90,1,9,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','2018-10-31 07:00:00','2018-10-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(91,1,10,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','0001-01-31 20:00:00','0001-01-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-12 17:11:24',-8,'2019-02-12 17:11:24',-8),(92,1,10,'Review all receivables for accuracy','ManualTaskBot','0001-01-31 20:00:00','0001-01-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-12 17:11:24',-8,'2019-02-12 17:11:24',-8),(93,1,10,'Compare total cash deposits to bank statement','ManualTaskBot','0001-01-31 20:00:00','0001-01-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-12 17:11:24',-8,'2019-02-12 17:11:24',-8),(94,1,10,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','0001-01-31 07:00:00','0001-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-12 17:11:24',-8,'2019-02-12 17:11:24',-8),(95,1,10,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','0001-01-31 07:00:00','0001-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-12 17:11:24',-8,'2019-02-12 17:11:24',-8),(96,1,10,'Make certain that all suspense accounts have been closed out','ManualTaskBot','0001-01-31 07:00:00','0001-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-12 17:11:24',-8,'2019-02-12 17:11:24',-8),(97,1,10,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','0001-01-31 07:00:00','0001-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-12 17:11:24',-8,'2019-02-12 17:11:24',-8),(98,1,10,'Print Rent Roll Activity Report','ManualTaskBot','0001-01-31 07:00:00','0001-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-12 17:11:24',-8,'2019-02-12 17:11:24',-8),(99,1,10,'Print Rent Roll Report','ManualTaskBot','0001-01-31 07:00:00','0001-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-12 17:11:24',-8,'2019-02-12 17:11:24',-8),(100,1,10,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','0001-01-31 07:00:00','0001-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-12 17:11:24',-8,'2019-02-12 17:11:24',-8),(101,1,11,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','2019-03-31 20:00:00','2019-03-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-03-01 18:41:00',-8,'2019-03-01 18:41:00',-8),(102,1,11,'Review all receivables for accuracy','ManualTaskBot','2019-03-31 20:00:00','2019-03-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-03-01 18:41:00',-8,'2019-03-01 18:41:00',-8),(103,1,11,'Compare total cash deposits to bank statement','ManualTaskBot','2019-03-31 20:00:00','2019-03-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-03-01 18:41:00',-8,'2019-03-01 18:41:00',-8),(104,1,11,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','2019-03-31 07:00:00','2019-03-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-03-01 18:41:00',-8,'2019-03-01 18:41:00',-8),(105,1,11,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','2019-03-31 07:00:00','2019-03-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-03-01 18:41:00',-8,'2019-03-01 18:41:00',-8),(106,1,11,'Make certain that all suspense accounts have been closed out','ManualTaskBot','2019-03-31 07:00:00','2019-03-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-03-01 18:41:00',-8,'2019-03-01 18:41:00',-8),(107,1,11,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','2019-03-31 07:00:00','2019-03-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-03-01 18:41:00',-8,'2019-03-01 18:41:00',-8),(108,1,11,'Print Rent Roll Activity Report','ManualTaskBot','2019-03-31 07:00:00','2019-03-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-03-01 18:41:00',-8,'2019-03-01 18:41:00',-8),(109,1,11,'Print Rent Roll Report','ManualTaskBot','2019-03-31 07:00:00','2019-03-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-03-01 18:41:00',-8,'2019-03-01 18:41:00',-8),(110,1,11,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','2019-03-31 07:00:00','2019-03-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-03-01 18:41:00',-8,'2019-03-01 18:41:00',-8);
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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TaskDescriptor`
--

LOCK TABLES `TaskDescriptor` WRITE;
/*!40000 ALTER TABLE `TaskDescriptor` DISABLE KEYS */;
INSERT INTO `TaskDescriptor` VALUES (1,1,1,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','2018-01-31 20:00:00','2018-01-20 20:00:00',6,'','2018-10-01 20:22:40',211,'2018-03-14 19:50:32',0),(2,1,1,'Review all receivables for accuracy','ManualTaskBot','2018-01-31 20:00:00','2018-01-20 20:00:00',6,'(provide comment for any receivables more than 30 days old','2018-10-01 20:22:44',211,'2018-03-14 19:50:32',0),(3,1,1,'Compare total cash deposits to bank statement','ManualTaskBot','2018-01-31 20:00:00','2018-01-20 20:00:00',6,'','2018-10-01 20:22:47',211,'2018-03-14 19:50:32',0),(5,1,1,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','2018-05-31 07:00:00','2018-05-20 07:00:00',6,' or make certain that you have a Report for any After-Lease Concessions occurring during the month','2018-10-01 20:22:50',211,'2018-05-29 18:24:26',211),(6,1,1,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','2018-05-31 07:00:00','2018-05-20 07:00:00',6,'','2018-10-01 20:23:02',211,'2018-05-29 18:25:05',211),(7,1,1,'Make certain that all suspense accounts have been closed out','ManualTaskBot','2018-05-31 07:00:00','2018-05-20 07:00:00',6,'','2018-10-01 20:23:06',211,'2018-05-29 18:25:30',211),(8,1,1,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','2018-05-31 07:00:00','2018-05-20 07:00:00',6,'','2018-10-01 20:23:09',211,'2018-05-29 18:25:57',211),(9,1,1,'Print Rent Roll Activity Report','ManualTaskBot','2018-05-31 07:00:00','2018-05-20 07:00:00',6,'','2018-10-01 20:23:12',211,'2018-05-29 18:30:54',211),(10,1,1,'Print Rent Roll Report','ManualTaskBot','2018-05-31 07:00:00','2018-05-20 07:00:00',6,'','2018-10-01 20:23:15',211,'2018-05-29 18:31:18',211),(11,1,1,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','2018-05-31 07:00:00','2018-05-20 07:00:00',6,'','2018-10-01 20:23:18',211,'2018-05-29 18:32:06',211);
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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TaskList`
--

LOCK TABLES `TaskList` WRITE;
/*!40000 ALTER TABLE `TaskList` DISABLE KEYS */;
INSERT INTO `TaskList` VALUES (1,1,0,1,'Monthly Close',6,'2018-02-28 17:00:00','2018-02-20 17:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','0000-00-00 00:00:00',86400000000000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(2,1,1,1,'Monthly Close',6,'2018-03-31 17:00:00','2018-03-20 17:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','0000-00-00 00:00:00',86400000000000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(3,1,1,1,'Monthly Close',6,'2018-04-30 17:00:00','2018-04-20 17:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','0000-00-00 00:00:00',86400000000000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(4,1,1,1,'Monthly Close',6,'2018-05-31 17:00:00','2018-05-20 17:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','0000-00-00 00:00:00',86400000000000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(5,1,1,1,'Monthly Close',6,'2018-06-30 17:00:00','2018-06-20 17:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','0000-00-00 00:00:00',86400000000000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(6,1,1,1,'Monthly Close',6,'2018-07-31 17:00:00','2018-07-20 17:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','0000-00-00 00:00:00',86400000000000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(7,1,1,1,'Monthly Close',6,'2018-08-31 17:00:00','2018-08-20 17:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','0000-00-00 00:00:00',86400000000000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(8,1,1,1,'Monthly Close',6,'2018-09-30 17:00:00','2018-09-20 17:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','0000-00-00 00:00:00',86400000000000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(9,1,1,1,'Monthly Close',6,'2018-10-31 17:00:00','2018-10-20 17:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','0000-00-00 00:00:00',86400000000000,'','2018-10-16 19:35:32',0,'2018-10-16 19:35:32',0),(10,1,1,1,'Monthly Close',6,'0001-01-31 17:00:00','0001-01-20 17:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','0000-00-00 00:00:00',86400000000000,'','2019-02-12 17:11:24',-8,'2019-02-12 17:11:24',-8),(11,1,1,1,'Monthly Close',6,'2019-03-31 17:00:00','2019-03-20 17:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','0000-00-00 00:00:00',86400000000000,'','2019-03-01 18:41:00',-8,'2019-03-01 18:41:00',-8);
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TaskListDefinition`
--

LOCK TABLES `TaskListDefinition` WRITE;
/*!40000 ALTER TABLE `TaskListDefinition` DISABLE KEYS */;
INSERT INTO `TaskListDefinition` VALUES (1,1,'Monthly Close',6,'2018-01-01 00:00:00','2018-01-31 17:00:00','2018-01-20 17:00:00',6,'',86400000000000,'','2018-05-29 18:39:32',211,'2018-03-14 19:50:32',0),(2,1,'Tucasa Apts Period Close',6,'2018-01-01 00:00:00','2018-01-31 00:00:00','2018-01-20 00:00:00',7,'bounce@simulator.amazonses.com',86400000000000,'','2018-05-29 18:15:32',0,'2018-05-29 18:15:32',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Transactant`
--

LOCK TABLES `Transactant` WRITE;
/*!40000 ALTER TABLE `Transactant` DISABLE KEYS */;
INSERT INTO `Transactant` VALUES (1,1,0,'Pablo','Chantell','Pearson','Shenita','Capital One Financial Corp.',0,'PabloPearson178@hotmail.com','PabloPearson203@aol.com','(597) 731-5597','(483) 719-8541','59256 Apache','','Chattanooga','VT','98081','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,1,0,'Gianna','Lamont','Bradley','Awilda','ONEOK Inc',0,'GBradley7523@hotmail.com','GiannaBradley652@bdiddy.com','(289) 548-9478','(387) 664-6906','51991 Johnson','','Burbank','KY','54425','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,1,0,'Dannette','Marth','Galloway','Freeman','Acterna Corp.',0,'DGalloway5134@bdiddy.com','DannetteG1424@abiz.com','(819) 591-4487','(853) 432-1974','75158 Third','','Leominster','NE','28162','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,1,0,'Floria','Rupert','Dean','Deborah','Unisource Energy Corp',0,'FloriaD7418@abiz.com','FDean2110@yahoo.com','(984) 141-4109','(950) 737-2368','22475 Elm','','Savannah','OR','23237','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(5,1,0,'Camilla','Lorilee','Britt','Marge','McLeodUSA Incorporated',0,'CBritt5735@abiz.com','CamillaBritt941@bdiddy.com','(666) 998-2343','(466) 519-4426','6056 Tenth','','Waco','VT','58047','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(6,1,0,'Melvin','Adelle','Norris','Isiah','Robert Half International Inc.',0,'MelvinNorris548@yahoo.com','MNorris7867@aol.com','(566) 401-2880','(313) 373-9214','19907 Jackson','','Pompano Beach','ME','95541','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(7,1,0,'Cornelia','Catarina','Schwartz','Shanon','Starwood Hotels & Resorts Worldwide Inc',0,'CorneliaSchwartz69@abiz.com','CSchwartz8734@hotmail.com','(352) 409-9550','(881) 114-2053','4938 6th','','Peoria','DC','41737','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(8,1,0,'Edris','Sook','Hamilton','Kate','Ecolab Inc.',0,'EdrisH4161@comcast.net','EHamilton6865@yahoo.com','(314) 756-1555','(749) 743-0627','23583 2nd','','Joliet','SD','65194','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(9,1,0,'Mose','Chantelle','Chambers','Celia','Southern Union Company',0,'MoseC3378@comcast.net','MoseChambers151@abiz.com','(517) 661-5174','(567) 173-5432','26165 Quail','','Monterey','AL','16159','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(10,1,0,'Librada','Belva','Orr','Evangelina','Rockwell Automation Inc',0,'LibradaO5688@bdiddy.com','LOrr5260@comcast.net','(609) 941-4223','(970) 654-7177','5448 Fifth','','Yakima','ID','13000','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(11,1,0,'Fatimah','Reena','Shaw','Indira','Shaw Group Inc',0,'FShaw2320@bdiddy.com','FatimahShaw843@hotmail.com','(363) 157-9153','(578) 633-8511','95872 Broadway','','Colorado Springs','IL','72451','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(12,1,0,'Oralee','Tisha','Castaneda','Mallie','URS Corporation',0,'OraleeC3486@yahoo.com','OCastaneda7359@comcast.net','(222) 455-3607','(172) 334-0248','10723 Broadway','','Scottsdale','AZ','75561','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(13,1,0,'Dionna','Joeann','Roberson','Nell','Longs Drug Stores Corporation',0,'DionnaRoberson964@bdiddy.com','DRoberson5456@gmail.com','(448) 741-1031','(540) 400-9516','44942 Seventh','','Kaneohe','OH','89002','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(14,1,0,'Heather','Jimmy','Powers','Royce','Smithfield Foods Inc',0,'HeatherPowers225@bdiddy.com','HPowers7736@yahoo.com','(402) 529-7514','(150) 143-3669','64654 River','','Redding','NY','67996','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(15,1,0,'Grant','Katheleen','Foley','Jaquelyn','GreenPoint Financial Corp.',0,'GrantF9366@abiz.com','GFoley2622@yahoo.com','(562) 401-6008','(682) 354-0287','40653 West Virginia','','Springdale','MN','79241','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(16,1,0,'Georgiana','Nobuko','Dale','Stephaine','The May Department Stores Company',0,'GDale6567@abiz.com','GeorgianaDale604@gmail.com','(575) 941-6702','(115) 219-5766','97420 6th','','Temecula','DE','98878','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(17,1,0,'Ying','Terina','David','Maricela','Audiovox Corporation',0,'YingD6128@abiz.com','YingDavid202@yahoo.com','(155) 161-5144','(626) 317-9918','19956 Evergreen','','Portsmouth','HI','30552','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(18,1,0,'Sherell','Winston','Hunter','Jacquelyne','Genuity Inc.',0,'SHunter4753@bdiddy.com','SherellH4138@gmail.com','(984) 843-5940','(473) 469-0227','57963 New Mexico','','St. Petersburg','CA','31515','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(19,1,0,'Laveta','Bea','Ford','Rebecca','Burlington Industries, Inc.',0,'LavetaF250@comcast.net','LFord5154@yahoo.com','(471) 263-8004','(907) 833-8380','99702 Ridge','','Pasadena','NY','58590','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(20,1,0,'Verdie','Tora','Tyson','Meri','Genuity Inc.',0,'VTyson1112@gmail.com','VTyson6090@yahoo.com','(819) 131-1609','(243) 446-4045','44968 10th','','Salt Lake City','WI','05384','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(21,1,0,'Roxy','Mohamed','Thomas','Lovetta','Abercrombie & Fitch Co.',0,'RoxyT2501@bdiddy.com','RoxyT3672@bdiddy.com','(266) 617-0842','(856) 319-8257','65510 Orchard','','Santa Maria','MD','97726','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(22,1,0,'Mildred','Dolores','Valencia','Jessica','First National of Nebraska Inc.',0,'MildredV6667@bdiddy.com','MildredValencia682@aol.com','(847) 529-8250','(217) 552-0230','63122 Hampton','','Antioch','MT','30493','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(23,1,0,'Humberto','Shara','Vincent','Lakisha','Albertson\'s, Inc.',0,'HVincent4578@bdiddy.com','HumbertoV1416@hotmail.com','(195) 211-7217','(153) 774-3174','96217 Fourth','','Port St. Lucie','MN','44885','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(24,1,0,'Katie','Alida','Carroll','Ciara','Jack In The Box Inc.',0,'KCarroll6457@gmail.com','KCarroll2799@bdiddy.com','(743) 460-9371','(860) 416-6498','19235 Magnolia','','Yonkers','CA','04384','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(25,1,0,'Val','Willodean','Lambert','Sun','Big Lots, Inc.',0,'ValLambert381@abiz.com','ValLambert289@aol.com','(861) 586-8006','(611) 369-3219','88545 Birch','','Waco','NE','03616','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(26,1,0,'Korey','Chasity','Watkins','Myung','Energy East Corporation',0,'KWatkins8471@comcast.net','KWatkins6006@abiz.com','(504) 563-9083','(681) 130-7731','38057 Lehua','','Saginaw','MD','53640','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(27,1,0,'Ruby','Lajuana','Greer','Suzanne','Hollywood Entertainment Corp.',0,'RubyG2904@bdiddy.com','RGreer2941@hotmail.com','(217) 883-3149','(903) 250-2244','30189 Pleasant','','Mesa','DE','54801','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(28,1,0,'Adrienne','Casey','Velasquez','Gabriele','Progressive Corporation',0,'AdrienneVelasquez284@bdiddy.com','AVelasquez7865@abiz.com','(978) 318-7634','(683) 199-2880','48217 Lee','','Albuquerque','AL','77578','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(29,1,0,'Shari','Sarai','Goodwin','Donette','Sara Lee Corp',0,'ShariGoodwin333@yahoo.com','ShariGoodwin454@abiz.com','(771) 472-6336','(876) 496-0112','93459 Hillside','','Santa Cruz','TX','47743','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(30,1,0,'Breann','Gertrudis','Alford','Ester','Emcor Group Inc.',0,'BreannAlford787@yahoo.com','BAlford5872@abiz.com','(996) 952-6930','(422) 218-3748','64966 Fifth','','West Valley City','FL','54904','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(31,1,0,'Karyn','Marquita','King','Blondell','Anheuser-Busch Companies, Inc.',0,'KarynK3477@hotmail.com','KKing9964@gmail.com','(611) 970-4441','(724) 974-4709','89347 Wilson','','Miami','NC','43430','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(32,1,0,'Kattie','Graham','Dickson','Beatris','Ecolab Inc.',0,'KDickson8283@hotmail.com','KattieDickson109@gmail.com','(794) 147-4480','(480) 558-1785','45773 Thirteenth','','Howell','MS','10870','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(33,1,0,'Malka','Blair','Jackson','Tomi','Lucent Technologies Inc.',0,'MJackson8189@bdiddy.com','MJackson6470@yahoo.com','(199) 329-2433','(768) 495-5768','59976 Third','','Louisville','VA','36829','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(34,1,0,'Ashly','Giuseppe','Barron','Roseanne','Pennzoil-Quaker State Company',0,'AshlyBarron751@yahoo.com','AshlyB7335@abiz.com','(429) 502-3448','(153) 154-8786','4581 North','','Mesa','IA','76200','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(35,1,0,'Fidelia','Nadene','Holmes','Darell','Peoples Energy Corp.',0,'FideliaH2174@bdiddy.com','FideliaHolmes711@comcast.net','(530) 858-0551','(982) 752-5503','34849 Lincoln','','Garland','MS','07886','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(36,1,0,'Alexis','Isidra','Whitley','Krissy','Audiovox Corporation',0,'AWhitley538@aol.com','AWhitley8905@gmail.com','(232) 601-9185','(993) 955-1016','91034 Main','','Elizabeth','FL','51888','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(37,1,0,'Quentin','Shante','Petty','Carolynn','Sears Roebuck & Co',0,'QPetty1239@bdiddy.com','QuentinPetty563@bdiddy.com','(195) 731-4732','(218) 802-4724','46459 Williams','','Kaneohe','AR','42019','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(38,1,0,'Herma','Darnell','Austin','Graig','Universal Corporation',0,'HAustin6341@bdiddy.com','HAustin864@abiz.com','(780) 591-1332','(978) 803-0251','52588 1st','','Howell','AL','21853','USA','',0,'','2018-10-17 06:23:18',211,'2018-10-16 19:35:31',0),(39,1,0,'Kandis','Romeo','Lancaster','Herb','Costco Wholesale Corp.',0,'KandisL8067@gmail.com','KandisL1206@hotmail.com','(563) 554-0112','(164) 104-1808','88792 North','','Clearwater','NM','09386','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(40,1,0,'Jacqueline','Lida','Cleveland','Ayesha','US Oncology Inc',0,'JacquelineC9488@aol.com','JacquelineCleveland4@abiz.com','(587) 868-2140','(730) 500-0148','87163 Malulani','','Lubbock','IN','81661','USA','',0,'','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(41,1,0,'Bill','M','Smith','Billy','',0,'bill@example.com','','123-456-7890','','','','','','','','',0,'','2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211);
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
INSERT INTO `User` VALUES (1,1,0,'1974-07-17','Sigrid Sanford','30059 Aspen,Indianapolis,VA 27887','(956) 631-6546','SSanford3254@comcast.net','35824 Sycamore,Cathedral City,TX 31847',0,0,304,26,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,1,0,'1962-05-18','Yong Reid','40338 Smith,Utica,MI 22540','(984) 974-7142','YongReid994@aol.com','8549 11th,Mesa,OR 40456',0,0,204,11,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,1,0,'1966-12-22','Zenaida Cooke','34653 North,Sebastian,NC 55089','(406) 574-3300','ZenaidaC8804@hotmail.com','52794 Cypress,Salem,NC 24728',0,0,187,43,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,1,0,'1955-04-14','Arline Stevens','65254 Redwood,Nashua,CT 39106','(157) 927-8775','AStevens9417@hotmail.com','39151 Johnson,Monroe,WV 40495',0,0,283,19,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(5,1,0,'1998-05-19','Earnest Mercer','69030 Hampton,Kissimmee,LA 79947','(666) 885-5602','EMercer7447@comcast.net','26783 Evergreen,San Jose,DC 38287',1,0,359,74,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(6,1,0,'1960-09-28','Tomeka Wright','37305 Kansas,Ogden,TN 80408','(481) 558-2223','TomekaW6684@comcast.net','64984 S 400,Melbourne,NE 07387',0,0,189,63,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(7,1,0,'1959-11-25','Jana Tate','16294 Wood,Santa Ana,OR 60631','(640) 168-9504','JTate3857@hotmail.com','20319 Elm,Kalamazoo,WV 11485',0,0,298,73,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(8,1,0,'1979-11-06','Jodee Burton','69537 Lehua,Tyler,AK 90563','(456) 662-0663','JodeeB4955@hotmail.com','64873 Sunset,Panama City,HI 12433',1,0,342,30,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(9,1,0,'1966-01-16','Tai Higgins','72933 Airport,Lewisville,NM 71353','(432) 599-5836','TaiH8185@gmail.com','28775 Washington,Leominster,UT 51957',0,0,291,45,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(10,1,0,'1975-08-29','Lucretia Riggs','80409 Canyon,Atlantic City,ME 38705','(762) 315-5279','LucretiaRiggs859@abiz.com','84652 New Hampshire,Garden Grove,SD 62888',0,0,218,46,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(11,1,0,'1997-08-27','Kristian Cummings','77137 Navajo,Honolulu,IL 08510','(554) 476-8539','KristianCummings460@gmail.com','65229 Third,New York City,PA 52605',1,0,355,19,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(12,1,0,'1992-10-11','Sherryl Ellison','49240 S 100,Des Moines,NJ 87202','(395) 489-9110','SherrylE6343@hotmail.com','45484 7th,Punta Gorda,AK 34783',0,0,365,19,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(13,1,0,'1970-08-21','Kathrine Henry','86135 Cottonwood,Salem,ID 29718','(839) 737-8076','KHenry1579@yahoo.com','49733 Hill,Hollywood,AZ 25447',1,0,189,31,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(14,1,0,'1982-12-31','Lowell Boone','64559 Williams,Oceanside,FL 06420','(173) 839-4430','LowellB1836@bdiddy.com','59490 Smith,Cathedral City,ND 18623',0,0,382,45,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(15,1,0,'1982-08-20','Barbara Burns','35894 Kukui,Augusta,IA 70059','(906) 482-9145','BarbaraBurns237@gmail.com','29393 North,Kalamazoo,MS 53033',1,0,262,68,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(16,1,0,'1969-10-15','Antionette Adkins','61755 Aspen,Trenton,CO 09336','(782) 149-7559','AntionetteAdkins376@hotmail.com','27469 Washington,Erie,NH 72546',1,0,224,52,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(17,1,0,'1989-02-07','Tobie Kemp','10086 Eighth,Murrieta,ND 99843','(484) 722-0209','TobieKemp959@yahoo.com','3555 9th,Peoria,ME 52205',1,0,290,18,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(18,1,0,'1961-05-20','Coreen Waters','3010 Bay,Buffalo,WA 89757','(739) 823-8591','CoreenW8092@yahoo.com','11105 Johnson,York,WA 03687',0,0,238,57,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(19,1,0,'1987-06-07','Mitzie Owens','10502 New Hampshire,Port Saint Lucie,OH 63632','(654) 515-0550','MOwens1397@hotmail.com','62742 Evergreen,Round Lake Beach,CO 33098',0,0,299,20,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(20,1,0,'1956-01-19','Edmundo Weeks','55088 Sixth,Shreveport,AR 11297','(480) 486-4510','EdmundoWeeks636@hotmail.com','1424 S 400,Pembroke Pines,CT 59267',0,0,176,7,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(21,1,0,'1976-02-13','Crystal Adkins','74352 Pioneer,Riverside,NJ 45802','(116) 388-3690','CrystalA8209@bdiddy.com','65505 Ninth,Newburgh,MN 03981',0,0,343,23,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(22,1,0,'1997-12-07','Janeen Acevedo','88462 New Hampshire,Tallahassee,VA 23086','(648) 840-0195','JAcevedo1988@hotmail.com','17995 Pinon,Elk Grove,MI 49819',1,0,191,21,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(23,1,0,'1993-02-18','Melodie Jarvis','91780 New York,Rockford,NY 45710','(155) 798-6797','MelodieJ4403@bdiddy.com','49372 Airport,Indianapolis,AL 51387',1,0,313,54,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(24,1,0,'1993-01-20','Jerica Atkinson','42459 Shore,Norfolk,NE 57903','(364) 558-0935','JAtkinson303@comcast.net','8997 12th,Las Cruces,IN 51224',0,0,264,6,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(25,1,0,'1996-09-02','Dan Flores','60209 Aspen,Chula Vista,WI 77839','(448) 655-9285','DanF9475@bdiddy.com','53146 Orchard,Charlotte,MN 90540',1,0,361,1,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(26,1,0,'1958-04-25','Jeromy Mcfadden','86357 3rd,Rockford,TN 57351','(446) 141-1375','JeromyM9352@hotmail.com','63298 Cottonwood,Nashua,MN 18844',1,0,186,56,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(27,1,0,'1979-09-07','Tawanna Cox','10983 Jackson,Lacey,DC 00090','(161) 773-3306','TawannaCox4@comcast.net','64946 10th,Scottsdale,NJ 61602',1,0,218,72,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(28,1,0,'1969-07-27','Verlene Hunter','81404 Lehua,Costa Mesa,TX 08154','(368) 759-6729','VerleneH9173@gmail.com','72192 Pioneer,Howell,TX 67822',0,0,286,19,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(29,1,0,'1959-02-14','Murray Mosley','39655 Ridge,Daly City,KS 31968','(815) 405-5921','MMosley2217@bdiddy.com','3403 Park,Murrieta,AL 01166',1,0,223,11,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(30,1,0,'1959-06-23','Maximo Cunningham','30950 Bay,Los Angeles,CO 93162','(401) 347-6884','MCunningham9904@hotmail.com','14613 Navajo,Indianapolis,MD 54657',1,0,181,26,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(31,1,0,'1975-09-01','Evan Hill','75003 S 400,Oceanside,GA 89513','(380) 744-8786','EvanHill79@aol.com','86922 Mesquite,Port Orange,MI 19700',1,0,256,66,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(32,1,0,'1972-12-12','Dovie Aguilar','97789 Eighth,Victorville,NC 12984','(541) 516-8420','DovieAguilar933@abiz.com','8453 North Carolina,New York,AL 68247',0,0,183,29,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(33,1,0,'1965-03-06','Olen Simpson','32035 8th,Havre de Grace,MD 07920','(336) 875-0422','OlenSimpson942@bdiddy.com','53909 Lake,Des Moines,KY 62048',1,0,295,4,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(34,1,0,'1970-07-09','Louie Winters','25796 Shore,McHenry,NE 19456','(255) 222-0450','LWinters2193@gmail.com','54868 9th,Kennewick,DC 56629',1,0,328,58,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(35,1,0,'1976-03-03','Billye French','45570 Tenth,Austin,NY 25320','(853) 555-1956','BillyeFrench305@abiz.com','28649 Williams,Charleston,SC 95399',0,0,213,46,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(36,1,0,'1986-11-25','Demetra Jenkins','56994 Airport,Vero Beach,IA 16611','(416) 644-2745','DJenkins3744@bdiddy.com','77889 Hickory,Jacksonville,RI 69103',0,0,365,7,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(37,1,0,'1974-11-15','Michelle Dunn','32543 13th,Santa Clarita,CO 97807','(367) 740-4045','MichelleD1855@yahoo.com','94765 New York,Laredo,OR 26157',1,0,311,15,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(38,1,0,'1972-05-08','Alex Welch','34446 North Carolina,Athens,OR 38795','(635) 818-5190','AWelch8114@yahoo.com','49875 S 400,Danbury,MA 10417',0,0,383,16,'2018-10-17 06:23:18',211,'2018-10-16 19:35:31',0),(39,1,0,'1959-08-05','Destiny Burns','63355 Meadow,Modesto,OR 42632','(548) 587-9530','DestinyB8283@bdiddy.com','78067 Center,Bryan,MS 82520',0,0,337,39,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(40,1,0,'1971-11-23','Love Sosa','76280 Mountain View,Baltimore,SD 70060','(803) 852-3717','LSosa7137@comcast.net','47837 Pecan,Kansas City,NY 52420',1,0,285,9,'2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(41,1,0,'1972-02-15','Jane Doe','321 Elm Street,  Sprinfield, MO 66666','456-123-7890','jane@example.com','',1,0,177,35,'2018-10-17 06:23:18',211,'2018-10-17 06:23:18',211);
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
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Vehicle`
--

LOCK TABLES `Vehicle` WRITE;
/*!40000 ALTER TABLE `Vehicle` DISABLE KEYS */;
INSERT INTO `Vehicle` VALUES (1,1,1,'car','Mazda','B-Series','Turquoise',1998,'WAU2F7X8N9DLXDEM','VA','CUH0864','2860465','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(2,1,1,'car','Ford','F-Series Super Duty','Navy',2011,'X1MO1JIW0ZPPN3PK','MA','C3Q722U','6386881','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(3,2,1,'car','Bentley','Continental Super','Blue',2010,'ZD423A5WRU7D6390','MT','84E5D1W','0115711','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(4,3,1,'car','Acura','RL','Turquoise',2004,'2WMT2YSOLB1LDFRL','CO','Y3XO905','5343673','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(5,3,1,'car','Infiniti','G','Aluminum',2003,'TRAYLWDE9HSD918P','MD','6UWX824','5385094','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(6,4,1,'car','MINI','Clubman','Yellow',2012,'6G1O6R6N6H0NRMP9','MT','IN4726U','4945075','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(7,5,1,'car','Chevrolet','Camaro','Bronze',1993,'JTQJCVU5WXLZBH5','MO','EL774N5','5629195','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(8,5,1,'car','GMC','Savana 1500','Tan',2005,'TMK6JET6UUFUQQ3L','LA','W000H4H','6248084','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(9,6,1,'car','Oldsmobile','Achieva','Yellow',1993,'SDBMT8TBFIK80CU2','NE','AC9R284','4137604','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(10,7,1,'car','Hyundai','Equus','Red',2013,'MNB7LRTZXB5ONR4E','MA','JLU2786','7960323','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(11,8,1,'car','Chrysler','Prowler','Pink',2001,'JTKVQ1NE7DM80NV','AR','Z9X60U4','0860201','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(12,9,1,'car','Honda','Fit','Bronze',2007,'U6Y6NJ04B2VXKFNC','OK','UKV4926','2548743','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(13,10,1,'car','MINI','Countryman','Tan',2012,'WAG3H30705PYGZS0','NV','5LGH883','7421448','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(14,11,1,'car','Bentley','Continental Flying Spur','Brown',2012,'LZMPGZ3B3F2RMQXB','MT','2S8U29P','8723603','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(15,12,1,'car','Volvo','S80','Rust',2008,'TSMYBM75Z7RH9IP8','MI','TL921S4','1460361','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(16,13,1,'car','Toyota','Tacoma Xtra','Beige',1997,'1G1C7YJ36LEN6TZ7','TN','DH89U06','9323671','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(17,14,1,'car','Dodge','Nitro','Tan',2007,'2FZ2HJOCCSVDLBPF','VA','78WL0Y3','1230593','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(18,15,1,'car','Ferrari','612 Scaglietti','Tan',2008,'MMSVXA8BRYSN90XV','IA','9Y0U7C5','3207892','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(19,16,1,'car','Aston Martin','V8 Vantage','White',2010,'WF0JSKFQRTGPO72Z','LA','G6G950Y','9362602','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(20,16,1,'car','Audi','Coupe GT','Gold',1986,'6U9KGOJSAT5PV50E','SC','D08LF15','5788310','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(21,17,1,'car','Volkswagen','Golf','Red',1986,'X4XB9K2SH7DJ0EEF','HI','L2A21Q8','9252297','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(22,18,1,'car','Chevrolet','Express 2500','Pink',2010,'NMTR1FABQ4H8A1TZ','WY','1E57GT3','9727623','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(23,19,1,'car','Chevrolet','Tahoe','Black',2008,'6T1XVF49J5WHS14P','AR','X2WA702','9996012','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(24,21,1,'car','Infiniti','QX56','Cream',2010,'UU3DUY7WQ9NQKR3Y','MI','D1S98I0','8311197','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(25,23,1,'car','Chevrolet','Monte Carlo','Blue',1997,'MB1QZ3L2MB3NBHXX','VT','TG87T52','6535783','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(26,24,1,'car','Land Rover','Range Rover Sport','Cream',2008,'MATAD5MJSSQDBAL2','VT','8E2C66G','6844428','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(27,25,1,'car','Chevrolet','2500','Silver',1994,'4TYS8XRCR75JW2Z','MO','51A2R2U','1779712','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(28,26,1,'car','Mazda','B-Series','Aluminum',2008,'ZLAXLL97NYO5CR0K','AK','YN2F372','9595660','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(29,27,1,'car','Saturn','S-Series','Gold',1992,'WMAO8HQJNDKXC335','NH','P5G2C02','6849804','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(30,28,1,'car','Dodge','Ram 3500','Copper',2002,'RFB0LJJXMJ4VDUN9','LA','YK5T008','9656435','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(31,30,1,'car','Lexus','LX','Rose',1999,'NNALA616J0QQWA8O','NM','P145B6L','6262785','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(32,31,1,'car','Chevrolet','S10','Yellow',1993,'1NERL2ZHZ92XQSR','WI','J1253XZ','9190989','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(33,32,1,'car','Lincoln','MKZ','Rust',2012,'W0LXP9URW865ZWL8','ND','DF0505I','7802741','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(34,33,1,'car','Chevrolet','Lumina','Cream',1992,'MS0GKZPG41W1VECL','SD','23IQX90','7727976','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(35,34,1,'car','Jaguar','XK Series','Pink',2013,'LAED3GCGSGXUSSXT','SD','Z1L3R05','5484367','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(36,35,1,'car','Mercury','Grand Marquis','Yellow',2010,'SCEL5SKDXWIT7KZ3','NM','6FR76E5','2308718','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(37,36,1,'car','Dodge','Caravan','Cream',1998,'NMC3ZWCFJ5JMH427','TX','260T1AZ','5764000','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(38,37,1,'car','Dodge','Dakota','Copper',1998,'W09N46TCCGS63G0N','AK','C57TU62','4011137','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(39,38,1,'car','Lexus','SC','Silver',2008,'SLPR6AM73VT2PDCD','IN','FO95A60','0806092','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(40,39,1,'car','Nissan','Quest','Pink',2012,'1G4K9E6S4PUMAP8N','WA','Z68G86L','5858524','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0),(41,40,1,'car','Saturn','Astra','Cream',2009,'LVZK5L5NXD3VLRZY','NV','D2753TE','9322362','2018-02-13','2020-03-01','2018-10-16 19:35:31',0,'2018-10-16 19:35:31',0);
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

-- Dump completed on 2019-03-22 23:34:39
