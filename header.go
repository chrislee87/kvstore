package kvstore

type FileHeader struct {
	CompressCodec uint64
	FileCapacity  uint64
	FileDataNums  uint64
}
