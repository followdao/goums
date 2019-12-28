package utils

import (
	"bytes"
	"encoding/binary"
	"strings"
)

// StrBuilder build string
func StrBuilder(args ...string) string {
	var str strings.Builder

	for _, v := range args {
		str.WriteString(v)
	}
	return str.String()
}

// ByteBuilder build []byte
func ByteBuilder(args ...[]byte) []byte {
	var buffer bytes.Buffer

	for _, v := range args {
		buffer.Write(v)
	}
	return buffer.Bytes()
}

// Int64ToBytes  int64 to byte
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// BytesToInt64 byte to int64
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
