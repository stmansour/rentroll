#!/bin/bash

#  Make sure that the test machine has the same dependent
#  databases as the one(s) we're using to test with...
/usr/local/bin/mysql --no-defaults accord < accord.sql 
