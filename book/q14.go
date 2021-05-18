package main

import "fmt"

// MaxSpeed 限速
const MaxSpeed = 120

func main() {
	// 小明车速
	speed := 0
	// 车速提升
	speed += 50
	// 判断超速没？
	if speed > MaxSpeed {
		fmt.Printf("速度：%d, 您已超速！\n", speed)
	}
	// 继续提速
	speed += 50
	// 再判断超速没？
	if speed > MaxSpeed {
		fmt.Printf("速度：%d, 您已超速！\n", speed)
	}
	// 继续提速，速度应该150了
	speed += 50
	// 再判断超速没？
	if speed > MaxSpeed {
		fmt.Printf("速度：%d, 您已超速！\n", speed)
	}
}
