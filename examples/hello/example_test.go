package main

import (
	"strings"
	"testing"

	"github.com/runetui/runetui"
	runetuitester "github.com/runetui/runetui/testing"
)

// helloComponent extracts the component function from the example
// so it can be tested independently of the Bubble Tea runtime.
func helloComponent() runetui.Component {
	return runetui.Box(
		runetui.BoxProps{
			Direction: runetui.Column,
			Padding:   runetui.SpacingAll(2),
			Border:    runetui.BorderSingle,
		},
		runetui.Text("Hello, RuneTUI!", runetui.TextProps{Bold: true}),
		runetui.Text("Press Ctrl+C to quit"),
	)
}

// TestHelloExample_RendersCorrectly verifies that the hello example
// renders correctly by comparing against a snapshot.
func TestHelloExample_RendersCorrectly(t *testing.T) {
	output := runetuitester.RenderToString(helloComponent, 80, 24)
	runetuitester.AssertSnapshot(t, "hello_snapshot", output)
}

// TestHelloExample_WithDifferentSizes verifies the example renders
// correctly at different terminal sizes.
func TestHelloExample_WithDifferentSizes(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
	}{
		{"SmallTerminal", 40, 10},
		{"StandardTerminal", 80, 24},
		{"LargeTerminal", 120, 40},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := runetuitester.RenderToString(helloComponent, tt.width, tt.height)
			// Verify output is not empty
			if output == "" {
				t.Error("expected non-empty output")
			}
			// Verify it contains expected text
			if !strings.Contains(output, "Hello, RuneTUI!") {
				t.Errorf("output should contain 'Hello, RuneTUI!'")
			}
		})
	}
}

// TestHelloExample_ResizeHandling verifies that the example handles
// terminal resizing correctly.
func TestHelloExample_ResizeHandling(t *testing.T) {
	app := runetuitester.NewTestApp(helloComponent)

	// Initial render
	output1 := app.View()
	if output1 == "" {
		t.Error("expected non-empty output")
	}

	// Resize and render again
	app.Resize(100, 50)
	output2 := app.View()
	if output2 == "" {
		t.Error("expected non-empty output after resize")
	}

	// Both outputs should contain the expected text
	if !strings.Contains(output1, "Hello, RuneTUI!") {
		t.Error("initial output should contain 'Hello, RuneTUI!'")
	}
	if !strings.Contains(output2, "Hello, RuneTUI!") {
		t.Error("resized output should contain 'Hello, RuneTUI!'")
	}
}
