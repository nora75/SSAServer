#!/bin/bash
cd /home/NORA/go/src/SSAServer
psn=`ps agx | grep ssa_server | grep -v grep | awk '{print $1}'`
if [ -z $psn ]; then
	./ssa_server > /home/NORA/go/src/SSAServer/log/debug.log &
fi