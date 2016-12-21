package redis

import (
	"github.com/garyburd/redigo/redis"
)

func Config(host string, port string, passward string, oper string, param string) ([]string, error) {
	
	//conn, err := redis.Dial("tcp", "10.213.12.74:12811")
	
	option := redis.DialPassword(passward)
	conn, err := redis.Dial("tcp", host+":"+port, option)
	if err != nil {
		return nil, err
	} 

	reply, err := conn.Do("config", oper, param)
	if err != nil {
		return nil, err
	} 
	
	infoStr, err := redis.Strings(reply, err)
	if err != nil {
		return nil, err
	} 
	
	return infoStr, nil
}
