package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Annotation struct {
	Values map[string]string `gorm:"type:LONGTEXT" json:"values"`
}

func (annotation *Annotation) Scan(val interface{}) error {
	switch val := val.(type) {
	case string:
		return json.Unmarshal([]byte(val), annotation)
	case []byte:
		return json.Unmarshal(val, annotation)
	default:
		return errors.New("not support")
	}
	return nil
}

func (annotation Annotation) Value() (driver.Value, error) {
	bytes, err := json.Marshal(annotation)
	return string(bytes), err
}
