package model

import "time"

//import "github.com/satori/go.uuid"

type Users struct {
	Id        string      `xorm:"pk UUID"`
	CreatedAt time.Time   `xorm:"created"`
	UpdatedAt time.Time   `xorm:"updated"`
	Phone     interface{} `xorm:"unique VARCHAR(16)"`
	Password  string      `xorm:"VARCHAR(16)"`
	JsonData  interface{} `xorm:"JSON"`
	Name      string      `xorm:"unique VARCHAR(32)"`
}
