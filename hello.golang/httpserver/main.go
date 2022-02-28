package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net/http"
	"os"
)

type HandleFnc func(w http.ResponseWriter, r *http.Request)

func init() {
	os.Setenv("VERSION", "1.0.0")
}

/**
编写一个 HTTP 服务器

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回200
*/
func main() {
	flag.Set("v", "4")
	http.HandleFunc("/", logPanics(simpleHandle))
	http.HandleFunc("/healthz", logPanics(healthz))
	glog.V(2).Info("start server...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

func logPanics(function HandleFnc) HandleFnc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic : %v", request.RemoteAddr, x)
				http.Error(writer, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		function(writer, request)
	}
}

func simpleHandle(w http.ResponseWriter, req *http.Request) {
	reqHeaderToRespHeader(w, req)
	getSvsVersion(w, req)
	myLog(w, req)
}

func reqHeaderToRespHeader(w http.ResponseWriter, req *http.Request) {
	if h := req.Header; h != nil {
		for s, strings := range h {
			if w.Header().Get(s) != "" {
				wsString := w.Header().Values(s)
				temp := make([]string, len(strings)+len(wsString))
				at := copy(temp, wsString)
				copy(temp[at:], strings)
				w.Header()[s] = temp
			} else {
				w.Header()[s] = strings
			}
		}
	}
	for s, strings := range w.Header() {
		fmt.Printf("response header key: %s ,value: %s", s, strings)
	}

}

func getSvsVersion(w http.ResponseWriter, req *http.Request) {
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	io.WriteString(w, "VERSION:"+w.Header().Get("VERSION"))
}

func myLog(w http.ResponseWriter, req *http.Request) {
	ip := req.RemoteAddr
	fmt.Printf("remmoteAddr: %s\n", ip)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
}
