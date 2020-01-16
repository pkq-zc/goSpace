package main

import (
	"fmt"
	"net"
)

func server(address string, exitChan chan int) {
	// 根据给定地址侦听
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}

	// 打印侦听地址
	fmt.Println("listen: " + address)
	// 延迟关闭侦听
	defer listener.Close()

	// 循环侦听
	for {
		//新连接没有来时,Accept阻塞
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// 根据连接开启会话
		go handleSession(conn, exitChan)
	}

}
