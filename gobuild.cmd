@echo off

set MODULE=s3service.ext

go build -ldflags "-H=windowsgui -s -w -extldflags -static" -a -i -o %MODULE% .
::go build -ldflags "-s -w -extldflags -static" -a -i -o %MODULE% .
