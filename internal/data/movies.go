package data

import (
	"time"
)

// It’s crucial to point out here that all the fields in our Movie struct are exported
// (i.e. start with a capital letter), which is necessary for them to be visible to Go’s encoding/json package.
// Any fields which aren’t exported won’t be included when encoding a struct to JSON.
type Movie struct {
	ID        int64     `json:"id"`         // Unique integer ID for the movie
	CreatedAt time.Time `json:"created_at"` // Timestamp for when the movie is added to our database
	Title     string    `json:"title"`      // Movie title
	Year      int32     `json:"year"`       // Movie release year
	Runtime   int32     `json:"runtime"`    // Movie runtime (in minutes)
	Genres    []string  `json:"genres"`     // Slice of genres for the movie (romance, comedy, etc.)
	Version   int32     `json:"version"`    // The version number starts at 1 and will be incremented each
	// time the movie information is updated
}
