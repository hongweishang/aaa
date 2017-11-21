package model

import "time"

//import "github.com/satori/go.uuid"

type Sessions struct {
	Id        string    `xorm:"pk UUID"`
	Userid    string    `xorm:"not null unique UUID"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	ExpiredAt time.Time
}
