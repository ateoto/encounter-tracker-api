#!/bin/bash

# This is intended to be run by Vagrant on a Ubuntu 14.04 box.

set -eo pipefail

# Install Postgres
sudo apt-get update && sudo apt-get install -y postgresql build-essential git

# Install Go
if [ ! -f /usr/local/go/bin/go ]; then
	curl -SLO https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz
	tar -C /usr/local -xzf go1.4.2.linux-amd64.tar.gz
fi

# Install encounter-tracker-api serivce
cp /vagrant/deployment/files/encounter-tracker-api.service /etc/systemd/system/encounter-tracker-api.service

# Create user and database
sudo -u postgres bash -c "psql -c \"CREATE USER etsa WITH PASSWORD 'etsa_secret';\"" || /bin/true
sudo -u postgres bash -c "psql -c \"CREATE DATABASE et WITH OWNER etsa;\"" || /bin/true

# Get go deps
mkdir -p /home/vagrant/go
echo "export GOPATH=/home/vagrant/go" >> /home/vagrant/.bashrc
echo "export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin" >> /home/vagrant/.bashrc

echo "Getting dependencies"
GOPATH=/home/vagrant/go /usr/local/go/bin/go get -u github.com/dgrijalva/jwt-go
GOPATH=/home/vagrant/go /usr/local/go/bin/go get -u github.com/labstack/echo
GOPATH=/home/vagrant/go /usr/local/go/bin/go get -u github.com/labstack/echo/middleware
GOPATH=/home/vagrant/go /usr/local/go/bin/go get -u github.com/lib/pq
GOPATH=/home/vagrant/go /usr/local/go/bin/go get -u github.com/mattes/migrate/migrate
GOPATH=/home/vagrant/go /usr/local/go/bin/go get -u golang.org/x/crypto/bcrypt

echo "Building"
pushd /vagrant/src/
	GOPATH=/home/vagrant/go /usr/local/go/bin/go build -o encounter-tracker-api
popd

mv /vagrant/src/encounter-tracker-api /usr/local/bin/encounter-tracker-api