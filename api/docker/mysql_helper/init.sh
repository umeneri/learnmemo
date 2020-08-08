#!/bin/sh

mysql --host=mysql --user=root --password=root --database=gin_test --execute='show databases;'
mysql --host=mysql --user=root --password=root --database=gin_test < init.sql
mysql --host=mysql --user=root --password=root --database=gin_test < data.sql
