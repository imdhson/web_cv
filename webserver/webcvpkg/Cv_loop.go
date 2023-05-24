package webcvpkg

import "time"

func Cv_loop(vs *[]VolatileStat, cv_ing *bool, cv_time *time.Time) {
	if len(*vs) > 0 && !*cv_ing { //cv가 하나라도 작동중이면 실행 안함
		filename, mode, _ := Vs_peek(vs)
		*cv_ing = true
		*cv_time = time.Now()
		Exec_cv(filename, mode)
		*cv_ing = false
		Vs_pop(vs)
		*cv_time = time.Now()
	}
}
