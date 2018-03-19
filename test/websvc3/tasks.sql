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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `AR`
--

LOCK TABLES `AR` WRITE;
/*!40000 ALTER TABLE `AR` DISABLE KEYS */;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Assessments`
--

LOCK TABLES `Assessments` WRITE;
/*!40000 ALTER TABLE `Assessments` DISABLE KEYS */;
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
INSERT INTO `Business` VALUES (1,'REX','JGM First, LLC',6,4,4,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `CustomAttr`
--

LOCK TABLES `CustomAttr` WRITE;
/*!40000 ALTER TABLE `CustomAttr` DISABLE KEYS */;
INSERT INTO `CustomAttr` VALUES (1,1,1,'Square Feet','5000','sqft','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(2,1,1,'Square Feet','692','sqft','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(3,1,1,'Square Feet','952','sqft','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(4,1,1,'Square Feet','2000','sqft','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0);
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Depository`
--

LOCK TABLES `Depository` WRITE;
/*!40000 ALTER TABLE `Depository` DISABLE KEYS */;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `GLAccount`
--

LOCK TABLES `GLAccount` WRITE;
/*!40000 ALTER TABLE `GLAccount` DISABLE KEYS */;
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
  `InvoicePayorID` bigint(20) NOT NULL AUTO_INCREMENT,
  `InvoiceNo` bigint(20) NOT NULL DEFAULT '0',
  `BID` bigint(20) NOT NULL DEFAULT '0',
  `PID` bigint(20) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Journal`
--

LOCK TABLES `Journal` WRITE;
/*!40000 ALTER TABLE `Journal` DISABLE KEYS */;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `JournalAllocation`
--

LOCK TABLES `JournalAllocation` WRITE;
/*!40000 ALTER TABLE `JournalAllocation` DISABLE KEYS */;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerEntry`
--

