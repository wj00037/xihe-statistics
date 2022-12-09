package mongodb

type UserWithRepoItem struct {
	UserName string `bson:"username" json:"username"`
	RepoName string `bson:"repo_name" json:"repo_name"`
	CreateAt int64  `bson:"create_at" json:"create_at"`
}

type BigModelRecordItem struct {
	UserName string `bson:"username" json:"username"`
	BigModel string `bson:"bigmodel" json:"bigmodel"`
	CreateAt int64  `bson:"create_at" json:"create_at"`
}
