package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"student_name"`
	Age  int
}

func main() {
	// 数组
	x := [5]int{1, 2, 3, 4, 5}
	s, err := json.Marshal(x)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(s))

	//map类型
	m := make(map[string]float64)
	m["zhangsan"] = 100.4
	s2, err2 := json.Marshal(m)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(string(s2))

	//json编码
	student := Student{"张三", 24}
	s3, err3 := json.Marshal(student)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(string(s3))

	var s4 interface{}
	json.Unmarshal([]byte(s3), &s4)
	fmt.Println("%v", s4)

}
