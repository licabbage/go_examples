package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second) // timer只会触发一次

	<-timer1.C
	fmt.Println("Timer 1 fired.")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired.")
	}()

	time.Sleep(2 * time.Second) //这里sleep一会，看Timer 2 fired会不会输出

	stop2 := timer2.Stop() // 如果成功停止闹钟, 返回ture,如果失败则返回false
	if stop2 {
		fmt.Println("Timer 2 stopped.")
	} else {
		fmt.Println("Timer 2 has fired.")
	}

	time.Sleep(2 * time.Second)
}
