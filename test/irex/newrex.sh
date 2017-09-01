#!/bin/bash
TOP="../.."
BINDIR="${TOP}/tmp/rentroll"

${BINDIR}/rrnewdb

function Rentables() {
	/usr/local/bin/mysql --no-defaults rentroll < tables/r.sql
	/usr/local/bin/mysql --no-defaults rentroll < tables/rtr.sql
	/usr/local/bin/mysql --no-defaults rentroll < tables/rt.sql
	/usr/local/bin/mysql --no-defaults rentroll < tables/rs.sql
	/usr/local/bin/mysql --no-defaults rentroll < tables/rmr.sql
}

function Accounts() {
	/usr/local/bin/mysql --no-defaults rentroll < tables/GLAccount.sql
	/usr/local/bin/mysql --no-defaults rentroll < tables/lm.sql
	/usr/local/bin/mysql --no-defaults rentroll < tables/ar.sql
}

function Transactants() {
	/usr/local/bin/mysql --no-defaults rentroll < tables/tc.sql
	/usr/local/bin/mysql --no-defaults rentroll < tables/payor.sql
	/usr/local/bin/mysql --no-defaults rentroll < tables/user.sql
	/usr/local/bin/mysql --no-defaults rentroll < tables/prospect.sql
}

function Agreements() {
	/usr/local/bin/mysql --no-defaults rentroll < tables/RentalAgreement.sql
	/usr/local/bin/mysql --no-defaults rentroll < tables/RentalAgreementRentables.sql
	/usr/local/bin/mysql --no-defaults rentroll < tables/RentalAgreementPayors.sql
	/usr/local/bin/mysql --no-defaults rentroll < tables/RentableUsers.sql
}

function base() {
	/usr/local/bin/mysql --no-defaults rentroll < tables/biz.sql
	/usr/local/bin/mysql --no-defaults rentroll < tables/pmt.sql
	/usr/local/bin/mysql --no-defaults rentroll < tables/dep.sql
}

base
Accounts
Agreements
Transactants
Rentables
