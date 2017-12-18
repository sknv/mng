package document

import "github.com/globalsign/mgo/bson"

type (
	IIdentifier interface {
		GetId() bson.ObjectId
	}

	Base struct {
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
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
