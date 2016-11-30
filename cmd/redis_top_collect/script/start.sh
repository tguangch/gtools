#!/bin/bash

user=`whoami`

if [ "$user" != "root" ]
then
  echo "Not permited, sudo su - root!!!"
  exit 1
fi

DIR_HOME=`pwd`

BASEDIR=$DIR_HOME/..

echo "#!/bin/sh" > $DIR_HOME/redis_top_task.sh
echo "nohup $DIR_HOME/../bin/redis_top_collect -h 10.209.16.113 -p 8086 -d redis -n 12 >> /dev/null 2>&1 &" >> $DIR_HOME/redis_top_task.sh

echo "* * * * * $DIR_HOME/redis_top_task.sh" >> /var/spool/cron/root
service crond restart