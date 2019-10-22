package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
)

//创建log
var log = logs.NewLogger()

//初始化，日志输出方式采用beego-logs study
func init() {
	//文件输出
	//_ = logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	//控制台输出
	_ = log.SetLogger(logs.AdapterConsole)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	GetBody := r.GetBody
	log.Info("这是一个很开心的脸 %s %v", GetBody, 3)
	_, _ = fmt.Fprint(w, "你好世界")
}

func main() {

	//===================test====================
	//an official log.Logger with prefix ORM
	log.Info("======================")
	log.Debug("my book is bought in the year of ", 2016)
	log.Info("this %s cat is %v years old", "yellow", 3)
	log.Warn("json is a type of kv like", map[string]int{"key": 2016})
	log.Error("1024", "is a very", "good game")
	log.Info("======================")
	//===================test====================

	//创建一个新的http路由管理器
	mux := http.NewServeMux()
	mux.HandleFunc("/index", indexHandler)
	//只是监听8080端口
	_ = http.ListenAndServe(":8080", mux)

	//Clinet -> Requests ->  [Multiplexer(router) -> handler  -> Response -> Clinet
	man := Man{name: "zhangbo", age: 15, length: 13}
	man.handleFunc(man.age, atLastHandleForAge)
	man.handleFunc(man.length, atLastHandleForLength)
}

type Man struct {
	name   string
	age    int
	length int
}
type NumHandler interface {
	//最终处理at last
	handle(num int)
}

type HandlerInt func(num int)

func (handlerInt HandlerInt) handle(num int) {
	handlerInt(num)
	//handlerInt.handle(num)
}

func atLastHandleForAge(age int) {
	logs.Info("我的年龄是：%v岁", age)
}
func atLastHandleForLength(length int) {
	logs.Info("我的长度是：%v米", length)
}

func (man Man) handleFunc(num int, atLastHandle func(num int)) {
	//第二个参数是把传入的函数atLastHandle 强转成 HandlerInt类型，这样atLastHandle就实现了NumHandler接口。
	//var handlerInt NumHandler  = HandlerInt(atLastHandle)
	atLastCall(num, HandlerInt(atLastHandle))
}
func atLastCall(int int, numHandler NumHandler) {
	numHandler.handle(int)
}
