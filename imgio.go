package imgio

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/mccutchen/palettor"
)

func init() {
	// damn important or else At(), Bounds() functions will
	// caused memory pointer error!!
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
}

// Paletted2Image converts image.Paletted to image.Image
func Paletted2Image(imagepaletted *image.Paletted) image.Image {
	size := imagepaletted.Bounds().Size()
	outputimage := image.NewRGBA(imagepaletted.Rect)
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			outputimage.Set(x, y, imagepaletted.At(x, y))
		}
	}
	return outputimage
}

// PngSave save png
func PngSave(path string, img image.Image) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
}

// JpgSave save jpg
func JpgSave(path string, img image.Image) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	err = jpeg.Encode(f, img, nil)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
}

// GifAnimSave save animated gif
func GifAnimSave(path string, images []image.Image, delay int, k int, maxIterations int, generatePalette bool) {
	gifPalette := color.Palette(palette.Plan9)
	if generatePalette == true {
		paletTemp, _ := palettor.Extract(k, maxIterations, images[0])
		gifPalette = paletTemp.Colors()
		// if err != nil {
		// 	log.Fatalf("image too small")
		// }
	}
	bound := images[0].Bounds()
	var imagesp []*image.Paletted
	intArray := []int{delay}
	for _, img := range images {
		palettedImage := image.NewPaletted(bound, gifPalette)
		draw.FloydSteinberg.Draw(palettedImage, bound, img, image.Point{})
		// draw.Draw(palettedImage, palettedImage.Rect, img, bounds.Min, draw.Over)
		imagesp = append(imagesp, palettedImage)
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: imagesp,
		Delay: intArray,
	})
}

// ImRead reads image from path
func ImRead(path string) image.Image {
	imgfile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(imgfile)
	if err != nil {
		log.Fatal(err)
	}
	imgfile.Close()
	return img
}

// GifRead reads gif from path
func GifRead(path string) *gif.GIF {
	imgfile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	frames, err := gif.DecodeAll(imgfile)
	if err != nil {
		log.Fatal(err)
	}
	imgfile.Close()
	return frames
}
