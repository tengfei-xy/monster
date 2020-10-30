package main
// golang lib
import(
	"html/template"
	"bytes"
)
// my lib
import(
	pnt "print"
)

// 作为被解析的html模板 结构体

func templateGo(r [30]rLine,c int) string{

	// 1 读取并解析 样本
	t, _ := template.ParseFiles("../web/template/eachline.html")
	var l rLine

	// 2 填充 HTML
	res := make([]byte,2000)
	w  := bytes.NewBuffer(res)
	for i:=0;i<=c;i++{
		l.Title		= r[i].Title
		l.Link		= r[i].Link
		l.Content	= r[i].Content
		// pnt.Info(r[i].Title)
		// pnt.Info(r[i].Link)
		// pnt.Info(r[i].Content)
		err := t.Execute(w, l)
		if err != nil{
			pnt.Error(err)
		}
	}
	//pnt.Info(string(w.Bytes()))

	return w.String()

}
