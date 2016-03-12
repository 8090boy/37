#!/bin/bash
mv ./backup/*  ./backupold
#Now=$(date +"%d_%m_%Y__%H.%M.%S")
Now=$(date +"%Y.%m%d.%H%M")
SavePath=/root/web/src/backup
File1=${SavePath}/sso.pro.${Now}.sql
File2=${SavePath}/interaction.pro.${Now}.sql
#
mysqldump -uroot -pKi_8%lIk5 sso_pro > ${File1}
mysqldump -uroot -pKi_8%lIk5 interaction_pro  > ${File2}
echo "========Database Backup Successfully Completed !=========="
