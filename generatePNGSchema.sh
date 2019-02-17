#!/bin/bash

sudo apt install graphviz unzip

unzip schemacrawler-15.04.01-distribution.zip

./schemacrawler-15.04.01-distribution/_schemacrawler/schemacrawler.sh -server sqlite -database ./test.db -user -password -infolevel standard -command schema -outputformat png -outputfile ./sql_schema.png

rm -r schemacrawler-15.04.01-distribution
