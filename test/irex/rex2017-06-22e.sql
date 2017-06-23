-- MySQL dump 10.13  Distrib 5.7.18, for Linux (x86_64)
--
-- Host: localhost    Database: rentroll
-- ------------------------------------------------------
-- Server version	5.7.18

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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`ARID`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `AR`
--

LOCK TABLES `AR` WRITE;
/*!40000 ALTER TABLE `AR` DISABLE KEYS */;
INSERT INTO `AR` VALUES (1,1,'Rent Non-Taxable',0,0,2,3,'','2017-06-13 00:00:00','2018-06-13 00:00:00','2017-06-13 19:18:47',0,'2017-06-14 18:26:45',0),(2,1,'Payment by Check',1,0,1,27,'','2017-06-13 00:00:00','2018-06-13 00:00:00','2017-06-13 19:20:49',0,'2017-06-14 18:26:45',0),(3,1,'Payment by WIRE',1,0,1,0,'','2017-06-13 00:00:00','2018-06-13 00:00:00','2017-06-13 19:21:57',0,'2017-06-14 18:26:45',0),(4,1,'Payment by WIRE',1,0,1,27,'','2017-06-13 00:00:00','2018-06-13 00:00:00','2017-06-13 19:22:40',0,'2017-06-14 18:26:45',0),(5,1,'Payment by ACH',1,0,1,27,'','2017-06-13 00:00:00','2018-06-13 00:00:00','2017-06-13 19:23:02',0,'2017-06-14 18:26:45',0),(6,1,'Payment by EFT',1,0,1,27,'','2017-06-13 00:00:00','2018-06-13 00:00:00','2017-06-13 19:23:17',0,'2017-06-14 18:26:45',0),(7,1,'Payment by Money Order',1,0,1,27,'','2017-06-13 00:00:00','2018-06-13 00:00:00','2017-06-13 19:23:43',0,'2017-06-14 18:26:45',0),(8,1,'Payment by Branch Deposit',1,0,1,27,'','2017-06-13 00:00:00','2018-06-13 00:00:00','2017-06-13 19:24:08',0,'2017-06-14 18:26:45',0),(9,1,'Tenant Security Deposit',0,0,2,7,'','2017-06-13 00:00:00','2018-06-13 00:00:00','2017-06-22 17:08:06',0,'2017-06-14 18:26:45',0),(10,1,'Loss to Lease',0,0,2,4,'','2017-06-13 00:00:00','2018-06-13 00:00:00','2017-06-22 17:06:37',0,'2017-06-14 18:26:45',0),(11,1,'Late Fee',0,2,2,31,'','2017-06-13 00:00:00','2018-06-13 00:00:00','2017-06-22 17:05:47',0,'2017-06-14 18:26:45',0),(12,1,'Rental Advertising/Commission',0,1,2,34,'','2014-01-01 00:00:00','9999-01-01 00:00:00','2017-06-22 17:07:19',0,'2017-06-20 18:14:32',0),(13,1,'NSF Fees',0,0,2,17,'NSF Fees','2017-06-22 00:00:00','2018-06-22 00:00:00','2017-06-22 16:52:38',0,'2017-06-22 16:52:38',0),(14,1,'Forfeited Security Deposit',0,0,2,18,'Forfeited Security Deposit','2017-06-22 00:00:00','2018-06-22 00:00:00','2017-06-22 16:53:31',0,'2017-06-22 16:53:31',0),(15,1,'Insurance Reimbursement',0,0,2,19,'Insurance Reimbursement','2017-06-22 00:00:00','2018-06-22 00:00:00','2017-06-22 16:53:59',0,'2017-06-22 16:53:59',0),(16,1,'Utility Reimbursements',0,0,2,35,'Utility Reimbursements','2017-06-22 00:00:00','2018-06-22 00:00:00','2017-06-22 16:57:46',0,'2017-06-22 16:57:46',0),(17,1,'Maintenance Reimbursements',0,0,2,20,'Maintenance Reimbursements','2017-06-22 00:00:00','2018-06-22 00:00:00','2017-06-22 16:58:23',0,'2017-06-22 16:58:23',0),(18,1,'Other Rental Income',0,0,2,21,'Other Rental Income','2017-06-22 00:00:00','2018-06-22 00:00:00','2017-06-22 16:59:11',0,'2017-06-22 16:59:11',0),(19,1,'Vacanty',0,0,2,5,'','2017-06-22 00:00:00','2018-06-22 00:00:00','2017-06-22 17:01:39',0,'2017-06-22 17:01:39',0),(20,1,'Rent Concession',0,0,2,22,'','2017-06-22 00:00:00','2018-06-22 00:00:00','2017-06-22 17:02:08',0,'2017-06-22 17:02:08',0),(21,1,'Offline Units',0,0,2,23,'','2017-06-22 00:00:00','2018-06-22 00:00:00','2017-06-22 17:02:33',0,'2017-06-22 17:02:33',0),(22,1,'Tenant Bad Debt',0,0,2,24,'','2017-06-22 00:00:00','2018-06-22 00:00:00','2017-06-22 17:03:06',0,'2017-06-22 17:03:06',0),(23,1,'Administrative Unit',0,0,2,25,'','2017-06-22 00:00:00','2018-06-22 00:00:00','2017-06-22 17:03:46',0,'2017-06-22 17:03:46',0),(24,1,'Other Income Offsets',0,0,2,26,'','2017-06-22 00:00:00','2018-06-22 00:00:00','2017-06-22 17:04:17',0,'2017-06-22 17:04:17',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Assessments`
--

LOCK TABLES `Assessments` WRITE;
/*!40000 ALTER TABLE `Assessments` DISABLE KEYS */;
INSERT INTO `Assessments` VALUES (1,0,1,1,0,1,7000.0000,'2014-03-01 00:00:00','2014-03-01 00:00:00',0,0,0,'',9,0,'','2017-06-20 16:10:25',0,'2017-06-20 16:10:25',0),(4,0,1,1,0,1,3750.0000,'2016-03-01 00:00:00','2018-03-01 00:00:00',6,4,0,'',1,0,'','2017-06-20 17:52:50',0,'2017-06-20 17:52:50',0),(5,0,1,0,0,0,4150.0000,'2016-07-01 00:00:00','2018-07-01 00:00:00',6,4,0,'',1,0,'','2017-06-20 17:53:52',0,'2017-06-20 17:53:52',0),(6,0,1,3,0,2,8300.0000,'2016-07-01 00:00:00','2018-07-01 00:00:00',0,0,0,'',9,0,'','2017-06-20 18:06:02',0,'2017-06-20 18:05:06',0),(7,0,1,3,0,2,4150.0000,'2016-07-01 00:00:00','2018-07-01 00:00:00',6,4,0,'',1,0,'','2017-06-20 18:05:37',0,'2017-06-20 18:05:37',0),(8,0,1,2,0,3,4000.0000,'2016-10-01 00:00:00','2018-01-01 00:00:00',6,4,0,'',1,0,'','2017-06-20 18:06:46',0,'2017-06-20 18:06:46',0),(9,4,1,1,0,1,3750.0000,'2016-03-01 00:00:00','2016-03-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(10,4,1,1,0,1,3750.0000,'2016-04-01 00:00:00','2016-04-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(11,4,1,1,0,1,3750.0000,'2016-05-01 00:00:00','2016-05-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(12,4,1,1,0,1,3750.0000,'2016-06-01 00:00:00','2016-06-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(13,4,1,1,0,1,3750.0000,'2016-07-01 00:00:00','2016-07-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(14,5,1,0,0,0,0.0000,'0000-00-00 00:00:00','0000-00-00 00:00:00',6,4,0,'',1,0,'Prorated: 0 days out of 0','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(15,7,1,3,0,2,4150.0000,'2016-07-01 00:00:00','2016-07-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(16,4,1,1,0,1,3750.0000,'2016-08-01 00:00:00','2016-08-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(17,5,1,0,0,0,0.0000,'0000-00-00 00:00:00','0000-00-00 00:00:00',6,4,0,'',1,0,'Prorated: 0 days out of 0','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(18,7,1,3,0,2,4150.0000,'2016-08-01 00:00:00','2016-08-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(19,4,1,1,0,1,3750.0000,'2016-09-01 00:00:00','2016-09-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(20,5,1,0,0,0,0.0000,'0000-00-00 00:00:00','0000-00-00 00:00:00',6,4,0,'',1,0,'Prorated: 0 days out of 0','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(21,7,1,3,0,2,4150.0000,'2016-09-01 00:00:00','2016-09-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(22,4,1,1,0,1,3750.0000,'2016-10-01 00:00:00','2016-10-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(23,5,1,0,0,0,0.0000,'0000-00-00 00:00:00','0000-00-00 00:00:00',6,4,0,'',1,0,'Prorated: 0 days out of 0','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(24,7,1,3,0,2,4150.0000,'2016-10-01 00:00:00','2016-10-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(25,8,1,2,0,3,4000.0000,'2016-10-01 00:00:00','2016-10-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(26,4,1,1,0,1,3750.0000,'2016-11-01 00:00:00','2016-11-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(27,5,1,0,0,0,0.0000,'0000-00-00 00:00:00','0000-00-00 00:00:00',6,4,0,'',1,0,'Prorated: 0 days out of 0','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(28,7,1,3,0,2,4150.0000,'2016-11-01 00:00:00','2016-11-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(29,8,1,2,0,3,4000.0000,'2016-11-01 00:00:00','2016-11-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(30,4,1,1,0,1,3750.0000,'2016-12-01 00:00:00','2016-12-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(31,5,1,0,0,0,0.0000,'0000-00-00 00:00:00','0000-00-00 00:00:00',6,4,0,'',1,0,'Prorated: 0 days out of 0','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(32,7,1,3,0,2,4150.0000,'2016-12-01 00:00:00','2016-12-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(33,8,1,2,0,3,4000.0000,'2016-12-01 00:00:00','2016-12-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(34,4,1,1,0,1,3750.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(35,5,1,0,0,0,0.0000,'0000-00-00 00:00:00','0000-00-00 00:00:00',6,4,0,'',1,0,'Prorated: 0 days out of 0','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(36,7,1,3,0,2,4150.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(37,8,1,2,0,3,4000.0000,'2017-01-01 00:00:00','2017-01-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(38,4,1,1,0,1,3750.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(39,5,1,0,0,0,0.0000,'0000-00-00 00:00:00','0000-00-00 00:00:00',6,4,0,'',1,0,'Prorated: 0 days out of 0','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(40,7,1,3,0,2,4150.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(41,8,1,2,0,3,4000.0000,'2017-02-01 00:00:00','2017-02-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(42,4,1,1,0,1,3750.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(43,5,1,0,0,0,0.0000,'0000-00-00 00:00:00','0000-00-00 00:00:00',6,4,0,'',1,0,'Prorated: 0 days out of 0','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(44,7,1,3,0,2,4150.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(45,8,1,2,0,3,4000.0000,'2017-03-01 00:00:00','2017-03-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(46,4,1,1,0,1,3750.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(47,5,1,0,0,0,0.0000,'0000-00-00 00:00:00','0000-00-00 00:00:00',6,4,0,'',1,0,'Prorated: 0 days out of 0','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(48,7,1,3,0,2,4150.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(49,8,1,2,0,3,4000.0000,'2017-04-01 00:00:00','2017-04-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(50,4,1,1,0,1,3750.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(51,5,1,0,0,0,0.0000,'0000-00-00 00:00:00','0000-00-00 00:00:00',6,4,0,'',1,0,'Prorated: 0 days out of 0','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(52,7,1,3,0,2,4150.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(53,8,1,2,0,3,4000.0000,'2017-05-01 00:00:00','2017-05-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(54,4,1,1,0,1,3750.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(55,5,1,0,0,0,0.0000,'0000-00-00 00:00:00','0000-00-00 00:00:00',6,4,0,'',1,0,'Prorated: 0 days out of 0','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(56,7,1,3,0,2,4150.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(57,8,1,2,0,3,4000.0000,'2017-06-01 00:00:00','2017-06-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(58,4,1,1,0,1,3750.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(59,5,1,0,0,0,0.0000,'0000-00-00 00:00:00','0000-00-00 00:00:00',6,4,0,'',1,0,'Prorated: 0 days out of 0','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(60,7,1,3,0,2,4150.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(61,8,1,2,0,3,4000.0000,'2017-07-01 00:00:00','2017-07-02 00:00:00',6,4,0,'',1,0,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(62,0,1,2,0,0,175.0000,'2017-02-01 00:00:00','2017-02-01 00:00:00',0,0,0,'',16,0,'','2017-06-22 18:05:22',0,'2017-06-22 18:05:22',0),(63,0,1,2,0,3,15.0000,'2016-11-01 00:00:00','2016-11-01 00:00:00',0,0,0,'',22,0,'incoming wire fee','2017-06-22 18:09:05',0,'2017-06-22 18:09:05',0),(64,0,1,2,0,3,175.0000,'2017-03-01 00:00:00','2017-03-01 00:00:00',0,0,0,'',16,0,'','2017-06-22 18:19:36',0,'2017-06-22 18:19:36',0),(65,0,1,2,0,3,15.0000,'2017-03-01 00:00:00','2017-03-01 00:00:00',0,0,0,'',22,0,'incoming wire fee','2017-06-22 18:20:20',0,'2017-06-22 18:20:20',0),(66,0,1,2,0,3,350.0000,'2017-04-01 00:00:00','2017-04-01 00:00:00',0,0,0,'',16,0,'One Month Utilities Reimbursement','2017-06-22 18:25:20',0,'2017-06-22 18:25:20',0),(67,0,1,2,0,3,81.7900,'2017-04-01 00:00:00','2017-04-01 00:00:00',0,0,0,'',16,0,'Retro Utilities Reimbursement','2017-06-22 18:26:05',0,'2017-06-22 18:26:05',0),(68,0,1,2,0,3,350.0000,'2017-05-01 00:00:00','2017-05-01 00:00:00',0,0,0,'',16,0,'One Month Utilities Reimbursement','2017-06-22 18:27:40',0,'2017-06-22 18:27:40',0),(69,0,1,2,0,3,628.4500,'2017-01-24 00:00:00','2017-01-24 00:00:00',0,0,0,'',16,0,'Utilities SEP-DEC','2017-06-22 18:51:21',0,'2017-06-22 18:51:21',0),(70,0,1,3,0,2,5976.0000,'2016-08-01 00:00:00','2016-08-01 00:00:00',0,0,0,'',12,0,'commission on 2-year lease','2017-06-22 19:01:12',0,'2017-06-22 19:01:12',0);
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
INSERT INTO `Business` VALUES (1,'REX','JGM First, LLC',6,4,4,'2017-06-13 05:39:46',0,'2017-06-14 18:26:46',0);
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
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
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
  `Name` varchar(50) NOT NULL DEFAULT '',
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
  `DID` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `RCPTID` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0'
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Depository`
--

LOCK TABLES `Depository` WRITE;
/*!40000 ALTER TABLE `Depository` DISABLE KEYS */;
INSERT INTO `Depository` VALUES (1,1,1,'First Republic Bank','80001054320','2017-06-13 16:48:45',0,'2017-06-14 18:26:48',0),(2,1,11,'First Republic Bank','80003196953','2017-06-13 16:55:54',0,'2017-06-14 18:26:48',0),(3,1,12,'Petty Cash','','2017-06-13 16:57:23',0,'2017-06-14 18:26:48',0);
/*!40000 ALTER TABLE `Depository` ENABLE KEYS */;
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
  `Type` smallint(6) NOT NULL DEFAULT '0',
  `Name` varchar(100) NOT NULL DEFAULT '',
  `AcctType` varchar(100) NOT NULL DEFAULT '',
  `RAAssociated` smallint(6) NOT NULL DEFAULT '0',
  `AllowPost` smallint(6) NOT NULL DEFAULT '0',
  `RARequired` smallint(6) NOT NULL DEFAULT '0',
  `ManageToBudget` smallint(6) NOT NULL DEFAULT '0',
  `Description` varchar(1024) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`LID`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `GLAccount`
--

LOCK TABLES `GLAccount` WRITE;
/*!40000 ALTER TABLE `GLAccount` DISABLE KEYS */;
INSERT INTO `GLAccount` VALUES (1,9,1,0,0,'10104',2,10,'FRB Operating','Bank',1,1,0,0,'80001054320','2017-06-13 16:45:34',0,'2017-06-14 18:26:48',0),(2,10,1,0,0,'11001',2,11,'Rent Roll Receivables','Accounts Receivables',1,1,0,0,'','2017-06-13 16:47:06',0,'2017-06-14 18:26:48',0),(3,14,1,0,0,'50201',2,12,'Gross Scheduled Rent','Income',1,1,0,0,'','2017-06-13 17:55:08',0,'2017-06-14 18:26:48',0),(4,0,1,0,0,'50255',2,0,'Loss to Lease','Income',2,0,0,0,'','2017-06-13 17:51:24',0,'2017-06-14 18:26:48',0),(5,0,1,0,0,'50251',2,0,'Vacancy','Income',2,0,0,0,'','2017-06-13 17:51:44',0,'2017-06-14 18:26:48',0),(6,0,1,0,0,'',2,15,'Security Deposit Receivable','',0,0,0,0,'','2017-06-13 05:39:46',0,'2017-06-14 18:26:48',0),(7,0,1,0,0,'33000',2,16,'Tenant Security Deposit','Other Current Liability',1,0,0,0,'','2017-06-13 17:11:47',0,'2017-06-14 18:26:48',0),(8,0,1,0,0,'41000',2,0,'Capital Accounts','Equity',1,0,0,0,'','2017-06-13 17:10:00',0,'2017-06-14 18:26:48',0),(9,0,1,0,0,'10000',2,0,'Cash Accounts','Bank',1,1,0,0,'','2017-06-13 16:55:16',0,'2017-06-14 18:26:48',0),(10,0,1,0,0,'11000',2,0,'Accounts Receivable','Accounts Receivable',1,1,0,0,'','2017-06-13 16:46:16',0,'2017-06-14 18:26:48',0),(11,0,1,0,0,'10105',2,0,'FRB Tenant Deposits ','Bank',1,1,0,0,'80003196953','2017-06-13 16:54:29',0,'2017-06-14 18:26:48',0),(12,9,1,0,0,'10100',2,0,'Petty Cash','Bank',1,1,0,0,'','2017-06-13 16:56:47',0,'2017-06-14 18:26:48',0),(13,0,1,0,0,'11099',2,0,'Undeposited Funds','Other Current Asset',1,1,0,0,'','2017-06-13 17:05:05',0,'2017-06-14 18:26:48',0),(14,28,1,0,0,'50200',2,0,'Rental Income','Income',1,1,0,0,'','2017-06-13 18:04:22',0,'2017-06-14 18:26:48',0),(15,14,1,0,0,'50210',2,0,'Additional Rental Income','Income',1,1,0,0,'','2017-06-13 17:56:19',0,'2017-06-14 18:26:48',0),(16,14,1,0,0,'50250',2,0,'Rental Income Offsets','Income',1,1,0,0,'','2017-06-13 17:57:10',0,'2017-06-14 18:26:48',0),(17,0,1,0,0,'50212',2,0,'NSF Fees','Income',2,1,0,0,'','2017-06-13 17:52:27',0,'2017-06-14 18:26:48',0),(18,0,1,0,0,'50213',2,0,'Forfeited Security Deposits','Income',2,1,0,0,'','2017-06-13 17:56:51',0,'2017-06-14 18:26:48',0),(19,0,1,0,0,'50214',2,0,'Insurance Reimbursement','Income',2,1,0,0,'','2017-06-13 17:52:17',0,'2017-06-14 18:26:48',0),(20,0,1,0,0,'50216',2,0,'Maintenance Reimbursements','Income',2,1,0,0,'','2017-06-13 17:52:12',0,'2017-06-14 18:26:48',0),(21,0,1,0,0,'50249',2,0,'Other Rental Income','Income',2,1,0,0,'','2017-06-13 17:52:07',0,'2017-06-14 18:26:48',0),(22,0,1,0,0,'50252',2,0,'Rent Concession','Income',2,1,0,0,'','2017-06-13 17:49:50',0,'2017-06-14 18:26:48',0),(23,0,1,0,0,'50253',2,0,'Offline Units','Income',1,1,0,0,'','2017-06-13 17:59:47',0,'2017-06-14 18:26:48',0),(24,0,1,0,0,'50254',2,0,'Tenant Bad Debt','Income',2,1,0,0,'','2017-06-13 17:50:38',0,'2017-06-14 18:26:48',0),(25,15,1,0,0,'50256',2,0,'Administrative Unit','Income',1,1,0,0,'','2017-06-13 17:43:58',0,'2017-06-14 18:26:48',0),(26,0,1,0,0,'50299',2,0,'Other Income Offsets','Income',2,1,0,0,'','2017-06-13 17:50:04',0,'2017-06-14 18:26:48',0),(27,14,1,0,0,'59998',2,0,'Income Suspense','Income',1,1,0,0,'','2017-06-13 17:46:18',0,'2017-06-14 18:26:48',0),(28,0,1,0,0,'50000',2,0,'Income','Income',1,1,0,0,'','2017-06-13 17:46:49',0,'2017-06-14 18:26:48',0),(30,7,1,0,0,'50001',2,0,'test','bla',2,1,0,0,'','2017-06-13 19:12:13',0,'2017-06-14 18:26:48',0),(31,15,1,0,0,'50211',2,0,'Late Fee','Income',2,1,0,0,'','2017-06-13 19:27:56',0,'2017-06-14 18:26:48',0),(32,0,1,0,0,'70000',2,0,'Business Expenses','Expnese',1,1,0,0,'','2017-06-20 18:11:07',0,'2017-06-20 18:11:07',0),(33,32,1,0,0,'73000',2,0,'Rental Expense','Expense',1,1,0,0,'','2017-06-20 18:11:38',0,'2017-06-20 18:11:38',0),(34,33,1,0,0,'73001',2,0,'Rental Advertising/Commissions','Expense',1,1,0,0,'','2017-06-20 18:12:05',0,'2017-06-20 18:12:05',0),(35,15,1,0,0,'50215',2,0,'Utility Reimbursements','Income',2,1,0,0,'Utility Reimbursements','2017-06-22 16:56:54',0,'2017-06-22 16:56:54',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Journal`
--

LOCK TABLES `Journal` WRITE;
/*!40000 ALTER TABLE `Journal` DISABLE KEYS */;
INSERT INTO `Journal` VALUES (1,1,'2014-03-01 00:00:00',7000.0000,1,1,'','2017-06-20 16:10:25',0,'2017-06-20 16:10:25',0),(2,1,'2016-07-01 00:00:00',8300.0000,1,2,'','2017-06-20 16:23:12',0,'2017-06-20 16:23:12',0),(3,1,'2014-03-01 00:00:00',7000.0000,2,1,'','2017-06-20 18:17:14',0,'2017-06-20 18:17:14',0),(4,1,'2016-03-01 00:00:00',3750.0000,1,9,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(5,1,'2016-04-01 00:00:00',3750.0000,1,10,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(6,1,'2016-05-01 00:00:00',3750.0000,1,11,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(7,1,'2016-06-01 00:00:00',3750.0000,1,12,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(8,1,'2016-07-01 00:00:00',3750.0000,1,13,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(9,1,'2016-07-01 00:00:00',0.0000,1,14,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(10,1,'2016-07-01 00:00:00',4150.0000,1,15,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(11,1,'2016-08-01 00:00:00',3750.0000,1,16,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(12,1,'2016-08-01 00:00:00',0.0000,1,17,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(13,1,'2016-08-01 00:00:00',4150.0000,1,18,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(14,1,'2016-09-01 00:00:00',3750.0000,1,19,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(15,1,'2016-09-01 00:00:00',0.0000,1,20,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(16,1,'2016-09-01 00:00:00',4150.0000,1,21,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(17,1,'2016-10-01 00:00:00',3750.0000,1,22,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(18,1,'2016-10-01 00:00:00',0.0000,1,23,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(19,1,'2016-10-01 00:00:00',4150.0000,1,24,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(20,1,'2016-10-01 00:00:00',4000.0000,1,25,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(21,1,'2016-11-01 00:00:00',3750.0000,1,26,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(22,1,'2016-11-01 00:00:00',0.0000,1,27,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(23,1,'2016-11-01 00:00:00',4150.0000,1,28,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(24,1,'2016-11-01 00:00:00',4000.0000,1,29,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(25,1,'2016-12-01 00:00:00',3750.0000,1,30,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(26,1,'2016-12-01 00:00:00',0.0000,1,31,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(27,1,'2016-12-01 00:00:00',4150.0000,1,32,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(28,1,'2016-12-01 00:00:00',4000.0000,1,33,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(29,1,'2017-01-01 00:00:00',3750.0000,1,34,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(30,1,'2017-01-01 00:00:00',0.0000,1,35,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(31,1,'2017-01-01 00:00:00',4150.0000,1,36,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(32,1,'2017-01-01 00:00:00',4000.0000,1,37,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(33,1,'2017-02-01 00:00:00',3750.0000,1,38,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(34,1,'2017-02-01 00:00:00',0.0000,1,39,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(35,1,'2017-02-01 00:00:00',4150.0000,1,40,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(36,1,'2017-02-01 00:00:00',4000.0000,1,41,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(37,1,'2017-03-01 00:00:00',3750.0000,1,42,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(38,1,'2017-03-01 00:00:00',0.0000,1,43,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(39,1,'2017-03-01 00:00:00',4150.0000,1,44,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(40,1,'2017-03-01 00:00:00',4000.0000,1,45,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(41,1,'2017-04-01 00:00:00',3750.0000,1,46,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(42,1,'2017-04-01 00:00:00',0.0000,1,47,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(43,1,'2017-04-01 00:00:00',4150.0000,1,48,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(44,1,'2017-04-01 00:00:00',4000.0000,1,49,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(45,1,'2017-05-01 00:00:00',3750.0000,1,50,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(46,1,'2017-05-01 00:00:00',0.0000,1,51,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(47,1,'2017-05-01 00:00:00',4150.0000,1,52,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(48,1,'2017-05-01 00:00:00',4000.0000,1,53,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(49,1,'2017-06-01 00:00:00',3750.0000,1,54,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(50,1,'2017-06-01 00:00:00',0.0000,1,55,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(51,1,'2017-06-01 00:00:00',4150.0000,1,56,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(52,1,'2017-06-01 00:00:00',4000.0000,1,57,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(53,1,'2017-07-01 00:00:00',3750.0000,1,58,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(54,1,'2017-07-01 00:00:00',0.0000,1,59,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(55,1,'2017-07-01 00:00:00',4150.0000,1,60,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(56,1,'2017-07-01 00:00:00',4000.0000,1,61,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(57,1,'2017-06-01 00:00:00',4150.0000,2,2,'','2017-06-22 17:20:34',0,'2017-06-22 17:20:34',0),(58,1,'2017-06-01 00:00:00',3750.0000,2,3,'','2017-06-22 17:21:37',0,'2017-06-22 17:21:37',0),(59,1,'2016-03-01 00:00:00',3750.0000,2,4,'','2017-06-22 17:26:32',0,'2017-06-22 17:26:32',0),(60,1,'2016-04-01 00:00:00',3750.0000,2,5,'','2017-06-22 17:27:11',0,'2017-06-22 17:27:11',0),(61,1,'2016-05-01 00:00:00',3750.0000,2,6,'','2017-06-22 17:27:52',0,'2017-06-22 17:27:52',0),(62,1,'2016-06-01 00:00:00',6750.0000,2,7,'','2017-06-22 17:29:21',0,'2017-06-22 17:29:21',0),(63,1,'2016-07-01 00:00:00',3750.0000,2,8,'','2017-06-22 17:30:18',0,'2017-06-22 17:30:18',0),(64,1,'2017-08-01 00:00:00',3750.0000,2,9,'','2017-06-22 17:30:44',0,'2017-06-22 17:30:44',0),(65,1,'2016-08-01 00:00:00',3750.0000,2,10,'','2017-06-22 17:31:20',0,'2017-06-22 17:31:20',0),(66,1,'2016-09-01 00:00:00',3750.0000,2,11,'','2017-06-22 17:32:16',0,'2017-06-22 17:32:16',0),(67,1,'2016-10-01 00:00:00',3750.0000,2,12,'','2017-06-22 17:32:37',0,'2017-06-22 17:32:37',0),(68,1,'2017-06-01 00:00:00',3750.0000,2,13,'','2017-06-22 17:33:06',0,'2017-06-22 17:33:06',0),(69,1,'2016-11-01 00:00:00',3750.0000,2,14,'','2017-06-22 17:34:00',0,'2017-06-22 17:34:00',0),(70,1,'2016-12-01 00:00:00',3750.0000,2,15,'','2017-06-22 17:34:29',0,'2017-06-22 17:34:29',0),(71,1,'2017-01-01 00:00:00',3750.0000,2,16,'','2017-06-22 17:35:16',0,'2017-06-22 17:35:16',0),(72,1,'2017-02-01 00:00:00',3750.0000,2,17,'','2017-06-22 17:35:40',0,'2017-06-22 17:35:40',0),(73,1,'2017-03-01 00:00:00',3750.0000,2,18,'','2017-06-22 17:36:00',0,'2017-06-22 17:36:00',0),(74,1,'2017-04-01 00:00:00',3750.0000,2,19,'','2017-06-22 17:36:24',0,'2017-06-22 17:36:24',0),(75,1,'2017-05-01 00:00:00',3750.0000,2,20,'','2017-06-22 17:36:56',0,'2017-06-22 17:36:56',0),(76,1,'2016-08-01 00:00:00',4150.0000,2,21,'','2017-06-22 17:48:45',0,'2017-06-22 17:48:45',0),(77,1,'2016-09-01 00:00:00',4150.0000,2,22,'','2017-06-22 17:49:15',0,'2017-06-22 17:49:15',0),(78,1,'2016-10-01 00:00:00',4150.0000,2,23,'','2017-06-22 17:50:25',0,'2017-06-22 17:50:25',0),(79,1,'2016-11-01 00:00:00',4150.0000,2,24,'','2017-06-22 17:51:26',0,'2017-06-22 17:51:26',0),(80,1,'2016-12-01 00:00:00',4150.0000,2,25,'','2017-06-22 17:51:52',0,'2017-06-22 17:51:52',0),(81,1,'2017-01-01 00:00:00',4150.0000,2,26,'','2017-06-22 17:52:26',0,'2017-06-22 17:52:26',0),(82,1,'2017-02-01 00:00:00',4150.0000,2,27,'','2017-06-22 17:52:52',0,'2017-06-22 17:52:52',0),(83,1,'2017-03-01 00:00:00',4150.0000,2,28,'','2017-06-22 17:53:32',0,'2017-06-22 17:53:32',0),(84,1,'2017-04-01 00:00:00',4150.0000,2,29,'','2017-06-22 17:54:12',0,'2017-06-22 17:54:12',0),(85,1,'2017-05-01 00:00:00',4150.0000,2,30,'','2017-06-22 17:54:46',0,'2017-06-22 17:54:46',0),(86,1,'2016-10-03 00:00:00',4000.0000,2,31,'','2017-06-22 18:00:00',0,'2017-06-22 18:00:00',0),(87,1,'2016-11-10 00:00:00',12000.0000,2,32,'','2017-06-22 18:01:58',0,'2017-06-22 18:01:58',0),(88,1,'2017-02-01 00:00:00',175.0000,1,62,'','2017-06-22 18:05:22',0,'2017-06-22 18:05:22',0),(89,1,'2016-11-01 00:00:00',15.0000,1,63,'','2017-06-22 18:09:05',0,'2017-06-22 18:09:05',0),(90,1,'2017-03-01 00:00:00',175.0000,1,64,'','2017-06-22 18:19:36',0,'2017-06-22 18:19:36',0),(91,1,'2017-03-01 00:00:00',15.0000,1,65,'','2017-06-22 18:20:20',0,'2017-06-22 18:20:20',0),(92,1,'2017-02-13 00:00:00',8335.0000,2,33,'','2017-06-22 18:22:37',0,'2017-06-22 18:22:37',0),(93,1,'2017-04-01 00:00:00',350.0000,1,66,'','2017-06-22 18:25:20',0,'2017-06-22 18:25:20',0),(94,1,'2017-04-01 00:00:00',81.7900,1,67,'','2017-06-22 18:26:05',0,'2017-06-22 18:26:05',0),(95,1,'2017-05-01 00:00:00',350.0000,1,68,'','2017-06-22 18:27:40',0,'2017-06-22 18:27:40',0),(96,1,'2017-05-12 00:00:00',13116.7900,2,34,'','2017-06-22 18:28:38',0,'2017-06-22 18:28:38',0),(97,1,'2017-01-24 00:00:00',628.4500,1,69,'','2017-06-22 18:51:21',0,'2017-06-22 18:51:21',0),(98,1,'2017-02-02 00:00:00',628.4500,2,35,'','2017-06-22 18:52:14',0,'2017-06-22 18:52:14',0),(99,1,'2016-08-01 00:00:00',5976.0000,1,70,'','2017-06-22 19:01:12',0,'2017-06-22 19:01:12',0),(100,1,'2016-07-01 00:00:00',6474.0000,2,36,'','2017-06-22 19:05:23',0,'2017-06-22 19:05:23',0);
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
  `Amount` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `ASMID` bigint(20) NOT NULL DEFAULT '0',
  `AcctRule` varchar(200) NOT NULL DEFAULT '',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`JAID`)
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `JournalAllocation`
--

LOCK TABLES `JournalAllocation` WRITE;
/*!40000 ALTER TABLE `JournalAllocation` DISABLE KEYS */;
INSERT INTO `JournalAllocation` VALUES (1,1,1,1,1,0,7000.0000,1,'d 10104 7000.00, c 33000 7000.00','2017-06-20 16:10:25',0),(2,1,2,3,3,0,8300.0000,2,'d 10104 8300.00, c 33000 8300.00','2017-06-20 16:23:12',0),(3,1,4,1,1,0,3750.0000,9,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:58',0),(4,1,5,1,1,0,3750.0000,10,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:58',0),(5,1,6,1,1,0,3750.0000,11,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:58',0),(6,1,7,1,1,0,3750.0000,12,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:58',0),(7,1,8,1,1,0,3750.0000,13,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:58',0),(8,1,9,0,0,0,0.0000,14,'d 11001 0.00, c 50201 0.00','2017-06-21 19:40:58',0),(9,1,10,3,2,0,4150.0000,15,'d 11001 4150.00, c 50201 4150.00','2017-06-21 19:40:58',0),(10,1,11,1,1,0,3750.0000,16,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:58',0),(11,1,12,0,0,0,0.0000,17,'d 11001 0.00, c 50201 0.00','2017-06-21 19:40:58',0),(12,1,13,3,2,0,4150.0000,18,'d 11001 4150.00, c 50201 4150.00','2017-06-21 19:40:58',0),(13,1,14,1,1,0,3750.0000,19,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:58',0),(14,1,15,0,0,0,0.0000,20,'d 11001 0.00, c 50201 0.00','2017-06-21 19:40:58',0),(15,1,16,3,2,0,4150.0000,21,'d 11001 4150.00, c 50201 4150.00','2017-06-21 19:40:58',0),(16,1,17,1,1,0,3750.0000,22,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:58',0),(17,1,18,0,0,0,0.0000,23,'d 11001 0.00, c 50201 0.00','2017-06-21 19:40:58',0),(18,1,19,3,2,0,4150.0000,24,'d 11001 4150.00, c 50201 4150.00','2017-06-21 19:40:58',0),(19,1,20,2,3,0,4000.0000,25,'d 11001 4000.00, c 50201 4000.00','2017-06-21 19:40:58',0),(20,1,21,1,1,0,3750.0000,26,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:58',0),(21,1,22,0,0,0,0.0000,27,'d 11001 0.00, c 50201 0.00','2017-06-21 19:40:58',0),(22,1,23,3,2,0,4150.0000,28,'d 11001 4150.00, c 50201 4150.00','2017-06-21 19:40:58',0),(23,1,24,2,3,0,4000.0000,29,'d 11001 4000.00, c 50201 4000.00','2017-06-21 19:40:58',0),(24,1,25,1,1,0,3750.0000,30,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:58',0),(25,1,26,0,0,0,0.0000,31,'d 11001 0.00, c 50201 0.00','2017-06-21 19:40:58',0),(26,1,27,3,2,0,4150.0000,32,'d 11001 4150.00, c 50201 4150.00','2017-06-21 19:40:58',0),(27,1,28,2,3,0,4000.0000,33,'d 11001 4000.00, c 50201 4000.00','2017-06-21 19:40:58',0),(28,1,29,1,1,0,3750.0000,34,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:59',0),(29,1,30,0,0,0,0.0000,35,'d 11001 0.00, c 50201 0.00','2017-06-21 19:40:59',0),(30,1,31,3,2,0,4150.0000,36,'d 11001 4150.00, c 50201 4150.00','2017-06-21 19:40:59',0),(31,1,32,2,3,0,4000.0000,37,'d 11001 4000.00, c 50201 4000.00','2017-06-21 19:40:59',0),(32,1,33,1,1,0,3750.0000,38,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:59',0),(33,1,34,0,0,0,0.0000,39,'d 11001 0.00, c 50201 0.00','2017-06-21 19:40:59',0),(34,1,35,3,2,0,4150.0000,40,'d 11001 4150.00, c 50201 4150.00','2017-06-21 19:40:59',0),(35,1,36,2,3,0,4000.0000,41,'d 11001 4000.00, c 50201 4000.00','2017-06-21 19:40:59',0),(36,1,37,1,1,0,3750.0000,42,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:59',0),(37,1,38,0,0,0,0.0000,43,'d 11001 0.00, c 50201 0.00','2017-06-21 19:40:59',0),(38,1,39,3,2,0,4150.0000,44,'d 11001 4150.00, c 50201 4150.00','2017-06-21 19:40:59',0),(39,1,40,2,3,0,4000.0000,45,'d 11001 4000.00, c 50201 4000.00','2017-06-21 19:40:59',0),(40,1,41,1,1,0,3750.0000,46,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:59',0),(41,1,42,0,0,0,0.0000,47,'d 11001 0.00, c 50201 0.00','2017-06-21 19:40:59',0),(42,1,43,3,2,0,4150.0000,48,'d 11001 4150.00, c 50201 4150.00','2017-06-21 19:40:59',0),(43,1,44,2,3,0,4000.0000,49,'d 11001 4000.00, c 50201 4000.00','2017-06-21 19:40:59',0),(44,1,45,1,1,0,3750.0000,50,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:59',0),(45,1,46,0,0,0,0.0000,51,'d 11001 0.00, c 50201 0.00','2017-06-21 19:40:59',0),(46,1,47,3,2,0,4150.0000,52,'d 11001 4150.00, c 50201 4150.00','2017-06-21 19:40:59',0),(47,1,48,2,3,0,4000.0000,53,'d 11001 4000.00, c 50201 4000.00','2017-06-21 19:40:59',0),(48,1,49,1,1,0,3750.0000,54,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:59',0),(49,1,50,0,0,0,0.0000,55,'d 11001 0.00, c 50201 0.00','2017-06-21 19:40:59',0),(50,1,51,3,2,0,4150.0000,56,'d 11001 4150.00, c 50201 4150.00','2017-06-21 19:40:59',0),(51,1,52,2,3,0,4000.0000,57,'d 11001 4000.00, c 50201 4000.00','2017-06-21 19:40:59',0),(52,1,53,1,1,0,3750.0000,58,'d 11001 3750.00, c 50201 3750.00','2017-06-21 19:40:59',0),(53,1,54,0,0,0,0.0000,59,'d 11001 0.00, c 50201 0.00','2017-06-21 19:40:59',0),(54,1,55,3,2,0,4150.0000,60,'d 11001 4150.00, c 50201 4150.00','2017-06-21 19:40:59',0),(55,1,56,2,3,0,4000.0000,61,'d 11001 4000.00, c 50201 4000.00','2017-06-21 19:40:59',0),(56,1,88,2,0,0,175.0000,62,'d 11001 175.00, c 50215 175.00','2017-06-22 18:05:22',0),(57,1,89,2,3,0,15.0000,63,'d 11001 15.00, c 50254 15.00','2017-06-22 18:09:05',0),(58,1,90,2,3,0,175.0000,64,'d 11001 175.00, c 50215 175.00','2017-06-22 18:19:36',0),(59,1,91,2,3,0,15.0000,65,'d 11001 15.00, c 50254 15.00','2017-06-22 18:20:20',0),(60,1,93,2,3,0,350.0000,66,'d 11001 350.00, c 50215 350.00','2017-06-22 18:25:20',0),(61,1,94,2,3,0,81.7900,67,'d 11001 81.79, c 50215 81.79','2017-06-22 18:26:05',0),(62,1,95,2,3,0,350.0000,68,'d 11001 350.00, c 50215 350.00','2017-06-22 18:27:40',0),(63,1,97,2,3,0,628.4500,69,'d 11001 628.45, c 50215 628.45','2017-06-22 18:51:21',0),(64,1,99,3,2,0,5976.0000,70,'d 11001 5976.00, c 73001 5976.00','2017-06-22 19:01:12',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `JournalMarker`
--

LOCK TABLES `JournalMarker` WRITE;
/*!40000 ALTER TABLE `JournalMarker` DISABLE KEYS */;
INSERT INTO `JournalMarker` VALUES (1,1,0,'2016-03-01 00:00:00','2016-04-01 00:00:00','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(2,1,0,'2016-04-01 00:00:00','2016-05-01 00:00:00','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(3,1,0,'2016-05-01 00:00:00','2016-06-01 00:00:00','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(4,1,0,'2016-06-01 00:00:00','2016-07-01 00:00:00','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(5,1,0,'2016-07-01 00:00:00','2016-08-01 00:00:00','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(6,1,0,'2016-08-01 00:00:00','2016-09-01 00:00:00','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(7,1,0,'2016-09-01 00:00:00','2016-10-01 00:00:00','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(8,1,0,'2016-10-01 00:00:00','2016-11-01 00:00:00','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(9,1,0,'2016-11-01 00:00:00','2016-12-01 00:00:00','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(10,1,0,'2016-12-01 00:00:00','2017-01-01 00:00:00','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(11,1,0,'2017-01-01 00:00:00','2017-02-01 00:00:00','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(12,1,0,'2017-02-01 00:00:00','2017-03-01 00:00:00','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(13,1,0,'2017-03-01 00:00:00','2017-04-01 00:00:00','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(14,1,0,'2017-04-01 00:00:00','2017-05-01 00:00:00','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(15,1,0,'2017-05-01 00:00:00','2017-06-01 00:00:00','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(16,1,0,'2017-06-01 00:00:00','2017-07-01 00:00:00','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(17,1,0,'2017-07-01 00:00:00','2017-08-01 00:00:00','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=105 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerEntry`
--

LOCK TABLES `LedgerEntry` WRITE;
/*!40000 ALTER TABLE `LedgerEntry` DISABLE KEYS */;
INSERT INTO `LedgerEntry` VALUES (1,1,1,1,1,1,1,0,'2014-03-01 00:00:00',7000.0000,'','2017-06-20 16:10:25',0,'2017-06-20 16:10:25',0),(2,1,1,1,7,1,1,0,'2014-03-01 00:00:00',-7000.0000,'','2017-06-20 16:10:25',0,'2017-06-20 16:10:25',0),(5,1,4,3,2,1,1,0,'2016-03-01 00:00:00',3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(6,1,4,3,3,1,1,0,'2016-03-01 00:00:00',-3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(7,1,5,4,2,1,1,0,'2016-04-01 00:00:00',3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(8,1,5,4,3,1,1,0,'2016-04-01 00:00:00',-3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(9,1,6,5,2,1,1,0,'2016-05-01 00:00:00',3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(10,1,6,5,3,1,1,0,'2016-05-01 00:00:00',-3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(11,1,7,6,2,1,1,0,'2016-06-01 00:00:00',3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(12,1,7,6,3,1,1,0,'2016-06-01 00:00:00',-3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(13,1,2,2,1,3,3,0,'2016-07-01 00:00:00',8300.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(14,1,2,2,7,3,3,0,'2016-07-01 00:00:00',-8300.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(15,1,8,7,2,1,1,0,'2016-07-01 00:00:00',3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(16,1,8,7,3,1,1,0,'2016-07-01 00:00:00',-3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(17,1,10,9,2,2,3,0,'2016-07-01 00:00:00',4150.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(18,1,10,9,3,2,3,0,'2016-07-01 00:00:00',-4150.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(19,1,11,10,2,1,1,0,'2016-08-01 00:00:00',3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(20,1,11,10,3,1,1,0,'2016-08-01 00:00:00',-3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(21,1,13,12,2,2,3,0,'2016-08-01 00:00:00',4150.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(22,1,13,12,3,2,3,0,'2016-08-01 00:00:00',-4150.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(23,1,14,13,2,1,1,0,'2016-09-01 00:00:00',3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(24,1,14,13,3,1,1,0,'2016-09-01 00:00:00',-3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(25,1,16,15,2,2,3,0,'2016-09-01 00:00:00',4150.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(26,1,16,15,3,2,3,0,'2016-09-01 00:00:00',-4150.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(27,1,17,16,2,1,1,0,'2016-10-01 00:00:00',3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(28,1,17,16,3,1,1,0,'2016-10-01 00:00:00',-3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(29,1,19,18,2,2,3,0,'2016-10-01 00:00:00',4150.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(30,1,19,18,3,2,3,0,'2016-10-01 00:00:00',-4150.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(31,1,20,19,2,3,2,0,'2016-10-01 00:00:00',4000.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(32,1,20,19,3,3,2,0,'2016-10-01 00:00:00',-4000.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(33,1,21,20,2,1,1,0,'2016-11-01 00:00:00',3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(34,1,21,20,3,1,1,0,'2016-11-01 00:00:00',-3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(35,1,23,22,2,2,3,0,'2016-11-01 00:00:00',4150.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(36,1,23,22,3,2,3,0,'2016-11-01 00:00:00',-4150.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(37,1,24,23,2,3,2,0,'2016-11-01 00:00:00',4000.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(38,1,24,23,3,3,2,0,'2016-11-01 00:00:00',-4000.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(39,1,25,24,2,1,1,0,'2016-12-01 00:00:00',3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(40,1,25,24,3,1,1,0,'2016-12-01 00:00:00',-3750.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(41,1,27,26,2,2,3,0,'2016-12-01 00:00:00',4150.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(42,1,27,26,3,2,3,0,'2016-12-01 00:00:00',-4150.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(43,1,28,27,2,3,2,0,'2016-12-01 00:00:00',4000.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(44,1,28,27,3,3,2,0,'2016-12-01 00:00:00',-4000.0000,'','2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(45,1,29,28,2,1,1,0,'2017-01-01 00:00:00',3750.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(46,1,29,28,3,1,1,0,'2017-01-01 00:00:00',-3750.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(47,1,31,30,2,2,3,0,'2017-01-01 00:00:00',4150.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(48,1,31,30,3,2,3,0,'2017-01-01 00:00:00',-4150.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(49,1,32,31,2,3,2,0,'2017-01-01 00:00:00',4000.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(50,1,32,31,3,3,2,0,'2017-01-01 00:00:00',-4000.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(51,1,33,32,2,1,1,0,'2017-02-01 00:00:00',3750.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(52,1,33,32,3,1,1,0,'2017-02-01 00:00:00',-3750.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(53,1,35,34,2,2,3,0,'2017-02-01 00:00:00',4150.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(54,1,35,34,3,2,3,0,'2017-02-01 00:00:00',-4150.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(55,1,36,35,2,3,2,0,'2017-02-01 00:00:00',4000.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(56,1,36,35,3,3,2,0,'2017-02-01 00:00:00',-4000.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(57,1,37,36,2,1,1,0,'2017-03-01 00:00:00',3750.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(58,1,37,36,3,1,1,0,'2017-03-01 00:00:00',-3750.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(59,1,39,38,2,2,3,0,'2017-03-01 00:00:00',4150.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(60,1,39,38,3,2,3,0,'2017-03-01 00:00:00',-4150.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(61,1,40,39,2,3,2,0,'2017-03-01 00:00:00',4000.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(62,1,40,39,3,3,2,0,'2017-03-01 00:00:00',-4000.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(63,1,41,40,2,1,1,0,'2017-04-01 00:00:00',3750.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(64,1,41,40,3,1,1,0,'2017-04-01 00:00:00',-3750.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(65,1,43,42,2,2,3,0,'2017-04-01 00:00:00',4150.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(66,1,43,42,3,2,3,0,'2017-04-01 00:00:00',-4150.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(67,1,44,43,2,3,2,0,'2017-04-01 00:00:00',4000.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(68,1,44,43,3,3,2,0,'2017-04-01 00:00:00',-4000.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(69,1,45,44,2,1,1,0,'2017-05-01 00:00:00',3750.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(70,1,45,44,3,1,1,0,'2017-05-01 00:00:00',-3750.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(71,1,47,46,2,2,3,0,'2017-05-01 00:00:00',4150.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(72,1,47,46,3,2,3,0,'2017-05-01 00:00:00',-4150.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(73,1,48,47,2,3,2,0,'2017-05-01 00:00:00',4000.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(74,1,48,47,3,3,2,0,'2017-05-01 00:00:00',-4000.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(75,1,49,48,2,1,1,0,'2017-06-01 00:00:00',3750.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(76,1,49,48,3,1,1,0,'2017-06-01 00:00:00',-3750.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(77,1,51,50,2,2,3,0,'2017-06-01 00:00:00',4150.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(78,1,51,50,3,2,3,0,'2017-06-01 00:00:00',-4150.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(79,1,52,51,2,3,2,0,'2017-06-01 00:00:00',4000.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(80,1,52,51,3,3,2,0,'2017-06-01 00:00:00',-4000.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(81,1,53,52,2,1,1,0,'2017-07-01 00:00:00',3750.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(82,1,53,52,3,1,1,0,'2017-07-01 00:00:00',-3750.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(83,1,55,54,2,2,3,0,'2017-07-01 00:00:00',4150.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(84,1,55,54,3,2,3,0,'2017-07-01 00:00:00',-4150.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(85,1,56,55,2,3,2,0,'2017-07-01 00:00:00',4000.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(86,1,56,55,3,3,2,0,'2017-07-01 00:00:00',-4000.0000,'','2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(87,1,88,56,2,0,2,0,'2017-02-01 00:00:00',175.0000,'','2017-06-22 18:05:22',0,'2017-06-22 18:05:22',0),(88,1,88,56,35,0,2,0,'2017-02-01 00:00:00',-175.0000,'','2017-06-22 18:05:22',0,'2017-06-22 18:05:22',0),(89,1,89,57,2,3,2,0,'2016-11-01 00:00:00',15.0000,'','2017-06-22 18:09:05',0,'2017-06-22 18:09:05',0),(90,1,89,57,24,3,2,0,'2016-11-01 00:00:00',-15.0000,'','2017-06-22 18:09:05',0,'2017-06-22 18:09:05',0),(91,1,90,58,2,3,2,0,'2017-03-01 00:00:00',175.0000,'','2017-06-22 18:19:36',0,'2017-06-22 18:19:36',0),(92,1,90,58,35,3,2,0,'2017-03-01 00:00:00',-175.0000,'','2017-06-22 18:19:36',0,'2017-06-22 18:19:36',0),(93,1,91,59,2,3,2,0,'2017-03-01 00:00:00',15.0000,'','2017-06-22 18:20:20',0,'2017-06-22 18:20:20',0),(94,1,91,59,24,3,2,0,'2017-03-01 00:00:00',-15.0000,'','2017-06-22 18:20:20',0,'2017-06-22 18:20:20',0),(95,1,93,60,2,3,2,0,'2017-04-01 00:00:00',350.0000,'','2017-06-22 18:25:20',0,'2017-06-22 18:25:20',0),(96,1,93,60,35,3,2,0,'2017-04-01 00:00:00',-350.0000,'','2017-06-22 18:25:20',0,'2017-06-22 18:25:20',0),(97,1,94,61,2,3,2,0,'2017-04-01 00:00:00',81.7900,'','2017-06-22 18:26:05',0,'2017-06-22 18:26:05',0),(98,1,94,61,35,3,2,0,'2017-04-01 00:00:00',-81.7900,'','2017-06-22 18:26:05',0,'2017-06-22 18:26:05',0),(99,1,95,62,2,3,2,0,'2017-05-01 00:00:00',350.0000,'','2017-06-22 18:27:40',0,'2017-06-22 18:27:40',0),(100,1,95,62,35,3,2,0,'2017-05-01 00:00:00',-350.0000,'','2017-06-22 18:27:40',0,'2017-06-22 18:27:40',0),(101,1,97,63,2,3,2,0,'2017-01-24 00:00:00',628.4500,'','2017-06-22 18:51:21',0,'2017-06-22 18:51:21',0),(102,1,97,63,35,3,2,0,'2017-01-24 00:00:00',-628.4500,'','2017-06-22 18:51:21',0,'2017-06-22 18:51:21',0),(103,1,99,64,2,2,3,0,'2016-08-01 00:00:00',5976.0000,'','2017-06-22 19:01:12',0,'2017-06-22 19:01:12',0),(104,1,99,64,34,2,3,0,'2016-08-01 00:00:00',-5976.0000,'','2017-06-22 19:01:12',0,'2017-06-22 19:01:12',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=596 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerMarker`
--

LOCK TABLES `LedgerMarker` WRITE;
/*!40000 ALTER TABLE `LedgerMarker` DISABLE KEYS */;
INSERT INTO `LedgerMarker` VALUES (1,1,1,0,0,0,'1999-12-31 00:00:00',0.0000,3,'2017-06-13 05:39:46',0,'2017-06-14 18:26:49',0),(2,2,1,0,0,0,'1999-12-31 00:00:00',0.0000,3,'2017-06-13 05:39:46',0,'2017-06-14 18:26:49',0),(3,3,1,0,0,0,'1999-12-31 00:00:00',0.0000,3,'2017-06-13 05:39:46',0,'2017-06-14 18:26:49',0),(4,4,1,0,0,0,'1999-12-31 00:00:00',0.0000,3,'2017-06-13 05:39:46',0,'2017-06-14 18:26:49',0),(5,5,1,0,0,0,'1999-12-31 00:00:00',0.0000,3,'2017-06-13 05:39:46',0,'2017-06-14 18:26:49',0),(6,6,1,0,0,0,'1999-12-31 00:00:00',0.0000,3,'2017-06-13 05:39:46',0,'2017-06-14 18:26:49',0),(7,7,1,0,0,0,'1999-12-31 00:00:00',0.0000,3,'2017-06-13 05:39:46',0,'2017-06-14 18:26:49',0),(8,8,1,0,0,0,'1999-12-31 00:00:00',0.0000,3,'2017-06-13 05:39:46',0,'2017-06-14 18:26:49',0),(9,9,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(10,10,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(11,11,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(12,12,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(13,13,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(14,14,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(15,15,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(16,16,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(17,17,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(18,18,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(19,19,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(20,20,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(21,21,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(22,22,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(23,23,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(24,24,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(25,25,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(26,26,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(27,27,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(28,28,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(29,29,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(30,30,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(31,31,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(32,32,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(33,33,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(34,34,1,0,0,0,'1970-01-01 00:00:00',0.0000,3,'2017-06-21 19:40:42',0,'2017-06-21 19:40:42',0),(35,6,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(36,9,1,0,0,0,'2016-04-01 00:00:00',7000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(37,12,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(38,1,1,0,0,0,'2016-04-01 00:00:00',7000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(39,11,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(40,10,1,0,0,0,'2016-04-01 00:00:00',3750.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(41,2,1,0,0,0,'2016-04-01 00:00:00',3750.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(42,13,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(43,7,1,0,0,0,'2016-04-01 00:00:00',-7000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(44,8,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(45,28,1,0,0,0,'2016-04-01 00:00:00',-3750.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(46,30,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(47,14,1,0,0,0,'2016-04-01 00:00:00',-3750.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(48,3,1,0,0,0,'2016-04-01 00:00:00',-3750.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(49,15,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(50,31,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(51,17,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(52,18,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(53,19,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(54,20,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(55,21,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(56,16,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(57,5,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(58,22,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(59,23,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(60,24,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(61,4,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(62,25,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(63,26,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(64,27,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(65,32,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(66,33,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(67,34,1,0,0,0,'2016-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(68,6,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(69,9,1,0,0,0,'2016-05-01 00:00:00',14000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(70,12,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(71,1,1,0,0,0,'2016-05-01 00:00:00',7000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(72,11,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(73,10,1,0,0,0,'2016-05-01 00:00:00',11250.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(74,2,1,0,0,0,'2016-05-01 00:00:00',7500.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(75,13,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(76,7,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(77,8,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(78,28,1,0,0,0,'2016-05-01 00:00:00',-15000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(79,30,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(80,14,1,0,0,0,'2016-05-01 00:00:00',-11250.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(81,3,1,0,0,0,'2016-05-01 00:00:00',-7500.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(82,15,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(83,31,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(84,17,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(85,18,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(86,19,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(87,20,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(88,21,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(89,16,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(90,5,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(91,22,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(92,23,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(93,24,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(94,4,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(95,25,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(96,26,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(97,27,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(98,32,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(99,33,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(100,34,1,0,0,0,'2016-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(101,6,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(102,9,1,0,0,0,'2016-06-01 00:00:00',21000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(103,12,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(104,1,1,0,0,0,'2016-06-01 00:00:00',7000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(105,11,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(106,10,1,0,0,0,'2016-06-01 00:00:00',22500.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(107,2,1,0,0,0,'2016-06-01 00:00:00',11250.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(108,13,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(109,7,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(110,8,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(111,28,1,0,0,0,'2016-06-01 00:00:00',-37500.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(112,30,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(113,14,1,0,0,0,'2016-06-01 00:00:00',-22500.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(114,3,1,0,0,0,'2016-06-01 00:00:00',-11250.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(115,15,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(116,31,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(117,17,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(118,18,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(119,19,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(120,20,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(121,21,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(122,16,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(123,5,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(124,22,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(125,23,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(126,24,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(127,4,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(128,25,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(129,26,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(130,27,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(131,32,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(132,33,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(133,34,1,0,0,0,'2016-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(134,6,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(135,9,1,0,0,0,'2016-07-01 00:00:00',28000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(136,12,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(137,1,1,0,0,0,'2016-07-01 00:00:00',7000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(138,11,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(139,10,1,0,0,0,'2016-07-01 00:00:00',37500.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(140,2,1,0,0,0,'2016-07-01 00:00:00',15000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(141,13,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(142,7,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(143,8,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(144,28,1,0,0,0,'2016-07-01 00:00:00',-75000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(145,30,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(146,14,1,0,0,0,'2016-07-01 00:00:00',-37500.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(147,3,1,0,0,0,'2016-07-01 00:00:00',-15000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(148,15,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(149,31,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(150,17,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(151,18,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(152,19,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(153,20,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(154,21,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(155,16,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(156,5,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(157,22,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(158,23,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(159,24,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(160,4,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(161,25,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(162,26,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(163,27,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(164,32,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(165,33,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(166,34,1,0,0,0,'2016-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(167,6,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(168,9,1,0,0,0,'2016-08-01 00:00:00',43300.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(169,12,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(170,1,1,0,0,0,'2016-08-01 00:00:00',15300.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(171,11,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(172,10,1,0,0,0,'2016-08-01 00:00:00',60400.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(173,2,1,0,0,0,'2016-08-01 00:00:00',22900.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(174,13,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(175,7,1,0,0,0,'2016-08-01 00:00:00',-8300.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(176,8,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(177,28,1,0,0,0,'2016-08-01 00:00:00',-135400.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(178,30,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(179,14,1,0,0,0,'2016-08-01 00:00:00',-60400.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(180,3,1,0,0,0,'2016-08-01 00:00:00',-22900.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(181,15,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(182,31,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(183,17,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(184,18,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(185,19,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(186,20,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(187,21,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(188,16,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(189,5,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(190,22,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(191,23,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(192,24,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(193,4,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(194,25,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(195,26,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(196,27,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(197,32,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(198,33,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(199,34,1,0,0,0,'2016-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(200,6,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(201,9,1,0,0,0,'2016-09-01 00:00:00',58600.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(202,12,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(203,1,1,0,0,0,'2016-09-01 00:00:00',15300.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(204,11,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(205,10,1,0,0,0,'2016-09-01 00:00:00',91200.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(206,2,1,0,0,0,'2016-09-01 00:00:00',30800.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(207,13,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(208,7,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(209,8,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(210,28,1,0,0,0,'2016-09-01 00:00:00',-226600.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(211,30,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(212,14,1,0,0,0,'2016-09-01 00:00:00',-91200.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(213,3,1,0,0,0,'2016-09-01 00:00:00',-30800.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(214,15,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(215,31,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(216,17,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(217,18,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(218,19,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(219,20,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(220,21,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(221,16,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(222,5,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(223,22,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(224,23,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(225,24,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(226,4,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(227,25,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(228,26,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(229,27,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(230,32,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(231,33,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(232,34,1,0,0,0,'2016-09-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(233,6,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(234,9,1,0,0,0,'2016-10-01 00:00:00',73900.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(235,12,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(236,1,1,0,0,0,'2016-10-01 00:00:00',15300.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(237,11,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(238,10,1,0,0,0,'2016-10-01 00:00:00',129900.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(239,2,1,0,0,0,'2016-10-01 00:00:00',38700.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(240,13,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(241,7,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(242,8,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(243,28,1,0,0,0,'2016-10-01 00:00:00',-356500.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(244,30,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(245,14,1,0,0,0,'2016-10-01 00:00:00',-129900.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(246,3,1,0,0,0,'2016-10-01 00:00:00',-38700.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(247,15,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(248,31,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(249,17,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(250,18,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(251,19,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(252,20,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(253,21,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(254,16,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(255,5,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(256,22,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(257,23,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(258,24,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(259,4,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(260,25,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(261,26,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(262,27,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(263,32,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(264,33,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(265,34,1,0,0,0,'2016-10-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(266,6,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(267,9,1,0,0,0,'2016-11-01 00:00:00',89200.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(268,12,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(269,1,1,0,0,0,'2016-11-01 00:00:00',15300.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(270,11,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(271,10,1,0,0,0,'2016-11-01 00:00:00',180500.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(272,2,1,0,0,0,'2016-11-01 00:00:00',50600.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(273,13,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(274,7,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(275,8,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(276,28,1,0,0,0,'2016-11-01 00:00:00',-537000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(277,30,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(278,14,1,0,0,0,'2016-11-01 00:00:00',-180500.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(279,3,1,0,0,0,'2016-11-01 00:00:00',-50600.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(280,15,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(281,31,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(282,17,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(283,18,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(284,19,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(285,20,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(286,21,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(287,16,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(288,5,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(289,22,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(290,23,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(291,24,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(292,4,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(293,25,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(294,26,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(295,27,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(296,32,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(297,33,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(298,34,1,0,0,0,'2016-11-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(299,6,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(300,9,1,0,0,0,'2016-12-01 00:00:00',104500.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(301,12,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(302,1,1,0,0,0,'2016-12-01 00:00:00',15300.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(303,11,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(304,10,1,0,0,0,'2016-12-01 00:00:00',243000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(305,2,1,0,0,0,'2016-12-01 00:00:00',62500.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(306,13,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(307,7,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(308,8,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(309,28,1,0,0,0,'2016-12-01 00:00:00',-780000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(310,30,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(311,14,1,0,0,0,'2016-12-01 00:00:00',-243000.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(312,3,1,0,0,0,'2016-12-01 00:00:00',-62500.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(313,15,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(314,31,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(315,17,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(316,18,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(317,19,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(318,20,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(319,21,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(320,16,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(321,5,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(322,22,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(323,23,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(324,24,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(325,4,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(326,25,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(327,26,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(328,27,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(329,32,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(330,33,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(331,34,1,0,0,0,'2016-12-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(332,6,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(333,9,1,0,0,0,'2017-01-01 00:00:00',119800.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(334,12,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(335,1,1,0,0,0,'2017-01-01 00:00:00',15300.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(336,11,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(337,10,1,0,0,0,'2017-01-01 00:00:00',317400.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(338,2,1,0,0,0,'2017-01-01 00:00:00',74400.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(339,13,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:58',0,'2017-06-21 19:40:58',0),(340,7,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(341,8,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(342,28,1,0,0,0,'2017-01-01 00:00:00',-1097400.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(343,30,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(344,14,1,0,0,0,'2017-01-01 00:00:00',-317400.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(345,3,1,0,0,0,'2017-01-01 00:00:00',-74400.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(346,15,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(347,31,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(348,17,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(349,18,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(350,19,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(351,20,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(352,21,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(353,16,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(354,5,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(355,22,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(356,23,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(357,24,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(358,4,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(359,25,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(360,26,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(361,27,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(362,32,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(363,33,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(364,34,1,0,0,0,'2017-01-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(365,6,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(366,9,1,0,0,0,'2017-02-01 00:00:00',135100.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(367,12,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(368,1,1,0,0,0,'2017-02-01 00:00:00',15300.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(369,11,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(370,10,1,0,0,0,'2017-02-01 00:00:00',403700.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(371,2,1,0,0,0,'2017-02-01 00:00:00',86300.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(372,13,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(373,7,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(374,8,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(375,28,1,0,0,0,'2017-02-01 00:00:00',-1501100.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(376,30,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(377,14,1,0,0,0,'2017-02-01 00:00:00',-403700.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(378,3,1,0,0,0,'2017-02-01 00:00:00',-86300.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(379,15,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(380,31,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(381,17,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(382,18,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(383,19,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(384,20,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(385,21,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(386,16,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(387,5,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(388,22,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(389,23,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(390,24,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(391,4,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(392,25,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(393,26,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(394,27,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(395,32,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(396,33,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(397,34,1,0,0,0,'2017-02-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(398,6,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(399,9,1,0,0,0,'2017-03-01 00:00:00',150400.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(400,12,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(401,1,1,0,0,0,'2017-03-01 00:00:00',15300.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(402,11,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(403,10,1,0,0,0,'2017-03-01 00:00:00',501900.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(404,2,1,0,0,0,'2017-03-01 00:00:00',98200.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(405,13,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(406,7,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(407,8,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(408,28,1,0,0,0,'2017-03-01 00:00:00',-2003000.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(409,30,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(410,14,1,0,0,0,'2017-03-01 00:00:00',-501900.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(411,3,1,0,0,0,'2017-03-01 00:00:00',-98200.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(412,15,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(413,31,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(414,17,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(415,18,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(416,19,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(417,20,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(418,21,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(419,16,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(420,5,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(421,22,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(422,23,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(423,24,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(424,4,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(425,25,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(426,26,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(427,27,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(428,32,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(429,33,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(430,34,1,0,0,0,'2017-03-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(431,6,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(432,9,1,0,0,0,'2017-04-01 00:00:00',165700.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(433,12,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(434,1,1,0,0,0,'2017-04-01 00:00:00',15300.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(435,11,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(436,10,1,0,0,0,'2017-04-01 00:00:00',612000.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(437,2,1,0,0,0,'2017-04-01 00:00:00',110100.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(438,13,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(439,7,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(440,8,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(441,28,1,0,0,0,'2017-04-01 00:00:00',-2615000.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(442,30,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(443,14,1,0,0,0,'2017-04-01 00:00:00',-612000.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(444,3,1,0,0,0,'2017-04-01 00:00:00',-110100.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(445,15,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(446,31,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(447,17,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(448,18,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(449,19,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(450,20,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(451,21,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(452,16,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(453,5,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(454,22,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(455,23,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(456,24,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(457,4,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(458,25,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(459,26,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(460,27,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(461,32,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(462,33,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(463,34,1,0,0,0,'2017-04-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(464,6,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(465,9,1,0,0,0,'2017-05-01 00:00:00',181000.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(466,12,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(467,1,1,0,0,0,'2017-05-01 00:00:00',15300.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(468,11,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(469,10,1,0,0,0,'2017-05-01 00:00:00',734000.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(470,2,1,0,0,0,'2017-05-01 00:00:00',122000.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(471,13,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(472,7,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(473,8,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(474,28,1,0,0,0,'2017-05-01 00:00:00',-3349000.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(475,30,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(476,14,1,0,0,0,'2017-05-01 00:00:00',-734000.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(477,3,1,0,0,0,'2017-05-01 00:00:00',-122000.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(478,15,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(479,31,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(480,17,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(481,18,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(482,19,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(483,20,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(484,21,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(485,16,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(486,5,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(487,22,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(488,23,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(489,24,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(490,4,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(491,25,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(492,26,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(493,27,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(494,32,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(495,33,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(496,34,1,0,0,0,'2017-05-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(497,6,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(498,9,1,0,0,0,'2017-06-01 00:00:00',196300.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(499,12,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(500,1,1,0,0,0,'2017-06-01 00:00:00',15300.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(501,11,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(502,10,1,0,0,0,'2017-06-01 00:00:00',867900.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(503,2,1,0,0,0,'2017-06-01 00:00:00',133900.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(504,13,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(505,7,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(506,8,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(507,28,1,0,0,0,'2017-06-01 00:00:00',-4216900.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(508,30,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(509,14,1,0,0,0,'2017-06-01 00:00:00',-867900.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(510,3,1,0,0,0,'2017-06-01 00:00:00',-133900.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(511,15,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(512,31,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(513,17,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(514,18,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(515,19,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(516,20,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(517,21,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(518,16,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(519,5,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(520,22,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(521,23,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(522,24,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(523,4,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(524,25,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(525,26,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(526,27,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(527,32,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(528,33,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(529,34,1,0,0,0,'2017-06-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(530,6,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(531,9,1,0,0,0,'2017-07-01 00:00:00',211600.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(532,12,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(533,1,1,0,0,0,'2017-07-01 00:00:00',15300.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(534,11,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(535,10,1,0,0,0,'2017-07-01 00:00:00',1013700.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(536,2,1,0,0,0,'2017-07-01 00:00:00',145800.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(537,13,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(538,7,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(539,8,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(540,28,1,0,0,0,'2017-07-01 00:00:00',-5230600.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(541,30,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(542,14,1,0,0,0,'2017-07-01 00:00:00',-1013700.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(543,3,1,0,0,0,'2017-07-01 00:00:00',-145800.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(544,15,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(545,31,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(546,17,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(547,18,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(548,19,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(549,20,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(550,21,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(551,16,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(552,5,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(553,22,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(554,23,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(555,24,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(556,4,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(557,25,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(558,26,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(559,27,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(560,32,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(561,33,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(562,34,1,0,0,0,'2017-07-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(563,6,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(564,9,1,0,0,0,'2017-08-01 00:00:00',226900.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(565,12,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(566,1,1,0,0,0,'2017-08-01 00:00:00',15300.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(567,11,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(568,10,1,0,0,0,'2017-08-01 00:00:00',1171400.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(569,2,1,0,0,0,'2017-08-01 00:00:00',157700.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(570,13,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(571,7,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(572,8,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(573,28,1,0,0,0,'2017-08-01 00:00:00',-6402000.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(574,30,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(575,14,1,0,0,0,'2017-08-01 00:00:00',-1171400.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(576,3,1,0,0,0,'2017-08-01 00:00:00',-157700.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(577,15,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(578,31,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(579,17,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(580,18,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(581,19,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(582,20,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(583,21,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(584,16,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(585,5,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(586,22,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(587,23,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(588,24,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(589,4,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(590,25,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(591,26,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(592,27,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(593,32,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(594,33,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0),(595,34,1,0,0,0,'2017-08-01 00:00:00',0.0000,0,'2017-06-21 19:40:59',0,'2017-06-21 19:40:59',0);
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`PMTID`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PaymentType`
--

LOCK TABLES `PaymentType` WRITE;
/*!40000 ALTER TABLE `PaymentType` DISABLE KEYS */;
INSERT INTO `PaymentType` VALUES (1,1,'Check','','2017-06-13 16:47:54',0,'2017-06-14 18:26:50',0),(2,1,'Credit Card','','2017-06-13 16:48:01',0,'2017-06-14 18:26:50',0),(3,1,'Money Order','','2017-06-13 16:48:14',0,'2017-06-14 18:26:50',0),(4,1,'EFT','','2017-06-13 16:51:56',0,'2017-06-14 18:26:50',0),(5,1,'ACH','','2017-06-13 16:52:03',0,'2017-06-14 18:26:50',0),(6,1,'WIRE','','2017-06-13 16:52:11',0,'2017-06-14 18:26:50',0),(7,1,'Branch Deposit','','2017-06-13 19:21:23',0,'2017-06-14 18:26:50',0);
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
INSERT INTO `Payor` VALUES (1,1,'',0.0000,0,1,'2017-06-13 19:39:18',0,'2017-06-14 18:26:50',0),(2,1,'',0.0000,0,1,'2017-06-13 19:40:59',0,'2017-06-14 18:26:50',0),(3,1,'',0.0000,0,1,'2017-06-15 16:35:44',0,'2017-06-15 16:35:44',0),(4,1,'',0.0000,0,1,'2017-06-15 16:36:27',0,'2017-06-15 16:36:27',0),(5,1,'',0.0000,0,1,'2017-06-15 16:38:32',0,'2017-06-15 16:38:32',0),(6,1,'',0.0000,0,1,'2017-06-15 16:50:13',0,'2017-06-15 16:50:13',0),(7,1,'',0.0000,0,1,'2017-06-22 18:59:11',0,'2017-06-22 18:59:11',0);
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
INSERT INTO `Prospect` VALUES (1,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-06-13 19:39:18',0,'2017-06-14 18:26:50',0),(2,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-06-13 19:40:59',0,'2017-06-14 18:26:50',0),(3,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-06-15 16:35:44',0,'2017-06-15 16:35:44',0),(4,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-06-15 16:36:27',0,'2017-06-15 16:36:27',0),(5,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-06-15 16:38:32',0,'2017-06-15 16:38:32',0),(6,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-06-15 16:50:13',0,'2017-06-15 16:50:13',0),(7,1,'','','','','','','','',0.0000,'1900-01-01',0,0,0,0,'','1900-01-01',0,0,0.0000,0,'2017-06-22 18:59:11',0,'2017-06-22 18:59:11',0);
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
  `AcctRuleApply` varchar(2048) NOT NULL DEFAULT '',
  `FLAGS` bigint(20) NOT NULL DEFAULT '0',
  `Comment` varchar(256) NOT NULL DEFAULT '',
  `OtherPayorName` varchar(128) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RCPTID`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Receipt`
--

LOCK TABLES `Receipt` WRITE;
/*!40000 ALTER TABLE `Receipt` DISABLE KEYS */;
INSERT INTO `Receipt` VALUES (1,0,1,1,1,0,0,'2014-03-01 00:00:00','9999',7000.0000,'',2,'',0,'','Kelemen Real Estate Agency','2017-06-20 18:17:14',0,'2017-06-20 18:17:14',0),(2,0,1,3,7,0,0,'2017-06-01 00:00:00','',4150.0000,'',8,'',0,'','','2017-06-22 17:20:34',0,'2017-06-22 17:20:34',0),(3,0,1,1,6,0,0,'2016-06-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:28:23',0,'2017-06-22 17:21:37',0),(4,0,1,1,6,0,0,'2016-03-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:26:32',0,'2017-06-22 17:26:32',0),(5,0,1,1,6,0,0,'2016-04-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:27:11',0,'2017-06-22 17:27:11',0),(6,0,1,1,6,0,0,'2016-05-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:27:52',0,'2017-06-22 17:27:52',0),(8,0,1,1,6,0,0,'2016-07-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:30:18',0,'2017-06-22 17:30:18',0),(9,0,1,1,6,0,0,'2017-08-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:30:44',0,'2017-06-22 17:30:44',0),(10,0,1,1,6,0,0,'2016-08-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:31:20',0,'2017-06-22 17:31:20',0),(11,0,1,1,6,0,0,'2016-09-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:32:16',0,'2017-06-22 17:32:16',0),(12,0,1,1,6,0,0,'2016-10-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:32:37',0,'2017-06-22 17:32:37',0),(13,0,1,1,6,0,0,'2017-06-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:33:06',0,'2017-06-22 17:33:06',0),(14,0,1,1,6,0,0,'2016-11-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:34:00',0,'2017-06-22 17:34:00',0),(15,0,1,1,6,0,0,'2016-12-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:34:29',0,'2017-06-22 17:34:29',0),(16,0,1,1,6,0,0,'2017-01-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:35:16',0,'2017-06-22 17:35:16',0),(17,0,1,1,6,0,0,'2017-02-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:35:40',0,'2017-06-22 17:35:40',0),(18,0,1,1,6,0,0,'2017-03-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:36:00',0,'2017-06-22 17:36:00',0),(19,0,1,1,6,0,0,'2017-04-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:36:24',0,'2017-06-22 17:36:24',0),(20,0,1,1,6,0,0,'2017-05-01 00:00:00','',3750.0000,'',3,'',0,'','','2017-06-22 17:36:56',0,'2017-06-22 17:36:56',0),(21,0,1,3,7,0,0,'2016-08-01 00:00:00','',4150.0000,'',8,'',0,'','','2017-06-22 17:48:45',0,'2017-06-22 17:48:45',0),(22,0,1,3,7,0,0,'2016-09-01 00:00:00','',4150.0000,'',8,'',0,'','','2017-06-22 17:49:15',0,'2017-06-22 17:49:15',0),(23,0,1,3,7,0,0,'2016-10-01 00:00:00','',4150.0000,'',8,'',0,'','','2017-06-22 17:50:25',0,'2017-06-22 17:50:25',0),(24,0,1,3,7,0,0,'2016-11-01 00:00:00','',4150.0000,'',8,'',0,'','','2017-06-22 17:51:26',0,'2017-06-22 17:51:26',0),(25,0,1,3,7,0,0,'2016-12-01 00:00:00','',4150.0000,'',8,'',0,'','','2017-06-22 17:51:52',0,'2017-06-22 17:51:52',0),(26,0,1,3,7,0,0,'2017-01-01 00:00:00','',4150.0000,'',8,'',0,'','','2017-06-22 17:52:26',0,'2017-06-22 17:52:26',0),(27,0,1,3,7,0,0,'2017-02-01 00:00:00','',4150.0000,'',8,'',0,'','','2017-06-22 17:52:52',0,'2017-06-22 17:52:52',0),(28,0,1,3,7,0,0,'2017-03-01 00:00:00','',4150.0000,'',8,'',0,'','','2017-06-22 17:53:32',0,'2017-06-22 17:53:32',0),(29,0,1,3,7,0,0,'2017-04-01 00:00:00','',4150.0000,'',8,'',0,'','','2017-06-22 17:54:12',0,'2017-06-22 17:54:12',0),(30,0,1,3,7,0,0,'2017-05-01 00:00:00','',4150.0000,'',8,'',0,'','','2017-06-22 17:54:46',0,'2017-06-22 17:54:46',0),(31,0,1,6,6,0,0,'2016-10-03 00:00:00','',4000.0000,'',3,'',0,'','Beaumont Partners LTD','2017-06-22 18:02:32',0,'2017-06-22 18:00:00',0),(32,0,1,6,6,0,0,'2016-11-10 00:00:00','',11985.0000,'',3,'',0,'Prepaid 3 monthls (NOV-JAN), less wire fee','','2017-06-22 18:10:35',0,'2017-06-22 18:01:58',0),(33,0,1,6,6,0,0,'2017-02-13 00:00:00','',8335.0000,'',3,'',0,'FEB/MAR Rent and Utilities less wire fee','','2017-06-22 18:22:37',0,'2017-06-22 18:22:37',0),(34,0,1,6,6,0,0,'2017-05-12 00:00:00','',13116.7900,'',4,'',0,'APR/MAY/JUN rents and utilities','','2017-06-22 18:28:38',0,'2017-06-22 18:28:38',0),(35,0,1,6,6,0,0,'2017-02-02 00:00:00','',628.4500,'',3,'',0,'SEP-DEC utilities','','2017-06-22 18:52:14',0,'2017-06-22 18:52:14',0),(36,0,1,7,1,0,0,'2016-07-01 00:00:00','',6474.0000,'',2,'',0,'net received from Kelemen','','2017-06-22 19:05:23',0,'2017-06-22 19:05:23',0);
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
  `AcctRule` varchar(150) DEFAULT NULL,
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RCPAID`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ReceiptAllocation`
--

LOCK TABLES `ReceiptAllocation` WRITE;
/*!40000 ALTER TABLE `ReceiptAllocation` DISABLE KEYS */;
INSERT INTO `ReceiptAllocation` VALUES (1,1,1,0,'2014-03-01 00:00:00',7000.0000,0,'d 10104 _, c 59998 _','2017-06-20 18:17:14',0),(2,2,1,0,'2017-06-01 00:00:00',4150.0000,0,'d 10104 _, c 59998 _','2017-06-22 17:20:34',0),(3,3,1,0,'2017-06-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:21:37',0),(4,4,1,0,'2016-03-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:26:32',0),(5,5,1,0,'2016-04-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:27:11',0),(6,6,1,0,'2016-05-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:27:52',0),(8,8,1,0,'2016-07-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:30:18',0),(9,9,1,0,'2017-08-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:30:44',0),(10,10,1,0,'2016-08-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:31:20',0),(11,11,1,0,'2016-09-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:32:16',0),(12,12,1,0,'2016-10-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:32:37',0),(13,13,1,0,'2017-06-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:33:06',0),(14,14,1,0,'2016-11-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:34:00',0),(15,15,1,0,'2016-12-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:34:29',0),(16,16,1,0,'2017-01-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:35:16',0),(17,17,1,0,'2017-02-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:35:40',0),(18,18,1,0,'2017-03-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:36:00',0),(19,19,1,0,'2017-04-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:36:24',0),(20,20,1,0,'2017-05-01 00:00:00',3750.0000,0,'d 10104 _, c  _','2017-06-22 17:36:56',0),(21,21,1,0,'2016-08-01 00:00:00',4150.0000,0,'d 10104 _, c 59998 _','2017-06-22 17:48:45',0),(22,22,1,0,'2016-09-01 00:00:00',4150.0000,0,'d 10104 _, c 59998 _','2017-06-22 17:49:15',0),(23,23,1,0,'2016-10-01 00:00:00',4150.0000,0,'d 10104 _, c 59998 _','2017-06-22 17:50:25',0),(24,24,1,0,'2016-11-01 00:00:00',4150.0000,0,'d 10104 _, c 59998 _','2017-06-22 17:51:26',0),(25,25,1,0,'2016-12-01 00:00:00',4150.0000,0,'d 10104 _, c 59998 _','2017-06-22 17:51:52',0),(26,26,1,0,'2017-01-01 00:00:00',4150.0000,0,'d 10104 _, c 59998 _','2017-06-22 17:52:26',0),(27,27,1,0,'2017-02-01 00:00:00',4150.0000,0,'d 10104 _, c 59998 _','2017-06-22 17:52:52',0),(28,28,1,0,'2017-03-01 00:00:00',4150.0000,0,'d 10104 _, c 59998 _','2017-06-22 17:53:32',0),(29,29,1,0,'2017-04-01 00:00:00',4150.0000,0,'d 10104 _, c 59998 _','2017-06-22 17:54:12',0),(30,30,1,0,'2017-05-01 00:00:00',4150.0000,0,'d 10104 _, c 59998 _','2017-06-22 17:54:46',0),(31,31,1,0,'2016-10-03 00:00:00',4000.0000,0,'d 10104 _, c  _','2017-06-22 18:00:00',0),(32,32,1,0,'2016-11-10 00:00:00',12000.0000,0,'d 10104 _, c  _','2017-06-22 18:01:58',0),(33,33,1,0,'2017-02-13 00:00:00',8335.0000,0,'d 10104 _, c  _','2017-06-22 18:22:37',0),(34,34,1,0,'2017-05-12 00:00:00',13116.7900,0,'d 10104 _, c 59998 _','2017-06-22 18:28:38',0),(35,35,1,0,'2017-02-02 00:00:00',628.4500,0,'d 10104 _, c  _','2017-06-22 18:52:14',0),(36,36,1,0,'2016-07-01 00:00:00',6474.0000,0,'d 10104 _, c 59998 _','2017-06-22 19:05:23',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Rentable`
--

LOCK TABLES `Rentable` WRITE;
/*!40000 ALTER TABLE `Rentable` DISABLE KEYS */;
INSERT INTO `Rentable` VALUES (1,1,'309 S Rexford',1,'2017-06-13 19:34:58',0,'2017-06-14 18:26:51',0),(2,1,'309 1/2 S Rexford',1,'2017-06-13 20:02:10',0,'2017-06-14 18:26:51',0),(3,1,'311 S Rexford',1,'2017-06-13 20:02:33',0,'2017-06-14 18:26:51',0),(4,1,'311 1/2 S Rexford',1,'2017-06-13 20:03:01',0,'2017-06-14 18:26:51',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableMarketRate`
--

LOCK TABLES `RentableMarketRate` WRITE;
/*!40000 ALTER TABLE `RentableMarketRate` DISABLE KEYS */;
INSERT INTO `RentableMarketRate` VALUES (1,1,1,3500.0000,'2014-01-01 00:00:00','3000-01-01 00:00:00','2017-06-14 18:26:52',0),(2,2,1,3550.0000,'2014-01-01 00:00:00','3000-01-01 00:00:00','2017-06-14 18:26:52',0),(3,3,1,4400.0000,'2014-01-01 00:00:00','3000-01-01 00:00:00','2017-06-14 18:26:52',0),(4,4,1,2500.0000,'2014-01-01 00:00:00','2017-06-15 05:09:09','2017-06-14 18:26:52',0),(51,4,1,0.0000,'2017-06-15 05:09:09','2017-06-15 05:17:56','2017-06-15 05:09:09',0),(52,4,1,2400.0000,'2017-06-15 05:17:56','9998-12-31 00:00:00','2017-06-15 05:17:55',0);
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
INSERT INTO `RentableStatus` VALUES (5,1,1,1,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-06-20 16:09:45',0,'2017-06-20 16:09:45',0),(6,2,1,1,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-06-20 16:20:18',0,'2017-06-20 16:20:18',0),(7,3,1,1,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-06-20 16:20:30',0,'2017-06-20 16:20:30',0),(8,4,1,4,'2014-01-01 00:00:00','9999-01-01 00:00:00','0000-00-00','2017-06-20 16:20:41',0,'2017-06-20 16:20:41',0);
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
INSERT INTO `RentableTypes` VALUES (1,1,'Rex1','309 Rexford',6,4,4,1,'2017-06-13 05:39:46',0,'2017-06-14 18:26:53',0),(2,1,'Rex2','309 1/2 Rexford',6,4,4,1,'2017-06-13 05:39:46',0,'2017-06-14 18:26:53',0),(3,1,'Rex3','311 Rexford',6,4,4,1,'2017-06-13 05:39:46',0,'2017-06-14 18:26:53',0),(4,1,'Rex4','311 1/2 Rexford',6,4,4,1,'2017-06-15 05:43:52',0,'2017-06-14 18:26:53',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableUsers`
--

LOCK TABLES `RentableUsers` WRITE;
/*!40000 ALTER TABLE `RentableUsers` DISABLE KEYS */;
INSERT INTO `RentableUsers` VALUES (1,1,1,2,'2014-03-01','2018-03-01','2017-06-14 18:26:53',0),(2,1,1,1,'2014-03-01','2018-03-01','2017-06-14 18:26:53',0),(3,3,1,3,'2016-07-01','2018-07-01','2017-06-15 16:44:09',0),(4,3,1,4,'2016-07-01','2018-07-01','2017-06-15 16:44:48',0),(5,2,1,5,'2016-10-01','2017-12-31','2017-06-15 16:52:18',0);
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
  `NextRateChange` date NOT NULL DEFAULT '1970-01-01',
  `PermittedUses` varchar(128) NOT NULL DEFAULT '',
  `ExclusiveUses` varchar(128) NOT NULL DEFAULT '',
  `ExtensionOption` varchar(128) NOT NULL DEFAULT '',
  `ExtensionOptionNotice` date NOT NULL DEFAULT '1970-01-01',
  `ExpansionOption` varchar(128) NOT NULL DEFAULT '',
  `ExpansionOptionNotice` date NOT NULL DEFAULT '1970-01-01',
  `RightOfFirstRefusal` varchar(128) NOT NULL DEFAULT '',
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
INSERT INTO `RentalAgreement` VALUES (1,0,1,0,'2014-03-01','2018-03-01','2014-03-01','2018-03-01','2014-03-01','2018-03-01','2017-07-01',0,1,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,'1900-01-01','','','','1900-01-01','','1900-01-01','','2017-06-15 16:41:28',0,'2017-06-14 18:26:53',0),(2,0,1,0,'2016-07-01','2018-07-01','2016-07-01','2018-07-01','2016-07-01','2018-07-01','2017-07-01',0,1,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,'1900-01-01','','','','1900-01-01','','1900-01-01','','2017-06-15 16:46:04',0,'2017-06-15 16:41:49',0),(3,0,1,0,'2016-10-01','2017-12-31','2016-10-01','2017-12-31','2016-10-01','2017-12-31','2017-07-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,'1900-01-01','','','','1900-01-01','','1900-01-01','','2017-06-15 16:48:19',0,'2017-06-15 16:46:12',0),(4,0,1,0,'2017-06-15','2017-06-15','2017-06-15','2018-06-15','2017-06-15','2018-06-15','2017-06-01',0,0,1,'',0,0,0.0000,'','1900-01-01','1900-01-01',0.0000,0.0000,'1900-01-01','','','','1900-01-01','','1900-01-01','','2017-06-20 18:28:47',0,'2017-06-15 16:50:49',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementPayors`
--

LOCK TABLES `RentalAgreementPayors` WRITE;
/*!40000 ALTER TABLE `RentalAgreementPayors` DISABLE KEYS */;
INSERT INTO `RentalAgreementPayors` VALUES (1,1,1,1,'2014-03-01','2018-03-01',0,'2017-06-14 18:26:53',0),(2,2,1,3,'2016-07-01','2018-07-01',0,'2017-06-15 16:43:15',0),(3,3,1,6,'2016-10-01','2017-12-31',0,'2017-06-15 16:51:41',0);
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
INSERT INTO `RentalAgreementRentables` VALUES (1,1,1,1,0,3750.0000,'2016-03-01','2018-03-01','2017-06-14 18:26:54',0),(4,2,1,3,0,4150.0000,'2016-07-01','2018-07-01','2017-06-20 17:59:59',0),(5,3,1,2,0,4000.0000,'2016-10-01','2018-01-01','2017-06-20 18:00:33',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Transactant`
--

LOCK TABLES `Transactant` WRITE;
/*!40000 ALTER TABLE `Transactant` DISABLE KEYS */;
INSERT INTO `Transactant` VALUES (1,1,0,'Aaron','','Read','','',0,'read.aaron@gmail.com','','','1-469-307-7095','','','','','','','','2017-06-15 16:33:59',0,'2017-06-14 18:26:55',0),(2,1,0,'Kirsten','','Read','','',0,'klmrda@gmail.com','','','1-469-693-9933','','','','','','','','2017-06-15 16:34:39',0,'2017-06-14 18:26:55',0),(3,1,0,'Kevin','','Mills','','',0,'kevinmillsesq@aol.com','','','1-424-234-3535','','','','','','','','2017-06-15 16:35:44',0,'2017-06-15 16:35:44',0),(4,1,0,'Lauren','','Beck','','',0,'laurensbeck@aol.com','','','1-310-948-6442','','','','','','','','2017-06-15 16:36:27',0,'2017-06-15 16:36:27',0),(5,1,0,'Alex','','Vahabzadeh','','Beaumont Partners LLC',0,'av@beaumont-partners.ch','','44-79-203-354-77','1-202-550-2477','118 Rue du Rhone','','1204 Geneva','','','Switzerland','','2017-06-15 16:50:29',0,'2017-06-15 16:38:32',0),(6,1,0,'','','','','Beaumont Partners LLC',1,'av@beaumont-partners.ch','scigler@bvgroup.com','44-79-203-354-77','1-202-550-2477','118 Rue du Rhone','','1204 Geneva','','','Switzerland','','2017-06-15 16:50:13',0,'2017-06-15 16:50:13',0),(7,1,0,'','','','','Kelemen Real Estate Agency',1,'','','','','','','','','','','','2017-06-22 18:59:11',0,'2017-06-22 18:59:11',0);
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
INSERT INTO `User` VALUES (1,1,0,'1900-01-01','','','','','',1,'',0,'2017-06-13 19:39:18',0,'2017-06-14 18:26:55',0),(2,1,0,'1900-01-01','','','','','',1,'',0,'2017-06-13 19:40:59',0,'2017-06-14 18:26:55',0),(3,1,0,'1900-01-01','','','','','',1,'',0,'2017-06-15 16:35:44',0,'2017-06-15 16:35:44',0),(4,1,0,'1900-01-01','','','','','',1,'',0,'2017-06-15 16:36:27',0,'2017-06-15 16:36:27',0),(5,1,0,'1900-01-01','','','','','',1,'',0,'2017-06-15 16:38:32',0,'2017-06-15 16:38:32',0),(6,1,0,'1900-01-01','','','','','',1,'',0,'2017-06-15 16:50:13',0,'2017-06-15 16:50:13',0),(7,1,0,'1900-01-01','','','','','',1,'',0,'2017-06-22 18:59:11',0,'2017-06-22 18:59:11',0);
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

-- Dump completed on 2017-06-22 19:05:37
