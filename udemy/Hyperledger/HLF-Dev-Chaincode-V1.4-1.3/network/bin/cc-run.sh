#!/bin/bash
#Executes the chaincode in dev mode testing


DIR="$( which $BASH_SOURCE)"
DIR="$(dirname $DIR)"

source $DIR/to_absolute_path.sh
source $DIR/cc.env.sh

echo "GOPATH=$GOPATH  Name=$CC_NAME"
# Executes chaincode in dev mode

# To use this script launch the dev env in dev mode
source dev-mode.sh
if [ "$PEER_MODE" == "net" ]; then
    echo "=====>Can't run chaincode in terminal in 'net' mode!!!"
    echo "=====>Please use 'dev-init.sh  -d'  to launch the env in DEV mode."

    exit 1
fi


echo "+++Building the GoLang chaincode"
#go build -o $CC_PATH
export GOCACHE=off

export CORE_CHAINCODE_ID_NAME=$CC_NAME:$CC_VERSION

if [ "$ORGANIZATION_CONTEXT" == "acme" ] ; then
    export CORE_PEER_ADDRESS=localhost:7052 
elif [ "$ORGANIZATION_CONTEXT" == "budget" ] ; then
    export CORE_PEER_ADDRESS=localhost:8052 
else
    echo "Unknown Organization Context = $ORGANIZATION_CONTEXT....Aborting!!!"
    exit 0
fi

go run  $GOPATH/src/$CC_PATH/*.go

echo "+ Launching the chaincode"

