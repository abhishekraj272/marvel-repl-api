package main

import (
	"main/cli"
)

func main() {	
  shell := cli.Init()
  shell.Run()
}
