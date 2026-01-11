# RuneTUI - AI Agent Rules

Rules for AI agents working on RuneTUI, a Go TUI framework.

## 1. Core Principles

- **Baby Steps**: One test, one change at a time. Never skip ahead.
- **TDD**: Always start with a failing test before implementation.
- **Simplicity**: Use the simplest working solution. No premature abstractions.
- **Small Code**: Functions ≤20 lines. Flag anything larger for refactoring.
- **Self-Documenting**: Clear names, no comments. Code explains itself.
- **100% Coverage**: All code must be tested. No exceptions.

## 2. Project Structure

RuneTUI is a **library** (not a service). Use flat structure:

```
runetui/
├── component.go      # Public API at root
├── text.go
├── box.go
├── layout.go
├── internal/         # Private helpers (if needed)
├── testing/          # Test utilities for users
└── examples/         # Usage examples
```

**No hexagonal architecture** - UI frameworks have no business/infrastructure split.
**No `utils` or `common`** - name packages by semantic purpose.

## 3. Error Handling

```go
// Export domain errors as variables
var ErrInvalidDimension = errors.New("invalid dimension")

// Always wrap with context
if err != nil {
    return fmt.Errorf("rendering box %s: %w", key, err)
}
```

- Core packages never read env vars or config
- Only adapter/examples wire configuration

## 4. TDD Workflow

1. Write ONE failing test
2. Run `make test` - verify it fails
3. Write minimal code to pass
4. Run `make test` - verify it passes
5. Refactor if needed
6. Repeat

**Test naming:** `TestComponent_Scenario_ExpectedBehavior`

```go
func TestText_WithBoldStyle_RendersBoldText(t *testing.T)
func TestBox_WithPadding_ReducesAvailableSpace(t *testing.T)
```

## 5. Task Runner (Makefile)

**NEVER** call tools directly. Always use Make:

```sh
# ✅ Good
make test
make lint
make fmt
make validate

# ❌ Bad
go test ./...
golangci-lint run
```

Run `make help` to discover available tasks.

## 6. Pre-Commit (MANDATORY)

Before EVERY commit:
1. Run `make validate`
2. Fix ALL errors
3. Only commit when passing

❌ Never: Commit → find errors → fix
✅ Always: Validate → fix → commit

## 7. Language & Communication

- **Conversation**: Any language (respond in user's language)
- **Artifacts**: Always English (code, commits, docs, tests)
- **Style**: Show reasoning, ask when unclear, be concise

## Quick Reference

1. 👣 Baby steps - one test at a time
2. ❌➡️✅ TDD - failing test first
3. 🔧 Use `make` - never call tools directly
4. 📏 Small code - ≤20 lines per function
5. ✅ Validate before every commit
6. 🧪 Run tests after every change
7. ❓ Ask when in doubt
