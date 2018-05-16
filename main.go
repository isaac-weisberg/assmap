package main

import (
	"fmt"
)

func main() {
	error := mapperMain()
	if error != nil {
		fmt.Println("Error:", error)
	}
	fmt.Println("Doen.")
}
