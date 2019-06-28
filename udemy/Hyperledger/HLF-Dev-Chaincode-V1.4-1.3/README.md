
####################### https://courses.pragmaticpaths.com ##############
# Part of a course 'Mastering Fabric Chaincode Development using GoLang'#
# http://www.bcmentors.com      raj@acloudfan.com                       #
#########################################################################

Change Log:
----------
HLF-Dev-Chaincode-V1.4-1.3      June 6, 2019
* Docker compose install failure (reported by Emir)
* Vagrantfile port mapping issue (7051:7052 instead of 7052:7052)

HLF-Dev-Chaincode-V1.4-1.2      May 17, 2019
* Install Fabric Script Change - Mac OS install was failing

HLF-Dev-Chaincode-V1.4-1.1      May 10, 2019

* Updated the shell script for updating the .bashrc for PATH
* Fixed errors in install-protoc.sh
  Updated the lecture "Hands On: Using functions for ...."
* Fixed the install-node-utils.sh for sporadic permission errors
* Default log level changed to info
* Added script [setup/update-git-repo.sh] for pulling latest sample code from GitHub
* Added script [update-sh-file-mode.sh] for updating mode for all of the shell scripts

HLF-Dev-Chaincode-V1.4-1.0      April 18, 2019


#############################################
# Setup Instructions: network/setup/README.md
#############################################



1. Initialize the Dev invironment
> dev-init.sh   -flag  -flag

To get help
> dev-init.sh -h

2. Explorer
- Start with the environment use the -e flag
> dev-init.sh -e
> exp-stop.sh

- At a later time
> exp-start.sh
> exp-stop.sh

3. Environment scripts
- Set the org context
>  .   set-env.sh    acme | budget
- To check current environment
> show-env.sh

4. Chaincode operations
- Set the chaincode environment
> set-chain-env.sh  -flag ..  -flag ..
- To check current chaincode parameters
> set-chain-env.sh   

5. Unit testing



# VBOX Mount issue
sudo service vboxadd-service stop
sudo date -s "2010-10-01 10:25:00"

Vendor package management
==========================

External package dependency
https://hyperledger-fabric.readthedocs.io/en/release-1.4/chaincode4ade.html?highlight=govendor


https://larry-price.com/blog/2018/04/26/quick-start-to-vendor-go-dependencies-with-govendor/
go get -u github.com/kardianos/govendor

someone faced this issue

Instantiate complains that it cannot assign the var to vendor/.. same package ...

to fix=>
govendor init
govendor add +external
govendor add github.com/hyperledger/fabric/core/chaincode/shim
govendor add github.com/hyperledger/fabric/protos/peer
govendor add github.com/hyperledger/fabric/protos/peer

