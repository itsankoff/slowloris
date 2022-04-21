#!/bin/bash

python3 -m http.server > .server.log 2>&1 &
echo $! > .server.pid
