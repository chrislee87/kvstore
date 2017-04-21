package kvstore

type Index struct {
	Key    int64 // for example, use date: 20160101
	Offset int64
	Size   int64
}
