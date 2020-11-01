package print
import(
	"fmt"
	"time"
)
func now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func Info(s interface{}){
	fmt.Println(now()," Info      ",s)
}
func Infof(f string,a ...interface{}){
	fmt.Println(now()," Info      ",fmt.Sprintf(f,a...))
}
func Request(s string){
	fmt.Println(now()," Request   ",s)
}
func Search(s string){
	fmt.Println(now()," Search    ",s)
}
func Searchf(f string,a ...interface{}){
	fmt.Println(now()," Search    ",fmt.Sprintf(f,a...))
}
func Space(){
	fmt.Println(now(),"           ")
}
func Error(s error){
	fmt.Println(now()," Error     ",s)
}