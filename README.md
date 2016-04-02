# 37
client 反向代理 interaction,hundred,sso;
## client 是静态的web,mobile界面；
## interaction 是一款游戏；
## hundred 是一款游戏，与interaction并列的；
## sso 是用户和登陆管理的系统。
# 环境与数据库
采用golang1.5和mysql 5+, utf8格式。 nginx1.8+ 代理直接运行在linux系统。
.sh 是编译，启动文件，在启动好nginx,mysql执行即可。
