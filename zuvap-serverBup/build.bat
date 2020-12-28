echo off
cls
SET ambiente=LINUX
REM
if "%ambiente%"=="LINUX" (
   del "C:\Users\Marcelo\Google Drive\Code\GoWork\bin\linux_amd64\Server"
   del "C:\Users\Marcelo\Google Drive\Code\GoWork\bin\linux_amd64\ServerTokens"
) else (
   del *.exe
   del "C:\Users\Marcelo\Google Drive\Code\GoWork\bin\Server.exe"
   del "C:\Users\Marcelo\Google Drive\Code\GoWork\bin\ServerTokens.exe"
)
del "C:\Users\Marcelo\AppData\Local\go-build\log.txt"
go clean -cache
if "%ambiente%"=="LINUX" (
   set GOARCH=amd64
   set GOOS=linux
   go install -v -a std
)
echo on
go build "github.com\Tokens\Data"
go build "github.com\Tokens\RestFul"
go build "github.com\Tokens\Server"
REM
go install "github.com\Tokens\Data"
go install "github.com\Tokens\RestFul"
go install "github.com\Tokens\Server"
REM
echo off
REM --------------
REM Copiar el .env
REM --------------
if "%ambiente%"=="LINUX" (
   cd "C:\Users\Marcelo\Google Drive\Code\GoWork\bin\linux_amd64"
   ren Server ServerTokens
) else (
   ren Server.exe ServerTokens.exe
)