package main

import (
	"flag"
	"fmt"
)

// 使用如下命令来查看结果
// $ ./command_line_flags -word=opt -numb=7 -fork -svar=flag
// $ ./command_line_flags -word=opt
// $ ./command_line_flags -word=opt a1 a2 a3
// $ ./command_line_flags -word=opt a1 a2 a3 -numb=7
// $ ./command_line_flags -h
// $ ./command_line_flags -wat

func main() {

	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
