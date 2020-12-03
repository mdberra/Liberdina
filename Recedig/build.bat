cls
del *.exe
REM del "C:\Users\Marcelo\Google Drive\Code\GoWork\bin\linux_amd64\Server"
REM del "C:\Users\Marcelo\Google Drive\Code\GoWork\bin\linux_amd64\ServerRecedig"
del "C:\Users\Marcelo\Google Drive\Code\GoWork\bin\Server.exe"
del "C:\Users\Marcelo\Google Drive\Code\GoWork\bin\ServerRecedig.exe"
del "C:\Users\Marcelo\AppData\Local\go-build\log.txt"
go clean -cache
REM
REM  Linux
REM set GOARCH=amd64
REM set GOOS=linux
REM go install -v -a std
REM
go build "github.com\Liberdina\Recedig\Base"
go build "github.com\Liberdina\Recedig\Data"
go build "github.com\Liberdina\Recedig\Services"
go build "github.com\Liberdina\Recedig\RestFul"
go build "github.com\Liberdina\Recedig\Server"
REM
go install "github.com\Liberdina\Recedig\Base"
go install "github.com\Liberdina\Recedig\Data"
go install "github.com\Liberdina\Recedig\Services"
go install "github.com\Liberdina\Recedig\RestFul"
go install "github.com\Liberdina\Recedig\Server"
REM
REM cd "C:\Users\Marcelo\Google Drive\Code\GoWork\bin\linux_amd64"
REM ren Server ServerRecedig
Server