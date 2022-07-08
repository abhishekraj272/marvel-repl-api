package cli

import (
	"fmt"
	"main/api"

	ishell "github.com/abiosoft/ishell/v2"
)

var marvel *api.Marvel = nil;

// Init is used to initialise the CLI
func Init() *ishell.Shell {
  shell := ishell.New()
  
  marvel = api.NewMarvel()

  shell.Println(`Marvel Characters CLI
Type 'help' to see all commands`)

  shell.AddCmd(&ishell.Cmd{
    Name: "next",
    Help: `Next command is used to see the next page.
        Eg. next`,
    Func: nextCommand,
  })

  shell.AddCmd(&ishell.Cmd{
    Name: "prev",
    Help: `Prev command is used to see the next page.
        Eg. prev`,
    Func: prevCommand,
  })

  shell.AddCmd(&ishell.Cmd{
    Name: "search",
    Help: `Search command is used to search Marvel
        characters from api. It finds string matching
        with starting letter of Marvel charachters.
        Eg 1. search iron
        Eg 2. search "iron man"`,
    Func: searchCommand,
  })
  
  return shell;
}

func searchCommand(c *ishell.Context) {
  marvel.GetCharacters(c.Args[0])
}

func nextCommand(c *ishell.Context) {
  if marvel.CanPaginate(api.Next) {
    marvel.SetPage(marvel.Page + 1)
    marvel.GetCharacters(marvel.Query)
    return;
  }
  fmt.Println("***** No more results *****")
}

func prevCommand(c *ishell.Context) {
  if marvel.CanPaginate(api.Prev) {
    fmt.Println(marvel.Page)
    marvel.SetPage(marvel.Page - 1)
    marvel.GetCharacters(marvel.Query)
    return;
  }
  fmt.Println("***** No more results, please search again *****")
}