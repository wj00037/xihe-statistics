package messages

type msgBigModelRecord struct {
	Type     string `json:"type"`
	UserName string `json:"username"`
	BigModel string `json:"bigmodel"`
	CreateAt int64  `json:"create_at"`
}

type msgRepoRecord struct {
	Type     string `json:"type"`
	UserName string `json:"username"`
	RepoName string `json:"repo_name"`
	CreateAt int64  `json:"create_at"`
}
