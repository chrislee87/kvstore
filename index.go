package kvstore

type Index struct {
	Key    uint64 // for example, use date: 20160101
	Offset uint64
	Size   uint64
}
