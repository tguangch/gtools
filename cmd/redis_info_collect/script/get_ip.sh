for i in `seq 10200 10900`;do
        ip=`dig +short "w$i.wdds.redis.com"`
        if [ "$ip" != "" ];then
                echo $i
        fi
done