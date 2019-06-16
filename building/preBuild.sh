#!/bin/bash

mkdir -p ../artifacts

echo "-- prepare"
apt update

echo "-- creating database"
go run ./initDB.go

echo "-- prepare schemacrawler"
sudo apt install -y graphviz unzip default-jre default-jdk
unzip schemacrawler-15.06.01-distribution.zip

echo "-- prepare linux build"
sudo apt install -y libssl-dev build-essential curl git

echo "-- prepare windows build"
sudo apt install -y mingw-w64

echo "-- install dependencies"
go get github.com/xeodou/go-sqlcipher
go get golang.org/x/crypto/bcrypt
