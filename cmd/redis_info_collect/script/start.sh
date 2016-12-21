#!/bin/bash

user=`whoami`

if [ "$user" != "root" ]
then
  echo "Not permited, sudo su - root!!!"
  exit 1
fi

##DIR_HOME=`pwd`
DIR_HOME=/opt/stats_redis_info/script

BASEDIR=$DIR_HOME/..

echo "#!/bin/sh" > $DIR_HOME/redis_info.sh
echo "nohup $DIR_HOME/../bin/redis_info_collect -host 10.209.16.113 -port 8086 -db redis -conf /opt/stats_redis_info/conf/ip.txt >> /dev/null 2>&1 &" >> $DIR_HOME/redis_info.sh

echo "* * * * * $DIR_HOME/redis_info.sh" >> /var/spool/cron/root
#service crond restart