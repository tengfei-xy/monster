package main

// golang lib
import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func getBaseDataResult(client *http.Client, key string) *goquery.Document {
	// 参数解释 https://blog.zfanw.com/google-search-url-parameters/
	// num=30 返回30条搜索结果
	// q = 关键词
	r, err := http.NewRequest("GET", url.QueryEscape(key), nil)
	if err != nil {
		panic(err)
	}
	r.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.193 Safari/537.36")
	r.Header.Add("Accept", "text/html")
	r.Header.Add("Origin", "https://www.hrloo.com")
	r.Header.Add("Host", "drm.vod2.myqcloud.com")
	r.Header.Add("sec-fetch-dest", "document")
	r.Header.Add("sec-fetch-mode", "navigate")
	r.Header.Add("sec-fetch-user", "?1")
	r.Header.Add("sec-ch-ua-mobile", "?0")
	r.Header.Add("Upgrade-Insecure-Requests", "1")

	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Print(res.Body)

	return doc

}
