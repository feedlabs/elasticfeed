#!/bin/sh

GOPATH_ENV=$1
SOURCE_DIR=$2

export GOPATH=$GOPATH_ENV

# create symlink of project into GOPATH
FEEDLABS_PACKAGE_DIR=$GOPATH/src/github.com/feedlabs
FEEDLABS_FEEDIFY_PATH=$FEEDLABS_PACKAGE_DIR/feedify
FEEDLABS_API_PATH=$FEEDLABS_PACKAGE_DIR/api

if [ ! -e $FEEDLABS_FEEDIFY_PATH ]; then
  mkdir -p $FEEDLABS_PACKAGE_DIR
  ln -s $SOURCE_DIR/../feedify $FEEDLABS_FEEDIFY_PATH
fi

if [ ! -e $FEEDLABS_API_PATH ]; then
  mkdir -p $FEEDLABS_PACKAGE_DIR
  ln -s $SOURCE_DIR $FEEDLABS_API_PATH
fi

# install dependencies
# go get github.com/bitly/go-nsq

# redis client
go get github.com/fzzy/radix/redis

# logger
go get github.com/barakmich/glog

# neo4j graph database client
go get github.com/jmcvetta/neoism

# javascript interpreter: scenario-engine ?
go get github.com/robertkrimen/otto

# memcache
go get github.com/bradfitz/gomemcache/memcache

# websoscket
go get github.com/gorilla/websocket

# sseserver
go get github.com/mroth/sseserver
