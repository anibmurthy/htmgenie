package mapper

import "fmt"

func Paragraph(in string) string {
	// As per the requirements empty lines need to be ignored.
	// Choosing to ignore empty content as well
	if in != "" && in != "\n" {
		return fmt.Sprintf("<p>%s</p>", in)
	}
	return in
}
