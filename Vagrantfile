# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  # The most common configuration options are documented and commented below.
  # For a complete reference, please see the online documentation at
  # https://docs.vagrantup.com.

  #############################################################################
  # VMs & Boxes
  # -----------
  #
  # Resources:
  # https://www.vagrantup.com/docs/multi-machine
  # https://www.vagrantup.com/docs/provisioning/shell
  #
  # Every Vagrant development environment requires a box. You can search for
  # boxes at https://vagrantcloud.com/search.

  $centos_setup = <<-HEREDOC
# Install nvm (https://github.com/nvm-sh/nvm):
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.36.0/install.sh | bash
HEREDOC

  # CentOS 8
  # config.vm.define "centos8" do |centos8|
  #   centos8.vm.box = "centos/8"
  #   centos8.vm.box_version = "1905.1"
  #   centos8.vm.provision :shell, inline: $centos_setup, privileged: false
  # end

  # CentOS 7
  # config.vm.define "centos7" do |centos7|
  #   centos7.vm.box = "centos/7"
  #   centos7.vm.provision :shell, inline: $centos_setup, privileged: false
  # end

  $ubuntu_setup = <<-HEREDOC
# Install node and npm:
sudo apt update -y
sudo apt install nodejs -y
sudo apt install npm -y

# Install nvm (https://github.com/nvm-sh/nvm):
# curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.36.0/install.sh | bash
HEREDOC

  # Ubuntu 20.10
  # config.vm.define "ubuntu2010" do |ubuntu2010|
  #   ubuntu2010.vm.box = "ubuntu/groovy64"
  #   ubuntu2010.vm.provision :shell, inline: $ubuntu_setup, privileged: false
  # end

  # Ubuntu 20.04 LTS
  # config.vm.define "ubuntu2004" do |ubuntu2004|
  #   ubuntu2004.vm.box = "ubuntu/focal64"
  #   ubuntu2004.vm.provision :shell, inline: $ubuntu_setup, privileged: false
  # end

  # Ubuntu 18.04 LTS
  # config.vm.define "ubuntu1804" do |ubuntu1804|
  #   ubuntu1804.vm.box = "hashicorp/bionic64"
  #   ubuntu1804.vm.provision :shell, inline: $ubuntu_setup, privileged: false
  # end

  #############################################################################

  # Disable automatic box update checking. If you disable this, then
  # boxes will only be checked for updates when the user runs
  # `vagrant box outdated`. This is not recommended.
  # config.vm.box_check_update = false

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine. In the example below,
  # accessing "localhost:8080" will access port 80 on the guest machine.
  # NOTE: This will enable public access to the opened port
  # config.vm.network "forwarded_port", guest: 80, host: 8080

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine and only allow access
  # via 127.0.0.1 to disable public access
  # config.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"

  # Create a private network, which allows host-only access to the machine
  # using a specific IP.
  # config.vm.network "private_network", ip: "192.168.33.10"

  # Create a public network, which generally matched to bridged network.
  # Bridged networks make the machine appear as another physical device on
  # your network.
  # config.vm.network "public_network"

  # Share an additional folder to the guest VM. The first argument is
  # the path on the host to the actual folder. The second argument is
  # the path on the guest to mount the folder. And the optional third
  # argument is a set of non-required options.
  # config.vm.synced_folder "../data", "/vagrant_data"

  # Provider-specific configuration so you can fine-tune various
  # backing providers for Vagrant. These expose provider-specific options.
  # Example for VirtualBox:
  #
  # config.vm.provider "virtualbox" do |vb|
  #   # Display the VirtualBox GUI when booting the machine
  #   vb.gui = true
  #
  #   # Customize the amount of memory on the VM:
  #   vb.memory = "1024"
  # end
  #
  # View the documentation for the provider you are using for more
  # information on available options.

  # Enable provisioning with a shell script. Additional provisioners such as
  # Ansible, Chef, Docker, Puppet and Salt are also available. Please see the
  # documentation for more information about their specific syntax and use.
  # config.vm.provision "shell", inline: <<-SHELL
  #   apt-get update
  #   apt-get install -y apache2
  # SHELL
end
