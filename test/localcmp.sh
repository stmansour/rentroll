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
usage() {
	cat <<EOF

SYNOPSIS
	$0 [-fs]

	Perform SQL schema compare of the schema defined in the roller source
	tree to the schema files in dbfiles.txt or the file specified in the
	-f option.

OPTIONS
	-f  filename
        Perform schema check on filename, then exit.
	-s  stop the program after encountering a schema mismatch

EXAMPLES
	To check all .sql files in the test directory and subdirectories:

		$ ./localcmp.sh

	To check all .sql files and stop on the first error found in a table
	schema:

		$ ./localcmp.sh -s

	To check a particular database file:

		$ ./localcmp.sh -f filename

EOF
}

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

###############################################################################
# schemacmp
#	Compare ${1} schema to the schema defined in the Roller database schema
#   source directory.
#
#	Parameters:
# 		$1 = filename on which to perform schema compare
###############################################################################
schemacmp() {
	BADSCHCOUNT=0
	missing=0
	extra=0
	echo "------------------------------------------------------------------------" | tee -a ${DBREPORT}

	#--------------------------------------------------------------------------
	# make sure it's a rentroll datebase...
	#--------------------------------------------------------------------------
	local DBNAME=$(head -5 ${1} | grep Database: | awk '{ print $5; }')
	if [ "${DBNAME}" != "rentroll" ]; then
		echo "SKIPPING: ${1}  (db name = ${DBNAME})"
		return 0
	fi

	echo "CHECKING SCHEMA: ${1}" | tee -a ${DBREPORT}
	echo "DROP DATABASE IF EXISTS rentroll; CREATE DATABASE rentroll;" | ${MYSQL}
	${MYSQL} rentroll < ${1}
	getSchema "${DEF2}" "${CHECKDB}"

	#--------------------------------------------------------------------------
	# The loop construction below is important.  If it is done this way:
	#
	#     BADSCHCOUNT=0
	#     ls -l ${REFDB} | awk '{print $9}' | while read f; do
	#         ...
	#         ((BADSCHCOUNT++))
	#         ...
	#     done
	#
	# then ${BADSCHCOUNT} will always be 0 when the loop terminates as it is
	# executed in a subshell.  The construction of the loop below is done so that
	# a subshell is not used, so ${BADSCHCOUNT} retains its value after the loop
	# ends.
	#--------------------------------------------------------------------------
	while read f; do
		if [ "x${f}" != "x" -a "${f}" != "TABLES" ]; then
			if [ -f "${REFDB}/${f}" -a -f "${CHECKDB}/${f}" ]; then
			    sort ${REFDB}/${f} > x ; cp x ${REFDB}/${f}
			    sort ${CHECKDB}/${f} > y ; cp y ${CHECKDB}/${f}
			    UDIFFS=$(diff x y | wc -l)
			    if [ ${UDIFFS} -ne 0 ]; then
			        echo "TABLE ${f} has differences:"  | tee -a ${DBREPORT}
			        diff x y >> ${DBREPORT}
					diff x y
			        ((BADSCHCOUNT++))
			        echo "Miscompare on TABLE ${f} = ${BADSCHCOUNT}" | tee -a ${DBREPORT}
			    fi
		    fi
		fi
	done << EOT
	$(ls -l ${REFDB} | awk '{print $9}')
EOT

	# sort ${CHECKDB}/TABLES > t ; mv t ${CHECKDB}/TABLES
	# sort ${REFDB}/TABLES > t ; mv t ${REFDB}
	missing=$(comm -13 ${CHECKDB}/TABLES ${REFDB}/TABLES | wc -l)
	if [ ${missing} -gt 0 ]; then
		echo "Miscompare MISSING TABLES" | tee -a ${DBREPORT}
		echo "The following tables are defined in the schema but are missing in ${1}:" | tee -a ${DBREPORT}
		comm -13 ${CHECKDB}/TABLES ${REFDB}/TABLES | tee -a ${DBREPORT}
		echo | tee -a ${DBREPORT}
	fi

	extra=$(comm -23 ${CHECKDB}/TABLES ${REFDB}/TABLES | wc -l)
	if [ ${extra} -gt 0 ]; then
		echo "Miscompare EXTRA TABLES" | tee -a ${DBREPORT}
		echo "The following tables exist in ${1} but not in defined schema:" | tee -a ${DBREPORT}
		comm -23 ${CHECKDB}/TABLES ${REFDB}/TABLES | tee -a ${DBREPORT}
		echo | tee -a ${DBREPORT}
	fi

	if [ "${BADSCHCOUNT}" -eq "0" -a ${missing} -eq 0 -a ${extra} -eq 0 ]; then
		echo "no issues found" | tee -a ${DBREPORT}
	else
		if [ ${STOPONERR} -eq "1" ]; then
			echo
			echo "Stopping on error(s). File = ${1}"
			exit 1
		fi
	fi
}

