# TASK-26: Align Generate Button with Site Theme

**Status:** COMPLETED
**Priority:** P2 (Medium)
**Dependencies:** TASK-25
**Estimated Time:** 15 minutes

## Objective
Refine the explicit **Generate ASCII Art** submit button so it visually matches the rest of the dark glassmorphism UI instead of reading like a browser-default control.

## Requirements
1. Keep the button as a real `type="submit"` control.
2. Preserve the button placement at the bottom of the left form card.
3. Align spacing, typography, borders, and hover/focus states with the existing theme.
4. Avoid changing backend behavior or request flow.

## Implementation Steps
- ✅ Refined `.generate-btn` styling in `static/style.css` to use inherited typography and full-width button layout.
- ✅ Removed browser-default button appearance to ensure consistent rendering.
- ✅ Updated hover, active, and focus states to match the existing accent treatment and panel chrome.
- ✅ Re-ran the test suite after the UI-only change.

## Acceptance Criteria
- [x] Generate button remains an explicit POST submit control.
- [x] Button visually matches the rest of the website theme.
- [x] Existing functionality remains intact.
- [x] Tests pass with `CGO_ENABLED=0 go test ./...`.
