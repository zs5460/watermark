package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/zs5460/watermark/syso"
)

const (
	maxWidth = 600
	quality  = 90
	version  = "0.2.1"
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
