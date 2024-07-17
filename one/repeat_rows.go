package one

import (
	"bufio"
	"fmt"
	"os"
)

// 1
// 输入：1
// 1
// 输入：1
// 1
// 输入：1
// exit
// 输入：exit
// 3	1
func repeatRows() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		counts[text]++
		fmt.Println("输入：" + text)
		if input.Text() == "exit" {
			break
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
