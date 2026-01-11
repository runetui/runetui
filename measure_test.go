package runetui

import "testing"

func TestResolveDimension_Fixed_ReturnsFixedValue(t *testing.T) {
	dim := DimensionFixed(100)
	result := resolveDimension(dim, 200)
	if result != 100 {
		t.Errorf("expected 100, got %d", result)
	}
}

func TestResolveDimension_Auto_ReturnsZero(t *testing.T) {
	dim := DimensionAuto()
	result := resolveDimension(dim, 200)
	if result != 0 {
		t.Errorf("expected 0 for Auto dimension, got %d", result)
	}
}

func TestResolveDimension_Percent_CalculatesPercentage(t *testing.T) {
	dim := DimensionPercent(50)
	result := resolveDimension(dim, 200)
	if result != 100 {
		t.Errorf("expected 100 (50%% of 200), got %d", result)
	}
}

func TestResolveDimension_Percent100_ReturnsFullAvailable(t *testing.T) {
	dim := DimensionPercent(100)
	result := resolveDimension(dim, 200)
	if result != 200 {
		t.Errorf("expected 200 (100%% of 200), got %d", result)
	}
}

func TestResolveDimension_PercentZero_ReturnsZero(t *testing.T) {
	dim := DimensionPercent(0)
	result := resolveDimension(dim, 200)
	if result != 0 {
		t.Errorf("expected 0 (0%% of 200), got %d", result)
	}
}

func TestMeasureText_SingleLine_ReturnsCorrectSize(t *testing.T) {
	size := measureText("hello", WrapNone, 100)
	if size.Width != 5 {
		t.Errorf("expected width 5, got %d", size.Width)
	}
	if size.Height != 1 {
		t.Errorf("expected height 1, got %d", size.Height)
	}
}

func TestMeasureText_EmptyString_ReturnsZeroSize(t *testing.T) {
	size := measureText("", WrapNone, 100)
	if size.Width != 0 || size.Height != 0 {
		t.Errorf("expected zero size, got width=%d height=%d", size.Width, size.Height)
	}
}

func TestMeasureText_MultiLine_ReturnsMaxWidth(t *testing.T) {
	size := measureText("hi\nhello\ngo", WrapNone, 100)
	if size.Width != 5 {
		t.Errorf("expected width 5 (max line length), got %d", size.Width)
	}
	if size.Height != 3 {
		t.Errorf("expected height 3 (three lines), got %d", size.Height)
	}
}

func TestMeasureText_Unicode_CountsRunes(t *testing.T) {
	size := measureText("こんにちは", WrapNone, 100)
	if size.Width != 5 {
		t.Errorf("expected width 5 (5 runes), got %d", size.Width)
	}
	if size.Height != 1 {
		t.Errorf("expected height 1, got %d", size.Height)
	}
}

func TestMeasureText_WrapTruncate_ConstrainsWidth(t *testing.T) {
	size := measureText("hello world", WrapTruncate, 5)
	if size.Width != 5 {
		t.Errorf("expected width 5 (truncated), got %d", size.Width)
	}
	if size.Height != 1 {
		t.Errorf("expected height 1, got %d", size.Height)
	}
}

func TestMeasureText_WrapWord_WrapsToMultipleLines(t *testing.T) {
	size := measureText("hello world", WrapWord, 5)
	if size.Width != 5 {
		t.Errorf("expected width 5 (constrained), got %d", size.Width)
	}
	if size.Height != 3 {
		t.Errorf("expected height 3 (11 chars / 5 = 3 lines), got %d", size.Height)
	}
}

func TestMeasureText_WrapChar_WrapsToMultipleLines(t *testing.T) {
	size := measureText("helloworld", WrapChar, 5)
	if size.Width != 5 {
		t.Errorf("expected width 5 (constrained), got %d", size.Width)
	}
	if size.Height != 2 {
		t.Errorf("expected height 2 (10 chars / 5 = 2 lines), got %d", size.Height)
	}
}

func TestMeasureBox_EmptyBox_ReturnsZeroSize(t *testing.T) {
	props := BoxProps{}
	size := measureBox(props, []Component{}, 100, 100)
	if size.Width != 0 || size.Height != 0 {
		t.Errorf("expected zero size for empty box, got width=%d height=%d", size.Width, size.Height)
	}
}

