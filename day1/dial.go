package day1

type Dial struct {
	Position int
}

func NewDial() *Dial {
	return &Dial{Position: 50}
}
