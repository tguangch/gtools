package redisModule

var NilItem = Item{}

type Item struct {
	host 		string
	port 		string
	infoMap		map[string] string
}
