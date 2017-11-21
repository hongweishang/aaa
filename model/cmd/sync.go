package main

import _ "github.com/lib/pq"
import "github.com/go-xorm/xorm"
import "fmt"
import "model"

func main() {
	engine, err := xorm.NewEngine("postgres", "user=postgres password=123456 host=127.0.0.1 port=5432 dbname=pj sslmode=disable")
	defer engine.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("engine ok")
	}
	err = engine.Sync2(new(model.UseridSessionidIndex), new(model.Users), new(model.UsernameUseridIndex), new(model.PhoneUseridIndex), new(model.Sessions))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sync2 ok")
	}
}
