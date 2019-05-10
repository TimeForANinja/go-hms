#!/bin/bash

echo "welcome to build.sh"
echo $PWD

cd ./building

echo "preBuild.sh"
./preBuild.sh

echo "generateHTMLSchema.sh"
./generateHTMLSchema.sh
echo "generatePNGSchema.sh"
./generatePNGSchema.sh
echo "linux.sh"
./linux.sh
echo "windows.sh"
./windows.sh

echo "postBuild.sh"
./postBuild.sh

cd ..
