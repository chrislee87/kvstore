package kvstore

type Config struct {
	FileName      string
	FileCapacity  int64 // only use for create file(default 10000); for reopen file, ignore
	CompressCodec int64 // only use for create file(default none) ; for reopen file, ignore
}

func (c *Config) GetCompressCodec() (string, error) {
	switch c.CompressCodec {
	case NoneCompress:
		return "None", nil
	case GzipCompress:
		return "Gzip", nil
	case SnappyCompress:
		return "Snappy", nil
	default:
		return "Error compress codec", ErrCompressCodec
	}
}
