package shortener

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrAlreadyExists = errors.New("there is already an entry for that redirect")
	ErrNotFound      = errors.New("no entry found for the alias")
)

type Store interface {
	Save(ctx context.Context, redirect Redirect) error
	Get(ctx context.Context, alias string) (Redirect, error)
	GetAll(ctx context.Context) ([]Redirect, error)
}

type InMemoryStore struct {
	redirects map[string]Redirect
	mu        sync.Mutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		redirects: make(map[string]Redirect),
	}
}

func (s *InMemoryStore) GetAll(ctx context.Context) ([]Redirect, error) {
	redirects := []Redirect{}
	for _, redirect := range s.redirects {
		redirects = append(redirects, redirect)
	}
	return redirects, nil
}

func (s *InMemoryStore) Save(ctx context.Context, redirect Redirect) error {
	if _, found := s.redirects[redirect.Alias]; found {
		return ErrAlreadyExists
	}

	s.mu.Lock()
	s.redirects[redirect.Alias] = redirect
	s.mu.Unlock()

	return nil
}

func (s *InMemoryStore) Get(ctx context.Context, alias string) (Redirect, error) {
	url, ok := s.redirects[alias]
	if !ok {
		return Redirect{}, ErrNotFound
	}

	return url, nil
}
