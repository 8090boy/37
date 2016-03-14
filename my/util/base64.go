package util

import (
"encoding/base64"
"fmt"
)

const base64Table = "1dfTsdYRsg6845tuu575wfddfrty3g4dgfggh345rkljklyyiz3254utysdfT645"

var coder = base64.NewEncoding(base64Table)

func Base64Encode(src  []byte ) string {
	fmt.Printf("加密前=  %v\n", src )
	str := coder.EncodeToString( src )
	fmt.Printf("加密后=  %v\n", str )
	return  str
}

func Base64Decode(src string) []byte {
	fmt.Printf("解密前=  %v\n", src )
	b, _ := coder.DecodeString(src)
      fmt.Printf("解密后=  %v\n", b )
	return b
}
