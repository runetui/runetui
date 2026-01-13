# Architecture Decision Records

This document captures important architectural decisions made in RuneTUI.

---

## ADR-0001: Hybrid Testing Approach with Golden Files

**Date:** 2026-01-13
**Status:** Accepted

### Context

Our test suite had a critical behavioral testing issue: tests were checking for ANSI code presence, not actual behavior. For example, a test for bold text would pass even if italic was applied instead, as long as some ANSI code existed.

This violated Kent Beck's Test Desiderata, specifically:
- **Behavioral**: Tests must verify actual output, not implementation details
- **Predictive**: Tests must catch real styling bugs
- **Specific**: Failures must clearly indicate what broke

We evaluated 5 approaches:
1. **Snapshot Testing (Golden Files)** - Behavioral but brittle
2. **ANSI Parser + Assertions** - Structure-insensitive but complex
3. **Test Helpers** - Easy to write but doesn't solve core problem
4. **Mock Renderer** - Over-engineering, violates CLAUDE.md principles ❌
5. **Hybrid Approach** - Balanced, pragmatic ✅

### Decision

We implement a **Hybrid Testing Approach** combining:

**Phase 1: Golden Files for Core Styles**
- Create `testdata/golden/*.golden` files for fundamental styles
- Test helpers: `loadGoldenFile()`, `updateGoldenFile()`, `compareWithGolden()`
- Use `-update` flag for regenerating golden files
- Cover: bold, italic, underline, colors, critical combinations

**Phase 2: Test Helpers for Variations**
- Assertion helpers: `assertHasANSICodes()`, `assertContainsText()`, `assertWidth()`
- Table-driven tests for style combinations
- Property-based testing (output has ANSI, text preserved, dimensions respected)

**Phase 3: ANSI Parser (Optional)**
- Only if Phase 1+2 prove insufficient
- Would add semantic assertions: `assertHasBoldStyle()`, `assertHasColor()`

### Consequences

**Positive:**
- ✅ Behavioral testing - verifies actual output
- ✅ Catches real bugs - wrong styles cause test failures
- ✅ Maintainable - `-update` flag for easy updates
- ✅ Follows CLAUDE.md - simple, incremental, pragmatic
- ✅ No over-engineering - uses standard Go patterns

**Negative:**
- ⚠️ Partial structure-sensitivity - golden files coupled to Lipgloss
- ⚠️ Two patterns to learn - developers must know when to use each
- ⚠️ Maintenance on Lipgloss updates - must regenerate golden files

**Neutral:**
- Standard Go testing patterns (well-understood by community)
- No external dependencies for Phase 1+2

### Implementation

See detailed plan in `docs/test-behavioral-improvement-plan.md`

---

## ADR-0002: Flat Project Structure for Library Code

**Date:** 2026-01-13
**Status:** Accepted

### Context

RuneTUI is a **library** (not a service). Many Go projects use hexagonal architecture with layers like `domain/`, `application/`, `infrastructure/`, but this adds unnecessary complexity for UI frameworks.

UI frameworks have no business/infrastructure split - they are purely presentational code.

### Decision

Use a **flat structure** at root level:

```
runetui/
├── component.go      # Public API at root
├── text.go
├── box.go
├── layout.go
├── stack.go
├── internal/         # Private helpers (if truly needed)
├── testing/          # Test utilities for library users
└── examples/         # Usage examples
```

**Rules:**
- No `pkg/` directory - this is already a package
- No `cmd/` directory - this is a library, not a service
- No `utils/` or `common/` - name by semantic purpose
- No hexagonal layers - no artificial domain/application/infrastructure split
- Use `internal/` only when necessary for truly private code

### Consequences

**Positive:**
- ✅ Simple - easy to navigate
- ✅ Clear API surface - root = public API
- ✅ Follows Go conventions for libraries
- ✅ Low cognitive load - no artificial boundaries

**Negative:**
- ⚠️ Root directory can grow - must manage file count
- ⚠️ Less structure - developers must use discipline

