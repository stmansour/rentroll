#!/bin/bash
#
# USAGE:
# 	rollback product
#
# SYNOPSIS: remove the currently released product from accord/air/latest
#           and replace it with the last archived version.
#
# DESCRIPTION:
# 1. Move accord/air/latest/prodname* to rollbacks/
# 2. Move accord/air/latest/BOM* to rollbacks/
#	 Update the bom to indicate the product and date and time of the rollback
# 3. Move the latest file in accord/air/archives/prodname*/ to accord/air/latest/
# 4. Create a BOM for accord/air/latest
# 
#------------------------------------------------------------------------------

TS=$(date +%Y%m%d_%H%M%S%Z)
TSFMT=$(date "+%B %d, %Y %l:%M:%S%P %Z")
PRODNAME=${1}
DEBUG=1
OS=$(uname)

if [ "x${PRODNAME}" == "x" ]; then
        echo "You must supply the bundle to release"
        exit 1
fi

function decho {
	if (( ${DEBUG} == 1 )); then
		echo
		echo "${1}"
	fi
}

#############################################################################
# pause()
#   Description:
#		Ask the user how to proceed.
#
#   Params:
#       ${1} = name of the file to move to gold/${1}.gold if 'm' is pressed
#############################################################################
pause() {
	echo
	read -p "Press [Enter] to continue, M to move ${2} to gold/${2}.gold, Q or X to quit..." x
	x=$(echo "${x}" | tr "[:upper:]" "[:lower:]")
	if [ "${x}" == "q" -o ${x} == "x" ]; then
		exit 0
	fi
}

#############################################################################
# readConfig
#   Description:
#   	Read the config.json file from the directory containing this script
#		to set some of the key values needed to access the artifactory repo.
#
#   Params:
#		none
#
#	Returns:
#		sets variables:  APIKEY, USER, and URLBASE
#
#############################################################################
function readConfig {
    APATH=$(cd `dirname "${BASH_SOURCE[0]}"` && pwd)
    CONF="${APATH}/config.json"
    USER=$(grep RepoUser ${CONF} | awk '{print $2;}' | sed 's/\"//g' | sed 's/,//')
    APIKEY=$(grep RepoPass ${CONF} | awk '{print $2;}' | sed 's/\"//g' | sed 's/,//')
    URLBASE=$(grep RepoURL ${CONF} | awk '{print $2;}' | sed 's/\"//g' | sed 's/,//')
}

#############################################################################
# bom
#   Description:
#               Create a bill-of-materials for the current contents of
#				accord/air/release/
#
#
#   Params:
#		none
#
#############################################################################
function bom {
	decho "*** Entering bom()"

    re="^([^_]+)_([^\.]+).*$"
    BOMTMP="bom"
    BOM="bom.txt"
    ReadableDate=$(date "+%B %d, %Y %l:%M:%S%P %Z")

    jfrog rt s "accord/air/release/*" | grep path | awk '{print $2;}' | sed 's/\"//g' > ${BOMTMP}

    #      20180123 _ 18  08  35  PST
    dtre="(........)_(..)(..)(..)(...)"

    echo "AIR Version 1.xxx" > ${BOM}
    echo >> ${BOM}
    echo "Product      Timestamp             Release Date" >> ${BOM}
    echo "---------------------------------------------------------------------" >> ${BOM}

    while read line; do
        filename=$(basename "$line")

        #------------------------------------------------------------
        # parse the filename into a product name and the timestamp
        #------------------------------------------------------------
        [[ $line =~ ${re} ]] && var1="${BASH_REMATCH[1]}" && dt="${BASH_REMATCH[2]}"
        prod=$(basename ${var1})

        #-----------------------------------------
        # parse the date into usable fields...
        #-----------------------------------------
        [[ ${dt} =~ ${dtre} ]] && dmy="${BASH_REMATCH[1]}" &&
            hr="${BASH_REMATCH[2]}" &&
             m="${BASH_REMATCH[3]}" &&
             s="${BASH_REMATCH[4]}" &&
            tz="${BASH_REMATCH[5]}"

        #----------------------------------------------------------
        # build an ISO 8601 formatted string that date can parse
        #----------------------------------------------------------
        dt2=$(printf '%s %s:%s:%s %s' ${dmy} ${hr} ${m} ${s} ${tz})

        #----------------------------------------------------------
        # create a human readable date.  Unfortunately, this is
        # OS dependent, but it's not too bad.
        #----------------------------------------------------------
		case ${OS} in
			Darwin)
				dt3=$(date -j -f "%Y%m%d_%H%M%S%Z" "${dmy}_${hr}${m}${s}${tz}" "+%B %d, %Y %l:%M:%S%P %Z")
				;;
			Linux)
				dt3=$(date -d "${dmy} ${hr}:${m}:${s} ${tz}" "+%B %d, %Y %l:%M:%S%P %Z")
				;;
		esac 

        printf '%-11s  %-20s  %-50s\n' "${prod}" "${dt}" "${dt3}" >> ${BOM}
    done < ${BOMTMP}

    echo "---------------------------------------------------------------------" >> ${BOM}
    echo "Release Timestamp: ${ReadableDate}" >> ${BOM}

    rm ${BOMTMP}

	decho "*** Exiting bom()"
}

