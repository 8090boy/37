package util

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)


type AesEncrypt struct{}

func (this *AesEncrypt) getKey() []byte {
	strKey := "s8.,OjH;`3kd6jfD5&$5q__jHfE"
	keyLen := len(strKey)
	if keyLen < 16 {
		panic("res key 长度不能小于16")
	}
	arrKey := []byte(strKey)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	//取前16个字节
	return arrKey[:16]
}

//加密
func (this *AesEncrypt) Encrypt(strMesg string) []byte {
	key := this.getKey()
	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMesg))
	return encrypted
}

//解密
func (this *AesEncrypt) Decrypt(src []byte) string {
	defer func() {
		//错误处理
		if e := recover(); e != nil {
			fmt.Println( e )
		}
	}()
	key := this.getKey()
	var iv = []byte(key)[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, src)
	return string(decrypted)
}
