package runetui

import (
	"testing"
)

func TestFlexChild_CanBeCreated_WithAllFields(t *testing.T) {
	component := Box(BoxProps{Key: "test"})
	size := Size{Width: 100, Height: 50}

	child := FlexChild{
		Component:  component,
		Size:       size,
		FlexGrow:   1.5,
		FlexShrink: 0.5,
	}

	if child.Component == nil {
		t.Error("Component should not be nil")
	}
	if child.Size.Width != 100 {
		t.Errorf("expected Size.Width=100, got %d", child.Size.Width)
	}
	if child.Size.Height != 50 {
		t.Errorf("expected Size.Height=50, got %d", child.Size.Height)
	}
	if child.FlexGrow != 1.5 {
		t.Errorf("expected FlexGrow=1.5, got %f", child.FlexGrow)
	}
	if child.FlexShrink != 0.5 {
		t.Errorf("expected FlexShrink=0.5, got %f", child.FlexShrink)
	}
}

func TestCalculateFlexGrow_SingleChild_DistributesAllExtraSpace(t *testing.T) {
	children := []FlexChild{
		{FlexGrow: 1.0},
	}

	result := calculateFlexGrow(children, 10)

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}
	if result[0] != 10 {
		t.Errorf("expected [10], got [%d]", result[0])
	}
}

func TestCalculateFlexGrow_TwoChildren_DistributesProportionally(t *testing.T) {
	children := []FlexChild{
		{FlexGrow: 2.0},
		{FlexGrow: 1.0},
	}

	result := calculateFlexGrow(children, 30)

	if len(result) != 2 {
		t.Fatalf("expected 2 results, got %d", len(result))
	}
	if result[0] != 20 {
		t.Errorf("expected result[0]=20, got %d", result[0])
	}
	if result[1] != 10 {
		t.Errorf("expected result[1]=10, got %d", result[1])
	}
}

func TestCalculateFlexGrow_ZeroExtraSpace_ReturnsAllZeros(t *testing.T) {
	children := []FlexChild{
		{FlexGrow: 1.0},
		{FlexGrow: 2.0},
	}

	result := calculateFlexGrow(children, 0)

	for i, val := range result {
		if val != 0 {
			t.Errorf("expected result[%d]=0, got %d", i, val)
		}
	}
}

func TestCalculateFlexGrow_NegativeExtraSpace_ReturnsAllZeros(t *testing.T) {
	children := []FlexChild{
		{FlexGrow: 1.0},
		{FlexGrow: 2.0},
	}

	result := calculateFlexGrow(children, -10)

	for i, val := range result {
		if val != 0 {
			t.Errorf("expected result[%d]=0, got %d", i, val)
		}
	}
}

func TestCalculateFlexGrow_AllZeroGrowValues_ReturnsAllZeros(t *testing.T) {
	children := []FlexChild{
		{FlexGrow: 0.0},
		{FlexGrow: 0.0},
	}

	result := calculateFlexGrow(children, 30)

	for i, val := range result {
		if val != 0 {
			t.Errorf("expected result[%d]=0, got %d", i, val)
		}
	}
}

func TestCalculateFlexShrink_SingleChild_ShrinksAllDeficit(t *testing.T) {
	children := []FlexChild{
		{FlexShrink: 1.0},
	}

	result := calculateFlexShrink(children, 10)

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}
	if result[0] != 10 {
		t.Errorf("expected [10], got [%d]", result[0])
	}
}

func TestCalculateFlexShrink_TwoChildren_ShrinksProportionally(t *testing.T) {
	children := []FlexChild{
		{FlexShrink: 2.0},
		{FlexShrink: 1.0},
	}

	result := calculateFlexShrink(children, 30)

	if len(result) != 2 {
		t.Fatalf("expected 2 results, got %d", len(result))
	}
	if result[0] != 20 {
		t.Errorf("expected result[0]=20, got %d", result[0])
	}
	if result[1] != 10 {
		t.Errorf("expected result[1]=10, got %d", result[1])
	}
}

