package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Program start")
	dir := "."
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if path == dir {
			return nil
		}
		fmt.Println(path)
		return nil
	})
	fmt.Println("Program end")
}
