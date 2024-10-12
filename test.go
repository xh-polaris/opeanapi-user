package main

import (
	"fmt"
	"github.com/xhpolaris/opeanapi-user/biz/infrastructure/util"
	"time"
)

func main() {
	id := "123456789876"
	userId := "987654321321"
	freshTime := time.Now()
	en, err := util.IssueKey(id, userId, freshTime, "xhpolarisopenapi")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(en)
	a, b, c, err := util.ParseKey(en)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
