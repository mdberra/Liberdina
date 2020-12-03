#!/bin/bash
clear
echo "Build..."
# cd /home/marce/Desktop/Code/Go/src/github.com/Liberdina/Servina
#
# Windows build to Linux
# del *.exe
# del "/home/marce/Desktop/Code/Go/bin/linux_amd64/Server"
# del "/home/marce/Desktop/Code/Go/bin/linux_amd64/ServerServina"
# go clean -cache
# get GOARCH=amd64
# set GOOS=linux
# go install -v -a std
#
# Linux build to Windows
# go clean -cache
# set GOOS=windows
# GOARCH=amd64 \
#   CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc
#
echo "build 1"
go build "github.com/Liberdina/Servina/data"
echo "build 2"
go build "github.com/Liberdina/Servina/services"
echo "build 3"
go build "github.com/Liberdina/Servina/restful"
echo "install"
go install "github.com/Liberdina/Servina/serverServina"
ls -al /home/marce/Desktop/Code/Go/bin/serverServina
#
echo "Run"
cd /home/marce/Desktop/Code/Go/bin
./serverServina