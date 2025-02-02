package utils

func BoolToInt(value bool) int {
	if value {
		return 1
	} else {
		return 0
	}
}

type Coords struct {
	X float32
	Y float32
}
