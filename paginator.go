package utility

import "math"

type Paginator struct {
	PerPage     int64 `json:"perPage"`
	HasPrevious bool  `json:"hasPrevious"`
	HasNext     bool  `json:"hasNext"`
	Offset      int64 `json:"offset"`
	CurrentPage int64 `json:"currentPage"`
	ItemCount   int64 `json:"itemCount"`
	TotalItems  int64 `json:"totalItems"`
	TotalPages  int64 `json:"totalPages"`
}

func NewPaginator(totalItems int64, currentPage int64, limit int64) *Paginator {
	paginator := &Paginator{
		TotalItems:  totalItems,
		CurrentPage: currentPage,
		PerPage:     limit,
	}
	paginator.Paginate()
	return paginator
}

func (p *Paginator) Paginate() {
	p.TotalPages = p.calculateTotalPages()
	p.CurrentPage = p.getCurrentPage()
	p.HasNext = p.hasNext()
	p.HasPrevious = p.hasPrevious()
	p.ItemCount = p.getItemCount()
	p.Offset = p.findOffset()
}

func (p *Paginator) getItemCount() int64 {
	return p.TotalItems - (p.TotalPages-1)*p.PerPage
}

func (p *Paginator) hasPrevious() bool {
	if p.CurrentPage >= 2 {
		return true
	}
	return false
}

func (p *Paginator) hasNext() bool {
	return p.CurrentPage < p.TotalPages
}

func (p *Paginator) calculateTotalPages() int64 {
	return int64(math.Ceil(float64(p.TotalItems) / float64(p.PerPage)))
}

func (p *Paginator) findOffset() int64 {
	return (p.getCurrentPage() - 1) * p.PerPage
}

func (p *Paginator) getCurrentPage() int64 {
	if p.CurrentPage <= 1 {
		return 1
	}
	return p.CurrentPage
}

func (p *Paginator) getNextPage() int64 {
	return p.getCurrentPage()
}

func (p *Paginator) findLastPage() int64 {
	return p.calculateTotalPages()
}
