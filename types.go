package kvstore

func (kv *KvStore) ReadAt(b []byte, offset uint64) (n int, err error) {
	return kv.File.ReadAt(b, int64(offset))
}

func (kv *KvStore) WriteAt(b []byte, offset uint64) (n int, err error) {
	return kv.File.WriteAt(b, int64(offset))
}
