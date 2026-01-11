# RuneTUI Implementation Plan

## Executive Summary

This plan outlines the implementation of RuneTUI v0.1, a declarative TUI framework for Go inspired by Ink (React for terminals). The approach focuses on **architecture validation** through incremental TDD implementation of core components, custom flexbox layout engine, Bubble Tea integration, and static zones for agent/logging use cases.

**Key Architectural Decisions:**
- **Component Model:** Stateless, pure functions with external state (Elm-like)
- **Layout Engine:** Custom flexbox-inspired algorithm, delegating styling to Lipgloss
- **Reconciliation:** Simple rebuild (no virtual tree diffing) for v0.1
- **Bubble Tea Integration:** Adapter pattern wrapping RuneTUI components as `tea.Model`
- **Testing:** TDD from day 1, snapshot testing, 100% coverage goal
- **Structure:** Flat library structure (not hexagonal) - UI frameworks have no business/infrastructure split

---

## Phase 1: Foundation - Component System

**Goal:** Establish core Component interface and basic types with TDD workflow.

### 1.1 Core Types and Component Interface

**File:** [component.go](../component.go)

**What to implement:**
- `Component` interface (Render, Children, Key, Measure methods)
- `ComponentFunc` type
- `Props` interface (marker)
- `Layout` struct (X, Y, Width, Height)
- `Size` struct (Width, Height, Min/Max constraints)

**Validation criteria:**
- All types have comprehensive unit tests
- Test coverage: 100% for type definitions
- Clear godoc comments for all public APIs

**Critical decisions to validate:**
- Is the `Component` interface minimal enough yet flexible?
- Does `Measure(availableWidth, availableHeight int) Size` provide sufficient information for layout?
- Is `Render(layout Layout) string` the right signature for rendering?

### 1.2 Shared Type Definitions

**File:** [types.go](../types.go)

**What to implement:**
- `Direction` enum (Column, Row)
- `Dimension` interface (Auto, Fixed, Percent implementations)
- `Spacing` struct (Top, Right, Bottom, Left)
- `Color` type (string alias or struct?)
- `BorderStyle` enum (None, Single, Double, Rounded, etc.)
- `Align` enum (Start, Center, End, Stretch)
- `Justify` enum (Start, Center, End, SpaceBetween, SpaceAround)
- `WrapMode` enum (None, Word, Char, Truncate)
- `TextAlign` enum (Left, Center, Right)

**Validation criteria:**
- Type-safe enums (iota-based constants)
- Clear constructors for Dimension types
- Spacing helper functions (e.g., `SpacingAll(n int)`, `SpacingVertical(v int)`, etc.)
- Tests for all type constructors and methods

**Critical decisions to validate:**
- Should `Color` be a string (hex/named) or a struct with R, G, B fields?
- Do we need separate `BorderWidth` or is `BorderStyle` sufficient?

### 1.3 Text Component (Simplest Leaf)

**File:** [text.go](../text.go), [text_test.go](../text_test.go)

**TDD workflow:**
1. Write failing test: basic text rendering
2. Implement minimal `Text` component
3. Write failing test: text with color styling
4. Add Lipgloss integration for colors
5. Write failing test: text with wrap/truncate
6. Implement text wrapping logic
7. Continue for bold, italic, underline, alignment

**What to implement:**
- `TextProps` struct (Color, Background, Bold, Italic, Underline, Strikethrough, Wrap, Align, Key)
- `Text(content string, props ...TextProps) Component` constructor
- `text` struct (private implementation)
- Implement `Render(layout Layout) string` using Lipgloss
- Implement `Measure(availableWidth, availableHeight int) Size`

**Dependencies:**
- `github.com/charmbracelet/lipgloss` for text styling

**Validation criteria:**
- Can render plain text
- Can apply colors (foreground, background)
- Can apply text styles (bold, italic, underline)
- Can wrap text to fit width
- Can truncate text with ellipsis
- Can align text (left, center, right)
- Test coverage: 100%

**Critical decisions to validate:**
- Is `Text(content string, props ...TextProps)` ergonomic? Alternative: `Text(props TextProps, content string)`
- How do we handle multi-line text in `Measure`?
- Should text wrapping be eager (in Measure) or lazy (in Render)?

### 1.4 Box Component (Simplest Container)

**File:** [box.go](../box.go), [box_test.go](../box_test.go)

