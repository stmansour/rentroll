#!/bin/bash
TOP="../.."
BINDIR="${TOP}/tmp/rentroll"

${BINDIR}/rrnewdb

/usr/local/bin/mysql --no-defaults rentroll < tables/biz.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/acct.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/pmt.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/dep.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/rt.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/ar.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/tc.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/payor.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/user.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/prospect.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/rmr.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/r.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/rtr.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/rs.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/lm.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/GLAccount.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/pmt.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/depository.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/RentalAgreement.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/RentalAgreementRentables.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/RentalAgreementPayors.sql
/usr/local/bin/mysql --no-defaults rentroll < tables/RentableUsers.sql
