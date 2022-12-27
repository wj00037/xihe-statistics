package messages

type msgStatistics struct {
	Type     string            `json:"type"`
	UserName string            `json:"username"`
	Info     map[string]string `json:"info"`
	CreateAt int64             `json:"create_at"`
}
