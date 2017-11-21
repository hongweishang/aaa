package model

//import "github.com/satori/go.uuid"

type UsernameUseridIndex struct {
	Username string `xorm:"pk VARCHAR(32)"`
	Userid   string `xorm:"not null UUID"`
}
