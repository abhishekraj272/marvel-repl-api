package api

type Direction int;

const (
  Prev Direction = iota
  Next 
)

func (d Direction) String() string {
    return []string{"left", "right"}[d];
}