// Form example demonstrates structured state management with RuneTUI.
// This example shows how to handle multiple input fields with navigation
// following the Elm Architecture pattern.
package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/runetui/runetui"
)

// formState holds the form data and current focus.
type formState struct {
	name    string
	email   string
	focused int
}

const numFields = 2

func main() {
	state := &formState{}
	rootFunc, updateFunc := createFormApp(state)

	app := runetui.New(rootFunc, runetui.WithUpdate(updateFunc))
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

// createFormApp creates a form application with state management.
func createFormApp(state *formState) (runetui.ComponentFunc, runetui.UpdateFunc) {
	rootFunc := func() runetui.Component {
		return runetui.Box(
			runetui.BoxProps{
				Direction: runetui.Column,
				Border:    runetui.BorderSingle,
				Padding:   runetui.SpacingAll(1),
			},
			runetui.Text("Form Example", runetui.TextProps{Bold: true}),
			runetui.Text(""),
			renderField("Name", state.name, state.focused == 0),
			renderField("Email", state.email, state.focused == 1),
			runetui.Text(""),
			runetui.Text("Tab: next field | Enter: submit | q: quit",
				runetui.TextProps{Italic: true}),
		)
	}

	updateFunc := func(msg tea.Msg) tea.Cmd {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyTab:
				state.focused = (state.focused + 1) % numFields
			case tea.KeyShiftTab:
				state.focused = (state.focused - 1 + numFields) % numFields
			case tea.KeyBackspace:
				switch state.focused {
				case 0:
					if len(state.name) > 0 {
						state.name = state.name[:len(state.name)-1]
					}
				case 1:
					if len(state.email) > 0 {
						state.email = state.email[:len(state.email)-1]
					}
				}
			case tea.KeyRunes:
				char := string(msg.Runes)
				switch state.focused {
				case 0:
					state.name += char
				case 1:
					state.email += char
				}
			}
			if msg.String() == "q" {
				return tea.Quit
			}
		}
		return nil
	}

	return rootFunc, updateFunc
}

func renderField(label, value string, focused bool) runetui.Component {
	prefix := "  "
	if focused {
		prefix = "> "
	}
	display := value
	if display == "" {
		display = "(empty)"
	}
	return runetui.Text(fmt.Sprintf("%s%s: %s", prefix, label, display))
}
