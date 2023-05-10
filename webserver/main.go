package main

import (
	"fmt"
	"net/http"
)

// 쿠키 설정 관련 함수 작성

func urlHandle(w http.ResponseWriter, r *http.Request) {
	sv_urlpath := r.URL.Path[1:]
	if sv_urlpath == "upload" {
		//파일 업로드 영역
		fmt.Fprintf(w, "%s", "파일 업로드를 이곳에서 하게됨")
	}
	fmt.Fprintf(w, "%s", sv_urlpath)
}

func main() {
	http.HandleFunc("/", urlHandle)

	http.ListenAndServe(":8080", nil)
}
