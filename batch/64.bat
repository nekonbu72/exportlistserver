set GOOS=windows
set GOARCH=amd64
go build -o ./dist/64/exportlistserver.exe
set GOOS=
set GOARCH=
robocopy static dist/64/static /is
robocopy settings dist/64/settings /is