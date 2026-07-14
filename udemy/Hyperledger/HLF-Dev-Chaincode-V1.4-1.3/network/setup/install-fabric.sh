#!/bin/bash

if [ -z $SUDO_USER ]; then
    echo "Script MUST be executed with 'sudo -E'"
    echo "Abroting!!!"
    exit 0
fi

if [ -z $GOPATH ]; then
    echo "GOPATH not set!!! You must use 'sudo -E ./install-fabric.sh'"
    echo "Aborting!!!"
    exit 0
fi

source ./to_absolute_path.sh

export PATH=$PATH:$GOROOT/bin

echo "GOPATH=$GOPATH"
echo "GOROOT=$GOROOT"

mkdir -p $GOPATH

# Execute in the setup folder
echo "=== Must execute in the setup folder ===="

rm -rf ./temp 2> /dev/null
# create temp directory
mkdir temp
cd temp

echo "====== Starting to Download Fabric ========"
# Documentation or Bootstrap script has an issue
# If -s is placed before the versions (per doc) then you will see a harmless error msg
#curl -sSL http://bit.ly/2ysbOFE | bash  1.3.0 1.3.0 0.4.10 -s

curl -sSL http://bit.ly/2ysbOFE -o bootstrap.sh

chmod 755 ./bootstrap.sh

# bash ./bootstrap.sh  1.3.0 1.3.0 0.4.10 -s
# bash ./bootstrap.sh  1.4.0 1.4.0 0.4.10 -s
bash ./bootstrap.sh  1.4.1 1.4.1 0.4.10 -s



echo "======= Copying the binaries to /usr/local/bin===="
cp bin/*    /usr/local/bin

# This downloads the shim code 
echo "======= Setting up the HLF Shim ===="
mkdir -p  $GOPATH/src/github.com/hyperledger
go get -u --tags nopkcs11 github.com/hyperledger/fabric/core/chaincode/shim

# The sample chaincode is under the subfolder go and need to come under gopath/src subfolder

cd ..
rm -rf temp

BIN_PATH=$PWD/../bin
to-absolute-path $BIN_PATH
BIN_PATH=$ABS_PATH

echo "export PATH=$PATH:$BIN_PATH:$GOPATH/bin" >> ~/.profile
echo "export PATH=$PATH:$BIN_PATH:$GOPATH/bin" >> ~/.bashrc

chmod u+x $BIN_PATH/*.sh


# Update /etc/hosts
source    ./manage_hosts.sh
HOSTNAME=acme-peer1.acme.com
removehost $HOSTNAME            &> /dev/null
addhost $HOSTNAME
HOSTNAME=budget-peer1.budget.com
removehost $HOSTNAME            &> /dev/null
addhost $HOSTNAME
HOSTNAME=orderer.acme.com
removehost $HOSTNAME            &> /dev/null
addhost $HOSTNAME
HOSTNAME=postgresql
removehost $HOSTNAME            &> /dev/null
addhost $HOSTNAME
HOSTNAME=explorer
removehost $HOSTNAME            &> /dev/null
addhost $HOSTNAME

echo "Done. Logout and Log back in !!"
