package repository

type CloneCount struct {
	Counts   int64
	CreateAt int64
}

type Gitlab interface {
	InsertCloneCount(*CloneCount) error
}
