package index

import (
	"bitcask-go/data"
	"bytes"

	"github.com/google/btree"
)

// Indexer抽象索引接口, 后续如果想接入其他的数据结构， 则实现这个接口即可
type Indexer interface {
	// Put 向索引中存放key对应的数据位置信息
	Put(key []byte, pos *data.LogRecordPos) bool
	// Get 根据key取出对应索引位置信息
	Get(key []byte) *data.LogRecordPos
	// Delete 根据key删除对应索引位置信息
	Delete(key []byte) bool
}

type Item struct {
	key []byte
	pos *data.LogRecordPos
}

func (i *Item) Less(bi btree.Item) bool {
	return bytes.Compare(i.key, bi.(*Item).key) == -1
}
