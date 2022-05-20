package BasicPractice

import (
	"fmt"
	"os"
	"runtime"
)

// RunTime 输出电脑操作系统，和变量路径
func RunTime() {
	var goos string = runtime.GOOS
	fmt.Printf("这台电脑的操作系统是：%s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("这台电脑的系统路径是：%s", path)
}

var a = "G"
var b string
func N() {
	print(a)
}

func M() {
	a := "O"
	print(a)
}

func F1() {
	b := "O"
	print(b)
	F2()
}
func F2()  {
	print(b)
}