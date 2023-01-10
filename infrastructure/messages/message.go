package messages

type msgStatistics struct {
	Type string            `json:"type"`
	User string            `json:"user"`
	Info map[string]string `json:"info"`
	When int64             `json:"when"`
}

type msgGitLab struct {
	ObjectKind string `json:"object_kind"`
	UserName   string `json:"user_name"`
	Project    `json:"project"`
	Commits    []Commits `json:"commits"`
}

type Project struct {
	Name string `json:"name"`
}

type Commits struct {
	TimeStamp string `json:"timestamp"`
}
