package cli

import (
	"fmt"

	ishell "github.com/abiosoft/ishell/v2"
)

func SearchCommand(c *ishell.Context) {
  c.Println(c.Args)
}

func NextCommand(c *ishell.Context) {
  c.Println(c.Args)
  fmt.Println("test")
}