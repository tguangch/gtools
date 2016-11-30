#!/bin/sh
nohup /opt/stats/bin/redis_top_collect -h 10.209.16.113 -p 8086 -d redis -n 12 >> /dev/null 2>&1 &