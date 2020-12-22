#!/bin/sh
touch ~/.pgpass
cp ./postgres-account/.pgpass ~/.pgpass
chmod 0600 ~/.pgpass

psql -h 127.0.0.1 -p 5432 -U postgres -d giniaccount -f ./postgresql-data/giniaccount.sql
