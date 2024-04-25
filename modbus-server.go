package main

import (
	"log"

	"github.com/tbrandon/mbserver"
)

func main() {
	// Create a new Modbus TCP server
	serv := mbserver.NewServer()
	defer serv.Close()

	// Set the listening address to localhost port 5020
	servAddress := "0.0.0.0:502"
	err := serv.ListenTCP(servAddress)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// Set up some default values
	for i := 0; i < 100; i++ {
		//写入4字节
		serv.HoldingRegisters[i] = uint16(i)
	}

	//打印服务器监听地址和端口
	log.Printf("Modbus server listening on %s", servAddress)
	log.Println("Modbus server started.")

	// Block forever
	select {}
}
