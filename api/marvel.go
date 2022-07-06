package api

import "os"

const MARVEL_BASE_URL string = "https://gateway.marvel.com/v1/public/characters";

type Marvel struct {
  publicKey string
  privateKey string
  query string
  currentPage int
}

func NewMarvel() *Marvel {
  var pubKey = os.Getenv("MARVEL_PUBLIC_KEY");
  var priKey = os.Getenv("MARVEL_PRIVATE_KEY");

  return &Marvel{pubKey, priKey, "", 0}
}

func (m *Marvel) getCharacters(query string) {
  
}