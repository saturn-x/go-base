package main

import "fmt"

type Human struct {
	Sex string
}

type Hero struct {
	Human
	Name  string
	Age   int
	Level int
}
type Book struct {
	title string
	auth  string
}

// 使用指针 就可以使用引用传递
func (this *Hero) getName() {
	fmt.Println(this.Name)
}
func (this *Hero) setName(name string) {
	this.Name = name
}

func object() {
	fmt.Println("使用结构体·····")
	// 简单使用结构体
	var mybook Book
	fmt.Println(mybook)
	mybook.title = "nihao"
	mybook.auth = "nihao"
	fmt.Println(mybook)
	// 使用成员函数
	hero := Hero{Name: "caixukun", Age: 10, Level: 1}
	hero.setName("lixiao")
	hero.getName()
	fmt.Println(hero)
	// 使用子类中的属性
	hero2 := Hero{Human{"nv"}, "cai", 10, 1}
	fmt.Println(hero2)
}
