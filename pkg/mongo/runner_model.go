package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/edward-kwok/go-run-api/pkg"
)

type RunnerModel struct {
	Id               bson.ObjectId `bson:"_id,omitempty"`
	OverallPosition  int32
	GenderPosition   int32
	CategoryPosition int32
	Category         string
	Name             string
	ChiName          string
	OfficalTime      int32
	ChipTime         int32
	Fivekm           int32
	FivekmToFinish   int32
}

func (u *RunnerModel) toRootRunner() *root.Runner {
	return &root.Runner{
		Id:               u.Id.Hex(),
		OverallPosition:  u.OverallPosition,
		GenderPosition:   u.GenderPosition,
		CategoryPosition: u.CategoryPosition,
		Category:         u.Category,
		Name:             u.Name,
		ChiName:          u.ChiName,
		OfficalTime:      u.OfficalTime,
		ChipTime:         u.ChipTime,
		Fivekm:           u.Fivekm,
		FivekmToFinish:   u.FivekmToFinish,
	}
}

func runnerModelIndex() mgo.Index{
	return mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

}
