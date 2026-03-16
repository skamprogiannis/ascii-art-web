# TASK-19: Web Server Implementation

**Status:** COMPLETED
**Priority:** P0 (Critical)
**Dependencies:** TASK-01 to TASK-18
**Estimated Time:** 45 minutes

## Objective
Set up the HTTP server and define core API routes for the ascii-art application.

## Requirements
1. Server must listen on port `8080`.
2. Endpoint `GET /` to serve the main frontend.
3. Endpoint `POST /ascii-art` to process text to ASCII conversions.
4. Correct HTTP status codes (200, 400, 404, 500) must be returned based on the scenario.

## TDD Workflow

### RED Phase
- Test `http.Get("http://localhost:8080")` and verify connection refused.

### GREEN Phase
- Use `net/http` to start the server.
- Implement basic route handlers.

### REFACTOR Phase
- Extract handler logic into dedicated internal functions.

## Acceptance Criteria
- [x] Server runs securely without panicking.
- [x] `GET /` returns 200 OK.
- [x] Invalid routes return 404 Not Found.
- [x] Malformed requests return 400 Bad Request.