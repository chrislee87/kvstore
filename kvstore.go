package kvstore

import (
	"os"
	"sync"
)

type KvStore struct {
	sync.RWMutex

	FileHeader
	Indexs     []Index
	File       *os.File
	Compress   func([]byte) []byte
	UnCompress func([]byte) []byte
}

func New(c *Config) (*KvStore, error) {
	if c == nil || c.FileName == "" {
		return nil, ErrWrongConfig
	}

	kv := new(KvStore)
	// File doesn't exist
	if _, err := os.Stat(c.FileName); os.IsNotExist(err) {
		if c.FileCapacity <= 0 {
			c.FileCapacity = DefaultFileCapacity

		}
		if _, err := c.GetCompressCodec(); err != nil {
			c.CompressCodec = 0
		}

		f, err := os.OpenFile(c.FileName, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return nil, ErrFailedCreateFile
		}

		kv.CompressCodec = c.CompressCodec
		kv.FileCapacity = c.FileCapacity
		kv.FileDataNums = 0
		kv.File = f
		kv.updateFileHeader()
		kv.setCompressFunc()
	} else {
		f, err := os.OpenFile(c.FileName, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return nil, ErrFailedCreateFile
		}
		kv.File = f
		kv.loadFileHeader()
		kv.setCompressFunc()
		kv.Indexs = make([]Index, kv.FileDataNums)
		kv.loadIndexes()
	}

	return kv, nil
}

func (kv *KvStore) updateFileHeader() error {
	kv.Lock()
	defer kv.Unlock()

	buf := make([]byte, SizeOfFileHeader)
	binary.LittleEndian.PutUint64(buf[:8], kv.CompressCodec)
	binary.LittleEndian.PutUint64(buf[8:16], kv.FileCapacity)
	binary.LittleEndian.PutUint64(buf[16:], kv.FileDataNums)

	_, err := kv.File.WriteAt(buf, 0)
	return err
}

func (kv *KvStore) loadFileHeader() {
	kv.RLock()
	defer kv.RUnlock()

	buf := make([]byte, SizeOfFileHeader)
	kv.File.ReadAt(buf, 0)
	kv.CompressCodec = binary.LittleEndian.Uint64(buf[:8])
	kv.FileCapacity = binary.LittleEndian.Uint64(buf[8:16])
	kv.FileDataNums = binary.LittleEndian.Uint64(buf[16:])
}

func (kv *KvStore) loadIndexes() {
	kv.RLock()
	defer kv.RUnlock()

	buf := make([]byte, SizeOfOneIndex*kv.FileDataNums)
	kv.File.ReadAt(buf, StartOffsetForIndexes)
	for i := 0; i < kv.FileDataNums; i++ {
		idx := Index{}
		offset := i * SizeOfOneIndex
		idx.Key = binary.LittleEndian.Uint64(buf[offset : offset+8])
		idx.Offset = binary.LittleEndian.Uint64(buf[offset+8 : offset+16])
		idx.Size = binary.LittleEndian.Uint64(buf[offset+16 : offset+24])
		kv.Indexs = append(kv.Indexs, idx)
	}
}

func (kv *KvStore) Get(key int64) []byte {

}

func (kv *KvStore) Append(key int64, buf []byte) error {

}
