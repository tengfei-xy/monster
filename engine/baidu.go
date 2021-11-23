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
	i				:= 0
	haskey			:= false
	doc 			:= getBaseDataResult(httpCli,key).Find("body").Find("div[class~=result]")

	// 输出搜索关键词
	// pnt.Search(key)
	// h,_ := doc.Html()
	// pnt.Info(h)
	doc.Each(func(j int ,s * goquery.Selection){
		// 获取标题
		r[i].Title = strings.TrimSpace(s.Find("h3>a").Text())

		// 过滤搜索结果:标题为空
		if r[i].Title != "" {

			// 获取链接
			r[i].Link,_ = s.Find("h3>a").Attr("href")

			// 获取内容
			r[i].Content = s.Find("div[class*=c-abstract]").Text()

			// 过滤搜索结果:内容不含关键词(怕是Money上榜)
			if r[i].Content != "" {
				for _,K := range strings.Split(strings.ToUpper(key)," "){
					if strings.Index(strings.ToUpper(r[i].Content),K) != -1{
						haskey=true
						break
					}
				}
			}

			// 补充搜索结果:限制抓取robots
			if r[i].Content == ""{
				r[i].Content = s.Find("p[class*=c-color-text]").Text()
			} 
			// 补充搜索结果:视频的搜索结果
			if r[i].Content == ""{
				r[i].Content = s.Find("div[class~=c-span9]>font>p+p").Text()
			}
			if r[i].Content == ""{
				// 打印当前html
				//pnt.Search(s.Html())
				pnt.Warn("没有发现内容!关键词:%s,标题:%s",key,r[i].Title)
				haskey = false
			}
			if haskey{i++}
				// 输出log
				// pnt.Searchf("关键词:%s 第%d条",key,i)
				// pnt.Search(r[i].Title)
				// pnt.Search(r[i].Link)
				// pnt.Search(r[i].Content)
				

		}
	})
	pnt.Search(fmt.Sprintf(`baidu "%s" 搜索到%d条`,key,i))
	return r,i
}