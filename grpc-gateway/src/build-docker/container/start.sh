#!/bin/bash
mv config.ini.tpl config.ini

echo ${TEST} |sed "s/\(\(\/\)\|\(\?\)\|\(\&\)\)/\\\\\1/g" |sed "s/\(\/\)\|\(\?\)\|\(\&\)\)/\\\\\1/g" |xargs -I {} sed "s/TEST/{}/g" -i config.ini

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