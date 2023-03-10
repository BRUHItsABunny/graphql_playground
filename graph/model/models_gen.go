// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Movie struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	ReleaseDate time.Time `json:"releaseDate"`
}

type NewMovie struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}
