package main

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/runetui/runetui"
	runtesting "github.com/runetui/runetui/testing"
)

// TestStreamingExample_RendersCorrectly verifies that the streaming example
// renders without errors and produces expected output.
func TestStreamingExample_RendersCorrectly(t *testing.T) {
	state := &streamState{
		logs: []string{
			"[12:00:00] Application started",
			"[12:00:00] Initializing components...",
			"[12:00:00] Ready!",
		},
		status: "Running... (3 entries)",
		ticks:  3,
	}

	rootFunc, _ := createStreamingApp(state)
	output := runtesting.RenderToString(rootFunc, 80, 24)

	runetui.AssertNotEmpty(t, output)
	runetui.AssertContainsText(t, output, "Streaming Logs Example")
	runetui.AssertContainsText(t, output, "Running...")

	runtesting.AssertSnapshot(t, "streaming_initial", output)
}

// TestStreamingExample_WithMultipleLogs verifies rendering with more log entries.
func TestStreamingExample_WithMultipleLogs(t *testing.T) {
	state := &streamState{
		logs: []string{
			"[12:00:00] Application started",
			"[12:00:01] Processing item 1",
			"[12:00:02] Processing item 2",
			"[12:00:03] Processing item 3",
			"[12:00:04] Processing item 4",
			"[12:00:05] All items processed",
		},
		status: "Complete! Press q to quit",
		ticks:  20,
	}

	rootFunc, _ := createStreamingApp(state)
	output := runtesting.RenderToString(rootFunc, 80, 24)

	runetui.AssertNotEmpty(t, output)
	runetui.AssertContainsText(t, output, "Complete!")

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
			state := &streamState{
				logs: []string{
					"[12:00:00] Application started",
					"[12:00:00] Ready!",
				},
				status: "Running...",
				ticks:  2,
			}

			rootFunc, _ := createStreamingApp(state)
			output := runtesting.RenderToString(rootFunc, tc.width, tc.height)

			runetui.AssertNotEmpty(t, output)
		})
	}
}

// TestStreamingExample_StaticBehavior verifies static zone behavior.
func TestStreamingExample_StaticBehavior(t *testing.T) {
	state := &streamState{
		logs: []string{
			"[12:00:00] Log 1",
			"[12:00:01] Log 2",
		},
		status: "Running...",
		ticks:  2,
	}

	rootFunc, _ := createStreamingApp(state)

	// First render
	output1 := runtesting.RenderToString(rootFunc, 80, 24)
	runetui.AssertNotEmpty(t, output1)

	// Add more logs
	state.logs = append(state.logs, "[12:00:02] Log 3")
	state.ticks = 3

	// Second render - should show new content
	output2 := runtesting.RenderToString(rootFunc, 80, 24)
	runetui.AssertNotEmpty(t, output2)
	runetui.AssertContainsText(t, output2, "Log 3")
}

// TestStreamingExample_UpdateFunc verifies the update function works.
func TestStreamingExample_UpdateFunc(t *testing.T) {
	state := &streamState{
		logs:   []string{"[12:00:00] Initial"},
		status: "Running...",
		ticks:  1,
	}

	_, updateFunc := createStreamingApp(state)

	// Send tick message
	cmd := updateFunc(tickMsg{})

	// State should be updated
	if len(state.logs) != 2 {
		t.Errorf("expected 2 logs after tick, got %d", len(state.logs))
	}

	// Command should be returned (to continue ticking)
	if cmd == nil {
		t.Error("expected tick command to be returned")
	}
}

// TestStreamingExample_QuitOnQ verifies quit functionality.
func TestStreamingExample_QuitOnQ(t *testing.T) {
	state := &streamState{
		logs:   []string{},
		status: "Running...",
		ticks:  0,
	}

	_, updateFunc := createStreamingApp(state)

	cmd := updateFunc(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})

	if cmd == nil {
		t.Error("expected quit command")
	}
}
