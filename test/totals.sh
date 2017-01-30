#!/bin/bash
echo "------------------------------------------------------------------------------" >> testreport.txt
egrep ' \d+\s+$' testreport.txt | perl -pe "s/(.*)(.......)$/\2/" | awk '{s+=$1} END {printf( "%-20.20s %57d\n", "Total Tests: ",s)}' >> testreport.txt