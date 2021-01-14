package main

import (
	"imgio"
	// "image"
	// "github.com/nfnt/resize"
)

func main() {

	// img := imgio.ImRead("_img/lenna.png")
	// img = resize.Resize(512, 0, img, resize.Bicubic)

	// images := []image.Image{img}

	// imgio.GifAnimSave("./_img/output16-4.gif", images, 100, 16, 4, true)

	gif := imgio.GifRead("_img/_input/flow.gif")

	// Writes first GIF frame
	imgio.PngSave("_img/_output/flow_frame_00.png", imgio.Paletted2Image(gif.Image[0]))
}
