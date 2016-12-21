#!/bin/sh
nohup /opt/stats_redis_info/bin/redis_info_collect -host 10.209.16.113 -port 8086 -db redis -conf /opt/stats_redis_info/conf/ip.txt >> /dev/null 2>&1 &