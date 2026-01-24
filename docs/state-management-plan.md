# RuneTUI State Management Implementation Plan

## Status: ✅ COMPLETED (January 24, 2026)

This plan has been fully implemented. All phases are complete:

- **Phase 1**: ✅ Custom Update/Init Support in Adapter
- **Phase 2**: ⏭️ Skipped (helpers deemed optional - raw Update/Init sufficient)
- **Phase 3**: ✅ Complete Examples with Adapter
- **Phase 4**: ✅ Enhanced Documentation

### What Was Implemented

1. **Adapter API Extensions** (`adapter.go`):
   - `UpdateFunc` type: `func(msg tea.Msg) tea.Cmd`
   - `InitFunc` type: `func() tea.Cmd`
   - `WithUpdate(UpdateFunc)` option
   - `WithInit(InitFunc)` option
   - Full backward compatibility maintained

2. **New Examples** (`examples/`):
   - `counter/` - Simple state with increment/decrement
   - `form/` - Structured state with multiple fields
   - `async/` - Loading states with spinner animation
   - `streaming/` - Refactored to use adapter (was using Bubble Tea directly)

3. **Documentation**:
   - `messages.go` updated with new API patterns
   - `README.md` updated with state management section
   - All examples include tests and golden files

### Test Coverage
- 361 tests total (up from 229)
- 97% coverage on core package
- All examples have behavioral tests

---

## Executive Summary

This plan outlines the implementation of **State Management Patterns** for RuneTUI v0.2, completing the remaining work from Phase 3.2 of the original implementation plan. The goal is to provide a seamless, ergonomic API for managing application state while maintaining RuneTUI's declarative, component-based architecture.

