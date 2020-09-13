package setting

var LotteryMap = map[string] string {
	"1": "pk10",
	"8": "pk10",
}

// 取得彩種 map
func GetLotteryMap(key string) string {
	val, ok := LotteryMap[key]
	if !ok {
		return LotteryMap["1"]
	}
	return val
}