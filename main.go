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
		// 使用爱心曲线方程生成轨迹
		// x = 16sin³(t), y = 13cos(t) - 5cos(2t) - 2cos(3t) - cos(4t)
		heartX := 16 * math.Pow(math.Sin(t), 3)
		heartY := 13*math.Cos(t) - 5*math.Cos(2*t) - 2*math.Cos(3*t) - math.Cos(4*t)

		// 调整数值范围以适应uint16类型和更好的可视化效果
		position := uint16(1000 + 40*heartX) // 模拟位置
		load := uint16(2000 + 100*heartY)    // 模拟载荷

		// 根据readme.md定义的设备参数，分配寄存器地址
		// 写入寄存器
		server.HoldingRegisters[0] = position // 位置/距离
		server.HoldingRegisters[1] = load     // 载荷
		server.HoldingRegisters[2] = 1500     // 电机转速 (mototspeed)
		server.HoldingRegisters[3] = 5        // 冲程数 (strokesnumber)
		server.HoldingRegisters[4] = 3000     // 冲程长度 (distance)
		server.HoldingRegisters[5] = 7850     // 抽油杆密度 (roddensity)
		server.HoldingRegisters[6] = 2000     // 传动比 (transmissionRatio)
		server.HoldingRegisters[7] = 500      // 截面积 (area)
		server.HoldingRegisters[8] = 3000     // 安装倾角 (inclination)
		server.HoldingRegisters[9] = 1200     // 泵下入深度 (pumpinsertiondepth)
		server.HoldingRegisters[10] = 850     // 原油密度 (oildensity)

		log.Printf("更新数据：Position=%d Load=%d Mototspeed=%d Strokesnumber=%d Distance=%d\n",
			position, load, server.HoldingRegisters[2], server.HoldingRegisters[3], server.HoldingRegisters[4])
		time.Sleep(500 * time.Millisecond)
	}
}

//roddensity: any,             // 抽油杆密度
//transmissionRatio: any,      // 传动比
//area: any,                   // 截面积
//inclination: any,            // 安装倾角
//pumpinsertiondepth: any,     // 泵下入深度
//oildensity: any              // 原油密度
