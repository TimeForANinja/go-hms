#!/bin/bash

cd building

./preBuild.sh

./generateHTMLSchema.sh
./generatePNGSchema.sh
./linux.sh
./windows.sh

./postBuild.sh

cd ..
