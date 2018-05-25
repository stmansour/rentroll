#!/bin/bash
# unlink existing symbolic link for pre-commit
# link pre-commit
RDIR=`pwd` # Assuming this file is placed at rentroll directory
GITHOOKDIR=${RDIR}/githooks
pushd .git/hooks
if [ -e "pre-commit" ]
then
    unlink pre-commit
fi
ln -s ${GITHOOKDIR}/pre-commit
popd