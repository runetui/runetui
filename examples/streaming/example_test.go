package main

import (
	"testing"

	"github.com/runetui/runetui"
	runtesting "github.com/runetui/runetui/testing"
)

// TestStreamingExample_RendersCorrectly verifies that the streaming example
// renders without errors and produces expected output.
func TestStreamingExample_RendersCorrectly(t *testing.T) {
	// Create model with initial state
	m := model{
		logs: []string{
			"[12:00:00] Application started",
			"[12:00:00] Initializing components...",
			"[12:00:00] Ready!",
		},
		status:        "Running... (3 entries)",
		ticks:         3,
		staticManager: runetui.NewStaticManager(),
		layoutEngine:  runetui.NewLayoutEngine(80, 24),
	}

	// Set static manager for rendering
	runetui.SetStaticManager(m.staticManager)
	defer runetui.SetStaticManager(nil)

	// Build component tree
	rootFunc := func() runetui.Component {
		return m.buildComponentTree()
	}

	// Render to string
	output := runtesting.RenderToString(rootFunc, 80, 24)

	// Verify output is not empty
	if output == "" {
		t.Error("expected non-empty output")
	}

	// Verify output contains key elements
	if len(output) < 100 {
		t.Errorf("expected substantial output, got %d characters", len(output))
	}

	// Snapshot test
	runtesting.AssertSnapshot(t, "streaming_initial", output)
}

// TestStreamingExample_WithMultipleLogs verifies rendering with more log entries.
func TestStreamingExample_WithMultipleLogs(t *testing.T) {
	// Create model with more logs
	m := model{
		logs: []string{
			"[12:00:00] Application started",
			"[12:00:01] Processing item 1",
			"[12:00:02] Processing item 2",
			"[12:00:03] Processing item 3",
			"[12:00:04] Processing item 4",
			"[12:00:05] All items processed",
		},
		status:        "Complete! Press Ctrl+C to quit",
		ticks:         20,
		staticManager: runetui.NewStaticManager(),
		layoutEngine:  runetui.NewLayoutEngine(80, 24),
	}

	// Set static manager for rendering
	runetui.SetStaticManager(m.staticManager)
	defer runetui.SetStaticManager(nil)

	// Build component tree
	rootFunc := func() runetui.Component {
		return m.buildComponentTree()
	}

	// Render to string
	output := runtesting.RenderToString(rootFunc, 80, 24)

	// Verify output contains all log entries
	if output == "" {
		t.Error("expected non-empty output")
	}

	// Snapshot test
	runtesting.AssertSnapshot(t, "streaming_multiple_logs", output)
}

// TestStreamingExample_ResizeHandling verifies rendering at different sizes.
func TestStreamingExample_ResizeHandling(t *testing.T) {
	testCases := []struct {
		name   string
		width  int
		height int
	}{
		{"SmallTerminal", 40, 10},
		{"StandardTerminal", 80, 24},
		{"LargeTerminal", 120, 40},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create model with initial state
			m := model{
				logs: []string{
					"[12:00:00] Application started",
					"[12:00:00] Ready!",
				},
				status:        "Running...",
				ticks:         2,
				staticManager: runetui.NewStaticManager(),
				layoutEngine:  runetui.NewLayoutEngine(tc.width, tc.height),
			}

			// Set static manager for rendering
			runetui.SetStaticManager(m.staticManager)
			defer runetui.SetStaticManager(nil)

			// Build component tree
			rootFunc := func() runetui.Component {
				return m.buildComponentTree()
			}

			// Render to string with specified dimensions
			output := runtesting.RenderToString(rootFunc, tc.width, tc.height)

			// Verify output is generated
			if output == "" {
				t.Errorf("%s: expected non-empty output", tc.name)
			}
		})
	}
}

// TestStreamingExample_StaticBehavior verifies static zone behavior.
func TestStreamingExample_StaticBehavior(t *testing.T) {
	// Create model with initial logs
	m := model{
		logs: []string{
			"[12:00:00] Log 1",
			"[12:00:01] Log 2",
		},
		status:        "Running...",
		ticks:         2,
		staticManager: runetui.NewStaticManager(),
		layoutEngine:  runetui.NewLayoutEngine(80, 24),
	}

	// Set static manager
	runetui.SetStaticManager(m.staticManager)
	defer runetui.SetStaticManager(nil)

	// First render
	rootFunc1 := func() runetui.Component {
		return m.buildComponentTree()
	}
	output1 := runtesting.RenderToString(rootFunc1, 80, 24)

	// Add more logs
	m.logs = append(m.logs, "[12:00:02] Log 3")
	m.ticks = 3

	// Second render - static content should accumulate
	rootFunc2 := func() runetui.Component {
		return m.buildComponentTree()
	}
	output2 := runtesting.RenderToString(rootFunc2, 80, 24)

	// Both renders should produce output
	if output1 == "" || output2 == "" {
		t.Error("expected non-empty output for both renders")
	}

	// Output should change when new logs are added
	if output1 == output2 {
		t.Error("expected output to differ when logs are added")
	}
}
