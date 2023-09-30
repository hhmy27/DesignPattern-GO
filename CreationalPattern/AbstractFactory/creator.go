package AbstractFactory

type ICreator interface {
	getTable() Table
	getSeat() Seat
}

type ModernCreator struct {
}

func (t *ModernCreator) getTable() Table {
	return &ModernTable{}
}

func (t *ModernCreator) getSeat() Seat {
	return &ModernSeat{}
}

type ArtCreator struct {
}

func (a *ArtCreator) getTable() Table {
	return &ArtTable{}
}

func (a *ArtCreator) getSeat() Seat {
	return &ArtSeat{}
}
