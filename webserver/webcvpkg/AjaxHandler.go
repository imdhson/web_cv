package webcvpkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func AjaxHanlder(w http.ResponseWriter, r *http.Request, vs *[]int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, _ := GetID(w, r)
	id_int, _ := strconv.Atoi(id)
	stat := check(id_int, vs)
	var data string
	if stat {
		data = "ajax: 프로세싱 중.."
	} else {
		data = "ajax: 프로세싱 완료"
	}
	data_json, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json Marshal 중 오류 발생")
	}
	w.Write(data_json)
}

func check(id int, vs *[]int) bool { //찾으면 true 반환
	target := id
	slice := *vs
	var rst bool = false
	for _, v := range slice {
		if v == target {
			rst = true
		}
	}
	//fmt.Println("AJAX: ",id, *vs, rst)
	return rst
}
