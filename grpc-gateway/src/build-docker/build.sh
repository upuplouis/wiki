cd `dirname $0`
cd ../..
export GOPATH=`pwd`:$GOPATH
cd -
cd ../protos
make
cd -
go build ../main.go

repo=`cat ../../Makefile |grep repo |awk -F '=' '{print $2}'`
tag=`cat ../../Makefile |grep tag |awk -F '=' '{print $2}'`

mv main container/
docker build --no-cache -t ${repo}/grpc-gateway:${tag} .
rm container/main