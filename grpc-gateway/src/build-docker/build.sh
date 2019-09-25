cd `dirname $0`
cd ../..
export GOPATH=`pwd`:$GOPATH
cd -
cd ../protos
make
cd -
go build ../main.go
mv main container/

:<<EOF
mkdir -p mysql
cd mysql
if [ ! -f mysql-community-release-el7-5.noarch.rpm ]; then
  wget http://repo.mysql.com/mysql-community-release-el7-5.noarch.rpm
fi
cd -
EOF

repo=`cat ../../Makefile |grep repo |awk -F '=' '{print $2}'`
tag=`cat ../../Makefile |grep tag |awk -F '=' '{print $2}'`

docker build --no-cache -t ${repo}/grpc-gateway:${tag} .
rm container/main