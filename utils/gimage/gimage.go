package gimage

import (
	"image"
)

type MyImage struct {
	Mode    int
	Width   float32
	Height  float32
	Quality int
	Format  []string
	Rotate  int
	Strip   []string
	Wm      WaterMark
}

type WaterMark struct {
	im       image.Image
	Dissolve int
	Gravity  []string
	x        float32
	y        float32
}

//var im = &image.RGBA{}
//
//func resize(i image.Image, w float32, h float32, mode int) image.RGBA {
//	f,err := os.Open("a.jpg")
//	png.Decode(f)
//	defer f.Close()
//	image.Pt()
//
//}
