package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

// 反序列化 json => 对象
func main() {
	jsonStr := "{\"Name\":\"tom\",\"Age\":18,\"Parents\":[\"Gomez\",\"Morticia\"]}"
	// 转化成Map
	var f interface{}
	err := json.Unmarshal([]byte(jsonStr), &f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)

	// 转化成对象
	var m FamilyMember
	err = json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", m)
}
