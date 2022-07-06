package api

type CharactersResponse struct {
	Code   int    `json:"code"`
	Data   Data   `json:"data"`
}

type Data struct {
	Total   int      `json:"total"`
	Results []Result `json:"results"`
}

type Result struct {
	Name        string     `json:"name"`
  Description string     `json:"description"`
  Thumbnail   Thumbnail  `json:"thumbnail"`
}

type Thumbnail struct {
  Path string `json:"path"`
}