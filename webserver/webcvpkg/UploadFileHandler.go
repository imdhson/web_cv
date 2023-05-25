package webcvpkg

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request, vs *[]VolatileStat) error {
	mode := r.FormValue("cvmode")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	file, fHeader, err := r.FormFile("originFile")
	filetype := DotFileType(fHeader.Filename)
	nowid := [2]string{SetID(w, r, filetype), filetype}
	willfileName := nowid[0] + "." + filetype
	if err != nil {
		fmt.Println("파일 수신 중 에러 발생", err)
		fmt.Fprintf(w, "Error 발생")
		return err
	} else {
		Vs_push(willfileName, mode, nowid[0], vs) //상태변수에 append
		fmt.Println(GetIP(r), "에게서 업로드된 파일이름: ", fHeader.Filename, "파일타입: ", filetype, "nowid: ", nowid)
	}
	defer file.Close()

	UploadHanlder(w, r) // upload.html을 불러옵니다.
	fileByte, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(nowid, "파일 수신 중 오류 발생")
		return err
	}
	willfilePath := "../files/" + willfileName
	//ioutil.WriteFile(willfilePath, fileByte, 0644) //deprecated
	os.WriteFile(willfilePath, fileByte, 0644) //랜덤한 id.확장자 형식으로 파일을 씁니다.
	return nil
}
