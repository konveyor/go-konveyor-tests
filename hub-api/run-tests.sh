#!/bin/bash
#
# Script for downloading Hub and executing its tests on running Konveyor.
#
# Usage: ./run-tests.sh <branch-name>
#
# ENV variables for HUB_BASE_URL and login/password are expected.
#

export BRANCH=$1
export HUB_TMP_DIR="/tmp/tackle2-hub-test"

# Upstream CI hack to use golang test ref input as branch for tests for releases.
if [ ! -z $GOLANG_TESTS_REF ] && [[ $GOLANG_TESTS_REF == release-* ]]; then
    echo "Using branch '${GOLANG_TESTS_REF}' for Hub API tests."
    export BRANCH=$GOLANG_TESTS_REF
fi

if [ -z $BRANCH ]; then
    echo "Assuming 'main' branch to be used for Hub."
    export BRANCH="main"
fi

if [ ! -d $HUB_TMP_DIR ]; then
    git clone https://github.com/konveyor/tackle2-hub.git $HUB_TMP_DIR
fi

cd $HUB_TMP_DIR
git fetch origin ${BRANCH}:_pr-branch
git checkout _pr-branch

make test-api
