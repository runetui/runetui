# Strategy for Verifiable Examples (E2E/Acceptance)

## Objective

Create a system that allows verifying that examples in `examples/` work correctly and produce expected output, similar to end-to-end or acceptance tests.

## Proposed Approaches

### Option 1: Example Tests with Snapshot Testing (Recommended)

**Advantages:**
- Simple to implement
- Reuses existing infrastructure (`testing` package)
- Easy to maintain
- Allows verifying visual changes

**Implementation:**
- Create `example_test.go` in each example directory
- Use `testing.RenderToString()` to capture output
- Use `testing.AssertSnapshot()` to compare with golden files
- Run with `go test ./examples/...`

**Structure:**
```
examples/
├── hello/
│   ├── main.go
│   ├── example_test.go
│   └── testdata/
│       └── hello_output.golden
└── streaming/
    ├── main.go
    ├── example_test.go
    └── testdata/
        └── streaming_output.golden
```

### Option 2: Execution Tests with Output Capture

**Advantages:**
- Verifies that the example actually runs without errors
- Can capture multiple frames (for interactive examples)
- Closer to real execution

**Disadvantages:**
- More complex (requires mocking terminal)
- Can be fragile with timing

**Implementation:**
- Run the example in a controlled environment
- Capture stdout/stderr
- Compare with golden files or assertions

### Option 3: Hybrid: Unit Tests + Execution Verification

**Advantages:**
- Combines the best of both approaches
- Verifies both the component and execution

**Implementation:**
- Unit tests to verify the component tree
- Execution tests to verify it runs without errors

## Recommendation: Option 1 (Snapshot Testing)

This option is the most practical because:
1. We already have the infrastructure (`testing` package)
2. Fast to execute (doesn't require real terminal)
3. Easy to maintain and update
4. Allows easily verifying visual changes

## Proposed Implementation

### File Structure

```
examples/
├── hello/
│   ├── main.go              # Original example
│   ├── example_test.go       # Example test
│   └── testdata/
│       └── hello_snapshot.golden
└── streaming/
    ├── main.go
    ├── example_test.go
    └── testdata/
        └── streaming_snapshot.golden
```

### Test Pattern

Each example would have a test that:
1. Extracts the component function from the example
2. Renders with standard dimensions
3. Compares with snapshot

### Implementation Example

See `examples/hello/example_test.go` for a complete example.

## Makefile Commands

Add targets to Makefile:

```makefile
test-examples: ## Run example tests
	go test ./examples/... -v

test-examples-update: ## Update example snapshots
	go test ./examples/... -update
```

## Advantages of This Approach

1. **Automatic Verification**: Examples are verified in CI/CD
2. **Living Documentation**: Snapshots serve as visual documentation
3. **Regression Detection**: Unexpected changes are automatically detected
4. **Easy Updates**: `-update` flag to update snapshots when intentionally changed
5. **Fast**: Doesn't require real terminal, only rendering

## Limitations

- Doesn't verify real user interaction
- Doesn't verify temporal behavior (for examples with state)
- Requires extracting the component function from the example

## Next Steps

1. Implement tests for `examples/hello` ✅ (Done)
2. Implement tests for `examples/streaming` (more complex, requires state mock)
3. Add targets to Makefile ✅ (Done)
4. Document in README how to run example tests ✅ (Done)
