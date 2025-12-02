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
	d.checkPosition()
}

func (d *Dial) TurnLeft(clicks int) {
	d.Position = (d.Position + 100 - clicks) % 100
	d.checkPosition()
}

func (d *Dial) checkPosition() {
	if d.Position == 0 {
		d.ZeroPointingCounts++
	}
}
