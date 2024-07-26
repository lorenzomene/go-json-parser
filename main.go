package main

import (
	"fmt"
)

func main() {
	res, err := from_string("{}")
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}
}

func from_string(jsonString string) (bool, error) {
	fmt.Println(jsonString)
	return true, nil
}
