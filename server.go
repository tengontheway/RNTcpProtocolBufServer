package main

import (
	"fmt"
	"net"

	"encoding/binary"

	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Starting the server...192.168.0.162:8888")
	//create listener
	listener, err := net.Listen("tcp", "192.168.0.162:8888")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	// listen and accept connections from clients:
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			return
		}
		//create a goroutine for each request.
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	fmt.Println("new connection:", conn.LocalAddr())
	for {
		buf := make([]byte, 1024)
		length, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		fmt.Println(buf[:length])

		// Unpack msg [MsgLen:2 + MsgID:1 + MsgData]
		// MsgLen: 1short = byte(MsgID).length + byte(MsgData).length
		// MsgID: 1byte MsgID
		// MsgData: byte[]
		totalLen := binary.BigEndian.Uint16(buf[0:2])
		fmt.Println("TotalLen", totalLen)

		id := binary.BigEndian.Uint16(buf[2:4])
		fmt.Println("ID", id)

		data := buf[4:totalLen]

		msgLogin := &Login{}
		err = proto.Unmarshal(data, msgLogin)
		if err != nil {
			fmt.Println("parse succeed! seed:", msgLogin.Seed)
		} else {
			fmt.Println("err", err)
		}

		fmt.Println("Receive data from client:", string(buf[:length]))
	}
}
