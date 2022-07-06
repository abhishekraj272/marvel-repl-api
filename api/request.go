package api

import (
	"io/ioutil"
	"net/http"
)

// Get is used to call api and get byte array from body
func Get(url string) ([]byte, error) {
  res, err := http.Get(url)

  if err != nil {
    return nil, err;
  }

  return ioutil.ReadAll(res.Body)
}

// UrlBuilder is used to generate URL, it accepts baseURL(string) and map of queries
func UrlBuilder(baseURL string, queries map[string]string) (string, error) {
  req, err := http.NewRequest("GET", baseURL, nil)

  if err != nil {
    return "", err;
  }
  
  q := req.URL.Query();
  for k, v := range queries {
    q.Add(k, v);
  }

  req.URL.RawQuery = q.Encode()

  return req.URL.String(), nil;
}