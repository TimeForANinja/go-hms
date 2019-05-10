#!/bin/bash

echo "welcome to build.sh"
echo $PWD

cd ./building

echo "preBuild.sh"
sudo ./preBuild.sh

echo "generateHTMLSchema.sh"
sudo ./generateHTMLSchema.sh
echo "generatePNGSchema.sh"
sudo ./generatePNGSchema.sh
echo "linux.sh"
sudo ./linux.sh
echo "windows.sh"
sudo ./windows.sh

echo "postBuild.sh"
sudo ./postBuild.sh

cd ..
