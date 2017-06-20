#!/bin/bash

BREWHOME="/home/linuxbrew/.linuxbrew"
BREWBIN="${BREWHOME}/bin"

# must execute as root

echo "Installing Linuxbrew"
ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Linuxbrew/install/master/install)"

echo "Updating PATH"
PATH="${BREWBIN}":$PATH"
echo 'export PATH="/home/linuxbrew/.linuxbrew/bin:$PATH"' >>~/.bash_profile
source ~/.bash_profile
brew update --force

echo "brew install mysql"
cd ${BREWBIN}
cp mysql.server /etc/init.d
chkconfig --add mysql.server

cd /usr/bin
for i in ${BREWBIN}/mysql* ; do ln -s $i; done

cd ${BREWBIN}
sudo mysql_upgrade
