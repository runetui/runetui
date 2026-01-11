// Package testing provides utilities for testing RuneTUI components.
//
// This package offers tools to render components without a terminal,
// perform snapshot testing with golden files, and simulate user interactions.
//
// Example usage:
//
//	func TestMyComponent(t *testing.T) {
//	    rootFunc := func() runetui.Component {
//	        return runetui.Text("Hello, World!")
//	    }
//
//	    // Render to string for assertions
//	    output := testing.RenderToString(rootFunc, 80, 24)
//	    if !strings.Contains(output, "Hello") {
//	        t.Error("expected output to contain 'Hello'")
//	    }
//
//	    // Or use snapshot testing
//	    testing.AssertSnapshot(t, "my_component", output)
//	}
//
//	func TestInteractiveComponent(t *testing.T) {
//	    rootFunc := func() runetui.Component {
//	        return runetui.Box(runetui.BoxProps{}, runetui.Text("Test"))
//	    }
//
//	    app := testing.NewTestApp(rootFunc)
//	    app.Resize(100, 50)
//	    app.SendKey("enter")
//	    view := app.View()
//
//	    testing.AssertSnapshot(t, "interactive_view", view)
//	}
package testing

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/runetui/runetui"
)

var updateGolden = flag.Bool("update", false, "update golden files")

// RenderToString renders a component tree to a string without starting a terminal.
// This is useful for testing components in non-interactive environments.
//
// The width and height parameters define the available space for rendering,
// similar to terminal dimensions.
//
// Example:
//
//	rootFunc := func() runetui.Component {
//	    return runetui.Text("Hello, World!")
//	}
//	output := testing.RenderToString(rootFunc, 80, 24)
//	fmt.Println(output) // "Hello, World!"
func RenderToString(rootFunc func() runetui.Component, width, height int) string {
	engine := runetui.NewLayoutEngine(width, height)
	root := rootFunc()
	tree := engine.CalculateLayout(root)
	return renderTree(tree)
}

// renderTree recursively renders a layout tree to a string.
func renderTree(tree *runetui.LayoutTree) string {
	if tree == nil {
		return ""
	}

	rendered := tree.Component.Render(tree.Layout)

	for _, child := range tree.Children {
		childOutput := renderTree(child)
		if childOutput != "" {
			rendered += childOutput
		}
	}

	return rendered
}

// AssertSnapshot compares the output string against a golden file.
// If the golden file doesn't exist, it creates a new golden file with the output.
// If the -update flag is set, it updates existing golden files with the new output.
// If the content differs from the golden file, the test fails with a diff.
//
// Golden files are stored in testdata/<name>.golden relative to the test file.
//
// Example:
//
//	output := testing.RenderToString(rootFunc, 80, 24)
//	testing.AssertSnapshot(t, "my_component_output", output)
//
// To update golden files when the output intentionally changes:
//
//	go test -update
func AssertSnapshot(t testing.TB, name string, output string) {
	t.Helper()

	goldenFile := filepath.Join("testdata", name+".golden")

	if *updateGolden {
		writeGoldenFile(t, goldenFile, output)
		return
	}

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		if os.IsNotExist(err) {
			writeGoldenFile(t, goldenFile, output)
			return
		}
		t.Fatalf("failed to read golden file: %v", err)
	}

	if string(expected) != output {
		t.Errorf("snapshot mismatch for %s:\nexpected:\n%s\n\ngot:\n%s\n\nrun with -update to update golden files", name, expected, output)
	}
}

func writeGoldenFile(t testing.TB, path string, content string) {
	if err := os.MkdirAll("testdata", 0755); err != nil {
		t.Fatalf("failed to create testdata directory: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write golden file: %v", err)
	}
}

// TestApp is a test wrapper that allows simulating user interactions
// with RuneTUI components without starting a terminal.
//
// Example:
//
//	rootFunc := func() runetui.Component {
//	    return runetui.Box(runetui.BoxProps{}, runetui.Text("Hello"))
//	}
//	app := testing.NewTestApp(rootFunc)
//	app.Resize(100, 50)
//	app.SendKey("enter")
//	view := app.View()
//	fmt.Println(view)
type TestApp struct {
	rootFunc func() runetui.Component
	width    int
	height   int
}

// NewTestApp creates a new TestApp for testing components.
// The default dimensions are 80x24 (standard terminal size).
func NewTestApp(rootFunc func() runetui.Component) *TestApp {
	return &TestApp{
		rootFunc: rootFunc,
		width:    80,
		height:   24,
	}
}

// Resize simulates a terminal resize event.
func (a *TestApp) Resize(width, height int) {
	a.width = width
	a.height = height
}

// View returns the current rendered view of the component tree.
func (a *TestApp) View() string {
	return RenderToString(a.rootFunc, a.width, a.height)
}

// SendKey simulates a keyboard input event.
// Note: This is a placeholder for future stateful component support.
// Currently, RuneTUI components are stateless, so this method stores the key
// but doesn't trigger any updates. When state management is added, this will
// trigger component updates and re-renders.
func (a *TestApp) SendKey(key string) {
	// Placeholder for future state management
	// Will be implemented when components support state
}
