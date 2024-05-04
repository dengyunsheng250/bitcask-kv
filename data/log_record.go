package data

// 数据内存索引，主要是描述数据在磁盘上的位置
type LogRecordPos struct {
	Fid    uint32 // 文件id, 存储在了哪个文件
	Offset int64  // 偏移
}
