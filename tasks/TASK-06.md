# Agent: previous | LLM: unknown

# TASK-06: Newline Handling

**Status:** COMPLETED  
**Priority:** P0 (Critical)  
**Golden Tests:** GT-011 to GT-013, GT-023  
**Dependencies:** TASK-01 to TASK-05  
**Estimated Time:** 20 minutes

## Objective
Handle newlines at different positions correctly.

## Requirements
1. Newline at start: blank line first, then 8 lines
2. Newline at end: 8 lines, then blank line
3. Newlines at both ends: blank + 8 + blank
4. Between text: 8 + blank + 8

## CRITICAL: strings.Split Behavior

**Understanding the Problem:**

When you split by `\n`, Go's `strings.Split` creates a trailing empty element:

```go
strings.Split("\n", "\n")           → ["", ""]           // 2 elements!
strings.Split("Hello\n", "\n")      → ["Hello", ""]      // trailing empty
strings.Split("Hello\nWorld", "\n") → ["Hello", "World"] // no empty between
strings.Split("A\n\nB", "\n")       → ["A", "", "B"]     // middle empty is real
```

**The Issue:**
- `"\n"` should print 1 blank line, but split gives 2 elements
- The last empty element is a split artifact, not a real newline
- BUT middle empty elements ARE real blank lines

**The Solution:**

1. **Skip the last empty element** (it's always a split artifact)
2. **Print empty elements as blank lines** (they represent `\n`)
3. **Add blank separator after non-empty segments** (if not last)

## TDD Workflow

### RED Phase

**Run integration tests (will fail):**
```bash
cd test && go test -v -run "TestGT01[1-3]|TestGT023"
```

### GREEN Phase

**Implement in `core/render_art.go`:**

1. **Split input by newline:**
   - Use `strings.Split(input, "\n")`
   - This creates array of segments

2. **Loop through segments with index:**
   - Need index to detect last element

3. **Skip last empty element:**
   - Check: `if i == len(lines)-1 && line == ""`
   - This is the split artifact, skip it

4. **Handle empty segments:**
   - Empty segment = blank line
   - Print: `fmt.Println()`

5. **Handle non-empty segments:**
   - Call printLine to render 8 lines
   - Add blank separator after (if not last segment)
   - Check: `if i < len(lines)-1`

**Key Logic Pattern:**
```
for i, line := range lines {
    // Skip last empty (split artifact)
    if i == len(lines)-1 && line == "" {
        continue
    }
    
    // Empty = blank line
    if line == "" {
        print blank
    }
    // Non-empty = 8 lines + separator
    else {
        print 8 lines
        if not last: print blank separator
    }
}
```

### REFACTOR Phase

**Verify:**
```bash
cd test && go test -v -run "TestGT01[1-3]|TestGT023"
go run . "\\nHello" → 9 lines (blank + 8)
go run . "Hello\\n" → 9 lines (8 + blank)
go run . "\\nHello\\n" → 10 lines (blank + 8 + blank)
go run . "Hello\\nWorld" → 17 lines (8 + blank + 8)
```

## Common Mistakes to Avoid

❌ **Don't skip ALL empty elements** - middle ones are real blank lines  
❌ **Don't add separator after empty elements** - they're already blanks  
❌ **Don't print the last empty element** - it's a split artifact  
✅ **Do skip ONLY the last empty element**  
✅ **Do add separator after non-empty segments**  

## Acceptance Criteria
- [ ] GT-011 to GT-013 pass
- [ ] GT-023 passes
- [ ] `"\n"` produces 1 blank (not 2)
- [ ] `"Hello\n"` produces 9 lines (8 + blank)
- [ ] `"Hello\nWorld"` produces 17 lines (8 + blank + 8)
- [ ] All previous tests still pass

## Next Task
TASK-07: Complex Newline Patterns
