package domain

type UserWithRepo struct {
	UserName Account `json:"username"`
	RepoName string `json:"repo_name"`
	CreateAt int64  `json:"create_at"`
}