#!/bin/bash
pathOut=/home/marce/Documents/Liberdina/Code/Go/bin
del $pathOut/ZuvapServer
mkdir $pathOut/ZuvapServer
#
go clean -cache
set GOARCH=amd64
set GOOS=linux
go install -v -a std
#
go build "github.com/Liberdina/zuvap-server/connect"
go build "github.com/Liberdina/zuvap-server/services"
go build "github.com/Liberdina/zuvap-server/server"
go install "github.com/Liberdina/zuvap-server/server"
mv $pathOut/server $pathOut/ZuvapServer/ZuvapServer
cp server/.env $pathOut/ZuvapServer/.env
#
go install "github.com/Liberdina/zuvap-server/server/client"
mv $pathOut/client $pathOut/ZuvapServer/client
#