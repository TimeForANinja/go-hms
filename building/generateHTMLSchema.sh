#!/bin/bash

echo "-- generateHTMLSchema"
./schemacrawler-15.04.01-distribution/_schemacrawler/schemacrawler.sh -loglevel=CONFIG -server sqlite -database ./test.db -user -password -infolevel standard -command schema -outputformat htmlx -outputfile ../artifacts/sql_schema.html
