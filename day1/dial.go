package day1

type Dial struct {
	Position           int
	ZeroPointingCounts int
}

func NewDial() *Dial {
	return &Dial{Position: 50}
}

func (d *Dial) TurnRight(clicks int) {
	d.Position = (d.Position + clicks) % 100
}

func (d *Dial) TurnLeft(clicks int) {
	d.Position = (d.Position + 100 - clicks) % 100
}
