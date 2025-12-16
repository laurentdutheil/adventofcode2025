package day10

type Button struct {
	LightsPosition []int
}

func (b Button) Toggle(input string) string {
	result := input
	for _, position := range b.LightsPosition {
		c := b.toggleChar(input, position)
		result = result[:position] + c + result[position+1:]
	}
	return result
}

func (b Button) toggleChar(input string, position int) string {
	c := "."
	if input[position] == '.' {
		c = "#"
	}
	return c
}

func NewButton(lightsPosition []int) Button {
	return Button{LightsPosition: lightsPosition}
}
