#!/bin/bash
ssopid=`pgrep sso.start`
kill -s 15 ${ssopid}
sleep 1
nohup ./sso.start > /dev/null 2>&1 &
sleep 1
#
interactionpid=`pgrep interaction.start`
kill -s 15 ${interactionpid}
sleep 1
nohup ./interaction.start > /dev/null 2>&1 &
sleep 1
ps aux | grep .start
echo "sso and 37 product init ok!"

