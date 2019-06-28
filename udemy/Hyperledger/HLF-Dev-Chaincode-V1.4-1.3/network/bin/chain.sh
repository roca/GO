#!/bin/bash
# Carries out desired chaincode operations
# Make sure that the environment variables are set before invoking this script
function usage {
    echo    'Usage: chain.sh  install | instantiate | query | invoke | list | upgrade | install-auto | upgrade-auto'
}


DIR="$( which set-chain-env.sh)"
DIR="$(dirname $DIR)"

# Read the current setup
source   $DIR/cc.env.sh

OPERATION=$1

if [ -z $OPERATION ];
then
    usage
elif  [ "$OPERATION" == "install" ]
then
    echo "==>Installing chaincode"
    peer chaincode install -l "$CC_LANGUAGE" -n "$CC_NAME" -p "$CC_PATH" -v "$CC_VERSION"
elif  [ "$OPERATION" == "instantiate" ] 
then
    EP_FLAG=""
    if [ "$CC_ENDORSEMENT_POLICY" != "" ]; then
        MSG="==>Instantiating chaincode with EP=$CC_ENDORSEMENT_POLICY"
        EP_FLAG="-P"
        # EP_FLAG=" -P  \"$CC_ENDORSEMENT_POLICY\""
        # echo $EP_FLAG
    else
        MSG="==>Instantiating chaincode"
    fi

    PDC_FLAG=""
    if [ "$CC_PRIVATE_DATA_JSON" != "" ]; then
        MSG="$MSG PDC=$CC_PATH/$CC_PRIVATE_DATA_JSON"
        PDC_FLAG="--collections-config"
    fi

    echo $MSG

    peer chaincode instantiate -C "$CC_CHANNEL_ID" -n "$CC_NAME" -v "$CC_VERSION" -c "$CC_CONSTRUCTOR"   -o "$ORDERER_ADDRESS" -l "$CC_LANGUAGE"  "$PDC_FLAG" "$GOPATH/src/$CC_PATH/$CC_PRIVATE_DATA_JSON" "$EP_FLAG" "$CC_ENDORSEMENT_POLICY"
# elif  [ "$OPERATION" == "instantiate-priv" ] 
# then
#     echo "==>Instantiating chaincode with private collections"
    
#     peer chaincode instantiate  -C "$CC_CHANNEL_ID" -n "$CC_NAME" -v "$CC_VERSION" -c "$CC_CONSTRUCTOR"  -o "$ORDERER_ADDRESS" -l "$CC_LANGUAGE" --collections-config "$GOPATH/src/$CC_PATH/$CC_PRIVATE_DATA_JSON"
elif  [ "$OPERATION" == "query" ] 
then
    echo "==>Querying chaincode"
    # Override if there is another parameter
    if [ "$2" != "" ]; then
        # Override the Invoke parameter
        CC_INVOKE_ARGS=$2
    fi
    QUERY_RESULT=$(peer chaincode query -C "$CC_CHANNEL_ID" -n "$CC_NAME"  -c "$CC_QUERY_ARGS")
    export QUERY_RESULT
    echo -e $QUERY_RESULT
elif  [ "$OPERATION" == "invoke" ]; then
    echo "==>Invoking chaincode"
    # Override if there is another parameter
    if [ "$2" != "" ]; then
        # Override the Invoke parameter
        CC_INVOKE_ARGS=$2
    fi
    peer chaincode invoke -o "$ORDERER_ADDRESS" -C "$CC_CHANNEL_ID" -n "$CC_NAME"  -c "$CC_INVOKE_ARGS"
elif  [ "$OPERATION" == "list" ]; then
    echo "==>Listing Installed chaincode"
    peer chaincode list --installed
    echo "==>Listing Instantiate chaincode on Channel: $CC_CHANNEL_ID"
    peer chaincode list --instantiated -C "$CC_CHANNEL_ID"
# Upgrade
elif  [ "$OPERATION" == "upgrade" ]; then
    EP_FLAG=""
    if [ "$CC_ENDORSEMENT_POLICY" != "" ]; then
        MSG="==>Instantiating chaincode with EP=$CC_ENDORSEMENT_POLICY"
        EP_FLAG="-P"
        # EP_FLAG=" -P  \"$CC_ENDORSEMENT_POLICY\""
        # echo $EP_FLAG
    else
        MSG="==>Instantiating chaincode"
    fi

    PDC_FLAG=""
    if [ "$CC_PRIVATE_DATA_JSON" != "" ]; then
        MSG="$MSG PDC=$CC_PATH/$CC_PRIVATE_DATA_JSON"
        PDC_FLAG="--collections-config"
    fi

    echo $MSG


    # Requires the manual install of new version
    # Use the constructor args
    peer chaincode upgrade -n "$CC_NAME" -v "$CC_VERSION" -C "$CC_CHANNEL_ID"  -c "$CC_CONSTRUCTOR"  "$PDC_FLAG" "$GOPATH/src/$CC_PATH/$CC_PRIVATE_DATA_JSON" "$EP_FLAG" "$CC_ENDORSEMENT_POLICY"

elif [ "$OPERATION" == "install-auto" ]; then
    # Increment the version
    NEW_VERSION="$CC_VERSION.1"
    . set-chain-env.sh -v $NEW_VERSION
    CC_VERSION=$NEW_VERSION

    # Now install with the new version
    echo "===>Setting version=$CC_VERSION"
    chain.sh install

elif  [ "$OPERATION" == "upgrade-auto" ]; then
    # Increment the version
    NEW_VERSION="$CC_VERSION.1"
    . set-chain-env.sh -v $NEW_VERSION
    CC_VERSION=$NEW_VERSION
    # Now install
    echo "===>Setting version=$CC_VERSION"
    chain.sh install

    EP_FLAG=""
    if [ "$CC_ENDORSEMENT_POLICY" != "" ]; then
        MSG="==>Instantiating chaincode with EP=$CC_ENDORSEMENT_POLICY"
        EP_FLAG="-P"
        # EP_FLAG=" -P  \"$CC_ENDORSEMENT_POLICY\""
        # echo $EP_FLAG
    else
        MSG="==>Auto Upgrading chaincode"
    fi

    PDC_FLAG=""
    if [ "$CC_PRIVATE_DATA_JSON" != "" ]; then
        MSG="$MSG PDC=$CC_PATH/$CC_PRIVATE_DATA_JSON"
        PDC_FLAG="--collections-config"
    fi

    echo $MSG

    peer chaincode upgrade -n "$CC_NAME" -v "$CC_VERSION" -C "$CC_CHANNEL_ID"  -c "$CC_CONSTRUCTOR" "$PDC_FLAG" "$GOPATH/src/$CC_PATH/$CC_PRIVATE_DATA_JSON" "$EP_FLAG" "$CC_ENDORSEMENT_POLICY"
    
else
    usage
    echo "Invalid operation!!!"
fi
