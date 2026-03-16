# TASK-20: HTML Templates & GUI Integration

**Status:** COMPLETED
**Priority:** P1 (High)
**Dependencies:** TASK-19
**Estimated Time:** 45 minutes

## Objective
Create a user-friendly web interface for generating ASCII art.

## Requirements
1. Create `templates/index.html` with a form containing a text input (textarea) and banner selection (radio or select).
2. Render the template for `GET /` requests.
3. Update `POST /ascii-art` to process the form data, generate the art, and return the populated template.
4. Ensure that the generated ASCII art is displayed in a `<pre>` tag to preserve formatting.

## TDD Workflow

### RED Phase
- Attempt to parse missing template files.

### GREEN Phase
- Add `index.html` in `templates/`.
- Use `html/template` to parse the files on application startup.

### REFACTOR Phase
- Centralize template parsing to avoid parsing per request.

## Acceptance Criteria
- [x] The web UI loads at `http://localhost:8080`.
- [x] Form submission correctly sends text and banner inputs.
- [x] Resulting ASCII art correctly maintains spacing via `<pre>` blocks.