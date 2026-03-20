package telegram

import (
	"testing"
)

func TestMarkdownToTelegramHTML_Think(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "closed think block",
			input:    "Intro\n<think>This is a thought</think>\nConclusion",
			expected: "Intro\n<blockquote expandable><b>💭 Thinking...</b><br>\nThis is a thought</blockquote>\nConclusion",
		},
		{
			name:     "unclosed think block (streaming)",
			input:    "<think>Thinking...",
			expected: "<blockquote expandable><b>💭 Thinking...</b><br>\nThinking...</blockquote>",
		},
		{
			name:     "multiple think blocks",
			input:    "<think>one</think> mid <think>two</think>",
			expected: "<blockquote expandable><b>💭 Thinking...</b><br>\none</blockquote> mid <blockquote expandable><b>💭 Thinking...</b><br>\ntwo</blockquote>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := markdownToTelegramHTML(tt.input)
			if got != tt.expected {
				t.Errorf("\nInput:    %q\nExpected: %q\nGot:      %q", tt.input, tt.expected, got)
			}
		})
	}
}
