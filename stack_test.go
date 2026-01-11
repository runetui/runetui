package runetui

import "testing"

// Step 1: Test that VStack creates a component
func TestVStack_WithNoChildren_CanBeCreated(t *testing.T) {
	stack := VStack()
	if stack == nil {
		t.Fatal("VStack() should not return nil")
	}
}

// Step 2: Test that VStack renders children vertically
func TestVStack_WithChildren_RendersVertically(t *testing.T) {
	child1 := &mockComponent{key: "child1", content: "Line1"}
	child2 := &mockComponent{key: "child2", content: "Line2"}

	stack := VStack(child1, child2)

	layout := Layout{X: 0, Y: 0, Width: 10, Height: 10}
	output := stack.Render(layout)

	expected := "Line1\nLine2"
	if output != expected {
		t.Errorf("Render() = %q, want %q", output, expected)
	}
}

// Step 3: Test that HStack can be created
func TestHStack_WithNoChildren_CanBeCreated(t *testing.T) {
	stack := HStack()
	if stack == nil {
		t.Fatal("HStack() should not return nil")
	}
}

// Step 4: Test that HStack renders children horizontally
func TestHStack_WithChildren_RendersHorizontally(t *testing.T) {
	child1 := &mockComponent{key: "child1", content: "A"}
	child2 := &mockComponent{key: "child2", content: "B"}

	stack := HStack(child1, child2)

	layout := Layout{X: 0, Y: 0, Width: 10, Height: 10}
	output := stack.Render(layout)

	expected := "AB"
	if output != expected {
		t.Errorf("Render() = %q, want %q", output, expected)
	}
}

// Step 5: Test that StackProps can be created
func TestStackProps_WithValues_StoresAllFields(t *testing.T) {
	props := StackProps{
		Gap:            5,
		Padding:        SpacingAll(2),
		AlignItems:     AlignCenter,
		JustifyContent: JustifyCenter,
		Width:          DimensionFixed(100),
		Height:         DimensionFixed(50),
		Key:            "test-stack",
	}

	if props.Gap != 5 {
		t.Errorf("Gap = %d, want 5", props.Gap)
	}
	if props.Padding.Top != 2 {
		t.Errorf("Padding.Top = %d, want 2", props.Padding.Top)
	}
	if props.AlignItems != AlignCenter {
		t.Error("AlignItems should be AlignCenter")
	}
	if props.JustifyContent != JustifyCenter {
		t.Error("JustifyContent should be JustifyCenter")
	}
	if props.Key != "test-stack" {
		t.Errorf("Key = %q, want %q", props.Key, "test-stack")
	}
}

// Step 6: Test VStackWithProps with gap
func TestVStackWithProps_WithGap_AppliesGap(t *testing.T) {
	child1 := &mockComponent{key: "child1", width: 10, height: 5}
	child2 := &mockComponent{key: "child2", width: 10, height: 5}

	props := StackProps{
		Gap: 3,
	}
	stack := VStackWithProps(props, child1, child2)

	// Measure to verify gap is applied
	size := stack.Measure(100, 100)

	// Height should be: 5 + 3 (gap) + 5 = 13
	expectedHeight := 13
	if size.Height != expectedHeight {
		t.Errorf("Measure().Height = %d, want %d", size.Height, expectedHeight)
	}
}

// Step 7: Test VStackWithProps with padding
func TestVStackWithProps_WithPadding_AppliesPadding(t *testing.T) {
	child := &mockComponent{key: "child", width: 10, height: 5}

	props := StackProps{
		Padding: SpacingAll(2),
	}
	stack := VStackWithProps(props, child)

	size := stack.Measure(100, 100)

	// Width should be: 10 + 2 (left) + 2 (right) = 14
	// Height should be: 5 + 2 (top) + 2 (bottom) = 9
	if size.Width != 14 {
		t.Errorf("Measure().Width = %d, want 14", size.Width)
	}
	if size.Height != 9 {
		t.Errorf("Measure().Height = %d, want 9", size.Height)
	}
}

