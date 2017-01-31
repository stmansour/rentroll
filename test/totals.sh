#!/bin/bash
echo "------------------------------------------------------------------------------" >> testreport.txt
egrep ' [0-9]+[ \t]+$' testreport.txt | perl -pe "s/(.*)(.......)$/\2/" | awk '{s+=$1} END {printf( "%-20.20s %57d\n", "Total Tests: ",s)}' >> testreport.txt