func TestMeasureBox_ColumnDirection_SumsHeight(t *testing.T) {
	props := BoxProps{Direction: Column}
	children := []Component{
		Text("line1"),
		Text("line2"),
		Text("line3"),
	}
	size := measureBox(props, children, 100, 100)
	if size.Height != 3 {
		t.Errorf("expected height 3 (sum of children), got %d", size.Height)
	}
	if size.Width != 5 {
		t.Errorf("expected width 5 (max of children), got %d", size.Width)
	}
}

func TestMeasureBox_RowDirection_SumsWidth(t *testing.T) {
	props := BoxProps{Direction: Row}
	children := []Component{
		Text("hi"),
		Text("go"),
	}
	size := measureBox(props, children, 100, 100)
	if size.Width != 4 {
		t.Errorf("expected width 4 (sum of children), got %d", size.Width)
	}
	if size.Height != 1 {
		t.Errorf("expected height 1 (max of children), got %d", size.Height)
	}
}

func TestSpacingWidth_WithSpacing_ReturnsTotalWidth(t *testing.T) {
	spacing := Spacing{Left: 2, Right: 3}
	width := spacingWidth(spacing)
	if width != 5 {
		t.Errorf("expected width 5 (2+3), got %d", width)
	}
}

func TestSpacingHeight_WithSpacing_ReturnsTotalHeight(t *testing.T) {
	spacing := Spacing{Top: 1, Bottom: 4}
	height := spacingHeight(spacing)
	if height != 5 {
		t.Errorf("expected height 5 (1+4), got %d", height)
	}
}

func TestBorderSize_WithNoBorder_ReturnsZero(t *testing.T) {
	width, height := borderSize(BorderNone)
	if width != 0 || height != 0 {
		t.Errorf("expected 0,0 for no border, got %d,%d", width, height)
	}
}

func TestBorderSize_WithBorder_ReturnsTwoByTwo(t *testing.T) {
	width, height := borderSize(BorderSingle)
	if width != 2 || height != 2 {
		t.Errorf("expected 2,2 for border, got %d,%d", width, height)
	}
}

func TestMeasureBox_WithPadding_AddsToPaddingToSize(t *testing.T) {
	props := BoxProps{
		Direction: Column,
		Padding:   SpacingAll(2),
	}
	children := []Component{Text("hi")}
	size := measureBox(props, children, 100, 100)
	if size.Width != 6 {
		t.Errorf("expected width 6 (2+2+2 padding), got %d", size.Width)
	}
	if size.Height != 5 {
		t.Errorf("expected height 5 (1+2+2 padding), got %d", size.Height)
	}
}

func TestMeasureBox_WithMargin_AddsMarginToSize(t *testing.T) {
	props := BoxProps{
		Direction: Column,
		Margin:    SpacingAll(1),
	}
	children := []Component{Text("hi")}
	size := measureBox(props, children, 100, 100)
	if size.Width != 4 {
		t.Errorf("expected width 4 (2+1+1 margin), got %d", size.Width)
	}
	if size.Height != 3 {
		t.Errorf("expected height 3 (1+1+1 margin), got %d", size.Height)
	}
}

func TestMeasureBox_WithBorder_AddsBorderToSize(t *testing.T) {
	props := BoxProps{
		Direction: Column,
		Border:    BorderSingle,
	}
	children := []Component{Text("hi")}
	size := measureBox(props, children, 100, 100)
	if size.Width != 4 {
		t.Errorf("expected width 4 (2+2 border), got %d", size.Width)
	}
	if size.Height != 3 {
		t.Errorf("expected height 3 (1+2 border), got %d", size.Height)
	}
}

func TestMeasureBox_WithGap_AddsGapBetweenChildren(t *testing.T) {
	props := BoxProps{
		Direction: Column,
		Gap:       2,
	}
	children := []Component{
		Text("a"),
		Text("b"),
		Text("c"),
	}
	size := measureBox(props, children, 100, 100)
	if size.Height != 7 {
		t.Errorf("expected height 7 (3 lines + 2*2 gap), got %d", size.Height)
	}
}

func TestApplyConstraints_WithMinWidth_ClampsToMin(t *testing.T) {
	size := Size{Width: 5, Height: 10}
	result := applyConstraints(size, 10, 0, 0, 0)
	if result.Width != 10 {
		t.Errorf("expected width 10 (clamped to min), got %d", result.Width)
	}
}

