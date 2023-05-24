package webcvpkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func AjaxHanlder(w http.ResponseWriter, r *http.Request, vs *[]VolatileStat, cv_time *time.Time) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	id_client, _ := GetID(w, r)
	data_time := strconv.Itoa(int(time.Since(*cv_time).Seconds())) + "초 지남" //cv_time 긴 시간을 초(실수)로 바꾸고 그걸 정수로 바꾸고 문자열로 바꾼것임
	stat, id_peek := check(id_client, vs)
	var data string
	if stat {
		_, _, id := Vs_peek(vs)
		data = "서버에서 " + id + "를 처리 중: " + data_time
	} else {
		data = "당신의 요청 처리 완료 - " + "서버는 현재 " + id + " 수행 중이며" + data_time
	}
	data_json, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json Marshal 중 오류 발생")
	}
	w.Write(data_json)
}

func check(id string, vs *[]VolatileStat) bool, string { //찾으면 true 반환
	target := id
	slice := *vs
	var rst bool = false
	for _, v := range slice {
		if v.CookieID == target {
			rst = true
		}
	}
	fmt.Println("AJAX: ", id, *vs, rst)
	return rst, Vs_peek(vs)
}
