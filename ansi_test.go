package runetui

import "testing"

func TestStripANSI_WithBoldCode_RemovesIt(t *testing.T) {
	input := "\x1b[1mBold\x1b[0m"
	want := "Bold"
	got := StripANSI(input)
	if got != want {
		t.Errorf("StripANSI(%q) = %q, want %q", input, got, want)
	}
}

func TestStripANSI_WithColorCodes_RemovesThem(t *testing.T) {
	input := "\x1b[38;2;255;0;0mRed\x1b[0m"
	want := "Red"
	got := StripANSI(input)
	if got != want {
		t.Errorf("StripANSI(%q) = %q, want %q", input, got, want)
	}
}

func TestStripANSI_WithPlainText_ReturnsUnchanged(t *testing.T) {
	input := "Hello World"
	want := "Hello World"
	got := StripANSI(input)
	if got != want {
		t.Errorf("StripANSI(%q) = %q, want %q", input, got, want)
	}
}

func TestStripANSI_WithEmptyString_ReturnsEmpty(t *testing.T) {
	input := ""
	want := ""
	got := StripANSI(input)
	if got != want {
		t.Errorf("StripANSI(%q) = %q, want %q", input, got, want)
	}
}

func TestStripANSI_WithMultipleCodes_RemovesAll(t *testing.T) {
	input := "\x1b[1m\x1b[3mBold Italic\x1b[0m"
	want := "Bold Italic"
	got := StripANSI(input)
	if got != want {
		t.Errorf("StripANSI(%q) = %q, want %q", input, got, want)
	}
}

func TestStripANSI_WithBackgroundColor_RemovesIt(t *testing.T) {
	input := "\x1b[48;2;0;255;0mGreen BG\x1b[0m"
	want := "Green BG"
	got := StripANSI(input)
	if got != want {
		t.Errorf("StripANSI(%q) = %q, want %q", input, got, want)
	}
}

// Tests for visualWidth helper

func TestVisualWidth_WithPlainText_ReturnsLength(t *testing.T) {
	input := "Hello"
	want := 5
	got := VisualWidth(input)
	if got != want {
		t.Errorf("VisualWidth(%q) = %d, want %d", input, got, want)
	}
}

func TestVisualWidth_WithANSICodes_ExcludesThem(t *testing.T) {
	input := "\x1b[1mHello\x1b[0m"
	want := 5
	got := VisualWidth(input)
	if got != want {
		t.Errorf("VisualWidth(%q) = %d, want %d", input, got, want)
	}
}

func TestVisualWidth_WithTrailingSpaces_IncludesTrailingSpaces(t *testing.T) {
	input := "Hi   "
	want := 5
	got := VisualWidth(input)
	if got != want {
		t.Errorf("VisualWidth(%q) = %d, want %d", input, got, want)
	}
}

func TestVisualWidth_EmptyString_ReturnsZero(t *testing.T) {
	input := ""
	want := 0
	got := VisualWidth(input)
	if got != want {
		t.Errorf("VisualWidth(%q) = %d, want %d", input, got, want)
	}
}

// Tests for visualHeight helper

func TestVisualHeight_SingleLine_ReturnsOne(t *testing.T) {
	input := "Hello"
	want := 1
	got := VisualHeight(input)
	if got != want {
		t.Errorf("VisualHeight(%q) = %d, want %d", input, got, want)
	}
}

func TestVisualHeight_MultipleLines_ReturnsLineCount(t *testing.T) {
	input := "Line1\nLine2\nLine3"
	want := 3
	got := VisualHeight(input)
	if got != want {
		t.Errorf("VisualHeight(%q) = %d, want %d", input, got, want)
	}
}

func TestVisualHeight_EmptyString_ReturnsZero(t *testing.T) {
	input := ""
	want := 0
	got := VisualHeight(input)
	if got != want {
		t.Errorf("VisualHeight(%q) = %d, want %d", input, got, want)
	}
}

// Integration tests that call assertion functions

func TestAssertHasANSICodes_CanBeCalled(t *testing.T) {
	output := "\x1b[1mBold\x1b[0m"
	AssertHasANSICodes(t, output)
}

func TestAssertContainsText_CanBeCalled(t *testing.T) {
	output := "\x1b[1mHello World\x1b[0m"
	AssertContainsText(t, output, "Hello")
	AssertContainsText(t, output, "World")
}

func TestAssertWidth_CanBeCalled(t *testing.T) {
	output := "Hello"
	AssertWidth(t, output, 5)
}

func TestAssertWidth_WithANSI_CanBeCalled(t *testing.T) {
	output := "\x1b[1mHello\x1b[0m"
	AssertWidth(t, output, 5)
}

func TestAssertHeight_CanBeCalled(t *testing.T) {
	output := "Line1\nLine2\nLine3"
	AssertHeight(t, output, 3)
}

func TestAssertNotEmpty_CanBeCalled(t *testing.T) {
	output := "Hello"
	AssertNotEmpty(t, output)
}