#############################################################################
# setChecksum
#   Description:
#   	Compute the checksum for the file ${1}
#
#   Params:
#		$1 = base product name.  Example:  roller
#
#	Returns:
#		md5Value  will be set to the checksum value
#       sha1Value will be set to the SHA1 value on Linux
#		xchkmd5 will be set to the value of the HTTP X-Checksum-Md5 header value
#		xchksha1 will be set to the value of the HTTP X-Checksum-Sha1 header value
#
#############################################################################
function setChecksum {
	decho "*** Entering setChecksum( ${1} )"
	xchkmd5=""
	xchksha1=""
	case "${OS}" in
	    Darwin)
			decho "*** Darwin"
	        MD5="md5 -r"
	        md5Value="$($MD5 "${1}")"
	        md5Value="${md5Value:0:32}"
	        echo "MD5 = ${md5Value} ${1}"
	        xchkmd5="-H \"X-Checksum-Md5: ${md5Value}\""
	        ;;  

	    "MINGW32_NT-6.2" | "CYGWIN_NT-6.2" )
	        echo "will attempt to md5sum ${1}"
	        MD5="md5sum"
	        which $MD5 || exit $?
	        md5Value="$($MD5 "${1}")"
	        md5Value="${md5Value:0:32}"
	        xchkmd5="-H \"X-Checksum-Md5: ${md5Value}\""
	        ;;  

	    *) 
	        echo "will attempt to md5sum ${1}"
	        MD5="md5sum"
	        which $MD5 || exit $?
	        md5Value="$($MD5 "${1}")"
	        md5Value="${md5Value:0:32}"
	        xchkmd5="-H \"X-Checksum-Md5: ${md5Value}\""
	        which sha1sum || exit $?
	        sha1Value="$(sha1sum "${1}")"
	        sha1Value="${sha1Value:0:40}"
	        echo "MD5 and SHA1 = ${md5Value} $sha1Value ${1}"
	        xchksha1="-H \"X-Checksum-Sha1: ${sha1Value}\""
	        ;;  
	esac

	decho "*** Exit setChecksum, xchkmd5 = ${xchkmd5}, xchksha1 = ${xchksha1}"
}

