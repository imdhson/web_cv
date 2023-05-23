package webcvpkg

import (
	"fmt"
	"net/http"
	"os"
)

func KitHanlder(w http.ResponseWriter, r *http.Request, path string) error {
	docsfile := DotFileType(path)
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
	case "mp4":
		w.Header().Set("Content-Type", "video/mp4; charset=utf-8")
	case "avi":
		w.Header().Set("Content-Type", "video/avi; charset=utf-8")
	case "mov":
		w.Header().Set("Content-Type", "video/mov; charset=utf-8")
	case "webm":
		w.Header().Set("Content-Type", "video/webm; charset=utf-8")
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
