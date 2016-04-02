#! /bin/bash
gopath="/usr/local/gopath"
nginxdir="/usr/local/nginx"


localbuild(){
cd  ${gopath}/src
rm -f ./interaction.start
rm -f ./sso.start
rm -f ./hundred.start
#
echo "-------kill current process-----------"
ssopid=`pgrep sso.start`
kill -15 ${ssopid}
sleep 2
inpid=`pgrep interaction.sta`
kill -15 ${inpid}
sleep 2
hundredpid=`pgrep hundred.start`
kill -15 ${hundredpid}
sleep 2
#
ps
echo "--------build  sso,hundred,interaction------------"
# build
export GOPATH=${gopath}
go build -i sso/sso.start.go
sleep 2
go build -i interaction/interaction.start.go
sleep 2
go build -i hundred/hundred.start.go
sleep 2
echo "--------------staring------------------"
# staring
nohup ./sso.start > /dev/null 2>&1 &
nohup ./interaction.start > /dev/null 2>&1 &
nohup ./hundred.start > /dev/null 2>&1 &
#
ps aux | grep *.start
#/usr/local/nginx/sbin/nginx
${nginxdir}/sbin/nginx -s reload
sleep 3
ps
echo "=======sso,interaction,hundred start ok!======"
}

localbuild
