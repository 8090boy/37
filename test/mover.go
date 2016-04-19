package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println(`Get index ...`)
	s, statusCode := Get("http://www.cnblogs.com/tt-0411/archive/2013/03/13/2958130.html")
	if statusCode != 200 {
		return
	}
	fmt.Println(s)
}

func Get(url string) (content string, statusCode int) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	content = string(data)
	return
}
