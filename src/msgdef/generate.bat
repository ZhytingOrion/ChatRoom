echo Build gen.exe...
set GOPATH=%~dp0\..\..
go build ../zeus/net/gen

gen.exe chat.toml
gofmt.exe -w gensvr

gofmt.exe -w genclt
pause