#!/bin/bash
#
# USAGE:
# 	release prodname
#
# SYNOPSIS: elevate airepo:/air/snapshot/<bundle> to release status in the 
#           repository
#
# DESCRIPTION:
# Archive the current product in air/accord/latest:
# 	1. Move air/accord/latest/prodname*.tar.gz  to air/archives/prodname/
# 	2. Move air/accord/latest/BOM* to air/archives
# 	3. If there are more than 10 files in air/accord/archives/prodname/ then
#      delete the oldest one so that we only keep the latest 10.
# 	4. Copy air/accord/snapshot/prodname.tar.gz to
#      air/latest/prodname_YYYYMMDD_hhmmssZZZ.tar.gz, where the time stamp is
#      set to the current time in the modified ISO 8601 format shown.
# 	5. Create a BOM_YYYMMDD_hhmmss.txt file with the system date containing 1
# 	   line with the version number associated with each bundle. Lines are of
#	   the format:  <bundle file name> <version stamp>
# 
# EXAMPLES:
#       ./release.sh rentroll
#------------------------------------------------------------------------------

TS=$(date +%Y%m%d_%H%M%S%Z)

PRODNAME=${1}
MaxArchiveDepth=10
DEBUG=1
OS=$(uname)
BOMMAX=100
PRODMAX=10

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
	if [ ! -f ${APATH}/config.json ]; then
		echo "config.json file was not found in ${APATH}"
		exit 1
	fi
    CONF="${APATH}/config.json"
    USER=$(grep RepoUser ${CONF} | awk '{print $2;}' | sed -e 's/[,"]//g')
    APIKEY=$(grep RepoPass ${CONF} | awk '{print $2;}' | sed -e 's/[,"]//g')
    URLBASE=$(grep RepoURL ${CONF} | awk '{print $2;}' | sed -e 's/[,"]//g')
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
# repocopy(from,to)
#   Description:
#               Copy the "from" repository file to the "to" repository file
#
#   Params:
#               $1 = source repo file
#               $2 = destination repo file
#
#############################################################################
function repocopy {
	decho "*** Entering repocopy( ${1} , ${2} )"

    dest=$(basename "${1}")
    curl -s -u "${USER}:${APIKEY}" ${URLBASE}${1} >${dest}
    upload ${dest} "${URLBASE}${2}"
    rm ${dest}

	decho "*** Exiting repocopy()"
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
    curl -s -u "${USER}:${APIKEY}" ${URLBASE}${1} >${dest}									# get the file to move
    upload ${dest} "${URLBASE}${2}"
    curl -i -X DELETE -u ${USER}:${APIKEY} ${URLBASE}${1}									# remove it from the old location
    rm ${dest}

	decho "*** Exiting repomove()"
}



#############################################################################
# setDefaultMax
#   Description:
#   	Create a max file for this archive directory.
#
#   Params:
#		$1 = base product name.  Example:  roller
#
#############################################################################
function setDefaultMax {
	if [ "${1}" = "bom" ]; then
		echo "${BOMMAX}" > qq.txt
	else
		echo "${PRODMAX}" > qq.txt
	fi
	upload qq.txt ${URLBASE}accord/air/archives/${1}/max.txt
	rm qq.txt
}

