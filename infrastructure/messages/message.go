package messages

type msgStatistics struct {
	Type string            `json:"type"`
	User string            `json:"user"`
	Info map[string]string `json:"info"`
	When int64             `json:"when"`
}
