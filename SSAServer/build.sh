#!/bin/bash
echo 'start build'
go build ./cmd/ssa_server
if [ $? -eq 0 ]; then
    echo -n 'start server'
    ./ssa_server > log/hoge &
fi
