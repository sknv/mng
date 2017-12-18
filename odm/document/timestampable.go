package document

import "time"

type Timestampable struct {
	Base `bson:",inline"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
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
