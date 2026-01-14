# RuneTUI

A declarative TUI framework for Go, inspired by Ink.

> ⚠️ **Early Access**: RuneTUI v0.1 is feature-complete and ready for early adopters. APIs may evolve based on feedback.

## Vision

RuneTUI brings React/Ink-style declarative UI to Go's terminal ecosystem. Built on top of [Bubble Tea](https://github.com/charmbracelet/bubbletea), it provides:

- 🎨 Declarative, component-based API
- 📦 Flexbox-inspired layouts  
- 🔄 Separation of static/dynamic zones
- 🧪 First-class testing support
- 🤖 Agent-friendly patterns (logs, panels, streaming)

## Quick Example

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

## Status

✅ **v0.1 - Core features complete:**

- [x] Core component system (Box, Text, VStack, HStack, Spacer, Static)
- [x] Flexbox-inspired layout engine (flex-grow, flex-shrink, alignment, justification)
- [x] Bubble Tea integration with adapter layer
- [x] Static zones for efficient log rendering (no flicker)
- [x] Testing utilities (RenderToString, snapshot testing, TestApp)
- [x] Example applications with tests
- [x] State management patterns documented
- [x] 229 tests with ~100% coverage

🚧 **Future (post-v0.1):**

- [ ] More component types (Input, List, Table, Spinner)
- [ ] Performance optimizations (memoization, virtual rendering)
- [ ] Advanced layout features (wrapping, scrolling)
- [ ] Theming system

## Features

- **Declarative Components**: Build UIs with composable, functional components
- **Flexbox Layouts**: Column/Row directions with flex properties, alignment, and spacing
- **Static Zones**: Efficient rendering for logs and streaming output (no flicker)
- **Rich Text Styling**: Colors, bold, italic, alignment, and text wrapping
- **Testing Support**: Snapshot testing and component rendering utilities
- **Bubble Tea Integration**: Built on the proven Elm Architecture pattern

### Available Components

- **Text** - Rich text with styling (color, bold, italic, underline, wrapping, alignment)
- **Box** - Flexible container with borders, padding, margin, and flex properties
- **VStack / HStack** - Convenient vertical/horizontal stack layouts
- **Spacer** - Fixed or flexible spacing between components
- **Static** - Accumulating zone for logs and streaming output (efficient, no re-render)

## Why RuneTUI?

Go is increasingly used for CLI tools and agent-based workflows, but building rich TUIs remains verbose. RuneTUI aims to provide the same DX that Ink brings to the Node.js ecosystem, leveraging Go's mature TUI libraries while offering a higher-level, declarative API.

Perfect for:
- CLI tools with rich interactive interfaces
- AI agent workflows with streaming output
- Development tools and dashboards
- Build tools and task runners

### Inspired by

- [Ink](https://github.com/vadimdemedes/ink) - React for CLIs
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - The Elm Architecture for Go
- [Ratatui](https://ratatui.rs/) - Rust TUI framework

## Testing

RuneTUI includes comprehensive testing utilities:

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

### Testing Examples

Examples in the `examples/` directory include tests that verify they render correctly:

```bash
# Run all tests (including examples)
make test

# Run only example tests
make test-examples

# Update snapshots when output changes intentionally
make test-examples-update
```

This ensures examples stay working and serve as living documentation.

## Contributing

RuneTUI v0.1 is now released! Contributions, ideas, and feedback are welcome! Please open an issue to discuss major changes before submitting PRs.

See [CHANGELOG.md](CHANGELOG.md) for release history.

## License

MIT
