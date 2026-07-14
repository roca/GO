#!/bin/bash

echo "Installing govendor....please wait"
# Installs the tools for Go
# https://github.com/kardianos/govendor/wiki/Govendor-CheatSheet
go get -u github.com/kardianos/govendor

# Install the package for the protocol
# ./install-protoc.sh

# echo "export PATH=$PATH:$GOPATH/bin" >> ~/.profile
./update-git-repo.sh

echo "Done. Logout and Log back in"
