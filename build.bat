@echo on
::命令提示
echo "go run sso,hundred,interaction !"
::启动
::E:\nginx-1.8.0\nginx.exe
start /b go run sso/sso.start.go
start /b go run interaction/interaction.start.go
go run hundred/hundred.start.go

:: 杀死某个进程
:: TASKKILL /F /IM notepad.exe /IM mspaint.exe
