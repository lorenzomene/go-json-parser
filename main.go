package main

import "fmt"

func main() {
	//read string
	res, err := from_string("{}")
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}
}

func from_string(jsonString string) (bool, error) {
	//lex
	//parse
	fmt.Println(jsonString)
	return true, nil
}
