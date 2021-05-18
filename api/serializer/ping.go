package serializer

import "ot/models"

// Ping 测试序列化器
type Ping struct {
	ID  int    `json:"id"`
	Msg string `json:"msg"`
}

//BuildPing 测试序列化器
func BuildPing(ping models.Ping) Ping {
	return Ping{
		ID:  int(ping.ID),
		Msg: ping.Msg,
	}
}
