package main

import (
	"database/sql/driver"
	"encoding/json"
)

type StringMap struct {
	Src   map[string]string
	Valid bool
}

func NewEmptyStringMap() *StringMap {
	return &StringMap{
		Src:   make(map[string]string),
		Valid: true,
	}
}

func NewStringMap(src map[string]string) *StringMap {
	return &StringMap{
		Src:   src,
		Valid: true,
	}
}

func (ls *StringMap) Scan(value interface{}) error {
	if value == nil {
		ls.Src, ls.Valid = make(map[string]string), false
		return nil
	}
	t := make(map[string]string)
	if e := json.Unmarshal(value.([]byte), &t); e != nil {
		return e
	}
	ls.Valid = true
	ls.Src = t
	return nil
}

func (ls *StringMap) Value() (driver.Value, error) {
	if ls == nil {
		return nil, nil
	}
	if !ls.Valid {
		return nil, nil
	}

	b, e := json.Marshal(ls.Src)
	return b, e
}