// Step 8: Test VStackWithProps with Key
func TestVStackWithProps_WithKey_SetsKey(t *testing.T) {
	props := StackProps{
		Key: "my-vstack",
	}
	stack := VStackWithProps(props)

	got := stack.Key()
	want := "my-vstack"
	if got != want {
		t.Errorf("Key() = %q, want %q", got, want)
	}
}

// Step 9: Test HStackWithProps with gap
func TestHStackWithProps_WithGap_AppliesGap(t *testing.T) {
	child1 := &mockComponent{key: "child1", width: 10, height: 5}
	child2 := &mockComponent{key: "child2", width: 10, height: 5}

	props := StackProps{
		Gap: 3,
	}
	stack := HStackWithProps(props, child1, child2)

	size := stack.Measure(100, 100)

	// Width should be: 10 + 3 (gap) + 10 = 23
	expectedWidth := 23
	if size.Width != expectedWidth {
		t.Errorf("Measure().Width = %d, want %d", size.Width, expectedWidth)
	}
}

// Step 10: Test HStackWithProps with padding
func TestHStackWithProps_WithPadding_AppliesPadding(t *testing.T) {
	child := &mockComponent{key: "child", width: 10, height: 5}

	props := StackProps{
		Padding: SpacingAll(3),
	}
	stack := HStackWithProps(props, child)

	size := stack.Measure(100, 100)

	// Width should be: 10 + 3 (left) + 3 (right) = 16
	// Height should be: 5 + 3 (top) + 3 (bottom) = 11
	if size.Width != 16 {
		t.Errorf("Measure().Width = %d, want 16", size.Width)
	}
	if size.Height != 11 {
		t.Errorf("Measure().Height = %d, want 11", size.Height)
	}
}

// Step 11: Test HStackWithProps with Key
func TestHStackWithProps_WithKey_SetsKey(t *testing.T) {
	props := StackProps{
		Key: "my-hstack",
	}
	stack := HStackWithProps(props)

	got := stack.Key()
	want := "my-hstack"
	if got != want {
		t.Errorf("Key() = %q, want %q", got, want)
	}
}

// Step 12: Test VStackWithProps with Width dimension
func TestVStackWithProps_WithFixedWidth_AppliesWidth(t *testing.T) {
	child := &mockComponent{key: "child", width: 10, height: 5}

	props := StackProps{
		Width: DimensionFixed(50),
	}
	stack := VStackWithProps(props, child)

	size := stack.Measure(100, 100)

	// Width should be 50 (fixed dimension)
	if size.Width != 50 {
		t.Errorf("Measure().Width = %d, want 50", size.Width)
	}
}

// Step 13: Test HStackWithProps with Height dimension
func TestHStackWithProps_WithFixedHeight_AppliesHeight(t *testing.T) {
	child := &mockComponent{key: "child", width: 10, height: 5}

	props := StackProps{
		Height: DimensionFixed(20),
	}
	stack := HStackWithProps(props, child)

	size := stack.Measure(100, 100)

	// Height should be 20 (fixed dimension)
	if size.Height != 20 {
		t.Errorf("Measure().Height = %d, want 20", size.Height)
	}
}

// Step 14: Test VStackWithProps with AlignItems center
func TestVStackWithProps_WithAlignCenter_AppliesAlignment(t *testing.T) {
	child := &mockComponent{key: "child", width: 10, height: 5}

	props := StackProps{
		AlignItems: AlignCenter,
	}
	stack := VStackWithProps(props, child)

	// Verify stack can be created with alignment
	// Actual alignment behavior is tested in layout tests
	if stack == nil {
		t.Fatal("VStackWithProps should not return nil")
	}
}

// Step 15: Test HStackWithProps with JustifyContent center
func TestHStackWithProps_WithJustifyCenter_AppliesJustification(t *testing.T) {
	child := &mockComponent{key: "child", width: 10, height: 5}

	props := StackProps{
		JustifyContent: JustifyCenter,
	}
	stack := HStackWithProps(props, child)

	// Verify stack can be created with justification
	// Actual justification behavior is tested in layout tests
	if stack == nil {
		t.Fatal("HStackWithProps should not return nil")
	}
}
