package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

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

	k, _ = registry.OpenKey(registry.CLASSES_ROOT, `SystemFileAssociations\.png\shell`, registry.ALL_ACCESS)
	registry.DeleteKey(k, `AddPwd\command`)
	registry.DeleteKey(k, "AddPwd")

	fmt.Println("Uninstall Success!")
}
