package main

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/runetui/runetui"
)

// This example demonstrates streaming logs using RuneTUI's Static component.
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
//
// This example simulates an agent or build process that:
// 1. Accumulates log entries in a static zone (no flicker, efficient)
// 2. Shows current status in a dynamic zone (updates on each render)
// 3. Adds new log entries automatically every 500ms
// 4. Handles terminal resize gracefully
// 5. Quits on Ctrl+C
//
// Use case: AI agents, build logs, streaming data, real-time monitoring
//
// State Management Pattern:
// This example uses the Elm architecture pattern where state lives in the
// model struct, messages flow through Update(), and View() is a pure function
// of state. This is the pattern RuneTUI will support once custom Init/Update
// handlers are exposed in the adapter.

// tickMsg is sent by the timer to trigger periodic log updates
type tickMsg time.Time

// model holds the application state
type model struct {
	// logs accumulate in the static zone
	logs []string
	// status shows in the dynamic zone
	status string
	// ticks counts the number of updates
	ticks int
	// staticManager tracks static content across renders
	staticManager *runetui.StaticManager
	// layoutEngine handles component layout
	layoutEngine *runetui.LayoutEngine
}

// Init initializes the model and returns the first command.
// This starts the periodic timer that adds log entries.
func (m model) Init() tea.Cmd {
	// Start ticking immediately to add log entries
	return tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Update handles incoming messages and updates state.
// This is where all state mutations happen in response to events.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		// Add new log entry with timestamp
		timestamp := time.Now().Format("15:04:05")
		m.logs = append(m.logs, fmt.Sprintf("[%s] Log entry %d", timestamp, m.ticks))
		m.ticks++

		// Update status based on progress
		if m.ticks < 20 {
			m.status = fmt.Sprintf("Running... (%d entries)", m.ticks)
			// Continue ticking to add more logs
			return m, tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
				return tickMsg(t)
			})
		} else {
			m.status = "Complete! Press Ctrl+C to quit"
			// Stop ticking after 20 entries
			return m, nil
		}

	case tea.WindowSizeMsg:
		// Handle terminal resize by recreating the layout engine
		m.layoutEngine = runetui.NewLayoutEngine(msg.Width, msg.Height)

	case tea.KeyMsg:
		// Handle keyboard input
		switch msg.String() {
		case "ctrl+c", "q":
			// Quit on Ctrl+C or 'q'
			return m, tea.Quit
		case " ":
			// Manually add a log entry on spacebar
			timestamp := time.Now().Format("15:04:05")
			m.logs = append(m.logs, fmt.Sprintf("[%s] Manual entry added", timestamp))
			m.status = fmt.Sprintf("Running... (%d entries)", len(m.logs))
		}
	}

	return m, nil
}

// View renders the component tree to a string.
// This is a pure function - it doesn't modify state, just renders it.
func (m model) View() string {
	// Set the static manager for this render cycle.
	// This allows the Static component to access the manager.
	runetui.SetStaticManager(m.staticManager)
	defer runetui.SetStaticManager(nil)

	// Build the component tree using RuneTUI components.
	// State is captured via closure in this view function.
	root := m.buildComponentTree()

	// Calculate layout for all components
	tree := m.layoutEngine.CalculateLayout(root)

	// Render static content (accumulated logs)
	staticContent := m.staticManager.RenderStatic()

	// Render dynamic content (status bar, etc.)
	dynamicContent := renderTree(tree)

	// Combine static and dynamic zones
	if staticContent == "" {
		return dynamicContent
	}
	if dynamicContent == "" {
		return staticContent
	}
	return staticContent + "\n" + dynamicContent
}

// buildComponentTree constructs the UI from current state.
// This is called on every render but the Static component ensures
// that old log entries are not re-rendered.
func (m model) buildComponentTree() runetui.Component {
	// Build log components from current state
	logComponents := make([]runetui.Component, 0, len(m.logs))
	for _, logLine := range m.logs {
		logComponents = append(logComponents,
			runetui.Text(logLine, runetui.TextProps{
				Color: "#888888",
			}),
		)
	}

	return runetui.VStack(
		// Title section
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

		// Static log zone - KEY FEATURE!
		// Old log entries are NOT re-rendered, only new ones are drawn.
		// This prevents flickering and improves performance dramatically.
		// The "logs" key identifies this static zone across renders.
		runetui.Static(runetui.StaticProps{Key: "logs"}, func() []runetui.Component {
			return logComponents
		}),

		// Separator line
		runetui.Text("────────────────────────────────────────", runetui.TextProps{
			Color: "#444444",
		}),

		// Dynamic status bar - updates on every render
		// This demonstrates the contrast with the static zone above.
		runetui.Box(
			runetui.BoxProps{
				Background: "#004455",
				Padding:    runetui.SpacingAll(1),
			},
			runetui.Text(m.status, runetui.TextProps{
				Color: "#FFFFFF",
				Bold:  true,
			}),
		),

		// Help text
		runetui.Text("Press SPACE to add entry | Ctrl+C to quit", runetui.TextProps{
			Color: "#666666",
		}),
	)
}

// renderTree recursively renders a layout tree to a string.
// This is a helper function that traverses the component tree.
func renderTree(tree *runetui.LayoutTree) string {
	if tree == nil {
		return ""
	}

	rendered := tree.Component.Render(tree.Layout)

	for _, child := range tree.Children {
		childOutput := renderTree(child)
		if childOutput != "" {
			rendered += childOutput
		}
	}

	return rendered
}

func main() {
	// Initialize the model with starting state
	m := model{
		logs: []string{
			"[" + time.Now().Format("15:04:05") + "] Application started",
			"[" + time.Now().Format("15:04:05") + "] Initializing components...",
			"[" + time.Now().Format("15:04:05") + "] Ready! Logs will stream automatically...",
		},
		status:        "Starting...",
		ticks:         3,
		staticManager: runetui.NewStaticManager(),
		layoutEngine:  runetui.NewLayoutEngine(80, 24),
	}

	// Create and run the Bubble Tea program
	// This is the direct Bubble Tea integration that RuneTUI will
	// eventually expose through its adapter with custom Init/Update support.
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	// Real-world use cases for this pattern:
	//
	// 1. AI Agent Output:
	//    - Agent thoughts/actions accumulate in static zone
	//    - Current step shows in dynamic status bar
	//    - No flickering of previous output
	//
	// 2. Build/Deploy Logs:
	//    - Build output accumulates as it's generated
	//    - Progress bar shows current build stage
	//    - Terminal remains readable and efficient
	//
	// 3. Test Runner:
	//    - Test results accumulate in static zone
	//    - Current test name shows in status bar
	//    - Fast rendering even with thousands of tests
	//
	// 4. Database Migration:
	//    - Migration logs accumulate
	//    - Current migration step shows in status
	//    - Clear audit trail of all operations
}
