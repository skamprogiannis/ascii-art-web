# TASK-28: Docker Documentation & Manual Commands

**Status:** COMPLETED
**Priority:** P3 (Low)
**Dependencies:** TASK-27
**Estimated Time:** 15 minutes

## Objective
Document the manual Docker workflow in the README and a dedicated Docker guide so auditors can build, run, inspect, and clean up the container without relying on `docker.sh`.

## Requirements
1. Keep the README Docker section concise and audit-friendly.
2. Add a dedicated Docker guide under `docs/`.
3. Include exact commands for build, run, inspect, exec, stop, and cleanup.
4. Explain the build-context requirement for `docker image build`.
5. Keep the instructions aligned with the current multi-stage Docker setup.

## Implementation Steps
- ✅ Updated `README.md` with a short Docker quick-start and follow-up commands.
- ✅ Added `docs/DOCKER.md` with the manual Docker workflow and audit-oriented checks.
- ✅ Documented the role of the trailing `.` as the build context.
- ✅ Updated project tracking files for the new documentation task.

## Acceptance Criteria
- [x] `README.md` includes concise Docker usage.
- [x] `docs/DOCKER.md` exists and lists raw Docker commands.
- [x] The guide explains build, run, inspect, exec, stop, and cleanup.
- [x] The instructions do not require `docker.sh`.
