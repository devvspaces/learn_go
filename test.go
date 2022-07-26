package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		str := recover()
		if str != nil {
			fmt.Println("Make sure you passed a file name")
		}
	}()

	allowed_flags := []string{
		"-f",
	}

	if len(os.Args) > 1 {
		flag := os.Args[1]
		found := false
		for _, value := range allowed_flags {
			if flag == value {
				found = true
				file_name := os.Args[2]
				fmt.Println(file_name)
				break
			}
		}
		if found == false {
			panic("Invalid Flag used")
		}
	}

}
