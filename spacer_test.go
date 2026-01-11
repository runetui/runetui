package runetui

import "testing"

func TestSpacer_WithFixedSize_CreatesBoxWithFixedWidth(t *testing.T) {
	spacer := Spacer(10)

	if spacer == nil {
		t.Fatal("Spacer should not return nil")
	}

	// Spacer should be a Box
	box, ok := spacer.(*box)
	if !ok {
		t.Fatal("Spacer should return a Box component")
	}

	// Width should be DimensionFixed(10)
	fixed, ok := box.props.Width.(dimensionFixed)
	if !ok {
		t.Fatal("Spacer width should be dimensionFixed")
	}

	if got := fixed.Value(); got != 10 {
		t.Errorf("expected width 10, got %d", got)
	}
}

func TestSpacer_WithFixedSize_CreatesBoxWithFixedHeight(t *testing.T) {
	spacer := Spacer(15)

	if spacer == nil {
		t.Fatal("Spacer should not return nil")
	}

	// Spacer should be a Box
	box, ok := spacer.(*box)
	if !ok {
		t.Fatal("Spacer should return a Box component")
	}

	// Height should be DimensionFixed(15)
	fixed, ok := box.props.Height.(dimensionFixed)
	if !ok {
		t.Fatal("Spacer height should be dimensionFixed")
	}

	if got := fixed.Value(); got != 15 {
		t.Errorf("expected height 15, got %d", got)
	}
}

func TestFlexSpacer_CreatesBoxWithFlexGrow(t *testing.T) {
	spacer := FlexSpacer()

	if spacer == nil {
		t.Fatal("FlexSpacer should not return nil")
	}

	// FlexSpacer should be a Box
	box, ok := spacer.(*box)
	if !ok {
		t.Fatal("FlexSpacer should return a Box component")
	}

	// FlexGrow should be 1.0
	if got := box.props.FlexGrow; got != 1.0 {
		t.Errorf("expected FlexGrow 1.0, got %f", got)
	}
}

func TestSpacer_WithZeroSize_CreatesBoxWithZeroDimensions(t *testing.T) {
	spacer := Spacer(0)

	if spacer == nil {
		t.Fatal("Spacer should not return nil")
	}

	box, ok := spacer.(*box)
	if !ok {
		t.Fatal("Spacer should return a Box component")
	}

	// Width should be DimensionFixed(0)
	fixedWidth, ok := box.props.Width.(dimensionFixed)
	if !ok {
		t.Fatal("Spacer width should be dimensionFixed")
	}

	if got := fixedWidth.Value(); got != 0 {
		t.Errorf("expected width 0, got %d", got)
	}

	// Height should be DimensionFixed(0)
	fixedHeight, ok := box.props.Height.(dimensionFixed)
	if !ok {
		t.Fatal("Spacer height should be dimensionFixed")
	}

	if got := fixedHeight.Value(); got != 0 {
		t.Errorf("expected height 0, got %d", got)
	}
}

func TestSpacer_WithNegativeSize_CreatesBoxWithNegativeDimensions(t *testing.T) {
	spacer := Spacer(-5)

	if spacer == nil {
		t.Fatal("Spacer should not return nil")
	}

	box, ok := spacer.(*box)
	if !ok {
		t.Fatal("Spacer should return a Box component")
	}

	// Width should be DimensionFixed(-5)
	fixedWidth, ok := box.props.Width.(dimensionFixed)
	if !ok {
		t.Fatal("Spacer width should be dimensionFixed")
	}

	if got := fixedWidth.Value(); got != -5 {
		t.Errorf("expected width -5, got %d", got)
	}

	// Height should be DimensionFixed(-5)
	fixedHeight, ok := box.props.Height.(dimensionFixed)
	if !ok {
		t.Fatal("Spacer height should be dimensionFixed")
	}

	if got := fixedHeight.Value(); got != -5 {
		t.Errorf("expected height -5, got %d", got)
	}
}

func TestSpacer_HasNoChildren(t *testing.T) {
	spacer := Spacer(10)

	if spacer == nil {
		t.Fatal("Spacer should not return nil")
	}

	children := spacer.Children()
	if children == nil {
		t.Fatal("Children() should return empty slice, not nil")
	}

	if got := len(children); got != 0 {
		t.Errorf("expected 0 children, got %d", got)
	}
}

func TestFlexSpacer_HasNoChildren(t *testing.T) {
	spacer := FlexSpacer()

	if spacer == nil {
		t.Fatal("FlexSpacer should not return nil")
	}

	children := spacer.Children()
	if children == nil {
		t.Fatal("Children() should return empty slice, not nil")
	}

	if got := len(children); got != 0 {
		t.Errorf("expected 0 children, got %d", got)
	}
}
