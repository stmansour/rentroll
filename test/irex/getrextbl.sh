#!/bin/bash

function Rentables() {
	/usr/local/bin/mysqldump --no-defaults rentroll RentableTypes > ./tables/rt.sql
	/usr/local/bin/mysqldump --no-defaults rentroll RentableMarketRate > ./tables/rmr.sql
	/usr/local/bin/mysqldump --no-defaults rentroll Rentable > ./tables/r.sql
	/usr/local/bin/mysqldump --no-defaults rentroll RentableTypeRef > ./tables/rtr.sql
	/usr/local/bin/mysqldump --no-defaults rentroll RentableStatus > ./tables/rs.sql
}

function Accounts() {
	/usr/local/bin/mysqldump --no-defaults rentroll GLAccount > ./tables/GLAccount.sql
	/usr/local/bin/mysqldump --no-defaults rentroll LedgerMarker > ./tables/lm.sql
	/usr/local/bin/mysqldump --no-defaults rentroll AR > ./tables/ar.sql
}

function Transactants() {
	/usr/local/bin/mysqldump --no-defaults rentroll Transactant > ./tables/tc.sql
	/usr/local/bin/mysqldump --no-defaults rentroll Payor > ./tables/payor.sql
	/usr/local/bin/mysqldump --no-defaults rentroll User > ./tables/user.sql
	/usr/local/bin/mysqldump --no-defaults rentroll Prospect > ./tables/prospect.sql
}

function Agreements() {
	/usr/local/bin/mysqldump --no-defaults rentroll RentalAgreement > ./tables/RentalAgreement.sql
	/usr/local/bin/mysqldump --no-defaults rentroll RentalAgreementRentables > ./tables/RentalAgreementRentables.sql
	/usr/local/bin/mysqldump --no-defaults rentroll RentalAgreementPayors > ./tables/RentalAgreementPayors.sql
	/usr/local/bin/mysqldump --no-defaults rentroll RentableUsers > ./tables/RentableUsers.sql
}

function base() {
	/usr/local/bin/mysqldump --no-defaults rentroll Business > ./tables/biz.sql
	/usr/local/bin/mysqldump --no-defaults rentroll PaymentType > ./tables/pmt.sql
	/usr/local/bin/mysqldump --no-defaults rentroll Depository > ./tables/dep.sql
}

base
Agreements
Transactants
Accounts
Rentables