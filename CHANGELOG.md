# Changelog

All notable changes to RuneTUI will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-01-14

### Added

**Core Components:**
- `Text` component with rich styling (bold, italic, underline, strikethrough, colors)
- `Box` component with borders, padding, margin, and flex properties
- `VStack` and `HStack` convenience wrappers for vertical/horizontal layouts
- `Spacer` component for fixed and flexible spacing
- `Static` component for efficient log/streaming output (no flicker)

**Layout System:**
- Flexbox-inspired layout engine with flex-grow, flex-shrink support
- Direction control (Column/Row)
- Alignment options (start, center, end, stretch)
- Justification options (start, center, end, space-between, space-around)
- Gap, padding, and margin support
- Border styles (single, double, rounded, thick, hidden)

**Testing Infrastructure:**
- `testing.RenderToString()` - render components without a terminal
- `testing.AssertSnapshot()` - golden file snapshot testing with `-update` flag
- `testing.NewTestApp()` - interactive test wrapper for simulating terminal events
- Assertion helpers (AssertHasANSICodes, AssertContainsText, AssertWidth, AssertHeight)
- Behavioral testing pattern with golden files for visual regression testing

**Bubble Tea Integration:**
- Full Elm Architecture support via Bubble Tea adapter
- Window size handling
- Keyboard event handling (Ctrl+C quit by default)
- Context-aware Run and RunContext methods

**Examples:**
- Hello World example with tests
- Streaming output example with tests
- Examples serve as living documentation

**Documentation:**
- Comprehensive README with quick start guide
- Testing utilities documentation in `testing/README.md`
- Architecture Decision Records (ADR.md)
- State management patterns documented
- Implementation plans and behavioral testing guides

### Testing

- 229 tests with ~100% code coverage
- All components include unit tests
- Examples include integration tests
- Golden file tests for visual regression testing
- Behavioral tests verify actual output, not just code presence

### Technical

- Built on [Bubble Tea](https://github.com/charmbracelet/bubbletea) for terminal rendering
- Uses [Lipgloss](https://github.com/charmbracelet/lipgloss) for styling
- Pure Go implementation
- No external dependencies beyond Charm libraries
- Follows TDD principles (baby steps, failing tests first)

---

## Release Philosophy

RuneTUI follows semantic versioning and maintains high test coverage. Each release includes:

- Full test coverage (~100%)
- Living documentation through tested examples
- Clear migration guides for breaking changes
- Architecture decision records for major choices

Future releases will focus on:
- Additional component types (Input, List, Table, Spinner)
- Performance optimizations
- Advanced layout features
- Theming system

[0.1.0]: https://github.com/runetui/runetui/releases/tag/v0.1.0
