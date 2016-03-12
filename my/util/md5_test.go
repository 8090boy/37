package util

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	fmt.Printf("18600006666 : mm = 3737.io  sql = %v\n", Md5Encode("3737.io"))
}
