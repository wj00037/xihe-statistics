package mongodb

import (
	"context"
	"project/xihe-statistics/infrastructure/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

func NewBigModelMapper(collection string) repositories.BigModelMapper {
	return bigModel{collection}
}

type bigModel struct {
	collectionName string
}

func (col bigModel) Add(
	b repositories.BigModelDO,
) (err error) {
	doc, err := col.toBigModelDoc(b)
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

func (col bigModel) Get(t string) (dos []repositories.BigModelDO, err error) {

	var items []BigModelRecordItem

	filter, err := typeFilter(t)
	if err != nil {
		return
	}

	f := func(ctx context.Context) error {
		return cli.filter(
			ctx, col.collectionName,
			filter, &items,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	dos = make([]repositories.BigModelDO, len(items))
	for j := range items {
		col.toBigModelDO(items[j], &dos[j])
	}

	return
}

func (col bigModel) GetAll() (dos []repositories.BigModelDO, err error) {
	var items []BigModelRecordItem

	f := func(ctx context.Context) error {
		return cli.filter(
			ctx, col.collectionName,
			bson.M{}, &items,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	dos = make([]repositories.BigModelDO, len(items))
	for j := range items {
		col.toBigModelDO(items[j], &dos[j])
	}

	return
}

func (col bigModel) toBigModelDoc(b repositories.BigModelDO) (bson.M, error) {
	docObj := BigModelRecordItem{
		UserName: b.UserName,
		BigModel: b.BigModel,
		CreateAt: b.CreateAt,
	}

	return genDoc(docObj)
}

func (col bigModel) toBigModelDO(v BigModelRecordItem, do *repositories.BigModelDO) {
	*do = repositories.BigModelDO{
		UserName: v.UserName,
		BigModel: v.BigModel,
		CreateAt: v.CreateAt,
	}
}