LOCK TABLES `LedgerEntry` WRITE;
/*!40000 ALTER TABLE `LedgerEntry` DISABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LedgerMarker`
--

LOCK TABLES `LedgerMarker` WRITE;
/*!40000 ALTER TABLE `LedgerMarker` DISABLE KEYS */;
INSERT INTO `LedgerMarker` VALUES (1,0,1,0,0,1,'1970-01-01 00:00:00',0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(2,0,1,0,0,2,'1970-01-01 00:00:00',0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(3,0,1,0,0,3,'1970-01-01 00:00:00',0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(4,0,1,0,0,4,'1970-01-01 00:00:00',0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(5,0,1,0,0,5,'1970-01-01 00:00:00',0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(6,0,1,0,0,6,'1970-01-01 00:00:00',0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(7,0,1,0,0,7,'1970-01-01 00:00:00',0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(8,0,1,0,0,8,'1970-01-01 00:00:00',0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(9,0,1,0,0,9,'1970-01-01 00:00:00',0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(10,0,1,0,0,10,'1970-01-01 00:00:00',0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(11,0,1,0,0,11,'1970-01-01 00:00:00',0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(12,0,1,0,0,12,'1970-01-01 00:00:00',0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(13,0,1,0,0,13,'1970-01-01 00:00:00',0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(14,0,1,0,0,14,'1970-01-01 00:00:00',0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(15,0,1,0,0,15,'1970-01-01 00:00:00',0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `NoteList`
--

LOCK TABLES `NoteList` WRITE;
/*!40000 ALTER TABLE `NoteList` DISABLE KEYS */;
INSERT INTO `NoteList` VALUES (1,1,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(2,1,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Notes`
--

LOCK TABLES `Notes` WRITE;
/*!40000 ALTER TABLE `Notes` DISABLE KEYS */;
INSERT INTO `Notes` VALUES (1,1,1,0,1,0,0,0,'Note for Aaron','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(2,1,2,0,1,0,0,0,'Note for Kirsten','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0);
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PaymentType`
--

LOCK TABLES `PaymentType` WRITE;
/*!40000 ALTER TABLE `PaymentType` DISABLE KEYS */;
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
INSERT INTO `Payor` VALUES (1,1,'',0.0000,1,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(2,1,'',0.0000,2,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(3,1,'',0.0000,3,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(4,1,'',0.0000,4,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(5,1,'',0.0000,5,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(6,1,'',0.0000,6,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(7,1,'',0.0000,7,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(8,1,'',0.0000,8,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(9,1,'',0.0000,9,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(10,1,'',0.0000,10,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(11,1,'',0.0000,0,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(12,1,'',0.0000,0,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(13,1,'',0.0000,0,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(14,1,'',0.0000,0,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(15,1,'',0.0000,0,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0);
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
INSERT INTO `Prospect` VALUES (1,1,'','','','','','','','',1000.7500,'0000-00-00',0,0,1,0,'','0000-00-00',0,0,0.0000,1,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(2,1,'','','','','','','','',0.0000,'0000-00-00',0,0,2,0,'','0000-00-00',0,0,0.0000,2,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(3,1,'','','','','','','','',1100.7500,'0000-00-00',0,0,3,0,'','0000-00-00',0,0,0.0000,3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(4,1,'','','','','','','','',0.0000,'0000-00-00',0,0,4,0,'','0000-00-00',0,0,0.0000,4,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(5,1,'','','','','','','','',0.0000,'0000-00-00',0,0,5,0,'','0000-00-00',0,0,0.0000,5,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(6,1,'','','','','','','','',1200.7500,'0000-00-00',0,0,6,0,'','0000-00-00',0,0,0.0000,6,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(7,1,'','','','','','','','',1300.7500,'0000-00-00',0,0,7,0,'','0000-00-00',0,0,0.0000,7,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(8,1,'','','','','','','','',0.0000,'0000-00-00',0,0,8,0,'','0000-00-00',0,0,0.0000,8,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(9,1,'','','','','','','','',0.0000,'0000-00-00',0,0,9,0,'','0000-00-00',0,0,0.0000,9,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(10,1,'','','','','','','','',0.0000,'0000-00-00',0,0,10,0,'','0000-00-00',0,0,0.0000,10,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(11,1,'','','','','','','','',0.0000,'0000-00-00',0,0,0,0,'','0000-00-00',0,0,0.0000,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(12,1,'','','','','','','','',0.0000,'0000-00-00',0,0,0,0,'','0000-00-00',0,0,0.0000,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(13,1,'','','','','','','','',0.0000,'0000-00-00',0,0,0,0,'','0000-00-00',0,0,0.0000,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(14,1,'','','','','','','','',0.0000,'0000-00-00',0,0,0,0,'','0000-00-00',0,0,0.0000,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(15,1,'','','','','','','','',0.0000,'0000-00-00',0,0,0,0,'','0000-00-00',0,0,0.0000,0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Receipt`
--

LOCK TABLES `Receipt` WRITE;
/*!40000 ALTER TABLE `Receipt` DISABLE KEYS */;
INSERT INTO `Receipt` VALUES (1,0,1,0,1,0,0,0,'2017-02-14 00:00:00','12345',4217000.0000,'',0,'',0,'','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0);
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ReceiptAllocation`
--

LOCK TABLES `ReceiptAllocation` WRITE;
/*!40000 ALTER TABLE `ReceiptAllocation` DISABLE KEYS */;
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
  PRIMARY KEY (`RID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Rentable`
--

LOCK TABLES `Rentable` WRITE;
/*!40000 ALTER TABLE `Rentable` DISABLE KEYS */;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableMarketRate`
--

LOCK TABLES `RentableMarketRate` WRITE;
/*!40000 ALTER TABLE `RentableMarketRate` DISABLE KEYS */;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableStatus`
--

LOCK TABLES `RentableStatus` WRITE;
/*!40000 ALTER TABLE `RentableStatus` DISABLE KEYS */;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableTypeRef`
--

LOCK TABLES `RentableTypeRef` WRITE;
/*!40000 ALTER TABLE `RentableTypeRef` DISABLE KEYS */;
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`RTID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableTypes`
--

LOCK TABLES `RentableTypes` WRITE;
/*!40000 ALTER TABLE `RentableTypes` DISABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentableUsers`
--

LOCK TABLES `RentableUsers` WRITE;
/*!40000 ALTER TABLE `RentableUsers` DISABLE KEYS */;
INSERT INTO `RentableUsers` VALUES (1,1,1,14,'2018-03-14','2020-02-14','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0);
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreement`
--

LOCK TABLES `RentalAgreement` WRITE;
/*!40000 ALTER TABLE `RentalAgreement` DISABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementPayors`
--

LOCK TABLES `RentalAgreementPayors` WRITE;
/*!40000 ALTER TABLE `RentalAgreementPayors` DISABLE KEYS */;
INSERT INTO `RentalAgreementPayors` VALUES (1,1,1,14,'2016-10-24','2017-02-14',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RentalAgreementRentables`
--

LOCK TABLES `RentalAgreementRentables` WRITE;
/*!40000 ALTER TABLE `RentalAgreementRentables` DISABLE KEYS */;
INSERT INTO `RentalAgreementRentables` VALUES (1,2,1,3,0,4500.0000,'2017-03-07','2018-07-08','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0);
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TWS`
--

LOCK TABLES `TWS` WRITE;
/*!40000 ALTER TABLE `TWS` DISABLE KEYS */;
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`TDID`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TaskDescriptor`
--

LOCK TABLES `TaskDescriptor` WRITE;
/*!40000 ALTER TABLE `TaskDescriptor` DISABLE KEYS */;
INSERT INTO `TaskDescriptor` VALUES (1,0,1,'Delinquency Report','Manual','2018-01-31 20:00:00','2018-01-20 20:00:00',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(2,0,1,'Walk the Units','Manual','2018-01-31 20:00:00','2018-01-20 20:00:00',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(3,0,1,'Generate Offsets','OffsetBot','2018-01-31 20:00:00','2018-01-20 20:00:00',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0);
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
  `Name` varchar(256) NOT NULL DEFAULT '',
  `Cycle` bigint(20) NOT NULL DEFAULT '0',
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
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` bigint(20) NOT NULL DEFAULT '0',
  `CreateTS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`TLDID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TaskListDefinition`
--

LOCK TABLES `TaskListDefinition` WRITE;
/*!40000 ALTER TABLE `TaskListDefinition` DISABLE KEYS */;
INSERT INTO `TaskListDefinition` VALUES (1,1,'Monthly Close',6,'2018-01-01 00:00:00','2018-01-31 17:00:00','2018-01-20 17:00:00',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Transactant`
--

LOCK TABLES `Transactant` WRITE;
/*!40000 ALTER TABLE `Transactant` DISABLE KEYS */;
INSERT INTO `Transactant` VALUES (1,1,1,'Joshua','Cudworth','Jones','Billy Bob','ABC Corporation',0,'JJones@abc.com','quintilian@nethersole.uk','','123-456-7890','123 South Rexford','','Beverly Hills','CA','90210','USA','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(2,1,2,'','','','','ABC Corporation',1,'info@abc.com','','','123-983-2244','10 S. ABC Pkwy','','Hollywood','CA','90220','USA','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(3,1,0,'Amanda','','Smith','','',0,'asmith@gmail.com','','','123-456-7891','2211 Hyannis Circle','','Glendale','CA','90210','USA','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(4,1,0,'Christoph','','Jones','','ARC Energy',0,'cjones@arcenergy.com','','','123-456-7892','1957 Billow Way','','LA','CA','90210','USA','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(5,1,0,'','','','','ARC Energy',1,'info@arcenergy.com','','','123-945-3715','9000 Zap Ave','','LA','CA','90220','USA','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(6,1,0,'Cindy','','Jones','','',0,'CindyJ@rexford.com','','','123-456-7893','309 1/2 South Rexford','','Beverly Hills','CA','90210','USA','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(7,1,0,'Edith','','Jones','','',0,'EdithJ@rexford.com','','','123-456-8516','309 1/2 South Rexford','','Beverly Hills','CA','90210','USA','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(8,1,0,'Ed','','Jones','','',0,'edj@rexford.com','','','123-456-7894','311 South Rexford','','Beverly Hills','CA','90210','USA','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(9,1,0,'Amanda','','Curry','','',0,'acurry@rexford.com','','','123-456-2337','311 South Rexford','','Beverly Hills','CA','90210','USA','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(10,1,0,'Jeb','','Night','','',0,'jnight@rexford.com','','','123-456-8519','311 South Rexford','','Beverly Hills','CA','90210','USA','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(11,1,0,'Chris','P','Bacon','','',0,'chris@bacon.com','','','123-456-6571','311 South Rexford','','Beverly Hills','CA','90210','USA','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(12,1,0,'Betty','','Diddit','','',0,'betty@diddit.com','','','123-456-9435','311 South Rexford','','Beverly Hills','CA','90210','USA','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(13,1,0,'Betty','','Diddent','','',0,'betty@diddent.com','','','123-456-0055','311 South Rexford','','Beverly Hills','CA','90210','USA','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(14,1,0,'Leeloo','','Kaixin','','',0,'leeloo@kaixin.com','','','123-456-0066','311 South Rexford','','Beverly Hills','CA','90210','USA','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(15,1,0,'','','','','School of Construction',1,'info@soc.com','','','123-456-9467','3000 Grand Avenue','','Glendale','CA','90210','USA','','2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0);
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
INSERT INTO `User` VALUES (1,1,10,'0000-00-00','Howard Hughes','Danvers State Mental Hospital, Massachusetts','BR549','','',0,'',1,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(2,1,20,'0000-00-00','','','','','',0,'',2,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(3,1,30,'0000-00-00','','','','','',0,'',3,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(4,1,40,'0000-00-00','','','','','',0,'',4,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(5,1,50,'0000-00-00','','','','','',0,'',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(6,1,60,'0000-00-00','','','','','',0,'',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(7,1,70,'0000-00-00','','','','','',0,'',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(8,1,80,'0000-00-00','','','','','',0,'',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(9,1,90,'0000-00-00','','','','','',0,'',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(10,1,100,'0000-00-00','','','','','',0,'',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(11,1,0,'0000-00-00','','','','','',0,'',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(12,1,0,'0000-00-00','','','','','',0,'',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(13,1,0,'0000-00-00','','','','','',0,'',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(14,1,0,'0000-00-00','','','','','',0,'',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0),(15,1,0,'0000-00-00','','','','','',0,'',0,'2018-03-14 19:50:32',0,'2018-03-14 19:50:32',0);
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

--
-- Table structure for table `classes`
--

DROP TABLE IF EXISTS `classes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `classes` (
  `ClassCode` mediumint(9) NOT NULL AUTO_INCREMENT,
  `CoCode` mediumint(9) NOT NULL DEFAULT '0',
  `Name` varchar(25) NOT NULL DEFAULT '',
  `Designation` char(3) NOT NULL DEFAULT '',
  `Description` varchar(256) NOT NULL DEFAULT '',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` mediumint(9) NOT NULL DEFAULT '0',
  PRIMARY KEY (`ClassCode`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `classes`
--

LOCK TABLES `classes` WRITE;
/*!40000 ALTER TABLE `classes` DISABLE KEYS */;
INSERT INTO `classes` VALUES (1,44,'L\'Objet USA LLC','USA','Owns all L\'Objet designs for North America, and sells at wholesale within North America.','2017-01-10 20:53:38',198),(2,42,'Texas Cedars, Ltd','TUC','Owns the Tucasa Townhomes in Irving, TX (rentals).','2017-01-10 20:53:19',198),(3,32,'NYC Luxury Reps, Inc.','NYR','Operates a showroom at 41 Madison for sale of L\'Objet products.','2017-01-10 20:49:57',198),(4,20,'Accord/OKC Members, LLC','OKC','Owns the Isola Bella Apartments in Oklahoma City (serviced apartments and transient accommodations).','2017-01-10 19:33:21',211),(5,4,'Accord Interests','AII','Primary operating company for all Accord activities.','2016-09-30 17:12:32',200),(6,19,'L\'Objet IOM Limited','IOM','Owns L\'Objet designs outside of North America, and sells at wholesale in all parts of the world outside of North America.','2017-01-10 20:49:36',198),(8,37,'Accord/PAC Members LLC','PAC','Owns the warehouse complex at 13571 Vaughn Street, San Fernando, CA (rentals).','2017-01-10 20:50:50',198),(9,14,'California Commerce Centr','CCC','Owners\' Association; owns and maintains common areas of warehouse complex at 13571 Vaughn Street, San Fernando, CA.','2017-01-10 20:49:13',198),(10,24,'309 South Rexford Drive','REX','Owns three townhomes and one apartment at 309 S. Rexford, Beverly Hills, CA (rentals).','2017-01-10 20:51:42',198),(11,24,'1775 Summitridge Drive','SUM','Owns the home at 1775 Summitridge Drive, Beverly Hills, CA (rental).','2017-01-10 20:52:57',198),(12,24,'JGM First, LLC','401','Owns the duplex property at 6401-6403 NW 63rd Street, Oklahoma City, OK (rentals).','2016-09-30 16:21:40',200),(13,76,'Accord/417 LLC','417','Owns the duplex property at 6417-6419 NW 63rd Street, Oklahoma City, OK (rentals).','2016-09-30 16:21:57',200),(14,3,'Acrosscity LDA','ACR','Owns the Portugal warehouse for storage and distribution of L\'Objet products.','2017-01-10 20:48:28',198),(15,34,'Accord/OKC Members LLC','OL2','Owns 231 West Olive Avenue, Burbank, CA; leased to Viacom for Nickelodeon Animation.\r\n','2017-01-10 20:50:24',198),(16,31,'8756 Wonderland Avenue','WON','Single Family Residence owned by MS1 and occupied by M. Sternbaum.','2017-01-10 20:54:06',198),(17,10,'Accord/BRO Members, LLC','BRO','Owns 203 W Olive Avenue, Burbank, CA; leased to Viacom for Nickelodeon Offices.','2016-09-30 16:15:11',200),(19,21,'JGM-BRA','BRA','424 Brandon Way\r\nAustin, TX 78733','2016-09-30 16:06:44',200),(20,21,'13656 Rayen Street','RAY','13656 W Rayen Street\r\nArleta, CA 91331','2017-01-10 20:52:32',198),(22,5,'Accord/JGM, LLC','AJM','Payroll Company','2016-09-30 20:36:35',200),(23,85,'Accord/EMP, LLC','EMP','Payroll Company','2016-09-30 16:29:26',200);
/*!40000 ALTER TABLE `classes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `companies`
--

DROP TABLE IF EXISTS `companies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `companies` (
  `CoCode` mediumint(9) NOT NULL AUTO_INCREMENT,
  `LegalName` varchar(50) NOT NULL DEFAULT '',
  `CommonName` varchar(50) NOT NULL DEFAULT '',
  `Address` varchar(35) NOT NULL DEFAULT '',
  `Address2` varchar(35) NOT NULL DEFAULT '',
  `City` varchar(25) NOT NULL DEFAULT '',
  `State` char(25) NOT NULL DEFAULT '',
  `PostalCode` varchar(10) NOT NULL DEFAULT '',
  `Country` varchar(25) NOT NULL DEFAULT '',
  `Phone` varchar(25) NOT NULL DEFAULT '',
  `Fax` varchar(25) NOT NULL DEFAULT '',
  `Email` varchar(35) NOT NULL DEFAULT '',
  `Designation` char(3) NOT NULL DEFAULT '',
  `Active` smallint(6) NOT NULL DEFAULT '0',
  `EmploysPersonnel` smallint(6) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` mediumint(9) NOT NULL DEFAULT '0',
  PRIMARY KEY (`CoCode`)
) ENGINE=InnoDB AUTO_INCREMENT=86 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `companies`
--

LOCK TABLES `companies` WRITE;
/*!40000 ALTER TABLE `companies` DISABLE KEYS */;
INSERT INTO `companies` VALUES (1,'Radford Place Associates','Inactive','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','RPA',0,0,'2016-02-17 21:21:18',200),(2,'Amal Brdoukan Mansour','','424 Brandon Way','','Austin','TX','78733','USA','818-516-9270','','amalbmansour@gmail.com','ABM',1,0,'2015-12-09 08:01:24',0),(3,'Acrosscity LDA','','Zona Industrial do Casal da Areia','Lote 55  Casal da Areia Distrito','Alcobaca','Freguesia','206 COZ','Portugal','351-1262-545-271','','','ACR',1,0,'2016-02-17 21:10:03',200),(4,'Accord Interests, LLC','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','323-512-0111','323-512-0105','djn@accordinterests.com','AII',1,1,'2016-09-30 17:11:44',200),(5,'Accord/JGM, LLC','','11719 Bee Cave Road','','Austin','TX','78738','','512-600-1880','','djn@accordinterests.com','AJM',1,1,'2016-09-30 17:13:55',200),(6,'Anna Joan Mansour','','424 Brandon Way','','Austin','TX','78738','USA','818-424-5534','','','ANN',1,0,'2015-12-09 08:01:24',0),(7,'Accord/BRO CLY, LLC','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','BCR',1,0,'2016-02-17 21:10:15',200),(8,'Accord/BRO, LLC','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','323-512-0111','323-512-0105','','BRG',1,0,'2016-09-30 15:47:03',200),(9,'Accord/BRO Equity, LLC','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','323-512-0111','323-512-0105','','BRE',1,0,'2016-09-30 15:47:17',200),(10,'Accord/BRO Members, LLC','203 W Olive Avenue Burbank CA','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','323-512-0111','323-512-0105','','BRO',1,0,'2016-09-30 15:46:31',200),(11,'BVS CBP Holdings, LLC','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','BVS',1,0,'2015-12-09 08:01:24',0),(12,'CBP Member LLC','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','CBM',1,0,'2016-02-17 21:10:40',200),(13,'Ceruzzi-Beaumont Partners, LLC','Inactive','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','CBP',1,0,'2016-09-30 15:34:20',200),(14,'California Commerce Centre Owners Association','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','CCC',1,0,'2016-02-17 21:10:58',200),(15,'Darla J. Nelson','','119 The Hills Drive','','Austin','TX','78738','USA','512-502-5486','','','DJN',1,0,'2015-12-09 08:01:24',0),(16,'Elizabeth Joan Mansour','','424 Brandon Way','','Austin','TX','78733','USA','310-987-1223','','','ELI',1,0,'2015-12-09 08:01:24',0),(17,'Elad Yifrach','','1203 David Drive','','Euless','TX','76040','USA','310-897-8498','','','EY',1,0,'2015-12-09 08:01:24',0),(18,'Accord Guaranty LP','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','','','','GUA',1,0,'2016-02-17 21:13:26',200),(19,'L\'Objet IOM Limited','','Ground Floor West Suite','54-58 Athol Street','Douglas','','1M11JD','Isle of Man','','','','IOM',1,1,'2016-02-17 21:14:32',200),(20,'Accord/OKC Members, LLC','Isola Bella','6303 NW 63rd Street','','Oklahoma City','OK','73132','USA','405-721-2191','405-603-4095','','OKC',1,1,'2016-09-30 15:32:42',200),(21,'Joseph G. Mansour','','424 Brandon Way','','Austin','TX','78733','USA','310-245-8220','','','JGM',1,0,'2015-12-09 08:01:24',0),(22,'Joan L.Mansour','','424 Brandon Way','','Austin','TX','78733','USA','310-600-1723','','','JLM',1,0,'2015-12-09 08:01:24',0),(23,'Jacob Louis Sternbaum','','1602 Lakeway Boulevard','','Lakeway','TX','78734','USA','818-231-1970','','','JLS',1,0,'2015-12-09 08:01:24',0),(24,'JGM First, LLC','SUM REX 417 Rentals','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','JM1',1,0,'2016-09-30 15:43:01',200),(25,'Jacob Sternbaum Trust','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','JS1',1,0,'2016-02-17 21:14:51',200),(26,'Joseph G. Mansour Family Trust 1','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','JT1',1,0,'2016-02-17 21:15:04',200),(27,'KLS First, LLC','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','KS1',1,0,'2016-02-17 21:15:33',200),(28,'L\'Objet UK Limited','','87-135 Brompton Road','Knightsbridge','London','','SW1X7XL','United Kingdom','','','','LUK',1,0,'2016-02-17 21:16:07',200),(29,'Maxxe E  Sternbaum','','8756 Wonderland Avenue','','Los Angeles','CA','90046','USA','818.461.4141','','','MES',1,0,'2016-02-17 21:16:43',200),(30,'Accord/MOR, LLC','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','MOR',1,0,'2016-02-17 21:09:38',200),(31,'Maxxe Sternbaum Trust','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','323.512.0111','323.512.0105','','MS1',1,0,'2016-02-17 21:16:58',200),(32,'NYC Luxury Reps, Inc.','','41 Madison Avenue','16th Floor','New York','NY','10010','USA','212-251-1011','212-251-1012','','NYR',1,0,'2016-02-17 21:17:42',200),(33,'Accord/OKC, LLC','Inactive','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','OKG',1,0,'2016-02-17 21:25:19',200),(34,'Accord/OLI Members, LLC','231 W Olive Avenue Burbank CA','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','OL2',1,0,'2016-02-17 21:19:15',200),(35,'Accord/OLI Equity, LLC','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','OLE',1,0,'2016-02-17 21:19:44',200),(36,'231 W Olive Partners, Ltd','Inactive','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','OLI',1,0,'2016-02-17 21:19:58',200),(37,'Accord/PAC Members, LLC','13571 Vaughn San Fernando CA','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','PAC',1,0,'2016-09-30 15:42:14',200),(38,'BH Rexford Group, Inc.','REX Owners Association','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','REH',1,0,'2015-12-09 08:01:24',0),(39,'SAT Isola Equity, LLC','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','SAE',1,0,'2016-02-17 21:26:05',200),(40,'SAT Isola Manager, LLC','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','SAG',1,0,'2016-02-17 21:26:23',200),(41,'SAT Isola Partners, LP','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','SAT',1,0,'2016-02-17 21:26:36',200),(42,'Texas Cedars, Ltd.','Tucasa Townhomes','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','TUC',1,1,'2016-02-17 21:26:54',200),(43,'Tucasa Corporation','Inactive','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','TUG',1,0,'2016-09-30 15:36:51',200),(44,'L\'Objet USA, LLC','','3515 Conflans Road','','Irving','TX','75061','USA','972-986-9575','972-767-4014','','USA',1,1,'2016-02-17 21:23:18',200),(45,'Accord/VAU Members, LLC','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','VAU',1,0,'2016-02-17 21:27:22',200),(46,'Accord/WIL Members, LLC','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','','WIL',1,0,'2016-02-17 21:27:35',200),(47,'Accord Group, Inc.','Inactive','','12345','','','','','','','','AGI',0,1,'2016-01-06 18:04:37',200),(48,'Accord Valet Services, LLC','Inactive','','','','','','','','','','AVS',0,0,'2016-01-06 18:07:08',200),(49,'Burbank Commerce Centre Industrial, Inc.','Inactive','','','','','','','','','','BCC',0,0,'2016-01-06 18:04:11',200),(50,'Accord/BEV, LLC','Inactive','','','','','','','','','','BEG',0,0,'2016-02-17 21:03:17',200),(51,'Accord/BEV Associates, L.P.','Inactive','','','','','','','','','','BEV',0,0,'2016-02-17 21:03:33',200),(52,'Cedar Interests, Inc.','Inactive','','','','','','','','','','CII',0,0,'2016-01-06 18:06:45',200),(53,'Accord/CLY, LLC','Inactive','','','','','','','','','','CLG',0,0,'2016-01-06 18:05:17',200),(54,'Accord/CLY Members, LLC','Inactive','','','','','','','','','','CLY',0,0,'2016-01-06 18:05:34',200),(55,'Accord/HOL Members, LLC','Inactive','','','','','','','','','','HOL',0,0,'2016-02-17 21:14:12',200),(56,'JGM Family Trust 2','','','','','','','','','','','JT2',0,0,'2016-02-17 21:15:15',200),(57,'Karl Louis Sternbaum','','','','','','','','','','','KLS',0,0,'2016-02-17 21:15:24',200),(58,'Estate of Karl Sternbaum','Inactive','','','','','','','','','','KSE',0,0,'2016-09-30 15:35:12',200),(59,'KLS Family Trust 1','Inactive','','','','','','','','','','KT1',0,0,'2016-09-30 15:34:59',200),(60,'MANCO','Inactive','','','','','','','','','','MCO',0,0,'2016-02-17 21:16:21',200),(61,'Mill Creek Group, Ltd. Co.','Inactive','','','','','','','','','','MCT',0,0,'2016-02-17 21:16:32',200),(62,'Maxxe Sternbaum Life Insurance Trust','Inactive','','','','','','','','','','MSL',0,0,'2016-02-17 21:18:46',200),(63,'Accord/NE2, LLC','Inactive','','','','','','','','','','NE2',0,0,'2016-02-17 21:17:17',200),(64,'Accord/NEW Manager, LLC','Inactive','','','','','','','','','','NEG',0,0,'2016-02-17 21:18:28',200),(65,'Accord/PAS, LLC','Inactive','','','','','','','','','','PAG',0,0,'2016-02-17 21:20:42',200),(66,'Pasadena Hotel Associates, LLC','Inactive','','','','','','','','','','PAS',0,0,'2016-02-17 21:20:58',200),(68,'Accord/SHO, LLC','Inactive','','','','','','','','','','SHG',0,0,'2016-02-17 21:22:04',200),(69,'Accord/SHO Members, LLC','Inactive','','','','','','','','','','SHO',0,0,'2016-09-30 15:35:56',200),(70,'Accord/SIE, LLC','Inactive','','','','','','','','','','SIE',0,0,'2016-09-30 15:33:20',200),(71,'Sternbaum General Partnership','Inactive','','','','','','','','','','STE',0,0,'2016-09-30 15:36:34',200),(72,'Accord/VAU, LLC','Inactive','','','','','','','','','','VAG',0,0,'2016-02-17 21:23:33',200),(73,'Accord/VIL, LLC','Inactive','','','','','','','','','','VIG',0,0,'2016-02-17 21:23:56',200),(74,'Accord/VIL Members, LLC','Inactive','','','','','','','','','','VIL',0,0,'2016-02-17 21:24:09',200),(75,'Accord/WIL, LLC','Inactive','','','','','','','','','','WIG',0,0,'2016-02-17 21:24:21',200),(76,'Accord/417, LLC','6417-19 NW 63rd Street, OKC, OK','11719 Bee Cave Road (FM2244)','Suite 301','Austin','TX','78738','USA','512-600-1880','323-512-0105','djn@accordinterests.com','417',1,0,'2016-09-30 20:24:02',211),(77,'SAL Carmichael, LP','Inactive','7411 Fair Oaks Boulevard','','Carmichael','CA','','USA','','','','CAR',1,0,'2016-09-30 15:39:30',200),(78,'SAL Citrus Heights, LP','Inactive','8220 Sunrise Boulevard','','Citrus Heights','CA','','USA','','','','CIT',1,0,'2016-09-30 15:39:03',200),(79,'SAL Assisted Living, LP','Inactive','SWC Bella Breeze and East Joiner','','Lincoln','CA','','USA','','','','LIN',1,0,'2016-09-30 15:35:26',200),(80,'SAL Kern Canyon, LP','Inactive','2902 AG Spanos Boulevard','','Stockton','CA','','USA','','','','STO',1,0,'2016-09-30 15:37:17',200),(81,'SAL Westgate, LP','Inactive','2305 Jefferson Boulevard','','West Sacramento','CA','','USA','','','','WSG',1,0,'2016-09-30 15:33:39',200),(82,'L\'Objet Retail, LLC','','3515 Conflans Road','','Irving','TX','75061','USA','','','','LOR',1,0,'2016-09-30 15:40:54',200),(83,'ANCO Building Services, LLC','','424 Brandon Way','','Austin','TX','78733','USA','','','','ABS',1,0,'2016-09-30 15:44:01',200),(84,'FAA Today, LLC','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','','','FAA',1,0,'2016-09-30 16:05:14',200),(85,'Accord/EMP, LLC','','11719 Bee Cave Road','Suite 301','Austin','TX','78738','USA','512-600-1880','','','EMP',1,1,'2016-09-30 17:12:17',200);
/*!40000 ALTER TABLE `companies` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `compensation`
--

DROP TABLE IF EXISTS `compensation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `compensation` (
  `UID` mediumint(9) NOT NULL,
  `Type` mediumint(9) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `compensation`
--

LOCK TABLES `compensation` WRITE;
/*!40000 ALTER TABLE `compensation` DISABLE KEYS */;
INSERT INTO `compensation` VALUES (95,0),(60,0),(107,0),(171,0),(35,0),(66,0),(130,0),(182,0),(7,0),(90,0),(153,0),(174,0),(25,0),(117,0),(84,0),(33,0),(67,0),(131,0),(75,0),(178,0),(16,0),(97,2),(106,2),(108,0),(168,0),(91,0),(137,0),(36,0),(119,0),(157,0),(177,0),(63,0),(74,0),(151,0),(128,0),(86,0),(18,0),(1,0),(104,0),(2,0),(28,0),(54,0),(80,0),(164,0),(165,0),(118,0),(134,0),(159,0),(154,0),(92,0),(87,0),(138,0),(170,0),(162,0),(148,0),(113,0),(68,0),(79,0),(123,0),(8,0),(11,0),(179,0),(152,0),(184,0),(5,0),(64,0),(78,0),(180,0),(112,0),(39,0),(158,0),(183,0),(61,0),(15,0),(32,0),(6,0),(115,0),(98,2),(102,0),(31,0),(127,0),(12,0),(140,0),(166,0),(44,0),(139,0),(65,0),(121,0),(150,0),(43,0),(135,0),(48,0),(62,0),(167,0),(58,2),(116,4),(132,4),(133,2),(22,2),(172,2),(185,2),(57,2),(29,2),(141,1),(214,1),(214,3),(19,1),(27,2),(103,1),(21,2),(46,2),(71,2),(72,2),(100,2),(110,2),(198,1),(3,2),(4,4),(204,2),(13,2),(59,1),(14,2),(50,2),(207,1),(30,2),(51,2),(120,2),(49,2),(93,2),(212,2),(105,2),(161,2),(186,2),(99,2),(156,2),(73,1),(37,2),(17,2),(42,2),(56,2),(160,2),(213,2),(124,2),(82,2),(26,2),(83,2),(85,1),(77,2),(136,2),(169,2),(197,2),(34,2),(122,2),(188,2),(81,2),(142,2),(181,2),(189,2),(20,2),(23,4),(41,4),(45,4),(52,4),(69,4),(94,4),(96,2),(101,4),(125,2),(129,4),(144,4),(145,2),(176,4),(187,4),(146,1),(147,1),(149,1),(206,1),(70,2),(109,2),(126,2),(199,1),(191,1),(192,1),(38,1),(193,1),(195,2),(190,2),(194,2),(196,2),(55,1),(88,2),(163,1),(208,1),(114,2),(173,1),(175,1),(203,1),(210,1),(209,1),(211,1),(200,1),(89,1),(201,1),(47,2),(265,1),(266,1),(40,2),(10,4),(53,2),(76,2),(111,2),(24,2),(205,2),(202,1),(9,2);
/*!40000 ALTER TABLE `compensation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `counters`
--

DROP TABLE IF EXISTS `counters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `counters` (
  `SearchPeople` bigint(20) NOT NULL DEFAULT '0',
  `SearchClasses` bigint(20) NOT NULL DEFAULT '0',
  `SearchCompanies` bigint(20) NOT NULL DEFAULT '0',
  `EditPerson` bigint(20) NOT NULL DEFAULT '0',
  `ViewPerson` bigint(20) NOT NULL DEFAULT '0',
  `ViewClass` bigint(20) NOT NULL DEFAULT '0',
  `ViewCompany` bigint(20) NOT NULL DEFAULT '0',
  `AdminEditPerson` bigint(20) NOT NULL DEFAULT '0',
  `AdminEditClass` bigint(20) NOT NULL DEFAULT '0',
  `AdminEditCompany` bigint(20) NOT NULL DEFAULT '0',
  `DeletePerson` bigint(20) NOT NULL DEFAULT '0',
  `DeleteClass` bigint(20) NOT NULL DEFAULT '0',
  `DeleteCompany` bigint(20) NOT NULL DEFAULT '0',
  `SignIn` bigint(20) NOT NULL DEFAULT '0',
  `Logoff` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `counters`
--

LOCK TABLES `counters` WRITE;
/*!40000 ALTER TABLE `counters` DISABLE KEYS */;
INSERT INTO `counters` VALUES (1921,284,415,117,1616,303,1414,186,76,139,0,2,4,627,159);
/*!40000 ALTER TABLE `counters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `deductionlist`
--

DROP TABLE IF EXISTS `deductionlist`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `deductionlist` (
  `DCode` mediumint(9) NOT NULL,
  `Name` varchar(25) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `deductionlist`
--

LOCK TABLES `deductionlist` WRITE;
/*!40000 ALTER TABLE `deductionlist` DISABLE KEYS */;
INSERT INTO `deductionlist` VALUES (0,'Unknown'),(1,'401K'),(2,'401K Loan'),(3,'Child Support'),(4,'Dental'),(5,'FSA'),(6,'GARN'),(7,'Group Life'),(8,'Housing'),(9,'Medical'),(10,'Miscded'),(11,'Taxes');
/*!40000 ALTER TABLE `deductionlist` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `deductions`
--

DROP TABLE IF EXISTS `deductions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `deductions` (
  `UID` mediumint(9) NOT NULL,
  `Deduction` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `deductions`
--

LOCK TABLES `deductions` WRITE;
/*!40000 ALTER TABLE `deductions` DISABLE KEYS */;
INSERT INTO `deductions` VALUES (58,11),(116,4),(116,9),(116,11),(132,4),(132,9),(132,11),(133,4),(133,9),(133,11),(22,8),(22,11),(172,3),(172,11),(185,8),(185,11),(57,4),(57,8),(57,9),(57,11),(29,4),(29,9),(29,11),(141,4),(141,9),(141,11),(0,1),(0,3),(0,6),(0,8),(0,9),(0,11),(214,1),(214,3),(214,4),(214,5),(214,7),(214,8),(214,9),(214,11),(27,11),(27,11),(103,11),(103,11),(21,1),(21,4),(21,9),(21,11),(46,1),(46,4),(46,9),(46,11),(71,11),(72,11),(100,11),(110,4),(110,9),(110,11),(198,4),(198,7),(198,9),(198,11),(3,11),(4,4),(4,9),(4,11),(204,8),(204,11),(13,4),(13,9),(13,11),(59,4),(59,8),(59,11),(14,4),(14,8),(14,9),(14,11),(50,11),(207,1),(207,2),(207,4),(207,5),(207,7),(207,9),(207,11),(30,4),(30,6),(30,9),(30,11),(51,11),(120,11),(257,11),(49,11),(93,11),(212,11),(105,8),(105,11),(161,11),(186,8),(186,11),(259,11),(99,4),(99,8),(99,9),(99,11),(156,4),(156,8),(156,9),(156,11),(73,1),(73,4),(73,9),(73,11),(37,4),(37,8),(37,9),(37,10),(37,11),(17,11),(42,4),(42,9),(42,11),(56,4),(56,9),(56,11),(160,11),(213,11),(124,4),(124,8),(124,9),(124,11),(82,4),(82,9),(82,11),(26,8),(26,11),(83,4),(83,9),(83,11),(261,8),(260,8),(85,11),(262,11),(77,4),(77,9),(77,11),(136,11),(169,8),(169,11),(197,1),(197,4),(197,7),(197,9),(197,11),(34,3),(34,4),(34,8),(34,9),(34,11),(122,11),(188,3),(188,4),(188,8),(188,9),(188,11),(81,3),(81,4),(81,9),(81,11),(142,4),(142,8),(142,9),(142,11),(181,8),(181,9),(181,11),(189,4),(189,8),(189,9),(189,11),(20,4),(20,9),(20,11),(23,4),(23,9),(23,11),(41,4),(41,9),(41,11),(45,4),(45,9),(45,11),(52,11),(69,4),(69,9),(69,11),(94,4),(94,9),(94,11),(96,4),(96,8),(96,9),(96,11),(101,4),(101,9),(101,11),(125,4),(125,9),(125,11),(129,4),(129,9),(129,11),(144,4),(144,9),(144,11),(145,11),(176,4),(176,9),(176,11),(187,4),(187,9),(187,11),(146,4),(146,5),(146,9),(146,11),(147,4),(147,9),(147,11),(149,11),(206,1),(206,2),(206,4),(206,5),(206,7),(206,9),(206,11),(70,4),(70,9),(70,11),(109,4),(109,9),(109,11),(126,4),(126,8),(126,9),(126,11),(199,11),(191,11),(192,1),(192,11),(38,4),(38,9),(38,11),(193,1),(193,11),(195,11),(190,11),(194,11),(196,11),(55,1),(55,4),(55,5),(55,9),(55,11),(88,11),(163,4),(163,9),(163,11),(208,11),(114,4),(114,9),(114,11),(173,4),(173,8),(173,9),(173,11),(175,4),(175,9),(175,11),(203,11),(210,11),(209,11),(211,11),(200,4),(200,7),(200,9),(200,11),(89,4),(89,8),(89,9),(89,11),(201,4),(201,7),(201,9),(201,10),(201,11),(47,1),(47,4),(47,9),(47,11),(265,4),(265,7),(265,9),(265,11),(266,11),(40,4),(40,9),(40,11),(10,4),(10,9),(10,11),(53,11),(76,4),(76,8),(76,9),(76,11),(111,8),(111,11),(24,8),(24,11),(205,4),(205,7),(205,8),(205,9),(205,11),(202,1),(202,4),(202,5),(202,7),(202,9),(202,11),(258,11),(9,4),(9,8),(9,9),(9,11);
/*!40000 ALTER TABLE `deductions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `departments`
--

DROP TABLE IF EXISTS `departments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `departments` (
  `DeptCode` mediumint(9) NOT NULL AUTO_INCREMENT,
  `Name` varchar(25) DEFAULT NULL,
  PRIMARY KEY (`DeptCode`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `departments`
--

LOCK TABLES `departments` WRITE;
/*!40000 ALTER TABLE `departments` DISABLE KEYS */;
INSERT INTO `departments` VALUES (1,'Accounting'),(2,'Administrative'),(3,'Capital Improvements'),(4,'Courtesy Patrol'),(5,'Customer Service'),(6,'Fitness Center'),(7,'Food & Beverage'),(8,'Guest Services'),(9,'Housekeeping'),(10,'Landscaping'),(11,'Maintenance'),(12,'Product Development'),(13,'Product Sales'),(14,'Serviced Apt Sales'),(15,'Trad Apt Sales'),(16,'Unknown'),(17,'Warehouse');
/*!40000 ALTER TABLE `departments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `fieldperms`
--

DROP TABLE IF EXISTS `fieldperms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `fieldperms` (
  `RID` mediumint(9) NOT NULL,
  `Elem` mediumint(9) NOT NULL,
  `Field` varchar(25) NOT NULL,
  `Perm` mediumint(9) NOT NULL,
  `Descr` varchar(256) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fieldperms`
--

LOCK TABLES `fieldperms` WRITE;
/*!40000 ALTER TABLE `fieldperms` DISABLE KEYS */;
INSERT INTO `fieldperms` VALUES (1,1,'Status',23,'Indicates whether the person is an active employee.'),(1,1,'EligibleForRehire',23,'Indicates whether a past employee can be rehired.'),(1,1,'UID',19,'A unique identifier associated with the employee. Once created, it never changes.'),(1,1,'Salutation',31,'\'Mr.\', \'Mrs.\', \'Ms.\', etc.'),(1,1,'FirstName',31,'The person\'s common name.'),(1,1,'MiddleName',31,'The person\'s middle name.'),(1,1,'LastName',31,'The person\'s surname or last name.'),(1,1,'PreferredName',95,'Less formal name but more commonly used, for example \'Mike\' rather than \'Michael\'.'),(1,1,'PrimaryEmail',95,'The primary email address to use for this person.'),(1,1,'OfficePhone',95,'This person\'s office telephone number.'),(1,1,'CellPhone',95,'This person\'s cellphone number.'),(1,1,'EmergencyContactName',95,'Name of someone to contact in the event of an emergency.'),(1,1,'EmergencyContactPhone',95,'Phone number for the emergency contact.'),(1,1,'HomeStreetAddress',95,'def'),(1,1,'HomeStreetAddress2',95,'def'),(1,1,'HomeCity',95,'def'),(1,1,'HomeState',95,'def'),(1,1,'HomePostalCode',95,'def'),(1,1,'HomeCountry',95,'def'),(1,1,'PrimaryEmail',95,'def'),(1,1,'SecondaryEmail',95,'def'),(1,1,'OfficePhone',95,'def'),(1,1,'OfficeFax',95,'def'),(1,1,'CellPhone',95,'def'),(1,1,'BirthDOM',31,'def'),(1,1,'BirthMonth',31,'def'),(1,1,'CoCode',31,'The company code associated with this user.'),(1,1,'JobCode',31,'def'),(1,1,'ClassCode',31,'def'),(1,1,'DeptCode',31,'def'),(1,1,'PositionControlNumber',31,'def'),(1,1,'MgrUID',31,'def'),(1,1,'Accepted401K',31,'def'),(1,1,'AcceptedDentalInsurance',31,'def'),(1,1,'AcceptedHealthInsurance',31,'def'),(1,1,'Hire',31,'def'),(1,1,'Termination',31,'def'),(1,1,'LastReview',31,'def'),(1,1,'NextReview',31,'def'),(1,1,'StateOfEmployment',31,'def'),(1,1,'CountryOfEmployment',31,'def'),(1,1,'Comps',31,'def'),(1,1,'Deductions',31,'def'),(1,1,'MyDeductions',31,'def'),(1,1,'RID',31,'def'),(1,1,'Role',31,'Permissions role'),(1,1,'ElemEntity',31,'Permissions to delete the entity'),(1,2,'CoCode',31,'def'),(1,2,'LegalName',31,'def'),(1,2,'CommonName',31,'def'),(1,2,'Address',31,'def'),(1,2,'Address2',31,'def'),(1,2,'City',31,'def'),(1,2,'State',31,'def'),(1,2,'PostalCode',31,'def'),(1,2,'Country',31,'def'),(1,2,'Phone',31,'def'),(1,2,'Fax',31,'def'),(1,2,'Email',31,'def'),(1,2,'Designation',31,'def'),(1,2,'Active',31,'def'),(1,2,'EmploysPersonnel',31,'def'),(1,2,'ElemEntity',31,'def'),(1,3,'ClassCode',31,'def'),(1,3,'CoCode',31,'The parent company for this business unit'),(1,3,'Name',31,'def'),(1,3,'Designation',31,'def'),(1,3,'Description',31,'def'),(1,3,'ElemEntity',31,'def'),(1,4,'Shutdown',256,'Permission to shutdown the service'),(1,4,'Restart',256,'Permission to restart the service'),(2,1,'Status',23,'Indicates whether the person is an active employee.'),(2,1,'EligibleForRehire',23,'Indicates whether a past employee can be rehired.'),(2,1,'UID',19,'A unique identifier associated with the employee. Once created, it never changes.'),(2,1,'Salutation',95,'\'Mr.\', \'Mrs.\', \'Ms.\', etc.'),(2,1,'FirstName',95,'The person\'s common name.'),(2,1,'MiddleName',95,'The person\'s middle name.'),(2,1,'LastName',95,'The person\'s surname or last name.'),(2,1,'PreferredName',95,'Less formal name but more commonly used, for example \'Mike\' rather than \'Michael\'.'),(2,1,'PrimaryEmail',95,'The primary email address to use for this person.'),(2,1,'OfficePhone',95,'This person\'s office telephone number.'),(2,1,'CellPhone',95,'This person\'s cellphone number.'),(2,1,'EmergencyContactName',95,'Name of someone to contact in the event of an emergency.'),(2,1,'EmergencyContactPhone',95,'Phone number for the emergency contact.'),(2,1,'HomeStreetAddress',95,'def'),(2,1,'HomeStreetAddress2',95,'def'),(2,1,'HomeCity',95,'def'),(2,1,'HomeState',95,'def'),(2,1,'HomePostalCode',95,'def'),(2,1,'HomeCountry',95,'def'),(2,1,'PrimaryEmail',95,'def'),(2,1,'SecondaryEmail',95,'def'),(2,1,'OfficePhone',95,'def'),(2,1,'OfficeFax',95,'def'),(2,1,'CellPhone',95,'def'),(2,1,'BirthDOM',31,'def'),(2,1,'BirthMonth',31,'def'),(2,1,'CoCode',31,'The company code associated with this user.'),(2,1,'JobCode',31,'def'),(2,1,'DeptCode',31,'def'),(2,1,'ClassCode',31,'def'),(2,1,'PositionControlNumber',31,'def'),(2,1,'MgrUID',31,'def'),(2,1,'Accepted401K',31,'def'),(2,1,'AcceptedDentalInsurance',31,'def'),(2,1,'AcceptedHealthInsurance',31,'def'),(2,1,'Hire',31,'def'),(2,1,'Termination',31,'def'),(2,1,'LastReview',31,'def'),(2,1,'NextReview',31,'def'),(2,1,'StateOfEmployment',31,'def'),(2,1,'CountryOfEmployment',31,'def'),(2,1,'Comps',31,'def'),(2,1,'Deductions',31,'def'),(2,1,'MyDeductions',31,'def'),(2,1,'Role',1,'Permissions Role'),(2,1,'RID',17,'def'),(2,1,'ElemEntity',0,'Permissions to delete the entity'),(2,2,'CoCode',17,'def'),(2,2,'LegalName',17,'def'),(2,2,'CommonName',17,'def'),(2,2,'Address',17,'def'),(2,2,'Address2',17,'def'),(2,2,'City',17,'def'),(2,2,'State',17,'def'),(2,2,'PostalCode',17,'def'),(2,2,'Country',17,'def'),(2,2,'Phone',17,'def'),(2,2,'Fax',17,'def'),(2,2,'Email',17,'def'),(2,2,'Designation',17,'def'),(2,2,'Active',17,'def'),(2,2,'EmploysPersonnel',17,'def'),(2,2,'ElemEntity',0,'def'),(2,3,'ClassCode',17,'def'),(2,3,'CoCode',17,'def'),(2,3,'Name',17,'The parent company for this business unit'),(2,3,'Designation',17,'def'),(2,3,'Description',17,'def'),(2,3,'ElemEntity',0,'def'),(2,4,'Shutdown',0,'Permission to shutdown the service'),(2,4,'Restart',0,'Permission to restart the service'),(3,1,'Status',17,'Indicates whether the person is an active employee.'),(3,1,'EligibleForRehire',23,'Indicates whether a past employee can be rehired.'),(3,1,'UID',19,'A unique identifier associated with the employee. Once created, it never changes.'),(3,1,'Salutation',17,'\'Mr.\', \'Mrs.\', \'Ms.\', etc.'),(3,1,'FirstName',17,'The person\'s common name.'),(3,1,'MiddleName',17,'The person\'s middle name.'),(3,1,'LastName',17,'The person\'s surname or last name.'),(3,1,'PreferredName',81,'Less formal name but more commonly used, for example \'Mike\' rather than \'Michael\'.'),(3,1,'PrimaryEmail',81,'The primary email address to use for this person.'),(3,1,'OfficePhone',81,'This person\'s office telephone number.'),(3,1,'CellPhone',81,'This person\'s cellphone number.'),(3,1,'EmergencyContactName',112,'Name of someone to contact in the event of an emergency.'),(3,1,'EmergencyContactPhone',112,'Phone number for the emergency contact.'),(3,1,'HomeStreetAddress',112,'def'),(3,1,'HomeStreetAddress2',112,'def'),(3,1,'HomeCity',112,'def'),(3,1,'HomeState',112,'def'),(3,1,'HomePostalCode',112,'def'),(3,1,'HomeCountry',112,'def'),(3,1,'PrimaryEmail',81,'def'),(3,1,'SecondaryEmail',81,'def'),(3,1,'OfficePhone',81,'def'),(3,1,'OfficeFax',81,'def'),(3,1,'CellPhone',81,'def'),(3,1,'BirthDOM',48,'def'),(3,1,'BirthMonth',48,'def'),(3,1,'CoCode',17,'The company code associated with this user.'),(3,1,'JobCode',17,'def'),(3,1,'DeptCode',17,'def'),(3,1,'ClassCode',17,'def'),(3,1,'MgrUID',17,'def'),(3,1,'Accepted401K',17,'def'),(3,1,'AcceptedDentalInsurance',17,'def'),(3,1,'AcceptedHealthInsurance',17,'def'),(3,1,'PositionControlNumber',17,'def'),(3,1,'Hire',48,'def'),(3,1,'Termination',17,'def'),(3,1,'LastReview',0,'def'),(3,1,'NextReview',0,'def'),(3,1,'StateOfEmployment',17,'def'),(3,1,'CountryOfEmployment',17,'def'),(3,1,'Comps',17,'def'),(3,1,'Deductions',17,'def'),(3,1,'MyDeductions',17,'def'),(3,1,'RID',0,'def'),(3,1,'Role',0,'Permissions Role'),(3,1,'ElemEntity',0,'Permissions to delete the entity'),(3,2,'CoCode',31,'def'),(3,2,'LegalName',31,'def'),(3,2,'CommonName',31,'def'),(3,2,'Address',31,'def'),(3,2,'Address2',31,'def'),(3,2,'City',31,'def'),(3,2,'State',31,'def'),(3,2,'PostalCode',31,'def'),(3,2,'Country',31,'def'),(3,2,'Phone',31,'def'),(3,2,'Fax',31,'def'),(3,2,'Email',31,'def'),(3,2,'Designation',31,'def'),(3,2,'Active',31,'def'),(3,2,'EmploysPersonnel',31,'def'),(3,2,'ElemEntity',0,'def'),(3,3,'ClassCode',31,'def'),(3,3,'CoCode',31,'The parent company for this business unit'),(3,3,'Name',31,'def'),(3,3,'Designation',31,'def'),(3,3,'Description',31,'def'),(3,3,'ElemEntity',0,'def'),(3,4,'Shutdown',0,'Permission to shutdown the service'),(3,4,'Restart',0,'Permission to restart the service'),(4,1,'Status',1,'Indicates whether the person is an active employee.'),(4,1,'EligibleForRehire',1,'Indicates whether a past employee can be rehired.'),(4,1,'UID',17,'A unique identifier associated with the employee. Once created, it never changes.'),(4,1,'Salutation',1,'\'Mr.\', \'Mrs.\', \'Ms.\', etc.'),(4,1,'FirstName',1,'The person\'s common name.'),(4,1,'MiddleName',1,'The person\'s middle name.'),(4,1,'LastName',1,'The person\'s surname or last name.'),(4,1,'PreferredName',65,'Less formal name but more commonly used, for example \'Mike\' rather than \'Michael\'.'),(4,1,'PrimaryEmail',65,'The primary email address to use for this person.'),(4,1,'OfficePhone',65,'This person\'s office telephone number.'),(4,1,'CellPhone',65,'This person\'s cellphone number.'),(4,1,'EmergencyContactName',193,'Name of someone to contact in the event of an emergency.'),(4,1,'EmergencyContactPhone',193,'Phone number for the emergency contact.'),(4,1,'HomeStreetAddress',193,'def'),(4,1,'HomeStreetAddress2',193,'def'),(4,1,'HomeCity',193,'def'),(4,1,'HomeState',193,'def'),(4,1,'HomePostalCode',193,'def'),(4,1,'HomeCountry',81,'def'),(4,1,'PrimaryEmail',81,'def'),(4,1,'SecondaryEmail',81,'def'),(4,1,'OfficePhone',81,'def'),(4,1,'OfficeFax',81,'def'),(4,1,'CellPhone',81,'def'),(4,1,'BirthDOM',160,'def'),(4,1,'BirthMonth',160,'def'),(4,1,'CoCode',160,'The company code associated with this user.'),(4,1,'JobCode',160,'def'),(4,1,'DeptCode',160,'def'),(4,1,'ClassCode',17,'def'),(4,1,'PositionControlNumber',160,'def'),(4,1,'MgrUID',17,'def'),(4,1,'Accepted401K',160,'def'),(4,1,'AcceptedDentalInsurance',160,'def'),(4,1,'AcceptedHealthInsurance',160,'def'),(4,1,'Hire',160,'def'),(4,1,'Termination',32,'def'),(4,1,'LastReview',32,'def'),(4,1,'NextReview',32,'def'),(4,1,'StateOfEmployment',160,'def'),(4,1,'CountryOfEmployment',160,'def'),(4,1,'Comps',160,'Compensation type(s) for this person.'),(4,1,'Deductions',160,'The deductions for this person.'),(4,1,'MyDeductions',160,'The deductions for this person.'),(4,1,'RID',17,'def'),(4,1,'Role',0,'Permissions Rol'),(4,1,'ElemEntity',0,'Permissions to delete the entity'),(4,2,'CoCode',1,'def'),(4,2,'LegalName',1,'def'),(4,2,'CommonName',1,'def'),(4,2,'Address',1,'def'),(4,2,'Address2',1,'def'),(4,2,'City',1,'def'),(4,2,'State',1,'def'),(4,2,'PostalCode',1,'def'),(4,2,'Country',1,'def'),(4,2,'Phone',1,'def'),(4,2,'Fax',1,'def'),(4,2,'Email',1,'def'),(4,2,'Designation',1,'def'),(4,2,'Active',1,'def'),(4,2,'EmploysPersonnel',1,'def'),(4,3,'ClassCode',1,'def'),(4,3,'CoCode',1,'The parent company for this business unit'),(4,3,'Name',1,'def'),(4,3,'Designation',1,'def'),(4,3,'Description',1,'def'),(4,3,'ElemEntity',0,'def'),(4,4,'Shutdown',0,'Permission to shutdown the service'),(4,4,'Restart',0,'Permission to restart the service'),(5,1,'Status',23,'Indicates whether the person is an active employee.'),(5,1,'EligibleForRehire',1,'Indicates whether a past employee can be rehired.'),(5,1,'UID',3,'A unique identifier associated with the employee. Once created, it never changes.'),(5,1,'Salutation',4,'\'Mr.\', \'Mrs.\', \'Ms.\', etc.'),(5,1,'FirstName',8,'The person\'s common name.'),(5,1,'MiddleName',16,'The person\'s middle name.'),(5,1,'LastName',0,'The person\'s surname or last name.'),(5,1,'PreferredName',17,'Less formal name but more commonly used, for example \'Mike\' rather than \'Michael\'.'),(5,1,'PrimaryEmail',1,'The primary email address to use for this person.'),(5,1,'OfficePhone',0,'This person\'s office telephone number.'),(5,1,'CellPhone',7,'This person\'s cellphone number.'),(5,1,'EmergencyContactName',0,'Name of someone to contact in the event of an emergency.'),(5,1,'EmergencyContactPhone',95,'Phone number for the emergency contact.'),(5,1,'HomeStreetAddress',95,'def'),(5,1,'HomeStreetAddress2',1,'def'),(5,1,'HomeCity',0,'def'),(5,1,'HomeState',95,'def'),(5,1,'HomePostalCode',0,'def'),(5,1,'HomeCountry',95,'def'),(5,1,'PrimaryEmail',95,'def'),(5,1,'SecondaryEmail',0,'def'),(5,1,'OfficePhone',95,'def'),(5,1,'OfficeFax',0,'def'),(5,1,'CellPhone',95,'def'),(5,1,'BirthDOM',0,'def'),(5,1,'BirthMonth',31,'def'),(5,1,'CoCode',0,'The company code associated with this user.'),(5,1,'JobCode',31,'def'),(5,1,'ClassCode',0,'def'),(5,1,'DeptCode',31,'def'),(5,1,'PositionControlNumber',0,'def'),(5,1,'MgrUID',31,'def'),(5,1,'Accepted401K',0,'def'),(5,1,'AcceptedDentalInsurance',31,'def'),(5,1,'AcceptedHealthInsurance',0,'def'),(5,1,'Hire',31,'def'),(5,1,'Termination',0,'def'),(5,1,'LastReview',31,'def'),(5,1,'NextReview',0,'def'),(5,1,'StateOfEmployment',31,'def'),(5,1,'CountryOfEmployment',0,'def'),(5,1,'Comps',31,'def'),(5,1,'Deductions',17,'def'),(5,1,'MyDeductions',17,'def'),(5,1,'RID',17,'def'),(5,1,'Role',0,'Permissions Rol'),(5,1,'ElemEntity',0,'Permissions to delete the entity'),(5,2,'CoCode',31,'def'),(5,2,'LegalName',0,'def'),(5,2,'CommonName',31,'def'),(5,2,'Address',31,'def'),(5,2,'Address2',0,'def'),(5,2,'City',31,'def'),(5,2,'State',31,'def'),(5,2,'PostalCode',31,'def'),(5,2,'Country',0,'def'),(5,2,'Phone',31,'def'),(5,2,'Fax',0,'def'),(5,2,'Email',31,'def'),(5,2,'Designation',31,'def'),(5,2,'Active',0,'def'),(5,2,'EmploysPersonnel',31,'def'),(5,2,'ElemEntity',31,'def'),(5,3,'ClassCode',31,'def'),(5,3,'CoCode',31,'The parent company for this business unit'),(5,3,'Name',31,'def'),(5,3,'Designation',31,'def'),(5,3,'Description',0,'def'),(5,3,'ElemEntity',31,'def'),(5,4,'Shutdown',256,'Permission to shutdown the service'),(5,4,'Restart',256,'Permission to restart the service'),(6,1,'Status',23,'Indicates whether the person is an active employee.'),(6,1,'EligibleForRehire',23,'Indicates whether a past employee can be rehired.'),(6,1,'UID',19,'A unique identifier associated with the employee. Once created, it never changes.'),(6,1,'Salutation',95,'\'Mr.\', \'Mrs.\', \'Ms.\', etc.'),(6,1,'FirstName',95,'The person\'s common name.'),(6,1,'MiddleName',95,'The person\'s middle name.'),(6,1,'LastName',95,'The person\'s surname or last name.'),(6,1,'PreferredName',95,'Less formal name but more commonly used, for example \'Mike\' rather than \'Michael\'.'),(6,1,'PrimaryEmail',95,'The primary email address to use for this person.'),(6,1,'OfficePhone',95,'This person\'s office telephone number.'),(6,1,'CellPhone',95,'This person\'s cellphone number.'),(6,1,'EmergencyContactName',95,'Name of someone to contact in the event of an emergency.'),(6,1,'EmergencyContactPhone',95,'Phone number for the emergency contact.'),(6,1,'HomeStreetAddress',95,'def'),(6,1,'HomeStreetAddress2',95,'def'),(6,1,'HomeCity',95,'def'),(6,1,'HomeState',95,'def'),(6,1,'HomePostalCode',95,'def'),(6,1,'HomeCountry',95,'def'),(6,1,'PrimaryEmail',95,'def'),(6,1,'SecondaryEmail',95,'def'),(6,1,'OfficePhone',95,'def'),(6,1,'OfficeFax',95,'def'),(6,1,'CellPhone',95,'def'),(6,1,'BirthDOM',31,'def'),(6,1,'BirthMonth',31,'def'),(6,1,'CoCode',31,'The company code associated with this user.'),(6,1,'JobCode',31,'def'),(6,1,'DeptCode',31,'def'),(6,1,'ClassCode',31,'def'),(6,1,'PositionControlNumber',31,'def'),(6,1,'MgrUID',31,'def'),(6,1,'Accepted401K',31,'def'),(6,1,'AcceptedDentalInsurance',31,'def'),(6,1,'AcceptedHealthInsurance',31,'def'),(6,1,'Hire',31,'def'),(6,1,'Termination',31,'def'),(6,1,'LastReview',31,'def'),(6,1,'NextReview',31,'def'),(6,1,'StateOfEmployment',31,'def'),(6,1,'CountryOfEmployment',31,'def'),(6,1,'Comps',31,'def'),(6,1,'Deductions',31,'def'),(6,1,'MyDeductions',31,'def'),(6,1,'Role',1,'Permissions Rol'),(6,1,'RID',17,'def'),(6,1,'ElemEntity',0,'Permissions to delete the entity'),(6,2,'CoCode',31,'def'),(6,2,'LegalName',31,'def'),(6,2,'CommonName',31,'def'),(6,2,'Address',31,'def'),(6,2,'Address2',31,'def'),(6,2,'City',31,'def'),(6,2,'State',31,'def'),(6,2,'PostalCode',31,'def'),(6,2,'Country',31,'def'),(6,2,'Phone',31,'def'),(6,2,'Fax',31,'def'),(6,2,'Email',31,'def'),(6,2,'Designation',31,'def'),(6,2,'Active',31,'def'),(6,2,'EmploysPersonnel',31,'def'),(6,2,'ElemEntity',0,'def'),(6,3,'ClassCode',31,'def'),(6,3,'CoCode',31,'The parent company for this business unit'),(6,3,'Name',31,'def'),(6,3,'Designation',31,'def'),(6,3,'Description',31,'def'),(6,3,'ElemEntity',0,'def'),(6,4,'Shutdown',0,'Permission to shutdown the service'),(6,4,'Restart',0,'Permission to restart the service'),(7,1,'Status',23,'Indicates whether the person is an active employee.'),(7,1,'EligibleForRehire',23,'Indicates whether a past employee can be rehired.'),(7,1,'UID',19,'A unique identifier associated with the employee. Once created, it never changes.'),(7,1,'Salutation',95,'\'Mr.\', \'Mrs.\', \'Ms.\', etc.'),(7,1,'FirstName',95,'The person\'s common name.'),(7,1,'MiddleName',95,'The person\'s middle name.'),(7,1,'LastName',95,'The person\'s surname or last name.'),(7,1,'PreferredName',95,'Less formal name but more commonly used, for example \'Mike\' rather than \'Michael\'.'),(7,1,'PrimaryEmail',95,'The primary email address to use for this person.'),(7,1,'OfficePhone',95,'This person\'s office telephone number.'),(7,1,'CellPhone',95,'This person\'s cellphone number.'),(7,1,'EmergencyContactName',95,'Name of someone to contact in the event of an emergency.'),(7,1,'EmergencyContactPhone',95,'Phone number for the emergency contact.'),(7,1,'HomeStreetAddress',95,'def'),(7,1,'HomeStreetAddress2',95,'def'),(7,1,'HomeCity',95,'def'),(7,1,'HomeState',95,'def'),(7,1,'HomePostalCode',95,'def'),(7,1,'HomeCountry',95,'def'),(7,1,'PrimaryEmail',95,'def'),(7,1,'SecondaryEmail',95,'def'),(7,1,'OfficePhone',95,'def'),(7,1,'OfficeFax',95,'def'),(7,1,'CellPhone',95,'def'),(7,1,'BirthDOM',31,'def'),(7,1,'BirthMonth',31,'def'),(7,1,'CoCode',31,'The company code associated with this user.'),(7,1,'JobCode',31,'def'),(7,1,'DeptCode',31,'def'),(7,1,'ClassCode',31,'def'),(7,1,'PositionControlNumber',31,'def'),(7,1,'MgrUID',31,'def'),(7,1,'Accepted401K',31,'def'),(7,1,'AcceptedDentalInsurance',31,'def'),(7,1,'AcceptedHealthInsurance',31,'def'),(7,1,'Hire',31,'def'),(7,1,'Termination',31,'def'),(7,1,'LastReview',31,'def'),(7,1,'NextReview',31,'def'),(7,1,'StateOfEmployment',31,'def'),(7,1,'CountryOfEmployment',31,'def'),(7,1,'Comps',31,'def'),(7,1,'Deductions',31,'def'),(7,1,'MyDeductions',31,'def'),(7,1,'Role',1,'Permissions Rol'),(7,1,'RID',17,'def'),(7,1,'ElemEntity',31,'Permissions to create/delete the entity'),(7,2,'CoCode',31,'def'),(7,2,'LegalName',31,'def'),(7,2,'CommonName',31,'def'),(7,2,'Address',31,'def'),(7,2,'Address2',31,'def'),(7,2,'City',31,'def'),(7,2,'State',31,'def'),(7,2,'PostalCode',31,'def'),(7,2,'Country',31,'def'),(7,2,'Phone',31,'def'),(7,2,'Fax',31,'def'),(7,2,'Email',31,'def'),(7,2,'Designation',31,'def'),(7,2,'Active',31,'def'),(7,2,'EmploysPersonnel',31,'def'),(7,2,'ElemEntity',31,'def'),(7,3,'ClassCode',31,'def'),(7,3,'CoCode',31,'The parent company of this business unit.'),(7,3,'Name',31,'def'),(7,3,'Designation',31,'def'),(7,3,'Description',31,'def'),(7,3,'ElemEntity',31,'def'),(7,4,'Shutdown',0,'Permission to shutdown the service'),(7,4,'Restart',0,'Permission to restart the service');
/*!40000 ALTER TABLE `fieldperms` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jobtitles`
--

DROP TABLE IF EXISTS `jobtitles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `jobtitles` (
  `JobCode` mediumint(9) NOT NULL AUTO_INCREMENT,
  `Title` varchar(40) NOT NULL DEFAULT '',
  `Descr` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`JobCode`)
) ENGINE=InnoDB AUTO_INCREMENT=87 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jobtitles`
--

LOCK TABLES `jobtitles` WRITE;
/*!40000 ALTER TABLE `jobtitles` DISABLE KEYS */;
INSERT INTO `jobtitles` VALUES (1,'Accounting Assistant',''),(2,'Accounting Associate',''),(3,'Accounting Manager',''),(4,'Administrative Assistant',''),(5,'Assistant Manager',''),(6,'Associate Developer',''),(7,'Chief Executive Officer',''),(8,'Chief Financial Officer',''),(9,'Chief Operating Officer',''),(10,'Chief Technology Officer',''),(11,'General Manager',''),(12,'HR & Payroll Manager',''),(13,'Intern',''),(14,'Night Auditor',''),(15,'Office Manager',''),(16,'Procurement Specialist',''),(17,'Special Projects Associate',''),(18,'Director of Procurements',''),(19,'Call Center Associate',''),(20,'Call Center Manager',''),(21,'Courtesy Patrol Driver',''),(22,'Courtesy Patrol Manager',''),(23,'Courtesy Patrol Officer',''),(24,'Courtesy Patrol Supervisor',''),(25,'Designer',''),(26,'Creative Director',''),(27,'Development Coordinator',''),(28,'Director of Fragrence',''),(29,'Director of Sales',''),(30,'Principal',''),(31,'Studio Manager',''),(32,'Visual Arts Director',''),(33,'Fitness Center Attendant',''),(34,'Fitness Center Manager',''),(35,'Food and Beverage Manager',''),(36,'Executive Chef',''),(37,'Food & Bev Associate',''),(38,'Bar Manager',''),(39,'Bartender',''),(40,'Host',''),(41,'Waitstaff',''),(42,'Guest Services Associate',''),(43,'Concierge Manager',''),(44,'Concierge',''),(45,'Guest Services Manager',''),(46,'Housekeeping Manager',''),(47,'Housekeeping Supervisor',''),(48,'Common Area Housekeeper',''),(49,'Laundry Associate',''),(50,'Laundry Attendant',''),(51,'Serviced Apt Housekeeper',''),(52,'Svc Apt Housekeeping Associate',''),(53,'Traditional Apt Housekeeper',''),(54,'Grounds Associate',''),(55,'Grounds Supervisor',''),(56,'Maintenance Associate',''),(57,'Maintenance Manager',''),(58,'Maintenance Supervisor',''),(59,'Makeready Associate',''),(60,'Cap Improvement Associate',''),(61,'Cap Improvement Supervisor',''),(62,'Customer Service Associate',''),(63,'Customer Service Manager',''),(64,'Product Sales Associate',''),(65,'Product Sales Manager',''),(66,'National Accounts Manager',''),(67,'Leasing Associate',''),(68,'Leasing Manager',''),(69,'Packer',''),(70,'Repair Room Staff',''),(71,'Seasonal Associate',''),(72,'Checker',''),(73,'Warehouse Associate',''),(74,'Warehouse Manager',''),(75,'Warehouse Supervisor',''),(76,'Retail Clerk',''),(77,'Store Manager',''),(78,'Asst Store Manager',''),(79,'Traditional Apt Housekeeping Assoc',''),(80,'Makeready Technician',''),(81,'Serviced Apt Sales Associate',''),(82,'Unknown',''),(83,'Construction Manager',''),(84,'Dishwasher',''),(85,'NY Retail Boutique Manager',''),(86,'Marketing & Sales Manager','');
/*!40000 ALTER TABLE `jobtitles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `people`
--

DROP TABLE IF EXISTS `people`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `people` (
  `UID` mediumint(9) NOT NULL AUTO_INCREMENT,
  `UserName` varchar(20) NOT NULL DEFAULT '',
  `LastName` varchar(25) NOT NULL DEFAULT '',
  `MiddleName` varchar(25) NOT NULL DEFAULT '',
  `FirstName` varchar(25) NOT NULL DEFAULT '',
  `PreferredName` varchar(25) NOT NULL DEFAULT '',
  `Salutation` varchar(10) NOT NULL DEFAULT '',
  `PositionControlNumber` varchar(10) NOT NULL DEFAULT '',
  `OfficePhone` varchar(25) NOT NULL DEFAULT '',
  `OfficeFax` varchar(25) NOT NULL DEFAULT '',
  `CellPhone` varchar(25) NOT NULL DEFAULT '',
  `PrimaryEmail` varchar(35) NOT NULL DEFAULT '',
  `SecondaryEmail` varchar(35) NOT NULL DEFAULT '',
  `BirthMonth` tinyint(4) NOT NULL DEFAULT '0',
  `BirthDoM` tinyint(4) NOT NULL DEFAULT '0',
  `HomeStreetAddress` varchar(35) NOT NULL DEFAULT '',
  `HomeStreetAddress2` varchar(25) NOT NULL DEFAULT '',
  `HomeCity` varchar(25) NOT NULL DEFAULT '',
  `HomeState` char(2) NOT NULL DEFAULT '',
  `HomePostalCode` varchar(10) NOT NULL DEFAULT '',
  `HomeCountry` varchar(25) NOT NULL DEFAULT '',
  `JobCode` mediumint(9) NOT NULL DEFAULT '0',
  `Hire` date NOT NULL DEFAULT '2000-01-01',
  `Termination` date NOT NULL DEFAULT '2000-01-01',
  `MgrUID` mediumint(9) NOT NULL DEFAULT '0',
  `DeptCode` mediumint(9) NOT NULL DEFAULT '0',
  `CoCode` mediumint(9) NOT NULL DEFAULT '0',
  `ClassCode` smallint(6) NOT NULL DEFAULT '0',
  `StateOfEmployment` varchar(25) NOT NULL DEFAULT '',
  `CountryOfEmployment` varchar(25) NOT NULL DEFAULT '',
  `EmergencyContactName` varchar(25) NOT NULL DEFAULT '',
  `EmergencyContactPhone` varchar(25) NOT NULL DEFAULT '',
  `Status` smallint(6) NOT NULL DEFAULT '0',
  `EligibleForRehire` smallint(6) NOT NULL DEFAULT '0',
  `AcceptedHealthInsurance` smallint(6) NOT NULL DEFAULT '0',
  `AcceptedDentalInsurance` smallint(6) NOT NULL DEFAULT '0',
  `Accepted401K` smallint(6) NOT NULL DEFAULT '0',
  `LastReview` date NOT NULL DEFAULT '2000-01-01',
  `NextReview` date NOT NULL DEFAULT '2000-01-01',
  `passhash` char(128) NOT NULL DEFAULT '',
  `RID` mediumint(9) NOT NULL DEFAULT '0',
  `LastModTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LastModBy` mediumint(9) NOT NULL DEFAULT '0',
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB AUTO_INCREMENT=271 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `people`
--

LOCK TABLES `people` WRITE;
/*!40000 ALTER TABLE `people` DISABLE KEYS */;
INSERT INTO `people` VALUES (1,'dacuna','Acuna','','Doris','','','','','','','','',10,21,'2400 S Macarther','Lote-321','Oklahoma City','OK','73128','',0,'2014-07-12','2014-07-24',0,0,23,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(2,'cadams','Adams','','Chase','','','','','','','','',1,15,'5961 N Seminole Rd','','Oklahoma City','OK','73132','',42,'2014-02-26','2014-09-05',0,7,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(3,'vagers','Agers','','Veronica','','','1066','','','','','bizee1roni-cagers@snet.net',10,23,'12600 N MacArthur Blvd','Apt 1619','Oklahoma City','OK','73142','',42,'2015-08-10','2015-10-22',49,1,20,4,'','','','',0,1,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',5,'2016-02-16 00:43:18',202),(4,'maguinaga','Aguinaga','','Maria','','','79','','','','','',12,1,'2139 NW 12th','','Oklahoma City','OK','73112','',52,'2011-08-13','2015-12-18',146,8,20,4,'Ok','US','','',0,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 20:39:23',202),(5,'maguinaga1','Aguinaga','','Maria0','','','79','','','','','',12,10,'2504 NW 33rd St','','Oklahoma City','OK','73112','',0,'2011-08-13','2015-04-29',146,0,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(6,'aalston','Alston','','Alivia','','','','','','','','',11,4,'6604 W. Edenborough Dr','Apt 210','Oklahoma City','OK','73132','',0,'2014-02-24','2014-03-09',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(7,'marchan','Archan','','Maria','','','','','','','','',7,2,'2309 SW 35th St','','Oklahoma City','OK','73119','',0,'2010-07-23','2014-03-11',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(8,'yarchan','Archan','','Yiovana','','','','','','','','',12,21,'1115 SW 31st','','Oklahoma City','OK','73109','',0,'2011-12-20','2014-03-24',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(9,'barriola','Arriola','','Brandi','','','832','','','','','',9,23,'6804 Lyrewood Ln','Apt 95','Oklahoma City','OK','73132','',45,'2014-08-13','2017-10-26',49,8,20,4,'OK','US','','',1,1,3,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2018-01-17 22:14:18',202),(10,'yavalos','Avalos','','Yuridia','','','783','','','','','',5,31,'5004 S Eastern','#402','Oklahoma City','OK','73129','',52,'2012-09-29','2016-01-22',146,9,20,4,'OK','US','','',0,0,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-07-25 19:37:41',202),(11,'tavery','Avery','','Tina','','','1005','','','','','',2,24,'6708 Lyrewood Ln','Apt 64','Oklahoma City','OK','73132','',0,'2014-10-15','2015-02-21',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(12,'jbarker','Barker Jr','','Jobe','','','','','','','','',5,29,'7168 Lyrewood Ln','Apt 275','Oklahoma City','OK','73132','',0,'2014-06-01','2014-09-02',0,0,20,0,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(13,'cbecerra','Becerra','','Cristina','','','604','','','','','',7,6,'7134 Lyrewood Ln','','Oklahoma City','OK','73112','',48,'2010-05-19','0000-00-00',146,9,20,4,'OK','USA','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 20:48:30',202),(14,'mbecerra','Becerra','','Maria','','','624','','','','','',9,9,'6904 Lyrewood Lane','Apt 157','Oklahoma City','OK','73132','',50,'2010-07-25','0000-00-00',146,9,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 20:49:03',202),(15,'ebetancort','Betancort','','Erika','','','','','','','','',4,17,'1928 NW 10th St','','Oklahoma City','OK','73106','',0,'2014-04-30','2014-05-08',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(16,'lbird','Bird','','Lindsey','','','','','','','','',10,31,'1317 Kenilworth Rd','','Oklahoma City','OK','73120','',0,'2014-10-14','2014-10-27',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(17,'sbradley','Bradley','','Seth','','','803','','','','','',2,16,'7172 Lyrewood Lane','Apt 286','Oklahoma City','OK','73132','',39,'2014-05-05','0000-00-00',56,7,20,4,'OK','US','bbb','',1,1,3,3,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:14:30',202),(18,'jbrazle','Brazle','','John','','','609','','','','','',12,29,'6912 Lyrewood Ln','Apt 176','Oklahoma City','OK','73127','',0,'2010-05-26','2015-05-26',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(19,'tbrogan','Brogan','','Terence','Terry','','786','323.512.0111 X308','323.512.0105','111-867-5309','tbrogan@accordinterests.com','',5,23,'1805 Forest Hill Dr','','Austin','TX','78745','',8,'2014-03-10','0000-00-00',198,1,4,5,'TX','US','','',0,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',211),(20,'pbuz','Buz','','Patricia','','','700','','','','','',2,14,'7400 NW 7th','','Oklahoma City','OK','73127','',52,'2012-09-12','0000-00-00',146,9,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:39:24',202),(21,'icardona','Cardona','','Isaac','','','434','','','','Isaac@l-objet.com','',5,26,'12608 Panorama Dr','','Burleson','TX','76028','',75,'2006-08-14','0000-00-00',199,17,44,1,'TX','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',202),(22,'icarrington','Carrington','','Imania','','','1067','','','','','',5,18,'7008 Lyrewood Lane','Apt 205','Oklahoma City','OK','73132','',54,'2015-08-14','2015-09-30',82,9,20,4,'OK','US','','',0,0,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(23,'acastillo','Castillo','','Anali','','','118','','','','','',12,9,'7014 NW 74TH','','BETHANY','OK','73008','',52,'2012-07-02','0000-00-00',146,9,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:39:40',202),(24,'lchadwick','Chadwick','','Lisa','','','782','405.721.2194 x1404','','405.596.5346','lchadwick@myisolabella.com','',4,5,'6608 Lyrewood','Apt 019','Oklahoma City','OK','73132','',20,'0000-00-00','2016-03-09',207,5,20,4,'OK','US','','',0,1,1,1,0,'0000-00-00','0000-00-00','36ddbdef8c37d5d01500914e3f83f5318549f6ed8d7d585be593c3474f35ebe78c98f871af6ca569badb9cebc1105676019c8ca1988ac89b1ebf7d4711dbfa20',4,'2016-07-25 19:41:54',202),(25,'kchandler','Chandler','','Kevin','','','1049','','','','kchandler5@uco.edu','',10,24,'930 S. Blvd #108','','Edmond','OK','73034','',0,'2015-07-08','2015-07-27',0,0,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(26,'ccompton','Compton','','Christopher','','','1040','','','','','chriscompton0140@gmail.com',5,4,'1111 N St Charles #8','','Oklahoma City','OK','73127','',54,'2015-06-03','0000-00-00',82,10,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:17:00',202),(27,'jcoury','Coury','','Joseph','','','1056','512.600.1880 x311','323.512.0105','610.574.1851','jcoury@accordinterests.com','couryjj@gmail.com',2,22,'5417 Cypress Ranch Blvd','','Spicewood','TX','78669','',16,'2015-07-29','0000-00-00',200,1,4,5,'TX','US','','',1,1,1,1,2,'0000-00-00','0000-00-00','e0b2aaa51ec0dc6c36caab1d7ea0099da17b92c6e969c28916df5f5785f9e084c2d12b32e953edd12a9af2b39bd1e7e649ef0caed0c82c020c84dc4e1d27568c',4,'2015-12-24 15:57:09',27),(28,'jdavis','Davis','','Jeremy','','','','','','','','',7,7,'9216 NE 46 Street','','Spencer','OK','73089','',0,'2013-10-30','2014-07-02',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(29,'sdavis','Davis','','Sonia','','','277','','','','Sonia@accordinterests.com','',2,14,'5202 Eagle Heights Dr','','Dallas','TX','75212','',5,'2004-09-14','0000-00-00',201,1,42,2,'TX','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','307e7f6b56d18894bcac22b7ca964cc2698cbbc00f0042602b7e0d265a66b2964ab740df2c0a8259c47093fc9e2015e6d05de5fa034ca591d2ee3ee6025b8fa9',4,'2016-02-02 21:04:48',0),(30,'adenson','Denson','','Ashley','','','720','405.721.2194 x 207','','405.627.2072','adenson@myisolabella.com','',5,15,'209 Tonhawa St','','Norman','OK','73069','',5,'2013-06-17','0000-00-00',207,2,20,4,'OK','US','Michele Bright','405.826.2384',1,1,1,1,0,'0000-00-00','0000-00-00','7dd0850607f60b8917b55a661a53ecdaf246a0271757824d8502912170729796a3ac399c5ffdd38520d351f45fd44021ddb1ae6136ffe5274d39bfa326b8c56b',4,'2016-02-02 20:48:01',30),(31,'bdixon','Dixon','','Billy','','','1048','','','','broyce86@hotmail.com','',9,28,'12401 N MacArthur #2513','','Oklahoma City','OK','73142','',0,'2015-07-01','2015-07-09',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(32,'ddoss','Doss','','Derrick','','','515','','','','','',5,4,'6321 West lane','','Oklahoma City','OK','73142','',0,'2009-03-30','2015-06-06',0,0,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(33,'jedgmon','Edgmon','','Jessica','','','','','','','','',2,21,'6604 Lyrewood Ln','Apt 14','Oklahoma City','OK','73132','',0,'2013-08-05','2014-02-27',0,0,20,0,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(34,'dedwards','Edwards','','David','Dave','','702','','','','maintenance@myisolabella.com','',4,22,'6606 Edenborough Dr  #115','','Oklahoma City','OK','73132','',58,'2012-09-12','0000-00-00',197,11,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','672f651261f1ed9785cbbee8dbc97a28e5f52d7d372eff82ab8c9ac0fd445e091276d04b8edc99f1339d35ec4c8c262373bbd5d1041df655ceddc7ed8ca07e5f',4,'2016-02-02 21:06:38',202),(35,'jedwards','Edwards','','Jacob','','','1050','','','','EJake6969@gmail.com','',1,28,'6606 Edenbourgh Dr','Apt 115','Oklahoma City','OK','73132','',0,'2015-06-29','2015-08-14',82,0,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(36,'zedwards','Edwards','','Zachary','','','','','','','','',5,10,'1020 Walsh Lane','','Yukon','OK','73099','',19,'2014-04-30','2014-10-18',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(37,'mespinoza','Espinoza','','Maria','','','584','','','','','',2,11,'6900 Lyrewood Land','Apt 47','Oklahoma City','OK','73132','',37,'2010-01-25','0000-00-00',73,7,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:14:13',202),(38,'mestella','Estella','','Milagros','Mia','','511','972.986.9575','','','mia@l-objet.com','',11,19,'1203 David Drive','','Euless','TX','76040','',15,'2006-05-01','0000-00-00',199,2,44,1,'TX','OK','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:57:39',202),(39,'aestrada','Estrada','','Ana','','','','','','','','',7,8,'5807 NW 36th','','Warr Acres','OK','73122','',0,'2010-09-25','2014-07-02',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(40,'jestrada','Estrada','','Jonathan','Josh','','1018','','','','','',7,15,'3000 NW 15th Street','','Oklahoma City','OK','73107','',17,'2015-01-26','0000-00-00',206,3,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-03-07 19:34:47',202),(41,'mestrada','Estrada','','Maria Consuelo','','','93','','','','','',11,4,'2139 NW 12th Street','','Oklahoma City','OK','73107','',51,'2011-12-24','0000-00-00',146,9,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:40:00',202),(42,'hflores','Flores','','Hector','','','1016','','','','','',3,19,'6159 NW 63rd','Apt B','Oklahoma City','OK','73132','',37,'2015-01-21','0000-00-00',73,7,20,4,'','','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:14:55',202),(43,'sfrazier','Frazier','','Sleece','','','','','','','','',11,17,'714 NE 17th','','Oklahoma City','OK','73105','',0,'2014-08-08','2014-09-19',0,0,20,0,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(44,'cfulbright','Fulbright','','Casey','','','','','','','','',10,26,'6416 Peniel','Apt 31','Oklahoma City','OK','73132','',0,'2014-01-06','2014-02-02',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(45,'cgarcia','Garcia-Torres','','Carmela','','','701','','','','','',3,20,'6217 NW 63rd #F','','Oklahoma City','OK','73132','',52,'2012-09-09','0000-00-00',146,9,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:40:11',202),(46,'jgarcia','Garcia','','Jose','','','303','','','','','',8,26,'214 W Union Bower','Apt #117','Irving','TX','75061','',73,'2006-12-11','0000-00-00',21,17,44,1,'TX','OK','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',202),(47,'jgarcia1','Garcia','O','Jose','','','588','','','','','',5,30,'1502 Carl Road','#116','Irving','TX','75061','',73,'0000-00-00','0000-00-00',21,17,44,1,'TX','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-02-17 21:53:18',200),(48,'agarrett','Garrett','','Alex','','','738','','','','','',11,19,'6714 Elk Canyon Rd','','Oklahoma City','OK','73162','',44,'2013-09-16','2015-06-08',49,7,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(49,'sgibbs','Gibbs','','Sheila','','','687','','','','sgibbs@myisolabella.com','',11,21,'12612 Bannockburn Pl','','Oklahoma City','OK','73142','',43,'2009-09-29','0000-00-00',207,8,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','4eb556a086eef507f5d2e78bf85a9a3a5d4a4e35ead718c2039ff3a3cbf011082f804398639659ea64ef2450476580767b6bdb797ad2ba7d4a912872780a3773',4,'2016-02-02 21:07:43',202),(50,'sgibbs1','Gibbs','','Sydni','','','1036','','','','','sydgibbs888@gmail.com',9,22,'12612 Bannockburn Pl','','Oklahoma City','OK','73142','',42,'2015-06-01','0000-00-00',49,8,20,4,'OK','US','','',0,1,3,3,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 20:49:28',202),(51,'dgilbert','Gilbert','','Derek','','','764','','','','','derekcallengilbert@gmail.com',4,27,'PO Box 355','','Edmond','OK','73083','',2,'2015-07-23','2015-11-25',30,2,20,4,'','','','',0,1,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 20:51:09',202),(52,'mgonzalez','Gonzalez','','Maria','','','1070','','','','','',2,11,'5913 NW 61st','','Warr Acres','OK','73122','',52,'2015-08-23','0000-00-00',146,9,20,4,'OK','US','','',1,1,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:40:24',202),(53,'bgraham','Graham','','Brittney','','','821','405-721-2191','','405-658-4922','bgraham@myisolabella.com','',3,19,'6608 Lyrewood Ln','Apt 24','Oklahoma City','OK','73132','',5,'0000-00-00','0000-00-00',207,1,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','6692e4db39b1a8f595cad9731c4ad8ae2ded96ba42825c8e4260c6e5414709fbdfb518e233bc10c1795b2ac3091c3a1e7d9bf143e1038e2c277f1a97bf373c98',4,'2018-01-12 21:44:42',53),(54,'jgriffin','Griffin-Moore','','Juan','','','1017','','','','jgriffinmoore@gmail.com','',5,12,'5560 Willow Cliff Rd','','Oklahoma City','OK','73122','',0,'2015-01-19','2015-01-21',73,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(55,'lguerra','Guerra','','Lorena','','','629','972.986.9575','','','lorena@l-objet.com','',9,30,'5736 Coventry Park Dr','','Haltom City','TX','76117','',2,'2010-09-13','0000-00-00',199,1,44,1,'TX','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:58:58',202),(56,'dhall','Hall','','Donal','','','717','','','','','',1,28,'6812 Lyrewood Lane #108','','Oklahoma City','OK','73132','',38,'2013-06-06','0000-00-00',207,7,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:15:12',202),(57,'thalliburton','Halliburton','','Thomas','','','799','','','','','',4,15,'6428 W Peniel Ave','','Oklahoma City','OK','73132','',23,'2014-04-16','2015-09-29',59,3,20,4,'OK','US','','',0,0,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(58,'rhammer','Hammer','','Ronald','','','760','','','','','',10,25,'1106 W Griggs Way','','Mustang','OK','73064','',56,'2013-05-22','2015-09-24',197,10,20,4,'OK','US','','',0,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(59,'mhankins','Hankins','','Michael','','','726','','','','mhankins@myisolabella.com','',9,18,'6600 Edenborough Dr','#201','Oklahoma City','OK','73132','',24,'2013-07-09','2015-11-25',207,3,20,4,'OK','US','','',0,1,2,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 20:48:50',202),(60,'bhanson','Hanson','','Brandy','','','','','','','','',4,4,'6321 NW Irwin','','Lawton','OK','73505','',0,'2014-02-25','2014-04-20',0,0,20,0,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(61,'jharrison','Harrison','','Johnny','','','1062','','','','','',3,1,'1111 N St Charles #19','','Oklahoma City','OK','73128','',54,'2015-07-30','2015-08-15',82,9,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(62,'mharshaw','Harshaw','','Michael','','','1041','','','','','',7,31,'7016 Lyrewood Ln Apt 218','','Oklahoma City','OK','73162','',54,'2015-06-12','2015-08-15',82,9,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(63,'shart','Hart','','Stephanie','','','1002','','','','stephaniekh@gmail.com','',8,21,'5812 Abilene Trail','','Austin','TX','78749','',13,'2014-10-21','2015-03-11',0,1,20,5,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(64,'wheard','Heard','','William','','','','','','','','',12,23,'2508 N Lee','','Oklahoma City','OK','73103','',0,'2014-05-20','2014-07-13',0,0,20,0,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(65,'ahenderson','Henderson','','Anthony','','','1025','','','','','',3,21,'3811 Cimarron Estate','','Oklahoma City','OK','73121','',0,'2015-03-26','2015-04-23',59,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(66,'mhenderson','Henderson','','Mar\'Ques','','','1022','','','','','',2,2,'6700 Lyrewood Lane #42','','Oklahoma City','OK','73132','',0,'2015-03-08','2015-04-23',59,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(67,'whenderson','Henderson','','Wayne','','','','','','','','',9,20,'6628 Lyrewood Ln','Apt 140','Oklahoma City','OK','73132','',0,'2013-01-01','2014-01-17',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(68,'dhernandez','Hernandez','','Daniela','','','','','','','','',2,15,'1613 1/2 SW 27th','','Oklahoma City','OK','73108','',0,'2013-08-27','2014-01-31',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(69,'ehernandez','Hernandez','','Eulalia','','','51','','','','','',10,9,'6725 Lancaster Cir','','Oklahoma City','OK','73132','',52,'2010-12-04','0000-00-00',146,9,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:50:48',202),(70,'jhernandez','Hernandez','','Jose','','','1013','','','','','',8,19,'7150 Lyrewood Ln','','Oklahoma City','OK','73132','',60,'2015-01-01','0000-00-00',206,3,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:54:26',202),(71,'thernandez','Hernandez','','Teresa','','','319','','','','','',5,15,'1300 Katy Dr','#217','Irving','TX','75061','',73,'2007-07-02','0000-00-00',21,17,44,1,'TX','US','','',1,1,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',202),(72,'yhernandez','Hernandez','','Yolanda','','','597','','','','','',9,11,'1236 N Britain Rd','Apt 134','Irving','TX','75061','',71,'2010-03-18','0000-00-00',21,17,44,1,'TX','US','','',1,1,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',202),(73,'dhibbard','Hibbard','','Daren','','','689','','','','dhibbard@myisolabella.com','',7,10,'4921 north Woodward Avenue','','Oklahoma City','OK','73112','',36,'2011-06-03','0000-00-00',207,7,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','2e3d8e6d5adf86be7c79d0b1727ad062208f2bbb8f337b68b1e1278204dd8217f99f6e81003cc52eeaf88fa1e670c71fb1ef7ce0917cd0265c5653ee5e7b26f3',4,'2016-03-28 17:05:38',73),(74,'khirouji','Hirouji','','Kyle','','','1024','','','','khiro20@utexas.edu','',12,12,'620 W 24th St #306','','Austin','TX','78705','',13,'2015-03-30','2015-08-28',0,1,4,5,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(75,'jhopkins','Hopkins','','Jaqueline','','','','','','','jaqueline.hopkins@yahoo.com','',4,22,'6442 W Wilshire','Apt D','Oklahoma City','OK','73132','',0,'2014-09-23','2014-11-11',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(76,'mhouse','House','','Mariah','','','1010','','','','','imariahhouse@gmail.com',8,8,'7000 Lyrewood Lane #191','','Oklahoma City','OK','73132','',19,'2014-12-22','0000-00-00',53,5,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-07-25 19:41:12',202),(77,'ohudson','Hudson','','Otha','','','1037','','','','','',1,15,'3308 SE 57th','','Oklahoma City','OK','73135','',23,'2015-06-01','0000-00-00',89,4,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:35:59',202),(78,'kirish','Irish','','Keith','','','','','','','','',9,30,'7100 Lyrewood','Apt 231','Oklahoma City','OK','73132','',0,'2014-01-24','2014-08-14',0,0,20,0,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(79,'jjacobs','Jacobs','','Joshua','','','','','','','','',1,26,'3002 NW 41st Street','','Oklahoma City','OK','73112','',0,'2014-04-28','2014-08-29',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(80,'sjohnson','Johnson','','Sharrell','','','','','','','','',9,29,'5552 Willow Cliff Rd','','Oklahoma City','OK','73122','',0,'2014-02-04','2014-02-13',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(81,'sjones','Jones','','Samuel','','','789','','','','','',10,20,'6708 Lyrewood Ln','Apt 63','Oklahoma City','OK','73132','',56,'2012-09-24','0000-00-00',197,11,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:37:56',202),(82,'ejulian','Julian','','Eric','','','1015','','','','grounds@myisolabella.com','julianeric@gmail.com',12,18,'209 E. Tonhawa','','Norman','OK','73070','',55,'2015-04-30','0000-00-00',207,10,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','84b925336a7c6f3b05710e74360265204a98b17fd5934f6d5664c259c65e6b0ab23feeef4644af825c0c50c58428bb72dc8e413715108a0d90313d6d356f9393',4,'2016-02-02 21:08:49',202),(83,'dkennedy','Kennedy','','Deundra','','','1045','','','','','DeundraKennedy7@gmail.com',10,9,'6026 S May Ave Apt 436','','Oklahoma City','OK','73159','',54,'2015-06-23','0000-00-00',82,10,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:17:13',202),(84,'rkissick','Kissick','','Rachelle','','','','','','','','',7,20,'1204 Templet Dr','','Oklahoma City','OK','73127','',0,'2014-04-29','2014-06-14',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(85,'kkoon','Koon','','Krista','Kristy','','155','405.721.2194','','405.537.6560 ','kkoon@myisolabella.com','',4,25,'10407 SE 23rd','','Midwest City','OK','73130','',29,'2010-11-17','0000-00-00',207,14,20,4,'OK','US','Jeff Koon ','405.365.0704',1,1,2,2,2,'0000-00-00','0000-00-00','c79e37e75de6b6cd9be8d214529ab2c9e46b688a4236ff1ee7f6ace473a4b4faab15639052a61ea37649eef06cc25af8b50c78d6fb4fa4db6416b48d91c8421b',4,'2016-02-03 01:59:50',85),(86,'slacue','Lacue','','Samuel','','','','','','','','',5,27,'5200 Randle Dr','','Spencer','OK','73084','',0,'2013-12-14','2014-05-13',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(87,'bleach','Leach','','Brittani','','','','','','','','',8,26,'2001 S MacArthur Blvd','Apt 95','Oklahoma City','OK','73128','',0,'2014-05-08','2014-09-20',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(88,'mledesma','Ledesma','','Monica','','','745','972.986.9575','','','','',10,7,'402 Skyline Rd','','Grand Prairie','TX','75050','',62,'2013-10-14','0000-00-00',199,5,44,1,'TX','US','','',1,1,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:59:16',202),(89,'mlee','Lee','','Marcus','','','756','','','4057270462','courtesypatrol@myisolabella.com','',10,31,'6824 Lyrewood Lane','Apt 132','Oklahoma City','OK','73132','',24,'0000-00-00','0000-00-00',207,4,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','ae53c381748a7e6e8d7766815d7a2008deb5b4536e9c344aaf98f8e2f3cdc47088c767c1a4e1863f9b5c452fb59c4bcc315a446cefc8a1158284e206208107d0',4,'2016-02-17 21:48:18',200),(90,'tleverette','Leverette','','Tristan','','','','','','','','',2,2,'9020 NW 86th St','','Yukon','OK','73099','',0,'2014-03-26','2014-09-20',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(91,'nlong','Long','','Nathan','','','775','','','','','',12,30,'6600 Lyrewood Lane #6','','Oklahoma City','OK','73132','',54,'2014-02-21','2015-06-12',0,9,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(92,'plong','Long','','Patrick','','','531','','','','','',3,6,'1812 Oxford Way','','Oklahoma City','OK','73120','',0,'2009-06-17','2015-02-23',0,0,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(93,'nlopez','Lopez Jr','','Nicolas','','','1042','','','','','nlbroberg@yahoo.com',12,23,'6416 N Peniel Ave Apt 32','','Oklahoma City','OK','73132','',42,'2015-06-15','2015-10-20',49,7,20,4,'OK','US','','',0,1,3,3,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 20:55:45',202),(94,'mlopez','Lopez','','Monica','','','116','','','','','',11,4,'4903 N Willow Ave','','Bethany','OK','73008','',52,'2012-06-01','0000-00-00',146,9,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:51:00',202),(95,'nlopez1','Lopez','','Nicolas','','','','','','','','',12,23,'6416 S. Lindsay Ave.','','Oklahoma City','OK','73149','',0,'2013-11-12','2014-05-15',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(96,'bmanning','Manning','','Brenda','','','602','','','','','',3,4,'6604 Lyrewood Lane','','Oklahoma City','OK','73132','',79,'2010-04-26','0000-00-00',146,9,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:51:13',202),(97,'amansour','Mansour','','Anna','','','820','','','','','',11,21,'424 Brandon Way','','Austin','TX','78738','',4,'2014-06-30','0000-00-00',19,1,4,5,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(98,'emansour','Mansour','','Elizabeth','','','1046','','','','liz1mansour@gmail.com','',1,11,'424 Brandon Way','','Austin','TX','78738','',4,'2015-06-25','0000-00-00',19,1,4,5,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(99,'smanzano','Manzano','','Stephanie','','','1053','','','','','manzanitasister@yahoo.com',8,11,'6800 Lyrewood Ln #84','','Oklahoma City','OK','73132','',67,'2015-07-13','0000-00-00',53,15,20,4,'OK','US','Michelle Venegas','405.520.8445',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:12:27',202),(100,'lmartinez','Martinez','','Liliana','','','316','','','','','',2,1,'1300 Katy','#227','Irving','TX','75061','',73,'2007-06-18','0000-00-00',21,17,44,1,'TX','US','','',1,1,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',202),(101,'rmartinez','Martinez','','Rosa','','','792','','','','','',9,20,'4909 N Grove Ave','','Warr Acres','OK','73122','',50,'2014-03-31','0000-00-00',146,9,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:51:36',202),(102,'smartinez','Martinez','','Sandra','','','830','','','','','',10,5,'121 SW 42nd','','Oklahoma City','OK','73109','',0,'2014-08-16','2015-07-03',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(103,'pmathews','Mathews','','Paula','','','795','323.512.0111 X310','323.51.20105','','pmathews@accordinterests.com','',4,12,'11400 W Parmer Ln','House # 29','Austin','TX','78613','',2,'2014-04-07','0000-00-00',200,1,4,5,'TX','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','4c63672675dfce5eaefd169e9dc860af0fc894d9b8d00a5dc5d57c1513dbb0cdfbf27904a512847d884994acf22f803317670848b3312b9778116a4886de1752',4,'2018-01-17 18:54:10',103),(104,'dmcclellan','McClellan','','Diamonte','','','1030','','','','dmcclellan93@gmail.com','',10,26,'13409 Pinehurst Rd','','Oklahoma City','OK','73120','',54,'2015-04-27','2015-05-25',153,9,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(105,'lmccoy','McCoy','','Laura','','','110','','','','','',6,15,'6003 Graham Lane','','Yukon','OK','73099','',44,'2013-10-10','0000-00-00',49,8,20,4,'OK','US','','',1,1,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 20:56:33',202),(106,'nmcmahon','McMahon','','Nathalie','','','1047','','','','natmcmahon99@gmail.com','',9,20,'2716 Barton Creek Blvd','Apt 2121','Austin','TX','78735','',4,'2015-06-25','0000-00-00',19,1,4,5,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(107,'cmighton','Mighton','','Charles','','','','','','','','',1,2,'2708 Tudor Rd','','Oklahoma City','OK','73127','',0,'2012-02-16','2014-06-14',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(108,'dmiller','Miller','','Desmond','','','','','','','','',8,20,'1434 NW 25th','','Oklahoma City','OK','73106','',0,'2014-02-06','2014-02-07',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(109,'vmiranda','Miranda','','Victor','','','100','','','','','',11,21,'3100 SW 22nd','','Oklahoma City','OK','73108','',60,'2012-02-13','0000-00-00',206,3,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:54:34',202),(110,'nmoncivaiz','Moncivaiz','','Nora','Nancy','','733','','','','','',3,17,'5202 Eagle Heights Dr.','','Dallas','TX','75212','',73,'2013-08-22','0000-00-00',21,17,44,1,'TX','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',202),(111,'dmonday','Monday','','Darrin','','','839','','','','','dweezy47@gmail.com',6,18,'6708 Lyrewood Lane #64','Apt 046','Oklahoma City','OK','73132','',19,'2014-09-16','0000-00-00',53,5,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-07-25 19:41:31',202),(112,'cmontgomery','Montgomery','','Christopher','','','1023','','','','chris_montgomery373@yahoo.com','',8,31,'6600 Lyrewood Ln Apt 8','','Oklahoma City','OK','73132','',54,'2015-03-17','2015-07-01',0,9,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(113,'amorales','Morales','','Amanda','','','840','','','','','',7,31,'5800 NW 86th Street','','Oklahoma City','OK','73132','',0,'2014-09-15','2015-05-15',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(114,'amoreno','Moreno','','Aurelia','','','56','','','','','',7,16,'1663 Tucasa Drive','Apt 215','Irving','TX','75061','',80,'2011-01-15','0000-00-00',201,11,42,2,'TX','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 22:00:23',202),(115,'lmoreno','Moreno','','Luis','','','','','','','','',12,12,'3021 SW 65th PL','','Oklahoma City','OK','73159','',0,'2013-01-31','2014-01-14',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(116,'rmoreno','Moreno','','Rosa','','','554','','','','','',3,30,'1024 SW 60th St','','Oklahoma City','OK','73139','',49,'2012-01-02','2015-09-16',146,8,20,4,'OK','US','','',0,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(117,'jmorrison','Morrison','','Jessica','','','1021','','','','jessicaxmorr@hotmail.com','',1,14,'5108 NE Haddington','','Lawton','OK','73507','',39,'2015-02-26','2015-08-15',73,6,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(118,'kmorrow','Morrow','','Kimberly','','','','','','','','',3,12,'6700 Lyrewood Lane','Apt 44','Oklahoma City','OK','73132','',0,'2014-03-10','2014-05-15',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(119,'mmyles','Myles','','Michael','','','834','','','','michael.myles15@gmail.com','',8,8,'6612 Lyrewood Lane #030','','Oklahoma city','OK','73132','',0,'2014-08-28','2015-06-12',0,0,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(120,'knguyen','Nguyen','','Kristine','','','1064','','','','','kristine.thet.nguyen@gmail.com',10,7,'1131 NW 30th St','','Oklahoma City','OK','73118','',1,'2015-07-27','0000-00-00',30,2,20,4,'','','','',0,1,3,3,0,'0000-00-00','2015-11-03','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 20:51:44',202),(121,'nnyambe','Nyambe','','Nyambe','','','','','','','','',5,25,'6219 SE Independence','Apt 130','Oklahoma City','OK','73159','',0,'2013-11-11','2014-05-24',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(122,'ro\'connor','O\'Connor','','Richard','','','1051','','','','','OConnor.dick@yahoo.com',6,5,'11701 Vicotria Pl','','Oklahoma City','OK','73120','',56,'2015-07-01','0000-00-00',34,11,20,4,'OK','US','','',1,1,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:37:24',202),(123,'aochoa','Ochoa','','Anthony','','','825','','','','','',1,13,'10005 Glascow Terrace','','Yukon','OK','73099','',0,'2014-08-02','2015-01-18',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(124,'sohanjanianse','Ohanjanianse','','Serjik','','','754','','','','','',5,13,'6616 Lyrewood Lane','Apt 34','Oklahoma City','OK','73132','',37,'2012-05-29','0000-00-00',73,7,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:16:11',202),(125,'holivera','Olivera-Ruelas','','Hermelinda','','','815','','','','','',1,13,'819 SW 47th Street','','Oklahoma City','OK','73109','',47,'2014-06-14','0000-00-00',146,9,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:51:47',202),(126,'mpacheco','Pacheco','','Miguel','','','58','','','','','',1,6,'6712 Lyrewood Lane','Apt 67','Oklahoma City','OK','73132','',60,'2011-03-07','0000-00-00',206,3,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:54:42',202),(127,'apalmer','Palmer','','Anthony','','','739','','','','','',5,15,'3921 SE 48th','','Oklahoma City','OK','73135','',54,'2013-09-17','2015-05-18',82,9,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(128,'jparker','Parker','','Jeremy','','','753','','','','','',10,28,'6712 Lyrewood Lane','Apt 68','Oklahoma City','OK','73132','',54,'2013-12-14','2015-06-22',0,9,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(129,'mparra','Parra','','Maria','','','97','','','','','',10,6,'2329 NW 36th Street','','Oklahoma City','OK','73112','',52,'2012-01-28','0000-00-00',146,9,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:51:57',202),(130,'sparsons','Parsons','','Steven','','','1014','','','','','',6,20,'6900 Lyrewood Ln #146','','Oklahoma City','OK','73132','',23,'2015-01-19','2015-08-04',59,3,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(131,'jpaslay','Paslay','','Jessica','','','1019','','','','jessicanicole30@gmail.com','',2,9,'6816 Lyrewood Lane #115','','Oklahoma City','OK','73132','',42,'2015-02-11','2015-08-28',24,7,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(132,'lperez','Perez','','Lourdes','','','831','','','','','',2,21,'124 SW 42nd','','Oklahoma City','OK','73109','',1,'2014-08-16','2015-08-31',146,8,20,4,'OK','US','','',0,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(133,'vperez','Perez','','Veronica','','','162','','','','','',10,25,'5704 NW 23rd Street','#112','Oklahoma City','OK','73127','',52,'2012-08-07','0000-00-00',146,8,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(134,'gpines','Pines','','Gerald','','','','','','','','',11,10,'6700 Lyrewood Lane','Apt 42','Oklahoma City','OK','73132','',0,'2014-06-16','2014-10-12',0,0,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(135,'dpost','Post','','Dydrell','','','1055','','','','','',11,4,'2513 N Rhode island','','Oklahoma City','OK','73111','',54,'2015-07-27','2015-08-28',82,9,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(136,'bprice','Price','','Brandon','','','1068','','','','','brandondewayneprice@yahoo.com',11,28,'12600 N Mc Arthur Blvd','Apt 719','Oklahoma City','OK','73142','',23,'2015-08-17','0000-00-00',89,4,20,4,'OK','US','','',1,1,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:36:10',202),(137,'rpuckett','Puckett','','Richard','','','','','','','','',10,2,'3121 NW 50th Street','','Oklahoma City','OK','93112','',0,'2014-01-13','2014-02-21',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(138,'lpugh','Pugh','','Lydarrion','','','','','','','','',4,5,'12701 N Pennsylvania',' apt 198','Oklahoma City','OK','73120','',0,'2013-12-14','2014-06-05',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(139,'mramirez','Ramirez','','Maria','','','','','','','','',4,22,'4333 NW 44th St','','Oklahoma City','OK','73106','',0,'2010-12-01','2014-05-06',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(140,'jredmon','Redmon','','John','','','1043','','','','jcredmon357@yahoo.com','',5,9,'12600 N McArthur Blvd Apt 1604','','Oklahoma City','OK','73142','',0,'2015-06-10','2015-06-28',0,0,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(141,'sroberts','Roberts','','Stephen','','','690','972.986.9575','','','steve@l-objet.com','',3,1,'7954 Dusty Way','','Ft Worth','TX','76121','',3,'2011-07-05','0000-00-00',199,1,44,1,'TX','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(142,'mrobison','Robison','','Michael','','','1034','','','','','whitemike201092@gmail.com',2,4,'7172 Lyrewood Ln','Apt 283','Oklahoma City','OK','73132','',56,'2015-05-29','0000-00-00',197,11,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:38:18',202),(143,'aroman','Roman','','Adriana','','','','','','','','',12,19,'2328 N Macarther Blvd','#2120','Irving','TX','75062','',73,'2010-07-21','2014-10-17',0,17,6,1,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',202),(144,'eromero','Romero','','Emma','','','95','','','','','',9,13,'7512 NW 11th','','Oklahoma City','OK','73127','',52,'2011-12-31','0000-00-00',146,9,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:52:28',202),(145,'sromero','Romero','','Susana','','','612','','','','','',4,27,'7042 Lyrewood Ln','','Oklahoma City','OK','73132','',47,'2010-06-05','0000-00-00',146,9,20,4,'OK','US','','',1,1,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:52:38',202),(146,'vroper','Roper','','Venetia','','','394','405-721-2194','','405-885-4468','vroper@myisolabella.com','',6,15,'5401 NW 41st Street','','Oklahoma City','OK','73122','',46,'2007-10-14','0000-00-00',207,9,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','b8f250b4f67f0e425b0e7c20c9d29c34c85e0de79b6de74a5124c42a364168d6971abd191e550767c5c26d86026fca905807f47a010f96f0651f06580de2dd49',4,'2016-02-08 20:14:39',146),(147,'aroundtree','Roundtree','','Alfred','Ace','','688','','','','aroundtree@myisolabella.com','',9,6,'6904 Lyrewood Ln','Apt 155','Oklahoma City','OK','73132','',34,'2010-04-09','0000-00-00',207,6,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','5ee698b3516ade1031acc65411d0ad6ce20c8fc5a9b0cd0924bd8660bfe6e6f205d425aabd2f77486612100c6d4ab23bf3e9d7218b32f0322bbf75847fafa321',4,'2016-02-02 21:06:16',202),(148,'jrouse','Rouse','','James','','','835','','','','','',4,10,'6820 Lyrewood Ln','Apt 128','Oklahoma City','OK','73132','',0,'2014-09-19','2014-11-23',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(149,'arunnels','Runnels','','Amanda','','','798','','','','agary@myisolabella.com','',8,28,'17316 Toledo Dr','','Oklahoma City','OK','73170','',81,'2012-09-24','0000-00-00',207,14,20,4,'OK','US','','',1,1,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:53:57',202),(150,'kryherd','Ryherd','','Kandice','','','1011','','','','xoxokandicejoyce@gmail.com','',10,8,'3108 West Park Place','','Oklahoma City','OK','73107','',0,'2014-12-22','2015-01-02',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(151,'jserrano','Serrano','','Jessica','','','','','','','','',1,25,'5001 NW 10th, Apt. 2702','','Oklahoma City','OK','73127','',0,'2014-05-08','2014-08-26',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(152,'ssibley','Sibley','','Sarah','','','1012','','','','sarahsibley@gmail.com','',6,25,'660 Edenborough Dr #103','','Oklahoma City','OK','73132','',42,'2014-12-22','2015-05-31',49,7,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(153,'dsims','Sims','','Dermond','','','114','','','','','',9,1,'6908 Lyrewood Ln Apt#165','','Oklahoma City','OK','73132','',55,'2012-05-14','2015-05-19',82,9,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(154,'dsmith','Smith','','Dennis','','','752','','','','','',9,6,'4317 SW 22nd St','Apt 907','Oklahoma City','OK','73108','',0,'2013-11-25','2015-01-14',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(155,'jsmith','Smith','','Jason','','','291','','','','','',3,6,'16324 Caney Fork Dr','','Justin','TX','76247','',73,'2006-06-25','2014-12-08',0,17,22,1,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',202),(156,'ssmith','Smith','','Stephanie','','','1039','','','','','stephaniesmith6062@gmail.com',3,16,'6828 Lyrewood Ln #144','','Oklahoma City','OK','73132','',67,'2015-06-12','0000-00-00',53,15,20,4,'OK','US','Dustyn Graham','480.951.0815',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:13:23',202),(157,'ssnow','Snow','','Sean','','','','','','','','',1,22,'808 Old Colony Rd','','Midwest City','OK','73130','',0,'2011-08-22','2014-05-16',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(158,'msolis','Solis','','Maria','','','','','','','','',10,27,'6201 NW 37th','','Bethany','OK','73008','',0,'2012-02-04','2014-06-27',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(159,'csooby','Sooby','','Christina','','','','','','','','',0,0,'','','','','','',0,'0000-00-00','2014-04-25',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(160,'ssouthwell','Southwell','','Shanna','','','1060','','','','','southwellshanna@gmail.com',6,30,'1308 N Independence Ave','','Oklahoma City','OK','73107','',39,'2015-07-27','0000-00-00',56,7,20,4,'OK','US','','',1,1,3,3,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:15:35',202),(161,'gsparks','Sparks','','Gabrielle','','','1059','','','','','gsparks88@gmail.com',9,10,'2505 N. Robinson Apt 3','','Oklahoma City','OK','73103','',44,'2015-08-03','2015-11-20',49,8,20,4,'OK','US','','',0,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 20:57:08',202),(162,'cstates','States','','Clifton','','','755','','','','','',2,15,'1716 SE 51st Street','','Oklahoma City','OK','73129','',54,'2013-12-14','2015-04-13',0,9,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(163,'jsteadman','Steadman','','Jonathan','','','1057','','','','Jsteadman@l-objet.com','jonathan.steadman@gmail.com',5,4,'808 Driggs Ave Apt 6D','','Brooklyn','NY','11211','',28,'2015-08-05','0000-00-00',199,12,44,1,'NY','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:59:40',202),(164,'sstill','Still','','Shay','','','','','','','','',3,3,'700 NE 122nd','Apt 234','Oklahoma City','OK','73023','',0,'2014-02-09','2014-07-01',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(165,'rstrand','Strand','','Rachel','','','1020','','','','','',6,5,'4320 NW 50th Street Apt 115','','Oklahoma City','OK','73112','',39,'2015-02-25','2015-04-17',73,6,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(166,'kstrawn','Strawn','','Kylee','','','1007','','','','strawnkylee@gmail.com','',2,23,'3441 NW 20th Street','','Okalhoma City','OK','73107','',19,'2014-11-17','2014-12-11',24,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(167,'jsullivan','Sullivan','','Jennifer','','','814','','','','jhouse82@gmail.com','',7,4,'4015 N Military Ave','','Oklahoma City','OK','73118','',0,'2014-06-09','2015-04-07',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(168,'xtaylor','Taylor','','Xavier','','','','','','','','',6,13,'6804 Lyrewood Lane','','Oklahoma City','OK','73132','',0,'2013-12-17','2014-04-29',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(169,'tthomas','Thomas','','Tracey','','','1031','','','','','tracey.thomas.3532507@gmail.com',4,17,'6916 Lyrewood Ln #184','','Oklahoma City','OK','73132','',23,'2015-05-12','0000-00-00',89,2,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:36:34',202),(170,'lthompson','Thompson','','Lakota','','','1032','','','','kotahope@gmail.com','',8,3,'2335 NW 30th Street','','Oklahoma City','OK','73112','',4,'2015-05-11','2015-07-15',0,1,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(171,'lthompson1','Thompson','','Laresia','','','1038','','','','','',10,17,'601 Vista Ln TRLR 44','','Edmond','OK','73034','',0,'2015-06-01','2015-08-15',49,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(172,'stierce','Tierce','','Steven','','','1063','','','','','',4,26,'3401 Lightner Ln','','Oklahoma City','OK','73179','',54,'2015-07-30','2015-10-29',82,9,20,4,'OK','US','','',0,0,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(173,'jtinoco','Tinoco','','Jose','','','683','','','','','',2,18,'1611 Tucasa Drive #148','','Irving','TX','75061','',57,'0000-00-00','0000-00-00',201,11,42,2,'TX','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 22:00:33',202),(174,'ctreece','Treece','','Christopher','','','1029','','','','','',2,9,'511 Meadow lake Dr','','Edmond','OK','73003','',0,'2015-04-27','2015-05-29',59,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(175,'rvalenciano','Valenciano','','Robert','','','293','','','','','',12,25,'1661 Darr St','#221','Irving','TX','75061','',22,'2006-06-25','0000-00-00',201,4,42,2,'TX','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 22:00:41',202),(176,'rvalenzuela','Valenzuela','','Ruth','','','590','','','','','',10,13,'12600 N Macarthur Blvd','Apt 1316','Oklahoma City','OK','73142','',52,'2010-02-15','0000-00-00',146,9,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:52:49',202),(177,'bvallejo','Vallejo','','Blanca','','','','','','','','',2,6,'1115 N Purdue','','Oklahoma City','OK','73127','',0,'2014-08-16','2014-09-12',0,0,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(178,'lvazquez','Vazquez','','Luz','','','','','','','','',10,8,'612 SE 14th','','Oklahoma City','OK','73129','',0,'2014-07-13','2014-07-24',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(179,'twalker','Walker','','Tory','','','1033','','','','','',4,20,'6712 Lyrewood Ln Apt 71','','Oklahoma City','OK','73132','',54,'2015-05-21','2015-06-08',82,9,20,4,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(180,'jwalla','Walla','','Jon','','','1009','','','','jwallajr@gmail.com','',1,29,'5725 NW 115th Street','','Oklahoma City','OK','73162','',19,'2014-12-22','2015-01-06',24,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(181,'twaller','Waller','','Timothy','Tim','','1069','','','','','t.waller@gmail.com',11,8,'6616 Lyrewood Lane','Apt 38','Oklahoma City','OK','73132','',56,'2015-08-10','0000-00-00',197,11,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:38:30',202),(182,'mwilde','Wilde','','Madison','','','','','','','','',7,28,'20520 Blackjack Ct','','Luther','OK','73054','',0,'2013-10-11','2014-01-24',0,0,20,0,'','','','',0,1,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(183,'dwilliams','Williams','','Desmond','','','1006','','','','','',6,18,'6912 Lyrewood Ln #174','','Oklahoma City','OK','73132','',0,'2014-10-10','2015-03-25',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(184,'jwilliams','Williams','','Jayson','','','1008','','','','','',4,5,'7000 Lyrewood Lane #190','','Oklahoma City','OK','73132','',0,'2014-11-24','2015-05-08',0,0,20,4,'','','','',0,0,0,0,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(185,'jwilliams1','Williams','','Jesse','','','1058','','','','','',12,24,'6912 Lyrewood Ln #170','','Oklahoma City','OK','73132','',23,'2015-07-29','2015-10-24',89,1,20,4,'OK','US','','',0,0,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2015-12-24 05:12:14',0),(186,'jworley','Worley','','Jennifer','','','759','','','','','',7,7,'6712 Lyrewood Lane','Apt 66','Oklahoma City','OK','73132','',42,'2013-12-14','0000-00-00',49,8,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 20:57:21',202),(187,'azapil','Zapil','','Ana','','','816','','','','','',7,24,'1500 N Drexel','','Oklahoma City','OK','73107','',52,'2014-06-14','0000-00-00',146,9,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:53:06',202),(188,'dzuber','Zuber','','David','','','1052','','','','','zuber_d@yahoo.com',4,26,'7108 Lyrewood Ln Apt 242','','Oklahoma City','OK','73132','',56,'2015-07-02','0000-00-00',34,11,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:37:37',202),(189,'fzuniga','Zuniga','','Fernando','','','766','','','','','',4,8,'6205 NW 63rd St','Apt B','Oklahoma City','OK','73132','',59,'2014-01-14','0000-00-00',197,11,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:38:40',202),(190,'rbegg','Begg','','Ryan','','','1061','','','','','',3,24,'1536 Majors Path','','Southhampton','NY','11968','',76,'2015-08-06','0000-00-00',195,13,44,3,'NY','US','','',1,1,3,3,3,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:58:19',202),(191,'jboucher','Boucher','Carole','Judy','Carole','','654','917.239.4595','212.251.1011','','carole@l-objet.com','',11,24,'502 Park Avenue','Apartment 18D','New York','NY','10022','',29,'2010-01-16','0000-00-00',199,13,44,3,'NY','US','','',1,1,1,1,2,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:56:21',202),(192,'sbrzozowski','Brzozowski','','Susan','','','665','212.251.1011','','','Susan@l-objet.com','',3,1,'31-35 31st St','Apt#503','Astoria','NY','11106','',32,'2012-04-01','0000-00-00',199,12,44,3,'NY','US','','',1,1,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:56:32',202),(193,'mfarrell','Farrell','','Maureen','','','80','516.353.8734','','','maureen@l-objet.com','',9,12,'333 E Broadway','5D','Long Beach','NY','11561','',66,'2011-10-17','0000-00-00',199,13,44,3,'NY','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:57:56',202),(194,'sgates','Gates','','Siena','','','1044','','','','','gatessiena@icloud.com',6,17,'63 Clearview Farm Road','','Southampton','NY','11968','',76,'2015-05-30','0000-00-00',195,13,44,3,'NY','US','','',1,1,3,3,3,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:58:29',202),(195,'jhale','Hale','','JoAnn','','','1026','','','','SHmanager@l-objet.com','JFHale2000@aol.com',9,7,'51 Meeting House Ln','','Southampton','NY','11968','',77,'2015-04-06','0000-00-00',193,13,44,3,'NY','US','','',1,1,2,2,2,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:58:07',202),(196,'ckoenitzer','Koenitzer','','Charlotte','','','1027','','','','','charlottekoenitzer@yahoo.com',8,18,'31 Winding Path Apt 6','','Manorville','NY','11949','',78,'2015-04-06','0000-00-00',193,13,44,3,'NY','OK','','',1,1,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:58:41',202),(197,'dlewis','Lewis','','Deborah','Deb','','413','(405)721-2194 ext 602','','(405) 802-2402','maintenance@myisolabella.com','dalokc@yahoo.com',11,12,'6436 N Peniel','Apt 68','Oklahoma City','OK','73132','',58,'2007-09-14','0000-00-00',207,11,20,4,'OK','US','Jerry Lewis','(405)273-3383',1,1,1,1,0,'0000-00-00','0000-00-00','28048efdfd209f3269674db1ea08fa1856dcb42733eba77a942cc5f2b035c5c9140792eb66618f26563ff4a2ed0e5bc1e7c88f2d774b35f8db848d77e4c1d980',4,'2016-02-08 20:17:15',197),(198,'jmansour','Mansour','','Joseph','Joe','','1004','323.512.0111 X303','3235120105','','jgm@accordinterests.com','',6,19,'424 Brandon Way','','Austin','TX','78733','',30,'2014-10-07','0000-00-00',0,2,4,5,'TX','US','','',1,1,1,1,3,'0000-00-00','0000-00-00','f6ee18c9d99e875c06010cc3dd1db8735ff71ac055eafc5f62b25160b347a767558ced86c573b2e9789e63c4850557348053a1708cacdc4758751a9b04fdd898',1,'2018-01-11 21:15:09',198),(199,'jmun','Mun','','James','','','87','972.986.9575','','','Jmun@l-objet.com','',12,22,'4561 Vista Knoll Dr','','Plano','TX','75093','',7,'2012-01-02','0000-00-00',198,2,44,1,'TX','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:56:07',202),(200,'dnelson','Nelson','','Darla','','','391','323.512.0111 X401','','512.709.4545','djn@accordinterests.com','',3,9,'119 The Hills Drive','','The Hills','TX','78738','',9,'0000-00-00','0000-00-00',198,1,4,5,'TX','US','David C. Miller','512.773.6506',1,1,1,1,2,'0000-00-00','0000-00-00','a0a83a936c3b5cee3edefb35a242d31214cb7879ea716a856f324d43257ab0ae54f7047379b1ff21803c34b288acc747271a862e643bd93a4b96324e98019e6b',7,'2018-01-11 18:29:00',200),(201,'doverbeck','Overbeck','','Dianna','','','686','214-878-0177','','214-878-0177','dlo@accordinterests.com','',1,27,'1711 Forest Band Lane','','Keller','TX','76248','',11,'0000-00-00','0000-00-00',198,2,42,2,'TX','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a4bcf25ebd2103ef5bf4368ddd1d6ffc9cb7c8b90e3f280439e19b5ca33fe31526540b79c84b738e4e39fb9370cbd9cb46a45f883dfec083178c5a1d49fc471d',4,'2016-02-17 21:50:03',200),(202,'sriccobene','Persinger','','Stacey','','','1001','323.512.0111 X309','323.512.0105','906.804.6945','spersinger@accordinterests.com','stacia2103@gmail.com',9,17,'9727 Anderson Village Dr','','Austin','TX','78729','',15,'2014-07-07','0000-00-00',200,1,4,5,'TX','US','Kenneth Persinger','512.577.0958',1,1,1,1,1,'2016-07-01','0000-00-00','afe1627a7a08b774a68380459bce39dea7ff616c193d3407ef04dffa5520095b6429701410eb6daa50d941bdb1b7568f36b8365d88b54256a6caa7e83a4f5253',6,'2017-05-02 17:55:20',202),(203,'tsantaniello','Santaniello','','Thomas','Tom','','773','323.512.0111 X313','323.512.0105','864.415.3703','tas@accordinterests.com','',8,28,'9001 Amberglen Blvd','Apt 12302','Austin','TX','78729','',6,'2014-02-10','0000-00-00',198,2,4,5,'TX','US','','',1,1,1,1,2,'0000-00-00','0000-00-00','29f8c72b5bc695f5cbac35d4ef5d1aafda6f81e99d2e68f30c568cebd4a140fc52e5ba64b3cd50c1f0790b292fa78244448e9b0ab776448b7fccf0b4f38cd4c4',4,'2016-02-02 18:06:16',203),(204,'dshannon','Shannon','','David','','','1054','','','','','',7,29,'6700 Lyrewood Ln #45','','Oklahoma City','OK','73132','',23,'2015-07-15','0000-00-00',89,4,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 20:46:20',202),(205,'rtatum','Tatum','Lee','Ronald','Ron','','663','','','','courtesypatrol@myisolabella.com','',6,8,'7048 Lyrewood Ln','','Oklahoma City','OK','73132','',24,'2010-12-18','0000-00-00',207,4,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-10-04 15:17:24',202),(206,'dwheeler','Wheeler','','Danny','Dan','','596','','','','maintenance@myisolabella.com','',5,15,'8309 NW 140th St','','Oklahoma City','OK','73142','',61,'2010-03-01','0000-00-00',207,3,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','4e53a5827a2da7341091255d6e91e3bfa8208e3779ccfabd626acea85ce1e7c259a7ff786e8381473ba46b38bde64aa750c4fb1cef83c68fa5eaf238e5cf0bd8',4,'2016-02-02 21:07:20',202),(207,'mwheeler','Wheeler','','Melissa','','','474','405.721.2194 x205','','405.812.7028','mwheeler@myisolabella.com','mwheeler1905@gmail.com',10,9,'8309 NW 140th St','','Oklahoma City','OK','73142','',11,'2008-08-26','0000-00-00',198,2,20,4,'OK','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','0848ef77863e706ac9281b91bc725e9473279209c65816a7fa7023dfbfc1b22115c17fd82151334598c26c7bfde35a2b6f985d4d6939c0bf0339bbdceac9fe3d',2,'2018-01-11 09:13:10',207),(208,'jwilliams2','Williams','','Jordan','','','838','212.251.1011','','','Jordan@l-objet.com','',11,29,'419 Hart St','Apt 1B','Brooklyn','NY','11221','',31,'2014-09-22','0000-00-00',199,2,44,3,'NY','US','','',1,1,2,2,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:59:53',202),(209,'eyirach','Yifrach','','Elad','','','450','','','','elad@l-objet.com','',10,3,'1203 Daivd Dr','','Euless','TX','76040','',26,'2008-10-01','0000-00-00',198,12,44,5,'TX','US','','',1,1,1,1,0,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 22:05:04',202),(210,'afeola','Feola','','Andrew','Andy','','1028','323.512.0111 x318','','818.535.7694','afeola@accordinterests.com','',10,31,'1417 Bruce Ave','','Glendale','CA','91202','',30,'2015-04-20','0000-00-00',198,2,4,5,'CA','US','','',1,1,2,1,0,'0000-00-00','0000-00-00','cc87b55169e9871233be4d38c063d9143ceacce502b492afb608845faac9439afec20f159b1bf8a54917cf41af52a1afb57804bc930c8891c8ab75ac161668d3',4,'2016-02-09 22:16:38',210),(211,'smansour','Mansour','F','Steven','Steve','','1065','323-512-0111 X305','','408-921-9957','sman@accordinterests.com','sman@stevemansour.com',10,24,'2215 Wellington Dr','','Milpitas','CA','95035','',10,'2015-08-17','0000-00-00',198,2,4,5,'CA','US','Sharon Yu','408-507-6062',1,1,2,2,2,'0000-00-00','0000-00-00','0848ef77863e706ac9281b91bc725e9473279209c65816a7fa7023dfbfc1b22115c17fd82151334598c26c7bfde35a2b6f985d4d6939c0bf0339bbdceac9fe3d',1,'2018-01-11 01:10:10',211),(212,'cmartin','Martin','','Christin','','','1073','','','','','christyok1117@gmail.com',11,17,'6401 N Asbury Ave','','Oklahoma City','OK','73132','USA',42,'0000-00-00','2015-11-20',49,8,20,5,'OK','USA','Karl Boutwell','405-885-2076',0,1,3,3,3,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 20:56:21',202),(213,'rwalters','Walters','','Richard','','','1074','','','','','',3,29,'5200 N Hales Dr','Apt 214','Oklahoma City','OK','73112','USA',39,'2015-10-26','0000-00-00',56,7,20,4,'OK','US','Jeppe Jakobsen','773-517-6678',1,1,2,3,3,'0000-00-00','0000-00-00','a14949f2168e476f1be751b900e74ec4ffad477c449c3940c9b2ddb25b309dd96a19452b28a5dcf1f3f7f19a11bafbebc51224243d0d7bdcd8f5a80134b2d875',4,'2016-01-21 21:15:49',202),(256,'sde loera','De Loera','','Sergio','','','','','','','','',4,4,'8800 NW 59th Ter','','Bethany','OK','73008','USA',54,'2016-01-12','0000-00-00',82,10,20,4,'OK','USA','Rosana De Loera','405.301.9900',1,1,0,0,0,'0000-00-00','0000-00-00','',4,'2016-01-21 20:43:24',202),(257,'hshelton','Shelton','','Hannah','','','1081','','','','','hannahshelton421@gmail.com',4,21,'6800 N Barr Ave','','Oklahoma City','OK','73132','USA',1,'2015-12-16','0000-00-00',30,2,20,4,'OK','USA','Debbie Shelton','405.642.2756',1,1,0,0,0,'0000-00-00','0000-00-00','',4,'2016-01-21 20:54:22',202),(258,'mfalls','Falls','','Michelle','','','510','405.721.2194 x2014','405.603.4095','405.473.8838','mfalls@myisolabella.com','',3,31,'2009 N Roff','','Oklahoma City','OK','73107','USA',44,'2015-10-07','0000-00-00',49,8,20,12,'OK','USA','','',1,1,1,1,2,'0000-00-00','0000-00-00','50eda3d9ca3e4e6c7d12e9cb3dd8e1cc3ff225fb90eb1612e8a32b665e5de66f06e2c2938af6d0ba64c7837e734312602f370d77d85b25b776b4e1565ab3499e',4,'2017-05-01 22:27:36',211),(259,'flewis','Lewis','','FeNecia','','','1079','','','','','feneciadoolin@gmail.com',3,22,'305 NW 79th','','Oklahoma City','OK','73114','USA',42,'2015-11-30','0000-00-00',49,8,20,4,'OK','USA','Enyo Lewis','405.859.0712',1,1,1,1,2,'0000-00-00','0000-00-00','',4,'2016-01-21 21:02:35',202),(260,'etierce','Tierce','','Erick','','','1075','','','','','',11,28,'6828 Lyrewood Ln Apt 140','','Oklahoma City','OK','73132','USA',54,'2015-11-02','0000-00-00',82,10,20,4,'OK','USA','Anita Tierce','405.306.8610',1,1,1,1,2,'0000-00-00','0000-00-00','',4,'2016-01-21 21:28:45',202),(261,'clambkin','Lambkin','','Cody','','','1076','','','','','',6,17,'7176 Lyrewood Ln #290','','Oklahoma City','OK','73132','USA',54,'2015-11-16','0000-00-00',82,10,20,4,'OK','USA','Lana Lambkin','405.250.0924',1,1,1,1,2,'0000-00-00','0000-00-00','',4,'2016-01-21 21:31:27',202),(262,'mmoore','Moore','','Maleek','','','1080','','','','','',7,22,'2224 Miramar Blvd','','Oklahoma City','OK','73111','USA',23,'2015-12-14','0000-00-00',89,4,20,4,'OK','USA','Ema Moore','918.378.3586',1,1,0,0,0,'0000-00-00','0000-00-00','',4,'2016-01-21 21:35:13',202),(263,'jstonge','St. Onge','','Jeffrey','Jeff','','','646.263.3862','','646.263.3862','jeff@stongecreative.com','',1,1,'1304 Mariposa Dr','#272','Austin','TX','78704','USA',25,'0000-00-00','0000-00-00',0,2,4,12,'','USA','','',1,1,0,0,0,'0000-00-00','0000-00-00','98420b73add4cf6843630356650fc682e8357bb0a458db89dc4bbf8a0fa7910daba55b52dbc3ed9e4aa70b1ff87c53e3e785d85cc9548ff3eb246983c353373e',4,'2016-02-04 17:54:59',263),(264,'drichard','Richard','','Danny','','Mr','','','','415-850-0479','danny@stongecreative.com','',1,1,'','','','','','USA',82,'0000-00-00','0000-00-00',0,1,4,12,'','USA','','',1,1,0,0,0,'0000-00-00','0000-00-00','ccbe0aaa93afb2a86237fc49583d2e849c55308fe952a1437216e6ea08b84c7fc01b380247169c2a6b8b88ff136d4a87eeb4bfdebc5f474a8073b51f793861e9',4,'2016-02-24 19:14:25',211),(265,'jcrow','Crow','','James','William','','1086','323.512.0111 X313','','5124231066','wcrow@accordinterests.com','',10,6,'1507 SUFFOLK DR','','AUSTIN','Te','78723','USA',6,'2016-02-15','0000-00-00',198,2,4,5,'TX','USA','','5124231066',1,1,1,1,2,'0000-00-00','0000-00-00','',4,'2016-02-23 16:08:31',202),(266,'joconnell','OConnell','','Jeffrey','Jeff','','1233','323.512.0111x312','323.512.0105','512.516.2134','joconnell@accordinterests.com','',10,10,'14501 Falconhead Blvd','#12','Austin','TX','78738','USA',6,'2016-03-07','0000-00-00',0,2,4,5,'TX','USA','Stephen O\'Connell','9178827549',1,1,1,1,0,'0000-00-00','0000-00-00','85e777d40a9b989056c907b0b4f8168432a6ca485b0fe5ab7fec016ef0158f4fa839d559f5532bb696a336d0e511258a0ed7502fd7051229a651f0433d2f175b',4,'2016-03-07 20:27:05',266),(267,'sraval','Raval','','Sudip','Sudip','','','','','','sudip@auberginesolutions.com','',2,22,'','','','','','India',25,'0000-00-00','0000-00-00',211,12,4,12,'','India','','',1,0,0,0,0,'0000-00-00','0000-00-00','2c81bf8caa9eb866a8bacfb116df1fdd1146d982b2c7a1f3c28b6b4a41dd1d43c99bc1013b3e73f9ecf52e9ed6eb789b2674df81ab841d3a1e21509bf01628df',4,'2018-01-19 04:51:46',267),(268,'abosamiya','Bosamiya','','Akshay','Akshay','','','','','','akshay@auberginesolutions.com','',1,1,'','','','','','India',25,'0000-00-00','0000-00-00',211,12,4,12,'','India','','',1,1,0,0,0,'0000-00-00','0000-00-00','f2429fccc82a694b05438a71615a5976f72f0437176a3db73d200113b67f1be9d542c1a872235919339e14bf21f00cef737d2320f1640e1b275fb3dfb94eab20',4,'2018-01-11 09:19:32',268),(269,'tester1','Tester','J','William','Billy','','','','','','billyjtester@example.com','',4,1,'123 Elm Street','','Springfield','MO','65619','USA',1,'0000-00-00','0000-00-00',211,12,4,5,'','','','',1,0,0,0,0,'0000-00-00','0000-00-00','782f22c6c71a849f8cdd354140d6a3f68af2d2b753049247540a85833432acb33a22cf227e3448ac1c7989f1d6fa78719a109e1be84f7ae8085d2c9bb63c9632',4,'2018-01-19 05:41:21',211),(270,'tester2','Tester','Q','Sally','Sally','','','','','','sallytester@example.com','',7,17,'1197 Ocello Street','','San Diego','CA','92103','USA',25,'0000-00-00','0000-00-00',211,12,4,12,'','','','',1,0,0,0,0,'0000-00-00','0000-00-00','782f22c6c71a849f8cdd354140d6a3f68af2d2b753049247540a85833432acb33a22cf227e3448ac1c7989f1d6fa78719a109e1be84f7ae8085d2c9bb63c9632',4,'2018-01-19 05:45:22',211);
/*!40000 ALTER TABLE `people` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roles` (
  `RID` mediumint(9) NOT NULL AUTO_INCREMENT,
  `Name` varchar(25) DEFAULT NULL,
  `Descr` varchar(512) DEFAULT NULL,
  PRIMARY KEY (`RID`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'Administrator','This role has permission to do everything'),(2,'Human Resources','This role has full permissions on people, read and print permissions for Companies and Classes.'),(3,'Finance','This role has full permissions on Companies and Classes, read and print permissions on People.'),(4,'Viewer','This role has read-only permissions on everything. Viewers can modify their own information.'),(5,'Tester','This role is for testing'),(6,'OfficeAdministrator','This role is both HR and Finance.'),(7,'OfficeInfoAdministrator','This role is like Office Administrator but also enables delete.');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sessions`
--

DROP TABLE IF EXISTS `sessions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sessions` (
  `UID` bigint(20) NOT NULL,
  `UserName` varchar(40) NOT NULL DEFAULT '',
  `Cookie` varchar(40) NOT NULL DEFAULT '',
  `DtExpire` timestamp NOT NULL DEFAULT '2000-01-01 08:00:00',
  `UserAgent` varchar(256) NOT NULL DEFAULT '',
  `IP` varchar(40) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sessions`
--

LOCK TABLES `sessions` WRITE;
/*!40000 ALTER TABLE `sessions` DISABLE KEYS */;
/*!40000 ALTER TABLE `sessions` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-03-16 15:12:22
