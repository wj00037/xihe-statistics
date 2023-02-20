package pgsql

import (
	"context"
	"project/xihe-statistics/infrastructure/repositories"
)

func NewGitLabRecordMapper(table GitLabRecord) repositories.GitLabRecordMapper {
	return gitlabRecord{table}
}

type gitlabRecord struct {
	table GitLabRecord
}

func (m gitlabRecord) InsertCloneCount(
	do *repositories.CloneCountDO,
) (err error) {
	col, err := m.toGitLabCloneRecord(do)
	if err != nil {
		return
	}

	f := func(ctx context.Context) error {
		return cli.create(
			ctx, m.table,
			col,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	return
}

func (m gitlabRecord) toGitLabCloneRecord(
	do *repositories.CloneCountDO,
) (GitLabRecord, error) {
	return GitLabRecord{
		Counts:   do.Counts,
		CreateAt: do.CreateAt,
	}, nil
}
