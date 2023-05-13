package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strconv"
)

func mainHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset = utf-8")
	wwwfile, err := ioutil.ReadFile("./www/main.html")
	if err != nil {
		fmt.Println("www/main.html 을 로드할 수 없음")
		log.Fatal(err)
		panic(err)
	} else {
		fmt.Fprintf(w, string(wwwfile))
	}
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

func getIp(r *http.Request) string {
	ip := r.Header.Get("X-REAL-IP")
	netIp := net.ParseIP(ip)
	return netIp.String()
}

func urlHandle(w http.ResponseWriter, r *http.Request) {
	sv_urlpath := r.URL.Path[1:] //sv_urlpath에 유저가 어떤 url을 요청했는지 저장됨
	if sv_urlpath == "" {
		fmt.Println("Path: /", "IP주소: ", getIp(r))
		mainHanlder(w, r)
	} else if sv_urlpath == "result" {
		fmt.Println("Path: ", sv_urlpath, "IP주소: ", getIp(r))
		//결과 표시해주는 창
	} else if sv_urlpath == "upload" {
		fmt.Println("Path: ", sv_urlpath, "IP주소: ", getIp(r))
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
