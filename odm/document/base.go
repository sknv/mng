package document

import "github.com/globalsign/mgo/bson"

type (
	IIdentifier interface {
		GetId() bson.ObjectId
	}

	Base struct {
		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	}
)

func (d *Base) GetId() bson.ObjectId {
	return d.Id
}

func (d *Base) BeforeInsert() {
	d.initId()
}

func (d *Base) initId() {
	if d.Id == "" {
		d.Id = bson.NewObjectId()
	}
}
