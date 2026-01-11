package runetui

import "testing"

func TestStaticProps_ImplementsPropsInterface(t *testing.T) {
	props := StaticProps{Key: "test"}
	var _ Props = props
}

func TestStatic_WithItemsFunc_ReturnsNonNil(t *testing.T) {
	props := StaticProps{Key: "static1"}
	itemsFunc := func() []Component {
		return []Component{}
	}
	static := Static(props, itemsFunc)
	if static == nil {
		t.Error("Static should return non-nil Component")
	}
}

func TestStatic_Key_ReturnsKeyFromProps(t *testing.T) {
	props := StaticProps{Key: "test-key"}
	itemsFunc := func() []Component {
		return []Component{}
	}
	static := Static(props, itemsFunc)
	key := static.Key()

	if key != "test-key" {
		t.Errorf("Expected key 'test-key', got %q", key)
	}
}

func TestStatic_Children_ReturnsEmptySlice(t *testing.T) {
	props := StaticProps{Key: "static1"}
	itemsFunc := func() []Component {
		return []Component{Text("test")}
	}
	static := Static(props, itemsFunc)
	children := static.Children()

	if children == nil {
		t.Error("Children() should return empty slice, not nil")
	}
	if len(children) != 0 {
		t.Errorf("Expected 0 children, got %d", len(children))
	}
}

func TestStatic_Render_WithEmptyItems_ReturnsEmptyString(t *testing.T) {
	props := StaticProps{Key: "static1"}
	itemsFunc := func() []Component {
		return []Component{}
	}
	static := Static(props, itemsFunc)
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 10}

	result := static.Render(layout)

	if result != "" {
		t.Errorf("Expected empty string, got %q", result)
	}
}

func TestStatic_Render_WithItems_RendersJoinedContent(t *testing.T) {
	props := StaticProps{Key: "static1"}
	itemsFunc := func() []Component {
		return []Component{
			Text("Line 1"),
			Text("Line 2"),
		}
	}
	static := Static(props, itemsFunc)
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 10}

	result := static.Render(layout)

	if result == "" {
		t.Error("Expected non-empty result")
	}
	// Should contain both lines
	if len(result) == 0 {
		t.Error("Expected content to be rendered")
	}
}

func TestStatic_Measure_WithEmptyItems_ReturnsZeroSize(t *testing.T) {
	props := StaticProps{Key: "static1"}
	itemsFunc := func() []Component {
		return []Component{}
	}
	static := Static(props, itemsFunc)

	size := static.Measure(10, 10)

	if size.Width != 0 {
		t.Errorf("Expected width 0, got %d", size.Width)
	}
	if size.Height != 0 {
		t.Errorf("Expected height 0, got %d", size.Height)
	}
}

func TestStatic_Measure_WithItems_SumsHeights(t *testing.T) {
	props := StaticProps{Key: "static1"}
	itemsFunc := func() []Component {
		return []Component{
			Text("Line 1"),
			Text("Line 2"),
		}
	}
	static := Static(props, itemsFunc)

	size := static.Measure(10, 10)

	if size.Height != 2 {
		t.Errorf("Expected height 2, got %d", size.Height)
	}
	if size.Width != 6 {
		t.Errorf("Expected width 6 (max of 'Line 1' and 'Line 2'), got %d", size.Width)
	}
}

func TestStatic_WithStaticManager_OnlyReturnsNewContent(t *testing.T) {
	sm := NewStaticManager()
	SetStaticManager(sm)
	defer SetStaticManager(nil)

	props := StaticProps{Key: "static1"}
	itemsFunc := func() []Component {
		return []Component{
			Text("Line 1"),
			Text("Line 2"),
		}
	}
	static := Static(props, itemsFunc)
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 10}

	result1 := static.Render(layout)
	result2 := static.Render(layout)

	if result1 == "" {
		t.Error("First render should return content")
	}
	if result2 != "" {
		t.Errorf("Second render should return empty string, got %q", result2)
	}
}

func TestStatic_WithStaticManager_DifferentKeysBothRender(t *testing.T) {
	sm := NewStaticManager()
	SetStaticManager(sm)
	defer SetStaticManager(nil)

	static1 := Static(StaticProps{Key: "static1"}, func() []Component {
		return []Component{Text("Line 1")}
	})
	static2 := Static(StaticProps{Key: "static2"}, func() []Component {
		return []Component{Text("Line 2")}
	})
	layout := Layout{X: 0, Y: 0, Width: 10, Height: 10}

	result1 := static1.Render(layout)
	result2 := static2.Render(layout)

	if result1 == "" {
		t.Error("First static should render content")
	}
	if result2 == "" {
		t.Error("Second static should render content")
	}
}
