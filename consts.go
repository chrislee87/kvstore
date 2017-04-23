package kvstore

const (
	DefaultFileCapacity   uint64 = 10000
	SizeOfFileHeader      uint64 = 24
	SizeOfOneIndex        uint64 = 24
	StartOffsetForIndexes uint64 = 24
)

const (
	NoneCompressCodec   uint64 = 0
	GzipCompressCodec   uint64 = 1
	SnappyCompressCodec uint64 = 2
)

func (kv *KvStore) setCompressFunc() {
	switch kv.CompressCodec {
	case NoneCompressCodec:
		kv.Compress = NoneCompress
		kv.UnCompress = NoneUnCompress
	case GzipCompressCodec:
		kv.Compress = GzipCompress
		kv.UnCompress = GzipUnCompress
	case SnappyCompressCodec:
		kv.Compress = SnappyCompress
		kv.UnCompress = SnappyUnCompress
	default:
		kv.Compress = NoneCompress
		kv.UnCompress = NoneUnCompress
	}
}
