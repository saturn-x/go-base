package main

import "fmt"

type man struct {
	name string
	sex  string
}

func change(m man) {
	m.sex = "nv"
}

func main() {
	m := man{"cai", "nan"}
	fmt.Println(m)
	change(m)
	fmt.Println(m)
	var t man
	fmt.Println(t)

}
