package main

import (
	"GoBlog/config"
	"GoBlog/models"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

/**
设置一个结构用来存储Json数据

*/
type IndexData struct {
	Title string `json:"title"` //``中的内容表示json格式显示的内容
	Desc  string `json:"desc"`
}

func isODD(num int) bool {
	return num%2 == 0
}

func getNextName(strs []string, index int) string {
	return strs[index+1]
}

func date(str string) string {
	return time.Now().Format(str)
}

func index(w http.ResponseWriter, r *http.Request) {
	//拿到当前路径
	path, _ := os.Getwd()
	t := template.New("index.html")
	//因为访问页面的时候有多个模板嵌套，因此需要把需要的数据全部解析
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	//解析文件
	t.Funcs(template.FuncMap{"isODD": isODD, "getNextName": getNextName, "date": date})
	t, err := t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)

	if err != nil {
		log.Println("解析模板出现错误", err)
	}
	//页面上涉及的数据必须有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	//返回页面
	t.Execute(w, hr)
}

func main() {
	//程序入口，一个程序只能有一个入口
	//web程序
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
