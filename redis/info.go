package redis

import (
	"strings"
	"github.com/garyburd/redigo/redis"
)

func InfoWithMap(host string, port string, passward string) (map[string]string, error){
	raw, err := Info(host, port, passward)
	if err != nil {
		return nil, err
	}
	
	return transform(raw), nil
}

func Info(host string, port string, passward string) (string, error) {
	
	//conn, err := redis.Dial("tcp", "10.213.12.74:12811")
	
	option := redis.DialPassword(passward)
	conn, err := redis.Dial("tcp", host+":"+port, option)
	if err != nil {
		return "", err
	} 

	reply, err := conn.Do("info")
	if err != nil {
		return "", err
	} 
	
	infoStr, err := redis.String(reply, err)
	if err != nil {
		return "", err
	} 
	
	return infoStr, nil
}

func transform(info string) map[string]string{
	tmp := strings.Split(info, "\r\n");
	
	infoMap := make(map[string]string)
	
	for _, element := range tmp {
		s := strings.Split(element, ":");
		if len(s)==2 && s[0]!="" && s[1]!="" {
			infoMap[s[0]] = s[1]
		}
	}
	
	return infoMap
}

	
