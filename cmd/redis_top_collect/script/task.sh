#!/bin/sh
/opt/stats/bin/collect -c /opt/stats/conf/collector.yaml -h 10.209.16.113 -p 8086 -db redis -n 5 >> /opt/stats/logs/cpu_stats.log 2>&1 &
