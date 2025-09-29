package paginator_test

import (
	"testing"

	"github.com/CodeInsiderIO/go-paginator"
)

type testCase struct {
	name string

	// Input
	totalItems  int64
	currentPage int64
	limit       int64

	// Expected Output
	expectedPerPage     int64
	expectedCurrentPage int64
	expectedTotalItems  int64
	expectedTotalPages  int64
	expectedOffset      int64
	expectedItemCount   int64
	expectedHasPrevious bool
	expectedHasNext     bool
	expectedPrevPage    int64
	expectedNextPage    int64
}

func checkPaginator(t *testing.T, p *paginator.Paginator, tc testCase) {
	if p.PerPage != tc.expectedPerPage {
		t.Errorf("PerPage got %d, want %d", p.PerPage, tc.expectedPerPage)
	}
	if p.CurrentPage != tc.expectedCurrentPage {
		t.Errorf("CurrentPage got %d, want %d", p.CurrentPage, tc.expectedCurrentPage)
	}
	if p.TotalItems != tc.expectedTotalItems {
		t.Errorf("TotalItems got %d, want %d", p.TotalItems, tc.expectedTotalItems)
	}
	if p.TotalPages != tc.expectedTotalPages {
		t.Errorf("TotalPages got %d, want %d", p.TotalPages, tc.expectedTotalPages)
	}
	if p.Offset != tc.expectedOffset {
		t.Errorf("Offset got %d, want %d", p.Offset, tc.expectedOffset)
	}
	if p.ItemCount != tc.expectedItemCount {
		t.Errorf("ItemCount got %d, want %d", p.ItemCount, tc.expectedItemCount)
	}
	if p.HasPrevious != tc.expectedHasPrevious {
		t.Errorf("HasPrevious got %t, want %t", p.HasPrevious, tc.expectedHasPrevious)
	}
	if p.HasNext != tc.expectedHasNext {
		t.Errorf("HasNext got %t, want %t", p.HasNext, tc.expectedHasNext)
	}
	if p.PrevPage != tc.expectedPrevPage {
		t.Errorf("PrevPage got %d, want %d", p.PrevPage, tc.expectedPrevPage)
	}
	if p.NextPage != tc.expectedNextPage {
		t.Errorf("NextPage got %d, want %d", p.NextPage, tc.expectedNextPage)
	}
}

func TestNewPaginator(t *testing.T) {
	testCases := []testCase{
		{
			name:       "Trường hợp 1: Trang đầu tiên, tổng cộng 100 items, limit 10",
			totalItems: 100, currentPage: 1, limit: 10,
			expectedPerPage: 10, expectedCurrentPage: 1, expectedTotalItems: 100,
			expectedTotalPages: 10, expectedOffset: 0, expectedItemCount: 10,
			expectedHasPrevious: false, expectedHasNext: true, expectedPrevPage: 0, expectedNextPage: 2,
		},
		{
			name:       "Trường hợp 2: Trang giữa, tổng cộng 100 items, limit 10",
			totalItems: 100, currentPage: 5, limit: 10,
			expectedPerPage: 10, expectedCurrentPage: 5, expectedTotalItems: 100,
			expectedTotalPages: 10, expectedOffset: 40, expectedItemCount: 10,
			expectedHasPrevious: true, expectedHasNext: true, expectedPrevPage: 4, expectedNextPage: 6,
		},
		{
			name:       "Trường hợp 3: Trang cuối cùng, tổng cộng 100 items, limit 10",
			totalItems: 100, currentPage: 10, limit: 10,
			expectedPerPage: 10, expectedCurrentPage: 10, expectedTotalItems: 100,
			expectedTotalPages: 10, expectedOffset: 90, expectedItemCount: 10,
			expectedHasPrevious: true, expectedHasNext: false, expectedPrevPage: 9, expectedNextPage: 0,
		},
		{
			name:       "Trường hợp 4: Tổng số lẻ items, trang cuối",
			totalItems: 42, currentPage: 5, limit: 10,
			expectedPerPage: 10, expectedCurrentPage: 5, expectedTotalItems: 42,
			expectedTotalPages: 5, expectedOffset: 40, expectedItemCount: 2,
			expectedHasPrevious: true, expectedHasNext: false, expectedPrevPage: 4, expectedNextPage: 0,
		},
		{
			name:       "Trường hợp 5: current page lớn hơn TotalPages",
			totalItems: 42, currentPage: 10, limit: 10,
			expectedPerPage: 10, expectedCurrentPage: 5, expectedTotalItems: 42,
			expectedTotalPages: 5, expectedOffset: 40, expectedItemCount: 2,
			expectedHasPrevious: true, expectedHasNext: false, expectedPrevPage: 4, expectedNextPage: 0,
		},
		{
			name:       "Trường hợp 6: current page nhỏ hơn 1",
			totalItems: 42, currentPage: 0, limit: 10,
			expectedPerPage: 10, expectedCurrentPage: 1, expectedTotalItems: 42,
			expectedTotalPages: 5, expectedOffset: 0, expectedItemCount: 10,
			expectedHasPrevious: false, expectedHasNext: true, expectedPrevPage: 0, expectedNextPage: 2,
		},
		{
			name:       "Trường hợp 7: TotalItems bằng 0",
			totalItems: 0, currentPage: 1, limit: 10,
			expectedPerPage: 10, expectedCurrentPage: 1, expectedTotalItems: 0,
			expectedTotalPages: 1, expectedOffset: 0, expectedItemCount: 0,
			expectedHasPrevious: false, expectedHasNext: false, expectedPrevPage: 0, expectedNextPage: 0,
		},
		{
			name:       "Trường hợp 8: limit bằng 0, nên được gán mặc định là 10",
			totalItems: 50, currentPage: 1, limit: 0,
			expectedPerPage: 10, expectedCurrentPage: 1, expectedTotalItems: 50,
			expectedTotalPages: 5, expectedOffset: 0, expectedItemCount: 10,
			expectedHasPrevious: false, expectedHasNext: true, expectedPrevPage: 0, expectedNextPage: 2,
		},
		{
			name:       "Trường hợp 9: Tổng 1 item, trang 1, limit 10",
			totalItems: 1, currentPage: 1, limit: 10,
			expectedPerPage: 10, expectedCurrentPage: 1, expectedTotalItems: 1,
			expectedTotalPages: 1, expectedOffset: 0, expectedItemCount: 1,
			expectedHasPrevious: false, expectedHasNext: false, expectedPrevPage: 0, expectedNextPage: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := paginator.NewPaginator(tc.totalItems, tc.currentPage, tc.limit)
			checkPaginator(t, p, tc)
		})
	}
}

