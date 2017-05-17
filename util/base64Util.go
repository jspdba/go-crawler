package fetcher

import "encoding/base64"
//用base64进行编码
func Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))

}

//用base64进行解码
func Base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))

}

func Base642String(src string)(string, error) {
	b,e:= Base64Decode([]byte(src))
	return string(b),e
}

func String2Base64(src string)(string) {
	b:= Base64Encode([]byte(src))
	return string(b)
}