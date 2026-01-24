// Streaming example demonstrates streaming logs using RuneTUI's Static component.
//
// This example shows how to use WithInit and WithUpdate for stateful applications
// with accumulating content following the Elm Architecture pattern.
//
// IMPORTANT CONCEPT - Static vs Dynamic Zones:
//
// Static zones (using Static component):
//   - Content accumulates across renders
//   - Old content is NOT re-rendered (efficient for logs)
//   - Only NEW content is rendered on each update
//   - Perfect for: logs, agent output, build output, streaming data
//
// Dynamic zones (regular components):
//   - Content is re-rendered on every frame
//   - Shows real-time updates
//   - Perfect for: status bars, progress indicators, current state
package main

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/runetui/runetui"
)

// streamState holds the streaming application state.
type streamState struct {
	logs   []string
	status string
	ticks  int
}

// tickMsg is sent by the timer to trigger periodic log updates.
type tickMsg time.Time

func main() {
	state := &streamState{
		logs: []string{
			"[" + time.Now().Format("15:04:05") + "] Application started",
			"[" + time.Now().Format("15:04:05") + "] Initializing components...",
			"[" + time.Now().Format("15:04:05") + "] Ready! Logs will stream automatically...",
		},
		status: "Starting...",
		ticks:  3,
	}

	rootFunc, updateFunc := createStreamingApp(state)
	initFunc := func() tea.Cmd {
		return tea.Tick(500*time.Millisecond, func(t time.Time) tea.Msg {
			return tickMsg(t)
		})
	}

	app := runetui.New(rootFunc,
		runetui.WithInit(initFunc),
		runetui.WithUpdate(updateFunc),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

// createStreamingApp creates a streaming log application.
func createStreamingApp(state *streamState) (runetui.ComponentFunc, runetui.UpdateFunc) {
	rootFunc := func() runetui.Component {
		logComponents := make([]runetui.Component, 0, len(state.logs))
		for _, logLine := range state.logs {
			logComponents = append(logComponents,
				runetui.Text(logLine, runetui.TextProps{Color: "#888888"}),
			)
		}

		return runetui.VStack(
			runetui.Box(
				runetui.BoxProps{
					Padding:    runetui.SpacingAll(1),
					Background: "#005577",
				},
				runetui.Text("Streaming Logs Example", runetui.TextProps{
					Color: "#FFFFFF",
					Bold:  true,
				}),
			),
			runetui.Static(runetui.StaticProps{Key: "logs"}, func() []runetui.Component {
				return logComponents
			}),
			runetui.Text("────────────────────────────────────────",
				runetui.TextProps{Color: "#444444"}),
			runetui.Box(
				runetui.BoxProps{
					Background: "#004455",
					Padding:    runetui.SpacingAll(1),
				},
				runetui.Text(state.status, runetui.TextProps{
					Color: "#FFFFFF",
					Bold:  true,
				}),
			),
			runetui.Text("Press SPACE to add entry | q to quit",
				runetui.TextProps{Color: "#666666"}),
		)
	}

	updateFunc := func(msg tea.Msg) tea.Cmd {
		switch msg := msg.(type) {
		case tickMsg:
			timestamp := time.Now().Format("15:04:05")
			state.logs = append(state.logs,
				fmt.Sprintf("[%s] Log entry %d", timestamp, state.ticks))
			state.ticks++

			if state.ticks < 20 {
				state.status = fmt.Sprintf("Running... (%d entries)", state.ticks)
				return tea.Tick(500*time.Millisecond, func(t time.Time) tea.Msg {
					return tickMsg(t)
				})
			}
			state.status = "Complete! Press q to quit"
			return nil

		case tea.KeyMsg:
			switch msg.String() {
			case "q":
				return tea.Quit
			case " ":
				timestamp := time.Now().Format("15:04:05")
				state.logs = append(state.logs,
					fmt.Sprintf("[%s] Manual entry added", timestamp))
				state.status = fmt.Sprintf("Running... (%d entries)", len(state.logs))
			}
		}
		return nil
	}

	return rootFunc, updateFunc
}
