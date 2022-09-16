package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	folders := []string{"."}
	if len(os.Args) > 1 {
		folders = os.Args[1:]
	}
	fmt.Printf("Folders: %s\n", folders)
	for _, folder := range folders {
		tree(folder, "  ")
	}
}

func tree(root string, indent string) {
	fileInfo, err := os.Stat(root)
	if err != nil {
		fmt.Printf("Error :%s", err.Error())
	}
	fmt.Println(fileInfo.Name())
	if !fileInfo.IsDir() {
		return
	}
	files, err := os.ReadDir(root)
	if err != nil {
		fmt.Printf("Error :%s", err.Error())
	}
	// skip the hidden files
	var names []string
	for _, fi := range files {
		if fi.Name()[0] != '.' {
			names = append(names, fi.Name())
		}
	}

	for i, name := range names {
		add := "│    "
		if i == len(names)-1 {
			fmt.Printf("%s└──", indent)
			add = "     "
		} else {
			fmt.Printf("%s├──", indent)
		}
		tree(filepath.Join(root, name), indent+add)
	}
}
