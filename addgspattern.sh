#!/bin/bash

# Add git secrets command to add pattern which should not be committed.
# e.g.,
# git secrets --add 'regex'
# git secrets --add --literal 'foo+bar' (Adds a string that is scanned for literally)
# Reference link: https://github.com/awslabs/git-secrets
