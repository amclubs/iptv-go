// Package list
// @Time:2024/07/22 10:00
// @File:mian.go
// @SoftWare:Goland
// @Author:AM科技
// @Contact:https://t.me/AM_CLUBS

package list

import (
	"io"
	"net/http"
)

type Yylist struct {
}

type DataElement struct {
	Avatar string `json:"avatar"`
	Biz    string `json:"biz"`
	Desc   string `json:"desc"`
	Sid    int    `json:"sid"`
}

type ApiResponse struct {
	Data struct {
		IsLastPage int           `json:"isLastPage"`
		Data       []DataElement `json:"data"`
	} `json:"data"`
}

func (y *Yylist) Yylb(requesturl string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requesturl, nil)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return string(body)
}