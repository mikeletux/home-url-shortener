package shortenservice

import "github.com/mikeletux/home-url-shortener/internal/core/domain"

type ShortenerService struct {
}

func NewShortenerService() *ShortenerService {
	return &ShortenerService{}
}

func (s *ShortenerService) Get(shortName string) (*domain.URLShortenEntry, error) {
	return nil, nil
}

func (s *ShortenerService) Create(url, shortName string) (*domain.URLShortenEntry, error) {
	return nil, nil
}

func (s *ShortenerService) Modify(id, url, shortName string) (*domain.URLShortenEntry, error) {
	return nil, nil
}
