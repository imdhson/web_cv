package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func mainHanlder(w http.ResponseWriter, r *http.Request) {
	w.Write("Test")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) error {
	setID(w, r)
	return nil
}

func resultHanlder() error {
	return nil
}

func getID(w http.ResponseWriter, r *http.Request) int {
	id, err := r.Cookie("id") //key, value로 쿠키를 가져옴
	if err != nil {
		//쿠키가 없으니 nil 리턴
		return -1
	} else {
		//쿠키를 기반으로 결과창으로 넘기기 위해 값을 리턴
		cvalue, verr := strconv.Atoi(id.Value) //쿠키를 가져와서 string을 int로 바꿈
		if verr != nil {                       //오류가 없으면 id값 반환
			return cvalue
		} else {
			panic(verr)
		}
	}
}

func setID(w http.ResponseWriter, r *http.Request) int {
	id := rand.Int()
	cookieid := http.Cookie{
		Name:     "id",
		Value:    strconv.Itoa(id),
		HttpOnly: true,
	}
	//w.Header().Set("Set-Cookie", cookieid.String())
	http.SetCookie(w, &cookieid)
	return id
}

func urlHandle(w http.ResponseWriter, r *http.Request) {
	sv_urlpath := r.URL.Path[1:] //sv_urlpath에 유저가 어떤 Url을 쳤는지 저장됨
	if sv_urlpath == "upload" {

	} else if sv_urlpath == "result" {
		//결과 표시해주는 창
	} else {
		fmt.Fprintf(w, "%s", sv_urlpath)
	}
}

func main() {
	server := http.NewServeMux()
	server.Handle("/", http.HandlerFunc(urlHandle))
	err := http.ListenAndServe(":8080", server)
	if err != nil { // http 서버 시작 중 문제 발생시
		log.Fatal(err)
		panic(err)
	}
}
