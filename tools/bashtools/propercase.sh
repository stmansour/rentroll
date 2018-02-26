#!/bin/bash

awk '{print tolower($0);}' $1 > x

while read -r line; do
	firstletter=$(echo $line | cut -c1 | tr ‘[[:lower:]]’ ‘[[:upper:]]’)
	otherletters=$(echo $line | cut -c2-)
	echo "${firstletter}${otherletters}"
done < x