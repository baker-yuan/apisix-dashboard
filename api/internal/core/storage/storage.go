package storage

import "context"

type Interface interface {
	Get(ctx context.Context, key string) (string, error)
	List(ctx context.Context, key string) ([]Keypair, error)
	Create(ctx context.Context, key, val string) error
	Update(ctx context.Context, key, val string) error
	BatchDelete(ctx context.Context, keys []string) error
	Watch(ctx context.Context, key string) <-chan WatchResponse
}

type WatchResponse struct {
	Events   []Event
	Error    error
	Canceled bool
}

type Keypair struct {
	Key   string
	Value string
}

type Event struct {
	Keypair
	Type EventType
}

type EventType string

var (
	EventTypePut    EventType = "put"
	EventTypeDelete EventType = "delete"
)
