package api

import (
	"context"
	"encoding/json"
)

// Bookmark is the API message for a bookmark.
type Bookmark struct {
	ID int `jsonapi:"primary,bookmark"`

	// Standard fields
	CreatorID int
	Creator   *Principal `jsonapi:"attr,creator"`
	CreatedTs int64      `jsonapi:"attr,createdTs"`
	UpdaterID int
	Updater   *Principal `jsonapi:"attr,updater"`
	UpdatedTs int64      `jsonapi:"attr,updatedTs"`

	// Domain specific fields
	Name string `jsonapi:"attr,name"`
	Link string `jsonapi:"attr,link"`
}

// BookmarkCreate is the API message for creating a bookmark.
type BookmarkCreate struct {
	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	CreatorID int

	// Domain specific fields
	Name string `jsonapi:"attr,name"`
	Link string `jsonapi:"attr,link"`
}

// BookmarkFind is the API message for finding bookmarks.
type BookmarkFind struct {
	ID *int

	// Standard fields
	CreatorID *int
}

func (find *BookmarkFind) String() string {
	str, err := json.Marshal(*find)
	if err != nil {
		return err.Error()
	}
	return string(str)
}

// BookmarkDelete is the API message for deleting a bookmark.
type BookmarkDelete struct {
	ID int

	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	DeleterID int
}

// BookmarkService is the service for bookmarks.
type BookmarkService interface {
	CreateBookmark(ctx context.Context, create *BookmarkCreate) (*Bookmark, error)
	FindBookmarkList(ctx context.Context, find *BookmarkFind) ([]*Bookmark, error)
	FindBookmark(ctx context.Context, find *BookmarkFind) (*Bookmark, error)
	DeleteBookmark(ctx context.Context, delete *BookmarkDelete) error
}
