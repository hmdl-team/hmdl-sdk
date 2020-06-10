package sdk

import (
	"bytes"
	"database/sql/driver"
	"errors"
)

type JSON []byte

func (m JSON) Value() (driver.Value, error) {
	if m.IsNull() {
		return nil, nil
	}
	return string(m), nil
}

func (m *JSON) Scan(value interface{}) error {
	if value == nil {
		*m = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		errors.New("Invalid Scan Source")
	}
	*m = append((*m)[0:0], s...)
	return nil
}

func (m JSON) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

func (m *JSON) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("null point exception")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

func (m JSON) IsNull() bool {
	return len(m) == 0 || string(m) == "null"
}

func (m JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(m), []byte(j1))
}
