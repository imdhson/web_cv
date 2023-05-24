package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webcv/webcvpkg"
)

var VolatileStat []int //전역 변수 아직에러

func urlHandler(w http.ResponseWriter, r *http.Request) {
	sv_urlpath := r.URL.Path[1:] //sv_urlpath에 유저가 어떤 url을 요청했는지 저장됨
	if sv_urlpath == "" {
		fmt.Println("Path: /", "IP주소: ", webcvpkg.GetIP(r))
		webcvpkg.MainHanlder(w, r)
	} else if len(sv_urlpath) > 1 && sv_urlpath[0:3] == "kit" { //첫 조건에 길이 확인이 있는 이유는 인덱스 초과 슬라이싱을 막기 위함.
		fmt.Println("Path:/kit" + sv_urlpath[3:]) // kit 폴더를 가져오게됨
		err := webcvpkg.KitHanlder(w, r, sv_urlpath[3:])
		if err != nil {
			webcvpkg.ErrHander(w, r)
		}
	} else if sv_urlpath == "upload" {
		fmt.Println("Path: ", sv_urlpath, "IP주소: ", webcvpkg.GetIP(r))
		webcvpkg.UploadHanlder(w, r)
	} else if sv_urlpath == "result" {
		fmt.Println("Path: ", sv_urlpath, "IP주소: ", webcvpkg.GetIP(r))
		webcvpkg.ResultHanlder(w, r)
	} else if sv_urlpath == "ajax" {
		fmt.Println("Path: ", sv_urlpath, "IP주소: ", webcvpkg.GetIP(r))
		webcvpkg.AjaxHanlder(w, r)
	} else if sv_urlpath == "result/file" {
		fmt.Println("Path: ", sv_urlpath, "IP주소: ", webcvpkg.GetIP(r))
		webcvpkg.ResultFileHanlder(w, r)
	} else if sv_urlpath == "uploadfile" { //upload/file 로 하면 쿠키 할당이 안되는 버그가있음. 원인 못찾음
		fmt.Println("Path: ", sv_urlpath, "IP주소: ", webcvpkg.GetIP(r))
		webcvpkg.UploadFileHandler(w, r)
	} else {
		fmt.Println("잘못된 Path: ", sv_urlpath, "IP주소: ", webcvpkg.GetIP(r))
		webcvpkg.ErrHander(w, r)
	}
}

func main() {
	const PORT int = 8080
	server := http.NewServeMux()
	server.Handle("/", http.HandlerFunc(urlHandler))
	fmt.Println("http://localhost:"+strconv.Itoa(PORT), "에서 요청을 기다리는 중:")
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), server)
	if err != nil { // http 서버 시작 중 문제 발생시
		log.Fatal(err)
		panic(err)
	}
}
