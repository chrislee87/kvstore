package kvstore

import (
	"encoding/binary"
	"os"
	"sort"
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
			return nil, ErrFailedOpenFile
		}
		kv.File = f
		kv.loadFileHeader()
		kv.setCompressFunc()
		kv.Indexs = make([]Index, kv.FileDataNums)
		kv.loadIndexes()
	}

	return kv, nil
}

func (kv *KvStore) Close() error {
	return kv.File.Close()
}

func (kv *KvStore) updateFileHeader() error {
	kv.Lock()
	defer kv.Unlock()

	return kv.updateFileHeaderNoLock()
}

func (kv *KvStore) updateFileHeaderNoLock() error {
	buf := make([]byte, SizeOfFileHeader)
	binary.LittleEndian.PutUint64(buf[:8], kv.CompressCodec)
	binary.LittleEndian.PutUint64(buf[8:16], kv.FileCapacity)
	binary.LittleEndian.PutUint64(buf[16:], kv.FileDataNums)

	_, err := kv.WriteAt(buf, 0)
	return err
}

func (kv *KvStore) appendLastIndexNoLock() error {
	idx := kv.Indexs[len(kv.Indexs)-1]
	buf := make([]byte, SizeOfOneIndex)
	binary.LittleEndian.PutUint64(buf[:8], idx.Key)
	binary.LittleEndian.PutUint64(buf[8:16], idx.Offset)
	binary.LittleEndian.PutUint64(buf[16:], idx.Size)

	_, err := kv.WriteAt(buf, StartOffsetForIndexes+SizeOfOneIndex*(uint64)(len(kv.Indexs)-1))
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
	kv.ReadAt(buf, StartOffsetForIndexes)
	for i := uint64(0); i < kv.FileDataNums; i++ {
		idx := Index{}
		offset := i * SizeOfOneIndex
		idx.Key = binary.LittleEndian.Uint64(buf[offset : offset+8])
		idx.Offset = binary.LittleEndian.Uint64(buf[offset+8 : offset+16])
		idx.Size = binary.LittleEndian.Uint64(buf[offset+16 : offset+24])
		kv.Indexs = append(kv.Indexs, idx)
	}
}

func (kv *KvStore) Get(key uint64) ([]byte, error) {
	kv.RLock()
	defer kv.RUnlock()

	idx := sort.Search(len(kv.Indexs), func(i int) bool { return kv.Indexs[i].Key >= key })
	if idx == len(kv.Indexs) || kv.Indexs[idx].Key != key {
		return nil, ErrDataNotFound
	}
	offset := kv.Indexs[idx].Offset
	size := kv.Indexs[idx].Size
	buf := make([]byte, size)
	kv.ReadAt(buf, offset)
	return kv.UnCompress(buf), nil
}

func (kv *KvStore) Append(key uint64, buf []byte) error {
	kv.Lock()
	defer kv.Unlock()

	if kv.FileDataNums >= kv.FileCapacity {
		return ErrOutOfCapacity
	}

	n := len(kv.Indexs)
	if n > 0 && kv.Indexs[n-1].Key >= key {
		return ErrAppendFail
	}

	wb := kv.Compress(buf)

	// FileHeader
	kv.FileDataNums += 1
	kv.updateFileHeaderNoLock()

	// Indexes
	var idx Index
	idx.Key = key
	idx.Size = uint64(len(wb))
	if n == 0 {
		idx.Offset = StartOffsetForIndexes + SizeOfOneIndex*DefaultFileCapacity
	} else {
		idx.Offset = kv.Indexs[n-1].Offset + kv.Indexs[n-1].Size
	}
	kv.Indexs = append(kv.Indexs, idx)
	kv.appendLastIndexNoLock()

	// FileBody
	kv.WriteAt(wb, idx.Offset)
	return nil
}
