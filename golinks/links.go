package golinks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	linksPath = "links"
)

// LinksService handles communication to Golinks Golinks API.
//
// API doc: https://docs.golinks.io/#19f87188-f2fb-4b75-bf69-83f0c0ca5029
type LinksService service

// Link object represents Golinks Golink.
//go:generate gomodifytags -file $GOFILE -struct Link -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct Link -add-tags json -w -transform camelcase
type Link struct {
	GID          int           `json:"gid"`
	CID          int           `json:"cid"`
	User         *User         `json:"user"`
	URL          string        `json:"url"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Tags         []string      `json:"tags"`
	Unlisted     *int          `json:"unlisted"`
	VariableLink *int          `json:"variable_link`
	Pinned       *int          `json:"pinned"`
	RedirectHits *RedirectHits `json:"redirect_hits"`
}

//go:generate gomodifytags -file $GOFILE -struct User -clear-tags -w
//go:generate gomodifytags --file $GOFILE --struct User -add-tags json -w -transform camelcase
type User struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	UserImageURL string `json:"user_image_url"`
}

//go:generate gomodifytags -file $GOFILE -struct v -clear-tags -w
//go:generate gomodifytags --file $GOFILE --structRedirectHits v -add-tags json -w -transform camelcase
type RedirectHits struct {
	Daily     *int `json:"daily"`
	Weekly    *int `json:"weekly"`
	Monthly   *int `json:"monthly"`
	AllTime   *int `json:"alltime"`
	CreatedAt *int `json:"created_at"`
	UpdatedAt *int `json:"updated_at"`
}

// Retrieve gets golinks by Golinks ID(GID).
//
// API doc: https://docs.golinks.io/#bc35cbc8-2424-438a-985f-c37cb69add7b
func (s *LinksService) Retrieve(ctx context.Context, id int) (*Link, error) {
	req, err := s.client.NewGetRequest(fmt.Sprintf("%s/%d", linksPath, id))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	if http.StatusBadRequest <= resp.StatusCode && resp.StatusCode <= http.StatusInsufficientStorage {
		return nil, fmt.Errorf("status code not expected, got:%d", resp.StatusCode)
	}

	t := &Link{}
	if err := json.NewDecoder(resp.Body).Decode(t); err != nil {
		return nil, err
	}

	return t, nil
}

// Create creates golinks by Golinks ID.
//
// API doc: https://docs.golinks.io/#bc35cbc8-2424-438a-985f-c37cb69add7b
func (s *LinksService) Create(ctx context.Context, link *Link) (*Link, error) {
	req, err := s.client.NewPostRequest(linksPath, link)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusAccepted {
		return nil, fmt.Errorf("status code not expected, got:%d want:201", resp.StatusCode)
	}

	t := &Link{}
	if err := json.NewDecoder(resp.Body).Decode(t); err != nil {
		return nil, err
	}

	return t, nil
}
