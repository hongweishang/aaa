package model

import "fmt"
import "config"
import "github.com/go-xorm/xorm"
import "encoding/binary"
import "github.com/satori/go.uuid"
import "os"
import _ "github.com/lib/pq"

const (
	UsersConstModel = iota
	UsernameUserIdIndexConstModel
	PhoneUseridIndexConstModel
	SessionsConstModel
	UseridSessionidIndexConstModel
	MaxConstModel
)

var enginesMap map[int][]*xorm.Engine

func loadEngines(tbmodel int, engine *xorm.Engine) {
	enginesMap[tbmodel] = append(enginesMap[tbmodel], engine)
}

func GetEngine(tbmodel int) *xorm.Engine {
	return enginesMap[tbmodel][0]
}

func DumpEngines() {
	fmt.Println(enginesMap)
}

func ShardUuid(tbmodel int, key uuid.UUID) *xorm.Engine {
	keyUint64 := binary.LittleEndian.Uint64(key[8:])
	shardId := int((keyUint64 % uint64(config.LogicDBNum)) / uint64((config.LogicDBNum / len(enginesMap[tbmodel]))))
	return enginesMap[tbmodel][shardId]
}

func init() {
	enginesMap = make(map[int][]*xorm.Engine, 6)
	engine1, err := xorm.NewEngine("postgres", "dbname=pj user=postgres password=123456 host=127.0.0.1 port=5432 sslmode=disable")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//defer engine1.Close()
	engine1.SetMaxIdleConns(10)
	engine1.SetMaxOpenConns(30)
	engine1.ShowSQL(true)
	loadEngines(UsersConstModel, engine1)
	loadEngines(UsernameUserIdIndexConstModel, engine1)
	loadEngines(PhoneUseridIndexConstModel, engine1)
	loadEngines(SessionsConstModel, engine1)
	loadEngines(UseridSessionidIndexConstModel, engine1)
}
