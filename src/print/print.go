package print
import(
	"fmt"
	"time"
)
func now() string {
	return time.Now().Format("2020-11-20 00:00:00")
}
func Info(s string){
	fmt.Println(now()," INFO      ",s)
}
func Request(s string){
	fmt.Println(now()," Request   ",s)
}
func Error(s error){
	fmt.Println(now()," Error     ",s)
}