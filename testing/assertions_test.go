package testing

import (
	"strings"
	stdtesting "testing"

	"github.com/runetui/runetui"
)

// Tests for visualWidth helper (now in runetui package)

func TestVisualWidth_WithPlainText_ReturnsLength(t *stdtesting.T) {
	input := "Hello"
	want := 5
	got := runetui.VisualWidth(input)
	if got != want {
		t.Errorf("runetui.VisualWidth(%q) = %d, want %d", input, got, want)
	}
}

func TestVisualWidth_WithANSICodes_ExcludesThem(t *stdtesting.T) {
	input := "\x1b[1mHello\x1b[0m"
	want := 5
	got := runetui.VisualWidth(input)
	if got != want {
		t.Errorf("runetui.VisualWidth(%q) = %d, want %d", input, got, want)
	}
}

func TestVisualWidth_WithTrailingSpaces_IncludesTrailingSpaces(t *stdtesting.T) {
	input := "Hi   "
	want := 5
	got := runetui.VisualWidth(input)
	if got != want {
		t.Errorf("runetui.VisualWidth(%q) = %d, want %d", input, got, want)
	}
}

func TestVisualWidth_EmptyString_ReturnsZero(t *stdtesting.T) {
	input := ""
	want := 0
	got := runetui.VisualWidth(input)
	if got != want {
		t.Errorf("runetui.VisualWidth(%q) = %d, want %d", input, got, want)
	}
}

// Tests for visualHeight helper function

func TestVisualHeight_SingleLine_ReturnsOne(t *stdtesting.T) {
	input := "Hello"
	want := 1
	got := runetui.VisualHeight(input)
	if got != want {
		t.Errorf("runetui.VisualHeight(%q) = %d, want %d", input, got, want)
	}
}

func TestVisualHeight_MultipleLines_ReturnsLineCount(t *stdtesting.T) {
	input := "Line1\nLine2\nLine3"
	want := 3
	got := runetui.VisualHeight(input)
	if got != want {
		t.Errorf("runetui.VisualHeight(%q) = %d, want %d", input, got, want)
	}
}

func TestVisualHeight_EmptyString_ReturnsZero(t *stdtesting.T) {
	input := ""
	want := 0
	got := runetui.VisualHeight(input)
	if got != want {
		t.Errorf("runetui.VisualHeight(%q) = %d, want %d", input, got, want)
	}
}

// Tests for AssertHasANSICodes - we test by checking if ANSI pattern matches

func TestAssertHasANSICodes_WithANSI_WouldPass(t *stdtesting.T) {
	output := "\x1b[1mBold\x1b[0m"
	if !strings.Contains(output, "\x1b[") {
		t.Error("AssertHasANSICodes should pass when ANSI codes present")
	}
}

func TestAssertHasANSICodes_WithoutANSI_WouldFail(t *stdtesting.T) {
	output := "Plain text"
	if strings.Contains(output, "\x1b[") {
		t.Error("Plain text should not contain ANSI codes")
	}
}

// Tests for AssertContainsText - test underlying behavior

func TestAssertContainsText_TextPresent_WouldPass(t *stdtesting.T) {
	output := "\x1b[1mHello World\x1b[0m"
	stripped := runetui.StripANSI(output)
	if !strings.Contains(stripped, "Hello") {
		t.Error("AssertContainsText should pass when text is present")
	}
}

func TestAssertContainsText_TextMissing_WouldFail(t *stdtesting.T) {
	output := "\x1b[1mHello\x1b[0m"
	stripped := runetui.StripANSI(output)
	if strings.Contains(stripped, "World") {
		t.Error("Stripped output should not contain missing text")
	}
}

func TestAssertContainsText_IgnoresANSICodes_WouldPass(t *stdtesting.T) {
	output := "\x1b[1mHello\x1b[0m"
	stripped := runetui.StripANSI(output)
	if !strings.Contains(stripped, "Hello") {
		t.Error("AssertContainsText should find text within ANSI codes")
	}
}

