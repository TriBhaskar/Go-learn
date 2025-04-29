package database

import (
	"go_project1/internal/models"
	"log"

	"github.com/hashicorp/go-memdb"
)

var DB *memdb.MemDB

func InitDB() error {
	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"blogpost": {
				Name: "blogpost",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"author": {
						Name:    "author",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Author"},
					},
					"published": {
						Name:    "published",
						Unique:  false,
						Indexer: &memdb.BoolFieldIndex{Field: "Published"},
					},
				},
			},
		},
	}
		// Create a new database
		db, err := memdb.NewMemDB(schema)
		if err != nil {
			return err
		}
		
		DB = db
		log.Println("Successfully connected to in-memory database")
		return nil
}

// GetAllBlogPosts retrieves all blog posts from the database
func GetAllBlogPosts() ([]models.BlogPost, error) {
	// Get a read transaction
	txn := DB.Txn(false)
	defer txn.Abort()

	// Get all blog posts
	it, err := txn.Get("blogpost", "id")
	if err != nil {
		return nil, err
	}

	var posts []models.BlogPost
	for obj := it.Next(); obj != nil; obj = it.Next() {
		posts = append(posts, obj.(models.BlogPost))
	}

	return posts, nil
}

// GetBlogPostByID retrieves a specific blog post by ID
func GetBlogPostByID(id string) (models.BlogPost, bool, error) {
	// Get a read transaction
	txn := DB.Txn(false)
	defer txn.Abort()

	// Get blog post by ID
	raw, err := txn.First("blogpost", "id", id)
	if err != nil {
		return models.BlogPost{}, false, err
	}

	if raw == nil {
		return models.BlogPost{}, false, nil
	}

	return raw.(models.BlogPost), true, nil
}

// CreateBlogPost adds a new blog post to the database
func CreateBlogPost(post models.BlogPost) error {
	// Start a write transaction
	txn := DB.Txn(true)
	defer txn.Abort()

	// Insert the blog post
	if err := txn.Insert("blogpost", post); err != nil {
		return err
	}

	// Commit the transaction
	txn.Commit()
	return nil
}

// UpdateBlogPost updates an existing blog post
func UpdateBlogPost(post models.BlogPost) error {
	// Start a write transaction
	txn := DB.Txn(true)
	defer txn.Abort()

	// Update the post
	if err := txn.Insert("blogpost", post); err != nil {
		return err
	}

	// Commit the transaction
	txn.Commit()
	return nil
}

// DeleteBlogPost deletes a blog post by ID
func DeleteBlogPost(id string) (bool, error) {
	// Start a write transaction
	txn := DB.Txn(true)
	defer txn.Abort()

	// Check if post exists
	raw, err := txn.First("blogpost", "id", id)
	if err != nil {
		return false, err
	}

	if raw == nil {
		return false, nil
	}

	// Delete the post
	if err := txn.Delete("blogpost", raw); err != nil {
		return false, err
	}

	// Commit the transaction
	txn.Commit()
	return true, nil
}