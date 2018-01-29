#!/bin/bash
#
# USAGE:
# 	snapshot bundle
#
# SYNOPSIS: push the supplied snapshot tar.gz file to the air snapshot repo
#
# DESCRIPTION:
#	Copy the supplied bundle to airep:/accord/air/snapshot/
# 
# params::
#   $1  the bundle to upload
#------------------------------------------------------------------------------

if [ "x${1}" == "x" ]; then
	echo "You must supply the bundle to release"
	exit 1
fi

jfrog rt u ${1} accord/air/snapshot/${1}
