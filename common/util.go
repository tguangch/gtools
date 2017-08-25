package common

import "strconv"

func AtoInt64(input string, _default int64) int64 {
	r, err := strconv.ParseInt(input, 10, 0)
	if(err == nil){
		return r
	}
	return _default
}

func AtoFloat64(input string, _default float64) float64{
	r, err := strconv.ParseFloat(input, 64)
	if(err == nil) {
		return r
	}
	return _default
}
