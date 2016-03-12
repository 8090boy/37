# interaction
my 37
[database]
url = tcp(127.0.0.1:3306)/interaction_dev?charset=utf8
dbtype = mysql
password = root
name = root

[sysinfo]
protocol = http://
ip = 192.168.1.2:100
url = http://192.168.1.2:100
serverName = interaction.kingbloc.com

[37client]
protocol = http://
serverName = 192.168.1.2
index = /index.html
37main = /37.html

[sso]
protocol = http://
serverName = 192.168.1.2:33
byToken = /byToken
loginPath = /login

[common]
initializedDatabase = true

[37conf]
dismissal = 7
mulriple = 3
