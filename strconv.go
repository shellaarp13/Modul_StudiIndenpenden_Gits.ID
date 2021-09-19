package main

import (
	"fmt"
	_ "fmt"
	"strconv"
	"strings"
)

func main() {
	var input, result, temp string
	var err error
	fmt.Print("Type a word ")
	fmt.Scanln(&input)
	var item = strings.Split(input, "")

	for i := 0; i < len(item); i++ {
		_, err = strconv.Atoi(item[i])
		if err != nil {
			result = result + string(item[i])
		} else {
      fmt.Println("Harus ada string")
    }
	}
	var item1 = strings.Split(result, "")
	for i := len(result) - 1; i >= 0; i-- {
		temp = temp + string(item1[i])

	}
	fmt.Println(temp)
}
