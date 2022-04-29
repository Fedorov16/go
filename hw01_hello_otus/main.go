package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	reversString := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(reversString)
}
