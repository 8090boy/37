#!/bin/bash
ssopid=`pgrep sso.start`
kill -s 15 ${ssopid}
sleep 1
nohup ./sso.start > /dev/null 2>&1 &
sleep 1
hundredpid=`pgrep hundred.start`
kill -s 15 ${hundredpid}
sleep 1
nohup ./hundred.start > /dev/null 2>&1 &
#
interactionpid=`pgrep interaction.start`
kill -s 15 ${interactionpid}
sleep 1
nohup ./interaction.start > /dev/null 2>&1 &
sleep 1
ps aux | grep .start
echo "sso,200,37 ok!"