func TestCalculateFlexShrink_ZeroDeficit_ReturnsAllZeros(t *testing.T) {
	children := []FlexChild{
		{FlexShrink: 1.0},
		{FlexShrink: 2.0},
	}

	result := calculateFlexShrink(children, 0)

	for i, val := range result {
		if val != 0 {
			t.Errorf("expected result[%d]=0, got %d", i, val)
		}
	}
}

func TestCalculateFlexShrink_NegativeDeficit_ReturnsAllZeros(t *testing.T) {
	children := []FlexChild{
		{FlexShrink: 1.0},
		{FlexShrink: 2.0},
	}

	result := calculateFlexShrink(children, -10)

	for i, val := range result {
		if val != 0 {
			t.Errorf("expected result[%d]=0, got %d", i, val)
		}
	}
}

func TestCalculateFlexShrink_AllZeroShrinkValues_ReturnsAllZeros(t *testing.T) {
	children := []FlexChild{
		{FlexShrink: 0.0},
		{FlexShrink: 0.0},
	}

	result := calculateFlexShrink(children, 30)

	for i, val := range result {
		if val != 0 {
			t.Errorf("expected result[%d]=0, got %d", i, val)
		}
	}
}

func TestAlignItems_AlignStart_Column_KeepsXAtZero(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 50, Height: 20}},
		{Layout: Layout{X: 0, Y: 20, Width: 30, Height: 20}},
	}
	props := BoxProps{Direction: Column, AlignItems: AlignStart}

	alignItems(children, props, 100)

	if children[0].Layout.X != 0 {
		t.Errorf("expected children[0].Layout.X=0, got %d", children[0].Layout.X)
	}
	if children[1].Layout.X != 0 {
		t.Errorf("expected children[1].Layout.X=0, got %d", children[1].Layout.X)
	}
}

func TestAlignItems_AlignCenter_Column_CentersOnXAxis(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 50, Height: 20}},
		{Layout: Layout{X: 0, Y: 20, Width: 30, Height: 20}},
	}
	props := BoxProps{Direction: Column, AlignItems: AlignCenter}

	alignItems(children, props, 100)

	if children[0].Layout.X != 25 {
		t.Errorf("expected children[0].Layout.X=25 (centered), got %d", children[0].Layout.X)
	}
	if children[1].Layout.X != 35 {
		t.Errorf("expected children[1].Layout.X=35 (centered), got %d", children[1].Layout.X)
	}
}

func TestAlignItems_AlignEnd_Column_AlignsToEndOnXAxis(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 50, Height: 20}},
		{Layout: Layout{X: 0, Y: 20, Width: 30, Height: 20}},
	}
	props := BoxProps{Direction: Column, AlignItems: AlignEnd}

	alignItems(children, props, 100)

	if children[0].Layout.X != 50 {
		t.Errorf("expected children[0].Layout.X=50 (100-50), got %d", children[0].Layout.X)
	}
	if children[1].Layout.X != 70 {
		t.Errorf("expected children[1].Layout.X=70 (100-30), got %d", children[1].Layout.X)
	}
}

func TestAlignItems_AlignStretch_Column_SetsWidthToContainer(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 50, Height: 20}},
		{Layout: Layout{X: 0, Y: 20, Width: 30, Height: 20}},
	}
	props := BoxProps{Direction: Column, AlignItems: AlignStretch}

	alignItems(children, props, 100)

	if children[0].Layout.Width != 100 {
		t.Errorf("expected children[0].Layout.Width=100 (stretched), got %d", children[0].Layout.Width)
	}
	if children[1].Layout.Width != 100 {
		t.Errorf("expected children[1].Layout.Width=100 (stretched), got %d", children[1].Layout.Width)
	}
}

func TestAlignItems_AlignStart_Row_KeepsYAtZero(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 50, Height: 20}},
		{Layout: Layout{X: 50, Y: 0, Width: 30, Height: 30}},
	}
	props := BoxProps{Direction: Row, AlignItems: AlignStart}

	alignItems(children, props, 100)

	if children[0].Layout.Y != 0 {
		t.Errorf("expected children[0].Layout.Y=0, got %d", children[0].Layout.Y)
	}
	if children[1].Layout.Y != 0 {
		t.Errorf("expected children[1].Layout.Y=0, got %d", children[1].Layout.Y)
	}
}

