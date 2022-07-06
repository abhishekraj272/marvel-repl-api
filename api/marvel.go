package api

import (
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
}

func NewMarvel() *Marvel {
  var pubKey = os.Getenv("MARVEL_PUBLIC_KEY");
  var priKey = os.Getenv("MARVEL_PRIVATE_KEY");

  return &Marvel{pubKey, priKey, "", 0, 10}
}

func (m *Marvel) setQuery(query string) {
  m.query = query;
}

func (m *Marvel) setPage(page int) {
  m.page = page;
}

func (m *Marvel) GetCharacters(searchQuery string) {
  queries := m.GetAuthQueries()

  queries["nameStartsWith"] = searchQuery;
  queries["limit"] = strconv.Itoa(m.limit);
  queries["offset"] = strconv.Itoa(m.limit * m.page);

  url, err := UrlBuilder(MARVEL_BASE_URL, queries)
  if err != nil {
    fmt.Errorf("Error: %v\n", err)
  }
  
  fmt.Println(url)
}