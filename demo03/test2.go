package main

import "fmt"

func test2() {
	fmt.Println("hello world")
	var arr1 [10]int
	var a = [10]int{1, 2, 3, 4, 5}
	b := [10]int{1, 2}

	for i := 0; i < len(a); i++ {
		fmt.Printf("%d \n", a[i])
	}
	fmt.Println("测试声明方式")
	for _, i := range arr1 {
		fmt.Println(i)
	}
	for _, i := range b {
		fmt.Println(i)
	}
	// 切片数组
	myarr := []int{1, 2, 3}
	var slice []int
	slice = make([]int, 3)
	slice[0] = 1
	slice2 := make([]int, 5)
	fmt.Printf("%v ： %v\n", myarr, slice)
	fmt.Printf("%v", slice2)

	// 切片追加 和 截取
	// len 3 cap 5 追加5之后就会进行重新开辟空间
	newslice := make([]int, 3, 5)
	// 追加
	newslice = append(newslice, 1)
	fmt.Println(len(newslice))
	fmt.Println(newslice)
	// 左闭又开
	fmt.Println(newslice[0:1])
	fmt.Println(newslice[0:2])
	// 截取全部
	fmt.Println(newslice[0:])
	newslice2 := make([]int, 4)
	copy(newslice2, newslice)
	fmt.Println(newslice2)

}
