package webcvpkg

import (
	"net"
	"net/http"
)

func GetIP(r *http.Request) string {
	ip := r.Header.Get("X-REAL-IP")
	netIp := net.ParseIP(ip)
	return netIp.String()
}
