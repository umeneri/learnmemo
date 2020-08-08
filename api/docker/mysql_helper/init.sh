#!/bin/sh

mysql --host=mysql --user=root --password=root --port=3307 --database=gin_test < init.sql
mysql --host=mysql --user=root --password=root --port=3307 --database=gin_test < data.sql
