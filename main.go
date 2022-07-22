package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/**
设置一个结构用来存储Json数据

*/
type IndexData struct {
	Title string `json:"title"` //``中的内容表示json格式显示的内容
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //设置返回的数据是Json格式
	var indexData IndexData
	indexData.Title = "GoBlog"
	indexData.Desc = "这是一个Go语言开发的博客"
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

func main() {
	//程序入口，一个程序只能有一个入口
	//web程序
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
