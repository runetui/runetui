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

import "github.com/runetui/runetui"

func App() runetui.Component {
    return runetui.Box(
        runetui.Text("Hello, RuneTUI!"),
        runetui.Style{Border: true},
    )
}
```

## Status

Currently in design phase. Core goals:

- [ ] Core component system
- [ ] Layout engine
- [ ] Bubble Tea integration
- [ ] Example applications
- [ ] Documentation

## Why RuneTUI?

Go is increasingly used for CLI tools and agent-based workflows, but building rich TUIs remains verbose. RuneTUI aims to provide the same DX that Ink brings to the Node.js ecosystem, leveraging Go's mature TUI libraries while offering a higher-level, declarative API.

### Inspired by

- [Ink](https://github.com/vadimdemedes/ink) - React for CLIs
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - The Elm Architecture for Go
- [Ratatui](https://ratatui.rs/) - Rust TUI framework

## Contributing

RuneTUI is in early development. Contributions, ideas, and feedback are welcome! Please open an issue to discuss major changes.

## License

MIT
