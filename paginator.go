package utility

type Paginator struct {
	PerPage     int64 `json:"perPage"`
	CurrentPage int64 `json:"currentPage"`
	TotalItems  int64 `json:"totalItems"`

	// Derived
	TotalPages  int64 `json:"totalPages"`
	Offset      int64 `json:"offset"`
	ItemCount   int64 `json:"itemCount"` // số item thực tế ở trang hiện tại
	HasPrevious bool  `json:"hasPrevious"`
	HasNext     bool  `json:"hasNext"`
	PrevPage    int64 `json:"prevPage"`
	NextPage    int64 `json:"nextPage"`
}

func NewPaginator(totalItems, currentPage, limit int64) *Paginator {
	p := &Paginator{
		TotalItems:  totalItems,
		PerPage:     limit,
		CurrentPage: currentPage,
	}
	p.recompute()
	return p
}

func (p *Paginator) recompute() {
	if p.PerPage <= 0 {
		p.PerPage = 10
	}

	if p.TotalItems <= 0 {
		p.TotalItems = 0
		p.TotalPages = 1
		p.CurrentPage = 1
		p.Offset = 0
		p.ItemCount = 0
		p.HasPrevious = false
		p.HasNext = false
		p.PrevPage = 0
		p.NextPage = 0
		return
	}

	p.TotalPages = (p.TotalItems + p.PerPage - 1) / p.PerPage

	if p.CurrentPage < 1 {
		p.CurrentPage = 1
	} else if p.CurrentPage > p.TotalPages {
		p.CurrentPage = p.TotalPages
	}

	p.Offset = (p.CurrentPage - 1) * p.PerPage

	if p.CurrentPage < p.TotalPages {
		p.ItemCount = p.PerPage
	} else {
		p.ItemCount = p.TotalItems - (p.PerPage * (p.TotalPages - 1))
	}

	p.HasPrevious = p.CurrentPage > 1
	p.HasNext = p.CurrentPage < p.TotalPages
	if p.HasPrevious {
		p.PrevPage = p.CurrentPage - 1
	} else {
		p.PrevPage = 0
	}
	if p.HasNext {
		p.NextPage = p.CurrentPage + 1
	} else {
		p.NextPage = 0
	}
}

func (p *Paginator) Set(totalItems, currentPage, limit int64) {
	p.TotalItems = totalItems
	p.CurrentPage = currentPage
	p.PerPage = limit
	p.recompute()
}
