#!/bin/bash
TOP=..
BINDIR="${TOP}/tmp/rentroll"
MYSQL="mysql --no-defaults"
DBNAME="rentroll"
DEF1="def1.sh"
DEF2="def2.sh"
S1="schema1"
REFDB="refdb"
CHECKDB="checkdb"
DBREPORT="dbreport.txt"

###############################################################################
# Compare the schemas in the current development source tree to the definition
# of the schema found in ../db/schema.  This source code assumes that it is
# being run in the ./test directory.
###############################################################################

###############################################################################
# getSchema
#	Contact the checkdb db server, get the table list, then the definition for
#   each table.
#
#	Parameters:
# 		$1 = def shell script name
#		$2 = directory name for table defs
###############################################################################
getSchema() {
	if [ "${2}" = "${REFDB}" ]; then
		${BINDIR}/rrnewdb > /dev/null
	fi

	rm -rf ${2}
	mkdir -p "${2}"

	${MYSQL} ${DBNAME} < cmds >t
	# echo "DONE"

cat > ${1} <<FEOF
#!/bin/bash
declare -a tables=(
FEOF

grep -v "Tables_in" t | sed -e 's/\(.*\)/	"\1"/' >> ${1}
cat >>${1} <<FEOF
)

echo  "TABLES" > ${2}/TABLES
for f in "\${tables[@]}"
do
	echo "\${f}" >> ${2}/TABLES
	echo "DESCRIBE \${f};" > t
	${MYSQL} ${DBNAME} < t > ${2}/\${f}
done
FEOF
	#------------------------------------------------
	# now execute the script to build the schema...
	#------------------------------------------------
	chmod +x ${1}
	./${1}
}

#-----------------------------------------
#  INITIALIZE...
#-----------------------------------------
start=$(date)
echo "DROP DATABASE IF EXISTS rentroll; CREATE DATABASE rentroll;" | ${MYSQL}
echo "show tables;" > cmds

#--------------------------------------------------
#  STEP 1  -- get refdb schema defs
#--------------------------------------------------
getSchema "${DEF1}" "${REFDB}"

#--------------------------------------------------
#  STEP 2  -- compare each table def and report diffs
#--------------------------------------------------
declare -a dblist=(
    '../tools/dbgen/empty.sql'
    'acctbal/baltest.sql'
    'payorstmt/pstmt.sql'
    'rfix/rcptfixed.sql'
    'rfix/receipts.sql'
    'roller/prodrr.sql'
    'rr/rr.sql'
    'tws/rr.sql'
    'webclient/webclientTest.sql'
    'websvc1/asmtest.sql'
    'websvc3/tasks.sql'
    'workerasm/rr.sql'
)

echo "SCHEMA DIFFS" > ${DBREPORT}
for db in ${dblist[@]}
do
	echo "------------------------------------------------------------------------" | tee -a ${DBREPORT}
	echo "CHECKING SCHEMA: ${db}" | tee -a ${DBREPORT}
	echo "DROP DATABASE IF EXISTS rentroll; CREATE DATABASE rentroll;" | ${MYSQL}
	${MYSQL} rentroll < ${db}
	getSchema "${DEF2}" "${CHECKDB}"
	BADSCHCOUNT=0
	ls -l ${REFDB} | awk '{print $9}' | while read f; do
		if [ "x${f}" != "x" -a "${f}" != "TABLES" ]; then
			if [ -f "${REFDB}/${f}" -a -f "${CHECKDB}/${f}" ]; then
			    sort ${REFDB}/${f} > x ; cp x ${REFDB}/${f}
			    sort ${CHECKDB}/${f} > y ; cp y ${CHECKDB}/${f}
			    UDIFFS=$(diff x y | wc -l)
			    if [ ${UDIFFS} -ne 0 ]; then
			        echo "TABLE ${f} has differences:"  | tee -a ${DBREPORT}
			        diff x y >> ${DBREPORT}
			        BADSCHCOUNT=$((BADSCHCOUNT + 1))
			        echo "Miscompare on TABLE ${f} = ${BADSCHCOUNT}" | tee -a ${DBREPORT}
			    fi
		    fi
		fi
	done

	# sort ${CHECKDB}/TABLES > t ; mv t ${CHECKDB}/TABLES
	# sort ${REFDB}/TABLES > t ; mv t ${REFDB}
	missing=$(comm -13 ${CHECKDB}/TABLES ${REFDB}/TABLES | wc -l)
	if [ ${missing} -gt 0 ]; then
		echo "Miscompare MISSING TABLES" | tee -a ${DBREPORT}
		echo "The following tables are defined in the schema but are missing in ${db}:" | tee -a ${DBREPORT}
		comm -13 ${CHECKDB}/TABLES ${REFDB}/TABLES | tee -a ${DBREPORT}
		echo | tee -a ${DBREPORT}
	fi

	extra=$(comm -23 ${CHECKDB}/TABLES ${REFDB}/TABLES | wc -l)
	if [ ${extra} -gt 0 ]; then
		echo "Miscompare EXTRA TABLES" | tee -a ${DBREPORT}
		echo "The following tables exist in ${db} but not in defined schema:" | tee -a ${DBREPORT}
		comm -23 ${CHECKDB}/TABLES ${REFDB}/TABLES | tee -a ${DBREPORT}
		echo | tee -a ${DBREPORT}
	fi

	if [ ${BADSCHCOUNT} -eq 0 -a ${missing} -eq 0 -a ${extra} -eq 0 ]; then
		echo "no issues found" | tee -a ${DBREPORT}
	fi
done
echo "------------------------------------------------------------------------" | tee -a ${DBREPORT}
stop=$(date)
echo "Start time:  ${start}" | tee -a ${DBREPORT}
echo "Stop time:   ${stop}" | tee -a ${DBREPORT}

if [ ${BADSCHCOUNT} -gt 0 -o ${missing} -gt 0 -o ${extra} -gt 0 ]; then
	exit 2
fi
exit 0
