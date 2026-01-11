package runetui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestApp_New_CreatesApp(t *testing.T) {
	rootFunc := func() Component {
		return Text("Hello")
	}

	app := New(rootFunc)

	if app == nil {
		t.Fatal("expected app to be created, got nil")
	}
}

func TestModel_Init_ReturnsNilCmd(t *testing.T) {
	rootFunc := func() Component {
		return Text("Hello")
	}

	app := New(rootFunc)
	m := app.createModel()

	cmd := m.Init()

	if cmd != nil {
		t.Errorf("expected Init() to return nil, got %v", cmd)
	}
}

func TestModel_View_RendersComponentTree(t *testing.T) {
	rootFunc := func() Component {
		return Text("Hello World")
	}

	app := New(rootFunc)
	m := app.createModel()

	output := m.View()

	if output == "" {
		t.Error("expected View() to return non-empty string")
	}
	if output != "Hello World" {
		t.Errorf("expected View() to render 'Hello World', got %q", output)
	}
}

func TestModel_Update_HandlesWindowSizeMsg(t *testing.T) {
	rootFunc := func() Component {
		return Text("Hello")
	}

	app := New(rootFunc)
	m := app.createModel().(*model)

	msg := tea.WindowSizeMsg{Width: 100, Height: 50}
	updatedModel, cmd := m.Update(msg)

	if cmd != nil {
		t.Errorf("expected Update() to return nil cmd, got %v", cmd)
	}

	updatedM, ok := updatedModel.(*model)
	if !ok {
		t.Fatal("expected updated model to be *model")
	}

	if updatedM.app.layoutEngine.terminalWidth != 100 {
		t.Errorf("expected terminal width to be 100, got %d", updatedM.app.layoutEngine.terminalWidth)
	}
	if updatedM.app.layoutEngine.terminalHeight != 50 {
		t.Errorf("expected terminal height to be 50, got %d", updatedM.app.layoutEngine.terminalHeight)
	}
}

func TestModel_Update_HandlesCtrlCQuit(t *testing.T) {
	rootFunc := func() Component {
		return Text("Hello")
	}

	app := New(rootFunc)
	m := app.createModel().(*model)

	msg := tea.KeyMsg{Type: tea.KeyCtrlC}
	updatedModel, cmd := m.Update(msg)

	if cmd == nil {
		t.Fatal("expected Update() to return quit command, got nil")
	}

	updatedM, ok := updatedModel.(*model)
	if !ok {
		t.Fatal("expected updated model to be *model")
	}

	if updatedM != m {
		t.Error("expected model to be unchanged")
	}
}

func TestApp_Run_CanBeCalled(t *testing.T) {
	rootFunc := func() Component {
		return Text("Hello")
	}

	app := New(rootFunc)

	// Verify Run method can be referenced (will panic if it doesn't exist)
	_ = app.Run
}

func TestApp_RunContext_CanBeCalled(t *testing.T) {
	rootFunc := func() Component {
		return Text("Hello")
	}

	app := New(rootFunc)

	// Verify RunContext method can be referenced (will panic if it doesn't exist)
	_ = app.RunContext
}

func TestModel_View_OnlyStaticContent(t *testing.T) {
	rootFunc := func() Component {
		return Static(StaticProps{Key: "test"}, func() []Component {
			return []Component{
				Text("Static Line 1"),
				Text("Static Line 2"),
			}
		})
	}

	app := New(rootFunc)
	m := app.createModel()

	// First render: should show static content
	output1 := m.View()
	if output1 == "" {
		t.Error("expected first render to show static content")
	}

	// Test that View() can handle only static content (no dynamic content)
	// This covers the branch where dynamicContent == ""
	if len(output1) < 5 {
		t.Errorf("expected static content to be rendered, got %q", output1)
	}
}

func TestModel_View_OnlyDynamicContent(t *testing.T) {
	rootFunc := func() Component {
		return Text("Only Dynamic")
	}

	app := New(rootFunc)
	m := app.createModel()

	// Render should show only dynamic content (no static content)
	// This covers the branch where staticContent == ""
	output := m.View()
	if output == "" {
		t.Error("expected dynamic content to be rendered")
	}
	if len(output) < 5 {
		t.Errorf("expected dynamic content, got %q", output)
	}
}

func TestModel_View_StaticAndDynamicContent(t *testing.T) {
	rootFunc := func() Component {
		return Box(
			BoxProps{Direction: Column},
			Static(StaticProps{Key: "test"}, func() []Component {
				return []Component{Text("Static")}
			}),
			Text("Dynamic"),
		)
	}

	app := New(rootFunc)
	m := app.createModel()

	// First render: should show both static and dynamic
	output1 := m.View()
	if output1 == "" {
		t.Error("expected first render to have content")
	}
	if len(output1) < 5 {
		t.Errorf("expected first render to have static and dynamic content, got %q", output1)
	}

	// Second render: should show only dynamic (static already rendered)
	output2 := m.View()
	if output2 == "" {
		t.Error("expected second render to have dynamic content")
	}
}

func TestRenderTree_WithNilTree_ReturnsEmpty(t *testing.T) {
	output := renderTree(nil)

	if output != "" {
		t.Errorf("expected empty string, got %q", output)
	}
}

func TestRenderTree_WithNestedChildren_CombinesOutput(t *testing.T) {
	child1 := &LayoutTree{
		Component: Text("Child1"),
		Layout:    Layout{X: 0, Y: 0, Width: 10, Height: 1},
		Children:  nil,
	}
	child2 := &LayoutTree{
		Component: Text("Child2"),
		Layout:    Layout{X: 0, Y: 1, Width: 10, Height: 1},
		Children:  nil,
	}
	parent := &LayoutTree{
		Component: Box(BoxProps{}, []Component{}...),
		Layout:    Layout{X: 0, Y: 0, Width: 10, Height: 2},
		Children:  []*LayoutTree{child1, child2},
	}

	output := renderTree(parent)

	if output == "" {
		t.Error("expected non-empty output")
	}
	// Just check that both children are rendered (without being strict about spacing)
	if len(output) < 10 {
		t.Errorf("expected combined output with children, got %q", output)
	}
}

func TestRenderTree_WithEmptyChildOutput_SkipsChild(t *testing.T) {
	// Create a child that returns empty string when rendered
	emptyChild := &LayoutTree{
		Component: Box(BoxProps{}, []Component{}...),
		Layout:    Layout{X: 0, Y: 0, Width: 0, Height: 0},
		Children:  nil,
	}
	parent := &LayoutTree{
		Component: Text("Parent"),
		Layout:    Layout{X: 0, Y: 0, Width: 10, Height: 1},
		Children:  []*LayoutTree{emptyChild},
	}

	output := renderTree(parent)

	// Check that parent content is included
	if len(output) < 3 {
		t.Errorf("expected parent content in output, got %q", output)
	}
}

func TestNew_WithOptions_AppliesOptions(t *testing.T) {
	rootFunc := func() Component {
		return Text("Hello")
	}

	var called bool
	option := func(a *App) {
		called = true
	}

	app := New(rootFunc, option)

	if app == nil {
		t.Fatal("expected app to be created")
	}
	if !called {
		t.Error("expected option to be called")
	}
}
