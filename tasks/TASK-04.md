# Agent: previous | LLM: unknown

# TASK-04: Character Rendering

**Status:** COMPLETED  
**Priority:** P0 (Critical)  
**Golden Tests:** GT-014 to GT-017  
**Dependencies:** TASK-01, TASK-02, TASK-03  
**Estimated Time:** 30 minutes

## Objective
Implement ASCII art rendering for single characters.

## Requirements
1. Print exactly 8 lines per character
2. Calculate correct banner index: `(char - 32) * 9 + 1 + row`
3. Support ASCII 32-126
4. Handle uppercase, lowercase, digits, space

## TDD Workflow

### RED Phase

**1. Run integration tests (should fail):**
```bash
cd test && go test -v -run TestTask05
```

### GREEN Phase

**1. Create `core/render_art.go`:**
- Package: `core`
- Function: `RenderArt(input string, banner []string)`
- Split input by `\n` using `strings.Split`
- For each line:
  - If empty: print blank line with `fmt.Println()`
  - If not empty: call helper to print 8 rows

**2. Add helper function in same file:**
- Function: `printLine(text string, banner []string)` (unexported)
- Loop: row from 0 to 7
- For each row:
  - Loop through each character in text
  - Calculate index: `(char - 32) * 9 + 1 + row`
  - Check bounds: `index >= 0 && index < len(banner)`
  - Print: `banner[index]` using `fmt.Print`
  - After all chars: `fmt.Println()` for newline

**3. Update `main.go`:**
- After loading banner, call `core.RenderArt(input, banner)`

### REFACTOR Phase

**Verify:**
```bash
cd test && go test -v -run "TestTask0[1-5]"
go run . "A" → prints 8 lines
go run . "a" → prints 8 lines
go run . "1" → prints 8 lines
go run . " " → prints 8 lines (space art)
```

## Acceptance Criteria
- [ ] `core/render_art.go` created
- [ ] RenderArt function implemented
- [ ] Helper function for printing lines
- [ ] GT-014 to GT-017 pass
- [ ] Index formula correct
- [ ] All previous tests still pass

## Next Task
TASK-05: Word Rendering