**TDD workflow:**
1. Write failing test: empty box rendering
2. Implement minimal `Box` component structure
3. Write failing test: box with single text child
4. Implement children rendering (no layout yet, just stacking)
5. Write failing test: box with border
6. Add Lipgloss border rendering
7. Write failing test: box with padding
8. Implement padding in render
9. Continue for margin, background, etc.

**What to implement:**
- `BoxProps` struct (Direction, Width, Height, Min/Max, FlexGrow/Shrink, AlignItems, JustifyContent, Padding, Margin, Gap, Border, BorderColor, Background, IsStatic, Key)
- `Box(props BoxProps, children ...Component) Component` constructor
- `box` struct (private implementation)
- Implement `Render(layout Layout) string` (basic stacking for now)
- Implement `Measure(availableWidth, availableHeight int) Size` (sum of children for now)
- Implement `Children() []Component`

**Validation criteria:**
- Can render empty box
- Can render box with children (stacked vertically)
- Can render border using Lipgloss
- Can apply padding (reduces available space for children)
- Can apply margin (affects parent layout)
- Can apply background color
- Test coverage: 100%

**Critical decisions to validate:**
- Should `Box` always require `BoxProps` or have a default? Alternative: `Box(children ...Component)` with optional `WithProps(BoxProps)`
- How does `Measure` handle percentage-based dimensions?
- Should border and padding be part of measured size or added afterward?

---

## Phase 2: Layout Engine

**Goal:** Implement flexbox-inspired layout calculation system.

### 2.1 Measure System

**File:** [measure.go](../measure.go), [measure_test.go](../measure_test.go)

**TDD workflow:**
1. Write failing test: measure fixed-size text
2. Implement intrinsic text measurement
3. Write failing test: measure auto-size box with children
4. Implement child measurement aggregation
5. Write failing test: measure box with min/max constraints
6. Implement constraint clamping
7. Continue for padding, margin, border calculations

**What to implement:**
- Helper functions for measuring components
- `measureText(content string, wrap WrapMode, availableWidth int) Size`
- `measureBox(props BoxProps, children []Component, availableWidth, availableHeight int) Size`
- Constraint resolution (min/max, fixed, percentage)
- Padding/margin/border size calculations

**Validation criteria:**
- Text measures correctly based on content and wrap mode
- Box measures as sum of children in Column direction
- Box measures as max of children in Row direction
- Constraints are applied correctly (min, max, fixed)
- Padding/margin/border add to measured size
- Test coverage: 100%

**Critical decisions to validate:**
- How do we handle percentage dimensions when parent size is unknown?
- Should measure be memoized/cached?
- How do we measure text without rendering (Unicode width calculation)?

### 2.2 Basic Layout Algorithm

**File:** [layout.go](../layout.go), [layout_test.go](../layout_test.go)

**TDD workflow:**
1. Write failing test: layout single text component
2. Implement basic position assignment
3. Write failing test: layout box with children in Column direction
4. Implement vertical stacking
5. Write failing test: layout box with children in Row direction
6. Implement horizontal placement
7. Write failing test: layout with padding
8. Adjust available space for padding
9. Continue for margin, alignment, justification

**What to implement:**
- `LayoutEngine` struct (terminalWidth, terminalHeight)
- `NewLayoutEngine(width, height int) *LayoutEngine`
- `CalculateLayout(root Component) *LayoutTree`
- `measureAndLayout(component Component, availableWidth, availableHeight, x, y int) *LayoutTree`
- `LayoutTree` struct (Component, Layout, Children)

**Validation criteria:**
- Single component positioned at (0, 0)
- Column layout stacks children vertically (Y increases)
- Row layout places children horizontally (X increases)
- Padding reduces available space for children
- Margin adds space around component
- Parent dimensions constrain children
- Test coverage: 100%

**Critical decisions to validate:**
- Should layout be two-pass (measure + position) or single-pass?
- How do we handle overflow (children larger than parent)?
- Should we use absolute coordinates or relative offsets?

### 2.3 Flexbox Properties Implementation

**File:** [flex.go](../flex.go), [flex_test.go](../flex_test.go)

**TDD workflow:**
1. Write failing test: flex-grow distributes extra space
2. Implement flex-grow calculation
3. Write failing test: flex-shrink when constrained
4. Implement flex-shrink calculation
5. Write failing test: align-items center
6. Implement cross-axis alignment
7. Write failing test: justify-content space-between
8. Implement main-axis justification
9. Continue for all alignment combinations

