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
  query string
  page int
  limit int
  totalLeft int
}

func NewMarvel() *Marvel {
  var pubKey = os.Getenv("MARVEL_PUBLIC_KEY");
  var priKey = os.Getenv("MARVEL_PRIVATE_KEY");

  return &Marvel{pubKey, priKey, "", 0, 10, -1}
}

func (m *Marvel) setQuery(query string) {
  m.query = query;
}

func (m *Marvel) setPage(page int) {
  m.page = page;
}

func (m *Marvel) setTotalLeft(total int) {
  m.totalLeft = total;
}

// GetCharacters is used to fetch Marvel characters based on the query.
func (m *Marvel) GetCharacters(searchQuery string) {
  // if the query is different than previous, set page to 0
  if (m.query != searchQuery) {
    m.setPage(0);
  }

  queries := m.GetAuthQueryParam();

  queries["nameStartsWith"] = searchQuery;
  queries["limit"] = strconv.Itoa(m.limit);
  queries["offset"] = strconv.Itoa(m.limit * m.page);

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
    handleReqFail(res.Code)
    return;
  }

  m.setPage(m.page + 1);
  m.setQuery(searchQuery);
  m.setTotalLeft(res.Data.Total)

  printCharacters(res.Data.Results)
}

func (m *Marvel) CanPaginate(direction Direction) bool {
  if (direction == Next && m.totalLeft > 0) {
    return true;
  }
  if (direction == Prev && m.page > 0) {
    return true;
  }

  return false;
}

func handleReqFail(code int) {
  
}

func printCharacters(results []Result) {
  for i, result := range results {
    fmt.Printf("%d.\n Name: %s\n Summary: %s\n Image: %s \n", i+1, result.Name, result.Description, result.Thumbnail.Path);
  }
}