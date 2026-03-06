# Color Feature Implementation Plan

## Task: Add --color flag to ascii-art

**Branch:** add-color  
**Objective:** Implement `--color=<color> <substring>` flag functionality
**Status:** Completed

---

## Requirements Analysis

### Input Format

- `go run . --color=red kit "a king kitten have kit"` → colors "kit" in "kitten" and "kit"
- If no substring specified: color entire string
- Must work with existing banner arguments: `go run . --color=red hello standard`

### Color Support

- Basic colors: red, green, yellow, blue, orange, etc.
- Bonus: `#RRGGBB`, `rgb(r,g,b)`, `hsl(h,s,l)` notations

### Edge Cases

- Invalid flag format → `Usage: go run . [OPTION] [STRING] [BANNER]`
- Invalid color → usage message
- Multiple flags → first valid flag wins (based on audit tests)
- Case insensitive color names

### Test Cases (from audit)

1. `--color red "banana"` → usage message (missing =)
2. `--color=red "hello world"` → colors entire string
3. `--color=green "1 + 1 = 2"` → colors entire string
4. `--color=yellow "(%&) ??"` → colors entire string
5. `--color=orange GuYs "HeY GuYs"` → colors "GuYs"
6. `--color=blue B "RGB()"` → colors "B"
7. `--color=blue B "RGB()"` → colors just B
8. Random strings with random colors

---

## Implementation Plan

### Phase 1: Input Parsing (Modify internal/parser/parse_input.go)

1. **Add flag detection:**
   - Check if any argument contains `--color=`
   - Parse color value and optional substring

2. **Parse color flag format:**
   - Valid: `--color=<color>` or `--color=<color> <substring>`
   - Invalid: `--color red` (missing =)

3. **Extract components:**
   - Color name/code
   - Substring to color (optional)
   - Main string to display

4. **Handle multi-argument cases:**
   - `--color=red hello` → color="red", substring="", text="hello"
   - `--color=red kit "a king kitten"` → color="red", substring="kit", text="a king kitten"

### Phase 2: Color Processing (New: internal/color/color.go)

1. **Color parsing:**
   - ANSI names: red, green, yellow, blue, magenta, cyan, white, black
   - Hex: #RRGGBB
   - RGB: rgb(r,g,b)
   - HSL: hsl(h,s,l)

2. **ANSI escape codes:**
   - Map color names to ANSI codes (foreground)
   - Support bright variants

3. **Color lookup:**
   ```go
   var colorMap = map[string]string{
       "red":   "\033[31m",
       "green": "\033[32m",
       "yellow": "\033[33m",
       "blue":  "\033[34m",
       "orange": "\033[38;5;208m",
       // ...
   }
   ```

### Phase 3: Rendering (Modify internal/render/render_art.go)

1. **Substring matching:**
   - Find all occurrences of substring in input
   - Case: exact match (based on example "kit" matches "kit" in "kitten")

2. **Build colored output:**
   - For each row of ASCII art:
     - Track character positions
     - Insert color codes at substring start
     - Insert reset codes at substring end
   - Use ANSI escape sequences: `\033[0m` to reset

3. **No substring specified:**
   - Color entire string (wrap all chars with color codes)

### Phase 4: Integration (main.go)

1. **Updated ParseInput returns:**
   - input string
   - banner name
   - color (optional)
   - substring to color (optional)
   - error

2. **Error handling:**
   - Invalid `--color` format → usage message
   - Unknown color → usage message

---

## File Changes

| File                             | Change                            |
| -------------------------------- | --------------------------------- |
| `internal/parser/parse_input.go` | Add flag parsing for `--color=`   |
| `internal/color/color.go`        | NEW: Color parsing and ANSI codes |
| `internal/render/render_art.go`  | Add substring coloring logic      |
| `main.go`                        | Update to pass color params       |

---

## Algorithm

### Substring Coloring

```
Input: "a king kitten have kit", substring="kit", color="red"

For each occurrence of "kit":
  - Position 9: "kit" in "kitten"
  - Position 25: "kit" in "kit"

Output (row example):
  [reset]a king [color]kit[reset]ten have [color]kit[reset]
```

### Edge Cases

- Overlapping substrings: first match wins
- Substring not found: no color applied
- Empty substring: color entire string

---

## Validation

Test with:

```bash
go run . --color=red "hello"              # whole string red
go run . --color=red kit "a kitten"       # "kit" in "kitten" red
go run . --color=red kit "kit kitten"    # both "kit"s red
go run . --color=blue B "RGB()"           # B is blue
go run . --color red "hello"              # usage (invalid format)
go run . --color=unknown "hello"         # usage (unknown color)
```

---

## Success Criteria

- [ ] All existing tests pass
- [ ] `--color=red "hello"` works
- [ ] `--color=red kit "a kitten"` colors "kit"
- [ ] `--color red "hello"` shows usage
- [ ] Works with banner: `--color=red hello standard`

---

## Completion Summary

- Added `internal/color/color.go` with ANSI and color-code parsing
- Updated `internal/parser/parse_input.go` for flag and banner parsing
- Updated `internal/render/render_art.go` for substring coloring
- Updated `main.go` for color validation and render path
- Added unit tests for parsing and color conversion
