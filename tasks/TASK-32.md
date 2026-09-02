# TASK-32: Restore CI Quality Gate

**Status:** IN PROGRESS
**Priority:** P1 (High)
**Dependencies:** TASK-30
**Estimated Time:** 10 minutes

## Objective

Make the Docker integration test match canonical Go formatting so the GitHub
Actions quality job can reach every validation stage, then replace deprecated
action runtimes reported by the runner.

## Requirements

1. Apply `gofmt` without changing test behavior.
2. Preserve the existing CI quality gate.
3. Use maintained GitHub Action majors backed by the current runner runtime.
4. Run the same formatting check locally before pushing.

## Implementation Steps

- [x] Reproduced the GitHub Actions failure from the failed-run log.
- [x] Applied canonical formatting to `docker_integration_test.go`.
- [x] Confirmed `gofmt -d .` produces no output.
- [x] Updated the official checkout, Go, and Node setup actions to v7.
- [x] Updated task tracking and the AI usage log.

## Acceptance Criteria

- [x] The Docker integration test has canonical indentation and a final newline.
- [x] Local formatting, tests, vet, build, and JavaScript syntax checks pass.
- [x] A replacement GitHub Actions run completes successfully.
- [ ] The final workflow run has no deprecated action-runtime annotations.
