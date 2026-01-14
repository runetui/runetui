package runetui

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBox_EmptyBox_CanBeCreated(t *testing.T) {
	props := BoxProps{
		Key: "test-box",
	}
	box := Box(props)
	if box == nil {
		t.Fatal("Box() should not return nil")
	}
}

func TestBoxProps_ImplementsProps(t *testing.T) {
	var _ Props = BoxProps{}
}

func TestBox_Key_ReturnsKeyFromProps(t *testing.T) {
	props := BoxProps{
		Key: "my-box",
	}
	box := Box(props)

	got := box.Key()
	want := "my-box"
	if got != want {
		t.Errorf("Key() = %q, want %q", got, want)
	}
}

func TestBox_Children_ReturnsEmptySliceWithNoChildren(t *testing.T) {
	props := BoxProps{Key: "box"}
	box := Box(props)

	children := box.Children()
	if children == nil {
		t.Error("Children() should not return nil, expected empty slice")
	}
	if len(children) != 0 {
		t.Errorf("Children() length = %d, want 0", len(children))
	}
}

func TestBox_Children_ReturnsProvidedChildren(t *testing.T) {
	child1 := &mockComponent{key: "child1"}
	child2 := &mockComponent{key: "child2"}

	props := BoxProps{Key: "box"}
	box := Box(props, child1, child2)

	children := box.Children()
	if len(children) != 2 {
		t.Fatalf("Children() length = %d, want 2", len(children))
	}
	if children[0] != child1 {
		t.Error("First child doesn't match expected")
	}
	if children[1] != child2 {
		t.Error("Second child doesn't match expected")
	}
}

func TestBox_Measure_EmptyBoxReturnsZeroSize(t *testing.T) {
	props := BoxProps{Key: "box"}
	box := Box(props)

	size := box.Measure(100, 100)

	if size.Width != 0 {
		t.Errorf("Measure().Width = %d, want 0", size.Width)
	}
	if size.Height != 0 {
		t.Errorf("Measure().Height = %d, want 0", size.Height)
	}
}

func TestBox_Measure_SumsChildrenVertically(t *testing.T) {
	child1 := &mockComponent{key: "child1", width: 10, height: 5}
	child2 := &mockComponent{key: "child2", width: 15, height: 8}

	props := BoxProps{
		Key:       "box",
		Direction: Column,
	}
	box := Box(props, child1, child2)

	size := box.Measure(100, 100)

	// For vertical stacking: width = max of children, height = sum of children
	wantWidth := 15  // max(10, 15)
	wantHeight := 13 // 5 + 8
	if size.Width != wantWidth {
		t.Errorf("Measure().Width = %d, want %d", size.Width, wantWidth)
	}
	if size.Height != wantHeight {
		t.Errorf("Measure().Height = %d, want %d", size.Height, wantHeight)
	}
}

func TestBox_Measure_SumsChildrenHorizontally(t *testing.T) {
	child1 := &mockComponent{key: "child1", width: 10, height: 5}
	child2 := &mockComponent{key: "child2", width: 15, height: 8}

	props := BoxProps{
		Key:       "box",
		Direction: Row,
	}
	box := Box(props, child1, child2)

	size := box.Measure(100, 100)

	// For horizontal stacking: width = sum of children, height = max of children
	wantWidth := 25 // 10 + 15
	wantHeight := 8 // max(5, 8)
	if size.Width != wantWidth {
		t.Errorf("Measure().Width = %d, want %d", size.Width, wantWidth)
	}
	if size.Height != wantHeight {
		t.Errorf("Measure().Height = %d, want %d", size.Height, wantHeight)
	}
}

func TestBox_Render_EmptyBoxReturnsEmptyString(t *testing.T) {
	props := BoxProps{Key: "box"}
	box := Box(props)

	layout := Layout{X: 0, Y: 0, Width: 10, Height: 10}
	output := box.Render(layout)

	if output != "" {
		t.Errorf("Render() = %q, want empty string", output)
	}
}

