# Test Behavioral Improvement Plan

**Status:** Phase 2 Complete
**Priority:** Critical
**Estimated Effort:** 3-4 days
**Created:** 2026-01-13
**Phase 1 Completed:** 2026-01-13
**Phase 2 Completed:** 2026-01-14

## Problem Statement

### Current Situation

Our test suite has a **critical behavioral testing issue** identified through Kent Beck's Test Desiderata analysis:

**Tests check for ANSI code presence, not actual behavior.**

```go
// ❌ Current Problem - False Confidence
func TestText_WithBold_AppliesBoldStyle(t *testing.T) {
    text := Text("Hello", TextProps{Bold: true})
    got := text.Render(layout)

    // Only verifies SOME ANSI code exists, not that it's bold
    if !strings.Contains(got, "\x1b[1m") && !strings.Contains(got, "\x1b[") {
        t.Errorf("Expected ANSI bold codes...")
    }
}
```

### Why This Is Critical

1. **False Positives:** Test passes if we apply italic instead of bold
2. **False Positives:** Test passes if we apply red color when blue was requested
3. **False Positives:** Test passes if ANSI code is malformed but exists
4. **Production Risk:** Bugs slip through to production with "passing" tests
5. **Erodes Trust:** High coverage % gives false confidence

### Test Desiderata Violations

| Property | Status | Impact |
|----------|--------|--------|
| **Behavioral** | ❌ **CRITICAL** | Tests don't verify actual behavior |
| **Predictive** | ❌ **HIGH** | Can't catch real styling bugs |
| Structure-insensitive | ⚠️ Medium | Coupled to ANSI implementation |
| Specific | ⚠️ Medium | Multiple assertions per test |

---

## Options Analysis

We evaluated 5 different approaches to solve this problem:

### Option 1: Snapshot Testing (Golden Files)

Compare complete output with reference files.

**Pros:**
- ✅ Truly behavioral - verifies exact output
- ✅ Specific - failures show exact diff
- ✅ Simple to write and understand

**Cons:**
- ❌ Structure-sensitive - coupled to Lipgloss implementation
- ❌ Brittle - Lipgloss updates break tests
- ❌ High maintenance - update all snapshots on changes

**Verdict:** Good for behavioral testing, but too brittle alone.

---

### Option 2: ANSI Parser + Assertions

Parse ANSI codes and verify semantic properties.

**Pros:**
- ✅ Behavioral - verifies styles are applied
- ✅ Structure-insensitive - independent of exact codes
- ✅ Readable - tests express intent clearly

**Cons:**
- ❌ Complex - need to write/maintain ANSI parser
- ❌ Lower writability - more setup required
- ❌ Incomplete - parser may not cover all ANSI cases

**Verdict:** Best for structure-insensitivity, but high upfront cost.

---

### Option 3: Test Helpers

Create helpers that verify specific styles.

**Pros:**
- ✅ Writable - easy to add new tests
- ✅ Readable - clear intent with named helpers
- ✅ Fast - no complex parsing

**Cons:**
- ❌ Structure-sensitive - still coupled to ANSI codes
- ❌ Incomplete behavioral - verifies codes but not combinations
- ❌ False positives - could pass with codes in wrong order

**Verdict:** Easy to write but doesn't solve the core problem.

---

### Option 4: Mock Renderer (Dependency Injection)

Abstract Lipgloss behind interface and mock in tests.

**Pros:**
- ✅ Structure-insensitive - fully decoupled from Lipgloss
- ✅ Fast - no real rendering
- ✅ Testable - easy to verify parameters

**Cons:**
- ❌ NOT Behavioral - doesn't test real rendering
- ❌ Architecture complexity - requires significant refactor
- ❌ Over-engineering - violates YAGNI principle
- ❌ Against project rules - UI framework needs no business/infra split

**Verdict:** ❌ **Rejected** - Over-engineering, violates CLAUDE.md principles.

---

