package helper

import (
	"fmt"
	"google-translator/constant"
	"google-translator/model"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Translate(obj model.Translate) ([]byte, error) {
	str2 := ""
	str2 = str2 + "q=" + obj.Text
	str2 = str2 + "&target=" + obj.TargetLanguage
	str2 = str2 + "&source=" + obj.SourceLanguage

	payload := strings.NewReader((str2))

	req, err := http.NewRequest(http.MethodPost, constant.TranslateURL, payload)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", os.Getenv("GoogleApiKey"))
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
