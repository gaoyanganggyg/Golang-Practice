package main

import (
	"net"
	"log"
	"io"
	"time"
)

func handleConn(c net.Conn)  {
	defer c.Close()
	for  {
		_, err := io.WriteString(c, time.Now().Format("2006-01-02 15:04:05Z07:00\n"))
		if nil != err{
			return
		}
		time.Sleep(1*time.Second)
	}
}


func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if nil != err{
		log.Fatal(err)
	}
	for  {
		conn, err := listener.Accept()
		if nil != err{
			log.Println(err)
			continue
		}
		handleConn(conn)
	}
}
