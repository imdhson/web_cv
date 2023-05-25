package webcvpkg

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func UploadHanlder(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//wwwfile, err := ioutil.ReadFile("./www/upload.html")
	wwwfile, err := os.ReadFile("./www/upload.html")
	if err != nil {
		fmt.Println("www/upload.html 을 로드할 수 없음")
		log.Fatal(err)
		panic(err)
	} else {
		w.Write(wwwfile)
	}
}
