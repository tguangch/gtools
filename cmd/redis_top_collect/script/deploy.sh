#!/bin/sh

##clean
function clean()
{
	stop $1
	ssh $1 sudo rm -fr /opt/stats >>/dev/null 2>>./error.log
}

function install()
{
	scp ./stats.zip tangguangchao@$1:/home/tangguangchao  >>/dev/null 2>>./error.log
	ssh $1 sudo unzip -o stats.zip -d /opt >>/dev/null 2>>./error.log
}

function start()
{
	ssh $1 sudo /opt/stats/script/start.sh >>/dev/null 2>>./error.log
}

function stop()
{
	ssh $1 sudo /opt/stats/script/stop.sh >>/dev/null 2>>./error.log
}

function all()
{
        clean $1
        install $1
        start $1
}

function deployAll()
{
for ip in `cat ips.txt`
do
	deploy $ip $1
done
}

function deploy()
{	
        echo $ip >> ./error.log
        if [ "$2" = "stop" ]; then
                stop $1
        elif [ "$2" = "clean" ]; then
                clean $1
        elif [ "$2" = "start" ]; then
                start $1
        elif [ "$2" = "install" ]; then
                install $1
        else
                all $1
        fi
}

flag=$1 ##all, clean
ip=$2

if [ "$ip" = "" ]; then
	deployAll $flag
else 
	deploy $ip $flag
fi


