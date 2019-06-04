package main

import (
	"net"
	"log"
	"fmt"
)

func main()  {
	port := "9091"
	Start(port)
}

func Start(port string)  {
	host := ":"+ port
	tcpAddr,err := net.ResolveTCPAddr("tcp",host)
	if err != nil {
		log.Printf("resolve tcp add failed : %v \n",err)
		return
	}

	//监听
	listener,err := net.ListenTCP("tcp",tcpAddr)
	log.Printf("listen tcp port failed:%v \n",err)

	//建立连接池，用于广播
	conns := make(map[net.Conn]string)

	//消息通道
	messageChan := make(chan string ,10)

	//广播消息

	for  {
		fmt.Printf(" listening port %s... \n",port)
		conn,err := listener.AcceptTCP()
		if err != nil {
			log.Printf("Accept failed:%v \n",err)
			return
		}
		conns[conn] = conn.RemoteAddr().String()
		fmt.Println(conns)
		//处理消息
		go Handler(conn,conns,messageChan)
	}
}

func Handler(conn net.Conn,conns map[net.Conn]string,messages chan string)  {

}
