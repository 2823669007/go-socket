package main

import (
	"net"
	"log"
	"github.com/go-ffmt/ffmt"
	"fmt"
	"os"
)

func main()  {
	Start(":9091")
}

func Start(tcpAddrStr string)  {
	tcpAddr,err := net.ResolveTCPAddr("tcp",tcpAddrStr)
	if err != nil {
		log.Printf("Resolve tcpAddr failed:%v \n",err)
		return
	}

	//向服务器拨号
	conn,err := net.DialTCP("tcp",nil,tcpAddr)
	if err != nil {
		log.Printf("Dial to server failed: %v \n",err)
		return
	}
	ffmt.P(conn)

	go SendMsg(conn)
}

func SendMsg(conn net.Conn)  {
	username := conn.RemoteAddr().String()

	for  {
		var input string
		//接收输入消息，放到input变量中
		fmt.Scanln(&input)

		if input == "/q" || input == "/quit" {
			fmt.Println("Byebye ...")
			conn.Close()
			os.Exit(0)
		}

		//只处理消息内容的消息
		if len(input) >0 {
			msg := username +  " say: " + "input"
			_,err := conn.Write([]byte(msg))
			if err != nil {
				conn.Close()
				break
			}
		}
	}
}