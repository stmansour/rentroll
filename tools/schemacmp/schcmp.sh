#!/bin/bash
MYSQL="mysql"
DBNAME="rentroll"
DEF1="def1.sh"
DEF2="def2.sh"
S1="schema1"
LOCAL="local"
REMOTE="remote"
DBREPORT="report.txt"

##############################################################################################
# getSchema
#	Contact the remote db server, get the table list, then the definition for each table.
#
#	Parameters:
# 		$1 = def shell script name
#		$2 = directory name for table defs
##############################################################################################
getSchema() {
	echo -n "Build Table List for ${2}..."
	if [ ! -f confprod.json ]; then
		/usr/local/accord/bin/getfile.sh accord/db/confprod.json
	fi
	if [ "${2}" = "${LOCAL}" ]; then
		OPTS="--no-defaults"
	else
		dbhost=$(grep "RRDbhost" confprod.json | awk '{print $2}' | sed -e 's/[",]//g')
		dbuser=$(grep "RRDbuser" confprod.json | awk '{print $2}' | sed -e 's/[",]//g')
		OPTS="-h ${dbhost} -u ${dbuser} -P 3306"
	fi
	mysql ${OPTS} ${DBNAME} < cmds >t
	echo "DONE"

cat > ${1} <<FEOF
#!/bin/bash
declare -a out_filters=(
FEOF

grep -v "Tables_in" t | sed -e 's/\(.*\)/	"\1"/' >> ${1}

cat >>${1} <<FEOF
)

echo  "TABLES" > ${2}/TABLES
for f in "\${out_filters[@]}"
do
	echo "\${f}" >> ${2}/TABLES
	echo "DESCRIBE \${f};" > t
	${MYSQL} ${OPTS} ${DBNAME} < t > ${2}/\${f}
done
FEOF
echo "DONE"

echo -n "Build schema definitions for ${2}..."
chmod +x ${1}
./${1}
echo "DONE"

}

#-----------------------------------------
#  INITIALIZE...
#-----------------------------------------
start=$(date)
rm -rf ${LOCAL} ${REMOTE}
mkdir ${LOCAL} ${REMOTE}
echo "show tables;" > cmds

#--------------------------------------------------
#  STEP 1  -- get local and remote schema defs
#--------------------------------------------------
getSchema "${DEF1}" "${LOCAL}"
getSchema "${DEF2}" "${REMOTE}"

#--------------------------------------------------
#  STEP 2  -- compare each table def and report diffs
#--------------------------------------------------
BADSCHCOUNT=0
echo "SCHEMA DIFFS" > ${DBREPORT}
ls -l ${LOCAL} | awk '{print $9}' | while read f; do
	if [ "x${f}" != "x" -a "${f}" != "TABLES" ]; then
	    sort ${LOCAL}/${f} > x ; cp x ${LOCAL}/${f}
	    sort ${REMOTE}/${f} > y ; cp y ${REMOTE}/${f}
	    UDIFFS=$(diff x y | wc -l)
	    if [ ${UDIFFS} -ne 0 ]; then
	        echo "TABLE ${f} has differences:" >> ${DBREPORT}
	        diff x y >> ${DBREPORT}
	        BADSCHCOUNT=$((BADSCHCOUNT + 1))
	        echo "Miscompare on TABLE ${f} = ${BADSCHCOUNT}"
	    fi
	fi
done

missing=$(comm -13 ${REMOTE}/TABLES ${LOCAL}/TABLES | wc -l)
if [ ${missing} -gt 0 ]; then
	echo "Miscompare MISSING TABLES" | tee -a ${DBREPORT}
	echo "The following tables are defined in the schema but are missing in ${db}:" | tee -a ${DBREPORT}
	comm -23 ${REMOTE}/TABLES ${LOCAL}/TABLES | tee -a ${DBREPORT}
	echo | tee -a ${DBREPORT}
fi

extra=$(comm -23 ${REMOTE}/TABLES ${LOCAL}/TABLES | wc -l)
if [ ${extra} -gt 0 ]; then
	echo "Miscompare EXTRA TABLES" | tee -a ${DBREPORT}
	echo "The following tables exist in ${db} but not in defined schema:" | tee -a ${DBREPORT}
	comm -23 ${REMOTE}/TABLES ${LOCAL}/TABLES | tee -a ${DBREPORT}
	echo | tee -a ${DBREPORT}
fi

stop=$(date)
echo "Start time:  ${start}"
echo "Stop time:   ${stop}"
