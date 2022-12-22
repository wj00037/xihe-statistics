package pgsql

import (
	"context"

	"project/xihe-statistics/infrastructure/repositories"
)

func NewUserWithRepoMapper(table UserWithRepo) repositories.UserWithRepoMapper {
	return userWithRepo{table}
}

type userWithRepo struct {
	table UserWithRepo
}

func (m userWithRepo) Add(
	u repositories.UserWithRepoDO,
) (err error) {
	col, err := m.toUserWithRepoCol(u)
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
		return err
	}

	return
}

func (m userWithRepo) Get() (
	r repositories.RepoRecordsDO,
	err error,
) {

	var records []interface{}

	f := func(ctx context.Context) error {
		err := cli.distinct(
			ctx, m.table,
			"username", &records,
		)

		return err
	}

	if err = withContext(f); err != nil {
		return
	}

	m.toRepoRecordsDO(records, &r)

	return
}

func (m userWithRepo) toUserWithRepoCol(u repositories.UserWithRepoDO) (UserWithRepo, error) {
	colObj := UserWithRepo{
		UserName: u.UserName,
		RepoName: u.RepoName,
		CreateAt: u.CreateAt,
	}

	return colObj, nil
}

func (m userWithRepo) toRepoRecordsDO(u []interface{}, do *repositories.RepoRecordsDO) {
	users := toArryString(u)
	*do = repositories.RepoRecordsDO{
		Users:  users,
		Counts: len(users),
	}
}

func toArryString(ar []interface{}) []string {
	arryString := make([]string, len(ar))

	for j, v := range ar {
		arryString[j] = v.(string)
	}
	return arryString
}
