package main

import "fmt"

type com struct {
	id   int
	name string
}

func main() {
	a := com{
		id:   123,
		name: "123",
	}
	b := a
	fmt.Printf("%p,%p\n", &a, &b)
	fmt.Println(b)
	bbb(&a)
	fmt.Println(b)
	fmt.Println(a)
}

func bbb(com2 *com) {
	b := *com2
	b.id = 456
	fmt.Printf("%p,%p", &b, &com2)
	fmt.Println("bbbbb0", b)
}
