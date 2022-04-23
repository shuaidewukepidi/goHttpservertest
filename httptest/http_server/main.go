package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

//访问/，返回Helloword
func Helloword(w http.ResponseWriter , r *http.Request){
	w.Write([]byte("Helloword"))
    
}
//访问/healthz,并返回200
func healthzFunc(w http.ResponseWriter , r *http.Request)  {
	healthcode := "200"
	w.Write([]byte(healthcode))
}

//接收客户端 request，并将 request 中带的 header 写入 response header
func  httpAccessFunc(w http.ResponseWriter , r *http.Request){
	if len(r.Header) > 0{
		for k, v := range r.Header {
			log.Printf("%s=%s",k ,v[0])
			//1. request header写入response header
			w.Header().Set(k,v[0])
		}
	}
	
}

//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
 func getversion(w http.ResponseWriter , r *http.Request){
	r.ParseForm() //解析所有请求数据，否则无法获取数据
	if len(r.Form) > 0 {
		for k, v := range r.Form {
			log.Printf("%s=%s", k, v[0])
		}
	}
	log.Printf("\n\n\n")

     os.Setenv("VERSION", "JDK version 1.11.0") //设置环境值的值
 
    //2. 获取环境变量"VERSION"
    name := os.Getenv("VERSION")
    log.Print("VERSION Env: ", name)
 }

//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
func getip(w http.ResponseWriter , r *http.Request){
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("err:", err)
	}
 
	if net.ParseIP(ip) != nil {
		fmt.Printf("ip ===>>%s\n", ip)
		log.Println(ip)
	}

}

func main ()  {
	//通过http.HandleFunc函数，注册一个请求处理器，第一个参数位请求路径，第二个参数位请求处理的函数主体func（ResponseWrite，*Request）
	http.HandleFunc("/",Helloword)
	http.HandleFunc("/healthz",healthzFunc)
	http.HandleFunc("/post",httpAccessFunc)
	http.HandleFunc("/version",getversion)
	http.HandleFunc("/localip",getip)

    //监听服务端口8080
	err :=http.ListenAndServe("127.0.0.1:8080",nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err.Error())
	}

}

	