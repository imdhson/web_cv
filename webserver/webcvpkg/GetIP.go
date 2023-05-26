package webcvpkg

import (
	"net"
	"net/http"
)

func GetIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	ip_2 := net.ParseIP(ip)
	return ip_2.String()
}
