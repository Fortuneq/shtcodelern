package main

import (
	"fmt"
	"net"
)

func main(){
	message := "Hello i am server"

	listener, err := net.Listen("tcp",":1234")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening")
	for{ conn,err:= listener.Accept()
	if err!= nil{
		fmt.Println(err)
		return
	}
	conn.Write([]byte(message))
	conn.Close()
}
}