#!/bin/bash

sudo tuned-adm profile virtual-guest

sudo yum -y update --nogpgcheck
sudo yum -y install epel-release --nogpgcheck
sudo yum -y install ansible --nogpgcheck

cd /vagrant
ansible-playbook ./ops/vagrant.yml -i ./ops/vagrant_inventory
