// brainfuck project main.go
package main

import (
	"fmt"
)

var (
	a     [6172]byte // VM内存
	p, pc int        // 数据指针, 指令计数器
	code  string     // 程序代码
)

func loop(inc int) {
	for i := inc; i != 0; pc += inc {
		switch code[pc+inc] {
		case '[':
			i++
		case ']':
			i--
		}
	}
}

func reset() {
	p, pc = 0, 0
	for i := 0; i < 6161; i++ {
		a[i] = 0
	}
}

func fuck() {
	var in byte
	for {
		// 指令系统
		switch code[pc] {
		case '>':
			p++
		case '<':
			p--
		case '+':
			a[p]++
		case '-':
			a[p]--
		case '.':
			fmt.Print(string(a[p]))
		case ',':
			fmt.Scanf("%c", &in)
			a[p] = in
		case '[':
			if a[p] == 0 {
				loop(1)
			}
		case ']':
			if a[p] != 0 {
				loop(-1)
			}
		default:
			fmt.Print("语法错误！")
		}
		pc++ // 指令计数器+1

		if pc == len(code) {
			return
		}
	}
}

func main() {
	for {
		reset()
		fmt.Print("BrainFuck > ")
		fmt.Scanln(&code)
		fuck()
	}
}
