package categorymodel

import (
	"github.com/hvchien216/golang-ecommerce/internal/orm"
	"time"
)

const EntityName = "CATEGORY"

type Category struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Parent      *Category   `json:"parent"`
	Children    []*Category `json:"children"`
	CreatedAt   *time.Time  `json:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at"`
}

type CategoryResult struct {
	Categories []*orm.Category `json:"categories"`
	PageInfo   *PageInfo       `json:"page_info"`
}

type PageInfo struct {
	Page       int     `json:"page"`
	Limit      int     `json:"limit"`
	OrderBy    string  `json:"order_by"`
	Asc        bool    `json:"asc"`
	Total      int64   `json:"total"`
	Cursor     *string `json:"cursor"`
	NextCursor string  `json:"next_cursor"`
}
