package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/lucasb-eyer/go-colorful"
)

func fabs(x int) int {
	if x > 0 {
		return x
	} else {
		return x * (-1)
	}

}

func main() {
	imgPath := "../image/"
	files, err := ioutil.ReadDir(imgPath)
	if err != nil {
		log.Fatalln(err)
	}

	file, _ := os.Open(imgPath + files[6].Name())
	img, _ := jpeg.Decode(file)

	rect := img.Bounds()
	outImage := image.NewRGBA(rect)

	for y := 0; y < rect.Max.Y; y++ {
		for x := 0; x < rect.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			c := colorful.FastLinearRgb(float64(r), float64(g), float64(b))
			h, _, _ := c.Hsv()
			//hsvのhの値が0~30
			//なおかつ青成分は140未満(白を弾くため)
			//その上rgbの差が少ないやつを弾く
			//fmt.Println(r & 0xff)
			col := color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 0xff}
			//if 0 <= h && h <= 35 && (float64(r) > float64(b)*1.8) && (r&0xff > 90) && (b&0xff > 10) {
			if (0 <= h && h <= 40 || 350 < h) &&
				(float64(r) > float64(b)*1.75) &&
				(r&0xff > 90) && (b&0xff > 10) &&
				(r&0xff-g&0xff > 50) && (r&0xff-g&0xff < 100) {
				//if 0 <= h && h <= 40 || 350 < h {
				//if 0 <= h && h <= 30 {
				outImage.Set(x, y, col)
			} else {
				col = color.RGBA{R: 0, G: 0, B: 0, A: 0xff}
				//col := color.RGBA{R: 0, G: 0, B: 0, A: 0}
				outImage.Set(x, y, col)
			}
			//fmt.Println(h)
		}
	}

	outFile, _ := os.Create("result_2.png")
	png.Encode(outFile, outImage)
}
