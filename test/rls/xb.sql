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
INSERT INTO `AR` VALUES (1,1,'Application Fee',0,0,0,9,46,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,15.0000,'2018-08-15 02:53:15',211,'2017-11-10 23:24:23',0,0,0),(2,1,'Application Fee (no assessment)',0,1,0,7,46,'Application fee taken, no assessment made','0000-00-00 00:00:00','9999-12-31 00:00:00',5,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(3,1,'Apply Payment',0,1,0,10,9,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 02:59:12',211,'2017-11-10 23:24:23',0,0,0),(4,1,'Bad Debt Write-Off',0,2,0,71,9,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(5,1,'Bank Service Fee (Deposit Account)',0,2,0,72,4,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(6,1,'Bank Service Fee (Operating Account)',0,2,0,72,3,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(7,1,'Broken Window charge',0,0,0,9,59,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:00:40',211,'2017-11-10 23:24:23',0,0,0),(8,1,'Damage Fee',0,0,0,9,59,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:01:10',211,'2017-11-10 23:24:23',0,0,0),(9,1,'Deposit to Deposit Account (FRB96953)',0,1,0,4,6,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(10,1,'Deposit to Operating Account (FRB54320)',0,1,0,3,6,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(11,1,'Electric Base Fee',0,0,0,9,36,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(12,1,'Electric Overage',0,0,0,9,37,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:01:56',211,'2017-11-10 23:24:23',0,0,0),(13,1,'Eviction Fee Reimbursement',0,0,0,9,56,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(14,1,'Auto-Generated Floating Deposit Assessment',0,3,0,9,12,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(15,1,'Receive Floating Security Deposit',0,1,0,6,9,'','0000-00-00 00:00:00','9999-12-31 00:00:00',13,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(16,1,'Gas Base Fee',0,0,0,9,40,'','1900-01-01 00:00:00','9999-12-29 00:00:00',2,50.0000,'2018-07-23 16:16:36',211,'2017-11-10 23:24:23',0,6,4),(17,1,'Gas Base Overage',0,0,0,9,41,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:02:36',211,'2017-11-10 23:24:23',0,0,0),(18,1,'Insufficient Funds Fee',0,0,0,9,48,'','1900-01-01 00:00:00','9999-12-29 00:00:00',64,25.0000,'2018-08-15 03:03:11',211,'2017-11-10 23:24:23',0,0,0),(19,1,'Late Fee',0,0,0,9,47,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,5.0000,'2018-08-15 03:03:22',211,'2017-11-10 23:24:23',0,0,0),(20,1,'Month to Month Fee',0,0,0,9,49,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(21,1,'No Show / Termination Fee',0,0,0,9,51,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:03:52',211,'2017-11-10 23:24:23',0,0,0),(22,1,'Other Special Tenant Charges',0,0,0,9,61,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(23,1,'Pet Fee',0,0,0,9,52,'','1900-01-01 00:00:00','9999-12-31 00:00:00',192,50.0000,'2018-07-04 04:13:35',211,'2017-11-10 23:24:23',0,0,0),(24,1,'Pet Rent',0,0,0,9,53,'','1900-01-01 00:00:00','9999-12-30 00:00:00',144,10.0000,'2018-07-23 16:16:52',211,'2017-11-10 23:24:23',0,6,4),(25,1,'Receive a Payment',0,1,0,6,10,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(26,1,'Rent Non-Taxable',0,0,0,9,18,'','1900-01-01 00:00:00','9999-12-30 00:00:00',16,0.0000,'2018-07-23 16:14:54',211,'2017-11-10 23:24:23',0,6,4),(27,1,'Rent Taxable',0,0,0,9,17,'','1900-01-01 00:00:00','9999-12-30 00:00:00',16,0.0000,'2018-07-23 16:15:28',211,'2017-11-10 23:24:23',0,4,0),(28,1,'Security Deposit Assessment',0,0,0,9,11,'normal deposit','1900-01-01 00:00:00','9999-12-30 00:00:00',96,0.0000,'2018-08-15 03:04:36',211,'2017-11-10 23:24:23',0,0,0),(29,1,'Security Deposit Forfeiture',0,0,0,11,58,'Forfeit','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(30,1,'Security Deposit Refund',0,0,0,11,5,'Refund','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(31,1,'Special Cleaning Fee',0,0,0,9,55,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(32,1,'Tenant Expense Chargeback',0,0,0,9,54,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(33,1,'Vending Income',0,1,0,7,65,'','0000-00-00 00:00:00','9999-12-31 00:00:00',5,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(34,1,'Water and Sewer Base Fee',0,0,0,9,38,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(35,1,'Water and Sewer Overage',0,0,0,9,39,'','1900-01-01 00:00:00','9999-12-30 00:00:00',64,0.0000,'2018-08-15 03:05:54',211,'2017-11-10 23:24:23',0,0,0),(36,1,'Auto-gen Application Fee Asmt',0,3,0,9,46,'','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(37,1,'Receive Application Fee (auto-gen asmt)',0,1,0,6,9,'Application fee taken, autogen asmt','0000-00-00 00:00:00','9999-12-31 00:00:00',13,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(38,1,'XFER  Operating to SecDep',0,2,0,4,3,'Move money from Operating acct to Sec Dep','0000-00-00 00:00:00','9999-12-31 00:00:00',0,0.0000,'2017-11-27 21:49:03',0,'2017-11-10 23:24:23',0,0,0),(39,1,'Vehicle Registration Fee',0,0,3,9,75,'','2018-01-01 00:00:00','9999-12-31 00:00:00',320,10.0000,'2018-07-04 04:14:56',211,'2018-07-03 02:47:47',211,0,0),(40,1,'Rent ST000',0,0,0,9,18,'Default rent assessment for rentable type RType000','1970-01-01 00:00:00','9999-12-30 00:00:00',16,1000.0000,'2018-08-15 03:07:17',211,'2018-07-27 06:49:30',0,6,4),(41,1,'Rent ST001',0,0,0,9,18,'Default rent assessment for rentable type RType001','1970-01-01 00:00:00','9999-12-30 00:00:00',16,1500.0000,'2018-08-15 03:07:28',211,'2018-07-27 06:49:30',0,6,4),(42,1,'Rent ST002',0,0,0,9,18,'Default rent assessment for rentable type RType002','1970-01-01 00:00:00','9999-12-30 00:00:00',16,1750.0000,'2018-08-15 03:07:40',211,'2018-07-27 06:49:30',0,6,4),(43,1,'Rent ST003',0,0,0,9,18,'Default rent assessment for rentable type RType003','1970-01-01 00:00:00','9999-12-30 00:00:00',16,2500.0000,'2018-08-15 03:16:30',211,'2018-07-27 06:49:30',0,6,4),(44,1,'Rent CP000',0,0,0,9,18,'Default rent assessment for rentable type Car Port 000','1970-01-01 00:00:00','9999-12-30 00:00:00',16,35.0000,'2018-08-15 03:16:45',211,'2018-07-27 06:49:30',0,6,4),(45,1,'Rent 1BR1BA',0,0,0,9,18,'Default rent assessment for rentable type Standard','1970-01-01 00:00:00','9999-12-31 00:00:00',20,1000.0000,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,6,4),(46,1,'Rent 2BR1BA',0,0,0,9,18,'Default rent assessment for rentable type Deluxe','1970-01-01 00:00:00','9999-12-31 00:00:00',20,1500.0000,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,6,4),(47,1,'Rent 2BR2BA',0,0,0,9,18,'Default rent assessment for rentable type Gold','1970-01-01 00:00:00','9999-12-31 00:00:00',20,1750.0000,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,6,4),(48,1,'Rent 3BR2BA',0,0,0,9,18,'Default rent assessment for rentable type Platinum','1970-01-01 00:00:00','9999-12-31 00:00:00',20,2500.0000,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,6,4),(49,1,'Rent H11',0,0,0,9,18,'Default rent assessment for rentable type HotelStd','1970-01-01 00:00:00','9999-12-31 00:00:00',20,65.0000,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,6,4),(50,1,'Rent H21',0,0,0,9,18,'Default rent assessment for rentable type HotelGold','1970-01-01 00:00:00','9999-12-31 00:00:00',20,75.0000,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,6,4),(51,1,'Rent H22',0,0,0,9,18,'Default rent assessment for rentable type HotelPlatinum','1970-01-01 00:00:00','9999-12-31 00:00:00',20,85.0000,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,6,4),(52,1,'Rent CP000',0,0,0,9,18,'Default rent assessment for rentable type Car Port 000','1970-01-01 00:00:00','9999-12-31 00:00:00',20,35.0000,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,0,0);
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
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Assessments`
--

LOCK TABLES `Assessments` WRITE;
/*!40000 ALTER TABLE `Assessments` DISABLE KEYS */;
INSERT INTO `Assessments` VALUES (1,0,0,0,1,1,0,0,1,1000.0000,'2019-02-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',26,0,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,1,0,0,1,1,0,0,1,1000.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',26,2,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(3,0,0,0,1,1,0,0,1,2000.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',28,2,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(4,0,0,0,1,1,0,0,1,935.4800,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',26,2,'prorated for 29 of 31 days','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(5,0,0,0,1,1,15,1,1,10.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',39,2,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(6,0,0,0,1,2,0,0,2,1000.0000,'2019-02-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',26,0,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(7,6,0,0,1,2,0,0,2,1000.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',26,2,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(8,0,0,0,1,2,0,0,2,2000.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',28,2,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(9,0,0,0,1,2,14,1,2,10.0000,'2019-02-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',24,8,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(10,9,0,0,1,2,14,1,2,10.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',24,2,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(11,0,0,0,1,2,0,0,2,935.4800,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',26,2,'prorated for 29 of 31 days','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(12,0,0,0,1,2,14,1,2,50.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',23,2,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(13,0,0,0,1,2,14,1,2,9.3500,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',24,2,'prorated for 29 of 31 days','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(14,0,0,0,1,2,15,2,2,10.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',39,2,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(15,0,0,0,1,3,0,0,3,1500.0000,'2019-02-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',26,0,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(16,15,0,0,1,3,0,0,3,1500.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',26,2,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(17,0,0,0,1,3,0,0,3,3000.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',28,2,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(18,0,0,0,1,3,14,2,3,10.0000,'2019-02-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',24,8,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(19,18,0,0,1,3,14,2,3,10.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',24,2,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(20,0,0,0,1,3,0,0,3,1403.2300,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',26,2,'prorated for 29 of 31 days','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(21,0,0,0,1,3,14,2,3,50.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',23,2,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(22,0,0,0,1,3,14,2,3,9.3500,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',24,2,'prorated for 29 of 31 days','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(23,0,0,0,1,3,15,3,3,10.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',39,2,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(24,0,0,0,1,4,0,0,4,1500.0000,'2019-02-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',26,0,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(25,24,0,0,1,4,0,0,4,1500.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',26,2,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(26,0,0,0,1,4,0,0,4,3000.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',28,2,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(27,0,0,0,1,4,14,3,4,10.0000,'2019-02-01 00:00:00','2020-03-01 00:00:00',6,4,0,'',24,8,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(28,27,0,0,1,4,14,3,4,10.0000,'2019-02-01 00:00:00','2019-02-01 00:00:00',6,4,0,'',24,2,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(29,0,0,0,1,4,0,0,4,1403.2300,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',26,2,'prorated for 29 of 31 days','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(30,0,0,0,1,4,14,3,4,50.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',23,2,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(31,0,0,0,1,4,14,3,4,9.3500,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',24,2,'prorated for 29 of 31 days','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(32,0,0,0,1,4,15,4,4,10.0000,'2019-01-03 00:00:00','2019-01-03 00:00:00',0,0,0,'',39,2,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0);
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
INSERT INTO `BusinessProperties` VALUES (1,1,'general',0,'{\"Epochs\": {\"Daily\": \"2017-01-01T00:00:00Z\", \"Weekly\": \"2017-01-01T00:00:00Z\", \"Yearly\": \"2017-01-01T00:00:00Z\", \"Monthly\": \"2017-01-01T00:00:00Z\", \"Quarterly\": \"2017-01-01T00:00:00Z\"}, \"PetFees\": [\"Pet Fee\", \"Pet Rent\"], \"VehicleFees\": [\"Vehicle Registration Fee\"]}','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Deposit`
--

LOCK TABLES `Deposit` WRITE;
/*!40000 ALTER TABLE `Deposit` DISABLE KEYS */;
INSERT INTO `Deposit` VALUES (1,1,2,1,'2019-01-03',4895.4700,0.0000,0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(2,1,1,1,'2019-01-03',10000.0000,0.0000,0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(3,1,2,1,'2019-02-01',5030.0000,0.0000,0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DepositPart`
--

LOCK TABLES `DepositPart` WRITE;
/*!40000 ALTER TABLE `DepositPart` DISABLE KEYS */;
INSERT INTO `DepositPart` VALUES (1,1,1,3,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(2,1,1,4,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(3,1,1,8,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(4,1,1,9,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(5,1,1,10,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(6,1,1,11,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(7,1,1,15,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(8,1,1,16,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(9,1,1,17,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(10,1,1,18,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(11,1,1,22,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(12,1,1,23,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(13,1,1,24,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(14,1,1,25,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(15,2,1,2,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(16,2,1,6,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(17,2,1,13,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(18,2,1,20,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(19,3,1,1,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(20,3,1,5,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(21,3,1,7,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(22,3,1,12,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(23,3,1,14,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(24,3,1,19,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(25,3,1,21,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=118 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Journal`
--

LOCK TABLES `Journal` WRITE;
/*!40000 ALTER TABLE `Journal` DISABLE KEYS */;
INSERT INTO `Journal` VALUES (1,1,'2019-02-01 00:00:00',1000.0000,1,2,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,1,'2019-01-03 00:00:00',2000.0000,1,3,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,1,'2019-01-03 00:00:00',935.4800,1,4,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,1,'2019-01-03 00:00:00',10.0000,1,5,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(5,1,'2019-02-01 00:00:00',1000.0000,1,7,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(6,1,'2019-01-03 00:00:00',2000.0000,1,8,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(7,1,'2019-02-01 00:00:00',10.0000,1,10,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(8,1,'2019-01-03 00:00:00',935.4800,1,11,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(9,1,'2019-01-03 00:00:00',50.0000,1,12,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(10,1,'2019-01-03 00:00:00',9.3500,1,13,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(11,1,'2019-01-03 00:00:00',10.0000,1,14,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(12,1,'2019-02-01 00:00:00',1500.0000,1,16,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(13,1,'2019-01-03 00:00:00',3000.0000,1,17,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(14,1,'2019-02-01 00:00:00',10.0000,1,19,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(15,1,'2019-01-03 00:00:00',1403.2300,1,20,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(16,1,'2019-01-03 00:00:00',50.0000,1,21,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(17,1,'2019-01-03 00:00:00',9.3500,1,22,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(18,1,'2019-01-03 00:00:00',10.0000,1,23,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(19,1,'2019-02-01 00:00:00',1500.0000,1,25,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(20,1,'2019-01-03 00:00:00',3000.0000,1,26,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(21,1,'2019-02-01 00:00:00',10.0000,1,28,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(22,1,'2019-01-03 00:00:00',1403.2300,1,29,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(23,1,'2019-01-03 00:00:00',50.0000,1,30,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(24,1,'2019-01-03 00:00:00',9.3500,1,31,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(25,1,'2019-01-03 00:00:00',10.0000,1,32,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(26,1,'2019-02-01 00:00:00',1000.0000,2,1,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(27,1,'2019-01-03 00:00:00',2000.0000,2,2,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(28,1,'2019-01-03 00:00:00',935.4800,2,3,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(29,1,'2019-01-03 00:00:00',10.0000,2,4,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(30,1,'2019-02-01 00:00:00',1000.0000,2,5,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(31,1,'2019-01-03 00:00:00',2000.0000,2,6,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(32,1,'2019-02-01 00:00:00',10.0000,2,7,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(33,1,'2019-01-03 00:00:00',935.4800,2,8,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(34,1,'2019-01-03 00:00:00',50.0000,2,9,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(35,1,'2019-01-03 00:00:00',9.3500,2,10,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(36,1,'2019-01-03 00:00:00',10.0000,2,11,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(37,1,'2019-02-01 00:00:00',1500.0000,2,12,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(38,1,'2019-01-03 00:00:00',3000.0000,2,13,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(39,1,'2019-02-01 00:00:00',10.0000,2,14,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(40,1,'2019-01-03 00:00:00',1403.2300,2,15,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(41,1,'2019-01-03 00:00:00',50.0000,2,16,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(42,1,'2019-01-03 00:00:00',9.3500,2,17,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(43,1,'2019-01-03 00:00:00',10.0000,2,18,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(44,1,'2019-02-01 00:00:00',1500.0000,2,19,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(45,1,'2019-01-03 00:00:00',3000.0000,2,20,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(46,1,'2019-02-01 00:00:00',10.0000,2,21,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(47,1,'2019-01-03 00:00:00',1403.2300,2,22,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(48,1,'2019-01-03 00:00:00',50.0000,2,23,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(49,1,'2019-01-03 00:00:00',9.3500,2,24,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(50,1,'2019-01-03 00:00:00',10.0000,2,25,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(51,1,'2019-01-03 00:00:00',9.3500,2,10,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(52,1,'2019-01-03 00:00:00',10.0000,2,11,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(53,1,'2019-01-03 00:00:00',50.0000,2,9,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(54,1,'2019-01-03 00:00:00',935.4800,2,8,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(55,1,'2019-01-03 00:00:00',995.1700,2,6,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(56,1,'2019-01-03 00:00:00',935.4800,2,6,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(57,1,'2019-01-03 00:00:00',50.0000,2,6,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(58,1,'2019-01-03 00:00:00',9.3500,2,6,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(59,1,'2019-01-03 00:00:00',10.0000,2,6,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(60,1,'2019-02-01 00:00:00',10.0000,2,7,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(61,1,'2019-02-01 00:00:00',990.0000,2,5,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(62,1,'2019-02-01 00:00:00',10.0000,2,5,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(63,1,'2019-01-03 00:00:00',9.3500,2,17,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(64,1,'2019-01-03 00:00:00',10.0000,2,18,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(65,1,'2019-01-03 00:00:00',50.0000,2,16,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(66,1,'2019-01-03 00:00:00',1403.2300,2,15,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(67,1,'2019-01-03 00:00:00',1527.4200,2,13,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(68,1,'2019-01-03 00:00:00',1403.2300,2,13,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(69,1,'2019-01-03 00:00:00',50.0000,2,13,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(70,1,'2019-01-03 00:00:00',9.3500,2,13,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(71,1,'2019-01-03 00:00:00',10.0000,2,13,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(72,1,'2019-02-01 00:00:00',10.0000,2,14,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(73,1,'2019-02-01 00:00:00',1490.0000,2,12,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(74,1,'2019-02-01 00:00:00',10.0000,2,12,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(75,1,'2019-01-03 00:00:00',9.3500,2,24,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(76,1,'2019-01-03 00:00:00',10.0000,2,25,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(77,1,'2019-01-03 00:00:00',50.0000,2,23,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(78,1,'2019-01-03 00:00:00',1403.2300,2,22,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(79,1,'2019-01-03 00:00:00',1527.4200,2,20,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(80,1,'2019-01-03 00:00:00',1403.2300,2,20,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(81,1,'2019-01-03 00:00:00',50.0000,2,20,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(82,1,'2019-01-03 00:00:00',9.3500,2,20,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(83,1,'2019-01-03 00:00:00',10.0000,2,20,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(84,1,'2019-02-01 00:00:00',10.0000,2,21,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(85,1,'2019-02-01 00:00:00',1490.0000,2,19,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(86,1,'2019-02-01 00:00:00',10.0000,2,19,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(87,1,'2019-01-03 00:00:00',10.0000,2,4,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(88,1,'2019-01-03 00:00:00',935.4800,2,3,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(89,1,'2019-01-03 00:00:00',1054.5200,2,2,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(90,1,'2019-01-03 00:00:00',935.4800,2,2,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(91,1,'2019-01-03 00:00:00',10.0000,2,2,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(92,1,'2019-02-01 00:00:00',1000.0000,2,1,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(93,1,'2019-01-03 00:00:00',935.4800,4,3,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(94,1,'2019-01-03 00:00:00',10.0000,4,4,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(95,1,'2019-01-03 00:00:00',935.4800,4,8,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(96,1,'2019-01-03 00:00:00',50.0000,4,9,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(97,1,'2019-01-03 00:00:00',9.3500,4,10,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(98,1,'2019-01-03 00:00:00',10.0000,4,11,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(99,1,'2019-01-03 00:00:00',1403.2300,4,15,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(100,1,'2019-01-03 00:00:00',50.0000,4,16,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(101,1,'2019-01-03 00:00:00',9.3500,4,17,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(102,1,'2019-01-03 00:00:00',10.0000,4,18,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(103,1,'2019-01-03 00:00:00',1403.2300,4,22,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(104,1,'2019-01-03 00:00:00',50.0000,4,23,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(105,1,'2019-01-03 00:00:00',9.3500,4,24,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(106,1,'2019-01-03 00:00:00',10.0000,4,25,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(107,1,'2019-01-03 00:00:00',2000.0000,4,2,'auto-transfer for deposit DEP-1','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(108,1,'2019-01-03 00:00:00',2000.0000,4,6,'auto-transfer for deposit DEP-1','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(109,1,'2019-01-03 00:00:00',3000.0000,4,13,'auto-transfer for deposit DEP-1','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(110,1,'2019-01-03 00:00:00',3000.0000,4,20,'auto-transfer for deposit DEP-1','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(111,1,'2019-02-01 00:00:00',1000.0000,4,1,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(112,1,'2019-02-01 00:00:00',1000.0000,4,5,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(113,1,'2019-02-01 00:00:00',10.0000,4,7,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(114,1,'2019-02-01 00:00:00',1500.0000,4,12,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(115,1,'2019-02-01 00:00:00',10.0000,4,14,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(116,1,'2019-02-01 00:00:00',1500.0000,4,19,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(117,1,'2019-02-01 00:00:00',10.0000,4,21,'auto-transfer for deposit DEP-2','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=118 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `JournalAllocation`
--

LOCK TABLES `JournalAllocation` WRITE;
/*!40000 ALTER TABLE `JournalAllocation` DISABLE KEYS */;
INSERT INTO `JournalAllocation` VALUES (1,1,1,1,1,0,0,1000.0000,2,0,'d 12001 1000.00, c 41001 1000.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,1,2,1,1,0,0,2000.0000,3,0,'d 12001 2000.00, c 30000 2000.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,1,3,1,1,0,0,935.4800,4,0,'d 12001 935.48, c 41001 935.48','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,1,4,1,1,0,0,10.0000,5,0,'d 12001 10.00, c 41416 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(5,1,5,2,2,0,0,1000.0000,7,0,'d 12001 1000.00, c 41001 1000.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(6,1,6,2,2,0,0,2000.0000,8,0,'d 12001 2000.00, c 30000 2000.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(7,1,7,2,2,0,0,10.0000,10,0,'d 12001 10.00, c 41408 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(8,1,8,2,2,0,0,935.4800,11,0,'d 12001 935.48, c 41001 935.48','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(9,1,9,2,2,0,0,50.0000,12,0,'d 12001 50.00, c 41407 50.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(10,1,10,2,2,0,0,9.3500,13,0,'d 12001 9.35, c 41408 9.35','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(11,1,11,2,2,0,0,10.0000,14,0,'d 12001 10.00, c 41416 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(12,1,12,3,3,0,0,1500.0000,16,0,'d 12001 1500.00, c 41001 1500.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(13,1,13,3,3,0,0,3000.0000,17,0,'d 12001 3000.00, c 30000 3000.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(14,1,14,3,3,0,0,10.0000,19,0,'d 12001 10.00, c 41408 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(15,1,15,3,3,0,0,1403.2300,20,0,'d 12001 1403.23, c 41001 1403.23','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(16,1,16,3,3,0,0,50.0000,21,0,'d 12001 50.00, c 41407 50.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(17,1,17,3,3,0,0,9.3500,22,0,'d 12001 9.35, c 41408 9.35','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(18,1,18,3,3,0,0,10.0000,23,0,'d 12001 10.00, c 41416 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(19,1,19,4,4,0,0,1500.0000,25,0,'d 12001 1500.00, c 41001 1500.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(20,1,20,4,4,0,0,3000.0000,26,0,'d 12001 3000.00, c 30000 3000.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(21,1,21,4,4,0,0,10.0000,28,0,'d 12001 10.00, c 41408 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(22,1,22,4,4,0,0,1403.2300,29,0,'d 12001 1403.23, c 41001 1403.23','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(23,1,23,4,4,0,0,50.0000,30,0,'d 12001 50.00, c 41407 50.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(24,1,24,4,4,0,0,9.3500,31,0,'d 12001 9.35, c 41408 9.35','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(25,1,25,4,4,0,0,10.0000,32,0,'d 12001 10.00, c 41416 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(26,1,26,0,0,1,0,1000.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(27,1,27,0,0,1,0,2000.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(28,1,28,0,0,1,0,935.4800,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(29,1,29,0,0,1,0,10.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(30,1,30,0,0,2,0,1000.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(31,1,31,0,0,2,0,2000.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(32,1,32,0,0,2,0,10.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(33,1,33,0,0,2,0,935.4800,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(34,1,34,0,0,2,0,50.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(35,1,35,0,0,2,0,9.3500,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(36,1,36,0,0,2,0,10.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(37,1,37,0,0,3,0,1500.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(38,1,38,0,0,3,0,3000.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(39,1,39,0,0,3,0,10.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(40,1,40,0,0,3,0,1403.2300,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(41,1,41,0,0,3,0,50.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(42,1,42,0,0,3,0,9.3500,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(43,1,43,0,0,3,0,10.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(44,1,44,0,0,4,0,1500.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(45,1,45,0,0,4,0,3000.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(46,1,46,0,0,4,0,10.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(47,1,47,0,0,4,0,1403.2300,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(48,1,48,0,0,4,0,50.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(49,1,49,0,0,4,0,9.3500,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(50,1,50,0,0,4,0,10.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(51,1,51,2,2,2,10,9.3500,8,0,'ASM(8) d 12999 9.35,c 12001 9.35','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(52,1,52,2,2,2,11,10.0000,8,0,'ASM(8) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(53,1,53,2,2,2,9,50.0000,8,0,'ASM(8) d 12999 50.00,c 12001 50.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(54,1,54,2,2,2,8,935.4800,8,0,'ASM(8) d 12999 935.48,c 12001 935.48','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(55,1,55,2,2,2,6,995.1700,8,0,'ASM(8) d 12999 995.17,c 12001 995.17','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(56,1,56,2,2,2,6,935.4800,11,0,'ASM(11) d 12999 935.48,c 12001 935.48','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(57,1,57,2,2,2,6,50.0000,12,0,'ASM(12) d 12999 50.00,c 12001 50.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(58,1,58,2,2,2,6,9.3500,13,0,'ASM(13) d 12999 9.35,c 12001 9.35','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(59,1,59,2,2,2,6,10.0000,14,0,'ASM(14) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(60,1,60,2,2,2,7,10.0000,7,0,'ASM(7) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(61,1,61,2,2,2,5,990.0000,7,0,'ASM(7) d 12999 990.00,c 12001 990.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(62,1,62,2,2,2,5,10.0000,10,0,'ASM(10) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(63,1,63,3,3,3,17,9.3500,17,0,'ASM(17) d 12999 9.35,c 12001 9.35','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(64,1,64,3,3,3,18,10.0000,17,0,'ASM(17) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(65,1,65,3,3,3,16,50.0000,17,0,'ASM(17) d 12999 50.00,c 12001 50.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(66,1,66,3,3,3,15,1403.2300,17,0,'ASM(17) d 12999 1403.23,c 12001 1403.23','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(67,1,67,3,3,3,13,1527.4200,17,0,'ASM(17) d 12999 1527.42,c 12001 1527.42','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(68,1,68,3,3,3,13,1403.2300,20,0,'ASM(20) d 12999 1403.23,c 12001 1403.23','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(69,1,69,3,3,3,13,50.0000,21,0,'ASM(21) d 12999 50.00,c 12001 50.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(70,1,70,3,3,3,13,9.3500,22,0,'ASM(22) d 12999 9.35,c 12001 9.35','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(71,1,71,3,3,3,13,10.0000,23,0,'ASM(23) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(72,1,72,3,3,3,14,10.0000,16,0,'ASM(16) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(73,1,73,3,3,3,12,1490.0000,16,0,'ASM(16) d 12999 1490.00,c 12001 1490.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(74,1,74,3,3,3,12,10.0000,19,0,'ASM(19) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(75,1,75,4,4,4,24,9.3500,26,0,'ASM(26) d 12999 9.35,c 12001 9.35','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(76,1,76,4,4,4,25,10.0000,26,0,'ASM(26) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(77,1,77,4,4,4,23,50.0000,26,0,'ASM(26) d 12999 50.00,c 12001 50.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(78,1,78,4,4,4,22,1403.2300,26,0,'ASM(26) d 12999 1403.23,c 12001 1403.23','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(79,1,79,4,4,4,20,1527.4200,26,0,'ASM(26) d 12999 1527.42,c 12001 1527.42','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(80,1,80,4,4,4,20,1403.2300,29,0,'ASM(29) d 12999 1403.23,c 12001 1403.23','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(81,1,81,4,4,4,20,50.0000,30,0,'ASM(30) d 12999 50.00,c 12001 50.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(82,1,82,4,4,4,20,9.3500,31,0,'ASM(31) d 12999 9.35,c 12001 9.35','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(83,1,83,4,4,4,20,10.0000,32,0,'ASM(32) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(84,1,84,4,4,4,21,10.0000,25,0,'ASM(25) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(85,1,85,4,4,4,19,1490.0000,25,0,'ASM(25) d 12999 1490.00,c 12001 1490.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(86,1,86,4,4,4,19,10.0000,28,0,'ASM(28) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(87,1,87,1,1,1,4,10.0000,3,0,'ASM(3) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(88,1,88,1,1,1,3,935.4800,3,0,'ASM(3) d 12999 935.48,c 12001 935.48','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(89,1,89,1,1,1,2,1054.5200,3,0,'ASM(3) d 12999 1054.52,c 12001 1054.52','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(90,1,90,1,1,1,2,935.4800,4,0,'ASM(4) d 12999 935.48,c 12001 935.48','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(91,1,91,1,1,1,2,10.0000,5,0,'ASM(5) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(92,1,92,1,1,1,1,1000.0000,2,0,'ASM(2) d 12999 1000.00,c 12001 1000.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(93,1,93,0,1,1,3,935.4800,3,0,'d 10105 935.4800, c 10999 935.4800','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(94,1,94,0,1,1,4,10.0000,3,0,'d 10105 10.0000, c 10999 10.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(95,1,95,0,2,2,8,935.4800,8,0,'d 10105 935.4800, c 10999 935.4800','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(96,1,96,0,2,2,9,50.0000,8,0,'d 10105 50.0000, c 10999 50.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(97,1,97,0,2,2,10,9.3500,8,0,'d 10105 9.3500, c 10999 9.3500','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(98,1,98,0,2,2,11,10.0000,8,0,'d 10105 10.0000, c 10999 10.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(99,1,99,0,3,3,15,1403.2300,17,0,'d 10105 1403.2300, c 10999 1403.2300','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(100,1,100,0,3,3,16,50.0000,17,0,'d 10105 50.0000, c 10999 50.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(101,1,101,0,3,3,17,9.3500,17,0,'d 10105 9.3500, c 10999 9.3500','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(102,1,102,0,3,3,18,10.0000,17,0,'d 10105 10.0000, c 10999 10.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(103,1,103,0,4,4,22,1403.2300,26,0,'d 10105 1403.2300, c 10999 1403.2300','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(104,1,104,0,4,4,23,50.0000,26,0,'d 10105 50.0000, c 10999 50.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(105,1,105,0,4,4,24,9.3500,26,0,'d 10105 9.3500, c 10999 9.3500','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(106,1,106,0,4,4,25,10.0000,26,0,'d 10105 10.0000, c 10999 10.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(107,1,107,0,1,1,2,2000.0000,0,0,'d 10104 2000.0000, c 10999 2000.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(108,1,108,0,2,2,6,2000.0000,0,0,'d 10104 2000.0000, c 10999 2000.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(109,1,109,0,3,3,13,3000.0000,0,0,'d 10104 3000.0000, c 10999 3000.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(110,1,110,0,4,4,20,3000.0000,0,0,'d 10104 3000.0000, c 10999 3000.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(111,1,111,0,1,1,1,1000.0000,2,0,'d 10105 1000.0000, c 10999 1000.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(112,1,112,0,2,2,5,1000.0000,0,0,'d 10105 1000.0000, c 10999 1000.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(113,1,113,0,2,2,7,10.0000,7,0,'d 10105 10.0000, c 10999 10.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(114,1,114,0,3,3,12,1500.0000,0,0,'d 10105 1500.0000, c 10999 1500.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(115,1,115,0,3,3,14,10.0000,16,0,'d 10105 10.0000, c 10999 10.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(116,1,116,0,4,4,19,1500.0000,0,0,'d 10105 1500.0000, c 10999 1500.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(117,1,117,0,4,4,21,10.0000,25,0,'d 10105 10.0000, c 10999 10.0000','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=235 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerEntry`
--

LOCK TABLES `LedgerEntry` WRITE;
/*!40000 ALTER TABLE `LedgerEntry` DISABLE KEYS */;
INSERT INTO `LedgerEntry` VALUES (1,1,1,1,9,1,1,0,'2019-02-01 00:00:00',1000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,1,1,1,18,1,1,0,'2019-02-01 00:00:00',-1000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,1,2,2,9,1,1,0,'2019-01-03 00:00:00',2000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,1,2,2,11,1,1,0,'2019-01-03 00:00:00',-2000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(5,1,3,3,9,1,1,0,'2019-01-03 00:00:00',935.4800,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(6,1,3,3,18,1,1,0,'2019-01-03 00:00:00',-935.4800,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(7,1,4,4,9,1,1,0,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(8,1,4,4,75,1,1,0,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(9,1,5,5,9,2,2,0,'2019-02-01 00:00:00',1000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(10,1,5,5,18,2,2,0,'2019-02-01 00:00:00',-1000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(11,1,6,6,9,2,2,0,'2019-01-03 00:00:00',2000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(12,1,6,6,11,2,2,0,'2019-01-03 00:00:00',-2000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(13,1,7,7,9,2,2,0,'2019-02-01 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(14,1,7,7,53,2,2,0,'2019-02-01 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(15,1,8,8,9,2,2,0,'2019-01-03 00:00:00',935.4800,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(16,1,8,8,18,2,2,0,'2019-01-03 00:00:00',-935.4800,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(17,1,9,9,9,2,2,0,'2019-01-03 00:00:00',50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(18,1,9,9,52,2,2,0,'2019-01-03 00:00:00',-50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(19,1,10,10,9,2,2,0,'2019-01-03 00:00:00',9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(20,1,10,10,53,2,2,0,'2019-01-03 00:00:00',-9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(21,1,11,11,9,2,2,0,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(22,1,11,11,75,2,2,0,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(23,1,12,12,9,3,3,0,'2019-02-01 00:00:00',1500.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(24,1,12,12,18,3,3,0,'2019-02-01 00:00:00',-1500.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(25,1,13,13,9,3,3,0,'2019-01-03 00:00:00',3000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(26,1,13,13,11,3,3,0,'2019-01-03 00:00:00',-3000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(27,1,14,14,9,3,3,0,'2019-02-01 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(28,1,14,14,53,3,3,0,'2019-02-01 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(29,1,15,15,9,3,3,0,'2019-01-03 00:00:00',1403.2300,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(30,1,15,15,18,3,3,0,'2019-01-03 00:00:00',-1403.2300,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(31,1,16,16,9,3,3,0,'2019-01-03 00:00:00',50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(32,1,16,16,52,3,3,0,'2019-01-03 00:00:00',-50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(33,1,17,17,9,3,3,0,'2019-01-03 00:00:00',9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(34,1,17,17,53,3,3,0,'2019-01-03 00:00:00',-9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(35,1,18,18,9,3,3,0,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(36,1,18,18,75,3,3,0,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(37,1,19,19,9,4,4,0,'2019-02-01 00:00:00',1500.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(38,1,19,19,18,4,4,0,'2019-02-01 00:00:00',-1500.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(39,1,20,20,9,4,4,0,'2019-01-03 00:00:00',3000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(40,1,20,20,11,4,4,0,'2019-01-03 00:00:00',-3000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(41,1,21,21,9,4,4,0,'2019-02-01 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(42,1,21,21,53,4,4,0,'2019-02-01 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(43,1,22,22,9,4,4,0,'2019-01-03 00:00:00',1403.2300,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(44,1,22,22,18,4,4,0,'2019-01-03 00:00:00',-1403.2300,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(45,1,23,23,9,4,4,0,'2019-01-03 00:00:00',50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(46,1,23,23,52,4,4,0,'2019-01-03 00:00:00',-50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(47,1,24,24,9,4,4,0,'2019-01-03 00:00:00',9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(48,1,24,24,53,4,4,0,'2019-01-03 00:00:00',-9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(49,1,25,25,9,4,4,0,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(50,1,25,25,75,4,4,0,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(51,1,26,26,6,0,0,1,'2019-02-01 00:00:00',1000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(52,1,26,26,10,0,0,1,'2019-02-01 00:00:00',-1000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(53,1,27,27,6,0,0,1,'2019-01-03 00:00:00',2000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(54,1,27,27,10,0,0,1,'2019-01-03 00:00:00',-2000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(55,1,28,28,6,0,0,1,'2019-01-03 00:00:00',935.4800,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(56,1,28,28,10,0,0,1,'2019-01-03 00:00:00',-935.4800,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(57,1,29,29,6,0,0,1,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(58,1,29,29,10,0,0,1,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(59,1,30,30,6,0,0,2,'2019-02-01 00:00:00',1000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(60,1,30,30,10,0,0,2,'2019-02-01 00:00:00',-1000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(61,1,31,31,6,0,0,2,'2019-01-03 00:00:00',2000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(62,1,31,31,10,0,0,2,'2019-01-03 00:00:00',-2000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(63,1,32,32,6,0,0,2,'2019-02-01 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(64,1,32,32,10,0,0,2,'2019-02-01 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(65,1,33,33,6,0,0,2,'2019-01-03 00:00:00',935.4800,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(66,1,33,33,10,0,0,2,'2019-01-03 00:00:00',-935.4800,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(67,1,34,34,6,0,0,2,'2019-01-03 00:00:00',50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(68,1,34,34,10,0,0,2,'2019-01-03 00:00:00',-50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(69,1,35,35,6,0,0,2,'2019-01-03 00:00:00',9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(70,1,35,35,10,0,0,2,'2019-01-03 00:00:00',-9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(71,1,36,36,6,0,0,2,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(72,1,36,36,10,0,0,2,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(73,1,37,37,6,0,0,3,'2019-02-01 00:00:00',1500.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(74,1,37,37,10,0,0,3,'2019-02-01 00:00:00',-1500.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(75,1,38,38,6,0,0,3,'2019-01-03 00:00:00',3000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(76,1,38,38,10,0,0,3,'2019-01-03 00:00:00',-3000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(77,1,39,39,6,0,0,3,'2019-02-01 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(78,1,39,39,10,0,0,3,'2019-02-01 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(79,1,40,40,6,0,0,3,'2019-01-03 00:00:00',1403.2300,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(80,1,40,40,10,0,0,3,'2019-01-03 00:00:00',-1403.2300,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(81,1,41,41,6,0,0,3,'2019-01-03 00:00:00',50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(82,1,41,41,10,0,0,3,'2019-01-03 00:00:00',-50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(83,1,42,42,6,0,0,3,'2019-01-03 00:00:00',9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(84,1,42,42,10,0,0,3,'2019-01-03 00:00:00',-9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(85,1,43,43,6,0,0,3,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(86,1,43,43,10,0,0,3,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(87,1,44,44,6,0,0,4,'2019-02-01 00:00:00',1500.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(88,1,44,44,10,0,0,4,'2019-02-01 00:00:00',-1500.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(89,1,45,45,6,0,0,4,'2019-01-03 00:00:00',3000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(90,1,45,45,10,0,0,4,'2019-01-03 00:00:00',-3000.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(91,1,46,46,6,0,0,4,'2019-02-01 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(92,1,46,46,10,0,0,4,'2019-02-01 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(93,1,47,47,6,0,0,4,'2019-01-03 00:00:00',1403.2300,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(94,1,47,47,10,0,0,4,'2019-01-03 00:00:00',-1403.2300,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(95,1,48,48,6,0,0,4,'2019-01-03 00:00:00',50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(96,1,48,48,10,0,0,4,'2019-01-03 00:00:00',-50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(97,1,49,49,6,0,0,4,'2019-01-03 00:00:00',9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(98,1,49,49,10,0,0,4,'2019-01-03 00:00:00',-9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(99,1,50,50,6,0,0,4,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(100,1,50,50,10,0,0,4,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(101,1,51,51,10,2,2,2,'2019-01-03 00:00:00',9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(102,1,51,51,9,2,2,2,'2019-01-03 00:00:00',-9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(103,1,52,52,10,2,2,2,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(104,1,52,52,9,2,2,2,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(105,1,53,53,10,2,2,2,'2019-01-03 00:00:00',50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(106,1,53,53,9,2,2,2,'2019-01-03 00:00:00',-50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(107,1,54,54,10,2,2,2,'2019-01-03 00:00:00',935.4800,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(108,1,54,54,9,2,2,2,'2019-01-03 00:00:00',-935.4800,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(109,1,55,55,10,2,2,2,'2019-01-03 00:00:00',995.1700,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(110,1,55,55,9,2,2,2,'2019-01-03 00:00:00',-995.1700,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(111,1,56,56,10,2,2,2,'2019-01-03 00:00:00',935.4800,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(112,1,56,56,9,2,2,2,'2019-01-03 00:00:00',-935.4800,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(113,1,57,57,10,2,2,2,'2019-01-03 00:00:00',50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(114,1,57,57,9,2,2,2,'2019-01-03 00:00:00',-50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(115,1,58,58,10,2,2,2,'2019-01-03 00:00:00',9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(116,1,58,58,9,2,2,2,'2019-01-03 00:00:00',-9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(117,1,59,59,10,2,2,2,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(118,1,59,59,9,2,2,2,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(119,1,60,60,10,2,2,2,'2019-02-01 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(120,1,60,60,9,2,2,2,'2019-02-01 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(121,1,61,61,10,2,2,2,'2019-02-01 00:00:00',990.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(122,1,61,61,9,2,2,2,'2019-02-01 00:00:00',-990.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(123,1,62,62,10,2,2,2,'2019-02-01 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(124,1,62,62,9,2,2,2,'2019-02-01 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(125,1,63,63,10,3,3,3,'2019-01-03 00:00:00',9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(126,1,63,63,9,3,3,3,'2019-01-03 00:00:00',-9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(127,1,64,64,10,3,3,3,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(128,1,64,64,9,3,3,3,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(129,1,65,65,10,3,3,3,'2019-01-03 00:00:00',50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(130,1,65,65,9,3,3,3,'2019-01-03 00:00:00',-50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(131,1,66,66,10,3,3,3,'2019-01-03 00:00:00',1403.2300,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(132,1,66,66,9,3,3,3,'2019-01-03 00:00:00',-1403.2300,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(133,1,67,67,10,3,3,3,'2019-01-03 00:00:00',1527.4200,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(134,1,67,67,9,3,3,3,'2019-01-03 00:00:00',-1527.4200,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(135,1,68,68,10,3,3,3,'2019-01-03 00:00:00',1403.2300,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(136,1,68,68,9,3,3,3,'2019-01-03 00:00:00',-1403.2300,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(137,1,69,69,10,3,3,3,'2019-01-03 00:00:00',50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(138,1,69,69,9,3,3,3,'2019-01-03 00:00:00',-50.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(139,1,70,70,10,3,3,3,'2019-01-03 00:00:00',9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(140,1,70,70,9,3,3,3,'2019-01-03 00:00:00',-9.3500,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(141,1,71,71,10,3,3,3,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(142,1,71,71,9,3,3,3,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(143,1,72,72,10,3,3,3,'2019-02-01 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(144,1,72,72,9,3,3,3,'2019-02-01 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(145,1,73,73,10,3,3,3,'2019-02-01 00:00:00',1490.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(146,1,73,73,9,3,3,3,'2019-02-01 00:00:00',-1490.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(147,1,74,74,10,3,3,3,'2019-02-01 00:00:00',10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(148,1,74,74,9,3,3,3,'2019-02-01 00:00:00',-10.0000,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(149,1,75,75,10,4,4,4,'2019-01-03 00:00:00',9.3500,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(150,1,75,75,9,4,4,4,'2019-01-03 00:00:00',-9.3500,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(151,1,76,76,10,4,4,4,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(152,1,76,76,9,4,4,4,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(153,1,77,77,10,4,4,4,'2019-01-03 00:00:00',50.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(154,1,77,77,9,4,4,4,'2019-01-03 00:00:00',-50.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(155,1,78,78,10,4,4,4,'2019-01-03 00:00:00',1403.2300,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(156,1,78,78,9,4,4,4,'2019-01-03 00:00:00',-1403.2300,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(157,1,79,79,10,4,4,4,'2019-01-03 00:00:00',1527.4200,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(158,1,79,79,9,4,4,4,'2019-01-03 00:00:00',-1527.4200,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(159,1,80,80,10,4,4,4,'2019-01-03 00:00:00',1403.2300,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(160,1,80,80,9,4,4,4,'2019-01-03 00:00:00',-1403.2300,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(161,1,81,81,10,4,4,4,'2019-01-03 00:00:00',50.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(162,1,81,81,9,4,4,4,'2019-01-03 00:00:00',-50.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(163,1,82,82,10,4,4,4,'2019-01-03 00:00:00',9.3500,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(164,1,82,82,9,4,4,4,'2019-01-03 00:00:00',-9.3500,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(165,1,83,83,10,4,4,4,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(166,1,83,83,9,4,4,4,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(167,1,84,84,10,4,4,4,'2019-02-01 00:00:00',10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(168,1,84,84,9,4,4,4,'2019-02-01 00:00:00',-10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(169,1,85,85,10,4,4,4,'2019-02-01 00:00:00',1490.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(170,1,85,85,9,4,4,4,'2019-02-01 00:00:00',-1490.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(171,1,86,86,10,4,4,4,'2019-02-01 00:00:00',10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(172,1,86,86,9,4,4,4,'2019-02-01 00:00:00',-10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(173,1,87,87,10,1,1,1,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(174,1,87,87,9,1,1,1,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(175,1,88,88,10,1,1,1,'2019-01-03 00:00:00',935.4800,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(176,1,88,88,9,1,1,1,'2019-01-03 00:00:00',-935.4800,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(177,1,89,89,10,1,1,1,'2019-01-03 00:00:00',1054.5200,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(178,1,89,89,9,1,1,1,'2019-01-03 00:00:00',-1054.5200,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(179,1,90,90,10,1,1,1,'2019-01-03 00:00:00',935.4800,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(180,1,90,90,9,1,1,1,'2019-01-03 00:00:00',-935.4800,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(181,1,91,91,10,1,1,1,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(182,1,91,91,9,1,1,1,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(183,1,92,92,10,1,1,1,'2019-02-01 00:00:00',1000.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(184,1,92,92,9,1,1,1,'2019-02-01 00:00:00',-1000.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(185,1,93,93,4,1,0,1,'2019-01-03 00:00:00',935.4800,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(186,1,93,93,6,1,0,1,'2019-01-03 00:00:00',-935.4800,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(187,1,94,94,4,1,0,1,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(188,1,94,94,6,1,0,1,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(189,1,95,95,4,2,0,2,'2019-01-03 00:00:00',935.4800,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(190,1,95,95,6,2,0,2,'2019-01-03 00:00:00',-935.4800,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(191,1,96,96,4,2,0,2,'2019-01-03 00:00:00',50.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(192,1,96,96,6,2,0,2,'2019-01-03 00:00:00',-50.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(193,1,97,97,4,2,0,2,'2019-01-03 00:00:00',9.3500,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(194,1,97,97,6,2,0,2,'2019-01-03 00:00:00',-9.3500,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(195,1,98,98,4,2,0,2,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(196,1,98,98,6,2,0,2,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(197,1,99,99,4,3,0,3,'2019-01-03 00:00:00',1403.2300,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(198,1,99,99,6,3,0,3,'2019-01-03 00:00:00',-1403.2300,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(199,1,100,100,4,3,0,3,'2019-01-03 00:00:00',50.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(200,1,100,100,6,3,0,3,'2019-01-03 00:00:00',-50.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(201,1,101,101,4,3,0,3,'2019-01-03 00:00:00',9.3500,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(202,1,101,101,6,3,0,3,'2019-01-03 00:00:00',-9.3500,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(203,1,102,102,4,3,0,3,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(204,1,102,102,6,3,0,3,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(205,1,103,103,4,4,0,4,'2019-01-03 00:00:00',1403.2300,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(206,1,103,103,6,4,0,4,'2019-01-03 00:00:00',-1403.2300,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(207,1,104,104,4,4,0,4,'2019-01-03 00:00:00',50.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(208,1,104,104,6,4,0,4,'2019-01-03 00:00:00',-50.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(209,1,105,105,4,4,0,4,'2019-01-03 00:00:00',9.3500,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(210,1,105,105,6,4,0,4,'2019-01-03 00:00:00',-9.3500,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(211,1,106,106,4,4,0,4,'2019-01-03 00:00:00',10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(212,1,106,106,6,4,0,4,'2019-01-03 00:00:00',-10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(213,1,107,107,3,1,0,1,'2019-01-03 00:00:00',2000.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(214,1,107,107,6,1,0,1,'2019-01-03 00:00:00',-2000.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(215,1,108,108,3,2,0,2,'2019-01-03 00:00:00',2000.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(216,1,108,108,6,2,0,2,'2019-01-03 00:00:00',-2000.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(217,1,109,109,3,3,0,3,'2019-01-03 00:00:00',3000.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(218,1,109,109,6,3,0,3,'2019-01-03 00:00:00',-3000.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(219,1,110,110,3,4,0,4,'2019-01-03 00:00:00',3000.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(220,1,110,110,6,4,0,4,'2019-01-03 00:00:00',-3000.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(221,1,111,111,4,1,0,1,'2019-02-01 00:00:00',1000.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(222,1,111,111,6,1,0,1,'2019-02-01 00:00:00',-1000.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(223,1,112,112,4,2,0,2,'2019-02-01 00:00:00',1000.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(224,1,112,112,6,2,0,2,'2019-02-01 00:00:00',-1000.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(225,1,113,113,4,2,0,2,'2019-02-01 00:00:00',10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(226,1,113,113,6,2,0,2,'2019-02-01 00:00:00',-10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(227,1,114,114,4,3,0,3,'2019-02-01 00:00:00',1500.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(228,1,114,114,6,3,0,3,'2019-02-01 00:00:00',-1500.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(229,1,115,115,4,3,0,3,'2019-02-01 00:00:00',10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(230,1,115,115,6,3,0,3,'2019-02-01 00:00:00',-10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(231,1,116,116,4,4,0,4,'2019-02-01 00:00:00',1500.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(232,1,116,116,6,4,0,4,'2019-02-01 00:00:00',-1500.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(233,1,117,117,4,4,0,4,'2019-02-01 00:00:00',10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(234,1,117,117,6,4,0,4,'2019-02-01 00:00:00',-10.0000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0);
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
INSERT INTO `LedgerMarker` VALUES (1,1,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(2,2,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(3,3,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(4,4,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(5,5,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(6,6,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(7,7,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(8,8,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(9,9,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(10,10,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(11,11,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(12,12,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(13,13,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(14,14,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(15,15,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(16,16,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(17,17,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(18,18,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(19,19,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(20,20,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(21,21,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(22,22,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(23,23,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(24,24,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(25,25,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(26,26,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(27,27,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(28,28,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(29,29,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(30,30,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(31,31,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(32,32,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(33,33,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(34,34,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(35,35,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(36,36,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(37,37,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(38,38,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(39,39,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(40,40,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(41,41,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(42,42,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(43,43,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(44,44,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(45,45,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(46,46,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(47,47,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(48,48,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(49,49,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(50,50,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(51,51,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(52,52,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(53,53,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(54,54,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(55,55,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(56,56,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(57,57,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(58,58,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(59,59,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(60,60,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(61,61,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(62,62,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(63,63,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(64,64,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(65,65,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(66,66,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(67,67,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(68,68,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(69,69,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(70,70,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(71,71,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(72,72,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(73,73,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(74,74,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-11-10 23:24:22',0,'2017-11-10 23:24:22',0),(75,75,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2018-07-03 02:45:37',211,'2018-07-03 02:45:37',211),(76,0,1,1,0,0,'2018-12-20 00:00:00',0.0000,3,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(77,0,1,1,1,0,'2018-12-20 00:00:00',0.0000,3,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(78,0,1,2,0,0,'2018-12-20 00:00:00',0.0000,3,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(79,0,1,2,2,0,'2018-12-20 00:00:00',0.0000,3,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(80,0,1,3,0,0,'2018-12-20 00:00:00',0.0000,3,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(81,0,1,3,3,0,'2018-12-20 00:00:00',0.0000,3,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(82,0,1,4,0,0,'2018-12-20 00:00:00',0.0000,3,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(83,0,1,4,4,0,'2018-12-20 00:00:00',0.0000,3,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
INSERT INTO `Payor` VALUES (1,1,'2ab49a1ce67b248bb147075f90607c8eaf44d7e0ef2b1f37bef4027df5470a15aeef641c',4880.0000,1,0,'ab1eeb696694048a2ab144036d2264a6be3c99762200011e002d1e544bb295163d54be82',30528.0000,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,1,'d1ebd75fbcaf96e3c0df0e6cf23b86979bed74dac73deeb4ba5dcc23da011a354655f544',7318.0000,1,0,'400744168bc690c5c65f0fbb5aa7a297f2fce8061adcc80fdcfcf07024ea6f1545615935',86846.0000,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,1,'b5f6f7eb0c024c59900d96258f3027e1cb820a219aebbe0eccb990287777c4e6cc5360fb',8778.0000,1,0,'38426a992455b79c30db73f0f4d941036c8fd4be2600ba8f64ef60c12c9f61d7eeeeba0f',142547.0000,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,1,'73ab9bdba2b7412b0b69f00f31f6b9dac5fbd6fec1fc3b83d946a034fd05587ca8f2a64f',16397.0000,1,0,'a67cc211a00059bf14868dc1d8ea091e190ca1deacf0c02da547ccd14d82dcbd640dd39b',69808.0000,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
INSERT INTO `Pets` VALUES (1,1,2,2,'dog','Alaskan Malamute','White',29.0000,'Benny','2019-01-03','2020-03-01','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,1,3,3,'dog','Labrador Retriever','Black',39.0000,'Blue','2019-01-03','2020-03-01','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,1,4,4,'dog','Clumber Spaniel','Grey',40.0000,'Rocco','2019-01-03','2020-03-01','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
INSERT INTO `Prospect` VALUES (1,1,'36036 Johnson','Huntington','NV','84059','LithiaMotorsIncHuntington943@yahoo.com','(651) 238-2188','ammunition and explosives operative (munitions worker)',0,'','','','','','0000-00-00','27276 Mountain View, Sarasota, VT 02081','Shellie Wheeler','(291) 783-0652',108,'9 years 3 months','39277 2nd, Jefferson, MO 41318','Graham Atkins','(376) 500-7795',102,'4 years 4 months','','Arron Swanson','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,1,'25745 Jackson','Houma','WA','03300','SystemaxIncH2496@abiz.com','(839) 826-1264','mechanical engineering production manager',0,'','','','','','0000-00-00','64503 Mountain View, Richmond, OK 10694','Sacha Davidson','(855) 139-0746',110,'6 months','41880 S 100, Anchorage, AL 78511','Terry England','(646) 268-8402',115,'8 months','','Roselee Fleming','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,1,'77733 Lee','Springfield','MN','33274','TheBearStearnsCompaniesIncSpringfield899@comcast.net','(330) 438-0738','psychotherapist',0,'','','','','','0000-00-00','11173 North Carolina, Saint Louis, ID 11211','Tamala Rush','(739) 600-7834',128,'3 years 7 months','13431 Second, Palm Springs, CT 31445','Noma Rush','(998) 168-8743',86,'7 years 7 months','','Jannette Foley','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,1,'16376 Malulani','GreenBay','IA','65466','StilwellFinancialIncG8460@gmail.com','(316) 487-1813','keeper of service animals',0,'','','','','','0000-00-00','28533 Second, Huntsville, MS 11528','Claudia Hardin','(513) 339-9468',126,'3 years 6 months','30184 Hampton, Temecula, KS 86258','Tatum Baldwin','(740) 569-9706',126,'8 years 7 months','','Samira Gould','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Receipt`
--

LOCK TABLES `Receipt` WRITE;
/*!40000 ALTER TABLE `Receipt` DISABLE KEYS */;
INSERT INTO `Receipt` VALUES (1,0,1,1,2,1,3,1,'2019-02-01 00:00:00','786623',1000.0000,'',25,'ASM(2) d 12999 1000.00,c 12001 1000.00',2,'payment for ASM-2','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(2,0,1,1,2,2,2,1,'2019-01-03 00:00:00','988459',2000.0000,'',25,'ASM(3) d 12999 1054.52,c 12001 1054.52,ASM(4) d 12999 935.48,c 12001 935.48,ASM(5) d 12999 10.00,c 12001 10.00',2,'payment for ASM-3','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(3,0,1,1,2,1,1,1,'2019-01-03 00:00:00','641803',935.4800,'',25,'ASM(3) d 12999 935.48,c 12001 935.48',2,'payment for ASM-4','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(4,0,1,1,2,1,1,1,'2019-01-03 00:00:00','361739',10.0000,'',25,'ASM(3) d 12999 10.00,c 12001 10.00',2,'payment for ASM-5','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(5,0,1,2,2,1,3,2,'2019-02-01 00:00:00','761883',1000.0000,'',25,'ASM(7) d 12999 990.00,c 12001 990.00,ASM(10) d 12999 10.00,c 12001 10.00',2,'payment for ASM-7','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(6,0,1,2,2,2,2,2,'2019-01-03 00:00:00','38907',2000.0000,'',25,'ASM(8) d 12999 995.17,c 12001 995.17,ASM(11) d 12999 935.48,c 12001 935.48,ASM(12) d 12999 50.00,c 12001 50.00,ASM(13) d 12999 9.35,c 12001 9.35,ASM(14) d 12999 10.00,c 12001 10.00',2,'payment for ASM-8','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(7,0,1,2,2,1,3,2,'2019-02-01 00:00:00','658932',10.0000,'',25,'ASM(7) d 12999 10.00,c 12001 10.00',2,'payment for ASM-10','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(8,0,1,2,2,1,1,2,'2019-01-03 00:00:00','387780',935.4800,'',25,'ASM(8) d 12999 935.48,c 12001 935.48',2,'payment for ASM-11','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(9,0,1,2,2,1,1,2,'2019-01-03 00:00:00','383115',50.0000,'',25,'ASM(8) d 12999 50.00,c 12001 50.00',2,'payment for ASM-12','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(10,0,1,2,2,1,1,2,'2019-01-03 00:00:00','742716',9.3500,'',25,'ASM(8) d 12999 9.35,c 12001 9.35',2,'payment for ASM-13','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(11,0,1,2,2,1,1,2,'2019-01-03 00:00:00','50020',10.0000,'',25,'ASM(8) d 12999 10.00,c 12001 10.00',2,'payment for ASM-14','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(12,0,1,3,2,1,3,3,'2019-02-01 00:00:00','623071',1500.0000,'',25,'ASM(16) d 12999 1490.00,c 12001 1490.00,ASM(19) d 12999 10.00,c 12001 10.00',2,'payment for ASM-16','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(13,0,1,3,2,2,2,3,'2019-01-03 00:00:00','134652',3000.0000,'',25,'ASM(17) d 12999 1527.42,c 12001 1527.42,ASM(20) d 12999 1403.23,c 12001 1403.23,ASM(21) d 12999 50.00,c 12001 50.00,ASM(22) d 12999 9.35,c 12001 9.35,ASM(23) d 12999 10.00,c 12001 10.00',2,'payment for ASM-17','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(14,0,1,3,2,1,3,3,'2019-02-01 00:00:00','281907',10.0000,'',25,'ASM(16) d 12999 10.00,c 12001 10.00',2,'payment for ASM-19','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(15,0,1,3,2,1,1,3,'2019-01-03 00:00:00','772719',1403.2300,'',25,'ASM(17) d 12999 1403.23,c 12001 1403.23',2,'payment for ASM-20','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(16,0,1,3,2,1,1,3,'2019-01-03 00:00:00','659062',50.0000,'',25,'ASM(17) d 12999 50.00,c 12001 50.00',2,'payment for ASM-21','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(17,0,1,3,2,1,1,3,'2019-01-03 00:00:00','663110',9.3500,'',25,'ASM(17) d 12999 9.35,c 12001 9.35',2,'payment for ASM-22','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(18,0,1,3,2,1,1,3,'2019-01-03 00:00:00','961397',10.0000,'',25,'ASM(17) d 12999 10.00,c 12001 10.00',2,'payment for ASM-23','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(19,0,1,4,2,1,3,4,'2019-02-01 00:00:00','813105',1500.0000,'',25,'ASM(25) d 12999 1490.00,c 12001 1490.00,ASM(28) d 12999 10.00,c 12001 10.00',2,'payment for ASM-25','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(20,0,1,4,2,2,2,4,'2019-01-03 00:00:00','569721',3000.0000,'',25,'ASM(26) d 12999 1527.42,c 12001 1527.42,ASM(29) d 12999 1403.23,c 12001 1403.23,ASM(30) d 12999 50.00,c 12001 50.00,ASM(31) d 12999 9.35,c 12001 9.35,ASM(32) d 12999 10.00,c 12001 10.00',2,'payment for ASM-26','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(21,0,1,4,2,1,3,4,'2019-02-01 00:00:00','585917',10.0000,'',25,'ASM(25) d 12999 10.00,c 12001 10.00',2,'payment for ASM-28','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(22,0,1,4,2,1,1,4,'2019-01-03 00:00:00','764712',1403.2300,'',25,'ASM(26) d 12999 1403.23,c 12001 1403.23',2,'payment for ASM-29','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(23,0,1,4,2,1,1,4,'2019-01-03 00:00:00','264449',50.0000,'',25,'ASM(26) d 12999 50.00,c 12001 50.00',2,'payment for ASM-30','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(24,0,1,4,2,1,1,4,'2019-01-03 00:00:00','989579',9.3500,'',25,'ASM(26) d 12999 9.35,c 12001 9.35',2,'payment for ASM-31','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(25,0,1,4,2,1,1,4,'2019-01-03 00:00:00','156333',10.0000,'',25,'ASM(26) d 12999 10.00,c 12001 10.00',2,'payment for ASM-32','','2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=93 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ReceiptAllocation`
--

LOCK TABLES `ReceiptAllocation` WRITE;
/*!40000 ALTER TABLE `ReceiptAllocation` DISABLE KEYS */;
INSERT INTO `ReceiptAllocation` VALUES (1,1,1,1,'2019-02-01 00:00:00',1000.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,2,1,1,'2019-01-03 00:00:00',2000.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,3,1,1,'2019-01-03 00:00:00',935.4800,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,4,1,1,'2019-01-03 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(5,5,1,2,'2019-02-01 00:00:00',1000.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(6,6,1,2,'2019-01-03 00:00:00',2000.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(7,7,1,2,'2019-02-01 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(8,8,1,2,'2019-01-03 00:00:00',935.4800,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(9,9,1,2,'2019-01-03 00:00:00',50.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(10,10,1,2,'2019-01-03 00:00:00',9.3500,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(11,11,1,2,'2019-01-03 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(12,12,1,3,'2019-02-01 00:00:00',1500.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(13,13,1,3,'2019-01-03 00:00:00',3000.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(14,14,1,3,'2019-02-01 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(15,15,1,3,'2019-01-03 00:00:00',1403.2300,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(16,16,1,3,'2019-01-03 00:00:00',50.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(17,17,1,3,'2019-01-03 00:00:00',9.3500,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(18,18,1,3,'2019-01-03 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(19,19,1,4,'2019-02-01 00:00:00',1500.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(20,20,1,4,'2019-01-03 00:00:00',3000.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(21,21,1,4,'2019-02-01 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(22,22,1,4,'2019-01-03 00:00:00',1403.2300,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(23,23,1,4,'2019-01-03 00:00:00',50.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(24,24,1,4,'2019-01-03 00:00:00',9.3500,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(25,25,1,4,'2019-01-03 00:00:00',10.0000,0,0,'d 10999 _, c 12999 _','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(26,10,1,2,'2019-01-03 00:00:00',9.3500,8,0,'ASM(8) d 12999 9.35,c 12001 9.35','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(27,11,1,2,'2019-01-03 00:00:00',10.0000,8,0,'ASM(8) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(28,9,1,2,'2019-01-03 00:00:00',50.0000,8,0,'ASM(8) d 12999 50.00,c 12001 50.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(29,8,1,2,'2019-01-03 00:00:00',935.4800,8,0,'ASM(8) d 12999 935.48,c 12001 935.48','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(30,6,1,2,'2019-01-03 00:00:00',995.1700,8,0,'ASM(8) d 12999 995.17,c 12001 995.17','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(31,6,1,2,'2019-01-03 00:00:00',935.4800,11,0,'ASM(11) d 12999 935.48,c 12001 935.48','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(32,6,1,2,'2019-01-03 00:00:00',50.0000,12,0,'ASM(12) d 12999 50.00,c 12001 50.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(33,6,1,2,'2019-01-03 00:00:00',9.3500,13,0,'ASM(13) d 12999 9.35,c 12001 9.35','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(34,6,1,2,'2019-01-03 00:00:00',10.0000,14,0,'ASM(14) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(35,7,1,2,'2019-02-01 00:00:00',10.0000,7,0,'ASM(7) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(36,5,1,2,'2019-02-01 00:00:00',990.0000,7,0,'ASM(7) d 12999 990.00,c 12001 990.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(37,5,1,2,'2019-02-01 00:00:00',10.0000,10,0,'ASM(10) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(38,17,1,3,'2019-01-03 00:00:00',9.3500,17,0,'ASM(17) d 12999 9.35,c 12001 9.35','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(39,18,1,3,'2019-01-03 00:00:00',10.0000,17,0,'ASM(17) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(40,16,1,3,'2019-01-03 00:00:00',50.0000,17,0,'ASM(17) d 12999 50.00,c 12001 50.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(41,15,1,3,'2019-01-03 00:00:00',1403.2300,17,0,'ASM(17) d 12999 1403.23,c 12001 1403.23','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(42,13,1,3,'2019-01-03 00:00:00',1527.4200,17,0,'ASM(17) d 12999 1527.42,c 12001 1527.42','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(43,13,1,3,'2019-01-03 00:00:00',1403.2300,20,0,'ASM(20) d 12999 1403.23,c 12001 1403.23','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(44,13,1,3,'2019-01-03 00:00:00',50.0000,21,0,'ASM(21) d 12999 50.00,c 12001 50.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(45,13,1,3,'2019-01-03 00:00:00',9.3500,22,0,'ASM(22) d 12999 9.35,c 12001 9.35','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(46,13,1,3,'2019-01-03 00:00:00',10.0000,23,0,'ASM(23) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(47,14,1,3,'2019-02-01 00:00:00',10.0000,16,0,'ASM(16) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(48,12,1,3,'2019-02-01 00:00:00',1490.0000,16,0,'ASM(16) d 12999 1490.00,c 12001 1490.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(49,12,1,3,'2019-02-01 00:00:00',10.0000,19,0,'ASM(19) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(50,24,1,4,'2019-01-03 00:00:00',9.3500,26,0,'ASM(26) d 12999 9.35,c 12001 9.35','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(51,25,1,4,'2019-01-03 00:00:00',10.0000,26,0,'ASM(26) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(52,23,1,4,'2019-01-03 00:00:00',50.0000,26,0,'ASM(26) d 12999 50.00,c 12001 50.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(53,22,1,4,'2019-01-03 00:00:00',1403.2300,26,0,'ASM(26) d 12999 1403.23,c 12001 1403.23','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(54,20,1,4,'2019-01-03 00:00:00',1527.4200,26,0,'ASM(26) d 12999 1527.42,c 12001 1527.42','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(55,20,1,4,'2019-01-03 00:00:00',1403.2300,29,0,'ASM(29) d 12999 1403.23,c 12001 1403.23','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(56,20,1,4,'2019-01-03 00:00:00',50.0000,30,0,'ASM(30) d 12999 50.00,c 12001 50.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(57,20,1,4,'2019-01-03 00:00:00',9.3500,31,0,'ASM(31) d 12999 9.35,c 12001 9.35','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(58,20,1,4,'2019-01-03 00:00:00',10.0000,32,0,'ASM(32) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(59,21,1,4,'2019-02-01 00:00:00',10.0000,25,0,'ASM(25) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(60,19,1,4,'2019-02-01 00:00:00',1490.0000,25,0,'ASM(25) d 12999 1490.00,c 12001 1490.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(61,19,1,4,'2019-02-01 00:00:00',10.0000,28,0,'ASM(28) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(62,4,1,1,'2019-01-03 00:00:00',10.0000,3,0,'ASM(3) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(63,3,1,1,'2019-01-03 00:00:00',935.4800,3,0,'ASM(3) d 12999 935.48,c 12001 935.48','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(64,2,1,1,'2019-01-03 00:00:00',1054.5200,3,0,'ASM(3) d 12999 1054.52,c 12001 1054.52','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(65,2,1,1,'2019-01-03 00:00:00',935.4800,4,0,'ASM(4) d 12999 935.48,c 12001 935.48','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(66,2,1,1,'2019-01-03 00:00:00',10.0000,5,0,'ASM(5) d 12999 10.00,c 12001 10.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(67,1,1,1,'2019-02-01 00:00:00',1000.0000,2,0,'ASM(2) d 12999 1000.00,c 12001 1000.00','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(68,3,1,1,'2019-01-03 00:00:00',935.4800,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(69,4,1,1,'2019-01-03 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(70,8,1,2,'2019-01-03 00:00:00',935.4800,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(71,9,1,2,'2019-01-03 00:00:00',50.0000,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(72,10,1,2,'2019-01-03 00:00:00',9.3500,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(73,11,1,2,'2019-01-03 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(74,15,1,3,'2019-01-03 00:00:00',1403.2300,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(75,16,1,3,'2019-01-03 00:00:00',50.0000,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(76,17,1,3,'2019-01-03 00:00:00',9.3500,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(77,18,1,3,'2019-01-03 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(78,22,1,4,'2019-01-03 00:00:00',1403.2300,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(79,23,1,4,'2019-01-03 00:00:00',50.0000,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(80,24,1,4,'2019-01-03 00:00:00',9.3500,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(81,25,1,4,'2019-01-03 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(82,2,1,1,'2019-01-03 00:00:00',2000.0000,0,0,'d 10104 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(83,6,1,2,'2019-01-03 00:00:00',2000.0000,0,0,'d 10104 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(84,13,1,3,'2019-01-03 00:00:00',3000.0000,0,0,'d 10104 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(85,20,1,4,'2019-01-03 00:00:00',3000.0000,0,0,'d 10104 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(86,1,1,1,'2019-02-01 00:00:00',1000.0000,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(87,5,1,2,'2019-02-01 00:00:00',1000.0000,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(88,7,1,2,'2019-02-01 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(89,12,1,3,'2019-02-01 00:00:00',1500.0000,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(90,14,1,3,'2019-02-01 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(91,19,1,4,'2019-02-01 00:00:00',1500.0000,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(92,21,1,4,'2019-02-01 00:00:00',10.0000,0,0,'d 10105 _, c 10999 _','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0);
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
INSERT INTO `Rentable` VALUES (1,1,0,'Rentable001',0,0,'0000-00-00 00:00:00','2019-02-23 00:42:37',211,'2019-02-23 00:40:52',0,''),(2,1,0,'Rentable002',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(3,1,0,'Rentable003',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(4,1,0,'Rentable004',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(5,1,0,'Rentable005',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(6,1,0,'Rentable006',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(7,1,0,'Rentable007',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(8,1,0,'Rentable008',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(9,1,0,'Rentable009',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(10,1,0,'Rentable010',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(11,1,0,'Rentable011',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(12,1,0,'Rentable012',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(13,1,0,'Rentable013',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(14,1,0,'Rentable014',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(15,1,0,'CP001',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(16,1,0,'CP002',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(17,1,0,'CP003',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(18,1,0,'CP004',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(19,1,0,'CP005',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(20,1,0,'CP006',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(21,1,0,'CP007',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,''),(22,1,0,'CP008',0,0,'0000-00-00 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0,'');
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
  `ConfirmationCode` varchar(20) NOT NULL DEFAULT '',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RLID`)
) ENGINE=InnoDB AUTO_INCREMENT=1080 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableLeaseStatus`
--

LOCK TABLES `RentableLeaseStatus` WRITE;
/*!40000 ALTER TABLE `RentableLeaseStatus` DISABLE KEYS */;
INSERT INTO `RentableLeaseStatus` VALUES (1,1,1,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,2,1,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,3,1,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,4,1,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(5,5,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(6,6,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(7,7,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(8,8,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(9,9,1,0,'2019-01-01 00:00:00','2019-01-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(10,10,1,0,'2019-01-01 00:00:00','2019-01-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(11,11,1,0,'2019-01-01 00:00:00','2019-01-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(12,12,1,0,'2019-01-01 00:00:00','2019-01-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(13,13,1,0,'2019-01-01 00:00:00','2019-01-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(14,14,1,0,'2019-01-01 00:00:00','2019-01-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(15,15,1,0,'2019-01-01 00:00:00','2019-02-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(16,16,1,0,'2019-01-01 00:00:00','2019-01-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(17,17,1,0,'2019-01-01 00:00:00','2019-01-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(18,18,1,0,'2019-01-01 00:00:00','2019-01-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(19,19,1,0,'2019-01-01 00:00:00','2019-01-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(20,20,1,0,'2019-01-01 00:00:00','2019-01-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(21,21,1,0,'2019-01-01 00:00:00','2019-01-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:52',0),(22,22,1,0,'2019-01-01 00:00:00','2019-01-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:52',0),(24,1,1,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(25,1,1,2,'2020-03-01 00:00:00','2020-03-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:42:41',211,'2019-02-23 00:40:52',0),(27,2,1,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(28,2,1,2,'2020-03-01 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(30,3,1,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(31,3,1,2,'2020-03-01 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(33,4,1,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(34,4,1,2,'2020-03-01 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(35,9,1,0,'2019-04-02 00:00:00','2019-04-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(36,9,1,2,'2019-03-30 00:00:00','2019-04-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(37,9,1,0,'2019-08-22 00:00:00','2019-09-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(38,9,1,2,'2019-08-18 00:00:00','2019-08-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(39,9,1,0,'2019-02-06 00:00:00','2019-02-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(40,9,1,2,'2019-01-29 00:00:00','2019-02-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(41,9,1,0,'2019-05-28 00:00:00','2019-06-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(42,9,1,2,'2019-05-24 00:00:00','2019-05-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(43,9,1,0,'2019-01-27 00:00:00','2019-01-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(44,9,1,2,'2019-01-16 00:00:00','2019-01-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(45,9,1,0,'2019-06-26 00:00:00','2019-07-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(46,9,1,2,'2019-06-25 00:00:00','2019-06-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(47,9,1,0,'2019-10-13 00:00:00','2019-10-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(48,9,1,2,'2019-10-08 00:00:00','2019-10-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(49,9,1,0,'2019-12-25 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(50,9,1,2,'2019-12-18 00:00:00','2019-12-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(51,9,1,0,'2019-03-21 00:00:00','2019-03-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(52,9,1,2,'2019-03-15 00:00:00','2019-03-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(53,9,1,0,'2019-08-07 00:00:00','2019-08-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(54,9,1,2,'2019-07-31 00:00:00','2019-08-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(55,9,1,2,'2019-03-20 00:00:00','2019-03-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(56,9,1,2,'2019-03-17 00:00:00','2019-03-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(57,9,1,0,'2019-10-27 00:00:00','2019-11-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(58,9,1,2,'2019-10-26 00:00:00','2019-10-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(59,9,1,0,'2019-04-13 00:00:00','2019-04-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(60,9,1,2,'2019-04-07 00:00:00','2019-04-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(61,9,1,0,'2019-11-20 00:00:00','2019-11-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(62,9,1,2,'2019-11-17 00:00:00','2019-11-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(63,9,1,0,'2019-05-03 00:00:00','2019-05-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(65,9,1,0,'2019-12-11 00:00:00','2019-12-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(66,9,1,2,'2019-12-08 00:00:00','2019-12-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(67,9,1,0,'2019-09-16 00:00:00','2019-09-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(68,9,1,2,'2019-09-13 00:00:00','2019-09-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(69,9,1,2,'2019-04-29 00:00:00','2019-05-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(70,9,1,0,'2019-11-14 00:00:00','2019-11-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(71,9,1,2,'2019-11-08 00:00:00','2019-11-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(72,9,1,0,'2019-02-20 00:00:00','2019-02-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(73,9,1,2,'2019-02-15 00:00:00','2019-02-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(74,9,1,2,'2019-12-21 00:00:00','2019-12-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(75,9,1,0,'2019-06-13 00:00:00','2019-06-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(76,9,1,2,'2019-06-09 00:00:00','2019-06-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(77,9,1,2,'2019-01-20 00:00:00','2019-01-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(78,9,1,2,'2019-11-01 00:00:00','2019-11-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(79,9,1,0,'2019-03-10 00:00:00','2019-03-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(80,9,1,2,'2019-03-08 00:00:00','2019-03-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(81,9,1,2,'2019-06-12 00:00:00','2019-06-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(82,9,1,2,'2019-03-02 00:00:00','2019-03-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(83,9,1,2,'2019-02-01 00:00:00','2019-02-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(84,9,1,0,'2019-09-21 00:00:00','2019-10-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(85,9,1,2,'2019-09-18 00:00:00','2019-09-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(86,9,1,0,'2019-07-03 00:00:00','2019-07-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(87,9,1,2,'2019-07-01 00:00:00','2019-07-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(88,9,1,0,'2019-01-12 00:00:00','2019-01-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(89,9,1,2,'2019-01-09 00:00:00','2019-01-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(90,9,1,2,'2019-10-19 00:00:00','2019-10-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(91,9,1,2,'2019-01-21 00:00:00','2019-01-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(92,9,1,0,'2019-11-30 00:00:00','2019-12-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(93,9,1,2,'2019-11-23 00:00:00','2019-11-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(94,9,1,2,'2019-10-24 00:00:00','2019-10-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(95,9,1,2,'2019-10-22 00:00:00','2019-10-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(96,9,1,0,'2019-07-14 00:00:00','2019-07-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(97,9,1,2,'2019-07-09 00:00:00','2019-07-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(98,9,1,2,'2019-11-11 00:00:00','2019-11-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(99,9,1,0,'2019-03-01 00:00:00','2019-03-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(100,9,1,2,'2019-02-23 00:00:00','2019-03-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(101,9,1,0,'2019-03-24 00:00:00','2019-03-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(102,9,1,2,'2019-03-22 00:00:00','2019-03-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(103,9,1,2,'2019-02-09 00:00:00','2019-02-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(104,9,1,2,'2019-02-03 00:00:00','2019-02-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(105,9,1,2,'2019-04-23 00:00:00','2019-04-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(106,10,1,0,'2019-06-28 00:00:00','2019-06-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(107,10,1,2,'2019-06-23 00:00:00','2019-06-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(108,10,1,0,'2019-03-09 00:00:00','2019-03-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(109,10,1,2,'2019-02-28 00:00:00','2019-03-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(111,10,1,2,'2019-07-07 00:00:00','2019-07-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(112,10,1,0,'2019-01-13 00:00:00','2019-01-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(113,10,1,2,'2019-01-07 00:00:00','2019-01-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(114,10,1,0,'2019-05-12 00:00:00','2019-05-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(115,10,1,2,'2019-05-09 00:00:00','2019-05-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(116,10,1,0,'2019-09-14 00:00:00','2019-09-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(119,10,1,0,'2019-11-30 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(120,10,1,2,'2019-11-26 00:00:00','2019-11-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(121,10,1,0,'2019-06-12 00:00:00','2019-06-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(122,10,1,2,'2019-06-01 00:00:00','2019-06-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(123,10,1,2,'2019-09-12 00:00:00','2019-09-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(124,10,1,0,'2019-11-03 00:00:00','2019-11-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(125,10,1,2,'2019-10-30 00:00:00','2019-11-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(126,10,1,2,'2019-11-28 00:00:00','2019-11-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(127,10,1,0,'2019-02-01 00:00:00','2019-02-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(128,10,1,2,'2019-01-22 00:00:00','2019-01-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(129,10,1,0,'2019-02-09 00:00:00','2019-02-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(130,10,1,2,'2019-02-02 00:00:00','2019-02-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(131,10,1,2,'2019-06-03 00:00:00','2019-06-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(133,10,1,2,'2019-08-04 00:00:00','2019-08-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(134,10,1,0,'2019-03-21 00:00:00','2019-03-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(135,10,1,2,'2019-03-20 00:00:00','2019-03-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(136,10,1,0,'2019-07-21 00:00:00','2019-08-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(137,10,1,2,'2019-07-15 00:00:00','2019-07-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(138,10,1,0,'2019-07-03 00:00:00','2019-07-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(139,10,1,2,'2019-06-30 00:00:00','2019-07-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(140,10,1,0,'2019-11-25 00:00:00','2019-11-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(141,10,1,2,'2019-11-20 00:00:00','2019-11-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(142,10,1,2,'2019-06-11 00:00:00','2019-06-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(143,10,1,2,'2019-01-11 00:00:00','2019-01-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(144,10,1,2,'2019-01-08 00:00:00','2019-01-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(145,10,1,2,'2019-06-06 00:00:00','2019-06-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(146,10,1,0,'2019-08-27 00:00:00','2019-09-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(147,10,1,2,'2019-08-20 00:00:00','2019-08-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(148,10,1,0,'2019-01-21 00:00:00','2019-01-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(149,10,1,2,'2019-01-17 00:00:00','2019-01-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(150,10,1,0,'2019-04-24 00:00:00','2019-05-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(151,10,1,2,'2019-04-19 00:00:00','2019-04-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(152,10,1,2,'2019-07-12 00:00:00','2019-07-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(153,10,1,2,'2019-06-07 00:00:00','2019-06-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(154,10,1,0,'2019-11-13 00:00:00','2019-11-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(155,10,1,2,'2019-11-06 00:00:00','2019-11-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(156,10,1,2,'2019-08-05 00:00:00','2019-08-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(157,10,1,2,'2019-09-09 00:00:00','2019-09-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(158,10,1,2,'2019-01-24 00:00:00','2019-01-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(159,10,1,0,'2019-09-20 00:00:00','2019-10-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(160,10,1,2,'2019-09-17 00:00:00','2019-09-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(161,10,1,0,'2019-08-19 00:00:00','2019-08-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(162,10,1,2,'2019-08-16 00:00:00','2019-08-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(163,10,1,2,'2019-08-09 00:00:00','2019-08-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(164,10,1,2,'2019-08-15 00:00:00','2019-08-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(165,10,1,0,'2019-03-30 00:00:00','2019-04-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(166,10,1,2,'2019-03-24 00:00:00','2019-03-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(167,10,1,2,'2019-07-01 00:00:00','2019-07-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(168,10,1,0,'2019-03-19 00:00:00','2019-03-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(169,10,1,2,'2019-03-17 00:00:00','2019-03-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(170,10,1,2,'2019-07-19 00:00:00','2019-07-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(171,10,1,2,'2019-03-29 00:00:00','2019-03-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(172,10,1,2,'2019-03-25 00:00:00','2019-03-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(173,10,1,2,'2019-01-30 00:00:00','2019-02-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(174,10,1,0,'2019-04-08 00:00:00','2019-04-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(175,10,1,2,'2019-04-07 00:00:00','2019-04-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(176,10,1,0,'2019-05-26 00:00:00','2019-06-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(177,10,1,2,'2019-05-24 00:00:00','2019-05-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(178,10,1,2,'2019-03-06 00:00:00','2019-03-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(179,11,1,0,'2019-10-27 00:00:00','2019-11-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(180,11,1,2,'2019-10-26 00:00:00','2019-10-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(181,11,1,0,'2019-09-08 00:00:00','2019-09-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(182,11,1,2,'2019-09-06 00:00:00','2019-09-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(183,11,1,0,'2019-09-24 00:00:00','2019-09-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(184,11,1,2,'2019-09-20 00:00:00','2019-09-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(185,11,1,0,'2019-09-02 00:00:00','2019-09-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(186,11,1,2,'2019-08-22 00:00:00','2019-08-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(187,11,1,0,'2019-11-03 00:00:00','2019-11-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(188,11,1,2,'2019-11-01 00:00:00','2019-11-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(189,11,1,0,'2019-06-03 00:00:00','2019-06-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(190,11,1,2,'2019-05-31 00:00:00','2019-06-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(191,11,1,0,'2019-07-07 00:00:00','2019-07-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(193,11,1,2,'2019-09-18 00:00:00','2019-09-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(194,11,1,0,'2019-05-30 00:00:00','2019-05-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(195,11,1,2,'2019-05-28 00:00:00','2019-05-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(196,11,1,0,'2019-10-13 00:00:00','2019-10-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(197,11,1,2,'2019-10-08 00:00:00','2019-10-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(198,11,1,0,'2019-09-17 00:00:00','2019-09-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(200,11,1,0,'2019-04-02 00:00:00','2019-04-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(201,11,1,2,'2019-03-24 00:00:00','2019-03-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(202,11,1,2,'2019-08-27 00:00:00','2019-09-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(203,11,1,0,'2019-02-05 00:00:00','2019-02-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(205,11,1,0,'2019-07-15 00:00:00','2019-08-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(206,11,1,2,'2019-07-09 00:00:00','2019-07-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(207,11,1,0,'2019-06-10 00:00:00','2019-06-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(208,11,1,2,'2019-06-05 00:00:00','2019-06-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(209,11,1,2,'2019-10-04 00:00:00','2019-10-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(210,11,1,0,'2019-01-13 00:00:00','2019-01-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(211,11,1,2,'2019-01-06 00:00:00','2019-01-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(212,11,1,0,'2019-06-24 00:00:00','2019-06-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(213,11,1,2,'2019-06-20 00:00:00','2019-06-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(214,11,1,0,'2019-05-15 00:00:00','2019-05-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(215,11,1,2,'2019-05-08 00:00:00','2019-05-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(216,11,1,2,'2019-05-21 00:00:00','2019-05-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(217,11,1,0,'2019-04-22 00:00:00','2019-04-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(218,11,1,2,'2019-04-16 00:00:00','2019-04-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(219,11,1,0,'2019-06-30 00:00:00','2019-07-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(220,11,1,2,'2019-06-29 00:00:00','2019-06-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(221,11,1,0,'2019-02-21 00:00:00','2019-03-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(222,11,1,2,'2019-02-20 00:00:00','2019-02-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(223,11,1,2,'2019-01-31 00:00:00','2019-02-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(224,11,1,2,'2019-07-02 00:00:00','2019-07-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(225,11,1,2,'2019-01-24 00:00:00','2019-01-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(226,11,1,2,'2019-03-30 00:00:00','2019-04-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(227,11,1,2,'2019-06-16 00:00:00','2019-06-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(228,11,1,2,'2019-09-10 00:00:00','2019-09-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(229,11,1,0,'2019-11-13 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(230,11,1,2,'2019-11-07 00:00:00','2019-11-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(231,11,1,2,'2019-03-19 00:00:00','2019-03-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(232,11,1,0,'2019-03-15 00:00:00','2019-03-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(233,11,1,2,'2019-03-11 00:00:00','2019-03-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(234,11,1,2,'2019-01-29 00:00:00','2019-02-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(235,11,1,0,'2019-10-19 00:00:00','2019-10-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(236,11,1,2,'2019-10-16 00:00:00','2019-10-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(237,11,1,0,'2019-04-10 00:00:00','2019-04-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(238,11,1,2,'2019-04-06 00:00:00','2019-04-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(239,11,1,0,'2019-05-07 00:00:00','2019-05-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(240,11,1,2,'2019-04-30 00:00:00','2019-05-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(241,11,1,0,'2019-10-02 00:00:00','2019-10-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(242,11,1,2,'2019-09-28 00:00:00','2019-10-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(243,11,1,2,'2019-04-21 00:00:00','2019-04-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(244,11,1,2,'2019-04-17 00:00:00','2019-04-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(245,12,1,0,'2019-11-19 00:00:00','2019-11-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(247,12,1,0,'2019-11-03 00:00:00','2019-11-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(249,12,1,0,'2019-07-27 00:00:00','2019-08-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(250,12,1,2,'2019-07-22 00:00:00','2019-07-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(251,12,1,0,'2019-12-03 00:00:00','2019-12-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(252,12,1,2,'2019-11-26 00:00:00','2019-12-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(253,12,1,2,'2019-11-11 00:00:00','2019-11-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(254,12,1,0,'2019-04-25 00:00:00','2019-04-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(255,12,1,2,'2019-04-22 00:00:00','2019-04-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(256,12,1,0,'2019-02-16 00:00:00','2019-02-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(257,12,1,2,'2019-02-13 00:00:00','2019-02-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(258,12,1,0,'2019-12-10 00:00:00','2019-12-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(259,12,1,2,'2019-12-05 00:00:00','2019-12-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(260,12,1,0,'2019-02-25 00:00:00','2019-03-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(261,12,1,2,'2019-02-21 00:00:00','2019-02-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(262,12,1,0,'2019-08-15 00:00:00','2019-08-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(263,12,1,2,'2019-08-12 00:00:00','2019-08-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(264,12,1,0,'2019-03-05 00:00:00','2019-03-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(265,12,1,2,'2019-03-04 00:00:00','2019-03-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(266,12,1,2,'2019-11-17 00:00:00','2019-11-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(267,12,1,0,'2019-05-02 00:00:00','2019-05-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(268,12,1,2,'2019-04-27 00:00:00','2019-05-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(269,12,1,2,'2019-10-26 00:00:00','2019-10-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(270,12,1,2,'2019-10-27 00:00:00','2019-11-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(271,12,1,2,'2019-11-14 00:00:00','2019-11-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(272,12,1,2,'2019-10-25 00:00:00','2019-10-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(273,12,1,0,'2019-03-26 00:00:00','2019-03-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(274,12,1,2,'2019-03-23 00:00:00','2019-03-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(275,12,1,0,'2019-08-23 00:00:00','2019-08-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(276,12,1,2,'2019-08-19 00:00:00','2019-08-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(277,12,1,0,'2019-06-26 00:00:00','2019-06-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(279,12,1,0,'2019-03-02 00:00:00','2019-03-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(280,12,1,2,'2019-03-01 00:00:00','2019-03-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(281,12,1,0,'2019-10-16 00:00:00','2019-10-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(282,12,1,2,'2019-10-08 00:00:00','2019-10-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(283,12,1,2,'2019-02-18 00:00:00','2019-02-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(284,12,1,2,'2019-04-24 00:00:00','2019-04-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(285,12,1,2,'2019-10-13 00:00:00','2019-10-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(286,12,1,0,'2019-03-22 00:00:00','2019-03-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(287,12,1,2,'2019-03-19 00:00:00','2019-03-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(288,12,1,0,'2019-07-05 00:00:00','2019-07-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(289,12,1,2,'2019-06-27 00:00:00','2019-07-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(290,12,1,0,'2019-05-07 00:00:00','2019-05-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(291,12,1,2,'2019-05-06 00:00:00','2019-05-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(292,12,1,0,'2019-05-31 00:00:00','2019-06-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(293,12,1,2,'2019-05-24 00:00:00','2019-05-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(294,12,1,0,'2019-09-23 00:00:00','2019-10-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(295,12,1,2,'2019-09-20 00:00:00','2019-09-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(296,12,1,0,'2019-05-16 00:00:00','2019-05-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(297,12,1,2,'2019-05-14 00:00:00','2019-05-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(298,12,1,0,'2019-08-27 00:00:00','2019-08-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(299,12,1,2,'2019-08-26 00:00:00','2019-08-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(300,12,1,0,'2019-01-24 00:00:00','2019-02-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(301,12,1,2,'2019-01-22 00:00:00','2019-01-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(302,12,1,2,'2019-12-02 00:00:00','2019-12-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(303,12,1,2,'2019-10-18 00:00:00','2019-10-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(304,12,1,2,'2019-10-23 00:00:00','2019-10-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(305,12,1,2,'2019-10-21 00:00:00','2019-10-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(306,12,1,2,'2019-06-21 00:00:00','2019-06-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(307,12,1,2,'2019-07-02 00:00:00','2019-07-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(308,12,1,0,'2019-09-02 00:00:00','2019-09-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(309,12,1,2,'2019-08-29 00:00:00','2019-09-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(310,12,1,0,'2019-07-15 00:00:00','2019-07-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(311,12,1,2,'2019-07-11 00:00:00','2019-07-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(312,12,1,2,'2019-01-17 00:00:00','2019-01-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(313,12,1,0,'2019-03-29 00:00:00','2019-04-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(314,12,1,2,'2019-03-28 00:00:00','2019-03-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(315,12,1,0,'2019-06-14 00:00:00','2019-06-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(316,12,1,2,'2019-06-13 00:00:00','2019-06-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(317,12,1,2,'2019-05-21 00:00:00','2019-05-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(318,12,1,0,'2019-12-23 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(319,12,1,2,'2019-12-15 00:00:00','2019-12-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(320,12,1,2,'2019-12-16 00:00:00','2019-12-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(321,12,1,0,'2019-06-10 00:00:00','2019-06-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(322,12,1,2,'2019-06-06 00:00:00','2019-06-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(323,12,1,2,'2019-08-11 00:00:00','2019-08-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(324,13,1,0,'2019-02-02 00:00:00','2019-02-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(326,13,1,0,'2019-06-01 00:00:00','2019-06-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(327,13,1,2,'2019-05-30 00:00:00','2019-06-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(328,13,1,2,'2019-01-25 00:00:00','2019-01-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(329,13,1,0,'2019-06-11 00:00:00','2019-06-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(330,13,1,2,'2019-06-10 00:00:00','2019-06-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(331,13,1,0,'2019-12-01 00:00:00','2019-12-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(333,13,1,0,'2019-08-30 00:00:00','2019-09-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(334,13,1,2,'2019-08-29 00:00:00','2019-08-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(335,13,1,0,'2019-11-08 00:00:00','2019-11-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(336,13,1,2,'2019-10-29 00:00:00','2019-11-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(337,13,1,0,'2019-07-12 00:00:00','2019-08-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(338,13,1,2,'2019-07-10 00:00:00','2019-07-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(339,13,1,0,'2019-08-04 00:00:00','2019-08-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(341,13,1,0,'2019-12-14 00:00:00','2019-12-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(342,13,1,2,'2019-12-05 00:00:00','2019-12-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(343,13,1,0,'2019-08-19 00:00:00','2019-08-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(345,13,1,0,'2019-06-25 00:00:00','2019-06-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(346,13,1,2,'2019-06-17 00:00:00','2019-06-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(347,13,1,0,'2019-04-20 00:00:00','2019-04-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(348,13,1,2,'2019-04-19 00:00:00','2019-04-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(349,13,1,2,'2019-11-02 00:00:00','2019-11-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(350,13,1,2,'2019-11-01 00:00:00','2019-11-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(351,13,1,0,'2019-07-02 00:00:00','2019-07-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(352,13,1,2,'2019-06-30 00:00:00','2019-07-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(353,13,1,2,'2019-08-03 00:00:00','2019-08-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(354,13,1,0,'2019-03-07 00:00:00','2019-03-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(355,13,1,2,'2019-02-28 00:00:00','2019-03-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(356,13,1,2,'2019-08-14 00:00:00','2019-08-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(357,13,1,2,'2019-01-30 00:00:00','2019-02-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(358,13,1,2,'2019-01-27 00:00:00','2019-02-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(359,13,1,0,'2020-01-01 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(360,13,1,2,'2019-12-31 00:00:00','2020-01-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(361,13,1,0,'2019-01-10 00:00:00','2019-01-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(362,13,1,2,'2019-01-06 00:00:00','2019-01-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(363,13,1,2,'2019-08-15 00:00:00','2019-08-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(364,13,1,2,'2019-07-03 00:00:00','2019-07-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(365,13,1,0,'2019-01-18 00:00:00','2019-01-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(366,13,1,2,'2019-01-17 00:00:00','2019-01-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(368,13,1,0,'2019-08-23 00:00:00','2019-08-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(369,13,1,2,'2019-08-22 00:00:00','2019-08-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(370,13,1,0,'2019-05-29 00:00:00','2019-05-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(371,13,1,2,'2019-05-22 00:00:00','2019-05-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(373,13,1,0,'2019-04-02 00:00:00','2019-04-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(374,13,1,2,'2019-03-26 00:00:00','2019-04-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(375,13,1,0,'2019-04-15 00:00:00','2019-04-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(376,13,1,2,'2019-04-12 00:00:00','2019-04-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(377,13,1,0,'2019-11-18 00:00:00','2019-11-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(378,13,1,2,'2019-11-15 00:00:00','2019-11-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(379,13,1,2,'2019-11-04 00:00:00','2019-11-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(380,13,1,2,'2019-08-08 00:00:00','2019-08-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(381,13,1,2,'2019-01-13 00:00:00','2019-01-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(382,13,1,2,'2019-12-07 00:00:00','2019-12-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(383,13,1,0,'2019-04-05 00:00:00','2019-04-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(384,13,1,2,'2019-04-03 00:00:00','2019-04-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(385,13,1,2,'2019-11-29 00:00:00','2019-12-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(386,13,1,0,'2019-03-15 00:00:00','2019-03-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(387,13,1,2,'2019-03-11 00:00:00','2019-03-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(388,13,1,2,'2019-06-05 00:00:00','2019-06-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(389,13,1,2,'2019-06-22 00:00:00','2019-06-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(390,13,1,0,'2019-12-25 00:00:00','2019-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(391,13,1,2,'2019-12-23 00:00:00','2019-12-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(392,13,1,0,'2019-09-26 00:00:00','2019-10-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(393,13,1,2,'2019-09-22 00:00:00','2019-09-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(394,13,1,0,'2019-10-08 00:00:00','2019-10-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(395,13,1,2,'2019-10-07 00:00:00','2019-10-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(396,13,1,2,'2019-10-03 00:00:00','2019-10-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(397,13,1,0,'2019-09-03 00:00:00','2019-09-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(398,13,1,2,'2019-09-02 00:00:00','2019-09-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(399,13,1,0,'2019-05-03 00:00:00','2019-05-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(400,13,1,2,'2019-04-30 00:00:00','2019-05-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(401,13,1,2,'2019-06-18 00:00:00','2019-06-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(402,14,1,0,'2019-12-16 00:00:00','2019-12-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(403,14,1,2,'2019-12-12 00:00:00','2019-12-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(404,14,1,0,'2019-03-28 00:00:00','2019-04-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(405,14,1,2,'2019-03-21 00:00:00','2019-03-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(406,14,1,0,'2019-07-13 00:00:00','2019-07-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(407,14,1,2,'2019-07-07 00:00:00','2019-07-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(408,14,1,0,'2019-11-28 00:00:00','2019-12-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(410,14,1,0,'2019-10-13 00:00:00','2019-10-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(411,14,1,2,'2019-10-08 00:00:00','2019-10-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(412,14,1,0,'2019-03-04 00:00:00','2019-03-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(413,14,1,2,'2019-02-18 00:00:00','2019-02-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(414,14,1,0,'2019-10-06 00:00:00','2019-10-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(415,14,1,2,'2019-10-03 00:00:00','2019-10-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(416,14,1,0,'2019-09-08 00:00:00','2019-10-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(417,14,1,2,'2019-09-01 00:00:00','2019-09-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(418,14,1,0,'2019-11-11 00:00:00','2019-11-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(419,14,1,2,'2019-11-04 00:00:00','2019-11-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(420,14,1,0,'2019-06-24 00:00:00','2019-06-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(421,14,1,2,'2019-06-22 00:00:00','2019-06-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(422,14,1,0,'2019-04-18 00:00:00','2019-04-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(423,14,1,2,'2019-04-13 00:00:00','2019-04-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(424,14,1,0,'2019-05-05 00:00:00','2019-05-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(425,14,1,2,'2019-04-29 00:00:00','2019-05-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(426,14,1,0,'2019-04-06 00:00:00','2019-04-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(427,14,1,2,'2019-04-05 00:00:00','2019-04-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(428,14,1,2,'2019-11-15 00:00:00','2019-11-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(429,14,1,0,'2019-01-09 00:00:00','2019-01-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(430,14,1,2,'2019-01-05 00:00:00','2019-01-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(431,14,1,0,'2019-07-04 00:00:00','2019-07-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(432,14,1,2,'2019-06-30 00:00:00','2019-07-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(433,14,1,2,'2019-11-18 00:00:00','2019-11-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(436,14,1,0,'2019-01-19 00:00:00','2019-01-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(437,14,1,2,'2019-01-13 00:00:00','2019-01-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(438,14,1,0,'2019-05-28 00:00:00','2019-06-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(440,14,1,2,'2019-11-10 00:00:00','2019-11-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(441,14,1,0,'2019-06-13 00:00:00','2019-06-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(443,14,1,2,'2019-11-19 00:00:00','2019-11-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(444,14,1,2,'2019-01-14 00:00:00','2019-01-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(445,14,1,2,'2019-02-19 00:00:00','2019-02-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(446,14,1,2,'2019-09-02 00:00:00','2019-09-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(447,14,1,0,'2019-01-26 00:00:00','2019-02-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(448,14,1,2,'2019-01-20 00:00:00','2019-01-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(449,14,1,2,'2019-06-06 00:00:00','2019-06-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(450,14,1,2,'2019-05-16 00:00:00','2019-05-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(451,14,1,0,'2019-04-28 00:00:00','2019-04-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(452,14,1,2,'2019-04-21 00:00:00','2019-04-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(453,14,1,0,'2019-10-24 00:00:00','2019-11-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(454,14,1,2,'2019-10-19 00:00:00','2019-10-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(455,14,1,0,'2019-12-27 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(456,14,1,2,'2019-12-22 00:00:00','2019-12-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(457,14,1,2,'2019-03-17 00:00:00','2019-03-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(458,14,1,0,'2019-02-05 00:00:00','2019-02-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(459,14,1,2,'2019-02-03 00:00:00','2019-02-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(461,14,1,2,'2019-02-26 00:00:00','2019-03-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(462,14,1,0,'2019-07-24 00:00:00','2019-09-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(463,14,1,2,'2019-07-20 00:00:00','2019-07-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(464,14,1,0,'2019-06-21 00:00:00','2019-06-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(465,14,1,2,'2019-06-16 00:00:00','2019-06-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(466,14,1,2,'2019-05-21 00:00:00','2019-05-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(467,14,1,0,'2019-02-17 00:00:00','2019-02-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(468,14,1,2,'2019-02-13 00:00:00','2019-02-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(469,14,1,2,'2019-11-22 00:00:00','2019-11-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(470,15,1,0,'2019-02-26 00:00:00','2019-02-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(471,15,1,2,'2019-02-22 00:00:00','2019-02-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(472,15,1,0,'2019-06-05 00:00:00','2019-06-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(473,15,1,2,'2019-06-02 00:00:00','2019-06-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(474,15,1,0,'2019-05-23 00:00:00','2019-05-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(475,15,1,2,'2019-05-22 00:00:00','2019-05-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(476,15,1,0,'2019-04-28 00:00:00','2019-05-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(477,15,1,2,'2019-04-22 00:00:00','2019-04-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(478,15,1,0,'2019-05-09 00:00:00','2019-05-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(481,15,1,2,'2019-07-02 00:00:00','2019-07-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(482,15,1,0,'2019-08-07 00:00:00','2019-08-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(483,15,1,2,'2019-07-31 00:00:00','2019-08-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(484,15,1,0,'2019-11-28 00:00:00','2019-12-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(485,15,1,2,'2019-11-27 00:00:00','2019-11-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(486,15,1,0,'2019-03-07 00:00:00','2019-03-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(487,15,1,2,'2019-02-28 00:00:00','2019-03-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(488,15,1,0,'2019-10-27 00:00:00','2019-11-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(489,15,1,2,'2019-10-18 00:00:00','2019-10-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(490,15,1,0,'2019-04-01 00:00:00','2019-04-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(491,15,1,2,'2019-03-26 00:00:00','2019-04-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(492,15,1,0,'2019-12-06 00:00:00','2019-12-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(493,15,1,2,'2019-12-05 00:00:00','2019-12-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(494,15,1,0,'2019-08-24 00:00:00','2019-09-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(495,15,1,2,'2019-08-21 00:00:00','2019-08-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(496,15,1,2,'2019-05-29 00:00:00','2019-06-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(497,15,1,0,'2019-07-16 00:00:00','2019-07-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(498,15,1,2,'2019-07-14 00:00:00','2019-07-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(500,15,1,0,'2019-09-30 00:00:00','2019-10-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(501,15,1,2,'2019-09-22 00:00:00','2019-09-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(502,15,1,0,'2019-06-16 00:00:00','2019-07-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(503,15,1,2,'2019-06-13 00:00:00','2019-06-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(504,15,1,2,'2019-09-24 00:00:00','2019-09-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(505,15,1,0,'2019-02-20 00:00:00','2019-02-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(506,15,1,2,'2019-02-14 00:00:00','2019-02-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(509,15,1,2,'2019-07-05 00:00:00','2019-07-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(510,15,1,2,'2019-04-27 00:00:00','2019-04-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(511,15,1,2,'2019-04-25 00:00:00','2019-04-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(512,15,1,0,'2019-09-14 00:00:00','2019-09-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(513,15,1,2,'2019-09-11 00:00:00','2019-09-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(514,15,1,0,'2019-03-14 00:00:00','2019-03-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(515,15,1,2,'2019-03-11 00:00:00','2019-03-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(516,15,1,0,'2019-10-16 00:00:00','2019-10-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(518,15,1,0,'2019-09-09 00:00:00','2019-09-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(519,15,1,2,'2019-09-02 00:00:00','2019-09-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(520,15,1,2,'2019-07-12 00:00:00','2019-07-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(521,15,1,2,'2019-10-20 00:00:00','2019-10-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(522,15,1,2,'2019-05-01 00:00:00','2019-05-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(523,15,1,2,'2019-07-06 00:00:00','2019-07-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(524,15,1,2,'2019-04-24 00:00:00','2019-04-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(525,15,1,2,'2019-04-23 00:00:00','2019-04-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(526,15,1,0,'2019-04-21 00:00:00','2019-04-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(527,15,1,2,'2019-04-20 00:00:00','2019-04-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(528,15,1,0,'2019-12-14 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(529,15,1,2,'2019-12-09 00:00:00','2019-12-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(530,15,1,0,'2019-02-10 00:00:00','2019-02-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(531,15,1,2,'2019-02-09 00:00:00','2019-02-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(532,15,1,2,'2019-11-22 00:00:00','2019-11-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(533,15,1,0,'2019-03-21 00:00:00','2019-03-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(534,15,1,2,'2019-03-16 00:00:00','2019-03-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(535,15,1,0,'2019-11-08 00:00:00','2019-11-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(536,15,1,2,'2019-11-02 00:00:00','2019-11-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(537,15,1,2,'2019-10-15 00:00:00','2019-10-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(538,15,1,2,'2019-09-27 00:00:00','2019-09-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(539,15,1,2,'2019-09-25 00:00:00','2019-09-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(540,15,1,0,'2019-09-19 00:00:00','2019-09-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(541,15,1,2,'2019-09-15 00:00:00','2019-09-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(542,15,1,2,'2019-05-03 00:00:00','2019-05-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(543,15,1,2,'2019-06-09 00:00:00','2019-06-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(544,16,1,0,'2019-09-13 00:00:00','2019-09-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(545,16,1,2,'2019-09-11 00:00:00','2019-09-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(546,16,1,0,'2019-11-07 00:00:00','2019-11-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(547,16,1,2,'2019-11-06 00:00:00','2019-11-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(548,16,1,0,'2019-07-04 00:00:00','2019-07-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(549,16,1,2,'2019-07-01 00:00:00','2019-07-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(551,16,1,2,'2019-01-15 00:00:00','2019-01-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(552,16,1,0,'2019-04-28 00:00:00','2019-05-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(553,16,1,2,'2019-04-27 00:00:00','2019-04-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(554,16,1,0,'2019-02-08 00:00:00','2019-02-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(555,16,1,2,'2019-02-06 00:00:00','2019-02-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(556,16,1,0,'2019-04-01 00:00:00','2019-04-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(557,16,1,2,'2019-03-25 00:00:00','2019-04-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(558,16,1,2,'2019-01-16 00:00:00','2019-01-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(559,16,1,0,'2019-12-29 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(560,16,1,2,'2019-12-20 00:00:00','2019-12-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(561,16,1,0,'2019-09-04 00:00:00','2019-09-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(562,16,1,2,'2019-08-28 00:00:00','2019-09-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(563,16,1,0,'2019-04-15 00:00:00','2019-04-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(564,16,1,2,'2019-04-06 00:00:00','2019-04-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(565,16,1,2,'2019-11-02 00:00:00','2019-11-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(566,16,1,2,'2019-12-26 00:00:00','2019-12-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(567,16,1,0,'2019-09-10 00:00:00','2019-09-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(568,16,1,2,'2019-09-06 00:00:00','2019-09-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(569,16,1,0,'2019-09-25 00:00:00','2019-10-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(570,16,1,2,'2019-09-21 00:00:00','2019-09-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(571,16,1,0,'2019-02-27 00:00:00','2019-03-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(572,16,1,2,'2019-02-22 00:00:00','2019-02-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(573,16,1,0,'2019-08-19 00:00:00','2019-08-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(574,16,1,2,'2019-08-14 00:00:00','2019-08-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(575,16,1,2,'2019-04-10 00:00:00','2019-04-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(576,16,1,0,'2019-07-21 00:00:00','2019-08-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(577,16,1,2,'2019-07-17 00:00:00','2019-07-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(578,16,1,0,'2019-11-16 00:00:00','2019-11-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(579,16,1,2,'2019-11-09 00:00:00','2019-11-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(580,16,1,2,'2019-09-18 00:00:00','2019-09-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(581,16,1,0,'2019-05-28 00:00:00','2019-06-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(582,16,1,2,'2019-05-23 00:00:00','2019-05-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(583,16,1,0,'2019-03-07 00:00:00','2019-03-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(585,16,1,0,'2019-05-21 00:00:00','2019-05-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(586,16,1,2,'2019-05-17 00:00:00','2019-05-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(587,16,1,0,'2019-04-17 00:00:00','2019-04-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(589,16,1,0,'2019-02-01 00:00:00','2019-02-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(591,16,1,2,'2019-12-22 00:00:00','2019-12-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(592,16,1,0,'2019-11-23 00:00:00','2019-12-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(593,16,1,2,'2019-11-19 00:00:00','2019-11-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(594,16,1,0,'2019-01-08 00:00:00','2019-01-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(595,16,1,2,'2019-01-04 00:00:00','2019-01-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(596,16,1,2,'2019-04-16 00:00:00','2019-04-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(597,16,1,0,'2019-06-27 00:00:00','2019-07-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(598,16,1,2,'2019-06-18 00:00:00','2019-06-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(599,16,1,2,'2019-06-20 00:00:00','2019-06-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(600,16,1,2,'2019-01-27 00:00:00','2019-02-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(601,16,1,0,'2019-05-03 00:00:00','2019-05-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(602,16,1,2,'2019-05-02 00:00:00','2019-05-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(603,16,1,0,'2019-03-15 00:00:00','2019-03-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(604,16,1,2,'2019-03-08 00:00:00','2019-03-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(605,16,1,0,'2019-10-19 00:00:00','2019-11-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(606,16,1,2,'2019-10-16 00:00:00','2019-10-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(607,16,1,0,'2019-05-10 00:00:00','2019-05-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(608,16,1,2,'2019-05-09 00:00:00','2019-05-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(609,16,1,0,'2019-02-21 00:00:00','2019-02-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(610,16,1,2,'2019-02-18 00:00:00','2019-02-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(611,16,1,2,'2019-01-20 00:00:00','2019-01-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(612,16,1,2,'2019-02-17 00:00:00','2019-02-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(613,16,1,2,'2019-03-01 00:00:00','2019-03-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(614,16,1,2,'2019-03-10 00:00:00','2019-03-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(615,17,1,0,'2019-04-08 00:00:00','2019-04-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(616,17,1,2,'2019-03-29 00:00:00','2019-04-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(617,17,1,0,'2019-05-17 00:00:00','2019-05-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(618,17,1,2,'2019-05-16 00:00:00','2019-05-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(621,17,1,0,'2019-12-21 00:00:00','2019-12-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(622,17,1,2,'2019-12-18 00:00:00','2019-12-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(623,17,1,0,'2019-06-25 00:00:00','2019-06-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(624,17,1,2,'2019-06-19 00:00:00','2019-06-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(625,17,1,2,'2019-12-13 00:00:00','2019-12-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(626,17,1,0,'2019-09-27 00:00:00','2019-10-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(627,17,1,2,'2019-09-26 00:00:00','2019-09-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(628,17,1,0,'2020-01-01 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(629,17,1,2,'2019-12-31 00:00:00','2020-01-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(630,17,1,0,'2019-09-16 00:00:00','2019-09-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(631,17,1,2,'2019-09-10 00:00:00','2019-09-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(633,17,1,0,'2019-06-11 00:00:00','2019-06-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(634,17,1,2,'2019-06-03 00:00:00','2019-06-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(635,17,1,0,'2019-04-20 00:00:00','2019-05-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(636,17,1,2,'2019-04-16 00:00:00','2019-04-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(637,17,1,0,'2019-03-28 00:00:00','2019-03-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(638,17,1,2,'2019-03-25 00:00:00','2019-03-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(639,17,1,0,'2019-09-23 00:00:00','2019-09-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(640,17,1,2,'2019-09-20 00:00:00','2019-09-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(641,17,1,0,'2019-02-08 00:00:00','2019-02-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(642,17,1,2,'2019-02-01 00:00:00','2019-02-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(643,17,1,0,'2019-07-04 00:00:00','2019-07-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(644,17,1,2,'2019-06-30 00:00:00','2019-07-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(645,17,1,2,'2019-02-04 00:00:00','2019-02-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(646,17,1,0,'2019-10-18 00:00:00','2019-10-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(647,17,1,2,'2019-10-14 00:00:00','2019-10-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(648,17,1,0,'2019-08-11 00:00:00','2019-09-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(650,17,1,0,'2019-05-30 00:00:00','2019-06-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(651,17,1,2,'2019-05-22 00:00:00','2019-05-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(652,17,1,2,'2019-06-09 00:00:00','2019-06-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(653,17,1,2,'2019-05-14 00:00:00','2019-05-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(654,17,1,2,'2019-04-04 00:00:00','2019-04-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(655,17,1,2,'2019-10-09 00:00:00','2019-10-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(656,17,1,0,'2019-01-25 00:00:00','2019-01-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(657,17,1,2,'2019-01-19 00:00:00','2019-01-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(658,17,1,0,'2019-07-08 00:00:00','2019-08-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(659,17,1,2,'2019-07-05 00:00:00','2019-07-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(660,17,1,2,'2019-10-10 00:00:00','2019-10-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(662,17,1,2,'2019-12-14 00:00:00','2019-12-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(663,17,1,2,'2019-12-15 00:00:00','2019-12-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(664,17,1,2,'2019-10-01 00:00:00','2019-10-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(665,17,1,2,'2019-12-24 00:00:00','2019-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(666,17,1,0,'2019-11-06 00:00:00','2019-11-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(667,17,1,2,'2019-11-04 00:00:00','2019-11-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(668,17,1,2,'2019-04-17 00:00:00','2019-04-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(669,17,1,2,'2019-08-09 00:00:00','2019-08-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(670,17,1,0,'2019-11-23 00:00:00','2019-12-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(671,17,1,2,'2019-11-16 00:00:00','2019-11-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(672,17,1,2,'2019-10-04 00:00:00','2019-10-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(673,17,1,0,'2019-05-21 00:00:00','2019-05-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(674,17,1,2,'2019-05-18 00:00:00','2019-05-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(675,17,1,2,'2019-10-30 00:00:00','2019-11-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(676,17,1,0,'2019-06-28 00:00:00','2019-06-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(677,17,1,2,'2019-06-26 00:00:00','2019-06-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(678,17,1,0,'2019-01-16 00:00:00','2019-01-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(679,17,1,2,'2019-01-11 00:00:00','2019-01-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(680,17,1,0,'2019-12-06 00:00:00','2019-12-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(681,17,1,2,'2019-12-02 00:00:00','2019-12-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(682,17,1,0,'2019-02-21 00:00:00','2019-03-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(683,17,1,2,'2019-02-16 00:00:00','2019-02-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(684,17,1,0,'2019-01-31 00:00:00','2019-02-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(685,17,1,2,'2019-01-28 00:00:00','2019-01-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(686,17,1,2,'2019-05-25 00:00:00','2019-05-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(687,18,1,0,'2019-04-18 00:00:00','2019-04-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(688,18,1,2,'2019-04-10 00:00:00','2019-04-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(689,18,1,0,'2019-06-23 00:00:00','2019-06-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(690,18,1,2,'2019-06-16 00:00:00','2019-06-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(691,18,1,0,'2019-04-01 00:00:00','2019-04-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(692,18,1,2,'2019-03-25 00:00:00','2019-03-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(693,18,1,0,'2020-01-01 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(694,18,1,2,'2019-12-29 00:00:00','2020-01-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(695,18,1,0,'2019-07-09 00:00:00','2019-07-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(696,18,1,2,'2019-07-03 00:00:00','2019-07-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(697,18,1,0,'2019-02-04 00:00:00','2019-02-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(698,18,1,2,'2019-01-30 00:00:00','2019-02-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(699,18,1,0,'2019-07-17 00:00:00','2019-07-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(700,18,1,2,'2019-07-16 00:00:00','2019-07-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(701,18,1,0,'2019-04-23 00:00:00','2019-04-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(702,18,1,2,'2019-04-19 00:00:00','2019-04-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(703,18,1,2,'2019-01-28 00:00:00','2019-01-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(704,18,1,0,'2019-11-16 00:00:00','2019-12-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(705,18,1,2,'2019-11-14 00:00:00','2019-11-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(706,18,1,0,'2019-01-22 00:00:00','2019-01-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(708,18,1,0,'2019-12-20 00:00:00','2019-12-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(709,18,1,2,'2019-12-19 00:00:00','2019-12-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(710,18,1,0,'2019-08-26 00:00:00','2019-09-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(711,18,1,2,'2019-08-21 00:00:00','2019-08-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(712,18,1,0,'2019-09-13 00:00:00','2019-09-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(713,18,1,2,'2019-09-10 00:00:00','2019-09-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(714,18,1,2,'2019-01-10 00:00:00','2019-01-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(715,18,1,2,'2019-06-27 00:00:00','2019-07-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(716,18,1,0,'2019-05-05 00:00:00','2019-05-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(717,18,1,2,'2019-04-30 00:00:00','2019-05-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(718,18,1,0,'2019-03-06 00:00:00','2019-03-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(719,18,1,2,'2019-02-27 00:00:00','2019-03-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(720,18,1,0,'2019-05-29 00:00:00','2019-06-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(721,18,1,2,'2019-05-17 00:00:00','2019-05-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(722,18,1,2,'2019-04-16 00:00:00','2019-04-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(723,18,1,0,'2019-03-22 00:00:00','2019-03-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(724,18,1,2,'2019-03-17 00:00:00','2019-03-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(725,18,1,0,'2019-06-06 00:00:00','2019-06-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(726,18,1,2,'2019-06-02 00:00:00','2019-06-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(727,18,1,2,'2019-01-15 00:00:00','2019-01-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(728,18,1,2,'2019-07-10 00:00:00','2019-07-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(729,18,1,2,'2019-01-21 00:00:00','2019-01-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(730,18,1,2,'2019-01-19 00:00:00','2019-01-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(731,18,1,0,'2019-06-14 00:00:00','2019-06-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(732,18,1,2,'2019-06-12 00:00:00','2019-06-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(733,18,1,2,'2019-11-07 00:00:00','2019-11-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(734,18,1,0,'2019-10-24 00:00:00','2019-11-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(735,18,1,2,'2019-10-20 00:00:00','2019-10-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(736,18,1,2,'2019-05-23 00:00:00','2019-05-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(737,18,1,0,'2019-07-27 00:00:00','2019-08-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(738,18,1,2,'2019-07-24 00:00:00','2019-07-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(739,18,1,2,'2019-06-22 00:00:00','2019-06-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(741,18,1,2,'2019-12-22 00:00:00','2019-12-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(742,18,1,0,'2019-02-24 00:00:00','2019-02-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(743,18,1,2,'2019-02-19 00:00:00','2019-02-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(744,18,1,2,'2019-02-20 00:00:00','2019-02-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(745,18,1,2,'2019-04-28 00:00:00','2019-04-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(746,18,1,0,'2019-10-06 00:00:00','2019-10-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(747,18,1,2,'2019-09-29 00:00:00','2019-10-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(748,18,1,2,'2019-12-23 00:00:00','2019-12-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(749,18,1,2,'2019-11-13 00:00:00','2019-11-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(750,18,1,2,'2019-11-09 00:00:00','2019-11-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(751,18,1,2,'2019-08-25 00:00:00','2019-08-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(752,18,1,2,'2019-11-11 00:00:00','2019-11-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(753,18,1,0,'2019-09-04 00:00:00','2019-09-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(754,18,1,2,'2019-09-03 00:00:00','2019-09-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(755,18,1,2,'2019-01-24 00:00:00','2019-01-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(756,18,1,2,'2019-03-27 00:00:00','2019-04-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(757,19,1,0,'2019-06-01 00:00:00','2019-06-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(758,19,1,2,'2019-05-31 00:00:00','2019-06-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(759,19,1,0,'2019-03-17 00:00:00','2019-03-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(760,19,1,2,'2019-03-16 00:00:00','2019-03-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(761,19,1,0,'2019-10-21 00:00:00','2019-10-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(762,19,1,2,'2019-10-18 00:00:00','2019-10-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(763,19,1,0,'2019-06-24 00:00:00','2019-06-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(764,19,1,2,'2019-06-19 00:00:00','2019-06-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(765,19,1,0,'2019-04-04 00:00:00','2019-04-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(766,19,1,2,'2019-03-19 00:00:00','2019-03-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(767,19,1,2,'2019-03-23 00:00:00','2019-03-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(768,19,1,0,'2019-11-21 00:00:00','2019-11-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(770,19,1,0,'2019-02-15 00:00:00','2019-02-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(771,19,1,2,'2019-02-08 00:00:00','2019-02-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(772,19,1,0,'2019-12-08 00:00:00','2019-12-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(773,19,1,2,'2019-12-07 00:00:00','2019-12-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(774,19,1,0,'2019-01-21 00:00:00','2019-01-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(775,19,1,2,'2019-01-19 00:00:00','2019-01-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(776,19,1,2,'2019-03-29 00:00:00','2019-04-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(778,19,1,0,'2019-01-14 00:00:00','2019-01-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(779,19,1,2,'2019-01-04 00:00:00','2019-01-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(780,19,1,0,'2019-08-30 00:00:00','2019-09-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(781,19,1,2,'2019-08-25 00:00:00','2019-08-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(782,19,1,2,'2019-01-07 00:00:00','2019-01-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(783,19,1,0,'2019-02-20 00:00:00','2019-02-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(784,19,1,2,'2019-02-17 00:00:00','2019-02-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(785,19,1,0,'2019-07-04 00:00:00','2019-07-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(786,19,1,2,'2019-07-03 00:00:00','2019-07-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(787,19,1,0,'2019-02-28 00:00:00','2019-03-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(788,19,1,2,'2019-02-21 00:00:00','2019-02-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(789,19,1,2,'2019-03-04 00:00:00','2019-03-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(790,19,1,0,'2019-11-13 00:00:00','2019-11-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(791,19,1,2,'2019-11-09 00:00:00','2019-11-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(792,19,1,2,'2019-06-29 00:00:00','2019-07-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(793,19,1,0,'2019-09-23 00:00:00','2019-09-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(794,19,1,2,'2019-09-18 00:00:00','2019-09-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(795,19,1,0,'2019-10-13 00:00:00','2019-10-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(796,19,1,2,'2019-10-09 00:00:00','2019-10-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(797,19,1,0,'2019-06-05 00:00:00','2019-06-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(798,19,1,2,'2019-06-04 00:00:00','2019-06-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(800,19,1,0,'2019-06-11 00:00:00','2019-06-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(801,19,1,2,'2019-06-10 00:00:00','2019-06-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(803,19,1,0,'2020-01-01 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(804,19,1,2,'2019-12-29 00:00:00','2020-01-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(805,19,1,0,'2019-11-30 00:00:00','2019-12-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(807,19,1,0,'2019-08-05 00:00:00','2019-08-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(808,19,1,2,'2019-08-04 00:00:00','2019-08-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(809,19,1,2,'2019-11-23 00:00:00','2019-11-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(810,19,1,0,'2019-07-21 00:00:00','2019-08-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(811,19,1,2,'2019-07-12 00:00:00','2019-07-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(812,19,1,2,'2019-06-27 00:00:00','2019-06-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(813,19,1,0,'2019-04-20 00:00:00','2019-05-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(814,19,1,2,'2019-04-16 00:00:00','2019-04-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(815,19,1,0,'2019-10-08 00:00:00','2019-10-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(816,19,1,2,'2019-10-06 00:00:00','2019-10-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(817,19,1,0,'2019-10-28 00:00:00','2019-11-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(818,19,1,2,'2019-10-27 00:00:00','2019-10-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(819,19,1,0,'2019-10-01 00:00:00','2019-10-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(820,19,1,2,'2019-09-24 00:00:00','2019-10-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(821,19,1,2,'2019-03-10 00:00:00','2019-03-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(822,19,1,0,'2019-08-16 00:00:00','2019-08-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(823,19,1,2,'2019-08-14 00:00:00','2019-08-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(824,19,1,2,'2019-05-31 00:00:00','2019-06-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(825,19,1,2,'2019-05-29 00:00:00','2019-05-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(826,19,1,2,'2019-05-29 00:00:00','2019-05-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(827,19,1,0,'2019-09-04 00:00:00','2019-09-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(828,19,1,2,'2019-09-03 00:00:00','2019-09-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(829,19,1,2,'2019-07-15 00:00:00','2019-07-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(830,19,1,0,'2019-01-24 00:00:00','2019-02-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(831,19,1,2,'2019-01-23 00:00:00','2019-01-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(832,19,1,2,'2019-11-14 00:00:00','2019-11-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(833,19,1,0,'2019-09-08 00:00:00','2019-09-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(834,19,1,2,'2019-09-07 00:00:00','2019-09-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(835,19,1,2,'2019-01-11 00:00:00','2019-01-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(836,19,1,2,'2019-04-18 00:00:00','2019-04-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(837,19,1,2,'2019-05-26 00:00:00','2019-05-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(838,20,1,0,'2019-02-26 00:00:00','2019-02-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(839,20,1,2,'2019-02-23 00:00:00','2019-02-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(840,20,1,0,'2019-10-06 00:00:00','2019-10-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(841,20,1,2,'2019-10-02 00:00:00','2019-10-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(842,20,1,0,'2019-01-26 00:00:00','2019-01-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(843,20,1,2,'2019-01-22 00:00:00','2019-01-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(844,20,1,0,'2019-05-21 00:00:00','2019-06-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(845,20,1,2,'2019-05-18 00:00:00','2019-05-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(846,20,1,0,'2019-12-19 00:00:00','2019-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(847,20,1,2,'2019-12-17 00:00:00','2019-12-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(848,20,1,0,'2019-02-08 00:00:00','2019-02-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(849,20,1,2,'2019-02-05 00:00:00','2019-02-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(850,20,1,0,'2019-01-15 00:00:00','2019-01-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(851,20,1,2,'2019-01-12 00:00:00','2019-01-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(852,20,1,0,'2019-04-15 00:00:00','2019-04-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(854,20,1,0,'2019-11-09 00:00:00','2019-11-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(855,20,1,2,'2019-11-08 00:00:00','2019-11-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(856,20,1,0,'2019-04-23 00:00:00','2019-04-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(858,20,1,2,'2019-10-30 00:00:00','2019-11-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(859,20,1,0,'2019-03-24 00:00:00','2019-03-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(860,20,1,2,'2019-03-23 00:00:00','2019-03-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(861,20,1,0,'2019-11-27 00:00:00','2019-12-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(862,20,1,2,'2019-11-20 00:00:00','2019-11-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(863,20,1,0,'2019-12-08 00:00:00','2019-12-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(864,20,1,2,'2019-12-02 00:00:00','2019-12-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(865,20,1,2,'2019-11-22 00:00:00','2019-11-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(867,20,1,0,'2019-05-07 00:00:00','2019-05-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(868,20,1,2,'2019-05-06 00:00:00','2019-05-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(870,20,1,2,'2019-04-26 00:00:00','2019-05-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(871,20,1,2,'2019-05-01 00:00:00','2019-05-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(872,20,1,0,'2019-07-13 00:00:00','2019-07-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(873,20,1,2,'2019-07-12 00:00:00','2019-07-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(874,20,1,0,'2019-03-04 00:00:00','2019-03-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(875,20,1,2,'2019-02-28 00:00:00','2019-03-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(876,20,1,0,'2019-06-05 00:00:00','2019-06-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(877,20,1,2,'2019-06-04 00:00:00','2019-06-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(878,20,1,0,'2020-01-02 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(879,20,1,2,'2019-12-31 00:00:00','2020-01-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(880,20,1,2,'2019-11-02 00:00:00','2019-11-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(881,20,1,2,'2019-11-03 00:00:00','2019-11-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(882,20,1,0,'2019-09-08 00:00:00','2019-10-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(883,20,1,2,'2019-09-04 00:00:00','2019-09-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(884,20,1,0,'2019-06-25 00:00:00','2019-06-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(885,20,1,2,'2019-06-18 00:00:00','2019-06-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(886,20,1,2,'2019-05-12 00:00:00','2019-05-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(887,20,1,0,'2019-07-22 00:00:00','2019-08-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(888,20,1,2,'2019-07-21 00:00:00','2019-07-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(889,20,1,2,'2019-06-23 00:00:00','2019-06-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(890,20,1,0,'2019-08-30 00:00:00','2019-09-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(891,20,1,2,'2019-08-25 00:00:00','2019-08-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(892,20,1,0,'2019-04-06 00:00:00','2019-04-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(893,20,1,2,'2019-04-05 00:00:00','2019-04-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(894,20,1,0,'2019-10-15 00:00:00','2019-10-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(895,20,1,2,'2019-10-08 00:00:00','2019-10-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(896,20,1,0,'2019-08-21 00:00:00','2019-08-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(897,20,1,2,'2019-08-15 00:00:00','2019-08-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(898,20,1,0,'2019-01-11 00:00:00','2019-01-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(899,20,1,2,'2019-01-09 00:00:00','2019-01-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(900,20,1,0,'2019-03-30 00:00:00','2019-04-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(901,20,1,2,'2019-03-27 00:00:00','2019-03-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(902,20,1,0,'2019-03-19 00:00:00','2019-03-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(903,20,1,2,'2019-03-16 00:00:00','2019-03-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(904,20,1,2,'2019-12-11 00:00:00','2019-12-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(905,20,1,0,'2019-01-29 00:00:00','2019-02-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(906,20,1,2,'2019-01-27 00:00:00','2019-01-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(907,20,1,2,'2019-01-08 00:00:00','2019-01-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(909,20,1,2,'2019-04-21 00:00:00','2019-04-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(910,20,1,2,'2019-01-20 00:00:00','2019-01-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(911,20,1,2,'2019-09-02 00:00:00','2019-09-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(912,20,1,0,'2019-07-04 00:00:00','2019-07-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(913,20,1,2,'2019-06-27 00:00:00','2019-07-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(914,20,1,2,'2019-04-14 00:00:00','2019-04-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(915,20,1,0,'2019-07-07 00:00:00','2019-07-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(916,20,1,2,'2019-07-06 00:00:00','2019-07-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(917,20,1,2,'2019-01-06 00:00:00','2019-01-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(918,20,1,2,'2019-10-13 00:00:00','2019-10-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(919,20,1,2,'2019-10-11 00:00:00','2019-10-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(920,20,1,2,'2019-07-08 00:00:00','2019-07-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(921,21,1,0,'2019-03-15 00:00:00','2019-03-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(922,21,1,2,'2019-03-10 00:00:00','2019-03-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(923,21,1,0,'2019-06-26 00:00:00','2019-07-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(924,21,1,2,'2019-06-25 00:00:00','2019-06-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(925,21,1,0,'2019-10-04 00:00:00','2019-10-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(926,21,1,2,'2019-09-30 00:00:00','2019-10-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(927,21,1,0,'2019-08-01 00:00:00','2019-08-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(928,21,1,2,'2019-07-31 00:00:00','2019-08-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(929,21,1,0,'2019-02-10 00:00:00','2019-02-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(930,21,1,2,'2019-02-09 00:00:00','2019-02-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(931,21,1,0,'2019-11-18 00:00:00','2019-11-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(932,21,1,2,'2019-11-17 00:00:00','2019-11-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(933,21,1,2,'2019-07-24 00:00:00','2019-07-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(934,21,1,0,'2019-06-17 00:00:00','2019-06-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(935,21,1,2,'2019-06-12 00:00:00','2019-06-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(939,21,1,2,'2019-06-18 00:00:00','2019-06-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(940,21,1,0,'2019-08-20 00:00:00','2019-09-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(941,21,1,2,'2019-08-17 00:00:00','2019-08-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(942,21,1,0,'2019-04-03 00:00:00','2019-04-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(943,21,1,2,'2019-03-24 00:00:00','2019-03-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(944,21,1,0,'2019-11-07 00:00:00','2019-11-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(945,21,1,2,'2019-11-06 00:00:00','2019-11-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(946,21,1,0,'2019-06-11 00:00:00','2019-06-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(947,21,1,2,'2019-06-07 00:00:00','2019-06-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(948,21,1,0,'2019-02-21 00:00:00','2019-02-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(949,21,1,2,'2019-02-14 00:00:00','2019-02-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(950,21,1,0,'2019-09-09 00:00:00','2019-09-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(951,21,1,2,'2019-09-06 00:00:00','2019-09-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(952,21,1,0,'2019-12-30 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(954,21,1,0,'2019-12-17 00:00:00','2019-12-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:53',0),(955,21,1,2,'2019-12-13 00:00:00','2019-12-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(956,21,1,0,'2019-03-07 00:00:00','2019-03-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(957,21,1,2,'2019-03-06 00:00:00','2019-03-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(958,21,1,0,'2019-04-12 00:00:00','2019-04-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(959,21,1,2,'2019-04-09 00:00:00','2019-04-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(960,21,1,0,'2019-05-03 00:00:00','2019-06-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(961,21,1,2,'2019-04-29 00:00:00','2019-05-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(962,21,1,0,'2019-07-13 00:00:00','2019-07-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(963,21,1,2,'2019-07-10 00:00:00','2019-07-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(964,21,1,0,'2019-10-22 00:00:00','2019-10-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(965,21,1,2,'2019-10-18 00:00:00','2019-10-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(966,21,1,2,'2019-03-28 00:00:00','2019-04-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(967,21,1,0,'2019-04-26 00:00:00','2019-04-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(969,21,1,0,'2019-12-11 00:00:00','2019-12-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(970,21,1,2,'2019-12-05 00:00:00','2019-12-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(971,21,1,0,'2019-11-28 00:00:00','2019-12-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(972,21,1,2,'2019-11-24 00:00:00','2019-11-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(973,21,1,0,'2019-03-03 00:00:00','2019-03-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(974,21,1,2,'2019-02-26 00:00:00','2019-03-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(975,21,1,2,'2019-06-18 00:00:00','2019-06-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(977,21,1,2,'2019-03-09 00:00:00','2019-03-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(978,21,1,2,'2019-03-21 00:00:00','2019-03-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(979,21,1,2,'2019-07-07 00:00:00','2019-07-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(980,21,1,2,'2019-04-02 00:00:00','2019-04-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(981,21,1,2,'2019-04-01 00:00:00','2019-04-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(982,21,1,0,'2019-01-15 00:00:00','2019-02-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(983,21,1,2,'2019-01-07 00:00:00','2019-01-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(984,21,1,2,'2019-02-16 00:00:00','2019-02-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(985,21,1,0,'2019-10-29 00:00:00','2019-11-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(986,21,1,2,'2019-10-23 00:00:00','2019-10-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(987,21,1,2,'2019-01-08 00:00:00','2019-01-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(988,21,1,2,'2019-04-20 00:00:00','2019-04-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(989,21,1,0,'2019-02-07 00:00:00','2019-02-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(990,21,1,2,'2019-02-02 00:00:00','2019-02-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(991,21,1,0,'2019-09-05 00:00:00','2019-09-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(992,21,1,2,'2019-09-04 00:00:00','2019-09-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(993,21,1,0,'2019-07-18 00:00:00','2019-07-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(994,21,1,2,'2019-07-16 00:00:00','2019-07-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(995,21,1,2,'2019-12-26 00:00:00','2019-12-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:53',0),(996,21,1,2,'2019-07-30 00:00:00','2019-07-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(997,21,1,2,'2019-07-25 00:00:00','2019-07-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(998,21,1,0,'2019-03-20 00:00:00','2019-03-21 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(999,21,1,2,'2019-03-19 00:00:00','2019-03-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(1000,21,1,2,'2019-12-20 00:00:00','2019-12-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1001,22,1,0,'2019-12-18 00:00:00','2019-12-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1002,22,1,2,'2019-12-14 00:00:00','2019-12-15 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1003,22,1,0,'2019-12-29 00:00:00','9999-12-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1005,22,1,0,'2019-09-02 00:00:00','2019-09-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1006,22,1,2,'2019-08-31 00:00:00','2019-09-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1007,22,1,0,'2019-05-27 00:00:00','2019-05-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1008,22,1,2,'2019-05-10 00:00:00','2019-05-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1009,22,1,0,'2019-02-09 00:00:00','2019-02-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1010,22,1,2,'2019-02-08 00:00:00','2019-02-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1012,22,1,2,'2019-01-10 00:00:00','2019-01-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1013,22,1,0,'2019-07-02 00:00:00','2019-07-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1014,22,1,2,'2019-06-27 00:00:00','2019-07-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1015,22,1,0,'2019-08-27 00:00:00','2019-08-31 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1016,22,1,2,'2019-08-25 00:00:00','2019-08-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1017,22,1,0,'2019-08-06 00:00:00','2019-08-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1018,22,1,2,'2019-07-30 00:00:00','2019-08-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1019,22,1,0,'2019-08-20 00:00:00','2019-08-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1020,22,1,2,'2019-08-12 00:00:00','2019-08-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1021,22,1,0,'2019-04-27 00:00:00','2019-04-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1022,22,1,2,'2019-04-23 00:00:00','2019-04-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1023,22,1,0,'2019-02-25 00:00:00','2019-03-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1024,22,1,2,'2019-02-23 00:00:00','2019-02-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1025,22,1,0,'2019-03-03 00:00:00','2019-03-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1026,22,1,2,'2019-03-01 00:00:00','2019-03-03 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1027,22,1,0,'2019-01-22 00:00:00','2019-02-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1029,22,1,0,'2019-05-07 00:00:00','2019-05-10 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1030,22,1,2,'2019-05-02 00:00:00','2019-05-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1031,22,1,0,'2019-11-07 00:00:00','2019-12-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1032,22,1,2,'2019-11-06 00:00:00','2019-11-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1033,22,1,0,'2019-03-11 00:00:00','2019-03-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1035,22,1,2,'2019-05-13 00:00:00','2019-05-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1036,22,1,2,'2019-08-26 00:00:00','2019-08-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1037,22,1,0,'2019-07-28 00:00:00','2019-07-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1038,22,1,2,'2019-07-24 00:00:00','2019-07-25 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1039,22,1,0,'2019-07-19 00:00:00','2019-07-24 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1040,22,1,2,'2019-07-16 00:00:00','2019-07-19 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1042,22,1,2,'2019-08-18 00:00:00','2019-08-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1043,22,1,2,'2019-01-20 00:00:00','2019-01-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1044,22,1,0,'2019-02-17 00:00:00','2019-02-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1045,22,1,2,'2019-02-10 00:00:00','2019-02-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1047,22,1,2,'2019-12-12 00:00:00','2019-12-14 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1048,22,1,2,'2019-08-02 00:00:00','2019-08-09 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1049,22,1,0,'2019-11-03 00:00:00','2019-11-06 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1051,22,1,2,'2019-01-18 00:00:00','2019-01-20 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1052,22,1,0,'2019-09-16 00:00:00','2019-11-01 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1053,22,1,2,'2019-09-11 00:00:00','2019-09-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1054,22,1,0,'2019-07-07 00:00:00','2019-07-13 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1055,22,1,2,'2019-07-03 00:00:00','2019-07-07 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1056,22,1,2,'2019-07-26 00:00:00','2019-07-28 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1057,22,1,2,'2019-07-25 00:00:00','2019-07-26 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1058,22,1,2,'2019-05-22 00:00:00','2019-05-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1059,22,1,2,'2019-12-23 00:00:00','2019-12-29 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1060,22,1,2,'2019-07-13 00:00:00','2019-07-16 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1061,22,1,0,'2019-02-07 00:00:00','2019-02-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1062,22,1,2,'2019-02-02 00:00:00','2019-02-05 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1063,22,1,2,'2019-05-19 00:00:00','2019-05-22 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1064,22,1,2,'2019-01-13 00:00:00','2019-01-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1067,22,1,0,'2019-04-30 00:00:00','2019-05-02 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1068,22,1,2,'2019-04-29 00:00:00','2019-04-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1069,22,1,2,'2019-03-05 00:00:00','2019-03-11 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1070,22,1,2,'2019-12-17 00:00:00','2019-12-18 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1071,22,1,2,'2019-12-15 00:00:00','2019-12-17 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1072,22,1,2,'2019-11-01 00:00:00','2019-11-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1073,22,1,0,'2019-03-30 00:00:00','2019-04-23 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1074,22,1,2,'2019-03-28 00:00:00','2019-03-30 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1075,22,1,2,'2019-02-05 00:00:00','2019-02-08 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1076,22,1,0,'2019-06-04 00:00:00','2019-06-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1077,22,1,2,'2019-05-28 00:00:00','2019-06-04 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1078,22,1,2,'2019-12-08 00:00:00','2019-12-12 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0),(1079,22,1,2,'2019-05-23 00:00:00','2019-05-27 00:00:00','','','','','','','','','','','','','','','','','',0,'2019-02-23 00:40:54',0,'2019-02-23 00:40:54',0);
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
INSERT INTO `RentableMarketRate` VALUES (1,1,0,1000.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,2,0,1500.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,3,0,1750.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,4,0,2500.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(5,5,0,65.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(6,6,0,75.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(7,7,0,85.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(8,8,0,35.0000,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
INSERT INTO `RentableTypeRef` VALUES (1,1,1,1,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,2,1,1,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,3,1,2,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,4,1,2,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(5,5,1,3,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(6,6,1,3,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(7,7,1,4,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(8,8,1,4,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(9,9,1,5,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(10,10,1,5,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(11,11,1,6,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(12,12,1,6,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(13,13,1,7,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(14,14,1,7,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(15,15,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(16,16,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(17,17,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(18,18,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(19,19,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(20,20,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(21,21,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(22,22,1,8,0,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
INSERT INTO `RentableTypes` VALUES (1,1,'1BR1BA','Standard',6,4,4,12,45,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,1,'2BR1BA','Deluxe',6,4,4,12,46,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,1,'2BR2BA','Gold',6,4,4,12,47,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,1,'3BR2BA','Platinum',6,4,4,12,48,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(5,1,'H11','HotelStd',4,0,0,4,49,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(6,1,'H21','HotelGold',4,0,0,4,50,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(7,1,'H22','HotelPlatinum',4,0,0,4,51,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(8,1,'CP000','Car Port 000',6,4,4,6,52,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
INSERT INTO `RentableUseStatus` VALUES (1,1,1,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,2,1,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,3,1,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,4,1,0,'2019-01-01 00:00:00','2019-01-03 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(5,5,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(6,6,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(7,7,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(8,8,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(9,9,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(10,10,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(11,11,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(12,12,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(13,13,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(14,14,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(15,15,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(16,16,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(17,17,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(18,18,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(19,19,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(20,20,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(21,21,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(22,22,1,0,'2019-01-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(23,1,1,0,'2020-03-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(24,1,1,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(25,2,1,0,'2020-03-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(26,2,1,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(27,3,1,0,'2020-03-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(28,3,1,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(29,4,1,0,'2020-03-01 00:00:00','9999-12-31 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(30,4,1,1,'2019-01-03 00:00:00','2020-03-01 00:00:00','','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableUsers`
--

LOCK TABLES `RentableUsers` WRITE;
/*!40000 ALTER TABLE `RentableUsers` DISABLE KEYS */;
INSERT INTO `RentableUsers` VALUES (1,1,1,1,'2019-01-03','2020-03-01','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,2,1,2,'2019-01-03','2020-03-01','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,3,1,3,'2019-01-03','2020-03-01','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,4,1,4,'2019-01-03','2020-03-01','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
INSERT INTO `RentalAgreement` VALUES (1,0,0,1,1,0,'2019-01-03 00:00:00','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-02-01',2,0,2,'',0,0,0.0000,'','0000-00-00','0000-00-00',0.0000,0.0000,124,'0000-00-00','','','','0000-00-00','','0000-00-00','','0000-00-00',0,52,94,'2019-01-03 00:00:00',40,'2019-01-03 00:00:00',0,200,'2019-01-03 00:00:00',0,288,'2019-01-03 00:00:00',82,'2019-01-03 00:00:00',0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00',0,'0000-00-00 00:00:00','1970-01-01 00:00:00',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,0,0,1,1,0,'2019-01-03 00:00:00','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-02-01',0,1,2,'',0,0,0.0000,'','0000-00-00','0000-00-00',0.0000,0.0000,235,'0000-00-00','','','','0000-00-00','','0000-00-00','','0000-00-00',0,52,271,'2019-01-03 00:00:00',296,'2019-01-03 00:00:00',0,94,'2019-01-03 00:00:00',0,171,'2019-01-03 00:00:00',122,'2019-01-03 00:00:00',0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00',0,'0000-00-00 00:00:00','1970-01-01 00:00:00',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,0,0,1,1,0,'2019-01-03 00:00:00','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-02-01',0,2,2,'',0,0,0.0000,'','0000-00-00','0000-00-00',0.0000,0.0000,247,'0000-00-00','','','','0000-00-00','','0000-00-00','','0000-00-00',0,52,215,'2019-01-03 00:00:00',174,'2019-01-03 00:00:00',0,47,'2019-01-03 00:00:00',0,104,'2019-01-03 00:00:00',186,'2019-01-03 00:00:00',0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00',0,'0000-00-00 00:00:00','1970-01-01 00:00:00',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,0,0,1,1,0,'2019-01-03 00:00:00','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-01-03','2020-03-01','2019-02-01',0,2,2,'',0,0,0.0000,'','0000-00-00','0000-00-00',0.0000,0.0000,180,'0000-00-00','','','','0000-00-00','','0000-00-00','','0000-00-00',0,52,244,'2019-01-03 00:00:00',193,'2019-01-03 00:00:00',0,66,'2019-01-03 00:00:00',0,159,'2019-01-03 00:00:00',217,'2019-01-03 00:00:00',0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00',0,'0000-00-00 00:00:00','1970-01-01 00:00:00',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
INSERT INTO `RentalAgreementPayors` VALUES (1,1,1,1,'2019-01-03','2020-03-01',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,2,1,2,'2019-01-03','2020-03-01',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,3,1,3,'2019-01-03','2020-03-01',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,4,1,4,'2019-01-03','2020-03-01',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
INSERT INTO `RentalAgreementRentables` VALUES (1,1,1,1,0,0,1000.0000,'2019-01-03','2020-03-01','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,2,1,2,0,0,1000.0000,'2019-01-03','2020-03-01','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,3,1,3,0,0,1500.0000,'2019-01-03','2020-03-01','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,4,1,4,0,0,1500.0000,'2019-01-03','2020-03-01','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TBind`
--

LOCK TABLES `TBind` WRITE;
/*!40000 ALTER TABLE `TBind` DISABLE KEYS */;
INSERT INTO `TBind` VALUES (1,1,1,1,15,1,'2019-01-03 00:00:00','9999-12-31 00:00:00',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,1,1,2,15,2,'2019-01-03 00:00:00','9999-12-31 00:00:00',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,1,1,2,14,1,'2019-01-03 00:00:00','9999-12-31 00:00:00',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,1,1,3,15,3,'2019-01-03 00:00:00','9999-12-31 00:00:00',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(5,1,1,3,14,2,'2019-01-03 00:00:00','9999-12-31 00:00:00',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(6,1,1,4,15,4,'2019-01-03 00:00:00','9999-12-31 00:00:00',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(7,1,1,4,14,3,'2019-01-03 00:00:00','9999-12-31 00:00:00',0,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
INSERT INTO `TWS` VALUES (1,'CreateAssessmentInstances','','CreateAssessmentInstances','2018-02-25 00:00:00','Steves-MacBook-Pro-2.local',4,'2018-02-24 01:19:54','2018-02-24 01:19:54','2017-11-10 15:24:21','2018-02-23 17:19:53'),(2,'CleanRARBalanceCache','','CleanRARBalanceCache','2018-02-24 05:09:45','Steves-MacBook-Pro-2.local',4,'2018-02-24 05:04:45','2018-02-24 05:04:45','2018-02-23 17:19:43','2018-02-23 21:04:45'),(3,'CleanSecDepBalanceCache','','CleanSecDepBalanceCache','2018-02-24 05:09:45','Steves-MacBook-Pro-2.local',4,'2018-02-24 05:04:45','2018-02-24 05:04:45','2018-02-23 17:19:43','2018-02-23 21:04:45'),(4,'CleanAcctSliceCache','','CleanAcctSliceCache','2018-02-24 05:09:45','Steves-MacBook-Pro-2.local',4,'2018-02-24 05:04:45','2018-02-24 05:04:45','2018-02-23 17:19:43','2018-02-23 21:04:45'),(5,'CleanARSliceCache','','CleanARSliceCache','2018-02-24 05:09:45','Steves-MacBook-Pro-2.local',4,'2018-02-24 05:04:45','2018-02-24 05:04:45','2018-02-23 17:19:43','2018-02-23 21:04:45'),(6,'RARBcacheBot','','RARBcacheBot','2019-02-23 00:56:12','Steves-MacBook-Pro-2.local',4,'2019-02-23 00:51:12','2019-02-23 00:51:12','2018-06-02 13:09:58','2019-02-22 16:51:11'),(7,'ARSliceCacheBot','','ARSliceCacheBot','2019-02-23 00:56:12','Steves-MacBook-Pro-2.local',4,'2019-02-23 00:51:12','2019-02-23 00:51:12','2018-06-02 13:09:58','2019-02-22 16:51:11'),(8,'TLReportBot','','TLReportBot','2019-02-23 00:53:22','Steves-MacBook-Pro-2.local',4,'2019-02-23 00:51:22','2019-02-23 00:51:22','2018-06-02 13:09:58','2019-02-22 16:51:21'),(9,'ManualTaskBot','','ManualTaskBot','2019-02-24 00:41:01','Steves-MacBook-Pro-2.local',4,'2019-02-23 00:41:01','2019-02-23 00:41:01','2018-06-02 13:09:58','2019-02-22 16:41:00'),(10,'AssessmentBot','','AssessmentBot','2019-02-24 00:41:01','Steves-MacBook-Pro-2.local',4,'2019-02-23 00:41:01','2019-02-23 00:41:01','2018-06-02 13:09:58','2019-02-22 16:41:00'),(11,'SecDepCacheBot','','SecDepCacheBot','2019-02-23 00:56:12','Steves-MacBook-Pro-2.local',4,'2019-02-23 00:51:12','2019-02-23 00:51:12','2018-06-02 13:09:58','2019-02-22 16:51:11'),(12,'AcctSliceCacheBot','','AcctSliceCacheBot','2019-02-23 00:56:12','Steves-MacBook-Pro-2.local',4,'2019-02-23 00:51:12','2019-02-23 00:51:12','2018-06-02 13:09:58','2019-02-22 16:51:11'),(13,'TLInstanceBot','','TLInstanceBot','2019-02-24 00:41:01','Steves-MacBook-Pro-2.local',4,'2019-02-23 00:41:01','2019-02-23 00:41:01','2018-06-02 13:09:58','2019-02-22 16:41:00');
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
INSERT INTO `Task` VALUES (1,1,1,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','2019-01-31 20:00:00','2019-01-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(2,1,1,'Review all receivables for accuracy','ManualTaskBot','2019-01-31 20:00:00','2019-01-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(3,1,1,'Compare total cash deposits to bank statement','ManualTaskBot','2019-01-31 20:00:00','2019-01-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(4,1,1,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','2019-01-31 07:00:00','2019-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(5,1,1,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','2019-01-31 07:00:00','2019-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(6,1,1,'Make certain that all suspense accounts have been closed out','ManualTaskBot','2019-01-31 07:00:00','2019-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(7,1,1,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','2019-01-31 07:00:00','2019-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(8,1,1,'Print Rent Roll Activity Report','ManualTaskBot','2019-01-31 07:00:00','2019-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(9,1,1,'Print Rent Roll Report','ManualTaskBot','2019-01-31 07:00:00','2019-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(10,1,1,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','2019-01-31 07:00:00','2019-01-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(11,1,2,'Tie closing SECDEP balance to bank SECDEP balance','ManualTaskBot','2019-02-28 20:00:00','2019-02-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(12,1,2,'Review all receivables for accuracy','ManualTaskBot','2019-02-28 20:00:00','2019-02-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(13,1,2,'Compare total cash deposits to bank statement','ManualTaskBot','2019-02-28 20:00:00','2019-02-20 20:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(14,1,2,'Confirm all Lease Concessions are document in resident\'s lease','ManualTaskBot','2019-02-28 07:00:00','2019-02-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(15,1,2,'Tie all Bar/Spa/F&B deposits in POS Lavu to Rent Roll Deposits','ManualTaskBot','2019-02-28 07:00:00','2019-02-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(16,1,2,'Make certain that all suspense accounts have been closed out','ManualTaskBot','2019-02-28 07:00:00','2019-02-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(17,1,2,'Compile all workpapers for the foregoing confirmations, and file as YYYY-MM-DD [3-letter property] Rent Roll Work Papers','ManualTaskBot','2019-02-28 07:00:00','2019-02-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(18,1,2,'Print Rent Roll Activity Report','ManualTaskBot','2019-02-28 07:00:00','2019-02-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(19,1,2,'Print Rent Roll Report','ManualTaskBot','2019-02-28 07:00:00','2019-02-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(20,1,2,'File PDFs for the reports as YYY-MM-DD [3-letter-property] Rent Roll','ManualTaskBot','2019-02-28 07:00:00','2019-02-20 07:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0);
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
INSERT INTO `TaskList` VALUES (1,1,0,1,'Monthly Close',6,'2019-01-31 17:00:00','2019-01-20 17:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','0000-00-00 00:00:00',86400000000000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0),(2,1,1,1,'Monthly Close',6,'2019-02-28 17:00:00','2019-02-20 17:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',6,0,0,'','0000-00-00 00:00:00',86400000000000,'','2019-02-23 00:40:53',0,'2019-02-23 00:40:53',0);
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
INSERT INTO `Transactant` VALUES (1,1,0,'Tawanna','Verda','Roberts','Kristi','Lithia Motors Inc.',0,'TRoberts3383@abiz.com','TawannaRoberts963@aol.com','(789) 819-3673','(893) 901-3042','59749 West Virginia','','Lakewood','LA','98081','USA','',0,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,1,0,'Dagmar','Richard','Rhodes','Earleen','Systemax Inc.',0,'DagmarRhodes11@comcast.net','DagmarRhodes306@abiz.com','(102) 281-2558','(109) 504-6572','28874 Highland','','Lafayette','FL','54425','USA','',0,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,1,0,'Lynna','Roslyn','Waller','Petronila','The Bear Stearns Companies Inc.',0,'LynnaW7407@yahoo.com','LynnaW2591@abiz.com','(244) 706-7428','(413) 158-9584','96205 Airport','','Cary','MA','28162','USA','',0,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,1,0,'Marylee','Ellamae','Ratliff','Margareta','Stilwell Financial Inc',0,'MaryleeRatliff950@bdiddy.com','MaryleeRatliff91@bdiddy.com','(501) 615-7319','(742) 618-6155','91917 13th','','Simi Valley','WA','23237','USA','',0,'','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
INSERT INTO `User` VALUES (1,1,0,'1980-04-30','Earlene Hester','17886 Smith,Santa Clara,NC 27887','(406) 916-0855','EarleneH5361@comcast.net','99394 Hampton,Muskegon,MA 31847',1,0,264,52,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,1,0,'1966-07-23','Sanda Flynn','73652 West,Virginia Beach,HI 22540','(837) 211-9297','SandaFlynn188@yahoo.com','52696 Dogwood,Alexandria,ND 40456',1,0,325,17,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,1,0,'1982-01-21','Cythia Workman','87779 Wilson,Chesapeake,WI 55089','(473) 288-6056','CWorkman8547@bdiddy.com','78764 Aspen,Bloomington,OH 24728',0,0,233,49,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,1,0,'1967-01-03','Jacklyn Atkins','86036 Cypress,Kenosha,AR 39106','(396) 576-8094','JAtkins1131@gmail.com','25790 11th,Santa Clara,MO 40495',1,0,316,8,'2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Vehicle`
--

LOCK TABLES `Vehicle` WRITE;
/*!40000 ALTER TABLE `Vehicle` DISABLE KEYS */;
INSERT INTO `Vehicle` VALUES (1,1,1,'car','Mazda','B-Series Plus','Navy',1998,'RFGF1ZJN3JS8PNCS','NH','H7Z4Y75','1360501','2019-01-03','2020-03-01','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(2,2,1,'car','Subaru','Tribeca','Copper',2006,'ZDMXAJEGOS56LI6O','AK','X860O0H','2350450','2019-01-03','2020-03-01','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(3,3,1,'car','Ford','F250','Navy',2003,'MA6CI6U17C2GO8AJ','ME','S1X022A','9517555','2019-01-03','2020-03-01','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0),(4,4,1,'car','Ford','EXP','Gray',1988,'3P3Q044QA3ZDP2XF','IL','1C5OR71','7652593','2019-01-03','2020-03-01','2019-02-23 00:40:52',0,'2019-02-23 00:40:52',0);
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

-- Dump completed on 2019-05-04 17:49:51
