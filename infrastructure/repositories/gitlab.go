package repositories

import (
	"project/xihe-statistics/domain/repository"
)

type GitLabRecordMapper interface {
	InsertCloneCount(*CloneCountDO) error
}

func NewGitLabRecordRepository(mapper GitLabRecordMapper) repository.Gitlab {
	return gitlabRecord{mapper}
}

type gitlabRecord struct {
	mapper GitLabRecordMapper
}

func (impl gitlabRecord) InsertCloneCount(cc *repository.CloneCount) error {
	do := impl.toCloneCountDO(cc)
	return impl.mapper.InsertCloneCount(&do)
}

type CloneCountDO struct {
	Counts   int64
	CreateAt int64
}

func (impl gitlabRecord) toCloneCountDO(cc *repository.CloneCount) CloneCountDO {
	return CloneCountDO{
		Counts:   cc.Counts,
		CreateAt: cc.CreateAt,
	}
}
