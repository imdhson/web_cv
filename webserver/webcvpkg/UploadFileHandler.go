package webcvpkg

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) error {
	mode := r.FormValue("cvmode")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	file, fHeader, err := r.FormFile("originFile")
	filetype := DotFileType(fHeader.Filename)
	nowid := [2]string{SetID(w, r, filetype), filetype}
	if err != nil {
		fmt.Println("파일 수신 중 에러 발생", err)
		fmt.Fprintf(w, "Error 발생")
		return err
	} else {
		w.Write([]byte("<meta http-equiv=\"refresh\" content=\"0;url=/upload\">"))
	}
	nowid_int, _ := strconv.Atoi(nowid[0])
	push(nowid_int, prs) //상태변수에 append

	fmt.Println(GetIP(r), "에게서 업로드된 파일이름: ", fHeader.Filename, "파일타입: ", filetype, "nowid: ", nowid)
	defer file.Close()
	fileByte, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(nowid, "파일 수신 중 오류 발생")
	}
	willfileName := nowid[0] + "." + filetype
	willfilePath := "../files/" + willfileName
	ioutil.WriteFile(willfilePath, fileByte, 0644) //랜덤한 id.확장자 형식으로 파일을 씁니다.

	success := Exec_cv(mode, willfileName)
	if success {
		pop(nowid_int, VolatileStat)
		fmt.Println("Opencv 처리완료")
	} else {
		fmt.Println("Opencv 처리 중 오류 발생")
	}
	return nil
}

func push(id int, slice []int) []int {
	slice = append(slice, id)
}
func pop(id int, slice []int) []int {
	target := id
	for i, v := range slice {
		if v == target {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
