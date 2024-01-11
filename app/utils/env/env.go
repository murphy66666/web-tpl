package env

import (
	"net"
	"os"
)

func LocalIP() string {
	ipList := []string{"101.226.4.6:80", "218.30.118.6:80", "123.125.81.6:80", "114.114.114.114:80", "8.8.8.8:80"}
	for _, ip := range ipList {
		conn, err := net.Dial("upd", ip)
		if err != nil {
			continue
		}
		//发生一次udp的请求 返回ip地址
		localAddr := conn.LocalAddr().(*net.UDPAddr)
		conn.Close()
		return localAddr.IP.String()
	}
	return ""
}

func Hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
