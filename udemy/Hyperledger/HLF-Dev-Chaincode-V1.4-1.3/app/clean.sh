#!/bin/bash
# Script for cleaning the sample code folder - to be executed in VM

rm -rf sdk/node_modules
rm sdk/package-lock.json

rm -rf sdk/client/credstore
rm -rf sdk/gateway/user-wallet



echo "Done."