#!/bin/bash
cd /root/web/src/
rm -f ./interaction.start
rm -f ./sso.start
sleep 1
#echo "build sso and interaction files !"
export GOPATH=/root/web/
go build -i sso/sso.start.go
sleep 1
go build -i interaction/interaction.start.go
sleep 1
#
echo "sso and 37 product init ok!"



