package machine

import (
	"os/exec"
	"strings"
	"strconv"
	"fmt"
	"log"
)

type TopItem struct {
	Port	int32
	Pid		int32
	user	string
	pr		string
	ni		string
	virt	string
	s		string
	Cpu		float32		//cpu load
	Mem		float32		//memory
	time	string
	command	string
}
//  PID USER      PR  NI  VIRT  RES  SHR S %CPU %MEM    TIME+  COMMAND

func Top() []TopItem{
	
	
	//1 step 获取 运行的redis  pid
             
	//	root     32342     1  0 Aug31 ?        00:42:43 /opt/redis/bin/redis-server *:10672              
	//	root     32344     1  0 Aug31 ?        00:39:01 /opt/redis/bin/redis-server *:10673
	//	root     30588 30583  0 10:06 pts/1    00:00:00 /bin/sh -c ps -ef | grep redis-server | awk '{print $2, $9}'
	//	root     30590 30588  0 10:06 pts/1    00:00:00 grep redis-server 
	//									|  |
	//									|  |
	//									\  /
	//									 \/
	//
	//								32342 *:10672
	//								32344 *:10673
	//								30588 -c
	//								30590 redis-server
	r, err := exec.Command("/bin/sh", "-c", "ps -ef | grep redis-server | awk '{print $2, $9}'").Output()
	if err != nil {
		log.Fatalln(err)
	}
	
	//2 step map[pid -> redisPort]
	ppMap := make(map[string]string)
	mArr := strings.Split(string(r), "\n") // linux下'\n', window下'\r\n'
	
	for _, m := range mArr {
		ppArr := strings.Split(m, " *:")
		if len(ppArr) == 2 {
			ppMap[ppArr[0]] = ppArr[1]
		}
	}
	
	//log.Println(ppMap)
	
	//3 step 获取top信息
	//top -b -n 1 | egrep "$(ps -ef | grep redis-server | awk '{print $2}')"
	//top -b -n 1 | egrep "port1|port2| ..."
	portConcat := ""
	for k, _ := range ppMap {
		if portConcat != "" {
			portConcat = portConcat + "|" + k
		} else {
			portConcat = portConcat + k
		}
	}
	
	//log.Println(portConcat)
	
	fr, err := exec.Command("/bin/sh", "-c", "top -b -n 1 | egrep -e '"+portConcat+"'").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	
	//log.Println(string(fr))
	
	//  PID USER      PR  NI  VIRT  RES  SHR S %CPU %MEM    TIME+  COMMAND                                                                                                   
	//	2405 root     20   0  234m 9292  816 S  2.0  0.1   7:24.46 redis-server 
	//   0    1       2    3   4     5    6  7   8    9     10      11
	
	stats := strings.Split(string(fr), "\n")
	
	items := make([]TopItem, 0, 0)
	
	for _, s := range stats {
		//log.Println(s)
		if s != "" {
			sm := strings.Fields(s)
			
			if len(sm) > 10 {
				//log.Println("sm", sm[0])
				var item TopItem
				
				iport, err := strconv.Atoi(ppMap[sm[0]])
				if err != nil {
					continue
				}
				item.Port = int32(iport)
				
				ipid, _ := strconv.Atoi(sm[0])
				item.Pid = int32(ipid)
				
				icpu, _ := strconv.ParseFloat(sm[8], 32)
				item.Cpu = float32(icpu)
				
				items = append(items, item)
			}
			
		}
	}
	
	return items
}
