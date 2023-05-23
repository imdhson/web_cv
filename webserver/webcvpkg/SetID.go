package webcvpkg

import (
	"math/rand"
	"net/http"
	"strconv"
)

func SetID(w http.ResponseWriter, r *http.Request, filetype string) string {
	id := rand.Intn(100000)
	cookieid := http.Cookie{
		Name:     "id",
		Value:    strconv.Itoa(id),
		HttpOnly: true,
	}
	cookietype := http.Cookie{
		Name:     "type",
		Value:    filetype,
		HttpOnly: true,
	}
	//w.Header().Set("Set-Cookie", cookieid.String())
	http.SetCookie(w, &cookieid)
	http.SetCookie(w, &cookietype)
	return strconv.Itoa(id)
}
