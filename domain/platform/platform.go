package platform

type ProjectId struct {
	Id int `json:"id"`
}

type CloneTotal struct {
	Total int64 `json:"total"`
}

type CloneCounts struct {
	CloneTotal

	CreateAt int64 `json:"create_at"`
}

type PlatForm interface {
	GetProjectId(int) ([]ProjectId, error)
	GetCloneTotal(int) (CloneTotal, error)
}
