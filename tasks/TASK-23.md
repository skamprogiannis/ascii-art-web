# TASK-23: JSON API Content Negotiation

**Status:** COMPLETED
**Priority:** P1 (High)
**Dependencies:** TASK-19, TASK-20, TASK-21, TASK-22
**Estimated Time:** 60 minutes

## Objective
Add structured JSON responses to `POST /ascii-art` using HTTP content negotiation while preserving existing HTML behavior for browser requests.

## Requirements
1. Keep `POST /ascii-art` as the single processing endpoint.
2. Return JSON when `Accept: application/json` is present.
3. Keep template-based HTML response mode for standard form submissions.
4. Return proper HTTP status codes (200, 400, 404, 500).
5. Centralize request parsing/validation to avoid duplicated logic.

## Implementation Steps
- ✅ Added shared request parsing and validation helpers in `main.go` (`parseASCIIArtRequest`, `validateASCIIArtRequest`, `respondError`).
- ✅ Implemented content negotiation (`wantsJSONResponse`) and JSON serializer (`writeJSON`).
- ✅ Added `APIResponse` payload with `input_text`, `banner`, `color`, `substr`, `ascii_art`, `ascii_art_html`, and `error`.
- ✅ Refactored rendering flow to produce plain and HTML art together via `asciiart.RenderBundle`.
- ✅ Added safer template handling (`tmplErr`) to avoid startup panic and return proper HTTP errors.

## Acceptance Criteria
- [x] `POST /ascii-art` returns JSON when `Accept: application/json`.
- [x] `POST /ascii-art` still returns HTML when JSON is not requested.
- [x] Invalid banner/input cases return HTTP 400 with structured error.
- [x] Missing banner files map to HTTP 404.
- [x] Server behavior remains compliant with required status codes.
