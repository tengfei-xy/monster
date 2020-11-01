package main
// golang lib
import(
	"net/http"
	"net/url"
	"strings"
	"fmt"
)
// my lib
import(
	pnt "print"
)
// third lib
import(
	"github.com/PuerkitoBio/goquery"
)

type rLine struct{
	Title		string
	Link		string
	Content		string
}

func getBaseDataResult(client *http.Client,key string) *goquery.Document {
	// 参数解释 https://blog.csdn.net/weixin_38796720/article/details/88991153
	// rn=30 返回30条搜索结果
	// tn=baidulocal 纯净结果、站内搜索(但会人机验证)
	// wd= 关键词
	r, err := http.NewRequest("GET", "https://www.baidu.com/s?rn=30&wd=" + url.QueryEscape(key), nil)
	if err != nil {
		panic(err)
	}
	r.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//r.Header.Add("Accept-Encoding", "gzip, deflate, br")
	r.Header.Add("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	r.Header.Add("Cache-Control", "no-cache")
	r.Header.Add("Connection", "keep-alive")
	r.Header.Add("Host", "www.baidu.com")
	r.Header.Add("Pragma", "no-cache")
	r.Header.Add("sec-fetch-dest", "document")
	r.Header.Add("sec-fetch-mode", "navigate")
	r.Header.Add("sec-fetch-site", "same-origin")
	r.Header.Add("sec-fetch-user", "?1")
	r.Header.Add("Upgrade-Insecure-Requests", "1")
	r.Header.Add("Cookie", "BIDUPSID=2F7E163138B8A6724F160141031D0E4F; PSTM=1603974261; BAIDUID=2F7E163138B8A6722AF9C467A87366E2:FG=1; BD_HOME=1; BD_UPN=123253; delPer=0; BD_CK_SAM=1; PSINO=5; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; H_PS_PSSID=32810_1452_32857_31660_32971_32706_7516_32115_32761_32917; BDSVRTM=1")

	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		panic(err)
	}

	return doc

}
func searchGo(key string) ([30]rLine,int) {
	r 				:= [30]rLine{}
	pnt.Search(key)

	doc := getBaseDataResult(httpCli,key).Find("body")
	baseLine := doc.Find("div[class~=result]")
	
	count	:= 0

	for i:=0;i< baseLine.Length();i++ {
		r[i].Title = strings.TrimSpace(baseLine.Find("h3>a").Eq(i).Text())

		// 如果 遇到 空 搜索结果
		// 则   跳过接下来可能的搜索条目
		if r[i].Title == "" {
			continue
		}

		r[i].Link,_ = baseLine.Find("h3>a").Eq(i).Attr("href")
		r[i].Content = baseLine.Find("div[class*=c-abstract]").Eq(i).Text()

		// 二次处理数据
		// 保证js的Json.parse()正常
		r[i].Content = strings.Replace(r[i].Content,`"`,`\\"`,-1)
		// pnt.Info(r[i].Title)
		// pnt.Info(r[i].Link)
		// pnt.Info(r[i].Content)

		// 搜索结果中包含标题的爬虫结果，才能+1
		count++

	}
	pnt.Search(fmt.Sprintf("%s 搜索到%d条",key,count))
	return r,count
}