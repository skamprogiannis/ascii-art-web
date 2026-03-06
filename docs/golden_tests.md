# Golden Tests Specification

## ASCII-Art Generator v1.3

### Purpose

This document defines the **source of truth** for ASCII-Art implementation. All code must pass these tests. If code fails any test, it is incorrect.

### Test Execution

```bash
cd ascii-art
go run . [ARGS]
```

### Test Summary

- **Total Tests:** 30
- **Level 1 (Audit):** 5 tests - Argument validation
- **Level 2 (Edge Cases):** 8 tests - Empty/newline scenarios
- **Level 3 (Core):** 10 tests - Basic rendering
- **Level 4 (Advanced):** 7 tests - Complex patterns

---

# LEVEL 1: AUDIT CASES (Argument Validation)

## Test #1: No Arguments

**ID:** GT-001  
**Category:** Audit Case  
**Difficulty:** Level 1  
**Priority:** P0 (Critical)

**Input:**

```bash
go run .
```

**Expected Output:**

```
Usage: go run . [COLOR] [SUBSTRING] [STRING] [BANNER]"
```

**Validation:**

- Exact usage message printed
- Exit code 0
- No crash or panic

---

## Test #2: Multiple Arguments (Valid)

**ID:** GT-002  
**Category:** Audit Case  
**Difficulty:** Level 1  
**Priority:** P0 (Critical)

**Input:**

```bash
go run . "Hello" "standard"
```

**Expected Output:**
(8 lines of ASCII art for "Hello")

**Validation:**

- 2 arguments are valid in v1.3
- Renders "Hello" using "standard" banner

---

## Test #3: Three Arguments (Invalid)

**ID:** GT-003  
**Category:** Audit Case  
**Difficulty:** Level 1  
**Priority:** P1 (High)

**Input:**

```bash
go run . "A" "B" "C"
```

**Expected Output:**

```
Usage: go run . [OPTION] [STRING] [BANNER]
```

**Validation:**

- Exact usage message printed
- Exit code 0

---

## Test #4: Empty Argument Plus Extra

**ID:** GT-004  
**Category:** Audit Case  
**Difficulty:** Level 1  
**Priority:** P1 (High)

**Input:**

```bash
go run . "" "standard"
```

**Expected Output:**
(No output)

**Validation:**

- 2 arguments are valid
- Empty string input produces no output

---

## Test #5: Many Arguments

**ID:** GT-005  
**Category:** Audit Case  
**Difficulty:** Level 1  
**Priority:** P2 (Medium)

**Input:**

```bash
go run . "1" "2" "3" "4" "5"
```

**Expected Output:**

```
Usage: go run . [OPTION] [STRING] [BANNER]
```

**Validation:**

- Any argument count > 3 rejected
- Exit code 0

---

# LEVEL 2: EDGE CASES (Empty/Newline Scenarios)

## Test #6: Empty String

**ID:** GT-006  
**Category:** Edge Case  
**Difficulty:** Level 2  
**Priority:** P0 (Critical)

**Input:**

```bash
go run . ""
```

**Expected Output:**
(No output - program exits immediately)

**Validation:**

- No output to stdout
- Exit code 0
- Program terminates without error

---

## Test #7: Single Newline

**ID:** GT-007  
**Category:** Edge Case  
**Difficulty:** Level 2  
**Priority:** P0 (Critical)

**Input:**

```bash
go run . "\n"
```

**Expected Output:**

```

```

(Single blank line)

**Validation:**

- Exactly 1 blank line printed
- No ASCII art rendered

---

## Test #8: Two Newlines

**ID:** GT-008  
**Category:** Edge Case  
**Difficulty:** Level 2  
**Priority:** P1 (High)

**Input:**

```bash
go run . "\n\n"
```

**Expected Output:**

```


```

(Two blank lines)

**Validation:**

- Exactly 2 blank lines printed
- No ASCII art rendered

---

## Test #9: Three Newlines

**ID:** GT-009  
**Category:** Edge Case  
**Difficulty:** Level 2  
**Priority:** P1 (High)

**Input:**

```bash
go run . "\n\n\n"
```

**Expected Output:**

```



```

(Three blank lines)

**Validation:**

- Exactly 3 blank lines printed
- No ASCII art rendered

---

## Test #10: Only Spaces

**ID:** GT-010  
**Category:** Edge Case  
**Difficulty:** Level 2  
**Priority:** P1 (High)

**Input:**

```bash
go run . "   "
```

**Expected Output:**
(8 lines with spaces rendered as empty columns)

**Validation:**

- 8 lines printed
- Each line contains space character art (empty columns)
- No visible characters

---

## Test #11: Newline at Start

**ID:** GT-011  
**Category:** Edge Case  
**Difficulty:** Level 2  
**Priority:** P1 (High)

