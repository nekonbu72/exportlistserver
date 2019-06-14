set GOOS=windows
set GOARCH=386
go build -o ./dist/32/exportlistserver32.exe
set GOOS=
set GOARCH=