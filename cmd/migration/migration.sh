#!/usr/bin/env bash

mysql -h phishql-mysql -u root -pkingofprussia -e "GRANT ALL PRIVILEGES ON *.* TO 'wilson'@'%'"
mysql -h phishql-mysql -u root -pkingofprussia -e "DROP DATABASE IF EXISTS phish"
mysql -h phishql-mysql -u root -pkingofprussia -e "CREATE DATABASE phish"
cat init.sql | mysql -h phishql-mysql -u root -pkingofprussia phish
