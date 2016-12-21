#!/bin/bash


user=`whoami`

if [ "$user" != "root" ]
then
  echo "Not permited, sudo su - root!!!"
  exit 1
fi

sed -i '/redis_info.sh/d' /var/spool/cron/root
#service crond restart