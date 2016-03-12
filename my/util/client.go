package util

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

var client = &http.Client{}

func GetUserInfo(toUrl string, headerToken string) []byte {

	request, _ := http.NewRequest("GET", toUrl, nil)
	request.Header.Add("token", headerToken)
	response, err := client.Do(request)
	if err != nil {

		return nil
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		byteArr, err := ioutil.ReadAll(response.Body)
		if err != nil {

			return nil
		}
		jsonStr := string(byteArr)

		if len(jsonStr) > 2 {
			return byteArr
		}
	}

	return nil
}

func Get(toUrl string) []byte {

	request, _ := http.NewRequest("GET", toUrl, nil)
	response, err := client.Do(request)
	if err != nil {
		return nil
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		byteArr, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil
		}
		jsonStr := string(byteArr)

		if len(jsonStr) > 2 {
			return byteArr
		}
	}
	return nil
}

func Post(toUrl string, param url.Values) ([]byte, string) {

	resp, err := client.PostForm(toUrl, param)
	if err != nil {

		return nil, ""
	}

	defer resp.Body.Close()

	token := resp.Header.Get("token")

	if resp.StatusCode == 200 {
		byteArr, _ := ioutil.ReadAll(resp.Body)
		jsonStr := string(byteArr)
		if len(jsonStr) > 2 {
			return byteArr, token
		}
	}

	return nil, ""
}