func TestAlignItems_AlignCenter_Row_CentersOnYAxis(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 50, Height: 20}},
		{Layout: Layout{X: 50, Y: 0, Width: 30, Height: 30}},
	}
	props := BoxProps{Direction: Row, AlignItems: AlignCenter}

	alignItems(children, props, 100)

	if children[0].Layout.Y != 40 {
		t.Errorf("expected children[0].Layout.Y=40 (centered), got %d", children[0].Layout.Y)
	}
	if children[1].Layout.Y != 35 {
		t.Errorf("expected children[1].Layout.Y=35 (centered), got %d", children[1].Layout.Y)
	}
}

func TestAlignItems_AlignEnd_Row_AlignsToEndOnYAxis(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 50, Height: 20}},
		{Layout: Layout{X: 50, Y: 0, Width: 30, Height: 30}},
	}
	props := BoxProps{Direction: Row, AlignItems: AlignEnd}

	alignItems(children, props, 100)

	if children[0].Layout.Y != 80 {
		t.Errorf("expected children[0].Layout.Y=80 (100-20), got %d", children[0].Layout.Y)
	}
	if children[1].Layout.Y != 70 {
		t.Errorf("expected children[1].Layout.Y=70 (100-30), got %d", children[1].Layout.Y)
	}
}

func TestAlignItems_AlignStretch_Row_SetsHeightToContainer(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 50, Height: 20}},
		{Layout: Layout{X: 50, Y: 0, Width: 30, Height: 30}},
	}
	props := BoxProps{Direction: Row, AlignItems: AlignStretch}

	alignItems(children, props, 100)

	if children[0].Layout.Height != 100 {
		t.Errorf("expected children[0].Layout.Height=100 (stretched), got %d", children[0].Layout.Height)
	}
	if children[1].Layout.Height != 100 {
		t.Errorf("expected children[1].Layout.Height=100 (stretched), got %d", children[1].Layout.Height)
	}
}

func TestJustifyContent_JustifyStart_Column_KeepsYPositions(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 10, Width: 50, Height: 20}},
		{Layout: Layout{X: 0, Y: 30, Width: 50, Height: 20}},
	}
	props := BoxProps{Direction: Column, JustifyContent: JustifyStart}

	justifyContent(children, props, 100)

	if children[0].Layout.Y != 10 {
		t.Errorf("expected children[0].Layout.Y=10 (unchanged), got %d", children[0].Layout.Y)
	}
	if children[1].Layout.Y != 30 {
		t.Errorf("expected children[1].Layout.Y=30 (unchanged), got %d", children[1].Layout.Y)
	}
}

func TestJustifyContent_JustifyCenter_Column_CentersGroupOnYAxis(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 50, Height: 20}},
		{Layout: Layout{X: 0, Y: 20, Width: 50, Height: 20}},
	}
	props := BoxProps{Direction: Column, JustifyContent: JustifyCenter}

	justifyContent(children, props, 100)

	if children[0].Layout.Y != 30 {
		t.Errorf("expected children[0].Layout.Y=30 (centered group), got %d", children[0].Layout.Y)
	}
	if children[1].Layout.Y != 50 {
		t.Errorf("expected children[1].Layout.Y=50 (centered group), got %d", children[1].Layout.Y)
	}
}

func TestJustifyContent_JustifyEnd_Column_PushesToEndOnYAxis(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 50, Height: 20}},
		{Layout: Layout{X: 0, Y: 20, Width: 50, Height: 20}},
	}
	props := BoxProps{Direction: Column, JustifyContent: JustifyEnd}

	justifyContent(children, props, 100)

	if children[0].Layout.Y != 60 {
		t.Errorf("expected children[0].Layout.Y=60 (pushed to end), got %d", children[0].Layout.Y)
	}
	if children[1].Layout.Y != 80 {
		t.Errorf("expected children[1].Layout.Y=80 (pushed to end), got %d", children[1].Layout.Y)
	}
}

