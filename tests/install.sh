#!/bin/sh

GOPATH_ENV=$1
SOURCE_DIR=$2

export GOPATH=$GOPATH_ENV

# create symlink of project into GOPATH
FEEDLABS_PACKAGE_DIR=$GOPATH/src/github.com/feedlabs
FEEDLABS_FEEDIFY_PATH=$FEEDLABS_PACKAGE_DIR/feedify

if [ ! -e $FEEDLABS_FEEDIFY_PATH ]; then
  mkdir -p $FEEDLABS_PACKAGE_DIR
  ln -s $SOURCE_DIR/.. $FEEDLABS_FEEDIFY_PATH
fi

# install dependencies
# go get github.com/bitly/go-nsq
