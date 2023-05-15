package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func exec_cv() {
	cmd := exec.Command("ls", "-al")
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func errHander(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	wwwfile, err := ioutil.ReadFile("./www/err.html")
	if err != nil {
		fmt.Println("www/err.html 을 로드할 수 없음")
		log.Fatal(err)
		panic(err)
	} else {
		fmt.Fprintf(w, string(wwwfile))
	}
}

func mainHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
	nowid := setID(w, r)
	mode := r.FormValue("cvmode")
	file, fHeader, err := r.FormFile("originFile")
	if err != nil {
		fmt.Println("파일 수신 중 에러 발생", err)
		fmt.Fprintf(w, "Error 발생")
		return err
	}
	exec_cv(mode)
	filetype := dotFileType(fHeader.Filename)
	fmt.Println(getIp(r), "에게서 업로드된 파일이름: ", fHeader.Filename, "파일타입: ", filetype)
	defer file.Close()
	fileByte, err := ioutil.ReadAll(file)
	willfilePath := "../files/" + strconv.Itoa(nowid) + "." + filetype
	ioutil.WriteFile(willfilePath, fileByte, 0644) //랜덤한 id.확장자 형식으로 파일을 씁니다.
	return nil
}

func resultHanlder(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	nowid := getID(w, r)
	fmt.Fprintf(w, "이 세션의 고유 번호: "+string(nowid))
	return nil
}

func dotFileType(in string) string { //파일 이름을 받으면 . 이후의 확장자만 리턴하여 줍니다.
	in2 := []rune(in)
	for i := len(in2) - 1; i >= 0; i-- { //파일 중간에 . 이 들어가는 경우가 있어서 뒤부터 순회
		v := string(in2[i])
		if v == "." {
			return string(in2[i+1:])
		}
	}
	return "None"
}

func getID(w http.ResponseWriter, r *http.Request) string {
	id, err := r.Cookie("id") //key to value로 쿠키를 가져옴
	if err != nil {
		//쿠키가 없으니 nil 리턴
		return "empty"
	} else {
		//쿠키를 기반으로 결과창으로 넘기기 위해 값을 리턴
		return id.Value
	}
}

func setID(w http.ResponseWriter, r *http.Request) int {
	id := rand.Intn(100000)
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
		resultHanlder(w, r)
	} else if sv_urlpath == "upload" {
		fmt.Println("Path: ", sv_urlpath, "IP주소: ", getIp(r))
		uploadHandler(w, r)
	} else {
		fmt.Println("잘못된 Path: ", sv_urlpath, "IP주소: ", getIp(r))
		errHander(w, r)
	}
}

func main() {
	const PORT int = 8080
	server := http.NewServeMux()
	server.Handle("/", http.HandlerFunc(urlHandle))
	fmt.Println("http://localhost:"+strconv.Itoa(PORT), "에서 요청을 기다리는 중:")
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), server)
	if err != nil { // http 서버 시작 중 문제 발생시
		log.Fatal(err)
		panic(err)
	}
}
