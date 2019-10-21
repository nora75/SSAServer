#!/bin/bash
cd /home/NORA/go/src/SSAServer
go build ./cmd/ssa_server
psn=`ps agx | grep ssa_server | grep -v grep | awk '{print $1}'`
if [ ! -z $psn ]; then
	kill $psn
fi
./ssa_server > /home/NORA/go/src/SSAServer/log/debug.log &
