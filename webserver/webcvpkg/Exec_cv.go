package webcvpkg

import (
	"fmt"
	"os"
	"os/exec"
)

func Exec_cv(filename string, mode string) bool {
	var vexec string
	switch mode {
	case "yolo":
		vexec = "yolo.py"
	case "haar":
		vexec = "haar.py"
	case "mosaic":
		vexec = "mosaic.py"
	case "edsr":
		vexec = "edsr.py"
	case "espcn":
		vexec = "espcn.py"
	case "grayscale":
		vexec = "grayscale.py"
	case "canny":
		vexec = "canny.py"
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
