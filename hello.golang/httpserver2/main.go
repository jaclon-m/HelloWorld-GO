package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"
)

/**
编写一个 HTTP 服务器

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回200
*/

func main() {
	flag.Set("V", "4")
	glog.V(2).Info("Starting server...")
	http.HandleFunc("/", SimpleHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "200")
}

func SimpleHandler(w http.ResponseWriter, req *http.Request) {
	reqHeaderToResponse(w, req)
	readVersion(w)
	myLog(w, req)
}

func myLog(w http.ResponseWriter, req *http.Request) {
	ip := req.RemoteAddr
	fmt.Printf("remmoteAddr: %s\n", ip)
}

func readVersion(w http.ResponseWriter) {
	version := os.Getenv("VERSION")
	if version != "" {
		w.Header().Add("VERSION", version)
	}
}

func reqHeaderToResponse(w http.ResponseWriter, req *http.Request) {
	for headName, headerValues := range req.Header {
		for _, v := range headerValues {
			w.Header().Add(headName, v)
		}
	}
}
