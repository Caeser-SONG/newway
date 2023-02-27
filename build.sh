#!/bin/sh

sqlite3 test.db < ./sql/create_db.sql

echo "init success"
