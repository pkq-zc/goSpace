package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	// 创建通道
	ch := make(chan string)

	go RPCServer(ch)
	recv, err := RPCClient(ch, "hi")
	if err != nil {
		//发生错误打印
		fmt.Println(err)
	} else {
		fmt.Println("client receive:", recv)
	}
}

func RPCClient(ch chan string, req string) (string, error) {
	//向服务器发送请求
	ch <- req
	//等待服务器返回
	select {
	//接收服务器返回数据
	case ack := <-ch:
		return ack, nil
	// 超时
	case <-time.After(time.Second):
		return "", errors.New("Time out")
	}
}

func RPCServer(ch chan string) {
	for true {
		//接收客户端请求
		data := <-ch

		//打印接收的数据
		fmt.Println("Server receive:", data)
		//向客户端反馈已收到
		time.Sleep(2 * time.Second)
		ch <- "roger"
	}
}
