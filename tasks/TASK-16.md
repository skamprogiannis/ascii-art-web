# TASK-16: Color Feature Implementation

**Status:** COMPLETED
**Priority:** P1 (High)
**Dependencies:** TASK-09
**Estimated Time:** 60 minutes

## Objective

Implement the `--color=<color> <substring>` flag to color ASCII art output while preserving existing behavior.

## Requirements

1. Support `--color=<color>` flag with optional substring argument
2. If no substring is provided, color the entire string
3. Accept color names and color codes (hex, rgb, hsl)
4. Invalid formats return `Usage: go run . [OPTION] [STRING] [BANNER]`
5. Preserve banner selection and single-argument execution

## TDD Workflow

### RED Phase

- Added audit-driven manual cases for color flags and substrings

### GREEN Phase

- Implemented color parsing in `internal/color/color.go`
- Updated `internal/parser/parse_input.go` for `--color=` parsing and substring handling
- Updated `internal/render/render_art.go` to apply ANSI colors per substring
- Updated `main.go` to validate color and orchestrate rendering

### REFACTOR Phase

- Simplified parsing flow and removed debug output
- Normalized rendering logic to preserve newline semantics

## Acceptance Criteria

- [x] Color flag renders ANSI colored output
- [x] Substring coloring works for multiple occurrences
- [x] Invalid formats print usage
- [x] Works with banner selection and single-string usage
- [x] Build succeeds with `go build`