func TestJustifyContent_JustifySpaceBetween_Column_DistributesSpace(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 50, Height: 20}},
		{Layout: Layout{X: 0, Y: 20, Width: 50, Height: 20}},
		{Layout: Layout{X: 0, Y: 40, Width: 50, Height: 20}},
	}
	props := BoxProps{Direction: Column, JustifyContent: JustifySpaceBetween}

	justifyContent(children, props, 100)

	if children[0].Layout.Y != 0 {
		t.Errorf("expected children[0].Layout.Y=0 (at start), got %d", children[0].Layout.Y)
	}
	if children[1].Layout.Y != 40 {
		t.Errorf("expected children[1].Layout.Y=40 (middle), got %d", children[1].Layout.Y)
	}
	if children[2].Layout.Y != 80 {
		t.Errorf("expected children[2].Layout.Y=80 (at end), got %d", children[2].Layout.Y)
	}
}

func TestJustifyContent_JustifySpaceAround_Column_DistributesAroundSpace(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 50, Height: 20}},
		{Layout: Layout{X: 0, Y: 20, Width: 50, Height: 20}},
	}
	props := BoxProps{Direction: Column, JustifyContent: JustifySpaceAround}

	justifyContent(children, props, 100)

	if children[0].Layout.Y != 15 {
		t.Errorf("expected children[0].Layout.Y=15 (with half-space before), got %d", children[0].Layout.Y)
	}
	if children[1].Layout.Y != 65 {
		t.Errorf("expected children[1].Layout.Y=65 (with space around), got %d", children[1].Layout.Y)
	}
}

func TestJustifyContent_JustifyStart_Row_KeepsXPositions(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 10, Y: 0, Width: 20, Height: 50}},
		{Layout: Layout{X: 30, Y: 0, Width: 20, Height: 50}},
	}
	props := BoxProps{Direction: Row, JustifyContent: JustifyStart}

	justifyContent(children, props, 100)

	if children[0].Layout.X != 10 {
		t.Errorf("expected children[0].Layout.X=10 (unchanged), got %d", children[0].Layout.X)
	}
	if children[1].Layout.X != 30 {
		t.Errorf("expected children[1].Layout.X=30 (unchanged), got %d", children[1].Layout.X)
	}
}

func TestJustifyContent_JustifyCenter_Row_CentersGroupOnXAxis(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 20, Height: 50}},
		{Layout: Layout{X: 20, Y: 0, Width: 20, Height: 50}},
	}
	props := BoxProps{Direction: Row, JustifyContent: JustifyCenter}

	justifyContent(children, props, 100)

	if children[0].Layout.X != 30 {
		t.Errorf("expected children[0].Layout.X=30 (centered group), got %d", children[0].Layout.X)
	}
	if children[1].Layout.X != 50 {
		t.Errorf("expected children[1].Layout.X=50 (centered group), got %d", children[1].Layout.X)
	}
}

func TestJustifyContent_JustifyEnd_Row_PushesToEndOnXAxis(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 20, Height: 50}},
		{Layout: Layout{X: 20, Y: 0, Width: 20, Height: 50}},
	}
	props := BoxProps{Direction: Row, JustifyContent: JustifyEnd}

	justifyContent(children, props, 100)

	if children[0].Layout.X != 60 {
		t.Errorf("expected children[0].Layout.X=60 (pushed to end), got %d", children[0].Layout.X)
	}
	if children[1].Layout.X != 80 {
		t.Errorf("expected children[1].Layout.X=80 (pushed to end), got %d", children[1].Layout.X)
	}
}

func TestJustifyContent_JustifySpaceBetween_Row_DistributesSpace(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 20, Height: 50}},
		{Layout: Layout{X: 20, Y: 0, Width: 20, Height: 50}},
		{Layout: Layout{X: 40, Y: 0, Width: 20, Height: 50}},
	}
	props := BoxProps{Direction: Row, JustifyContent: JustifySpaceBetween}

	justifyContent(children, props, 100)

	if children[0].Layout.X != 0 {
		t.Errorf("expected children[0].Layout.X=0 (at start), got %d", children[0].Layout.X)
	}
	if children[1].Layout.X != 40 {
		t.Errorf("expected children[1].Layout.X=40 (middle), got %d", children[1].Layout.X)
	}
	if children[2].Layout.X != 80 {
		t.Errorf("expected children[2].Layout.X=80 (at end), got %d", children[2].Layout.X)
	}
}

