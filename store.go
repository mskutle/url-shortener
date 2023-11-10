package main

import (
	"errors"
	"sync"
)

var (
	ErrAlreadyExists = errors.New("there is already an entry for that url")
	ErrNotFound      = errors.New("no entry found for the alias")
)

type URLStore interface {
	Save(url URL) error
	Get(alias string) (URL, error)
	GetAll() ([]URL, error)
}

type InMemoryUrlStore struct {
	urls map[string]URL
	mu   sync.Mutex
}

func NewInMemoryUrlStore() *InMemoryUrlStore {
	return &InMemoryUrlStore{
		urls: make(map[string]URL),
	}
}

func (s *InMemoryUrlStore) GetAll() ([]URL, error) {
	urls := []URL{}
	for _, url := range s.urls {
		urls = append(urls, url)
	}
	return urls, nil
}

func (s *InMemoryUrlStore) Save(url URL) error {
	if _, found := s.urls[url.Alias]; found {
		return ErrAlreadyExists
	}

	s.mu.Lock()
	s.urls[url.Alias] = url
	s.mu.Unlock()

	return nil
}

func (s *InMemoryUrlStore) Get(alias string) (URL, error) {
	url, ok := s.urls[alias]
	if !ok {
		return URL{}, ErrNotFound
	}

	return url, nil
}
