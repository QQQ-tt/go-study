package one

import (
	"fmt"
	"os"
	"strings"
)

func osTest() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// range 产生一对值 索引和值
func osTest2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	fmt.Println(s)
}

func osTest3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// 练习 1.1： 修改 echo 程序，使其能够打印 os.Args[0]，即被执行命令本身的名字。
//
// 练习 1.2： 修改 echo 程序，使其打印每个参数的索引和值，每个一行。
func osTest4() {
	for index, arg := range os.Args {
		if index == 0 {
			fmt.Println(arg)
		} else {
			fmt.Println(index, "-", arg)
		}
	}
}
