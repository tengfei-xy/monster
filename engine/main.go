package main
// golang lib
import(
	"net/http"
	"golang.org/x/net/websocket"
	"encoding/json"
	"bytes"
)
// my lib
import(
	pnt "print"
)

var httpCli			= &http.Client{}

func main (){
	pnt.Info("Monster Start!")
	http.Handle("/", websocket.Handler(wsIndex))
	pnt.Info(http.ListenAndServe("0.0.0.0:1765", nil))
}


// web socket Index function
func wsIndex(ws *websocket.Conn) {
	var err error
	for {
		var wsmsg string
		if err = websocket.Message.Receive(ws, &wsmsg); err != nil {
			ws.Close()
			break
		}
		// 根据关键词搜索,返回包含标题、链接、内容的r结构体，以及c数量
		r,c := searchGo(wsmsg)

		// 将搜索结果转化成Json格式，并放回字符串string格式
		wsres := merge(r,c)

		// 打印返回结果
		//pnt.Info(wsres)

		if err = websocket.Message.Send(ws, wsres); err != nil {
			ws.Close()
			break
		}
	}
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