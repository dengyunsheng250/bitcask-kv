package fio

const (
	DataFilePerm = 0644
)

// 抽象IO管理器的接口
type IOManager interface {
	// Read从文件的指定位置读取相应的数据
	Read([]byte, int64) (int, error)
	// Write写入字节数组到文件中
	Write([]byte) (int, error)
	// Sync 将内存缓冲区的数据持久化到内存中
	Sync() error
	// Close 关闭文件
	Close() error
}
