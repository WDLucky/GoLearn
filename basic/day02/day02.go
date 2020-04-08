package main

import "fmt"

func main() {

	type Dog struct {
		Name string
	}

	var (
		a = 100
		b = Dog{Name: "白狗"}
		c = 0xFF
		d = 0xff
		e = 99.9
	)

	fmt.Printf("%T\n", a) //显示相应的值得类型
	fmt.Printf("%d\n", a) //十进制显示值
	fmt.Printf("%b\n", a) //二进制显示值
	fmt.Printf("%o\n", a) //八进制显示值

	//十六进制是以0x开头
	fmt.Printf("%x\n", d) //十六进制表示，字母形式为小写 a-f
	fmt.Printf("%X\n", c) //十六进制表示，字母形式为大写 A-F

	fmt.Printf("++++++++++++++++++++++++++++\n")
	fmt.Printf("%v\n", b)  //显示相应默认的格式
	fmt.Printf("%+v\n", b) //显示相应的字段名
	fmt.Printf("%#v\n", b) //相应值的类型语法表示
	fmt.Printf("%f\n", e)  //浮点数
}
