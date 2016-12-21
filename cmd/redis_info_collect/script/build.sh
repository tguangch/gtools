#!/bin/sh

base=$1
if [ "$base" == "" ]; then
        base=stats_redis_info
fi
_pwd=`pwd`
basedir=$_pwd/$base

bin=/$basedir/bin
script=/$basedir/script
conf=/$basedir/conf
logs=/$basedir/logs

mkdir $bin -p
mkdir $script -p
mkdir $conf -p
mkdir $logs -p

##cp bin file
cd ..
go build
cp redis* $bin

## cp sh script
cp $_pwd/start.sh $script
cp $_pwd/stop.sh $script
cp $_pwd/redis_info.sh $script
chmod 755 $script/*.sh

## make zip
cd $_pwd
zip -r stats_redis_info.zip $base >> /dev/null

## clean
rm -fr $basedir