#!/bin/bash
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
#       sets variables:  APIKEY, USER, and URLBASE
#
#############################################################################
readConfig() {
    RELDIR=$(cd `dirname "${BASH_SOURCE[0]}"` && pwd)
    CONF="${RELDIR}/config.json"
    USER=$(grep RepoUser ${CONF} | awk '{print $2;}' | sed -e 's/[,"]//g')
    APIKEY=$(grep RepoPass ${CONF} | awk '{print $2;}' | sed -e 's/[,"]//g')
    URLBASE=$(grep RepoURL ${CONF} | awk '{print $2;}' | sed -e 's/[,"]//g')

    # add a trailing / if it does not have one...
    if [ "${URLBASE: -1}" != "/" ]; then
        URLBASE="${URLBASE}/"
    fi

    echo "USER = ${USER}"
    echo "APIKEY = ${APIKEY}"
    echo "URLBASE = ${URLBASE}"
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
        curl -s -u "${USER}:${APIKEY}" ${URLBASE}accord/tools/jfrog > ~ec2-user/bin/jfrog
        chown ec2-user:ec2-user ~ec2-user/bin/jfrog
        chmod +x ~ec2-user/bin/jfrog
    fi
    if [ ! -d ~ec2-user/.jfrog ]; then
        curl -s -u "${USER}:${APIKEY}" ${URLBASE}accord/tools/jfrogconf.tar > ~ec2-user/jfrogconf.tar
        pushd ~ec2-user
        tar xvf jfrogconf.tar
        rm jfrogconf.tar
        chown ec2-user:ec2-user ~ec2-user/bin/jfrog
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
    echo "preparing to load release bundle into directory ${d}"
    t=$(basename ${f})
    curl -s -u "${USER}:${APIKEY}" ${URLBASE}${f} > ${t}
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

echo -n "Shutting down rentroll server."; $(./activate.sh stop) >/dev/null 2>&1
echo -n "."
echo -n "."; cd ${RELDIR}/..
echo -n "."; rm -f rentroll*.tar
echo
echo -n "Retrieving latest released Rentroll..."

GetLatestRepoRelease "rentroll"

echo "Installing.."
echo -n "."; cd ${RELDIR}/..
#echo -n "."; rm -f rentroll*.tar
echo -n "."; gunzip -f rentroll*.tar.gz
echo -n "."; tar xf rentroll*.tar
echo -n "."; chown -R ec2-user:ec2-user rentroll
echo -n "."; rm -f rentroll*.tar*
echo -n "."; cd ${RELDIR}
echo

echo -n "Installation complete.  Launching..."
echo -n "."; ./activate.sh start
echo -n "."; sleep 2
echo -n "."; status=$(./activate.sh ready)
echo -n "."; ./installman.sh >installman.log 2>&1  # a task to perform while activation is running
echo
if [ "${status}" = "OK" ]; then
    echo "Activation successful"
else
    echo "Problems activating rentroll.  Status = ${status}"
fi
