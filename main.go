package main

import (
	"Go-Stack/stack"
	"fmt"
)

func main() {
	var s = stack.New(3)
	s.Push(5)
	s.Push(1)
	s.Push(4)
	s.Push(2)
	for !s.Empty() {
		fmt.Println(s.Pop())
	}
	fmt.Println("--------------------------")
	s.Push("I Love You")
	s.Push("无法不爱你Baby~")
	s.Push("说你也爱我")
	s.Push(1234)
	s.Push("ILoveYou~~~~")
	for !s.Empty() {
		fmt.Println(s.Pop())
	}
}
