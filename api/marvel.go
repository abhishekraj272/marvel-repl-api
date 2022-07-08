package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const MARVEL_BASE_URL string = "https://gateway.marvel.com/v1/public/characters";

type Marvel struct {
  publicKey string
  privateKey string
  Query string
  Page int
  limit int
  totalLeft int
}

func NewMarvel() *Marvel {
  var pubKey = os.Getenv("MARVEL_PUBLIC_KEY");
  var priKey = os.Getenv("MARVEL_PRIVATE_KEY");

  return &Marvel{pubKey, priKey, "", 0, 10, -1}
}

func (m *Marvel) setQuery(query string) {
  m.Query = query;
}

func (m *Marvel) SetPage(page int) {
  m.Page = page;
}

func (m *Marvel) setTotalLeft(total int) {
  m.totalLeft = total;
}

// GetCharacters is used to fetch Marvel characters based on the query.
func (m *Marvel) GetCharacters(searchQuery string) {
  // if the query is different than previous, set page to 0
  if (m.Query != searchQuery) {
    m.SetPage(0);
  }

  queries := m.GetAuthQueryParam();

  queries["nameStartsWith"] = searchQuery;
  queries["limit"] = strconv.Itoa(m.limit);
  queries["offset"] = strconv.Itoa(m.limit * m.Page);

  url, err := UrlBuilder(MARVEL_BASE_URL, queries)
  if err != nil {
    fmt.Errorf("Error: %v\n", err)
    return;
  }

  body, err := Get(url);
  if err != nil {
    fmt.Errorf("Error: %v\n", err)
    return;
  }

  var res CharactersResponse;
  json.Unmarshal(body, &res);

  if res.Code != 200 {
    HandleReqFail(res.Code)
    return;
  }

  m.setQuery(searchQuery);
  m.setTotalLeft(res.Data.Total)

  m.printCharacters(res.Data.Results)
}

// CanPaginate is used to check if pagination is possible.
func (m *Marvel) CanPaginate(direction Direction) bool {
  if (direction == Next && m.totalLeft > 0) {
    return true;
  }
  if (direction == Prev && m.Page > 0) {
    return true;
  }

  return false;
}

// printCharacters is used to print result of search
func (m *Marvel) printCharacters(results []Result) {
  if (len(results) == 0) {
    fmt.Println("No such character found")
    return;
  }
  
  for i, result := range results {
    index := (m.Page * m.limit) + i + 1;
    fmt.Printf("%d.\n Name: %s\n Summary: %s\n Image: %s \n", index, result.Name, result.Description, result.Thumbnail.Path);
  }
}