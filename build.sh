#!/bin/bash
cd /root/web/src/
rm -f ./interaction.start
rm -f ./sso.start
rm -f ./hundred.start
sleep 1
#echo "-------------------"
export GOPATH=/root/web/
go build -i sso/sso.start.go
sleep 1
go build -i interaction/interaction.start.go
sleep 1
go build -i hundred/hundred.start.go
sleep 1
#
echo "sso,200,37 go build ok!"



