package parser_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/anibmurthy/htmgenie/pkg/parser"
	"go.uber.org/zap"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "simple test 1",
			input: `# Header one
Hello there

How are you?
What's going on?

## Another Header

This is a paragraph [with an inline link](http://google.com). Neat, eh?

## This is a header [with a link](http://yahoo.com)
`,
			want: `<h1>Header one</h1>

<p>Hello there</p>

<p>How are you?
What's going on?</p>

<h2>Another Header</h2>

<p>This is a paragraph <a href="http://google.com">with an inline link</a>. Neat, eh?</p>

<h2>This is a header <a href="http://yahoo.com">with a link</a></h2>`,
		},
		{
			name: "simple test 2",
			input: `# Sample Document

Hello!

This is sample markdown for the [Mailchimp](https://www.mailchimp.com) homework assignment.
`,
			want: `<h1>Sample Document</h1>

<p>Hello!</p>

<p>This is sample markdown for the <a href="https://www.mailchimp.com">Mailchimp</a> homework assignment.</p>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			p := parser.New(strings.NewReader(test.input), &buf, &zap.Logger{})
			p.Generate()
			got := buf.String()

			if compare(test.want, got) {
				t.Errorf("Expected %v; got %v", test.want, got)
			}
		})
	}
}

func compare(s1 string, s2 string) bool {
	s1 = strings.ReplaceAll(s1, "\n", "")
	s2 = strings.ReplaceAll(s2, "\n", "")

	return s1 != s2
}
