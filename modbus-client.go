package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/goburrow/modbus"
)

func main() {
	// 检查命令行参数的数量
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s [register address] [number of registers]\n", os.Args[0])
	}

	// 读取命令行参数
	registerAddress, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid register address: %v", err)
	}
	numRegisters, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid number of registers: %v", err)
	}

	// 创建 Modbus TCP 客户端连接到服务器
	handler := modbus.NewTCPClientHandler("localhost:502")
	// handler.Timeout = modbus.SerialTimeout // 5 seconds
	handler.SlaveId = 1
	defer handler.Close()
	if err := handler.Connect(); err != nil {
		log.Fatalf("Error connecting: %v", err)
	}

	client := modbus.NewClient(handler)

	// 根据提供的寄存器地址和数量读取 Holding Registers
	results, err := client.ReadHoldingRegisters(uint16(registerAddress), uint16(numRegisters))
	if err != nil {
		log.Fatalf("Error reading from server: %v", err)
	}

	// 遍历数组并打印每个字节的十六进制表示
	for i, b := range results {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%02X", b)
	}
	fmt.Println()
}
