// Counter example demonstrates state management with RuneTUI.
// This example shows how to use WithUpdate to handle keyboard input
// and update application state following the Elm Architecture pattern.
package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/runetui/runetui"
)

func main() {
	count := 0
	rootFunc, updateFunc := createCounterApp(&count)

	app := runetui.New(rootFunc, runetui.WithUpdate(updateFunc))
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

// createCounterApp creates a counter application with state management.
// It returns the root component function and the update function.
// The count pointer allows the state to be modified externally (useful for testing).
func createCounterApp(count *int) (runetui.ComponentFunc, runetui.UpdateFunc) {
	rootFunc := func() runetui.Component {
		return runetui.Box(
			runetui.BoxProps{
				Direction: runetui.Column,
				Border:    runetui.BorderSingle,
				Padding:   runetui.SpacingAll(1),
			},
			runetui.Text("Counter", runetui.TextProps{Bold: true}),
			runetui.Text(fmt.Sprintf("Count: %d", *count)),
			runetui.Text(""),
			runetui.Text("Press k/↑ to increment", runetui.TextProps{Italic: true}),
			runetui.Text("Press j/↓ to decrement", runetui.TextProps{Italic: true}),
			runetui.Text("Press q to quit", runetui.TextProps{Italic: true}),
		)
	}

	updateFunc := func(msg tea.Msg) tea.Cmd {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "k", "up":
				*count++
			case "j", "down":
				*count--
			case "q":
				return tea.Quit
			}
		}
		return nil
	}

	return rootFunc, updateFunc
}
