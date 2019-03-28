package mongo

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/edward-kwok/go-run-api/pkg"
)

type RunnerService struct {
	collection *mgo.Collection
}

func (p *RunnerService) GetByUsername(username string) (*root.Runner, error) {
	model := RunnerModel{Id: ""}
	err := p.collection.Find(bson.M{"username": username}).One(&model)
	return model.toRootRunner(), err
}

func NewRunnerService(session *Session, dbName string, collectionName string) *RunnerService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(runnerModelIndex())
	return &RunnerService {collection}
}
