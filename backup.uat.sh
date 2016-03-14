#/bin/bash
cd /root/uat/src
mv ./backup/*  ./backupold
#Now=$(date +"%d_%m_%Y__%H.%M.%S")
Now=$(date +"%Y.%m%d.%H%M")
SavePath=/root/uat/src/backup
File1=${SavePath}/sso.uat.${Now}.sql
File2=${SavePath}/interaction.uat.${Now}.sql
#
mysqldump -uroot -pKi_8%lIk5 sso_uat > ${File1}
mysqldump -uroot -pKi_8%lIk5 interaction_uat  > ${File2}
echo "backup done !"
