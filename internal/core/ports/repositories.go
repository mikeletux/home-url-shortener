package ports

import "github.com/mikeletux/home-url-shortener/internal/core/domain"

type ShortenURLRepository interface {
	Get(shortName string) (*domain.URLShortenEntry, error)
	Save(*domain.URLShortenEntry) error
	Delete(shortName string) error
}
