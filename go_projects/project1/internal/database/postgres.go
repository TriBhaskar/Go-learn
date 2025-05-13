package database

import (
	"database/sql"
	"fmt"
	"go_project1/internal/models"

	"github.com/lib/pq"
)

type PostgresDB struct {
    db *sql.DB
}

func NewPostgresDB() *PostgresDB {
    return &PostgresDB{}
}

func (p *PostgresDB) Initialize() error {
    connStr := "postgres://bhaskar:bhaskar@localhost:5432/testgodb?sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return fmt.Errorf("error connecting to database: %v", err)
    }

    if err = db.Ping(); err != nil {
        return fmt.Errorf("error connecting to the database: %v", err)
    }else{
		fmt.Println("Successfully connected to PostgreSQL database")
	}

    p.db = db
    return p.createTables()
}

func (p *PostgresDB) createTables() error {
    query := `
        CREATE TABLE IF NOT EXISTS blog_posts (
            id VARCHAR(255) PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            content TEXT,
            author VARCHAR(255),
            tags TEXT[], -- PostgreSQL array type
            published BOOLEAN DEFAULT false,
            created_at TIMESTAMP NOT NULL,
            updated_at TIMESTAMP NOT NULL
        );`

    _, err := p.db.Exec(query)
    return err
}

func (p *PostgresDB) GetAllPosts() ([]models.BlogPost, error) {
    query := `SELECT id, title, content, author, tags, published, created_at, updated_at FROM blog_posts`
    rows, err := p.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []models.BlogPost
    for rows.Next() {
        var post models.BlogPost
        var tags pq.StringArray // Temporary variable for scanning tags

        err := rows.Scan(
            &post.ID,
            &post.Title,
            &post.Content,
            &post.Author,
            &tags, // Scan into pq.StringArray
            &post.Published,
            &post.CreatedAt,
            &post.UpdatedAt,
        )
        if err != nil {
            return nil, fmt.Errorf("error scanning row: %v", err)
        }

        post.Tags = []string(tags) // Convert pq.StringArray to []string
        posts = append(posts, post)
    }

    // Check for errors from iterating over rows
    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating over rows: %v", err)
    }

    return posts, nil
}

func (p *PostgresDB) CreatePost(post models.BlogPost) error {
    query := `
        INSERT INTO blog_posts (id, title, content, author, tags, published, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id`

    // Convert []string to pq.StringArray
    tags := pq.StringArray(post.Tags)

    _, err := p.db.Exec(query,
        post.ID,
        post.Title,
        post.Content,
        post.Author,
        tags, // Use the converted tags
        post.Published,
        post.CreatedAt,
        post.UpdatedAt,
    )
    
    if err != nil {
        return fmt.Errorf("error creating blog post: %v", err)
    }

    return nil
}

func (p *PostgresDB) UpdatePost(post models.BlogPost) error {
    query := `
        UPDATE blog_posts 
        SET title = $1, content = $2, author = $3, tags = $4, published = $5, updated_at = $6
        WHERE id = $7`

    result, err := p.db.Exec(query,
        post.Title,
        post.Content,
        post.Author,
        post.Tags,
        post.Published,
        post.UpdatedAt,
        post.ID,
    )

    if err != nil {
        return fmt.Errorf("error updating blog post: %v", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking update result: %v", err)
    }

    if rows == 0 {
        return fmt.Errorf("post not found")
    }

    return nil
}

func (p *PostgresDB) DeletePost(id string) (bool, error) {
    query := `DELETE FROM blog_posts WHERE id = $1`
    
    result, err := p.db.Exec(query, id)
    if err != nil {
        return false, fmt.Errorf("error deleting blog post: %v", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return false, fmt.Errorf("error checking delete result: %v", err)
    }

    return rows > 0, nil
}

func (p *PostgresDB) GetPostByID(id string) (models.BlogPost, bool, error) {
    query := `
        SELECT id, title, content, author, tags, published, created_at, updated_at 
        FROM blog_posts 
        WHERE id = $1`

    var post models.BlogPost
    var tags pq.StringArray
    
    err := p.db.QueryRow(query, id).Scan(
        &post.ID,
        &post.Title,
        &post.Content,
        &post.Author,
        &tags,
        &post.Published,
        &post.CreatedAt,
        &post.UpdatedAt,
    )

    post.Tags = []string(tags) // Convert back to []string

    if err == sql.ErrNoRows {
        return models.BlogPost{}, false, nil
    }

    if err != nil {
        return models.BlogPost{}, false, fmt.Errorf("error fetching blog post: %v", err)
    }

    return post, true, nil
}