#############################################################################
# upload(from,to)
#   Description:
#               Copy the local file to the "to" repository file. Use
#				md5 and sha1 checksums if they are available
#
#   Params:
#               $1 = source repo file
#               $2 = destination repo file
#
#############################################################################
function upload {
	setChecksum ${1}
	if [ "x${xchkmd5}" = "x" -a "x${xchksha1}" = "x" ]; then
		echo "no checksum"
		curl -s -u "${USER}:${APIKEY}" -T ${1} ${2}
	elif [ "x${xchkmd5}" != "x" -a "x${xchksha1}" = "x" ]; then
		echo "sending md5 checksum"
		curl -s -u "${USER}:${APIKEY}" -H "X-Checksum-Md5: ${md5Value}" -T ${1} ${2}
	elif [ "x${xchkmd5}" = "x" -a "x${xchksha1}" != "x" ]; then
		echo "sending sha1 checksum"
		curl -s -u "${USER}:${APIKEY}" -H "X-Checksum-Sha1: ${sha1Value}" -T ${1} ${2}
	else
		echo "sending both md5 and sha1 checksums"
		curl -s -u "${USER}:${APIKEY}" -H "X-Checksum-Md5: ${md5Value}" -H "X-Checksum-Sha1: ${sha1Value}" -T ${1} ${2}
	fi
}

#############################################################################
# repomove(from,to)
#   Description:
#               Move the "from" repository file to the "to" repository file
#
#   Params:
#               $1 = source repo file
#               $2 = destination directory for the repo file
#
#############################################################################
function repomove {
	decho "*** Entering repomove( ${1} , ${2} )"
    dest=$(basename "${1}")
    curl -s -u "${USER}:${APIKEY}" ${URLBASE}${1} >${dest}  # get the file to move

cmd="upload ${dest} ${URLBASE}${2}"
echo "cmd = ${cmd}"
ls -l mojo*
pause

    upload ${dest} "${URLBASE}${2}"                         # upload to destination
    curl -i -X DELETE -u ${USER}:${APIKEY} ${URLBASE}${1}	# remove it from the old location
    rm ${dest}
	decho "*** Exiting repomove()"
}

readConfig


#-------------------------------------------------------
# Make sure that there's something to roll back to...
#-------------------------------------------------------
prodarchive=($(jfrog rt s "accord/air/archives/${PRODNAME}/${PRODNAME}*" | grep path| awk '{print $2;}' | sed 's/\"//g'))
prodarchiveLen=${#prodarchive[@]}
if (( ${prodarchiveLen}  == 0 )); then
	echo "*** ERROR ***"
	echo "Cannot rollback because accord/air/archives/${PRODNAME} is empty"
	exit 1
fi

#---------------------------------------------------
# Move accord/air/latest/prodname* to rollbacks/
#---------------------------------------------------
rmme=$(jfrog rt s "accord/air/release/${PRODNAME}*" | grep "${PRODNAME}" | awk '{print $2;}' | sed 's/"//g')
rmbase=$(basename ${rmme})
repomove ${rmme} "accord/air/rollbacks/${rmbase}"

#---------------------------------------------------
# Move accord/air/latest/BOM* to rollbacks/
#---------------------------------------------------
bompath=$(jfrog rt s "accord/air/release/bom*" | grep "bom" | awk '{print $2;}' | sed 's/"//g')
bom=$(basename ${bompath})
curl -s -u "${USER}:${APIKEY}" ${URLBASE}${bompath} >${bom}
echo "Rolled back ${PRODNAME} on ${TSFMT}" >> ${bom}		## mark the rollback date
upload ${bom} "${URLBASE}accord/air/rollbacks/${bom}"
rm -f ${bom}
curl -i -X DELETE -u ${USER}:${APIKEY} ${URLBASE}accord/air/release/${bom}

#------------------------------------------------------------
# Move the latest file in accord/air/archives/prodname*/ to
# accord/air/latest/ .  The latest file in the archive will be
# the last element in array 
#------------------------------------------------------------
idx=$(( prodarchiveLen - 1 ))
prevrel=${prodarchive[${idx}]}
dest=$(basename ${prevrel})

repomove ${prevrel} "accord/air/release/${dest}"

#-------------------------------------------------------------------
# Create a new bill of materials
#-------------------------------------------------------------------
bom
jfrog rt upload bom.txt "accord/air/release/bom_${TS}.txt"
rm -f bom.txt

exit 0