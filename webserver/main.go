package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func exec_cv(mode string, filename string) bool {
	var vexec string
	if mode == "hog" {
		vexec = "main.py" //나중에 수정
	} else if mode == "haar" { //나중에 수정
		vexec = "haar.py" //나중에 수정
	}
	arg1 := "../opencv/" + vexec
	arg2 := "../files/" + filename
	//arg3 := dotFileType(filename)
	fmt.Println("실행중: python3", arg1, arg2)
	cmd := exec.Command("python3", arg1, arg2) // 예시: python ../opencv/main.py ive.jpeg 를 터미널에서 실행하는 것과 같은 효과임
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
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
		fmt.Println("www/main/main.html 을 로드할 수 없음")
		log.Fatal(err)
		panic(err)
	} else {
		fmt.Fprintf(w, string(wwwfile))
	}
}

func kitHanlder(w http.ResponseWriter, r *http.Request, path string) error {
	docsfile := dotFileType(path)
	switch docsfile {
	case "jpg":
		w.Header().Set("Content-Type", "image/jpg; charset=utf-8")
	case "png":
		w.Header().Set("Content-Type", "image/png; charset=utf-8")
	case "jpeg":
		w.Header().Set("Content-Type", "image/jpeg; charset=utf-8")
	case "js":
		w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	case "css":
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	case "scss":
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	case "html":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	}
	wwwfile, err := os.ReadFile("./www/kit" + path)
	if err != nil {
		fmt.Println("www/kit 을 로드할 수 없음", path)
		return err
	} else {
		w.Write(wwwfile)
		return nil
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) error {
	mode := r.FormValue("cvmode")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	file, fHeader, err := r.FormFile("originFile")
	if err != nil {
		fmt.Println("파일 수신 중 에러 발생", err)
		fmt.Fprintf(w, "Error 발생")
		return err
	}
	filetype := dotFileType(fHeader.Filename)
	nowid := [2]string{setID(w, r, filetype), filetype}
	fmt.Println(getIp(r), "에게서 업로드된 파일이름: ", fHeader.Filename, "파일타입: ", filetype, "nowid: ", nowid)
	defer file.Close()
	fileByte, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(nowid, "파일 수신 중 오류 발생")
	}
	willfileName := nowid[0] + "." + filetype
	willfilePath := "../files/" + willfileName
	ioutil.WriteFile(willfilePath, fileByte, 0644) //랜덤한 id.확장자 형식으로 파일을 씁니다.
	success := exec_cv(mode, willfileName)
	if success {
		fmt.Fprintf(w, "<meta http-equiv=\"refresh\" content=\"0;url=/result\">")
	}
	return nil
}

func resultHanlder(w http.ResponseWriter, r *http.Request) error {

	nowid, filetype := getID(w, r)
	willfileName := nowid + "." + filetype
	willfilePath := "../files/" + willfileName
	file, err := ioutil.ReadFile(willfilePath)
	if err != nil {
		fmt.Println(willfilePath, " 로드할 수 없음")
		//log.Fatal(err)
	} else {
		w.Header().Set("Content-Type", "image/jpeg; charset=utf-8") // 이미지인 경우
		w.Write(file)
	}
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

func getID(w http.ResponseWriter, r *http.Request) (string, string) {
	id, err := r.Cookie("id")    //key to value로 쿠키를 가져옴
	ctype, _ := r.Cookie("type") //key to value로 쿠키를 가져옴
	if err != nil {
		//쿠키가 없으니 nil 리턴
		return "", ""
	} else {
		//쿠키를 기반으로 결과창으로 넘기기 위해 값을 리턴
		return id.Value, ctype.Value
	}
}

func setID(w http.ResponseWriter, r *http.Request, filetype string) string {
	id := rand.Intn(100000)
	cookieid := http.Cookie{
		Name:     "id",
		Value:    strconv.Itoa(id),
		HttpOnly: true,
	}
	cookietype := http.Cookie{
		Name:     "type",
		Value:    filetype,
		HttpOnly: true,
	}
	//w.Header().Set("Set-Cookie", cookieid.String())
	http.SetCookie(w, &cookieid)
	http.SetCookie(w, &cookietype)
	return strconv.Itoa(id)
}

func getIp(r *http.Request) string {
	ip := r.Header.Get("X-REAL-IP")
	netIp := net.ParseIP(ip)
	return netIp.String()
}

func urlHandler(w http.ResponseWriter, r *http.Request) {
	sv_urlpath := r.URL.Path[1:] //sv_urlpath에 유저가 어떤 url을 요청했는지 저장됨
	if sv_urlpath == "" {
		fmt.Println("Path: /", "IP주소: ", getIp(r))
		mainHanlder(w, r)
	} else if len(sv_urlpath) > 1 && sv_urlpath[0:3] == "kit" { //첫 조건에 길이 확인이 있는 이유는 인덱스 초과 슬라이싱을 막기 위함.
		fmt.Println("Path:/kit" + sv_urlpath[3:]) // kit 폴더를 가져오게됨
		err := kitHanlder(w, r, sv_urlpath[3:])
		if err != nil {
			errHander(w, r)
		}
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
	server.Handle("/", http.HandlerFunc(urlHandler))
	fmt.Println("http://localhost:"+strconv.Itoa(PORT), "에서 요청을 기다리는 중:")
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), server)
	if err != nil { // http 서버 시작 중 문제 발생시
		log.Fatal(err)
		panic(err)
	}
}
