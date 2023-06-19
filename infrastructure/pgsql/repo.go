package pgsql

import (
	"context"
	"errors"

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

	if err = m.toRepoRecordsDO(records, &r); err != nil {
		return
	}

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

func (m userWithRepo) toRepoRecordsDO(u []interface{}, do *repositories.RepoRecordsDO) error {
	users, err := toArryString(u)
	if err != nil {
		return err
	}

	*do = repositories.RepoRecordsDO{
		Users:  users,
		Counts: len(users),
	}

	return nil
}

func toArryString(ar []interface{}) ([]string, error) {
	arryString := make([]string, len(ar))

	var ok bool
	for j, v := range ar {
		if arryString[j], ok = v.(string); !ok {
			return nil, errors.New("assertion error")
		}
	}
	return arryString, nil
}
