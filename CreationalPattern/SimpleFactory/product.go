package SimpleFactory

import "fmt"

type IProduct interface {
	use() string
}

type Table struct {
}

func (t *Table) use() string {
	return fmt.Sprintf("Using Table")
}

type Seat struct {
}

func (s *Seat) use() string {
	return fmt.Sprintf("Using Seat")
}
