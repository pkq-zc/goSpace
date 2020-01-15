package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Address   []*Address
	Remark    string
}

// 序列化 对象 => json
func main() {
	sz := &Address{"private", "SenZhen", "China"}
	home := &Address{"home", "HengYang", "China"}
	vc := VCard{"Zeng", "Chao", []*Address{sz, home}, "none"}

	// JSON format
	jsonStr, _ := json.Marshal(vc)
	fmt.Printf("JSON format : %s\n", jsonStr)
	//Write to File 1
	error := ioutil.WriteFile("vard.json", jsonStr, 0666)
	if error != nil {
		log.Fatal(error)
	}
	//Write to file 2
	file, _ := os.OpenFile("vard2.json", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	encoder := json.NewEncoder(file)
	error = encoder.Encode(vc)
	if error != nil {
		log.Println("Error in encoding json")
	}

}