func TestBox_Render_SingleChildRendersChildContent(t *testing.T) {
	child := &mockComponent{key: "child", content: "Hello"}

	props := BoxProps{Key: "box"}
	box := Box(props, child)

	layout := Layout{X: 0, Y: 0, Width: 10, Height: 10}
	output := box.Render(layout)

	if output != "Hello" {
		t.Errorf("Render() = %q, want %q", output, "Hello")
	}
}

func TestBox_Render_MultipleChildrenStackVertically(t *testing.T) {
	child1 := &mockComponent{key: "child1", content: "Line 1"}
	child2 := &mockComponent{key: "child2", content: "Line 2"}

	props := BoxProps{
		Key:       "box",
		Direction: Column,
	}
	box := Box(props, child1, child2)

	layout := Layout{X: 0, Y: 0, Width: 10, Height: 10}
	output := box.Render(layout)

	expected := "Line 1\nLine 2"
	if output != expected {
		t.Errorf("Render() = %q, want %q", output, expected)
	}
}

func TestBox_Render_MultipleChildrenStackHorizontally(t *testing.T) {
	child1 := &mockComponent{key: "child1", content: "A"}
	child2 := &mockComponent{key: "child2", content: "B"}

	props := BoxProps{
		Key:       "box",
		Direction: Row,
	}
	box := Box(props, child1, child2)

	layout := Layout{X: 0, Y: 0, Width: 10, Height: 10}
	output := box.Render(layout)

	expected := "AB"
	if output != expected {
		t.Errorf("Render() = %q, want %q", output, expected)
	}
}

func TestBox_Render_WithBorderAppliesLipgloss(t *testing.T) {
	child := &mockComponent{key: "child", content: "Text"}

	props := BoxProps{
		Key:    "box",
		Border: BorderSingle,
	}
	box := Box(props, child)

	layout := Layout{X: 0, Y: 0, Width: 20, Height: 10}
	got := box.Render(layout)

	// Use golden file to verify actual single border rendering
	compareWithGoldenBox(t, "box_border_single", got)
}

func TestBox_Render_WithBackgroundAppliesColor(t *testing.T) {
	child := &mockComponent{key: "child", content: "Text"}

	props := BoxProps{
		Key:        "box",
		Background: "#FF0000",
	}
	box := Box(props, child)

	layout := Layout{X: 0, Y: 0, Width: 20, Height: 10}
	got := box.Render(layout)

	// Use golden file to verify actual red background rendering
	compareWithGoldenBox(t, "box_background_red", got)
}

func TestBox_Render_WithDoubleBorder(t *testing.T) {
	child := &mockComponent{key: "child", content: "X"}

	props := BoxProps{
		Key:    "box",
		Border: BorderDouble,
	}
	box := Box(props, child)

	layout := Layout{X: 0, Y: 0, Width: 20, Height: 10}
	got := box.Render(layout)

	// Use golden file to verify actual double border rendering
	compareWithGoldenBox(t, "box_border_double", got)
}

func TestBox_Render_WithRoundedBorder(t *testing.T) {
	child := &mockComponent{key: "child", content: "Y"}

	props := BoxProps{
		Key:    "box",
		Border: BorderRounded,
	}
	box := Box(props, child)

	layout := Layout{X: 0, Y: 0, Width: 20, Height: 10}
	got := box.Render(layout)

	// Use golden file to verify actual rounded border rendering
	compareWithGoldenBox(t, "box_border_rounded", got)
}

func TestBox_Render_WithBorderColor(t *testing.T) {
	child := &mockComponent{key: "child", content: "Z"}

	props := BoxProps{
		Key:         "box",
		Border:      BorderSingle,
		BorderColor: "#00FF00",
	}
	box := Box(props, child)

	layout := Layout{X: 0, Y: 0, Width: 20, Height: 10}
	got := box.Render(layout)

	// Use golden file to verify actual green border color rendering
	compareWithGoldenBox(t, "box_border_color_green", got)
}

