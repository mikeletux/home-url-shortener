package ports

import "github.com/mikeletux/home-url-shortener/internal/core/domain"

type URLShortenerService interface {
	Get(shortName string) (*domain.URLShortenEntry, error)
	Create(url, shortName string) (*domain.URLShortenEntry, error)
	Modify(id, url, shortName string) (*domain.URLShortenEntry, error)
}
