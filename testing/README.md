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

**Examples:**

Text component:
```go
func TestText_WithBold_AppliesBoldStyle(t *testing.T) {
    text := Text("Hello", TextProps{Bold: true})
    layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

    got := text.Render(layout)
    compareWithGolden(t, "text_bold", got)  // ✅ Verifies exact bold output
}
```

Box component:
```go
func TestBox_Render_WithBorderAppliesLipgloss(t *testing.T) {
    child := &mockComponent{key: "child", content: "Text"}
    props := BoxProps{Key: "box", Border: BorderSingle}
    box := Box(props, child)

    layout := Layout{X: 0, Y: 0, Width: 20, Height: 10}
    got := box.Render(layout)

    compareWithGoldenBox(t, "box_border_single", got)  // ✅ Verifies exact border output
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

### Text Component

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

### Box Component

**Border Styles (3):**
- `box_border_single.golden` - Single line border
- `box_border_double.golden` - Double line border
- `box_border_rounded.golden` - Rounded corners border

**Colors (2):**
- `box_background_red.golden` - Red background (#FF0000)
- `box_border_color_green.golden` - Green border color (#00FF00)

**Combinations (1):**
- `box_border_background.golden` - Border + background combination

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

## Assertion Helpers (Phase 2)

For non-critical tests where exact output isn't important, use assertion helpers.
These helpers verify **properties** of the output without being coupled to exact ANSI codes.

### Available Helpers

Import from the testing package:

```go
import runetesting "github.com/runetui/runetui/testing"
```

#### `AssertHasANSICodes(t, output)`

Verifies that output contains ANSI escape sequences (styled).

```go
runetesting.AssertHasANSICodes(t, got)  // Fails if no ANSI codes present
```

#### `AssertContainsText(t, output, text)`

Verifies that visible text content is present, ignoring ANSI codes.

```go
runetesting.AssertContainsText(t, got, "Hello")  // Finds "Hello" even with styling
```

#### `AssertWidth(t, output, expected)`

Verifies the visible width of output, excluding ANSI codes.

```go
runetesting.AssertWidth(t, got, 10)  // Checks visible width is 10
```

#### `AssertHeight(t, output, expected)`

Verifies the number of lines in output.

```go
runetesting.AssertHeight(t, got, 3)  // Checks output has 3 lines
```

#### `AssertNotEmpty(t, output)`

Verifies output has visible content (not just whitespace or ANSI codes).

```go
runetesting.AssertNotEmpty(t, got)  // Fails if output is empty/whitespace
```

### When to Use Helpers vs Golden Files

| Scenario | Use |
|----------|-----|
| Verify exact bold rendering | Golden file |
| Verify output has styling | `AssertHasANSICodes` |
| Verify text content preserved | `AssertContainsText` |
| Verify layout dimensions | `AssertWidth`, `AssertHeight` |
| Sanity check output exists | `AssertNotEmpty` |
| Test style combinations | Helpers (table-driven) |
| Test critical single styles | Golden file |

### Table-Driven Tests with Helpers

For testing many variations, use table-driven tests with helpers:

**Text component example:**
```go
func TestText_StyleCombinations_ProducesValidOutput(t *testing.T) {
    tests := []struct {
        name    string
        props   TextProps
        content string
    }{
        {"bold_only", TextProps{Bold: true}, "Hello"},
        {"italic_only", TextProps{Italic: true}, "Hello"},
        {"bold_italic", TextProps{Bold: true, Italic: true}, "Hello"},
        {"color_only", TextProps{Color: "#FF0000"}, "Red"},
        {"all_styles", TextProps{Bold: true, Italic: true, Underline: true}, "Test"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            text := Text(tt.content, tt.props)
            layout := Layout{Width: 20, Height: 1}

            got := text.Render(layout)

            // Verify properties, not exact output
            runetesting.AssertHasANSICodes(t, got)
            runetesting.AssertContainsText(t, got, tt.content)
            runetesting.AssertNotEmpty(t, got)
        })
    }
}
```

**Box component example:**
```go
func TestBox_StyleCombinations_ProducesValidOutput(t *testing.T) {
    tests := []struct {
        name  string
        props BoxProps
    }{
        {"border_only", BoxProps{Key: "box", Border: BorderSingle}},
        {"background_only", BoxProps{Key: "box", Background: "#FF0000"}},
        {"border_and_background", BoxProps{Key: "box", Border: BorderSingle, Background: "#00FF00"}},
        {"all_styles", BoxProps{Key: "box", Border: BorderRounded, BorderColor: "#FF00FF", Background: "#FFFF00"}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            child := &mockComponent{key: "child", content: "Test"}
            box := Box(tt.props, child)
            layout := Layout{Width: 20, Height: 10}

            got := box.Render(layout)

            // Verify properties, not exact output
            AssertContainsText(t, got, "Test")
            AssertNotEmpty(t, got)

            // Verify ANSI codes when color is applied
            if tt.props.Background != "" || tt.props.BorderColor != "" {
                AssertHasANSICodes(t, got)
            }
        })
    }
}
```

### Benefits of Helpers

1. **Less Brittle** - Don't break when Lipgloss changes ANSI codes
2. **More Readable** - Clear intent from helper names
3. **Faster to Write** - No golden files to manage
4. **Good for Variations** - Test many combinations quickly

### When to Add New Golden Files

Only add golden files when:

1. You need to verify **exact** output (specific ANSI codes)
2. The test covers a **critical** styling behavior
3. You want to catch **any** change in output

For other cases, use assertion helpers.

## References

- **Test Desiderata:** https://testdesiderata.com/
- **CLAUDE.md:** Project rules and TDD principles
- **Test Plan:** docs/test-behavioral-improvement-plan.md

---

**Status:** Phase 2 Complete + Box Component Extended ✅

**Implemented:**
- Golden file infrastructure (Phase 1)
- Assertion helpers: `AssertHasANSICodes`, `AssertContainsText`, `AssertWidth`, `AssertHeight`, `AssertNotEmpty` (Phase 2)
- Table-driven tests for style combinations (Phase 2)
- Box component golden files for borders, colors, and combinations
- Box component table-driven tests using assertion helpers