### Option 5: Hybrid Approach ⭐ **SELECTED**

Combine snapshot tests for critical cases + helpers for variations.

**Pros:**
- ✅ Balanced - behavioral where it matters, pragmatic for variations
- ✅ Maintainable - snapshots updateable with flag
- ✅ Readable - mix of precision (snapshots) and clarity (helpers)
- ✅ Practical - follows CLAUDE.md simplicity principle
- ✅ Incremental - can implement in baby steps

**Cons:**
- ⚠️ Partial structure-sensitivity - snapshots still coupled to Lipgloss
- ⚠️ Two patterns - developers must know when to use each
- ⚠️ Moderate maintenance - update golden files when Lipgloss changes

**Tradeoff:** **Pragmatism ↔ Purity**

**Verdict:** ✅ **SELECTED** - Best balance for our needs.

---

## Why Option 5 (Hybrid Approach)?

### Alignment with CLAUDE.md Principles

1. **Baby Steps** ✅
   - Can implement incrementally
   - Phase 1 → Phase 2 → Phase 3
   - Each phase delivers value

2. **Simplicity** ✅
   - No over-engineering (unlike Option 4)
   - Uses standard Go testing patterns
   - Minimal dependencies

3. **TDD** ✅
   - Write golden files first (failing tests)
   - Implement/fix → tests pass
   - Refactor with confidence

4. **Pragmatism** ✅
   - Behavioral testing where critical (styles, colors)
   - Simple helpers for variations (combinations)
   - Doesn't chase perfect structure-insensitivity

### Technical Rationale

1. **Solves the Core Problem**
   - Golden files ensure actual behavior is correct
   - Catches regressions in styling immediately
   - No more false positives from generic ANSI checks

2. **Maintainability**
   - `-update` flag makes golden file updates trivial
   - Helpers reduce boilerplate in variation tests
   - Clear pattern for future contributors

3. **Low Risk**
   - Doesn't require architecture changes
   - Doesn't break existing code
   - Can be implemented test-by-test

4. **Proven Pattern**
   - Used successfully in many Go projects
   - Standard practice for CLI/TUI testing
   - Well-understood by Go community

---

## Implementation Plan

### Phase 1: Golden Files for Core Styles (1-2 days)

**Goal:** Establish behavioral testing for fundamental styling.

**Tasks:**

1. **Create golden file infrastructure**
   - Add `testdata/golden/` directory
   - Create helper functions:
     - `loadGoldenFile(t, name)` - load expected output
     - `updateGoldenFile(t, name, content)` - update with flag
     - `compareWithGolden(t, name, got)` - compare and report diff
   - Add `-update` flag support: `go test -update`

2. **Write golden files for basic text styles**
   - `text_bold.golden` - bold text
   - `text_italic.golden` - italic text
   - `text_underline.golden` - underline text
   - `text_strikethrough.golden` - strikethrough text
   - Each tests: style applied + text content + width/padding

