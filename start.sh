#!/bin/bash
echo "--------------------------"
ssopid=`pgrep sso.start`
kill -15 ${ssopid}
sleep 1
nohup ./sso.start > /dev/null 2>&1 &
sleep 2
#
#
hundredpid=`pgrep hundred.start`
kill -15 ${hundredpid}
sleep 1
nohup ./hundred.start > /dev/null 2>&1 &
sleep 2
#
#
interPID=`pgrep interaction.s`
kill -15 ${interPID}
sleep 1
nohup ./interaction.start > /dev/null 2>&1 &
sleep 2
ps
echo "-----------sso,200,37 starting ok!------------"

