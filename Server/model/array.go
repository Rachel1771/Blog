package model

import "database/sql/driver"

import (
	"strings"
)

type Array []string

// Scan 从数据库读取数据后，对其进行处理，获得Go类型的变量
func (m *Array) Scan(val interface{}) error {
	s := val.([]uint8)
	ss := strings.Split(string(s), "|")
	*m = ss
	return nil
}

func (m Array) Value() (driver.Value, error) {
	str := strings.Join(m, "|")
	return str, nil
}
