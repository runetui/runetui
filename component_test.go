package runetui

import "testing"

func TestLayout_WithValues_StoresAllFields(t *testing.T) {
	layout := Layout{X: 10, Y: 20, Width: 100, Height: 50}

	if layout.X != 10 {
		t.Errorf("expected X=10, got %d", layout.X)
	}
	if layout.Y != 20 {
		t.Errorf("expected Y=20, got %d", layout.Y)
	}
	if layout.Width != 100 {
		t.Errorf("expected Width=100, got %d", layout.Width)
	}
	if layout.Height != 50 {
		t.Errorf("expected Height=50, got %d", layout.Height)
	}
}

func TestLayout_ZeroValue_CreatesZeroLayout(t *testing.T) {
	layout := Layout{}

	if layout.X != 0 {
		t.Errorf("expected X=0, got %d", layout.X)
	}
	if layout.Y != 0 {
		t.Errorf("expected Y=0, got %d", layout.Y)
	}
	if layout.Width != 0 {
		t.Errorf("expected Width=0, got %d", layout.Width)
	}
	if layout.Height != 0 {
		t.Errorf("expected Height=0, got %d", layout.Height)
	}
}

func TestLayout_WithNegativeValues_AllowsNegative(t *testing.T) {
	layout := Layout{X: -5, Y: -10, Width: -20, Height: -30}

	if layout.X != -5 {
		t.Errorf("expected X=-5, got %d", layout.X)
	}
	if layout.Y != -10 {
		t.Errorf("expected Y=-10, got %d", layout.Y)
	}
	if layout.Width != -20 {
		t.Errorf("expected Width=-20, got %d", layout.Width)
	}
	if layout.Height != -30 {
		t.Errorf("expected Height=-30, got %d", layout.Height)
	}
}

func TestSize_WithValues_StoresAllFields(t *testing.T) {
	size := Size{
		Width:     100,
		Height:    50,
		MinWidth:  10,
		MinHeight: 5,
		MaxWidth:  200,
		MaxHeight: 100,
	}

	if size.Width != 100 {
		t.Errorf("expected Width=100, got %d", size.Width)
	}
	if size.Height != 50 {
		t.Errorf("expected Height=50, got %d", size.Height)
	}
	if size.MinWidth != 10 {
		t.Errorf("expected MinWidth=10, got %d", size.MinWidth)
	}
	if size.MinHeight != 5 {
		t.Errorf("expected MinHeight=5, got %d", size.MinHeight)
	}
	if size.MaxWidth != 200 {
		t.Errorf("expected MaxWidth=200, got %d", size.MaxWidth)
	}
	if size.MaxHeight != 100 {
		t.Errorf("expected MaxHeight=100, got %d", size.MaxHeight)
	}
}

func TestSize_ZeroValue_CreatesZeroSize(t *testing.T) {
	size := Size{}

	if size.Width != 0 {
		t.Errorf("expected Width=0, got %d", size.Width)
	}
	if size.Height != 0 {
		t.Errorf("expected Height=0, got %d", size.Height)
	}
	if size.MinWidth != 0 {
		t.Errorf("expected MinWidth=0, got %d", size.MinWidth)
	}
	if size.MinHeight != 0 {
		t.Errorf("expected MinHeight=0, got %d", size.MinHeight)
	}
	if size.MaxWidth != 0 {
		t.Errorf("expected MaxWidth=0, got %d", size.MaxWidth)
	}
	if size.MaxHeight != 0 {
		t.Errorf("expected MaxHeight=0, got %d", size.MaxHeight)
	}
}

func TestSize_WithNegativeValues_AllowsNegative(t *testing.T) {
	size := Size{
		Width:     -100,
		Height:    -50,
		MinWidth:  -10,
		MinHeight: -5,
		MaxWidth:  -200,
		MaxHeight: -100,
	}

	if size.Width != -100 {
		t.Errorf("expected Width=-100, got %d", size.Width)
	}
	if size.Height != -50 {
		t.Errorf("expected Height=-50, got %d", size.Height)
	}
	if size.MinWidth != -10 {
		t.Errorf("expected MinWidth=-10, got %d", size.MinWidth)
	}
	if size.MinHeight != -5 {
		t.Errorf("expected MinHeight=-5, got %d", size.MinHeight)
	}
	if size.MaxWidth != -200 {
		t.Errorf("expected MaxWidth=-200, got %d", size.MaxWidth)
	}
	if size.MaxHeight != -100 {
		t.Errorf("expected MaxHeight=-100, got %d", size.MaxHeight)
	}
}

