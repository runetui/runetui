package runetui

// State Management and Message Handling Patterns
//
// RuneTUI follows the Elm architecture pattern where:
// - State lives OUTSIDE components (in closures or structs)
// - Components are pure functions of state
// - Messages flow through Update functions to modify state
// - State changes trigger re-renders automatically
//
// RuneTUI provides WithUpdate and WithInit options to integrate
// custom state management with the application lifecycle.
//
// This file documents the recommended patterns for managing state
// and handling messages in RuneTUI applications.

// Using WithUpdate and WithInit
//
// WithUpdate allows you to handle all messages in your application:
//
//	updateFunc := func(msg tea.Msg) tea.Cmd {
//	    switch msg := msg.(type) {
//	    case tea.KeyMsg:
//	        // Handle key presses
//	    case customMsg:
//	        // Handle custom messages
//	    }
//	    return nil
//	}
//
//	app := New(rootFunc, WithUpdate(updateFunc))
//
// WithInit runs once when the application starts and can return
// an initial command (useful for starting timers or loading data):
//
//	initFunc := func() tea.Cmd {
//	    return tea.Tick(time.Second, func(t time.Time) tea.Msg {
//	        return tickMsg{}
//	    })
//	}
//
//	app := New(rootFunc, WithInit(initFunc), WithUpdate(updateFunc))

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
//	    // Root component captures state via closure
//	    rootFunc := func() Component {
//	        return VStack(
//	            Text(fmt.Sprintf("Count: %d", count)),
//	            Text("Press k to increment, j to decrement"),
//	        )
//	    }
//
//	    // Update function modifies state based on messages
//	    updateFunc := func(msg tea.Msg) tea.Cmd {
//	        switch msg := msg.(type) {
//	        case tea.KeyMsg:
//	            switch msg.String() {
//	            case "k", "up":
//	                count++
//	            case "j", "down":
//	                count--
//	            case "q":
//	                return tea.Quit
//	            }
//	        }
//	        return nil
//	    }
//
//	    app := New(rootFunc, WithUpdate(updateFunc))
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
//	        name    string
//	        email   string
//	        focused int
//	    }
//	    state := &formState{}
//
//	    // Root component renders form based on state
//	    rootFunc := func() Component {
//	        return VStack(
//	            renderField("Name", state.name, state.focused == 0),
//	            renderField("Email", state.email, state.focused == 1),
//	            Text("Tab: next field | q: quit"),
//	        )
//	    }
//
//	    // Update function handles keyboard input
//	    updateFunc := func(msg tea.Msg) tea.Cmd {
//	        switch msg := msg.(type) {
//	        case tea.KeyMsg:
//	            switch msg.Type {
//	            case tea.KeyTab:
//	                state.focused = (state.focused + 1) % 2
//	            case tea.KeyRunes:
//	                char := string(msg.Runes)
//	                switch state.focused {
//	                case 0:
//	                    state.name += char
//	                case 1:
//	                    state.email += char
//	                }
//	            }
//	            if msg.String() == "q" {
//	                return tea.Quit
//	            }
//	        }
//	        return nil
//	    }
//
//	    app := New(rootFunc, WithUpdate(updateFunc))
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
//	        frame   int
//	    }
//	    state := &appState{loading: true}
//
//	    // Custom message types
//	    type tickMsg struct{}
//	    type dataLoadedMsg string
//
//	    spinnerFrames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
//
//	    // Root component shows loading or data
//	    rootFunc := func() Component {
//	        if state.loading {
//	            frame := spinnerFrames[state.frame%len(spinnerFrames)]
//	            return Text(fmt.Sprintf("%s Loading...", frame))
//	        }
//	        return Text(fmt.Sprintf("Data: %s", state.data))
//	    }
//
//	    // Update function handles async messages
//	    updateFunc := func(msg tea.Msg) tea.Cmd {
//	        switch msg := msg.(type) {
//	        case tickMsg:
//	            if state.loading {
//	                state.frame++
//	                return tick()
//	            }
//	        case dataLoadedMsg:
//	            state.loading = false
//	            state.data = string(msg)
//	        case tea.KeyMsg:
//	            if msg.String() == "q" {
//	                return tea.Quit
//	            }
//	        }
//	        return nil
//	    }
//
//	    // Init function starts the timer and data load
//	    initFunc := func() tea.Cmd {
//	        return tea.Batch(loadData(), tick())
//	    }
//
//	    app := New(rootFunc, WithInit(initFunc), WithUpdate(updateFunc))
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
//    - type aliases for simple data (type dataLoadedMsg string)
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
//    - Use pointers for shared state that needs mutation
//
// 2. State Updates:
//    - All state updates happen in the UpdateFunc
//    - UpdateFunc is called for every message
//    - State changes automatically trigger re-renders
//
// 3. State Sharing:
//    - Multiple components can read the same state via closure
//    - UpdateFunc is the single source of state mutations
//    - No need for context or props drilling
//
// Update Function Signature
//
// The UpdateFunc follows Bubble Tea's signature:
//
//	type UpdateFunc func(msg tea.Msg) tea.Cmd
//
//	updateFunc := func(msg tea.Msg) tea.Cmd {
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
// Init Function Signature
//
// The InitFunc runs once at startup:
//
//	type InitFunc func() tea.Cmd
//
//	initFunc := func() tea.Cmd {
//	    return tea.Batch(loadData(), startTimer())
//	}
//
// - Returns tea.Cmd (optional initial command)
// - Use tea.Batch() to run multiple commands
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
//	// Multiple commands
//	return tea.Batch(cmd1, cmd2, cmd3)
//
// Common Patterns
//
// Conditional rendering based on state:
//
//	rootFunc := func() Component {
//	    if state.loading {
//	        return Text("Loading...")
//	    }
//	    return Text(state.data)
//	}
//
// List rendering:
//
//	rootFunc := func() Component {
//	    children := make([]Component, len(items))
//	    for i, item := range items {
//	        children[i] = Text(item)
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
//	    case "q":
//	        return tea.Quit
//	    }
//
// Testing Patterns
//
// Test state updates independently:
//
//	func TestUpdate_IncrementKey_UpdatesCount(t *testing.T) {
//	    count := 0
//	    updateFunc := func(msg tea.Msg) tea.Cmd {
//	        if keyMsg, ok := msg.(tea.KeyMsg); ok {
//	            if keyMsg.String() == "k" {
//	                count++
//	            }
//	        }
//	        return nil
//	    }
//
//	    updateFunc(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'k'}})
//	    if count != 1 {
//	        t.Errorf("expected count=1, got %d", count)
//	    }
//	}
//
// Test component rendering with runetui/testing:
//
//	func TestRootFunc_WithState_RendersCorrectly(t *testing.T) {
//	    count := 5
//	    rootFunc := func() Component {
//	        return Text(fmt.Sprintf("Count: %d", count))
//	    }
//
//	    output := testing.RenderToString(rootFunc, 80, 24)
//	    testing.AssertContainsText(t, output, "Count: 5")
//	}
//
// For full working examples, see the examples/ directory:
// - examples/counter - Simple counter with increment/decrement
// - examples/form - Multi-field form with navigation
// - examples/async - Async loading with spinner
// - examples/streaming - Static log zones with accumulating content