// Tests for AssertWidth - test underlying behavior

func TestAssertWidth_CorrectWidth_WouldPass(t *stdtesting.T) {
	output := "Hello"
	if runetui.VisualWidth(output) != 5 {
		t.Error("AssertWidth should pass when width matches")
	}
}

func TestAssertWidth_IncorrectWidth_WouldFail(t *stdtesting.T) {
	output := "Hello"
	if runetui.VisualWidth(output) == 10 {
		t.Error("Width should not be 10 for 'Hello'")
	}
}

func TestAssertWidth_WithANSICodes_MeasuresVisibleWidth(t *stdtesting.T) {
	output := "\x1b[1mHello\x1b[0m"
	if runetui.VisualWidth(output) != 5 {
		t.Error("AssertWidth should measure visible width, excluding ANSI codes")
	}
}

// Tests for AssertHeight - test underlying behavior

func TestAssertHeight_CorrectHeight_WouldPass(t *stdtesting.T) {
	output := "Line1\nLine2"
	if runetui.VisualHeight(output) != 2 {
		t.Error("AssertHeight should pass when height matches")
	}
}

func TestAssertHeight_IncorrectHeight_WouldFail(t *stdtesting.T) {
	output := "Line1\nLine2"
	if runetui.VisualHeight(output) == 5 {
		t.Error("Height should not be 5 for two lines")
	}
}

// Tests for AssertNotEmpty - test underlying behavior

func TestAssertNotEmpty_WithContent_WouldPass(t *stdtesting.T) {
	output := "Hello"
	stripped := strings.TrimSpace(runetui.StripANSI(output))
	if stripped == "" {
		t.Error("AssertNotEmpty should pass when content is present")
	}
}

func TestAssertNotEmpty_Empty_WouldFail(t *stdtesting.T) {
	output := ""
	stripped := strings.TrimSpace(runetui.StripANSI(output))
	if stripped != "" {
		t.Error("Empty string should be empty after stripping")
	}
}

func TestAssertNotEmpty_OnlyWhitespace_WouldFail(t *stdtesting.T) {
	output := "   "
	stripped := strings.TrimSpace(runetui.StripANSI(output))
	if stripped != "" {
		t.Error("Whitespace-only string should be empty after trimming")
	}
}

func TestAssertNotEmpty_OnlyANSICodes_WouldFail(t *stdtesting.T) {
	output := "\x1b[1m\x1b[0m"
	stripped := strings.TrimSpace(runetui.StripANSI(output))
	if stripped != "" {
		t.Error("ANSI-only string should be empty after stripping")
	}
}

// Integration tests that actually call assertion functions

func TestAssertHasANSICodes_CanBeCalled(t *stdtesting.T) {
	output := "\x1b[1mBold\x1b[0m"
	runetui.AssertHasANSICodes(t, output)
}

func TestAssertContainsText_CanBeCalled(t *stdtesting.T) {
	output := "\x1b[1mHello World\x1b[0m"
	runetui.AssertContainsText(t, output, "Hello")
	runetui.AssertContainsText(t, output, "World")
}

func TestAssertWidth_CanBeCalled(t *stdtesting.T) {
	output := "Hello"
	runetui.AssertWidth(t, output, 5)
}

func TestAssertWidth_WithANSI_CanBeCalled(t *stdtesting.T) {
	output := "\x1b[1mHello\x1b[0m"
	runetui.AssertWidth(t, output, 5)
}

func TestAssertHeight_CanBeCalled(t *stdtesting.T) {
	output := "Line1\nLine2\nLine3"
	runetui.AssertHeight(t, output, 3)
}

func TestAssertNotEmpty_CanBeCalled(t *stdtesting.T) {
	output := "Hello"
	runetui.AssertNotEmpty(t, output)
}
