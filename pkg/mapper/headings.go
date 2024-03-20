package mapper

import (
	"fmt"
)

// Currently using implementation to assume the whole line is passed.
// This can be revised to just receive the value to be embedded instead.
// TODO: Evaluate what's the best way forward.
func Heading1(in string) string {
	return fmt.Sprintf("<h1>%s</h1>", in)
}

func Heading2(in string) string {
	return fmt.Sprintf("<h2>%s</h2>", in)
}

func Heading3(in string) string {
	return fmt.Sprintf("<h3>%s</h3>", in)
}

func Heading4(in string) string {
	return fmt.Sprintf("<h4>%s</h4>", in)
}

func Heading5(in string) string {
	return fmt.Sprintf("<h5>%s</h5>", in)
}

func Heading6(in string) string {
	return fmt.Sprintf("<h6>%s</h6>", in)
}
