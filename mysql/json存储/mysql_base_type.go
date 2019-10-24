package main

import (
	"database/sql/driver"
	"encoding/json"
)

/****************使gorm支持[]string结构*******************/
type Strings []string

func (c Strings) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Strings) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

/****************使gorm支持[]string结构*******************/

/****************使gorm支持[]int64结构*******************/
type Int64s []int64

func (c Int64s) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Int64s) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

/****************使gorm支持[]int64结构*******************/
