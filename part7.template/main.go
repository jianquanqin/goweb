package main

import (
	"html/template"
	"net/http"
	"time"
)

func main() {

	//server := http.Server{
	//	Addr: "localhost:8080",
	//}
	//http.HandleFunc("/process", process)
	//server.ListenAndServe()

	//第一种方法
	//t,_ := template.ParseFiles("tmpl01.html") //调用解析函数，先把模版文件里面的内容读成字符串，调用模版上面的parse方法解析字符串
	//第二种方法
	//t := template.New("tmpl01.html")
	//t, _ = t.ParseFiles("tmpl01.html")
	//
	//t, _ := template.ParseGlob("*.html") //*是通配符，会找到所有文件名，模版名为第一个文件名

	//调用加载模版代码
	//templates := loadTemplates()
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fileName := r.URL.Path[1:]      //截取后边的部分，获取文件名
	//	t := templates.Lookup(fileName) //通过文件名寻找模版
	//	if t != nil {
	//		err := t.Execute(w, nil) //执行模版
	//		if err != nil {
	//			log.Fatalln(err.Error())
	//		}
	//	} else {
	//		w.WriteHeader(http.StatusNotFound)
	//	}
	//})
	//http.Handle("/css/", http.FileServer(http.Dir("wwwroot")))
	//http.Handle("/img/", http.FileServer(http.Dir("wwwroot")))
	//
	//http.ListenAndServe("localhost:8080", nil)

	//Action
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()

}

func process(w http.ResponseWriter, r *http.Request) {

	//t, _ := template.ParseFiles("tmpl.html") //text/template库，解析模版
	//t.Execute(w, "Hello World")              //执行传入ResponseWriter 和 上下文数据

	//执行模版的两个方法
	//t, _ := template.ParseFiles("t1.html")
	//t.Execute(w, "Hello world")
	//
	//ts, _ := template.ParseFiles("t1.html", "t2.html")
	//ts.ExecuteTemplate(w, "t2.html", "Hello World")

	//Action
	//条件
	//t, _ := template.ParseFiles("tmpl02.html")
	//rand.Seed(time.Now().Unix())    //从1970年1月日到现在的秒数
	//t.Execute(w, rand.Intn(10) > 5) //生成随机数的范围是0-9，判断是否大于5

	//Action
	//遍历
	//t, _ := template.ParseFiles("tmpl03.html")
	////daysofWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sum"}
	//daysofWeek := []string{}
	//t.Execute(w, daysofWeek)

	// //Action
	// //设置
	// t, _ := template.ParseFiles("tmpl04.html")
	// t.Execute(w, "hello")

	//包含
	//t,_ := template.ParseFiles("tmpl05.html","tmpl06.html")
	//t.Execute(w, "hello world")

	//函数与管道
	//t, _ := template.ParseFiles("tmpl07.html")
	//t.Execute(w, nil)

	//编写自定义的函数
	funcMap := template.FuncMap{"fdate": formatDate} //放入FuncMap
	t := template.New("tmpl08.html").Funcs(funcMap)  //创建新模版
	t.ParseFiles("tmpl08.html")                      //解析
	t.Execute(w, time.Now())                         //执行，操作顺序非常重要
}
func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

////解析模版
//func loadTemplates() *template.Template {
//	result := template.New("templates")                 //创建一个空的模版
//	template.Must(result.ParseGlob("templates/*.html")) //程序刚开始启动时就应该加载，失败的话终止程序,模版名就是文件名
//	return result
//}
