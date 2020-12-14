package main
// golang lib
import(
	"net/http"
	"net/url"
	"strings"
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
	// 参数解释 https://blog.zfanw.com/google-search-url-parameters/
	// num=30 返回30条搜索结果
	// q = 关键词
	r, err := http.NewRequest("GET", "https://www.google.com/search?num=30&q=" + url.QueryEscape(key), nil)
	if err != nil {
		panic(err)
	}
	r.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.193 Safari/537.36")
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	r.Header.Add("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	r.Header.Add("Cache-Control", "no-cache")
	r.Header.Add("Pragma", "no-cache")
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
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		panic(err)
	}

	return doc

}
func searchGo(key string) ([30]rLine,int) {
	r 				:= [30]rLine{}
	i				:= 0
	haslink			:= false
	doc 			:= getBaseDataResult(httpCli,key).Find("div[class=g]")

	// 输出搜索关键词
	// pnt.Search(key)
	doc.Each(func(j int ,s * goquery.Selection){
		// 获取标题
		r[i].Title = strings.TrimSpace(s.Find("h3").Text())

		// 获取连接
		r[i].Link,haslink = s.Find("a").Attr("href")

		// 获取内容
		r[i].Content = s.Find("div>div+div").Text()

		if haslink{
			//`r[i].Link = strings.Split(r[i].Link,"url=")[1]
			i++
			// pnt.Result("标题:",r[i].Title )
			// pnt.Result("链接",r[i].Link)
			// pnt.Result("内容",r[i].Content)
		}

	})
	pnt.Searchf(`google "%s" 搜索到%d条`,key,i)
	return r,i
}