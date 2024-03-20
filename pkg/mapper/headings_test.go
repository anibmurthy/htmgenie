package mapper_test

import (
	"testing"

	"github.com/anibmurthy/htmgenie/pkg/mapper"
)

func TestHeading1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "valid heading",
			input: "header1",
			want:  `<h1>header1</h1>`,
		},
		{
			name:  "valid heading with space",
			input: "header1    ",
			want:  `<h1>header1    </h1>`,
		},
		{
			name:  "valid heading with special character",
			input: "#header",
			want:  `<h1>#header</h1>`,
		},
		{
			name:  "valid heading with alphanumeric",
			input: "#header123",
			want:  `<h1>#header123</h1>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := mapper.Heading1(test.input)

			if test.want != got {
				t.Errorf("Expected %v; got %v", test.want, got)
			}
		})
	}
}

func TestHeading2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "valid heading",
			input: "header1",
			want:  `<h2>header1</h2>`,
		},
		{
			name:  "valid heading with space",
			input: "header1    ",
			want:  `<h2>header1    </h2>`,
		},
		{
			name:  "valid heading with special character",
			input: "#header",
			want:  `<h2>#header</h2>`,
		},
		{
			name:  "valid heading with alphanumeric",
			input: "#header123",
			want:  `<h2>#header123</h2>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := mapper.Heading2(test.input)

			if test.want != got {
				t.Errorf("Expected %v; got %v", test.want, got)
			}
		})
	}
}

func TestHeading3(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "valid heading",
			input: "header3",
			want:  `<h3>header3</h3>`,
		},
		{
			name:  "valid heading with space",
			input: "header1    ",
			want:  `<h3>header1    </h3>`,
		},
		{
			name:  "valid heading with special character",
			input: "#header",
			want:  `<h3>#header</h3>`,
		},
		{
			name:  "valid heading with alphanumeric",
			input: "#header123",
			want:  `<h3>#header123</h3>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := mapper.Heading3(test.input)

			if test.want != got {
				t.Errorf("Expected %v; got %v", test.want, got)
			}
		})
	}
}

func TestHeading4(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "valid heading",
			input: "header3",
			want:  `<h4>header3</h4>`,
		},
		{
			name:  "valid heading with space",
			input: "header1    ",
			want:  `<h4>header1    </h4>`,
		},
		{
			name:  "valid heading with special character",
			input: "#header",
			want:  `<h4>#header</h4>`,
		},
		{
			name:  "valid heading with alphanumeric",
			input: "#header123",
			want:  `<h4>#header123</h4>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := mapper.Heading4(test.input)

			if test.want != got {
				t.Errorf("Expected %v; got %v", test.want, got)
			}
		})
	}
}

func TestHeading5(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "valid heading",
			input: "header3",
			want:  `<h5>header3</h5>`,
		},
		{
			name:  "valid heading with space",
			input: "header1    ",
			want:  `<h5>header1    </h5>`,
		},
		{
			name:  "valid heading with special character",
			input: "#header",
			want:  `<h5>#header</h5>`,
		},
		{
			name:  "valid heading with alphanumeric",
			input: "#header123",
			want:  `<h5>#header123</h5>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := mapper.Heading5(test.input)

			if test.want != got {
				t.Errorf("Expected %v; got %v", test.want, got)
			}
		})
	}
}
func TestHeading6(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "valid heading",
			input: "header3",
			want:  `<h6>header3</h6>`,
		},
		{
			name:  "valid heading with space",
			input: "header1    ",
			want:  `<h6>header1    </h6>`,
		},
		{
			name:  "valid heading with special character",
			input: "#header",
			want:  `<h6>#header</h6>`,
		},
		{
			name:  "valid heading with alphanumeric",
			input: "#header123",
			want:  `<h6>#header123</h6>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := mapper.Heading6(test.input)

			if test.want != got {
				t.Errorf("Expected %v; got %v", test.want, got)
			}
		})
	}
}
