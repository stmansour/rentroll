TOP=../..
RELDIR=/usr/local/accord/bin
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR="bashtools"

bashtools:
	@echo "*** ${THISDIR}: completed build ***"

package:
	@echo "*** ${THISDIR}: package build ***"

clean:
	rm -f *.json 
install:
	cp release.sh snapshot.sh rollback.sh /usr/local/accord/bin/
