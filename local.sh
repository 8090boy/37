#! /bin/bash
echo "build sso and interaction..."
cd   /usr/local/gopath/src/
rm -f ./interaction.start
rm -f ./sso.start

#
ssopid=`pgrep sso.start`
kill -9 ${ssopid}
sleep 5
inpid=`pgrep interaction.start`
kill -9 ${inpid}
sleep 5
echo "status--------------------"
ps aux | grep .start
echo "status--------------------"
#
#
export GOPATH=/usr/local/gopath
go build -i sso/sso.start.go
sleep 1
go build -i interaction/interaction.start.go
sleep 1


nohup ./sso.start > /dev/null 2>&1 &
nohup ./interaction.start > /dev/null 2>&1 &

ps aux | grep .start
/usr/local/nginx/sbin/nginx
/usr/local/nginx/sbin/nginx -s reload
ps aux | grep nginx
echo "=======sso and interaction start ok!======"
