#!/bin/bash

echo "-- generateHTMLSchema"
./schemacrawler-15.06.01-distribution/_schemacrawler/schemacrawler.sh -server sqlite -database ./test.db -user -password -infolevel standard -command schema -outputformat htmlx -outputfile ../artifacts/sql_schema.html
