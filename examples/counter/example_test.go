package main

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/runetui/runetui"
	rtest "github.com/runetui/runetui/testing"
)

func TestCounterExample_RendersInitialState(t *testing.T) {
	count := 0
	rootFunc, _ := createCounterApp(&count)

	output := rtest.RenderToString(rootFunc, 40, 10)

	runetui.AssertContainsText(t, output, "0")
	runetui.AssertContainsText(t, output, "Counter")
}

func TestCounterExample_IncrementOnKeyUp(t *testing.T) {
	count := 0
	rootFunc, updateFunc := createCounterApp(&count)

	// Initial render
	output1 := rtest.RenderToString(rootFunc, 40, 10)
	runetui.AssertContainsText(t, output1, "0")

	// Simulate pressing 'k' (up)
	updateFunc(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'k'}})

	// Render after increment
	output2 := rtest.RenderToString(rootFunc, 40, 10)
	runetui.AssertContainsText(t, output2, "1")
}

func TestCounterExample_DecrementOnKeyDown(t *testing.T) {
	count := 5
	rootFunc, updateFunc := createCounterApp(&count)

	// Initial render
	output1 := rtest.RenderToString(rootFunc, 40, 10)
	runetui.AssertContainsText(t, output1, "5")

	// Simulate pressing 'j' (down)
	updateFunc(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})

	// Render after decrement
	output2 := rtest.RenderToString(rootFunc, 40, 10)
	runetui.AssertContainsText(t, output2, "4")
}

func TestCounterExample_Snapshot(t *testing.T) {
	count := 42
	rootFunc, _ := createCounterApp(&count)

	output := rtest.RenderToString(rootFunc, 40, 10)

	rtest.AssertSnapshot(t, "counter_display", output)
}
