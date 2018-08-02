package testcrypto

import (
    "crypto/sha1"
    "encoding/hex"
	"fmt"
	"crypto/md5"
	"encoding/base64"
)

func testSha1(data string) string {
   dataSum := sha1.Sum([]byte(data))
   return hex.EncodeToString(dataSum[:])
}

func testMD5(data string) string {
	dataSum := md5.Sum([]byte(data))
	return hex.EncodeToString(dataSum[:])
}

//有不同的解码器 这里分Url Std 主要是数据的不同
//RawURLEncoding 解析的数据包含 "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
// RawStdEncoding 解析的数据包含 "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
//

//测试加密base64
func testEncodeBase64(data string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(data))
}

//测试解密base64
func testDecodeBase64(data string) string {
	bytes, err := base64.RawURLEncoding.DecodeString(data)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func TestCrypto() {
    fmt.Println(testSha1("hello123"))
    //aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d

	fmt.Println(testMD5("hello123"))
    //3cabc71b96c403c7d34e72c4fa0d615f

	encodeString := testEncodeBase64("hello123")
	fmt.Println(encodeString)
	//aGVsbG8xPzEx

    decodeString := testDecodeBase64(encodeString)
	fmt.Println(decodeString)
    //hello1?11
}