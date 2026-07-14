# Removes the file cc.env.sh
# Since the Chaincode parameters are persisted in the file
# Sometimes it may lead to unintended impact. As a good practice
# Reset the chaincode environment before testing
DIR="$( which set-chain-env.sh)"
DIR="$(dirname $DIR)"

rm $DIR/cc.env.sh

unset CC_INVOKE_ARGS
unset CC_QUERY_ARGS
unset CC_PRIVATE_DATA_JSON
unset CC_PATH
unset CC_CHANNEL_ID
unset CC_LANGUAGE
unset CC_NAME
unset CC_ENDORSEMENT_POLICY
unset CC_CONSTRUCTOR
unset CC_VERSION

echo "# Chaincode Environment Initialized!!!" > $DIR/cc.env.sh
set-chain-env.sh -n token -C airlinechannel

echo "DELETED   cc.env.sh !!!"
echo "Please set the environment for chaincode using set-chain-env.sh"