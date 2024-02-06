package main

import (
	"fmt"

	"github.com/golang/example/stringutil" //nolint: depguard
)

func main() {
	fmt.Println(stringutil.Reverse("Hello, OTUS!"))
}
