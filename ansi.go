package runetui

import (
	"regexp"
	"strings"
	"testing"
)

// ansiPattern matches ANSI escape sequences with letter terminators.
// Primarily covers SGR codes (m terminator) for colors and styles.
var ansiPattern = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)

// StripANSI removes all ANSI escape sequences from a string.
// Returns the visible text content only.
func StripANSI(s string) string {
	return ansiPattern.ReplaceAllString(s, "")
}

// VisualWidth calculates the visible width of a string,
// excluding ANSI escape codes.
func VisualWidth(s string) int {
	return len(StripANSI(s))
}

// VisualHeight returns the number of lines in the output.
// Empty string returns 0; otherwise counts lines.
func VisualHeight(s string) int {
	if s == "" {
		return 0
	}
	return strings.Count(s, "\n") + 1
}

// AssertHasANSICodes verifies that the output contains ANSI escape sequences.
// Useful for ensuring styled output is being rendered.
func AssertHasANSICodes(t testing.TB, output string) {
	t.Helper()
	if !strings.Contains(output, "\x1b[") {
		t.Errorf("expected output to contain ANSI escape codes, got: %q", output)
	}
}

// AssertContainsText verifies that the visible text content contains
// the expected substring, ignoring ANSI codes.
func AssertContainsText(t testing.TB, output, text string) {
	t.Helper()
	stripped := StripANSI(output)
	if !strings.Contains(stripped, text) {
		t.Errorf("expected output to contain text %q, got: %q", text, stripped)
	}
}

// AssertWidth verifies that the visible width of the output matches expected.
// ANSI codes are excluded from the width calculation.
func AssertWidth(t testing.TB, output string, expected int) {
	t.Helper()
	width := VisualWidth(output)
	if width != expected {
		t.Errorf("expected width %d, got %d", expected, width)
	}
}

// AssertHeight verifies that the output has the expected number of lines.
func AssertHeight(t testing.TB, output string, expected int) {
	t.Helper()
	height := VisualHeight(output)
	if height != expected {
		t.Errorf("expected height %d, got %d", expected, height)
	}
}

// AssertNotEmpty verifies that the output has visible content,
// not just whitespace or ANSI codes.
func AssertNotEmpty(t testing.TB, output string) {
	t.Helper()
	stripped := strings.TrimSpace(StripANSI(output))
	if stripped == "" {
		t.Errorf("expected non-empty output, got: %q", output)
	}
}
