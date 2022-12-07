package mongodb

type UserItem struct {
	UserName string `bson:"username" json:"username"`
	UpdateAt int64  `bson:"update_at" json:"update_at"`
}

type UserWithRepoItem struct {
	Type     string     `bson:"type" json:"type"`
	Users    []UserItem `bson:"users" json:"users"`
	Counts   int        `bson:"counts" json:"counts"`
	UpdateAt int64      `bson:"update_at" json:"update_at"`
}
