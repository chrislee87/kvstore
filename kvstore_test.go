package kvstore

import (
	"os"
	"reflect"
	"testing"
)

var (
	key1 uint64 = 20150110
	key2 uint64 = 20150120
	key3 uint64 = 20150130

	keylesser uint64 = 20150101
	keymiss   uint64 = 20150115
	keymore   uint64 = 20150131

	str0 string = "hello world."
	str1 string = "hello world 1."
	str2 string = "hello world 2."
	str3 string = "hello world 3."

	buf0 []byte = []byte(str0)
	buf1 []byte = []byte(str1)
	buf2 []byte = []byte(str2)
	buf3 []byte = []byte(str3)
)

func TestWrongConfig(t *testing.T) {
	c := &Config{}
	_, err := New(c)

	if err != ErrWrongConfig {
		t.Errorf("Config without filename, can't create kvstore!")
	}
}

func TestCreateEmptyFileAndGet(t *testing.T) {
	os.Remove("./data/empty.gzip.file")

	c := &Config{
		FileName:      "./data/empty.gzip.file",
		FileCapacity:  10,
		CompressCodec: 1}

	kv, err := New(c)
	if err != nil {
		t.Errorf("Failed to create kv, reason: %v\n", err)
	}

	_, err = kv.Get(key1)
	if err != ErrDataNotFound {
		t.Errorf("Get empty file, return value should be nil!")
	}
}

func TestAppendAndGet(t *testing.T) {
	os.Remove("./data/gzip.file")

	c := &Config{
		FileName:      "./data/gzip.file",
		FileCapacity:  3,
		CompressCodec: 1}

	kv, err := New(c)
	if err != nil {
		t.Errorf("Failed to create kv, reason: %v\n", err)
	}

	err = kv.Append(key1, buf1)
	if err != nil {
		t.Errorf("Failed to create key1")
	}

	err = kv.Append(key1, buf1)
	if err != ErrAppendFail {
		t.Errorf("key1 equal key1, shouldn't been appended.")
	}

	err = kv.Append(keylesser, buf0)
	if err != ErrAppendFail {
		t.Errorf("keylesser less than key1, shouldn't been appended.")
	}

	err = kv.Append(key2, buf2)
	if err != nil {
		t.Errorf("Failed to create key2, reason: %v\n", err)
	}

	err = kv.Append(key3, buf3)
	if err != nil {
		t.Errorf("Failed to create key3, reason: %v\n", err)
	}

	err = kv.Append(keymore, buf0)
	if err != ErrOutOfCapacity {
		t.Errorf("Out out the capacity, should not be appended. Reason: %v\n", err)
	}

	buf, err := kv.Get(key1)
	if !reflect.DeepEqual(buf, buf1) {
		t.Errorf("buf should equal buf1")
	}

	buf, err = kv.Get(key2)
	if !reflect.DeepEqual(buf, buf2) {
		t.Errorf("buf should equal buf2")
	}

	buf, err = kv.Get(key3)
	if !reflect.DeepEqual(buf, buf3) {
		t.Errorf("buf should equal buf3")
	}

	buf, err = kv.Get(keylesser)
	if err != ErrDataNotFound {
		t.Errorf("key is lesser, should not be found!")
	}

	buf, err = kv.Get(keymiss)
	if err != ErrDataNotFound {
		t.Errorf("key is miss, should not be found!")
	}

	buf, err = kv.Get(keymore)
	if err != ErrDataNotFound {
		t.Errorf("key is bigger, should not be found!")
	}
}
