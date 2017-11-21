package model

import (
	"github.com/go-xorm/xorm"
	"testing"
	"time"
)

func Test_loadEngines(t *testing.T) {
	enginesMap = make(map[int][]*xorm.Engine, 4)
	engine, err := xorm.NewEngine("postgres", "dbname=pj user=postgres password=123456 host=127.0.0.1 port=5432 sslmode=disable")
	if err != nil {
		t.Error("fails")
	}
	defer engine.Close()
	loadEngines(UsersConstModel, engine)
	if len(enginesMap) != 1 {
		t.Error("fails")
	}
	if len(enginesMap[UsersConstModel]) != 1 {
		t.Error("fails")
	}
	engine = enginesMap[UsersConstModel][0]
	t.Log("ok")
}

func BenchmarkLoadEngines(b *testing.B) {
	b.StopTimer()
	b.Log("begin", time.Now())
	engine, _ := xorm.NewEngine("postgres", "dbname=pj user=postgres password=123456 host=127.0.0.1 port=5432 sslmode=disable")
	b.Log("end", time.Now())
	defer engine.Close()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		loadEngines(UsersConstModel, engine)
	}
}
func BenchmarkDumpEngines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DumpEngines()
	}
}
