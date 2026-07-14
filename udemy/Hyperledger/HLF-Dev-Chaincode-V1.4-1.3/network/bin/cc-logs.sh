#!/bin/bash

# Provides easy access to the chaincode container logs

usage() {
    echo "cc-logs.sh   [-o ORG_NAME default=acme]  [-p PEER_NAME default=acme-peer1] [-f   Follow] [-t  Number]"
    echo "              Shows the logs for the chaincode container. "
    echo "              -f follows the logs, useful when debugging in net mode "
    echo "              -t flag is equivalent to --tail flag for docker logs"
    echo "              -h  shows the usage"
}

CONTAINER_PREFIX=dev
ORG_NAME=acme
PEER_NAME='acme-peer1'
source cc.env.sh
LOG_OPTIONS=""
while getopts "o:p:ft:h" OPTION; do
    case $OPTION in
    o)
        ORG_NAME=${OPTARG}
        ;;
    p)
        PEER_NAME=${OPTARG}
        ;;
    f)
        LOG_OPTIONS="-f"
        ;;
    t)
        LOG_OPTIONS="--tail ${OPTARG}"
        ;;
    h)
        usage
        exit 1
        ;;
    *)
        echo "Incorrect options provided"
        exit 1
        ;;
    esac
done

CC_CONTAINER_NAME="$CONTAINER_PREFIX-$PEER_NAME.$ORG_NAME.com-$CC_NAME-$CC_VERSION"

echo "Logs for container:  $CC_CONTAINER_NAME     Usage:   dev-logs.sh   -u"



docker logs $LOG_OPTIONS $CC_CONTAINER_NAME 

