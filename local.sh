#! /bin/bash
cd  /usr/local/gopath/src/
rm -f ./interaction.start
rm -f ./sso.start
rm -f ./hundred.start
#
ssopid=`pgrep ./sso.start`
kill -9 ${ssopid}
inpid=`pgrep ./interaction.start`
kill -9 ${inpid}
hundredpid=`pgrep ./hundred.start`
kill -9 ${hundredpid}

echo "-----------------------"
ps aux | grep .start
echo "--------build------------"
# build
export GOPATH=/usr/local/gopath
go build -i sso/sso.start.go
sleep 1
go build -i interaction/interaction.start.go
sleep 1
go build -i hundred/hundred.start.go
sleep 1
echo "--------staring-----------"
# staring
nohup ./sso.start > /dev/null 2>&1 &
nohup ./interaction.start > /dev/null 2>&1 &
nohup ./hundred.start > /dev/null 2>&1 &
#
ps aux | grep *.start
/usr/local/nginx/sbin/nginx
/usr/local/nginx/sbin/nginx -s reload
ps aux | grep nginx
echo "=======sso,interaction,hundred start ok!======"


