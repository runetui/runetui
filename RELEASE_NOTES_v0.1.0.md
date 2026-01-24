# RuneTUI v0.1.0 - Initial Release 🎉

A declarative TUI framework for Go, inspired by Ink.

## What is RuneTUI?

RuneTUI brings React/Ink-style declarative UI to Go's terminal ecosystem. Built on top of [Bubble Tea](https://github.com/charmbracelet/bubbletea), it provides a high-level, component-based API for building rich terminal user interfaces.

Perfect for:
- 🛠️ CLI tools with interactive interfaces
- 🤖 AI agent workflows with streaming output
- 📊 Development tools and dashboards
- 🏗️ Build tools and task runners

## Quick Start

```go
package main

import (
    "log"
    "github.com/runetui/runetui"
)

func main() {
    app := runetui.New(func() runetui.Component {
        return runetui.Box(
            runetui.BoxProps{
                Direction: runetui.Column,
                Padding:   runetui.SpacingAll(2),
                Border:    runetui.BorderSingle,
            },
            runetui.Text("Hello, RuneTUI!", runetui.TextProps{Bold: true}),
            runetui.Text("Press Ctrl+C to quit"),
        )
    })

    if err := app.Run(); err != nil {
        log.Fatal(err)
    }
}
```

## Features

### Core Components

- **Text** - Rich text with styling (color, bold, italic, underline, wrapping, alignment)
- **Box** - Flexible container with borders, padding, margin, and flex properties
- **VStack / HStack** - Convenient vertical/horizontal stack layouts
- **Spacer** - Fixed or flexible spacing between components
- **Static** - Accumulating zone for logs and streaming output (efficient, no flicker)

### Layout System

- 📦 Flexbox-inspired layout engine
- 🔄 Direction control (Column/Row)
- 📐 Flex properties (flex-grow, flex-shrink)
- 🎯 Alignment & justification
- 📏 Gap, padding, margin support
- 🎨 Border styles (single, double, rounded, thick, hidden)

### Testing Infrastructure

```go
// Render components without a terminal
output := testing.RenderToString(rootFunc, 80, 24)

// Snapshot testing with golden files
testing.AssertSnapshot(t, "my-component", output)

// Interactive test wrapper
app := testing.NewTestApp(rootFunc)
app.Resize(100, 30)
view := app.View()
```

- ✅ Golden file snapshot testing with `-update` flag
- ✅ Assertion helpers for behavioral testing
- ✅ Component rendering utilities
- ✅ 229 tests with ~100% coverage

### Bubble Tea Integration

- Full Elm Architecture support
- Window size handling
- Keyboard event handling
- Context-aware Run and RunContext methods

## Technical Highlights

- 🧪 **Test-Driven Development**: Every feature built with TDD (baby steps, failing tests first)
- 📚 **Living Documentation**: Examples with tests serve as living documentation
- 🎯 **100% Coverage**: ~100% test coverage maintained across all components
- 🏗️ **Behavioral Testing**: Golden files verify actual visual output, not just code presence
- 📖 **Documented Decisions**: Architecture Decision Records (ADR.md) track major choices

## Documentation

- [README.md](README.md) - Quick start and overview
- [CHANGELOG.md](CHANGELOG.md) - Detailed changelog
- [testing/README.md](testing/README.md) - Testing utilities documentation
- [ADR.md](ADR.md) - Architecture decision records
- [docs/](docs/) - Implementation plans and guides
- [examples/](examples/) - Working examples with tests

## Installation

```bash
go get github.com/runetui/runetui
```

## What's Next?

Post-v0.1 roadmap:

- 📝 More component types (Input, List, Table, Spinner)
- ⚡ Performance optimizations (memoization, virtual rendering)
- 🎨 Advanced layout features (wrapping, scrolling)
- 🎭 Theming system

## Contributing

Contributions, ideas, and feedback are welcome! Please open an issue to discuss major changes before submitting PRs.

## Built With

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - The Elm Architecture for Go
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions for terminal applications

## License

MIT

---

**Full Changelog**: https://github.com/runetui/runetui/blob/main/CHANGELOG.md
