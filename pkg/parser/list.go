package parser

type List struct {
	data string
}

func (l *List) Add(s string) bool {
	// Add logic here to determine the end of chunk

	return true
}

func (l *List) Format() string {
	// Add formatting specific logic here

	return l.data
}

func (l *List) Reset() {
	l.data = ""
}

func (l *List) Active() bool {
	return l.data != ""
}
