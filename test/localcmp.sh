#!/bin/bash
TOP=..
BINDIR="${TOP}/tmp/rentroll"
MYSQL="mysql"
DBNAME="rentroll"
DEF1="def1.sh"
DEF2="def2.sh"
S1="schema1"
REFDB="refdb"
CHECKDB="checkdb"

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
		${BINDIR}/rrnewdb
	fi

	rm -rf ${2}
	mkdir -p "${2}"

	mysql --no-defaults ${DBNAME} < cmds >t
	# echo "DONE"

cat > ${1} <<FEOF
#!/bin/bash
declare -a out_filters=(
FEOF

grep -v "Tables_in" t | sed -e 's/\(.*\)/	"\1"/' >> ${1}
cat >>${1} <<FEOF
)

for f in "\${out_filters[@]}"
do
	echo "DESCRIBE \${f};" > t
	${MYSQL} --no-defaults ${DBNAME} < t > ${2}/\${f}
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
echo "show tables;" > cmds

#--------------------------------------------------
#  STEP 1  -- get refdb schema defs
#--------------------------------------------------
getSchema "${DEF1}" "${REFDB}"

#--------------------------------------------------
#  STEP 2  -- compare each table def and report diffs
#--------------------------------------------------
declare -a dblist=(
    # 'acctbal/baltest.sql'
    # 'payorstmt/pstmt.sql'
    # 'rfix/rcptfixed.sql'
    # 'rfix/receipts.sql'
    # 'roller/prodrr.sql'
    # 'rr/rr.sql'
    # 'webclient/webclientTest.sql'
    # 'websvc1/asmtest.sql'
    # 'websvc3/tasks.sql'
    'workerasm/rr.sql'
)

echo "SCHEMA DIFFS" > report.txt
for db in ${dblist[@]}
do
	echo "------------------------------------------------------------------------" | tee -a report.txt
	echo "CHECKING SCHEMA: ${db}" | tee -a report.txt
	echo | tee -a report.txt
	mysql --no-defaults rentroll < ${db}
	getSchema "${DEF2}" "${CHECKDB}"
	BADSCHCOUNT=0
	ls -l ${REFDB} | awk '{print $9}' | while read f; do
		# echo "checking ${f}"
		if [ "x${f}" != "x" ]; then
			if [ -f "${REFDB}/${f}" -a -f "${CHECKDB}/${f}" ]; then
			    sort ${REFDB}/${f} > x ; cp x ${REFDB}/${f}
			    sort ${CHECKDB}/${f} > y ; cp y ${CHECKDB}/${f}
			    UDIFFS=$(diff x y | wc -l)
			    if [ ${UDIFFS} -ne 0 ]; then
			        echo "TABLE ${f} has differences:"  | tee -a report.txt
			        diff x y >> report.txt
			        BADSCHCOUNT=$((BADSCHCOUNT + 1))
			        echo "Miscompare on TABLE ${f} = ${BADSCHCOUNT}" | tee -a report.txt
			    fi
			else
				if [ -f "{REFDB}/${f}" ]; then
					echo "Table \"${f}\" was not found in ${db}" | tee -a report.txt
				else
					echo "Table \"${f}\" was found in ${db} but does not exist in the db created by ${BINDIR}/rrdbnew" | tee -a report.txt
				fi
		    fi
		fi
	done
done
echo "------------------------------------------------------------------------" | tee -a report.txt

stop=$(date)
echo "Start time:  ${start}" | tee -a report.txt
echo "Stop time:   ${stop}" | tee -a report.txt
