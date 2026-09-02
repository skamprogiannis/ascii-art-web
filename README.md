# ASCII Art Web

ASCII Art Web is a standard-library Go application that turns text into
browser-rendered ASCII art. It combines the Zone01 stylize, Dockerize, and
export milestones in one final project: generate colored art interactively,
consume the same renderer as JSON, or download the result in three formats.

## Highlights

- Three banner styles: `standard`, `shadow`, and `thinkertoy`
- Whole-text or substring coloring with named, hex, RGB, and HSL colors
- Responsive browser interface with debounced previews and an explicit submit
  action
- HTML and JSON responses from the same `POST /ascii-art` endpoint
- Server-generated `.txt`, `.html`, and `.json` downloads from `POST /export`
- Explicit `Content-Type`, `Content-Length`, and `Content-Disposition` export
  headers
- Multi-stage Docker image running as an unprivileged user
- Standard-library backend with no third-party Go dependencies

## Architecture

```text
Browser or API client
        |
        v
Go net/http routes: /, /ascii-art, /export
        |
        v
Request validation and content negotiation
        |
        v
internal/generator
   |             |
banner loader    plain/HTML renderer + color parser
        |
        v
HTML page, JSON response, or downloadable file
```

The web server in `main.go` owns routing, validation, response headers, and
export serialization. `internal/generator` coordinates banner loading and the
plain/HTML renderers. The `asciiart` package exposes the same generator through
a small reusable Go API. Frontend behavior and styles live under `static/`,
while `templates/index.html` retains a normal HTML form fallback.

## Run Locally

Use the Go version declared in `go.mod`, then start the server from the project
root so it can find the templates, static assets, and banners:

```bash
go run .
```

If the local environment has no C toolchain, disable cgo:

```bash
CGO_ENABLED=0 go run .
```

Open <http://localhost:8080>.

### Docker

```bash
docker image build -t ascii-art-web .
docker container run --rm -p 8080:8080 --name ascii-art-web ascii-art-web
```

The container uses a multi-stage build and serves the application as a non-root
user. See [Docker Guide](docs/DOCKER.md) for inspection and cleanup commands.

## HTTP Interface

### Generate JSON

`POST /ascii-art` returns HTML by default. Request JSON through content
negotiation:

```bash
curl -X POST http://localhost:8080/ascii-art \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  --data-urlencode 'text=Hello' \
  --data-urlencode 'banner=standard' \
  --data-urlencode 'color=#8b5cf6' \
  --data-urlencode 'substr=He'
```

The JSON object contains the normalized request fields, plain `ascii_art`,
colored `ascii_art_html`, and an `error` field.

### Export a File

```bash
curl -X POST http://localhost:8080/export \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  --data-urlencode 'text=Hello' \
  --data-urlencode 'banner=shadow' \
  --data-urlencode 'format=html' \
  -OJ
```

Supported formats:

| Format | Media type | Contents |
| --- | --- | --- |
| `txt` | `text/plain` | Plain ASCII art |
| `html` | `text/html` | Standalone, styled HTML document |
| `json` | `application/json` | Request metadata plus plain and HTML art |

Omitting `format` defaults to `txt`.

## Test and Build

Run the fast suite without Docker integration:

```bash
go test -short ./...
go test -race -short ./...
go vet ./...
CGO_ENABLED=0 go build ./...
node --check static/app.js
```

When Docker is available, run the container build and HTTP smoke tests:

```bash
./docker_integration_test.sh
```

The integration script uses port `8081` and cleans up its test container and
image. Additional testing details are in
[Docker Integration Tests](docs/DOCKER_INTEGRATION_TESTS.md).

## Security and Operational Limits

- Banner names and export formats are selected from fixed allowlists.
- Input is limited to 1,000 bytes and printable ASCII plus newlines.
- User text is escaped for templates and HTML exports; colors must pass the
  dedicated color parser before becoming style values.
- The HTTP server sets read, write, and idle timeouts.
- The runtime container drops root privileges.

This remains an educational portfolio service, not an internet-facing security
boundary. It does not provide TLS, authentication, rate limiting, or a custom
whole-request body limit. A public deployment should add those controls at a
trusted reverse proxy and run with resource limits.

## Status

This repository is the canonical final version, reconstructed from the team’s
stylize, Dockerize, and export work while retaining original commit authorship.
The project is feature-complete for its Zone01 audit scope and maintained as a
portfolio demonstration rather than a general-purpose text rendering service.

Design and requirement history remains available in
[Architecture](docs/architecture.MD), [Product Requirements](docs/PRD.md), and
the append-only [AI Implementation Log](docs/AI_LOG.md).

## Team Contributions

- **gelafros** — built the original web experience, rendering/color features,
  request validation, and interactive frontend foundation.
- **Stefanos Kamprogiannis (`skamprog`)** — added JSON content negotiation,
  frontend API integration, explicit form submission, internal generator
  boundaries, Docker support, and multi-format exports.
- **emanola** — contributed container maintenance and the Docker integration
  test suite.

The Git history preserves the original author and committer metadata for each
stage of the project.
