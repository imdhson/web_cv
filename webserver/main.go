package main

import (
	"fmt"
	"net/http"
	"os"
)

// 파일 처리
func upload() error {

	return nil
}
func getCookie(w http.ResponseWriter, r *http.Request) int{
	saved, err := r.Cookie("") //key, value로 쿠키를 가져옴
	if err != nil {
		//쿠키가 없으니 nil 리턴
		return nil
	} else{
		//쿠키를 기반으로 결과창으로 넘기기 위해 값을 리턴
		return 1
}
func setCookie(w http.ResponseWriter, r *http.Request) int{
	idnumber := math.rand()
	willsave := http.Cookie{
		Name:	"idnumber",
		Value:	math.rand()
	}
}

// 쿠키 설정 관련 함수 작성

func urlHandle(w http.ResponseWriter, r *http.Request) {
	sv_urlpath := r.URL.Path[1:]
	if sv_urlpath == "upload" {
		upload()
	} else if sv_urlpath == "result" {
		//결과 표시해주는 창 
	} else {
		fmt.Fprintf(w, "%s", sv_urlpath) 
	}
}

func main() {
	http.HandleFunc("/", urlHandle)
	http.ListenAndServe(":8080", nil)
}
