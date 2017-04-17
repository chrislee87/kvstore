package kvstore

const (
	FileCapacity = 10000
)

const (
	NoneCompress = iota
	GzipCompress
	SnappyCompress
	Lz4Compress
)
