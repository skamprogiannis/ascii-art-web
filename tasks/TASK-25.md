# TASK-25: Restore Explicit Form Submit Button

**Status:** COMPLETED
**Priority:** P1 (High)
**Dependencies:** TASK-24
**Estimated Time:** 15 minutes

## Objective
Re-add an explicit **Generate** submit button to the left form panel so the UI directly satisfies the requirement: *"a button that sends a POST request to `/ascii-art`."*

## Requirements
1. Add a visible submit button inside the existing form on the left card.
2. Keep the form target `action="/ascii-art"` and `method="POST"`.
3. Preserve current interactive behavior and styling consistency.
4. Ensure no regression in backend/API behavior.

## Implementation Steps
- ✅ Added `type="submit"` Generate button at the bottom of the left form card in `templates/index.html`.
- ✅ Added `.generate-btn` styles in `static/style.css` to match existing visual language.
- ✅ Kept existing progressive enhancement behavior (JS preview + normal form semantics).
- ✅ Re-ran test suite after change.

## Acceptance Criteria
- [x] Main page contains a clear button that submits POST to `/ascii-art`.
- [x] Button appears at the bottom of the left card.
- [x] Existing functionality remains intact.
- [x] Tests pass with `CGO_ENABLED=0 go test ./...`.
