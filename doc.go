// Package runetui provides a declarative TUI framework for Go.
//
// RuneTUI brings React/Ink-style component composition to terminal applications,
// built on top of Bubble Tea. It offers a higher-level, declarative API for building
// rich terminal user interfaces with less boilerplate.
//
// # Design Philosophy
//
// RuneTUI is inspired by Ink (https://github.com/vadimdemedes/ink), which brought
// React's declarative paradigm to Node.js CLIs. The goal is to provide Go developers
// with a similar experience while leveraging Go's mature TUI ecosystem, particularly
// Bubble Tea (https://github.com/charmbracelet/bubbletea).
//
// Key principles:
//   - Declarative, component-based API
//   - Flexbox-inspired layout system
//   - Clear separation between static and dynamic UI zones
//   - First-class testing support
//   - Patterns optimized for CLI agents and streaming workflows
//
// # Status
//
// ⚠️ RuneTUI is in early development. APIs will change as we explore the design space.
// Feedback and contributions are welcome at https://github.com/runetui/runetui
//
// # Basic Example
//
// A simple "Hello World" application:
//
//	package main
//
//	import "github.com/runetui/runetui"
//
//	func main() {
//		app := runetui.New(App)
//		if err := app.Run(); err != nil {
//			panic(err)
//		}
//	}
//
//	func App() runetui.Component {
//		return runetui.Box(
//			runetui.Text("Hello, RuneTUI!"),
//			runetui.Style{
//				Border:      true,
//				BorderColor: runetui.ColorCyan,
//				Padding:     1,
//			},
//		)
//	}
//
// # Architecture
//
// RuneTUI components are functions that return a Component interface. The framework
// handles reconciliation, layout calculation, and rendering to the terminal through
// the underlying Bubble Tea runtime.
//
// Components can be composed to build complex UIs:
//
//	func Dashboard() runetui.Component {
//		return runetui.Box(
//			runetui.VStack(
//				Header(),
//				runetui.HStack(
//					Sidebar(),
//					MainContent(),
//				),
//				Footer(),
//			),
//		)
//	}
//
// # Use Cases
//
// RuneTUI is designed for:
//   - CLI tools with rich interactive interfaces
//   - Development tools and dashboards
//   - Agent-based workflows with streaming output
//   - Build tools and task runners
//   - Any scenario where you want Ink-like DX in Go
//
// For lower-level control or simpler UIs, consider using Bubble Tea directly.
//
// # Related Projects
//
//   - Bubble Tea: The Elm Architecture for Go terminals
//   - Lip Gloss: Styles for terminal output
//   - Bubbles: Common TUI components for Bubble Tea
//   - Ink: React for interactive command-line apps (Node.js)
//   - Ratatui: TUI framework for Rust
package runetui
