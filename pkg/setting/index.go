package setting

var LotteryQueueMap = map[string] string {
	"1": "pk10",
	"8": "pk10",
}

var QueueName = []string {"pk10"}

// 取得彩種 map
func GetQueueNameByLottery(key string) string {
	val, ok := LotteryQueueMap[key]
	if !ok {
		return LotteryQueueMap["1"]
	}
	return val
}