#!/bin/bash
# Installs node JS and the utilities
sudo ./node.sh

rm -rf temp
mkdir temp
cd temp

# If owner is not changed we get a Warning (Harmless but annoying)
mkdir -p  $HOME/.config
# sudo chown -R $(whoami) $HOME/.config
sudo chown -R $USER $HOME/.config
sudo chown -R $USER $HOME/.npm
git clone  https://github.com/acloudfan/HLFChaincode_Utils.git

# Remove the previous version if any
sudo rm -rf $HOME/HLFChaincode_Utils &> /dev/null
sleep 2s
# Move the latest version 
mv HLFChaincode_Utils  $HOME

# Change directory and do an npm install
cd $HOME/HLFChaincode_Utils
npm install

cd ..
rm -rf temp 