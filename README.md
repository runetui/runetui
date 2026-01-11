# RuneTUI

A declarative TUI framework for Go, inspired by Ink.

> ⚠️ **Work in Progress**: RuneTUI is in early development. APIs will change.

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

✅ **Core features implemented:**

- [x] Core component system (Box, Text, Stack)
- [x] Flexbox-inspired layout engine
- [x] Bubble Tea integration
- [x] Static zones for efficient log rendering
- [x] Testing utilities with snapshot support
- [x] Example applications

🚧 **In progress:**

- [ ] State management patterns
- [ ] More component types
- [ ] Enhanced documentation
- [ ] Performance optimizations

## Features

- **Declarative Components**: Build UIs with composable, functional components
- **Flexbox Layouts**: Column/Row directions with flex properties, alignment, and spacing
- **Static Zones**: Efficient rendering for logs and streaming output (no flicker)
- **Rich Text Styling**: Colors, bold, italic, alignment, and text wrapping
- **Testing Support**: Snapshot testing and component rendering utilities
- **Bubble Tea Integration**: Built on the proven Elm Architecture pattern

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

## Testing Examples

Examples in the `examples/` directory include tests that verify they render correctly:

```bash
# Run example tests
make test-examples

# Update snapshots when output changes intentionally
make test-examples-update
```

This ensures examples stay working and serve as living documentation.

## Contributing

RuneTUI is in early development. Contributions, ideas, and feedback are welcome! Please open an issue to discuss major changes.

## License

MIT
