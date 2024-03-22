package parser

import (
	"fmt"

	"github.com/anibmurthy/htmgenie/pkg/mapper"
)

type Content interface {
	Add(string) bool
	Format() string
	Reset()
	Active() bool
}

func NewContent(s string) Content {
	return &Paragraph{
		data: s,
	}
}

type Paragraph struct {
	data string
}

func (p *Paragraph) Add(s string) bool {
	if p.data == "" {
		p.data = s
	} else {
		p.data = fmt.Sprintf("%s\n%s", p.data, s)
	}

	return true
}

func (p *Paragraph) Format() string {
	// Take care of url references
	res := mapper.Url(p.data)

	// Since there is no heading, consider this to be <p></p> content
	res = mapper.Paragraph(res)

	return res
}

func (p *Paragraph) Reset() {
	p.data = ""
}

func (p *Paragraph) Active() bool {
	return p.data != ""
}
