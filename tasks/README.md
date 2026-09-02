# Tasks - Modular TDD Implementation Guide

## Overview

32 tasks that build and present ASCII-Art-Web using **modular architecture** and
**strict TDD** for product behavior, followed by validated documentation and
delivery work.

## Critical Rules

### 1. Sequential Execution

Tasks MUST be completed in order (01 → 31). Each builds on previous tasks.

### 2. TDD Cycle (Mandatory)

Every task follows Red-Green-Refactor:

- **RED:** Write tests, verify they FAIL
- **GREEN:** Write minimal code to pass
- **REFACTOR:** Clean up, verify still passes

### 3. Modular Structure

Code organized in packages:

- `internal/` - Business logic (testable functions)
- `main.go` - Entry point (thin wrapper)
- Each internal function has unit tests

### 4. No Code in Tasks

Tasks contain:

- Requirements (what to achieve)
- Instructions (how to achieve it)
- Validation (how to verify)
- NO actual code (prevents copy-paste)

## Task Execution Order

| Task | Module                   | Golden Tests                     | Time  |
| ---- | ------------------------ | -------------------------------- | ----- |
| 01   | parse_input              | GT-001 to GT-005                 | 20min |
| 02   | parse_input (update)     | GT-006 to GT-010                 | 15min |
| 03   | load_banner              | Infrastructure                   | 20min |
| 04   | render_art               | GT-014 to GT-017                 | 30min |
| 05   | render_art (verify)      | GT-018 to GT-022                 | 15min |
| 06   | render_art (fix)         | GT-011 to GT-013, GT-023         | 20min |
| 07   | render_art (refine)      | GT-007 to GT-009, GT-024, GT-025 | 20min |
| 08   | Verification             | GT-026 to GT-029                 | 15min |
| 09   | Final Validation         | GT-030 + All                     | 30min |
| 10   | Project Structure        | Guidelines & Banners             | 15min |
| 11   | Banner Path Integration  | Multi-style support              | 20min |
| 12   | Style Selection          | Optional Arguments               | 30min |
| 13   | Audit & Cleanup          | Final Polish                     | 30min |
| 14   | Color Repo Setup         | Error Handling Standardization   | 15min |
| 15   | AGENTS.md Go Conventions | Update to follow Go idioms       | 10min |
| 16   | Color Feature            | Implement --color flag           | 60min |
| 17   | Docs Updates             | PRD + Architecture refresh       | 20min |
| 18   | Internal Refactor        | Move core to internal packages   | 30min |
| 19   | Web Server               | HTTP routing and status codes    | 45min |
| 20   | Web Templates            | HTML GUI integration             | 45min |
| 21   | Web UX & Validation      | SPA, UI Polish, ASCII Check      | 45min |
| 22   | Auto-Layout & Cleanup    | Responsive wrap, remove buttons  | 30min |
| 23   | JSON API Negotiation     | Structured `/ascii-art` responses | 60min |
| 24   | Frontend JSON + Docs     | JSON UI flow, tests, docs polish | 60min |
| 25   | Explicit Form Button     | Add required POST submit button   | 15min |
| 26   | Generate Button Theme    | Match submit button to site UI    | 15min |
| 27   | Internal Render Refactor | Decouple web from public package  | 15min |
| 28   | Docker Docs              | Audit workflow                   | 15min |
| 29   | Export Downloads         | Audit headers + multi-format export | 60min |
| 30   | Portfolio Polish & CI    | Documentation + automated validation | 45min |
| 31   | Responsive & A11y Polish | Mobile, keyboard + reduced motion     | 45min |
| 32   | CI Formatting Repair     | Canonical Go formatting               | 10min |

**Total Time:** ~15 hours 20 minutes
**Total Tests:** 30 golden tests + color, handler, container, and unit tests

## File Structure (Final)

```
ascii-art/
├── asciiart/
│   └── asciiart.go
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── internal/
│   ├── parser/
│   ├── banner/
│   ├── render/
│   └── color/
├── main.go
├── go.mod
├── docs/
│   ├── PRD.md
│   ├── DOCKER.md
│   ├── architecture.MD
│   └── golden_tests.md
├── tasks/
│   ├── README.md
│   ├── TASK-01.md
│   └── ...
└── README.md
```

## How to Use Tasks

### Step 1: Read Documentation

Before starting TASK-01:

1. `docs/PRD.md` - Requirements
2. `docs/architecture.MD` - Design
3. `docs/golden_tests.md` - Acceptance criteria
4. `TDD_HYBRID.md` - Modular approach

### Step 2: Execute Each Task

For each task (01-32):

**1. Open task file:**

```bash
cat tasks/TASK-01.md
```

**2. RED Phase:**

```bash
cd test
go test -v -run TestTask01  # See it fail
```

**3. GREEN Phase:**

- Create/update files in `internal/`
- Create unit tests
- Update `main.go` if needed
- Follow task instructions

**4. Verify:**

```bash
go test ./internal/... -v       # Unit tests
cd ../test && go test -v -run TestTask01  # Integration tests
```

**5. REFACTOR:**

- Clean up code
- Run `gofmt`
- Verify all tests still pass

**6. Regression:**

```bash
cd test && go test -v  # All previous tasks still pass
```

### Step 3: Final Validation

After TASK-09:

```bash
cd test && go test -v  # All 30 tests pass
go build              # Build succeeds
gofmt -d .            # No formatting issues
```

## Validation After Each Task

- [ ] Current task tests pass
- [ ] All previous tests still pass (regression)
- [ ] Code formatted with `gofmt`
- [ ] No compiler warnings
- [ ] Unit tests pass
- [ ] Integration tests pass

## Benefits of This Approach

✅ **Modular:** Each function in separate file  
✅ **Testable:** Unit tests for each component  
✅ **Maintainable:** Clear separation of concerns  
✅ **Professional:** Real-world Go structure  
✅ **TDD:** Tests drive implementation  
✅ **Compliant:** All 30 golden tests pass

## Quick Start

```bash
# Navigate to project
cd ascii-art

# Start with TASK-01
cat tasks/TASK-01.md

# Follow RED-GREEN-REFACTOR cycle
cd test && go test -v -run TestTask01  # RED
# ... implement code ...
cd test && go test -v -run TestTask01  # GREEN

# Move to TASK-02
# Repeat until TASK-09 complete
```

## Completion Criteria

Project is complete when:

- [ ] All 32 tasks marked DONE
- [ ] All 30 golden tests pass
- [ ] All unit tests pass
- [ ] `go build` succeeds
- [ ] No external dependencies
- [ ] Code formatted with `gofmt`
- [ ] Modular structure maintained

## Support

If stuck:

- Re-read task instructions
- Check `docs/golden_tests.md` for expected output
- Review `TDD_HYBRID.md` for architecture
- Ensure previous tasks pass (regression)
