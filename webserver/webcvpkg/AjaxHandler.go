package webcvpkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func AjaxHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, _ := GetID(w, r)
	stat := check(id, VolatileStat)
	var data string
	if stat {
		data = "프로세싱 중.."
	} else {
		data = "프로세싱 완료"
	}
	data_json, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json marshal중 오류 발생")
	}
	w.Write(data_json)
}

func check(id int, slice []int) bool { //찾으면 true 반환
	target := id
	var rst bool = false
	for i, v := range slice {
		if v == target {
			rst = true
		}
	}
	return rst
}
