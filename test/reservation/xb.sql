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
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `AR`
--

LOCK TABLES `AR` WRITE;
/*!40000 ALTER TABLE `AR` DISABLE KEYS */;
INSERT INTO `AR` VALUES (1,1,'Application Fee',0,0,0,9,46,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,15.0000,'2018-08-15 02:53:15',211,'2017-11-10 23:24:23',0,0,0),(2,1,'Application Fee (no assessment)',0,1,0,7,46,'Application fee taken, no assessment made','0000-00-00 00:00:00','9999-12-31 00:00:00',5,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(3,1,'Apply Payment',0,1,0,10,9,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 02:59:12',211,'2017-11-10 23:24:23',0,0,0),(4,1,'Bad Debt Write-Off',0,2,0,71,9,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(5,1,'Bank Service Fee (Deposit Account)',0,2,0,72,4,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(6,1,'Bank Service Fee (Operating Account)',0,2,0,72,3,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(7,1,'Broken Window charge',0,0,0,9,59,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:00:40',211,'2017-11-10 23:24:23',0,0,0),(8,1,'Damage Fee',0,0,0,9,59,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:01:10',211,'2017-11-10 23:24:23',0,0,0),(9,1,'Deposit to Deposit Account (FRB96953)',0,1,0,4,6,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(10,1,'Deposit to Operating Account (FRB54320)',0,1,0,3,6,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(11,1,'Electric Base Fee',0,0,0,9,36,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(12,1,'Electric Overage',0,0,0,9,37,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:01:56',211,'2017-11-10 23:24:23',0,0,0),(13,1,'Eviction Fee Reimbursement',0,0,0,9,56,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(14,1,'Auto-Generated Floating Deposit Assessment',0,3,0,9,12,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(15,1,'Receive Floating Security Deposit',0,1,0,6,9,'','0000-00-00 00:00:00','9999-12-31 00:00:00',13,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(16,1,'Gas Base Fee',0,0,0,9,40,'','1900-01-01 00:00:00','9999-12-29 00:00:00',2,50.0000,'2018-07-23 16:16:36',211,'2017-11-10 23:24:23',0,6,4),(17,1,'Gas Base Overage',0,0,0,9,41,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:02:36',211,'2017-11-10 23:24:23',0,0,0),(18,1,'Insufficient Funds Fee',0,0,0,9,48,'','1900-01-01 00:00:00','9999-12-29 00:00:00',64,25.0000,'2018-08-15 03:03:11',211,'2017-11-10 23:24:23',0,0,0),(19,1,'Late Fee',0,0,0,9,47,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,5.0000,'2018-08-15 03:03:22',211,'2017-11-10 23:24:23',0,0,0),(20,1,'Month to Month Fee',0,0,0,9,49,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(21,1,'No Show / Termination Fee',0,0,0,9,51,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:03:52',211,'2017-11-10 23:24:23',0,0,0),(22,1,'Other Special Tenant Charges',0,0,0,9,61,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(23,1,'Pet Fee',0,0,0,9,52,'','1900-01-01 00:00:00','9999-12-31 00:00:00',192,50.0000,'2018-07-04 04:13:35',211,'2017-11-10 23:24:23',0,0,0),(24,1,'Pet Rent',0,0,0,9,53,'','1900-01-01 00:00:00','9999-12-30 00:00:00',144,10.0000,'2018-07-23 16:16:52',211,'2017-11-10 23:24:23',0,6,4),(25,1,'Receive a Payment',0,1,0,6,10,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(26,1,'Rent Non-Taxable',0,0,0,9,18,'','1900-01-01 00:00:00','9999-12-30 00:00:00',16,0.0000,'2018-07-23 16:14:54',211,'2017-11-10 23:24:23',0,6,4),(27,1,'Rent Taxable',0,0,0,9,17,'','1900-01-01 00:00:00','9999-12-30 00:00:00',16,0.0000,'2018-07-23 16:15:28',211,'2017-11-10 23:24:23',0,4,0),(28,1,'Security Deposit Assessment',0,0,0,9,11,'normal deposit','1900-01-01 00:00:00','9999-12-30 00:00:00',96,0.0000,'2018-08-15 03:04:36',211,'2017-11-10 23:24:23',0,0,0),(29,1,'Security Deposit Forfeiture',0,0,0,11,58,'Forfeit','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(30,1,'Security Deposit Refund',0,0,0,11,5,'Refund','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(31,1,'Special Cleaning Fee',0,0,0,9,55,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(32,1,'Tenant Expense Chargeback',0,0,0,9,54,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(33,1,'Vending Income',0,1,0,7,65,'','0000-00-00 00:00:00','9999-12-31 00:00:00',5,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(34,1,'Water and Sewer Base Fee',0,0,0,9,38,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(35,1,'Water and Sewer Overage',0,0,0,9,39,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:05:54',211,'2017-11-10 23:24:23',0,0,0),(36,1,'Auto-gen Application Fee Asmt',0,3,0,9,46,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(37,1,'Receive Application Fee (auto-gen asmt)',0,1,0,6,9,'Application fee taken, autogen asmt','0000-00-00 00:00:00','9999-12-31 00:00:00',13,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(38,1,'XFER  Operating to SecDep',0,2,0,4,3,'Move money from Operating acct to Sec Dep','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(39,1,'Vehicle Registration Fee',0,0,3,9,75,'','2018-01-01 00:00:00','9999-12-31 00:00:00',320,10.0000,'2018-07-04 04:14:56',211,'2018-07-03 02:47:47',211,0,0),(40,1,'Rent ST000',0,0,0,9,18,'Default rent assessment for rentable type RType000','1970-01-01 00:00:00','9999-12-30 00:00:00',16,1000.0000,'2018-08-15 03:07:17',211,'2018-07-27 06:49:30',0,6,4),(41,1,'Rent ST001',0,0,0,9,18,'Default rent assessment for rentable type RType001','1970-01-01 00:00:00','9999-12-30 00:00:00',16,1500.0000,'2018-08-15 03:07:28',211,'2018-07-27 06:49:30',0,6,4),(42,1,'Rent ST002',0,0,0,9,18,'Default rent assessment for rentable type RType002','1970-01-01 00:00:00','9999-12-30 00:00:00',16,1750.0000,'2018-08-15 03:07:40',211,'2018-07-27 06:49:30',0,6,4),(43,1,'Rent ST003',0,0,0,9,18,'Default rent assessment for rentable type RType003','1970-01-01 00:00:00','9999-12-30 00:00:00',16,2500.0000,'2018-08-15 03:16:30',211,'2018-07-27 06:49:30',0,6,4),(44,1,'Rent CP000',0,0,0,9,18,'Default rent assessment for rentable type Car Port 000','1970-01-01 00:00:00','9999-12-30 00:00:00',16,35.0000,'2018-08-15 03:16:45',211,'2018-07-27 06:49:30',0,6,4),(45,1,'Rent 1BR1BA',0,0,0,9,18,'Default rent assessment for rentable type Standard','1970-01-01 00:00:00','9999-12-31 00:00:00',20,1000.0000,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,6,4),(46,1,'Rent 2BR1BA',0,0,0,9,18,'Default rent assessment for rentable type Deluxe','1970-01-01 00:00:00','9999-12-31 00:00:00',20,1500.0000,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,6,4),(47,1,'Rent 2BR2BA',0,0,0,9,18,'Default rent assessment for rentable type Gold','1970-01-01 00:00:00','9999-12-31 00:00:00',20,1750.0000,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,6,4),(48,1,'Rent 3BR2BA',0,0,0,9,18,'Default rent assessment for rentable type Platinum','1970-01-01 00:00:00','9999-12-31 00:00:00',20,2500.0000,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,6,4),(49,1,'Rent H11',0,0,0,9,18,'Default rent assessment for rentable type HotelStd','1970-01-01 00:00:00','9999-12-31 00:00:00',20,65.0000,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,6,4),(50,1,'Rent H21',0,0,0,9,18,'Default rent assessment for rentable type HotelGold','1970-01-01 00:00:00','9999-12-31 00:00:00',20,75.0000,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,6,4),(51,1,'Rent H22',0,0,0,9,18,'Default rent assessment for rentable type HotelPlatinum','1970-01-01 00:00:00','9999-12-31 00:00:00',20,85.0000,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,6,4),(52,1,'Rent CP000',0,0,0,9,18,'Default rent assessment for rentable type Car Port 000','1970-01-01 00:00:00','9999-12-31 00:00:00',20,35.0000,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,0,0);
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
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Assessments`
--

LOCK TABLES `Assessments` WRITE;
/*!40000 ALTER TABLE `Assessments` DISABLE KEYS */;
INSERT INTO `Assessments` VALUES (1,0,0,0,1,1,0,0,1,1000.0000,'2019-02-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',26,0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,0,0,0,1,1,0,0,1,2000.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',28,2,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,0,0,0,1,1,14,1,1,10.0000,'2019-02-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',24,8,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,0,0,0,1,1,0,0,1,935.4800,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',26,2,'prorated for 29 of 31 days','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(5,0,0,0,1,1,14,1,1,50.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',23,2,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(6,0,0,0,1,1,14,1,1,9.3500,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',24,2,'prorated for 29 of 31 days','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(7,0,0,0,1,1,15,1,1,10.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',39,2,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(8,0,0,0,1,2,0,0,2,1000.0000,'2019-02-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',26,0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(9,0,0,0,1,2,0,0,2,2000.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',28,2,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(10,0,0,0,1,2,14,2,2,10.0000,'2019-02-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',24,8,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(11,0,0,0,1,2,0,0,2,935.4800,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',26,2,'prorated for 29 of 31 days','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(12,0,0,0,1,2,14,2,2,50.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',23,2,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(13,0,0,0,1,2,14,2,2,9.3500,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',24,2,'prorated for 29 of 31 days','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(14,0,0,0,1,2,15,2,2,10.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',39,2,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(15,0,0,0,1,3,0,0,3,1500.0000,'2019-02-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',26,0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(16,0,0,0,1,3,0,0,3,3000.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',28,2,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(17,0,0,0,1,3,0,0,3,1403.2300,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',26,2,'prorated for 29 of 31 days','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(18,0,0,0,1,4,0,0,4,1500.0000,'2019-02-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',26,0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(19,0,0,0,1,4,0,0,4,3000.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',28,2,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(20,0,0,0,1,4,14,3,4,10.0000,'2019-02-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',24,8,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(21,0,0,0,1,4,0,0,4,1403.2300,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',26,2,'prorated for 29 of 31 days','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(22,0,0,0,1,4,14,3,4,50.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',23,2,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(23,0,0,0,1,4,14,3,4,9.3500,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',24,2,'prorated for 29 of 31 days','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(24,0,0,0,1,4,15,3,4,10.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',39,2,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(25,15,0,0,1,3,0,0,3,1500.0000,'2019-05-01 00:00:00','2019-05-01 00:00:00',6,4,0,'',26,0,'','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(26,18,0,0,1,4,0,0,4,1500.0000,'2019-05-01 00:00:00','2019-05-01 00:00:00',6,4,0,'',26,0,'','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(27,1,0,0,1,1,0,0,1,1000.0000,'2019-05-01 00:00:00','2019-05-01 00:00:00',6,4,0,'',26,0,'','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(28,8,0,0,1,2,0,0,2,1000.0000,'2019-05-01 00:00:00','2019-05-01 00:00:00',6,4,0,'',26,0,'','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(29,3,0,0,1,1,14,1,1,10.0000,'2019-05-01 00:00:00','2019-05-01 00:00:00',6,4,0,'',24,0,'','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(30,10,0,0,1,2,14,2,2,10.0000,'2019-05-01 00:00:00','2019-05-01 00:00:00',6,4,0,'',24,0,'','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(31,20,0,0,1,4,14,3,4,10.0000,'2019-05-01 00:00:00','2019-05-01 00:00:00',6,4,0,'',24,0,'','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1);
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
INSERT INTO `BusinessProperties` VALUES (1,1,'general',0,'{\"Epochs\": {\"Daily\": \"2017-01-01T00:00:00Z\", \"Weekly\": \"2017-01-01T00:00:00Z\", \"Yearly\": \"2017-01-01T00:00:00Z\", \"Monthly\": \"2017-01-01T00:00:00Z\", \"Quarterly\": \"2017-01-01T00:00:00Z\"}, \"PetFees\": [\"Pet Fee\", \"Pet Rent\"], \"VehicleFees\": [\"Vehicle Registration Fee\"]}','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
INSERT INTO `Deposit` VALUES (1,1,2,1,'2019-01-03',4885.4700,0.0000,0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,1,1,1,'2019-01-03',10000.0000,0.0000,0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DepositPart`
--

LOCK TABLES `DepositPart` WRITE;
/*!40000 ALTER TABLE `DepositPart` DISABLE KEYS */;
INSERT INTO `DepositPart` VALUES (1,1,1,2,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,1,1,3,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,1,1,4,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,1,1,5,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(5,1,1,7,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(6,1,1,8,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(7,1,1,9,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(8,1,1,10,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(9,1,1,12,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(10,1,1,14,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(11,1,1,15,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(12,1,1,16,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(13,1,1,17,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(14,2,1,1,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(15,2,1,6,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(16,2,1,11,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(17,2,1,13,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Journal`
--

LOCK TABLES `Journal` WRITE;
/*!40000 ALTER TABLE `Journal` DISABLE KEYS */;
INSERT INTO `Journal` VALUES (1,1,'2019-01-03 00:00:00',2000.0000,1,2,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,1,'2019-01-03 00:00:00',935.4800,1,4,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,1,'2019-01-03 00:00:00',50.0000,1,5,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,1,'2019-01-03 00:00:00',9.3500,1,6,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(5,1,'2019-01-03 00:00:00',10.0000,1,7,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(6,1,'2019-01-03 00:00:00',2000.0000,1,9,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(7,1,'2019-01-03 00:00:00',935.4800,1,11,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(8,1,'2019-01-03 00:00:00',50.0000,1,12,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(9,1,'2019-01-03 00:00:00',9.3500,1,13,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(10,1,'2019-01-03 00:00:00',10.0000,1,14,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(11,1,'2019-01-03 00:00:00',3000.0000,1,16,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(12,1,'2019-01-03 00:00:00',1403.2300,1,17,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(13,1,'2019-01-03 00:00:00',3000.0000,1,19,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(14,1,'2019-01-03 00:00:00',1403.2300,1,21,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(15,1,'2019-01-03 00:00:00',50.0000,1,22,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(16,1,'2019-01-03 00:00:00',9.3500,1,23,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(17,1,'2019-01-03 00:00:00',10.0000,1,24,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(18,1,'2019-01-03 00:00:00',2000.0000,2,1,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(19,1,'2019-01-03 00:00:00',935.4800,2,2,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(20,1,'2019-01-03 00:00:00',50.0000,2,3,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(21,1,'2019-01-03 00:00:00',9.3500,2,4,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(22,1,'2019-01-03 00:00:00',10.0000,2,5,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(23,1,'2019-01-03 00:00:00',2000.0000,2,6,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(24,1,'2019-01-03 00:00:00',935.4800,2,7,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(25,1,'2019-01-03 00:00:00',50.0000,2,8,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(26,1,'2019-01-03 00:00:00',9.3500,2,9,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(27,1,'2019-01-03 00:00:00',10.0000,2,10,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(28,1,'2019-01-03 00:00:00',3000.0000,2,11,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(29,1,'2019-01-03 00:00:00',1403.2300,2,12,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(30,1,'2019-01-03 00:00:00',3000.0000,2,13,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(31,1,'2019-01-03 00:00:00',1403.2300,2,14,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(32,1,'2019-01-03 00:00:00',50.0000,2,15,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(33,1,'2019-01-03 00:00:00',9.3500,2,16,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(34,1,'2019-01-03 00:00:00',10.0000,2,17,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(35,1,'2019-01-03 00:00:00',2000.0000,2,1,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(36,1,'2019-01-03 00:00:00',935.4800,2,2,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(37,1,'2019-01-03 00:00:00',50.0000,2,3,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(38,1,'2019-01-03 00:00:00',9.3500,2,4,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(39,1,'2019-01-03 00:00:00',10.0000,2,5,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(40,1,'2019-01-03 00:00:00',2000.0000,2,6,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(41,1,'2019-01-03 00:00:00',935.4800,2,7,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(42,1,'2019-01-03 00:00:00',50.0000,2,8,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(43,1,'2019-01-03 00:00:00',9.3500,2,9,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(44,1,'2019-01-03 00:00:00',10.0000,2,10,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(45,1,'2019-01-03 00:00:00',3000.0000,2,11,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(46,1,'2019-01-03 00:00:00',1403.2300,2,12,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(47,1,'2019-01-03 00:00:00',3000.0000,2,13,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(48,1,'2019-01-03 00:00:00',1403.2300,2,14,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(49,1,'2019-01-03 00:00:00',50.0000,2,15,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(50,1,'2019-01-03 00:00:00',9.3500,2,16,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(51,1,'2019-01-03 00:00:00',10.0000,2,17,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(52,1,'2019-01-03 00:00:00',935.4800,4,2,'auto-transfer for deposit DEP-2','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(53,1,'2019-01-03 00:00:00',50.0000,4,3,'auto-transfer for deposit DEP-2','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(54,1,'2019-01-03 00:00:00',9.3500,4,4,'auto-transfer for deposit DEP-2','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(55,1,'2019-01-03 00:00:00',10.0000,4,5,'auto-transfer for deposit DEP-2','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(56,1,'2019-01-03 00:00:00',935.4800,4,7,'auto-transfer for deposit DEP-2','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(57,1,'2019-01-03 00:00:00',50.0000,4,8,'auto-transfer for deposit DEP-2','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(58,1,'2019-01-03 00:00:00',9.3500,4,9,'auto-transfer for deposit DEP-2','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(59,1,'2019-01-03 00:00:00',10.0000,4,10,'auto-transfer for deposit DEP-2','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(60,1,'2019-01-03 00:00:00',1403.2300,4,12,'auto-transfer for deposit DEP-2','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(61,1,'2019-01-03 00:00:00',1403.2300,4,14,'auto-transfer for deposit DEP-2','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(62,1,'2019-01-03 00:00:00',50.0000,4,15,'auto-transfer for deposit DEP-2','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(63,1,'2019-01-03 00:00:00',9.3500,4,16,'auto-transfer for deposit DEP-2','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(64,1,'2019-01-03 00:00:00',10.0000,4,17,'auto-transfer for deposit DEP-2','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(65,1,'2019-01-03 00:00:00',2000.0000,4,1,'auto-transfer for deposit DEP-1','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(66,1,'2019-01-03 00:00:00',2000.0000,4,6,'auto-transfer for deposit DEP-1','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(67,1,'2019-01-03 00:00:00',3000.0000,4,11,'auto-transfer for deposit DEP-1','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(68,1,'2019-01-03 00:00:00',3000.0000,4,13,'auto-transfer for deposit DEP-1','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(69,1,'2019-05-01 00:00:00',0.0000,1,25,'','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(70,1,'2019-05-01 00:00:00',0.0000,1,26,'','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(71,1,'2019-05-01 00:00:00',0.0000,1,27,'','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(72,1,'2019-05-01 00:00:00',0.0000,1,28,'','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(73,1,'2019-05-01 00:00:00',0.0000,1,29,'','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(74,1,'2019-05-01 00:00:00',0.0000,1,30,'','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(75,1,'2019-05-01 00:00:00',0.0000,1,31,'','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1);
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
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `JournalAllocation`
--

LOCK TABLES `JournalAllocation` WRITE;
/*!40000 ALTER TABLE `JournalAllocation` DISABLE KEYS */;
INSERT INTO `JournalAllocation` VALUES (1,1,1,1,1,0,0,2000.0000,2,0,'d 12001 2000.00, c 30000 2000.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,1,2,1,1,0,0,935.4800,4,0,'d 12001 935.48, c 41001 935.48','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,1,3,1,1,0,0,50.0000,5,0,'d 12001 50.00, c 41407 50.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,1,4,1,1,0,0,9.3500,6,0,'d 12001 9.35, c 41408 9.35','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(5,1,5,1,1,0,0,10.0000,7,0,'d 12001 10.00, c 41416 10.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(6,1,6,2,2,0,0,2000.0000,9,0,'d 12001 2000.00, c 30000 2000.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(7,1,7,2,2,0,0,935.4800,11,0,'d 12001 935.48, c 41001 935.48','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(8,1,8,2,2,0,0,50.0000,12,0,'d 12001 50.00, c 41407 50.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(9,1,9,2,2,0,0,9.3500,13,0,'d 12001 9.35, c 41408 9.35','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(10,1,10,2,2,0,0,10.0000,14,0,'d 12001 10.00, c 41416 10.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(11,1,11,3,3,0,0,3000.0000,16,0,'d 12001 3000.00, c 30000 3000.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(12,1,12,3,3,0,0,1403.2300,17,0,'d 12001 1403.23, c 41001 1403.23','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(13,1,13,4,4,0,0,3000.0000,19,0,'d 12001 3000.00, c 30000 3000.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(14,1,14,4,4,0,0,1403.2300,21,0,'d 12001 1403.23, c 41001 1403.23','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(15,1,15,4,4,0,0,50.0000,22,0,'d 12001 50.00, c 41407 50.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(16,1,16,4,4,0,0,9.3500,23,0,'d 12001 9.35, c 41408 9.35','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(17,1,17,4,4,0,0,10.0000,24,0,'d 12001 10.00, c 41416 10.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(18,1,18,0,0,1,0,2000.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(19,1,19,0,0,1,0,935.4800,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(20,1,20,0,0,1,0,50.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(21,1,21,0,0,1,0,9.3500,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(22,1,22,0,0,1,0,10.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(23,1,23,0,0,2,0,2000.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(24,1,24,0,0,2,0,935.4800,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(25,1,25,0,0,2,0,50.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(26,1,26,0,0,2,0,9.3500,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(27,1,27,0,0,2,0,10.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(28,1,28,0,0,3,0,3000.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(29,1,29,0,0,3,0,1403.2300,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(30,1,30,0,0,4,0,3000.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(31,1,31,0,0,4,0,1403.2300,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(32,1,32,0,0,4,0,50.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(33,1,33,0,0,4,0,9.3500,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(34,1,34,0,0,4,0,10.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(35,1,35,1,1,1,1,2000.0000,2,0,'ASM(2) d 12999 2000.00,c 12001 2000.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(36,1,36,1,1,1,2,935.4800,4,0,'ASM(4) d 12999 935.48,c 12001 935.48','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(37,1,37,1,1,1,3,50.0000,5,0,'ASM(5) d 12999 50.00,c 12001 50.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(38,1,38,1,1,1,4,9.3500,6,0,'ASM(6) d 12999 9.35,c 12001 9.35','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(39,1,39,1,1,1,5,10.0000,7,0,'ASM(7) d 12999 10.00,c 12001 10.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(40,1,40,2,2,2,6,2000.0000,9,0,'ASM(9) d 12999 2000.00,c 12001 2000.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(41,1,41,2,2,2,7,935.4800,11,0,'ASM(11) d 12999 935.48,c 12001 935.48','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(42,1,42,2,2,2,8,50.0000,12,0,'ASM(12) d 12999 50.00,c 12001 50.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(43,1,43,2,2,2,9,9.3500,13,0,'ASM(13) d 12999 9.35,c 12001 9.35','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(44,1,44,2,2,2,10,10.0000,14,0,'ASM(14) d 12999 10.00,c 12001 10.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(45,1,45,3,3,3,11,3000.0000,16,0,'ASM(16) d 12999 3000.00,c 12001 3000.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(46,1,46,3,3,3,12,1403.2300,17,0,'ASM(17) d 12999 1403.23,c 12001 1403.23','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(47,1,47,4,4,4,13,3000.0000,19,0,'ASM(19) d 12999 3000.00,c 12001 3000.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(48,1,48,4,4,4,14,1403.2300,21,0,'ASM(21) d 12999 1403.23,c 12001 1403.23','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(49,1,49,4,4,4,15,50.0000,22,0,'ASM(22) d 12999 50.00,c 12001 50.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(50,1,50,4,4,4,16,9.3500,23,0,'ASM(23) d 12999 9.35,c 12001 9.35','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(51,1,51,4,4,4,17,10.0000,24,0,'ASM(24) d 12999 10.00,c 12001 10.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(52,1,52,0,1,1,2,935.4800,4,0,'d 10105 935.4800, c 10999 935.4800','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(53,1,53,0,1,1,3,50.0000,5,0,'d 10105 50.0000, c 10999 50.0000','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(54,1,54,0,1,1,4,9.3500,6,0,'d 10105 9.3500, c 10999 9.3500','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(55,1,55,0,1,1,5,10.0000,7,0,'d 10105 10.0000, c 10999 10.0000','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(56,1,56,0,2,2,7,935.4800,11,0,'d 10105 935.4800, c 10999 935.4800','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(57,1,57,0,2,2,8,50.0000,12,0,'d 10105 50.0000, c 10999 50.0000','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(58,1,58,0,2,2,9,9.3500,13,0,'d 10105 9.3500, c 10999 9.3500','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(59,1,59,0,2,2,10,10.0000,14,0,'d 10105 10.0000, c 10999 10.0000','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(60,1,60,0,3,3,12,1403.2300,17,0,'d 10105 1403.2300, c 10999 1403.2300','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(61,1,61,0,4,4,14,1403.2300,21,0,'d 10105 1403.2300, c 10999 1403.2300','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(62,1,62,0,4,4,15,50.0000,22,0,'d 10105 50.0000, c 10999 50.0000','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(63,1,63,0,4,4,16,9.3500,23,0,'d 10105 9.3500, c 10999 9.3500','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(64,1,64,0,4,4,17,10.0000,24,0,'d 10105 10.0000, c 10999 10.0000','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(65,1,65,0,1,1,1,2000.0000,2,0,'d 10104 2000.0000, c 10999 2000.0000','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(66,1,66,0,2,2,6,2000.0000,9,0,'d 10104 2000.0000, c 10999 2000.0000','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(67,1,67,0,3,3,11,3000.0000,16,0,'d 10104 3000.0000, c 10999 3000.0000','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(68,1,68,0,4,4,13,3000.0000,19,0,'d 10104 3000.0000, c 10999 3000.0000','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(69,1,69,3,3,0,0,0.0000,25,0,'d 12001 0.00, c 41001 0.00','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(70,1,70,4,4,0,0,0.0000,26,0,'d 12001 0.00, c 41001 0.00','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(71,1,71,1,1,0,0,0.0000,27,0,'d 12001 0.00, c 41001 0.00','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(72,1,72,2,2,0,0,0.0000,28,0,'d 12001 0.00, c 41001 0.00','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(73,1,73,1,1,0,0,0.0000,29,0,'d 12001 0.00, c 41408 0.00','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(74,1,74,2,2,0,0,0.0000,30,0,'d 12001 0.00, c 41408 0.00','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1),(75,1,75,4,4,0,0,0.0000,31,0,'d 12001 0.00, c 41408 0.00','2019-05-16 08:22:11',-1,'2019-05-16 08:22:11',-1);
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
) ENGINE=InnoDB AUTO_INCREMENT=137 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerEntry`
--

LOCK TABLES `LedgerEntry` WRITE;
/*!40000 ALTER TABLE `LedgerEntry` DISABLE KEYS */;
INSERT INTO `LedgerEntry` VALUES (1,1,1,1,9,1,1,0,'2019-01-03 00:00:00',2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,1,1,1,11,1,1,0,'2019-01-03 00:00:00',-2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,1,2,2,9,1,1,0,'2019-01-03 00:00:00',935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,1,2,2,18,1,1,0,'2019-01-03 00:00:00',-935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(5,1,3,3,9,1,1,0,'2019-01-03 00:00:00',50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(6,1,3,3,52,1,1,0,'2019-01-03 00:00:00',-50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(7,1,4,4,9,1,1,0,'2019-01-03 00:00:00',9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(8,1,4,4,53,1,1,0,'2019-01-03 00:00:00',-9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(9,1,5,5,9,1,1,0,'2019-01-03 00:00:00',10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(10,1,5,5,75,1,1,0,'2019-01-03 00:00:00',-10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(11,1,6,6,9,2,2,0,'2019-01-03 00:00:00',2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(12,1,6,6,11,2,2,0,'2019-01-03 00:00:00',-2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(13,1,7,7,9,2,2,0,'2019-01-03 00:00:00',935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(14,1,7,7,18,2,2,0,'2019-01-03 00:00:00',-935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(15,1,8,8,9,2,2,0,'2019-01-03 00:00:00',50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(16,1,8,8,52,2,2,0,'2019-01-03 00:00:00',-50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(17,1,9,9,9,2,2,0,'2019-01-03 00:00:00',9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(18,1,9,9,53,2,2,0,'2019-01-03 00:00:00',-9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(19,1,10,10,9,2,2,0,'2019-01-03 00:00:00',10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(20,1,10,10,75,2,2,0,'2019-01-03 00:00:00',-10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(21,1,11,11,9,3,3,0,'2019-01-03 00:00:00',3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(22,1,11,11,11,3,3,0,'2019-01-03 00:00:00',-3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(23,1,12,12,9,3,3,0,'2019-01-03 00:00:00',1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(24,1,12,12,18,3,3,0,'2019-01-03 00:00:00',-1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(25,1,13,13,9,4,4,0,'2019-01-03 00:00:00',3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(26,1,13,13,11,4,4,0,'2019-01-03 00:00:00',-3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(27,1,14,14,9,4,4,0,'2019-01-03 00:00:00',1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(28,1,14,14,18,4,4,0,'2019-01-03 00:00:00',-1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(29,1,15,15,9,4,4,0,'2019-01-03 00:00:00',50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(30,1,15,15,52,4,4,0,'2019-01-03 00:00:00',-50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(31,1,16,16,9,4,4,0,'2019-01-03 00:00:00',9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(32,1,16,16,53,4,4,0,'2019-01-03 00:00:00',-9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(33,1,17,17,9,4,4,0,'2019-01-03 00:00:00',10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(34,1,17,17,75,4,4,0,'2019-01-03 00:00:00',-10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(35,1,18,18,6,0,0,1,'2019-01-03 00:00:00',2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(36,1,18,18,10,0,0,1,'2019-01-03 00:00:00',-2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(37,1,19,19,6,0,0,1,'2019-01-03 00:00:00',935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(38,1,19,19,10,0,0,1,'2019-01-03 00:00:00',-935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(39,1,20,20,6,0,0,1,'2019-01-03 00:00:00',50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(40,1,20,20,10,0,0,1,'2019-01-03 00:00:00',-50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(41,1,21,21,6,0,0,1,'2019-01-03 00:00:00',9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(42,1,21,21,10,0,0,1,'2019-01-03 00:00:00',-9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(43,1,22,22,6,0,0,1,'2019-01-03 00:00:00',10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(44,1,22,22,10,0,0,1,'2019-01-03 00:00:00',-10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(45,1,23,23,6,0,0,2,'2019-01-03 00:00:00',2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(46,1,23,23,10,0,0,2,'2019-01-03 00:00:00',-2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(47,1,24,24,6,0,0,2,'2019-01-03 00:00:00',935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(48,1,24,24,10,0,0,2,'2019-01-03 00:00:00',-935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(49,1,25,25,6,0,0,2,'2019-01-03 00:00:00',50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(50,1,25,25,10,0,0,2,'2019-01-03 00:00:00',-50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(51,1,26,26,6,0,0,2,'2019-01-03 00:00:00',9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(52,1,26,26,10,0,0,2,'2019-01-03 00:00:00',-9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(53,1,27,27,6,0,0,2,'2019-01-03 00:00:00',10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(54,1,27,27,10,0,0,2,'2019-01-03 00:00:00',-10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(55,1,28,28,6,0,0,3,'2019-01-03 00:00:00',3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(56,1,28,28,10,0,0,3,'2019-01-03 00:00:00',-3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(57,1,29,29,6,0,0,3,'2019-01-03 00:00:00',1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(58,1,29,29,10,0,0,3,'2019-01-03 00:00:00',-1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(59,1,30,30,6,0,0,4,'2019-01-03 00:00:00',3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(60,1,30,30,10,0,0,4,'2019-01-03 00:00:00',-3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(61,1,31,31,6,0,0,4,'2019-01-03 00:00:00',1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(62,1,31,31,10,0,0,4,'2019-01-03 00:00:00',-1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(63,1,32,32,6,0,0,4,'2019-01-03 00:00:00',50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(64,1,32,32,10,0,0,4,'2019-01-03 00:00:00',-50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(65,1,33,33,6,0,0,4,'2019-01-03 00:00:00',9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(66,1,33,33,10,0,0,4,'2019-01-03 00:00:00',-9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(67,1,34,34,6,0,0,4,'2019-01-03 00:00:00',10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(68,1,34,34,10,0,0,4,'2019-01-03 00:00:00',-10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(69,1,35,35,10,1,1,1,'2019-01-03 00:00:00',2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(70,1,35,35,9,1,1,1,'2019-01-03 00:00:00',-2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(71,1,36,36,10,1,1,1,'2019-01-03 00:00:00',935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(72,1,36,36,9,1,1,1,'2019-01-03 00:00:00',-935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(73,1,37,37,10,1,1,1,'2019-01-03 00:00:00',50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(74,1,37,37,9,1,1,1,'2019-01-03 00:00:00',-50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(75,1,38,38,10,1,1,1,'2019-01-03 00:00:00',9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(76,1,38,38,9,1,1,1,'2019-01-03 00:00:00',-9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(77,1,39,39,10,1,1,1,'2019-01-03 00:00:00',10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(78,1,39,39,9,1,1,1,'2019-01-03 00:00:00',-10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(79,1,40,40,10,2,2,2,'2019-01-03 00:00:00',2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(80,1,40,40,9,2,2,2,'2019-01-03 00:00:00',-2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(81,1,41,41,10,2,2,2,'2019-01-03 00:00:00',935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(82,1,41,41,9,2,2,2,'2019-01-03 00:00:00',-935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(83,1,42,42,10,2,2,2,'2019-01-03 00:00:00',50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(84,1,42,42,9,2,2,2,'2019-01-03 00:00:00',-50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(85,1,43,43,10,2,2,2,'2019-01-03 00:00:00',9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(86,1,43,43,9,2,2,2,'2019-01-03 00:00:00',-9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(87,1,44,44,10,2,2,2,'2019-01-03 00:00:00',10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(88,1,44,44,9,2,2,2,'2019-01-03 00:00:00',-10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(89,1,45,45,10,3,3,3,'2019-01-03 00:00:00',3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(90,1,45,45,9,3,3,3,'2019-01-03 00:00:00',-3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(91,1,46,46,10,3,3,3,'2019-01-03 00:00:00',1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(92,1,46,46,9,3,3,3,'2019-01-03 00:00:00',-1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(93,1,47,47,10,4,4,4,'2019-01-03 00:00:00',3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(94,1,47,47,9,4,4,4,'2019-01-03 00:00:00',-3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(95,1,48,48,10,4,4,4,'2019-01-03 00:00:00',1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(96,1,48,48,9,4,4,4,'2019-01-03 00:00:00',-1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(97,1,49,49,10,4,4,4,'2019-01-03 00:00:00',50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(98,1,49,49,9,4,4,4,'2019-01-03 00:00:00',-50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(99,1,50,50,10,4,4,4,'2019-01-03 00:00:00',9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(100,1,50,50,9,4,4,4,'2019-01-03 00:00:00',-9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(101,1,51,51,10,4,4,4,'2019-01-03 00:00:00',10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(102,1,51,51,9,4,4,4,'2019-01-03 00:00:00',-10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(103,1,52,52,4,1,0,1,'2019-01-03 00:00:00',935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(104,1,52,52,6,1,0,1,'2019-01-03 00:00:00',-935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(105,1,53,53,4,1,0,1,'2019-01-03 00:00:00',50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(106,1,53,53,6,1,0,1,'2019-01-03 00:00:00',-50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(107,1,54,54,4,1,0,1,'2019-01-03 00:00:00',9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(108,1,54,54,6,1,0,1,'2019-01-03 00:00:00',-9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(109,1,55,55,4,1,0,1,'2019-01-03 00:00:00',10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(110,1,55,55,6,1,0,1,'2019-01-03 00:00:00',-10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(111,1,56,56,4,2,0,2,'2019-01-03 00:00:00',935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(112,1,56,56,6,2,0,2,'2019-01-03 00:00:00',-935.4800,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(113,1,57,57,4,2,0,2,'2019-01-03 00:00:00',50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(114,1,57,57,6,2,0,2,'2019-01-03 00:00:00',-50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(115,1,58,58,4,2,0,2,'2019-01-03 00:00:00',9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(116,1,58,58,6,2,0,2,'2019-01-03 00:00:00',-9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(117,1,59,59,4,2,0,2,'2019-01-03 00:00:00',10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(118,1,59,59,6,2,0,2,'2019-01-03 00:00:00',-10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(119,1,60,60,4,3,0,3,'2019-01-03 00:00:00',1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(120,1,60,60,6,3,0,3,'2019-01-03 00:00:00',-1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(121,1,61,61,4,4,0,4,'2019-01-03 00:00:00',1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(122,1,61,61,6,4,0,4,'2019-01-03 00:00:00',-1403.2300,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(123,1,62,62,4,4,0,4,'2019-01-03 00:00:00',50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(124,1,62,62,6,4,0,4,'2019-01-03 00:00:00',-50.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(125,1,63,63,4,4,0,4,'2019-01-03 00:00:00',9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(126,1,63,63,6,4,0,4,'2019-01-03 00:00:00',-9.3500,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(127,1,64,64,4,4,0,4,'2019-01-03 00:00:00',10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(128,1,64,64,6,4,0,4,'2019-01-03 00:00:00',-10.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(129,1,65,65,3,1,0,1,'2019-01-03 00:00:00',2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(130,1,65,65,6,1,0,1,'2019-01-03 00:00:00',-2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(131,1,66,66,3,2,0,2,'2019-01-03 00:00:00',2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(132,1,66,66,6,2,0,2,'2019-01-03 00:00:00',-2000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(133,1,67,67,3,3,0,3,'2019-01-03 00:00:00',3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(134,1,67,67,6,3,0,3,'2019-01-03 00:00:00',-3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(135,1,68,68,3,4,0,4,'2019-01-03 00:00:00',3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(136,1,68,68,6,4,0,4,'2019-01-03 00:00:00',-3000.0000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerMarker`
--

LOCK TABLES `LedgerMarker` WRITE;
/*!40000 ALTER TABLE `LedgerMarker` DISABLE KEYS */;
INSERT INTO `LedgerMarker` VALUES (1,1,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(2,2,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(3,3,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(4,4,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(5,5,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(6,6,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(7,7,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(8,8,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(9,9,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(10,10,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(11,11,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(12,12,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(13,13,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(14,14,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(15,15,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(16,16,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(17,17,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(18,18,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(19,19,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(20,20,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(21,21,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(22,22,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(23,23,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(24,24,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(25,25,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(26,26,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(27,27,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(28,28,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(29,29,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(30,30,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(31,31,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(32,32,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(33,33,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(34,34,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(35,35,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(36,36,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(37,37,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(38,38,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(39,39,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(40,40,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(41,41,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(42,42,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(43,43,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(44,44,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(45,45,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(46,46,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(47,47,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(48,48,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(49,49,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(50,50,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(51,51,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(52,52,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(53,53,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(54,54,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(55,55,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(56,56,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(57,57,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(58,58,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(59,59,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(60,60,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(61,61,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(62,62,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(63,63,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(64,64,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(65,65,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(66,66,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(67,67,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(68,68,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(69,69,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(70,70,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(71,71,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(72,72,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(73,73,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(74,74,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(75,75,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-07-03 02:45:37',211,'2018-07-03 02:45:37',211),(76,0,1,1,0,0,'2018-12-20 00:00:00',0.0000,3,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(77,0,1,1,1,0,'2018-12-20 00:00:00',0.0000,3,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(78,0,1,2,0,0,'2018-12-20 00:00:00',0.0000,3,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(79,0,1,2,2,0,'2018-12-20 00:00:00',0.0000,3,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(80,0,1,3,0,0,'2018-12-20 00:00:00',0.0000,3,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(81,0,1,3,3,0,'2018-12-20 00:00:00',0.0000,3,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(82,0,1,4,0,0,'2018-12-20 00:00:00',0.0000,3,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(83,0,1,4,4,0,'2018-12-20 00:00:00',0.0000,3,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
INSERT INTO `Payor` VALUES (1,1,'c84aa5a39fed16758596d37925bd0730f30b52b0794f5d3b3bb776c235b3c428f35bbd78',14710.0000,1,0,'c2e18691979503b073f62b07114bc8a179e51e168ad970da746fd39e58247ba4f79289f5',127392.0000,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,1,'beec17b9623770d879432b28af139605c22c651fcf562de2ed97d5c5e3fc81d8b092caed',5893.0000,1,0,'66f41d59994969b66b3ef8fd6a629b6a2192b2b1828b7528dcd7662a13ab625f6b3b34c5',51694.0000,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,1,'b0e8471c028fd06c9d8f098d03360918c657827f8b847803ccfdc98b1f99f90eed067d3a',23417.0000,1,0,'76b6710bbc03139735613e13f062ef2584e70cd9f95feb41767a5b35296a44b148f5b20f',99126.0000,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,1,'a00fb20fffa9ad34ccaba4436e2a0a51c646b331525dc5ed5f82025158ecee8f1e158634',6524.0000,1,0,'85a33b81af012149bb4e0df5e2adf3328078d7085bf60b73f0a1105af3f7bfcd52bbd9ca',123991.0000,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Pets`
--

LOCK TABLES `Pets` WRITE;
/*!40000 ALTER TABLE `Pets` DISABLE KEYS */;
INSERT INTO `Pets` VALUES (1,1,1,1,'cat','Chantilly-Tiffany','Harlequin',16.0000,'Precious ','2019-01-03','2020-03-01','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,1,2,2,'dog','Lagotto Romagnolo','Blue',10.0000,'Thor','2019-01-03','2020-03-01','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,1,4,4,'cat','Arabian Mau','Dilute',23.0000,'Belle ','2019-01-03','2020-03-01','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
INSERT INTO `Prospect` VALUES (1,1,'70253 Willow','Flint','OR','84059','3ComCorpF2374@comcast.net','(468) 821-4643','special needs teacher (learning/behavioural difficulties)',0,'','','','','','0000-00-00','61751 Apache, Harrisburg, SC 02081','Nyla Wilkinson','(581) 543-7787',99,'9 years 8 months','72677 Cottonwood, Wichita, OR 41318','Nena Lindsey','(931) 293-9053',88,'3 years 5 months','','Ozella Dillon','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,1,'83557 Aloha','Arvada','GA','03300','BArvada2715@aol.com','(175) 279-4173','art photographer',0,'','','','','','0000-00-00','18019 North Dakota, Elizabeth, OH 10694','Myra Bond','(460) 124-1381',118,'6 years 4 months','76655 Mesquite, North Port, WI 78511','Katelynn Roberts','(158) 123-0235',94,'2 months','','Annetta Robles','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,1,'32506 Pecan','Oxnard','MS','33274','JOxnard6337@bdiddy.com','(178) 395-2640','lecturer/researcher in linguistics',0,'','','','','','0000-00-00','90225 Juniper, Norwich, NJ 11211','Caterina Greer','(874) 368-1450',119,'10 years 5 months','73054 Apache, Lorain, NH 31445','Alexis Munoz','(835) 825-7669',127,'2 years 5 months','','Francesco Wilkins','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,1,'60883 Cottonwood','Miramar','TN','65466','IMiramar5412@abiz.com','(901) 276-1315','flight attendant (steward/ess, cabin staff)',0,'','','','','','0000-00-00','88092 Wilson, Glendale, TN 11528','Gale Garner','(561) 321-9584',129,'3 years 7 months','83749 Church, Buffalo, MT 86258','Chiquita Stein','(166) 808-2254',127,'6 years 8 months','','Lakita Rice','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Receipt`
--

LOCK TABLES `Receipt` WRITE;
/*!40000 ALTER TABLE `Receipt` DISABLE KEYS */;
INSERT INTO `Receipt` VALUES (1,0,1,1,2,2,2,1,'2019-01-03 00:00:00','786623',2000.0000,'',25,'ASM(2) d 12999 2000.00,c 12001 2000.00',2,'payment for ASM-2','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,0,1,1,2,1,1,1,'2019-01-03 00:00:00','988459',935.4800,'',25,'ASM(4) d 12999 935.48,c 12001 935.48',2,'payment for ASM-4','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,0,1,1,2,1,1,1,'2019-01-03 00:00:00','641803',50.0000,'',25,'ASM(5) d 12999 50.00,c 12001 50.00',2,'payment for ASM-5','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,0,1,1,2,1,1,1,'2019-01-03 00:00:00','361739',9.3500,'',25,'ASM(6) d 12999 9.35,c 12001 9.35',2,'payment for ASM-6','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(5,0,1,1,2,1,1,1,'2019-01-03 00:00:00','761883',10.0000,'',25,'ASM(7) d 12999 10.00,c 12001 10.00',2,'payment for ASM-7','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(6,0,1,2,2,2,2,2,'2019-01-03 00:00:00','38907',2000.0000,'',25,'ASM(9) d 12999 2000.00,c 12001 2000.00',2,'payment for ASM-9','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(7,0,1,2,2,1,1,2,'2019-01-03 00:00:00','658932',935.4800,'',25,'ASM(11) d 12999 935.48,c 12001 935.48',2,'payment for ASM-11','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(8,0,1,2,2,1,1,2,'2019-01-03 00:00:00','387780',50.0000,'',25,'ASM(12) d 12999 50.00,c 12001 50.00',2,'payment for ASM-12','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(9,0,1,2,2,1,1,2,'2019-01-03 00:00:00','383115',9.3500,'',25,'ASM(13) d 12999 9.35,c 12001 9.35',2,'payment for ASM-13','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(10,0,1,2,2,1,1,2,'2019-01-03 00:00:00','742716',10.0000,'',25,'ASM(14) d 12999 10.00,c 12001 10.00',2,'payment for ASM-14','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(11,0,1,3,2,2,2,3,'2019-01-03 00:00:00','50020',3000.0000,'',25,'ASM(16) d 12999 3000.00,c 12001 3000.00',2,'payment for ASM-16','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(12,0,1,3,2,1,1,3,'2019-01-03 00:00:00','623071',1403.2300,'',25,'ASM(17) d 12999 1403.23,c 12001 1403.23',2,'payment for ASM-17','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(13,0,1,4,2,2,2,4,'2019-01-03 00:00:00','134652',3000.0000,'',25,'ASM(19) d 12999 3000.00,c 12001 3000.00',2,'payment for ASM-19','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(14,0,1,4,2,1,1,4,'2019-01-03 00:00:00','281907',1403.2300,'',25,'ASM(21) d 12999 1403.23,c 12001 1403.23',2,'payment for ASM-21','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(15,0,1,4,2,1,1,4,'2019-01-03 00:00:00','772719',50.0000,'',25,'ASM(22) d 12999 50.00,c 12001 50.00',2,'payment for ASM-22','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(16,0,1,4,2,1,1,4,'2019-01-03 00:00:00','659062',9.3500,'',25,'ASM(23) d 12999 9.35,c 12001 9.35',2,'payment for ASM-23','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(17,0,1,4,2,1,1,4,'2019-01-03 00:00:00','663110',10.0000,'',25,'ASM(24) d 12999 10.00,c 12001 10.00',2,'payment for ASM-24','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ReceiptAllocation`
--

LOCK TABLES `ReceiptAllocation` WRITE;
/*!40000 ALTER TABLE `ReceiptAllocation` DISABLE KEYS */;
INSERT INTO `ReceiptAllocation` VALUES (1,1,1,1,'2019-01-03 00:00:00',2000.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,2,1,1,'2019-01-03 00:00:00',935.4800,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,3,1,1,'2019-01-03 00:00:00',50.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,4,1,1,'2019-01-03 00:00:00',9.3500,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(5,5,1,1,'2019-01-03 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(6,6,1,2,'2019-01-03 00:00:00',2000.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(7,7,1,2,'2019-01-03 00:00:00',935.4800,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(8,8,1,2,'2019-01-03 00:00:00',50.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(9,9,1,2,'2019-01-03 00:00:00',9.3500,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(10,10,1,2,'2019-01-03 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(11,11,1,3,'2019-01-03 00:00:00',3000.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(12,12,1,3,'2019-01-03 00:00:00',1403.2300,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(13,13,1,4,'2019-01-03 00:00:00',3000.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(14,14,1,4,'2019-01-03 00:00:00',1403.2300,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(15,15,1,4,'2019-01-03 00:00:00',50.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(16,16,1,4,'2019-01-03 00:00:00',9.3500,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(17,17,1,4,'2019-01-03 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(18,1,1,1,'2019-01-03 00:00:00',2000.0000,2,0,'ASM(2) d 12999 2000.00,c 12001 2000.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(19,2,1,1,'2019-01-03 00:00:00',935.4800,4,0,'ASM(4) d 12999 935.48,c 12001 935.48','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(20,3,1,1,'2019-01-03 00:00:00',50.0000,5,0,'ASM(5) d 12999 50.00,c 12001 50.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(21,4,1,1,'2019-01-03 00:00:00',9.3500,6,0,'ASM(6) d 12999 9.35,c 12001 9.35','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(22,5,1,1,'2019-01-03 00:00:00',10.0000,7,0,'ASM(7) d 12999 10.00,c 12001 10.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(23,6,1,2,'2019-01-03 00:00:00',2000.0000,9,0,'ASM(9) d 12999 2000.00,c 12001 2000.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(24,7,1,2,'2019-01-03 00:00:00',935.4800,11,0,'ASM(11) d 12999 935.48,c 12001 935.48','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(25,8,1,2,'2019-01-03 00:00:00',50.0000,12,0,'ASM(12) d 12999 50.00,c 12001 50.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(26,9,1,2,'2019-01-03 00:00:00',9.3500,13,0,'ASM(13) d 12999 9.35,c 12001 9.35','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(27,10,1,2,'2019-01-03 00:00:00',10.0000,14,0,'ASM(14) d 12999 10.00,c 12001 10.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(28,11,1,3,'2019-01-03 00:00:00',3000.0000,16,0,'ASM(16) d 12999 3000.00,c 12001 3000.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(29,12,1,3,'2019-01-03 00:00:00',1403.2300,17,0,'ASM(17) d 12999 1403.23,c 12001 1403.23','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(30,13,1,4,'2019-01-03 00:00:00',3000.0000,19,0,'ASM(19) d 12999 3000.00,c 12001 3000.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(31,14,1,4,'2019-01-03 00:00:00',1403.2300,21,0,'ASM(21) d 12999 1403.23,c 12001 1403.23','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(32,15,1,4,'2019-01-03 00:00:00',50.0000,22,0,'ASM(22) d 12999 50.00,c 12001 50.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(33,16,1,4,'2019-01-03 00:00:00',9.3500,23,0,'ASM(23) d 12999 9.35,c 12001 9.35','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(34,17,1,4,'2019-01-03 00:00:00',10.0000,24,0,'ASM(24) d 12999 10.00,c 12001 10.00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(35,2,1,1,'2019-01-03 00:00:00',935.4800,0,0,'d 10105 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(36,3,1,1,'2019-01-03 00:00:00',50.0000,0,0,'d 10105 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(37,4,1,1,'2019-01-03 00:00:00',9.3500,0,0,'d 10105 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(38,5,1,1,'2019-01-03 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(39,7,1,2,'2019-01-03 00:00:00',935.4800,0,0,'d 10105 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(40,8,1,2,'2019-01-03 00:00:00',50.0000,0,0,'d 10105 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(41,9,1,2,'2019-01-03 00:00:00',9.3500,0,0,'d 10105 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(42,10,1,2,'2019-01-03 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(43,12,1,3,'2019-01-03 00:00:00',1403.2300,0,0,'d 10105 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(44,14,1,4,'2019-01-03 00:00:00',1403.2300,0,0,'d 10105 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(45,15,1,4,'2019-01-03 00:00:00',50.0000,0,0,'d 10105 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(46,16,1,4,'2019-01-03 00:00:00',9.3500,0,0,'d 10105 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(47,17,1,4,'2019-01-03 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(48,1,1,1,'2019-01-03 00:00:00',2000.0000,0,0,'d 10104 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(49,6,1,2,'2019-01-03 00:00:00',2000.0000,0,0,'d 10104 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(50,11,1,3,'2019-01-03 00:00:00',3000.0000,0,0,'d 10104 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(51,13,1,4,'2019-01-03 00:00:00',3000.0000,0,0,'d 10104 _, c 10999 _','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Rentable`
--

LOCK TABLES `Rentable` WRITE;
/*!40000 ALTER TABLE `Rentable` DISABLE KEYS */;
INSERT INTO `Rentable` VALUES (1,1,0,'Rentable001',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,''),(2,1,0,'Rentable002',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,''),(3,1,0,'Rentable003',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,''),(4,1,0,'Rentable004',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,''),(5,1,0,'Rentable005',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,''),(6,1,0,'Rentable006',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,''),(7,1,0,'Rentable007',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,''),(8,1,0,'Rentable008',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,''),(9,1,0,'Rentable009',0,0,'0000-00-00 00:00:00','2019-05-16 08:23:00',211,'2019-01-17 01:12:42',0,''),(10,1,0,'Rentable010',0,0,'0000-00-00 00:00:00','2019-05-16 08:22:29',211,'2019-01-17 01:12:42',0,''),(11,1,0,'Rentable011',0,0,'0000-00-00 00:00:00','2019-05-16 08:23:24',211,'2019-01-17 01:12:42',0,''),(12,1,0,'Rentable012',0,0,'0000-00-00 00:00:00','2019-05-16 08:23:37',211,'2019-01-17 01:12:42',0,''),(13,1,0,'Rentable013',0,0,'0000-00-00 00:00:00','2019-05-16 08:23:54',211,'2019-01-17 01:12:42',0,''),(14,1,0,'Rentable014',0,0,'0000-00-00 00:00:00','2019-05-16 08:24:08',211,'2019-01-17 01:12:42',0,''),(15,1,0,'CP001',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,''),(16,1,0,'CP002',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,''),(17,1,0,'CP003',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,''),(18,1,0,'CP004',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,''),(19,1,0,'CP005',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,''),(20,1,0,'CP006',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,''),(21,1,0,'CP007',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,''),(22,1,0,'CP008',0,0,'0000-00-00 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0,'');
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
  `RAID` bigint(20) NOT NULL DEFAULT '0',
  `LeaseStatus` smallint(6) NOT NULL DEFAULT '0',
  `DtStart` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `DtStop` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `Comment` varchar(2048) NOT NULL DEFAULT '',
  `ConfirmationCode` varchar(20) NOT NULL DEFAULT '',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RLID`)
) ENGINE=InnoDB AUTO_INCREMENT=857 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableLeaseStatus`
--

LOCK TABLES `RentableLeaseStatus` WRITE;
/*!40000 ALTER TABLE `RentableLeaseStatus` DISABLE KEYS */;
INSERT INTO `RentableLeaseStatus` VALUES (1,1,1,0,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,2,1,0,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,3,1,0,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,4,1,0,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(5,5,1,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(6,6,1,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(7,7,1,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(8,8,1,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(9,9,1,0,0,'2019-01-01 00:00:00','2019-01-16 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(10,10,1,0,0,'2019-01-01 00:00:00','2019-01-04 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(11,11,1,0,0,'2019-01-01 00:00:00','2019-01-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(12,12,1,0,0,'2019-01-01 00:00:00','2019-01-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(13,13,1,0,0,'2019-01-01 00:00:00','2019-01-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(14,14,1,0,0,'2019-01-01 00:00:00','2019-01-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(15,15,1,0,0,'2019-01-01 00:00:00','2019-01-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(16,16,1,0,0,'2019-01-01 00:00:00','2019-01-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(17,17,1,0,0,'2019-01-01 00:00:00','2019-01-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(18,18,1,0,0,'2019-01-01 00:00:00','2019-01-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(19,19,1,0,0,'2019-01-01 00:00:00','2019-01-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(20,20,1,0,0,'2019-01-01 00:00:00','2019-02-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(21,21,1,0,0,'2019-01-01 00:00:00','2019-01-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(22,22,1,0,0,'2019-01-01 00:00:00','2019-01-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(23,1,1,0,2,'2020-03-01 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(24,1,1,0,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(25,2,1,0,2,'2020-03-01 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(26,2,1,0,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(27,3,1,0,2,'2020-03-01 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(28,3,1,0,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(29,4,1,0,2,'2020-03-01 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(30,4,1,0,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(31,9,1,0,0,'2019-02-07 00:00:00','2019-02-14 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(32,9,1,0,2,'2019-01-31 00:00:00','2019-02-07 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(33,9,1,0,0,'2019-12-12 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(34,9,1,0,2,'2019-12-07 00:00:00','2019-12-08 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(35,9,1,0,0,'2019-02-20 00:00:00','2019-02-24 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(36,9,1,0,2,'2019-02-16 00:00:00','2019-02-20 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(37,9,1,0,0,'2019-02-15 00:00:00','2019-02-16 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(38,9,1,0,2,'2019-02-14 00:00:00','2019-02-15 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(39,9,1,0,0,'2019-07-09 00:00:00','2019-07-16 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(40,9,1,0,2,'2019-07-07 00:00:00','2019-07-09 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(41,9,1,0,0,'2019-10-06 00:00:00','2019-10-11 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(42,9,1,0,2,'2019-10-04 00:00:00','2019-10-06 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(43,9,1,0,0,'2019-07-20 00:00:00','2019-07-31 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(44,9,1,0,2,'2019-07-16 00:00:00','2019-07-20 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(45,9,1,0,0,'2019-03-20 00:00:00','2019-03-24 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(46,9,1,0,2,'2019-03-15 00:00:00','2019-03-20 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(47,9,1,0,0,'2019-10-01 00:00:00','2019-10-04 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(48,9,1,0,2,'2019-09-24 00:00:00','2019-09-27 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(49,9,1,0,0,'2019-02-27 00:00:00','2019-03-15 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(50,9,1,0,2,'2019-02-24 00:00:00','2019-02-27 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(51,9,1,0,0,'2019-07-06 00:00:00','2019-07-07 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(52,9,1,0,2,'2019-07-05 00:00:00','2019-07-06 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(53,9,1,0,0,'2019-09-22 00:00:00','2019-09-24 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(54,9,1,0,2,'2019-09-21 00:00:00','2019-09-22 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(55,9,1,0,0,'2019-03-28 00:00:00','2019-04-28 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(56,9,1,0,2,'2019-03-24 00:00:00','2019-03-28 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(57,9,1,0,0,'2019-09-18 00:00:00','2019-09-21 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(58,9,1,0,2,'2019-09-12 00:00:00','2019-09-18 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(59,9,1,0,0,'2019-08-01 00:00:00','2019-08-20 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(60,9,1,0,2,'2019-07-31 00:00:00','2019-08-01 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(61,9,1,0,2,'2019-12-10 00:00:00','2019-12-12 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(62,9,1,0,2,'2019-12-08 00:00:00','2019-12-10 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(63,9,1,0,0,'2019-11-05 00:00:00','2019-12-07 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(64,9,1,0,2,'2019-10-29 00:00:00','2019-11-05 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(65,9,1,0,0,'2019-05-04 00:00:00','2019-05-09 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(66,9,1,0,2,'2019-04-28 00:00:00','2019-05-04 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(67,9,1,0,2,'2019-09-29 00:00:00','2019-10-01 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(68,9,1,0,2,'2019-09-27 00:00:00','2019-09-29 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(69,9,1,0,0,'2019-10-18 00:00:00','2019-10-22 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(70,9,1,0,2,'2019-10-11 00:00:00','2019-10-18 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(71,9,1,0,0,'2019-08-22 00:00:00','2019-09-12 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(72,9,1,0,2,'2019-08-20 00:00:00','2019-08-22 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(73,9,1,0,0,'2019-01-17 00:00:00','2019-01-21 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(74,9,1,0,2,'2019-01-16 00:00:00','2019-01-17 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(75,9,1,0,0,'2019-01-23 00:00:00','2019-01-26 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(76,9,1,0,2,'2019-01-21 00:00:00','2019-01-23 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(77,9,1,0,0,'2019-10-27 00:00:00','2019-10-29 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(78,9,1,0,2,'2019-10-22 00:00:00','2019-10-27 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(79,9,1,0,0,'2019-05-16 00:00:00','2019-06-07 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(80,9,1,0,2,'2019-05-09 00:00:00','2019-05-16 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(81,9,1,0,0,'2019-01-29 00:00:00','2019-01-31 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(82,9,1,0,2,'2019-01-26 00:00:00','2019-01-29 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(83,9,1,0,0,'2019-06-09 00:00:00','2019-07-05 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(84,9,1,0,2,'2019-06-07 00:00:00','2019-06-09 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(85,10,1,0,0,'2019-05-20 00:00:00','2019-05-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(86,10,1,0,2,'2019-05-13 00:00:00','2019-05-20 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(87,10,1,0,0,'2019-04-21 00:00:00','2019-04-27 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(88,10,1,0,2,'2019-04-15 00:00:00','2019-04-21 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(89,10,1,0,0,'2019-09-24 00:00:00','2019-10-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(90,10,1,0,2,'2019-09-17 00:00:00','2019-09-24 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(91,10,1,0,0,'2019-12-14 00:00:00','2019-12-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(92,10,1,0,2,'2019-12-12 00:00:00','2019-12-14 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(93,10,1,0,0,'2019-11-12 00:00:00','2019-11-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(94,10,1,0,2,'2019-11-05 00:00:00','2019-11-12 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(95,10,1,0,0,'2019-07-10 00:00:00','2019-08-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(96,10,1,0,2,'2019-07-07 00:00:00','2019-07-10 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(97,10,1,0,0,'2019-01-08 00:00:00','2019-01-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(98,10,1,0,2,'2019-01-04 00:00:00','2019-01-08 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(99,10,1,0,0,'2019-05-02 00:00:00','2019-05-13 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(100,10,1,0,2,'2019-04-27 00:00:00','2019-05-02 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(101,10,1,0,0,'2019-09-06 00:00:00','2019-09-09 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(102,10,1,0,2,'2019-09-03 00:00:00','2019-09-06 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(103,10,1,0,0,'2019-09-14 00:00:00','2019-09-17 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(104,10,1,0,2,'2019-09-09 00:00:00','2019-09-14 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(105,10,1,0,0,'2019-06-30 00:00:00','2019-07-07 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(106,10,1,0,2,'2019-06-24 00:00:00','2019-06-26 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(107,10,1,0,0,'2019-03-17 00:00:00','2019-03-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(108,10,1,0,2,'2019-03-11 00:00:00','2019-03-17 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(109,10,1,0,0,'2019-10-23 00:00:00','2019-11-05 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(110,10,1,0,2,'2019-10-17 00:00:00','2019-10-23 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(111,10,1,0,0,'2019-06-15 00:00:00','2019-06-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:42',0),(112,10,1,0,2,'2019-06-08 00:00:00','2019-06-15 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(113,10,1,0,2,'2019-06-29 00:00:00','2019-06-30 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(114,10,1,0,2,'2019-06-26 00:00:00','2019-06-29 00:00:00','','',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(115,10,1,0,0,'2019-12-25 00:00:00','2019-12-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(116,10,1,0,2,'2019-12-22 00:00:00','2019-12-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(117,10,1,0,0,'2019-01-15 00:00:00','2019-03-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(118,10,1,0,2,'2019-01-11 00:00:00','2019-01-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(119,10,1,0,0,'2019-08-07 00:00:00','2019-08-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(120,10,1,0,2,'2019-08-03 00:00:00','2019-08-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(121,10,1,0,0,'2019-08-20 00:00:00','2019-09-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(122,10,1,0,2,'2019-08-19 00:00:00','2019-08-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(123,10,1,0,0,'2019-04-11 00:00:00','2019-04-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(124,10,1,0,2,'2019-04-10 00:00:00','2019-04-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(125,10,1,0,0,'2019-06-19 00:00:00','2019-06-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(126,10,1,0,2,'2019-06-18 00:00:00','2019-06-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(127,10,1,0,0,'2019-10-11 00:00:00','2019-10-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(128,10,1,0,2,'2019-10-08 00:00:00','2019-10-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(129,10,1,0,0,'2019-12-31 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(130,10,1,0,2,'2019-12-30 00:00:00','2019-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(131,10,1,0,0,'2019-10-07 00:00:00','2019-10-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(132,10,1,0,2,'2019-10-05 00:00:00','2019-10-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(133,10,1,0,0,'2019-05-31 00:00:00','2019-06-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(134,10,1,0,2,'2019-05-30 00:00:00','2019-05-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(135,10,1,0,0,'2019-03-28 00:00:00','2019-04-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(136,10,1,0,2,'2019-03-24 00:00:00','2019-03-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(137,10,1,0,0,'2019-12-06 00:00:00','2019-12-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(138,10,1,0,2,'2019-11-30 00:00:00','2019-12-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(139,11,1,0,0,'2019-10-27 00:00:00','2019-11-01 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(140,11,1,0,2,'2019-10-25 00:00:00','2019-10-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(141,11,1,0,0,'2019-06-08 00:00:00','2019-06-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(142,11,1,0,2,'2019-06-05 00:00:00','2019-06-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(143,11,1,0,0,'2019-04-12 00:00:00','2019-04-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(144,11,1,0,2,'2019-04-10 00:00:00','2019-04-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(145,11,1,0,0,'2019-06-26 00:00:00','2019-07-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(146,11,1,0,2,'2019-06-23 00:00:00','2019-06-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(147,11,1,0,0,'2019-01-14 00:00:00','2019-01-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(148,11,1,0,2,'2019-01-08 00:00:00','2019-01-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(149,11,1,0,0,'2019-10-04 00:00:00','2019-10-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(150,11,1,0,2,'2019-09-27 00:00:00','2019-10-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(151,11,1,0,0,'2019-03-20 00:00:00','2019-04-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(152,11,1,0,2,'2019-03-17 00:00:00','2019-03-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(153,11,1,0,0,'2019-03-15 00:00:00','2019-03-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(154,11,1,0,2,'2019-03-13 00:00:00','2019-03-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(155,11,1,0,0,'2019-05-11 00:00:00','2019-05-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(156,11,1,0,2,'2019-05-06 00:00:00','2019-05-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(157,11,1,0,0,'2019-11-26 00:00:00','2019-11-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(158,11,1,0,2,'2019-11-23 00:00:00','2019-11-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(159,11,1,0,0,'2019-03-05 00:00:00','2019-03-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(160,11,1,0,2,'2019-03-02 00:00:00','2019-03-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(161,11,1,0,0,'2019-03-07 00:00:00','2019-03-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(162,11,1,0,2,'2019-03-06 00:00:00','2019-03-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(163,11,1,0,0,'2019-06-03 00:00:00','2019-06-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(164,11,1,0,2,'2019-05-30 00:00:00','2019-06-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(165,11,1,0,0,'2019-04-05 00:00:00','2019-04-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(166,11,1,0,2,'2019-04-02 00:00:00','2019-04-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(167,11,1,0,0,'2019-04-28 00:00:00','2019-05-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(168,11,1,0,2,'2019-04-23 00:00:00','2019-04-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(169,11,1,0,0,'2019-07-11 00:00:00','2019-07-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(170,11,1,0,2,'2019-07-07 00:00:00','2019-07-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(171,11,1,0,0,'2019-11-12 00:00:00','2019-11-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(172,11,1,0,2,'2019-11-08 00:00:00','2019-11-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(173,11,1,0,0,'2019-05-16 00:00:00','2019-05-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(174,11,1,0,2,'2019-05-12 00:00:00','2019-05-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(175,11,1,0,0,'2019-08-13 00:00:00','2019-09-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(176,11,1,0,2,'2019-08-06 00:00:00','2019-08-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(177,11,1,0,0,'2019-06-18 00:00:00','2019-06-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(178,11,1,0,2,'2019-06-12 00:00:00','2019-06-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(179,11,1,0,0,'2019-04-20 00:00:00','2019-04-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(180,11,1,0,2,'2019-04-17 00:00:00','2019-04-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(181,11,1,0,0,'2019-10-10 00:00:00','2019-10-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(182,11,1,0,2,'2019-10-09 00:00:00','2019-10-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(183,11,1,0,0,'2019-01-28 00:00:00','2019-02-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(184,11,1,0,2,'2019-01-25 00:00:00','2019-01-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(185,11,1,0,0,'2019-02-13 00:00:00','2019-02-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(186,11,1,0,2,'2019-02-12 00:00:00','2019-02-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(187,11,1,0,0,'2019-02-22 00:00:00','2019-03-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(188,11,1,0,2,'2019-02-17 00:00:00','2019-02-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(189,11,1,0,0,'2019-12-03 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(190,11,1,0,2,'2019-11-30 00:00:00','2019-12-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(191,11,1,0,0,'2019-07-23 00:00:00','2019-08-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(192,11,1,0,2,'2019-07-18 00:00:00','2019-07-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(193,11,1,0,0,'2019-11-07 00:00:00','2019-11-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(194,11,1,0,2,'2019-11-01 00:00:00','2019-11-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(195,11,1,0,2,'2019-05-09 00:00:00','2019-05-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(196,11,1,0,2,'2019-05-08 00:00:00','2019-05-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(197,12,1,0,0,'2019-02-26 00:00:00','2019-03-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(198,12,1,0,2,'2019-02-23 00:00:00','2019-02-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(199,12,1,0,0,'2019-07-31 00:00:00','2019-08-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(200,12,1,0,2,'2019-07-25 00:00:00','2019-07-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(201,12,1,0,0,'2019-07-08 00:00:00','2019-07-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(202,12,1,0,2,'2019-07-03 00:00:00','2019-07-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(203,12,1,0,0,'2019-12-04 00:00:00','2019-12-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(204,12,1,0,2,'2019-11-27 00:00:00','2019-12-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(205,12,1,0,0,'2019-05-28 00:00:00','2019-07-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(206,12,1,0,2,'2019-05-27 00:00:00','2019-05-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(207,12,1,0,0,'2019-02-14 00:00:00','2019-02-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(208,12,1,0,2,'2019-02-10 00:00:00','2019-02-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(209,12,1,0,0,'2019-12-27 00:00:00','2019-12-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(210,12,1,0,2,'2019-12-22 00:00:00','2019-12-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(211,12,1,0,0,'2019-07-13 00:00:00','2019-07-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(212,12,1,0,2,'2019-07-12 00:00:00','2019-07-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(213,12,1,0,0,'2019-10-13 00:00:00','2019-10-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(214,12,1,0,2,'2019-10-07 00:00:00','2019-10-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(215,12,1,0,0,'2019-08-22 00:00:00','2019-09-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(216,12,1,0,2,'2019-08-21 00:00:00','2019-08-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(217,12,1,0,0,'2019-10-01 00:00:00','2019-10-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(218,12,1,0,2,'2019-09-24 00:00:00','2019-10-01 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(219,12,1,0,0,'2019-01-14 00:00:00','2019-02-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(220,12,1,0,2,'2019-01-08 00:00:00','2019-01-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(221,12,1,0,0,'2019-04-03 00:00:00','2019-04-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(222,12,1,0,2,'2019-03-31 00:00:00','2019-04-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(223,12,1,0,0,'2019-05-09 00:00:00','2019-05-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(224,12,1,0,2,'2019-05-03 00:00:00','2019-05-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(225,12,1,0,0,'2019-04-22 00:00:00','2019-05-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(226,12,1,0,2,'2019-04-21 00:00:00','2019-04-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(227,12,1,0,0,'2019-10-23 00:00:00','2019-11-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(228,12,1,0,2,'2019-10-16 00:00:00','2019-10-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(229,12,1,0,0,'2019-08-14 00:00:00','2019-08-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(230,12,1,0,2,'2019-08-09 00:00:00','2019-08-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(231,12,1,0,0,'2019-11-19 00:00:00','2019-11-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(232,12,1,0,2,'2019-11-17 00:00:00','2019-11-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(233,12,1,0,0,'2019-09-16 00:00:00','2019-09-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(234,12,1,0,2,'2019-09-14 00:00:00','2019-09-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(235,12,1,0,0,'2020-01-02 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(236,12,1,0,2,'2019-12-30 00:00:00','2020-01-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(237,12,1,0,0,'2019-05-24 00:00:00','2019-05-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(238,12,1,0,2,'2019-05-21 00:00:00','2019-05-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(239,12,1,0,0,'2019-11-14 00:00:00','2019-11-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(240,12,1,0,2,'2019-11-08 00:00:00','2019-11-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(241,12,1,0,2,'2019-02-12 00:00:00','2019-02-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(242,12,1,0,2,'2019-02-11 00:00:00','2019-02-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(243,12,1,0,2,'2019-05-07 00:00:00','2019-05-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(244,12,1,0,2,'2019-05-04 00:00:00','2019-05-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(245,12,1,0,0,'2019-05-17 00:00:00','2019-05-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(246,12,1,0,2,'2019-05-14 00:00:00','2019-05-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(247,13,1,0,0,'2019-09-06 00:00:00','2019-09-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(248,13,1,0,2,'2019-09-04 00:00:00','2019-09-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(249,13,1,0,0,'2019-01-29 00:00:00','2019-02-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(250,13,1,0,2,'2019-01-27 00:00:00','2019-01-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(251,13,1,0,0,'2019-10-09 00:00:00','2019-10-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(252,13,1,0,2,'2019-10-07 00:00:00','2019-10-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(253,13,1,0,0,'2019-03-03 00:00:00','2019-03-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(254,13,1,0,2,'2019-02-26 00:00:00','2019-03-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(255,13,1,0,0,'2019-03-30 00:00:00','2019-04-01 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(256,13,1,0,2,'2019-03-25 00:00:00','2019-03-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(257,13,1,0,0,'2019-04-14 00:00:00','2019-04-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(258,13,1,0,2,'2019-04-09 00:00:00','2019-04-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(259,13,1,0,0,'2019-10-24 00:00:00','2019-11-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(260,13,1,0,2,'2019-10-23 00:00:00','2019-10-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(261,13,1,0,0,'2019-01-19 00:00:00','2019-01-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(262,13,1,0,2,'2019-01-12 00:00:00','2019-01-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(263,13,1,0,0,'2019-09-24 00:00:00','2019-09-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(264,13,1,0,2,'2019-09-21 00:00:00','2019-09-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(265,13,1,0,0,'2019-02-22 00:00:00','2019-02-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(266,13,1,0,2,'2019-02-20 00:00:00','2019-02-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(267,13,1,0,0,'2019-08-13 00:00:00','2019-08-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(268,13,1,0,2,'2019-08-06 00:00:00','2019-08-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(269,13,1,0,0,'2019-05-27 00:00:00','2019-06-01 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(270,13,1,0,2,'2019-05-21 00:00:00','2019-05-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(271,13,1,0,0,'2019-05-05 00:00:00','2019-05-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(272,13,1,0,2,'2019-05-04 00:00:00','2019-05-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(273,13,1,0,0,'2019-04-02 00:00:00','2019-04-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(274,13,1,0,2,'2019-04-01 00:00:00','2019-04-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(275,13,1,0,0,'2019-02-07 00:00:00','2019-02-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(276,13,1,0,2,'2019-02-02 00:00:00','2019-02-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(277,13,1,0,0,'2019-12-28 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(278,13,1,0,2,'2019-12-24 00:00:00','2019-12-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(279,13,1,0,0,'2019-05-16 00:00:00','2019-05-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(280,13,1,0,2,'2019-05-09 00:00:00','2019-05-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(281,13,1,0,0,'2019-03-15 00:00:00','2019-03-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(282,13,1,0,2,'2019-03-09 00:00:00','2019-03-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(283,13,1,0,0,'2019-12-07 00:00:00','2019-12-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(284,13,1,0,2,'2019-12-03 00:00:00','2019-12-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(285,13,1,0,0,'2019-07-19 00:00:00','2019-07-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(286,13,1,0,2,'2019-07-16 00:00:00','2019-07-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(287,13,1,0,0,'2019-05-03 00:00:00','2019-05-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(288,13,1,0,2,'2019-04-29 00:00:00','2019-05-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(289,13,1,0,0,'2019-06-16 00:00:00','2019-06-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(290,13,1,0,2,'2019-06-10 00:00:00','2019-06-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(291,13,1,0,0,'2019-08-25 00:00:00','2019-09-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(292,13,1,0,2,'2019-08-21 00:00:00','2019-08-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(293,13,1,0,0,'2019-11-18 00:00:00','2019-12-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(294,13,1,0,2,'2019-11-13 00:00:00','2019-11-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(295,13,1,0,0,'2019-10-19 00:00:00','2019-10-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(296,13,1,0,2,'2019-10-15 00:00:00','2019-10-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(297,13,1,0,0,'2019-03-23 00:00:00','2019-03-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(298,13,1,0,2,'2019-03-19 00:00:00','2019-03-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(299,13,1,0,0,'2019-02-09 00:00:00','2019-02-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(300,13,1,0,2,'2019-02-08 00:00:00','2019-02-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(301,13,1,0,0,'2019-04-05 00:00:00','2019-04-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(302,13,1,0,2,'2019-04-04 00:00:00','2019-04-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(303,13,1,0,0,'2019-07-12 00:00:00','2019-07-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(304,13,1,0,2,'2019-07-07 00:00:00','2019-07-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(305,13,1,0,0,'2019-10-04 00:00:00','2019-10-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(306,13,1,0,2,'2019-09-28 00:00:00','2019-10-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(307,13,1,0,0,'2019-08-03 00:00:00','2019-08-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(308,13,1,0,2,'2019-07-27 00:00:00','2019-08-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(309,13,1,0,2,'2019-04-11 00:00:00','2019-04-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(310,13,1,0,2,'2019-04-10 00:00:00','2019-04-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(311,13,1,0,0,'2019-06-02 00:00:00','2019-06-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(312,13,1,0,2,'2019-06-01 00:00:00','2019-06-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(313,13,1,0,0,'2019-07-03 00:00:00','2019-07-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(314,13,1,0,2,'2019-06-29 00:00:00','2019-07-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(315,13,1,0,0,'2019-09-14 00:00:00','2019-09-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(316,13,1,0,2,'2019-09-07 00:00:00','2019-09-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(317,14,1,0,0,'2019-10-18 00:00:00','2019-10-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(318,14,1,0,2,'2019-10-14 00:00:00','2019-10-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(319,14,1,0,0,'2019-12-02 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(320,14,1,0,2,'2019-11-26 00:00:00','2019-12-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(321,14,1,0,0,'2019-10-28 00:00:00','2019-10-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(322,14,1,0,2,'2019-10-25 00:00:00','2019-10-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(323,14,1,0,0,'2019-06-23 00:00:00','2019-07-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(324,14,1,0,2,'2019-06-19 00:00:00','2019-06-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(325,14,1,0,0,'2019-10-07 00:00:00','2019-10-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(326,14,1,0,2,'2019-09-30 00:00:00','2019-10-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(327,14,1,0,0,'2019-09-02 00:00:00','2019-09-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(328,14,1,0,2,'2019-08-27 00:00:00','2019-09-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(329,14,1,0,0,'2019-05-29 00:00:00','2019-05-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(330,14,1,0,2,'2019-05-26 00:00:00','2019-05-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(331,14,1,0,0,'2019-04-27 00:00:00','2019-05-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(332,14,1,0,2,'2019-04-21 00:00:00','2019-04-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(333,14,1,0,0,'2019-09-10 00:00:00','2019-09-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(334,14,1,0,2,'2019-09-03 00:00:00','2019-09-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(335,14,1,0,0,'2019-09-28 00:00:00','2019-09-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(336,14,1,0,2,'2019-09-26 00:00:00','2019-09-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(337,14,1,0,0,'2019-06-10 00:00:00','2019-06-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(338,14,1,0,2,'2019-06-05 00:00:00','2019-06-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(339,14,1,0,0,'2019-03-14 00:00:00','2019-03-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(340,14,1,0,2,'2019-03-08 00:00:00','2019-03-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(341,14,1,0,0,'2019-10-24 00:00:00','2019-10-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(342,14,1,0,2,'2019-10-23 00:00:00','2019-10-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(343,14,1,0,0,'2019-07-29 00:00:00','2019-08-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(344,14,1,0,2,'2019-07-26 00:00:00','2019-07-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(345,14,1,0,0,'2019-08-17 00:00:00','2019-08-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(346,14,1,0,2,'2019-08-15 00:00:00','2019-08-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(347,14,1,0,0,'2019-06-04 00:00:00','2019-06-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(348,14,1,0,2,'2019-05-31 00:00:00','2019-06-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(349,14,1,0,0,'2019-01-08 00:00:00','2019-01-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(350,14,1,0,2,'2019-01-04 00:00:00','2019-01-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(351,14,1,0,0,'2019-02-02 00:00:00','2019-02-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(352,14,1,0,2,'2019-01-26 00:00:00','2019-02-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(353,14,1,0,0,'2019-03-20 00:00:00','2019-04-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(354,14,1,0,2,'2019-03-19 00:00:00','2019-03-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(355,14,1,0,0,'2019-01-23 00:00:00','2019-01-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(356,14,1,0,2,'2019-01-21 00:00:00','2019-01-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(357,14,1,0,0,'2019-01-20 00:00:00','2019-01-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(358,14,1,0,2,'2019-01-18 00:00:00','2019-01-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(359,14,1,0,0,'2019-05-24 00:00:00','2019-05-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(360,14,1,0,2,'2019-05-18 00:00:00','2019-05-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(361,14,1,0,0,'2019-11-03 00:00:00','2019-11-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(362,14,1,0,2,'2019-10-30 00:00:00','2019-11-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(363,14,1,0,0,'2019-11-20 00:00:00','2019-11-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(364,14,1,0,2,'2019-11-15 00:00:00','2019-11-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(365,14,1,0,0,'2019-02-10 00:00:00','2019-03-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(366,14,1,0,2,'2019-02-04 00:00:00','2019-02-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(367,15,1,0,0,'2019-11-07 00:00:00','2019-12-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(368,15,1,0,2,'2019-10-31 00:00:00','2019-11-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(369,15,1,0,0,'2019-06-03 00:00:00','2019-06-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(370,15,1,0,2,'2019-05-28 00:00:00','2019-06-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(371,15,1,0,0,'2019-05-21 00:00:00','2019-05-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(372,15,1,0,2,'2019-05-14 00:00:00','2019-05-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(373,15,1,0,0,'2019-08-15 00:00:00','2019-08-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(374,15,1,0,2,'2019-08-09 00:00:00','2019-08-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(375,15,1,0,0,'2019-09-28 00:00:00','2019-10-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(376,15,1,0,2,'2019-09-24 00:00:00','2019-09-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(377,15,1,0,0,'2019-04-06 00:00:00','2019-04-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(378,15,1,0,2,'2019-03-30 00:00:00','2019-04-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(379,15,1,0,0,'2019-02-16 00:00:00','2019-02-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(380,15,1,0,2,'2019-02-12 00:00:00','2019-02-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(381,15,1,0,0,'2019-10-11 00:00:00','2019-10-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(382,15,1,0,2,'2019-10-10 00:00:00','2019-10-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(383,15,1,0,0,'2019-08-29 00:00:00','2019-08-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(384,15,1,0,2,'2019-08-25 00:00:00','2019-08-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(385,15,1,0,0,'2019-07-20 00:00:00','2019-07-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(386,15,1,0,2,'2019-07-14 00:00:00','2019-07-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(387,15,1,0,0,'2019-12-11 00:00:00','2019-12-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(388,15,1,0,2,'2019-12-10 00:00:00','2019-12-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(389,15,1,0,0,'2019-09-13 00:00:00','2019-09-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(390,15,1,0,2,'2019-09-06 00:00:00','2019-09-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(391,15,1,0,0,'2019-02-23 00:00:00','2019-03-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(392,15,1,0,2,'2019-02-17 00:00:00','2019-02-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(393,15,1,0,0,'2019-06-22 00:00:00','2019-07-01 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(394,15,1,0,2,'2019-06-17 00:00:00','2019-06-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(395,15,1,0,2,'2019-07-19 00:00:00','2019-07-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(396,15,1,0,2,'2019-07-15 00:00:00','2019-07-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(397,15,1,0,0,'2019-09-01 00:00:00','2019-09-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(398,15,1,0,2,'2019-08-30 00:00:00','2019-09-01 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(399,15,1,0,0,'2019-05-09 00:00:00','2019-05-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(400,15,1,0,2,'2019-05-06 00:00:00','2019-05-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(401,15,1,0,0,'2019-10-28 00:00:00','2019-10-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(402,15,1,0,2,'2019-10-23 00:00:00','2019-10-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(403,15,1,0,0,'2019-08-08 00:00:00','2019-08-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(404,15,1,0,2,'2019-08-04 00:00:00','2019-08-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(405,15,1,0,0,'2019-04-24 00:00:00','2019-05-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(406,15,1,0,2,'2019-04-23 00:00:00','2019-04-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(407,15,1,0,0,'2019-01-10 00:00:00','2019-02-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(408,15,1,0,2,'2019-01-05 00:00:00','2019-01-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(409,15,1,0,0,'2019-02-03 00:00:00','2019-02-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(410,15,1,0,2,'2019-02-02 00:00:00','2019-02-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(411,15,1,0,0,'2019-07-07 00:00:00','2019-07-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(412,15,1,0,2,'2019-07-01 00:00:00','2019-07-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(413,15,1,0,0,'2019-08-23 00:00:00','2019-08-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(414,15,1,0,2,'2019-08-21 00:00:00','2019-08-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(415,15,1,0,0,'2019-02-10 00:00:00','2019-02-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(416,15,1,0,2,'2019-02-08 00:00:00','2019-02-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(417,15,1,0,0,'2019-08-02 00:00:00','2019-08-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(418,15,1,0,2,'2019-07-28 00:00:00','2019-08-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(419,15,1,0,0,'2019-12-23 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(420,15,1,0,2,'2019-12-16 00:00:00','2019-12-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(421,15,1,0,0,'2019-06-10 00:00:00','2019-06-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(422,15,1,0,2,'2019-06-09 00:00:00','2019-06-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(423,16,1,0,0,'2019-12-31 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(424,16,1,0,2,'2019-12-26 00:00:00','2019-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(425,16,1,0,0,'2019-09-30 00:00:00','2019-10-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(426,16,1,0,2,'2019-09-24 00:00:00','2019-09-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(427,16,1,0,0,'2019-05-15 00:00:00','2019-05-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(428,16,1,0,2,'2019-05-10 00:00:00','2019-05-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(429,16,1,0,0,'2019-02-03 00:00:00','2019-02-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(430,16,1,0,2,'2019-01-27 00:00:00','2019-02-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(431,16,1,0,0,'2019-04-23 00:00:00','2019-05-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(432,16,1,0,2,'2019-04-17 00:00:00','2019-04-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(433,16,1,0,0,'2019-11-21 00:00:00','2019-11-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(434,16,1,0,2,'2019-11-20 00:00:00','2019-11-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(435,16,1,0,0,'2019-02-20 00:00:00','2019-02-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(436,16,1,0,2,'2019-02-17 00:00:00','2019-02-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(437,16,1,0,0,'2019-07-07 00:00:00','2019-07-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(438,16,1,0,2,'2019-07-06 00:00:00','2019-07-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(439,16,1,0,0,'2019-07-03 00:00:00','2019-07-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(440,16,1,0,2,'2019-07-02 00:00:00','2019-07-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(441,16,1,0,0,'2019-12-07 00:00:00','2019-12-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(442,16,1,0,2,'2019-12-06 00:00:00','2019-12-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(443,16,1,0,0,'2019-07-26 00:00:00','2019-08-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(444,16,1,0,2,'2019-07-23 00:00:00','2019-07-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(445,16,1,0,0,'2019-05-03 00:00:00','2019-05-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(446,16,1,0,2,'2019-05-02 00:00:00','2019-05-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(447,16,1,0,0,'2019-07-10 00:00:00','2019-07-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(448,16,1,0,2,'2019-07-09 00:00:00','2019-07-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(449,16,1,0,0,'2019-07-18 00:00:00','2019-07-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(450,16,1,0,2,'2019-07-13 00:00:00','2019-07-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(451,16,1,0,0,'2019-05-31 00:00:00','2019-06-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(452,16,1,0,2,'2019-05-29 00:00:00','2019-05-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(453,16,1,0,0,'2019-08-31 00:00:00','2019-09-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(454,16,1,0,2,'2019-08-29 00:00:00','2019-08-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(455,16,1,0,0,'2019-10-08 00:00:00','2019-10-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(456,16,1,0,2,'2019-10-06 00:00:00','2019-10-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(457,16,1,0,0,'2019-08-14 00:00:00','2019-08-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(458,16,1,0,2,'2019-08-13 00:00:00','2019-08-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(459,16,1,0,0,'2019-06-19 00:00:00','2019-06-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(460,16,1,0,2,'2019-06-17 00:00:00','2019-06-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(461,16,1,0,0,'2019-05-24 00:00:00','2019-05-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(462,16,1,0,2,'2019-05-23 00:00:00','2019-05-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(463,16,1,0,2,'2019-09-29 00:00:00','2019-09-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(464,16,1,0,2,'2019-09-27 00:00:00','2019-09-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(465,16,1,0,0,'2019-01-12 00:00:00','2019-01-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(466,16,1,0,2,'2019-01-06 00:00:00','2019-01-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(467,16,1,0,0,'2019-10-16 00:00:00','2019-11-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(468,16,1,0,2,'2019-10-12 00:00:00','2019-10-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(469,16,1,0,0,'2019-06-08 00:00:00','2019-06-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(470,16,1,0,2,'2019-06-07 00:00:00','2019-06-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(471,16,1,0,0,'2019-06-25 00:00:00','2019-06-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(472,16,1,0,2,'2019-06-22 00:00:00','2019-06-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(473,16,1,0,0,'2019-03-08 00:00:00','2019-04-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(474,16,1,0,2,'2019-03-02 00:00:00','2019-03-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(475,16,1,0,0,'2019-01-25 00:00:00','2019-01-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(476,16,1,0,2,'2019-01-19 00:00:00','2019-01-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(477,16,1,0,0,'2019-11-13 00:00:00','2019-11-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(478,16,1,0,2,'2019-11-11 00:00:00','2019-11-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(479,16,1,0,0,'2019-04-12 00:00:00','2019-04-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(480,16,1,0,2,'2019-04-09 00:00:00','2019-04-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(481,16,1,0,0,'2019-06-30 00:00:00','2019-07-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(482,16,1,0,2,'2019-06-29 00:00:00','2019-06-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(483,16,1,0,0,'2019-12-13 00:00:00','2019-12-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(484,16,1,0,2,'2019-12-08 00:00:00','2019-12-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(485,16,1,0,0,'2019-06-16 00:00:00','2019-06-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(486,16,1,0,2,'2019-06-10 00:00:00','2019-06-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(487,16,1,0,0,'2019-02-27 00:00:00','2019-03-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(488,16,1,0,2,'2019-02-23 00:00:00','2019-02-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(489,16,1,0,0,'2019-02-10 00:00:00','2019-02-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(490,16,1,0,2,'2019-02-05 00:00:00','2019-02-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(491,16,1,0,0,'2019-08-04 00:00:00','2019-08-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(492,16,1,0,2,'2019-08-02 00:00:00','2019-08-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(493,16,1,0,0,'2019-11-07 00:00:00','2019-11-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(494,16,1,0,2,'2019-11-05 00:00:00','2019-11-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(495,16,1,0,0,'2019-11-26 00:00:00','2019-12-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(496,16,1,0,2,'2019-11-24 00:00:00','2019-11-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(497,16,1,0,2,'2019-04-22 00:00:00','2019-04-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(498,16,1,0,2,'2019-04-19 00:00:00','2019-04-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(499,17,1,0,0,'2019-03-10 00:00:00','2019-03-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(500,17,1,0,2,'2019-03-08 00:00:00','2019-03-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(501,17,1,0,0,'2019-08-25 00:00:00','2019-08-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(502,17,1,0,2,'2019-08-21 00:00:00','2019-08-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(503,17,1,0,0,'2019-08-10 00:00:00','2019-08-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(504,17,1,0,2,'2019-08-07 00:00:00','2019-08-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(505,17,1,0,0,'2019-05-05 00:00:00','2019-05-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(506,17,1,0,2,'2019-05-03 00:00:00','2019-05-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(507,17,1,0,0,'2019-12-04 00:00:00','2019-12-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(508,17,1,0,2,'2019-12-03 00:00:00','2019-12-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(509,17,1,0,0,'2019-09-12 00:00:00','2019-10-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(510,17,1,0,2,'2019-09-08 00:00:00','2019-09-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(511,17,1,0,0,'2019-12-16 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(512,17,1,0,2,'2019-12-12 00:00:00','2019-12-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(513,17,1,0,0,'2019-10-05 00:00:00','2019-10-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(514,17,1,0,2,'2019-10-04 00:00:00','2019-10-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(515,17,1,0,0,'2019-05-28 00:00:00','2019-06-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(516,17,1,0,2,'2019-05-26 00:00:00','2019-05-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(517,17,1,0,0,'2019-07-13 00:00:00','2019-08-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(518,17,1,0,2,'2019-07-12 00:00:00','2019-07-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(519,17,1,0,0,'2019-06-22 00:00:00','2019-06-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(520,17,1,0,2,'2019-06-21 00:00:00','2019-06-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(521,17,1,0,0,'2019-06-19 00:00:00','2019-06-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(522,17,1,0,2,'2019-06-15 00:00:00','2019-06-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(523,17,1,0,0,'2019-01-19 00:00:00','2019-02-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(524,17,1,0,2,'2019-01-14 00:00:00','2019-01-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(525,17,1,0,0,'2019-11-11 00:00:00','2019-12-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(526,17,1,0,2,'2019-11-09 00:00:00','2019-11-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(527,17,1,0,0,'2019-08-31 00:00:00','2019-09-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(528,17,1,0,2,'2019-08-30 00:00:00','2019-08-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(529,17,1,0,0,'2019-06-28 00:00:00','2019-07-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(530,17,1,0,2,'2019-06-24 00:00:00','2019-06-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(531,17,1,0,0,'2019-05-14 00:00:00','2019-05-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(532,17,1,0,2,'2019-05-07 00:00:00','2019-05-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(533,17,1,0,0,'2019-09-05 00:00:00','2019-09-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(534,17,1,0,2,'2019-09-02 00:00:00','2019-09-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(535,17,1,0,0,'2019-03-04 00:00:00','2019-03-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(536,17,1,0,2,'2019-02-28 00:00:00','2019-03-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(537,17,1,0,0,'2019-10-08 00:00:00','2019-10-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(538,17,1,0,2,'2019-10-06 00:00:00','2019-10-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(539,17,1,0,0,'2019-04-16 00:00:00','2019-05-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(540,17,1,0,2,'2019-04-13 00:00:00','2019-04-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(541,17,1,0,0,'2019-05-24 00:00:00','2019-05-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(542,17,1,0,2,'2019-05-23 00:00:00','2019-05-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(543,17,1,0,0,'2019-02-09 00:00:00','2019-02-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(544,17,1,0,2,'2019-02-07 00:00:00','2019-02-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(545,17,1,0,0,'2019-06-09 00:00:00','2019-06-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(546,17,1,0,2,'2019-06-05 00:00:00','2019-06-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(547,17,1,0,0,'2019-10-24 00:00:00','2019-11-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(548,17,1,0,2,'2019-10-21 00:00:00','2019-10-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(549,17,1,0,0,'2019-04-03 00:00:00','2019-04-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(550,17,1,0,2,'2019-04-02 00:00:00','2019-04-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(551,17,1,0,0,'2019-10-19 00:00:00','2019-10-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(552,17,1,0,2,'2019-10-15 00:00:00','2019-10-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(553,17,1,0,0,'2019-03-13 00:00:00','2019-04-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(554,17,1,0,2,'2019-03-12 00:00:00','2019-03-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(555,18,1,0,0,'2019-04-22 00:00:00','2019-05-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(556,18,1,0,2,'2019-04-17 00:00:00','2019-04-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(557,18,1,0,0,'2019-03-29 00:00:00','2019-04-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(558,18,1,0,2,'2019-03-25 00:00:00','2019-03-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(559,18,1,0,0,'2019-11-17 00:00:00','2019-11-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(560,18,1,0,2,'2019-11-10 00:00:00','2019-11-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(561,18,1,0,0,'2019-01-14 00:00:00','2019-01-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(562,18,1,0,2,'2019-01-13 00:00:00','2019-01-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(563,18,1,0,0,'2019-06-11 00:00:00','2019-06-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(564,18,1,0,2,'2019-06-04 00:00:00','2019-06-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(565,18,1,0,0,'2019-04-06 00:00:00','2019-04-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(566,18,1,0,2,'2019-04-04 00:00:00','2019-04-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(567,18,1,0,0,'2019-02-26 00:00:00','2019-03-01 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(568,18,1,0,2,'2019-02-22 00:00:00','2019-02-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(569,18,1,0,0,'2019-10-26 00:00:00','2019-11-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(570,18,1,0,2,'2019-10-22 00:00:00','2019-10-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(571,18,1,0,0,'2019-07-31 00:00:00','2019-08-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(572,18,1,0,2,'2019-07-28 00:00:00','2019-07-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(573,18,1,0,0,'2019-07-20 00:00:00','2019-07-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(574,18,1,0,2,'2019-07-18 00:00:00','2019-07-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(575,18,1,0,0,'2019-04-14 00:00:00','2019-04-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(576,18,1,0,2,'2019-04-07 00:00:00','2019-04-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(577,18,1,0,0,'2019-01-22 00:00:00','2019-01-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(578,18,1,0,2,'2019-01-19 00:00:00','2019-01-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(579,18,1,0,0,'2019-08-30 00:00:00','2019-09-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(580,18,1,0,2,'2019-08-29 00:00:00','2019-08-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(581,18,1,0,0,'2019-12-22 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(582,18,1,0,2,'2019-12-20 00:00:00','2019-12-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(583,18,1,0,0,'2019-06-02 00:00:00','2019-06-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(584,18,1,0,2,'2019-06-01 00:00:00','2019-06-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(585,18,1,0,0,'2019-08-05 00:00:00','2019-08-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(586,18,1,0,2,'2019-08-03 00:00:00','2019-08-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(587,18,1,0,0,'2019-09-29 00:00:00','2019-10-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(588,18,1,0,2,'2019-09-23 00:00:00','2019-09-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(589,18,1,0,0,'2019-09-10 00:00:00','2019-09-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(590,18,1,0,2,'2019-09-05 00:00:00','2019-09-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(591,18,1,0,0,'2019-10-21 00:00:00','2019-10-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(592,18,1,0,2,'2019-10-18 00:00:00','2019-10-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(593,18,1,0,0,'2019-03-24 00:00:00','2019-03-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(594,18,1,0,2,'2019-03-19 00:00:00','2019-03-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(595,18,1,0,0,'2019-03-18 00:00:00','2019-03-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(596,18,1,0,2,'2019-03-11 00:00:00','2019-03-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(597,18,1,0,0,'2019-06-19 00:00:00','2019-06-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(598,18,1,0,2,'2019-06-15 00:00:00','2019-06-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(599,18,1,0,0,'2019-08-20 00:00:00','2019-08-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(600,18,1,0,2,'2019-08-19 00:00:00','2019-08-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(601,18,1,0,0,'2019-11-23 00:00:00','2019-12-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(602,18,1,0,2,'2019-11-22 00:00:00','2019-11-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(603,18,1,0,0,'2019-03-07 00:00:00','2019-03-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(604,18,1,0,2,'2019-03-01 00:00:00','2019-03-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(605,18,1,0,0,'2019-07-16 00:00:00','2019-07-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(606,18,1,0,2,'2019-07-13 00:00:00','2019-07-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(607,18,1,0,0,'2019-02-04 00:00:00','2019-02-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(608,18,1,0,2,'2019-01-29 00:00:00','2019-02-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(609,18,1,0,0,'2019-05-25 00:00:00','2019-06-01 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(610,18,1,0,2,'2019-05-22 00:00:00','2019-05-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(611,18,1,0,0,'2019-01-09 00:00:00','2019-01-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(612,18,1,0,2,'2019-01-04 00:00:00','2019-01-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(613,18,1,0,0,'2019-06-26 00:00:00','2019-07-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(614,18,1,0,2,'2019-06-21 00:00:00','2019-06-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(615,18,1,0,0,'2019-05-08 00:00:00','2019-05-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(616,18,1,0,2,'2019-05-03 00:00:00','2019-05-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(617,18,1,0,0,'2019-09-17 00:00:00','2019-09-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(618,18,1,0,2,'2019-09-11 00:00:00','2019-09-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(619,18,1,0,0,'2019-12-17 00:00:00','2019-12-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(620,18,1,0,2,'2019-12-12 00:00:00','2019-12-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(621,18,1,0,2,'2019-03-16 00:00:00','2019-03-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(622,18,1,0,2,'2019-03-15 00:00:00','2019-03-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(623,18,1,0,0,'2019-05-12 00:00:00','2019-05-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(624,18,1,0,2,'2019-05-10 00:00:00','2019-05-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(625,18,1,0,0,'2019-08-10 00:00:00','2019-08-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(626,18,1,0,2,'2019-08-06 00:00:00','2019-08-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(627,19,1,0,0,'2019-05-27 00:00:00','2019-06-01 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(628,19,1,0,2,'2019-05-20 00:00:00','2019-05-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(629,19,1,0,0,'2019-12-27 00:00:00','2019-12-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(630,19,1,0,2,'2019-12-23 00:00:00','2019-12-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(631,19,1,0,0,'2019-04-27 00:00:00','2019-05-01 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(632,19,1,0,2,'2019-04-26 00:00:00','2019-04-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(633,19,1,0,0,'2019-10-10 00:00:00','2019-10-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(634,19,1,0,2,'2019-10-09 00:00:00','2019-10-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(635,19,1,0,0,'2019-05-04 00:00:00','2019-05-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(636,19,1,0,2,'2019-05-01 00:00:00','2019-05-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(637,19,1,0,0,'2019-05-13 00:00:00','2019-05-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(638,19,1,0,2,'2019-05-12 00:00:00','2019-05-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(639,19,1,0,0,'2019-10-27 00:00:00','2019-10-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(640,19,1,0,2,'2019-10-25 00:00:00','2019-10-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(641,19,1,0,0,'2019-07-29 00:00:00','2019-08-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(642,19,1,0,2,'2019-07-24 00:00:00','2019-07-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(643,19,1,0,0,'2019-08-22 00:00:00','2019-08-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(644,19,1,0,2,'2019-08-19 00:00:00','2019-08-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(645,19,1,0,0,'2019-02-05 00:00:00','2019-03-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(646,19,1,0,2,'2019-01-31 00:00:00','2019-02-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(647,19,1,0,0,'2019-10-30 00:00:00','2019-12-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(648,19,1,0,2,'2019-10-28 00:00:00','2019-10-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(649,19,1,0,0,'2019-08-14 00:00:00','2019-08-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(650,19,1,0,2,'2019-08-08 00:00:00','2019-08-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(651,19,1,0,0,'2019-10-04 00:00:00','2019-10-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(652,19,1,0,2,'2019-09-28 00:00:00','2019-10-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(653,19,1,0,0,'2019-04-02 00:00:00','2019-04-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(654,19,1,0,2,'2019-03-29 00:00:00','2019-04-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(655,19,1,0,0,'2019-12-09 00:00:00','2019-12-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(656,19,1,0,2,'2019-12-05 00:00:00','2019-12-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(657,19,1,0,0,'2019-09-09 00:00:00','2019-09-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(658,19,1,0,2,'2019-09-03 00:00:00','2019-09-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(659,19,1,0,0,'2019-07-04 00:00:00','2019-07-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(660,19,1,0,2,'2019-06-29 00:00:00','2019-07-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(661,19,1,0,0,'2019-06-22 00:00:00','2019-06-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(662,19,1,0,2,'2019-06-16 00:00:00','2019-06-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(663,19,1,0,0,'2019-04-15 00:00:00','2019-04-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(664,19,1,0,2,'2019-04-11 00:00:00','2019-04-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(665,19,1,0,0,'2019-06-08 00:00:00','2019-06-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(666,19,1,0,2,'2019-06-05 00:00:00','2019-06-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(667,19,1,0,0,'2019-01-26 00:00:00','2019-01-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(668,19,1,0,2,'2019-01-20 00:00:00','2019-01-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(669,19,1,0,0,'2019-06-03 00:00:00','2019-06-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(670,19,1,0,2,'2019-06-01 00:00:00','2019-06-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(671,19,1,0,0,'2019-10-19 00:00:00','2019-10-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(672,19,1,0,2,'2019-10-12 00:00:00','2019-10-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(673,19,1,0,0,'2019-12-30 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(674,19,1,0,2,'2019-12-29 00:00:00','2019-12-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(675,19,1,0,0,'2019-04-25 00:00:00','2019-04-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(676,19,1,0,2,'2019-04-22 00:00:00','2019-04-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(677,19,1,0,0,'2019-09-02 00:00:00','2019-09-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(678,19,1,0,2,'2019-08-31 00:00:00','2019-09-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(679,19,1,0,0,'2019-12-22 00:00:00','2019-12-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(680,19,1,0,2,'2019-12-16 00:00:00','2019-12-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(681,19,1,0,0,'2019-09-23 00:00:00','2019-09-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(682,19,1,0,2,'2019-09-19 00:00:00','2019-09-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(683,19,1,0,2,'2019-05-25 00:00:00','2019-05-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(684,19,1,0,2,'2019-05-21 00:00:00','2019-05-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(685,19,1,0,0,'2019-09-14 00:00:00','2019-09-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(686,19,1,0,2,'2019-09-10 00:00:00','2019-09-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(687,19,1,0,0,'2019-12-13 00:00:00','2019-12-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(688,19,1,0,2,'2019-12-11 00:00:00','2019-12-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(689,20,1,0,0,'2019-12-16 00:00:00','2019-12-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(690,20,1,0,2,'2019-12-13 00:00:00','2019-12-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(691,20,1,0,0,'2019-09-05 00:00:00','2019-09-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(692,20,1,0,2,'2019-09-03 00:00:00','2019-09-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(693,20,1,0,0,'2019-08-28 00:00:00','2019-09-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(694,20,1,0,2,'2019-08-21 00:00:00','2019-08-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(695,20,1,0,0,'2019-02-12 00:00:00','2019-02-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(696,20,1,0,2,'2019-02-11 00:00:00','2019-02-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(697,20,1,0,0,'2019-07-30 00:00:00','2019-08-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(698,20,1,0,2,'2019-07-26 00:00:00','2019-07-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(699,20,1,0,0,'2019-06-13 00:00:00','2019-06-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(700,20,1,0,2,'2019-06-06 00:00:00','2019-06-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(701,20,1,0,0,'2019-10-15 00:00:00','2019-10-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(702,20,1,0,2,'2019-10-14 00:00:00','2019-10-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(703,20,1,0,0,'2019-05-07 00:00:00','2019-05-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(704,20,1,0,2,'2019-05-04 00:00:00','2019-05-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(705,20,1,0,0,'2019-07-17 00:00:00','2019-07-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(706,20,1,0,2,'2019-07-14 00:00:00','2019-07-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(707,20,1,0,0,'2019-03-02 00:00:00','2019-03-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(708,20,1,0,2,'2019-02-23 00:00:00','2019-02-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(709,20,1,0,0,'2019-07-01 00:00:00','2019-07-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(710,20,1,0,2,'2019-06-28 00:00:00','2019-07-01 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(711,20,1,0,0,'2019-05-17 00:00:00','2019-05-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(712,20,1,0,2,'2019-05-15 00:00:00','2019-05-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(713,20,1,0,0,'2019-12-22 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(714,20,1,0,2,'2019-12-19 00:00:00','2019-12-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(715,20,1,0,0,'2019-11-30 00:00:00','2019-12-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(716,20,1,0,2,'2019-11-24 00:00:00','2019-11-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(717,20,1,0,0,'2019-05-14 00:00:00','2019-05-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(718,20,1,0,2,'2019-05-10 00:00:00','2019-05-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(719,20,1,0,0,'2019-04-09 00:00:00','2019-04-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(720,20,1,0,2,'2019-04-05 00:00:00','2019-04-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(721,20,1,0,0,'2019-04-04 00:00:00','2019-04-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(722,20,1,0,2,'2019-03-28 00:00:00','2019-04-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(723,20,1,0,2,'2019-02-26 00:00:00','2019-03-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(724,20,1,0,2,'2019-02-25 00:00:00','2019-02-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(725,20,1,0,0,'2019-06-22 00:00:00','2019-06-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(726,20,1,0,2,'2019-06-17 00:00:00','2019-06-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(727,20,1,0,0,'2019-12-05 00:00:00','2019-12-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(728,20,1,0,2,'2019-12-02 00:00:00','2019-12-05 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(729,20,1,0,0,'2019-11-02 00:00:00','2019-11-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(730,20,1,0,2,'2019-10-31 00:00:00','2019-11-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(731,20,1,0,0,'2019-03-09 00:00:00','2019-03-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(732,20,1,0,2,'2019-03-07 00:00:00','2019-03-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(733,20,1,0,0,'2019-12-12 00:00:00','2019-12-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(734,20,1,0,2,'2019-12-09 00:00:00','2019-12-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(735,20,1,0,0,'2019-05-30 00:00:00','2019-06-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(736,20,1,0,2,'2019-05-24 00:00:00','2019-05-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(737,20,1,0,0,'2019-04-27 00:00:00','2019-05-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(738,20,1,0,2,'2019-04-20 00:00:00','2019-04-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(739,20,1,0,0,'2019-07-08 00:00:00','2019-07-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(740,20,1,0,2,'2019-07-05 00:00:00','2019-07-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(741,20,1,0,0,'2019-08-11 00:00:00','2019-08-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(742,20,1,0,2,'2019-08-08 00:00:00','2019-08-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(743,20,1,0,0,'2019-07-22 00:00:00','2019-07-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(744,20,1,0,2,'2019-07-18 00:00:00','2019-07-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(745,20,1,0,0,'2019-09-23 00:00:00','2019-10-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(746,20,1,0,2,'2019-09-18 00:00:00','2019-09-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(747,21,1,0,0,'2019-03-25 00:00:00','2019-04-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(748,21,1,0,2,'2019-03-18 00:00:00','2019-03-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(749,21,1,0,0,'2019-09-14 00:00:00','2019-09-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(750,21,1,0,2,'2019-09-08 00:00:00','2019-09-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(751,21,1,0,0,'2019-01-22 00:00:00','2019-01-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(752,21,1,0,2,'2019-01-18 00:00:00','2019-01-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(753,21,1,0,0,'2019-10-02 00:00:00','2019-10-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(754,21,1,0,2,'2019-09-26 00:00:00','2019-10-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(755,21,1,0,0,'2019-06-20 00:00:00','2019-06-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(756,21,1,0,2,'2019-06-17 00:00:00','2019-06-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(757,21,1,0,0,'2019-05-30 00:00:00','2019-06-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(758,21,1,0,2,'2019-05-24 00:00:00','2019-05-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(759,21,1,0,0,'2019-05-03 00:00:00','2019-05-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(760,21,1,0,2,'2019-04-28 00:00:00','2019-05-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(761,21,1,0,0,'2019-03-07 00:00:00','2019-03-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(762,21,1,0,2,'2019-03-01 00:00:00','2019-03-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(763,21,1,0,0,'2019-12-16 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(764,21,1,0,2,'2019-12-13 00:00:00','2019-12-16 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(765,21,1,0,2,'2019-05-29 00:00:00','2019-05-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(766,21,1,0,2,'2019-05-28 00:00:00','2019-05-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(767,21,1,0,0,'2019-11-13 00:00:00','2019-12-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(768,21,1,0,2,'2019-11-09 00:00:00','2019-11-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(769,21,1,0,0,'2019-08-03 00:00:00','2019-08-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(770,21,1,0,2,'2019-07-31 00:00:00','2019-08-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(771,21,1,0,0,'2019-05-19 00:00:00','2019-05-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(772,21,1,0,2,'2019-05-15 00:00:00','2019-05-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(773,21,1,0,0,'2019-12-08 00:00:00','2019-12-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(774,21,1,0,2,'2019-12-07 00:00:00','2019-12-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(775,21,1,0,0,'2019-02-15 00:00:00','2019-02-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(776,21,1,0,2,'2019-02-08 00:00:00','2019-02-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(777,21,1,0,2,'2019-09-12 00:00:00','2019-09-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(778,21,1,0,2,'2019-09-10 00:00:00','2019-09-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(779,21,1,0,0,'2019-06-04 00:00:00','2019-06-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(780,21,1,0,2,'2019-06-02 00:00:00','2019-06-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(781,21,1,0,0,'2019-06-11 00:00:00','2019-06-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(782,21,1,0,2,'2019-06-07 00:00:00','2019-06-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(783,21,1,0,0,'2019-02-04 00:00:00','2019-02-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(784,21,1,0,2,'2019-01-28 00:00:00','2019-02-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(785,21,1,0,0,'2019-06-28 00:00:00','2019-07-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(786,21,1,0,2,'2019-06-25 00:00:00','2019-06-28 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(787,21,1,0,0,'2019-08-31 00:00:00','2019-09-08 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(788,21,1,0,2,'2019-08-24 00:00:00','2019-08-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(789,21,1,0,0,'2019-02-27 00:00:00','2019-03-01 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(790,21,1,0,2,'2019-02-21 00:00:00','2019-02-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(791,21,1,0,0,'2019-10-15 00:00:00','2019-11-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(792,21,1,0,2,'2019-10-11 00:00:00','2019-10-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(793,21,1,0,0,'2019-07-18 00:00:00','2019-07-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(794,21,1,0,2,'2019-07-11 00:00:00','2019-07-18 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(795,22,1,0,0,'2019-07-19 00:00:00','2019-07-21 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(796,22,1,0,2,'2019-07-17 00:00:00','2019-07-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(797,22,1,0,0,'2019-06-26 00:00:00','2019-07-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(798,22,1,0,2,'2019-06-22 00:00:00','2019-06-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(799,22,1,0,0,'2019-11-03 00:00:00','2019-11-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(800,22,1,0,2,'2019-10-27 00:00:00','2019-11-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(801,22,1,0,0,'2019-11-20 00:00:00','2019-12-02 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(802,22,1,0,2,'2019-11-13 00:00:00','2019-11-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(803,22,1,0,0,'2019-06-15 00:00:00','2019-06-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(804,22,1,0,2,'2019-06-14 00:00:00','2019-06-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(805,22,1,0,0,'2019-07-13 00:00:00','2019-07-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(806,22,1,0,2,'2019-07-10 00:00:00','2019-07-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(807,22,1,0,0,'2019-09-20 00:00:00','2019-10-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(808,22,1,0,2,'2019-09-15 00:00:00','2019-09-20 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(809,22,1,0,0,'2019-04-27 00:00:00','2019-05-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(810,22,1,0,2,'2019-04-26 00:00:00','2019-04-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(811,22,1,0,0,'2019-05-22 00:00:00','2019-06-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(812,22,1,0,2,'2019-05-19 00:00:00','2019-05-22 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(813,22,1,0,0,'2019-04-07 00:00:00','2019-04-12 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(814,22,1,0,2,'2019-04-06 00:00:00','2019-04-07 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(815,22,1,0,0,'2019-12-04 00:00:00','2019-12-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(816,22,1,0,2,'2019-12-02 00:00:00','2019-12-04 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(817,22,1,0,0,'2019-08-14 00:00:00','2019-08-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(818,22,1,0,2,'2019-08-10 00:00:00','2019-08-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(819,22,1,0,0,'2019-08-03 00:00:00','2019-08-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(820,22,1,0,2,'2019-07-27 00:00:00','2019-07-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(821,22,1,0,0,'2019-03-24 00:00:00','2019-03-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(822,22,1,0,2,'2019-03-23 00:00:00','2019-03-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(823,22,1,0,0,'2019-04-17 00:00:00','2019-04-26 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(824,22,1,0,2,'2019-04-12 00:00:00','2019-04-17 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(825,22,1,0,0,'2019-03-15 00:00:00','2019-03-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(826,22,1,0,2,'2019-03-13 00:00:00','2019-03-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(827,22,1,0,0,'2019-12-31 00:00:00','9999-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(828,22,1,0,2,'2019-12-30 00:00:00','2019-12-31 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(829,22,1,0,0,'2019-04-01 00:00:00','2019-04-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(830,22,1,0,2,'2019-03-26 00:00:00','2019-04-01 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(831,22,1,0,0,'2019-01-14 00:00:00','2019-01-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(832,22,1,0,2,'2019-01-09 00:00:00','2019-01-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(833,22,1,0,2,'2019-07-30 00:00:00','2019-08-03 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(834,22,1,0,2,'2019-07-29 00:00:00','2019-07-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(835,22,1,0,0,'2019-06-11 00:00:00','2019-06-14 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(836,22,1,0,2,'2019-06-09 00:00:00','2019-06-11 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(837,22,1,0,0,'2019-08-23 00:00:00','2019-08-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(838,22,1,0,2,'2019-08-19 00:00:00','2019-08-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(839,22,1,0,0,'2019-05-10 00:00:00','2019-05-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(840,22,1,0,2,'2019-05-04 00:00:00','2019-05-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(841,22,1,0,0,'2019-12-19 00:00:00','2019-12-23 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(842,22,1,0,2,'2019-12-13 00:00:00','2019-12-19 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(843,22,1,0,0,'2019-12-10 00:00:00','2019-12-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(844,22,1,0,2,'2019-12-09 00:00:00','2019-12-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(845,22,1,0,0,'2019-11-10 00:00:00','2019-11-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(846,22,1,0,2,'2019-11-06 00:00:00','2019-11-10 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(847,22,1,0,0,'2019-01-06 00:00:00','2019-01-09 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(848,22,1,0,2,'2019-01-04 00:00:00','2019-01-06 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(849,22,1,0,0,'2019-07-24 00:00:00','2019-07-27 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(850,22,1,0,2,'2019-07-21 00:00:00','2019-07-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(851,22,1,0,0,'2019-12-24 00:00:00','2019-12-30 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(852,22,1,0,2,'2019-12-23 00:00:00','2019-12-24 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(853,22,1,0,0,'2019-08-29 00:00:00','2019-09-15 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(854,22,1,0,2,'2019-08-24 00:00:00','2019-08-29 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(855,22,1,0,0,'2019-01-25 00:00:00','2019-03-13 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0),(856,22,1,0,2,'2019-01-23 00:00:00','2019-01-25 00:00:00','','',0,'2019-01-17 01:12:43',0,'2019-01-17 01:12:43',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableMarketRate`
--

LOCK TABLES `RentableMarketRate` WRITE;
/*!40000 ALTER TABLE `RentableMarketRate` DISABLE KEYS */;
INSERT INTO `RentableMarketRate` VALUES (1,1,0,1000.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,2,0,1500.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,3,0,1750.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,4,0,2500.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(5,5,0,65.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(6,6,0,75.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(7,7,0,85.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(8,8,0,35.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableTypeRef`
--

LOCK TABLES `RentableTypeRef` WRITE;
/*!40000 ALTER TABLE `RentableTypeRef` DISABLE KEYS */;
INSERT INTO `RentableTypeRef` VALUES (1,1,1,1,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,2,1,1,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,3,1,2,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,4,1,2,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(5,5,1,3,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(6,6,1,3,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(7,7,1,4,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(8,8,1,4,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(9,9,1,5,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(10,10,1,5,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(11,11,1,6,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(12,12,1,6,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(13,13,1,7,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(14,14,1,7,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(15,15,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(16,16,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(17,17,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(18,18,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(19,19,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(20,20,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(21,21,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(22,22,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableTypes`
--

LOCK TABLES `RentableTypes` WRITE;
/*!40000 ALTER TABLE `RentableTypes` DISABLE KEYS */;
INSERT INTO `RentableTypes` VALUES (1,1,'1BR1BA','Standard',6,4,4,12,45,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,1,'2BR1BA','Deluxe',6,4,4,12,46,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,1,'2BR2BA','Gold',6,4,4,12,47,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,1,'3BR2BA','Platinum',6,4,4,12,48,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(5,1,'H11','HotelStd',4,0,0,4,49,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(6,1,'H21','HotelGold',4,0,0,4,50,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(7,1,'H22','HotelPlatinum',4,0,0,4,51,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(8,1,'CP000','Car Port 000',6,4,4,6,52,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableUseStatus`
--

LOCK TABLES `RentableUseStatus` WRITE;
/*!40000 ALTER TABLE `RentableUseStatus` DISABLE KEYS */;
INSERT INTO `RentableUseStatus` VALUES (1,1,1,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,2,1,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,3,1,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,4,1,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(5,5,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(6,6,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(7,7,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(8,8,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(9,9,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(10,10,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(11,11,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(12,12,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(13,13,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(14,14,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(15,15,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(16,16,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(17,17,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(18,18,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(19,19,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(20,20,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(21,21,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(22,22,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(23,1,1,0,'2020-03-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(24,1,1,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(25,2,1,0,'2020-03-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(26,2,1,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(27,3,1,0,'2020-03-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(28,3,1,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(29,4,1,0,'2020-03-01 00:00:00','9999-12-31 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(30,4,1,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableUseType`
--

LOCK TABLES `RentableUseType` WRITE;
/*!40000 ALTER TABLE `RentableUseType` DISABLE KEYS */;
INSERT INTO `RentableUseType` VALUES (1,10,1,100,'2018-01-01 00:00:00','9999-12-31 00:00:00','','2019-05-16 08:22:30',211,'2019-05-16 08:22:30',211),(2,9,1,100,'2018-01-01 00:00:00','9999-12-31 00:00:00','','2019-05-16 08:23:00',211,'2019-05-16 08:23:00',211),(3,11,1,100,'2018-01-01 00:00:00','9999-12-31 00:00:00','','2019-05-16 08:23:24',211,'2019-05-16 08:23:24',211),(4,12,1,100,'2018-01-01 00:00:00','9999-12-31 00:00:00','','2019-05-16 08:23:37',211,'2019-05-16 08:23:37',211),(5,13,1,100,'2018-01-01 00:00:00','9999-12-31 00:00:00','','2019-05-16 08:23:54',211,'2019-05-16 08:23:54',211),(6,14,1,100,'2018-01-01 00:00:00','9999-12-31 00:00:00','','2019-05-16 08:24:08',211,'2019-05-16 08:24:08',211);
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableUsers`
--

LOCK TABLES `RentableUsers` WRITE;
/*!40000 ALTER TABLE `RentableUsers` DISABLE KEYS */;
INSERT INTO `RentableUsers` VALUES (1,1,1,1,'2019-01-03','2020-03-01','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,2,1,2,'2019-01-03','2020-03-01','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,3,1,3,'2019-01-03','2020-03-01','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,4,1,4,'2019-01-03','2020-03-01','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
  `TerminationStarted` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
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
INSERT INTO `RentalAgreement` VALUES (1,0,0,1,1,0,'2019-01-03 00:00:00','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-02-01',2,0,2,'',0,0,0.0000,'','0000-00-00','0000-00-00',0.0000,0.0000,235,'0000-00-00','','','','0000-00-00','','0000-00-00','','0000-00-00',0,52,285,'2019-01-03 00:00:00',202,'2019-01-03 00:00:00',0,196,'2019-01-03 00:00:00',0,160,'2019-01-03 00:00:00',286,'2019-01-03 00:00:00',0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00',0,'0000-00-00 00:00:00','1970-01-01 00:00:00',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,0,0,1,1,0,'2019-01-03 00:00:00','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-02-01',0,1,2,'',0,0,0.0000,'','0000-00-00','0000-00-00',0.0000,0.0000,78,'0000-00-00','','','','0000-00-00','','0000-00-00','','0000-00-00',0,52,141,'2019-01-03 00:00:00',143,'2019-01-03 00:00:00',0,63,'2019-01-03 00:00:00',0,57,'2019-01-03 00:00:00',38,'2019-01-03 00:00:00',0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00',0,'0000-00-00 00:00:00','1970-01-01 00:00:00',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,0,0,1,1,0,'2019-01-03 00:00:00','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-02-01',0,2,2,'',0,0,0.0000,'','0000-00-00','0000-00-00',0.0000,0.0000,281,'0000-00-00','','','','0000-00-00','','0000-00-00','','0000-00-00',0,52,108,'2019-01-03 00:00:00',279,'2019-01-03 00:00:00',0,274,'2019-01-03 00:00:00',0,58,'2019-01-03 00:00:00',147,'2019-01-03 00:00:00',0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00',0,'0000-00-00 00:00:00','1970-01-01 00:00:00',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,0,0,1,1,0,'2019-01-03 00:00:00','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-02-01',0,2,2,'',0,0,0.0000,'','0000-00-00','0000-00-00',0.0000,0.0000,92,'0000-00-00','','','','0000-00-00','','0000-00-00','','0000-00-00',0,52,99,'2019-01-03 00:00:00',218,'2019-01-03 00:00:00',0,158,'2019-01-03 00:00:00',0,141,'2019-01-03 00:00:00',30,'2019-01-03 00:00:00',0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00',0,'0000-00-00 00:00:00','1970-01-01 00:00:00',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementPayors`
--

LOCK TABLES `RentalAgreementPayors` WRITE;
/*!40000 ALTER TABLE `RentalAgreementPayors` DISABLE KEYS */;
INSERT INTO `RentalAgreementPayors` VALUES (1,1,1,1,'2019-01-03','2020-03-01',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,2,1,2,'2019-01-03','2020-03-01',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,3,1,3,'2019-01-03','2020-03-01',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,4,1,4,'2019-01-03','2020-03-01',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementRentables`
--

LOCK TABLES `RentalAgreementRentables` WRITE;
/*!40000 ALTER TABLE `RentalAgreementRentables` DISABLE KEYS */;
INSERT INTO `RentalAgreementRentables` VALUES (1,1,1,1,0,0,1000.0000,'2019-01-03','2020-03-01','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,2,1,2,0,0,1000.0000,'2019-01-03','2020-03-01','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,3,1,3,0,0,1500.0000,'2019-01-03','2020-03-01','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,4,1,4,0,0,1500.0000,'2019-01-03','2020-03-01','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
INSERT INTO `SLString` VALUES (1,1,1,'4Walls','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(2,1,1,'Apartment Finder Blue Book','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(3,1,1,'Apartment Guide','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(4,1,1,'Apartment Locator','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(5,1,1,'Apartment Map','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(6,1,1,'ApartmentFinder.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(7,1,1,'ApartmentGuide.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(8,1,1,'ApartmentGuyze.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(9,1,1,'ApartmentHomeLiving.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(10,1,1,'ApartmentLints.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(11,1,1,'ApartmentMag.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(12,1,1,'ApartmentMarketer.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(13,1,1,'ApartmentMatching.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(14,1,1,'ApartmentRatings.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(15,1,1,'ApartmentSearch.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(16,1,1,'ApartmentShowcase.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(17,1,1,'Apartments.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(18,1,1,'Apartments24-7.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(19,1,1,'ApartmentsNationwide.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(20,1,1,'ApartmentsPlus.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(21,1,1,'Brochure/Flyer','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(22,1,1,'CitySearch.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(23,1,1,'CollegeRentals.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(24,1,1,'CraigsList.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(25,1,1,'Current resident','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(26,1,1,'Direct Mail - Conventional','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(27,1,1,'Direct Mail - FullService','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(28,1,1,'Drive by','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(29,1,1,'EasyRent.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(30,1,1,'El Nacional','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(31,1,1,'EliteRenting.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(32,1,1,'For Rent Magazine','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(33,1,1,'ForRent.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(34,1,1,'Google Internet Program','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(35,1,1,'Google.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(36,1,1,'HotPads.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(37,1,1,'LivingChoices.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(38,1,1,'Local Line Rolloer','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(39,1,1,'Locator Service','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(40,1,1,'Move.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(41,1,1,'MoveForFree.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(42,1,1,'MyNewPlace.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(43,1,1,'Oklahoma Gazette','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(44,1,1,'Oodle.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(45,1,1,'Other','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(46,1,1,'Other','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(47,1,1,'Other OneSite property','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(48,1,1,'Other property','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(49,1,1,'Other publication','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(50,1,1,'Other site','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(51,1,1,'PMC-owned Website','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(52,1,1,'PeopleWithPets.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(53,1,1,'Preferred employer program','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(54,1,1,'Prior resident','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(55,1,1,'Property website','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(56,1,1,'Radio Advertising','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(57,1,1,'Referral companies/merchants','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(58,1,1,'Rent.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(59,1,1,'RentAndMove.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(60,1,1,'RentClicks.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(61,1,1,'RentJungle.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(62,1,1,'RentNet.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(63,1,1,'RentWiki.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(64,1,1,'Rentals.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(65,1,1,'Rentping.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(66,1,1,'Roomster.net','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(67,1,1,'Senior Living Magazine','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(68,1,1,'Site-owned website','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(69,1,1,'TV Advertising','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(70,1,1,'Tinker Take Off','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(71,1,1,'UMoveFree.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(72,1,1,'Unknown/Would not give','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(73,1,1,'Yahoo.com','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(74,1,1,'Yellow pages','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(75,1,2,'Criminal background','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(76,1,2,'No credit history','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(77,1,2,'No employment history','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(78,1,2,'Poor credit history','2018-10-31 18:47:17',0,'2018-06-12 18:01:26',0),(79,1,2,'Poor employment history','2018-10-31 18:47:17',0,'2018-06-12 18:01:26',0),(80,1,2,'Poor rental history','2018-10-31 18:47:17',0,'2018-06-12 18:01:26',0),(81,1,2,'No rental history','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(82,1,2,'Other','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(83,1,3,'Abandoned Apartment','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(84,1,3,'Acquired a pet','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(85,1,3,'Added a roommate','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(86,1,3,'Amenities lacking','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(87,1,3,'Bought condominium','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(88,1,3,'Bought home','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(89,1,3,'Bought townhome','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(90,1,3,'Changed jobs','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(91,1,3,'Closer to airport','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(92,1,3,'Closer to town/city','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(93,1,3,'Closer to work','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(94,1,3,'Corporate or short term lease only','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(95,1,3,'Death or illness','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(96,1,3,'Dissatisfied for another reason','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(97,1,3,'Divorce','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(98,1,3,'Employment transfer','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(99,1,3,'Evicted for another reason','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(100,1,3,'Evicted for criminal reasons','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(101,1,3,'Evicted for non-compliance with community policies','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(102,1,3,'Evicted for non-payment of rent','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(103,1,3,'Generally unhappy with property','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(104,1,3,'Getting married','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(105,1,3,'High utility costs','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(106,1,3,'Leaving/graduating school','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(107,1,3,'Lifestyle change for another reason','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(108,1,3,'Loss of employment from the PMC','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(109,1,3,'Lost a job','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(110,1,3,'Lost a roommate','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(111,1,3,'Marital status change','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(112,1,3,'Military transfer','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(113,1,3,'Money problems','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(114,1,3,'Moving closer to home','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(115,1,3,'Moving home','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(116,1,3,'No reason given','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(117,1,3,'Noise problem','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(118,1,3,'Non-renewal of lease','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(119,1,3,'Other','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(120,1,3,'Parking problems','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(121,1,3,'Personal reasons/concerns','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(122,1,3,'Property disaster','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(123,1,3,'Rental increase','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(124,1,3,'Rentin home','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(125,1,3,'Returning/going to school','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(126,1,3,'Road construction','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(127,1,3,'Selling/old house','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(128,1,3,'Skipped during eviction process','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(129,1,3,'Skipped without notice','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(130,1,4,'ADA accessible','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(131,1,4,'Amenities lacking','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(132,1,4,'Color palette','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(133,1,4,'Drive up appeal','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(134,1,4,'Furniture','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(135,1,4,'Lease term','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(136,1,4,'Location to employment','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(137,1,4,'Location to family','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(138,1,4,'Location to shopping and entertainment','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(139,1,4,'Meets square footage needs','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(140,1,4,'Personnel','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(141,1,4,'Pet allowances','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(142,1,4,'Point of lease e-commerce offers','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(143,1,4,'Priing','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(144,1,4,'Public transportation','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(145,1,4,'School district','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(146,1,4,'Special','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(147,1,5,'Amenities ^ Amenities lacking','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(148,1,5,'Amenities ^ Bedroom size','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(149,1,5,'Amenities ^ Color scheme','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(150,1,5,'Amenities ^ Competition has better amenities','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(151,1,5,'Amenities ^ Objection to floor plan','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(152,1,5,'Cost ^ Competition is less expensive','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(153,1,5,'Cost ^ No specials/concessions','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(154,1,5,'Cost ^ Too expensive','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(155,1,5,'Inactive ^ Inactive','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(156,1,5,'Location ^ Location','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(157,1,5,'Location ^ Road construction','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(158,1,5,'Location ^ Too close to highway','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(159,1,5,'Not available ^ Unit/floor plan not available','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(160,1,5,'Not interested ^ Bought/rented house instead','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(161,1,5,'Not interested ^ Changed their mind','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(162,1,5,'Not interested ^ Not interested','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(163,1,5,'Not qualified ^ Credit rating below standard','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(164,1,5,'Not qualified ^ Criminal history not allowed','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(165,1,5,'Not qualified ^ Does not meet property criteria','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(166,1,5,'Not qualified ^ Oversized/unallowable pet','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(167,1,5,'Not qualified ^ Rental history not allowed','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(168,1,5,'Not qualified ^ Roommate/spouse unqualified','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(169,1,5,'Not qualified ^ Too many occupants','2018-06-12 18:01:26',0,'2018-06-12 18:01:26',0),(170,1,2,'Application declined','2018-06-30 00:53:45',0,'2018-06-30 00:53:45',0),(171,1,6,'Application was declined','2018-07-23 16:14:19',0,'2018-07-23 16:14:19',-99998),(172,1,6,'Rental Agreement was updated','2018-07-23 16:14:19',0,'2018-07-23 16:14:19',-99998),(175,1,7,'Accountants','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(176,1,7,'Advertising/Public Relations','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(177,1,7,'Aerospace','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(178,1,7,'Agribusiness','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(179,1,7,'Agricultural Services & Products','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(180,1,7,'Agriculture','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(181,1,7,'Air Transport','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(182,1,7,'Air Transport Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(183,1,7,'Airlines','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(184,1,7,'Alcoholic Beverages','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(185,1,7,'Alternative Energy Production & Services','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(186,1,7,'Architectural Services','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(187,1,7,'Attorneys/Law Firms','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(188,1,7,'Auto Dealers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(189,1,7,'Auto Dealers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(190,1,7,'Auto Manufacturers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(191,1,7,'Automotive','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(192,1,7,'Banking','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(193,1,7,'Banks','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(194,1,7,'Banks','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(195,1,7,'Bars & Restaurants','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(196,1,7,'Beer','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(197,1,7,'Books','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(198,1,7,'Broadcasters','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(199,1,7,'Builders/General Contractors','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(200,1,7,'Builders/Residential','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(201,1,7,'Building Materials & Equipment','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(202,1,7,'Building Trade Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(203,1,7,'Business Associations','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(204,1,7,'Business Services','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(205,1,7,'Cable & Satellite TV Production & Distribution','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(206,1,7,'Candidate Committees','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(207,1,7,'Candidate Committees','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(208,1,7,'Candidate Committees','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(209,1,7,'Car Dealers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(210,1,7,'Car Dealers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(211,1,7,'Car Manufacturers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(212,1,7,'Casinos / Gambling','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(213,1,7,'Cattle Ranchers/Livestock','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(214,1,7,'Chemical & Related Manufacturing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(215,1,7,'Chiropractors','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(216,1,7,'Civil Servants/Public Officials','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(217,1,7,'Clergy & Religious Organizations','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(218,1,7,'Clothing Manufacturing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(219,1,7,'Coal Mining','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(220,1,7,'Colleges','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(221,1,7,'Commercial Banks','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(222,1,7,'Commercial TV & Radio Stations','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(223,1,7,'Communications/Electronics','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(224,1,7,'Computer Software','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(225,1,7,'Conservative/Republican','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(226,1,7,'Construction','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(227,1,7,'Construction Services','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(228,1,7,'Construction Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(229,1,7,'Credit Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(230,1,7,'Crop Production & Basic Processing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(231,1,7,'Cruise Lines','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(232,1,7,'Cruise Ships & Lines','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(233,1,7,'Dairy','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(234,1,7,'Defense','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(235,1,7,'Defense Aerospace','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(236,1,7,'Defense Electronics','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(237,1,7,'Defense/Foreign Policy Advocates','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(238,1,7,'Democratic Candidate Committees','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(239,1,7,'Democratic Leadership PACs','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(240,1,7,'Democratic/Liberal','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(241,1,7,'Dentists','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(242,1,7,'Doctors & Other Health Professionals','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(243,1,7,'Drug Manufacturers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(244,1,7,'Education','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(245,1,7,'Electric Utilities','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(246,1,7,'Electronics Manufacturing & Equipment','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(247,1,7,'Electronics','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(248,1,7,'Energy & Natural Resources','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(249,1,7,'Entertainment Industry','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(250,1,7,'Environment','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(251,1,7,'Farm Bureaus','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(252,1,7,'Farming','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(253,1,7,'Finance / Credit Companies','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(254,1,7,'Finance','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(255,1,7,'Food & Beverage','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(256,1,7,'Food Processing & Sales','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(257,1,7,'Food Products Manufacturing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(258,1,7,'Food Stores','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(259,1,7,'For-profit Education','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(260,1,7,'For-profit Prisons','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(261,1,7,'Foreign & Defense Policy','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(262,1,7,'Forestry & Forest Products','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(263,1,7,'Foundations','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(264,1,7,'Funeral Services','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(265,1,7,'Gambling & Casinos','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(266,1,7,'Gambling','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(267,1,7,'Garbage Collection/Waste Management','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(268,1,7,'Gas & Oil','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(269,1,7,'Gay & Lesbian Rights & Issues','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(270,1,7,'General Contractors','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(271,1,7,'Government Employee Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(272,1,7,'Government Employees','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(273,1,7,'Gun Control','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(274,1,7,'Gun Rights','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(275,1,7,'Health','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(276,1,7,'Health Professionals','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(277,1,7,'Health Services/HMOs','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(278,1,7,'Hedge Funds','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(279,1,7,'HMOs & Health Care Services','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(280,1,7,'Home Builders','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(281,1,7,'Hospitals & Nursing Homes','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(282,1,7,'Hotels','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(283,1,7,'Human Rights','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(284,1,7,'Ideological/Single-Issue','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(285,1,7,'Indian Gaming','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(286,1,7,'Industrial Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(287,1,7,'Insurance','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(288,1,7,'Internet','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(289,1,7,'Israel Policy','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(290,1,7,'Labor','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(291,1,7,'Lawyers & Lobbyists','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(292,1,7,'Lawyers / Law Firms','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(293,1,7,'Leadership PACs','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(294,1,7,'Liberal/Democratic','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(295,1,7,'Liquor','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(296,1,7,'Livestock','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(297,1,7,'Lobbyists','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(298,1,7,'Lodging / Tourism','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(299,1,7,'Logging','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(300,1,7,'Manufacturing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(301,1,7,'Marine Transport','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(302,1,7,'Meat processing & products','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(303,1,7,'Medical Supplies','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(304,1,7,'Mining','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(305,1,7,'Misc Business','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(306,1,7,'Misc Finance','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(307,1,7,'Misc Manufacturing & Distributing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(308,1,7,'Misc Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(309,1,7,'Miscellaneous Defense','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(310,1,7,'Miscellaneous Services','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(311,1,7,'Mortgage Bankers & Brokers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(312,1,7,'Motion Picture Production & Distribution','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(313,1,7,'Music Production','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(314,1,7,'Natural Gas Pipelines','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(315,1,7,'Newspaper','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(316,1,7,'Non-profits','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(317,1,7,'Nurses','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(318,1,7,'Nursing Homes/Hospitals','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(319,1,7,'Nutritional & Dietary Supplements','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(320,1,7,'Oil & Gas','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(321,1,7,'Payday Lenders','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(322,1,7,'Pharmaceutical Manufacturing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(323,1,7,'Pharmaceuticals / Health Products','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(324,1,7,'Phone Companies','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(325,1,7,'Physicians & Other Health Professionals','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(326,1,7,'Postal Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(327,1,7,'Poultry & Eggs','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(328,1,7,'Power Utilities','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(329,1,7,'Printing & Publishing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(330,1,7,'Private Equity & Investment Firms','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(331,1,7,'Pro-Israel','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(332,1,7,'Professional Sports','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(333,1,7,'Progressive/Democratic','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(334,1,7,'Public Employees','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(335,1,7,'Public Sector Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(336,1,7,'Publishing & Printing','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(337,1,7,'Radio/TV Stations','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(338,1,7,'Railroads','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(339,1,7,'Real Estate','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(340,1,7,'Record Companies/Singers','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(341,1,7,'Recorded Music & Music Production','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(342,1,7,'Recreation / Live Entertainment','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(343,1,7,'Religious Organizations/Clergy','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(344,1,7,'Republican Candidate Committees','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(345,1,7,'Republican Leadership PACs','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(346,1,7,'Republican/Conservative','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(347,1,7,'Residential Construction','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(348,1,7,'Restaurants & Drinking Establishments','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(349,1,7,'Retail Sales','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(350,1,7,'Retired','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(351,1,7,'Savings & Loans','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(352,1,7,'Schools/Education','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(353,1,7,'Sea Transport','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(354,1,7,'Securities & Investment','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(355,1,7,'Special Trade Contractors','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(356,1,7,'Sports','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(357,1,7,'Steel Production','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(358,1,7,'Stock Brokers/Investment Industry','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(359,1,7,'Student Loan Companies','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(360,1,7,'Sugar Cane & Sugar Beets','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(361,1,7,'Teachers Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(362,1,7,'Teachers/Education','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(363,1,7,'Telecom Services & Equipment','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(364,1,7,'Telephone Utilities','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(365,1,7,'Textiles','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(366,1,7,'Timber','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(367,1,7,'Tobacco','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(368,1,7,'Transportation','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(369,1,7,'Transportation Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(370,1,7,'Trash Collection/Waste Management','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(371,1,7,'Trucking','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(372,1,7,'TV / Movies / Music','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(373,1,7,'TV Production','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(374,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(375,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(376,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(377,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(378,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(379,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(380,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(381,1,7,'Unions','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(382,1,7,'Universities','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(383,1,7,'Vegetables & Fruits','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(384,1,7,'Venture Capital','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(385,1,7,'Waste Management','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(386,1,7,'Wine','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0),(387,1,7,'Women\'s Issues','2018-07-24 09:09:29',0,'2018-07-24 09:09:29',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TBind`
--

LOCK TABLES `TBind` WRITE;
/*!40000 ALTER TABLE `TBind` DISABLE KEYS */;
INSERT INTO `TBind` VALUES (1,1,1,1,15,1,'2019-01-03 00:00:00','9999-12-31 00:00:00',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,1,1,1,14,1,'2019-01-03 00:00:00','9999-12-31 00:00:00',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,1,1,2,15,2,'2019-01-03 00:00:00','9999-12-31 00:00:00',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,1,1,2,14,2,'2019-01-03 00:00:00','9999-12-31 00:00:00',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(5,1,1,4,15,3,'2019-01-03 00:00:00','9999-12-31 00:00:00',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(6,1,1,4,14,3,'2019-01-03 00:00:00','9999-12-31 00:00:00',0,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
INSERT INTO `TWS` VALUES (6,'RARBcacheBot','','RARBcacheBot','2019-05-16 08:27:11','Steves-MacBook-Pro-2.local',4,'2019-05-16 08:22:11','2019-05-16 08:22:11','2018-06-02 13:09:58','2019-05-16 01:22:11'),(7,'ARSliceCacheBot','','ARSliceCacheBot','2019-05-16 08:27:11','Steves-MacBook-Pro-2.local',4,'2019-05-16 08:22:11','2019-05-16 08:22:11','2018-06-02 13:09:58','2019-05-16 01:22:11'),(9,'ManualTaskBot','','ManualTaskBot','2019-05-17 08:22:11','Steves-MacBook-Pro-2.local',4,'2019-05-16 08:22:11','2019-05-16 08:22:11','2018-06-02 13:09:58','2019-05-16 01:22:11'),(11,'SecDepCacheBot','','SecDepCacheBot','2019-05-16 08:27:11','Steves-MacBook-Pro-2.local',4,'2019-05-16 08:22:11','2019-05-16 08:22:11','2018-06-02 13:09:58','2019-05-16 01:22:11'),(12,'AcctSliceCacheBot','','AcctSliceCacheBot','2019-05-16 08:27:11','Steves-MacBook-Pro-2.local',4,'2019-05-16 08:22:11','2019-05-16 08:22:11','2018-06-02 13:09:58','2019-05-16 01:22:11');
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
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Task`
--

LOCK TABLES `Task` WRITE;
/*!40000 ALTER TABLE `Task` DISABLE KEYS */;
INSERT INTO `Task` VALUES (1,1,1,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','2019-01-31 20:00:00','2019-01-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,1,1,'Review all receivables for accuracy','ManualTaskBot','2019-01-31 20:00:00','2019-01-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,1,1,'Compare total cash deposits to bank statement','ManualTaskBot','2019-01-31 20:00:00','2019-01-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,1,1,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','2019-01-31 07:00:00','2019-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(5,1,1,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','2019-01-31 07:00:00','2019-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(6,1,1,'Make certain that all suspense accounts have been closed out','ManualTaskBot','2019-01-31 07:00:00','2019-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(7,1,1,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','2019-01-31 07:00:00','2019-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(8,1,1,'Print Rent Roll Activity Report','ManualTaskBot','2019-01-31 07:00:00','2019-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(9,1,1,'Print Rent Roll Report','ManualTaskBot','2019-01-31 07:00:00','2019-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(10,1,1,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','2019-01-31 07:00:00','2019-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(11,1,2,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','0001-01-31 20:00:00','0001-01-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-05-16 08:22:11',-8,'2019-05-16 08:22:11',-8),(12,1,2,'Review all receivables for accuracy','ManualTaskBot','0001-01-31 20:00:00','0001-01-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-05-16 08:22:11',-8,'2019-05-16 08:22:11',-8),(13,1,2,'Compare total cash deposits to bank statement','ManualTaskBot','0001-01-31 20:00:00','0001-01-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-05-16 08:22:11',-8,'2019-05-16 08:22:11',-8),(14,1,2,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','0001-01-31 07:00:00','0001-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-05-16 08:22:11',-8,'2019-05-16 08:22:11',-8),(15,1,2,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','0001-01-31 07:00:00','0001-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-05-16 08:22:11',-8,'2019-05-16 08:22:11',-8),(16,1,2,'Make certain that all suspense accounts have been closed out','ManualTaskBot','0001-01-31 07:00:00','0001-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-05-16 08:22:11',-8,'2019-05-16 08:22:11',-8),(17,1,2,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','0001-01-31 07:00:00','0001-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-05-16 08:22:11',-8,'2019-05-16 08:22:11',-8),(18,1,2,'Print Rent Roll Activity Report','ManualTaskBot','0001-01-31 07:00:00','0001-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-05-16 08:22:11',-8,'2019-05-16 08:22:11',-8),(19,1,2,'Print Rent Roll Report','ManualTaskBot','0001-01-31 07:00:00','0001-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-05-16 08:22:11',-8,'2019-05-16 08:22:11',-8),(20,1,2,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','0001-01-31 07:00:00','0001-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-05-16 08:22:11',-8,'2019-05-16 08:22:11',-8);
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TaskList`
--

LOCK TABLES `TaskList` WRITE;
/*!40000 ALTER TABLE `TaskList` DISABLE KEYS */;
INSERT INTO `TaskList` VALUES (1,1,0,1,'Monthly Close',6,'2019-01-31 17:00:00','2019-01-20 17:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','0000-00-00 00:00:00',86400000000000,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,1,1,1,'Monthly Close',6,'0001-01-31 17:00:00','0001-01-20 17:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','0000-00-00 00:00:00',86400000000000,'','2019-05-16 08:22:11',-8,'2019-05-16 08:22:11',-8);
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Transactant`
--

LOCK TABLES `Transactant` WRITE;
/*!40000 ALTER TABLE `Transactant` DISABLE KEYS */;
INSERT INTO `Transactant` VALUES (1,1,0,'Marcela','Kai','Middleton','Beverlee','3Com Corp',0,'MarcelaM1710@abiz.com','MarcelaM1275@bdiddy.com','(186) 537-1629','(536) 973-8865','42730 Pioneer','','St. Petersburg','TN','98081','USA','',0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,1,0,'Caprice','Nell','Wilder','Evia','BJ\'s Wholesale Club, Inc.',0,'CWilder5966@yahoo.com','CapriceWilder863@aol.com','(878) 257-3627','(582) 544-6390','19064 Third','','Peoria','WI','54425','USA','',0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,1,0,'Arcelia','Brandie','Hoffman','Genaro','John Hancock Financial Services Inc.',0,'AHoffman7623@bdiddy.com','ArceliaHoffman723@hotmail.com','(833) 257-5947','(650) 994-0781','91978 Main','','Miramar','FL','28162','USA','',0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,1,0,'Kayleen','Kiana','Harmon','Patti','Insight Enterprises Inc.',0,'KayleenHarmon963@comcast.net','KHarmon128@aol.com','(313) 804-1019','(215) 137-3719','13323 Apache','','Portsmouth','IA','23237','USA','',0,'','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
INSERT INTO `User` VALUES (1,1,0,'1965-02-10','Easter Wood','29674 Mountain View,Murrieta,MI 27887','(235) 121-8853','EasterW339@abiz.com','52918 Tenth,Minneapolis,OH 31847',1,0,234,15,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,1,0,'1992-06-28','Ginette Spencer','53732 6th,Myrtle Beach,WV 22540','(700) 620-9989','GinetteS6350@aol.com','80192 3rd,Spokane,AZ 40456',0,0,295,49,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,1,0,'1993-11-04','Sherlene Odom','83870 Kansas,Knoxville,OR 55089','(691) 626-8868','SherleneOdom69@bdiddy.com','33534 Pine,Bradenton,NM 24728',1,0,371,41,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(4,1,0,'1982-10-20','Floria Thompson','28953 Hillside,Hialeah,CO 39106','(712) 579-4606','FThompson7316@abiz.com','21126 New Jersey,Las Cruces,WI 40495',0,0,350,62,'2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Vehicle`
--

LOCK TABLES `Vehicle` WRITE;
/*!40000 ALTER TABLE `Vehicle` DISABLE KEYS */;
INSERT INTO `Vehicle` VALUES (1,1,1,'car','BMW','M5','Copper',2000,'ZD0TYM6QOQF7RQPE','MO','Y5J1O61','9200179','2019-01-03','2020-03-01','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(2,2,1,'car','Pontiac','GTO','Blue',2005,'YCMC3FF4RP9QGXED','WA','5ZJJ753','6361638','2019-01-03','2020-03-01','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0),(3,4,1,'car','Ford','F-Series','Gray',1996,'VLUGK260U3LRTQTY','MD','M5V5C43','9152724','2019-01-03','2020-03-01','2019-01-17 01:12:42',0,'2019-01-17 01:12:42',0);
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

-- Dump completed on 2019-06-17 14:11:28
