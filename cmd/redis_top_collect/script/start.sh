
#!/bin/bash

user=`whoami`

if [ "$user" != "root" ]
then
  echo "Not permited, sudo su - root!!!"
  exit 1
fi

DIR_HOME=`pwd`

echo "* * * * * /opt/stats/bin/task.sh" >> /var/spool/cron/root
service crond restart


