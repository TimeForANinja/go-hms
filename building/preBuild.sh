#!/bin/bash

mkdir -p ../artifacts

# echo "-- prepare apt"
# sudo apt update > /dev/null

echo "-- creating database"
go run ./initDB.go

echo "-- prepare schemacrawler"
sudo apt install -y graphviz unzip default-jre default-jdk > /dev/null
unzip schemacrawler-15.06.01-distribution.zip > /dev/null

echo "-- prepare linux build"
sudo apt install -y libssl-dev build-essential curl git > /dev/null

echo "-- prepare windows build"
sudo apt install -y mingw-w64 > /dev/null

echo "-- install dependencies"
go get github.com/xeodou/go-sqlcipher
go get golang.org/x/crypto/bcrypt
