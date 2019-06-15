set GOOS=windows
set GOARCH=386
go build -o ./dist/32/exportlistserver32.exe
set GOOS=
set GOARCH=
robocopy static dist/32/static /is
robocopy settings dist/32/settings /is