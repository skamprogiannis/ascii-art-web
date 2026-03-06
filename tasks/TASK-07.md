# Agent: previous | LLM: unknown

# TASK-07: Complex Newline Patterns

**Status:** COMPLETED  
**Priority:** P1 (High)  
**Golden Tests:** GT-007 to GT-009, GT-024, GT-025  
**Dependencies:** TASK-01 to TASK-06  
**Estimated Time:** 20 minutes

## Objective
Handle multiple consecutive newlines and multiple text segments.

## Requirements
1. Single `\n` → 1 blank line
2. Double `\\n\\n` → 2 blank lines
3. Triple `\\n\\n\\n` → 3 blank lines
4. Three segments: `A\\nB\\nC` → 8 + blank + 8 + blank + 8
5. Consecutive newlines: `A\\n\\nB` → 8 + 2 blanks + 8

## TDD Workflow

### RED Phase

**Run integration tests (should fail):**
```bash
cd test && go test -v -run "TestGT00[7-9]|TestGT02[4-5]"
```

### GREEN Phase

**Analyze split behavior:**
- `"\\n"` → splits to `["", ""]` → should print 1 blank (not 2)
- `"\\n\\n"` → splits to `["", "", ""]` → should print 2 blanks (not 3)
- `"A\\n\\nB"` → splits to `["A", "", "B"]` → 8 + blank + blank + 8

**Fix logic:**
- Problem: Last empty element from split causes extra blank
- Solution: Skip last element if it's empty AND input ends with `\n`
- OR: Different approach - handle blank lines between segments correctly

**Test each case carefully:**
- Pure newlines vs text with newlines behave differently
- May need special handling for trailing empty element

### REFACTOR Phase

**Verify:**
```bash
cd test && go test -v -run "TestGT00[7-9]|TestGT02[4-5]"
go run . "\\n" → 1 blank line
go run . "\\n\\n" → 2 blank lines
go run . "A\\nB\\nC" → 8 + blank + 8 + blank + 8
go run . "A\\n\\nB" → 8 + 2 blanks + 8
```

## Acceptance Criteria
- [ ] GT-007 to GT-009 pass
- [ ] GT-024 to GT-025 pass
- [ ] Single newline = 1 blank
- [ ] Multiple newlines = multiple blanks
- [ ] Multiple segments separated correctly
- [ ] All previous tests still pass

## Next Task
TASK-08: Special Characters
