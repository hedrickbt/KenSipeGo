from https://github.com/kensipe/go-labs/

ctrl+shift+a - increase font size

Setting up Powershell environment
$env:Path += ";C:\go\sdk\go1.18.2\bin;C:\go\bin"
$env:GOROOT="C:\go\sdk\go1.18.2"
$env:GOPATH="D:\Data\SpiderOak\projects\KenSipeGo"
$env:GOBIN = "C:\go\bin"
D:
cd D:\Data\SpiderOak\projects\KenSipeGo\src\github.com\hedrickbt\go-labs\wman
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.1