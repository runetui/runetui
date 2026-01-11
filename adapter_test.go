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
