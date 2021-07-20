package domain

import "time"

type URLShortenEntry struct {
	// ID is an identifier that references this entry
	ID string
	// ActualURL holds the actual URL to be returned when the short one is queried
	ActualURL string
	// ShortName is the name that the user will be referencing when trying to get the actual URL
	ShortName string
	// CreatedTime is the time the entry was created
	CreatedTime time.Time
	// UpdateTime is the last time the user updated the entry
	UpdateTime time.Time
	// LastAccessTime is the time this entry was last accessed by a user
	LastAccessTime time.Time
}

func NewURLShortenEntry(id, actualURL, shortName string) *URLShortenEntry {
	timeNow := time.Now()
	return &URLShortenEntry{
		ID:             id,
		ActualURL:      actualURL,
		ShortName:      shortName,
		CreatedTime:    timeNow,
		UpdateTime:     timeNow,
		LastAccessTime: timeNow,
	}
}

func (u *URLShortenEntry) ModifyURLShortenEntry(actualURL, shortName string) {
	var updated bool = false

	if len(actualURL) > 0 {
		u.ActualURL = actualURL
		updated = true
	}

	if len(shortName) > 0 {
		u.ShortName = shortName
		updated = true
	}

	if updated {
		u.UpdateTime = time.Now()
	}
}
