kvstore
==============
Kvstore for time series data, with indexing and data compression.

Features
---------------
* Store data with a uint64 key, for example use date(20160101) as key
* Support Get and Append
* Support data compression, both gzip and snappy

Architectures
---------------
CompressCodec |	FileCapacity | FileDataNums
-----------------|------|---------
Index(Key, Offset, Size)   | Index	| ... ...
Data Block    | Data Block	| ... ...

Each datafile contains three parts:
* File header
	* File header contains CompressCodec, FileCapacity and FileDataNums.
* Indexes
	* Indexes are sorted by key. Each index contains Key, Offset and Size for the corresponding data block.
* Data blocks
	* Data blocks can be compressed, support gzip and snappy.

Usages
---------------
```
var (
	key1 uint64 = 20150110
	key2 uint64 = 20150120
	key3 uint64 = 20150130

	buf1 []byte = []byte("hello world 1.")
	buf2 []byte = []byte("hello world 2.")
	buf3 []byte = []byte("hello world 3.")
)

func main() {
	c := &Config{
		FileName:      "./data/gzip.file",
		FileCapacity:  3,
		CompressCodec: 1}

	kv, _ := New(c)

	kv.Append(key1, buf1)
	kv.Append(key2, buf2)
	kv.Append(key3, buf3)

	buf, _ := kv.Get(key1)

	kv.Close()
}

```
