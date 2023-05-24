package webcvpkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ResultFileHanlder(w http.ResponseWriter, r *http.Request) error {

	nowid, filetype := GetID(w, r)
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