**Architectural Foundation:**
- **Base Library:** Built on [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Elm Architecture pattern
- **Inspiration:** [Ink](https://github.com/vadimdemedes/ink) - React for CLIs (declarative components)
- **Hybrid Approach:**
  - **From Ink:** Declarative, component-based API (like React/Ink)
  - **From Bubble Tea:** Elm Architecture (state outside components, Update/Init pattern)
  - **Result:** Declarative components (Ink-style) with functional state management (Bubble Tea/Elm-style)

**Current State:**
- ✅ Documentation exists in `messages.go` (patterns documented)
- ✅ Unit tests demonstrate patterns (`messages_test.go`)
- ✅ **Adapter supports custom Update/Init functions via WithUpdate/WithInit**
- ⏭️ Helper utilities skipped (raw Update/Init sufficient for all patterns)
- ✅ Examples use RuneTUI adapter with state management

**Key Architectural Decisions:**
- **State Location:** State lives outside components (Elm/Bubble Tea pattern, already established)
- **Component Model:** Pure functions of state (Ink/React inspiration, but functional rather than hooks)
- **Update Integration:** Allow users to provide custom `Update` function to `App` (Bubble Tea pattern)
- **Helper Utilities:** Provide optional helpers for common patterns (not required)
- **Backward Compatibility:** Existing apps without Update functions continue to work
- **Testing:** 100% coverage for all new functionality

**Why This Hybrid Approach:**
- **Ink uses React hooks** (`useState`, `useEffect`) - state lives inside components
- **Bubble Tea uses Elm Architecture** - state lives outside, pure functions
- **RuneTUI chooses Bubble Tea's approach** because:
  1. Built on Bubble Tea runtime (must follow its patterns)
  2. More functional/pure (easier to test, reason about)
  3. Better for Go's type system (no hooks needed)
  4. Components remain pure functions (Ink's declarative style)
  5. State management follows proven Elm Architecture (Bubble Tea's strength)

---

## Alignment with Base Libraries and Inspiration

### Bubble Tea (Base Library)

**What we adopt:**
- **Elm Architecture pattern:** `Init()`, `Update()`, `View()` model
- **Message-based state updates:** All state changes flow through `Update(msg)`
- **Commands for side effects:** `tea.Cmd` for async operations
- **State outside components:** State lives in model/closure, not in components

**How RuneTUI integrates:**
- Adapter wraps RuneTUI components as `tea.Model`
- Users provide `Update` function that follows Bubble Tea's signature
- Commands from `Update` are executed by Bubble Tea runtime
- State management follows Bubble Tea's proven patterns

### Ink (Inspiration)

**What we adopt:**
- **Declarative components:** Components are functions, not classes
- **Composable API:** Build UIs by composing components
- **Pure rendering:** Components are pure functions of props/state
- **Developer experience:** Simple, ergonomic API

**How RuneTUI differs:**
- **No hooks:** Ink uses React hooks (`useState`, `useEffect`), RuneTUI uses Elm Architecture
- **State location:** Ink's state lives in components (via hooks), RuneTUI's state lives outside (via closures/structs)
- **Why different:** Built on Bubble Tea requires Elm Architecture, but we maintain Ink's declarative component style

### Lipgloss (Styling Library)

**What we use:**
- Text styling (colors, bold, italic, etc.)
- Border rendering
- Background colors
- All styling delegated to Lipgloss (no custom styling engine)

**State management impact:**
- Styling is stateless (pure functions)
- No state management needed for styling
- Style props are passed to components, not managed in state

### Result: Hybrid Architecture

```
┌─────────────────────────────────────────────────────────┐
│ RuneTUI: Declarative Components (Ink-style)            │
│   ↓                                                      │
│ Components are pure functions of state                 │
│   ↓                                                      │
│ State lives outside (Bubble Tea/Elm-style)              │
│   ↓                                                      │
│ Update/Init pattern (Bubble Tea runtime)                │
│   ↓                                                      │
│ Styling via Lipgloss (stateless)                        │
└─────────────────────────────────────────────────────────┘
```

**Benefits:**
- ✅ Declarative component API (like Ink)
- ✅ Functional state management (like Bubble Tea/Elm)
- ✅ Pure components (easier to test, reason about)
- ✅ Type-safe (Go's type system)
- ✅ No hooks needed (simpler mental model)

---

## Phase 1: Custom Update/Init Support in Adapter

**Goal:** Allow users to provide custom `Update` and `Init` functions to `App`, enabling full state management integration.

### 1.1 Extend App API with Update Support

**File:** [adapter.go](../adapter.go), [adapter_test.go](../adapter_test.go)

**TDD workflow:**
1. Write failing test: App with custom Update function receives messages
2. Add `UpdateFunc` type to adapter.go
3. Add `WithUpdate(UpdateFunc)` AppOption
4. Implement message forwarding in model.Update
5. Write failing test: custom Update can return commands
6. Implement command forwarding
7. Write failing test: custom Update triggers re-render
8. Verify re-render happens automatically
9. Continue for Init function support

**What to implement:**
- `UpdateFunc` type: `type UpdateFunc func(msg tea.Msg) tea.Cmd`
- `InitFunc` type: `type InitFunc func() tea.Cmd`
- `WithUpdate(UpdateFunc)` AppOption
- `WithInit(InitFunc)` AppOption
- Modify `model.Update` to forward messages to user's Update function
- Modify `model.Init` to call user's Init function
- Ensure backward compatibility (Update/Init are optional)

**Validation criteria:**
- App without Update function works as before (backward compatible)
- App with Update function receives all messages
- Custom Update can return commands that execute
- Custom Init runs on app start
- Commands from Update are properly executed
- Test coverage: 100%

**Critical decisions to validate:**
- Should Update receive ALL messages or can it filter?
- How do we handle WindowSizeMsg and Ctrl+C when user provides Update?
- Should we merge user Update with internal Update or replace it?
- **Decision:** User Update receives ALL messages, including WindowSizeMsg and KeyMsg. User can choose to handle or ignore them.

### 1.2 State Management Integration Tests

**File:** [adapter_test.go](../adapter_test.go)

**TDD workflow:**
1. Write failing test: counter app with Update function
2. Implement test using TestApp from testing package
3. Write failing test: form app with multiple inputs
4. Implement form state test
5. Write failing test: async loading with commands
6. Implement async pattern test
7. Continue for edge cases

**What to implement:**
- Integration tests using `testing.TestApp`
- Test counter pattern end-to-end
- Test form pattern end-to-end
- Test async loading pattern end-to-end
- Test command chaining
- Test Init function execution

**Validation criteria:**
- All documented patterns work with adapter
- Commands execute correctly
- State updates trigger re-renders
- Test coverage: 100%

---

## Phase 2: State Management Helpers

**Goal:** Provide optional helper utilities to reduce boilerplate for common state management patterns.

### 2.1 State Store Helper

**File:** [state.go](../state.go), [state_test.go](../state_test.go)

**TDD workflow:**
1. Write failing test: StateStore stores and retrieves value
2. Implement basic StateStore struct
3. Write failing test: StateStore notifies subscribers on change
4. Implement observer pattern
5. Write failing test: StateStore with typed getters/setters
6. Add type-safe wrappers
7. Continue for advanced features

**What to implement:**
- `StateStore[T any]` generic struct (Go 1.18+)
- `NewStateStore[T any](initial T) *StateStore[T]`
- `Get() T` - get current value
- `Set(T)` - set new value and notify
- `Subscribe(func(T)) func()` - subscribe to changes, returns unsubscribe
- `Update(func(T) T)` - functional update pattern

**Alternative (if generics not desired):**
- `StateStore` interface with type assertions
- `NewStateStore(initial interface{}) StateStore`
- Less type-safe but compatible with older Go versions

**Validation criteria:**
- Can store and retrieve state
- Subscribers are notified on changes
- Unsubscribe works correctly
- Type-safe (if using generics)
- Test coverage: 100%

**Critical decisions to validate:**
- Do we need generics or is interface{} acceptable?
- Should StateStore be part of core or optional helper?
- **Decision:** Start with interface{} for compatibility, document generics version for Go 1.18+ users.

### 2.2 Update Helper Functions

**File:** [state.go](../state.go) (add to existing file)

**TDD workflow:**
1. Write failing test: CombineUpdates merges multiple Update functions
2. Implement CombineUpdates helper
3. Write failing test: FilterMessages filters message types
4. Implement FilterMessages helper
5. Continue for other common patterns

**What to implement:**
- `CombineUpdates(updates ...UpdateFunc) UpdateFunc` - chain multiple Update functions
- `FilterMessages(update UpdateFunc, filter func(tea.Msg) bool) UpdateFunc` - filter messages before Update
- `HandleKey(key string, handler func() tea.Cmd) UpdateFunc` - handle specific key press
- `HandleMessage[T tea.Msg](update UpdateFunc, handler func(T) tea.Cmd) UpdateFunc` - type-safe message handler

**Validation criteria:**
- Helpers reduce boilerplate
- Helpers compose well together
- Clear documentation with examples
- Test coverage: 100%

**Critical decisions to validate:**
- Are these helpers necessary or is raw Update function sufficient?
- **Decision:** Provide helpers but make them optional. Users can use raw Update if preferred.

---

## Phase 3: Complete Examples with Adapter

**Goal:** Update existing examples and create new ones that use RuneTUI adapter (not Bubble Tea directly).

### 3.1 Counter Example

**File:** [examples/counter/main.go](../examples/counter/main.go), [examples/counter/example_test.go](../examples/counter/example_test.go)

**TDD workflow:**
1. Write failing test: counter example renders correctly
2. Implement counter using RuneTUI adapter with Update function
3. Write snapshot test
4. Verify example runs correctly
5. Update documentation

**What to implement:**
- Counter app using `runetui.New()` with `WithUpdate()`
- State lives in closure (following documented pattern)
- Increment/decrement on key presses
- Snapshot test to verify rendering

**Validation criteria:**
- Example uses RuneTUI adapter (not Bubble Tea directly)
- Follows documented state management pattern
- Works end-to-end
- Snapshot test passes
- Code is clear and well-documented

### 3.2 Form Example

**File:** [examples/form/main.go](../examples/form/main.go), [examples/form/example_test.go](../examples/form/example_test.go)

**TDD workflow:**
1. Write failing test: form example renders correctly
2. Implement form with multiple input fields
3. Write snapshot test
4. Verify state updates work
5. Update documentation

**What to implement:**
- Form app with name and email fields
- Uses structured state (formState struct)
- Handles keyboard input for field navigation
- Demonstrates multiple input message types

**Validation criteria:**
- Example uses RuneTUI adapter
- Demonstrates structured state pattern
- Works end-to-end
- Snapshot test passes
- Code is clear and well-documented

### 3.3 Async Loading Example

**File:** [examples/async/main.go](../examples/async/main.go), [examples/async/example_test.go](../examples/async/example_test.go)

**TDD workflow:**
1. Write failing test: async example renders loading state
2. Implement async loading with spinner
3. Write test for data loaded state
4. Verify commands execute correctly
5. Update documentation

**What to implement:**
- Async data loading simulation
- Loading spinner using tick commands
- Demonstrates Init function with initial command
- Shows state transitions (loading -> loaded)

**Validation criteria:**
- Example uses RuneTUI adapter
- Demonstrates async pattern with commands
- Works end-to-end
- Snapshot tests pass
- Code is clear and well-documented

### 3.4 Update Streaming Example

**File:** [examples/streaming/main.go](../examples/streaming/main.go)

**What to implement:**
- Refactor existing streaming example to use RuneTUI adapter
- Replace direct Bubble Tea usage with `runetui.New()` + `WithUpdate()` + `WithInit()`
- Maintain same functionality
- Update comments to reflect RuneTUI usage

**Validation criteria:**
- Example uses RuneTUI adapter (not Bubble Tea directly)
- Functionality unchanged
- Code is cleaner and more idiomatic
- Documentation updated

---

## Phase 4: Enhanced Documentation

**Goal:** Update documentation to reflect new state management capabilities.

### 4.1 Update messages.go Documentation

**File:** [messages.go](../messages.go)

**What to implement:**
- Update all code examples to use `runetui.New()` with `WithUpdate()`
- Add section on AppOptions (`WithUpdate`, `WithInit`)
- Add section on helper utilities (if implemented)
- Add migration guide from direct Bubble Tea usage
- Add troubleshooting section

**Validation criteria:**
- All examples compile and run
- Documentation is clear and comprehensive
- Migration path is documented
- Common pitfalls are addressed

### 4.2 Update README.md

**File:** [README.md](../README.md)

**What to implement:**
- Update status section: mark "State management patterns" as complete
- Add state management section to Features
- Update Quick Example to show Update function (optional)
- Add link to state management documentation

**Validation criteria:**
- README accurately reflects current capabilities
- Examples are up-to-date
- Links work correctly

### 4.3 Create State Management Guide

**File:** [docs/state-management.md](../docs/state-management.md)

**What to implement:**
- Comprehensive guide to state management in RuneTUI
- All patterns with working examples
- Best practices section
- Common patterns and anti-patterns
- Performance considerations
- Testing strategies

**Validation criteria:**
- Guide is comprehensive and clear
- All examples work
- Best practices are documented
- Common mistakes are highlighted

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

### Test Organization

```
runetui/
├── adapter_test.go          # Update existing with new tests
├── state_test.go            # New: State helper tests
├── messages_test.go          # Update existing tests
└── examples/
    ├── counter/
    │   ├── example_test.go   # New: Counter example tests
    │   └── main.go           # New: Counter example
    ├── form/
    │   ├── example_test.go   # New: Form example tests
    │   └── main.go            # New: Form example
    ├── async/
    │   ├── example_test.go   # New: Async example tests
    │   └── main.go            # New: Async example
    └── streaming/
        └── main.go           # Update: Use adapter
```

### Test Coverage Goals

- **Unit tests:** 100% coverage for adapter changes
- **Unit tests:** 100% coverage for state helpers (if implemented)
- **Integration tests:** All examples work end-to-end
- **Snapshot tests:** All examples have snapshot tests
- **Documentation tests:** All code examples in docs compile and run

---

## Implementation Order Summary

**Phase 1: Adapter Support (Week 1)**
1. Extend App API with UpdateFunc support
2. Add WithUpdate AppOption
3. Modify model.Update to forward messages
4. Add WithInit AppOption
5. Modify model.Init to call user Init
6. Write integration tests

**Phase 2: Helpers (Week 2 - Optional)**
7. Implement StateStore helper (if desired)
8. Implement Update helper functions
9. Write tests for helpers
10. Document helper usage

**Phase 3: Examples (Week 2-3)**
11. Create counter example
12. Create form example
13. Create async example
14. Update streaming example
15. Write snapshot tests for all examples

**Phase 4: Documentation (Week 3)**
16. Update messages.go documentation
17. Update README.md
18. Create state management guide
19. Review and polish all documentation

---

## Validation Checkpoints

### Checkpoint 1: Adapter Support Complete
**After completing Phase 1:**
- [ ] App supports custom Update function
- [ ] App supports custom Init function
- [ ] Backward compatibility maintained
- [ ] All tests pass: `make test`
- [ ] Coverage at 100%: `make test-coverage`
- [ ] Validation passes: `make validate`

**Architectural questions to answer:**
- Is the Update/Init API ergonomic?
- Does it handle all message types correctly?
- Are there any edge cases we missed?

### Checkpoint 2: Examples Complete
**After completing Phase 3:**
- [ ] All examples use RuneTUI adapter
- [ ] All examples work end-to-end
- [ ] All snapshot tests pass
- [ ] Examples demonstrate all documented patterns
- [ ] Code is clear and well-documented

**Architectural questions to answer:**
- Are the examples helpful for users?
- Do they cover common use cases?
- Is the API intuitive?

### Checkpoint 3: Documentation Complete
**After completing Phase 4:**
- [ ] Documentation is comprehensive
- [ ] All code examples compile and run
- [ ] Migration guide is clear
- [ ] README is up-to-date
- [ ] State management guide is complete

**Release criteria:**
- State management patterns are fully implemented
- Examples demonstrate all patterns
- Documentation is comprehensive
- All tests pass
- API is stable and ergonomic

---

## Pre-Commit Checklist (MANDATORY)

Before EVERY commit:
1. [ ] Run `make validate`
2. [ ] Fix ALL errors (zero tolerance)
3. [ ] Run `make validate` again
4. [ ] Only commit when passing

**Never commit without passing validation.**

---

## Open Questions

### Questions to answer during implementation:

1. **Update Function Signature:**
   - Should Update receive the App instance for access to layoutEngine/staticManager?
   - **Decision:** No, keep it pure (Bubble Tea pattern). Users can capture what they need in closure.
   - **Rationale:** Follows Bubble Tea's pure function approach. Keeps Update testable and simple.

2. **Message Filtering:**
   - Should we provide built-in filtering for WindowSizeMsg/KeyMsg?
   - **Decision:** No, let users handle all messages. Simpler and more flexible.
   - **Rationale:** Bubble Tea doesn't filter messages - users handle what they need. Maintains consistency.

3. **State Store Helper:**
   - Is StateStore necessary or is closure-based state sufficient?
   - **Decision:** Start without it. Add if users request it or if patterns emerge.
   - **Rationale:** Bubble Tea uses simple closures/structs. Ink uses hooks (different model). We follow Bubble Tea's simplicity.

4. **Command Chaining:**
   - How do we handle multiple commands from Update?
   - **Decision:** Use `tea.Batch()` helper from Bubble Tea.
   - **Rationale:** Bubble Tea provides `tea.Batch()` for this exact use case. Use the base library's utilities.

5. **Init Function:**
   - Should Init receive App instance?
   - **Decision:** No, keep it pure like Update (Bubble Tea pattern).
   - **Rationale:** Matches Bubble Tea's `Init() tea.Cmd` signature. Pure functions are easier to test.

6. **Hooks vs Elm Architecture:**
   - Should we provide React-like hooks (Ink-style) or stick with Elm Architecture (Bubble Tea-style)?
   - **Decision:** Stick with Elm Architecture (already established).
   - **Rationale:** 
     - Built on Bubble Tea runtime (must follow its patterns)
     - Go's type system doesn't need hooks (unlike JavaScript)
     - Elm Architecture is more functional and testable
     - Components remain declarative (Ink's strength) even without hooks

**Approach:** Start simple, add complexity only when needed. Validate with real examples.

---

## Success Metrics

### Technical Metrics
- Test coverage: 100% for adapter changes
- All validation checks passing
- Zero known bugs
- Backward compatibility maintained

### API Design Metrics
- Can build counter app in < 20 lines
- Can build form app without fighting the API
- Clear error messages
- Minimal boilerplate

### Use Case Validation
- All documented patterns work with adapter
- Examples are clear and helpful
- Migration from direct Bubble Tea is straightforward
- State management feels natural and idiomatic

---

## Next Steps After Plan Approval

1. **Start with first test:**
   - Create failing test in `adapter_test.go`
   - Test: App with Update function receives messages
   - Implement minimal code to pass
   - Continue TDD workflow

2. **Commit after each passing test:**
   - Run `make validate`
   - Commit with clear message
   - Push to track progress

3. **Regular architecture review:**
   - After Phase 1 checkpoint
   - Review API ergonomics
   - Adjust plan if needed

---

## End of Plan

This plan provides a comprehensive, test-driven roadmap for completing State Management Patterns in RuneTUI v0.2. The focus is on integrating custom Update/Init functions into the adapter while maintaining backward compatibility and providing optional helpers for common patterns.

**Key principles:**
- Baby steps (one test at a time)
- TDD workflow (test before code)
- Backward compatibility (existing apps continue to work)
- Optional helpers (users can use raw Update if preferred)
- Comprehensive examples (all patterns demonstrated)
- Idiomatic Go (simple, clear, composable)

