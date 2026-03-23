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

### Features

- Web-based GUI for ASCII art generation
- Single Page Application (SPA) experience using JavaScript `fetch`
- Live typing (debounced auto-generation)
- Three banner styles: standard, shadow, thinkertoy
- Real-time result display on the same page
- Dark/Light mode canvas toggle
- Download generated art as `.txt` file
- Proper HTTP status code handling (200, 400, 404, 500)
- Strict ASCII range validation (32-126) for secure input

## Implementation Details

### Algorithm

1. **HTTP Server Setup**: Server listens on port 8080 with two endpoints
   - `GET /`: Serves the main HTML page with form
   - `POST /ascii-art`: Processes form data and returns result

2. **Request Processing**:
   - Parse form data (text input and banner selection)
   - Validate banner type (standard, shadow, thinkertoy)
   - Return 400 Bad Request for invalid input

3. **ASCII Art Generation**:
   - Load selected banner file from `banners/` directory
   - Each banner contains 8 lines per character (ASCII 32-126)
   - Calculate character position: `(ASCII_code - 32) * 9 + 1 + rowNumber`
   - Render each line by concatenating character art horizontally
   - Handle newlines (`\n`) by processing text in segments

4. **Response Handling**:
   - 200 OK: Successful generation
   - 400 Bad Request: Invalid input or banner
   - 404 Not Found: Invalid route or missing template
   - 500 Internal Server Error: Template execution or rendering errors

### Project Structure

```
ascii-art-web/
├── templates/           # HTML templates
│   └── index.html
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
