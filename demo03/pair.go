package main

import (
	"fmt"
)

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type human struct {
	name string
}

func (this *human) ReadBook() {
	fmt.Println("human  read  a book")
}

func (this *human) WriteBook() {
	fmt.Println("human  write  a book")
}

func pair() {

	var r Reader
	fmt.Println(r)

	man := &human{"hello"}
	r = man
	fmt.Println(r)

	var w Writer
	w = r.(Writer) // 强制类型转换
	w.WriteBook()

}
