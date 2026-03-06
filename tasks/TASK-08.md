# Agent: previous | LLM: unknown

# TASK-08: Special Characters

**Status:** COMPLETED  
**Priority:** P1 (High)  
**Golden Tests:** GT-026 to GT-029  
**Dependencies:** TASK-01 to TASK-07  
**Estimated Time:** 15 minutes

## Objective
Verify all special characters and symbols render correctly.

## Requirements
1. Punctuation: `!"#$%&'()*+,-./`
2. Symbols: `:;<=>?@`
3. Brackets: `[]{}()`
4. Special: `[\]^_`{|}~` (including backslash)

## TDD Workflow

### RED Phase

**Run integration tests (should pass if index formula correct):**
```bash
cd test && go test -v -run "TestGT02[6-9]"
```

### GREEN Phase

**Verify index formula:**
- Formula: `(char - 32) * 9 + 1 + row`
- ASCII 33 ('!'): (33-32)*9+1 = 10 → lines 10-17 ✓
- ASCII 126 ('~'): (126-32)*9+1 = 847 → lines 847-854 ✓
- Banner has 855 lines (0-854) ✓

**If tests fail:**
- Check bounds validation: `index >= 0 && index < len(banner)`
- Verify no special case handling for specific characters
- Ensure formula is exactly as specified

**No code changes needed if:**
- Index formula correct from TASK-04
- Bounds checking in place
- No character filtering

### REFACTOR Phase

**Verify:**
```bash
cd test && go test -v -run "TestGT02[6-9]"
go run . "!@#$%^&*()" → 8 lines
go run . ":;<=>?@" → 8 lines
go run . "[]{}()" → 8 lines
go run . "[\\\\]^_\`{|}~" → 8 lines (note: backslash escaped)
```

## Acceptance Criteria
- [ ] GT-026 to GT-029 pass
- [ ] All punctuation renders
- [ ] All symbols render
- [ ] Brackets render
- [ ] Backslash renders
- [ ] No crashes for any printable ASCII
- [ ] All previous tests still pass

## Next Task
TASK-09: Full ASCII Range
