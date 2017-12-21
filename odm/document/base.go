package document

import "github.com/globalsign/mgo/bson"

type (
	IIdentifier interface {
		GetId() bson.ObjectId
	}

	BaseDocument struct {
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	}
)

func (d *BaseDocument) GetId() bson.ObjectId {
	return d.Id
}

func (d *BaseDocument) BeforeInsert() {
	d.initId()
}

func (d *BaseDocument) initId() {
	if d.Id == "" {
		d.Id = bson.NewObjectId()
	}
}
