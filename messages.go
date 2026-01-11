package runetui

// State Management and Message Handling Patterns
//
// RuneTUI follows the Elm architecture pattern where:
// - State lives OUTSIDE components (in closures or structs)
// - Components are pure functions of state
// - Messages flow through Update functions to modify state
// - State changes trigger re-renders automatically
//
// This file documents the recommended patterns for managing state
// and handling messages in RuneTUI applications.

// Pattern 1: Counter App (Simple State)
//
// This pattern demonstrates basic state management with increment/decrement actions.
//
// Example:
//
//	func main() {
//	    // State lives in main
//	    count := 0
//
//	    // Custom message types
//	    type incrementMsg struct{}
//	    type decrementMsg struct{}
//
//	    // Update function modifies state based on messages
//	    update := func(msg tea.Msg) tea.Cmd {
//	        switch msg.(type) {
//	        case incrementMsg:
//	            count++
//	        case decrementMsg:
//	            count--
//	        case tea.KeyMsg:
//	            key := msg.(tea.KeyMsg)
//	            if key.String() == "+" {
//	                return func() tea.Msg { return incrementMsg{} }
//	            }
//	            if key.String() == "-" {
//	                return func() tea.Msg { return decrementMsg{} }
//	            }
//	        }
//	        return nil
//	    }
//
//	    // Root component captures state via closure
//	    rootFunc := func() Component {
//	        return VStack(
//	            Text(fmt.Sprintf("Count: %d", count), TextProps{}),
//	            Text("Press + to increment, - to decrement", TextProps{}),
//	        )
//	    }
//
//	    app := New(rootFunc)
//	    app.Run()
//	}

// Pattern 2: Form with Multiple Inputs (Structured State)
//
// This pattern demonstrates managing structured state with multiple fields.
//
// Example:
//
//	func main() {
//	    // Structured state
//	    type formState struct {
//	        name     string
//	        email    string
//	        focusedField int
//	    }
//	    state := formState{}
//
//	    // Custom message types with data
//	    type nameInputMsg struct{ value string }
//	    type emailInputMsg struct{ value string }
//
//	    // Update function handles different message types
//	    update := func(msg tea.Msg) tea.Cmd {
//	        switch msg := msg.(type) {
//	        case nameInputMsg:
//	            state.name = msg.value
//	        case emailInputMsg:
//	            state.email = msg.value
//	        case tea.KeyMsg:
//	            // Handle keyboard input
//	            if msg.Type == tea.KeyTab {
//	                state.focusedField = (state.focusedField + 1) % 2
//	            }
//	        }
//	        return nil
//	    }
//
//	    // Root component renders form based on state
//	    rootFunc := func() Component {
//	        return VStack(
//	            Text(fmt.Sprintf("Name: %s", state.name), TextProps{}),
//	            Text(fmt.Sprintf("Email: %s", state.email), TextProps{}),
//	        )
//	    }
//
//	    app := New(rootFunc)
//	    app.Run()
//	}

// Pattern 3: Async Operations with Commands (Loading States)
//
// This pattern demonstrates async operations like data loading with spinners.
//
// Example:
//
//	func main() {
//	    // State with loading indicator
//	    type appState struct {
//	        loading bool
//	        data    string
//	        ticks   int
//	    }
//	    state := appState{loading: true}
//
//	    // Custom message types
//	    type tickMsg struct{}
//	    type dataLoadedMsg struct{ data string }
//
//	    // Update function handles async messages
//	    update := func(msg tea.Msg) tea.Cmd {
//	        switch msg := msg.(type) {
//	        case tickMsg:
//	            state.ticks++
//	            // Return command to tick again
//	            return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
//	                return tickMsg{}
//	            })
//	        case dataLoadedMsg:
//	            state.loading = false
//	            state.data = msg.data
//	            return nil
//	        case tea.KeyMsg:
//	            if msg.Type == tea.KeyEnter && state.loading {
//	                // Simulate data load
//	                return func() tea.Msg {
//	                    time.Sleep(time.Second * 2)
//	                    return dataLoadedMsg{data: "Hello World"}
//	                }
//	            }
//	        }
//	        return nil
//	    }
//
//	    // Root component shows loading or data
//	    rootFunc := func() Component {
//	        if state.loading {
//	            spinner := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
//	            frame := spinner[state.ticks%len(spinner)]
//	            return Text(fmt.Sprintf("%s Loading...", frame), TextProps{})
//	        }
//	        return Text(fmt.Sprintf("Data: %s", state.data), TextProps{})
//	    }
//
//	    app := New(rootFunc)
//	    app.Run()
//	}

