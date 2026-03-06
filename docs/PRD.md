# Product Requirements Document (PRD)

## ASCII-Art Generator v1.3

### 1. Overview

**Product Name:** ASCII-Art Generator  
**Version:** 1.3  
**Date:** 2026  
**Owner:** Zone01 Athens

### 2. Problem Statement

Users need a command-line tool to convert text strings into large ASCII art representations for terminal displays, documentation, and creative text formatting.

### 3. Goals & Objectives

- Provide instant text-to-ASCII-art conversion
- Support multiple banner styles (standard, shadow, thinkertoy)
- Maintain zero external dependencies
- Ensure deterministic, reproducible output
- Provide a library API for programmatic use

### 4. Target Users

- Developers creating CLI tools
- System administrators customizing terminal outputs
- Students learning Go programming
- Anyone needing ASCII art for text-based interfaces

### 5. Functional Requirements

#### FR-1: Input Handling

- **FR-1.1:** Accept command-line arguments with optional `--color=<color>` flag
- **FR-1.2:** Support literal `\n` for newlines within the string
- **FR-1.3:** Handle empty strings (no output)
- **FR-1.4:** Reject invalid argument formats with usage message `"Usage: go run . [COLOR] [SUBSTRING] [STRING] [BANNER]"`

#### FR-2: ASCII Art Generation

- **FR-2.1:** Convert each character to 8-line ASCII art representation
- **FR-2.2:** Support ASCII characters 32-126 (printable characters)
- **FR-2.3:** Concatenate characters horizontally per line
- **FR-2.4:** Print exactly 8 lines per text segment

#### FR-3: Newline Processing

- **FR-3.1:** Split input by newline characters
- **FR-3.2:** Print blank line for each newline separator
- **FR-3.3:** Handle multiple consecutive newlines correctly

#### FR-4: Banner File Support

- **FR-4.1:** Load banner file at startup (default: standard.txt)
- **FR-4.2:** Support optional banner styles: standard, shadow, thinkertoy
- **FR-4.3:** Exit with error if banner file missing

#### FR-5: Color Output

- **FR-5.1:** Support `--color=<color>` flag to color output
- **FR-5.2:** If substring is provided, color only matching substrings
- **FR-5.3:** If no substring is provided, color the entire string
- **FR-5.4:** Accept color names and color codes (hex, rgb, hsl)
- **FR-5.5:** Invalid flag formats return usage message

### 6. Non-Functional Requirements

#### NFR-1: Performance

- Load banner file once (no repeated I/O)
- Process and print line-by-line (no full buffering)
- Handle strings up to 1000 characters efficiently

#### NFR-2: Reliability

- Deterministic output (same input = same output)
- Graceful error handling for missing files
- No crashes on invalid input

#### NFR-3: Maintainability

- Modular implementation with internal packages and public library package
- Standard library only (no external dependencies)
- Code compatible with `gofmt`

#### NFR-4: Portability

- Cross-platform (Linux, macOS, Windows)
- Go 1.16+ compatibility
- No OS-specific dependencies

### 7. Constraints

- **Technical:** Standard library only, no external packages
- **Architectural:** Modular design with internal packages for business logic
- **Operational:** Banner files must be in `banners/` directory

### 8. Success Metrics

- All 30 golden tests pass (GT-001 through GT-030) for standard banner
- Multi-banner selection works as expected
- Build succeeds with `go build`
- Zero external dependencies in `go.mod`
- Code passes `gofmt` validation

### 9. Out of Scope (v1.3)

- Custom banner file creation
- File input/output
- Configuration files
- Unicode support beyond ASCII 32-126

### 10. Acceptance Criteria

1. ✓ Valid argument formats accepted (with or without color flag)
2. ✓ Literal `\n` converted to newlines
3. ✓ 8 lines printed per text segment
4. ✓ Empty string produces no output
5. ✓ Newlines create blank line separators
6. ✓ Invalid arguments show usage message (`Usage: go run . [OPTION] [STRING] [BANNER]`)
7. ✓ Banner file errors handled gracefully
8. ✓ All 30 golden tests pass (GT-001 through GT-030)
9. ✓ All ASCII 32-126 characters supported
10. ✓ All styles (standard, shadow, thinkertoy) render correctly
11. ✓ Color flag renders colored output
12. ✓ Substring coloring works as expected

### 11. Dependencies

- Go 1.16+ runtime
- Banner files: standard.txt, shadow.txt, thinkertoy.txt

### 12. Risks & Mitigations

| Risk                 | Impact | Mitigation                        |
| -------------------- | ------ | --------------------------------- |
| Missing banner files | High   | Error handling with clear message |
| Invalid characters   | Medium | Skip or handle gracefully         |
| Large input strings  | Low    | Line-by-line processing           |

### 13. Timeline

- **Phase 1:** Core implementation (`main.go`) - 1 day
- **Phase 2:** Testing against golden tests - 1 day
- **Phase 3:** Multi-banner and cleanup - 1 day

### 14. Approval

This PRD serves as the source of truth for ASCII-Art v1.3 implementation.
