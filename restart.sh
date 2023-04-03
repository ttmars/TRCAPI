#!/bin/bash

kill -9 $(ps -ef|grep trcApi|grep -v grep|awk '{print $2}')
git pull origin master
go build -o trcApi ./main.go
chmod +x ./trcApi
nohup ./trcApi &
