package main

import (
	// "main/api"
	"main/cli"
	ishell "github.com/abiosoft/ishell/v2"
)

func main() {
  // var m = api.NewMarvel()
	
  shell := ishell.New()

  // m.GetCharacters("iron man")
  
  shell.Println("Sample Interactive Shell")

  shell.AddCmd(&ishell.Cmd{
    Name: "search",
    Help: `Search command is used to search Marvel
        characters from api. It finds string matching
        with starting letter of Marvel charachters.
        Eg 1. search iron
        Eg 2. search "iron man"`,
    Func: cli.SearchCommand,
  })

  shell.AddCmd(&ishell.Cmd{
    Name: "next",
    Help: `Next command is used to see the next page.
        Eg. next`,
    Func: cli.NextCommand,
  })

  shell.AddCmd(&ishell.Cmd{
    Name: "prev",
    Help: `Prev command is used to see the next page.
        Eg. prev`,
    Func: cli.NextCommand,
  })
  
  shell.Run()
}