func TestJustifyContent_JustifySpaceAround_Row_DistributesAroundSpace(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 0, Width: 20, Height: 50}},
		{Layout: Layout{X: 20, Y: 0, Width: 20, Height: 50}},
	}
	props := BoxProps{Direction: Row, JustifyContent: JustifySpaceAround}

	justifyContent(children, props, 100)

	if children[0].Layout.X != 15 {
		t.Errorf("expected children[0].Layout.X=15 (with half-space before), got %d", children[0].Layout.X)
	}
	if children[1].Layout.X != 65 {
		t.Errorf("expected children[1].Layout.X=65 (with space around), got %d", children[1].Layout.X)
	}
}

func TestJustifyContent_EmptyChildren_DoesNothing(t *testing.T) {
	children := []*LayoutTree{}
	props := BoxProps{Direction: Column, JustifyContent: JustifyCenter}

	justifyContent(children, props, 100)

	if len(children) != 0 {
		t.Errorf("expected children to remain empty")
	}
}

func TestJustifyContent_JustifySpaceBetween_SingleChild_DoesNothing(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 10, Width: 50, Height: 20}},
	}
	props := BoxProps{Direction: Column, JustifyContent: JustifySpaceBetween}

	justifyContent(children, props, 100)

	if children[0].Layout.Y != 10 {
		t.Errorf("expected children[0].Layout.Y=10 (unchanged), got %d", children[0].Layout.Y)
	}
}

func TestJustifyContent_JustifySpaceBetween_SingleChild_Row_DoesNothing(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 10, Y: 0, Width: 20, Height: 50}},
	}
	props := BoxProps{Direction: Row, JustifyContent: JustifySpaceBetween}

	justifyContent(children, props, 100)

	if children[0].Layout.X != 10 {
		t.Errorf("expected children[0].Layout.X=10 (unchanged), got %d", children[0].Layout.X)
	}
}

func TestGetTotalHeight_WithEmptyChildren_ReturnsZero(t *testing.T) {
	children := []*LayoutTree{}

	result := getTotalHeight(children)

	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
}

func TestGetTotalHeight_WithSingleChild_ReturnsChildHeight(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 10, Width: 50, Height: 20}},
	}

	result := getTotalHeight(children)

	if result != 20 {
		t.Errorf("expected 20 (10 + 20 - 10), got %d", result)
	}
}

func TestGetTotalHeight_WithMultipleChildren_ReturnsSpan(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 0, Y: 10, Width: 50, Height: 20}},
		{Layout: Layout{X: 0, Y: 35, Width: 50, Height: 15}},
	}

	result := getTotalHeight(children)

	if result != 40 {
		t.Errorf("expected 40 (35 + 15 - 10), got %d", result)
	}
}

func TestGetTotalWidth_WithEmptyChildren_ReturnsZero(t *testing.T) {
	children := []*LayoutTree{}

	result := getTotalWidth(children)

	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
}

func TestGetTotalWidth_WithSingleChild_ReturnsChildWidth(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 10, Y: 0, Width: 20, Height: 50}},
	}

	result := getTotalWidth(children)

	if result != 20 {
		t.Errorf("expected 20 (10 + 20 - 10), got %d", result)
	}
}

func TestGetTotalWidth_WithMultipleChildren_ReturnsSpan(t *testing.T) {
	children := []*LayoutTree{
		{Layout: Layout{X: 10, Y: 0, Width: 20, Height: 50}},
		{Layout: Layout{X: 35, Y: 0, Width: 15, Height: 50}},
	}

	result := getTotalWidth(children)

	if result != 40 {
		t.Errorf("expected 40 (35 + 15 - 10), got %d", result)
	}
}