#############################################################################
# prune
#   Description:
#   	If there are more than MaxArchiveDepth files in the archive for
#		the supplied product, then remove the oldest files until it contains
#		the max.
#
#   Params:
#		$1 = base product name.  Example:  roller
#
#############################################################################
function prune {
	decho "*** Entering prune( ${1} )"

	#-------------------------------------------------------------------------
	# this statement creates an array of filenames that are in accord/air/
	#-------------------------------------------------------------------------
	n=($(jfrog rt s "accord/air/archives/${1}/${1}*" | grep path| awk '{print $2;}' | sed 's/\"//g'))

	#-------------------------------------------------------------------------
	# if there are files in the archive dir, then read "max.txt" to determine
	# how we need to prune it.
	#-------------------------------------------------------------------------
	max=${MaxArchiveDepth}  ## default value
	if (( ${#n[@]} > 0 )); then
		# first make sure that we actually have a max file...
		n1=$(jfrog rt s "accord/air/archives/${1}/max.txt" | grep "path:" | wc -l)
		if (( ${n1} == 0 )); then
			setDefaultMax "${1}"
		fi
		curl -s -u "${USER}:${APIKEY}" ${URLBASE}accord/air/archives/${1}/max.txt > max.txt
		max=$(cat max.txt)
		rm -f max.txt
	fi
	
	decho "*** PRUNE ${1} - Max = ${MaxArchiveDepth} -- found ${#n[@]} files"
	if (( ${#n[@]} > ${max} )); then
		LIMIT=$(( ${#n[@]} - MaxArchiveDepth ));
		decho "*** PRUNE ${1} -  will PRUNE ${LIMIT} files"
		i=0
		while [[ $i -lt ${LIMIT} ]]; do
			echo "Removing archived file ${n[${i}]}"
			curl -i -X DELETE -u ${USER}:${APIKEY} ${URLBASE}${n[${i}]}
			i=$(( i + 1 ))
		done
	fi
	decho "*** Exiting prune()"
}


#############################################################################
# archive
#   Description:
#   	Move the contents of accord/air/release/${1}* 
#		to air/archives/${1}/
#
#   Params:
#		$1 = base product name.  Example: roller
#
#############################################################################
function archive {
	decho "*** Entering archive( ${1} )"
	#-----------------------------------------------------------------
	# Move all versions of ${1} out of accord/air/release and
	# into accord/air/archives/${1}/
	# There should not be multiple versions, but this is a good
	# place to ensure that there's only 1 copy of the product in the
	# release directory
	#-----------------------------------------------------------------
	t="accord/air/archives/${1}/"
	jfrog rt s "accord/air/release/${1}*" | grep path | awk '{print $2;}' | sed 's/\"//g' | while read -r line; do
		repomove "${line}" "${t}"
	done
	
	t="accord/air/archives/bom/"
	jfrog rt s "accord/air/release/bom*" | grep path | awk '{print $2;}' | sed 's/\"//g' | while read -r line; do
		repomove "${line}" "${t}"
	done

	#----------------------------------------------------------------------
	# If there are more than 10 files in air/accord/archives/prodname/
	# then delete the oldest one so that we only keep the latest 10.
	#----------------------------------------------------------------------
	prune ${1}
	prune bom
	decho "*** Exiting archive"
}

decho "*** BEGIN RELEASE"
readConfig

#----------------------------------------------------------------------
# Archive the product bundle being updated.
# Move air/accord/latest/prodname*.tar.gz  to air/archives/prodname/
# Move air/accord/latest/BOM* to air/archives as well
#----------------------------------------------------------------------
archive "${PRODNAME}"

#----------------------------------------------------------------------
# Copy air/accord/snapshot/prodname.tar.gz to
# air/latest/prodname_YYYYMMDD_hhmmssZZZ.tar.gz to
# air/accord/latest/ where the time stamp is set to the current time
# in the modified ISO 8601 format shown.
#----------------------------------------------------------------------
f="accord/air/snapshot/${PRODNAME}.tar.gz"
t="accord/air/release/${PRODNAME}_${TS}.tar.gz"
repocopy "${f}" "${t}"

#----------------------------------------------------------------------
# Create a BOM_YYYMMDD_hhmmss.txt file with the system date containing
# 1 line with the version number associated with each bundle. Lines
# are of the format:  <bundle file name> <version stamp>
#----------------------------------------------------------------------
bom
jfrog rt upload bom.txt "accord/air/release/bom_${TS}.txt"
rm -f bom.txt

decho "*** END RELEASE"
