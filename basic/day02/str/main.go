package main

import (
	"fmt"
	"strings"
)

//字符串基础的的一些操作方法
func main() {
	name := "我的狗叫小白"
	ps := "这是第三只叫小白的狗了"

	//字符串的长度
	fmt.Println(len(name)) //18

	//字符串的拼接
	ss := name + ps
	fmt.Println(ss)                      //这里是两个字符串拼接的
	ss1 := fmt.Sprintf("%s%s", name, ps) //这个返回的是一个字符串
	fmt.Println(ss1)

	//字符串的分割
	split := `file:\\\home\adsfsaft\Documents\GoWorkRoom/herasdf/basic\day02/str`
	ret := strings.Split(split, "\\")
	fmt.Println(ret) //[file:   home adsfsaft Documents GoWorkRoom/herasdf/basic day02/str]

	//某个字符出现的次数
	countsT := strings.Count(split, "a")
	countsF := strings.Count(split, "z")
	fmt.Println(countsT) //5
	fmt.Println(countsF) //0

	//字符串包含
	fmt.Println(strings.Contains(ss1, "小白"))  //找到了就返回true
	fmt.Println(strings.Contains(ss1, "大大啊")) //如果没有找到则返回false
	//判断前缀
	fmt.Println(strings.HasPrefix(ss1, "你")) //false
	//判断后缀
	fmt.Println(strings.HasSuffix(ss1, "了")) //true

	//判断字符串位置
	fmt.Println(strings.Index(ss1, "小"))     //12
	fmt.Println(strings.Index(ss1, "抽"))     //如果没有找到则返回-1
	fmt.Println(strings.LastIndex(ss1, "白")) //39
	fmt.Println(strings.LastIndex(ss1, "抽")) //如果没找到则返回-1
	//join操作
	fmt.Println(strings.Join(ret, "+")) //file:+++home+adsfsaft+Documents+GoWorkRoom/herasdf/basic+day02/str

	str3 := strings.Join([]string{"hello", "world"}, " ")
	fmt.Println("join : ", str3) //join :  hello world

}
