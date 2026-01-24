// Async example demonstrates asynchronous operations with RuneTUI.
// This example shows how to use WithInit for initial commands and
// handle async responses following the Elm Architecture pattern.
package main

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/runetui/runetui"
)

// asyncState holds the async operation state.
type asyncState struct {
	loading bool
	data    string
	err     string
	frame   int
}

// Messages for async operations.
type dataLoadedMsg string
type errorMsg string
type tickMsg struct{}

var spinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

func main() {
	state := &asyncState{loading: true}
	rootFunc, updateFunc := createAsyncApp(state)

	initFunc := func() tea.Cmd {
		return tea.Batch(loadData(), tick())
	}

	app := runetui.New(rootFunc,
		runetui.WithInit(initFunc),
		runetui.WithUpdate(updateFunc),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

// createAsyncApp creates an async loading application.
func createAsyncApp(state *asyncState) (runetui.ComponentFunc, runetui.UpdateFunc) {
	rootFunc := func() runetui.Component {
		var content runetui.Component

		if state.loading {
			spinner := spinnerFrames[state.frame%len(spinnerFrames)]
			content = runetui.Text(spinner + " Loading...")
		} else if state.err != "" {
			content = runetui.VStack(
				runetui.Text("Error!", runetui.TextProps{Bold: true}),
				runetui.Text(state.err),
			)
		} else {
			content = runetui.VStack(
				runetui.Text("Data Loaded!", runetui.TextProps{Bold: true}),
				runetui.Text(state.data),
			)
		}

		return runetui.Box(
			runetui.BoxProps{
				Direction: runetui.Column,
				Border:    runetui.BorderSingle,
				Padding:   runetui.SpacingAll(1),
			},
			runetui.Text("Async Example", runetui.TextProps{Bold: true}),
			runetui.Text(""),
			content,
			runetui.Text(""),
			runetui.Text("Press q to quit", runetui.TextProps{Italic: true}),
		)
	}

	updateFunc := func(msg tea.Msg) tea.Cmd {
		switch msg := msg.(type) {
		case dataLoadedMsg:
			state.loading = false
			state.data = string(msg)
		case errorMsg:
			state.loading = false
			state.err = string(msg)
		case tickMsg:
			if state.loading {
				state.frame++
				return tick()
			}
		case tea.KeyMsg:
			if msg.String() == "q" {
				return tea.Quit
			}
		}
		return nil
	}

	return rootFunc, updateFunc
}

func loadData() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(2 * time.Second)
		return dataLoadedMsg("Hello from the server!")
	}
}

func tick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}
