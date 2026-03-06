# TASK-01: Argument Validation

**Agent:** previous | **LLM:** unknown

**Status:** COMPLETED
**Priority:** P0 (Critical)
**Golden Tests:** GT-001 to GT-005

## Objective

Implement argument validation using modular architecture with unit tests.

## Requirements

1. Accept 1 or 2 command-line arguments (Updated in TASK-12)
2. Reject invalid argument counts
3. Return usage message when invalid
4. Use modular structure (core package)

## TDD Workflow

### RED Phase

- Write tests for argument counts.

### GREEN Phase

- Implement `ParseInput` in `core/parse_input.go`.
- Update `main.go` to handle usage message.

### REFACTOR Phase

- Ensure code follows Go conventions.

## Acceptance Criteria

- [x] `core/parse_input.go` exists.
- [x] Unit tests pass.
- [x] GT-001 to GT-005 pass.
