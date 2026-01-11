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

### Test Quality Rules

**Baby steps ≠ Superficial tests.** Every test must validate actual behavior.

#### What Makes a Good Test

**✅ DO:**
- Verify concrete values, not just "exists" or "not nil"
- Test edge cases (zero, negative, boundary values)
- Validate invariants (enums don't change order, ranges respected)
- Check actual behavior (stored value == retrieved value)
- Test error conditions if applicable

**❌ DON'T:**
- Only check `!= nil` (except for constructors that return interfaces)
- Only check values are "different" without verifying which is which
- Skip boundary validation (negative, overflow, empty)
- Trust implementation without verifying observable behavior

#### Examples

**❌ Superficial (BAD):**
```go
func TestDirection_Values_ExistAndAreDifferent(t *testing.T) {
    if Column == Row {  // Doesn't verify which is 0 or 1
        t.Error("should be different")
    }
}
```

**✅ Robust (GOOD):**
```go
func TestDirection_Column_IsZero(t *testing.T) {
    if Column != 0 {  // Protects against reordering
        t.Errorf("Column should be 0, got %d", Column)
    }
}

func TestDirection_Row_IsOne(t *testing.T) {
    if Row != 1 {
        t.Errorf("Row should be 1, got %d", Row)
    }
}
```

**❌ Superficial (BAD):**
```go
func TestDimensionFixed_WithValue_CanBeCreated(t *testing.T) {
    dim := DimensionFixed(100)
    if dim == nil {  // Doesn't verify value is stored
        t.Error("should not be nil")
    }
}
```

**✅ Robust (GOOD):**
```go
// First expose the value through interface or method
type fixedDimension interface {
    Dimension
    Value() int
}

func TestDimensionFixed_StoresValue(t *testing.T) {
    dim := DimensionFixed(100)
    fixed, ok := dim.(fixedDimension)
    if !ok {
        t.Fatal("should implement fixedDimension")
    }
    if got := fixed.Value(); got != 100 {
        t.Errorf("expected 100, got %d", got)
    }
}

func TestDimensionFixed_NegativeValue_IsError(t *testing.T) {
    dim := DimensionFixed(-10)
    // Should either: panic, return error, or clamp to 0
    // Test the actual behavior
}
```

#### Test Checklist (per type)

**Enums:**
- [ ] Verify concrete values (0, 1, 2...) not just "different"
- [ ] Protects against reordering

**Structs:**
- [ ] All fields can be set and retrieved
- [ ] Zero values behave correctly
- [ ] Edge cases (negative, max, empty) handled

**Functions/Constructors:**
- [ ] Return values match inputs
- [ ] Edge cases return expected results
- [ ] Invalid inputs handled (error/panic/clamp)

**Interfaces:**
- [ ] Concrete implementations satisfy interface
- [ ] Method contracts verified through behavior
- [ ] Can't test internal state? Design is wrong - expose needed behavior

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
3. 💪 Robust tests - verify concrete values, test edge cases, validate invariants
4. 🔧 Use `make` - never call tools directly
5. 📏 Small code - ≤20 lines per function
6. ✅ Validate before every commit
7. 🧪 Run tests after every change
8. ❓ Ask when in doubt
