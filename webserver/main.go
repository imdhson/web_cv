package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"webcv/webcvpkg"
)

var vs []webcvpkg.VolatileStat //전역변수로 휘발성 템프 변수 만듬
var cv_ing bool = false
var cv_time time.Time

func urlHandler(w http.ResponseWriter, r *http.Request) {
	sv_urlpath := r.URL.Path[1:] //sv_urlpath에 유저가 어떤 url을 요청했는지 저장됨
	if sv_urlpath == "" {
		fmt.Println("Path: /", "IP주소: ", webcvpkg.GetIP(r))
		webcvpkg.MainHanlder(w, r)
	} else if len(sv_urlpath) > 1 && sv_urlpath[0:3] == "kit" { //첫 조건에 길이 확인이 있는 이유는 인덱스 초과 슬라이싱을 막기 위함.
		//fmt.Println("Path:/kit" + sv_urlpath[3:]) // kit 폴더를 가져오게됨
		err := webcvpkg.KitHanlder(w, r, sv_urlpath[3:])
		if err != nil {
			webcvpkg.ErrHander(w, r)
		}
	} else if sv_urlpath == "upload" {
		fmt.Println("Path: ", sv_urlpath, "IP주소: ", webcvpkg.GetIP(r))
		webcvpkg.UploadFileHandler(w, r, &vs)
	} else if sv_urlpath == "result" {
		fmt.Println("Path: ", sv_urlpath, "IP주소: ", webcvpkg.GetIP(r))
		webcvpkg.ResultHanlder(w, r)
	} else if sv_urlpath == "ajax" {
		fmt.Println("Path: ", sv_urlpath, "IP주소: ", webcvpkg.GetIP(r))

		go webcvpkg.Cv_loop(&vs, &cv_ing, &cv_time) // 요청시 마다 빈 것이 있는지&&cv가 비활성화인지 확인 go키워드로 백그라운드로 보내버림

		webcvpkg.AjaxHanlder(w, r, &vs, &cv_time)
	} else if sv_urlpath == "result/file" {
		fmt.Println("Path: ", sv_urlpath, "IP주소: ", webcvpkg.GetIP(r))
		webcvpkg.ResultFileHanlder(w, r)
	} else {
		fmt.Println("잘못된 Path: ", sv_urlpath, "IP주소: ", webcvpkg.GetIP(r))
		webcvpkg.ErrHander(w, r)
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
