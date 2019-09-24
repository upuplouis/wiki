#!/bin/bash

echo `date +[%x"-"%T]`": init data"
cd sql
for f in `ls`
do
mysql -uroot -p123456 < $f
done