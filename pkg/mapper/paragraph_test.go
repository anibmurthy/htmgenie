package mapper_test

import (
	"testing"

	"github.com/anibmurthy/htmgenie/pkg/mapper"
)

func TestParagraph(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple paragraph",
			input: "simple paragraph",
			want:  `<p>simple paragraph</p>`,
		},
		{
			name:  "empty line",
			input: "",
			want:  ``,
		},
		{
			name:  "space lines",
			input: "   ",
			want:  `<p>   </p>`,
		},
		{
			name:  "newline char",
			input: "\n",
			want: `
`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := mapper.Paragraph(test.input)

			if test.want != got {
				t.Errorf("Expected %v; got %v", test.want, got)
			}
		})
	}
}
