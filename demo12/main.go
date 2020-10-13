package main

import "fmt"

//今有雉兔同笼，上有三十五头，下有九十四足，问雉兔各几何？

const RabbitFeet = 4   //兔子有四只脚
const PheasantFeet = 2 //雉有两只脚

func main() {
	totalHead := 35 //总头数
	totalFoot := 94 //总脚数

	var x int //兔

	//4x + 2(35 - x) = 94
	//4x + 2 * 35 - 2x = 94
	//x = (94 - 2 * 35) / (4 - 2)

	x = (totalFoot - PheasantFeet*totalHead) / (RabbitFeet - PheasantFeet)

	//输出结果
	fmt.Println(fmt.Sprintf("兔 %d，雉 %d", x, totalHead-x))
}
