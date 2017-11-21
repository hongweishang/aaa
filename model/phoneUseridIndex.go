package model

//import "github.com/satori/go.uuid"

type PhoneUseridIndex struct {
	Phone  string `xorm:"pk VARCHAR(16)"`
	Userid string `xorm:"not null UUID"`
}
