#!/bin/bash

BASE_DIR=$(cd "$(dirname "$0")/..";pwd)

cd $BASE_DIR

pid = `cat app.pid`
kill -9 $pid
