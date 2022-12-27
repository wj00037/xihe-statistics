package messages

type msgStatistics struct {
	Type     string `json:"type"`
	UserName string `json:"username"`
	Info     string `json:"info"`
	CreateAt int64  `json:"create_at"`
}
