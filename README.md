
# storage-memorydb

基于内存的key-value数据存储

## 目录
 序号 | Go文件/函数或方法 | 作用 
---|---|---
 1 | memorydb.go | 该文件主要实现基于内存的数据库操作方法
 2 | `Single`  | 单例初始化
 3 | `NewMemDB`| 初始化数据库
 4 | `Put`     | 存放键值对
 5 | `Get`     | 根据key获取value
 6 | `Has`     | 判断键是否存在
 7 | `Delete`  | 删除置顶key
 8 | `GetAll`  | 获取全部列表


## 单元测试 

序号 | Go文件/测试用例方法 | 说明
---|---|---
 1 | memorydb_test.go | 数据库操作测试用例
 2 | `TestMemoryDB_PutGet`  | 测试存取
 3 | `TestMemoryDB_Has`     | 测试是否存在
 4 | `TestMemoryDB_Delete`  | 测试删除
 5 | `TestMemoryDB_GetAll`  | 测试获取所有