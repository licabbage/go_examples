package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// 一个url的结构为 scheme://host:port/path?query#fragment
	// schema 为协议，如ftp, http, https
	// host为主机，可以是域名或者Ip地址
	// port为端口，这个是可选参数，如果未指定，则使用默认端口，如http默认端口为80， https的默认端口为443
	// path是服务器上资源的位置
	// query是查询参数，以?开头
	// fragment是片段, 它用于指定资源中的特定部分或锚点。通常以#开头，用于跳转到页面内的特定位置。

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
