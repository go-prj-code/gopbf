package main

import (
	"encoding/json"
	"fmt"
	"gopbf/user"
	"log"

	"google.golang.org/protobuf/proto"
	// "git.chinaase.com/stu_self/gopbf/user"
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
	u.Password = "123456"
	u.UserName = "star"
	b, _ := json.Marshal(u)
	fmt.Println(string(b))

	//序列化
	out, err := proto.Marshal(&u)
	if err != nil {
		log.Fatalln("failed to:", err)
	}
	fmt.Println(out)

	//反序列化
	i := user.User{}
	err = proto.Unmarshal(out, &i)
	if err != nil {
		log.Fatalln("failed to:", err)

	}
	b2, _ := json.Marshal(i)
	fmt.Println(b2)
}
