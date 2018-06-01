#!/usr/bin/env bash

ulimit -n 50000


sudo yum update
sudo yum -y install git


#
#ZIP
#
sudo yum install -y zip
sudo yum install -y unzip


#Apache
sudo yum install -y  httpd
sudo chkconfig --levels 235 httpd on
sudo rm -rf /var/www
sudo ln -fs /vagrant /var/www
sudo /bin/systemctl restart  httpd.service

#
# GO
#
sudo yum -y install go
sudo tar -C /usr/local -xzf /vagrant/lib/go/go1.4.2.linux-amd64.tar.gz
echo 'export GOPATH=$HOME/go' >> /home/vagrant/.bash_profile
echo 'export PATH=$PATH:/usr/local/go/bin' >> /home/vagrant/.bash_profile
echo 'export PKG_CONFIG_PATH=/vagrant/lib/go/config' >> /home/vagrant/.bash_profile
#
# MYSQL
#
# There are numerous prompts for root password
#
sudo yum -y localinstall /vagrant/lib/mysql/mysql-community-release-el7-5.noarch.rpm
sudo yum -y install mysql-server
sudo yum -y install mysql-devel
sudo yum -y install libaio
sudo /sbin/service mysqld restart
sudo yum -y install gcc
sudo yum -y install java
sudo yum install python-devel


#
# ORACLE: Add this to the .bash_profile
#
echo 'export LD_LIBRARY_PATH=/vagrant/lib/oracle/instantclient_12_1/lib:$LD_LIBRARY_PATH' >> /home/vagrant/.bash_profile
echo 'export PATH=/vagrant/lib/oracle/instantclient_12_1/bin:$PATH' >> /home/vagrant/.bash_profile
echo 'export ORACLE_HOME=/vagrant/lib/oracle/instantclient_12_1' >> /home/vagrant/.bash_profile
echo 'export NLS_LANG=AMERICAN_AMERICA.UTF8' >> /home/vagrant/.bash_profile
echo 'export GIT_SSL_NO_VERIFY=true' >> /home/vagrant/.bash_profile

# Docker
sudo yum install -y docker

#
# RVM
#
su vagrant -c  'gpg --keyserver hkp://keys.gnupg.net --recv-keys D39DC0E3'
su vagrant -c  '\curl -sSL https://get.rvm.io | bash'
echo 'export PATH="$PATH:$HOME/.rvm/bin"' >>  /home/vagrant/.bash_profile
echo '[[ -s "$HOME/.rvm/scripts/rvm" ]] && source "$HOME/.rvm/scripts/rvm"' >> /home/vagrant/.bash_profile


export PATH="$PATH:/home/vagrant/.rvm/bin"
[[ -s "/home/vagrant/.rvm/scripts/rvm" ]] && source "/home/vagrant/.rvm/scripts/rvm"

# #
# # RUBY
# #
rvm install 1.9.2
rvm install 1.9.3
rvm install ruby-1.9.3-p392
rvm install 2.0.0
rvm install 2.2.2
rvm use 2.2.2 --default
rvm default
sudo yum install libxml2-devel
sudo yum install libxslt-devel


# #
# # RAILS
# #
gem install rails
gem install bundler

#
# AWS SDK
#
gem install pry
gem install aws-sdk

#
#  pip
#
sudo python /vagrant/lib/pip/get-pip.py

#
# AWS CLI
#
sudo /vagrant/lib/aws-cli/awscli-bundle/install -i /usr/local/aws -b /usr/local/bin/aws
sudo pip install awsebcli --ignore-installed six
#EC2
sudo mkdir /usr/local/ec2
sudo unzip /vagrant/lib/aws-cli/ec2-api-tools.zip -d /usr/local/ec2
echo 'export EC2_HOME=/usr/local/ec2/ec2-api-tools-1.7.3.0' >> /home/vagrant/.bash_profile
sudo yum install java
echo 'export JAVA_HOME=/usr/lib/jvm/java-1.7.0-openjdk-1.7.0.75-2.5.4.2.el7_0.x86_64/jre' >> /home/vagrant/.bash_profile
sudo unzip /vagrant/lib/aws-cli/ec2-ami-tools.zip -d /usr/local/ec2
echo 'export EC2_AMITOOL_HOME=/usr/local/ec2/ec2-ami-tools-1.5.6' >> /home/vagrant/.bash_profile


# #
# # Passanger
# #
gem install passenger
sudo yum -y install curl-devel httpd-devel
passenger-install-apache2-module

sudo chmod 755 /home/vagrant

# extras
sudo yum install dos2unix


#
# PKG_CONFIG_PATH#
#
sudo yum -y install postgresql
sudo yum install postgresql-server
sudo yum -y install postgresql-devel

# PhantomJS
sudo tar -xzf /vagrant/lib/phantomjs-1.9.8-linux-x86_64.tar.bz2
sudo cp /vagrant/lib/phantomjs-1.9.8-linux-x86_64/bin/phantomjs /usr/bin
rm -rf /vagrant/lib/phantomjs-1.9.8-linux-x86_64


sudo tar -xzf /vagrant/lib/phantomjs-2.1.1-linux-x86_64.tar.bz2
sudo cp /vagrant/lib/phantomjs-2.1.1-linux-x86_64/bin/phantomjs /usr/bin
rm -rf /vagrant/lib/phantomjs-2.1.1-linux-x86_64

#
# Node
#
# Must run as root
#
# curl -sL https://rpm.nodesource.com/setup | sudo bash -
# sudo yum install -y nodejs
#
# Express
#
# sudo npm install -g express
# sudo npm install -g express-generator@4
# sudo npm install jspm -g
#

#
#VNC setup
#
sudo yum -y groupinstall "Desktop" "Desktop Platform" "X Window System" "Fonts"
sudo yum -y install tigervnc-server
sudo yum -y install xorg-x11-fonts-Type1