func TestApplyConstraints_WithMaxWidth_ClampsToMax(t *testing.T) {
	size := Size{Width: 20, Height: 10}
	result := applyConstraints(size, 0, 0, 15, 0)
	if result.Width != 15 {
		t.Errorf("expected width 15 (clamped to max), got %d", result.Width)
	}
}

func TestApplyConstraints_WithMinHeight_ClampsToMin(t *testing.T) {
	size := Size{Width: 10, Height: 3}
	result := applyConstraints(size, 0, 5, 0, 0)
	if result.Height != 5 {
		t.Errorf("expected height 5 (clamped to min), got %d", result.Height)
	}
}

func TestApplyConstraints_WithMaxHeight_ClampsToMax(t *testing.T) {
	size := Size{Width: 10, Height: 20}
	result := applyConstraints(size, 0, 0, 0, 15)
	if result.Height != 15 {
		t.Errorf("expected height 15 (clamped to max), got %d", result.Height)
	}
}

func TestMeasureBox_WithMinWidth_AppliesMinConstraint(t *testing.T) {
	props := BoxProps{
		Direction: Column,
		MinWidth:  20,
	}
	children := []Component{Text("hi")}
	size := measureBox(props, children, 100, 100)
	if size.Width != 20 {
		t.Errorf("expected width 20 (min constraint), got %d", size.Width)
	}
}

func TestMeasureBox_WithFixedWidth_UsesFixedValue(t *testing.T) {
	props := BoxProps{
		Direction: Column,
		Width:     DimensionFixed(50),
	}
	children := []Component{Text("hi")}
	size := measureBox(props, children, 100, 100)
	if size.Width != 50 {
		t.Errorf("expected width 50 (fixed), got %d", size.Width)
	}
}

func TestMeasureBox_WithPercentWidth_CalculatesFromAvailable(t *testing.T) {
	props := BoxProps{
		Direction: Column,
		Width:     DimensionPercent(50),
	}
	children := []Component{Text("hi")}
	size := measureBox(props, children, 100, 100)
	if size.Width != 50 {
		t.Errorf("expected width 50 (50%% of 100), got %d", size.Width)
	}
}

func TestMeasureText_WrapTruncate_NoTruncateNeeded(t *testing.T) {
	size := measureText("hi", WrapTruncate, 10)
	if size.Width != 2 {
		t.Errorf("expected width 2 (no truncate), got %d", size.Width)
	}
	if size.Height != 1 {
		t.Errorf("expected height 1, got %d", size.Height)
	}
}

func TestMeasureText_WrapWord_NoWrapNeeded(t *testing.T) {
	size := measureText("hi", WrapWord, 10)
	if size.Width != 2 {
		t.Errorf("expected width 2 (no wrap), got %d", size.Width)
	}
	if size.Height != 1 {
		t.Errorf("expected height 1, got %d", size.Height)
	}
}

func TestMeasureBox_RowDirection_WithGap(t *testing.T) {
	props := BoxProps{
		Direction: Row,
		Gap:       1,
	}
	children := []Component{
		Text("a"),
		Text("b"),
	}
	size := measureBox(props, children, 100, 100)
	if size.Width != 3 {
		t.Errorf("expected width 3 (1+1+1 gap), got %d", size.Width)
	}
}

func TestMeasureBox_WithAutoHeight_UsesIntrinsicHeight(t *testing.T) {
	props := BoxProps{
		Direction: Column,
		Height:    DimensionAuto(),
	}
	children := []Component{Text("hi")}
	size := measureBox(props, children, 100, 100)
	if size.Height != 1 {
		t.Errorf("expected height 1 (intrinsic), got %d", size.Height)
	}
}

func TestMeasureBox_ComplexScenario_AllFeaturesWork(t *testing.T) {
	props := BoxProps{
		Direction: Row,
		Gap:       2,
		Padding:   SpacingAll(1),
		Margin:    SpacingAll(1),
		Border:    BorderSingle,
		MinWidth:  0,
		MinHeight: 0,
		MaxWidth:  0,
		MaxHeight: 0,
	}
	children := []Component{
		Text("a"),
		Text("b"),
	}
	size := measureBox(props, children, 100, 100)
	expected := 1 + 1 + 2 + 2 + 2 + 2
	if size.Width != expected {
		t.Errorf("expected width %d (1+1+2gap+2pad+2mar+2bor), got %d", expected, size.Width)
	}
}
