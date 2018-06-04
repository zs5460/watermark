package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	_ "watermark/syso"

	"github.com/golang/sys/windows/registry"
	"github.com/nfnt/resize"
)

const (
	maxWidth = 576
	quality  = 90
	version  = "0.2.0"
)

func main() {

	var arg1 string
	if len(os.Args) > 1 {
		arg1 = strings.ToLower(os.Args[1])
	} else {
		install()
	}

	if strings.HasSuffix(arg1, ".jpg") || strings.HasSuffix(arg1, ".png") {
		mark(arg1)
	} else if arg1 == "-i" {
		install()
	} else if arg1 == "-u" {
		uninstall()
	} else if arg1 == "-h" {
		help()
	}

}

func help() {
	msg := "Songz Watermark Tool v" + version +
		"\nAvailable at https://github.com/zs5460/watermark\n\n" +
		"Copyright © 2018 zs5460 <zs5460@gmail.com>\n" +
		"Distributed under the Simplified BSD License\n" +
		"See website for details\n\n" +
		"Usage:\n" +
		"  %s [options] [inputfile]\n\n" +
		"Options:\n" +
		"-i\n" +
		"  install to Registry (run as administrator)\n" +
		"-u\n" +
		"  uninstall and remove from Registry (run as administrator)\n\n"

	fmt.Printf(msg, os.Args[0])
}

func mark(fn string) {
	ext := path.Ext(fn)
	switch ext {
	case ".jpg":
		markJPG(fn)
	case ".png":
		markPNG(fn)
	}

}

func markJPG(fn string) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	imgfile, err := os.Open(fn)

	img, _ := jpeg.Decode(imgfile)
	defer imgfile.Close()
	if err != nil {
		panic(err)
	}

	wmfile, err := os.Open(dir + `\watermark.png`)
	if err != nil {
		panic(err)
	}
	watermark, _ := png.Decode(wmfile)
	defer wmfile.Close()

	b := img.Bounds()
	//如果超尺寸则resize
	if b.Dx() > maxWidth {
		w := uint(maxWidth)
		h := uint(maxWidth * b.Dy() / b.Dx())
		img = resize.Resize(w, h, img, resize.Lanczos3)
	}

	b = img.Bounds()
	//把水印写到右下角，边距20像素
	offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-20, img.Bounds().Dy()-watermark.Bounds().Dy()-20)
	m := image.NewNRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	imgOutput, _ := os.Create(strings.Replace(fn, ".jpg", ".wm.jpg", -1))
	defer imgOutput.Close()
	jpeg.Encode(imgOutput, m, &jpeg.Options{Quality: quality})

}

func markPNG(fn string) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	imgfile, err := os.Open(fn)

	img, _ := png.Decode(imgfile)
	defer imgfile.Close()
	if err != nil {
		panic(err)
	}

	wmfile, err := os.Open(dir + `\watermark.png`)
	if err != nil {
		panic(err)
	}
	watermark, _ := png.Decode(wmfile)
	defer wmfile.Close()

	b := img.Bounds()
	//如果超尺寸则resize
	if b.Dx() > maxWidth {
		w := uint(maxWidth)
		h := uint(maxWidth * b.Dy() / b.Dx())
		img = resize.Resize(w, h, img, resize.Lanczos3)
	}

	b = img.Bounds()
	//把水印写到右下角，边距20像素
	offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-20, img.Bounds().Dy()-watermark.Bounds().Dy()-20)
	m := image.NewNRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	imgOutput, _ := os.Create(strings.Replace(fn, ".png", ".wm.png", -1))
	defer imgOutput.Close()
	png.Encode(imgOutput, m)

}

func install() {
	full, _ := filepath.Abs(os.Args[0])
	k, err := registry.OpenKey(registry.CLASSES_ROOT, `SystemFileAssociations\.jpg\shell`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	sk1, _, _ := registry.CreateKey(k, "AddPwd", registry.ALL_ACCESS)
	defer sk1.Close()
	sk2, _, _ := registry.CreateKey(k, `AddPwd\command`, registry.ALL_ACCESS)
	defer sk2.Close()

	sk1.SetStringValue("", "裁剪并打水印")
	sk2.SetStringValue("", `"`+full+`" "%1"`)
	k.Close()

	k, err = registry.OpenKey(registry.CLASSES_ROOT, `SystemFileAssociations\.png\shell`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	sk1, _, _ = registry.CreateKey(k, "AddPwd", registry.ALL_ACCESS)
	defer sk1.Close()
	sk2, _, _ = registry.CreateKey(k, `AddPwd\command`, registry.ALL_ACCESS)
	defer sk2.Close()

	sk1.SetStringValue("", "裁剪并打水印")
	sk2.SetStringValue("", `"`+full+`" "%1"`)
	k.Close()

	fmt.Println("Install Success!")

}

func uninstall() {
	k, _ := registry.OpenKey(registry.CLASSES_ROOT, `SystemFileAssociations\.jpg\shell`, registry.ALL_ACCESS)
	registry.DeleteKey(k, `AddPwd\command`)
	registry.DeleteKey(k, "AddPwd")
	fmt.Println("Uninstall Success!")
}
