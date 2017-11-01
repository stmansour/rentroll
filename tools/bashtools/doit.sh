#!/bin/bash
find . -name "*.go" -o -name "*.sh" -o -name "Makefile" -o -name "*.js" -o -name "*.csv" -o -name "*.gold" | tar -cf rr.tar -T -
