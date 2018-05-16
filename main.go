package main

import (
	"fmt"
)

func main() {
	error := mapperMain()
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("Doen.")
}
