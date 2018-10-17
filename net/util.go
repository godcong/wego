package net

import (
	"encoding/json"
	"net"
	"net/http"
)

/*WriteJSON 将参数obj作为json返回 */
func WriteJSON(w http.ResponseWriter, status int, obj interface{}) error {
	w.WriteHeader(status)
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	w.Write(jsonBytes)
	return nil
}

/*GetServerIP 获取服务端IP */
func GetServerIP() string {
	adds, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}

	for _, address := range adds {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return "127.0.0.1"
}

/*GetClientIP 取得客户端IP */
func GetClientIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil && ip != "127.0.0.1" {
		return ip
	}
	ip = r.Header.Get("X-Forwarded-For")
	if ip == "" {
		return "127.0.0.1"
	}
	return ip
}
