package main

import (
	"log"
	"math"
	"time"

	"github.com/tbrandon/mbserver"
)

func main() {
	server := mbserver.NewServer()

	// 启动监听 502 端口
	go func() {
		if err := server.ListenTCP("0.0.0.0:5020"); err != nil {
			log.Fatalf("启动服务器失败: %v", err)
		}
	}()
	log.Println("虚拟PLC已启动，监听 :502")

	// 模拟周期性更新寄存器数据
	for {
		t := float64(time.Now().Unix()%60) / 60 * 2 * math.Pi
		position := uint16(1000 + 500*math.Sin(t)) // 模拟位置
		load := uint16(2000 + 1000*math.Cos(t))    // 模拟载荷

		// 写入寄存器
		server.HoldingRegisters[0] = position
		server.HoldingRegisters[1] = load

		log.Printf("更新数据：Position=%d Load=%d\n", position, load)
		time.Sleep(500 * time.Millisecond)
	}
}
