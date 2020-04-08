package main

import (
	"fmt"
)

func main() {
	s := "hello小白"
	//104(h) 101(e) 108(l) 108(l) 111(o) 229(å) 176(°) 143() 231(ç) 153() 189(½)
	for i := 0; i < len(s); i++ {
		//byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()

	//104(h) 101(e) 108(l) 108(l) 111(o) 23567(小) 30333(白)
	for _, r := range s {
		//rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()

	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1)) //pig

	runeS2 := []rune(s)
	fmt.Println(len(runeS2))
	runeS2[6] = '黑'
	fmt.Println(string(runeS2)) //hello小黑
	fmt.Println(runeS2)         //[104 101 108 108 111 23567 40657]

	str := "hello沙河小王子"
	count := 0
	for _, s := range str {
		fmt.Printf("%c\n", s)
		if len(string(s)) >= 3 {
			count++
		}
	}
	fmt.Printf("\"%s\"中的汉字数量是:%d", str, count)

}
