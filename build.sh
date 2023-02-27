#!/bin/sh

sqlite3 ./test.db
sqlite3 -init ./sql/create_db.sql

echo "init success"
