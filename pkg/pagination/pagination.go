package pagination

import "math"

type (
	PaginationRequest struct {
		Page     int    `form:"page" binding:"omitempty,min=1"`
		PageSize int    `form:"pageSize" binding:"omitempty,min=1"`
		Sort     string `form:"sort" binding:"omitempty,oneof=asc desc"`
		SortBy   string `form:"sortBy" binding:"omitempty"`
	}

	PaginationResponse struct {
		TotalData   int64 `json:"totalData"`
		TotalPages  int   `json:"totalPage"`
		CurrentPage int   `json:"currentPage"`
		CurrentSize int   `json:"currentSize"`
	}
)

func (p *PaginationRequest) GetPage() int {
	if p.Page <= 0 {
		return 1
	}
	return p.Page
}

func (p *PaginationRequest) GetLimit() int {
	if p.PageSize <= 0 {
		return 10
	}
	return p.PageSize
}

func (p *PaginationRequest) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *PaginationRequest) GetSortBy() string {
	return p.SortBy
}

func (p *PaginationRequest) GetSortOrder() string {
	return p.Sort
}

func (p *PaginationRequest) CreatePaginationResponse(totalData int64, currentSize int) PaginationResponse {
	totalPages := int(math.Ceil(float64(totalData) / float64(p.GetLimit())))
	return PaginationResponse{
		TotalData:   totalData,
		TotalPages:  totalPages,
		CurrentPage: p.GetPage(),
		CurrentSize: currentSize,
	}
}
