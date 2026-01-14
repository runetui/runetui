# RuneTUI Testing Guide

This guide explains how to write behavioral tests for RuneTUI components using golden files.

## Problem: Why Golden Files?

Previously, our tests only checked for the **presence** of ANSI codes, not **actual behavior**:

```go
// ❌ BAD - False positives possible
if !strings.Contains(got, "\x1b[1m") && !strings.Contains(got, "\x1b[") {
    t.Errorf("Expected ANSI bold codes...")
}
```

**Problems:**
- Test passes if we apply italic instead of bold
- Test passes if we apply wrong color
- Test passes if ANSI code is malformed

**Solution: Golden files** verify the **complete output**, not just code presence.

## Golden File Testing

### How It Works

Golden files store the expected output of a component render. Tests compare actual output against these "golden" reference files.

```go
func TestText_WithBold_AppliesBoldStyle(t *testing.T) {
    text := Text("Hello", TextProps{Bold: true})
    layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

    got := text.Render(layout)

    // Verify actual output matches golden file
    compareWithGolden(t, "text_bold", got)
}
```

### First Run

On the first run, golden files are **automatically created**:

```bash
$ go test -run TestText_WithBold
=== RUN   TestText_WithBold_AppliesBoldStyle
    text_test.go:71: Created golden file: text_bold
--- PASS: TestText_WithBold_AppliesBoldStyle (0.00s)
```

### Subsequent Runs

Tests compare output against the golden file:

```go
// ✅ PASS - output matches golden file
// ❌ FAIL - output differs, shows diff
```

### Updating Golden Files

When you **intentionally** change output (e.g., Lipgloss upgrade), regenerate golden files:

```bash
go test -update
```

**⚠️ Warning:** Only use `-update` when changes are intentional. Review diffs carefully.

## When to Use Golden Files

Use golden files for **critical behavioral tests**:

### ✅ Use Golden Files For:

1. **Core styles** - Bold, italic, underline, strikethrough
2. **Colors** - Specific colors (red, blue, green)
3. **Critical combinations** - Bold + color, italic + underline
4. **Layout dimensions** - Width, height, padding behavior
5. **Text wrapping** - Word wrap, character wrap, truncate

**Example:**
```go
func TestText_WithBold_AppliesBoldStyle(t *testing.T) {
    text := Text("Hello", TextProps{Bold: true})
    layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

    got := text.Render(layout)
    compareWithGolden(t, "text_bold", got)  // ✅ Verifies exact bold output
}
```

### ❌ Don't Use Golden Files For:

1. **Non-critical variations** - All possible style combinations
2. **Property testing** - Verify text content preserved
3. **Sanity checks** - Output is not empty
4. **Internal logic** - Business logic without rendering

**Use helpers instead** (Phase 2 - coming soon).

## Decision Tree

```
Does the test verify critical styling behavior?
├─ YES → Use golden file (compareWithGolden)
│         Examples: bold, italic, colors, combinations
│
└─ NO → Use assertion helpers (Phase 2)
          Examples: text content, dimensions, non-empty
```

## Golden File Helpers

### `compareWithGolden(t, name, got)`

Compares output against golden file. Auto-creates if missing.

**Parameters:**
- `t` - testing.T
- `name` - golden file name (without .golden extension)
- `got` - actual output to compare

**Example:**
```go
compareWithGolden(t, "text_bold", got)
// Compares against: testdata/text_bold.golden
```

### `updateGoldenFile(t, name, content)`

Manually update a golden file (rarely needed).

### `loadGoldenFile(t, name)`

Load golden file contents (for advanced testing).

## Current Golden Files

**Basic Styles (4):**
- `text_bold.golden` - Bold text
- `text_italic.golden` - Italic text
- `text_underline.golden` - Underlined text
- `text_strikethrough.golden` - Strikethrough text

**Colors (3):**
- `text_foreground_red.golden` - Red text (#FF0000)
- `text_foreground_blue.golden` - Blue text (#0000FF)
- `text_background_green.golden` - Green background (#00FF00)

**Combinations (3):**
- `text_bold_red.golden` - Bold + red color
- `text_italic_underline.golden` - Italic + underline
- `text_bold_background.golden` - Bold + blue background

## Best Practices

### 1. One Behavior Per Test

```go
// ✅ Good - tests one thing
func TestText_WithBold_AppliesBoldStyle(t *testing.T) {
    text := Text("Hello", TextProps{Bold: true})
    layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}
    got := text.Render(layout)
    compareWithGolden(t, "text_bold", got)
}

// ❌ Bad - tests multiple things
func TestText_StylesWork(t *testing.T) {
    // Tests bold, italic, and colors in one test
}
```

### 2. Clear Test Names

Follow pattern: `TestComponent_Scenario_ExpectedBehavior`

```go
TestText_WithBold_AppliesBoldStyle
TestBox_WithPadding_ReducesContentArea
TestLayout_WithFixedWidth_RespectsConstraints
```

### 3. Keep Golden Files Small

Use minimal dimensions for faster tests:

```go
// ✅ Good - small, focused
layout := Layout{Width: 10, Height: 1}

// ❌ Bad - unnecessarily large
layout := Layout{Width: 100, Height: 50}
```

### 4. Review Golden Files in Git

When golden files change, **always review the diff**:

```bash
git diff testdata/
```

Ensure changes match your intention.

## Troubleshooting

### Test Fails After Lipgloss Upgrade

**Solution:** Regenerate golden files after confirming changes are expected:

```bash
# Review changes first
go test ./... | grep FAIL

# Update golden files
go test -update

# Verify tests pass
go test ./...

# Review diffs
git diff testdata/

# Commit if correct
git add testdata/
git commit -m "test: update golden files for Lipgloss vX.Y.Z"
```

### Golden File Doesn't Exist

**Solution:** Tests auto-create golden files on first run. Just run the test:

```bash
go test -run TestYourNewTest
```

### Test Output Mismatch

**Error message shows:**
```
got:
"\x1b[1mHello\x1b[0m     "

want:
"\x1b[3mHello\x1b[0m     "
```

**This means:** Test caught a bug! Bold (\x1b[1m) was applied instead of italic (\x1b[3m).

**Fix the code**, don't update the golden file (unless the change is intentional).

## Next Steps

### Phase 2: Assertion Helpers (Coming Soon)

For non-critical tests, we'll add helpers:

```go
// Phase 2 helpers (not yet implemented)
assertHasANSICodes(t, got)
assertContainsText(t, got, "Hello")
assertWidth(t, got, 10)
```

Use golden files for now. We'll migrate some tests to helpers in Phase 2.

## References

- **Test Desiderata:** https://testdesiderata.com/
- **CLAUDE.md:** Project rules and TDD principles
- **Test Plan:** docs/test-behavioral-improvement-plan.md

---

**Status:** Phase 1 Complete ✅
**Next:** Phase 2 - Assertion Helpers
