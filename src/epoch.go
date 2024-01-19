package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	fmt.Println(time.Unix(now.Unix(), 0)) // time.Unix，第一个参数是秒，第二个参数是nano seconds（纳秒）
	fmt.Println(time.Unix(0, now.UnixNano()))
}
