package webcvpkg

import "net/http"

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