SINGLEFILE=
STOPONERR=0
#==============================================================================
#
# 				SCRIPT BEGINS HERE...
#
#==============================================================================
while getopts "f:F:s" o; do
	echo "o = ${o}"
	case "${o}" in
		f | F)
			SINGLEFILE=${OPTARG}
			;;
		s)	STOPONERR=1
			echo "Option -s selected.  Will stop on error."
			;;
		*) 	usage
			exit 1
			;;
	esac
done
shift $((OPTIND-1))

#==============================================================================
# First see if we need to actually do this comparison.  We'll save a file
# after we complete a comparison to have the date of the last compare. If
# the db schema file is older than our last comparison we don't need to
# do the compare.
#==============================================================================
SCHEMA="../db/schema/schema.sql"
# LASTCHECK="./lastschemacheck"
# NEEDCHECK=0
# if [ ! -f "${LASTCHECK}" ]; then
#     NEEDCHECK=1
#     touch "${LASTCHECK}"
# fi
# if [ NEEDCHECK = "0" -a "${SCHEMA}" -nt "${LASTCHECK}" ]; then
#     NEEDCHECK=1
# fi
#
# if [ ${NEEDCHECK} = "0" ]; then
#     exit 0
# fi

#==============================================================================
# Create dbfiles.txt.  It will contain all the .sql files in the test
# directory and subdirectories, plus any sql files hardcoded below.
#==============================================================================
rm -f "dbqqqmods.sql"   # this is the temp db from dbmod.sql
mv dbfiles.txt dbfiles.old
cat >dbfiles.txt <<EOF
../tools/dbgen/empty.sql
EOF
find . -name "*.sql" >> dbfiles.txt

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

echo "SCHEMA DIFFS" > ${DBREPORT}

if [ "x${SINGLEFILE}" != "x" ]; then
	schemacmp "${SINGLEFILE}"
else
	#==============================================================================
	# Explanation of the loop
	#     IFS=''
	#         (or IFS=) prevents leading/trailing whitespace from being trimmed.
	#     -r
	#         prevents backslash escapes from being interpreted.
	#     || [[ -n ${1} ]]
	#         prevents the last line from being ignored if it doesn't end with
	#         a \n (since  read returns a non-zero exit code when it encounters
	#         EOF).
	#
	# for db in ${dblist[@]}
	#==============================================================================
	while IFS='' read -r db || [[ -n "${db}" ]]; do
		schemacmp "${db}"
	done < dbfiles.txt
fi

echo "------------------------------------------------------------------------" | tee -a ${DBREPORT}
stop=$(date)
echo "Start time:  ${start}" | tee -a ${DBREPORT}
echo "Stop time:   ${stop}" | tee -a ${DBREPORT}

if [ ${BADSCHCOUNT} -gt 0 -o ${missing} -gt 0 -o ${extra} -gt 0 ]; then
	exit 2
fi
exit 0
