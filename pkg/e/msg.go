package e

var MsgFiles = map[int]string {
	SUCCESS: "成功",
	ERROR: "錯誤",
	NOT_FOUND: "找不到資料",
}

func GetMsg(code int) string {
	msg, ok := MsgFiles[code]
	if ok {
		return msg
	}
	return MsgFiles[ERROR]
}
