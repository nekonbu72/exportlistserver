set GOOS=windows
set GOARCH=amd64
go build -o ./dist/64/exportlistserver.exe
set GOOS=
set GOARCH=