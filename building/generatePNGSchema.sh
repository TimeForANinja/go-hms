#!/bin/bash

echo "-- generatePNGSchema"
./schemacrawler-15.06.01-distribution/_schemacrawler/schemacrawler.sh -loglevel=CONFIG -server sqlite -database ./test.db -user -password -infolevel standard -command schema -outputformat png -outputfile ../artifacts/sql_schema.png
