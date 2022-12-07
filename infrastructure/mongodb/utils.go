package mongodb

import (
	"context"
	"encoding/json"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	mongoCmdAll         = "$all"
	mongoCmdSet         = "$set"
	mongoCmdInc         = "$inc"
	mongoCmdPush        = "$push"
	mongoCmdPull        = "$pull"
	mongoCmdMatch       = "$match"
	mongoCmdFilter      = "$filter"
	mongoCmdProject     = "$project"
	mongoCmdAddToSet    = "$addToSet"
	mongoCmdElemMatch   = "$elemMatch"
	mongoCmdSetOnInsert = "$setOnInsert"
)

var errDocNotExists = errors.New("doc doesn't exist")

func isDocNotExists(err error) bool {
	return errors.Is(err, errDocNotExists)
}

func genDoc(doc interface{}) (m bson.M, err error) {
	v, err := json.Marshal(doc)
	if err != nil {
		return
	}

	err = json.Unmarshal(v, &m)

	return
}

func typeFilter(statisticsType string) (bson.M, error) {
	return bson.M{
		"type": statisticsType,
	}, nil
}

func doubleTypeFilter(statisticsType string, bigmodelType string) (bson.M, error) {
	return bson.M{
		"type":          statisticsType,
		"bigmodel_type": bigmodelType,
	}, nil
}

func (cli *client) updateDoc(
	ctx context.Context, collection string,
	filterOfDoc, update bson.M, op string,
) error {
	r, err := cli.collection(collection).UpdateOne(
		ctx, filterOfDoc,
		bson.M{
			op: update,
		},
	)

	if err != nil {
		return err
	}

	if r.MatchedCount == 0 {
		return errDocNotExists
	}

	return nil
}
