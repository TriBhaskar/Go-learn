package database

import (
	"go_project1/internal/models"
	"log"

	"github.com/hashicorp/go-memdb"
)

type MemoryDB struct {
    db *memdb.MemDB
}

func NewMemoryDB() *MemoryDB {
    return &MemoryDB{}
}

func (m *MemoryDB) Initialize() error {
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

    db, err := memdb.NewMemDB(schema)
    if err != nil {
        return err
    }

    m.db = db
    log.Println("Successfully connected to in-memory database")
    return nil
}

func (m *MemoryDB) GetAllPosts() ([]models.BlogPost, error) {
    txn := m.db.Txn(false)
    defer txn.Abort()

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

func (m *MemoryDB) GetPostByID(id string) (models.BlogPost, bool, error) {
    txn := m.db.Txn(false)
    defer txn.Abort()

    raw, err := txn.First("blogpost", "id", id)
    if err != nil {
        return models.BlogPost{}, false, err
    }

    if raw == nil {
        return models.BlogPost{}, false, nil
    }

    return raw.(models.BlogPost), true, nil
}

func (m *MemoryDB) CreatePost(post models.BlogPost) error {
    txn := m.db.Txn(true)
    defer txn.Abort()

    if err := txn.Insert("blogpost", post); err != nil {
        return err
    }

    txn.Commit()
    return nil
}

func (m *MemoryDB) UpdatePost(post models.BlogPost) error {
    txn := m.db.Txn(true)
    defer txn.Abort()

    if err := txn.Insert("blogpost", post); err != nil {
        return err
    }

    txn.Commit()
    return nil
}

func (m *MemoryDB) DeletePost(id string) (bool, error) {
    txn := m.db.Txn(true)
    defer txn.Abort()

    raw, err := txn.First("blogpost", "id", id)
    if err != nil {
        return false, err
    }

    if raw == nil {
        return false, nil
    }

    if err := txn.Delete("blogpost", raw); err != nil {
        return false, err
    }

    txn.Commit()
    return true, nil
}