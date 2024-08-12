package yqk

import (
  "Golang/list"
  "Golang/utils"
  "fmt"
  "encoding/json"
  "net/http"
  "strconv"
  "log"
)

// vercel 平台会将请求传递给该函数，这个函数名随意，但函数参数必须按照该规则。
// go语言大写就是公开，所以首字母必须大写
func Handler(w http.ResponseWriter, r *http.Request)  {
	path := r.URL.Path
	switch path {
	  case "/yqk/huyayqk.m3u":
		yaobj := &list.HuyaYqk{}
		res, _ := yaobj.HuYaYqk("https://live.cdn.huya.com/liveHttpUI/getLiveList?iGid=2135")
		var result list.YaResponse
		json.Unmarshal(res, &result)
		pageCount := result.ITotalPage
		pageSize := result.IPageSize
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=huyayqk.m3u")
		utils.GetTestVideoUrl(w)

		for i := 1; i <= pageCount; i++ {
			apiRes, _ := yaobj.HuYaYqk(fmt.Sprintf("https://live.cdn.huya.com/liveHttpUI/getLiveList?iGid=2135&iPageNo=%d&iPageSize=%d", i, pageSize))
			var res list.YaResponse
			json.Unmarshal(apiRes, &res)
			data := res.VList
			for _, value := range data {
				fmt.Fprintf(w, "#EXTINF:-1 tvg-logo=\"%s\" group-title=\"%s\", %s\n", value.SAvatar180, value.SGameFullName, value.SNick)
				fmt.Fprintf(w, "%s/huya/%v\n", utils.GetLivePrefix(r), value.LProfileRoom)
			}
		}
	  case "/yqk/douyuyqk.m3u":
		yuobj := &list.DouYuYqk{}
		resAPI, _ := yuobj.Douyuyqk("https://www.douyu.com/gapi/rkc/directory/mixList/2_208/list")

		var result list.DouYuResponse
		json.Unmarshal(resAPI, &result)
		pageCount := result.Data.Pgcnt

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=douyuyqk.m3u")
		utils.GetTestVideoUrl(w)

		for i := 1; i <= pageCount; i++ {
			apiRes, _ := yuobj.Douyuyqk("https://www.douyu.com/gapi/rkc/directory/mixList/2_208/" + strconv.Itoa(i))

			var res list.DouYuResponse
			json.Unmarshal(apiRes, &res)
			data := res.Data.Rl

			for _, value := range data {
				fmt.Fprintf(w, "#EXTINF:-1 tvg-logo=\"https://apic.douyucdn.cn/upload/%s_big.jpg\" group-title=\"%s\", %s\n", value.Av, value.C2name, value.Nn)
				fmt.Fprintf(w, "%s/douyu/%v\n", utils.GetLivePrefix(r), value.Rid)
			}
		}
	  case "/yqk/yylunbo.m3u":
		yylistobj := &list.Yylist{}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=yylunbo.m3u")
		utils.GetTestVideoUrl(w)

		i := 1
		for {
			apiRes := yylistobj.Yylb(fmt.Sprintf("http://rubiks-ipad.yy.com/nav/other/idx/213?channel=appstore&ispType=0&model=iPad8,6&netType=2&os=iOS&osVersion=17.2&page=%d&uid=0&yyVersion=6.17.0", i))
			var res list.ApiResponse
			json.Unmarshal([]byte(apiRes), &res)
			for _, value := range res.Data.Data {
				fmt.Fprintf(w, "#EXTINF:-1 tvg-logo=\"%s\" group-title=\"%s\", %s\n", value.Avatar, value.Biz, value.Desc)
				fmt.Fprintf(w, "%s/yy/%v\n", utils.GetLivePrefix(r), value.Sid)
			}
			if res.Data.IsLastPage == 1 {
				break
			}
			i++
		}
	  default:
		log.Println("Invalid path:", path)
		fmt.Fprintf(w, "<h1>链接错误!</h1>")
	}
}