// Message Type Guidelines
//
// 1. Use Bubble Tea's built-in message types when possible:
//    - tea.KeyMsg for keyboard input
//    - tea.MouseMsg for mouse events
//    - tea.WindowSizeMsg for terminal resize
//
// 2. Define custom message types for domain events:
//    - struct{} for simple events (incrementMsg, saveMsg)
//    - struct{ data T } for events with data (inputMsg, loadedMsg)
//
// 3. Message naming convention:
//    - Use descriptive past-tense names for completed actions (dataLoadedMsg)
//    - Use imperative names for commands (incrementMsg, saveMsg)
//
// State Management Guidelines
//
// 1. State Location:
//    - State lives in main() or in the scope where New() is called
//    - Components capture state via closures in rootFunc
//    - Never store state inside Component structs
//
// 2. State Updates:
//    - All state updates happen in the Update function
//    - Update function is called for every message
//    - State changes automatically trigger re-renders
//
// 3. State Sharing:
//    - Multiple components can read the same state via closure
//    - Update function is the single source of state mutations
//    - No need for context or props drilling
//
// Update Function Signature
//
// The Update function follows Bubble Tea's signature:
//
//	func update(msg tea.Msg) tea.Cmd {
//	    switch msg := msg.(type) {
//	    case CustomMsg:
//	        // Handle custom message
//	        return nil
//	    }
//	    return nil
//	}
//
// - Takes a tea.Msg (any message type)
// - Returns tea.Cmd (optional command to execute)
// - Use type switch to handle different message types
//
// Returning Commands (Side Effects)
//
// Commands are functions that return messages asynchronously:
//
//	// Simple command (immediate)
//	return func() tea.Msg {
//	    return myMsg{}
//	}
//
//	// Async command (with delay)
//	return func() tea.Msg {
//	    time.Sleep(time.Second)
//	    return myMsg{}
//	}
//
//	// Bubble Tea helper (tick)
//	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
//	    return tickMsg{}
//	})
//
// Common Patterns
//
// Conditional rendering based on state:
//
//	rootFunc := func() Component {
//	    if state.loading {
//	        return Text("Loading...", TextProps{})
//	    }
//	    return Text(state.data, TextProps{})
//	}
//
// List rendering:
//
//	rootFunc := func() Component {
//	    children := make([]Component, len(items))
//	    for i, item := range items {
//	        children[i] = Text(item, TextProps{})
//	    }
//	    return VStack(children...)
//	}
//
// Event handlers via keyboard:
//
//	case tea.KeyMsg:
//	    switch msg.String() {
//	    case "enter":
//	        return submitCommand()
//	    case "esc":
//	        return tea.Quit
//	    }
//
// Testing Patterns
//
// Test state updates independently:
//
//	func TestUpdate_IncrementMsg_UpdatesCount(t *testing.T) {
//	    count := 0
//	    update := func(msg tea.Msg) tea.Cmd {
//	        if _, ok := msg.(incrementMsg); ok {
//	            count++
//	        }
//	        return nil
//	    }
//
//	    update(incrementMsg{})
//	    if count != 1 {
//	        t.Errorf("expected count=1, got %d", count)
//	    }
//	}
//
// Test component rendering:
//
//	func TestRootFunc_WithState_RendersCorrectly(t *testing.T) {
//	    count := 5
//	    rootFunc := func() Component {
//	        return Text(fmt.Sprintf("Count: %d", count), TextProps{})
//	    }
//
//	    comp := rootFunc()
//	    rendered := comp.Render(Layout{Width: 8, Height: 1})
//	    if rendered != "Count: 5" {
//	        t.Errorf("unexpected render: %q", rendered)
//	    }
//	}
//
// For full working examples, see the test file messages_test.go.
