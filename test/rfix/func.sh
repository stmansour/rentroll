#!/bin/bash

ssh ec2-user@dir3 '/usr/bin/mysqldump -h phbk.cjkdwqbdvxyu.us-east-1.rds.amazonaws.com -P 3306 receipts > receipts.sql'
scp -i ~/.ssh/smanAWS1.pem dir3:~/receipts.sql .
mysql --no-defaults rentroll <receipts.sql
./rfix
mysqldump --no-defaults rentroll >rcptfixed.sql
scp -i ~/.ssh/smanAWS1.pem rcptfixed.sql dir3:~/rcptfixed.sql
