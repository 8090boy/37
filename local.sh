#! /bin/bash
echo "build sso and interaction..."
cd   /usr/local/gopath/src/
rm -f ./interaction.start
rm -f ./sso.start

export GOPATH=/usr/local/gopath
go build -i sso/sso.start.go
sleep 1
go build -i interaction/interaction.start.go
sleep 1
#
echo "build ok."
echo "starting..."
sopid=`pgrep ./sso.start`
kill -s 15 ${ssopid}
nohup ./sso.start > /dev/null 2>&1 &
sleep 1
#
interactionpid=`pgrep ./interaction.start`
kill -s 15 ${interactionpid}
nohup ./interaction.start > /dev/null 2>&1 &
sleep 1
ps aux | grep .start
/usr/local/nginx/sbin/nginx
/usr/local/nginx/sbin/nginx -s reload
ps aux | grep nginx
echo "=======sso and interaction start ok!======"
