package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"
	"reflect"
)

func main() {
	f, err := os.Open("src.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, formatName, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(formatName)
	fmt.Println(img.Bounds())
	fmt.Println(img.ColorModel())

	// 新图片
	reflect.TypeOf(img)
	reImg := img.(*image.NRGBA)
	cutByHeight(reImg)

}

// 图片剪切并保存
func cutByHeight(srcImg *image.NRGBA) {
	startX, startY := 0, 0
	// 高度剪切偏移量
	v := 0
	// 剪切高度 16:9
	height := srcImg.Bounds().Max.X / 16 * 9

	for i := 0; i < srcImg.Bounds().Max.Y-height; i++ {
		subImg := srcImg.SubImage(image.Rect(startX, startY+v, srcImg.Bounds().Max.X, startY+v+height))
		saveImage(fmt.Sprintf("images/_%09d.png", i), subImg)
		v++
	}

}

// 保存图片内容文件
func saveImage(pngFilename string, subImg image.Image) {
	nf, err := os.Create(pngFilename)
	if err != nil {
		panic(err)
	}
	defer nf.Close()

	err = png.Encode(nf, subImg)
	if err != nil {
		panic(err)
	}
}
