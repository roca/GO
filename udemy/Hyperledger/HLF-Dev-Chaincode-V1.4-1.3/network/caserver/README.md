export FABRIC_CFG_PATH=$PWD/config
configtxgen -outputBlock  ./config/airlinegenesis.block -channelID ordererchannel  -profile AirlineOrdererGenesis
configtxgen -outputCreateChannelTx  ./config/airlinechannel.tx -channelID airlinechannel  -profile AirlineChannel