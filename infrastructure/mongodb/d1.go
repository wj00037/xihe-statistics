package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"project/xihe-statistics/infrastructure/repositories"
)

func NewUserWithRepoMapper(name string) repositories.UserWithRepoMapper {
	return userWithRepo{name}
}

type userWithRepo struct {
	collectionName string
}

func (col userWithRepo) AddUser(
	u *repositories.UserWithRepoDO,
	p *repositories.UserDO,
) (err error) {
	doc, err := col.toUserWithRepoDoc(u, p)
	if err != nil {
		return
	}

	filter, err := typeFilter(u.Type)
	if err != nil {
		return
	}

	f := func(ctx context.Context) error {
		return cli.updateDoc(
			ctx, col.collectionName,
			filter, doc, mongoCmdSet,
		)
	}

	if err = withContext(f); err != nil && isDocNotExists(err) {
		err = repositories.NewErrorConcurrentUpdating(err)
	}

	return
}

func (col userWithRepo) Get() (do repositories.UserWithRepoDO, err error) {

}

func (col userWithRepo) toUserItem(u repositories.UserDO) UserItem {
	return UserItem{
		UserName: u.UserName,
		UpdateAt: u.UpdateAt,
	}
}

func (col userWithRepo) toUserItemArry(u []repositories.UserDO) (ui []UserItem) {
	for _, v := range u {
		ui = append(ui, col.toUserItem(v))
	}
	return
}

func (col userWithRepo) toUserWithRepoDoc(uwr *repositories.UserWithRepoDO, u repositories.UserDO) (bson.M, error) {
	users := append(uwr.Users, u)
	counts := uwr.Counts + 1
	docObj := UserWithRepoItem{
		Type:     uwr.Type,
		Users:    col.toUserItemArry(users),
		Counts:   counts,
		UpdateAt: time.Now().Unix(),
	}

	return genDoc(docObj)
}
