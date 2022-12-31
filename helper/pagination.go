package helper

type Pagination struct {
	Page       int `json:"page,omitempty"`
	Limit      int `json:"limit,omitempty"`
	TotalRows  int `json:"total_rows"`
	TotalPages int `json:"total_pages"`
}

func (p *Pagination) GetPage() int {
	if p.Page < 1 {
		p.Page = 1
	}

	return p.Page
}

func (p *Pagination) GetLimit() int {
	if p.Limit < 1 {
		p.Limit = 10
	}

	return p.Limit
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}
