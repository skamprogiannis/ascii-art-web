# ASCII-Art-Web

## Description

ASCII-Art-Web is a web application that converts text into ASCII art using different banner styles. Built with Go and standard library only, it provides a simple web interface to generate artistic text representations in three styles: standard, shadow, and thinkertoy.

## Authors

- [gelafros](https://platform.zone01.gr/git/gelafros/)
- [skamprog](https://platform.zone01.gr/git/skamprog/)
- [emanola](https://platform.zone01.gr/git/emanola/)

## Usage

### How to Run

1. Ensure Go 1.21 or higher is installed
2. Navigate to the project directory
3. Run the server:

```bash
go run .
```

If your environment is missing a C compiler (`gcc`) and `go run .` fails with a cgo error, run:

```bash
CGO_ENABLED=0 go run .
```

4. Open your browser and visit: `http://localhost:8080`
5. Enter text, select a banner style, and click "Generate"

### Docker

You can build and run the project with plain Docker commands from the project root:

```bash
docker image build -t ascii-art-web-docker .
docker container run -d -p 8080:8080 --name dockerize ascii-art-web-docker
```

The final `.` passes the current directory as the Docker build context.

Useful follow-up commands:

```bash
docker images
docker ps -a
docker exec -it dockerize /bin/bash
docker image inspect ascii-art-web-docker --format '{{json .Config.Labels}}'
docker stop dockerize
docker rm dockerize
```

For the full manual Docker workflow, see [`docs/DOCKER.md`](docs/DOCKER.md).

## Testing

### Unit Tests

Run the unit tests for the web handlers:

```bash
go test .
```

### Docker Integration Tests

Comprehensive integration tests for the Docker containerization:

```bash
# Run all Docker integration tests (handles permissions automatically)
./run_docker_tests.sh

# Or run directly (requires Docker permissions)
./docker_integration_test.sh
```

These tests verify:
- ✅ Docker build success
- ✅ Best practices (linting, security)
- ✅ Runtime health (container stays alive)
- ✅ Port accessibility (web server responds)

See [docs/DOCKER_INTEGRATION_TESTS.md](docs/DOCKER_INTEGRATION_TESTS.md) for detailed documentation.

### Features

- Web-based GUI for ASCII art generation
- Single Page Application (SPA) experience using JavaScript `fetch`
- JSON response mode on `POST /ascii-art` via `Accept: application/json`
- Live typing (debounced auto-generation)
- Three banner styles: standard, shadow, thinkertoy
- Real-time result display on the same page
- Color picker and substring coloring support
- Dark/Light mode canvas toggle
- Download generated art as `.txt` file
- Proper HTTP status code handling (200, 400, 404, 500)
- Strict ASCII range validation (32-126) for secure input

### Security Measures

- **DoS Protection**: Strict 1000-character payload limit on the backend to prevent CPU and memory exhaustion.
- **Slowloris Mitigation**: Custom `http.Server` configured with strict 10-second read/write timeouts.
- **Path Traversal Prevention**: Hardcoded whitelisting for banner file selection (blocks `../../` injections).
- **XSS Prevention**: Safe HTML escaping using Go's `html/template` package.

## Implementation Details

### Algorithm

1. **HTTP Server Setup**: Server listens on port 8080 with two endpoints
    - `GET /`: Serves the main HTML page with form
    - `POST /ascii-art`: Processes form data and returns HTML or JSON (content negotiation)

2. **Request Processing**:
    - Parse form data (text input and banner selection)
    - Validate banner type (standard, shadow, thinkertoy)
    - Validate input length and printable ASCII characters
    - Return 400 Bad Request for invalid input

3. **ASCII Art Generation**:
    - Load selected banner file from `banners/` directory
   - Each banner contains 8 lines per character (ASCII 32-126)
   - Calculate character position: `(ASCII_code - 32) * 9 + 1 + rowNumber`
   - Render each line by concatenating character art horizontally
   - Handle newlines (`\n`) by processing text in segments

4. **Response Handling**:
    - `Accept: application/json` → JSON payload with `ascii_art`, `ascii_art_html`, and `error`
    - Other `Accept` values → HTML template response for browser rendering
    - 200 OK: Successful generation
    - 400 Bad Request: Invalid input or banner
    - 404 Not Found: Invalid route or missing template/banner
    - 500 Internal Server Error: Rendering or internal server errors

### JSON Usage Example

```bash
curl -X POST "http://localhost:8080/ascii-art" \
  -H "Accept: application/json" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  --data-urlencode "text=Hello" \
  --data-urlencode "banner=standard" \
  --data-urlencode "color=#8b5cf6" \
  --data-urlencode "substr=He"
```

Example JSON fields:
- `input_text`, `banner`, `color`, `substr`
- `ascii_art` (plain text)
- `ascii_art_html` (HTML with optional span coloring)
- `error` (empty on success)

### Project Structure

```
ascii-art-web/
├── templates/           # HTML templates
│   └── index.html
├── static/              # CSS and frontend JS assets
│   ├── style.css
│   ├── layout.css
│   └── app.js
├── banners/            # Banner files
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── asciiart/           # Public library
│   └── asciiart.go
├── internal/           # Internal packages
│   ├── banner/
│   ├── parser/
│   └── render/
├── main.go             # Server entry point
└── go.mod              # Go module
```

### Technical Stack

- Language: Go (standard library only)
- Server: net/http
- Templating: html/template
- Port: 8080
