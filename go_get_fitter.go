package main

import (
	"fmt"
	"go-get-fitter.com/app/data"
)

func main() {
	result := data.SplitByComma("a, b")
	fmt.Println(result)
}