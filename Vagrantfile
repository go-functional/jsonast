# -*- mode: ruby -*-
# vi: set ft=ruby :

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#
# Place this Vagrantfile in your src folder and run:
#
#     vagrant up
#
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #

# modified from https://github.com/nathany/vagrant-gopher

# Vagrantfile API/syntax version.
VAGRANTFILE_API_VERSION = "2"

Vagrant.require_version ">= 1.5.0"

# location of this Vagrantfile
def gopath
  ENV['GOPATH']
end

# shell script to bootstrap Go
def bootstrap()
  install = "apt-get update -qq; apt-get install -qq -y git mercurial bzr curl"
  # See http://dl.golang.org/dl/
  archive = "go1.4.2.linux-amd64.tar.gz"

  profile = <<-PROFILE
  export GOPATH=$HOME/go
  export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
  export CDPATH=.:$GOPATH/src/github.com:$GOPATH/src/code.google.com/p:$GOPATH/src/bitbucket.org:$GOPATH/src/launchpad.net
  PROFILE

  # g++ installation stuff taken from https://gist.github.com/omnus/6404505
  <<-SCRIPT
  #{install}
  if ! [ -f /home/vagrant/#{archive} ]; then
    response=$(curl -O# https://storage.googleapis.com/golang/#{archive})
  fi
  tar -C /usr/local -xzf #{archive}
  echo '#{profile}' >> /home/vagrant/.profile

  apt-get -y install gcc
  apt-get -y install build-essential
  apt-get -y install libgflags-dev
  apt-get -y install libsnappy-dev
  apt-get -y install zlib1g-dev
  apt-get -y install libbz2-dev

  apt-get -y install docker.io
  SCRIPT
end

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.define "linux" do |linux|
    linux.vm.box = "ubuntu/trusty64"
    linux.vm.synced_folder gopath, "/home/vagrant/go"
    linux.vm.provision :shell, :inline => bootstrap()
  end

  config.vm.provider "virtualbox" do |v|
    v.cpus = 2
    v.memory = 2048
  end
end
