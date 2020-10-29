package main
// golang lib
import(
	"net/http"
	"path/filepath"
	"io/ioutil"
)
// my lib
import(
	pnt "print"
)

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
	case "r":
		res = monsterResult(reqURL)
	}
	w.Write(res)

}
func monsterRoot()[]byte{
	
	index, err := ioutil.ReadFile("../web/index.html")
	if err != nil {
		pnt.Error(err)
	}
	return index

}
func monsterResult(key string)[]byte{
	getLine(key)
	return nil
}