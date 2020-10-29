package main
// golang lib
import(
	"net/http"
)
// my lib
import(
	pnt "print"
)

var httpCli			= &http.Client{}

func main (){

	pnt.Info("Monster Start!")

	http.HandleFunc("/", monsterIndex)

	http.ListenAndServe("0.0.0.0:80", nil)
	//http.ListenAndServeTLS("0.0.0.0:443", "ssl/ssl.crt", "ssl/ssl.key", nil)
		
}