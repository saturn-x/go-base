package main

import "fmt"

// 测试多态
type Animal interface {
	eat()
	call() string
}

type dog struct {
	name  string
	color string
}

func (this *dog) eat() {
	fmt.Println("狗在吃")
}
func (this *dog) call() string {
	fmt.Println("狗在叫")
	return "Dog"
}

type cat struct {
	name  string
	color string
}

func (this *cat) eat() {
	fmt.Println("猫在吃")
}
func (this *cat) call() string {
	fmt.Println("猫在叫")
	return "Cat"
}

func polymorphic() {
	// 指针
	var animal Animal
	mydog := dog{"cai", "red"}
	// 必须使用指针
	animal = &mydog
	animal.eat()
	animal.call()
	animal = &cat{"li", "green"}
	animal.eat()

}
