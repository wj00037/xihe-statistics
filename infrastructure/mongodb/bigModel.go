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

func (col bigModel) toBigModelDoc(b repositories.BigModelDO) (bson.M, error) {
	docObj := BigModelRecordItem{
		UserName: b.UserName,
		BigModel: b.BigModel,
		CreateAt: b.CreateAt,
	}

	return genDoc(docObj)
}
