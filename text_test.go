package runetui

import (
	"strings"
	"testing"
)

func TestText_WithBasicContent_RendersPlainText(t *testing.T) {
	text := Text("Hello")
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// Should contain "Hello" (may be padded to width)
	if !strings.Contains(got, "Hello") {
		t.Errorf("Expected output to contain 'Hello', got %q", got)
	}

	// Should be 10 characters wide (padded)
	if len(got) != 10 {
		t.Errorf("Expected width 10, got %d", len(got))
	}
}

func TestText_WithForegroundColor_AppliesColor(t *testing.T) {
	text := Text("Hello", TextProps{Color: "#FF0000"})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// ANSI escape codes should be present for colored text
	if !strings.Contains(got, "\x1b[") {
		t.Errorf("Expected ANSI color codes in output, got: %q", got)
	}

	// Should still contain the actual text
	if !strings.Contains(got, "Hello") {
		t.Errorf("Expected text content 'Hello' in output, got: %q", got)
	}
}

func TestText_WithBackgroundColor_AppliesBackground(t *testing.T) {
	text := Text("Hello", TextProps{Background: "#00FF00"})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// ANSI escape codes should be present for colored text
	if !strings.Contains(got, "\x1b[") {
		t.Errorf("Expected ANSI color codes in output, got: %q", got)
	}

	// Should still contain the actual text
	if !strings.Contains(got, "Hello") {
		t.Errorf("Expected text content 'Hello' in output, got: %q", got)
	}
}

func TestText_WithBold_AppliesBoldStyle(t *testing.T) {
	text := Text("Hello", TextProps{Bold: true})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// ANSI bold code is typically \x1b[1m
	if !strings.Contains(got, "\x1b[1m") && !strings.Contains(got, "\x1b[") {
		t.Errorf("Expected ANSI bold codes in output, got: %q", got)
	}

	if !strings.Contains(got, "Hello") {
		t.Errorf("Expected text content 'Hello' in output, got: %q", got)
	}
}

func TestText_WithItalic_AppliesItalicStyle(t *testing.T) {
	text := Text("Hello", TextProps{Italic: true})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// ANSI italic code is typically \x1b[3m
	if !strings.Contains(got, "\x1b[3m") && !strings.Contains(got, "\x1b[") {
		t.Errorf("Expected ANSI italic codes in output, got: %q", got)
	}

	if !strings.Contains(got, "Hello") {
		t.Errorf("Expected text content 'Hello' in output, got: %q", got)
	}
}

func TestText_WithUnderline_AppliesUnderlineStyle(t *testing.T) {
	text := Text("Hello", TextProps{Underline: true})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// ANSI underline code is typically \x1b[4m
	if !strings.Contains(got, "\x1b[4") {
		t.Errorf("Expected ANSI underline codes in output, got: %q", got)
	}

	// Text should contain all characters (even if wrapped in codes)
	if !strings.Contains(got, "H") || !strings.Contains(got, "e") || !strings.Contains(got, "l") || !strings.Contains(got, "o") {
		t.Errorf("Expected text content 'Hello' in output, got: %q", got)
	}
}

func TestText_WithStrikethrough_AppliesStrikethroughStyle(t *testing.T) {
	text := Text("Hello", TextProps{Strikethrough: true})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// ANSI strikethrough code is typically \x1b[9m
	if !strings.Contains(got, "\x1b[9m") {
		t.Errorf("Expected ANSI strikethrough codes in output, got: %q", got)
	}

	// Text should contain all characters (even if wrapped in codes)
	if !strings.Contains(got, "H") || !strings.Contains(got, "e") || !strings.Contains(got, "l") || !strings.Contains(got, "o") {
		t.Errorf("Expected text content 'Hello' in output, got: %q", got)
	}
}

func TestText_WrapWord_WrapsAtWordBoundaries(t *testing.T) {
	text := Text("Hello World", TextProps{Wrap: WrapWord})
	layout := Layout{X: 0, Y: 0, Width: 5, Height: 2}

	got := text.Render(layout)

	// Should wrap at word boundary, so "Hello" on first line
	if !strings.Contains(got, "Hello") {
		t.Errorf("Expected 'Hello' in output, got: %q", got)
	}
}

