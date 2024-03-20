package parser

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/anibmurthy/htmgenie/pkg/mapper"
	"go.uber.org/zap"
)

type Parser struct {
	filePath io.Reader
	Log      *zap.Logger
	Writer   io.Writer
}

type Chunk struct {
	content string
	Ready   bool
}

func (c *Chunk) Add(s string) bool {
	// Empty spaces are considered an empty line
	if strings.TrimSpace(s) == "" {
		c.Ready = true
		return false
	} else {
		c.Ready = false
	}
	if c.content == "" {
		c.content = s
	} else {
		c.content = fmt.Sprintf("%s\n%s", c.content, s)
	}
	return true
}

func (c *Chunk) Format() string {
	// Take care of url references
	res := mapper.Url(c.content)

	// Since there is no heading, consider this to be <p></p> content
	res = mapper.Paragraph(res)

	c.Reset()

	return res
}

func (c *Chunk) Reset() {
	c.Ready = false
	c.content = ""
}

func (c *Chunk) Active() bool {
	return c.content != ""
}

func New(fp io.Reader, w io.Writer, l *zap.Logger) *Parser {
	return &Parser{
		filePath: fp,
		Log:      l,
		Writer:   w,
	}
}

func (p *Parser) Generate() {
	// Read each line of the mark down and do the needful
	scanner := bufio.NewScanner(p.filePath)
	chunk := Chunk{}

	for scanner.Scan() {
		// Each line is run here
		line := scanner.Text()

		// Handle headings
		if strings.HasPrefix(line, "#") {
			if left, heading, ok := strings.Cut(line, " "); ok {
				// Parse the heading at last.
				if handle, ok := mapper.Mapper[left]; ok {
					// Now Headings might also be made of url references.
					heading = mapper.Url(heading)
					line = handle(heading)
				} else {
					// His is an unknown header type. Treat it as normal text
					// Nothing to do
					p.Log.Sugar().Warnf("Line is starting with '#' but no matching header definition found: %s", line)
				}
			}
			fmt.Fprintln(p.Writer, line)
			continue
		} else {
			if !chunk.Add(line) {
				fmt.Fprintln(p.Writer, chunk.Format())
			} else {
				// Create a new chunk if necessary depending on future scope.
				// Ex: If a new element is present without a newline separation.
			}
		}
		if chunk.Ready {
			fmt.Fprintln(p.Writer, chunk.Format())
		}
	}

	if chunk.Active() {
		fmt.Fprintln(p.Writer, chunk.Format())
	}

	// Scanner gets out of loop only in two cases:
	// 1. If file end was reached
	// 2. If there was an error in reading. Handle error use case
	if err := scanner.Err(); err != nil {
		p.Log.Fatal(fmt.Sprintf("Could not read from the input: %s", "something"),
			zap.String("Error:", err.Error()))
	}
}