**Neutral:**
- Standard pattern for Go libraries (cobra, viper, chi, etc.)

---

## ADR-0003: Lipgloss as Styling Engine

**Date:** 2026-01-13
**Status:** Accepted

### Context

TUI applications need terminal styling (colors, bold, italic, etc.). We need a reliable way to generate ANSI escape codes that work across different terminals.

Options:
1. **Write custom ANSI generator** - full control but high maintenance
2. **Use low-level library** (termenv) - flexible but verbose
3. **Use Lipgloss** - high-level, battle-tested, maintained by Charm

### Decision

Use **Lipgloss** (`github.com/charmbracelet/lipgloss`) as our styling engine.

Lipgloss provides:
- High-level styling API
- Cross-platform ANSI code generation
- Color support with automatic terminal capability detection
- Layout primitives (alignment, padding, borders)
- Battle-tested by Charm ecosystem (Bubble Tea, Glow, etc.)

### Consequences

**Positive:**
- ✅ Reliable - battle-tested in production TUI apps
- ✅ Maintained - active development by Charm
- ✅ Cross-platform - handles terminal differences
- ✅ High-level API - easier to use than raw ANSI codes
- ✅ Ecosystem - integrates well with other Charm tools

**Negative:**
- ⚠️ External dependency - adds to go.mod
- ⚠️ Coupling - our API somewhat coupled to Lipgloss
- ⚠️ Breaking changes - Lipgloss updates may require golden file updates

**Neutral:**
- Standard choice for Go TUI frameworks
- Trade-off: convenience vs. control

### Mitigation

- Pin Lipgloss version in `go.mod`
- Use golden files to detect breaking changes
- Update golden files when upgrading (using `-update` flag)
- Keep our public API abstract enough to potentially swap implementations

---

## ADR-0004: TDD with Baby Steps

**Date:** 2026-01-13
**Status:** Accepted

### Context

Software development approaches vary widely. We need a consistent methodology that ensures code quality, prevents bugs, and keeps the codebase maintainable.

### Decision

Follow **Test-Driven Development (TDD) with Baby Steps**:

**Workflow:**
1. Write ONE failing test
2. Run `make test` - verify it fails
3. Write minimal code to pass
4. Run `make test` - verify it passes
5. Refactor if needed
6. Repeat

**Test Quality Rules:**
- Functions ≤20 lines (flag for refactoring if larger)
- 100% test coverage - no exceptions
- Baby steps ≠ superficial tests
- Verify concrete values, not just "exists" or "not nil"
- Test edge cases (zero, negative, boundary values)
- Validate invariants (enums don't change order, ranges respected)

**Task Runner:**
- NEVER call tools directly (`go test`, `golangci-lint`)
- ALWAYS use Make: `make test`, `make lint`, `make fmt`, `make validate`
- Ensures consistent tooling and flags across team

### Consequences

**Positive:**
- ✅ High confidence - every line tested
- ✅ Prevents regressions - tests catch breaking changes
- ✅ Better design - TDD leads to more testable code
- ✅ Documentation - tests show how to use the API
- ✅ Consistent workflow - `make` commands standardize process

**Negative:**
- ⚠️ Slower initial development - writing tests takes time
- ⚠️ Discipline required - must resist skipping tests

**Neutral:**
- Tests become first-class code (not afterthought)
- Refactoring becomes safe and confident

### Rationale

Documented in `CLAUDE.md` as core principle. This approach has proven effective for maintaining code quality and enabling fearless refactoring.

---

## Template for New ADRs

When adding a new ADR, use this format:

```markdown
## ADR-XXXX: [Title]

**Date:** YYYY-MM-DD
**Status:** [Proposed | Accepted | Rejected | Deprecated]

### Context

What problem are we solving? What forces are at play?

### Decision

What did we decide? Be specific.

### Consequences

**Positive:**
- What becomes easier?

**Negative:**
- What becomes harder?

**Neutral:**
- What changes without clear positive/negative?

### Alternatives Considered (optional)

What other options did we evaluate? Why rejected?
```
