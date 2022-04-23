#!/bin/bash

python3 -m http.server --bind localhost 8080 > .server.log 2>&1 &
echo $! > .server.pid
