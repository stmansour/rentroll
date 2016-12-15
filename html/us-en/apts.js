/* rentroll 1.0 (c) Accord Interests, LLC, sman@accordinterests.com */
/************************************************
*  Strings for RentRoll 1.0
*
*  Language:  us-en
*  Industry:  Apartment Rental  (apts)
*
*  This module defines the values for use with
*  the US-English Apartment Rental template.
*
*  Conventions:
*  - All strings begin wth 's' followed by camel-case description
*  - Wherever possible, all camel-cased variable names should be the term
*    used in the Accord Glossary
************************************************/

function plural(s) {
	// todo - if last char is s, plural is almost always + 'es'
	return s + 's'
}

var sTransactant = "Person";
var sRentable = "Unit";
var sUser = "Renter";
var sPayor = "Payor";
var sProspect = "Prospect";
var sRentalAgreement = "Rental Agreement";
var sAssessment ="Charge";
var sReceipt ="Payment";
var sRatePlan ="Rate Plan";
var sService ="Service";
var sBusinessUnit="Business Unit";
var sAssessment="Charge"
var sReceipt="Payment"