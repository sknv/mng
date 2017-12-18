package mng

import (
	"encoding/json"

	"github.com/globalsign/mgo/bson"
)

// ParseQuery parses given string
// and returns parsed Mongo query and an error if any raised.
func ParseQuery(str string) (bson.M, error) {
	var result bson.M
	err := parseString(str, &result)
	return result, err
}

// ParseSort parses given string
// and returns parsed slice of sorting rules and an error if any raised.
func ParseSort(str string) ([]string, error) {
	var result []string
	err := parseString(str, &result)
	return result, err
}

func parseString(str string, result interface{}) error {
	if str == "" {
		return nil
	}
	return json.Unmarshal([]byte(str), result)
}
