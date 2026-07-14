#!/bin/bash

# Install Node JS
# sudo ./node.sh

# Create a temp folder
mkdir temp

# Download the 8.9.x distribution
# wget http://nodejs.org/dist/v8.9.4/node-v8.9.4-linux-x64.tar.gz

cd temp

wget http://nodejs.org/dist/v8.15.1/node-v8.15.1-linux-x64.tar.gz

# Untar it to /usr/local
tar -C /usr/local --strip-components 1 -xzf node-v8.15.1-linux-x64.tar.gz

# Remove the temp folder
cd ..
rm -rf ./temp

# Show the versions
echo -n "Installed Node version:"
node -v
echo -n "Installed Npm version:"
npm -v
echo "Done."

# rm ./node-v8.15.1-linux-x64.tar.gz &> /dev/null

#Removing the Node
# find / -name "node"
# sudo rm -rf /usr/local/bin/npm /usr/local/share/man/man1/node* /usr/local/lib/dtrace/node.d ~/.npm ~/.node-gyp /opt/local/bin/node opt/local/include/node /opt/local/lib/node_modules 
# sudo rm -rf /usr/local/lib/node*
# sudo rm -rf /usr/local/include/node*
# sudo rm -rf /usr/local/bin/node*
# sudo rm -rf /usr/local/share/doc/node