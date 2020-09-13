package setting

// 取得彩種 map
func GetQueueNameByLottery(key string) string {
	val, ok := LotteryQueueMap[key]
	if !ok {
		return LotteryQueueMap["1"]
	}
	return val
}