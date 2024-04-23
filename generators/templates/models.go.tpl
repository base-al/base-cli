package {{.ModuleName}}

import "time"

// {{.EntityTitle}} represents a single CMS {{.EntityName}} with comprehensive details including metadata and status.
type {{.EntityTitle}} struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PublishedAt time.Time `json:"published_at"`
	Status      string    `json:"status"`    // e.g., "draft", "published", "archived"
	AuthorID    int       `json:"author_id"` // ID of the user who created or last updated the {{.EntityName}}
}

// Create{{.EntityTitle}}Request model for the create API call, now includes AuthorID.
type Create{{.EntityTitle}}Request struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int    `json:"author_id"`
}

// Create{{.EntityTitle}}Response model for the create response.
type Create{{.EntityTitle}}Response struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

// Read{{.EntityTitle}}Request model for fetching a {{.EntityName}}.
type Read{{.EntityTitle}}Request struct {
	ID int `json:"id"`
}

// Read{{.EntityTitle}}Response model for returning a {{.EntityName}}.
type Read{{.EntityTitle}}Response struct {
	{{.EntityTitle}}
}

// Update{{.EntityTitle}}Request model for updating {{.EntityName}} details.
type Update{{.EntityTitle}}Request struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
	Status      string    `json:"status"`
	AuthorID    int       `json:"author_id"`
}

// Update{{.EntityTitle}}Response model for the update response.
type Update{{.EntityTitle}}Response struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

// Delete{{.EntityTitle}}Request model for deleting a {{.EntityName}}.
type Delete{{.EntityTitle}}Request struct {
	ID int `json:"id"`
}

// Delete{{.EntityTitle}}Response model for the delete response.
type Delete{{.EntityTitle}}Response struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

// List{{.EntityTitle}}Request model for listing {{.EntityName}}s, can include filters for status.
type List{{.EntityTitle}}Request struct {
	{{.EntityTitle}}     int    `json:"{{.EntityName}}"`
	Limit    int    `json:"limit"`
	Status   string `json:"status"`
	AuthorID int    `json:"author_id"`
}

// List{{.EntityTitle}}Response model for the list response.
type List{{.EntityTitle}}Response struct {
	{{.EntityTitle}}s []{{.EntityTitle}} `json:"{{.EntityName}}s"`
}
