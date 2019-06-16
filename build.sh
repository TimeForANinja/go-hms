#!/bin/bash

echo "welcome to build.sh"
echo $PWD

cd ./building

echo "preBuild.sh"
bash ./preBuild.sh

echo "generateHTMLSchema.sh"
bash ./generateHTMLSchema.sh
echo "generatePNGSchema.sh"
bash ./generatePNGSchema.sh
echo "linux.sh"
bash ./linux.sh
echo "windows.sh"
bash ./windows.sh

echo "postBuild.sh"
bash ./postBuild.sh

cd ..
