package mapper_test

import (
	"testing"

	"github.com/anibmurthy/htmgenie/pkg/mapper"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple url",
			input: "[link](https://www.example.com)",
			want:  `<a href="https://www.example.com">link</a>`,
		},
		{
			name:  "simple url2",
			input: "[link](http://www.example.com)",
			want:  `<a href="http://www.example.com">link</a>`,
		},
		{
			name:  "url - variety1",
			input: "[link](https://www.more.example.com)",
			want:  `<a href="https://www.more.example.com">link</a>`,
		},
		{
			name:  "url - variety2",
			input: "[link](https://www.more123.example.com)",
			want:  `<a href="https://www.more123.example.com">link</a>`,
		},
		{
			name:  "url - variety3",
			input: "[link ](https://www.more.example.com)",
			want:  `<a href="https://www.more.example.com">link </a>`,
		},
		{
			name:  "url - variety4",
			input: "[link text with space](https://www.more.example.com)",
			want:  `<a href="https://www.more.example.com">link text with space</a>`,
		},
		{
			name:  "url - variety5",
			input: "[link#123](https://www.more.example.com)",
			want:  `<a href="https://www.more.example.com">link#123</a>`,
		},
		{
			name:  "url - variety6",
			input: "[link](https://www.more.example.com/some/path/value/123)",
			want:  `<a href="https://www.more.example.com/some/path/value/123">link</a>`,
		},
		{
			name:  "url - variety7",
			input: "[link/text](https://www.example.com)",
			want:  `<a href="https://www.example.com">link/text</a>`,
		},
		{
			name:  "url - variety8",
			input: `[link\text](https://www.example.com)`,
			want:  `<a href="https://www.example.com">link\text</a>`,
		},
		{
			name:  "url - variety9",
			input: `[link.text](https://www.example.com)`,
			want:  `<a href="https://www.example.com">link.text</a>`,
		},
		{
			name:  "url - variety10",
			input: `[link.](https://www.example.com)`,
			want:  `<a href="https://www.example.com">link.</a>`,
		},
		{
			name:  "url - variety11",
			input: `[link....](https://www.example.com)`,
			want:  `<a href="https://www.example.com">link....</a>`,
		},
		{
			name:  "url - variety12",
			input: `[link>](https://www.example.com)`,
			want:  `<a href="https://www.example.com">link></a>`,
		},
		{
			name:  "url - variety13",
			input: `[link<](https://www.example.com)`,
			want:  `<a href="https://www.example.com">link<</a>`,
		},
		{
			name:  "url - variety13",
			input: `[link](https://www.example.com/#L23)`,
			want:  `<a href="https://www.example.com/#L23">link</a>`,
		},
		{
			name:  "url - variety14 - with text",
			input: `junk text [link](https://www.example.com/#L23)`,
			want:  `junk text <a href="https://www.example.com/#L23">link</a>`,
		},
		{
			name:  "invalid url1",
			input: "[link][link123](https://www.example.com)",
			want:  `[link]<a href="https://www.example.com">link123</a>`,
		},
		{
			name:  "invalid url2",
			input: "[[link]](https://www.more.example.com)",
			want:  `[[link]](https://www.more.example.com)`,
		},
		{
			name:  "invalid url3",
			input: "[link](https://www.more.example.com https://www.more.example.com)",
			want:  `[link](https://www.more.example.com https://www.more.example.com)`,
		},
		{
			name:  "invalid url4",
			input: "[link]random text(https://www.more.example.com)",
			want:  `[link]random text(https://www.more.example.com)`,
		},
		{
			name:  "With non link string in the line",
			input: "I am here testing this [link](https://www.example.com)",
			want:  `I am here testing this <a href="https://www.example.com">link</a>`,
		},
		{
			name:  "With multiple links",
			input: "This is first [link](https://www.example.com). And this one is second [link](https://www.example.com)",
			want:  `This is first <a href="https://www.example.com">link</a>. And this one is second <a href="https://www.example.com">link</a>`,
		},
		{
			name:  "With multiple links - one of them invalid",
			input: "This is first [link](https://www.example.com). And this one is second [link](https://www.example.com). Third is [wrong](www.wrong.com)",
			want:  `This is first <a href="https://www.example.com">link</a>. And this one is second <a href="https://www.example.com">link</a>. Third is [wrong](www.wrong.com)`,
		},
		{
			name: "With multiple links - multiple lines",
			input: `This is first [link](https://www.example.com).
And this one is second [link](https://www.example.com).
Third is [wrong](www.wrong.com)`,
			want: `This is first <a href="https://www.example.com">link</a>.
And this one is second <a href="https://www.example.com">link</a>.
Third is [wrong](www.wrong.com)`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := mapper.Url(test.input)

			if test.want != got {
				t.Errorf("Expected %v; got %v", test.want, got)
			}
		})
	}
}