type testProps struct {
	value string
}

func (testProps) isProps() {}

func TestProps_ConcreteType_ImplementsInterface(t *testing.T) {
	var _ Props = testProps{}
}

func TestProps_ConcreteType_StoresData(t *testing.T) {
	props := testProps{value: "test"}
	if props.value != "test" {
		t.Errorf("expected value=test, got %s", props.value)
	}
}

type testComponent struct {
	key      string
	children []Component
}

func (c testComponent) Render(layout Layout) string {
	return "rendered"
}

func (c testComponent) Children() []Component {
	return c.children
}

func (c testComponent) Key() string {
	return c.key
}

func (c testComponent) Measure(availableWidth, availableHeight int) Size {
	return Size{Width: availableWidth, Height: availableHeight}
}

func TestComponent_ConcreteType_ImplementsInterface(t *testing.T) {
	var _ Component = testComponent{}
}

func TestComponent_Render_ReturnsString(t *testing.T) {
	comp := testComponent{key: "test"}
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 5}
	result := comp.Render(layout)

	if result != "rendered" {
		t.Errorf("expected 'rendered', got %s", result)
	}
}

func TestComponent_Key_ReturnsKey(t *testing.T) {
	comp := testComponent{key: "my-key"}
	if comp.Key() != "my-key" {
		t.Errorf("expected 'my-key', got %s", comp.Key())
	}
}

func TestComponent_Children_ReturnsNilForNoChildren(t *testing.T) {
	comp := testComponent{key: "test"}
	children := comp.Children()

	if children != nil {
		t.Errorf("expected nil slice, got %v", children)
	}
}

func TestComponent_Children_ReturnsChildren(t *testing.T) {
	child1 := testComponent{key: "child1"}
	child2 := testComponent{key: "child2"}
	comp := testComponent{key: "parent", children: []Component{child1, child2}}

	children := comp.Children()
	if len(children) != 2 {
		t.Errorf("expected 2 children, got %d", len(children))
	}
	if children[0].Key() != "child1" {
		t.Errorf("expected first child key 'child1', got %s", children[0].Key())
	}
	if children[1].Key() != "child2" {
		t.Errorf("expected second child key 'child2', got %s", children[1].Key())
	}
}

func TestComponent_Measure_ReturnsSize(t *testing.T) {
	comp := testComponent{key: "test"}
	size := comp.Measure(100, 50)

	if size.Width != 100 {
		t.Errorf("expected Width=100, got %d", size.Width)
	}
	if size.Height != 50 {
		t.Errorf("expected Height=50, got %d", size.Height)
	}
}

func TestComponentFunc_ImplementsComponent(t *testing.T) {
	var _ Component = ComponentFunc(nil)
}

func TestComponentFunc_Render_DelegatesToFunction(t *testing.T) {
	inner := testComponent{key: "inner"}
	fn := ComponentFunc(func() Component {
		return inner
	})

	layout := Layout{X: 0, Y: 0, Width: 10, Height: 5}
	result := fn.Render(layout)

	if result != "rendered" {
		t.Errorf("expected 'rendered', got %s", result)
	}
}

func TestComponentFunc_Key_DelegatesToFunction(t *testing.T) {
	inner := testComponent{key: "func-key"}
	fn := ComponentFunc(func() Component {
		return inner
	})

	if fn.Key() != "func-key" {
		t.Errorf("expected 'func-key', got %s", fn.Key())
	}
}

func TestComponentFunc_Children_DelegatesToFunction(t *testing.T) {
	child := testComponent{key: "child"}
	inner := testComponent{key: "inner", children: []Component{child}}
	fn := ComponentFunc(func() Component {
		return inner
	})

	children := fn.Children()
	if len(children) != 1 {
		t.Errorf("expected 1 child, got %d", len(children))
	}
	if children[0].Key() != "child" {
		t.Errorf("expected child key 'child', got %s", children[0].Key())
	}
}

func TestComponentFunc_Measure_DelegatesToFunction(t *testing.T) {
	inner := testComponent{key: "inner"}
	fn := ComponentFunc(func() Component {
		return inner
	})

	size := fn.Measure(80, 40)
	if size.Width != 80 {
		t.Errorf("expected Width=80, got %d", size.Width)
	}
	if size.Height != 40 {
		t.Errorf("expected Height=40, got %d", size.Height)
	}
}
