package webcvpkg

import (
	"fmt" //ioutill 대체하기
	"log"
	"net/http"
	"os"
)

func ErrHander(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//wwwfile, err := ioutil.ReadFile("./www/err.html")
	wwwfile, err := os.ReadFile("./www/err.html")
	if err != nil {
		fmt.Println("www/err.html 을 로드할 수 없음")
		log.Fatal(err)
		panic(err)
	} else {
		w.Write(wwwfile)
	}
}
