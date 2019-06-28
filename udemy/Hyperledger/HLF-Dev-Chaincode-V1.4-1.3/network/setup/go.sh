#!/bin/bash

if [ -z $SUDO_USER ]
then
    echo "===== Script need to be executed with sudo ===="
    echo "Change directory to 'setup'"
    echo "Usage: sudo ./caserver.sh"
    exit 0
fi

source ./to_absolute_path.sh

echo "=======Set up go======"
apt-get update
apt-get -y install golang-1.10-go

mv /usr/lib/go-1.10 /usr/local
mv /usr/local/go-1.10   /usr/local/go




# If GOROOT already set then DO Not set it again
if [ -z $GOROOT ]
then
    echo "export GOROOT=/usr/local/go" >> ~/.profile
    echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile

    GOPATH=$PWD/../../gocc
    to-absolute-path $GOPATH
    GOPATH=$ABS_PATH

    echo "export GOPATH=$GOPATH" >> ~/.profile
    echo "======== Updated .profile with GOROOT/GOPATH/PATH===="

    echo "export GOROOT=/usr/local/go" >> ~/.bashrc
    echo "export GOPATH=$GOPATH" >> ~/.bashrc
    echo "======== Updated .profile with GOROOT/GOPATH/PATH===="

    echo "export GOCACHE=off" >> ~/.bashrc


else
    echo "======== No Change made to .profile ====="
fi

echo "======= Done. PLEASE LOG OUT & LOG Back In ===="
echo "Then validate by executing    'go version'"