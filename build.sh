#!/bin/bash



set -x 

TAG=${GITHUB_REF##*/}
Machine=`uname -m`

BINARY=${GITHUB_REPOSITORY##*/}
# BINARY=dnsx

CGO_ENABLED=0 GOOS=darwin go build -v -o bin/${BINARY}_${TAG}_Darwin_${Machine} .
CGO_ENABLED=0 GOOS=linux go build -v -o bin/${BINARY}_${TAG}_Linux_${Machine} .



