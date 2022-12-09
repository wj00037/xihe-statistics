package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"project/xihe-statistics/infrastructure/repositories"
)

func NewUserWithRepoMapper(name string) repositories.UserWithRepoMapper {
	return userWithRepo{name}
}

type userWithRepo struct {
	collectionName string
}

func (col userWithRepo) Add(
	u repositories.UserWithRepoDO,
) (err error) {
	doc, err := col.toUserWithRepoDoc(&u)
	if err != nil {
		return
	}

	f := func(ctx context.Context) error {
		return cli.createDoc(
			ctx, col.collectionName,
			doc,
		)
	}

	if err = withContext(f); err != nil {
		return err
	}

	return
}

func (col userWithRepo) toUserWithRepoDoc(u *repositories.UserWithRepoDO) (bson.M, error) {
	docObj := UserWithRepoItem{
		UserName: u.UserName,
		RepoName: u.RepoName,
		CreateAt: u.CreateAt,
	}

	return genDoc(docObj)
}
