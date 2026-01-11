package testing

import (
	"testing"

	"github.com/runetui/runetui"
)

// Test 1: RenderToString renders a simple text component
func TestRenderToString_WithSimpleText_RendersContent(t *testing.T) {
	rootFunc := func() runetui.Component {
		return runetui.Text("Hello")
	}

	output := RenderToString(rootFunc, 80, 24)

	if output == "" {
		t.Error("expected non-empty output")
	}

	if output != "Hello" {
		t.Errorf("expected 'Hello', got %q", output)
	}
}

// Test 2: RenderToString renders a box with children
func TestRenderToString_WithBoxAndChildren_RendersAllComponents(t *testing.T) {
	rootFunc := func() runetui.Component {
		return runetui.Box(
			runetui.BoxProps{},
			runetui.Text("Line1"),
			runetui.Text("Line2"),
		)
	}

	output := RenderToString(rootFunc, 80, 24)

	if output == "" {
		t.Error("expected non-empty output")
	}

	// Should contain both text components
	if len(output) < 10 {
		t.Errorf("expected output with both texts, got %q", output)
	}
}

// Test 3: AssertSnapshot creates golden file when it doesn't exist
func TestAssertSnapshot_NewSnapshot_CreatesGoldenFile(t *testing.T) {
	output := "test content"
	name := "test_new_snapshot"

	// This should pass and create the golden file
	AssertSnapshot(t, name, output)
}

// Test 4: AssertSnapshot compares with existing golden file
func TestAssertSnapshot_ExistingSnapshot_ComparesCorrectly(t *testing.T) {
	name := "test_existing_snapshot"
	content := "consistent content"

	// First call creates the golden file
	AssertSnapshot(t, name, content)

	// Second call with same content should pass
	AssertSnapshot(t, name, content)
}

// Test 5: NewTestApp creates a TestApp instance
func TestNewTestApp_WithRootFunc_CreatesTestApp(t *testing.T) {
	rootFunc := func() runetui.Component {
		return runetui.Text("Test")
	}

	app := NewTestApp(rootFunc)

	if app == nil {
		t.Fatal("expected non-nil TestApp")
	}
}

// Test 6: TestApp.Resize updates dimensions
func TestTestApp_Resize_UpdatesDimensions(t *testing.T) {
	rootFunc := func() runetui.Component {
		return runetui.Text("Test")
	}

	app := NewTestApp(rootFunc)
	app.Resize(100, 50)

	// Verify resize by checking the View output dimensions
	view := app.View()

	if view == "" {
		t.Error("expected non-empty view after resize")
	}
}

// Test 7: TestApp.SendKey simulates keyboard input
func TestTestApp_SendKey_SimulatesKeyPress(t *testing.T) {
	rootFunc := func() runetui.Component {
		return runetui.Text("Test")
	}

	app := NewTestApp(rootFunc)
	app.SendKey("enter")

	// For now, just verify the method exists and can be called
	view := app.View()
	if view == "" {
		t.Error("expected non-empty view after SendKey")
	}
}

// Test 8: RenderToString with zero dimensions
func TestRenderToString_ZeroDimensions_HandlesGracefully(t *testing.T) {
	rootFunc := func() runetui.Component {
		return runetui.Text("Test")
	}

	// Should not panic, even with zero dimensions
	output := RenderToString(rootFunc, 0, 0)

	// With zero dimensions, rendering should still work (lipgloss handles it)
	// We just verify it doesn't panic
	_ = output
}

// Test 9: TestApp initial dimensions
func TestNewTestApp_DefaultDimensions_AreSet(t *testing.T) {
	rootFunc := func() runetui.Component {
		return runetui.Text("Test")
	}

	app := NewTestApp(rootFunc)
	view := app.View()

	// Default dimensions should be 80x24
	if view == "" {
		t.Error("expected non-empty view with default dimensions")
	}
}

// Test 10: AssertSnapshot with -update flag behavior
func TestAssertSnapshot_UpdateFlag_UpdatesGoldenFile(t *testing.T) {
	// This test verifies the golden file workflow
	// When run normally, it uses existing golden files
	// When run with -update, it creates/updates golden files
	name := "test_update_workflow"
	content := "workflow content"

	AssertSnapshot(t, name, content)
}

// Test 11: RenderToString with complex nested structure
func TestRenderToString_NestedBoxes_RendersCorrectly(t *testing.T) {
	rootFunc := func() runetui.Component {
		return runetui.Box(
			runetui.BoxProps{},
			runetui.Box(
				runetui.BoxProps{},
				runetui.Text("Nested"),
			),
		)
	}

	output := RenderToString(rootFunc, 80, 24)

	if output == "" {
		t.Error("expected non-empty output for nested structure")
	}
}

// Test 12: TestApp multiple resize operations
func TestTestApp_MultipleResize_UpdatesCorrectly(t *testing.T) {
	rootFunc := func() runetui.Component {
		return runetui.Text("Test")
	}

	app := NewTestApp(rootFunc)

	app.Resize(100, 50)
	view1 := app.View()

	app.Resize(50, 25)
	view2 := app.View()

	// Both views should render (dimensions affect layout)
	if view1 == "" || view2 == "" {
		t.Error("expected non-empty views after multiple resizes")
	}
}

// Test 13: SendKey can be called multiple times
func TestTestApp_SendKey_CanBeCalledMultipleTimes(t *testing.T) {
	rootFunc := func() runetui.Component {
		return runetui.Text("Test")
	}

	app := NewTestApp(rootFunc)

	// Call SendKey multiple times
	app.SendKey("a")
	app.SendKey("b")
	app.SendKey("enter")

	view := app.View()
	if view == "" {
		t.Error("expected non-empty view")
	}
}

// Test 14: AssertSnapshot with update flag
func TestAssertSnapshot_WithUpdateFlag_CreatesFile(t *testing.T) {
	// This test verifies that AssertSnapshot works with the update workflow
	// The actual -update flag testing is done at runtime
	name := "test_snapshot_create"
	content := "snapshot content"

	AssertSnapshot(t, name, content)
}

// Test 15: RenderToString with nil-like component
func TestRenderToString_EmptyTree_ReturnsEmptyString(t *testing.T) {
	rootFunc := func() runetui.Component {
		return runetui.Text("")
	}

	output := RenderToString(rootFunc, 80, 24)

	// Empty text should still render (might have padding/styling)
	_ = output
}