func TestBox_Render_WithBorderAndBackground(t *testing.T) {
	child := &mockComponent{key: "child", content: "Combo"}

	props := BoxProps{
		Key:         "box",
		Border:      BorderSingle,
		BorderColor: "#0000FF",
		Background:  "#FFFF00",
	}
	box := Box(props, child)

	layout := Layout{X: 0, Y: 0, Width: 20, Height: 10}
	got := box.Render(layout)

	// Use golden file to verify border + background combination
	compareWithGoldenBox(t, "box_border_background", got)
}

func TestBox_StyleCombinations_ProducesValidOutput(t *testing.T) {
	tests := []struct {
		name       string
		props      BoxProps
		childCount int
	}{
		{
			name:       "border_only",
			props:      BoxProps{Key: "box", Border: BorderSingle},
			childCount: 1,
		},
		{
			name:       "background_only",
			props:      BoxProps{Key: "box", Background: "#FF0000"},
			childCount: 1,
		},
		{
			name:       "border_and_background",
			props:      BoxProps{Key: "box", Border: BorderSingle, Background: "#00FF00"},
			childCount: 1,
		},
		{
			name:       "border_with_color",
			props:      BoxProps{Key: "box", Border: BorderDouble, BorderColor: "#0000FF"},
			childCount: 1,
		},
		{
			name:       "all_borders",
			props:      BoxProps{Key: "box", Border: BorderRounded, BorderColor: "#FF00FF", Background: "#FFFF00"},
			childCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			child := &mockComponent{key: "child", content: "Test"}
			box := Box(tt.props, child)
			layout := Layout{X: 0, Y: 0, Width: 20, Height: 10}

			got := box.Render(layout)

			// Verify properties using assertion helpers
			AssertContainsText(t, got, "Test")
			AssertNotEmpty(t, got)

			// Verify ANSI codes present when color is applied (background or border color)
			if tt.props.Background != "" || tt.props.BorderColor != "" {
				AssertHasANSICodes(t, got)
			}
		})
	}
}

func TestBox_DirectionVariations_ProducesValidOutput(t *testing.T) {
	tests := []struct {
		name      string
		direction Direction
		expected  string
	}{
		{
			name:      "row_direction",
			direction: Row,
			expected:  "AB",
		},
		{
			name:      "column_direction",
			direction: Column,
			expected:  "A\nB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			child1 := &mockComponent{key: "child1", content: "A"}
			child2 := &mockComponent{key: "child2", content: "B"}

			props := BoxProps{
				Key:       "box",
				Direction: tt.direction,
			}
			box := Box(props, child1, child2)
			layout := Layout{X: 0, Y: 0, Width: 20, Height: 10}

			got := box.Render(layout)

			// Verify expected layout
			if got != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, got)
			}

			// Verify properties
			AssertContainsText(t, got, "A")
			AssertContainsText(t, got, "B")
			AssertNotEmpty(t, got)
		})
	}
}

// Golden file helpers for behavioral testing

func loadGoldenFileBox(t *testing.T, name string) string {
	t.Helper()
	path := filepath.Join("testdata", name+".golden")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read golden file %s: %v", name, err)
	}
	return string(data)
}

func updateGoldenFileBox(t *testing.T, name, content string) {
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

func compareWithGoldenBox(t *testing.T, name, got string) {
	t.Helper()

	if *updateGolden {
		updateGoldenFileBox(t, name, got)
		t.Logf("Updated golden file: %s", name)
		return
	}

	// Check if golden file exists
	path := filepath.Join("testdata", name+".golden")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create golden file on first run
		updateGoldenFileBox(t, name, got)
		t.Logf("Created golden file: %s", name)
		return
	}

	want := loadGoldenFileBox(t, name)
	if got != want {
		t.Errorf("Output doesn't match golden file %s:\ngot:\n%q\n\nwant:\n%q\n\nRun 'go test -update' to update golden files", name, got, want)
	}
}

// mockComponent is a simple Component implementation for testing.
type mockComponent struct {
	key     string
	width   int
	height  int
	content string
}

func (m *mockComponent) Render(layout Layout) string { return m.content }
func (m *mockComponent) Children() []Component       { return nil }
func (m *mockComponent) Key() string                 { return m.key }
func (m *mockComponent) Measure(w, h int) Size {
	return Size{Width: m.width, Height: m.height}
}
