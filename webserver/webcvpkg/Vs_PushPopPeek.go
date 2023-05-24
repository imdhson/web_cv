package webcvpkg

import "fmt"

func Vs_push(path string, mode string, id string, vs *[]VolatileStat) {
	slice := *vs
	fmt.Println("VS Before push: ", *vs)
	slice = append(slice, VolatileStat{path, mode, id})
	*vs = slice
	fmt.Println("VS After push: ", *vs)
}

func Vs_pop(vs *[]VolatileStat) (string, string, string) {
	slice := *vs
	fmt.Println("VS Before pop: ", *vs)
	i := 0
	a, b, c := slice[i].Filename, slice[i].CvMode, slice[i].CookieID
	slice = slice[0 : len(slice)-1]
	*vs = slice
	fmt.Println("VS After pop: ", *vs)
	return a, b, c
}

func Vs_peek(vs *[]VolatileStat) (string, string, string) {
	vs_p := *vs
	i := 0
	a, b, c := vs_p[i].Filename, vs_p[i].CvMode, vs_p[i].CookieID
	return a, b, c
}
