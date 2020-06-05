package sdk

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Array []interface{}

//Value json Marshal to byte
func (a Array) Value() (driver.Value, error) {
	bytes, err := json.Marshal(a)
	return string(bytes), err
}
func (a Array) ArrayString() ([]string, error) {
	var arr []string
	bytes, err := json.Marshal(a)
	err = json.Unmarshal(bytes, &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil

}
//Scan string or byte Unmarshal to json
func (a *Array) Scan(src interface{}) error {
	switch value := src.(type) {
	case string:
		return json.Unmarshal([]byte(value), a)
	case []byte:
		return json.Unmarshal(value, a)
	default:
		return errors.New("Not support!")
	}
}