**What to implement:**
- `calculateFlexGrow(children []FlexChild, extraSpace int) []int` - distribute extra space
- `calculateFlexShrink(children []FlexChild, deficit int) []int` - shrink when constrained
- `alignItems(children []*LayoutTree, props BoxProps, crossSize int)` - cross-axis alignment
- `justifyContent(children []*LayoutTree, props BoxProps, mainSize int)` - main-axis distribution
- `FlexChild` struct (Component, Size, FlexGrow, FlexShrink)

**Validation criteria:**
- Flex-grow distributes extra space proportionally
- Flex-shrink reduces size when constrained
- AlignItems positions children on cross-axis (Start, Center, End, Stretch)
- JustifyContent distributes children on main-axis (Start, Center, End, SpaceBetween, SpaceAround)
- Gap adds space between children
- Test coverage: 100% for all flex scenarios

**Critical decisions to validate:**
- How do we handle flex-basis (intrinsic size before grow/shrink)?
- Should we support flex-wrap or keep it simple?
- How do nested flex containers interact?

---

## Phase 3: Bubble Tea Integration

**Goal:** Create adapter layer to run RuneTUI components in Bubble Tea runtime.

### 3.1 Basic Adapter Implementation

**File:** [adapter.go](../adapter.go), [adapter_test.go](../adapter_test.go)

**TDD workflow:**
1. Write failing test: App.Run() initializes Bubble Tea program
2. Implement basic App struct and Run method
3. Write failing test: model.View() renders component tree
4. Implement View using LayoutEngine
5. Write failing test: model.Update() handles WindowSizeMsg
6. Implement terminal resize handling
7. Write failing test: model.Update() handles KeyMsg (Ctrl+C to quit)
8. Implement basic keyboard input
9. Continue for other Bubble Tea messages

**What to implement:**
- `App` struct (rootFunc, model, program, staticMgr, layoutEngine)
- `New(rootFunc ComponentFunc, opts ...AppOption) *App`
- `Run() error` - starts Bubble Tea program
- `RunContext(ctx context.Context) error` - starts with context for graceful shutdown
- `AppOption` type and common options
- `model` struct (implements `tea.Model`)
- `Init() tea.Cmd`
- `Update(msg tea.Msg) (tea.Model, tea.Cmd)`
- `View() string`

**Dependencies:**
- `github.com/charmbracelet/bubbletea`

**Note:** Use `context.Context` for operations that may need cancellation (e.g., `RunContext`).

