echo off
cls
SET ambiente=WINDOWS
REM
if "%ambiente%"=="LINUX" (
   del "C:\Users\Marcelo\Google Drive\Code\GoWork\bin\linux_amd64\Seguridad"
) else (
   del *.exe
   del "C:\Users\Marcelo\Google Drive\Code\GoWork\bin\Seguridad.exe"
)
del "C:\Users\Marcelo\AppData\Local\go-build\log.txt"
go clean -cache
if "%ambiente%"=="LINUX" (
   set GOARCH=amd64
   set GOOS=linux
   go install -v -a std
)
echo on
go build "github.com\Liberdina\Seguridad\Data"
go build "github.com\Liberdina\Seguridad\Services"
go build "github.com\Liberdina\Seguridad\Seguridad"
REM
go install "github.com\Liberdina\Seguridad\Data"
go install "github.com\Liberdina\Seguridad\Services"
go install "github.com\Liberdina\Seguridad\Seguridad"