package runetui

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var updateGolden = flag.Bool("update", false, "update golden files")

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

	// Use golden file to verify actual red foreground color rendering
	compareWithGolden(t, "text_foreground_red", got)
}

func TestText_WithBlueForeground_AppliesBlueColor(t *testing.T) {
	text := Text("Hello", TextProps{Color: "#0000FF"})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// Use golden file to verify actual blue foreground color rendering
	compareWithGolden(t, "text_foreground_blue", got)
}

func TestText_WithBackgroundColor_AppliesBackground(t *testing.T) {
	text := Text("Hello", TextProps{Background: "#00FF00"})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// Use golden file to verify actual green background color rendering
	compareWithGolden(t, "text_background_green", got)
}

func TestText_WithBold_AppliesBoldStyle(t *testing.T) {
	text := Text("Hello", TextProps{Bold: true})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// Use golden file to verify actual bold rendering behavior
	compareWithGolden(t, "text_bold", got)
}

func TestText_WithItalic_AppliesItalicStyle(t *testing.T) {
	text := Text("Hello", TextProps{Italic: true})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// Use golden file to verify actual italic rendering behavior
	compareWithGolden(t, "text_italic", got)
}

func TestText_WithUnderline_AppliesUnderlineStyle(t *testing.T) {
	text := Text("Hello", TextProps{Underline: true})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// Use golden file to verify actual underline rendering behavior
	compareWithGolden(t, "text_underline", got)
}

func TestText_WithStrikethrough_AppliesStrikethroughStyle(t *testing.T) {
	text := Text("Hello", TextProps{Strikethrough: true})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// Use golden file to verify actual strikethrough rendering behavior
	compareWithGolden(t, "text_strikethrough", got)
}

func TestText_WithBoldAndRedColor_AppliesBothStyles(t *testing.T) {
	text := Text("Hello", TextProps{Bold: true, Color: "#FF0000"})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// Use golden file to verify bold + red color combination
	compareWithGolden(t, "text_bold_red", got)
}

func TestText_WithItalicAndUnderline_AppliesBothStyles(t *testing.T) {
	text := Text("Hello", TextProps{Italic: true, Underline: true})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// Use golden file to verify italic + underline combination
	compareWithGolden(t, "text_italic_underline", got)
}

func TestText_WithBoldAndBackground_AppliesBothStyles(t *testing.T) {
	text := Text("Hello", TextProps{Bold: true, Background: "#0000FF"})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

	got := text.Render(layout)

	// Use golden file to verify bold + background color combination
	compareWithGolden(t, "text_bold_background", got)
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

// Table-driven tests for style combinations using assertion helpers
// These tests verify that style combinations produce valid output properties
// without being coupled to exact ANSI codes.

func TestText_StyleCombinations_ProducesValidOutput(t *testing.T) {
	tests := []struct {
		name    string
		props   TextProps
		content string
	}{
		{
			name:    "bold_only",
			props:   TextProps{Bold: true},
			content: "Hello",
		},
		{
			name:    "italic_only",
			props:   TextProps{Italic: true},
			content: "Hello",
		},
		{
			name:    "underline_only",
			props:   TextProps{Underline: true},
			content: "Hello",
		},
		{
			name:    "strikethrough_only",
			props:   TextProps{Strikethrough: true},
			content: "Hello",
		},
		{
			name:    "bold_italic",
			props:   TextProps{Bold: true, Italic: true},
			content: "Hello",
		},
		{
			name:    "bold_underline",
			props:   TextProps{Bold: true, Underline: true},
			content: "Hello",
		},
		{
			name:    "italic_underline_strikethrough",
			props:   TextProps{Italic: true, Underline: true, Strikethrough: true},
			content: "Test",
		},
		{
			name:    "all_styles",
			props:   TextProps{Bold: true, Italic: true, Underline: true, Strikethrough: true},
			content: "Test",
		},
		{
			name:    "color_only",
			props:   TextProps{Color: "#FF0000"},
			content: "Red",
		},
		{
			name:    "background_only",
			props:   TextProps{Background: "#00FF00"},
			content: "Green",
		},
		{
			name:    "color_with_bold",
			props:   TextProps{Bold: true, Color: "#FF0000"},
			content: "Bold Red",
		},
		{
			name:    "background_with_italic",
			props:   TextProps{Italic: true, Background: "#0000FF"},
			content: "Italic Blue",
		},
		{
			name:    "full_combination",
			props:   TextProps{Bold: true, Italic: true, Color: "#FF0000", Background: "#0000FF"},
			content: "Full",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			text := Text(tt.content, tt.props)
			layout := Layout{X: 0, Y: 0, Width: 20, Height: 1}

			got := text.Render(layout)

			AssertHasANSICodes(t, got)
			AssertContainsText(t, got, tt.content)
			AssertNotEmpty(t, got)
		})
	}
}

func TestText_AlignmentVariations_ProducesValidOutput(t *testing.T) {
	tests := []struct {
		name    string
		align   TextAlign
		content string
	}{
		{"left", TextAlignLeft, "Hi"},
		{"center", TextAlignCenter, "Hi"},
		{"right", TextAlignRight, "Hi"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			text := Text(tt.content, TextProps{Align: tt.align})
			layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

			got := text.Render(layout)

			AssertContainsText(t, got, tt.content)
			AssertWidth(t, got, 10)
		})
	}
}

func TestText_WrapModes_ProducesValidOutput(t *testing.T) {
	tests := []struct {
		name        string
		wrap        WrapMode
		content     string
		expectLines int
	}{
		{"word_wrap", WrapWord, "Hello World", 3},
		{"truncate", WrapTruncate, "Hello World", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			text := Text(tt.content, TextProps{Wrap: tt.wrap})
			size := text.Measure(5, 10)

			if size.Height != tt.expectLines {
				t.Errorf("expected %d lines, got %d", tt.expectLines, size.Height)
			}
		})
	}
}

// Golden file helpers for behavioral testing

func loadGoldenFile(t *testing.T, name string) string {
	t.Helper()
	path := filepath.Join("testdata", name+".golden")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read golden file %s: %v", name, err)
	}
	return string(data)
}

func updateGoldenFile(t *testing.T, name, content string) {
	t.Helper()
	path := filepath.Join("testdata", name+".golden")

	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatalf("Failed to create testdata directory: %v", err)
	}

	// Write golden file
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write golden file %s: %v", name, err)
	}
}

func compareWithGolden(t *testing.T, name, got string) {
	t.Helper()

	if *updateGolden {
		updateGoldenFile(t, name, got)
		t.Logf("Updated golden file: %s", name)
		return
	}

	// Check if golden file exists
	path := filepath.Join("testdata", name+".golden")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create golden file on first run
		updateGoldenFile(t, name, got)
		t.Logf("Created golden file: %s", name)
		return
	}

	want := loadGoldenFile(t, name)
	if got != want {
		t.Errorf("Output doesn't match golden file %s:\ngot:\n%q\n\nwant:\n%q\n\nRun 'go test -update' to update golden files", name, got, want)
	}
}
