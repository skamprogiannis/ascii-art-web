# Agent: previous | LLM: unknown

# TASK-02: Input Processing

**Status:** COMPLETED  
**Priority:** P0 (Critical)  
**Golden Tests:** GT-006 to GT-010  
**Dependencies:** TASK-01  
**Estimated Time:** 15 minutes

## Objective
Add newline conversion and empty string handling.

## Requirements
1. Convert literal `\\n` to actual newline character
2. Handle empty string input (no output)
3. Maintain argument validation from TASK-01

## TDD Workflow

### RED Phase

**1. Add unit tests to `core/parse_input_test.go`:**
- Test: Input `"Hello\\nWorld"` converts to `"Hello\nWorld"`
- Test: Empty string `""` is valid input (returns true)

**2. Run tests (should fail):**
```bash
cd core && go test -v
```

### GREEN Phase

**1. Update `core/parse_input.go`:**
- Import: Add `strings` package
- After getting `os.Args[1]`, use `strings.ReplaceAll` to convert `\\n` to `\n`
- Return converted string

**2. Update `main.go`:**
- After ParseInput succeeds, check if input is empty string
- If empty, return immediately (no output)
- Otherwise continue (will implement rendering later)

### REFACTOR Phase

**Verify:**
```bash
cd core && go test -v
cd ../test && go test -v -run "TestTask0[1-2]"
go run . "" → no output
go run . "\\n" → will handle in next task
```

## Acceptance Criteria
- [ ] Literal `\\n` converts to actual newline
- [ ] Empty string produces no output
- [ ] GT-006 to GT-010 pass
- [ ] All TASK-01 tests still pass
- [ ] Unit tests updated and passing

## Next Task
TASK-03: Banner Loading
