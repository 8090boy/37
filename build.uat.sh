#!/bin/bash
cd /root/uat/src/
rm -f ./interaction.uat
rm -f ./sso.uat
#
export GOPATH=/root/uat/
go build -i sso/sso.start.go
go build -i interaction/interaction.start.go
#
mv -f ./sso.start ./sso.uat
mv -f ./interaction.start ./interaction.uat
ssopid=`pgrep sso.uat`
kill -s 15 ${ssopid}
interactionpid=`pgrep interaction.uat`
kill -s 15 ${interactionpid}
# start sso.start
# start interaction.start
nohup ./sso.uat > /dev/null 2>&1 &
sleep 1
nohup ./interaction.uat > /dev/null 2>&1 &
sleep 1
ps aux | grep .uat
echo "uat init ok!"

