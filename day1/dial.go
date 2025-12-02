package day1

type Dial struct {
	Position           int
	ZeroPointingCounts int
}

func NewDial() *Dial {
	return &Dial{Position: 50}
}

func (d *Dial) TurnRight(clicks int) {
	d.Position = d.Position + clicks
	for d.Position >= 100 {
		d.ZeroPointingCounts++
		d.Position = d.Position - 100
	}
}

func (d *Dial) TurnLeft(clicks int) {
	rawClickCount := d.Position - clicks
	for rawClickCount <= 0 {
		d.ZeroPointingCounts++
		rawClickCount = rawClickCount + 100
	}
	if d.Position == 0 {
		d.ZeroPointingCounts--
	}
	d.Position = rawClickCount % 100
}
