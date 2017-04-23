package kvstore

type Config struct {
	FileName      string
	FileCapacity  uint64 // only use for create file(default 10000); for reopen file, ignore
	CompressCodec uint64 // only use for create file(default none) ; for reopen file, ignore
}

func (c *Config) GetCompressCodec() (string, error) {
	switch c.CompressCodec {
	case NoneCompressCodec:
		return "None", nil
	case GzipCompressCodec:
		return "Gzip", nil
	case SnappyCompressCodec:
		return "Snappy", nil
	default:
		return "Error compress codec", ErrCompressCodec
	}
}
