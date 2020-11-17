package session

import "context"

// Store interface for session management
type Store interface {
	// Save session to store
	Save(context.Context, *Session) error
	// Fetch session from store
	Get(ctx context.Context, sessionID []byte) (*Session, error)
	// Delete session by id
	Delete(ctx context.Context, sessionID []byte) error
	// Cleanup old sessions
	RemoveExpired(context.Context) error
	// Fetch sessions with the following value set
	// Useful for looking up all sessions by IP address, User ID, etc...
	// Hard to implement with certain stores
	// FetchByValue(ctx context.Context, key string, value interface{}) ([]*Session, error)
}
