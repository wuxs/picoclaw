package telegram

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/md2_all_formats.txt
var md2AllFormats string

func Test_markdownToTelegramMarkdownV2(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "heading -> bolding",
			input:    `## HeadingH2 #`,
			expected: "*HeadingH2 \\#*",
		},
		{
			name:     "strikethrough",
			input:    "~strikethroughMD~",
			expected: "~strikethroughMD~",
		},
		{
			name:     "inline URL",
			input:    "[inline URL](http://www.example.com/)",
			expected: "[inline URL](http://www.example.com/)",
		},
		{
			name:     "all telegram formats",
			input:    md2AllFormats,
			expected: md2AllFormats,
		},
		{
			name:     "empty",
			input:    "",
			expected: "",
		},
		{
			name:     "one letter",
			input:    "o",
			expected: "o",
		},
		{
			name:     "",
			input:    "*Last update: ~10 24h*",
			expected: "*Last update: \\~10 24h*",
		},
		{
			name:     "",
			input:    "<Market Capitalization>",
			expected: "\\<Market Capitalization\\>",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := markdownToTelegramMarkdownV2(tc.input)

			require.EqualValues(t, tc.expected, actual)
		})
	}
}

func TestMarkdownToTelegramMarkdownV2_Think(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "closed think block",
			in:   "Intro\n<think>This is a thought\nLine2</think>\nConclusion",
			want: "Intro\n**> 💭 Thinking...**\n> \n> This is a thought\n> Line2\nConclusion",
		},
		{
			name: "unclosed think block",
			in:   "<think>Streaming...",
			want: "**> 💭 Thinking...**\n> \n> Streaming...",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := markdownToTelegramMarkdownV2(tt.in)
			if got != tt.want {
				t.Errorf("\nInput:    %q\nExpected: %q\nGot:      %q", tt.in, tt.want, got)
			}
		})
	}
}
