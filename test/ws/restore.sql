DROP DATABASE IF EXISTS rentroll; CREATE DATABASE rentroll; USE rentroll;
source rentrolldb.sql
GRANT ALL PRIVILEGES ON rentroll TO 'ec2-user'@'localhost' WITH GRANT OPTION;
