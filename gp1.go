package main

import (
	"fmt"
	"http://git.chinaase.com/stu_self/gopbf/user"
)

type sp string

func main01() {
	var s sp
	s = "hi"
	fmt.Println(&s, "\n", s)
	s1 := &s
	*s1 = "get"
	s2 := &s1
	fmt.Println("s内存地址:", &s, "s的值:", s, "s1的值:", s1, "s1的内存地址:", &s1, "s2的内存地址:", &s2, "s2的值:", &s1)
}

func main() {
	u := user.User{}
}
