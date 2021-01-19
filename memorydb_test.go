package memorydb

import (
	"encoding/hex"
	"fmt"
	"testing"
)

var key = []byte("123456")
var value = []byte("456789")

func TestMemoryDB_PutGet(t *testing.T) {
	db := Single()
	db.Put(key, value)
	values, err := db.Get(key)
	if err != nil {
		panic("获取失败：" + err.Error())
	}
	t.Logf("获取成功：%s", string(values))
}

func TestMemoryDB_Has(t *testing.T) {
	db := Single()
	db.Put(key, value)
	t.Log(db.Has(key))
}

func TestMemoryDB_Delete(t *testing.T) {
	db := Single()
	db.Put(key, value)
	t.Log(db.Has(key))
	db.Delete(key)
	t.Log(db.Has(key))
}

func TestMemoryDB_GetAll(t *testing.T) {
	db := Single()
	db.Put([]byte("111111"), value)
	db.Put([]byte("222222"), value)
	db.Put([]byte("333333"), value)
	db.Put([]byte("444444"), value)

	pool := db.GetAll()
	for k, val := range pool {
		key, _ := hex.DecodeString(k)
		t.Log(fmt.Sprintf("key=>%s, val=>%s", string(key), string(val)))
	}
}
