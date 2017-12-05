package document

import "time"

type Timestampable struct {
	Base `bson:",inline"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (t *Timestampable) BeforeInsert() {
	t.Base.BeforeInsert()

	t.initTimestamp()
}

func (t *Timestampable) BeforeSave() {
	t.updateTimestamp()
}

func (t *Timestampable) initTimestamp() {
	t.CreatedAt = time.Now()
}

func (t *Timestampable) updateTimestamp() {
	t.UpdatedAt = time.Now()
}
