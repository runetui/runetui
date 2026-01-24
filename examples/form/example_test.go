package main

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/runetui/runetui"
	rtest "github.com/runetui/runetui/testing"
)

func TestFormExample_RendersInitialState(t *testing.T) {
	state := &formState{focused: 0}
	rootFunc, _ := createFormApp(state)

	output := rtest.RenderToString(rootFunc, 50, 15)

	runetui.AssertContainsText(t, output, "Name")
	runetui.AssertContainsText(t, output, "Email")
}

func TestFormExample_ShowsFieldValues(t *testing.T) {
	state := &formState{
		name:    "John",
		email:   "john@example.com",
		focused: 0,
	}
	rootFunc, _ := createFormApp(state)

	output := rtest.RenderToString(rootFunc, 50, 15)

	runetui.AssertContainsText(t, output, "John")
	runetui.AssertContainsText(t, output, "john@example.com")
}

func TestFormExample_NavigateFields(t *testing.T) {
	state := &formState{focused: 0}
	rootFunc, updateFunc := createFormApp(state)

	// Initial: first field focused
	output1 := rtest.RenderToString(rootFunc, 50, 15)
	runetui.AssertContainsText(t, output1, "> Name")

	// Press tab to move to next field
	updateFunc(tea.KeyMsg{Type: tea.KeyTab})

	output2 := rtest.RenderToString(rootFunc, 50, 15)
	runetui.AssertContainsText(t, output2, "> Email")
}

func TestFormExample_TypeInField(t *testing.T) {
	state := &formState{focused: 0}
	rootFunc, updateFunc := createFormApp(state)

	// Type some characters
	for _, r := range "Alice" {
		updateFunc(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}})
	}

	output := rtest.RenderToString(rootFunc, 50, 15)
	runetui.AssertContainsText(t, output, "Alice")
}

func TestFormExample_Snapshot(t *testing.T) {
	state := &formState{
		name:    "Test User",
		email:   "test@test.com",
		focused: 1,
	}
	rootFunc, _ := createFormApp(state)

	output := rtest.RenderToString(rootFunc, 50, 15)

	rtest.AssertSnapshot(t, "form_display", output)
}
