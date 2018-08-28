#!/bin/bash

if [ "x${ACTIONSBIN}" = "x" ]; then
	ACTIONSBIN="."
else
	echo "ACTIONSBIN was pre-set to:  \"${ACTIONSBIN}\""
fi

execute() {
	cmd="${ACTIONSBIN}/actions -f $1"
	echo "${cmd}"
    ${ACTIONSBIN}/actions -f $1
}

##########################################################################
# outputCheck()
#   Compares the supplied file $1 to gold/$1.gold
#   Specially for this perticular tests
#	Parameters:
# 		$1 = base file name
#		$2 = app options to reproduce
# 		$3 = title
##########################################################################
outputCheck() {
	TESTCOUNT=$((TESTCOUNT + 1))
	printf "PHASE %2s  %3s  %s... " ${TESTCOUNT} $1 $3

	checkPause
	if [ ! -d ${GOLD} ]; then
		echo "${GOLD} directory is missing. Creating it..."
		mkdir ${GOLD}
		echo "replace with known-good output" > ${GOLD}/${1}.gold
	fi
	if [ "${FORCEGOOD}" = "1" ]; then
		goldpath
		cp ${1} ${GOLD}/${1}.gold
		echo "DONE"
	elif [ "${SKIPCOMPARE}" = "0" ]; then
		if [ ! -f ${GOLD}/${1}.gold -o ! -f ${1} ]; then
			echo "Missing file -- Required files for this check: ${1} and ${GOLD}/${1}.gold"
			failmsg
			exit 1
		fi

		#--------------------------------------------------------------------
		# The actual data has timestamp information that changes every run.
		# The timestamp can be filtered out for purposes of testing whether
		# or not the web service could be called and can return the expected
		# data.
		#--------------------------------------------------------------------
		declare -a out_filters=(
			's/(^[ \t]+"LastModTime":).*/$1 TIMESTAMP/'
			's/(^[ \t]+"CreateTS":).*/$1 TIMESTAMP/'
			's/(^[ \t]+"ApplicationReadyDate":).*/$1 TIMESTAMP/'
			's/(^[ \t]+"DecisionDate1":).*/$1 TIMESTAMP/'
			's/(^[ \t]+"DecisionDate2":).*/$1 TIMESTAMP/'
			's/(^[ \t]+"MoveInDate":).*/$1 TIMESTAMP/'
			's/(^[ \t]+"ActiveDate":).*/$1 TIMESTAMP/'
			's/(^[ \t]+"NoticeToMoveReported":).*/$1 TIMESTAMP/'
			's/(^[ \t]+"TerminationDate":).*/$1 TIMESTAMP/'
			's/(^[ \t]+"UserRefNo":).*/$1 USEREFNO/'
			's/(^[ \t]+"AgreementStart":).*/$1 TIMESTAMP/'
			's/(^[ \t]+"AgreementStop":).*/$1 TIMESTAMP/'
			's/(^[ \t]+"RentStart":).*/$1 TIMESTAMP/'
			's/(^[ \t]+"RentStop":).*/$1 TIMESTAMP/'
			's/(^[ \t]+"PossessionStart":).*/$1 TIMESTAMP/'
			's/(^[ \t]+"PossessionStop":).*/$1 TIMESTAMP/'
		)
		cp gold/${1}.gold qqx
		cp ${1} qqy
		for f in "${out_filters[@]}"
		do
			perl -pe "$f" qqx > qqx1; mv qqx1 qqx
			perl -pe "$f" qqy > qqy1; mv qqy1 qqy
		done

		UDIFFS=$(diff qqx qqy | wc -l)
		if [ ${UDIFFS} -eq 0 ]; then
			echo "PASSED"
			passmsg "${3}"
		else
			echo "FAILED:  differences are as follows:" >> ${ERRFILE}
			diff qqx qqy >> ${ERRFILE}
			echo >> ${ERRFILE}
			echo "If the new output is correct:  mv ${1} ${GOLD}/${1}.gold" >> ${ERRFILE}
			cat ${ERRFILE}
			failmsg "${3}"
			if [ "${ASKBEFOREEXIT}" = "1" ]; then
				pause ${1}
			else
				exit 1
			fi
		fi
	else
		echo
	fi
	rm -f qqx qqy
}
##########################################################################
