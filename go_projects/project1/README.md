# Project 1

## Description

This is a learning project to explore Go programming concepts and best practices.

## Getting Started

### Prerequisites

- Go 1.20 or higher installed on your system
- Basic understanding of Go programming
- MySQL or PostgreSQL database installed
- Your preferred code editor (VS Code recommended)

### Installation

1. Clone the repository:
   ```bash
   git clone [repository-url]
   ```
2. Navigate to the project directory:
   ```bash
   cd project1
   ```
3. Initialize the Go module:
   ```bash
   go mod init project1
   ```
4. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

This CRUD application allows you to:

- Create new blog posts
- Read existing posts
- Update post content
- Delete posts

Run the project:

```bash
go run main.go
```

Access the API endpoints at `http://localhost:8080`:

- `POST /posts` - Create a new post
- `GET /posts` - List all posts
- `GET /posts/{id}` - Get a specific post
- `PUT /posts/{id}` - Update a post
- `DELETE /posts/{id}` - Delete a post

## Contributing

This is a personal learning project, but suggestions and feedback are welcome through issues and discussions.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
