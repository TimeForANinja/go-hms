#!/bin/bash

mkdir -p ../artifacts

echo "-- creating database"
go run ./initDB.go

echo "-- prepare schemacrawler"
sudo apt install -y graphviz unzip default-jre default-jdk
unzip schemacrawler-15.06.01-distribution.zip
