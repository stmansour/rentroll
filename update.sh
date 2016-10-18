PASS=AP3wHZhcQQCvkC4GVCCZzPcqe3L
ART=http://ec2-52-91-201-195.compute-1.amazonaws.com/artifactory
USR=accord

EXTERNAL_HOST_NAME=$(curl -s http://169.254.169.254/latest/meta-data/public-hostname)
#${EXTERNAL_HOST_NAME:?"Need to set EXTERNAL_HOST_NAME non-empty"}

#--------------------------------------------------------------
#  Routine to download files from Artifactory
#--------------------------------------------------------------
artf_get() {
    echo "Downloading $1/$2"
    wget -O "$2" --user=$USR --password=$PASS ${ART}/"$1"/"$2"
}

loadAccordTools() {
    #--------------------------------------------------------------
    #  Let's get our tools in place...
    #--------------------------------------------------------------
    artf_get ext-tools/utils accord-linux.tar.gz
    echo "Installing /usr/local/accord" >>${LOGFILE}
    cd /usr/local
    tar xzf ~ec2-user/accord-linux.tar.gz
    chown -R ec2-user:ec2-user accord
    cd ~ec2-user/
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

echo -n "Shutting down rentroll server."; $(./activate.sh stop) >/dev/null 2>&1
echo -n "."
echo -n "."; 
echo -n "."; cd ..
echo
echo -n "Retrieving latest development snapshot of rentroll..."
/usr/local/accord/bin/getfile.sh jenkins-snapshot/rentroll/latest/rentroll.tar.gz
echo
echo -n "."; gunzip -f rentroll.tar.gz
echo -n "."; tar xf rentroll.tar
echo -n "."; chown -R ec2-user:ec2-user rentroll
echo -n "."; cd rentroll/
echo -n "."; echo -n "starting..."
echo -n "."; ./activate.sh start
echo -n "."; sleep 3
echo -n "."; status=$(./activate.sh ready)
echo
if [ "${status}" = "OK" ]; then
    echo "Activation successful"
else
    echo "Problems activating rentroll.  Status = ${status}"
fi
