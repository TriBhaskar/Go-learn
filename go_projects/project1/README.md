# Project 1

## Description

This is a learning project implementing a RESTful Blog Post API using Go with an in-memory database.

## Features

- RESTful API endpoints for blog post management
- In-memory database using hashicorp/go-memdb
- Graceful server shutdown
- Structured project layout
- JSON response handling

## Prerequisites

- Go 1.20 or higher installed on your system
- Basic understanding of Go programming
- Your preferred code editor (VS Code recommended)

## Project Structure

```
project1/
├── cmd/
│   └── api/
│       └── main.go         # Application entry point
├── internal/
│   ├── database/          # Database operations
│   ├── handlers/          # HTTP request handlers
│   ├── models/            # Data structures
│   └── routes/            # Route definitions
└── README.md
```

## Installation

1. Clone the repository:
   ```bash
   git clone [repository-url]
   ```
2. Navigate to the project directory:
   ```bash
   cd project1
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

Run the project:

```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:8090`

### API Endpoints

| Method | Endpoint          | Description                  |
| ------ | ----------------- | ---------------------------- |
| GET    | `/api/posts`      | Get all blog posts           |
| GET    | `/api/posts/{id}` | Get a specific blog post     |
| POST   | `/api/posts`      | Create a new blog post       |
| PUT    | `/api/posts/{id}` | Update an existing blog post |
| DELETE | `/api/posts/{id}` | Delete a blog post           |

### Blog Post Structure

```json
{
  "id": "string",
  "title": "string",
  "content": "string",
  "author": "string",
  "tags": ["string"],
  "published": boolean,
  "created_at": "string",
  "updated_at": "string"
}
```

## Database

The application uses an in-memory database (`hashicorp/go-memdb`) which means:

- Data persists only during runtime
- Perfect for development and testing
- Data is lost when the server stops

## Contributing

This is a personal learning project, but suggestions and feedback are welcome through issues and discussions.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
