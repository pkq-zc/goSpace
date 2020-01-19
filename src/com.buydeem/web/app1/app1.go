package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	//设置访问路由
	http.HandleFunc("/",sayHelloName)
	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ",err)
	}


}

// sayHello
func sayHelloName(w http.ResponseWriter, r *http.Request) {
	//解析参数
	r.ParseForm()
	//
	fmt.Println(r)
	fmt.Println("path:",r.URL.Path)
	fmt.Println("scheme:",r.URL.Scheme)
	fmt.Println(r.Form["rul_long"])

	for k, v := range r.Form {
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello astaxie")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:",r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("src/com.buydeem/web/app1/login.gtpl")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(t.Execute(w,nil))
	}else {
		err := r.ParseForm()
		if err != nil{
			log.Fatal("ParseForm",err)
		}
		fmt.Println("username:",r.Form["username"])
		fmt.Println("password",r.Form["password"])
	}
}
