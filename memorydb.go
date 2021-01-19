package memorydb

import (
	"encoding/hex"
	"errors"
	"sync"
)

// 结构体
type MemoryDB struct {
	lock sync.Mutex
	db   map[string][]byte
}

var (
	memoryDB *MemoryDB
	memoryOnce sync.Once
)

// Single 单例模式初始化数据库
func Single() *MemoryDB {
	memoryOnce.Do(func() {
		memoryDB = NewMemDB()
	})
	return memoryDB
}

// NewMemDB 初始化数据库
func NewMemDB() *MemoryDB {
	return &MemoryDB{
		db: make(map[string][]byte),
	}
}

// Put 写方法
func (db *MemoryDB) Put(key []byte, value []byte) {
	db.lock.Lock()
	db.db[hex.EncodeToString(key)] = value
	db.lock.Unlock()
}

// Get 读方法
func (db *MemoryDB) Get(key []byte) ([]byte, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	if value, ok := db.db[hex.EncodeToString(key)]; ok {
		return value, nil
	}
	return nil, errors.New("key is not found")
}

// Has 判断是否存在
func (db *MemoryDB) Has(key []byte) bool {
	db.lock.Lock()
	defer db.lock.Unlock()

	if _, ok := db.db[hex.EncodeToString(key)]; ok {
		return true
	}
	return false
}

// Delete 删除制定键值
func (db *MemoryDB) Delete(key []byte) {
	db.lock.Lock()
	delete(db.db, hex.EncodeToString(key))
	db.lock.Unlock()
}

// GetAll 获取所有键
func (db *MemoryDB) GetAll() map[string][]byte {
	db.lock.Lock()
	defer db.lock.Unlock()
	return db.db
}

