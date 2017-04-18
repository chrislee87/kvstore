package kvstore

type Config struct {
	FileName      string
	FileCapacity  int64 // only use for create file(default 10000); for reopen file, ignore
	CompressCodec int64
}
