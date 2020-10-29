package print
import(
	"fmt"
	"time"
)
func now() string {
	return time.Now().Format("2006-01-02 15:04:05.999999999")
}
func Info(s string){
	fmt.Print(now()," INFO   ",s,"\n")
}