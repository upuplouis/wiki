#!/bin/bash
cd `dirname $0`

repo=`cat ../Makefile |grep repo |awk -F '=' '{print $2}'`
tag=`cat ../Makefile |grep tag |awk -F '=' '{print $2}'`

docker build --no-cache -t ${repo}/grpc-gateway-mysql:${tag} .