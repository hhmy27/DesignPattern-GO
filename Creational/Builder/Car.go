package Builder

type CarUnit struct {
	Seats    int
	Engine   string
	Computer bool
	GPS      bool
}

type Car struct {
	CarUnit
}
