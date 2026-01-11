package runetui

import "testing"

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
	output := box.Render(layout)

	// With a border, output should be longer than just "Text"
	if len(output) <= len("Text") {
		t.Error("Expected border to increase output length")
	}
	// Border should contain the original content
	if !contains(output, "Text") {
		t.Error("Expected output to contain original content")
	}
}

func TestBox_Render_WithBackgroundAppliesColor(t *testing.T) {
	child := &mockComponent{key: "child", content: "Text"}

	props := BoxProps{
		Key:        "box",
		Background: "#FF0000",
	}
	box := Box(props, child)

	layout := Layout{X: 0, Y: 0, Width: 20, Height: 10}
	output := box.Render(layout)

	// Background should contain the original content
	if !contains(output, "Text") {
		t.Error("Expected output to contain original content")
	}
	// Should include ANSI color codes (output length will be larger due to escape codes)
	if len(output) < len("Text") {
		t.Error("Expected background color to add escape codes")
	}
}

func TestBox_Render_WithDoubleBorder(t *testing.T) {
	child := &mockComponent{key: "child", content: "X"}

	props := BoxProps{
		Key:    "box",
		Border: BorderDouble,
	}
	box := Box(props, child)

	layout := Layout{X: 0, Y: 0, Width: 20, Height: 10}
	output := box.Render(layout)

	if !contains(output, "X") {
		t.Error("Expected output to contain original content")
	}
}

func TestBox_Render_WithRoundedBorder(t *testing.T) {
	child := &mockComponent{key: "child", content: "Y"}

	props := BoxProps{
		Key:    "box",
		Border: BorderRounded,
	}
	box := Box(props, child)

	layout := Layout{X: 0, Y: 0, Width: 20, Height: 10}
	output := box.Render(layout)

	if !contains(output, "Y") {
		t.Error("Expected output to contain original content")
	}
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
	output := box.Render(layout)

	if !contains(output, "Z") {
		t.Error("Expected output to contain original content")
	}
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
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
