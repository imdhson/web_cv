package webcvpkg

func DotFileType(in string) string { //파일 이름을 받으면 . 이후의 확장자만 리턴하여 줍니다.
	in2 := []rune(in)
	for i := len(in2) - 1; i >= 0; i-- { //파일 중간에 . 이 들어가는 경우가 있어서 뒤부터 순회
		v := string(in2[i])
		if v == "." {
			return string(in2[i+1:])
		}
	}
	return "None"
}
