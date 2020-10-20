package main

import (
	"flag"
	"fmt"

	"github.com/brandNamer"
)

var num = flag.Int("num", 3, "number of letters for brand name")

func main() {
	flag.Parse()

	fmt.Println("Your new brand name is: ", brandNamer.Name(*num))
}
