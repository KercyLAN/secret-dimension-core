package cipher

import "encoding/base64"


// 对数据进行Base64编码
func Base64Byte(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// 对数据进行Base64解码
func Base64Decoded(data string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	return decoded, err
}
