#!/bin/bash
#############################################################################
# decho
#   Description:
#   	Use this function like echo. If DEBUG is 1 then it will echo the
#		the output to the terminal. Otherwise it will just return without
#       echoing anything.
#
#   Params:
#		The string to echo
#
#	Returns:
#
#############################################################################
function decho {
	if (( ${DEBUG} == 1 )); then
		echo
		echo "${1}"
	fi
}

#############################################################################
# readConfig
#   Description:
#       Read the config.json file from the directory containing this script
#       to set some of the key values needed to access the artifactory repo.
#
#       Upon returning, URLBASE will end with the character "/".
#
#   Params:
#       none
#
#   Returns:
#       sets variables:  APIKEY, REPOUSER, and URLBASE
#
#############################################################################
readConfig() {
    RELDIR=$(cd `dirname "${BASH_SOURCE[0]}"` && pwd)
    CONF="${RELDIR}/config.json"
    REPOUSER=$(grep RepoUser ${CONF} | awk '{print $2;}' | sed -e 's/[,"]//g')
    APIKEY=$(grep RepoPass ${CONF} | awk '{print $2;}' | sed -e 's/[,"]//g')
    URLBASE=$(grep RepoURL ${CONF} | awk '{print $2;}' | sed -e 's/[,"]//g')

    # add a trailing / if it does not have one...
    if [ "${URLBASE: -1}" != "/" ]; then
        URLBASE="${URLBASE}/"
    fi

    decho "RELDIR = ${RELDIR}"
    decho "REPOUSER = ${REPOUSER}"
    decho "APIKEY = ${APIKEY}"
    decho "URLBASE = ${URLBASE}"
}

#############################################################################
# configure
#   Description:
#       The only configuration needed is the jfrog cli environment. Just
#       make sure we have it in the path. If it is not present, then
#       get it.
#
#   Params:
#       none
#
#   Returns:
#       nothing
#
#############################################################################
configure() {
    #---------------------------------------------------------
    # the user's bin directory is not created by default...
    #---------------------------------------------------------
    if [ ! -d ~ec2-user/bin ]; then
        mkdir ~ec2-user/bin
    fi

    #---------------------------------------------------------
    # now make sure that we have jfrog...
    #---------------------------------------------------------
    if [ ! -f ~ec2-user/bin/jfrog ]; then
        curl -s -u "${REPOUSER}:${APIKEY}" ${URLBASE}accord/tools/jfrog > ~ec2-user/bin/jfrog
        chown ec2-user:ec2-user ~ec2-user/bin/jfrog
        chmod +x ~ec2-user/bin/jfrog
    fi
    if [ ! -d ~ec2-user/.jfrog ]; then
        curl -s -u "${REPOUSER}:${APIKEY}" ${URLBASE}accord/tools/jfrogconf.tar > ~ec2-user/jfrogconf.tar
        pushd ~ec2-user
        tar xvf jfrogconf.tar
        rm jfrogconf.tar
        chown ec2-user:ec2-user ~ec2-user/bin/jfrog
        popd
    fi
    if [ ! -d ~root/.jfrog ]; then
        curl -s -u "${REPOUSER}:${APIKEY}" ${URLBASE}accord/tools/jfrogconf.tar > ~root/jfrogconf.tar
        pushd ~root
        tar xvf jfrogconf.tar
        rm jfrogconf.tar
        popd
    fi
}

#############################################################################
# GetLatestProductRelease
#   Description:
#       The only configuration needed is the jfrog cli environment. Just
#       make sure we have it in the path. If it is not present, then
#       get it.
#
#   Params:
#       ${1} = base name of product (rentroll, phonebook, mojo, ...)
#
#   Returns:
#       nothing
#
#############################################################################
GetLatestRepoRelease() {
    f=$(~ec2-user/bin/jfrog rt s "accord/air/release/*" | grep ${1} | awk '{print $2}' | sed 's/"//g')
    if [ "x${f}" = "x" ]; then
        echo "There are no product releases for ${f}"
        exit 1
    fi
    cd ${RELDIR}/..
    d=$(pwd)
    t=$(basename ${f})
    curl -s -u "${REPOUSER}:${APIKEY}" ${URLBASE}${f} > ${t}
}

#----------------------------------------------
#  ensure that we're in the rentroll directory...
#----------------------------------------------
dir=${PWD##*/}
if [ ${dir} != "rentroll" ]; then
    echo "This script must execute in the rentroll directory."
    exit 1
fi

user=$(whoami)
if [ ${user} != "root" ]; then
    echo "This script must execute as root.  Try sudo !!"
    exit 1
fi

readConfig
configure

echo -n "Shut down rentroll server: ";
$(./activate.sh stop) >/dev/null 2>&1
echo "done"

cd ${RELDIR}/..
rm -f rentroll*.tar*
echo "Distribution download to:  ${PWD}"
GetLatestRepoRelease "rentroll"

echo -n "Extracting: "
cd ${RELDIR}/..
tar xzf rentroll*.tar.gz
chown -R ec2-user:ec2-user rentroll
rm -f rentroll*.tar*
cd ${RELDIR}
echo "done"

echo -n "Activating: "
stat=$(./activate.sh start)
sleep 2
status=$(./activate.sh ready)
./installman.sh >installman.log 2>&1  # a task to perform while activation is running
if [ "${status}" = "OK" ]; then
    echo "Success!"
else
    echo "error: tatus = ${status}"
    echo "output from ./activate.sh -b start "
    echo "${stat}"
fi