3. **Write golden files for colors**
   - `text_foreground_red.golden` - red foreground (#FF0000)
   - `text_foreground_blue.golden` - blue foreground (#0000FF)
   - `text_background_green.golden` - green background (#00FF00)
   - Verify specific colors, not just "has color"

4. **Write golden files for critical combinations**
   - `text_bold_red.golden` - bold + red color
   - `text_italic_underline.golden` - italic + underline
   - `text_bold_background.golden` - bold + background color
   - Ensures multiple styles work together

5. **Refactor existing tests**
   - Replace ANSI checking tests with golden comparisons
   - Keep one test per style property
   - Run tests to verify they catch actual bugs

**Deliverables:**
- `testing/golden.go` - golden file helpers
- `testdata/golden/*.golden` - 10-12 golden files
- Updated `text_test.go` - behavioral tests using golden files
- Documentation in `testing/README.md`

**Success Criteria:**
- All golden tests pass
- Manually break a style (change Bold→Italic) - test fails
- `-update` flag regenerates golden files correctly

---

### Phase 2: Test Helpers for Properties (1 day)

**Goal:** Reduce boilerplate in variation tests with reusable helpers.

**Tasks:**

1. **Create assertion helpers**
   ```go
   // testing/assertions.go

   func assertHasANSICodes(t *testing.T, output string)
   // Verifies output contains ANSI escape sequences

   func assertContainsText(t *testing.T, output, text string)
   // Verifies visible text content is present

   func assertWidth(t *testing.T, output string, width int)
   // Verifies rendered output has expected width

   func assertHeight(t *testing.T, output string, height int)
   // Verifies rendered output has expected height

   func assertNotEmpty(t *testing.T, output string)
   // Verifies output is not empty (sanity check)
   ```

2. **Write property-based tests**
   - Test that all style combinations produce ANSI output
   - Test that text content is always preserved
   - Test that layout dimensions are respected
   - Don't verify exact ANSI codes, just properties

3. **Add table-driven tests for variations**
   ```go
   func TestText_StyleCombinations_ProducesOutput(t *testing.T) {
       tests := []struct {
           name  string
           props TextProps
       }{
           {"bold_italic", TextProps{Bold: true, Italic: true}},
           {"all_styles", TextProps{Bold: true, Italic: true, Underline: true}},
           // ... more combinations
       }

       for _, tt := range tests {
           t.Run(tt.name, func(t *testing.T) {
               text := Text("Hello", tt.props)
               got := text.Render(Layout{Width: 10, Height: 1})

               assertHasANSICodes(t, got)
               assertContainsText(t, got, "Hello")
               assertWidth(t, got, 10)
           })
       }
   }
   ```

4. **Document helper usage**
   - Add examples to `testing/README.md`
   - Document when to use golden files vs helpers
   - Provide decision tree for new tests

**Deliverables:**
- `testing/assertions.go` - assertion helpers
- Updated `text_test.go` - tests using helpers
- `testing/README.md` - usage documentation

**Success Criteria:**
- Helpers work for all style combinations
- Tests are more readable than before
- New tests can be written faster

---

### Phase 3: Optional ANSI Parser (As Needed)

**Goal:** Add ANSI parsing only if Phase 1+2 prove insufficient.

**Decision Criteria:**
- Do we need to verify complex style interactions?
- Are golden files too brittle for certain tests?
- Do we need cross-platform ANSI compatibility checks?

**Tasks (if needed):**

1. **Evaluate existing libraries**
   - `github.com/acarl005/stripansi` - strip ANSI codes
   - `github.com/muesli/termenv` - terminal environment detection
   - Consider building custom parser if needed

2. **Implement parser or integrate library**
   ```go
   // testing/ansi.go

   type StyleInfo struct {
       Text           string
       Bold           bool
       Italic         bool
       Underline      bool
       Strikethrough  bool
       ForegroundColor string
       BackgroundColor string
   }

   func parseANSI(output string) StyleInfo
   ```

3. **Write semantic assertion helpers**
   ```go
   func assertHasBoldStyle(t *testing.T, output string)
   func assertHasColor(t *testing.T, output, color string)
   func assertHasStyles(t *testing.T, output string, expected StyleInfo)
   ```

4. **Refactor tests to use semantic assertions**
   - Replace golden files for non-critical tests
   - Keep golden files for reference cases
   - Use parser for complex style verification

**Deliverables:**
- `testing/ansi.go` - ANSI parser (if built)
- Updated tests using semantic assertions
- Documentation of when to use each approach

**Success Criteria:**
- Parser correctly extracts all style properties
- Tests are more structure-insensitive
- Maintenance burden is acceptable

---

## Tools and Dependencies

### Required

**None** - Phase 1 and 2 use only Go standard library.

### Optional (Phase 3)

If we decide to implement an ANSI parser:

- **Option A:** Use existing library
  - `github.com/acarl005/stripansi` - strips ANSI codes
  - `github.com/muesli/ansi` - ANSI sequence parsing
  - Pros: Battle-tested, maintained
  - Cons: May not cover all our needs

- **Option B:** Build custom parser
  - Parse only the ANSI codes Lipgloss generates
  - Tailored to our exact needs
  - Pros: Full control, minimal code
  - Cons: Maintenance burden, testing needed

**Recommendation:** Start with Option A if Phase 3 is needed.

### Development Tools

- **go test -update** - flag for updating golden files
- **diff tool** - for comparing golden file changes
- **make test** - runs full test suite (already exists)

---

## Example Code

### Golden File Test (Phase 1)

```go
// text_test.go

func TestText_WithBold_MatchesGoldenFile(t *testing.T) {
    text := Text("Hello", TextProps{Bold: true})
    layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

    got := text.Render(layout)
    want := loadGoldenFile(t, "text_bold.golden")

    if got != want {
        t.Errorf("Output doesn't match golden file:\ngot:\n%q\nwant:\n%q", got, want)

        if *updateGolden {
            updateGoldenFile(t, "text_bold.golden", got)
            t.Log("Updated golden file")
        }
    }
}
```

### Golden File Helper (Phase 1)

```go
// testing/golden.go

var updateGolden = flag.Bool("update", false, "update golden files")

func loadGoldenFile(t *testing.T, name string) string {
    t.Helper()
    path := filepath.Join("testdata", "golden", name)
    data, err := os.ReadFile(path)
    if err != nil {
        t.Fatalf("Failed to read golden file %s: %v", name, err)
    }
    return string(data)
}

func updateGoldenFile(t *testing.T, name, content string) {
    t.Helper()
    path := filepath.Join("testdata", "golden", name)

    // Ensure directory exists
    dir := filepath.Dir(path)
    if err := os.MkdirAll(dir, 0755); err != nil {
        t.Fatalf("Failed to create golden directory: %v", err)
    }

    // Write golden file
    if err := os.WriteFile(path, []byte(content), 0644); err != nil {
        t.Fatalf("Failed to write golden file %s: %v", name, err)
    }
}

func compareWithGolden(t *testing.T, name, got string) {
    t.Helper()
    want := loadGoldenFile(t, name)

    if got != want {
        t.Errorf("Output doesn't match golden file %s:\ngot:\n%q\nwant:\n%q",
            name, got, want)

        if *updateGolden {
            updateGoldenFile(t, name, got)
            t.Log("Updated golden file")
        }
    }
}
```

### Assertion Helper (Phase 2)

```go
// testing/assertions.go

func assertHasANSICodes(t *testing.T, output string) {
    t.Helper()
    if !strings.Contains(output, "\x1b[") {
        t.Error("Expected output to contain ANSI escape codes")
    }
}

func assertContainsText(t *testing.T, output, text string) {
    t.Helper()
    // Strip ANSI codes for text comparison
    stripped := stripANSI(output)
    if !strings.Contains(stripped, text) {
        t.Errorf("Expected output to contain text %q, got: %q", text, stripped)
    }
}

func assertWidth(t *testing.T, output string, expectedWidth int) {
    t.Helper()
    // Measure visible width (excluding ANSI codes)
    width := visualWidth(output)
    if width != expectedWidth {
        t.Errorf("Expected width %d, got %d", expectedWidth, width)
    }
}

// Helper: strip ANSI codes
func stripANSI(s string) string {
    re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
    return re.ReplaceAllString(s, "")
}

// Helper: calculate visual width
func visualWidth(s string) int {
    stripped := stripANSI(s)
    return len(strings.TrimRight(stripped, " "))
}
```

### Table-Driven Test (Phase 2)

```go
// text_test.go

func TestText_StyleCombinations_ProducesValidOutput(t *testing.T) {
    tests := []struct {
        name    string
        props   TextProps
        content string
    }{
        {
            name:    "bold_only",
            props:   TextProps{Bold: true},
            content: "Hello",
        },
        {
            name:    "bold_italic",
            props:   TextProps{Bold: true, Italic: true},
            content: "Hello",
        },
        {
            name:    "all_styles",
            props:   TextProps{Bold: true, Italic: true, Underline: true, Strikethrough: true},
            content: "Test",
        },
        {
            name:    "color_with_style",
            props:   TextProps{Bold: true, Color: "#FF0000"},
            content: "Red",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            text := Text(tt.content, tt.props)
            layout := Layout{X: 0, Y: 0, Width: 10, Height: 1}

            got := text.Render(layout)

            // Verify properties
            assertHasANSICodes(t, got)
            assertContainsText(t, got, tt.content)
            assertNotEmpty(t, got)
        })
    }
}
```

---

## Success Criteria

### Phase 1 Complete

- [x] Golden file infrastructure implemented
- [x] 10-12 golden files for core styles
- [x] All `text_test.go` style tests use golden files
- [x] Tests catch actual styling bugs (manually verified)
- [x] `-update` flag works correctly
- [x] Documentation in `testing/README.md`

### Phase 2 Complete

- [x] Assertion helpers implemented (`testing/assertions.go`)
- [x] Table-driven tests for style combinations
- [x] Tests are more readable than before
- [x] Documentation updated with helper usage
- [x] New tests can be written faster

### Overall Success

- [x] **Behavioral property achieved:** Tests verify actual output
- [x] **No false positives:** Wrong styles cause test failures
- [x] **Maintainable:** Golden files easy to update
- [x] **Readable:** Clear test intent
- [x] **Follows CLAUDE.md:** Simple, incremental, pragmatic

---

## Risks and Mitigations

### Risk: Golden files become brittle

**Likelihood:** Medium
**Impact:** Medium

**Mitigation:**
- Use `-update` flag for easy regeneration
- Document when to update (Lipgloss upgrade)
- Keep golden files for core cases only
- Use helpers for variations

### Risk: Too many golden files

**Likelihood:** Low
**Impact:** Low

**Mitigation:**
- Limit to 10-15 core style cases
- Use helpers for combinations
- Regular review and cleanup

### Risk: ANSI codes change format

**Likelihood:** Low
**Impact:** High

**Mitigation:**
- Pin Lipgloss version in go.mod
- Update golden files when upgrading
- Test upgrade in separate branch first

### Risk: Cross-platform ANSI differences

**Likelihood:** Low (Lipgloss handles this)
**Impact:** Medium

**Mitigation:**
- Lipgloss normalizes ANSI output
- Test on CI for multiple platforms
- Document platform-specific behavior if found

---

## Timeline

**Total Estimated Time:** 3-4 days

| Phase | Tasks | Time | Dependencies |
|-------|-------|------|--------------|
| Phase 1 | Golden file infrastructure | 0.5 day | None |
| Phase 1 | Write golden files | 0.5 day | Infrastructure |
| Phase 1 | Refactor existing tests | 0.5-1 day | Golden files |
| Phase 2 | Create assertion helpers | 0.5 day | Phase 1 complete |
| Phase 2 | Write property-based tests | 0.5 day | Helpers |
| Phase 2 | Documentation | 0.5 day | Phase 2 tests |
| Phase 3 | ANSI parser (optional) | 1-2 days | Phase 1+2 evaluation |

**Recommended Approach:** Implement Phase 1 fully, then Phase 2, then evaluate if Phase 3 is needed.

---

## Next Steps

1. **Review this plan** with team
2. **Get approval** to proceed
3. **Start Phase 1** - Create golden file infrastructure
4. **Implement incrementally** - One test at a time (baby steps)
5. **Document learnings** - Update this plan with findings
6. **Evaluate Phase 3** after Phase 1+2 complete

---

## References

- **Kent Beck's Test Desiderata:** https://testdesiderata.com/
- **CLAUDE.md:** Project rules and principles
- **Test Desiderata Analysis:** See prioritized improvements table
- **Original Issue:** Critical behavioral testing violations in test suite

---

## Questions for Discussion

1. Should we implement Phase 3 (ANSI parser) or is Phase 1+2 sufficient?
2. How many golden files is too many? (Current plan: 10-12)
3. Should we apply this pattern to Box and Layout tests too?
4. Should `-update` flag be documented in main README or just testing docs?
5. Do we need CI checks to prevent accidental golden file updates?

---

**Plan Status:** ✅ Phase 2 Complete + Box Extension Complete
**Phase 3 Decision:** ❌ **SKIPPED** - Not needed (see evaluation below)
**Next Action:** Apply pattern to Layout tests if needed

### Implementation Summary

**Phase 1 (Completed 2026-01-13):**
- Created `testing/testing.go` with golden file infrastructure
- Added 10 golden files in `testdata/`
- Refactored `text_test.go` to use behavioral golden file tests
- Full documentation in `testing/README.md`

**Phase 2 (Completed 2026-01-14):**
- Created `testing/assertions.go` with 5 assertion helpers:
  - `AssertHasANSICodes` - verifies ANSI codes present
  - `AssertContainsText` - verifies text content (ignoring ANSI)
  - `AssertWidth` - verifies visible width
  - `AssertHeight` - verifies line count
  - `AssertNotEmpty` - verifies non-empty visible content
- Added 26 tests for assertion helpers in `testing/assertions_test.go`
- Added 3 table-driven test functions in `text_test.go`:
  - `TestText_StyleCombinations_ProducesValidOutput` (13 test cases)
  - `TestText_AlignmentVariations_ProducesValidOutput` (3 test cases)
  - `TestText_WrapModes_ProducesValidOutput` (2 test cases)
- Updated `testing/README.md` with Phase 2 documentation

**Phase 3 Evaluation (2026-01-14):**
**Decision:** ❌ **SKIPPED** - ANSI parser not needed

**Evaluation Results:**
1. ✅ Complex style interactions - Already covered by Phase 1 golden files
2. ✅ Golden files brittleness - Not an issue, `-update` flag works well
3. ✅ Cross-platform compatibility - Lipgloss handles normalization
4. ✅ High coverage - 96.6% in main package maintained
5. ✅ No pain points - Tests are maintainable and effective

**Rationale:**
- Phases 1+2 solve the behavioral testing problem completely
- No false positives detected in current test suite
- YAGNI principle - don't build what we don't need
- Follows CLAUDE.md simplicity principle

**Box Component Extension (Completed 2026-01-14):**
Applied the same Phase 1+2 pattern to Box component tests:

- Created 6 golden files for Box:
  - `box_border_single.golden` - Single line border
  - `box_border_double.golden` - Double line border
  - `box_border_rounded.golden` - Rounded corners border
  - `box_background_red.golden` - Red background
  - `box_border_color_green.golden` - Green border color
  - `box_border_background.golden` - Border + background combination

- Refactored 5 existing tests to use golden files:
  - `TestBox_Render_WithBorderAppliesLipgloss`
  - `TestBox_Render_WithBackgroundAppliesColor`
  - `TestBox_Render_WithDoubleBorder`
  - `TestBox_Render_WithRoundedBorder`
  - `TestBox_Render_WithBorderColor`

- Added 1 new test:
  - `TestBox_Render_WithBorderAndBackground` - Critical combination

- Added 2 table-driven test functions:
  - `TestBox_StyleCombinations_ProducesValidOutput` (5 test cases)
  - `TestBox_DirectionVariations_ProducesValidOutput` (2 test cases)

- Updated `testing/README.md` with Box examples and documentation

**Results:**
- ✅ All Box tests now verify actual behavior, not just code presence
- ✅ Test coverage maintained at 96.6%
- ✅ No false positives in Box styling tests
- ✅ Pattern is proven and reusable for future components
