package common

type Paging struct {
	Page    *int    `json:"page" form:"page"`
	Limit   *int    `json:"limit" form:"limit"`
	OrderBy *string `json:"order_by" form:"order_by"`
	Asc     *bool   `json:"asc" form:"asc"`
	//Suport cursor with UID
	Cursor *string `json:"cursor" form:"cursor"`
	//NextCursor *string `json:"next_cursor"`
}

func (p *Paging) Fulfill() {
	if p.Page == nil || *p.Page <= 0 {
		page := 1
		p.Page = &page
	}

	if p.Limit == nil || *p.Limit <= 0 {
		limit := 50
		p.Limit = &limit
	}
	//
	//if p.NextCursor != nil {
	//	*p.NextCursor = strings.TrimSpace(*p.NextCursor)
	//}
}
