package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

func mark(fn string) {
	ext := path.Ext(fn)
	switch ext {
	case ".jpg":
		markJPG(fn)
	case ".png":
		markPNG(fn)
	}

}

func addMark(img image.Image, mark image.Image) image.Image {

	b := img.Bounds()

	if b.Dx() > maxWidth {
		w := uint(maxWidth)
		h := uint(maxWidth * b.Dy() / b.Dx())
		img = resize.Resize(w, h, img, resize.Lanczos3)
	}

	b = img.Bounds()

	offset := image.Pt(b.Dx()-mark.Bounds().Dx()-20, b.Dy()-mark.Bounds().Dy()-20)
	newimg := image.NewNRGBA(b)

	draw.Draw(newimg, b, img, image.ZP, draw.Src)
	draw.Draw(newimg, mark.Bounds().Add(offset), mark, image.ZP, draw.Over)
	return newimg
}

func getMark() image.Image {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	wmfile, err := os.Open(dir + `\watermark.png`)
	if err != nil {
		panic(err)
	}
	watermark, _ := png.Decode(wmfile)
	defer wmfile.Close()
	return watermark
}

func markJPG(fn string) {

	imgfile, err := os.Open(fn)
	defer imgfile.Close()
	if err != nil {
		panic(err)
	}
	img, _ := jpeg.Decode(imgfile)

	watermark := getMark()
	m := addMark(img, watermark)

	ext := filepath.Ext(fn)
	imgOutput, _ := os.Create(strings.Replace(fn, ext, ".marked"+ext, -1))
	defer imgOutput.Close()
	jpeg.Encode(imgOutput, m, &jpeg.Options{Quality: quality})

}

func markPNG(fn string) {
	imgfile, err := os.Open(fn)
	defer imgfile.Close()
	if err != nil {
		panic(err)
	}
	img, _ := png.Decode(imgfile)

	watermark := getMark()
	m := addMark(img, watermark)

	ext := filepath.Ext(fn)
	imgOutput, _ := os.Create(strings.Replace(fn, ext, ".marked"+ext, -1))
	defer imgOutput.Close()
	png.Encode(imgOutput, m)

}
