package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	our "webcv/webcvpkg"
)

var vs []our.VolatileStat //전역변수로 휘발성 템프 변수 만듬
var cv_ing bool = false
var cv_time time.Time

func urlHandler(w http.ResponseWriter, r *http.Request) {
	sv_urlpath := r.URL.Path[1:] //sv_urlpath에 유저가 어떤 url을 요청했는지 저장됨
	if sv_urlpath == "" {
		fmt.Printf("%s: ", time.Now().Local().Format("01-02 15:04"))
		fmt.Println(our.GetIP(r), "에서 /")
		our.MainHanlder(w, r)
	} else if len(sv_urlpath) > 1 && sv_urlpath[0:3] == "kit" { //첫 조건에 길이 확인이 있는 이유는 인덱스 초과 슬라이싱을 막기 위함.
		//fmt.Printf("%s: ", time.Now().Local().Format("01-02 15:04"))
		//fmt.Println(our.GetIP(r), "에서 /kit"+sv_urlpath[3:], "요청")
		err := our.KitHanlder(w, r, sv_urlpath[3:])
		if err != nil {
			our.ErrHandler(w, r)
		}
	} else if sv_urlpath == "upload" {
		fmt.Printf("%s: ", time.Now().Local().Format("01-02 15:04"))
		fmt.Println(our.GetIP(r), "에서 ", sv_urlpath, "요청")
		our.UploadFileHandler(w, r, &vs)
	} else if sv_urlpath == "result" {
		fmt.Printf("%s: ", time.Now().Local().Format("01-02 15:04"))
		fmt.Println(our.GetIP(r), "에서 ", sv_urlpath, "요청")
		our.ResultHanlder(w, r)
	} else if sv_urlpath == "ajax" {
		fmt.Printf("%s: ", time.Now().Local().Format("01-02 15:04"))
		fmt.Println(our.GetIP(r), "에서 ", sv_urlpath, "요청")

		go our.Cv_loop(&vs, &cv_ing, &cv_time) // 요청시 마다 빈 것이 있는지&&cv가 비활성화인지 확인 go키워드로 백그라운드로 보내버림

		our.AjaxHanlder(w, r, &vs, &cv_time)
	} else if sv_urlpath == "result/file" {
		fmt.Printf("%s: ", time.Now().Local().Format("01-02 15:04"))
		fmt.Println(our.GetIP(r), "에서 ", sv_urlpath, "요청")
		our.ResultFileHanlder(w, r)
	} else {
		fmt.Printf("%s: ", time.Now().Local().Format("01-02 15:04"))
		fmt.Println(our.GetIP(r), "에서 ", sv_urlpath, "요청. 그리고 !에러! 발생")
		our.ErrHandler(w, r)
	}
}

func main() {
	var Port int = 0
	fmt.Println("사용할 Port 번호를 입력해주세요: ")
	fmt.Scanf("%d", &Port)
	server := http.NewServeMux()
	server.Handle("/", http.HandlerFunc(urlHandler))
	fmt.Println("http://localhost:"+strconv.Itoa(Port), "에서 요청을 기다리는 중:")
	err := http.ListenAndServe(":"+strconv.Itoa(Port), server)
	if err != nil { // http 서버 시작 중 문제 발생시
		log.Fatal(err)
		panic(err)
	}
}
