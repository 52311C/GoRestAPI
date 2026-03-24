
# Go To-Do API (Docker + PostgreSQL)

A simple REST API built with Go and PostgreSQL, containerized with Docker Compose.

## Overview

This project supports full CRUD for items:

- Create items
- Read items
- Update items
- Delete items

The app includes:

- JSON REST API endpoints
- PostgreSQL persistence
- Dockerized local development
- A simple browser UI for interacting with the API

## Tech Stack

- Go
- PostgreSQL
- Docker / Docker Compose
- Go modules: `net/http`, `github.com/lib/pq`, `github.com/joho/godotenv`

## Project Structure

```text
go-todo-api/
├── main.go
├── go.mod
├── Dockerfile
├── docker-compose.yml
├── handlers/
│   └── itemHandler.go
├── models/
│   └── item.go
├── database/
│   └── db.go
├── routes/
│   └── routes.go
└── static/
    └── index.html
```

## Setup

1. Clone the repo and enter it:

```bash
git clone <your-repo-url>
cd go-todo-api
```

2. Create `.env` from your template (or create it manually):

```bash
cp .env.example .env
```

3. Start services:

```bash
docker compose up --build
```

## URLs

- Web UI: `http://localhost:8090/`
- Health: `http://localhost:8090/health`
- Items API: `http://localhost:8090/items`

## Web UI Usage

Open `http://localhost:8090/` in your browser.

From the UI you can:

- Create an item
- Update an item by ID
- Delete an item by ID
- Refresh the list (GET `/items`)

## API Endpoints

| Method | Endpoint | Description | Body |
| --- | --- | --- | --- |
| GET | `/items` | Get all items | N/A |
| POST | `/items` | Create an item | `{ "name": "Task" }` |
| PUT | `/items?id=1` | Update an item by ID | `{ "name": "Updated task" }` |
| DELETE | `/items?id=1` | Delete an item by ID | N/A |

## API Examples

Create:

```bash
curl -X POST http://localhost:8090/items \
  -H "Content-Type: application/json" \
  -d '{"name":"Buy groceries"}'
```

Get all:

```bash
curl http://localhost:8090/items
```

Update:

```bash
curl -X PUT "http://localhost:8090/items?id=1" \
  -H "Content-Type: application/json" \
  -d '{"name":"Buy groceries and milk"}'
```

Delete:

```bash
curl -X DELETE "http://localhost:8090/items?id=1"
```

## Cleanup

Stop containers:

```bash
docker compose down
```

Stop and remove volumes (reset DB data):

```bash
docker compose down -v
```

## Future Improvements

- Add timestamps for created/updated tasks
- Add input validation
- Add pgAdmin container for DB inspection
