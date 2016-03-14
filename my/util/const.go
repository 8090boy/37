package util

import (
	"os"
	"regexp"
	"time"
)

var Loca *time.Location

func init() {

	Loca, _ = time.LoadLocation(LoadLocation)
}

var (
	RegexpMobile   = regexp.MustCompile(`^[1|0]{1}[35789]{1}[0-9]{8,}$`)
	RegexpEmail    = regexp.MustCompile(`.+@.+\.[a-zA-Z]+$`)
	RegexpCommon   = regexp.MustCompile(`^[a-zA-Z\d_]+$`)
	RegexpPassword = regexp.MustCompile(`^[a-zA-Z\d_]{6,30}$`)
	RegexpWeChat   = regexp.MustCompile(`^[a-zA-Z\d_]{5,30}$`)
	RegexpFileter  = regexp.MustCompile(`\||\<|\>|\"|\s|\'|!`)
	PageOK         = "/OK.html"
	PageFailed     = "/failed.html"
	PageIndex      = "/index.html"
	TimeFormat     = "2006-01-02 15:04:05"
	LoadLocation   = "Asia/Shanghai"
)

func GetSysSplit() string {
	var spit = "/"

	if os.IsPathSeparator('\\') {
		spit = "\\"
	}
	return spit
}