**Input:**

```bash
go run . "\nHello"
```

**Expected Output:**

```

(blank line, then 8 lines of "Hello")
```

**Validation:**

- Blank line first (total 9 lines)
- Then 8 lines of ASCII art for "Hello"

---

## Test #12: Newline at End

**ID:** GT-012  
**Category:** Edge Case  
**Difficulty:** Level 2  
**Priority:** P1 (High)

**Input:**

```bash
go run . "Hello\n"
```

**Expected Output:**
(8 lines of "Hello")

**Validation:**

- Exactly 8 lines of ASCII art
- Trailing newline handled by shell

---

## Test #13: Newlines at Both Ends

**ID:** GT-013  
**Category:** Edge Case  
**Difficulty:** Level 2  
**Priority:** P2 (Medium)

**Input:**

```bash
go run . "\nHello\n"
```

**Expected Output:**

```

(blank, then 8 lines of "Hello")
```

**Validation:**

- Blank line first
- 8 lines of ASCII art
- Total 9 lines

---

# LEVEL 3: CORE FUNCTIONALITY (Basic Rendering)

## Test #14: Single Character - Uppercase

**ID:** GT-014  
**Category:** Core Functionality  
**Difficulty:** Level 3  
**Priority:** P0 (Critical)

**Input:**

```bash
go run . "A"
```

**Expected Output:**
(8 lines of ASCII art for "A")

**Validation:**

- Exactly 8 lines printed
- Character "A" rendered correctly
- No trailing newline after line 8

---

## Test #15: Single Character - Lowercase

**ID:** GT-015  
**Category:** Core Functionality  
**Difficulty:** Level 3  
**Priority:** P0 (Critical)

**Input:**

```bash
go run . "a"
```

**Expected Output:**
(8 lines of ASCII art for "a")

**Validation:**

- Exactly 8 lines printed
- Lowercase "a" rendered correctly

---

## Test #16: Single Character - Digit

**ID:** GT-016  
**Category:** Core Functionality  
**Difficulty:** Level 3  
**Priority:** P0 (Critical)

**Input:**

```bash
go run . "1"
```

**Expected Output:**
(8 lines of ASCII art for "1")

**Validation:**

- Exactly 8 lines printed
- Digit "1" rendered correctly

---

## Test #17: Single Character - Space

**ID:** GT-017  
**Category:** Core Functionality  
**Difficulty:** Level 3  
**Priority:** P0 (Critical)

**Input:**

```bash
go run . " "
```

**Expected Output:**
(8 lines with space character art)

**Validation:**

- Exactly 8 lines printed
- Space (ASCII 32) rendered as empty columns

---

## Test #18: Single Word

**ID:** GT-018  
**Category:** Core Functionality  
**Difficulty:** Level 3  
**Priority:** P0 (Critical)

**Input:**

```bash
go run . "Hello"
```

**Expected Output:**
(8 lines of ASCII art for "Hello")

**Validation:**

- Exactly 8 lines printed
- Characters concatenated horizontally
- No trailing newline after line 8

---

## Test #19: Two Words with Space

**ID:** GT-019  
**Category:** Core Functionality  
**Difficulty:** Level 3  
**Priority:** P0 (Critical)

**Input:**

```bash
go run . "Hello World"
```

**Expected Output:**
(8 lines with space between "Hello" and "World")

**Validation:**

- Space (ASCII 32) rendered as empty columns
- Words separated correctly
- 8 lines total

---

## Test #20: Lowercase Word

**ID:** GT-020  
**Category:** Core Functionality  
**Difficulty:** Level 3  
**Priority:** P1 (High)

**Input:**

```bash
go run . "hello"
```

**Expected Output:**
(8 lines of ASCII art for "hello")

**Validation:**

- Lowercase letters rendered correctly
- ASCII 97-122 supported

---

## Test #21: Mixed Case

**ID:** GT-021  
**Category:** Core Functionality  
**Difficulty:** Level 3  
**Priority:** P1 (High)

**Input:**

```bash
go run . "HeLLo"
```

**Expected Output:**
(8 lines of ASCII art for "HeLLo")

**Validation:**

- Mixed case rendered correctly
- Each character uses correct case

---

## Test #22: All Digits

**ID:** GT-022  
**Category:** Core Functionality  
**Difficulty:** Level 3  
**Priority:** P1 (High)

**Input:**

```bash
go run . "0123456789"
```

**Expected Output:**
(8 lines of ASCII art for all digits)

**Validation:**

- All digits 0-9 rendered correctly
- ASCII 48-57 supported

---

## Test #23: Simple Newline

**ID:** GT-023  
**Category:** Core Functionality  
**Difficulty:** Level 3  
**Priority:** P0 (Critical)

