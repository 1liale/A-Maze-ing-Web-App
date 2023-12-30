#!bin/bash

echo "Installing Java JDK 11"
yum update -y
yum install -y java-11-amazon-correcto-devel jq

echo "Installing Docker engine"
yum install docker -y
usermod -aG docker ec2-user
systemctl enable docker

echo "Installing git"
yum install -y git