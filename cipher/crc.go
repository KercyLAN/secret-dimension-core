package cipher

import (
	"hash/crc32"
)

// 对字符串进行CRC加密并返回其结果。
func CRC32String(str string) uint32{
	return CRC32Data([]byte(str))
}

// 对字节数组进行CRC加密并返回其结果。
func CRC32Data(data []byte) uint32{
	return crc32.ChecksumIEEE(data)
}

