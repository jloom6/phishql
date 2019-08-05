#!/usr/bin/env bash

mysql -h phishql-mysql -u root -pkingofprussia -e "GRANT ALL PRIVILEGES ON *.* TO 'wilson'@'%'"
mysql -h phishql-mysql -u root -pkingofprussia -e "DROP DATABASE IF EXISTS phish"
mysql -h phishql-mysql -u root -pkingofprussia -e "CREATE DATABASE phish"
cat init.sql | mysql -h phishql-mysql -u root -pkingofprussia phish
mysql -h phishql-mysql -u root -pkingofprussia -e "SET collation_connection = 'utf8_general_ci'"
mysql -h phishql-mysql -u root -pkingofprussia -e "ALTER DATABASE phish CHARACTER SET utf8 COLLATE utf8_general_ci"
mysql -h phishql-mysql -u root -pkingofprussia phish -e "ALTER TABLE artists CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci"
mysql -h phishql-mysql -u root -pkingofprussia phish -e "ALTER TABLE set_songs CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci"
mysql -h phishql-mysql -u root -pkingofprussia phish -e "ALTER TABLE sets CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci"
mysql -h phishql-mysql -u root -pkingofprussia phish -e "ALTER TABLE shows CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci"
mysql -h phishql-mysql -u root -pkingofprussia phish -e "ALTER TABLE songs CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci"
mysql -h phishql-mysql -u root -pkingofprussia phish -e "ALTER TABLE tags CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci"
mysql -h phishql-mysql -u root -pkingofprussia phish -e "ALTER TABLE tours CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci"
mysql -h phishql-mysql -u root -pkingofprussia phish -e "ALTER TABLE venues CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci"
