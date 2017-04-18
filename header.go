package kvstore

type FileHeader struct {
	CompressCodec int64
	FileCapacity  int64
	FileDataNums  int64
}
