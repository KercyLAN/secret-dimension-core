package read

import (
	"io/ioutil"
	"os"
)

// 单次读取文件
//
// 一次性对整个文件进行读取，小文件读取可以很方便的一次性将文件内容读取出来，而大文件读取会造成性能影响。
func ReadOnce(filePath string) ([]byte, error) {
	if file, err := os.Open(filePath); err != nil {
		return nil, err
	}else {
		defer file.Close()
		return ioutil.ReadAll(file)
	}
}
