package main

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
)

func main() {
	dst := image.NewRGBA(image.Rect(0, 0, 640, 480))
	sr := image.Rect(0, 0, 200, 200)
	src := image.NewRGBA(image.Rect(0, 0, 640, 480))
	//dp := image.Point{100, 100}

	// 设置50% 透明度
	for i := 0; i < src.Bounds().Dx(); i++ {
		for j := 0; j < src.Bounds().Dy(); j++ {
			src.SetRGBA(i, j, color.RGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 128,
			})
		}
	}

	// RECT2 OMIT
	r := sr.Sub(sr.Min)
	draw.Draw(dst, r, src, sr.Min, draw.Src)
	// STOP OMIT

	// TODO
	nf, err := os.Create("nrgba.png")
	if err != nil {
		panic(err)
	}
	defer nf.Close()

	err = png.Encode(nf, dst)
	if err != nil {
		panic(err)
	}
}
