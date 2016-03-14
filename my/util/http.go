package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"

	"net/http"
)

func WriteJSON(w http.ResponseWriter, arg interface{}) {
	var byteArr []byte = []byte{}
	if arg == nil {
		w.Write(byteArr)
		return
	}
	b, err := json.Marshal(arg)
	if err != nil {

		w.Write(byteArr)
		return
	}
	w.Write(b[:])
}

func WriteJSONP(w http.ResponseWriter, arg string) {
	w.Header().Set("Content-Type", "appliaction/x-javascript")
	var byteArr []byte = []byte{}
	if arg == "" {
		w.Write(byteArr)
		return
	}

	w.Write([]byte(arg))
}

func enCode(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	//   fmt.Println(cipherStr) //准备
	md5EnCode := hex.EncodeToString(cipherStr)
	return md5EnCode
}

func dnCode(md5EnCode string) string {
	md5DnCode, _ := hex.DecodeString(md5EnCode)
	return string(md5DnCode)
}

func getSha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	//切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	//SHA1 值经常以 16 进制输出

	return string(bs)
}
