package runetui

import (
	"fmt"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

// TestCounterPattern_IncrementMsg_UpdatesCount demonstrates state management pattern
// where state lives outside components (Elm-like architecture).
func TestCounterPattern_IncrementMsg_UpdatesCount(t *testing.T) {
	// Counter state (lives outside component)
	count := 0

	// Custom message type
	type incrementMsg struct{}

	// Update function handles messages and updates state
	update := func(msg tea.Msg) tea.Cmd {
		switch msg.(type) {
		case incrementMsg:
			count++
			return nil
		}
		return nil
	}

	// Root component function that captures state via closure
	rootFunc := func() Component {
		return Text(fmt.Sprintf("Count: %d", count), TextProps{})
	}

	// Initial state
	if count != 0 {
		t.Errorf("expected initial count to be 0, got %d", count)
	}

	// Process increment message
	update(incrementMsg{})

	if count != 1 {
		t.Errorf("expected count to be 1 after increment, got %d", count)
	}

	// Verify component reflects new state
	comp := rootFunc()
	layout := Layout{Width: 8, Height: 1} // Width matches "Count: 1"
	rendered := comp.Render(layout)

	expected := "Count: 1"
	if rendered != expected {
		t.Errorf("expected rendered text to be %q, got %q", expected, rendered)
	}
}

// TestCounterPattern_DecrementMsg_UpdatesCount demonstrates decrement message handling.
func TestCounterPattern_DecrementMsg_UpdatesCount(t *testing.T) {
	// Counter state
	count := 5

	// Custom message types
	type incrementMsg struct{}
	type decrementMsg struct{}

	// Update function handles both increment and decrement
	update := func(msg tea.Msg) tea.Cmd {
		switch msg.(type) {
		case incrementMsg:
			count++
		case decrementMsg:
			count--
		}
		return nil
	}

	// Root component function
	rootFunc := func() Component {
		return Text(fmt.Sprintf("Count: %d", count), TextProps{})
	}

	// Initial state
	if count != 5 {
		t.Errorf("expected initial count to be 5, got %d", count)
	}

	// Process decrement message
	update(decrementMsg{})

	if count != 4 {
		t.Errorf("expected count to be 4 after decrement, got %d", count)
	}

	// Process increment message
	update(incrementMsg{})

	if count != 5 {
		t.Errorf("expected count to be 5 after increment, got %d", count)
	}

	// Verify component reflects new state
	comp := rootFunc()
	layout := Layout{Width: 8, Height: 1}
	rendered := comp.Render(layout)

	expected := "Count: 5"
	if rendered != expected {
		t.Errorf("expected rendered text to be %q, got %q", expected, rendered)
	}
}

// TestFormPattern_InputMsg_StoresMultipleValues demonstrates form state management
// with multiple input fields.
func TestFormPattern_InputMsg_StoresMultipleValues(t *testing.T) {
	// Form state
	type formState struct {
		name  string
		email string
	}
	form := formState{}

	// Custom message types
	type nameInputMsg struct{ value string }
	type emailInputMsg struct{ value string }

	// Update function handles different input types
	update := func(msg tea.Msg) tea.Cmd {
		switch msg := msg.(type) {
		case nameInputMsg:
			form.name = msg.value
		case emailInputMsg:
			form.email = msg.value
		}
		return nil
	}

	// Root component function showing form
	rootFunc := func() Component {
		return VStack(
			Text(fmt.Sprintf("Name: %s", form.name), TextProps{}),
			Text(fmt.Sprintf("Email: %s", form.email), TextProps{}),
		)
	}

	// Initial state
	if form.name != "" {
		t.Errorf("expected initial name to be empty, got %q", form.name)
	}
	if form.email != "" {
		t.Errorf("expected initial email to be empty, got %q", form.email)
	}

	// Process name input
	update(nameInputMsg{value: "Alice"})

	if form.name != "Alice" {
		t.Errorf("expected name to be 'Alice', got %q", form.name)
	}

	// Process email input
	update(emailInputMsg{value: "alice@example.com"})

	if form.email != "alice@example.com" {
		t.Errorf("expected email to be 'alice@example.com', got %q", form.email)
	}

	// Verify component reflects state
	comp := rootFunc()
	if comp == nil {
		t.Fatal("rootFunc should return non-nil component")
	}

	// Verify children count
	children := comp.Children()
	if len(children) != 2 {
		t.Errorf("expected 2 children, got %d", len(children))
	}
}

// TestAsyncPattern_LoadingState_WithTickMsg demonstrates async state management
// with commands and effects (like data loading with spinner).
func TestAsyncPattern_LoadingState_WithTickMsg(t *testing.T) {
	// Async state
	type asyncState struct {
		loading bool
		data    string
		ticks   int
	}
	state := asyncState{loading: true}

	// Custom message types
	type tickMsg struct{}
	type dataLoadedMsg struct{ data string }

	// Update function handles async messages
	update := func(msg tea.Msg) tea.Cmd {
		switch msg := msg.(type) {
		case tickMsg:
			state.ticks++
			// In real app, would return tea.Tick command
			return nil
		case dataLoadedMsg:
			state.loading = false
			state.data = msg.data
			return nil
		}
		return nil
	}

	// Root component function showing loading state
	rootFunc := func() Component {
		if state.loading {
			spinner := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
			frame := spinner[state.ticks%len(spinner)]
			return Text(fmt.Sprintf("%s Loading...", frame), TextProps{})
		}
		return Text(fmt.Sprintf("Data: %s", state.data), TextProps{})
	}

	// Initial state
	if !state.loading {
		t.Error("expected initial loading to be true")
	}
	if state.data != "" {
		t.Errorf("expected initial data to be empty, got %q", state.data)
	}
	if state.ticks != 0 {
		t.Errorf("expected initial ticks to be 0, got %d", state.ticks)
	}

	// Process tick messages (spinner animation)
	update(tickMsg{})
	if state.ticks != 1 {
		t.Errorf("expected ticks to be 1, got %d", state.ticks)
	}

	update(tickMsg{})
	if state.ticks != 2 {
		t.Errorf("expected ticks to be 2, got %d", state.ticks)
	}

	// Verify loading state renders spinner
	comp := rootFunc()
	layout := Layout{Width: 12, Height: 1} // Width matches "⠹ Loading..."
	rendered := comp.Render(layout)
	expected := "⠹ Loading..."
	if rendered != expected {
		t.Errorf("expected rendered text to be %q, got %q", expected, rendered)
	}

	// Process data loaded message
	update(dataLoadedMsg{data: "Hello World"})

	if state.loading {
		t.Error("expected loading to be false after data loaded")
	}
	if state.data != "Hello World" {
		t.Errorf("expected data to be 'Hello World', got %q", state.data)
	}

	// Verify loaded state renders data
	comp = rootFunc()
	layout = Layout{Width: 17, Height: 1}
	rendered = comp.Render(layout)
	expected = "Data: Hello World"
	if rendered != expected {
		t.Errorf("expected rendered text to be %q, got %q", expected, rendered)
	}
}
