package SimpleFactory

type ICreator interface {
	createProduct() IProduct
}

type SeatCreator struct {
}

func (s *SeatCreator) createProduct() IProduct {
	return &Seat{}
}

type TableCreator struct {
}

func (t *TableCreator) createProduct() IProduct {
	return &Table{}
}
