package day1

type Dial struct {
	Position int
}

func NewDial() *Dial {
	return &Dial{Position: 50}
}

func (d *Dial) TurnRight(clicks int) {
	d.Position += clicks
}

func (d *Dial) TurnLeft(clicks int) {
	d.Position -= clicks
}
