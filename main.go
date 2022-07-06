package main

import (
	"fmt"
	"os"
)

func main() {
  var t string = os.Getenv("MARVEL_PUBLIC_KEY")
	fmt.Println(t)
}
