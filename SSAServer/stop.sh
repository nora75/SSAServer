#!/bin/bash
cd /home/NORA/go/src/SSAServer
psn=`ps agx | grep ssa_server | grep -v grep | awk '{print $1}'`
if [ ! -z $psn ]; then
	kill $psn
	echo "Stop the server"
else
	echo "The server isn't running."
fi
