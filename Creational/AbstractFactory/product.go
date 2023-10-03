package AbstractFactory

type Table interface {
	UseTable() string
}

type Seat interface {
	UseSeat() string
}

type ModernTable struct {
}

func (m *ModernTable) UseTable() string {
	return "Use ModernTable"
}

type ArtTable struct {
}

func (a *ArtTable) UseTable() string {
	return "Use ArtTable"
}

type ModernSeat struct {
}

func (m *ModernSeat) UseSeat() string {
	return "Use ModernSeat"
}

type ArtSeat struct {
}

func (a *ArtSeat) UseSeat() string {
	return "Use ArtSeat"
}
