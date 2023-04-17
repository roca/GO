# -*- mode: ruby -*-
# vi: set ft=ruby :

$script = <<-SCRIPT
  echo "Provisioning..."
  apt update
  DEBIAN_FRONTEND=noninteractive apt -y dist-upgrade
  DEBIAN_FRONTEND=noninteractive apt install -y --install-recommends build-essential git make libelf-dev clang strace tar bpfcc-tools linux-headers-$(uname -r) gcc-multilib
  git clone --depth 1 git://kernel.ubuntu.com/ubuntu/ubuntu-bionic.git /kernel-src
  cd /kernel-src/tools/lib/bpf
  sudo make && sudo make install prefix=/usr/local
  sudo mv /usr/local/lib64/libbpf.* /lib/x86_64-linux-gnu/
  date > /etc/vagrant_provisioned_at
SCRIPT

Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/bionic64"
  config.vm.provision "shell", inline: $script
end
