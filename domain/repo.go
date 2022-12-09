package domain

type UserWithRepo struct {
	UserName string `json:"username"`
	RepoName string `json:"repo_name"`
	CreateAt int64  `json:"create_at"`
}