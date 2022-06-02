package main

import "fmt"

func Demo03Base() {
	fmt.Println("测试一些基本数据")
	fmt.Println("变量声明")

	var a int8 = 59
	fmt.Printf("%b\n", a) //输出二进制：111011
	fmt.Printf("%d\n", a) //输出十进制：59
	fmt.Printf("%o\n", a) //输出八进制：73
	fmt.Printf("%O\n", a) //输出八进制(带0o前缀)：0o73
	fmt.Printf("%x\n", a) //输出十六进制(小写)：3b
	fmt.Printf("%X\n", a) //输出十六进制(大写)：3B

	str := "hello world"
	str = "你好，世界！"
	fmt.Printf("len:%d", len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d ~ %c\n", i, str[i])
	}

	for i, t := range str {
		fmt.Printf("%d ~ %c\n", i, t)
	}
}
