package main
// golang lib
import(
	"net/http"
	"path/filepath"
	"io/ioutil"
	"golang.org/x/net/websocket"
	"encoding/json"
	"bytes"

)
// my lib
import(
	pnt "print"
)

// 
//
// web Index function
//
//

func monsterIndex(w http.ResponseWriter, r *http.Request) {
	var reqURLExt string = filepath.Ext(r.URL.Path)
	var reqURL string = r.URL.Path
	var res []byte
	//var err error
	//var ip string = r.RemoteAddr

	switch reqURLExt {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".ico":
		w.Header().Set("Content-Type", "image/x-ico")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	case ".jpg":
		w.Header().Set("Content-Type", "image/jpeg")
	default:
		w.Header().Set("Content-Type", "text/html")
	}
	pnt.Request(reqURL)
	switch reqURL {
	case `/`:
		res = monsterRoot()
	// case "r":
	// 	res = monsterResult(reqURL)
	default:
		res = monsterFile("../web"+reqURL)
	}
	w.Write(res)

}
func monsterRoot() []byte{
	index, err := ioutil.ReadFile("../web/index.html")
	if err != nil {
		pnt.Error(err)
	}
	return index
}
func monsterFile(file string) []byte{

	index, err := ioutil.ReadFile(file)
	if err != nil {
		pnt.Error(err)
	}
	return index
}

// 
//
// web socket Index function
//
//

func wsIndex(ws *websocket.Conn) {
	var err error
	for {
			var wsmsg string
			if err = websocket.Message.Receive(ws, &wsmsg); err != nil {
					ws.Close()
					break
			}
			wsres := wsMain(wsmsg)
			// 打印返回结果
			//pnt.Info(wsres)

			if err = websocket.Message.Send(ws, wsres); err != nil {
					ws.Close()
					break
			}
	}
}
func wsMain(key string) string{
	
	// 根据关键词搜索,返回包含标题、链接、内容的r结构体，以及c数量
	r,c := searchGo(key)

	// 将搜索结果转化成Json格式，并放回字符串string格式
	return merge(r,c)
}
 
func merge(r [30]rLine,c int) string{
	var jResult []byte
	buf := bytes.NewBuffer(jResult)
	// 单引号，表示总体是Json的规范
	// [],Json数组
	buf.WriteString("[")
	for i:=0;i<c;i++ {
		json,err := json.Marshal(r[i])
		if err !=nil {
			pnt.Error(err)
		}

		buf.Write(json)
		// 对于最后一个搜索结果的条目，不加逗号
		if (i!=c-1){
			buf.WriteString(",")
		}
	}
	buf.WriteString("]")

	return buf.String()
}