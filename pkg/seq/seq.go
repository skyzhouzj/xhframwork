package seq

import (
	"github.com/go-redis/redis/v7"
	"github.com/xormplus/xorm"
	"sync"
)

type seq struct {
}

var Seq = new(seq)

var mutex sync.Mutex

var db *xorm.Engine

const (
	RedisTableSEQ = "SEQ"
)

//根据表单名称获取最新id
func (*seq) GetSeq(redis *redis.Client, db *xorm.Engine, tableName string) int {

	//判断是否存在这个seq
	err := redis.Exists(RedisTableSEQ + ":" + tableName)

	if err.Val() > 0 {
		id, _ := redis.Incr(RedisTableSEQ + ":" + tableName).Uint64()
		return int(id)
	} else {
		//创建SEQ
		t := new(T)

		//db.Table(tableName).Select("ID").Last(&t)
		id := int(t.Id) + 1
		redis.Set(RedisTableSEQ+":"+tableName, id, 0)
		return id
	}
}

type T struct {
	Id uint `json:"ID" gorm:"column:ID"`
}