func TestPaginator_Set(t *testing.T) {
	p := paginator.NewPaginator(10, 1, 5)

	testCases := []testCase{
		{
			name:       "Trường hợp 1: Thay đổi thông số, tổng 50, trang 3, limit 5",
			totalItems: 50, currentPage: 3, limit: 5,
			expectedPerPage: 5, expectedCurrentPage: 3, expectedTotalItems: 50,
			expectedTotalPages: 10, expectedOffset: 10, expectedItemCount: 5,
			expectedHasPrevious: true, expectedHasNext: true, expectedPrevPage: 2, expectedNextPage: 4,
		},
		{
			name:       "Trường hợp 2: Thay đổi sang trường hợp TotalItems=0",
			totalItems: 0, currentPage: 1, limit: 10,
			expectedPerPage: 10, expectedCurrentPage: 1, expectedTotalItems: 0,
			expectedTotalPages: 1, expectedOffset: 0, expectedItemCount: 0,
			expectedHasPrevious: false, expectedHasNext: false, expectedPrevPage: 0, expectedNextPage: 0,
		},
		{
			name:       "Trường hợp 3: Thay đổi sang trường hợp limit=0",
			totalItems: 25, currentPage: 2, limit: 0,
			expectedPerPage: 10, expectedCurrentPage: 2, expectedTotalItems: 25,
			expectedTotalPages: 3, expectedOffset: 10, expectedItemCount: 10,
			expectedHasPrevious: true, expectedHasNext: true, expectedPrevPage: 1, expectedNextPage: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p.Set(tc.totalItems, tc.currentPage, tc.limit)
			checkPaginator(t, p, tc)
		})
	}
}

func TestPaginator_recompute_LimitDefaults(t *testing.T) {
	tc := testCase{
		name:       "NewPaginator với limit âm",
		totalItems: 100, currentPage: 1, limit: -5,
		expectedPerPage: 10, expectedCurrentPage: 1, expectedTotalItems: 100,
		expectedTotalPages: 10, expectedOffset: 0, expectedItemCount: 10,
		expectedHasPrevious: false, expectedHasNext: true, expectedPrevPage: 0, expectedNextPage: 2,
	}
	t.Run(tc.name, func(t *testing.T) {
		pNegativeLimit := paginator.NewPaginator(tc.totalItems, tc.currentPage, tc.limit)
		checkPaginator(t, pNegativeLimit, tc)
	})
}
