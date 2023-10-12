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

make test-api
