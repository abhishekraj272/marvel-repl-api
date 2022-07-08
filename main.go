package main

import (
	"fmt"
	"main/cli"
	"os"

	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil  {
    fmt.Println("Error loading .env file")
    return;
  }
  
  shell := cli.Init()
  shell.Run()
}