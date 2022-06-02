package main

import "fmt"

func test03() {
	fmt.Println("Test Map Using")
	var mymap map[int]string
	if mymap == nil {
		fmt.Println("map 为 null")
	}
	// 第一种声明方式
	mymap = make(map[int]string, 10)
	mymap[1] = "hello"
	mymap[2] = "world"
	mymap[-1] = "tmptmp"
	fmt.Println(mymap)
	// 第二种声明方式
	mymap2 := map[string]string{
		"hello": "hello",
		"test":  "trest",
	}
	fmt.Println(mymap2)
	mymap3 := make(map[string]string, 10)
	mymap3["no1"] = "java"
	mymap3["no2"] = "go"
	mymap3["no3"] = "c"
	fmt.Printf("长度：%d\n", len(mymap3))
	// 遍历
	for key, value := range mymap3 {
		fmt.Printf("key：%s value：%s\n", key, value)
	}
	// 删除
	delete(mymap3, "no3")
	fmt.Printf("长度：%d\n", len(mymap3))
}
