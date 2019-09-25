#!/bin/bash
mv config.ini.tpl config.ini

echo ${TEST} |sed "s/\(\(\/\)\|\(\?\)\|\(\&\)\)/\\\\\1/g" |sed "s/\(\/\)\|\(\?\)\|\(\&\)\)/\\\\\1/g" |xargs -I {} sed "s/TEST/{}/g" -i config.ini

:<<EOF
if [ ! -d /root/logs ]; then
  mkdir -p /root/logs
fi

if [ -d /root/mysql ]; then
  yum install -y /root/mysql/*.rpm
  nohup mysqld --user=root > /root/logs/mysql.log 2>&1 &
  sleep 10
  echo "set password=password('123456');" |mysql -uroot -p`cat /root/.mysql_secret |grep "#" |awk -F ' ' '{print $NF}'` --connect-expired-password -b
  mysql -uroot -p123456</root/grpc_gateway.sql
  rm -rf /root/mysql
  rm -f /root/grpc_gateway.sql
else
  nohup mysqld --user=root > /root/logs/mysql.log 2>&1 &
  sleep 10
fi
EOF

while true
do
  curl 127.0.0.1:3306
  if [ "$?" = "0" ]
  then
    echo `date +[%x"-"%T]`: Mysql is online
    break
  else
    echo `date +[%x"-"%T]`: Waiting mysql start
    sleep 1
  fi
done

while true; do
    ./main
    echo `date +[%x"-"%T]`: service down. restart 10 seconds later
    sleep 10
done