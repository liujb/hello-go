package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/go-logger/logger"
)

const TIME_LAYOUT = "2006-01-2 15:04:05"
const HOST_PORT = ":7777"

func main() {
	logger.SetConsole(false)
	logger.SetRollingDaily("/Users/liujb/logs/", "aaaaa.log")
	logger.Info("fucking")

	createTcpServer()
}

// 创建TCP的server
func createTcpServer() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", HOST_PORT)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for { // 死循环监听
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// 并发处理请求
		go handleclientData(conn)

		// 并发写数据
		go writeDataToClient(conn)

		// 关闭连接
		defer conn.Close()
	}
}

/**
 * 处理请求
 */
func handleclientData(conn net.Conn) {
	reqData := make([]byte, 512)
	_, err := conn.Read(reqData)

	checkError(err)
	logger.Info(string(reqData))

	fmt.Println(string(reqData))
}

/**
 * 往客户端写数据
 */
func writeDataToClient(conn net.Conn) {
	if conn != nil {
		dayTime := time.Now().Format(TIME_LAYOUT)
		conn.Write([]byte(dayTime))

		logger.Info(dayTime)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, " Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
