#!/usr/bin/env bash

if [  $# != 2 ]
then
    echo "Pass the bosh package name and the golang package name"
    exit 1
fi

BOSH_PACKAGE_NAME=$1
GOLANG_PACKAGE_NAME=$2

pushd $(dirname $0)
  export GOPATH=$(pwd)/..
  export SPEC_FILE=$(pwd)/../packages/$BOSH_PACKAGE_NAME/spec

  echo "---
# This file was autogenerated using ./scripts/make_all_golang_project_spec_files.sh
# dont modify this file, run that command instead, DO IT!
name: $BOSH_PACKAGE_NAME

dependencies:
- backup-and-restore-release-golang

files:
- github.com/cloudfoundry-incubator/$GOLANG_PACKAGE_NAME/cmd/$GOLANG_PACKAGE_NAME/*.go" > $SPEC_FILE

  pushd ../src/github.com/cloudfoundry-incubator/$GOLANG_PACKAGE_NAME/cmd/$GOLANG_PACKAGE_NAME


  go list -f {{.Deps}} | tr -d '[]' | tr -s ' ' "\n" | \
    grep "github.com/cloudfoundry-incubator/$GOLANG_PACKAGE_NAME" | \
    xargs -IN echo - N/*.go >> $SPEC_FILE
  popd
popd