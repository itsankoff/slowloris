#!/bin/bash

echo $$ > .server.pid
python3 -m http.server
