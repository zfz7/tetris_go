#!/bin/bash

# Require script to be run as root (or with sudo)
function super-user-check() {
  if [ "$EUID" -ne 0 ]; then
    echo "You need to run this script as super user."
    exit
  fi
}

# Check for root
super-user-check

#Update System
apt update
apt upgrade -y

#Install Docker
apt install apt-transport-https ca-certificates curl software-properties-common -y
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable"
apt update
apt-cache policy docker-ce
apt install docker-ce -y

#Install Docker-compose
if [ "$1" == 'DO' ]; then
	DOCKER_CONFIG=/root/.docker
else
	DOCKER_CONFIG=/home/ubuntu/.docker
	groupadd docker
	usermod -aG docker ubuntu
fi
mkdir -p $DOCKER_CONFIG/cli-plugins
curl -SL https://github.com/docker/compose/releases/download/v2.10.2/docker-compose-linux-x86_64 -o $DOCKER_CONFIG/cli-plugins/docker-compose
chmod +x $DOCKER_CONFIG/cli-plugins/docker-compose

sudo chmod 666 /var/run/docker.sock

#Install certbot
snap install core
snap refresh core
snap install --classic certbot

#Update and remove  System
apt update
apt upgrade -y
apt autoremove -y