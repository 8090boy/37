@echo on
::������ʾ
echo "go run sso,hundred,interaction !"
::����
start /b go run sso/sso.start.go
start /b go run interaction/interaction.start.go
go run hundred/hundred.start.go

:: ɱ��ĳ������
:: TASKKILL /F /IM notepad.exe /IM mspaint.exe
