// zenCodeRender project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 1 { // 通过参数方式输入待处理字符串
		for _, str := range os.Args[1:] {
			fmt.Println(Expression(str))
		}
	} else { // 通过基本输入方式输入待处理字符串
		bio := bufio.NewReader(os.Stdin)
		line, _, err := bio.ReadLine()
		for {
			fmt.Println(Expression(string(line)))
			if err == nil {
				line, _, err = bio.ReadLine()
			} else {
				break
			}
		}
	}
}
