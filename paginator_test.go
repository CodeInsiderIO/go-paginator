package utility

import (
	"testing"
)

func TestNewPaginator(t *testing.T) {
	totalItems := int64(100)
	currentPage := int64(1)
	limit := int64(10)

	paginator := NewPaginator(totalItems, currentPage, limit)

	if paginator.TotalItems != totalItems {
		t.Errorf("Expected TotalItems to be %d, but got %d", totalItems, paginator.TotalItems)
	}

	if paginator.CurrentPage != currentPage {
		t.Errorf("Expected CurrentPage to be %d, but got %d", currentPage, paginator.CurrentPage)
	}

	if paginator.PerPage != limit {
		t.Errorf("Expected PerPage to be %d, but got %d", limit, paginator.PerPage)
	}

	// Test pagination
	paginator.Paginate()

	expectedTotalPages := int64(10)
	if paginator.TotalPages != expectedTotalPages {
		t.Errorf("Expected TotalPages to be %d, but got %d", expectedTotalPages, paginator.TotalPages)
	}

	expectedHasPrevious := false
	if paginator.HasPrevious != expectedHasPrevious {
		t.Errorf("Expected HasPrevious to be %v, but got %v", expectedHasPrevious, paginator.HasPrevious)
	}

	expectedHasNext := true
	if paginator.HasNext != expectedHasNext {
		t.Errorf("Expected HasNext to be %v, but got %v", expectedHasNext, paginator.HasNext)
	}

	expectedItemCount := int64(10)
	if paginator.ItemCount != expectedItemCount {
		t.Errorf("Expected ItemCount to be %d, but got %d", expectedItemCount, paginator.ItemCount)
	}

	expectedOffset := int64(0)
	if paginator.Offset != expectedOffset {
		t.Errorf("Expected Offset to be %d, but got %d", expectedOffset, paginator.Offset)
	}
}

func TestPaginator_calculateTotalPages(t *testing.T) {
	totalItems := int64(20)
	perPage := int64(5)

	paginator := Paginator{
		TotalItems: totalItems,
		PerPage:    perPage,
	}

	expectedTotalPages := int64(4)
	actualTotalPages := paginator.calculateTotalPages()

	if actualTotalPages != expectedTotalPages {
		t.Errorf("Expected TotalPages to be %d, but got %d", expectedTotalPages, actualTotalPages)
	}
}

func TestPaginator_findOffset(t *testing.T) {
	currentPage := int64(3)
	perPage := int64(10)

	paginator := Paginator{
		CurrentPage: currentPage,
		PerPage:     perPage,
	}

	expectedOffset := int64(20)
	actualOffset := paginator.findOffset()

	if actualOffset != expectedOffset {
		t.Errorf("Expected Offset to be %d, but got %d", expectedOffset, actualOffset)
	}
}

func TestPaginator_getCurrentPage(t *testing.T) {
	currentPage := int64(0)

	paginator := Paginator{
		CurrentPage: currentPage,
	}

	expectedCurrentPage := int64(1)
	actualCurrentPage := paginator.getCurrentPage()

	if actualCurrentPage != expectedCurrentPage {
		t.Errorf("Expected CurrentPage to be %d, but got %d", expectedCurrentPage, actualCurrentPage)
	}
}

func TestPaginator_getItemCount(t *testing.T) {
	totalItems := int64(100)
	totalPages := int64(10)
	perPage := int64(10)

	paginator := Paginator{
		TotalItems: totalItems,
		TotalPages: totalPages,
		PerPage:    perPage,
	}

	expectedItemCount := int64(10)
	actualItemCount := paginator.getItemCount()

	if actualItemCount != expectedItemCount {
		t.Errorf("Expected ItemCount to be %d, but got %d", expectedItemCount, actualItemCount)
	}
}

func TestPaginator_hasPrevious(t *testing.T) {
	currentPage := int64(1)

	paginator := Paginator{
		CurrentPage: currentPage,
	}

	actualHasPrevious := paginator.hasPrevious()

	if actualHasPrevious != false {
		t.Errorf("Expected HasPrevious to be %v, but got %v", false, actualHasPrevious)
	}
}

func TestPaginator_hasNext(t *testing.T) {
	currentPage := int64(1)
	totalPages := int64(10)

	paginator := Paginator{
		CurrentPage: currentPage,
		TotalPages:  totalPages,
	}

	actualHasNext := paginator.hasNext()

	if actualHasNext != true {
		t.Errorf("Expected HasNext to be %v, but got %v", true, actualHasNext)
	}
}
