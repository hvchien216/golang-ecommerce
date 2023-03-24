package categorymodel

type NewCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentID    *int   `json:"parent_id"`
}
