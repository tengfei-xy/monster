package main
// golang lib
import(
	"net/http"
	"path/filepath"
	"io/ioutil"
	"golang.org/x/net/websocket"

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

			if err = websocket.Message.Send(ws, wsres); err != nil {
					ws.Close()
					break
			}
	}
}
func wsMain(key string) string{
	pnt.Search(key)
	r,c := searchGo(key)
	return templateGo(r,c)
}