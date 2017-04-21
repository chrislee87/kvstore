package kvstore

import (
	"bytes"
	"compress/gzip"
	"io"

	"github.com/golang/snappy"
)

// None
func NoneCompress(bs []byte) []byte {
	return bs
}

func NoneUnCompress(bs []byte) []byte {
	return bs
}

// Gzip
func GzipCompress(bs []byte) []byte {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	zw.Write(bs)
	zw.Flush()
	zw.Close()

	return buf.Bytes()
}

func GzipUnCompress(bs []byte) []byte {
	var buf bytes.Buffer
	bsBuf := bytes.NewBuffer(bs)
	zr, _ := gzip.NewReader(bsBuf)
	io.Copy(&buf, zr)
	zr.Close()

	return buf.Bytes()
}

// Snappy
func SnappyCompress(bs []byte) []byte {
	return snappy.Encode(nil, bs)
}

func SnappyUnCompress(bs []byte) []byte {
	buf, _ := snappy.Decode(nil, bs)
	return buf
}
