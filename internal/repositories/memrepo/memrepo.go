package memrepo

import "github.com/mikeletux/home-url-shortener/internal/core/domain"

type memRepo struct {
	repository []domain.URLShortenEntry
}

func NewMemRepo() *memRepo {
	return &memRepo{
		repository: []domain.URLShortenEntry{},
	}
}

func (m *memRepo) Get(shortName string) (*domain.URLShortenEntry, error) {
	return nil, nil
}

func (m *memRepo) Save(*domain.URLShortenEntry) error {
	return nil
}

func (m *memRepo) Delete(shortName string) error {
	return nil
}
