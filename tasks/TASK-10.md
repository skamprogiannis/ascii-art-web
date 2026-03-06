# TASK-10: Update Project Guidelines & Structure

**Status:** COMPLETED
**Priority:** P1 (High)
**Dependencies:** TASK-01 to TASK-09
**Estimated Time:** 15 minutes

## Objective

Update AGENTS.md to reflect new simplicity rules and move banner files to a dedicated directory.

## Requirements

1. Updated AGENTS.md: Removed "no closures, no callbacks" from Direct Logic.
2. Updated AGENTS.md: Removed restriction on multiple arguments.
3. Updated AGENTS.md: Updated file structure to include `banners/` directory.
4. Move banner files (standard.txt, shadow.txt, thinkertoy.txt) to `banners/`.

## TDD Workflow

### RED Phase

N/A (Documentation and structure change)

### GREEN Phase

- Check AGENTS.md for consistency.
- Check `banners/` directory exists and contains all files.

### REFACTOR Phase

- Ensure no code relies on banner files being in the root directory (will be handled in TASK-11).

## Acceptance Criteria

- [x] AGENTS.md updated.
- [x] Banner files moved to `banners/`.
- [x] AI_LOG.md updated.
