#!/bin/bash
FTYPES=("*.go" "*.js" "*.sh" "*.csv" "Makefile" "*.html")
DIR1=../..
DIR2=../../../rentroll.new

dodiff() {
    find ${DIR1} -name "${1}" | while read -r fname ; do
    	fname2=${fname/$DIR1/$DIR2}
        if [ ! -f ${fname2} ]; then
            echo "missing: ${fname2}"
        else
            UDIFFS=$(diff ${fname} ${fname2} | wc -l)
            if [ ${UDIFFS} -gt 0 ]; then
                echo "meld  ${fname}   ${fname2}"
            fi
        fi
    done
    find ${DIR2} -name "${1}" | while read -r fname ; do
    	fname2=${fname/$DIR2/$DIR1}
        if [ ! -f ${fname2} ]; then
            echo "additional: ${fname}"
        fi
    done
}

for i in "${FTYPES[@]}"
do
   dodiff "${i}"
done
