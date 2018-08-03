package utility

import (
	"strings"
)

// StringCentered returns str centered in a string of length width, padding is done using space
func StringCentered(str string, width int) string {
	l := len(str)
	front, rear := width-l, 0
	if front < 0 {
		front = 0
	}
	front = front / 2
	rear = width - l - front
	if rear < 0 {
		rear = 0
	}
	return strings.Repeat(" ", front) + str + strings.Repeat(" ", rear)
}
