kvstore
==============
Kvstore for time series data, with indexing and data compression.

Features
---------------
Store data with a uint64 key, for example use date(20160101) as key
Support Get and Append
Support data compression, both gzip and snappy

Architectures
---------------

CompressCodec |	FileCapacity | FileDataNums
-----------------|------|---------
Index(Key, Offset, Size)   | Index	| ... ...
Data Block    | Data Block	| ... ...

Each datafile contains three parts:
* File header
* Indexes
* Data blocks

File header contains CompressCodec, FileCapacity and FileDataNums.
Indexes are sorted by key. Each index contains Key, Offset and Size for the corresponding data block.
Data blocks can be compressed.
