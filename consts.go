package kvstore

const (
	DefaultFileCapacity   int64 = 10000
	SizeOfFileHeader      int64 = 24
	SizeOfOneIndex        int64 = 24
	StartOffsetForIndexes int64 = 24
)

const (
	NoneCompress   int64 = 0
	GzipCompress   int64 = 1
	SnappyCompress int64 = 2
)

func (kv *KvStore) setCompressFunc() {
	switch kv.CompressCodec {
	case NoneCompress:
		kv.Compress = NoneCompress
		kv.UnCompress = NoneUnCompress
	case GzipCompress:
		kv.Compress = GzipCompress
		kv.UnCompress = GzipUnCompress
	case SnappyCompress:
		kv.Compress = SnappyCompress
		kv.UnCompress = SnappyUnCompress()
	default:
		kv.Compress = NoneCompress
		kv.UnCompress = NoneUnCompress
	}
}
