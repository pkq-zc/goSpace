package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started:")
	//建立一个网络连接读取器
	reader := bufio.NewReader(conn)
	//接收数据的循环
	for {
		str, err := reader.ReadString('\n')
		if err == nil {
			str = strings.TrimSpace(str)
			//处理telnet指令
			if !processTelnetCommand(str, exitChan) {
				conn.Close()
				break
			}
			//Echo 逻辑,发什么数据,原样返回
			conn.Write([]byte(str + "\r\n"))
		} else {
			//发生错误
			fmt.Println("Session close")
			conn.Close()
			break
		}
	}
}
