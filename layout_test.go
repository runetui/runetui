package runetui

import "testing"

func TestLayoutEngine_SingleTextComponent_PositionedAtOrigin(t *testing.T) {
	engine := NewLayoutEngine(80, 24)
	textComponent := Text("Hello")

	layoutTree := engine.CalculateLayout(textComponent)

	if layoutTree.Layout.X != 0 {
		t.Errorf("expected X=0, got %d", layoutTree.Layout.X)
	}
	if layoutTree.Layout.Y != 0 {
		t.Errorf("expected Y=0, got %d", layoutTree.Layout.Y)
	}
	if layoutTree.Component != textComponent {
		t.Error("expected component to match input")
	}
}

func TestLayoutEngine_BoxWithColumnChildren_StacksVertically(t *testing.T) {
	engine := NewLayoutEngine(80, 24)
	child1 := Text("First")
	child2 := Text("Second")
	boxComponent := Box(BoxProps{Direction: Column}, child1, child2)

	layoutTree := engine.CalculateLayout(boxComponent)

	if len(layoutTree.Children) != 2 {
		t.Fatalf("expected 2 children, got %d", len(layoutTree.Children))
	}

	firstChild := layoutTree.Children[0]
	secondChild := layoutTree.Children[1]

	if firstChild.Layout.X != 0 {
		t.Errorf("first child X: expected 0, got %d", firstChild.Layout.X)
	}
	if firstChild.Layout.Y != 0 {
		t.Errorf("first child Y: expected 0, got %d", firstChild.Layout.Y)
	}

	if secondChild.Layout.X != 0 {
		t.Errorf("second child X: expected 0, got %d", secondChild.Layout.X)
	}
	expectedY := firstChild.Layout.Height
	if secondChild.Layout.Y != expectedY {
		t.Errorf("second child Y: expected %d, got %d", expectedY, secondChild.Layout.Y)
	}
}

func TestLayoutEngine_BoxWithRowChildren_PlacesHorizontally(t *testing.T) {
	engine := NewLayoutEngine(80, 24)
	child1 := Text("First")
	child2 := Text("Second")
	boxComponent := Box(BoxProps{Direction: Row}, child1, child2)

	layoutTree := engine.CalculateLayout(boxComponent)

	if len(layoutTree.Children) != 2 {
		t.Fatalf("expected 2 children, got %d", len(layoutTree.Children))
	}

	firstChild := layoutTree.Children[0]
	secondChild := layoutTree.Children[1]

	if firstChild.Layout.X != 0 {
		t.Errorf("first child X: expected 0, got %d", firstChild.Layout.X)
	}
	if firstChild.Layout.Y != 0 {
		t.Errorf("first child Y: expected 0, got %d", firstChild.Layout.Y)
	}

	expectedX := firstChild.Layout.Width
	if secondChild.Layout.X != expectedX {
		t.Errorf("second child X: expected %d, got %d", expectedX, secondChild.Layout.X)
	}
	if secondChild.Layout.Y != 0 {
		t.Errorf("second child Y: expected 0, got %d", secondChild.Layout.Y)
	}
}

func TestLayoutEngine_BoxWithPadding_AdjustsChildPosition(t *testing.T) {
	engine := NewLayoutEngine(80, 24)
	child := Text("Child")
	padding := Spacing{Top: 2, Right: 3, Bottom: 2, Left: 5}
	boxComponent := Box(BoxProps{Direction: Column, Padding: padding}, child)

	layoutTree := engine.CalculateLayout(boxComponent)

	if len(layoutTree.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(layoutTree.Children))
	}

	childLayout := layoutTree.Children[0].Layout

	expectedX := padding.Left
	if childLayout.X != expectedX {
		t.Errorf("child X: expected %d (padding.Left), got %d", expectedX, childLayout.X)
	}

	expectedY := padding.Top
	if childLayout.Y != expectedY {
		t.Errorf("child Y: expected %d (padding.Top), got %d", expectedY, childLayout.Y)
	}
}

func TestLayoutEngine_BoxWithMargin_OffsetsBoxPosition(t *testing.T) {
	engine := NewLayoutEngine(80, 24)
	child := Text("Child")
	margin := Spacing{Top: 3, Right: 0, Bottom: 0, Left: 4}
	boxComponent := Box(BoxProps{Direction: Column, Margin: margin}, child)

	layoutTree := engine.CalculateLayout(boxComponent)

	expectedX := margin.Left
	if layoutTree.Layout.X != expectedX {
		t.Errorf("box X: expected %d (margin.Left), got %d", expectedX, layoutTree.Layout.X)
	}

	expectedY := margin.Top
	if layoutTree.Layout.Y != expectedY {
		t.Errorf("box Y: expected %d (margin.Top), got %d", expectedY, layoutTree.Layout.Y)
	}

	if len(layoutTree.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(layoutTree.Children))
	}

	childLayout := layoutTree.Children[0].Layout
	if childLayout.X != margin.Left {
		t.Errorf("child X: expected %d, got %d", margin.Left, childLayout.X)
	}
	if childLayout.Y != margin.Top {
		t.Errorf("child Y: expected %d, got %d", margin.Top, childLayout.Y)
	}
}

func TestLayoutEngine_BoxWithBorder_AdjustsChildPosition(t *testing.T) {
	engine := NewLayoutEngine(80, 24)
	child := Text("Child")
	boxComponent := Box(BoxProps{Direction: Column, Border: BorderSingle}, child)

	layoutTree := engine.CalculateLayout(boxComponent)

	if len(layoutTree.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(layoutTree.Children))
	}

	childLayout := layoutTree.Children[0].Layout

	expectedX := 1
	if childLayout.X != expectedX {
		t.Errorf("child X: expected %d (border offset), got %d", expectedX, childLayout.X)
	}

	expectedY := 1
	if childLayout.Y != expectedY {
		t.Errorf("child Y: expected %d (border offset), got %d", expectedY, childLayout.Y)
	}
}

func TestLayoutEngine_BoxWithGap_AddsSpaceBetweenChildren(t *testing.T) {
	engine := NewLayoutEngine(80, 24)
	child1 := Text("First")
	child2 := Text("Second")
	gap := 5
	boxComponent := Box(BoxProps{Direction: Column, Gap: gap}, child1, child2)

	layoutTree := engine.CalculateLayout(boxComponent)

	if len(layoutTree.Children) != 2 {
		t.Fatalf("expected 2 children, got %d", len(layoutTree.Children))
	}

	firstChild := layoutTree.Children[0]
	secondChild := layoutTree.Children[1]

	expectedSecondY := firstChild.Layout.Height + gap
	if secondChild.Layout.Y != expectedSecondY {
		t.Errorf("second child Y: expected %d (first height + gap), got %d", expectedSecondY, secondChild.Layout.Y)
	}
}

func TestLayoutEngine_BoxWithGapRow_AddsSpaceBetweenChildren(t *testing.T) {
	engine := NewLayoutEngine(80, 24)
	child1 := Text("First")
	child2 := Text("Second")
	gap := 3
	boxComponent := Box(BoxProps{Direction: Row, Gap: gap}, child1, child2)

	layoutTree := engine.CalculateLayout(boxComponent)

	if len(layoutTree.Children) != 2 {
		t.Fatalf("expected 2 children, got %d", len(layoutTree.Children))
	}

	firstChild := layoutTree.Children[0]
	secondChild := layoutTree.Children[1]

	expectedSecondX := firstChild.Layout.Width + gap
	if secondChild.Layout.X != expectedSecondX {
		t.Errorf("second child X: expected %d (first width + gap), got %d", expectedSecondX, secondChild.Layout.X)
	}
}
