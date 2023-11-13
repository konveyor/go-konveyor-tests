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

if [ -d $HUB_TMP_DIR ]; then
    cd $HUB_TMP_DIR
    git pull origin ${BRANCH}
else
    git clone --branch ${BRANCH} https://github.com/konveyor/tackle2-hub.git $HUB_TMP_DIR && cd $_
fi

echo "Running Hub API tests (printing FAILs only).."
set -o pipefail
make test-api | grep -B 20  FAIL
