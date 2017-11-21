package model

//import "github.com/satori/go.uuid"

type UseridSessionidIndex struct {
	Userid    string `xorm:"pk UUID"`
	Sessionid string `xorm:"not null UUID"`
}
