#!/bin/bash

echo "-- remove schemacrawler"
rm -r schemacrawler-15.04.01-distribution

echo "-- remove database"
rm ./test.db
