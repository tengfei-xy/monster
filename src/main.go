package main
// golang lib
import(
	"net/http"
	"golang.org/x/net/websocket"
)
// my lib
import(
	pnt "print"
)

var httpCli			= &http.Client{}

func main (){

	pnt.Info("Monster Start!")

	http.HandleFunc("/", monsterIndex)
	http.Handle("/ws", websocket.Handler(wsIndex))
	go http.ListenAndServe("0.0.0.0:80", nil)
	http.ListenAndServeTLS("0.0.0.0:443", "../ssl/ssl.crt", "../ssl/ssl.key", nil)
		
}
