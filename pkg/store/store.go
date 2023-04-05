package store

import (
	"VulTracks/pkg/globals"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
)

// StoreType is the session store, it contains the redis and sqlite3 storage apis
type StoreType struct {
	Storage  *sqlite3.Storage
	Sessions *session.Store
}

// NewStore creates a new session store based on the storage type
func NewStore() *StoreType {
	store := &StoreType{}
	storage := sqlite3.New(sqlite3.Config{
		Database: globals.DatabaseLocation,
		Table:    "sessions",
	})
	store.Storage = storage
	store.Sessions = session.New(session.Config{
		Storage: storage,
	})
	return store
}

// Store is the session store, it needs to be set before use
var Store *StoreType