**Input:**

```bash
go run . "Hello\nWorld"
```

**Expected Output:**
(8 lines for "Hello", then 8 lines for "World")

**Validation:**

- Total 16 lines of output
- Literal `\n` converted to actual newline

---

# LEVEL 4: ADVANCED FEATURES (Complex Patterns)

## Test #24: Three Lines

**ID:** GT-024  
**Category:** Advanced  
**Difficulty:** Level 4  
**Priority:** P1 (High)

**Input:**

```bash
go run . "A\nB\nC"
```

**Expected Output:**
(8 lines for "A", 8 lines for "B", 8 lines for "C")

**Validation:**

- Three segments of 8 lines each (Total 24 lines)

---

## Test #25: Multiple Consecutive Newlines

**ID:** GT-025  
**Category:** Advanced  
**Difficulty:** Level 4  
**Priority:** P1 (High)

**Input:**

```bash
go run . "A\n\nB"
```

**Expected Output:**
(8 lines for "A", 1 blank line, 8 lines for "B")

**Validation:**

- Total 17 lines of output
- `\n\n` results in ONE blank line between art blocks

---

## Test #26: All Punctuation

**ID:** GT-026  
**Category:** Advanced  
**Difficulty:** Level 4  
**Priority:** P1 (High)

**Input:**

```bash
go run . "!\"#$%&'()*+,-./"
```

**Expected Output:**
(8 lines of ASCII art for punctuation marks)

**Validation:**

- All punctuation ASCII 33-47 rendered
- Characters concatenated correctly

---

## Test #27: Special Symbols

**ID:** GT-027  
**Category:** Advanced  
**Difficulty:** Level 4  
**Priority:** P1 (High)

**Input:**

```bash
go run . ":;<=>?@"
```

**Expected Output:**
(8 lines of ASCII art for symbols)

**Validation:**

- ASCII 58-64 rendered correctly
- Symbols concatenated horizontally

---

## Test #28: Brackets and Braces

**ID:** GT-028  
**Category:** Advanced  
**Difficulty:** Level 4  
**Priority:** P1 (High)

**Input:**

```bash
go run . "[]{}()"
```

**Expected Output:**
(8 lines of ASCII art for brackets)

**Validation:**

- All bracket types rendered correctly
- Proper pairing visible

---

## Test #29: Backslash and Special

**ID:** GT-029  
**Category:** Advanced  
**Difficulty:** Level 4  
**Priority:** P1 (High)

**Input:**

```bash
go run . "[\\]^_`{|}~"
```

**Expected Output:**
(8 lines of ASCII art for special characters)

**Validation:**

- Backslash (ASCII 92) rendered
- Pipe (ASCII 124) rendered
- Tilde (ASCII 126) rendered
- No escape sequence interpretation

---

## Test #30: Full ASCII Range

**ID:** GT-030  
**Category:** Advanced  
**Difficulty:** Level 4  
**Priority:** P2 (Medium)

**Input:**

```bash
go run . " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
```

**Expected Output:**
(8 lines of ASCII art for all printable ASCII characters 32-126)

**Validation:**

- All ASCII 32-126 characters rendered
- Extremely long horizontal output
- No crashes or truncation
- Complete character set coverage

---

## TDD Workflow

### Red-Green-Refactor Cycle

**RED Phase:**

1. Write test case
2. Run test
3. Verify it FAILS

**GREEN Phase:**

1. Write minimal code to pass
2. Run test
3. Verify it PASSES

**REFACTOR Phase:**

1. Clean up code
2. Run test again
3. Verify still PASSES

---

## Color Extension Tests

### Test C-01: Full String Color

**Input:**

```bash
go run . --color=red "hello"
```

**Expected Output:**

- All ASCII art characters wrapped with ANSI red color codes
- Output shape identical to non-colored "hello"

### Test C-02: Substring Color

**Input:**

```bash
go run . --color=red kit "a king kitten have kit"
```

**Expected Output:**

- Only occurrences of "kit" are colored
- Non-matching characters remain uncolored

### Test C-03: Invalid Flag Format

**Input:**

```bash
go run . --color red "banana"
```

**Expected Output:**

```
Usage: go run . [OPTION] [STRING] [BANNER]
```

---

## Success Criteria

All 30 tests must pass with:

- Exact output match
- Correct exit codes
- No crashes
- Deterministic results (same input = same output every time)

---

## Notes

- All tests use `standard.txt` banner by default
- Banner files exist in `banners/` directory (Updated in v1.3)
- Output must match byte-for-byte (no approximations)
- ASCII range: 32-126 (95 printable characters)
- Banner format: 8 lines per character, 855 total lines
- Index formula: `(ASCII_code - 32) * 9 + 1 + row`
