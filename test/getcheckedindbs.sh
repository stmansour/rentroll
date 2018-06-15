#!/bin/bash

# Remove whatever .sql files we have and checkout the versions
# that are currently checked in

while IFS='' read -r f || [[ -n "${f}" ]]; do
	rm -f ${f}
	git checkout ${f}
done < dbfiles.txt
