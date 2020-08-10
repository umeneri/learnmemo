#!/bin/sh

mysql -h 127.0.0.1 -uroot -proot -P 3306 gin < ./init.sql
mysql -h 127.0.0.1 -uroot -proot -P 3306 gin < ./data.sql
mysql -h 127.0.0.1 -uroot -proot -P 3306 gin_test < ./init.sql
mysql -h 127.0.0.1 -uroot -proot -P 3306 gin_test < ./data.sql
mysql -h 127.0.0.1 -uroot -proot -P 3307 gin_test < ./init.sql
mysql -h 127.0.0.1 -uroot -proot -P 3307 gin_test < ./data.sql
