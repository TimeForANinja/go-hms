#!/bin/bash

echo "-- build windows openssl"
git clone https://github.com/openssl/openssl /tmp/openssl > /dev/null
cd /tmp/openssl/
./Configure mingw64 shared --cross-compile-prefix=x86_64-w64-mingw32- --prefix=/usr/x86_64-w64-mingw32 > /dev/null
make > /dev/null
sudo make install > /dev/null
cd $GOPATH/src/github.com/timeforaninja/go-hms/building
rm -rf /tmp/openssl

echo "-- build windows"
CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o ../artifacts/hms-windows.exe ../main.go