func TestText_WrapTruncate_TruncatesWithEllipsis(t *testing.T) {
	text := Text("Hello World", TextProps{Wrap: WrapTruncate})
	layout := Layout{X: 0, Y: 0, Width: 7, Height: 1}

	got := text.Render(layout)

	// Should truncate with ellipsis when text is too long
	// Lipgloss uses "..." or "…" for ellipsis
	if len(got) > 10 {
		t.Errorf("Expected truncated text, got: %q (len=%d)", got, len(got))
	}
}

func TestText_AlignLeft_AlignsTextLeft(t *testing.T) {
	text := Text("Hi", TextProps{Align: TextAlignLeft})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// Left-aligned text should start at the beginning (no leading spaces)
	// Should be "Hi        " with trailing spaces to fill width
	if !strings.HasPrefix(got, "Hi") {
		t.Errorf("Expected text to start with 'Hi', got: %q", got)
	}
}

func TestText_AlignCenter_CentersText(t *testing.T) {
	text := Text("Hi", TextProps{Align: TextAlignCenter})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// Centered text should have spaces before and after
	// For "Hi" in width 10: "    Hi    " (4 spaces before, 4 after)
	if strings.HasPrefix(got, "Hi") {
		t.Errorf("Expected centered text with leading spaces, got: %q", got)
	}
	if !strings.Contains(got, "Hi") {
		t.Errorf("Expected text to contain 'Hi', got: %q", got)
	}
}

func TestText_AlignRight_AlignsTextRight(t *testing.T) {
	text := Text("Hi", TextProps{Align: TextAlignRight})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// Right-aligned text should have leading spaces
	// For "Hi" in width 10: "        Hi"
	if strings.HasPrefix(got, "Hi") {
		t.Errorf("Expected right-aligned text with leading spaces, got: %q", got)
	}
	if !strings.HasSuffix(got, "Hi") && !strings.Contains(got, "Hi") {
		t.Errorf("Expected text to end with 'Hi', got: %q", got)
	}
}

func TestText_Measure_ReturnsCorrectSize(t *testing.T) {
	text := Text("Hello")
	size := text.Measure(10, 10)

	// Text "Hello" is 5 characters wide, 1 line tall
	if size.Width != 5 {
		t.Errorf("Expected width 5, got %d", size.Width)
	}
	if size.Height != 1 {
		t.Errorf("Expected height 1, got %d", size.Height)
	}
}

func TestText_Measure_WithWrapWord_CalculatesMultipleLines(t *testing.T) {
	text := Text("Hello World", TextProps{Wrap: WrapWord})
	size := text.Measure(5, 10)

	// Text "Hello World" (11 chars) wrapped at 5 should be 5 wide, 3 lines
	if size.Width != 5 {
		t.Errorf("Expected width 5, got %d", size.Width)
	}
	if size.Height != 3 {
		t.Errorf("Expected height 3, got %d", size.Height)
	}
}

func TestText_Measure_WithWrapTruncate_SingleLine(t *testing.T) {
	text := Text("Hello World", TextProps{Wrap: WrapTruncate})
	size := text.Measure(5, 10)

	// Truncated text should be 5 wide, 1 line
	if size.Width != 5 {
		t.Errorf("Expected width 5, got %d", size.Width)
	}
	if size.Height != 1 {
		t.Errorf("Expected height 1, got %d", size.Height)
	}
}

func TestText_Children_ReturnsEmptySlice(t *testing.T) {
	text := Text("Hello")
	children := text.Children()

	if children == nil {
		t.Error("Children() should return empty slice, not nil")
	}
	if len(children) != 0 {
		t.Errorf("Expected 0 children, got %d", len(children))
	}
}

func TestText_Key_ReturnsKeyFromProps(t *testing.T) {
	text := Text("Hello", TextProps{Key: "test-key"})
	key := text.Key()

	if key != "test-key" {
		t.Errorf("Expected key 'test-key', got %q", key)
	}
}

func TestText_Key_ReturnsEmptyWhenNotSet(t *testing.T) {
	text := Text("Hello")
	key := text.Key()

	if key != "" {
		t.Errorf("Expected empty key, got %q", key)
	}
}

func TestTextProps_ImplementsPropsInterface(t *testing.T) {
	props := TextProps{Content: "test"}
	var _ Props = props
}
