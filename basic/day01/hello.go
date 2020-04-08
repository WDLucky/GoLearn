package main // 声明 main 包，表明当前是一个可执行程序
import (
	"fmt"
	"strings"
) // 导入内置 fmt 包

func main() { // main函数，是程序执行的入口
	a := "我是狗abcdefg1"
	fmt.Println(len(a))                //17
	fmt.Println(strings.Index(a, "我")) //0
	fmt.Println(strings.Index(a, "是")) //3
	fmt.Println(strings.Index(a, "狗")) //6
	fmt.Println(strings.Index(a, "a")) //9
	fmt.Println(strings.Index(a, "1")) //16

}
