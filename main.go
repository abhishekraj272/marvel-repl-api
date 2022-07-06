package main

import (
	"main/api"
)

func main() {
  var m = api.NewMarvel()
	m.GetCharacters("iron")
}
