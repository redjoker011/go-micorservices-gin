#!/bin/bash
export GOOS=linux
export CGO_ENABLED=0

go build -o microservices-gin-linux-amd64;echo built `pwd`;

export GOOS=darwin

docker build -t microservices-gin:0.1.0 .
