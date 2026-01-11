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
//	import (
//		"log"
//		"github.com/runetui/runetui"
//	)
//
//	func main() {
//		app := runetui.New(func() runetui.Component {
//			return runetui.Box(
//				runetui.BoxProps{
//					Direction: runetui.Column,
//					Padding:   runetui.SpacingAll(2),
//					Border:    runetui.BorderSingle,
//				},
//				runetui.Text("Hello, RuneTUI!", runetui.TextProps{Bold: true}),
//				runetui.Text("Press Ctrl+C to quit"),
//			)
//		})
//
//		if err := app.Run(); err != nil {
//			log.Fatal(err)
//		}
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
//			runetui.BoxProps{Direction: runetui.Column},
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
// # Available Components
//
// Core components:
//   - Box: Container with flexbox-like layout (Column/Row direction)
//   - Text: Text rendering with styling (colors, bold, italic, alignment, wrapping)
//   - VStack/HStack: Convenience wrappers for vertical/horizontal stacks
//   - Static: Accumulates content across renders (ideal for logs and streaming output)
//   - Spacer/FlexSpacer: Space management utilities
//
// # Layout System
//
// RuneTUI uses a flexbox-inspired layout system:
//   - Direction: Column (vertical) or Row (horizontal)
//   - Dimensions: Auto, Fixed, or Percentage-based sizing
//   - Flex properties: FlexGrow and FlexShrink for flexible sizing
//   - Alignment: AlignItems (cross-axis) and JustifyContent (main-axis)
//   - Spacing: Padding, Margin, and Gap support
//   - Borders: Single, Double, or Rounded border styles
//
// # Static vs Dynamic Zones
//
// RuneTUI distinguishes between static and dynamic UI zones:
//   - Static zones (Static component): Content accumulates across renders,
//     old content is not re-rendered (efficient for logs, agent output, streaming)
//   - Dynamic zones (regular components): Content is re-rendered on every frame
//     (perfect for status bars, progress indicators, current state)
//
// # Testing
//
// The testing package provides utilities for testing components:
//   - RenderToString: Render components without a terminal
//   - AssertSnapshot: Snapshot testing with golden files
//   - TestApp: Simulate user interactions and terminal resizes
//
// Example:
//
//	import "github.com/runetui/runetui/testing"
//
//	func TestMyComponent(t *testing.T) {
//		output := testing.RenderToString(func() runetui.Component {
//			return runetui.Text("Hello, World!")
//		}, 80, 24)
//		testing.AssertSnapshot(t, "my_component", output)
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
