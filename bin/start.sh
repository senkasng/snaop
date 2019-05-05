#!/bin/bash

BASE_DIR=$(cd "$(dirname "$0")/..";pwd)

cd $BASE_DIR

nohup bin/main &  
echo $! > app.pid
