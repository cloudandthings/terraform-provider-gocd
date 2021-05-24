#!/usr/bin/env bash -e

ROOT_DIR=$(pwd)/
COVERAGE_PATH=${ROOT_DIR}/coverage.txt

echo "" > ${COVERAGE_PATH}

for d in $(go list ./... | grep -v vendor | grep -v gocd-response-links); do
    go test ${TESTARGS} -v $d
    r=$?
    if [ $r -eq 1 ]; then
        # Don't exit here, run all tests
        echo "[go-test] FAIL ${d}"
    fi
    
    if [ -f profile.out ]; then
        cat profile.out >> ${COVERAGE_PATH}
        rm profile.out
    fi
done