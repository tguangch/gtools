package main

import (
	"net"
	"flag"
	"fmt"
	"errors"
)

var cmd = `

$i: 端口
	prod(线上)域名	:	w$i.wdds.redis.com	
						r$i.wdds.redis.com
						
	sit(集成)域名	:	w$i.sit.wdds.redis.com
						r$i.sit.wdds.redis.com

	test(测试)域名	:	w$i.test.wdds.redis.com
						r$i.test.wdds.redis.com

	uat(测试)域名	:	w$i.uat.wdds.redis.com
						r$i.uat.wdds.redis.com

该程序目的：探测 prod 10300 - 10900 的redis端口对应的域名，找出uat中不存在的端口及域名

`

var (
	start = flag.Int("start", 0, "开始端口")
	end = flag.Int("end", 0, "开始端口")
)

func main(){
	
	flag.Parse()
	
	if *start <= 0 {
		panic("开始端口必须是大于0的整数")
	}
	if *end <= 0 {
		panic("结束端口必须是大于0的整数")
	}
	if *start > *end {
		panic("开始端口必须小于结束端口");
	}
	
	for i:=*start; i<=*end; i++ {
		p := fmt.Sprintf("w%d.wdds.redis.com", i)
		u := fmt.Sprintf("w%d.uat.wdds.redis.com", i)
//		pName, perr := remoteipv4(p)
//		uName ,uerr := remoteipv4(u)
		
//		fmt.Println("prod", pName, perr)
//		fmt.Println("uat", uName, uerr)
		
		_, perr := remoteipv4(p)
		_, uerr := remoteipv4(u)
		
		if perr==nil && uerr != nil {
			fmt.Println(i)
		}

	}
}

func remoteipv4(host string) (string, error) {
	ips, _ := net.LookupIP(host)
	for _, ip := range ips {
		if !ip.IsLoopback() && len(ip.To4())==4 {
			return ip.String(), nil
		}
	}
	
	return "", errors.New("invalid ip(ipv4) address")
}