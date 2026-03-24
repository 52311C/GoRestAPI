# Building a Go REST API with Docker and PostgreSQL - Beginner Guide

## 1. Title and Objective

### Title
Building a Go REST API with Docker and PostgreSQL - A Beginner's Guide

### Technology Chosen
The project uses these core technologies:

- Go for the backend API
- PostgreSQL for persistent storage
- Docker Compose for running the API and database together
- A simple HTML, CSS, and JavaScript page for browser-based interaction

### Why These Technologies Were Chosen

- Go is a good fit for small REST APIs because the standard library already includes `net/http`, which keeps the code simple.
- PostgreSQL is a reliable relational database and is commonly used in production systems.
- Docker Compose makes local setup repeatable because the API and database can be started with one command.
- A plain browser UI avoids adding a full frontend framework for a small CRUD app.

### End Goal
The end goal is to run a working CRUD application where a user can:

- Create items
- Read all items
- Update an existing item
- Delete an item
- Use either the browser UI or direct HTTP requests

## 2. Quick Summary of the Technology

Go is a compiled programming language designed for simplicity, readability, and strong support for network services. In this project it is used to build a REST API that listens for HTTP requests and returns JSON responses.

### Where It Is Used

- Backend services
- REST APIs
- Cloud services
- CLI tools
- Microservices

### Real-World Example
Companies such as Google, Docker, and Kubernetes-related projects use Go heavily for backend and infrastructure tools. Docker itself is written primarily in Go.

## 3. System Requirements

### Supported Operating Systems

- Linux
- macOS
- Windows

### Required Tools

- VS Code or another text editor
- Docker Desktop or Docker Engine with Docker Compose
- Git
- A web browser such as Chrome or Firefox

### Optional Local Tools

- Go 1.22 or later if you want to run the app outside Docker
- curl for direct API testing
- Postman or Insomnia for API exploration

### Packages and Dependencies

This project uses these Go packages:

- `github.com/lib/pq`
- `github.com/joho/godotenv`

The project dependencies are installed automatically during Docker image build with:

```bash
go mod download
```

## 4. Installation and Setup Instructions

### Step 1: Clone the repository

```bash
git clone <your-repo-url>
cd go-todo-api
```

### Step 2: Create the environment file

```bash
cp .env.example .env
```

Example `.env` values:

```env
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=todo_db
```

### Step 3: Start the application stack

```bash
docker compose up --build
```

### Expected terminal output

A successful startup should look similar to this:

```text
[+] Running 2/2
 ✔ Container postgres_db Running
 ✔ Container go_api      Started
```

The API should now be available at:

- `http://localhost:8090/`
- `http://localhost:8090/health`
- `http://localhost:8090/items`

### Step 4: Open the browser UI

Open this URL in your browser:

`http://localhost:8090/`

From the UI you can create, update, delete, and refresh items without using terminal commands.

## 5. Minimal Working Example

### What the Example Does

This example shows a minimal Go HTTP endpoint that returns a health status JSON response. It demonstrates the same style used in the project for lightweight API routes.

### Example Code

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func main() {
    // Register a GET endpoint at /health.
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        // Reject unsupported HTTP methods.
        if r.Method != http.MethodGet {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        // Tell the client the response is JSON.
        w.Header().Set("Content-Type", "application/json")

        // Send a small JSON payload that confirms the service is alive.
        _ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
    })

    // Start the HTTP server on port 8090.
    log.Fatal(http.ListenAndServe(":8090", nil))
}
```

### Expected Output

Request:

```bash
curl http://localhost:8090/health
```

Response:

```json
{"status":"ok"}
```

## 6. AI Prompt Journal

This section records example prompts used while building and improving the project.

### Prompt 1

- Prompt used: "Help me build a simple Go REST API with PostgreSQL using Docker Compose."
- Curriculum link: Not provided in this repository. Add your course or assignment link here if required.
- AI response summary: Outlined a basic project structure with Go files, a Dockerfile, and a Compose file.
- Helpful response excerpt: "Use Docker Compose to run both the API container and the PostgreSQL container together so local setup stays consistent."
- Evaluation: Helpful for initial scaffolding because it reduced setup time and suggested a reasonable folder structure.

### Prompt 2

- Prompt used: "Why does curl to localhost:8090 fail with connection reset by peer?"
- Curriculum link: Not provided in this repository. Add your course or assignment link here if required.
- AI response summary: Helped identify that the app and container port configuration had drifted and that the running image needed to be rebuilt.
- Helpful response excerpt: "Check whether the app is actually listening on the same port that Docker publishes, and rebuild the container if the image is stale."
- Evaluation: Very helpful because it pointed directly to a real runtime issue instead of only reviewing source code.

### Prompt 3

- Prompt used: "Add a PUT endpoint so an item can be updated by ID."
- Curriculum link: Not provided in this repository. Add your course or assignment link here if required.
- AI response summary: Added the update handler, registered the route, and validated it with a live POST -> PUT -> GET flow.
- Helpful response excerpt: "Return 404 when the update affects zero rows so the client can distinguish missing records from successful updates."
- Evaluation: Helpful because it improved both functionality and API behavior.

### Prompt 4

- Prompt used: "Add a simple UI so users can interact with the API without curl."
- Curriculum link: Not provided in this repository. Add your course or assignment link here if required.
- AI response summary: Added a static HTML page, served it from the Go server, and connected it to the API using `fetch`.
- Helpful response excerpt: "Serve a small static page at the root path and call the existing `/items` endpoints from browser JavaScript."
- Evaluation: Helpful because it delivered a usable interface without introducing unnecessary frontend complexity.

## 7. Common Issues and Fixes

### Issue 1: Port mismatch between app and Docker

Problem:

- The app and container were not aligned on the same port during testing, which caused failed connections.

Symptom:

```text
curl: (56) Recv failure: Connection reset by peer
```

Fix:

- Ensure the Go server listens on `:8090`.
- Ensure Docker Compose maps `8090:8090`.
- Rebuild the app container:

```bash
docker compose up -d --build app
```

Helpful references:

- https://docs.docker.com/compose/
- https://go.dev/doc/

### Issue 2: Stale container image after source changes

Problem:

- The code in the container did not match the latest local source.

Fix:

```bash
docker compose up -d --build app
```

Why it worked:

- This forces Docker to rebuild the app image with the current Go files.

### Issue 3: Root path did not provide a usable page

Problem:

- Visiting `/` returned `404 page not found` before a UI route was added.

Fix:

- Add a static UI page under `static/index.html`.
- Serve static files from the root path in the router.

### Issue 4: Missing update support

Problem:

- The API originally supported GET, POST, and DELETE, but not PUT.

Fix:

- Add a `PUT /items?id=1` route.
- Validate the `id`.
- Return `404` when no rows are updated.

## 8. References

### Official Documentation

- Go documentation: https://go.dev/doc/
- Go `net/http` package: https://pkg.go.dev/net/http
- Go `database/sql` package: https://pkg.go.dev/database/sql
- Docker Compose documentation: https://docs.docker.com/compose/
- PostgreSQL documentation: https://www.postgresql.org/docs/
- `lib/pq` package: https://pkg.go.dev/github.com/lib/pq
- `godotenv` package: https://pkg.go.dev/github.com/joho/godotenv

### Helpful Learning Resources

- Go by Example: https://gobyexample.com/
- MDN Fetch API: https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API
- Docker Get Started: https://docs.docker.com/get-started/

### Optional Video Search Topics

If you need videos for a report or presentation, search these topics on YouTube:

- "Go REST API tutorial"
- "Docker Compose beginner tutorial"
- "PostgreSQL Docker tutorial"
