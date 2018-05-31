TOP=./..
BINDIR=${TOP}/tmp/rentroll

webclient: node_modules
	@touch fail
	if [ ! -d node_modules ]; then \
		echo "node_modules does not exits.  Will attempt to create..."; \
		sh -c "ln -s `npm root -g`"; \
	fi ;\
	if [ ! -d node_modules ]; then \
		echo "Could not make node_modules directory" ;\
		exit 1 ;\
	fi
	grunt
	jshint --extract=always ./html/*.html ./html/test/*.html ./js/elems/*.js
	@rm -f fail
	@echo "*** Grunt done ***"

clean:
	@touch fail
	grunt clean
	rm -f *log *.sql
	@rm -f fail
	@echo "*** Clean client ***"

package:
	@touch fail
	mkdir -p ${BINDIR}/webclient/
	mkdir -p ${BINDIR}/webclient/js/
	mkdir -p ${BINDIR}/webclient/html/
	cp -r ./html ${BINDIR}/webclient/
	cp ./js/*.js ${BINDIR}/webclient/js/
	cp ./js/*.css ${BINDIR}/webclient/js/
	@rm -f fail
	@echo "*** PACKAGE completed in clean ***"

node_modules:
	ln -s `npm root -g`

jshint:
	@touch fail
	jshint --extract=always ./html/*.html ./html/test/*.html ./js/elems/*.js
	@rm -f fail
	# @echo "*** jshint in client ***"

jshint-unused:
	@touch fail
	jshint --extract=always --show-non-errors ./html/*.html ./html/test/*.html ./js/elems/*.js
	@rm -f fail
	# @echo "*** jshint with --show-non-errors in client ***"


secure:
	@rm -f config.json confdev.json confprod.json
