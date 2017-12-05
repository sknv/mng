package mng

import (
	"encoding/json"

	"github.com/globalsign/mgo/bson"
)

// ParseQuery parses given string
// and returns parsed Mongo query and an error if any raised.
func ParseQuery(str string) (bson.M, error) {
	q := bson.M{}
	if str == "" {
		return q, nil
	}

	if err := json.Unmarshal([]byte(str), &q); err != nil {
		return nil, err
	}

	return q, nil
}
