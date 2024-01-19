package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

// 使用 $ go run context-in-http-servers.go 来启动服务
// 接着另一个终端中访问该服务 $ curl localhost:8090/hello
// 在请求未收到返回时按ctrl+C退出，可以看到服务端的输出。
func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8098", nil)
}
