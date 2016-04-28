#! /bin/bash
# ./build.sh dev
# ./build.sh pro
clearStart(){
rm -f ./interaction.start
rm -f ./sso.start
rm -f ./hundred.start
}
#
stop(){
ssoPID=`pgrep sso.start`
interPID=`pgrep interaction.st`
hundredPID=`pgrep hundred.st`
kill -15 ${ssoPID}
kill -15 ${interPID}
kill -15 ${hundredPID}
}
#
buildItem(){
export GOPATH=${1}
go build -i sso/sso.start.go
sleep 1
go build -i interaction/interaction.start.go
sleep 1
go build -i hundred/hundred.start.go
sleep 1
}
#
starting(){
nohup ./sso.start > /dev/null 2>&1 &
nohup ./hundred.start > /dev/null 2>&1 &
nohup ./interaction.start > /dev/null 2>&1 &
echo "--------------------------"
ps
echo "-----------sso,200,37 starting ok!------------"
}
#
backup(){
mv ./backup/*  ./backupold
#Now=$(date +"%d_%m_%Y__%H.%M.%S")
Now=$(date +"%Y.%m%d.%H%M")
SavePath=/root/web/src/backup
File1=${SavePath}/sso.pro.${Now}.sql
File2=${SavePath}/hundred.pro.${Now}.sql
File3=${SavePath}/interaction.pro.${Now}.sql
mysqldump -uroot -pKi_8%lIk5 sso_pro > ${File1}
mysqldump -uroot -pKi_8%lIk5 hundred_pro  > ${File2}
mysqldump -uroot -pKi_8%lIk5 interaction_pro  > ${File3}
echo "========Database Backup Successfully Completed ! =========="
}
#
nginxreload(){
	`${1}/sbin/nginx`
	`${1}/sbin/nginx -s reload`
}
###########
###########
#
local_env(){
gopath="/usr/local/gopath"
nginxdir="/usr/local/nginx"
clearStart
#stop
killall *.start
buildItem ${gopath}
starting
nginxreload ${nginxdir}
}
#
product_env(){
gopath="/root/web"
nginxdir="/usr/local/nginx"
backup
clearStart
#stop
killall *.start
buildItem ${gopath}
starting
nginxreload ${nginxdir}
}

########################
######## START ########
########################
envParam=${1}
isBack=""

if test ${envParam} == ${isBack}
then
exit 0
else
echo "..."
fi

case ${envParam} in
"dev")
gopath="/usr/local/gopath"
cd  ${gopath}/src
local_env
;;
"pro")
gopath="/root/web"
cd  ${gopath}/src
product_env
;;
*)
echo "cannot use '${envParam}' match env !"
esac

