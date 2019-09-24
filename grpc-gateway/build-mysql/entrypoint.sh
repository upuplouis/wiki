#!/bin/bash
export MYSQL_ROOT_PASSWORD=123456
bash ./setup.sh
mysqld

count=`echo "select count(*) from grpc_gateway.data" |mysql -uroot -p123456 2>/dev/null |tail -1`
if [ "$count" == "" ]; then
  bash loadDatabase.sh
fi

while :
do
  sleep 1000
done