**Validation criteria:**
- Can create and run a RuneTUI app
- Terminal resize updates layout
- Ctrl+C quits application
- Component tree renders correctly in View
- Test coverage: 100% (using Bubble Tea's testing utilities)

**Critical decisions to validate:**
- Should `rootFunc` be called every View or cached?
- How do we pass Bubble Tea messages to components?
- Should we expose `tea.Cmd` directly or wrap it?

### 3.2 Message Handling and State Management

**File:** [messages.go](../messages.go), [messages_test.go](../messages_test.go)

**TDD workflow:**
1. Write failing test: custom message type propagates to component
2. Implement message passing mechanism
3. Write failing test: state update triggers re-render
4. Implement state management pattern
5. Write failing test: multiple components share state
6. Document state sharing patterns
7. Continue for async commands and effects

**What to implement:**
- Common message types (if any)
- State management patterns (document how users should manage state)
- Examples of Update patterns
- Helper functions for message handling

**Validation criteria:**
- Clear documentation on state management
- Example showing counter app (classic state demo)
- Example showing form with multiple inputs
- Test coverage: 100% for message utilities

**Critical decisions to validate:**
- Do we need custom message types or just use Bubble Tea's?
- How do components communicate with each other?
- Should we provide a state management helper or leave it to users?

### 3.3 First Working Example

**File:** [examples/basic/main.go](../examples/basic/main.go)

**What to implement:**
- Simple "Hello World" application
- Uses Box and Text components
- Demonstrates border and styling
- Handles Ctrl+C to quit

**Validation criteria:**
- Example runs without errors
- Displays styled text in a bordered box
- Can quit cleanly
- Code is clear and documented

**Critical decisions to validate:**
- Is the API ergonomic for simple use cases?
- Are error messages helpful?
- Is the learning curve acceptable?

---

## Phase 4: Static Zones

**Goal:** Implement static/dynamic split for accumulating logs and agent output.

### 4.1 Static Manager

**File:** [static_manager.go](../static_manager.go), [static_manager_test.go](../static_manager_test.go)

**TDD workflow:**
1. Write failing test: append static content
2. Implement basic buffer and append
3. Write failing test: key-based tracking prevents duplication
4. Implement key tracking
5. Write failing test: render static buffer
6. Implement RenderStatic
7. Write failing test: clear static buffer
8. Implement Clear
9. Continue for edge cases (empty buffer, duplicate keys, etc.)

**What to implement:**
- `StaticManager` struct (staticBuffer, staticKeys)
- `NewStaticManager() *StaticManager`
- `AppendStatic(key string, content []string) int` - append new lines
- `RenderStatic() string` - return accumulated content
- `Clear()` - reset buffer

**Validation criteria:**
- Can accumulate lines of text
- Tracks last rendered index per key
- Only renders new lines on subsequent calls
- Can clear buffer
- Test coverage: 100%

**Critical decisions to validate:**
- Should static content be line-based or string-based?
- How do we handle very large static buffers (memory concerns)?
- Should we support removing items from static buffer?

### 4.2 Static Component

**File:** [static.go](../static.go), [static_test.go](../static_test.go)

**TDD workflow:**
1. Write failing test: static component accumulates items
2. Implement Static component structure
3. Write failing test: static items don't re-render
4. Integrate with StaticManager
5. Write failing test: static + dynamic zones combine correctly
6. Implement rendering split
7. Continue for edge cases

**What to implement:**
- `StaticProps` struct (Key)
- `Static(props StaticProps, itemsFunc func() []Component) Component`
- `static` struct (private implementation)
- Integration with StaticManager in adapter
- Rendering logic to combine static + dynamic

**Validation criteria:**
- Static items accumulate above dynamic content
- Static items don't re-render on every frame
- New static items append correctly
- Static + dynamic combination looks correct
- Test coverage: 100%

**Critical decisions to validate:**
- Should `itemsFunc` return all items or just new items?
- How do we handle terminal resize with static content?
- Should static content scroll or be truncated?

### 4.3 Streaming Example

**File:** [examples/streaming/main.go](../examples/streaming/main.go)

**What to implement:**
- Application with static log zone (top)
- Dynamic status zone (bottom)
- Simulated streaming logs (timer-based)
- Demonstrates key use case for agents

**Validation criteria:**
- Logs accumulate in static zone
- Status updates in dynamic zone
- No flickering or re-rendering of static content
- Handles terminal resize gracefully
- Code demonstrates best practices

**Critical decisions to validate:**
- Is the static/dynamic API intuitive?
- Does it handle real-world streaming scenarios?
- Are there any visual artifacts or performance issues?

---

## Phase 5: Helpers and Ergonomics

**Goal:** Add convenience components and improve developer experience.

### 5.1 Stack Helpers

**File:** [stack.go](../stack.go), [stack_test.go](../stack_test.go)

**TDD workflow:**
1. Write failing test: VStack renders children vertically
2. Implement VStack as Box wrapper
3. Write failing test: HStack renders children horizontally
4. Implement HStack as Box wrapper
5. Write failing test: stack with gap
6. Add gap support
7. Continue for alignment and other props

**What to implement:**
- `VStack(children ...Component) Component` - vertical stack
- `HStack(children ...Component) Component` - horizontal stack
- Overload variants: `VStackWithProps(props StackProps, children ...Component)`
- `StackProps` struct (simplified BoxProps)

**Validation criteria:**
- VStack and HStack work as expected
- Props like gap, alignment, padding work
- API is more ergonomic than raw Box
- Test coverage: 100%

**Critical decisions to validate:**
- Should stacks be separate components or Box wrappers?
- What subset of BoxProps should StackProps expose?

### 5.2 Spacer Component

**File:** [spacer.go](../spacer.go), [spacer_test.go](../spacer_test.go)

**TDD workflow:**
1. Write failing test: Spacer with fixed size
2. Implement Spacer component
3. Write failing test: Spacer with flex-grow
4. Add flex support
5. Continue for different use cases

**What to implement:**
- `Spacer(size int) Component` - fixed-size spacer
- `FlexSpacer() Component` - flexible spacer (flex-grow: 1)

**Validation criteria:**
- Fixed spacer creates exact space
- Flex spacer fills available space
- Works in both vertical and horizontal layouts
- Test coverage: 100%

**Critical decisions to validate:**
- Is Spacer a separate component or just an empty Box?
- Do we need both fixed and flex variants?

### 5.3 Testing Utilities

**File:** [testing/testing.go](../testing/testing.go), [testing/testing_test.go](../testing/testing_test.go)

**What to implement:**
- `RenderToString(rootFunc ComponentFunc, width, height int) string` - render without terminal
- `AssertSnapshot(t *testing.T, name string, output string)` - golden file comparison
- `NewTestApp(rootFunc ComponentFunc) *TestApp` - test wrapper for App
- `TestApp.Resize(width, height int)` - simulate resize
- `TestApp.SendKey(key string)` - simulate keyboard input
- `TestApp.View() string` - get current view

**Validation criteria:**
- Can render components without Bubble Tea
- Snapshot testing works (golden files)
- Can simulate user interactions
- Test utilities are well-documented
- Test coverage: 100%

**Critical decisions to validate:**
- Should test utilities be in a separate package?
- How do we handle snapshot updates (flag to regenerate)?
- Should we integrate with testify or keep it stdlib-only?

---

## Testing Strategy

### TDD Workflow (Project Rule)

**Process for every feature:**
1. Write ONE failing test
2. Run test to verify it fails: `make test`
3. Write minimal code to make it pass
4. Run test to verify it passes: `make test`
5. Refactor if needed
6. Run test again to ensure still passing
7. Repeat for next test

**Never:**
- Write multiple tests at once
- Write code before writing test
- Skip running tests after changes

### Test Organization

```
runetui/
├── component_test.go
├── types_test.go
├── text_test.go
├── box_test.go
├── measure_test.go
├── layout_test.go
├── flex_test.go
├── adapter_test.go
├── messages_test.go
├── static_manager_test.go
├── static_test.go
├── stack_test.go
├── spacer_test.go
├── internal/              # (optional) Private helpers not part of public API
│   └── ansi/              # Example: ANSI escape utilities
└── testing/
    └── testing_test.go
```

**Note on `internal/`:** Only use for truly private implementation details that should not be part of the public API. Most code stays at the root as public API.

### Test Coverage Goals

- **Unit tests:** 100% coverage for all packages
- **Integration tests:** Full App lifecycle (init, update, view, quit)
- **Snapshot tests:** All example applications
- **No E2E tests:** Not needed for v0.1 (TUI framework doesn't need browser-like E2E)

### Test Naming Convention

```go
func TestComponentName_Scenario_ExpectedBehavior(t *testing.T)
```

Examples:
- `TestText_WithBoldStyle_RendersBoldText`
- `TestBox_WithPadding_ReducesAvailableSpace`
- `TestLayoutEngine_ColumnDirection_StacksChildrenVertically`
- `TestFlexGrow_WithExtraSpace_DistributesProportionally`

---

## Critical Files

### Package Naming

- **No `utils` or `common` packages**: Name packages by their semantic purpose
- Root package `runetui`: All public API (components, types, adapter)
- `runetui/testing`: Test utilities for users of the framework
- `runetui/internal/*`: Private implementation details (if needed)

### Core Implementation Files (in dependency order)

1. [types.go](../types.go) - Shared type definitions (Direction, Dimension, Spacing, Color, etc.)
2. [component.go](../component.go) - Component interface and fundamental types
3. [text.go](../text.go) - Text leaf component
4. [box.go](../box.go) - Box container component
5. [measure.go](../measure.go) - Size measurement system
6. [layout.go](../layout.go) - Layout engine and algorithm
7. [flex.go](../flex.go) - Flexbox properties implementation
8. [static_manager.go](../static_manager.go) - Static content accumulation
9. [static.go](../static.go) - Static zone component
10. [adapter.go](../adapter.go) - Bubble Tea integration
11. [messages.go](../messages.go) - Message handling and state patterns
12. [stack.go](../stack.go) - VStack and HStack helpers
13. [spacer.go](../spacer.go) - Spacer component

### Test Files (one per implementation file)

- [types_test.go](../types_test.go)
- [component_test.go](../component_test.go)
- [text_test.go](../text_test.go)
- [box_test.go](../box_test.go)
- [measure_test.go](../measure_test.go)
- [layout_test.go](../layout_test.go)
- [flex_test.go](../flex_test.go)
- [static_manager_test.go](../static_manager_test.go)
- [static_test.go](../static_test.go)
- [adapter_test.go](../adapter_test.go)
- [messages_test.go](../messages_test.go)
- [stack_test.go](../stack_test.go)
- [spacer_test.go](../spacer_test.go)

### Testing Utilities

- [testing/testing.go](../testing/testing.go) - Test helpers and snapshot utilities
- [testing/testing_test.go](../testing/testing_test.go) - Tests for test utilities

### Example Applications

- [examples/basic/main.go](../examples/basic/main.go) - Hello World
- [examples/streaming/main.go](../examples/streaming/main.go) - Static logs + dynamic status

### Documentation (as features are completed)

- Update [doc.go](../doc.go) - Add examples as they work
- Update [README.md](../README.md) - Check off completed features

---

## Dependencies to Add

Update [go.mod](../go.mod):

```go
require (
    github.com/charmbracelet/bubbletea v0.25.0
    github.com/charmbracelet/lipgloss v0.9.1
)
```

Run after updating:
```bash
go mod tidy
go mod download
```

---

## Validation Checkpoints

### Checkpoint 1: Foundation Complete
**After completing Phase 1:**
- [ ] Can create Text and Box components
- [ ] Components have well-defined interfaces
- [ ] All tests pass: `make test`
- [ ] Coverage at 100%: `make test-coverage`
- [ ] Code formatted: `make fmt`
- [ ] Linter passes: `make lint`
- [ ] Validation passes: `make validate`

**Architectural questions to answer:**
- Is the Component interface flexible enough?
- Is the API ergonomic?
- Are there any obvious design flaws?

### Checkpoint 2: Layout Engine Complete
**After completing Phase 2:**
- [ ] Layout engine calculates positions correctly
- [ ] Flexbox properties work as expected
- [ ] Complex nested layouts render correctly
- [ ] All tests pass: `make test`
- [ ] Coverage at 100%: `make test-coverage`
- [ ] Validation passes: `make validate`

**Architectural questions to answer:**
- Is the layout algorithm efficient enough?
- Does it handle edge cases (overflow, constraints)?
- Are there any performance issues?

### Checkpoint 3: Bubble Tea Integration Complete
**After completing Phase 3:**
- [ ] Can run a working RuneTUI application
- [ ] Basic example works end-to-end
- [ ] Terminal resize handled correctly
- [ ] Keyboard input works
- [ ] All tests pass: `make test`
- [ ] Coverage at 100%: `make test-coverage`
- [ ] Validation passes: `make validate`

**Architectural questions to answer:**
- Is the adapter layer clean and maintainable?
- Is the state management pattern clear?
- Are there any Bubble Tea integration issues?

### Checkpoint 4: Static Zones Complete
**After completing Phase 4:**
- [ ] Static zones accumulate content correctly
- [ ] No flickering or re-rendering issues
- [ ] Streaming example works smoothly
- [ ] Static + dynamic split looks correct
- [ ] All tests pass: `make test`
- [ ] Coverage at 100%: `make test-coverage`
- [ ] Validation passes: `make validate`

**Architectural questions to answer:**
- Is the static/dynamic API intuitive?
- Does it handle real-world agent scenarios?
- Are there memory or performance concerns?

### Checkpoint 5: v0.1 Release Ready
**After completing Phase 5:**
- [ ] All helper components work
- [ ] Testing utilities functional
- [ ] All examples work
- [ ] Documentation complete
- [ ] All tests pass: `make test`
- [ ] Coverage at 100%: `make test-coverage`
- [ ] Validation passes: `make validate`
- [ ] API feels ergonomic and consistent

**Release criteria:**
- No known bugs
- API is stable enough for early adopters
- Documentation is comprehensive
- Examples demonstrate key use cases

---

## Pre-Commit Checklist (MANDATORY)

Before EVERY commit:
1. [ ] Run `make validate`
2. [ ] Fix ALL errors (zero tolerance)
3. [ ] Run `make validate` again
4. [ ] Only commit when passing

**Never commit without passing validation.**

---

## Implementation Order Summary

**Phase 1: Foundation (Weeks 1-2)**
1. types.go + tests
2. component.go + tests
3. text.go + tests
4. box.go + tests

**Phase 2: Layout Engine (Weeks 3-4)**
5. measure.go + tests
6. layout.go + tests (column direction)
7. layout.go + tests (row direction)
8. flex.go + tests (grow/shrink/align/justify)

**Phase 3: Bubble Tea Integration (Week 5)**
9. adapter.go + tests
10. messages.go + tests
11. examples/basic/main.go (first working app)

**Phase 4: Static Zones (Week 6)**
12. static_manager.go + tests
13. static.go + tests
14. examples/streaming/main.go (agent use case)

**Phase 5: Helpers (Week 7)**
15. stack.go + tests
16. spacer.go + tests
17. testing/testing.go + tests
18. Update documentation

---

## Error Handling Guidelines

Based on team conventions, follow these patterns:

### Domain Errors

Export business errors as package-level variables:

```go
// errors.go
var (
    ErrInvalidDimension = errors.New("invalid dimension value")
    ErrLayoutOverflow   = errors.New("children exceed available space")
    ErrComponentNil     = errors.New("component cannot be nil")
)
```

### Error Wrapping

Wrap errors with context to preserve the chain:

```go
// Good: adds context without losing original cause
func (e *LayoutEngine) CalculateLayout(root Component) (*LayoutTree, error) {
    size, err := root.Measure(e.width, e.height)
    if err != nil {
        return nil, fmt.Errorf("measuring root component: %w", err)
    }
    // ...
}
```

### Configuration Boundaries

- **Core components** (`text.go`, `box.go`, `layout.go`): Never read environment variables or external config
- **Adapter layer** (`adapter.go`): May read config and inject values into components via constructors
- **Examples** (`examples/*/main.go`): Read config/env vars and wire dependencies

This keeps components pure and testable.

---

## Open Architectural Questions

### Questions to answer during implementation:

1. **Component Lifecycle:**
   - Do we need lifecycle methods (mount, unmount, update)?
   - Or is pure render enough?

2. **Performance Optimization:**
   - Do we need component memoization (shouldComponentUpdate)?
   - Or is simple rebuild fast enough?

3. **Error Handling:** *(See Error Handling Guidelines section)*
   - Export domain errors as variables for type checking
   - Wrap errors with context using `fmt.Errorf("context: %w", err)`
   - Render methods: return error for recoverable issues, panic for programmer bugs

4. **Event Handling:**
   - How do components handle keyboard/mouse input?
   - Through Bubble Tea messages only?

5. **Accessibility:**
   - Do we need ARIA-like annotations for screen readers?
   - Or is it too early?

6. **Theming:**
   - Do we need a theme system (like Lipgloss themes)?
   - Or leave it to users?

7. **State Management:**
   - Do we provide helpers (like React Context)?
   - Or keep it pure Bubble Tea?

**Approach:** Start simple, add complexity only when needed. Validate with real examples.

---

## Success Metrics

### Technical Metrics
- Test coverage: 100%
- All validation checks passing
- Zero known bugs
- Clean, idiomatic Go code

### API Design Metrics
- Can build Hello World in < 10 lines
- Can build complex layouts without fighting the API
- Clear error messages
- Minimal boilerplate

### Use Case Validation
- Static/dynamic split works for agent logs
- Flexbox layout handles complex UIs
- Bubble Tea integration is seamless
- Examples are clear and helpful

---

## Next Steps After Plan Approval

1. **Set up dependencies:**
   ```bash
   go get github.com/charmbracelet/bubbletea@v0.25.0
   go get github.com/charmbracelet/lipgloss@v0.9.1
   ```

2. **Start with first test:**
   - Create `types_test.go`
   - Write first failing test for `Direction` enum
   - Implement types.go to make it pass
   - Continue TDD workflow

3. **Commit after each passing test:**
   - Run `make validate`
   - Commit with clear message
   - Push to track progress

4. **Regular architecture review:**
   - After each phase checkpoint
   - Review architectural questions
   - Adjust plan if needed

---

## End of Plan

This plan provides a comprehensive, test-driven roadmap for implementing RuneTUI v0.1. The focus is on validating architectural decisions through incremental implementation, with strict TDD discipline and 100% test coverage.

**Key principles:**
- Baby steps (one test at a time)
- TDD workflow (test before code)
- Architecture validation (answer questions through implementation)
- Quality over speed (100% coverage, zero errors)
- Idiomatic Go (simple, clear, composable)
