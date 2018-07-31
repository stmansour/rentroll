#!/bin/bash

# Remove whatever .sql files we have and checkout the versions
# that are currently checked in

while IFS='' read -r f || [[ -n "${f}" ]]; do
    if [[ ! $f =~ .*sqlschema.* ]]; then
	rm -f ${f}
	git checkout ${f}
    fi
done < dbfiles.txt
