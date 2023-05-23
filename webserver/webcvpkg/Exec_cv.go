package webcvpkg

import (
	"fmt"
	"os"
	"os/exec"
)

func Exec_cv(mode string, filename string) bool {
	var vexec string
	if mode == "yolo" {
		vexec = "yolo.py" //나중에 수정
	} else if mode == "haar" { //나중에 수정
		vexec = "haar.py" //나중에 수정
	}
	arg1 := "../opencv/" + vexec
	arg2 := "../files/" + filename
	//arg3 := dotFileType(filename)
	fmt.Println("실행중: python", arg1, arg2)
	cmd := exec.Command("python", arg1, arg2) // 예시: python ../opencv/main.py ive.jpeg 를 터미널에서 실행하는 것과 같은 효과임
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}
